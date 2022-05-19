# \MembershipsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MembershipsCreate**](MembershipsApi.md#MembershipsCreate) | **Post** /api/v1/memberships/ | 
[**MembershipsDestroy**](MembershipsApi.md#MembershipsDestroy) | **Delete** /api/v1/memberships/{id}/ | 
[**MembershipsList**](MembershipsApi.md#MembershipsList) | **Get** /api/v1/memberships/ | 
[**MembershipsPartialUpdate**](MembershipsApi.md#MembershipsPartialUpdate) | **Patch** /api/v1/memberships/{id}/ | 
[**MembershipsRetrieve**](MembershipsApi.md#MembershipsRetrieve) | **Get** /api/v1/memberships/{id}/ | 
[**MembershipsUpdate**](MembershipsApi.md#MembershipsUpdate) | **Put** /api/v1/memberships/{id}/ | 



## MembershipsCreate

> Membership MembershipsCreate(ctx).MembershipCreate(membershipCreate).Execute()



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
    membershipCreate := *openapiclient.NewMembershipCreate("User_example", "TODO") // MembershipCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.MembershipsCreate(context.Background()).MembershipCreate(membershipCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.MembershipsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MembershipsCreate`: Membership
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.MembershipsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMembershipsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membershipCreate** | [**MembershipCreate**](MembershipCreate.md) |  | 

### Return type

[**Membership**](Membership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembershipsDestroy

> MembershipsDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.MembershipsApi.MembershipsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.MembershipsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMembershipsDestroyRequest struct via the builder pattern


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


## MembershipsList

> PaginatedMembershipList MembershipsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).User(user).Execute()



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
    role := "role_example" // string | The role that the user has in the organization. (optional)
    user := "user_example" // string | The unique identifier of a user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.MembershipsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.MembershipsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MembershipsList`: PaginatedMembershipList
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.MembershipsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMembershipsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **role** | **string** | The role that the user has in the organization. | 
 **user** | **string** | The unique identifier of a user. | 

### Return type

[**PaginatedMembershipList**](PaginatedMembershipList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembershipsPartialUpdate

> Membership MembershipsPartialUpdate(ctx, id).PatchedMembership(patchedMembership).Execute()



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
    patchedMembership := *openapiclient.NewPatchedMembership() // PatchedMembership |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.MembershipsPartialUpdate(context.Background(), id).PatchedMembership(patchedMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.MembershipsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MembershipsPartialUpdate`: Membership
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.MembershipsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembershipsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedMembership** | [**PatchedMembership**](PatchedMembership.md) |  | 

### Return type

[**Membership**](Membership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembershipsRetrieve

> Membership MembershipsRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.MembershipsApi.MembershipsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.MembershipsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MembershipsRetrieve`: Membership
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.MembershipsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembershipsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Membership**](Membership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembershipsUpdate

> Membership MembershipsUpdate(ctx, id).Membership(membership).Execute()



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
    membership := *openapiclient.NewMembership("Url_example", "Id_example", "User_example", "Organization_example", "TODO", time.Now(), time.Now()) // Membership | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembershipsApi.MembershipsUpdate(context.Background(), id).Membership(membership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembershipsApi.MembershipsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MembershipsUpdate`: Membership
    fmt.Fprintf(os.Stdout, "Response from `MembershipsApi.MembershipsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembershipsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membership** | [**Membership**](Membership.md) |  | 

### Return type

[**Membership**](Membership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

