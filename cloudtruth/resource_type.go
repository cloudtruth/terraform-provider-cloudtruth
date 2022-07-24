package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceType() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Tag - unique per environment, defaulting to the 'default' environment",

		CreateContext: resourceTypeCreate,
		ReadContext:   resourceTypeRead,
		UpdateContext: resourceTypeUpdate,
		DeleteContext: resourceTypeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Tag, unique per project, per environment",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Tag",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// rules TypeList
		},
	}
}

func resourceTypeCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeCreate")
	return nil
}

func resourceTypeRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeRead")
	return dataCloudTruthTagRead(ctx, d, meta)
}

func resourceTypeUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeUpdate")
	return resourceTypeRead(ctx, d, meta)
}

func resourceTypeDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeDelete")
	return nil
}
