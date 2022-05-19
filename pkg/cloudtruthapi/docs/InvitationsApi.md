# \InvitationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationsAcceptCreate**](InvitationsApi.md#InvitationsAcceptCreate) | **Post** /api/v1/invitations/{id}/accept/ | Accept an invitation.
[**InvitationsCreate**](InvitationsApi.md#InvitationsCreate) | **Post** /api/v1/invitations/ | Create an invitation.
[**InvitationsDestroy**](InvitationsApi.md#InvitationsDestroy) | **Delete** /api/v1/invitations/{id}/ | 
[**InvitationsList**](InvitationsApi.md#InvitationsList) | **Get** /api/v1/invitations/ | 
[**InvitationsPartialUpdate**](InvitationsApi.md#InvitationsPartialUpdate) | **Patch** /api/v1/invitations/{id}/ | 
[**InvitationsResendCreate**](InvitationsApi.md#InvitationsResendCreate) | **Post** /api/v1/invitations/{id}/resend/ | Resend an invitation.
[**InvitationsRetrieve**](InvitationsApi.md#InvitationsRetrieve) | **Get** /api/v1/invitations/{id}/ | 
[**InvitationsUpdate**](InvitationsApi.md#InvitationsUpdate) | **Put** /api/v1/invitations/{id}/ | 



## InvitationsAcceptCreate

> Invitation InvitationsAcceptCreate(ctx, id).Execute()

Accept an invitation.



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The invitation ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.InvitationsAcceptCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsAcceptCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsAcceptCreate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsAcceptCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The invitation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsAcceptCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invitation**](Invitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsCreate

> Invitation InvitationsCreate(ctx).InvitationCreate(invitationCreate).Execute()

Create an invitation.



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
    invitationCreate := *openapiclient.NewInvitationCreate("Email_example", "TODO") // InvitationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.InvitationsCreate(context.Background()).InvitationCreate(invitationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsCreate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationCreate** | [**InvitationCreate**](InvitationCreate.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsDestroy

> InvitationsDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.InvitationsApi.InvitationsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInvitationsDestroyRequest struct via the builder pattern


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


## InvitationsList

> PaginatedInvitationList InvitationsList(ctx).Email(email).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).State(state).Execute()



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
    email := "email_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    role := "role_example" // string | The role that the user will have in the organization, should the user accept. (optional)
    state := "state_example" // string | The current state of the invitation. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.InvitationsList(context.Background()).Email(email).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsList`: PaginatedInvitationList
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **role** | **string** | The role that the user will have in the organization, should the user accept. | 
 **state** | **string** | The current state of the invitation. | 

### Return type

[**PaginatedInvitationList**](PaginatedInvitationList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsPartialUpdate

> Invitation InvitationsPartialUpdate(ctx, id).PatchedInvitation(patchedInvitation).Execute()



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
    patchedInvitation := *openapiclient.NewPatchedInvitation() // PatchedInvitation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.InvitationsPartialUpdate(context.Background(), id).PatchedInvitation(patchedInvitation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsPartialUpdate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedInvitation** | [**PatchedInvitation**](PatchedInvitation.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsResendCreate

> Invitation InvitationsResendCreate(ctx, id).Execute()

Resend an invitation.



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The invitation ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.InvitationsResendCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsResendCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsResendCreate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsResendCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The invitation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsResendCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invitation**](Invitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsRetrieve

> Invitation InvitationsRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.InvitationsApi.InvitationsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsRetrieve`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invitation**](Invitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsUpdate

> Invitation InvitationsUpdate(ctx, id).Invitation(invitation).Execute()



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
    invitation := *openapiclient.NewInvitation("Url_example", "Id_example", "Email_example", "TODO", "Inviter_example", "InviterName_example", "State_example", "StateDetail_example", "Membership_example", "Organization_example") // Invitation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsApi.InvitationsUpdate(context.Background(), id).Invitation(invitation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.InvitationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsUpdate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.InvitationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitation** | [**Invitation**](Invitation.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

