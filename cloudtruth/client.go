package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"net/http"
	"sync"
	"time"
)

const (
	// Allow for CloudTruth CLI API key definitions in addition to the standard TF_VAR_var_name approach
	apiKeyVarName = "CLOUDTRUTH_API_KEY"

	// Optional variables with default values, settable/overridable per Terraform state
	domainVarName   = "CLOUDTRUTH_DOMAIN"
	defaultDomain   = "api.cloudtruth.io"
	protocolVarName = "CLOUDTRUTH_PROTOCOL"
	defaultProtocol = "https"

	// Optional config variables with no defaults - settable/overridable in HCL
	projectVarName     = "CLOUDTRUTH_PROJECT"
	environmentVarName = "CLOUDTRUTH_ENVIRONMENT"

	// Hard code the CloudTruth API major version for now
	apiVersion = "v1"

	// Number of times to attempt reloading project and env caches
	loadCacheRetries = 5
	retrySleep       = 1
	cacheCount       = 7
)

type clientConfig struct {
	APIKey      string
	Project     string
	Environment string
	UserAgent   string
	Domain      string
	Protocol    string
	BaseURL     string
}

func configureClient(ctx context.Context, conf clientConfig) (*cloudTruthClient, diag.Diagnostics) {
	tflog.Debug(ctx, "configureClient")
	tflog.Debug(ctx, fmt.Sprintf("%+v", conf))
	apiConfig := cloudtruthapi.NewConfiguration()
	apiConfig.Host = conf.Domain
	apiConfig.Scheme = conf.Protocol
	apiConfig.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", conf.APIKey))
	apiConfig.AddDefaultHeader("User-Agent", conf.UserAgent)
	conf.BaseURL = fmt.Sprintf("%s://%s/api/%s", conf.Protocol, conf.Domain,
		apiVersion)

	client := cloudTruthClient{
		config:        conf,
		openAPIClient: cloudtruthapi.NewAPIClient(apiConfig),
	}

	loadErrors := make(chan error, cacheCount)
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		loadErrors <- client.loadProjectNameCache(ctx)
		loadErrors <- client.loadProjectIDCache(ctx)
	}()
	go func() {
		defer wg.Done()
		loadErrors <- client.loadEnvNameCache(ctx)
		loadErrors <- client.loadEnvIDCache(ctx)
	}()
	go func() {
		defer wg.Done()
		loadErrors <- client.loadUserCache(ctx)
	}()
	go func() {
		defer wg.Done()
		loadErrors <- client.loadGroupCache(ctx)
	}()
	go func() {
		defer wg.Done()
		loadErrors <- client.loadTypeCache(ctx)
	}()

	go func() {
		wg.Wait()
		close(loadErrors)
	}()
	for err := range loadErrors {
		if err != nil {
			return nil, diag.FromErr(fmt.Errorf("configureClient: %w", err))
		}
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
	users         map[string]cloudtruthapi.User
	groups        map[string]cloudtruthapi.Group
	types         map[string]cloudtruthapi.ParameterType
}

// Utility function
// With an input k -> v map, returns the "reverse" v -> k map
func convertMap(m map[string]string) map[string]string {
	inv := make(map[string]string)
	for k, v := range m {
		inv[v] = k
	}
	return inv
}

// Look up a project identifier first as a name, returning its ID, then as an ID, returning its name
func (c *cloudTruthClient) lookupProject(ctx context.Context, projNameOrID string) (*string, error) {
	tflog.Debug(ctx, fmt.Sprintf("lookupProject: looking up project with name/ID %s", projNameOrID))
	if projNameOrID == "" {
		if c.config.Project != "" {
			projNameOrID = c.config.Project
		} else {
			return nil, errors.New("lookupProject: the CloudTruth project must be specified at the provider or resource level")
		}
	}
	if val, ok := c.projectNames[projNameOrID]; ok {
		tflog.Debug(ctx, fmt.Sprintf("lookupProject: found project by name %s, with id %s", projNameOrID, val))
		return &val, nil
	} else {
		if val, ok := c.projectIDs[projNameOrID]; ok {
			tflog.Debug(ctx, fmt.Sprintf("lookupProject: found project by ID %s with name %s", projNameOrID, val))
			return &val, nil
		}
	}
	return nil, fmt.Errorf("lookupProject: project with name/ID %s not found", projNameOrID)
}

