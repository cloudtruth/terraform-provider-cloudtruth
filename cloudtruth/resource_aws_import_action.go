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
		Description: `A CloudTruth AWS import action.`,

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
			"integration": {
				Description:  "The name (using the format AWS_ROLE@AWS_ACCOUNT_ID) of the CloudTruth integration corresponding to this pull action",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: isValidAWSIntegrationName,
			},
			"integration_id": {
				Description: "The ID of the CloudTruth integration corresponding to this pull action",
				Type:        schema.TypeString,
				Computed:    true,
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
			"resource": {
				Description: "The mustache style or JMESPath resource string specifying the environment, project, and parameter",
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

func resourceAWSImportActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSImportActionCreate")
	defer tflog.Debug(ctx, "exiting resourceAWSImportActionCreate")
	c := meta.(*cloudTruthClient)
	importActionCreate := cloudtruthapi.NewAwsPullWithDefaults()
	importActionName := d.Get("name").(string)
	awsIntegrationName := d.Get("integration").(string)
	awsIntegrationID, err := lookupAWSIntegration(ctx, awsIntegrationName, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	importActionCreate.SetName(importActionName)
	importActionCreate.SetDescription(d.Get("description").(string))
	importActionMode, err := cloudtruthapi.NewModeEnumFromValue(d.Get("mode").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	importActionCreate.SetMode(*importActionMode)
	region, err := cloudtruthapi.NewAwsRegionEnumFromValue(d.Get("region").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	service, err := cloudtruthapi.NewAwsServiceEnumFromValue(d.Get("service").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	resourcePath := d.Get("resource").(string)
	importActionCreate.SetRegion(*region)
	importActionCreate.SetService(*service)
	importActionCreate.SetResource(resourcePath)

	var awsPull *cloudtruthapi.AwsPull
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		awsPull, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsCreate(ctx, *awsIntegrationID).AwsPull(*importActionCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSImportActionCreate: error creating AWS import action %s", importActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err = d.Set("integration_id", awsIntegrationID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(awsPull.GetId())
	return nil
}

func resourceAWSImportActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSImportActionRead")
	defer tflog.Debug(ctx, "exiting resourceAWSImportActionRead")
	c := meta.(*cloudTruthClient)
	awsIntegrationID := d.Get("integration_id").(string)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()

	var awsPull *cloudtruthapi.AwsPull
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		awsPull, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsRetrieve(ctx, awsIntegrationID, importActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSImportActionRead: error reading AWS import action %s", importActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// Read & set all properties from the deployed resource
	err := setAWSImportReadProps(awsPull, d)
	if retryError != nil {
		return diag.FromErr(err)
	}
	d.SetId(awsPull.GetId())
	return nil
}

// todo: figure how and if we should read integration_id, it's not a top level property on Pull objects
// omitting region and service for now since they are not settable in the UI
// Also, the disparate return types make it less appealing to handle all of these Set() calls in a loop or two
func setAWSImportReadProps(pull *cloudtruthapi.AwsPull, d *schema.ResourceData) error {
	err := d.Set("name", pull.GetName())
	if err != nil {
		return err
	}
	err = d.Set("description", pull.GetDescription())
	if err != nil {
		return err
	}
	err = d.Set("resource", pull.GetResource())
	if err != nil {
		return err
	}
	err = d.Set("mode", pull.GetMode())
	if err != nil {
		return err
	}
	err = d.Set("create_projects", pull.GetCreateProjects())
	if err != nil {
		return err
	}
	err = d.Set("mode", pull.GetMode())
	if err != nil {
		return err
	}
	err = d.Set("create_environments", pull.GetCreateEnvironments())
	if err != nil {
		return err
	}
	return nil
}

// Some fields (e.g. integration id/name) are not settable in the UI, we are skipping them for now
func resourceAWSImportActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSImportActionUpdate")
	defer tflog.Debug(ctx, "exiting resourceAWSImportActionUpdate")
	c := meta.(*cloudTruthClient)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)
	patchedAWSPull := cloudtruthapi.PatchedAwsPull{}
	hasChange := false
	props := map[string]func(v string){
		"name":        patchedAWSPull.SetName,
		"description": patchedAWSPull.SetDescription,
		"resource":    patchedAWSPull.SetResource,
	}
	for prop := range props {
		if d.HasChange(prop) {
			props[prop](d.Get(prop).(string))
			hasChange = true
		}
	}
	if d.HasChange("mode") {
		importActionMode, err := cloudtruthapi.NewModeEnumFromValue(d.Get("mode").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		patchedAWSPull.SetMode(*importActionMode)
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
				return handleAPIError(fmt.Sprintf("resourceAWSImportActionUpdate: error updating AWS import action %s", importActionName), r, err)
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
	tflog.Debug(ctx, "entering resourceAWSImportActionDelete")
	defer tflog.Debug(ctx, "exiting resourceAWSImportActionDelete")
	c := meta.(*cloudTruthClient)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPullsDestroy(ctx, awsIntegrationID, importActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSImportActionDelete: error destroying AWS import action %s", importActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
