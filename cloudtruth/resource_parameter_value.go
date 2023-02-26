package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strings"
	"time"
)

const importRetries = 5

func resourceParameterValue() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Parameter and environment specific value (defaulting to the 'default' environment)",

		CreateContext: resourceParameterValueCreate,
		ReadContext:   resourceParameterValueRead,
		UpdateContext: resourceParameterValueUpdate,
		DeleteContext: resourceParameterValueDelete,

		Importer: &schema.ResourceImporter{
			StateContext: paramValueImportHelper,
		},

		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The name of the CloudTruth project where the Parameter Value is defined",
				Type:        schema.TypeString,
				Optional:    true, // But must be set via this property or the CLOUDTRUTH_PROJECT env variable
			},
			"environment": {
				Description: "The name of the CloudTruth environment where the Parameter Value will be added. Defaults to the 'default' environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"parameter_name": {
				Description: "The name of the CloudTruth Parameter which will store this value, unique per project",
				Type:        schema.TypeString,
				Required:    true,
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
				Description: "The location (path/uri) of the secret value, required for external parameters",
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
			"parameter_id": {
				Description: "The ID of the CloudTruth Parameter which holds this value",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceParameterValueCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceParameterValueCreate")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterValueCreate: %w", err))
	}

	// Check for the parameter
	paramName := d.Get("parameter_name").(string)
	var paramListResp *cloudtruthapi.PaginatedParameterList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		paramListResp, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersList(ctx, *projID).Environment(*envID).Name(paramName).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterValueCreate: error looking up parameter %s", paramName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	if paramListResp.GetCount() != 1 {
		return diag.FromErr(fmt.Errorf("resourceParameterValueCreate: failed to create the parameter value, could not find the parameter %s in project %s",
			paramName, project))
	}

	var valueID string
	valueCreate, err := valueCreateConfig(*envID, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterValueCreate: %w", err))
	}
	var value *cloudtruthapi.Value
	retryError = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		value, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersValuesCreate(ctx, paramListResp.GetResults()[0].GetId(),
			*projID).ValueCreate(*valueCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterValueCreate: error creating the value for parameter %s in the %s environment",
				paramName, environment), r, err)
		}
		valueID = value.GetId()
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(valueID)
	return nil
}

func valueCreateConfig(envID string, d *schema.ResourceData) (*cloudtruthapi.ValueCreate, error) {
	name := d.Get("parameter_name").(string)
	value := d.Get("value").(string)
	dynamic := d.Get("dynamic").(bool)
	valueCreate := cloudtruthapi.NewValueCreate(value)
	external := d.Get("external").(bool)
	if external && value != "" {
		return nil, fmt.Errorf("you cannot specify a value (%s) with the external parameter %s", value, name)
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

func parseParamValueIDInfo(ctx context.Context, c *cloudTruthClient, paramValueIDInfo string) (*string, *string, *string, *string, error) {
	tflog.Debug(ctx, "entering parseParamValueIDInfo")
	defer tflog.Debug(ctx, "exiting parseParamValueIDInfo")

	projParamParamValueIDs := strings.Split(paramValueIDInfo, ".")
	if len(projParamParamValueIDs) < 3 || len(projParamParamValueIDs) > 4 {
		return nil, nil, nil, nil,
			fmt.Errorf(`invalid import ID format: %s, you must use the format 'PROJECT_NAME.PARAMETER_ID.PARAMETER_VALUE_ID' or 
'PROJECT_NAME.ENVIRONMENT_NAME.PARAMETER_ID.PARAMETER_VALUE_ID', when you omit the environment name, import will use the 'default' environment`, paramValueIDInfo)
	}
	var projName, envName, paramID, paramValueID string
	if len(projParamParamValueIDs) == 3 {
		projName, paramID, paramValueID = projParamParamValueIDs[0], projParamParamValueIDs[1], projParamParamValueIDs[2]
	} else {
		projName, envName, paramID, paramValueID = projParamParamValueIDs[0], projParamParamValueIDs[1], projParamParamValueIDs[2], projParamParamValueIDs[3]
	}
	return &projName, &envName, &paramID, &paramValueID, nil
}

func paramValueImportHelper(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "entering paramValueImportHelper")
	defer tflog.Debug(ctx, "exiting paramValueImportHelper")
	c := meta.(*cloudTruthClient)

	projName, envName, paramID, paramValueID, err := parseParamValueIDInfo(ctx, c, d.Id())
	if err != nil {
		return nil, err
	}

	// We have the project ID, look up its name
	projID, err := c.lookupProject(ctx, *projName)
	if err != nil {
		return nil, err
	}

	var paramName string
	retryCount := 0
	var apiError error
	for retryCount < importRetries {
		var param *cloudtruthapi.Parameter
		var r *http.Response
		param, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersRetrieve(ctx, *paramID, *projID).Execute()
		if (r == nil) || (r.StatusCode >= 400 && r.StatusCode < 500) {
			apiError = err
			break
		} else if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("paramValueImportHelper: %s", err))
			retryCount++
		} else {
			paramName = param.GetName()
			break
		}
		time.Sleep(retrySleep * time.Second)
	}
	if apiError != nil {
		return nil, fmt.Errorf("paramValueImportHelper: %w", apiError)
	}
	d.SetId(*paramValueID)
	err = d.Set("project", *projName)
	if err != nil {
		return nil, err
	}
	err = d.Set("environment", *envName)
	if err != nil {
		return nil, err
	}
	err = d.Set("parameter_name", paramName)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceParameterValueRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterRead")
	return dataCloudTruthParameterValueRead(ctx, d, meta)
}

func updateParameterValue(ctx context.Context, paramID, paramValueID, projID string, d *schema.ResourceData, c *cloudTruthClient) (*http.Response, error) {
	updateValue := cloudtruthapi.NewValueWithDefaults()
	hasParamValueChange := false
	if d.HasChange("value") {
		value := d.Get("value").(string)
		updateValue.SetInternalValue(value)
		paramEnv := d.Get("environment").(string)
		envID, err := c.lookupEnvironment(ctx, paramEnv)
		if err != nil {
			return nil, err
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
	var r *http.Response
	var err error
	if hasParamValueChange {
		// No retry logic here, the caller handles that
		_, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersValuesUpdate(ctx, paramValueID, paramID,
			projID).Value(*updateValue).Execute()
		if err != nil {
			return r, err
		}
	}
	return nil, nil
}

func resourceParameterValueUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceParameterValueUpdate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterValueUpdate: %w", err))
	}
	environment := d.Get("environment").(string)

	paramID := d.Get("parameter_id").(string)
	paramName := d.Get("parameter_name").(string)

	// Then update Parameter Value level changes
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := updateParameterValue(ctx, paramID, d.Id(), *projID, d, c)
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterValueUpdate: error updating the parameter value config for parameter %s in the %s environment",
				paramName, environment), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(d.Id())
	return resourceParameterValueRead(ctx, d, meta)
}

func resourceParameterValueDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterValueDelete")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterValueDelete: %w", err))
	}
	paramID := d.Get("parameter_id").(string)
	paramValueID := d.Id()
	paramName := d.Get("parameter_name").(string)
	environment := d.Get("environment").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsApi.ProjectsParametersValuesDestroy(ctx, paramValueID,
			paramID, *projID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterValueDelete: error deleting the value from parameter %s in the %s environment",
				paramName, environment), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
