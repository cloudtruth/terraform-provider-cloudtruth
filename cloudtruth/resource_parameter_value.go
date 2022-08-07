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
)

func resourceParameterValue() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Parameter and environment specific value (defaulting to the 'default' environment)",

		CreateContext: resourceParameterValueCreate,
		ReadContext:   resourceParameterValueRead,
		UpdateContext: resourceParameterValueUpdate,
		DeleteContext: resourceParameterValueDelete,

		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The CloudTruth project where the Parameter Value is defined",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment where the Parameter Value will be added. Defaults to the 'default' environment",
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
			"parameter_id": {
				Description: "The ID of the CloudTruth Parameter which stores this value",
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
			outErr := fmt.Errorf("resourceParameterValueCreate: error looking up parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	if paramListResp.GetCount() != 1 {
		return diag.FromErr(fmt.Errorf("resourceParameterRead: expected 1 value for parameter %s, found %d instead",
			paramName, paramListResp.GetCount()))
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
			outErr := fmt.Errorf("resourceParameterValueCreate: error creating the value for parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
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

func resourceParameterValueRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterRead")
	return dataCloudTruthParameterValueRead(ctx, d, meta)
}

// No retry logic here, the caller handles that
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
	tflog.Debug(ctx, "resourceParameterUpdate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
	}

	paramID := d.Get("parameter_id").(string)
	paramName := d.Get("parameter_name").(string)

	// Then update Parameter Value level changes
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := updateParameterValue(ctx, paramID, d.Id(), *projID, d, c)
		if err != nil {
			outErr := fmt.Errorf("resourceParameterUpdate: error updating the parameter value config for parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
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
	tflog.Debug(ctx, "resourceParameterDelete")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterDelete: %w", err))
	}
	paramID := d.Get("parameter_id").(string)
	paramValueID := d.Id()
	paramName := d.Get("parameter_name").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsApi.ProjectsParametersValuesDestroy(ctx, paramValueID,
			paramID, *projID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceParameterDelete: error deleting parameter value %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
