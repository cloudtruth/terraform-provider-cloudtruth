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

func resourceAWSPushAction() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth push action.`,

		CreateContext: resourceAWSPushActionCreate,
		ReadContext:   resourceAWSPushActionRead,
		UpdateContext: resourceAWSPushActionUpdate,
		DeleteContext: resourceAWSPushActionDelete,

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
			"resource_pattern": {
				Description: "The regex or mustache resource pattern specifying the environment, project, and parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceAWSPushActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSPushActionCreate")
	c := meta.(*cloudTruthClient)
	pushActionCreate := cloudtruthapi.NewAwsPushWithDefaults()
	pushActionName := d.Get("name").(string)
	awsIntegrationID := d.Get("integration_id").(string)
	pushActionDesc := d.Get("description").(string)
	pushActionCreate.SetName(pushActionName)
	pushActionCreate.SetDescription(pushActionDesc)
	region, err := cloudtruthapi.NewAwsRegionEnumFromValue(d.Get("region").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	// todo:
	// add the remaining push action properties and refactor
	service, err := cloudtruthapi.NewAwsServiceEnumFromValue(d.Get("service").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	resourcePath := d.Get("resource_pattern").(string)
	pushActionCreate.SetRegion(*region)
	pushActionCreate.SetService(*service)
	pushActionCreate.SetResource(resourcePath)
	rawProjects := d.Get("projects").([]interface{})
	projects := make([]string, len(rawProjects))
	for i, v := range projects {
		projects[i] = fmt.Sprint(v)
	}
	pushActionCreate.SetProjects(projects)
	rawTags := d.Get("tags").([]interface{})
	tags := make([]string, len(rawTags))
	for i, v := range tags {
		tags[i] = fmt.Sprint(v)
	}
	pushActionCreate.SetTags(tags)

	var resp *cloudtruthapi.AwsPush
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPushesCreate(ctx, awsIntegrationID).AwsPush(*pushActionCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceAWSPushActionCreate: error creating AWS push action %s: %w", pushActionName, err)
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

	d.SetId(resp.GetId())
	return nil
}

func resourceAWSPushActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSPushActionRead")
	return nil
}

func resourceAWSPushActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSPushActionUpdate")
	return resourceAWSPushActionRead(ctx, d, meta)
}

func resourceAWSPushActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSPushActionDelete")
	c := meta.(*cloudTruthClient)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPushesDestroy(ctx, awsIntegrationID, pushActionID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceAWSPushActionDelete: error destroying AWS push action %s: %w", pushActionName, err)
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
