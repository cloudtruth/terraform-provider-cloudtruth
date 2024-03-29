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

func resourceAzureImportAction() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth Azure import action.`,

		CreateContext: resourceAzureImportActionCreate,
		ReadContext:   resourceAzureImportActionRead,
		UpdateContext: resourceAzureImportActionUpdate,
		DeleteContext: resourceAzureImportActionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the import action",
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
			"resource": {
				Description: "The regex or mustache resource string specifying the environment, project, and parameter",
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

func resourceAzureImportActionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionCreate")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionCreate")
	c := meta.(*cloudTruthClient)
	importActionCreate := cloudtruthapi.NewAzureKeyVaultPullWithDefaults()
	importActionName := d.Get("name").(string)
	azureIntegrationName := d.Get("integration").(string)

	azureIntegrationID, err := lookupAzureIntegration(ctx, azureIntegrationName, c, d)
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
	resourcePath := d.Get("resource").(string)
	importActionCreate.SetResource(resourcePath)

	var azureImport *cloudtruthapi.AzureKeyVaultPull
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		azureImport, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPullsCreate(ctx, *azureIntegrationID).AzureKeyVaultPull(*importActionCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzureImportActionCreate: error creating Azure import action %s", importActionName), r, err)
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
	d.SetId(azureImport.GetId())
	return nil
}

// todo: figure how/if we should read integration_id, it's not a top level property on Pull objects
func setAzureKeyVaultImportReadProps(pull *cloudtruthapi.AzureKeyVaultPull, d *schema.ResourceData) error {
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

func resourceAzureImportActionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionRead")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionRead")
	c := meta.(*cloudTruthClient)
	azureIntegrationID := d.Get("integration_id").(string)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()

	var azureKeyVaultPull *cloudtruthapi.AzureKeyVaultPull
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error
		azureKeyVaultPull, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPullsRetrieve(ctx, azureIntegrationID, importActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzureImportActionRead: error reading Azure import action %s with ID %s",
				importActionName, importActionID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	err := setAzureKeyVaultImportReadProps(azureKeyVaultPull, d)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(azureKeyVaultPull.GetId())
	return nil
}

func resourceAzureImportActionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionUpdate")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionUpdate")
	c := meta.(*cloudTruthClient)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()
	azureIntegrationID := d.Get("integration_id").(string)
	patchedAzureKeyVaultPull := cloudtruthapi.PatchedAzureKeyVaultPull{}
	hasChange := false
	props := map[string]func(v string){
		"name":        patchedAzureKeyVaultPull.SetName,
		"description": patchedAzureKeyVaultPull.SetDescription,
		"resource":    patchedAzureKeyVaultPull.SetResource,
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
		patchedAzureKeyVaultPull.SetMode(*importActionMode)
	}
	if d.HasChange("create_environments") {
		patchedAzureKeyVaultPull.SetCreateEnvironments(d.Get("create_environments").(bool))
		hasChange = true
	}
	if d.HasChange("create_projects") {
		patchedAzureKeyVaultPull.SetCreateProjects(d.Get("create_projects").(bool))
		hasChange = true
	}

	if hasChange {
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPullsPartialUpdate(ctx, azureIntegrationID,
				importActionID).PatchedAzureKeyVaultPull(patchedAzureKeyVaultPull).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceAzureImportActionUpdate: error updating Azure import action %s", importActionName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}

	d.SetId(importActionID)
	return resourceAzureImportActionRead(ctx, d, meta)
}

func resourceAzureImportActionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureImportActionDelete")
	defer tflog.Debug(ctx, "exiting resourceAzureImportActionDelete")
	c := meta.(*cloudTruthClient)
	importActionName := d.Get("name").(string)
	importActionID := d.Id()
	azureIntegrationID := d.Get("integration_id").(string)

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPullsDestroy(ctx, azureIntegrationID, importActionID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzureImportActionDelete: error destroying Azure import action %s", importActionName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