// Look up a project identifier first as a name, then as an ID
func (c *cloudTruthClient) getProjectURL(ctx context.Context, projNameOrID string) (*string, error) {
	tflog.Debug(ctx, fmt.Sprintf("lookupProject: looking up project with name/ID %s", projNameOrID))
	projID, err := c.lookupProject(ctx, projNameOrID)
	if err != nil {
		return nil, err
	}
	projURL := fmt.Sprintf("%s/projects/%s/", c.config.BaseURL, *projID)
	return &projURL, nil
}

// Map of CloudTruth project names -> project IDs
func (c *cloudTruthClient) loadProjectNameCache(ctx context.Context) error {
	if c.projectNames == nil {
		tflog.Debug(ctx, "loadProjectNameCache: fetching project names")
		retryCount := 0
		var apiError error
		// We cannot use the TF Provider SDK's retry functionality because it only works with state change events.
		// Therefore, we employ a simple retry loop instead.
		for retryCount < loadCacheRetries {
			projectList, r, err := c.openAPIClient.ProjectsAPI.ProjectsList(ctx).Execute()
			if (r == nil) || (r.StatusCode >= 400 && r.StatusCode < 500) {
				apiError = err
				break
			}
			if r.StatusCode >= 500 {
				tflog.Debug(ctx, fmt.Sprintf("loadProjectNameCache: %s", err))
				apiError = err
				retryCount++
			} else {
				c.projectNames = make(map[string]string)
				for _, p := range projectList.Results {
					c.projectNames[p.Name] = p.Id
				}
				apiError = nil
				break
			}
			time.Sleep(retrySleep * time.Second)
		}
		if apiError != nil {
			return fmt.Errorf("loadProjectNameCache: %w", apiError)
		}
	}
	return nil
}

// Map of CloudTruth project IDs -> project names
func (c *cloudTruthClient) loadProjectIDCache(ctx context.Context) error {
	if c.projectIDs == nil {
		tflog.Debug(ctx, "loadProjectIDCache: fetching project IDs")
		err := c.loadProjectNameCache(ctx)
		if err != nil {
			return fmt.Errorf("loadProjectIDCache: %w", err)
		}
		c.projectIDs = convertMap(c.projectNames)
	}
	return nil
}

func (c *cloudTruthClient) addNewProjectToCaches(projectName, projectID string) {
	c.projectNames[projectName] = projectID
	c.projectIDs[projectID] = projectName
}

func (c *cloudTruthClient) removeProjectFromCaches(projectName, projectID string) {
	delete(c.projectNames, projectName)
	delete(c.projectIDs, projectID)
}

// Look up an environment identifier first as a name and return its ID, then if not found as an ID and return its name
func (c *cloudTruthClient) lookupEnvironment(ctx context.Context, envNameOrID string) (*string, error) {
	tflog.Debug(ctx, fmt.Sprintf("lookupEnvironment: looking up environment with name/ID %s", envNameOrID))
	if val, ok := c.envNames[envNameOrID]; ok {
		tflog.Debug(ctx, fmt.Sprintf("found environment by name %s, with id %s", envNameOrID, val))
		return &val, nil
	} else {
		if val, ok := c.envIDs[envNameOrID]; ok {
			tflog.Debug(ctx, fmt.Sprintf("found environment by ID %s with name %s", envNameOrID, val))
			return &val, nil
		}
	}
	return nil, fmt.Errorf("lookupEnvironment: environment with name/ID %s not found", envNameOrID)
}

// Map of CloudTruth environment names -> environment IDs
func (c *cloudTruthClient) loadEnvNameCache(ctx context.Context) error {
	if c.envNames == nil {
		tflog.Debug(ctx, "loadEnvNameCache: fetching environment names")
		retryCount := 0
		var apiError error
		// We cannot use the TF Provider SDK's retry functionality because it only works with state change events.
		// Therefore, we employ a simple retry loop instead.
		for retryCount < loadCacheRetries {
			envList, r, err := c.openAPIClient.EnvironmentsAPI.EnvironmentsList(ctx).Execute()
			if (r == nil) || (r.StatusCode >= 400 && r.StatusCode < 500) {
				apiError = err
				break
			}
			if r.StatusCode >= 500 {
				tflog.Debug(ctx, fmt.Sprintf("loadEnvNameCache: %s", err))
				apiError = err
				retryCount++
			} else if envList != nil {
				c.envNames = make(map[string]string)
				for _, p := range envList.Results {
					c.envNames[p.Name] = p.Id
				}
				apiError = nil
				break
			}
			time.Sleep(retrySleep * time.Second)
		}
		if apiError != nil {
			return fmt.Errorf("loadEnvNameCache: %w", apiError)
		}
	}
	return nil
}

