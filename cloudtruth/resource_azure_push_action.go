package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzurePushAction() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth push action.`,

		CreateContext: resourceAzurePushActionCreate,
		ReadContext:   resourceAzurePushActionRead,
		UpdateContext: resourceAzurePushActionUpdate,
		DeleteContext: resourceAzurePushActionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the push action",
				Type:        schema.TypeString,
				Required:    true,
			},
			"integration_id": {
				Description: "The ID of the CloudTruth integration corresponding to this pull action",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description of the push action",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceAzurePushActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionCreate")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionCreate")
	return nil
}

func resourceAzurePushActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionRead")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionRead")
	return nil
}

func resourceAzurePushActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionUpdate")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionUpdate")
	return resourceAzurePushActionRead(ctx, d, meta)
}

func resourceAzurePushActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionDelete")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionDelete")
	return nil
}
