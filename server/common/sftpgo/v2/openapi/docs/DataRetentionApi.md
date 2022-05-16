# \DataRetentionApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsersRetentionChecks**](DataRetentionApi.md#GetUsersRetentionChecks) | **Get** /retention/users/checks | Get retention checks
[**StartUserRetentionCheck**](DataRetentionApi.md#StartUserRetentionCheck) | **Post** /retention/users/{username}/check | Start a retention check



## GetUsersRetentionChecks

> []RetentionCheck GetUsersRetentionChecks(ctx).Execute()

Get retention checks



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
    resp, r, err := apiClient.DataRetentionApi.GetUsersRetentionChecks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionApi.GetUsersRetentionChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersRetentionChecks`: []RetentionCheck
    fmt.Fprintf(os.Stdout, "Response from `DataRetentionApi.GetUsersRetentionChecks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRetentionChecksRequest struct via the builder pattern


### Return type

[**[]RetentionCheck**](RetentionCheck.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartUserRetentionCheck

> ApiResponse StartUserRetentionCheck(ctx, username).FolderRetention(folderRetention).Notifications(notifications).Execute()

Start a retention check



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
    username := "username_example" // string | the username
    folderRetention := []openapiclient.FolderRetention{*openapiclient.NewFolderRetention()} // []FolderRetention | Defines virtual paths to check and their retention time in hours
    notifications := []openapiclient.RetentionCheckNotification{openapiclient.RetentionCheckNotification("Hook")} // []RetentionCheckNotification | specify how to notify results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataRetentionApi.StartUserRetentionCheck(context.Background(), username).FolderRetention(folderRetention).Notifications(notifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionApi.StartUserRetentionCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartUserRetentionCheck`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DataRetentionApi.StartUserRetentionCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the username | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartUserRetentionCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderRetention** | [**[]FolderRetention**](FolderRetention.md) | Defines virtual paths to check and their retention time in hours | 
 **notifications** | [**[]RetentionCheckNotification**](RetentionCheckNotification.md) | specify how to notify results | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

