# \IntegrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsAwsCreate**](IntegrationsApi.md#IntegrationsAwsCreate) | **Post** /api/v1/integrations/aws/ | Establishes an AWS Integration.
[**IntegrationsAwsDestroy**](IntegrationsApi.md#IntegrationsAwsDestroy) | **Delete** /api/v1/integrations/aws/{id}/ | Delete an AWS integration.
[**IntegrationsAwsList**](IntegrationsApi.md#IntegrationsAwsList) | **Get** /api/v1/integrations/aws/ | 
[**IntegrationsAwsPartialUpdate**](IntegrationsApi.md#IntegrationsAwsPartialUpdate) | **Patch** /api/v1/integrations/aws/{id}/ | 
[**IntegrationsAwsPullsCreate**](IntegrationsApi.md#IntegrationsAwsPullsCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pulls/ | 
[**IntegrationsAwsPullsDestroy**](IntegrationsApi.md#IntegrationsAwsPullsDestroy) | **Delete** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPullsList**](IntegrationsApi.md#IntegrationsAwsPullsList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/ | 
[**IntegrationsAwsPullsPartialUpdate**](IntegrationsApi.md#IntegrationsAwsPullsPartialUpdate) | **Patch** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPullsRetrieve**](IntegrationsApi.md#IntegrationsAwsPullsRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPullsSyncCreate**](IntegrationsApi.md#IntegrationsAwsPullsSyncCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/sync/ | 
[**IntegrationsAwsPullsTasksList**](IntegrationsApi.md#IntegrationsAwsPullsTasksList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/ | 
[**IntegrationsAwsPullsTasksRetrieve**](IntegrationsApi.md#IntegrationsAwsPullsTasksRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/{id}/ | 
[**IntegrationsAwsPullsTasksStepsList**](IntegrationsApi.md#IntegrationsAwsPullsTasksStepsList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/{awspulltask_pk}/steps/ | 
[**IntegrationsAwsPullsTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsAwsPullsTasksStepsRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{awspull_pk}/tasks/{awspulltask_pk}/steps/{id}/ | 
[**IntegrationsAwsPullsUpdate**](IntegrationsApi.md#IntegrationsAwsPullsUpdate) | **Put** /api/v1/integrations/aws/{awsintegration_pk}/pulls/{id}/ | 
[**IntegrationsAwsPushesCreate**](IntegrationsApi.md#IntegrationsAwsPushesCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pushes/ | 
[**IntegrationsAwsPushesDestroy**](IntegrationsApi.md#IntegrationsAwsPushesDestroy) | **Delete** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsPushesList**](IntegrationsApi.md#IntegrationsAwsPushesList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/ | 
[**IntegrationsAwsPushesPartialUpdate**](IntegrationsApi.md#IntegrationsAwsPushesPartialUpdate) | **Patch** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsPushesRetrieve**](IntegrationsApi.md#IntegrationsAwsPushesRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsPushesSyncCreate**](IntegrationsApi.md#IntegrationsAwsPushesSyncCreate) | **Post** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/sync/ | 
[**IntegrationsAwsPushesTasksList**](IntegrationsApi.md#IntegrationsAwsPushesTasksList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/ | 
[**IntegrationsAwsPushesTasksRetrieve**](IntegrationsApi.md#IntegrationsAwsPushesTasksRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/{id}/ | 
[**IntegrationsAwsPushesTasksStepsList**](IntegrationsApi.md#IntegrationsAwsPushesTasksStepsList) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/{awspushtask_pk}/steps/ | 
[**IntegrationsAwsPushesTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsAwsPushesTasksStepsRetrieve) | **Get** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{awspush_pk}/tasks/{awspushtask_pk}/steps/{id}/ | 
[**IntegrationsAwsPushesUpdate**](IntegrationsApi.md#IntegrationsAwsPushesUpdate) | **Put** /api/v1/integrations/aws/{awsintegration_pk}/pushes/{id}/ | 
[**IntegrationsAwsRetrieve**](IntegrationsApi.md#IntegrationsAwsRetrieve) | **Get** /api/v1/integrations/aws/{id}/ | Get details of an AWS Integration.
[**IntegrationsAwsScanCreate**](IntegrationsApi.md#IntegrationsAwsScanCreate) | **Post** /api/v1/integrations/aws/{id}/scan/ | Evaluate a potential pull pattern and understand what it will match.
[**IntegrationsAwsUpdate**](IntegrationsApi.md#IntegrationsAwsUpdate) | **Put** /api/v1/integrations/aws/{id}/ | 
[**IntegrationsAzureKeyVaultCreate**](IntegrationsApi.md#IntegrationsAzureKeyVaultCreate) | **Post** /api/v1/integrations/azure/key_vault/ | Establishes an Azure Key Vault Integration.
[**IntegrationsAzureKeyVaultDestroy**](IntegrationsApi.md#IntegrationsAzureKeyVaultDestroy) | **Delete** /api/v1/integrations/azure/key_vault/{id}/ | Delete an AWS integration.
[**IntegrationsAzureKeyVaultList**](IntegrationsApi.md#IntegrationsAzureKeyVaultList) | **Get** /api/v1/integrations/azure/key_vault/ | 
[**IntegrationsAzureKeyVaultPartialUpdate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPartialUpdate) | **Patch** /api/v1/integrations/azure/key_vault/{id}/ | 
[**IntegrationsAzureKeyVaultPullsCreate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsCreate) | **Post** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/ | 
[**IntegrationsAzureKeyVaultPullsDestroy**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsDestroy) | **Delete** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{id}/ | 
[**IntegrationsAzureKeyVaultPullsList**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsList) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/ | 
[**IntegrationsAzureKeyVaultPullsPartialUpdate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsPartialUpdate) | **Patch** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{id}/ | 
[**IntegrationsAzureKeyVaultPullsRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{id}/ | 
[**IntegrationsAzureKeyVaultPullsSyncCreate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsSyncCreate) | **Post** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{id}/sync/ | 
[**IntegrationsAzureKeyVaultPullsTasksList**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsTasksList) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{akvpull_pk}/tasks/ | 
[**IntegrationsAzureKeyVaultPullsTasksRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsTasksRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{akvpull_pk}/tasks/{id}/ | 
[**IntegrationsAzureKeyVaultPullsTasksStepsList**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsTasksStepsList) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{akvpull_pk}/tasks/{akvpulltask_pk}/steps/ | 
[**IntegrationsAzureKeyVaultPullsTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsTasksStepsRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{akvpull_pk}/tasks/{akvpulltask_pk}/steps/{id}/ | 
[**IntegrationsAzureKeyVaultPullsUpdate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPullsUpdate) | **Put** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pulls/{id}/ | 
[**IntegrationsAzureKeyVaultPushesCreate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesCreate) | **Post** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/ | 
[**IntegrationsAzureKeyVaultPushesDestroy**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesDestroy) | **Delete** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{id}/ | 
[**IntegrationsAzureKeyVaultPushesList**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesList) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/ | 
[**IntegrationsAzureKeyVaultPushesPartialUpdate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesPartialUpdate) | **Patch** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{id}/ | 
[**IntegrationsAzureKeyVaultPushesRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{id}/ | 
[**IntegrationsAzureKeyVaultPushesSyncCreate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesSyncCreate) | **Post** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{id}/sync/ | 
[**IntegrationsAzureKeyVaultPushesTasksList**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesTasksList) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{akvpush_pk}/tasks/ | 
[**IntegrationsAzureKeyVaultPushesTasksRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesTasksRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{akvpush_pk}/tasks/{id}/ | 
[**IntegrationsAzureKeyVaultPushesTasksStepsList**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesTasksStepsList) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{akvpush_pk}/tasks/{akvpushtask_pk}/steps/ | 
[**IntegrationsAzureKeyVaultPushesTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesTasksStepsRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{akvpush_pk}/tasks/{akvpushtask_pk}/steps/{id}/ | 
[**IntegrationsAzureKeyVaultPushesUpdate**](IntegrationsApi.md#IntegrationsAzureKeyVaultPushesUpdate) | **Put** /api/v1/integrations/azure/key_vault/{akvintegration_pk}/pushes/{id}/ | 
[**IntegrationsAzureKeyVaultRetrieve**](IntegrationsApi.md#IntegrationsAzureKeyVaultRetrieve) | **Get** /api/v1/integrations/azure/key_vault/{id}/ | Get details of an Azure Key Vault Integration.
[**IntegrationsAzureKeyVaultScanCreate**](IntegrationsApi.md#IntegrationsAzureKeyVaultScanCreate) | **Post** /api/v1/integrations/azure/key_vault/{id}/scan/ | Evaluate a potential pull pattern and understand what it will match.
[**IntegrationsAzureKeyVaultUpdate**](IntegrationsApi.md#IntegrationsAzureKeyVaultUpdate) | **Put** /api/v1/integrations/azure/key_vault/{id}/ | 
[**IntegrationsExploreList**](IntegrationsApi.md#IntegrationsExploreList) | **Get** /api/v1/integrations/explore/ | Retrieve third-party integration data for the specified FQN.
[**IntegrationsGithubCreate**](IntegrationsApi.md#IntegrationsGithubCreate) | **Post** /api/v1/integrations/github/ | Establishes a GitHub Integration.
[**IntegrationsGithubDestroy**](IntegrationsApi.md#IntegrationsGithubDestroy) | **Delete** /api/v1/integrations/github/{id}/ | Delete a GitHub integration.
[**IntegrationsGithubList**](IntegrationsApi.md#IntegrationsGithubList) | **Get** /api/v1/integrations/github/ | 
[**IntegrationsGithubPullsList**](IntegrationsApi.md#IntegrationsGithubPullsList) | **Get** /api/v1/integrations/github/{githubintegration_pk}/pulls/ | 
[**IntegrationsGithubPullsPartialUpdate**](IntegrationsApi.md#IntegrationsGithubPullsPartialUpdate) | **Patch** /api/v1/integrations/github/{githubintegration_pk}/pulls/{id}/ | 
[**IntegrationsGithubPullsRetrieve**](IntegrationsApi.md#IntegrationsGithubPullsRetrieve) | **Get** /api/v1/integrations/github/{githubintegration_pk}/pulls/{id}/ | 
[**IntegrationsGithubPullsSyncCreate**](IntegrationsApi.md#IntegrationsGithubPullsSyncCreate) | **Post** /api/v1/integrations/github/{githubintegration_pk}/pulls/{id}/sync/ | 
[**IntegrationsGithubPullsTasksList**](IntegrationsApi.md#IntegrationsGithubPullsTasksList) | **Get** /api/v1/integrations/github/{githubintegration_pk}/pulls/{githubpull_pk}/tasks/ | 
[**IntegrationsGithubPullsTasksRetrieve**](IntegrationsApi.md#IntegrationsGithubPullsTasksRetrieve) | **Get** /api/v1/integrations/github/{githubintegration_pk}/pulls/{githubpull_pk}/tasks/{id}/ | 
[**IntegrationsGithubPullsTasksStepsList**](IntegrationsApi.md#IntegrationsGithubPullsTasksStepsList) | **Get** /api/v1/integrations/github/{githubintegration_pk}/pulls/{githubpull_pk}/tasks/{githubpulltask_pk}/steps/ | 
[**IntegrationsGithubPullsTasksStepsRetrieve**](IntegrationsApi.md#IntegrationsGithubPullsTasksStepsRetrieve) | **Get** /api/v1/integrations/github/{githubintegration_pk}/pulls/{githubpull_pk}/tasks/{githubpulltask_pk}/steps/{id}/ | 
[**IntegrationsGithubPullsUpdate**](IntegrationsApi.md#IntegrationsGithubPullsUpdate) | **Put** /api/v1/integrations/github/{githubintegration_pk}/pulls/{id}/ | 
[**IntegrationsGithubRetrieve**](IntegrationsApi.md#IntegrationsGithubRetrieve) | **Get** /api/v1/integrations/github/{id}/ | Get details of a GitHub Integration.



## IntegrationsAwsCreate

> AwsIntegration IntegrationsAwsCreate(ctx).AwsIntegrationCreate(awsIntegrationCreate).Execute()

Establishes an AWS Integration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsIntegrationCreate := *openapiclient.NewAwsIntegrationCreate("AwsAccountId_example", []openapiclient.AwsRegionEnum{openapiclient.AwsRegionEnum("af-south-1")}, []openapiclient.AwsServiceEnum{openapiclient.AwsServiceEnum("s3")}, "AwsRoleName_example") // AwsIntegrationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsCreate(context.Background()).AwsIntegrationCreate(awsIntegrationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsCreate`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsIntegrationCreate** | [**AwsIntegrationCreate**](AwsIntegrationCreate.md) |  | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsDestroy

> IntegrationsAwsDestroy(ctx, id).InUse(inUse).Execute()

Delete an AWS integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    inUse := "inUse_example" // string | (Optional) Desired behavior if the integration has in-use values.  - `fail` will return HTTP error 409 if there are any values using the integration. - `leave` (default) will leave values in place and future queries may fail; you can control future value query behavior with the `lookup_error` query parameter on those requests. - `remove` will remove the all values using the integration when the integration is removed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsDestroy(context.Background(), id).InUse(inUse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inUse** | **string** | (Optional) Desired behavior if the integration has in-use values.  - &#x60;fail&#x60; will return HTTP error 409 if there are any values using the integration. - &#x60;leave&#x60; (default) will leave values in place and future queries may fail; you can control future value query behavior with the &#x60;lookup_error&#x60; query parameter on those requests. - &#x60;remove&#x60; will remove the all values using the integration when the integration is removed. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsList

> PaginatedAwsIntegrationList IntegrationsAwsList(ctx).AwsAccountId(awsAccountId).AwsRoleName(awsRoleName).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsAccountId := "awsAccountId_example" // string |  (optional)
    awsRoleName := "awsRoleName_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsList(context.Background()).AwsAccountId(awsAccountId).AwsRoleName(awsRoleName).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsList`: PaginatedAwsIntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAccountId** | **string** |  | 
 **awsRoleName** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAwsIntegrationList**](PaginatedAwsIntegrationList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPartialUpdate

> AwsIntegration IntegrationsAwsPartialUpdate(ctx, id).PatchedAwsIntegration(patchedAwsIntegration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedAwsIntegration := *openapiclient.NewPatchedAwsIntegration() // PatchedAwsIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPartialUpdate(context.Background(), id).PatchedAwsIntegration(patchedAwsIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPartialUpdate`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAwsIntegration** | [**PatchedAwsIntegration**](PatchedAwsIntegration.md) |  | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsCreate

> AwsPull IntegrationsAwsPullsCreate(ctx, awsintegrationPk).AwsPull(awsPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsPull := *openapiclient.NewAwsPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []openapiclient.Value{*openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "TODO", "TODO", "TODO", "Resource_example") // AwsPull | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsCreate(context.Background(), awsintegrationPk).AwsPull(awsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsCreate`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsPull** | [**AwsPull**](AwsPull.md) |  | 

### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsDestroy

> IntegrationsAwsPullsDestroy(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsDestroy(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsList

> PaginatedAwsPullList IntegrationsAwsPullsList(ctx, awsintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsList(context.Background(), awsintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsList`: PaginatedAwsPullList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAwsPullList**](PaginatedAwsPullList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsPartialUpdate

> AwsPull IntegrationsAwsPullsPartialUpdate(ctx, awsintegrationPk, id).PatchedAwsPull(patchedAwsPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedAwsPull := *openapiclient.NewPatchedAwsPull() // PatchedAwsPull |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsPartialUpdate(context.Background(), awsintegrationPk, id).PatchedAwsPull(patchedAwsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsPartialUpdate`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedAwsPull** | [**PatchedAwsPull**](PatchedAwsPull.md) |  | 

### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsRetrieve

> AwsPull IntegrationsAwsPullsRetrieve(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsRetrieve(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsRetrieve`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsSyncCreate

> IntegrationsAwsPullsSyncCreate(ctx, awsintegrationPk, id).AwsPullSyncActionRequest(awsPullSyncActionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsPullSyncActionRequest := *openapiclient.NewAwsPullSyncActionRequest() // AwsPullSyncActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsSyncCreate(context.Background(), awsintegrationPk, id).AwsPullSyncActionRequest(awsPullSyncActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPullSyncActionRequest** | [**AwsPullSyncActionRequest**](AwsPullSyncActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksList

> PaginatedAwsPullTaskList IntegrationsAwsPullsTasksList(ctx, awsintegrationPk, awspullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsTasksList(context.Background(), awsintegrationPk, awspullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksList`: PaginatedAwsPullTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspullPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedAwsPullTaskList**](PaginatedAwsPullTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksRetrieve

> AwsPullTask IntegrationsAwsPullsTasksRetrieve(ctx, awsintegrationPk, awspullPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsTasksRetrieve(context.Background(), awsintegrationPk, awspullPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksRetrieve`: AwsPullTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspullPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AwsPullTask**](AwsPullTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksStepsList

> PaginatedAwsPullTaskStepList IntegrationsAwsPullsTasksStepsList(ctx, awsintegrationPk, awspullPk, awspulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspulltaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsTasksStepsList(context.Background(), awsintegrationPk, awspullPk, awspulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksStepsList`: PaginatedAwsPullTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspullPk** | **string** |  | 
**awspulltaskPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedAwsPullTaskStepList**](PaginatedAwsPullTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsTasksStepsRetrieve

> AwsPullTaskStep IntegrationsAwsPullsTasksStepsRetrieve(ctx, awsintegrationPk, awspullPk, awspulltaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspulltaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsTasksStepsRetrieve(context.Background(), awsintegrationPk, awspullPk, awspulltaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsTasksStepsRetrieve`: AwsPullTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspullPk** | **string** |  | 
**awspulltaskPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AwsPullTaskStep**](AwsPullTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPullsUpdate

> AwsPull IntegrationsAwsPullsUpdate(ctx, awsintegrationPk, id).AwsPull(awsPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsPull := *openapiclient.NewAwsPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []openapiclient.Value{*openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "TODO", "TODO", "TODO", "Resource_example") // AwsPull | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPullsUpdate(context.Background(), awsintegrationPk, id).AwsPull(awsPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPullsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPullsUpdate`: AwsPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPullsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPullsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPull** | [**AwsPull**](AwsPull.md) |  | 

### Return type

[**AwsPull**](AwsPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesCreate

> AwsPush IntegrationsAwsPushesCreate(ctx, awsintegrationPk).AwsPush(awsPush).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsPush := *openapiclient.NewAwsPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, "TODO", "TODO", "Resource_example") // AwsPush | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesCreate(context.Background(), awsintegrationPk).AwsPush(awsPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesCreate`: AwsPush
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsPush** | [**AwsPush**](AwsPush.md) |  | 

### Return type

[**AwsPush**](AwsPush.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesDestroy

> IntegrationsAwsPushesDestroy(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesDestroy(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesList

> PaginatedAwsPushList IntegrationsAwsPushesList(ctx, awsintegrationPk).DescriptionIcontains(descriptionIcontains).Environment(environment).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Project(project).Tag(tag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    environment := "environment_example" // string | Filter results to pushes that operate on a tag in the given environment. (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    project := "project_example" // string | Filter results to pushes that operate on the given project. (optional)
    tag := "tag_example" // string | Filter results to pushes that operate on the given tag. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesList(context.Background(), awsintegrationPk).DescriptionIcontains(descriptionIcontains).Environment(environment).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Project(project).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesList`: PaginatedAwsPushList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **environment** | **string** | Filter results to pushes that operate on a tag in the given environment. | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **project** | **string** | Filter results to pushes that operate on the given project. | 
 **tag** | **string** | Filter results to pushes that operate on the given tag. | 

### Return type

[**PaginatedAwsPushList**](PaginatedAwsPushList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesPartialUpdate

> AwsPushUpdate IntegrationsAwsPushesPartialUpdate(ctx, awsintegrationPk, id).PatchedAwsPushUpdate(patchedAwsPushUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedAwsPushUpdate := *openapiclient.NewPatchedAwsPushUpdate() // PatchedAwsPushUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesPartialUpdate(context.Background(), awsintegrationPk, id).PatchedAwsPushUpdate(patchedAwsPushUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesPartialUpdate`: AwsPushUpdate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedAwsPushUpdate** | [**PatchedAwsPushUpdate**](PatchedAwsPushUpdate.md) |  | 

### Return type

[**AwsPushUpdate**](AwsPushUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesRetrieve

> AwsPush IntegrationsAwsPushesRetrieve(ctx, awsintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesRetrieve(context.Background(), awsintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesRetrieve`: AwsPush
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AwsPush**](AwsPush.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesSyncCreate

> IntegrationsAwsPushesSyncCreate(ctx, awsintegrationPk, id).AwsPush(awsPush).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsPush := *openapiclient.NewAwsPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, "TODO", "TODO", "Resource_example") // AwsPush | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesSyncCreate(context.Background(), awsintegrationPk, id).AwsPush(awsPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPush** | [**AwsPush**](AwsPush.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksList

> PaginatedAwsPushTaskList IntegrationsAwsPushesTasksList(ctx, awsintegrationPk, awspushPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesTasksList(context.Background(), awsintegrationPk, awspushPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksList`: PaginatedAwsPushTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspushPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedAwsPushTaskList**](PaginatedAwsPushTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksRetrieve

> AwsPushTask IntegrationsAwsPushesTasksRetrieve(ctx, awsintegrationPk, awspushPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesTasksRetrieve(context.Background(), awsintegrationPk, awspushPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksRetrieve`: AwsPushTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspushPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AwsPushTask**](AwsPushTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksStepsList

> PaginatedAwsPushTaskStepList IntegrationsAwsPushesTasksStepsList(ctx, awsintegrationPk, awspushPk, awspushtaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspushtaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesTasksStepsList(context.Background(), awsintegrationPk, awspushPk, awspushtaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksStepsList`: PaginatedAwsPushTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspushPk** | **string** |  | 
**awspushtaskPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedAwsPushTaskStepList**](PaginatedAwsPushTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesTasksStepsRetrieve

> AwsPushTaskStep IntegrationsAwsPushesTasksStepsRetrieve(ctx, awsintegrationPk, awspushPk, awspushtaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awspushtaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesTasksStepsRetrieve(context.Background(), awsintegrationPk, awspushPk, awspushtaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesTasksStepsRetrieve`: AwsPushTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**awspushPk** | **string** |  | 
**awspushtaskPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AwsPushTaskStep**](AwsPushTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsPushesUpdate

> AwsPushUpdate IntegrationsAwsPushesUpdate(ctx, awsintegrationPk, id).AwsPushUpdate(awsPushUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    awsintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsPushUpdate := *openapiclient.NewAwsPushUpdate("Name_example", []string{"Projects_example"}, []string{"Tags_example"}, "Resource_example") // AwsPushUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsPushesUpdate(context.Background(), awsintegrationPk, id).AwsPushUpdate(awsPushUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsPushesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsPushesUpdate`: AwsPushUpdate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsPushesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsPushesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsPushUpdate** | [**AwsPushUpdate**](AwsPushUpdate.md) |  | 

### Return type

[**AwsPushUpdate**](AwsPushUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsRetrieve

> AwsIntegration IntegrationsAwsRetrieve(ctx, id).RefreshStatus(refreshStatus).Execute()

Get details of an AWS Integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    refreshStatus := true // bool | Trigger a refresh of the integration status before returning the details. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsRetrieve(context.Background(), id).RefreshStatus(refreshStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsRetrieve`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshStatus** | **bool** | Trigger a refresh of the integration status before returning the details. | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsScanCreate

> DiscoveryResult IntegrationsAwsScanCreate(ctx, id).AwsIntegrationScan(awsIntegrationScan).Execute()

Evaluate a potential pull pattern and understand what it will match.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsIntegrationScan := *openapiclient.NewAwsIntegrationScan("TODO", "TODO", "Resource_example") // AwsIntegrationScan | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsScanCreate(context.Background(), id).AwsIntegrationScan(awsIntegrationScan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsScanCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsScanCreate`: DiscoveryResult
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsScanCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsScanCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsIntegrationScan** | [**AwsIntegrationScan**](AwsIntegrationScan.md) |  | 

### Return type

[**DiscoveryResult**](DiscoveryResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAwsUpdate

> AwsIntegration IntegrationsAwsUpdate(ctx, id).AwsIntegration(awsIntegration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    awsIntegration := *openapiclient.NewAwsIntegration("Url_example", "Id_example", "Name_example", "TODO", "StatusDetail_example", time.Now(), time.Now(), time.Now(), "Fqn_example", "Type_example", "AwsAccountId_example", []openapiclient.AwsRegionEnum{openapiclient.AwsRegionEnum("af-south-1")}, []openapiclient.AwsServiceEnum{openapiclient.AwsServiceEnum("s3")}, "AwsExternalId_example", "AwsRoleName_example") // AwsIntegration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAwsUpdate(context.Background(), id).AwsIntegration(awsIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAwsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAwsUpdate`: AwsIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAwsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAwsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsIntegration** | [**AwsIntegration**](AwsIntegration.md) |  | 

### Return type

[**AwsIntegration**](AwsIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultCreate

> AzureKeyVaultIntegration IntegrationsAzureKeyVaultCreate(ctx).AzureKeyVaultIntegrationCreate(azureKeyVaultIntegrationCreate).Execute()

Establishes an Azure Key Vault Integration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    azureKeyVaultIntegrationCreate := *openapiclient.NewAzureKeyVaultIntegrationCreate("VaultName_example", "TenantId_example") // AzureKeyVaultIntegrationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultCreate(context.Background()).AzureKeyVaultIntegrationCreate(azureKeyVaultIntegrationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultCreate`: AzureKeyVaultIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureKeyVaultIntegrationCreate** | [**AzureKeyVaultIntegrationCreate**](AzureKeyVaultIntegrationCreate.md) |  | 

### Return type

[**AzureKeyVaultIntegration**](AzureKeyVaultIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultDestroy

> IntegrationsAzureKeyVaultDestroy(ctx, id).InUse(inUse).Execute()

Delete an AWS integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    inUse := "inUse_example" // string | (Optional) Desired behavior if the integration has in-use values.  - `fail` will return HTTP error 409 if there are any values using the integration. - `leave` (default) will leave values in place and future queries may fail; you can control future value query behavior with the `lookup_error` query parameter on those requests. - `remove` will remove the all values using the integration when the integration is removed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultDestroy(context.Background(), id).InUse(inUse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inUse** | **string** | (Optional) Desired behavior if the integration has in-use values.  - &#x60;fail&#x60; will return HTTP error 409 if there are any values using the integration. - &#x60;leave&#x60; (default) will leave values in place and future queries may fail; you can control future value query behavior with the &#x60;lookup_error&#x60; query parameter on those requests. - &#x60;remove&#x60; will remove the all values using the integration when the integration is removed. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultList

> PaginatedAzureKeyVaultIntegrationList IntegrationsAzureKeyVaultList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).TenantId(tenantId).VaultName(vaultName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    vaultName := "vaultName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).TenantId(tenantId).VaultName(vaultName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultList`: PaginatedAzureKeyVaultIntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **tenantId** | **string** |  | 
 **vaultName** | **string** |  | 

### Return type

[**PaginatedAzureKeyVaultIntegrationList**](PaginatedAzureKeyVaultIntegrationList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPartialUpdate

> AzureKeyVaultIntegration IntegrationsAzureKeyVaultPartialUpdate(ctx, id).PatchedAzureKeyVaultIntegration(patchedAzureKeyVaultIntegration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedAzureKeyVaultIntegration := *openapiclient.NewPatchedAzureKeyVaultIntegration() // PatchedAzureKeyVaultIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPartialUpdate(context.Background(), id).PatchedAzureKeyVaultIntegration(patchedAzureKeyVaultIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPartialUpdate`: AzureKeyVaultIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAzureKeyVaultIntegration** | [**PatchedAzureKeyVaultIntegration**](PatchedAzureKeyVaultIntegration.md) |  | 

### Return type

[**AzureKeyVaultIntegration**](AzureKeyVaultIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsCreate

> AzureKeyVaultPull IntegrationsAzureKeyVaultPullsCreate(ctx, akvintegrationPk).AzureKeyVaultPull(azureKeyVaultPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultPull := *openapiclient.NewAzureKeyVaultPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []openapiclient.Value{*openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "TODO", "Resource_example") // AzureKeyVaultPull | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsCreate(context.Background(), akvintegrationPk).AzureKeyVaultPull(azureKeyVaultPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsCreate`: AzureKeyVaultPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureKeyVaultPull** | [**AzureKeyVaultPull**](AzureKeyVaultPull.md) |  | 

### Return type

[**AzureKeyVaultPull**](AzureKeyVaultPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsDestroy

> IntegrationsAzureKeyVaultPullsDestroy(ctx, akvintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsDestroy(context.Background(), akvintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsList

> PaginatedAzureKeyVaultPullList IntegrationsAzureKeyVaultPullsList(ctx, akvintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsList(context.Background(), akvintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsList`: PaginatedAzureKeyVaultPullList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedAzureKeyVaultPullList**](PaginatedAzureKeyVaultPullList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsPartialUpdate

> AzureKeyVaultPull IntegrationsAzureKeyVaultPullsPartialUpdate(ctx, akvintegrationPk, id).PatchedAzureKeyVaultPull(patchedAzureKeyVaultPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedAzureKeyVaultPull := *openapiclient.NewPatchedAzureKeyVaultPull() // PatchedAzureKeyVaultPull |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsPartialUpdate(context.Background(), akvintegrationPk, id).PatchedAzureKeyVaultPull(patchedAzureKeyVaultPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsPartialUpdate`: AzureKeyVaultPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedAzureKeyVaultPull** | [**PatchedAzureKeyVaultPull**](PatchedAzureKeyVaultPull.md) |  | 

### Return type

[**AzureKeyVaultPull**](AzureKeyVaultPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsRetrieve

> AzureKeyVaultPull IntegrationsAzureKeyVaultPullsRetrieve(ctx, akvintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsRetrieve(context.Background(), akvintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsRetrieve`: AzureKeyVaultPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AzureKeyVaultPull**](AzureKeyVaultPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsSyncCreate

> IntegrationsAzureKeyVaultPullsSyncCreate(ctx, akvintegrationPk, id).AzureKeyVaultPullSyncActionRequest(azureKeyVaultPullSyncActionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultPullSyncActionRequest := *openapiclient.NewAzureKeyVaultPullSyncActionRequest() // AzureKeyVaultPullSyncActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsSyncCreate(context.Background(), akvintegrationPk, id).AzureKeyVaultPullSyncActionRequest(azureKeyVaultPullSyncActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureKeyVaultPullSyncActionRequest** | [**AzureKeyVaultPullSyncActionRequest**](AzureKeyVaultPullSyncActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsTasksList

> PaginatedAzureKeyVaultPullTaskList IntegrationsAzureKeyVaultPullsTasksList(ctx, akvintegrationPk, akvpullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksList(context.Background(), akvintegrationPk, akvpullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsTasksList`: PaginatedAzureKeyVaultPullTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpullPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedAzureKeyVaultPullTaskList**](PaginatedAzureKeyVaultPullTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsTasksRetrieve

> AzureKeyVaultPullTask IntegrationsAzureKeyVaultPullsTasksRetrieve(ctx, akvintegrationPk, akvpullPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksRetrieve(context.Background(), akvintegrationPk, akvpullPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsTasksRetrieve`: AzureKeyVaultPullTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpullPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AzureKeyVaultPullTask**](AzureKeyVaultPullTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsTasksStepsList

> PaginatedAzureKeyVaultPullTaskStepList IntegrationsAzureKeyVaultPullsTasksStepsList(ctx, akvintegrationPk, akvpullPk, akvpulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpulltaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksStepsList(context.Background(), akvintegrationPk, akvpullPk, akvpulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsTasksStepsList`: PaginatedAzureKeyVaultPullTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpullPk** | **string** |  | 
**akvpulltaskPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedAzureKeyVaultPullTaskStepList**](PaginatedAzureKeyVaultPullTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsTasksStepsRetrieve

> AzureKeyVaultPullTaskStep IntegrationsAzureKeyVaultPullsTasksStepsRetrieve(ctx, akvintegrationPk, akvpullPk, akvpulltaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpulltaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksStepsRetrieve(context.Background(), akvintegrationPk, akvpullPk, akvpulltaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsTasksStepsRetrieve`: AzureKeyVaultPullTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpullPk** | **string** |  | 
**akvpulltaskPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AzureKeyVaultPullTaskStep**](AzureKeyVaultPullTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPullsUpdate

> AzureKeyVaultPull IntegrationsAzureKeyVaultPullsUpdate(ctx, akvintegrationPk, id).AzureKeyVaultPull(azureKeyVaultPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultPull := *openapiclient.NewAzureKeyVaultPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []openapiclient.Value{*openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "TODO", "Resource_example") // AzureKeyVaultPull | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPullsUpdate(context.Background(), akvintegrationPk, id).AzureKeyVaultPull(azureKeyVaultPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPullsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPullsUpdate`: AzureKeyVaultPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPullsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPullsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureKeyVaultPull** | [**AzureKeyVaultPull**](AzureKeyVaultPull.md) |  | 

### Return type

[**AzureKeyVaultPull**](AzureKeyVaultPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesCreate

> AzureKeyVaultPush IntegrationsAzureKeyVaultPushesCreate(ctx, akvintegrationPk).AzureKeyVaultPush(azureKeyVaultPush).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultPush := *openapiclient.NewAzureKeyVaultPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, "Resource_example") // AzureKeyVaultPush | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesCreate(context.Background(), akvintegrationPk).AzureKeyVaultPush(azureKeyVaultPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesCreate`: AzureKeyVaultPush
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureKeyVaultPush** | [**AzureKeyVaultPush**](AzureKeyVaultPush.md) |  | 

### Return type

[**AzureKeyVaultPush**](AzureKeyVaultPush.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesDestroy

> IntegrationsAzureKeyVaultPushesDestroy(ctx, akvintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesDestroy(context.Background(), akvintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesList

> PaginatedAzureKeyVaultPushList IntegrationsAzureKeyVaultPushesList(ctx, akvintegrationPk).DescriptionIcontains(descriptionIcontains).Environment(environment).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Project(project).Tag(tag).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    environment := "environment_example" // string | Filter results to pushes that operate on a tag in the given environment. (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    project := "project_example" // string | Filter results to pushes that operate on the given project. (optional)
    tag := "tag_example" // string | Filter results to pushes that operate on the given tag. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesList(context.Background(), akvintegrationPk).DescriptionIcontains(descriptionIcontains).Environment(environment).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Project(project).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesList`: PaginatedAzureKeyVaultPushList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **environment** | **string** | Filter results to pushes that operate on a tag in the given environment. | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **project** | **string** | Filter results to pushes that operate on the given project. | 
 **tag** | **string** | Filter results to pushes that operate on the given tag. | 

### Return type

[**PaginatedAzureKeyVaultPushList**](PaginatedAzureKeyVaultPushList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesPartialUpdate

> AzureKeyVaultPushUpdate IntegrationsAzureKeyVaultPushesPartialUpdate(ctx, akvintegrationPk, id).PatchedAzureKeyVaultPushUpdate(patchedAzureKeyVaultPushUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedAzureKeyVaultPushUpdate := *openapiclient.NewPatchedAzureKeyVaultPushUpdate() // PatchedAzureKeyVaultPushUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesPartialUpdate(context.Background(), akvintegrationPk, id).PatchedAzureKeyVaultPushUpdate(patchedAzureKeyVaultPushUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesPartialUpdate`: AzureKeyVaultPushUpdate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedAzureKeyVaultPushUpdate** | [**PatchedAzureKeyVaultPushUpdate**](PatchedAzureKeyVaultPushUpdate.md) |  | 

### Return type

[**AzureKeyVaultPushUpdate**](AzureKeyVaultPushUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesRetrieve

> AzureKeyVaultPush IntegrationsAzureKeyVaultPushesRetrieve(ctx, akvintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesRetrieve(context.Background(), akvintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesRetrieve`: AzureKeyVaultPush
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AzureKeyVaultPush**](AzureKeyVaultPush.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesSyncCreate

> IntegrationsAzureKeyVaultPushesSyncCreate(ctx, akvintegrationPk, id).AzureKeyVaultPush(azureKeyVaultPush).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultPush := *openapiclient.NewAzureKeyVaultPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, "Resource_example") // AzureKeyVaultPush | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesSyncCreate(context.Background(), akvintegrationPk, id).AzureKeyVaultPush(azureKeyVaultPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureKeyVaultPush** | [**AzureKeyVaultPush**](AzureKeyVaultPush.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesTasksList

> PaginatedAzureKeyVaultPushTaskList IntegrationsAzureKeyVaultPushesTasksList(ctx, akvintegrationPk, akvpushPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksList(context.Background(), akvintegrationPk, akvpushPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesTasksList`: PaginatedAzureKeyVaultPushTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpushPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedAzureKeyVaultPushTaskList**](PaginatedAzureKeyVaultPushTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesTasksRetrieve

> AzureKeyVaultPushTask IntegrationsAzureKeyVaultPushesTasksRetrieve(ctx, akvintegrationPk, akvpushPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksRetrieve(context.Background(), akvintegrationPk, akvpushPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesTasksRetrieve`: AzureKeyVaultPushTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpushPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AzureKeyVaultPushTask**](AzureKeyVaultPushTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesTasksStepsList

> PaginatedAzureKeyVaultPushTaskStepList IntegrationsAzureKeyVaultPushesTasksStepsList(ctx, akvintegrationPk, akvpushPk, akvpushtaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpushtaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksStepsList(context.Background(), akvintegrationPk, akvpushPk, akvpushtaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesTasksStepsList`: PaginatedAzureKeyVaultPushTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpushPk** | **string** |  | 
**akvpushtaskPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedAzureKeyVaultPushTaskStepList**](PaginatedAzureKeyVaultPushTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesTasksStepsRetrieve

> AzureKeyVaultPushTaskStep IntegrationsAzureKeyVaultPushesTasksStepsRetrieve(ctx, akvintegrationPk, akvpushPk, akvpushtaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpushPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    akvpushtaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksStepsRetrieve(context.Background(), akvintegrationPk, akvpushPk, akvpushtaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesTasksStepsRetrieve`: AzureKeyVaultPushTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**akvpushPk** | **string** |  | 
**akvpushtaskPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AzureKeyVaultPushTaskStep**](AzureKeyVaultPushTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultPushesUpdate

> AzureKeyVaultPushUpdate IntegrationsAzureKeyVaultPushesUpdate(ctx, akvintegrationPk, id).AzureKeyVaultPushUpdate(azureKeyVaultPushUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    akvintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultPushUpdate := *openapiclient.NewAzureKeyVaultPushUpdate("Name_example", []string{"Projects_example"}, []string{"Tags_example"}, "Resource_example") // AzureKeyVaultPushUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultPushesUpdate(context.Background(), akvintegrationPk, id).AzureKeyVaultPushUpdate(azureKeyVaultPushUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultPushesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultPushesUpdate`: AzureKeyVaultPushUpdate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultPushesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**akvintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultPushesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureKeyVaultPushUpdate** | [**AzureKeyVaultPushUpdate**](AzureKeyVaultPushUpdate.md) |  | 

### Return type

[**AzureKeyVaultPushUpdate**](AzureKeyVaultPushUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultRetrieve

> AzureKeyVaultIntegration IntegrationsAzureKeyVaultRetrieve(ctx, id).RefreshStatus(refreshStatus).Execute()

Get details of an Azure Key Vault Integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    refreshStatus := true // bool | Trigger a refresh of the integration status before returning the details. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultRetrieve(context.Background(), id).RefreshStatus(refreshStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultRetrieve`: AzureKeyVaultIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshStatus** | **bool** | Trigger a refresh of the integration status before returning the details. | 

### Return type

[**AzureKeyVaultIntegration**](AzureKeyVaultIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultScanCreate

> DiscoveryResult IntegrationsAzureKeyVaultScanCreate(ctx, id).AzureKeyVaultIntegrationScan(azureKeyVaultIntegrationScan).Execute()

Evaluate a potential pull pattern and understand what it will match.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultIntegrationScan := *openapiclient.NewAzureKeyVaultIntegrationScan("Resource_example") // AzureKeyVaultIntegrationScan | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultScanCreate(context.Background(), id).AzureKeyVaultIntegrationScan(azureKeyVaultIntegrationScan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultScanCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultScanCreate`: DiscoveryResult
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultScanCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultScanCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureKeyVaultIntegrationScan** | [**AzureKeyVaultIntegrationScan**](AzureKeyVaultIntegrationScan.md) |  | 

### Return type

[**DiscoveryResult**](DiscoveryResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsAzureKeyVaultUpdate

> AzureKeyVaultIntegration IntegrationsAzureKeyVaultUpdate(ctx, id).AzureKeyVaultIntegration(azureKeyVaultIntegration).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    azureKeyVaultIntegration := *openapiclient.NewAzureKeyVaultIntegration("Url_example", "Id_example", "Name_example", "TODO", "StatusDetail_example", time.Now(), time.Now(), time.Now(), "Fqn_example", "Type_example", "VaultName_example", "TenantId_example") // AzureKeyVaultIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsAzureKeyVaultUpdate(context.Background(), id).AzureKeyVaultIntegration(azureKeyVaultIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsAzureKeyVaultUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsAzureKeyVaultUpdate`: AzureKeyVaultIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsAzureKeyVaultUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsAzureKeyVaultUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureKeyVaultIntegration** | [**AzureKeyVaultIntegration**](AzureKeyVaultIntegration.md) |  | 

### Return type

[**AzureKeyVaultIntegration**](AzureKeyVaultIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsExploreList

> PaginatedIntegrationNodeList IntegrationsExploreList(ctx).Fqn(fqn).Jmes(jmes).Ordering(ordering).Page(page).PageSize(pageSize).Execute()

Retrieve third-party integration data for the specified FQN.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fqn := "fqn_example" // string | FQN (URL-like) for third-party integration. (optional)
    jmes := "jmes_example" // string | JMES-path qualifier (key within file) (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsExploreList(context.Background()).Fqn(fqn).Jmes(jmes).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsExploreList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsExploreList`: PaginatedIntegrationNodeList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsExploreList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsExploreListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fqn** | **string** | FQN (URL-like) for third-party integration. | 
 **jmes** | **string** | JMES-path qualifier (key within file) | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedIntegrationNodeList**](PaginatedIntegrationNodeList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubCreate

> GitHubIntegration IntegrationsGithubCreate(ctx).GitHubIntegrationCreate(gitHubIntegrationCreate).Execute()

Establishes a GitHub Integration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gitHubIntegrationCreate := *openapiclient.NewGitHubIntegrationCreate(int32(123)) // GitHubIntegrationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubCreate(context.Background()).GitHubIntegrationCreate(gitHubIntegrationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubCreate`: GitHubIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitHubIntegrationCreate** | [**GitHubIntegrationCreate**](GitHubIntegrationCreate.md) |  | 

### Return type

[**GitHubIntegration**](GitHubIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubDestroy

> IntegrationsGithubDestroy(ctx, id).InUse(inUse).Execute()

Delete a GitHub integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    inUse := "inUse_example" // string | (Optional) Desired behavior if the integration has in-use values.  - `fail` will return HTTP error 409 if there are any values using the integration. - `leave` (default) will leave values in place and future queries may fail; you can control future value query behavior with the `lookup_error` query parameter on those requests. - `remove` will remove the all values using the integration when the integration is removed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubDestroy(context.Background(), id).InUse(inUse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inUse** | **string** | (Optional) Desired behavior if the integration has in-use values.  - &#x60;fail&#x60; will return HTTP error 409 if there are any values using the integration. - &#x60;leave&#x60; (default) will leave values in place and future queries may fail; you can control future value query behavior with the &#x60;lookup_error&#x60; query parameter on those requests. - &#x60;remove&#x60; will remove the all values using the integration when the integration is removed. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubList

> PaginatedGitHubIntegrationList IntegrationsGithubList(ctx).GhOrganizationSlug(ghOrganizationSlug).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ghOrganizationSlug := "ghOrganizationSlug_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubList(context.Background()).GhOrganizationSlug(ghOrganizationSlug).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubList`: PaginatedGitHubIntegrationList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ghOrganizationSlug** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedGitHubIntegrationList**](PaginatedGitHubIntegrationList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsList

> PaginatedGitHubPullList IntegrationsGithubPullsList(ctx, githubintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsList(context.Background(), githubintegrationPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsList`: PaginatedGitHubPullList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedGitHubPullList**](PaginatedGitHubPullList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsPartialUpdate

> GitHubPull IntegrationsGithubPullsPartialUpdate(ctx, githubintegrationPk, id).PatchedGitHubPull(patchedGitHubPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedGitHubPull := *openapiclient.NewPatchedGitHubPull() // PatchedGitHubPull |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsPartialUpdate(context.Background(), githubintegrationPk, id).PatchedGitHubPull(patchedGitHubPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsPartialUpdate`: GitHubPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedGitHubPull** | [**PatchedGitHubPull**](PatchedGitHubPull.md) |  | 

### Return type

[**GitHubPull**](GitHubPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsRetrieve

> GitHubPull IntegrationsGithubPullsRetrieve(ctx, githubintegrationPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsRetrieve(context.Background(), githubintegrationPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsRetrieve`: GitHubPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GitHubPull**](GitHubPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsSyncCreate

> IntegrationsGithubPullsSyncCreate(ctx, githubintegrationPk, id).GitHubPull(gitHubPull).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    gitHubPull := *openapiclient.NewGitHubPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []openapiclient.Value{*openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "TODO") // GitHubPull | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsSyncCreate(context.Background(), githubintegrationPk, id).GitHubPull(gitHubPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsSyncCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitHubPull** | [**GitHubPull**](GitHubPull.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsTasksList

> PaginatedGitHubPullTaskList IntegrationsGithubPullsTasksList(ctx, githubintegrationPk, githubpullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    githubpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    state := "state_example" // string | The current state of this task. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsTasksList(context.Background(), githubintegrationPk, githubpullPk).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Ordering(ordering).Page(page).PageSize(pageSize).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsTasksList`: PaginatedGitHubPullTaskList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**githubpullPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **state** | **string** | The current state of this task. | 

### Return type

[**PaginatedGitHubPullTaskList**](PaginatedGitHubPullTaskList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsTasksRetrieve

> GitHubPullTask IntegrationsGithubPullsTasksRetrieve(ctx, githubintegrationPk, githubpullPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    githubpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsTasksRetrieve(context.Background(), githubintegrationPk, githubpullPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsTasksRetrieve`: GitHubPullTask
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**githubpullPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GitHubPullTask**](GitHubPullTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsTasksStepsList

> PaginatedGitHubPullTaskStepList IntegrationsGithubPullsTasksStepsList(ctx, githubintegrationPk, githubpullPk, githubpulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    githubpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    githubpulltaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    fqn := "fqn_example" // string |  (optional)
    modifiedAt := time.Now() // time.Time |  (optional)
    modifiedAtGte := time.Now() // time.Time |  (optional)
    modifiedAtLte := time.Now() // time.Time |  (optional)
    operation := "operation_example" // string | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    success := true // bool |  (optional)
    venueId := "venueId_example" // string |  (optional)
    venueIdIcontains := "venueIdIcontains_example" // string |  (optional)
    venueName := "venueName_example" // string |  (optional)
    venueNameIcontains := "venueNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsTasksStepsList(context.Background(), githubintegrationPk, githubpullPk, githubpulltaskPk).Fqn(fqn).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLte(modifiedAtLte).Operation(operation).Ordering(ordering).Page(page).PageSize(pageSize).Success(success).VenueId(venueId).VenueIdIcontains(venueIdIcontains).VenueName(venueName).VenueNameIcontains(venueNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsTasksStepsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsTasksStepsList`: PaginatedGitHubPullTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsTasksStepsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**githubpullPk** | **string** |  | 
**githubpulltaskPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsTasksStepsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fqn** | **string** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLte** | **time.Time** |  | 
 **operation** | **string** | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **success** | **bool** |  | 
 **venueId** | **string** |  | 
 **venueIdIcontains** | **string** |  | 
 **venueName** | **string** |  | 
 **venueNameIcontains** | **string** |  | 

### Return type

[**PaginatedGitHubPullTaskStepList**](PaginatedGitHubPullTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsTasksStepsRetrieve

> GitHubPullTaskStep IntegrationsGithubPullsTasksStepsRetrieve(ctx, githubintegrationPk, githubpullPk, githubpulltaskPk, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    githubpullPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    githubpulltaskPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsTasksStepsRetrieve(context.Background(), githubintegrationPk, githubpullPk, githubpulltaskPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsTasksStepsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsTasksStepsRetrieve`: GitHubPullTaskStep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsTasksStepsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**githubpullPk** | **string** |  | 
**githubpulltaskPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsTasksStepsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GitHubPullTaskStep**](GitHubPullTaskStep.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubPullsUpdate

> GitHubPull IntegrationsGithubPullsUpdate(ctx, githubintegrationPk, id).GitHubPull(gitHubPull).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    githubintegrationPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    gitHubPull := *openapiclient.NewGitHubPull("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []openapiclient.Value{*openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "TODO") // GitHubPull | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubPullsUpdate(context.Background(), githubintegrationPk, id).GitHubPull(gitHubPull).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubPullsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubPullsUpdate`: GitHubPull
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubPullsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubintegrationPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubPullsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitHubPull** | [**GitHubPull**](GitHubPull.md) |  | 

### Return type

[**GitHubPull**](GitHubPull.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGithubRetrieve

> GitHubIntegration IntegrationsGithubRetrieve(ctx, id).RefreshStatus(refreshStatus).Execute()

Get details of a GitHub Integration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    refreshStatus := true // bool | Refresh the integration status before returning the details. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.IntegrationsGithubRetrieve(context.Background(), id).RefreshStatus(refreshStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.IntegrationsGithubRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationsGithubRetrieve`: GitHubIntegration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.IntegrationsGithubRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGithubRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshStatus** | **bool** | Refresh the integration status before returning the details. | 

### Return type

[**GitHubIntegration**](GitHubIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

