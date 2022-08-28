package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzureImportAction() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth Azure import action.`,

		CreateContext: resourceAzureImportActionCreate,
		ReadContext:   resourceAzureImportActionRead,
		UpdateContext: resourceAzureImportActionUpdate,
		DeleteContext: resourceAzureImportActionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the import action",
				Type:        schema.TypeString,
				Required:    true,
			},
			"integration_id": {
				Description: "The ID of the CloudTruth integration corresponding to this pull action",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description of the import action",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"create_environments": {
				Description: "Auto create environments",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"create_projects": {
				Description: "Auto create projects",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"resource": {
				Description: "The regex or mustache resource string specifying the environment, project, and parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
			"mode": {
				Description: "The resource type: 'pattern' for mustache-style regexes and 'mapped' for JMESpath expression",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceAzureImportActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionCreate")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionCreate")
	return nil
}

func resourceAzureImportActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionRead")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionRead")
	return nil
}

func resourceAzureImportActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionUpdate")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionUpdate")
	return resourceAzureImportActionRead(ctx, d, meta)
}

func resourceAzureImportActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionDelete")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionDelete")
	return nil
}
