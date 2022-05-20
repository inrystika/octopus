# \ConnectionsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseConnection**](ConnectionsApi.md#CloseConnection) | **Delete** /connections/{connectionID} | Close connection
[**GetConnections**](ConnectionsApi.md#GetConnections) | **Get** /connections | Get connections details



## CloseConnection

> ApiResponse CloseConnection(ctx, connectionID).Execute()

Close connection



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
    connectionID := "connectionID_example" // string | ID of the connection to close

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.CloseConnection(context.Background(), connectionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CloseConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseConnection`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CloseConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionID** | **string** | ID of the connection to close | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnections

> []ConnectionStatus GetConnections(ctx).Execute()

Get connections details



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
    resp, r, err := apiClient.ConnectionsApi.GetConnections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnections`: []ConnectionStatus
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


### Return type

[**[]ConnectionStatus**](ConnectionStatus.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