func (c *cloudTruthClient) addNewEnvToCaches(envName, envID string) {
	c.envNames[envName] = envID
	c.envIDs[envID] = envName
}

func (c *cloudTruthClient) removeEnvFromCaches(envName, envID string) {
	delete(c.envNames, envName)
	delete(c.envIDs, envID)
}

// Map of CloudTruth environment IDs -> project names
func (c *cloudTruthClient) loadEnvIDCache(ctx context.Context) error {
	if c.envIDs == nil {
		tflog.Debug(ctx, "loadEnvIDCache: fetching environment IDs")
		err := c.loadEnvNameCache(ctx)
		if err != nil {
			return fmt.Errorf("loadEnvIDCache: %w", err)
		}
		c.envIDs = convertMap(c.envNames)
	}
	return nil
}

// Look up an environment identifier first as a name, then as an ID
func (c *cloudTruthClient) lookupEnvProj(ctx context.Context, envNameOrID, projNameOrID string) (*string, *string, error) {
	tflog.Debug(ctx, fmt.Sprintf("lookupEnvProj: looking up environment %s and project %s", envNameOrID, projNameOrID))
	envID, err := c.lookupEnvironment(ctx, envNameOrID)
	if err != nil {
		return nil, nil, err
	}
	projID, err := c.lookupProject(ctx, projNameOrID)
	if err != nil {
		return nil, nil, err
	}
	return envID, projID, nil
}

// Map of ALL usernames -> user structs
// We also add email -> user struct entries for easy lookups by name or email
func (c *cloudTruthClient) loadUserCache(ctx context.Context) error {
	if c.users == nil {
		tflog.Debug(ctx, "loadUserCache: fetching user accounts")
		c.users = make(map[string]cloudtruthapi.User)
		var pageNum int32 = 0
		for {
			userListRequest := c.openAPIClient.UsersAPI.UsersList(ctx)
			var userList *cloudtruthapi.PaginatedUserList
			if pageNum > 0 {
				userListRequest = userListRequest.Page(pageNum)
			}
			pageNum++
			retryCount := 0
			var apiError, err error
			var r *http.Response
			for retryCount < loadCacheRetries {
				userList, r, err = userListRequest.Execute()
				if (r == nil) || (r.StatusCode >= 400 && r.StatusCode < 500) {
					apiError = err
					break
				}
				if r.StatusCode >= 500 {
					tflog.Debug(ctx, fmt.Sprintf("loadUserCache: %s", err))
					retryCount++
				} else {
					for _, u := range userList.Results {
						name := u.GetName()
						if _, ok := c.users[name]; ok {
							tflog.Error(ctx,
								fmt.Sprintf("loadUserCache: duplicate users found with the name '%s', specify user emails instead to disambiguate user data source", name))
						}
						c.users[name] = u
						if u.GetEmail() != "" {
							c.users[u.GetEmail()] = u // An email key pointing to the same User (pointer)
						}
					}
					err = nil
					break
				}
				time.Sleep(retrySleep * time.Second)
			}
			if apiError != nil {
				return fmt.Errorf("loadUserCache: %w", apiError)
			}
			if userList.GetNext() == "" {
				break
			}
		}
	}
	return nil
}

// Look up a user by email or name
func (c *cloudTruthClient) lookupUser(ctx context.Context, userNameOrEmail string) (*cloudtruthapi.User, error) {
	tflog.Debug(ctx, fmt.Sprintf("lookupUser: looking up the user account for %s", userNameOrEmail))
	if user, ok := c.users[userNameOrEmail]; ok {
		return &user, nil
	}
	return nil, nil
}

