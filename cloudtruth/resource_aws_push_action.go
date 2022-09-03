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
			"integration": {
				Description: "The name of the CloudTruth integration corresponding to this pull action",
				Type:        schema.TypeString,
				Required:    true,
			},
			"integration_id": {
				Description: "The ID of the CloudTruth integration corresponding to this pull action",
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
				Description: "The mustache style resource string specifying the environment, project, and parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceAWSPushActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSPushActionCreate")
	defer tflog.Debug(ctx, "exiting resourceAWSPushActionCreate")
	c := meta.(*cloudTruthClient)
	pushActionCreate := cloudtruthapi.NewAwsPushWithDefaults()
	pushActionName := d.Get("name").(string)
	awsIntegrationID := d.Get("integration").(string)
	pushActionDesc := d.Get("description").(string)
	pushActionCreate.SetName(pushActionName)
	pushActionCreate.SetDescription(pushActionDesc)
	region, err := cloudtruthapi.NewAwsRegionEnumFromValue(d.Get("region").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	service, err := cloudtruthapi.NewAwsServiceEnumFromValue(d.Get("service").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	resourcePath := d.Get("resource").(string)
	pushActionCreate.SetRegion(*region)
	pushActionCreate.SetService(*service)
	pushActionCreate.SetResource(resourcePath)
	rawProjects := d.Get("projects").([]interface{})
	projects := make([]string, len(rawProjects))
	for i, v := range projects {
		projID, err := c.lookupProject(ctx, fmt.Sprint(v))
		if err != nil {
			return diag.FromErr(err)
		}
		projURL := fmt.Sprintf("https://api.cloudtruth.io/api/v1/projects/%s/", *projID)
		projects[i] = projURL
	}
	pushActionCreate.SetProjects(projects)
	rawTags := d.Get("tags").([]interface{})
	tags := make([]string, len(rawTags))
	for i, v := range rawTags {
		tags[i], err = lookupEnvTag(ctx, d, c, fmt.Sprint(v))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	pushActionCreate.SetTags(tags)
	pushActionCreate.SetIncludeParameters(d.Get("parameters").(bool))
	pushActionCreate.SetIncludeSecrets(d.Get("secrets").(bool))
	pushActionCreate.SetIncludeTemplates(d.Get("templates").(bool))
	pushActionCreate.SetCoerceParameters(d.Get("coerce").(bool))
	pushActionCreate.SetForce(d.Get("force").(bool))
	pushActionCreate.SetLocal(d.Get("local").(bool))
	pushActionCreate.SetDryRun(d.Get("dry_run").(bool))

	var resp *cloudtruthapi.AwsPush
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPushesCreate(ctx, awsIntegrationID).AwsPush(*pushActionCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSPushActionCreate: error creating AWS push action %s", pushActionName), r, err)
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
	tflog.Debug(ctx, "entering resourceAWSPushActionRead")
	defer tflog.Debug(ctx, "exiting resourceAWSPushActionRead")
	c := meta.(*cloudTruthClient)
	awsIntegrationID := d.Get("integration_id").(string)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()

	var awsPush *cloudtruthapi.AwsPush
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		awsPush, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPushesRetrieve(ctx, awsIntegrationID, pushActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSPushActionRead: error reading AWS push action %s", pushActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	props := map[string]func() string{
		"name":        awsPush.GetName,
		"description": awsPush.GetDescription,
		"resource":    awsPush.GetResource,
	}
	var err error
	for prop := range props {
		err = d.Set(prop, props[prop]())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	boolProps := map[string]func() bool{
		"parameters": awsPush.GetIncludeParameters,
		"secrets":    awsPush.GetIncludeSecrets,
		"templates":  awsPush.GetIncludeTemplates,
		"coerce":     awsPush.GetCoerceParameters,
		"force":      awsPush.GetForce,
		"local":      awsPush.GetLocal,
		"dry_run":    awsPush.GetDryRun,
	}
	for prop := range boolProps {
		val := boolProps[prop]()
		err = d.Set(prop, val)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(awsPush.GetId())
	return nil
}

func resourceAWSPushActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSPushActionUpdate")
	defer tflog.Debug(ctx, "exiting resourceAWSPushActionUpdate")
	c := meta.(*cloudTruthClient)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)
	patchedAWSPush := cloudtruthapi.PatchedAwsPushUpdate{}
	hasChange := false
	props := map[string]func(v string){
		"name":        patchedAWSPush.SetName,
		"description": patchedAWSPush.SetDescription,
		"resource":    patchedAWSPush.SetResource,
	}
	for prop := range props {
		if d.HasChange(prop) {
			props[prop](d.Get(prop).(string))
			hasChange = true
		}
	}
	boolProps := map[string]func(v bool){
		"parameters": patchedAWSPush.SetIncludeParameters,
		"secrets":    patchedAWSPush.SetIncludeSecrets,
		"templates":  patchedAWSPush.SetIncludeTemplates,
		"coerce":     patchedAWSPush.SetCoerceParameters,
		"force":      patchedAWSPush.SetForce,
		"local":      patchedAWSPush.SetLocal,
		"dry_run":    patchedAWSPush.SetDryRun,
	}
	for prop := range boolProps {
		if d.HasChange(prop) {
			boolProps[prop](d.Get(prop).(bool))
			hasChange = true
		}
	}

	if d.HasChange("tags") {
		rawTags := d.Get("tags").([]interface{})
		tags := make([]string, len(rawTags))
		var err error
		for i, v := range rawTags {
			tags[i], err = lookupEnvTag(ctx, d, c, fmt.Sprint(v))
			if err != nil {
				return diag.FromErr(err)
			}
		}
		patchedAWSPush.SetTags(tags)
		hasChange = true
	}

	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPushesPartialUpdate(ctx, awsIntegrationID,
				pushActionID).PatchedAwsPushUpdate(patchedAWSPush).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceAWSPushActionUpdate: error updating AWS push action %s", pushActionName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	d.SetId(pushActionID)
	return resourceAWSPushActionRead(ctx, d, meta)
}

func resourceAWSPushActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSPushActionDelete")
	defer tflog.Debug(ctx, "exiting resourceAWSPushActionDelete")
	c := meta.(*cloudTruthClient)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()
	awsIntegrationID := d.Get("integration_id").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsPushesDestroy(ctx, awsIntegrationID, pushActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSPushActionDelete: error destroying AWS push action %s", pushActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
