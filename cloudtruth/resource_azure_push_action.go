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
			"integration": {
				Description:  "The name (using the format VAULT_NAME@TENANT_ID) of the CloudTruth integration corresponding to this import action",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: isValidAzureIntegrationName,
			},
			"integration_id": {
				Description: "The ID of the CloudTruth integration corresponding to this import action",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "A description of the push action",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"parameters": {
				Description: "Include parameters (non-secrets) when pushing, defaults to true",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"secrets": {
				Description: "Include secrets when pushing, defaults to true",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"templates": {
				Description: "Include templates when pushing, defaults to true",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"coerce": {
				Description: "Include secrets/parameters even if the upstream destination doesn't allow them (e.g. non-secrets in AWS SecretsManager), defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"force": {
				Description: "Allow CloudTruth to take ownership and overwrite any pre-existing items, defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"local": {
				Description: "Only send the parameters defined directly in the specified projects (not inherited), defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"dry_run": {
				Description: "When true, the action only reports what it would push without actually pushing changes, defaults to true",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"projects": {
				Description: "The projects containing the parameters to pushed to the AWS destination",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Description: "Tags specified in the form 'environment_name:tag_name' indicating the sync point for parameters to be pushed (multiple tags allowed but only one per environment)",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"resource": {
				Description: "The mustache style resource string specifying the environment, project, and parameter",
				Type:        schema.TypeString,
				Required:    true,
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
