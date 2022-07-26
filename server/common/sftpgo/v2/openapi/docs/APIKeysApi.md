# \APIKeysApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApiKey**](APIKeysApi.md#AddApiKey) | **Post** /apikeys | Add API key
[**DeleteApiKey**](APIKeysApi.md#DeleteApiKey) | **Delete** /apikeys/{id} | Delete API key
[**GetApiKeyById**](APIKeysApi.md#GetApiKeyById) | **Get** /apikeys/{id} | Find API key by id
[**GetApiKeys**](APIKeysApi.md#GetApiKeys) | **Get** /apikeys | Get API keys
[**UpdateApiKey**](APIKeysApi.md#UpdateApiKey) | **Put** /apikeys/{id} | Update API key



## AddApiKey

> InlineResponse201 AddApiKey(ctx).AuthAPIKey(authAPIKey).Execute()

Add API key



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
    authAPIKey := *openapiclient.NewAuthAPIKey() // AuthAPIKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.AddApiKey(context.Background()).AuthAPIKey(authAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.AddApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApiKey`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.AddApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authAPIKey** | [**AuthAPIKey**](AuthAPIKey.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKey

> ApiResponse DeleteApiKey(ctx, id).Execute()

Delete API key



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
    id := "id_example" // string | the key id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.DeleteApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiKey`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyById

> AuthAPIKey GetApiKeyById(ctx, id).Execute()

Find API key by id



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
    id := "id_example" // string | the key id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.GetApiKeyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.GetApiKeyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyById`: AuthAPIKey
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.GetApiKeyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthAPIKey**](AuthAPIKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeys

> []AuthAPIKey GetApiKeys(ctx).Offset(offset).Limit(limit).Order(order).Execute()

Get API keys



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
    order := "ASC" // string | Ordering API keys by id. Default ASC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.GetApiKeys(context.Background()).Offset(offset).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.GetApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeys`: []AuthAPIKey
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.GetApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** | The maximum number of items to return. Max value is 500, default is 100 | [default to 100]
 **order** | **string** | Ordering API keys by id. Default ASC | 

### Return type

[**[]AuthAPIKey**](AuthAPIKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKey

> ApiResponse UpdateApiKey(ctx, id).AuthAPIKey(authAPIKey).Execute()

Update API key



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
    id := "id_example" // string | the key id
    authAPIKey := *openapiclient.NewAuthAPIKey() // AuthAPIKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysApi.UpdateApiKey(context.Background(), id).AuthAPIKey(authAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.UpdateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiKey`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authAPIKey** | [**AuthAPIKey**](AuthAPIKey.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

