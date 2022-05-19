# \BackupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupSnapshotCreate**](BackupApi.md#BackupSnapshotCreate) | **Post** /api/v1/backup/snapshot/ | Get a snapshot of all Projects with parameters



## BackupSnapshotCreate

> BackupDataSnapshot BackupSnapshotCreate(ctx).Execute()

Get a snapshot of all Projects with parameters

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.BackupSnapshotCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.BackupSnapshotCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupSnapshotCreate`: BackupDataSnapshot
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.BackupSnapshotCreate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBackupSnapshotCreateRequest struct via the builder pattern


### Return type

[**BackupDataSnapshot**](BackupDataSnapshot.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWTAuth](../README.md#JWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

