package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAWSIntegration() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth AWS integration.`,

		CreateContext: resourceAWSIntegrationCreate,
		ReadContext:   resourceAWSIntegrationRead,
		UpdateContext: resourceAWSIntegrationUpdate,
		DeleteContext: resourceAWSIntegrationDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "The ID of the source AWS account",
				Type:        schema.TypeString,
				Required:    true,
			},
			"role": {
				Description: "The name of the role which CloudTruth will assume in the AWS account",
				Type:        schema.TypeString,
				Required:    true,
			},
			"kms_key_id": {
				Description: "The name of the role which CloudTruth will assume in the AWS account",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			// services + write access + regions
		},
	}
}

func resourceAWSIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationCreate")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationCreate")
	return nil
}

func resourceAWSIntegrationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationRead")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationRead")
	return nil
}

func resourceAWSIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationUpdate")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationUpdate")
	return resourceAWSIntegrationRead(ctx, d, meta)
}

func resourceAWSIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationDelete")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationDelete")
	return nil
}
