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

func resourceParameter() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A Cloudtruth Parameter.",

		CreateContext: resourceParameterCreate,
		ReadContext:   resourceParameterRead,
		UpdateContext: resourceParameterUpdate,
		DeleteContext: resourceParameterDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Parameter",
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
				Description: "The CloudTruth environment where the Parameter's value is defined. Defaults to empty string",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"value": {
				Description: "The value of the CloudTruth Parameter, specific to an Environment (which can be overridden/inherited)",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"secret": {
				Description: "Whether or not the Parameter is a secret, defaults to false/non-secret",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"evaluate": {
				Description: "Whether to run template evaluation on the Parameter's value aka 'dynamic', incompatible with secret parameters",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"external": {
				Description: "Whether the value is a reference to a value in an external system or defined in CloudTruth, defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

// In Terraform terms, a 'resource_parameter' represents a CloudTruth Parameter AND its environment specific value
// Therefore, when this provider destroys a parameter resource, it only removes the per-environment value unless
// the parameter is only defined in one environment, in which case it is destroyed entirely
func resourceParameterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterCreate")
	var diags diag.Diagnostics
	// First create the parameter
	c := meta.(*cloudTruthClient)
	paramName := d.Get("name").(string)
	paramCreate := cloudtruthapi.NewParameterCreate(paramName)
	paramDesc := d.Get("description").(string)
	if paramDesc != "" {
		paramCreate.SetDescription(paramDesc)
	}
	paramIsSecret := d.Get("secret").(bool)
	paramCreate.SetSecret(paramIsSecret)

	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	paramEnv := d.Get("environment").(string)
	paramEnvID, err := c.lookupEnvironment(ctx, paramEnv)
	if err != nil {
		return diag.FromErr(err)
	}

	// NOTE: a parameter will exist in any/all environments however, its value may be set/overridden/inherited across
	// multiple environments
	var paramID, valueID string
	// First check to see if it exists, if it does grab the parameter's ID
	lookupResp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
		*projID).Environment(*paramEnvID).Name(paramName).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("error looking up parameter %s: %+v", paramName, r))
	}
	value := d.Get("value").(string)
	if lookupResp.GetCount() == 1 {
		paramID = lookupResp.GetResults()[0].GetId()

	} else {
		// if good, we use this
		// else we set
		paramCreateResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersCreate(context.Background(),
			*projID).ParameterCreate(*paramCreate).Execute()
		if err != nil {
			return diag.FromErr(err)
		}
		paramID = paramCreateResp.GetId()
	}
	external := d.Get("external").(bool)
	evaluate := d.Get("evaluate").(bool)
	valueCreate := cloudtruthapi.NewValueCreate(value)
	valueCreate.SetInternalValue(value)
	valueCreate.SetEnvironment(*paramEnvID)
	valueCreate.SetExternal(external)
	valueCreate.SetInterpolated(evaluate)
	valueResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesCreate(context.Background(),
		paramID, *projID).ValueCreate(*valueCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	valueID = valueResp.GetId()

	// Then add its value in the specified environment
	// Composite ID - <PARAMETER_ID>:<PARAMETER_VALUE_ID>
	internalID := fmt.Sprintf("%s:%s", paramID, valueID)
	d.SetId(internalID)
	return diags
}

func resourceParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterRead")
	return dataCloudTruthParameterRead(ctx, d, meta)
}

func resourceParameterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterUpdate")
	c := meta.(*cloudTruthClient)
	paramCompositeID := d.Id()
	ids := strings.Split(paramCompositeID, ":")
	if len(ids) != 2 {
		return diag.FromErr(fmt.Errorf("failed to extract the Parameter and Parameter Value IDs from %s",
			paramCompositeID))
	}
	paramID, paramValueID := ids[0], ids[1]
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	// Two concerns:
	// 1. Update Parameter level changes
	// 2. Update Paramater Value level changes
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
			return diag.FromErr(err)
		}
	}

	updateValue := cloudtruthapi.NewValueWithDefaults()
	hasParamValueChange := false
	if d.HasChange("value") {
		value := d.Get("value").(string)
		updateValue.SetInternalValue(value)
		paramEnv := d.Get("environment").(string)
		paramEnvID, err := c.lookupEnvironment(ctx, paramEnv)
		if err != nil {
			return diag.FromErr(err)
		}
		updateValue.SetEnvironment(*paramEnvID)
		hasParamValueChange = true
	}
	if d.HasChange("external") {
		external := d.Get("external").(bool)
		updateValue.SetExternal(external)
		hasParamValueChange = true
	}
	if d.HasChange("evaluate") {
		evalValue := d.Get("evaluate").(bool)
		updateValue.SetInterpolated(evalValue)
		hasParamValueChange = true
	}
	if hasParamValueChange {
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesUpdate(ctx, paramValueID, paramID,
			*projID).Value(*updateValue).Execute()
		if err != nil {
			return diag.FromErr(err)
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
	paramCompositeID := d.Id()
	ids := strings.Split(paramCompositeID, ":")
	if len(ids) != 2 {
		return diag.FromErr(fmt.Errorf("failed to extract the Parameter and Parameter Value IDs from %s",
			paramCompositeID))
	}
	paramID, paramValueID := ids[0], ids[1]
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersRetrieve(context.Background(), paramID, *projID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	// We iterate over the values for the parameter, where the keys are the URLs of the corresponding
	// environments. If they are all null, we delete the parameter. If they are all null except the one matching
	// the specified environment, we delete the parameter.
	// Otherwise, we delete the specific value defined in the target environment
	paramEnv := d.Get("environment").(string)
	paramEnvID, err := c.lookupEnvironment(ctx, paramEnv)
	if err != nil {
		return diag.FromErr(err)
	}
	targetEnvURL := fmt.Sprintf("%s/environments/%s/", c.config.BaseURL, *paramEnvID)
	definedOutsideTargetEnv := false
	values := resp.GetValues()
	for envURL := range values {
		value := values[envURL]
		if (envURL != targetEnvURL) && value.HasInternalValue() {
			if value.GetEnvironmentName() != paramEnv {
				definedOutsideTargetEnv = true
				break
			}
		}
	}
	if definedOutsideTargetEnv {
		// delete the specific env value
		_, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesDestroy(context.Background(), paramValueID,
			paramID, *projID).Execute()
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// delete the parameter entirely
		_, err := c.openAPIClient.ProjectsApi.ProjectsParametersDestroy(context.Background(), paramID,
			*projID).Execute()
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}
