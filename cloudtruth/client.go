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

// Map of CloudTruth project names -> project IDs
func (c *cloudTruthClient) projectNameCache() map[string]string {
	if c.projectIDCache == nil {
		tflog.Debug(ctx, "Fetching project names")
		resp, r, err := c.openAPIClient.ProjectsApi.ProjectsList(context.Background()).Execute()
		if err != nil {
			// todo: diag logging
			fmt.Printf("%+v", r)
		}
		c.projectIDCache = make(map[string]string)
		for _, p := range *resp.Results {
			c.projectNameCache[p.Name] = p.Id
		}
	} else {
		tflog.Debug(ctx, "Returning cached project names")
	}
	return c.projectNameCache
}

// Map of CloudTruth project IDs -> project names
func (c *cloudTruthClient) projectIDCache() map[string]string {
	if c.projectIDCache == nil {
		tflog.Debug(ctx, "Fetching project IDs")
		c.projectIDCache = convertMap(c.projectNameCache())
	} else {
		tflog.Debug(ctx, "Returning cached project IDs")
	}
	return c.projectIDCache
}

func convertMap(m map[int]string) map[string]string {
	inv := make(map[string]string)
	for k, v := range m {
		inv[v] = k
	}
	return inv
}

// Map of CloudTruth environment names -> environment IDs
func (c *cloudTruthClient) envNameCache() map[string]string {
	if c.envIDCache == nil {
		tflog.Debug(ctx, "Fetching projects")
		resp, r, err := c.openAPIClient.EnvironmentsApi.EnvironmentsList(context.Background()).Execute()
		if err != nil {
			// todo: diag logging
			fmt.Printf("%+v", r)
		}
		c.envIDCache = make(map[string]string)
		for _, p := range *resp.Results {
			c.envIDCache[p.Name] = p.Id
		}
	}
	return c.envIDCache
}

func (c *cloudTruthClient) Get(url string) (resp *http.Response, err error) {
	sResp, _, sErr := c.openAPIClient.EnvironmentsApi.EnvironmentsList(context.Background()).Execute()
	if sErr != nil {
		fmt.Print("todo")
	}
	for _, p := range sResp.Results {
		fmt.Println(p.Name)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err

	}
	// The API returns a 400 when the project isn't known
	// fmt.Println(formatRequest(req))
	resp, respErr := c.Do(req)
	if err != nil {
		return nil, respErr
	}

	return resp, respErr
}
