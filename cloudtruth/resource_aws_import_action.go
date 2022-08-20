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
	importActionName := d.Get("name").(string)
	awsIntegrationID := d.Get("integration_id").(string)
	importActionCreate := cloudtruthapi.NewAwsPullWithDefaults()
	importActionCreate.SetName(importActionName)
	importActionDesc := d.Get("description").(string)
	if importActionDesc != "" {
		importActionCreate.SetDescription(importActionDesc)
	}
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
