package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
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
				Description: "Include secrets/parameters even if the upstream destination doesn't allow them, defaults to false",
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
				Description: "The projects containing the parameters to pushed to the Azure destination",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Description: "Tags specified in the form 'environment_name:tag_name' indicating the sync point for parameters to be pushed (multiple tags allowed but only one per environment)",
				Type:        schema.TypeSet,
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

// todo: use generics and make one function for both AWS and Azure types
func setAzurePushActionBoolProps(pushAction *cloudtruthapi.AzureKeyVaultPush, d *schema.ResourceData) {
	pushAction.SetIncludeParameters(d.Get("parameters").(bool))
	pushAction.SetIncludeSecrets(d.Get("secrets").(bool))
	pushAction.SetIncludeTemplates(d.Get("templates").(bool))
	pushAction.SetCoerceParameters(d.Get("coerce").(bool))
	pushAction.SetForce(d.Get("force").(bool))
	pushAction.SetLocal(d.Get("local").(bool))
	pushAction.SetDryRun(d.Get("dry_run").(bool))
}

func resourceAzurePushActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionCreate")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionCreate")
	c := meta.(*cloudTruthClient)
	pushActionCreate := cloudtruthapi.NewAzureKeyVaultPushWithDefaults()
	pushActionName := d.Get("name").(string)
	azureIntegrationName := d.Get("integration").(string)
	azureIntegrationID, err := lookupAzureIntegration(ctx, azureIntegrationName, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	pushActionDesc := d.Get("description").(string)
	pushActionCreate.SetName(pushActionName)
	pushActionCreate.SetDescription(pushActionDesc)
	resourcePath := d.Get("resource").(string)

	pushActionCreate.SetResource(resourcePath)
	setAzurePushActionBoolProps(pushActionCreate, d)
	rawProjects := d.Get("projects").(*schema.Set)
	projects, err := getProjectURLs(ctx, c, rawProjects.List())
	if err != nil {
		return diag.FromErr(err)
	}
	pushActionCreate.SetProjects(projects)
	rawTags := d.Get("tags").(*schema.Set)
	tags, err := getEnvTags(ctx, d, c, rawTags.List())
	pushActionCreate.SetTags(tags)
	if err != nil {
		return diag.FromErr(err)
	}

	var azurePush *cloudtruthapi.AzureKeyVaultPush
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		azurePush, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPushesCreate(ctx, *azureIntegrationID).AzureKeyVaultPush(*pushActionCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzurePushActionCreate: error creating Azure push action %s", pushActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	err = d.Set("integration_id", azureIntegrationID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(azurePush.GetId())
	return nil
}

func resourceAzurePushActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionRead")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionRead")
	c := meta.(*cloudTruthClient)
	azureIntegrationID := d.Get("integration_id").(string)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()

	var azurePush *cloudtruthapi.AzureKeyVaultPush
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error
		azurePush, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPushesRetrieve(ctx, azureIntegrationID, pushActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzurePushActionRead: error reading Azure push action %s", pushActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// Read & set all properties from the deployed resource
	props := map[string]func() string{
		"name":        azurePush.GetName,
		"description": azurePush.GetDescription,
		"resource":    azurePush.GetResource,
	}
	var err error
	for prop := range props {
		err = d.Set(prop, props[prop]())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	boolProps := map[string]func() bool{
		"parameters": azurePush.GetIncludeParameters,
		"secrets":    azurePush.GetIncludeSecrets,
		"templates":  azurePush.GetIncludeTemplates,
		"coerce":     azurePush.GetCoerceParameters,
		"force":      azurePush.GetForce,
		"local":      azurePush.GetLocal,
		"dry_run":    azurePush.GetDryRun,
	}
	for prop := range boolProps {
		val := boolProps[prop]()
		err = d.Set(prop, val)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	tags, err := getTags(ctx, azurePush.GetTags(), meta)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("tags", tags)
	if err != nil {
		return diag.FromErr(err)
	}
	projects, err := getProjects(ctx, azurePush.GetProjects(), meta)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("projects", projects)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(azurePush.GetId())
	return nil
}

func resourceAzurePushActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionUpdate")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionUpdate")
	c := meta.(*cloudTruthClient)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()
	azureIntegrationID := d.Get("integration_id").(string)
	patchedAzurePush := cloudtruthapi.PatchedAzureKeyVaultPushUpdate{}
	hasChange := false
	props := map[string]func(v string){
		"name":        patchedAzurePush.SetName,
		"description": patchedAzurePush.SetDescription,
		"resource":    patchedAzurePush.SetResource,
	}
	for prop := range props {
		if d.HasChange(prop) {
			props[prop](d.Get(prop).(string))
			hasChange = true
		}
	}
	boolProps := map[string]func(v bool){
		"parameters": patchedAzurePush.SetIncludeParameters,
		"secrets":    patchedAzurePush.SetIncludeSecrets,
		// todo: determine if this a backend bug or if we need to update the Go API bindings
		//"templates":  patchedAzurePush.SetIncludeTemplates,
		"coerce":  patchedAzurePush.SetCoerceParameters,
		"force":   patchedAzurePush.SetForce,
		"local":   patchedAzurePush.SetLocal,
		"dry_run": patchedAzurePush.SetDryRun,
	}
	for prop := range boolProps {
		if d.HasChange(prop) {
			boolProps[prop](d.Get(prop).(bool))
			hasChange = true
		}
	}

	if d.HasChange("tags") {
		rawTags := d.Get("tags").(*schema.Set)
		tags, err := getEnvTags(ctx, d, c, rawTags.List())
		if err != nil {
			return diag.FromErr(err)
		}
		patchedAzurePush.SetTags(tags)
		hasChange = true
	}

	if d.HasChange("projects") {
		rawProjects := d.Get("projects").([]interface{})
		projects, err := getProjectURLs(ctx, c, rawProjects)
		if err != nil {
			return diag.FromErr(err)
		}
		patchedAzurePush.SetProjects(projects)
		hasChange = true
	}

	if hasChange {
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPushesPartialUpdate(ctx, azureIntegrationID,
				pushActionID).PatchedAzureKeyVaultPushUpdate(patchedAzurePush).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceAzurePushActionUpdate: error updating Azure push action %s", pushActionName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	d.SetId(pushActionID)
	return resourceAzurePushActionRead(ctx, d, meta)
}

func resourceAzurePushActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzurePushActionDelete")
	defer tflog.Debug(ctx, "exiting resourceAzurePushActionDelete")
	c := meta.(*cloudTruthClient)
	pushActionName := d.Get("name").(string)
	pushActionID := d.Id()
	azureIntegrationID := d.Get("integration_id").(string)

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPushesDestroy(ctx, azureIntegrationID, pushActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzurePushActionDelete: error destroying Azure push action %s", pushActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
