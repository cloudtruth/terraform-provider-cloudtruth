package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

/* Notes:
   At the Parameter level, only the secret property and description are mutable. We disallow changing
   the parameter name, project and environment and instead instruct the user to destroy and recreate
   the Parameter.

	At the Parameter Value level, we allow toggling internal/external + dynamic/non-dynamic as well as
    updating the value itself. However, we must block certain invalid combinations such as:
    * changing a value from internal -> external when the Parameter is a secret Parameter
    * changing a value from non-dynamic -> dynamic when the Value is external
    There are likely many corner cases to consider here.
*/

func resourceParameter() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A CloudTruth Parameter and environment specific value (defaulting to the 'default' environment)",

		CreateContext: resourceParameterCreate,
		ReadContext:   resourceParameterRead,
		UpdateContext: resourceParameterUpdate,
		DeleteContext: resourceParameterDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Parameter, unique per project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Parameter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "The CloudTruth project where the Parameter is defined",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment where the Parameter's value is defined. Defaults to the 'default' environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"secret": {
				Description: "Whether or not the Parameter is a secret, defaults to false (non-secret)",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"value": {
				Description: "The value of the CloudTruth Parameter, specific to an environment (can be overridden/inherited relative to other environments)",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				// With external parameters, the value is computed but should not be used to diff/detect changes
				// It is only useful in read operations. A change to the location and/or filter field would indicate a resource change
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					external := d.Get("external").(bool)
					return external
				},
			},
			"dynamic": {
				Description: "Whether or not to evaluate/interpolate the Parameter's value (incompatible with secret parameters)",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"external": {
				Description: "Whether or not the value is external, defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"location": {
				Description: "The location of the secret value, required for external parameters",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"filter": {
				Description: "An optional filter (path/query), optional and used only with external parameters",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
		},
	}
}

// In Terraform terms, a 'resource_parameter' represents a CloudTruth Parameter AND its environment specific value
// Therefore, when this provider destroys a parameter resource, it only removes the per-environment value unless
// the parameter is only defined in one environment, in which case it is destroyed entirely
func resourceParameterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceParameterCreate")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterCreate: %w", err))
	}

	paramName := d.Get("name").(string)
	var paramID, valueID string
	lookupResp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(ctx,
		*projID).Environment(*envID).Name(paramName).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterCreate: error looking up parameter %s: %+v", paramName, r))
	}

	if lookupResp.GetCount() == 1 { // Named parameter already exists
		paramID = lookupResp.GetResults()[0].GetId()
	} else {
		paramCreate := paramCreateConfig(d)
		paramCreateResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersCreate(ctx,
			*projID).ParameterCreate(*paramCreate).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceParameterCreate: %w", err))
		}
		paramID = paramCreateResp.GetId()
	}
	valueCreate, err := valueCreateConfig(*envID, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterCreate: %w", err))
	}
	valueResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesCreate(ctx,
		paramID, *projID).ValueCreate(*valueCreate).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterCreate: %w", err))
	}
	valueID = valueResp.GetId()

	// Then add its value in the specified environment with composite ID - <PARAMETER_ID>:<PARAMETER_VALUE_ID>
	internalID := fmt.Sprintf("%s:%s", paramID, valueID)
	d.SetId(internalID)
	return nil
}

func paramCreateConfig(d *schema.ResourceData) *cloudtruthapi.ParameterCreate {
	paramName := d.Get("name").(string)
	paramCreate := cloudtruthapi.NewParameterCreate(paramName)
	paramDesc := d.Get("description").(string)
	if paramDesc != "" {
		paramCreate.SetDescription(paramDesc)
	}
	isSecret := d.Get("secret").(bool)
	paramCreate.SetSecret(isSecret)
	return paramCreate
}

func valueCreateConfig(envID string, d *schema.ResourceData) (*cloudtruthapi.ValueCreate, error) {
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	dynamic := d.Get("dynamic").(bool)
	valueCreate := cloudtruthapi.NewValueCreate(value)
	external := d.Get("external").(bool)
	if external && value != "" {
		return nil, fmt.Errorf("you cannot specify a value (%s)with the external parameter %s", value, name)
	}
	if value != "" {
		valueCreate.SetInternalValue(value)
		valueCreate.SetInterpolated(dynamic)
	} else {
		valueCreate.SetExternal(external)
		location, filter := d.Get("location").(string), d.Get("filter").(string)
		if location == "" {
			return nil, fmt.Errorf("you must specify the location of the external parameter %s", name)
		}
		valueCreate.SetExternalFqn(location)
		if filter != "" {
			valueCreate.SetExternalFilter(filter)
		}
	}
	valueCreate.SetEnvironment(envID)
	return valueCreate, nil
}

func resourceParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterRead")
	return dataCloudTruthParameterRead(ctx, d, meta)
}

func resourceParameterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceParameterUpdate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
	}

	paramCompositeID := d.Id()
	ids := strings.Split(paramCompositeID, ":")
	if len(ids) != 2 {
		return diag.FromErr(fmt.Errorf("resourceParameterUpdate: ;failed to extract the Parameter and Parameter Value IDs from %s",
			paramCompositeID))
	}
	paramID, paramValueID := ids[0], ids[1]

	// First update Parameter level changes
	patchedParam := cloudtruthapi.PatchedParameter{}
	hasParamChange := false
	if d.HasChange("description") {
		paramDesc := d.Get("description").(string)
		patchedParam.SetDescription(paramDesc)
		hasParamChange = true
	}
	if d.HasChange("secret") {
		paramIsSecret := d.Get("secret").(bool)
		patchedParam.SetSecret(paramIsSecret)
		hasParamChange = true
	}
	if hasParamChange {
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersPartialUpdate(ctx, paramID, *projID).
			PatchedParameter(patchedParam).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
		}
	}

	// Then update Parameter Value level changes
	updateValue := cloudtruthapi.NewValueWithDefaults()
	hasParamValueChange := false
	if d.HasChange("value") {
		value := d.Get("value").(string)
		updateValue.SetInternalValue(value)
		paramEnv := d.Get("environment").(string)
		envID, err := c.lookupEnvironment(ctx, paramEnv)
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
		}
		updateValue.SetEnvironment(*envID)
		hasParamValueChange = true
	}
	if d.HasChange("external") {
		external := d.Get("external").(bool)
		updateValue.SetExternal(external)
		hasParamValueChange = true
	}
	if d.HasChange("dynamic") {
		evalValue := d.Get("dynamic").(bool)
		updateValue.SetInterpolated(evalValue)
		hasParamValueChange = true
	}
	if hasParamValueChange {
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesUpdate(ctx, paramValueID, paramID,
			*projID).Value(*updateValue).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
		}
	}
	d.SetId(paramCompositeID)
	return resourceParameterRead(ctx, d, meta)
}

// If a parameter is defined in the target environment only, we destroy it
// If it is defined in one or more other environments, we only destroy the value
// defined in the target environment
func resourceParameterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterDelete")
	c := meta.(*cloudTruthClient)
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterDelete: %w", err))
	}

	paramCompositeID := d.Id()
	ids := strings.Split(paramCompositeID, ":")
	if len(ids) != 2 {
		return diag.FromErr(fmt.Errorf("resourceParameterDelete: failed to extract the Parameter and Parameter Value IDs from %s",
			paramCompositeID))
	}
	paramID, paramValueID := ids[0], ids[1]
	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersRetrieve(ctx, paramID, *projID).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterDelete: %w", err))
	}

	// We iterate over the values for the parameter, where the keys are the URLs of the corresponding
	// environments. If they are all null, we delete the parameter. If they are all null except the one matching
	// the specified environment, we delete the parameter.
	// Otherwise, we delete the specific value defined in the target environment
	multiEnv := false
	values := resp.GetValues()
	count := 1 // the param value is defined in the target environment
	// check to see if it's explicitly defined in any others
	for envURL := range values {
		value := values[envURL]
		if (value.GetEnvironmentName() != environment) && value.HasInternalValue() {
			count++
			if count > 1 {
				multiEnv = true
				break
			}
		}
	}

	if multiEnv {
		// delete the specific env value only, because there are explicit definitions in one or more other envs
		_, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesDestroy(ctx, paramValueID,
			paramID, *projID).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceParameterDelete: %w", err))
		}
	} else {
		// delete the parameter entirely
		_, err := c.openAPIClient.ProjectsApi.ProjectsParametersDestroy(ctx, paramID,
			*projID).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceParameterDelete: %w", err))
		}
	}
	return nil
}
