# \ServiceaccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceaccountsCreate**](ServiceaccountsApi.md#ServiceaccountsCreate) | **Post** /api/v1/serviceaccounts/ | Create a ServiceAccount user.
[**ServiceaccountsDestroy**](ServiceaccountsApi.md#ServiceaccountsDestroy) | **Delete** /api/v1/serviceaccounts/{id}/ | 
[**ServiceaccountsList**](ServiceaccountsApi.md#ServiceaccountsList) | **Get** /api/v1/serviceaccounts/ | 
[**ServiceaccountsPartialUpdate**](ServiceaccountsApi.md#ServiceaccountsPartialUpdate) | **Patch** /api/v1/serviceaccounts/{id}/ | 
[**ServiceaccountsRetrieve**](ServiceaccountsApi.md#ServiceaccountsRetrieve) | **Get** /api/v1/serviceaccounts/{id}/ | 
[**ServiceaccountsUpdate**](ServiceaccountsApi.md#ServiceaccountsUpdate) | **Put** /api/v1/serviceaccounts/{id}/ | 



## ServiceaccountsCreate

> ServiceAccountCreateResponse ServiceaccountsCreate(ctx).ServiceAccountCreateRequest(serviceAccountCreateRequest).Execute()

Create a ServiceAccount user.



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
    serviceAccountCreateRequest := *openapiclient.NewServiceAccountCreateRequest("Name_example") // ServiceAccountCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceaccountsApi.ServiceaccountsCreate(context.Background()).ServiceAccountCreateRequest(serviceAccountCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ServiceaccountsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceaccountsCreate`: ServiceAccountCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.ServiceaccountsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceaccountsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountCreateRequest** | [**ServiceAccountCreateRequest**](ServiceAccountCreateRequest.md) |  | 

### Return type

[**ServiceAccountCreateResponse**](ServiceAccountCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceaccountsDestroy

> ServiceaccountsDestroy(ctx, id).Execute()



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
    id := "id_example" // string | A unique value identifying this service account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceaccountsApi.ServiceaccountsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ServiceaccountsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique value identifying this service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceaccountsDestroyRequest struct via the builder pattern


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


## ServiceaccountsList

> PaginatedServiceAccountList ServiceaccountsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceaccountsApi.ServiceaccountsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ServiceaccountsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceaccountsList`: PaginatedServiceAccountList
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.ServiceaccountsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceaccountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedServiceAccountList**](PaginatedServiceAccountList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceaccountsPartialUpdate

> ServiceAccount ServiceaccountsPartialUpdate(ctx, id).PatchedServiceAccount(patchedServiceAccount).Execute()



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
    id := "id_example" // string | A unique value identifying this service account.
    patchedServiceAccount := *openapiclient.NewPatchedServiceAccount() // PatchedServiceAccount |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceaccountsApi.ServiceaccountsPartialUpdate(context.Background(), id).PatchedServiceAccount(patchedServiceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ServiceaccountsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceaccountsPartialUpdate`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.ServiceaccountsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique value identifying this service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceaccountsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedServiceAccount** | [**PatchedServiceAccount**](PatchedServiceAccount.md) |  | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceaccountsRetrieve

> ServiceAccount ServiceaccountsRetrieve(ctx, id).Execute()



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
    id := "id_example" // string | A unique value identifying this service account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceaccountsApi.ServiceaccountsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ServiceaccountsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceaccountsRetrieve`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.ServiceaccountsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique value identifying this service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceaccountsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceaccountsUpdate

> ServiceAccount ServiceaccountsUpdate(ctx, id).ServiceAccount(serviceAccount).Execute()



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
    id := "id_example" // string | A unique value identifying this service account.
    serviceAccount := *openapiclient.NewServiceAccount("Url_example", "Id_example", *openapiclient.NewPatchedServiceAccountUser("Url_example", "Id_example", "Name_example", "OrganizationName_example", "MembershipId_example", "Role_example", "Email_example", "PictureUrl_example", time.Now(), time.Now()), time.Now(), time.Now(), time.Now()) // ServiceAccount |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceaccountsApi.ServiceaccountsUpdate(context.Background(), id).ServiceAccount(serviceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ServiceaccountsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceaccountsUpdate`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.ServiceaccountsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique value identifying this service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceaccountsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccount** | [**ServiceAccount**](ServiceAccount.md) |  | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

