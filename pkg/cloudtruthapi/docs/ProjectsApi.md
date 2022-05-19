# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsCreate**](ProjectsApi.md#ProjectsCreate) | **Post** /api/v1/projects/ | 
[**ProjectsDestroy**](ProjectsApi.md#ProjectsDestroy) | **Delete** /api/v1/projects/{id}/ | 
[**ProjectsList**](ProjectsApi.md#ProjectsList) | **Get** /api/v1/projects/ | 
[**ProjectsParameterExportList**](ProjectsApi.md#ProjectsParameterExportList) | **Get** /api/v1/projects/{project_pk}/parameter-export/ | 
[**ProjectsParametersCreate**](ProjectsApi.md#ProjectsParametersCreate) | **Post** /api/v1/projects/{project_pk}/parameters/ | 
[**ProjectsParametersDestroy**](ProjectsApi.md#ProjectsParametersDestroy) | **Delete** /api/v1/projects/{project_pk}/parameters/{id}/ | 
[**ProjectsParametersDualityList**](ProjectsApi.md#ProjectsParametersDualityList) | **Get** /api/v1/projects/{project_pk}/parameters/duality/ | 
[**ProjectsParametersList**](ProjectsApi.md#ProjectsParametersList) | **Get** /api/v1/projects/{project_pk}/parameters/ | 
[**ProjectsParametersPartialUpdate**](ProjectsApi.md#ProjectsParametersPartialUpdate) | **Patch** /api/v1/projects/{project_pk}/parameters/{id}/ | 
[**ProjectsParametersPushesList**](ProjectsApi.md#ProjectsParametersPushesList) | **Get** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/pushes/ | List push operations.
[**ProjectsParametersRetrieve**](ProjectsApi.md#ProjectsParametersRetrieve) | **Get** /api/v1/projects/{project_pk}/parameters/{id}/ | 
[**ProjectsParametersRulesCreate**](ProjectsApi.md#ProjectsParametersRulesCreate) | **Post** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/rules/ | 
[**ProjectsParametersRulesDestroy**](ProjectsApi.md#ProjectsParametersRulesDestroy) | **Delete** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/rules/{id}/ | 
[**ProjectsParametersRulesList**](ProjectsApi.md#ProjectsParametersRulesList) | **Get** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/rules/ | 
[**ProjectsParametersRulesPartialUpdate**](ProjectsApi.md#ProjectsParametersRulesPartialUpdate) | **Patch** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/rules/{id}/ | 
[**ProjectsParametersRulesRetrieve**](ProjectsApi.md#ProjectsParametersRulesRetrieve) | **Get** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/rules/{id}/ | 
[**ProjectsParametersRulesUpdate**](ProjectsApi.md#ProjectsParametersRulesUpdate) | **Put** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/rules/{id}/ | 
[**ProjectsParametersTimelineRetrieve**](ProjectsApi.md#ProjectsParametersTimelineRetrieve) | **Get** /api/v1/projects/{project_pk}/parameters/{id}/timeline/ | 
[**ProjectsParametersTimelinesRetrieve**](ProjectsApi.md#ProjectsParametersTimelinesRetrieve) | **Get** /api/v1/projects/{project_pk}/parameters/timelines/ | 
[**ProjectsParametersUpdate**](ProjectsApi.md#ProjectsParametersUpdate) | **Put** /api/v1/projects/{project_pk}/parameters/{id}/ | 
[**ProjectsParametersValuesCreate**](ProjectsApi.md#ProjectsParametersValuesCreate) | **Post** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/values/ | Set a value.
[**ProjectsParametersValuesDestroy**](ProjectsApi.md#ProjectsParametersValuesDestroy) | **Delete** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/values/{id}/ | Destroy a value.
[**ProjectsParametersValuesList**](ProjectsApi.md#ProjectsParametersValuesList) | **Get** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/values/ | Retrieve values.
[**ProjectsParametersValuesPartialUpdate**](ProjectsApi.md#ProjectsParametersValuesPartialUpdate) | **Patch** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/values/{id}/ | Update a value.
[**ProjectsParametersValuesRetrieve**](ProjectsApi.md#ProjectsParametersValuesRetrieve) | **Get** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/values/{id}/ | Retrieve a value.
[**ProjectsParametersValuesUpdate**](ProjectsApi.md#ProjectsParametersValuesUpdate) | **Put** /api/v1/projects/{project_pk}/parameters/{parameter_pk}/values/{id}/ | Update a value.
[**ProjectsPartialUpdate**](ProjectsApi.md#ProjectsPartialUpdate) | **Patch** /api/v1/projects/{id}/ | 
[**ProjectsRetrieve**](ProjectsApi.md#ProjectsRetrieve) | **Get** /api/v1/projects/{id}/ | 
[**ProjectsTemplatePreviewCreate**](ProjectsApi.md#ProjectsTemplatePreviewCreate) | **Post** /api/v1/projects/{project_pk}/template-preview/ | 
[**ProjectsTemplatesCreate**](ProjectsApi.md#ProjectsTemplatesCreate) | **Post** /api/v1/projects/{project_pk}/templates/ | 
[**ProjectsTemplatesDestroy**](ProjectsApi.md#ProjectsTemplatesDestroy) | **Delete** /api/v1/projects/{project_pk}/templates/{id}/ | 
[**ProjectsTemplatesList**](ProjectsApi.md#ProjectsTemplatesList) | **Get** /api/v1/projects/{project_pk}/templates/ | 
[**ProjectsTemplatesPartialUpdate**](ProjectsApi.md#ProjectsTemplatesPartialUpdate) | **Patch** /api/v1/projects/{project_pk}/templates/{id}/ | 
[**ProjectsTemplatesRetrieve**](ProjectsApi.md#ProjectsTemplatesRetrieve) | **Get** /api/v1/projects/{project_pk}/templates/{id}/ | 
[**ProjectsTemplatesTimelineRetrieve**](ProjectsApi.md#ProjectsTemplatesTimelineRetrieve) | **Get** /api/v1/projects/{project_pk}/templates/{id}/timeline/ | 
[**ProjectsTemplatesTimelinesRetrieve**](ProjectsApi.md#ProjectsTemplatesTimelinesRetrieve) | **Get** /api/v1/projects/{project_pk}/templates/timelines/ | 
[**ProjectsTemplatesUpdate**](ProjectsApi.md#ProjectsTemplatesUpdate) | **Put** /api/v1/projects/{project_pk}/templates/{id}/ | 
[**ProjectsUpdate**](ProjectsApi.md#ProjectsUpdate) | **Put** /api/v1/projects/{id}/ | 



## ProjectsCreate

> Project ProjectsCreate(ctx).ProjectCreate(projectCreate).Execute()



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
    projectCreate := *openapiclient.NewProjectCreate("Name_example") // ProjectCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreate(context.Background()).ProjectCreate(projectCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreate** | [**ProjectCreate**](ProjectCreate.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDestroy

> ProjectsDestroy(ctx, id).Execute()



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
    resp, r, err := apiClient.ProjectsApi.ProjectsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsDestroyRequest struct via the builder pattern


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


## ProjectsList

> PaginatedProjectList ProjectsList(ctx).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsList(context.Background()).DescriptionIcontains(descriptionIcontains).Name(name).NameIcontains(nameIcontains).Ordering(ordering).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsList`: PaginatedProjectList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descriptionIcontains** | **string** |  | 
 **name** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedProjectList**](PaginatedProjectList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParameterExportList

> ParameterExport ProjectsParameterExportList(ctx, projectPk).AsOf(asOf).Contains(contains).Endswith(endswith).Environment(environment).ExplicitExport(explicitExport).MaskSecrets(maskSecrets).Ordering(ordering).Output(output).Startswith(startswith).Tag(tag).Wrap(wrap).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    contains := "contains_example" // string | Only include parameters with names that contain the provided string. (optional)
    endswith := "endswith_example" // string | Only include parameters with names that end with the provided string. (optional)
    environment := "environment_example" // string | Name or id of the environment to use to retrieve parameter values. (optional)
    explicitExport := true // bool | Explicitly marks parameters with export, e.g. `export FOO=bar`. (optional) (default to false)
    maskSecrets := true // bool | Masks all secrets in the template with `*****`. (optional) (default to false)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    output := "output_example" // string | Format to output: One of 'docker', 'dotenv', 'shell'. (optional)
    startswith := "startswith_example" // string | Only include parameters with names that start with the provided string. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)
    wrap := true // bool | Indicates all secrets are wrapped. For more information on secret wrapping, see the documentation. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParameterExportList(context.Background(), projectPk).AsOf(asOf).Contains(contains).Endswith(endswith).Environment(environment).ExplicitExport(explicitExport).MaskSecrets(maskSecrets).Ordering(ordering).Output(output).Startswith(startswith).Tag(tag).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParameterExportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParameterExportList`: ParameterExport
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParameterExportList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParameterExportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **contains** | **string** | Only include parameters with names that contain the provided string. | 
 **endswith** | **string** | Only include parameters with names that end with the provided string. | 
 **environment** | **string** | Name or id of the environment to use to retrieve parameter values. | 
 **explicitExport** | **bool** | Explicitly marks parameters with export, e.g. &#x60;export FOO&#x3D;bar&#x60;. | [default to false]
 **maskSecrets** | **bool** | Masks all secrets in the template with &#x60;*****&#x60;. | [default to false]
 **ordering** | **string** | Which field to use when ordering the results. | 
 **output** | **string** | Format to output: One of &#39;docker&#39;, &#39;dotenv&#39;, &#39;shell&#39;. | 
 **startswith** | **string** | Only include parameters with names that start with the provided string. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **wrap** | **bool** | Indicates all secrets are wrapped. For more information on secret wrapping, see the documentation. | [default to false]

### Return type

[**ParameterExport**](ParameterExport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersCreate

> Parameter ProjectsParametersCreate(ctx, projectPk).ParameterCreate(parameterCreate).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    parameterCreate := *openapiclient.NewParameterCreate("Name_example") // ParameterCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersCreate(context.Background(), projectPk).ParameterCreate(parameterCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersCreate`: Parameter
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameterCreate** | [**ParameterCreate**](ParameterCreate.md) |  | 

### Return type

[**Parameter**](Parameter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersDestroy

> ProjectsParametersDestroy(ctx, id, projectPk).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersDestroy(context.Background(), id, projectPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersDestroyRequest struct via the builder pattern


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


## ProjectsParametersDualityList

> PaginatedParameterDualityList ProjectsParametersDualityList(ctx, projectPk).Difference(difference).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIstartswith(nameIstartswith).NameStartswith(nameStartswith).Ordering(ordering).Page(page).PageSize(pageSize).Secret(secret).T1AsOf(t1AsOf).T1Tag(t1Tag).T2AsOf(t2AsOf).T2Tag(t2Tag).Values(values).Wrap(wrap).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    difference := true // bool | Compare the values at `t1` and `t2` and if they are the same, omit the parameter from the result. (optional)
    environment := "environment_example" // string | Name or id (uuid) of the environment to get parameter values for. Cannot be used with `values`. (optional)
    evaluate := true // bool | If `true`, runs template evaluation on this parameter's values.  If `false`, returns the value's template. Has no effect on values that are not interpolated. (optional) (default to true)
    maskSecrets := true // bool | If true, replaces all secrets with `*****`. (optional) (default to false)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    nameIstartswith := "nameIstartswith_example" // string |  (optional)
    nameStartswith := "nameStartswith_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    secret := true // bool |  (optional)
    t1AsOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `t1_tag`. (optional)
    t1Tag := "t1Tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `t1_as_of`. Requires `environment`. (optional)
    t2AsOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `t2_tag`. (optional)
    t2Tag := "t2Tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `t2_as_of`. Requires `environment`. (optional)
    values := true // bool | If false, values are not returned: the `values` array will have no entries. This speeds up retrieval if value content is not needed. Cannot be used with `environment`. (optional) (default to true)
    wrap := true // bool | Wrap secrets. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersDualityList(context.Background(), projectPk).Difference(difference).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIstartswith(nameIstartswith).NameStartswith(nameStartswith).Ordering(ordering).Page(page).PageSize(pageSize).Secret(secret).T1AsOf(t1AsOf).T1Tag(t1Tag).T2AsOf(t2AsOf).T2Tag(t2Tag).Values(values).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersDualityList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersDualityList`: PaginatedParameterDualityList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersDualityList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersDualityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **difference** | **bool** | Compare the values at &#x60;t1&#x60; and &#x60;t2&#x60; and if they are the same, omit the parameter from the result. | 
 **environment** | **string** | Name or id (uuid) of the environment to get parameter values for. Cannot be used with &#x60;values&#x60;. | 
 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter&#39;s values.  If &#x60;false&#x60;, returns the value&#39;s template. Has no effect on values that are not interpolated. | [default to true]
 **maskSecrets** | **bool** | If true, replaces all secrets with &#x60;*****&#x60;. | [default to false]
 **name** | **string** |  | 
 **nameContains** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **nameIstartswith** | **string** |  | 
 **nameStartswith** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **secret** | **bool** |  | 
 **t1AsOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;t1_tag&#x60;. | 
 **t1Tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;t1_as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **t2AsOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;t2_tag&#x60;. | 
 **t2Tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;t2_as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **values** | **bool** | If false, values are not returned: the &#x60;values&#x60; array will have no entries. This speeds up retrieval if value content is not needed. Cannot be used with &#x60;environment&#x60;. | [default to true]
 **wrap** | **bool** | Wrap secrets. | [default to false]

### Return type

[**PaginatedParameterDualityList**](PaginatedParameterDualityList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersList

> PaginatedParameterList ProjectsParametersList(ctx, projectPk).AsOf(asOf).DescriptionIcontains(descriptionIcontains).Difference(difference).Environment(environment).Evaluate(evaluate).IdIn(idIn).MaskSecrets(maskSecrets).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIstartswith(nameIstartswith).NameStartswith(nameStartswith).Ordering(ordering).Page(page).PageSize(pageSize).ProjectName(projectName).ProjectNameContains(projectNameContains).ProjectNameIcontains(projectNameIcontains).ProjectNameIexact(projectNameIexact).ProjectNameIstartswith(projectNameIstartswith).ProjectNameStartswith(projectNameStartswith).Secret(secret).Tag(tag).Values(values).Wrap(wrap).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    descriptionIcontains := "descriptionIcontains_example" // string |  (optional)
    difference := "difference_example" // string | Specify a list of comma-separated environment names or ids (uuid) to compare values. Only returns a parameter if there is a difference. Cannot be used with `environment`, `wrap` or `values`. If used with `mask_secrets` then no secret parameters will be included in the result. (optional)
    environment := "environment_example" // string | Name or id (uuid) of the environment to get parameter values for. Cannot be used with `values`. (optional)
    evaluate := true // bool | If `true`, runs template evaluation on this parameter's values.  If `false`, returns the value's template. Has no effect on values that are not interpolated. (optional) (default to true)
    idIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    maskSecrets := true // bool | If true, replaces all secrets with `*****`. (optional) (default to false)
    name := "name_example" // string |  (optional)
    nameContains := "nameContains_example" // string |  (optional)
    nameIcontains := "nameIcontains_example" // string |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    nameIstartswith := "nameIstartswith_example" // string |  (optional)
    nameStartswith := "nameStartswith_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    projectName := "projectName_example" // string |  (optional)
    projectNameContains := "projectNameContains_example" // string |  (optional)
    projectNameIcontains := "projectNameIcontains_example" // string |  (optional)
    projectNameIexact := "projectNameIexact_example" // string |  (optional)
    projectNameIstartswith := "projectNameIstartswith_example" // string |  (optional)
    projectNameStartswith := "projectNameStartswith_example" // string |  (optional)
    secret := true // bool |  (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)
    values := true // bool | If false, values are not returned: the `values` array will have no entries. This speeds up retrieval if value content is not needed. Cannot be used with `environment`. (optional) (default to true)
    wrap := true // bool | Wrap secrets. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersList(context.Background(), projectPk).AsOf(asOf).DescriptionIcontains(descriptionIcontains).Difference(difference).Environment(environment).Evaluate(evaluate).IdIn(idIn).MaskSecrets(maskSecrets).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIstartswith(nameIstartswith).NameStartswith(nameStartswith).Ordering(ordering).Page(page).PageSize(pageSize).ProjectName(projectName).ProjectNameContains(projectNameContains).ProjectNameIcontains(projectNameIcontains).ProjectNameIexact(projectNameIexact).ProjectNameIstartswith(projectNameIstartswith).ProjectNameStartswith(projectNameStartswith).Secret(secret).Tag(tag).Values(values).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersList`: PaginatedParameterList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **descriptionIcontains** | **string** |  | 
 **difference** | **string** | Specify a list of comma-separated environment names or ids (uuid) to compare values. Only returns a parameter if there is a difference. Cannot be used with &#x60;environment&#x60;, &#x60;wrap&#x60; or &#x60;values&#x60;. If used with &#x60;mask_secrets&#x60; then no secret parameters will be included in the result. | 
 **environment** | **string** | Name or id (uuid) of the environment to get parameter values for. Cannot be used with &#x60;values&#x60;. | 
 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter&#39;s values.  If &#x60;false&#x60;, returns the value&#39;s template. Has no effect on values that are not interpolated. | [default to true]
 **idIn** | **[]string** | Multiple values may be separated by commas. | 
 **maskSecrets** | **bool** | If true, replaces all secrets with &#x60;*****&#x60;. | [default to false]
 **name** | **string** |  | 
 **nameContains** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **nameIstartswith** | **string** |  | 
 **nameStartswith** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **projectName** | **string** |  | 
 **projectNameContains** | **string** |  | 
 **projectNameIcontains** | **string** |  | 
 **projectNameIexact** | **string** |  | 
 **projectNameIstartswith** | **string** |  | 
 **projectNameStartswith** | **string** |  | 
 **secret** | **bool** |  | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **values** | **bool** | If false, values are not returned: the &#x60;values&#x60; array will have no entries. This speeds up retrieval if value content is not needed. Cannot be used with &#x60;environment&#x60;. | [default to true]
 **wrap** | **bool** | Wrap secrets. | [default to false]

### Return type

[**PaginatedParameterList**](PaginatedParameterList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersPartialUpdate

> Parameter ProjectsParametersPartialUpdate(ctx, id, projectPk).PatchedParameter(patchedParameter).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedParameter := *openapiclient.NewPatchedParameter() // PatchedParameter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersPartialUpdate(context.Background(), id, projectPk).PatchedParameter(patchedParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersPartialUpdate`: Parameter
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedParameter** | [**PatchedParameter**](PatchedParameter.md) |  | 

### Return type

[**Parameter**](Parameter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersPushesList

> PaginatedTaskStepList ProjectsParametersPushesList(ctx, parameterPk, projectPk).AsOf(asOf).Ordering(ordering).Page(page).PageSize(pageSize).Tag(tag).Execute()

List push operations.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersPushesList(context.Background(), parameterPk, projectPk).AsOf(asOf).Ordering(ordering).Page(page).PageSize(pageSize).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersPushesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersPushesList`: PaginatedTaskStepList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersPushesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterPk** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersPushesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

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


## ProjectsParametersRetrieve

> Parameter ProjectsParametersRetrieve(ctx, id, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Tag(tag).Values(values).Wrap(wrap).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id (uuid) of the environment to get parameter values for. Cannot be used with `values`. (optional)
    evaluate := true // bool | If `true`, runs template evaluation on this parameter's values.  If `false`, returns the value's template. Has no effect on values that are not interpolated. (optional) (default to true)
    maskSecrets := true // bool | If true, replaces all secrets with `*****`. (optional) (default to false)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)
    values := true // bool | If false, values are not returned: the `values` array will have no entries. This speeds up retrieval if value content is not needed. Cannot be used with `environment`. (optional) (default to true)
    wrap := true // bool | Wrap secrets. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRetrieve(context.Background(), id, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Tag(tag).Values(values).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersRetrieve`: Parameter
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id (uuid) of the environment to get parameter values for. Cannot be used with &#x60;values&#x60;. | 
 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter&#39;s values.  If &#x60;false&#x60;, returns the value&#39;s template. Has no effect on values that are not interpolated. | [default to true]
 **maskSecrets** | **bool** | If true, replaces all secrets with &#x60;*****&#x60;. | [default to false]
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **values** | **bool** | If false, values are not returned: the &#x60;values&#x60; array will have no entries. This speeds up retrieval if value content is not needed. Cannot be used with &#x60;environment&#x60;. | [default to true]
 **wrap** | **bool** | Wrap secrets. | [default to false]

### Return type

[**Parameter**](Parameter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersRulesCreate

> ParameterRule ProjectsParametersRulesCreate(ctx, parameterPk, projectPk).ParameterRuleCreate(parameterRuleCreate).Execute()



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    parameterRuleCreate := *openapiclient.NewParameterRuleCreate(openapiclient.ParameterRuleTypeEnum("min"), "Constraint_example") // ParameterRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRulesCreate(context.Background(), parameterPk, projectPk).ParameterRuleCreate(parameterRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersRulesCreate`: ParameterRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parameterRuleCreate** | [**ParameterRuleCreate**](ParameterRuleCreate.md) |  | 

### Return type

[**ParameterRule**](ParameterRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersRulesDestroy

> ProjectsParametersRulesDestroy(ctx, id, parameterPk, projectPk).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter rule.
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRulesDestroy(context.Background(), id, parameterPk, projectPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter rule. | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRulesDestroyRequest struct via the builder pattern


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


## ProjectsParametersRulesList

> PaginatedParameterRuleList ProjectsParametersRulesList(ctx, parameterPk, projectPk).Ordering(ordering).Page(page).PageSize(pageSize).Type_(type_).Execute()



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRulesList(context.Background(), parameterPk, projectPk).Ordering(ordering).Page(page).PageSize(pageSize).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersRulesList`: PaginatedParameterRuleList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **type_** | **string** |  | 

### Return type

[**PaginatedParameterRuleList**](PaginatedParameterRuleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersRulesPartialUpdate

> ParameterRule ProjectsParametersRulesPartialUpdate(ctx, id, parameterPk, projectPk).PatchedParameterRule(patchedParameterRule).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter rule.
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    patchedParameterRule := *openapiclient.NewPatchedParameterRule() // PatchedParameterRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRulesPartialUpdate(context.Background(), id, parameterPk, projectPk).PatchedParameterRule(patchedParameterRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersRulesPartialUpdate`: ParameterRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter rule. | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedParameterRule** | [**PatchedParameterRule**](PatchedParameterRule.md) |  | 

### Return type

[**ParameterRule**](ParameterRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersRulesRetrieve

> ParameterRule ProjectsParametersRulesRetrieve(ctx, id, parameterPk, projectPk).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter rule.
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRulesRetrieve(context.Background(), id, parameterPk, projectPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRulesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersRulesRetrieve`: ParameterRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter rule. | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ParameterRule**](ParameterRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersRulesUpdate

> ParameterRule ProjectsParametersRulesUpdate(ctx, id, parameterPk, projectPk).ParameterRule(parameterRule).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this parameter rule.
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    parameterRule := *openapiclient.NewParameterRule("Url_example", "Id_example", "Parameter_example", openapiclient.ParameterRuleTypeEnum("min"), "Constraint_example", time.Now(), time.Now()) // ParameterRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersRulesUpdate(context.Background(), id, parameterPk, projectPk).ParameterRule(parameterRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersRulesUpdate`: ParameterRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this parameter rule. | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parameterRule** | [**ParameterRule**](ParameterRule.md) |  | 

### Return type

[**ParameterRule**](ParameterRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersTimelineRetrieve

> ParameterTimeline ProjectsParametersTimelineRetrieve(ctx, id, projectPk).AsOf(asOf).Tag(tag).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersTimelineRetrieve(context.Background(), id, projectPk).AsOf(asOf).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersTimelineRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersTimelineRetrieve`: ParameterTimeline
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersTimelineRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersTimelineRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

### Return type

[**ParameterTimeline**](ParameterTimeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersTimelinesRetrieve

> ParameterTimeline ProjectsParametersTimelinesRetrieve(ctx, projectPk).AsOf(asOf).Tag(tag).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersTimelinesRetrieve(context.Background(), projectPk).AsOf(asOf).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersTimelinesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersTimelinesRetrieve`: ParameterTimeline
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersTimelinesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersTimelinesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

### Return type

[**ParameterTimeline**](ParameterTimeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersUpdate

> Parameter ProjectsParametersUpdate(ctx, id, projectPk).Parameter(parameter).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    parameter := *openapiclient.NewParameter("Url_example", "Id_example", "Name_example", []openapiclient.ParameterRule{*openapiclient.NewParameterRule("Url_example", "Id_example", "Parameter_example", openapiclient.ParameterRuleTypeEnum("min"), "Constraint_example", time.Now(), time.Now())}, "Project_example", "ProjectName_example", []string{"ReferencingTemplates_example"}, []string{"ReferencingValues_example"}, map[string]ParameterValuesValue{"key": *openapiclient.NewParameterValuesValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now())}, "Overrides_example", time.Now(), time.Now()) // Parameter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersUpdate(context.Background(), id, projectPk).Parameter(parameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersUpdate`: Parameter
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parameter** | [**Parameter**](Parameter.md) |  | 

### Return type

[**Parameter**](Parameter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersValuesCreate

> Value ProjectsParametersValuesCreate(ctx, parameterPk, projectPk).ValueCreate(valueCreate).Evaluate(evaluate).Wrap(wrap).Execute()

Set a value.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    valueCreate := *openapiclient.NewValueCreate("Environment_example") // ValueCreate | 
    evaluate := true // bool | If `true`, runs template evaluation on this parameter.  If `false`, returns the value's template. No effect on values that are not interpolated. (optional) (default to true)
    wrap := true // bool | Indicates the `internal_value` is a wrapped secret. For more information on secret wrapping, see the documentation.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersValuesCreate(context.Background(), parameterPk, projectPk).ValueCreate(valueCreate).Evaluate(evaluate).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersValuesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersValuesCreate`: Value
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersValuesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersValuesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **valueCreate** | [**ValueCreate**](ValueCreate.md) |  | 
 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter.  If &#x60;false&#x60;, returns the value&#39;s template. No effect on values that are not interpolated. | [default to true]
 **wrap** | **bool** | Indicates the &#x60;internal_value&#x60; is a wrapped secret. For more information on secret wrapping, see the documentation.  | 

### Return type

[**Value**](Value.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersValuesDestroy

> ProjectsParametersValuesDestroy(ctx, id, parameterPk, projectPk).Evaluate(evaluate).Execute()

Destroy a value.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    evaluate := true // bool | If `true`, runs template evaluation on this parameter.  If `false`, returns the value's template. No effect on values that are not interpolated. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersValuesDestroy(context.Background(), id, parameterPk, projectPk).Evaluate(evaluate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersValuesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersValuesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter.  If &#x60;false&#x60;, returns the value&#39;s template. No effect on values that are not interpolated. | [default to true]

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


## ProjectsParametersValuesList

> PaginatedValueList ProjectsParametersValuesList(ctx, parameterPk, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).Exclude(exclude).Include(include).MaskSecrets(maskSecrets).Page(page).PageSize(pageSize).Tag(tag).Wrap(wrap).Execute()

Retrieve values.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id of the environment to limit the result to. If this is not specified then the result will contain a value for any environment in which it is set. You cannot use this option to retrieve the _effective_ value of a parameter in an environment for which is is not explicitly set. To see _effective_ values use the Parameters API (see the `values` field). (optional)
    evaluate := true // bool | If `true`, runs template evaluation on this parameter.  If `false`, returns the value's template. No effect on values that are not interpolated. (optional) (default to true)
    exclude := "exclude_example" // string | A comma-separated list of field names to exclude from the response. (optional)
    include := "include_example" // string | A comma-separated list of field names to include in the response. (optional)
    maskSecrets := true // bool | Mask secret values in responses with `*****`. (optional) (default to false)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)
    wrap := true // bool | For writes, indicates `internal_value` is wrapped; for reads, indicates `value` is wrapped. For more information on secret wrapping, see the documentation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersValuesList(context.Background(), parameterPk, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).Exclude(exclude).Include(include).MaskSecrets(maskSecrets).Page(page).PageSize(pageSize).Tag(tag).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersValuesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersValuesList`: PaginatedValueList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersValuesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersValuesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id of the environment to limit the result to. If this is not specified then the result will contain a value for any environment in which it is set. You cannot use this option to retrieve the _effective_ value of a parameter in an environment for which is is not explicitly set. To see _effective_ values use the Parameters API (see the &#x60;values&#x60; field). | 
 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter.  If &#x60;false&#x60;, returns the value&#39;s template. No effect on values that are not interpolated. | [default to true]
 **exclude** | **string** | A comma-separated list of field names to exclude from the response. | 
 **include** | **string** | A comma-separated list of field names to include in the response. | 
 **maskSecrets** | **bool** | Mask secret values in responses with &#x60;*****&#x60;. | [default to false]
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **wrap** | **bool** | For writes, indicates &#x60;internal_value&#x60; is wrapped; for reads, indicates &#x60;value&#x60; is wrapped. For more information on secret wrapping, see the documentation.  | [default to false]

### Return type

[**PaginatedValueList**](PaginatedValueList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersValuesPartialUpdate

> Value ProjectsParametersValuesPartialUpdate(ctx, id, parameterPk, projectPk).Evaluate(evaluate).Wrap(wrap).PatchedValue(patchedValue).Execute()

Update a value.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    evaluate := true // bool | If `true`, runs template evaluation on this parameter.  If `false`, returns the value's template. No effect on values that are not interpolated. (optional) (default to true)
    wrap := true // bool | Indicates the `internal_value` is a wrapped secret. For more information on secret wrapping, see the documentation.  (optional)
    patchedValue := *openapiclient.NewPatchedValue() // PatchedValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersValuesPartialUpdate(context.Background(), id, parameterPk, projectPk).Evaluate(evaluate).Wrap(wrap).PatchedValue(patchedValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersValuesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersValuesPartialUpdate`: Value
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersValuesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersValuesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter.  If &#x60;false&#x60;, returns the value&#39;s template. No effect on values that are not interpolated. | [default to true]
 **wrap** | **bool** | Indicates the &#x60;internal_value&#x60; is a wrapped secret. For more information on secret wrapping, see the documentation.  | 
 **patchedValue** | [**PatchedValue**](PatchedValue.md) |  | 

### Return type

[**Value**](Value.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersValuesRetrieve

> Value ProjectsParametersValuesRetrieve(ctx, id, parameterPk, projectPk).AsOf(asOf).Evaluate(evaluate).Exclude(exclude).Include(include).MaskSecrets(maskSecrets).Tag(tag).Wrap(wrap).Execute()

Retrieve a value.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    evaluate := true // bool | If `true`, runs template evaluation on this parameter.  If `false`, returns the value's template. No effect on values that are not interpolated. (optional) (default to true)
    exclude := "exclude_example" // string | A comma-separated list of field names to exclude from the response. (optional)
    include := "include_example" // string | A comma-separated list of field names to include in the response. (optional)
    maskSecrets := true // bool | Mask secret values in responses with `*****`. (optional) (default to false)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)
    wrap := true // bool | For writes, indicates `internal_value` is wrapped; for reads, indicates `value` is wrapped. For more information on secret wrapping, see the documentation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersValuesRetrieve(context.Background(), id, parameterPk, projectPk).AsOf(asOf).Evaluate(evaluate).Exclude(exclude).Include(include).MaskSecrets(maskSecrets).Tag(tag).Wrap(wrap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersValuesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersValuesRetrieve`: Value
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersValuesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersValuesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter.  If &#x60;false&#x60;, returns the value&#39;s template. No effect on values that are not interpolated. | [default to true]
 **exclude** | **string** | A comma-separated list of field names to exclude from the response. | 
 **include** | **string** | A comma-separated list of field names to include in the response. | 
 **maskSecrets** | **bool** | Mask secret values in responses with &#x60;*****&#x60;. | [default to false]
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **wrap** | **bool** | For writes, indicates &#x60;internal_value&#x60; is wrapped; for reads, indicates &#x60;value&#x60; is wrapped. For more information on secret wrapping, see the documentation.  | [default to false]

### Return type

[**Value**](Value.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsParametersValuesUpdate

> Value ProjectsParametersValuesUpdate(ctx, id, parameterPk, projectPk).Evaluate(evaluate).Wrap(wrap).Value(value).Execute()

Update a value.



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
    parameterPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The parameter id.
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project id.
    evaluate := true // bool | If `true`, runs template evaluation on this parameter.  If `false`, returns the value's template. No effect on values that are not interpolated. (optional) (default to true)
    wrap := true // bool | Indicates the `internal_value` is a wrapped secret. For more information on secret wrapping, see the documentation.  (optional)
    value := *openapiclient.NewValue("Url_example", "Id_example", "Environment_example", "EnvironmentName_example", "EarliestTag_example", "Parameter_example", "ExternalError_example", "TODO", "Value_example", false, false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, time.Now(), time.Now()) // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsParametersValuesUpdate(context.Background(), id, parameterPk, projectPk).Evaluate(evaluate).Wrap(wrap).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersValuesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsParametersValuesUpdate`: Value
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsParametersValuesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**parameterPk** | **string** | The parameter id. | 
**projectPk** | **string** | The project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsParametersValuesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **evaluate** | **bool** | If &#x60;true&#x60;, runs template evaluation on this parameter.  If &#x60;false&#x60;, returns the value&#39;s template. No effect on values that are not interpolated. | [default to true]
 **wrap** | **bool** | Indicates the &#x60;internal_value&#x60; is a wrapped secret. For more information on secret wrapping, see the documentation.  | 
 **value** | [**Value**](Value.md) |  | 

### Return type

[**Value**](Value.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPartialUpdate

> Project ProjectsPartialUpdate(ctx, id).PatchedProject(patchedProject).Execute()



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
    patchedProject := *openapiclient.NewPatchedProject() // PatchedProject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsPartialUpdate(context.Background(), id).PatchedProject(patchedProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsPartialUpdate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedProject** | [**PatchedProject**](PatchedProject.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsRetrieve

> Project ProjectsRetrieve(ctx, id).Execute()



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
    resp, r, err := apiClient.ProjectsApi.ProjectsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsRetrieve`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatePreviewCreate

> TemplatePreview ProjectsTemplatePreviewCreate(ctx, projectPk).TemplatePreviewCreateRequest(templatePreviewCreateRequest).AsOf(asOf).Environment(environment).MaskSecrets(maskSecrets).Tag(tag).Template(template).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    templatePreviewCreateRequest := *openapiclient.NewTemplatePreviewCreateRequest("Body_example") // TemplatePreviewCreateRequest | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id of the environment to use to instantiate this template. If not specified then the default environment is used. (optional)
    maskSecrets := true // bool | Masks all secrets in the template with `*****`. (optional) (default to false)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)
    template := "template_example" // string | ID of the template parameter being previewed.  If not specified, this is assumed to be a not-yet-created parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatePreviewCreate(context.Background(), projectPk).TemplatePreviewCreateRequest(templatePreviewCreateRequest).AsOf(asOf).Environment(environment).MaskSecrets(maskSecrets).Tag(tag).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatePreviewCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatePreviewCreate`: TemplatePreview
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatePreviewCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatePreviewCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templatePreviewCreateRequest** | [**TemplatePreviewCreateRequest**](TemplatePreviewCreateRequest.md) |  | 
 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id of the environment to use to instantiate this template. If not specified then the default environment is used. | 
 **maskSecrets** | **bool** | Masks all secrets in the template with &#x60;*****&#x60;. | [default to false]
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 
 **template** | **string** | ID of the template parameter being previewed.  If not specified, this is assumed to be a not-yet-created parameter. | 

### Return type

[**TemplatePreview**](TemplatePreview.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesCreate

> Template ProjectsTemplatesCreate(ctx, projectPk).TemplateCreate(templateCreate).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    templateCreate := *openapiclient.NewTemplateCreate("Name_example") // TemplateCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesCreate(context.Background(), projectPk).TemplateCreate(templateCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesCreate`: Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateCreate** | [**TemplateCreate**](TemplateCreate.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesDestroy

> ProjectsTemplatesDestroy(ctx, id, projectPk).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesDestroy(context.Background(), id, projectPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesDestroyRequest struct via the builder pattern


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


## ProjectsTemplatesList

> PaginatedTemplateList ProjectsTemplatesList(ctx, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Tag(tag).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. (optional)
    evaluate := true // bool | If `true`, evaluates the template's body.  If `false`, returns the unevaluated template body.  (optional) (default to false)
    maskSecrets := true // bool | Masks all secrets in the template with `*****`. (optional) (default to false)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesList(context.Background(), projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesList`: PaginatedTemplateList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. | 
 **evaluate** | **bool** | If &#x60;true&#x60;, evaluates the template&#39;s body.  If &#x60;false&#x60;, returns the unevaluated template body.  | [default to false]
 **maskSecrets** | **bool** | Masks all secrets in the template with &#x60;*****&#x60;. | [default to false]
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

### Return type

[**PaginatedTemplateList**](PaginatedTemplateList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesPartialUpdate

> Template ProjectsTemplatesPartialUpdate(ctx, id, projectPk).PatchedTemplate(patchedTemplate).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    patchedTemplate := *openapiclient.NewPatchedTemplate() // PatchedTemplate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesPartialUpdate(context.Background(), id, projectPk).PatchedTemplate(patchedTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesPartialUpdate`: Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedTemplate** | [**PatchedTemplate**](PatchedTemplate.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesRetrieve

> Template ProjectsTemplatesRetrieve(ctx, id, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Tag(tag).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. (optional)
    evaluate := true // bool | If `true`, evaluates the template's body.  If `false`, returns the unevaluated template body.  (optional) (default to true)
    maskSecrets := true // bool | Masks all secrets in the template with `*****`. (optional) (default to false)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesRetrieve(context.Background(), id, projectPk).AsOf(asOf).Environment(environment).Evaluate(evaluate).MaskSecrets(maskSecrets).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesRetrieve`: Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. | 
 **evaluate** | **bool** | If &#x60;true&#x60;, evaluates the template&#39;s body.  If &#x60;false&#x60;, returns the unevaluated template body.  | [default to true]
 **maskSecrets** | **bool** | Masks all secrets in the template with &#x60;*****&#x60;. | [default to false]
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesTimelineRetrieve

> TemplateTimeline ProjectsTemplatesTimelineRetrieve(ctx, id, projectPk).AsOf(asOf).Environment(environment).Tag(tag).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesTimelineRetrieve(context.Background(), id, projectPk).AsOf(asOf).Environment(environment).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesTimelineRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesTimelineRetrieve`: TemplateTimeline
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesTimelineRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesTimelineRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

### Return type

[**TemplateTimeline**](TemplateTimeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesTimelinesRetrieve

> TemplateTimeline ProjectsTemplatesTimelinesRetrieve(ctx, projectPk).AsOf(asOf).Environment(environment).Tag(tag).Execute()





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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    asOf := time.Now() // time.Time | Specify a point in time to retrieve configuration from. Cannot be specified with `tag`. (optional)
    environment := "environment_example" // string | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. (optional)
    tag := "tag_example" // string | Specify a tag to retrieve configuration from. Cannot be specified with `as_of`. Requires `environment`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesTimelinesRetrieve(context.Background(), projectPk).AsOf(asOf).Environment(environment).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesTimelinesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesTimelinesRetrieve`: TemplateTimeline
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesTimelinesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesTimelinesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | Specify a point in time to retrieve configuration from. Cannot be specified with &#x60;tag&#x60;. | 
 **environment** | **string** | Name or id of the environment to use to evaluate this template. If not specified then the original content is returned in the body. | 
 **tag** | **string** | Specify a tag to retrieve configuration from. Cannot be specified with &#x60;as_of&#x60;. Requires &#x60;environment&#x60;. | 

### Return type

[**TemplateTimeline**](TemplateTimeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTemplatesUpdate

> Template ProjectsTemplatesUpdate(ctx, id, projectPk).Template(template).Execute()



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
    projectPk := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    template := *openapiclient.NewTemplate("Url_example", "Id_example", "Name_example", false, []string{"ReferencedParameters_example"}, []string{"ReferencedTemplates_example"}, []string{"ReferencingTemplates_example"}, []string{"ReferencingValues_example"}, false, time.Now(), time.Now()) // Template | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsTemplatesUpdate(context.Background(), id, projectPk).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsTemplatesUpdate`: Template
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsTemplatesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **template** | [**Template**](Template.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdate

> Project ProjectsUpdate(ctx, id).Project(project).Execute()



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
    project := *openapiclient.NewProject("Url_example", "Id_example", "Name_example", []string{"Dependents_example"}, "TODO", []openapiclient.AwsPush{*openapiclient.NewAwsPush("Url_example", "Id_example", "Name_example", "TODO", time.Now(), time.Now(), []string{"Projects_example"}, []string{"Tags_example"}, "TODO", "TODO", "Resource_example")}, []string{"PushUrls_example"}, time.Now(), time.Now()) // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsUpdate(context.Background(), id).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsUpdate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

