package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	// Allow for CloudTruth CLI API key definitions in addition to TF_VAR_var_name approach
	// no default
	apiKeyVarName = "CLOUDTRUTH_API_KEY"

	// Optional variables with default values, settable/overridable per Terraform state
	domainVarName   = "CLOUDTRUTH_DOMAIN"
	defaultDomain   = "api.cloudtruth.io"
	protocolVarName = "CLOUDTRUTH_PROTOCOL"
	defaultProtocol = "https"

	// Optional config variables with no defaults - settable/overridable per Terraform state
	projectVarName     = "CLOUDTRUTH_PROJECT"
	environmentVarName = "CLOUDTRUTH_ENVIRONMENT"
)

type clientConfig struct {
	APIKey      string
	Project     string
	Environment string
	UserAgent   string
	Domain      string
	Protocol    string
}

func configureClient(ctx context.Context, conf clientConfig) (*cloudTruthClient, diag.Diagnostics) {
	tflog.Debug(ctx, "configureClient")
	tflog.Debug(ctx, fmt.Sprintf("%+v", conf))
	apiConfig := cloudtruthapi.NewConfiguration()
	apiConfig.Host = conf.Domain
	apiConfig.Scheme = conf.Protocol
	apiConfig.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", conf.APIKey))
	apiConfig.AddDefaultHeader("UserAgent", conf.UserAgent)
	client := cloudTruthClient{
		config:        conf,
		openAPIClient: cloudtruthapi.NewAPIClient(apiConfig),
	}

	// populate & load caches
	err := client.loadProjectNameCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	err = client.loadProjectIDCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	err = client.loadEnvNameCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	err = client.loadEnvIDCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return &client, nil
}

type cloudTruthClient struct {
	config        clientConfig
	openAPIClient *cloudtruthapi.APIClient
	envNames      map[string]string
	envIDs        map[string]string
	projectNames  map[string]string
	projectIDs    map[string]string
}

// For convenient + fast lookups by ID and name, use this utility function
// to create a reverse val -> key map
func convertMap(m map[string]string) map[string]string {
	inv := make(map[string]string)
	for k, v := range m {
		inv[v] = k
	}
	return inv
}

// Look up a project identifier first as a name, then as an ID
func (c *cloudTruthClient) lookupProject(ctx context.Context, projNameOrID string) (*string, error) {
	tflog.Debug(ctx, fmt.Sprintf("Looking up project with name/ID %s", projNameOrID))
	if val, ok := c.projectNames[projNameOrID]; ok {
		tflog.Debug(ctx, fmt.Sprintf("Found project by name %s, with id %s", projNameOrID, val))
		return &val, nil
	} else {
		if val, ok := c.projectIDs[projNameOrID]; ok {
			tflog.Debug(ctx, fmt.Sprintf("Found project by ID %s with name %s", projNameOrID, val))
			return &projNameOrID, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Project with name/ID %s not found", projNameOrID))
}

// Map of CloudTruth project names -> project IDs
func (c *cloudTruthClient) loadProjectNameCache(ctx context.Context) error {
	if c.projectNames == nil {
		tflog.Debug(ctx, "Fetching project names")
		resp, _, err := c.openAPIClient.ProjectsApi.ProjectsList(context.Background()).Execute()
		if err != nil {
			return err
		}
		c.projectNames = make(map[string]string)
		for _, p := range resp.Results {
			c.projectNames[p.Name] = p.Id
		}
	}
	return nil
}

// Map of CloudTruth project IDs -> project names
func (c *cloudTruthClient) loadProjectIDCache(ctx context.Context) error {
	if c.projectIDs == nil {
		tflog.Debug(ctx, "Fetching project IDs")
		err := c.loadProjectNameCache(ctx)
		if err != nil {
			return err
		}
		c.projectIDs = convertMap(c.projectNames)
	}
	return nil
}

// Look up an environment identifier first as a name, then as an ID
func (c *cloudTruthClient) lookupEnvironment(ctx context.Context, envNameOrID string) (*string, error) {
	tflog.Debug(ctx, fmt.Sprintf("looking up environment with name/ID %s", envNameOrID))
	if val, ok := c.envNames[envNameOrID]; ok {
		tflog.Debug(ctx, fmt.Sprintf("found environment by name %s, with id %s", envNameOrID, val))
		return &val, nil
	} else {
		if val, ok := c.envIDs[envNameOrID]; ok {
			tflog.Debug(ctx, fmt.Sprintf("found environment by ID %s with name %s", envNameOrID, val))
			return &envNameOrID, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("environment with name/ID %s not found", envNameOrID))
}

func (c *cloudTruthClient) getEnvironmentName(ctx context.Context, envID string) (*string, error) {
	tflog.Debug(ctx, fmt.Sprintf("fetching the name of the environment with ID %s", envID))
	if val, ok := c.envIDs[envID]; ok {
		tflog.Debug(ctx, fmt.Sprintf("found environment with name %s, using id %s", envID, val))
		return &val, nil
	}
	return nil, errors.New(fmt.Sprintf("environment with ID %s not found", envID))
}

// Map of CloudTruth environment names -> environment IDs
func (c *cloudTruthClient) loadEnvNameCache(ctx context.Context) error {
	if c.envNames == nil {
		tflog.Debug(ctx, "fetching environment names")
		resp, _, err := c.openAPIClient.EnvironmentsApi.EnvironmentsList(context.Background()).Execute()
		if err != nil {
			return err
		}
		c.envNames = make(map[string]string)
		for _, p := range resp.Results {
			c.envNames[p.Name] = p.Id
		}
	}
	return nil
}

// Map of CloudTruth environment IDs -> project names
func (c *cloudTruthClient) loadEnvIDCache(ctx context.Context) error {
	if c.envIDs == nil {
		tflog.Debug(ctx, "Fetching environment IDs")
		err := c.loadEnvNameCache(ctx)
		if err != nil {
			return err
		}
		c.envIDs = convertMap(c.envNames)
	}
	return nil
}
