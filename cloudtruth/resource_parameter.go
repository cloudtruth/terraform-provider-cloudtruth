package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"strings"
)

// todo:
// force_delete support
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
			},
			"secret": {
				Description: "Whether or not the Parameter is a secret, defaults to false/non-secret",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"type": { // todo: handle this
				Description: "Whether or not the Parameter is a secret, defaults to false/non-secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"evaluate": { // todo: handle this
				Description: "Whether or not the Parameter is a secret, defaults to false/non-secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"wrap": { // todo: handle this
				Description: "Whether or not the Parameter is a secret, defaults to false/non-secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"force_delete": { // todo: handle this
				Description: "Whether to allow Terraform to delete the CloudTruth Parameter or not",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

// We treat the parameter and its per environment value as one provider object
func resourceParameterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterCreate")
	var diags diag.Diagnostics
	// First create the parameter
	c := meta.(*cloudTruthClient)
	paramName := d.Get("name").(string)
	paramDesc := d.Get("description").(string)
	paramCreate := cloudtruthapi.NewParameterCreate(paramName)
	if paramDesc != "" {
		paramCreate.SetDescription(paramDesc)
	}

	// todo: refactor this to be a shared function
	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	// NOTE: a parameter will exist in any/all environments however, its value may be set/overridden/inherited across
	// multiple environments
	paramResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersCreate(context.Background(),
		*projID).ParameterCreate(*paramCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	paramID := paramResp.GetId()

	// Then add its value in the specified environment
	// guaranteed to be set to "default" if not explicitly specified
	paramEnv := d.Get("environment").(string)
	paramEnvID, err := c.lookupEnvironment(ctx, paramEnv)
	if err != nil {
		return diag.FromErr(err)
	}

	// todo: handle multiple param/value/env instances, possibly with a map
	// if unset, we default to an empty string/nil
	value := d.Get("value").(string)
	valueCreate := cloudtruthapi.NewValueCreate(value)
	valueCreate.SetInternalValue(value)
	valueCreate.SetEnvironment(*paramEnvID)
	// todo: someday support external values
	valueCreate.SetExternal(false)
	//todo: evaluate
	//todo: wrap
	valueResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesCreate(context.Background(),
		paramID, *projID).ValueCreate(*valueCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	// We use a composite ID - <PARAMATER_ID>:<PARAMETER_VALUE_ID>
	d.SetId(fmt.Sprintf("%s:%s", paramID, valueResp.GetId()))
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
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
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
	if hasParamChange {
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersPartialUpdate(ctx, paramID, *projID).
			PatchedParameter(patchedParam).Execute()
		if err != nil {
			return diag.FromErr(err)
		}
	}

	patchedParamValue := cloudtruthapi.NewPatchedValue()
	hasParamValueChange := false
	if d.HasChange("value") {
		paramValue := d.Get("value").(string)
		patchedParamValue.SetValue(paramValue)
		hasParamValueChange = true
	}
	if hasParamValueChange {
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesPartialUpdate(ctx, paramValueID, paramID,
			*projID).PatchedValue(*patchedParamValue).Execute()
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
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersRetrieve(context.Background(), paramID, *projID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	for _, v := range resp.GetValues() {
		v.
	}


	return nil
}