func (c *cloudTruthClient) loadGroupCache(ctx context.Context) error {
	if c.groups == nil {
		tflog.Debug(ctx, "loadGroupCache: fetching groups")
		c.groups = make(map[string]cloudtruthapi.Group)
		var pageNum int32 = 0
		for {
			groupListRequest := c.openAPIClient.GroupsAPI.GroupsList(ctx)
			var groupList *cloudtruthapi.PaginatedGroupList
			if pageNum > 0 {
				groupListRequest = groupListRequest.Page(pageNum)
			}
			pageNum++
			retryCount := 0
			var apiError, err error
			var r *http.Response
			for retryCount < loadCacheRetries {
				groupList, r, err = groupListRequest.Execute()
				if (r == nil) || (r.StatusCode >= 400 && r.StatusCode < 500) {
					apiError = err
					break
				}
				if r.StatusCode >= 500 {
					tflog.Debug(ctx, fmt.Sprintf("loadGroupCache: %s", err))
					apiError = err
					retryCount++
				} else if groupList != nil { // No groups exist
					for _, g := range groupList.Results {
						c.groups[g.GetName()] = g
					}
					apiError = nil
					break
				}
				time.Sleep(retrySleep * time.Second)
			}
			if apiError != nil {
				return fmt.Errorf("loadGroupCache: %w", apiError)
			}
			if groupList.GetNext() == "" {
				break
			}
		}
	}
	return nil
}

func (c *cloudTruthClient) lookupGroup(ctx context.Context, groupName string) (*cloudtruthapi.Group, error) {
	tflog.Debug(ctx, fmt.Sprintf("lookupUser: looking up the group %s", groupName))
	if group, ok := c.groups[groupName]; ok {
		return &group, nil
	}
	return nil, nil
}

func (c *cloudTruthClient) loadTypeCache(ctx context.Context) error {
	if c.types == nil {
		tflog.Debug(ctx, "loadTypeCache: fetching custom and builtin types")
		c.types = make(map[string]cloudtruthapi.ParameterType)
		var pageNum int32 = 0 // Pagination is likely overkill here
		for {
			typeListRequest := c.openAPIClient.TypesAPI.TypesList(ctx).PageSize(pageNum)
			var typeList *cloudtruthapi.PaginatedParameterTypeList
			if pageNum > 0 {
				typeListRequest = typeListRequest.Page(pageNum)
			}
			pageNum++
			retryCount := 0
			var apiError, err error
			var r *http.Response
			for retryCount < loadCacheRetries {
				typeList, r, err = typeListRequest.Execute()
				if (r == nil) || (r.StatusCode >= 400 && r.StatusCode < 500) {
					apiError = err
					break
				}
				if r.StatusCode >= 500 {
					tflog.Debug(ctx, fmt.Sprintf("loadTypeCache: %s", err))
					apiError = err
					retryCount++
				} else if typeList != nil {
					for _, t := range typeList.Results {
						c.types[t.GetName()] = t
					}
					apiError = nil
					break
				}
				time.Sleep(retrySleep * time.Second)
			}
			if apiError != nil {
				return fmt.Errorf("loadTypeCache: %w", apiError)
			}
			if typeList.GetNext() == "" {
				break
			}
		}
	}
	return nil
}

// Look up a type (custom or builtin) by name
func (c *cloudTruthClient) lookupType(ctx context.Context, typeName string) *cloudtruthapi.ParameterType {
	tflog.Debug(ctx, fmt.Sprintf("lookupType: looking up the parameter type %s", typeName))
	if paramType, ok := c.types[typeName]; ok {
		return &paramType
	}
	return nil
}

/*
Helper function for making API errors more useful
API/server side errors are generally retryable
client side errors are generally not retryable, except for 409 responses
which can indicate that we have a race condition and should retry the operation
In non-retryable cases, weextract the underlying API error so that we can present
it to the end user
*/
func handleAPIError(msg string, r *http.Response, err error) *retry.RetryError {
	outErr := fmt.Errorf(msg, err)
	if r == nil { // Can be nil when the context has been cancelled
		return retry.NonRetryableError(outErr)
	} else if r.StatusCode >= 409 { // Conflict, we should try again
		return retry.RetryableError(outErr)
	} else if r.StatusCode >= 500 { // server side error
		return retry.RetryableError(outErr)
	} else { // All other 4xx errors
		outErr = fmt.Errorf("%s - %d client error: %s", msg, r.StatusCode, r.Body)
		return retry.NonRetryableError(outErr)
	}
}
