package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"net/http"
	"regexp"
	"strings"
)

// isValidAzureVaultName validates Azure Key Vault names according to Azure requirements:
// - Must be 1-24 characters long
// - Must contain only alphanumeric characters and hyphens
// - Cannot start or end with a hyphen
// - Cannot contain consecutive hyphens
func isValidAzureVaultName(val any, key string) (warns []string, errs []error) {
	v := val.(string)

	if len(v) < 1 || len(v) > 24 {
		errs = append(errs, fmt.Errorf("%q must be between 1 and 24 characters, got: %d", key, len(v)))
	}

	// Check for valid characters (alphanumeric and hyphens only)
	validPattern := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !validPattern.MatchString(v) {
		errs = append(errs, fmt.Errorf("%q must contain only alphanumeric characters and hyphens", key))
	}

	// Check for leading or trailing hyphens
	if len(v) > 0 && (v[0] == '-' || v[len(v)-1] == '-') {
		errs = append(errs, fmt.Errorf("%q cannot start or end with a hyphen", key))
	}

	// Check for consecutive hyphens
	consecutiveHyphens := regexp.MustCompile(`--`)
	if consecutiveHyphens.MatchString(v) {
		errs = append(errs, fmt.Errorf("%q cannot contain consecutive hyphens", key))
	}

	return warns, errs
}

