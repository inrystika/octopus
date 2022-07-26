# \PublicSharesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadShareFile**](PublicSharesApi.md#DownloadShareFile) | **Get** /shares/{id}/files | Download a single file
[**GetShare**](PublicSharesApi.md#GetShare) | **Get** /shares/{id} | Download shared files and folders as a single zip file
[**GetShareDirContents**](PublicSharesApi.md#GetShareDirContents) | **Get** /shares/{id}/dirs | Read directory contents
[**UploadSingleToShare**](PublicSharesApi.md#UploadSingleToShare) | **Post** /shares/{id}/{fileName} | Upload a single file to the shared path
[**UploadToShare**](PublicSharesApi.md#UploadToShare) | **Post** /shares/{id} | Upload one or more files to the shared path



## DownloadShareFile

> *os.File DownloadShareFile(ctx, id).Path(path).Inline(inline).Execute()

Download a single file



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
    id := "id_example" // string | the share id
    path := "path_example" // string | Path to the file to download. It must be URL encoded, for example the path \"my dir/àdir/file.txt\" must be sent as \"my%20dir%2F%C3%A0dir%2Ffile.txt\"
    inline := "inline_example" // string | If set, the response will not have the Content-Disposition header set to `attachment` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSharesApi.DownloadShareFile(context.Background(), id).Path(path).Inline(inline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSharesApi.DownloadShareFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadShareFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PublicSharesApi.DownloadShareFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadShareFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path to the file to download. It must be URL encoded, for example the path \&quot;my dir/àdir/file.txt\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir%2Ffile.txt\&quot; | 
 **inline** | **string** | If set, the response will not have the Content-Disposition header set to &#x60;attachment&#x60; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShare

> *os.File GetShare(ctx, id).Compress(compress).Execute()

Download shared files and folders as a single zip file



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
    id := "id_example" // string | the share id
    compress := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSharesApi.GetShare(context.Background(), id).Compress(compress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSharesApi.GetShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShare`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PublicSharesApi.GetShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **compress** | **bool** |  | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareDirContents

> []DirEntry GetShareDirContents(ctx, id).Path(path).Execute()

Read directory contents



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
    id := "id_example" // string | the share id
    path := "path_example" // string | Path to the folder to read. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\". If empty or missing the root folder is assumed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSharesApi.GetShareDirContents(context.Background(), id).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSharesApi.GetShareDirContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShareDirContents`: []DirEntry
    fmt.Fprintf(os.Stdout, "Response from `PublicSharesApi.GetShareDirContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareDirContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path to the folder to read. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot;. If empty or missing the root folder is assumed | 

### Return type

[**[]DirEntry**](DirEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSingleToShare

> []ApiResponse UploadSingleToShare(ctx, id, fileName).Body(body).XSFTPGOMTIME(xSFTPGOMTIME).Execute()

Upload a single file to the shared path



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
    id := "id_example" // string | the share id
    fileName := "fileName_example" // string | the name of the new file. It must be path encoded. Sub directories are not accepted
    body := os.NewFile(1234, "some_file") // *os.File | 
    xSFTPGOMTIME := int32(56) // int32 | File modification time as unix timestamp in milliseconds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSharesApi.UploadSingleToShare(context.Background(), id, fileName).Body(body).XSFTPGOMTIME(xSFTPGOMTIME).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSharesApi.UploadSingleToShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadSingleToShare`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicSharesApi.UploadSingleToShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 
**fileName** | **string** | the name of the new file. It must be path encoded. Sub directories are not accepted | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSingleToShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 
 **xSFTPGOMTIME** | **int32** | File modification time as unix timestamp in milliseconds | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/_*, text/_*, image/_*, audio/_*, video/_*
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToShare

> []ApiResponse UploadToShare(ctx, id).Filenames(filenames).Execute()

Upload one or more files to the shared path



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
    id := "id_example" // string | the share id
    filenames := []*os.File{"TODO"} // []*os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSharesApi.UploadToShare(context.Background(), id).Filenames(filenames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSharesApi.UploadToShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadToShare`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicSharesApi.UploadToShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadToShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filenames** | **[]*os.File** |  | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

