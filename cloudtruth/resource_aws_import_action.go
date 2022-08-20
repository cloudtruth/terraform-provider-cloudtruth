package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAWSImportAction() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth import action..`,

		CreateContext: resourceAWSImportActionCreate,
		ReadContext:   resourceAWSImportActionRead,
		UpdateContext: resourceAWSImportActionUpdate,
		DeleteContext: resourceAWSImportActionDelete,

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
			},
			"create_environments": {
				Description: "Auto create environments",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
			},
			"create_projects": {
				Description: "Auto create projects",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
			},
			"region": {
				Description: "The target AWS region",
				Type:        schema.TypeString,
				Required:    true,
			},
			"service": {
				Description: "The AWS service to pull from: s3|ssm|secretsmanager",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource": {
				Description: "The regex or mustache resource pattern specifying the environment, project, and parameter",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func resourceAWSImportActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSImportActionCreate")
	return nil
}

func resourceAWSImportActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSImportActionRead")
	return nil
}

func resourceAWSImportActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSImportActionUpdate")
	return resourceAWSImportActionRead(ctx, d, meta)
}

func resourceAWSImportActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSImportActionDelete")
	return nil
}