func resourceAzureIntegration() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth Azure Key Vault integration.`,

		CreateContext: resourceAzureIntegrationCreate,
		ReadContext:   resourceAzureIntegrationRead,
		UpdateContext: resourceAzureIntegrationUpdate,
		DeleteContext: resourceAzureIntegrationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceAzureIntegrationImport,
		},

		Schema: map[string]*schema.Schema{
			"vault_name": {
				Description:  "The Azure Key Vault name (1-24 alphanumeric characters and hyphens)",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: isValidAzureVaultName,
			},
			"tenant_id": {
				Description:  "The Azure Tenant ID (UUID format)",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
			"description": {
				Description: "An optional description for the integration",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"writable": {
				Description: "Whether or not the CloudTruth integration can write to Azure Key Vault, defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"resource_tags": {
				Description: "A list of tags to be set on all integration resources",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Default: nil,
			},
			"id": {
				Description: "The unique identifier for the integration",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The integration name in format 'vault_name@tenant_id'",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"fqn": {
				Description: "The fully qualified name of the integration (e.g., 'akv://vault_name@tenant_id/secrets/')",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "The connection status of the integration (checking, connected, errored, unknown)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status_detail": {
				Description: "Additional details about the integration status, especially error information",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status_last_checked_at": {
				Description: "The timestamp when the integration status was last checked",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "The timestamp when the integration was created",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"modified_at": {
				Description: "The timestamp when the integration was last modified",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of integration (always 'akv' for Azure Key Vault)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Description: "The API URL for this integration resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceAzureIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureIntegrationCreate")
	defer tflog.Debug(ctx, "exiting resourceAzureIntegrationCreate")
	c := meta.(*cloudTruthClient)

	vaultName := d.Get("vault_name").(string)
	tenantID := d.Get("tenant_id").(string)

	intCreate := cloudtruthapi.NewAzureKeyVaultIntegrationCreateWithDefaults()
	intCreate.SetVaultName(vaultName)
	intCreate.SetTenantId(tenantID)

	// Optional properties
	if v, ok := d.GetOk("description"); ok {
		intCreate.SetDescription(v.(string))
	}
	intCreate.SetWritable(d.Get("writable").(bool))

	if resourceTags, ok := d.GetOk("resource_tags"); ok {
		intCreate.SetResourceTags(resourceTags.(map[string]any))
	}

	var integration *cloudtruthapi.AzureKeyVaultIntegration
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		integration, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultCreate(ctx).AzureKeyVaultIntegrationCreate(*intCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzureIntegrationCreate: error creating integration %s@%s", vaultName, tenantID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// Set the ID and all computed fields
	d.SetId(integration.GetId())
	if err := setAzureIntegrationComputedFields(d, integration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceAzureIntegrationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureIntegrationRead")
	defer tflog.Debug(ctx, "exiting resourceAzureIntegrationRead")
	c := meta.(*cloudTruthClient)

	integrationID := d.Id()
	vaultName := d.Get("vault_name").(string)
	tenantID := d.Get("tenant_id").(string)

	var integration *cloudtruthapi.AzureKeyVaultIntegration
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error

		// Query by vault_name and tenant_id to find the integration
		integrations, r, err := c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultList(ctx).VaultName(vaultName).TenantId(tenantID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAzureIntegrationRead: error reading Azure integration %s@%s with ID %s",
				vaultName, tenantID, integrationID), r, err)
		}

		// Find the integration with matching ID
		results := integrations.GetResults()
		if len(results) == 0 {
			// Integration not found - it was deleted outside Terraform
			d.SetId("")
			return nil
		}

		// Find matching integration by ID
		found := false
		for _, integ := range results {
			if integ.GetId() == integrationID {
				integration = &integ
				found = true
				break
			}
		}

		if !found {
			// Integration with this ID not found
			d.SetId("")
			return nil
		}

		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// If integration was not found (ID cleared), return early
	if d.Id() == "" {
		return nil
	}

	// Set all fields from the API response
	if err := d.Set("description", integration.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("writable", integration.GetWritable()); err != nil {
		return diag.FromErr(err)
	}

	resourceTags := integration.GetResourceTags()
	if len(resourceTags) > 0 {
		if err := d.Set("resource_tags", resourceTags); err != nil {
			return diag.FromErr(err)
		}
	}

	// Set computed fields
	if err := setAzureIntegrationComputedFields(d, integration); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceAzureIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureIntegrationUpdate")
	defer tflog.Debug(ctx, "exiting resourceAzureIntegrationUpdate")
	c := meta.(*cloudTruthClient)

	integrationID := d.Id()
	vaultName := d.Get("vault_name").(string)
	tenantID := d.Get("tenant_id").(string)

	patchedIntegration := cloudtruthapi.PatchedAzureKeyVaultIntegration{}
	hasChange := false

	if d.HasChange("description") {
		patchedIntegration.SetDescription(d.Get("description").(string))
		hasChange = true
	}

	if d.HasChange("writable") {
		patchedIntegration.SetWritable(d.Get("writable").(bool))
		hasChange = true
	}

	if d.HasChange("resource_tags") {
		resourceTagMap := d.Get("resource_tags").(map[string]any)
		patchedIntegration.SetResourceTags(resourceTagMap)
		hasChange = true
	}

	if hasChange {
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultPartialUpdate(ctx, integrationID).PatchedAzureKeyVaultIntegration(patchedIntegration).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceAzureIntegrationUpdate: error updating Azure integration %s@%s",
					vaultName, tenantID), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}

	return resourceAzureIntegrationRead(ctx, d, meta)
}

func resourceAzureIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAzureIntegrationDelete")
	defer tflog.Debug(ctx, "exiting resourceAzureIntegrationDelete")
	c := meta.(*cloudTruthClient)

	vaultName := d.Get("vault_name").(string)
	tenantID := d.Get("tenant_id").(string)
	intID := d.Id()

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultDestroy(ctx, intID).Execute()
		if err != nil {
			// Special case: If the integration still has push actions, the API returns 400 with this message
			// This is a timing issue - the actions may have just been deleted but the API hasn't caught up yet
			// Treat this as retryable so we wait for the API to process the action deletions
			if r != nil && r.StatusCode == 400 && strings.Contains(fmt.Sprintf("%s", r.Body), "Cannot delete an integration with pushes") {
				errMsg := fmt.Errorf("resourceAzureIntegrationDelete: Azure integration %s@%s still has push actions, waiting for API to process deletions", vaultName, tenantID)
				tflog.Warn(ctx, "Retrying integration deletion after push actions cleanup", map[string]any{
					"integration": fmt.Sprintf("%s@%s", vaultName, tenantID),
					"error":       errMsg.Error(),
				})
				return retry.RetryableError(errMsg)
			}
			return handleAPIError(fmt.Sprintf("resourceAzureIntegrationDelete: error destroying Azure integration %s@%s", vaultName, tenantID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	return nil
}

func resourceAzureIntegrationImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "entering resourceAzureIntegrationImport")
	defer tflog.Debug(ctx, "exiting resourceAzureIntegrationImport")
	c := meta.(*cloudTruthClient)

	importID := d.Id()

	// Try to import by integration name format (vault_name@tenant_id)
	if integrationID, err := lookupAzureIntegration(ctx, importID, c, d); err == nil && integrationID != nil {
		d.SetId(*integrationID)

		// Perform a read to populate all fields
		diags := resourceAzureIntegrationRead(ctx, d, meta)
		if diags.HasError() {
			return nil, fmt.Errorf("error reading Azure integration during import: %v", diags)
		}

		return []*schema.ResourceData{d}, nil
	}

	// Otherwise, assume it's a UUID and try to read directly
	d.SetId(importID)

	// We need to get vault_name and tenant_id from the API first
	var integration *cloudtruthapi.AzureKeyVaultIntegration
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error

		// List all integrations and find by ID
		integrations, r, err := c.openAPIClient.IntegrationsAPI.IntegrationsAzureKeyVaultList(ctx).Execute()
		if err != nil {
			return handleAPIError("resourceAzureIntegrationImport: error listing Azure integrations", r, err)
		}

		// Find the integration with matching ID
		results := integrations.GetResults()
		found := false
		for _, integ := range results {
			if integ.GetId() == importID {
				integration = &integ
				found = true
				break
			}
		}

		if !found {
			return retry.NonRetryableError(fmt.Errorf("Azure integration with ID %s not found", importID))
		}

		return nil
	})
	if retryError != nil {
		return nil, retryError
	}

	// Set vault_name and tenant_id from the API response
	if err := d.Set("vault_name", integration.GetVaultName()); err != nil {
		return nil, err
	}
	if err := d.Set("tenant_id", integration.GetTenantId()); err != nil {
		return nil, err
	}

	// Perform a read to populate all remaining fields
	diags := resourceAzureIntegrationRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading Azure integration during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}

// setAzureIntegrationComputedFields sets all computed fields from an Azure integration API response
func setAzureIntegrationComputedFields(d *schema.ResourceData, integration *cloudtruthapi.AzureKeyVaultIntegration) error {
	if err := d.Set("name", integration.GetName()); err != nil {
		return err
	}
	if err := d.Set("fqn", integration.GetFqn()); err != nil {
		return err
	}
	if err := d.Set("status", integration.GetStatus()); err != nil {
		return err
	}
	if err := d.Set("status_detail", integration.GetStatusDetail()); err != nil {
		return err
	}
	if statusLastChecked, ok := integration.GetStatusLastCheckedAtOk(); ok && statusLastChecked != nil {
		if err := d.Set("status_last_checked_at", statusLastChecked.String()); err != nil {
			return err
		}
	}
	if err := d.Set("created_at", integration.GetCreatedAt().String()); err != nil {
		return err
	}
	if modifiedAt, ok := integration.GetModifiedAtOk(); ok && modifiedAt != nil {
		if err := d.Set("modified_at", modifiedAt.String()); err != nil {
			return err
		}
	}
	if err := d.Set("type", integration.GetType()); err != nil {
		return err
	}
	if err := d.Set("url", integration.GetUrl()); err != nil {
		return err
	}
	return nil
}
