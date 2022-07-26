# \QuotaApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FolderQuotaUpdateUsage**](QuotaApi.md#FolderQuotaUpdateUsage) | **Put** /quotas/folders/{name}/usage | Update folder quota usage limits
[**FolderQuotaUpdateUsageDeprecated**](QuotaApi.md#FolderQuotaUpdateUsageDeprecated) | **Put** /folder-quota-update | Update folder quota limits
[**GetFoldersQuotaScans**](QuotaApi.md#GetFoldersQuotaScans) | **Get** /quotas/folders/scans | Get active folder quota scans
[**GetFoldersQuotaScansDeprecated**](QuotaApi.md#GetFoldersQuotaScansDeprecated) | **Get** /folder-quota-scans | Get folders quota scans
[**GetUsersQuotaScans**](QuotaApi.md#GetUsersQuotaScans) | **Get** /quotas/users/scans | Get active user quota scans
[**GetUsersQuotaScansDeprecated**](QuotaApi.md#GetUsersQuotaScansDeprecated) | **Get** /quota-scans | Get quota scans
[**StartFolderQuotaScan**](QuotaApi.md#StartFolderQuotaScan) | **Post** /quotas/folders/{name}/scan | Start a folder quota scan
[**StartFolderQuotaScanDeprecated**](QuotaApi.md#StartFolderQuotaScanDeprecated) | **Post** /folder-quota-scans | Start a folder quota scan
[**StartUserQuotaScan**](QuotaApi.md#StartUserQuotaScan) | **Post** /quotas/users/{username}/scan | Start a user quota scan
[**StartUserQuotaScanDeprecated**](QuotaApi.md#StartUserQuotaScanDeprecated) | **Post** /quota-scans | Start user quota scan
[**UserQuotaUpdateUsage**](QuotaApi.md#UserQuotaUpdateUsage) | **Put** /quotas/users/{username}/usage | Update disk quota usage limits
[**UserQuotaUpdateUsageDeprecated**](QuotaApi.md#UserQuotaUpdateUsageDeprecated) | **Put** /quota-update | Update quota usage limits
[**UserTransferQuotaUpdateUsage**](QuotaApi.md#UserTransferQuotaUpdateUsage) | **Put** /quotas/users/{username}/transfer-usage | Update transfer quota usage limits



## FolderQuotaUpdateUsage

> ApiResponse FolderQuotaUpdateUsage(ctx, name).QuotaUsage(quotaUsage).Mode(mode).Execute()

Update folder quota usage limits



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
    quotaUsage := *openapiclient.NewQuotaUsage() // QuotaUsage | If used_quota_size and used_quota_files are missing they will default to 0, this means that if mode is \"add\" the current value, for the missing field, will remain unchanged, if mode is \"reset\" the missing field is set to 0
    mode := "reset" // string | the update mode specifies if the given quota usage values should be added or replace the current ones (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.FolderQuotaUpdateUsage(context.Background(), name).QuotaUsage(quotaUsage).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.FolderQuotaUpdateUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FolderQuotaUpdateUsage`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.FolderQuotaUpdateUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | folder name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderQuotaUpdateUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quotaUsage** | [**QuotaUsage**](QuotaUsage.md) | If used_quota_size and used_quota_files are missing they will default to 0, this means that if mode is \&quot;add\&quot; the current value, for the missing field, will remain unchanged, if mode is \&quot;reset\&quot; the missing field is set to 0 | 
 **mode** | **string** | the update mode specifies if the given quota usage values should be added or replace the current ones | 

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


## FolderQuotaUpdateUsageDeprecated

> ApiResponse FolderQuotaUpdateUsageDeprecated(ctx).BaseVirtualFolder(baseVirtualFolder).Mode(mode).Execute()

Update folder quota limits



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
    baseVirtualFolder := *openapiclient.NewBaseVirtualFolder() // BaseVirtualFolder | The only folder mandatory fields are mapped_path,used_quota_size and used_quota_files. Please note that if the used quota fields are missing they will default to 0, this means that if mode is \"add\" the current value will remain unchanged, if mode is \"reset\" the missing field is set to 0
    mode := "reset" // string | the update mode specifies if the given quota usage values should be added or replace the current ones (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.FolderQuotaUpdateUsageDeprecated(context.Background()).BaseVirtualFolder(baseVirtualFolder).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.FolderQuotaUpdateUsageDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FolderQuotaUpdateUsageDeprecated`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.FolderQuotaUpdateUsageDeprecated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFolderQuotaUpdateUsageDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseVirtualFolder** | [**BaseVirtualFolder**](BaseVirtualFolder.md) | The only folder mandatory fields are mapped_path,used_quota_size and used_quota_files. Please note that if the used quota fields are missing they will default to 0, this means that if mode is \&quot;add\&quot; the current value will remain unchanged, if mode is \&quot;reset\&quot; the missing field is set to 0 | 
 **mode** | **string** | the update mode specifies if the given quota usage values should be added or replace the current ones | 

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


## GetFoldersQuotaScans

> []FolderQuotaScan GetFoldersQuotaScans(ctx).Execute()

Get active folder quota scans



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
    resp, r, err := apiClient.QuotaApi.GetFoldersQuotaScans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.GetFoldersQuotaScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFoldersQuotaScans`: []FolderQuotaScan
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.GetFoldersQuotaScans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFoldersQuotaScansRequest struct via the builder pattern


### Return type

[**[]FolderQuotaScan**](FolderQuotaScan.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFoldersQuotaScansDeprecated

> []FolderQuotaScan GetFoldersQuotaScansDeprecated(ctx).Execute()

Get folders quota scans



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
    resp, r, err := apiClient.QuotaApi.GetFoldersQuotaScansDeprecated(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.GetFoldersQuotaScansDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFoldersQuotaScansDeprecated`: []FolderQuotaScan
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.GetFoldersQuotaScansDeprecated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFoldersQuotaScansDeprecatedRequest struct via the builder pattern


### Return type

[**[]FolderQuotaScan**](FolderQuotaScan.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersQuotaScans

> []QuotaScan GetUsersQuotaScans(ctx).Execute()

Get active user quota scans



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
    resp, r, err := apiClient.QuotaApi.GetUsersQuotaScans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.GetUsersQuotaScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersQuotaScans`: []QuotaScan
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.GetUsersQuotaScans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersQuotaScansRequest struct via the builder pattern


### Return type

[**[]QuotaScan**](QuotaScan.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersQuotaScansDeprecated

> []QuotaScan GetUsersQuotaScansDeprecated(ctx).Execute()

Get quota scans



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
    resp, r, err := apiClient.QuotaApi.GetUsersQuotaScansDeprecated(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.GetUsersQuotaScansDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersQuotaScansDeprecated`: []QuotaScan
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.GetUsersQuotaScansDeprecated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersQuotaScansDeprecatedRequest struct via the builder pattern


### Return type

[**[]QuotaScan**](QuotaScan.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFolderQuotaScan

> ApiResponse StartFolderQuotaScan(ctx, name).Execute()

Start a folder quota scan



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
    resp, r, err := apiClient.QuotaApi.StartFolderQuotaScan(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.StartFolderQuotaScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartFolderQuotaScan`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.StartFolderQuotaScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | folder name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFolderQuotaScanRequest struct via the builder pattern


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


## StartFolderQuotaScanDeprecated

> ApiResponse StartFolderQuotaScanDeprecated(ctx).BaseVirtualFolder(baseVirtualFolder).Execute()

Start a folder quota scan



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
    resp, r, err := apiClient.QuotaApi.StartFolderQuotaScanDeprecated(context.Background()).BaseVirtualFolder(baseVirtualFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.StartFolderQuotaScanDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartFolderQuotaScanDeprecated`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.StartFolderQuotaScanDeprecated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartFolderQuotaScanDeprecatedRequest struct via the builder pattern


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


## StartUserQuotaScan

> ApiResponse StartUserQuotaScan(ctx, username).Execute()

Start a user quota scan



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.StartUserQuotaScan(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.StartUserQuotaScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartUserQuotaScan`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.StartUserQuotaScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the username | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartUserQuotaScanRequest struct via the builder pattern


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


## StartUserQuotaScanDeprecated

> ApiResponse StartUserQuotaScanDeprecated(ctx).User(user).Execute()

Start user quota scan



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
    user := *openapiclient.NewUser() // User | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.StartUserQuotaScanDeprecated(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.StartUserQuotaScanDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartUserQuotaScanDeprecated`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.StartUserQuotaScanDeprecated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartUserQuotaScanDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

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


## UserQuotaUpdateUsage

> ApiResponse UserQuotaUpdateUsage(ctx, username).QuotaUsage(quotaUsage).Mode(mode).Execute()

Update disk quota usage limits



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
    quotaUsage := *openapiclient.NewQuotaUsage() // QuotaUsage | If used_quota_size and used_quota_files are missing they will default to 0, this means that if mode is \"add\" the current value, for the missing field, will remain unchanged, if mode is \"reset\" the missing field is set to 0
    mode := "reset" // string | the update mode specifies if the given quota usage values should be added or replace the current ones (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.UserQuotaUpdateUsage(context.Background(), username).QuotaUsage(quotaUsage).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.UserQuotaUpdateUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserQuotaUpdateUsage`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.UserQuotaUpdateUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserQuotaUpdateUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quotaUsage** | [**QuotaUsage**](QuotaUsage.md) | If used_quota_size and used_quota_files are missing they will default to 0, this means that if mode is \&quot;add\&quot; the current value, for the missing field, will remain unchanged, if mode is \&quot;reset\&quot; the missing field is set to 0 | 
 **mode** | **string** | the update mode specifies if the given quota usage values should be added or replace the current ones | 

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


## UserQuotaUpdateUsageDeprecated

> ApiResponse UserQuotaUpdateUsageDeprecated(ctx).User(user).Mode(mode).Execute()

Update quota usage limits



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
    user := *openapiclient.NewUser() // User | The only user mandatory fields are username, used_quota_size and used_quota_files. Please note that if the quota fields are missing they will default to 0, this means that if mode is \"add\" the current value will remain unchanged, if mode is \"reset\" the missing field is set to 0
    mode := "reset" // string | the update mode specifies if the given quota usage values should be added or replace the current ones (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.UserQuotaUpdateUsageDeprecated(context.Background()).User(user).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.UserQuotaUpdateUsageDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserQuotaUpdateUsageDeprecated`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.UserQuotaUpdateUsageDeprecated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserQuotaUpdateUsageDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) | The only user mandatory fields are username, used_quota_size and used_quota_files. Please note that if the quota fields are missing they will default to 0, this means that if mode is \&quot;add\&quot; the current value will remain unchanged, if mode is \&quot;reset\&quot; the missing field is set to 0 | 
 **mode** | **string** | the update mode specifies if the given quota usage values should be added or replace the current ones | 

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


## UserTransferQuotaUpdateUsage

> ApiResponse UserTransferQuotaUpdateUsage(ctx, username).TransferQuotaUsage(transferQuotaUsage).Mode(mode).Execute()

Update transfer quota usage limits



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
    transferQuotaUsage := *openapiclient.NewTransferQuotaUsage() // TransferQuotaUsage | If used_upload_data_transfer and used_download_data_transfer are missing they will default to 0, this means that if mode is \"add\" the current value, for the missing field, will remain unchanged, if mode is \"reset\" the missing field is set to 0
    mode := "reset" // string | the update mode specifies if the given quota usage values should be added or replace the current ones (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.UserTransferQuotaUpdateUsage(context.Background(), username).TransferQuotaUsage(transferQuotaUsage).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.UserTransferQuotaUpdateUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTransferQuotaUpdateUsage`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.UserTransferQuotaUpdateUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTransferQuotaUpdateUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferQuotaUsage** | [**TransferQuotaUsage**](TransferQuotaUsage.md) | If used_upload_data_transfer and used_download_data_transfer are missing they will default to 0, this means that if mode is \&quot;add\&quot; the current value, for the missing field, will remain unchanged, if mode is \&quot;reset\&quot; the missing field is set to 0 | 
 **mode** | **string** | the update mode specifies if the given quota usage values should be added or replace the current ones | 

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

