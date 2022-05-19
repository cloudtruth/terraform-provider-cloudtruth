# \TypesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TypesCreate**](TypesApi.md#TypesCreate) | **Post** /api/v1/types/ | 
[**TypesDestroy**](TypesApi.md#TypesDestroy) | **Delete** /api/v1/types/{id}/ | 
[**TypesList**](TypesApi.md#TypesList) | **Get** /api/v1/types/ | 
[**TypesPartialUpdate**](TypesApi.md#TypesPartialUpdate) | **Patch** /api/v1/types/{id}/ | 
[**TypesRetrieve**](TypesApi.md#TypesRetrieve) | **Get** /api/v1/types/{id}/ | 
[**TypesRulesCreate**](TypesApi.md#TypesRulesCreate) | **Post** /api/v1/types/{parametertype_pk}/rules/ | 
[**TypesRulesDestroy**](TypesApi.md#TypesRulesDestroy) | **Delete** /api/v1/types/{parametertype_pk}/rules/{id}/ | 
[**TypesRulesList**](TypesApi.md#TypesRulesList) | **Get** /api/v1/types/{parametertype_pk}/rules/ | 
[**TypesRulesPartialUpdate**](TypesApi.md#TypesRulesPartialUpdate) | **Patch** /api/v1/types/{parametertype_pk}/rules/{id}/ | 
[**TypesRulesRetrieve**](TypesApi.md#TypesRulesRetrieve) | **Get** /api/v1/types/{parametertype_pk}/rules/{id}/ | 
[**TypesRulesUpdate**](TypesApi.md#TypesRulesUpdate) | **Put** /api/v1/types/{parametertype_pk}/rules/{id}/ | 
[**TypesUpdate**](TypesApi.md#TypesUpdate) | **Put** /api/v1/types/{id}/ | 



## TypesCreate

> ParameterType TypesCreate(ctx).ParameterTypeCreate(parameterTypeCreate).Execute()



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
    parameterTypeCreate := *openapiclient.NewParameterTypeCreate("Name_example") // ParameterTypeCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesCreate(context.Background()).ParameterTypeCreate(parameterTypeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesCreate`: ParameterType
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTypesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parameterTypeCreate** | [**ParameterTypeCreate**](ParameterTypeCreate.md) |  | 

### Return type

[**ParameterType**](ParameterType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesDestroy

> TypesDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.TypesApi.TypesDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTypesDestroyRequest struct via the builder pattern


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


## TypesList

> PaginatedParameterTypeList TypesList(ctx).DescriptionIcontains(descriptionIcontains).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).ParentNameIcontains(parentNameIcontains).ParentNameIexact(parentNameIexact).Execute()



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
    nameIcontains := "nameIcontains_example" // string |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    parentNameIcontains := "parentNameIcontains_example" // string |  (optional)
    parentNameIexact := "parentNameIexact_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesList(context.Background()).DescriptionIcontains(descriptionIcontains).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).ParentNameIcontains(parentNameIcontains).ParentNameIexact(parentNameIexact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesList`: PaginatedParameterTypeList
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descriptionIcontains** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **parentNameIcontains** | **string** |  | 
 **parentNameIexact** | **string** |  | 

### Return type

[**PaginatedParameterTypeList**](PaginatedParameterTypeList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesPartialUpdate

> ParameterType TypesPartialUpdate(ctx, id).PatchedParameterType(patchedParameterType).Execute()



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
    patchedParameterType := *openapiclient.NewPatchedParameterType() // PatchedParameterType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesPartialUpdate(context.Background(), id).PatchedParameterType(patchedParameterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesPartialUpdate`: ParameterType
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedParameterType** | [**PatchedParameterType**](PatchedParameterType.md) |  | 

### Return type

[**ParameterType**](ParameterType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesRetrieve

> ParameterType TypesRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.TypesApi.TypesRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesRetrieve`: ParameterType
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParameterType**](ParameterType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesRulesCreate

> ParameterTypeRule TypesRulesCreate(ctx, parametertypePk).ParameterTypeRuleCreate(parameterTypeRuleCreate).Execute()



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
    parametertypePk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter type id.
    parameterTypeRuleCreate := *openapiclient.NewParameterTypeRuleCreate(openapiclient.ParameterRuleTypeEnum("min"), "Constraint_example") // ParameterTypeRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesRulesCreate(context.Background(), parametertypePk).ParameterTypeRuleCreate(parameterTypeRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesRulesCreate`: ParameterTypeRule
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parametertypePk** | **string** | The parameter type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameterTypeRuleCreate** | [**ParameterTypeRuleCreate**](ParameterTypeRuleCreate.md) |  | 

### Return type

[**ParameterTypeRule**](ParameterTypeRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesRulesDestroy

> TypesRulesDestroy(ctx, id, parametertypePk).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter type rule.
    parametertypePk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter type id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesRulesDestroy(context.Background(), id, parametertypePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter type rule. | 
**parametertypePk** | **string** | The parameter type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRulesDestroyRequest struct via the builder pattern


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


## TypesRulesList

> PaginatedParameterTypeRuleList TypesRulesList(ctx, parametertypePk).Ordering(ordering).Page(page).PageSize(pageSize).Type_(type_).Execute()



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
    parametertypePk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter type id.
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesRulesList(context.Background(), parametertypePk).Ordering(ordering).Page(page).PageSize(pageSize).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesRulesList`: PaginatedParameterTypeRuleList
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parametertypePk** | **string** | The parameter type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **type_** | **string** |  | 

### Return type

[**PaginatedParameterTypeRuleList**](PaginatedParameterTypeRuleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesRulesPartialUpdate

> ParameterTypeRule TypesRulesPartialUpdate(ctx, id, parametertypePk).PatchedParameterTypeRule(patchedParameterTypeRule).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter type rule.
    parametertypePk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter type id.
    patchedParameterTypeRule := *openapiclient.NewPatchedParameterTypeRule() // PatchedParameterTypeRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesRulesPartialUpdate(context.Background(), id, parametertypePk).PatchedParameterTypeRule(patchedParameterTypeRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesRulesPartialUpdate`: ParameterTypeRule
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter type rule. | 
**parametertypePk** | **string** | The parameter type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedParameterTypeRule** | [**PatchedParameterTypeRule**](PatchedParameterTypeRule.md) |  | 

### Return type

[**ParameterTypeRule**](ParameterTypeRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesRulesRetrieve

> ParameterTypeRule TypesRulesRetrieve(ctx, id, parametertypePk).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter type rule.
    parametertypePk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter type id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesRulesRetrieve(context.Background(), id, parametertypePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRulesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesRulesRetrieve`: ParameterTypeRule
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter type rule. | 
**parametertypePk** | **string** | The parameter type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParameterTypeRule**](ParameterTypeRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesRulesUpdate

> ParameterTypeRule TypesRulesUpdate(ctx, id, parametertypePk).ParameterTypeRule(parameterTypeRule).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter type rule.
    parametertypePk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter type id.
    parameterTypeRule := *openapiclient.NewParameterTypeRule("Url_example", "Id_example", "ParameterType_example", openapiclient.ParameterRuleTypeEnum("min"), "Constraint_example", time.Now(), time.Now()) // ParameterTypeRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesRulesUpdate(context.Background(), id, parametertypePk).ParameterTypeRule(parameterTypeRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesRulesUpdate`: ParameterTypeRule
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter type rule. | 
**parametertypePk** | **string** | The parameter type id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parameterTypeRule** | [**ParameterTypeRule**](ParameterTypeRule.md) |  | 

### Return type

[**ParameterTypeRule**](ParameterTypeRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TypesUpdate

> ParameterType TypesUpdate(ctx, id).ParameterType(parameterType).Execute()



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
    parameterType := *openapiclient.NewParameterType("Url_example", "Id_example", "Name_example", []openapiclient.ParameterTypeRule{*openapiclient.NewParameterTypeRule("Url_example", "Id_example", "ParameterType_example", openapiclient.ParameterRuleTypeEnum("min"), "Constraint_example", time.Now(), time.Now())}, "ParentName_example", time.Now(), time.Now()) // ParameterType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TypesApi.TypesUpdate(context.Background(), id).ParameterType(parameterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.TypesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TypesUpdate`: ParameterType
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.TypesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTypesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameterType** | [**ParameterType**](ParameterType.md) |  | 

### Return type

[**ParameterType**](ParameterType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

