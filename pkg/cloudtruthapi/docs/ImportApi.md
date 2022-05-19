# \ImportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportCreate**](ImportApi.md#ImportCreate) | **Post** /api/v1/import/ | 



## ImportCreate

> ImportCreateResponse ImportCreate(ctx).ImportCreateRequest(importCreateRequest).MaskSecrets(maskSecrets).Preview(preview).Execute()





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
    importCreateRequest := *openapiclient.NewImportCreateRequest("Project_example", "Body_example", []string{"Secrets_example"}, []string{"Ignore_example"}) // ImportCreateRequest | 
    maskSecrets := true // bool | Mask secrets in the response (optional) (default to false)
    preview := true // bool | Do not create any objects (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportApi.ImportCreate(context.Background()).ImportCreateRequest(importCreateRequest).MaskSecrets(maskSecrets).Preview(preview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.ImportCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportCreate`: ImportCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.ImportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importCreateRequest** | [**ImportCreateRequest**](ImportCreateRequest.md) |  | 
 **maskSecrets** | **bool** | Mask secrets in the response | [default to false]
 **preview** | **bool** | Do not create any objects | [default to false]

### Return type

[**ImportCreateResponse**](ImportCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

