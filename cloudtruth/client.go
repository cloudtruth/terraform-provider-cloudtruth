package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	//"sync"
)

// todo:
// set the default env and project
// read Parameter -> data source, determine what properties to expose in HCL
// read Parameters based on search/filter, but scoped to one env and project
// project CRUD
// env CRUD
// param CRUD
// template CRUD
// import support for all resources

/* sync.Once can only be used one time per function
var (
	projNameCacheOnce sync.Once
	projIDCacheOnce   sync.Once
	envNameCacheOnce  sync.Once
	envIDCacheOnce    sync.Once
        )*/

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

// Look up a project ID first by name, then by the ID itself
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

// Map of CloudTruth environment names -> environment IDs
func (c *cloudTruthClient) loadEnvNameCache(ctx context.Context) error {
	if c.envNames == nil {
		tflog.Debug(ctx, "Fetching environment names")
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
