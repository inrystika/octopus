# \MaintenanceApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Dumpdata**](MaintenanceApi.md#Dumpdata) | **Get** /dumpdata | Dump data
[**GetStatus**](MaintenanceApi.md#GetStatus) | **Get** /status | Get status
[**GetVersion**](MaintenanceApi.md#GetVersion) | **Get** /version | Get version details
[**LoaddataFromFile**](MaintenanceApi.md#LoaddataFromFile) | **Get** /loaddata | Load data from path
[**LoaddataFromRequestBody**](MaintenanceApi.md#LoaddataFromRequestBody) | **Post** /loaddata | Load data



## Dumpdata

> BackupData Dumpdata(ctx).OutputFile(outputFile).OutputData(outputData).Indent(indent).Execute()

Dump data



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
    outputFile := "outputFile_example" // string | Path for the file to write the JSON serialized data to. This path is relative to the configured \"backups_path\". If this file already exists it will be overwritten. To return the backup as response body set `output_data` to true instead. (optional)
    outputData := int32(56) // int32 | output data:   * `0` or any other value != 1, the backup will be saved to a file on the server, `output_file` is required   * `1` the backup will be returned as response body  (optional)
    indent := int32(56) // int32 | indent:   * `0` no indentation. This is the default   * `1` format the output JSON  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceApi.Dumpdata(context.Background()).OutputFile(outputFile).OutputData(outputData).Indent(indent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.Dumpdata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Dumpdata`: BackupData
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.Dumpdata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDumpdataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outputFile** | **string** | Path for the file to write the JSON serialized data to. This path is relative to the configured \&quot;backups_path\&quot;. If this file already exists it will be overwritten. To return the backup as response body set &#x60;output_data&#x60; to true instead. | 
 **outputData** | **int32** | output data:   * &#x60;0&#x60; or any other value !&#x3D; 1, the backup will be saved to a file on the server, &#x60;output_file&#x60; is required   * &#x60;1&#x60; the backup will be returned as response body  | 
 **indent** | **int32** | indent:   * &#x60;0&#x60; no indentation. This is the default   * &#x60;1&#x60; format the output JSON  | 

### Return type

[**BackupData**](BackupData.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> ServicesStatus GetStatus(ctx).Execute()

Get status



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
    resp, r, err := apiClient.MaintenanceApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: ServicesStatus
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**ServicesStatus**](ServicesStatus.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> VersionInfo GetVersion(ctx).Execute()

Get version details



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
    resp, r, err := apiClient.MaintenanceApi.GetVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: VersionInfo
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.GetVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


### Return type

[**VersionInfo**](VersionInfo.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoaddataFromFile

> ApiResponse LoaddataFromFile(ctx).InputFile(inputFile).ScanQuota(scanQuota).Mode(mode).Execute()

Load data from path



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
    inputFile := "inputFile_example" // string | Path for the file to read the JSON serialized data from. This can be an absolute path or a path relative to the configured \"backups_path\". The max allowed file size is 10MB
    scanQuota := int32(56) // int32 | Quota scan:   * `0` no quota scan is done, the imported users/folders will have used_quota_size and used_quota_files = 0 or the existing values if they already exists. This is the default   * `1` scan quota   * `2` scan quota if the user has quota restrictions required: false  (optional)
    mode := int32(56) // int32 | Mode:   * `0` New objects are added, existing ones are updated. This is the default   * `1` New objects are added, existing ones are not modified   * `2` New objects are added, existing ones are updated and connected users are disconnected and so forced to use the new configuration  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceApi.LoaddataFromFile(context.Background()).InputFile(inputFile).ScanQuota(scanQuota).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.LoaddataFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoaddataFromFile`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.LoaddataFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoaddataFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFile** | **string** | Path for the file to read the JSON serialized data from. This can be an absolute path or a path relative to the configured \&quot;backups_path\&quot;. The max allowed file size is 10MB | 
 **scanQuota** | **int32** | Quota scan:   * &#x60;0&#x60; no quota scan is done, the imported users/folders will have used_quota_size and used_quota_files &#x3D; 0 or the existing values if they already exists. This is the default   * &#x60;1&#x60; scan quota   * &#x60;2&#x60; scan quota if the user has quota restrictions required: false  | 
 **mode** | **int32** | Mode:   * &#x60;0&#x60; New objects are added, existing ones are updated. This is the default   * &#x60;1&#x60; New objects are added, existing ones are not modified   * &#x60;2&#x60; New objects are added, existing ones are updated and connected users are disconnected and so forced to use the new configuration  | 

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


## LoaddataFromRequestBody

> ApiResponse LoaddataFromRequestBody(ctx).BackupData(backupData).ScanQuota(scanQuota).Mode(mode).Execute()

Load data



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
    backupData := *openapiclient.NewBackupData() // BackupData | 
    scanQuota := int32(56) // int32 | Quota scan:   * `0` no quota scan is done, the imported users/folders will have used_quota_size and used_quota_files = 0 or the existing values if they already exists. This is the default   * `1` scan quota   * `2` scan quota if the user has quota restrictions required: false  (optional)
    mode := int32(56) // int32 | Mode:   * `0` New objects are added, existing ones are updated. This is the default   * `1` New objects are added, existing ones are not modified   * `2` New objects are added, existing ones are updated and connected users are disconnected and so forced to use the new configuration  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceApi.LoaddataFromRequestBody(context.Background()).BackupData(backupData).ScanQuota(scanQuota).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceApi.LoaddataFromRequestBody``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoaddataFromRequestBody`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceApi.LoaddataFromRequestBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoaddataFromRequestBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupData** | [**BackupData**](BackupData.md) |  | 
 **scanQuota** | **int32** | Quota scan:   * &#x60;0&#x60; no quota scan is done, the imported users/folders will have used_quota_size and used_quota_files &#x3D; 0 or the existing values if they already exists. This is the default   * &#x60;1&#x60; scan quota   * &#x60;2&#x60; scan quota if the user has quota restrictions required: false  | 
 **mode** | **int32** | Mode:   * &#x60;0&#x60; New objects are added, existing ones are updated. This is the default   * &#x60;1&#x60; New objects are added, existing ones are not modified   * &#x60;2&#x60; New objects are added, existing ones are updated and connected users are disconnected and so forced to use the new configuration  | 

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

