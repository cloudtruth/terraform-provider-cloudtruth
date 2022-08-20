package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAWSPushAction() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth import action..`,

		CreateContext: resourceAWSPushActionCreate,
		ReadContext:   resourceAWSPushActionRead,
		UpdateContext: resourceAWSPushActionUpdate,
		DeleteContext: resourceAWSPushActionDelete,

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
			"coerce_parameters": {
				Description: "Auto create environments",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"include_parameters": {
				Description: "Include parameters (non-secrets) when pushing",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"coerce": {
				Description: "Include secrets/parameters even if the upstream destination doesn't allow them (e.g. non-secrets in AWS SecretsManager)",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"force": {
				Description: "Allow CloudTruth to take ownership and overwrite any pre-existing items",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"local": {
				Description: "Only send the parameters defined directly in the specified projects (not inherited)",
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
				Description: "The environment tags specifying the sync point for parameters to be pushed (multiple allowed, but only one per environment)",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Required:    true,
			},
		},
	}
}

func resourceAWSPushActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSPushActionCreate")
	return nil
}

func resourceAWSPushActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSPushActionRead")
	return nil
}

func resourceAWSPushActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSPushActionUpdate")
	return resourceAWSPushActionRead(ctx, d, meta)
}

func resourceAWSPushActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAWSPushActionDelete")
	return nil
}
