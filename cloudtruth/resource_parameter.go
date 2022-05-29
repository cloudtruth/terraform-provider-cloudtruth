package cloudtruth

import (
	"context"
	"errors"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Description: "Description of the environment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"force_delete": {
				Description: "Whether to allow Terraform to delete the environment or not",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

// We combine the parameter and its value as one object as far as the provider is concerned
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

	// todo: default environment lookup needed here likely
	//environment := d.Get("environment").(string)
	paramResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersCreate(context.Background(), *projID).ParameterCreate(*paramCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	paramID := paramResp.GetId()

	// Then add the value
	// todo: handle failure here? if we create the param but fail to add the value, subsequent param creates will fail
	value := d.Get("value").(string)
	valueCreate := cloudtruthapi.NewValueCreate(value) // ValueCreate |
	//todo: evaluate
	//todo: wrap
	valueResp, _, err := c.openAPIClient.ProjectsApi.ProjectsParametersValuesCreate(context.Background(), paramID, *projID).ValueCreate(*valueCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(valueResp.GetId())
	return diags
}

func resourceParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterRead")
	return dataCloudTruthParametersRead(ctx, d, meta)
}

func resourceParameterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterUpdate")
	// Update parameter level changes if they occurred

	// Update value level changes if they occurred

	return resourceParameterRead(ctx, d, meta)
}

func resourceParameterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterDelete")
	// todo:
	// destroy the parameter and the value?

	return nil
}
