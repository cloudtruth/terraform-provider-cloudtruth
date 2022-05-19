# \UtilsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UtilsGeneratePasswordCreate**](UtilsApi.md#UtilsGeneratePasswordCreate) | **Post** /api/v1/utils/generate_password/ | Get a randomly generated password using AWS Secrets Manager, with fallback to /dev/urandom.



## UtilsGeneratePasswordCreate

> GeneratedPasswordResponse UtilsGeneratePasswordCreate(ctx).Length(length).RequireHardwareGeneration(requireHardwareGeneration).RequireLowercase(requireLowercase).RequireNumbers(requireNumbers).RequireSpaces(requireSpaces).RequireSymbols(requireSymbols).RequireUppercase(requireUppercase).Execute()

Get a randomly generated password using AWS Secrets Manager, with fallback to /dev/urandom.



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
    length := int32(56) // int32 | The length of the password to generate.  Minimum of 8, maximum of 4095.
    requireHardwareGeneration := true // bool | Default behavior is to fallback to /dev/urandom if we fail to get a random password from AWS Secrets Manager.  If set to 'True', we will not fallback to local password generation using /dev/urandom.  Default: False (optional) (default to false)
    requireLowercase := true // bool | The password must include lowercase letters [a-z]. Default: True. (optional) (default to true)
    requireNumbers := true // bool | The password must include numbers [0-9].  Default: True. (optional) (default to true)
    requireSpaces := true // bool | The password must include spaces [ ].  Default: False. (optional) (default to false)
    requireSymbols := true // bool | The password must include symbols [!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~].  Default: False. (optional) (default to false)
    requireUppercase := true // bool | The password must include uppercase letters [A-Z].  Default: True. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilsApi.UtilsGeneratePasswordCreate(context.Background()).Length(length).RequireHardwareGeneration(requireHardwareGeneration).RequireLowercase(requireLowercase).RequireNumbers(requireNumbers).RequireSpaces(requireSpaces).RequireSymbols(requireSymbols).RequireUppercase(requireUppercase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.UtilsGeneratePasswordCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UtilsGeneratePasswordCreate`: GeneratedPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `UtilsApi.UtilsGeneratePasswordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUtilsGeneratePasswordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int32** | The length of the password to generate.  Minimum of 8, maximum of 4095. | 
 **requireHardwareGeneration** | **bool** | Default behavior is to fallback to /dev/urandom if we fail to get a random password from AWS Secrets Manager.  If set to &#39;True&#39;, we will not fallback to local password generation using /dev/urandom.  Default: False | [default to false]
 **requireLowercase** | **bool** | The password must include lowercase letters [a-z]. Default: True. | [default to true]
 **requireNumbers** | **bool** | The password must include numbers [0-9].  Default: True. | [default to true]
 **requireSpaces** | **bool** | The password must include spaces [ ].  Default: False. | [default to false]
 **requireSymbols** | **bool** | The password must include symbols [!\&quot;#$%&amp;&#39;()*+,-./:;&lt;&#x3D;&gt;?@[\\]^_&#x60;{|}~].  Default: False. | [default to false]
 **requireUppercase** | **bool** | The password must include uppercase letters [A-Z].  Default: True. | [default to true]

### Return type

[**GeneratedPasswordResponse**](GeneratedPasswordResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

