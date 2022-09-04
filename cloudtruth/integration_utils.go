package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strings"
)

// Utility functions for integrations

func validAWSIntegrationName(v interface{}, attributeName string) (warns []string, errs []error) {
	intName := v.(string)
	nameSegments := strings.Split(intName, "@")
	if len(nameSegments) != 2 {
		errs = append(errs, fmt.Errorf("the %s property %s must use this formate: AWS_ROLE@AWS_ACCOUNT_ID", attributeName, intName))
	}
	return
}

func lookupAWSIntegration(ctx context.Context, intName string, c *cloudTruthClient, d *schema.ResourceData) (*string, error) {
	nameSegments := strings.Split(intName, "@")
	role, awsAccountID := nameSegments[0], nameSegments[1]
	var integrations *cloudtruthapi.PaginatedAwsIntegrationList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		integrations, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsList(ctx).AwsAccountId(awsAccountID).AwsRoleName(role).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("lookupAWSIntegration: error looking up AWS integration %s", intName), r, err)
		}
		if *integrations.Count != 1 {
			return resource.NonRetryableError(fmt.Errorf("lookupAWSIntegration: unexpected number of AWS integrations found for %s", intName))
		}
		return nil
	})
	if retryError != nil {
		return nil, retryError
	}
	awsIntegrationID := integrations.GetResults()[0].GetId()
	return &awsIntegrationID, nil
}

func validAzureIntegrationName(v interface{}, attributeName string) (warns []string, errs []error) {
	intName := v.(string)
	var nameSegments []string
	if strings.Contains(intName, "@") {
		nameSegments = strings.Split(intName, "@")
	}
	if len(nameSegments) != 2 {
		errs = append(errs, fmt.Errorf("the %s property %s must use this format: VAULT_NAME@TENANT_ID", attributeName, intName))
	}
	return
}

func lookupAzureIntegration(ctx context.Context, intName string, c *cloudTruthClient, d *schema.ResourceData) (*string, error) {
	nameSegments := strings.Split(intName, "@")
	vaultName, tenantID := nameSegments[0], nameSegments[1]
	var integrations *cloudtruthapi.PaginatedAzureKeyVaultIntegrationList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		integrations, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAzureKeyVaultList(ctx).VaultName(vaultName).TenantId(tenantID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("lookupAzureIntegration: error looking up Azure integration %s", intName), r, err)
		}
		if *integrations.Count != 1 {
			return resource.NonRetryableError(fmt.Errorf("lookupAzureIntegration: unexpected number of Azure integrations found for %s", intName))
		}
		return nil
	})
	if retryError != nil {
		return nil, retryError
	}
	azureIntegrationID := integrations.GetResults()[0].GetId()
	return &azureIntegrationID, nil
}
