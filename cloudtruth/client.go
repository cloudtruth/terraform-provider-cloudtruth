package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type cloudTruthClient struct {
	config        clientConfig
	openAPIClient *cloudtruthapi.APIClient
	envNames      map[string]string
	envIDs        map[string]string
	projectNames  map[string]string
	projectIDs    map[string]string
}

// todo:
// set the default env and project
// read Parameter -> data source, determine what properties to expose in HCL
// read Parameters based on search/filter, but scoped to one env and project
// project CRUD
// env CRUD
// param CRUD
// template CRUD
// import support for all resources

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
func (c *cloudTruthClient) lookupProject(ctx context.Context, projNameOrID string) *string {
	projectNames := c.projectNameCache(ctx)
	if val, ok := projectNames[projNameOrID]; ok {
		tflog.Debug(ctx, fmt.Sprintf("Found project by name %s, with id %s", projNameOrID, val))
		return &val
	} else {
		projectIDs := c.projectIDCache(ctx)
		if val, ok := projectIDs[projNameOrID]; ok {
			tflog.Debug(ctx, fmt.Sprintf("Found project by ID %s with name %s", projNameOrID, val))
			return &projNameOrID
		}
	}
	return nil
}

// Map of CloudTruth project names -> project IDs
func (c *cloudTruthClient) projectNameCache(ctx context.Context) map[string]string {
	if c.projectNames == nil {
		tflog.Debug(ctx, "Fetching project names")
		resp, r, err := c.openAPIClient.ProjectsApi.ProjectsList(context.Background()).Execute()
		if err != nil {
			// todo: diag logging
			fmt.Printf("%+v", r)
		}
		c.projectNames = make(map[string]string)
		for _, p := range resp.Results {
			c.projectNames[p.Name] = p.Id
		}
	} else {
		tflog.Debug(ctx, "Returning cached project names")
	}
	return c.projectNames
}

// Map of CloudTruth project IDs -> project names
func (c *cloudTruthClient) projectIDCache(ctx context.Context) map[string]string {
	if c.projectIDs == nil {
		tflog.Debug(ctx, "Fetching project IDs")
		c.projectIDs = convertMap(c.projectNameCache(ctx))
	} else {
		tflog.Debug(ctx, "Returning cached project IDs")
	}
	return c.projectIDs
}

// Map of CloudTruth environment names -> environment IDs
func (c *cloudTruthClient) envNameCache(ctx context.Context) map[string]string {
	if c.envNames == nil {
		tflog.Debug(ctx, "Fetching environment names")
		resp, r, err := c.openAPIClient.EnvironmentsApi.EnvironmentsList(context.Background()).Execute()
		if err != nil {
			// todo: diag logging
			fmt.Printf("%+v", r)
		}
		c.envNames = make(map[string]string)
		for _, p := range resp.Results {
			c.envNames[p.Name] = p.Id
		}
	}
	return c.envNames
}

// Map of CloudTruth environment IDs -> project names
func (c *cloudTruthClient) envIDCache(ctx context.Context) map[string]string {
	if c.envIDs == nil {
		tflog.Debug(ctx, "Fetching environment IDs")
		c.envIDs = convertMap(c.envNameCache(ctx))
	} else {
		tflog.Debug(ctx, "Returning cached environment IDs")
	}
	return c.envIDs
}
