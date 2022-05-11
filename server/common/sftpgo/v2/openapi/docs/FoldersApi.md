# \FoldersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFolder**](FoldersApi.md#AddFolder) | **Post** /folders | Add folder
[**DeleteFolder**](FoldersApi.md#DeleteFolder) | **Delete** /folders/{name} | Delete folder
[**GetFolderByName**](FoldersApi.md#GetFolderByName) | **Get** /folders/{name} | Find folders by name
[**GetFolders**](FoldersApi.md#GetFolders) | **Get** /folders | Get folders
[**UpdateFolder**](FoldersApi.md#UpdateFolder) | **Put** /folders/{name} | Update folder



## AddFolder

> BaseVirtualFolder AddFolder(ctx).BaseVirtualFolder(baseVirtualFolder).Execute()

Add folder



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
    baseVirtualFolder := *openapiclient.NewBaseVirtualFolder() // BaseVirtualFolder | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.AddFolder(context.Background()).BaseVirtualFolder(baseVirtualFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.AddFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFolder`: BaseVirtualFolder
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.AddFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseVirtualFolder** | [**BaseVirtualFolder**](BaseVirtualFolder.md) |  | 

### Return type

[**BaseVirtualFolder**](BaseVirtualFolder.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFolder

> ApiResponse DeleteFolder(ctx, name).Execute()

Delete folder



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
    name := "name_example" // string | folder name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.DeleteFolder(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.DeleteFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFolder`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | folder name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


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


## GetFolderByName

> BaseVirtualFolder GetFolderByName(ctx, name).Execute()

Find folders by name



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
    name := "name_example" // string | folder name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.GetFolderByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.GetFolderByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderByName`: BaseVirtualFolder
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.GetFolderByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | folder name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseVirtualFolder**](BaseVirtualFolder.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolders

> []BaseVirtualFolder GetFolders(ctx).Offset(offset).Limit(limit).Order(order).Execute()

Get folders



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
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of items to return. Max value is 500, default is 100 (optional) (default to 100)
    order := "ASC" // string | Ordering folders by path. Default ASC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.GetFolders(context.Background()).Offset(offset).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.GetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolders`: []BaseVirtualFolder
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.GetFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** | The maximum number of items to return. Max value is 500, default is 100 | [default to 100]
 **order** | **string** | Ordering folders by path. Default ASC | 

### Return type

[**[]BaseVirtualFolder**](BaseVirtualFolder.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolder

> ApiResponse UpdateFolder(ctx, name).BaseVirtualFolder(baseVirtualFolder).Execute()

Update folder



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
    name := "name_example" // string | folder name
    baseVirtualFolder := *openapiclient.NewBaseVirtualFolder() // BaseVirtualFolder | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.UpdateFolder(context.Background(), name).BaseVirtualFolder(baseVirtualFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.UpdateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFolder`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.UpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | folder name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseVirtualFolder** | [**BaseVirtualFolder**](BaseVirtualFolder.md) |  | 

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

