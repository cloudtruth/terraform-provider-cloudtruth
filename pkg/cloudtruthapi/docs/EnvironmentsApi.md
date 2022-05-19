# \EnvironmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsCreate**](EnvironmentsApi.md#EnvironmentsCreate) | **Post** /api/v1/environments/ | 
[**EnvironmentsDestroy**](EnvironmentsApi.md#EnvironmentsDestroy) | **Delete** /api/v1/environments/{id}/ | 
[**EnvironmentsList**](EnvironmentsApi.md#EnvironmentsList) | **Get** /api/v1/environments/ | 
[**EnvironmentsPartialUpdate**](EnvironmentsApi.md#EnvironmentsPartialUpdate) | **Patch** /api/v1/environments/{id}/ | 
[**EnvironmentsPushesList**](EnvironmentsApi.md#EnvironmentsPushesList) | **Get** /api/v1/environments/{environment_pk}/pushes/ | List push operations.
[**EnvironmentsRetrieve**](EnvironmentsApi.md#EnvironmentsRetrieve) | **Get** /api/v1/environments/{id}/ | 
[**EnvironmentsTagsCreate**](EnvironmentsApi.md#EnvironmentsTagsCreate) | **Post** /api/v1/environments/{environment_pk}/tags/ | 
[**EnvironmentsTagsDestroy**](EnvironmentsApi.md#EnvironmentsTagsDestroy) | **Delete** /api/v1/environments/{environment_pk}/tags/{id}/ | 
[**EnvironmentsTagsList**](EnvironmentsApi.md#EnvironmentsTagsList) | **Get** /api/v1/environments/{environment_pk}/tags/ | 
[**EnvironmentsTagsPartialUpdate**](EnvironmentsApi.md#EnvironmentsTagsPartialUpdate) | **Patch** /api/v1/environments/{environment_pk}/tags/{id}/ | 
[**EnvironmentsTagsRetrieve**](EnvironmentsApi.md#EnvironmentsTagsRetrieve) | **Get** /api/v1/environments/{environment_pk}/tags/{id}/ | 
[**EnvironmentsTagsUpdate**](EnvironmentsApi.md#EnvironmentsTagsUpdate) | **Put** /api/v1/environments/{environment_pk}/tags/{id}/ | 
[**EnvironmentsUpdate**](EnvironmentsApi.md#EnvironmentsUpdate) | **Put** /api/v1/environments/{id}/ | 



## EnvironmentsCreate

> Environment EnvironmentsCreate(ctx).EnvironmentCreate(environmentCreate).Execute()



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
    environmentCreate := *openapiclient.NewEnvironmentCreate("Name_example") // EnvironmentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsCreate(context.Background()).EnvironmentCreate(environmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsCreate`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCreate** | [**EnvironmentCreate**](EnvironmentCreate.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDestroy

> EnvironmentsDestroy(ctx, id).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsDestroyRequest struct via the builder pattern


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


## EnvironmentsList

> PaginatedEnvironmentList EnvironmentsList(ctx).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).ParentName(parentName).ParentNameIcontains(parentNameIcontains).Execute()



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
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    parentName := "parentName_example" // string |  (optional)
    parentNameIcontains := "parentNameIcontains_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsList(context.Background()).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).ParentName(parentName).ParentNameIcontains(parentNameIcontains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsList`: PaginatedEnvironmentList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **parentName** | **string** |  | 
 **parentNameIcontains** | **string** |  | 

### Return type

[**PaginatedEnvironmentList**](PaginatedEnvironmentList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsPartialUpdate

> Environment EnvironmentsPartialUpdate(ctx, id).PatchedEnvironment(patchedEnvironment).Execute()



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
    patchedEnvironment := *openapiclient.NewPatchedEnvironment() // PatchedEnvironment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsPartialUpdate(context.Background(), id).PatchedEnvironment(patchedEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsPartialUpdate`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEnvironment** | [**PatchedEnvironment**](PatchedEnvironment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsPushesList

> PaginatedTaskStepList EnvironmentsPushesList(ctx, environmentPk).Ordering(ordering).Page(page).PageSize(pageSize).Execute()

List push operations.



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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsPushesList(context.Background(), environmentPk).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsPushesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsPushesList`: PaginatedTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsPushesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsPushesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedTaskStepList**](PaginatedTaskStepList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsRetrieve

> Environment EnvironmentsRetrieve(ctx, id).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsRetrieve`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsTagsCreate

> Tag EnvironmentsTagsCreate(ctx, environmentPk).TagCreate(tagCreate).Execute()





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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagCreate := *openapiclient.NewTagCreate("Name_example") // TagCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsTagsCreate(context.Background(), environmentPk).TagCreate(tagCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsTagsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsTagsCreate`: Tag
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsTagsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsTagsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagCreate** | [**TagCreate**](TagCreate.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsTagsDestroy

> EnvironmentsTagsDestroy(ctx, environmentPk, id).Execute()





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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsTagsDestroy(context.Background(), environmentPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsTagsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsTagsDestroyRequest struct via the builder pattern


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


## EnvironmentsTagsList

> PaginatedTagList EnvironmentsTagsList(ctx, environmentPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Timestamp(timestamp).TimestampGte(timestampGte).TimestampLte(timestampLte).Execute()





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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    timestamp := time.Now() // time.Time |  (optional)
    timestampGte := time.Now() // time.Time |  (optional)
    timestampLte := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsTagsList(context.Background(), environmentPk).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Timestamp(timestamp).TimestampGte(timestampGte).TimestampLte(timestampLte).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsTagsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsTagsList`: PaginatedTagList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsTagsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsTagsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **timestamp** | **time.Time** |  | 
 **timestampGte** | **time.Time** |  | 
 **timestampLte** | **time.Time** |  | 

### Return type

[**PaginatedTagList**](PaginatedTagList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsTagsPartialUpdate

> TagUpdate EnvironmentsTagsPartialUpdate(ctx, environmentPk, id).PatchedTagUpdate(patchedTagUpdate).Execute()





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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedTagUpdate := *openapiclient.NewPatchedTagUpdate() // PatchedTagUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsTagsPartialUpdate(context.Background(), environmentPk, id).PatchedTagUpdate(patchedTagUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsTagsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsTagsPartialUpdate`: TagUpdate
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsTagsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsTagsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedTagUpdate** | [**PatchedTagUpdate**](PatchedTagUpdate.md) |  | 

### Return type

[**TagUpdate**](TagUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsTagsRetrieve

> Tag EnvironmentsTagsRetrieve(ctx, environmentPk, id).Execute()





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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsTagsRetrieve(context.Background(), environmentPk, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsTagsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsTagsRetrieve`: Tag
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsTagsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsTagsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsTagsUpdate

> TagUpdate EnvironmentsTagsUpdate(ctx, environmentPk, id).TagUpdate(tagUpdate).Execute()





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
    environmentPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagUpdate := *openapiclient.NewTagUpdate("Id_example", "Name_example") // TagUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsTagsUpdate(context.Background(), environmentPk, id).TagUpdate(tagUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsTagsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsTagsUpdate`: TagUpdate
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsTagsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentPk** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsTagsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagUpdate** | [**TagUpdate**](TagUpdate.md) |  | 

### Return type

[**TagUpdate**](TagUpdate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsUpdate

> Environment EnvironmentsUpdate(ctx, id).Environment(environment).Execute()



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
    environment := *openapiclient.NewEnvironment("Url_example", "Id_example", "Name_example", []string{"Children_example"}, "TODO", time.Now(), time.Now()) // Environment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsUpdate(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsUpdate`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**Environment**](Environment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

