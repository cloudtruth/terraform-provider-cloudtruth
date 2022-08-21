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
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func resourceAWSImportActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSImportActionCreate")
	c := meta.(*cloudTruthClient)
	importActionCreate := cloudtruthapi.NewAwsPullWithDefaults()
	importActionName := d.Get("name").(string)
	awsIntegrationID := d.Get("integration_id").(string)
	importActionDesc := d.Get("description").(string)
	importActionCreate.SetName(importActionName)
	importActionCreate.SetDescription(importActionDesc)
	region, err := cloudtruthapi.NewAwsRegionEnumFromValue(d.Get("region").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	service, err := cloudtruthapi.NewAwsServiceEnumFromValue(d.Get("service").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	resourcePath := d.Get("resource_pattern").(string)
	importActionCreate.SetRegion(*region)
	importActionCreate.SetService(*service)
	importActionCreate.SetResource(resourcePath)

	var resp *cloudtruthapi.AwsPull
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsCreate(ctx, awsIntegrationID).AwsPull(*importActionCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceAWSImportActionCreate: error creating AWS import action %s: %w", importActionName, err)
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

func resourceAWSImportActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSImportActionRead")
	c := meta.(*cloudTruthClient)
	awsIntegrationID := d.Get("integration_id").(string)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()

	var resp *cloudtruthapi.AwsPull
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsRetrieve(ctx, awsIntegrationID, importActionID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceAccessGrantRead: error reading AWS import action %s: %w", importActionName, err)
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

func resourceAWSImportActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSImportActionUpdate")
	c := meta.(*cloudTruthClient)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)
	patchedAWSPull := cloudtruthapi.PatchedAwsPull{}
	hasChange := false
	props := map[string]func(v string){
		"name":             patchedAWSPull.SetName,
		"description":      patchedAWSPull.SetDescription,
		"resource_pattern": patchedAWSPull.SetResource,
	}
	for prop := range props {
		if d.HasChange(prop) {
			props[prop](d.Get(prop).(string))
			hasChange = true
		}
	}
	if d.HasChange("create_environments") {
		patchedAWSPull.SetCreateEnvironments(d.Get("create_environments").(bool))
		hasChange = true
	}
	if d.HasChange("create_projects") {
		patchedAWSPull.SetCreateEnvironments(d.Get("create_projects").(bool))
		hasChange = true
	}

	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsPartialUpdate(ctx, awsIntegrationID,
				importActionID).PatchedAwsPull(patchedAWSPull).Execute()
			if err != nil {
				outErr := fmt.Errorf("resourceEnvironmentUpdate: error updating environment %s: %w", importActionName, err)
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
	}
	d.SetId(importActionID)
	return resourceAWSImportActionRead(ctx, d, meta)
}

func resourceAWSImportActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAWSImportActionDelete")
	c := meta.(*cloudTruthClient)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsDestroy(ctx, awsIntegrationID, importActionID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceAWSImportActionDelete: error destroying AWS import action %s: %w", importActionName, err)
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
