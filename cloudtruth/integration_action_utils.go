package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strings"
)

// Utility functions for integrations and actions

func isValidAWSIntegrationName(v interface{}, attributeName string) (warns []string, errs []error) {
	intName := v.(string)
	nameSegments := strings.Split(intName, "@")
	if len(nameSegments) != 2 {
		errs = append(errs, fmt.Errorf("the %s property %s must use this format: AWS_ROLE@AWS_ACCOUNT_ID", attributeName, intName))
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

func isValidAzureIntegrationName(v interface{}, attributeName string) (warns []string, errs []error) {
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

func isValidEnvTagSet(v interface{}, attributeName string) (warns []string, errs []error) {
	envTags := v.([]string)
	seenTags := make(map[string]bool)
	for _, envTag := range envTags {
		nameSegments := strings.Split(envTag, "@")
		if len(nameSegments) != 2 {
			errs = append(errs, fmt.Errorf("the %s env tag must use this format: ENVIRONMENT_NAME:TAG_NAME", envTag))
		}
		if _, ok := seenTags[envTag]; ok {
			errs = append(errs, fmt.Errorf("duplicate env tag %s, duplicate tags are not allowed", envTag))
		}
		seenTags[envTag] = true
	}
	return
}

func getProjectURLs(ctx context.Context, c *cloudTruthClient, rawProjects []interface{}) ([]string, error) {
	projects := make([]string, len(rawProjects))
	for i, v := range projects {
		projID, err := c.lookupProject(ctx, fmt.Sprint(v))
		if err != nil {
			return nil, err
		}
		projURL := fmt.Sprintf("https://api.cloudtruth.io/api/v1/projects/%s/", *projID)
		projects[i] = projURL
	}
	return projects, nil
}

func getEnvTags(ctx context.Context, d *schema.ResourceData, c *cloudTruthClient, rawTags []interface{}) ([]string, error) {
	tags := make([]string, len(rawTags))
	var err error
	for i, v := range rawTags {
		tags[i], err = lookupEnvTag(ctx, d, c, fmt.Sprint(v))
		if err != nil {
			return nil, err
		}
	}
	return tags, nil
}

// Helper function for push actions
// todo: consider implementing a tag cache
func lookupEnvTag(ctx context.Context, d *schema.ResourceData, c *cloudTruthClient, envTag string) (string, error) {
	tflog.Debug(ctx, "lookupEnvTag")
	if !strings.Contains(envTag, ":") {
		return "", fmt.Errorf("the tag %s is invalid, it must be specified using the form 'environment_name:tag_name'", envTag)
	}
	envTagSegments := strings.Split(envTag, ":")
	env, tagName := envTagSegments[0], envTagSegments[1]
	envID, err := c.lookupEnvironment(ctx, env)
	if err != nil {
		return "", err
	}

	var resp *cloudtruthapi.PaginatedTagList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		resp, r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsTagsList(ctx, *envID).Name(tagName).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("lookupEnvTag: error looking up tag %s", tagName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return "", retryError
	}

	// There should only be one tag with the specified per environment
	results := resp.GetResults()
	tag := results[0]
	tagID := tag.GetId()
	return tagID, nil
}
