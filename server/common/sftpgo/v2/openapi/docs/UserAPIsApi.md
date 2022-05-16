# \UserAPIsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddShare**](UserAPIsApi.md#AddShare) | **Post** /user/shares | Add a share
[**ChangeUserPassword**](UserAPIsApi.md#ChangeUserPassword) | **Put** /user/changepwd | Change user password
[**CreateUserDir**](UserAPIsApi.md#CreateUserDir) | **Post** /user/dirs | Create a directory
[**CreateUserFile**](UserAPIsApi.md#CreateUserFile) | **Post** /user/files/upload | Upload a single file
[**CreateUserFiles**](UserAPIsApi.md#CreateUserFiles) | **Post** /user/files | Upload files
[**DeleteUserDir**](UserAPIsApi.md#DeleteUserDir) | **Delete** /user/dirs | Delete a directory
[**DeleteUserFile**](UserAPIsApi.md#DeleteUserFile) | **Delete** /user/files | Delete a file
[**DeleteUserShare**](UserAPIsApi.md#DeleteUserShare) | **Delete** /user/shares/{id} | Delete share
[**DownloadUserFile**](UserAPIsApi.md#DownloadUserFile) | **Get** /user/files | Download a single file
[**GenerateUserRecoveryCodes**](UserAPIsApi.md#GenerateUserRecoveryCodes) | **Post** /user/2fa/recoverycodes | Generate recovery codes
[**GenerateUserTotpSecret**](UserAPIsApi.md#GenerateUserTotpSecret) | **Post** /user/totp/generate | Generate a new TOTP secret
[**GetUserDirContents**](UserAPIsApi.md#GetUserDirContents) | **Get** /user/dirs | Read directory contents
[**GetUserFile**](UserAPIsApi.md#GetUserFile) | **Get** /user/file | Download a single file
[**GetUserFolderContents**](UserAPIsApi.md#GetUserFolderContents) | **Get** /user/folder | Read folders contents
[**GetUserProfile**](UserAPIsApi.md#GetUserProfile) | **Get** /user/profile | Get user profile
[**GetUserPublicKeys**](UserAPIsApi.md#GetUserPublicKeys) | **Get** /user/publickeys | Get the user&#39;s public keys
[**GetUserRecoveryCodes**](UserAPIsApi.md#GetUserRecoveryCodes) | **Get** /user/2fa/recoverycodes | Get recovery codes
[**GetUserShareById**](UserAPIsApi.md#GetUserShareById) | **Get** /user/shares/{id} | Get share by id
[**GetUserShares**](UserAPIsApi.md#GetUserShares) | **Get** /user/shares | List user shares
[**GetUserTotpConfigs**](UserAPIsApi.md#GetUserTotpConfigs) | **Get** /user/totp/configs | Get available TOTP configuration
[**RenameUserDir**](UserAPIsApi.md#RenameUserDir) | **Patch** /user/dirs | Rename a directory
[**RenameUserFile**](UserAPIsApi.md#RenameUserFile) | **Patch** /user/files | Rename a file
[**SaveUserTotpConfig**](UserAPIsApi.md#SaveUserTotpConfig) | **Post** /user/totp/save | Save a TOTP config
[**SetUserPublicKeys**](UserAPIsApi.md#SetUserPublicKeys) | **Put** /user/publickeys | Set the user&#39;s public keys
[**SetpropsUserFile**](UserAPIsApi.md#SetpropsUserFile) | **Patch** /user/files/metadata | Set metadata for a file/directory
[**Streamzip**](UserAPIsApi.md#Streamzip) | **Post** /user/streamzip | Download multiple files and folders as a single zip file
[**UpdateUserProfile**](UserAPIsApi.md#UpdateUserProfile) | **Put** /user/profile | Update user profile
[**UpdateUserShare**](UserAPIsApi.md#UpdateUserShare) | **Put** /user/shares/{id} | Update share
[**ValidateUserTotpSecret**](UserAPIsApi.md#ValidateUserTotpSecret) | **Post** /user/totp/validate | Validate a one time authentication code



## AddShare

> ApiResponse AddShare(ctx).Share(share).Execute()

Add a share



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
    share := *openapiclient.NewShare() // Share | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.AddShare(context.Background()).Share(share).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.AddShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddShare`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.AddShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **share** | [**Share**](Share.md) |  | 

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


## ChangeUserPassword

> ApiResponse ChangeUserPassword(ctx).PwdChange(pwdChange).Execute()

Change user password



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
    pwdChange := *openapiclient.NewPwdChange() // PwdChange | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.ChangeUserPassword(context.Background()).PwdChange(pwdChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.ChangeUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeUserPassword`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.ChangeUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pwdChange** | [**PwdChange**](PwdChange.md) |  | 

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


## CreateUserDir

> []ApiResponse CreateUserDir(ctx).Path(path).MkdirParents(mkdirParents).Execute()

Create a directory



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
    path := "path_example" // string | Path to the folder to create. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\"
    mkdirParents := true // bool | Create parent directories if they do not exist? (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.CreateUserDir(context.Background()).Path(path).MkdirParents(mkdirParents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.CreateUserDir``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserDir`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.CreateUserDir`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserDirRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the folder to create. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot; | 
 **mkdirParents** | **bool** | Create parent directories if they do not exist? | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserFile

> []ApiResponse CreateUserFile(ctx).Path(path).Body(body).MkdirParents(mkdirParents).XSFTPGOMTIME(xSFTPGOMTIME).Execute()

Upload a single file



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
    path := "path_example" // string | Full file path. It must be path encoded, for example the path \"my dir/àdir/file.txt\" must be sent as \"my%20dir%2F%C3%A0dir%2Ffile.txt\". The parent directory must exist. If a file with the same name already exists, it will be overwritten
    body := os.NewFile(1234, "some_file") // *os.File | 
    mkdirParents := true // bool | Create parent directories if they do not exist? (optional)
    xSFTPGOMTIME := int32(56) // int32 | File modification time as unix timestamp in milliseconds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.CreateUserFile(context.Background()).Path(path).Body(body).MkdirParents(mkdirParents).XSFTPGOMTIME(xSFTPGOMTIME).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.CreateUserFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserFile`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.CreateUserFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Full file path. It must be path encoded, for example the path \&quot;my dir/àdir/file.txt\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir%2Ffile.txt\&quot;. The parent directory must exist. If a file with the same name already exists, it will be overwritten | 
 **body** | ***os.File** |  | 
 **mkdirParents** | **bool** | Create parent directories if they do not exist? | 
 **xSFTPGOMTIME** | **int32** | File modification time as unix timestamp in milliseconds | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/_*, text/_*, image/_*, audio/_*, video/_*
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserFiles

> []ApiResponse CreateUserFiles(ctx).Path(path).MkdirParents(mkdirParents).Filenames(filenames).Execute()

Upload files



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
    path := "path_example" // string | Parent directory for the uploaded files. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\". If empty or missing the root path is assumed. If a file with the same name already exists, it will be overwritten (optional)
    mkdirParents := true // bool | Create parent directories if they do not exist? (optional)
    filenames := []*os.File{"TODO"} // []*os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.CreateUserFiles(context.Background()).Path(path).MkdirParents(mkdirParents).Filenames(filenames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.CreateUserFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserFiles`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.CreateUserFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Parent directory for the uploaded files. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot;. If empty or missing the root path is assumed. If a file with the same name already exists, it will be overwritten | 
 **mkdirParents** | **bool** | Create parent directories if they do not exist? | 
 **filenames** | **[]*os.File** |  | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserDir

> []ApiResponse DeleteUserDir(ctx).Path(path).Execute()

Delete a directory



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
    path := "path_example" // string | Path to the folder to delete. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.DeleteUserDir(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.DeleteUserDir``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserDir`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.DeleteUserDir`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDirRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the folder to delete. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot; | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserFile

> []ApiResponse DeleteUserFile(ctx).Path(path).Execute()

Delete a file



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
    path := "path_example" // string | Path to the file to delete. It must be URL encoded

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.DeleteUserFile(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.DeleteUserFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserFile`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.DeleteUserFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the file to delete. It must be URL encoded | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserShare

> ApiResponse DeleteUserShare(ctx, id).Execute()

Delete share



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.DeleteUserShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.DeleteUserShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserShare`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.DeleteUserShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserShareRequest struct via the builder pattern


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


## DownloadUserFile

> *os.File DownloadUserFile(ctx).Path(path).Inline(inline).Execute()

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
    path := "path_example" // string | Path to the file to download. It must be URL encoded, for example the path \"my dir/àdir/file.txt\" must be sent as \"my%20dir%2F%C3%A0dir%2Ffile.txt\"
    inline := "inline_example" // string | If set, the response will not have the Content-Disposition header set to `attachment` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.DownloadUserFile(context.Background()).Path(path).Inline(inline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.DownloadUserFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadUserFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.DownloadUserFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadUserFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the file to download. It must be URL encoded, for example the path \&quot;my dir/àdir/file.txt\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir%2Ffile.txt\&quot; | 
 **inline** | **string** | If set, the response will not have the Content-Disposition header set to &#x60;attachment&#x60; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUserRecoveryCodes

> []string GenerateUserRecoveryCodes(ctx).Execute()

Generate recovery codes



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
    resp, r, err := apiClient.UserAPIsApi.GenerateUserRecoveryCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GenerateUserRecoveryCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUserRecoveryCodes`: []string
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GenerateUserRecoveryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUserRecoveryCodesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUserTotpSecret

> InlineResponse200 GenerateUserTotpSecret(ctx).InlineObject6(inlineObject6).Execute()

Generate a new TOTP secret



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
    inlineObject6 := *openapiclient.NewInlineObject6() // InlineObject6 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.GenerateUserTotpSecret(context.Background()).InlineObject6(inlineObject6).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GenerateUserTotpSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUserTotpSecret`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GenerateUserTotpSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUserTotpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject6** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDirContents

> []DirEntry GetUserDirContents(ctx).Path(path).Execute()

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
    path := "path_example" // string | Path to the folder to read. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\". If empty or missing the root folder is assumed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.GetUserDirContents(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserDirContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDirContents`: []DirEntry
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserDirContents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDirContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the folder to read. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot;. If empty or missing the root folder is assumed | 

### Return type

[**[]DirEntry**](DirEntry.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFile

> *os.File GetUserFile(ctx).Path(path).Execute()

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
    path := "path_example" // string | Path to the file to download. It must be URL encoded, for example the path \"my dir/àdir/file.txt\" must be sent as \"my%20dir%2F%C3%A0dir%2Ffile.txt\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.GetUserFile(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the file to download. It must be URL encoded, for example the path \&quot;my dir/àdir/file.txt\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir%2Ffile.txt\&quot; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFolderContents

> []DirEntry GetUserFolderContents(ctx).Path(path).Execute()

Read folders contents



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
    path := "path_example" // string | Path to the folder to read. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\". If empty or missing the root folder is assumed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.GetUserFolderContents(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserFolderContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFolderContents`: []DirEntry
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserFolderContents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFolderContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the folder to read. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot;. If empty or missing the root folder is assumed | 

### Return type

[**[]DirEntry**](DirEntry.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserProfile

> UserProfile GetUserProfile(ctx).Execute()

Get user profile



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
    resp, r, err := apiClient.UserAPIsApi.GetUserProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserProfile`: UserProfile
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfileRequest struct via the builder pattern


### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPublicKeys

> []string GetUserPublicKeys(ctx).Execute()

Get the user's public keys



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
    resp, r, err := apiClient.UserAPIsApi.GetUserPublicKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserPublicKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPublicKeys`: []string
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserPublicKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPublicKeysRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRecoveryCodes

> []RecoveryCode GetUserRecoveryCodes(ctx).Execute()

Get recovery codes



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
    resp, r, err := apiClient.UserAPIsApi.GetUserRecoveryCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserRecoveryCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRecoveryCodes`: []RecoveryCode
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserRecoveryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRecoveryCodesRequest struct via the builder pattern


### Return type

[**[]RecoveryCode**](RecoveryCode.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserShareById

> Share GetUserShareById(ctx, id).Execute()

Get share by id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.GetUserShareById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserShareById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserShareById`: Share
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserShareById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserShareByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Share**](Share.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserShares

> []Share GetUserShares(ctx).Offset(offset).Limit(limit).Order(order).Execute()

List user shares



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
    order := "ASC" // string | Ordering shares by ID. Default ASC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.GetUserShares(context.Background()).Offset(offset).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserShares`: []Share
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** | The maximum number of items to return. Max value is 500, default is 100 | [default to 100]
 **order** | **string** | Ordering shares by ID. Default ASC | 

### Return type

[**[]Share**](Share.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTotpConfigs

> []TOTPConfig GetUserTotpConfigs(ctx).Execute()

Get available TOTP configuration



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
    resp, r, err := apiClient.UserAPIsApi.GetUserTotpConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.GetUserTotpConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTotpConfigs`: []TOTPConfig
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.GetUserTotpConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTotpConfigsRequest struct via the builder pattern


### Return type

[**[]TOTPConfig**](TOTPConfig.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameUserDir

> []ApiResponse RenameUserDir(ctx).Path(path).Target(target).Execute()

Rename a directory



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
    path := "path_example" // string | Path to the folder to rename. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\"
    target := "target_example" // string | New name. It must be URL encoded, for example the path \"my dir/àdir\" must be sent as \"my%20dir%2F%C3%A0dir\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.RenameUserDir(context.Background()).Path(path).Target(target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.RenameUserDir``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameUserDir`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.RenameUserDir`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameUserDirRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the folder to rename. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot; | 
 **target** | **string** | New name. It must be URL encoded, for example the path \&quot;my dir/àdir\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir\&quot; | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameUserFile

> []ApiResponse RenameUserFile(ctx).Path(path).Target(target).Execute()

Rename a file



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
    path := "path_example" // string | Path to the file to rename. It must be URL encoded
    target := "target_example" // string | New name. It must be URL encoded

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.RenameUserFile(context.Background()).Path(path).Target(target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.RenameUserFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameUserFile`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.RenameUserFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameUserFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path to the file to rename. It must be URL encoded | 
 **target** | **string** | New name. It must be URL encoded | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUserTotpConfig

> ApiResponse SaveUserTotpConfig(ctx).UserTOTPConfig(userTOTPConfig).Execute()

Save a TOTP config



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
    userTOTPConfig := *openapiclient.NewUserTOTPConfig() // UserTOTPConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.SaveUserTotpConfig(context.Background()).UserTOTPConfig(userTOTPConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.SaveUserTotpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveUserTotpConfig`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.SaveUserTotpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveUserTotpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userTOTPConfig** | [**UserTOTPConfig**](UserTOTPConfig.md) |  | 

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


## SetUserPublicKeys

> ApiResponse SetUserPublicKeys(ctx).RequestBody(requestBody).Execute()

Set the user's public keys



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
    requestBody := []string{"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIPVILdH2u3yV5SAeE6XksD1z1vXRg0E4hJUov8ITDAZ2 user@host"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.SetUserPublicKeys(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.SetUserPublicKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUserPublicKeys`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.SetUserPublicKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetUserPublicKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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


## SetpropsUserFile

> []ApiResponse SetpropsUserFile(ctx).Path(path).InlineObject9(inlineObject9).Execute()

Set metadata for a file/directory



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
    path := "path_example" // string | Full file/directory path. It must be URL encoded, for example the path \"my dir/àdir/file.txt\" must be sent as \"my%20dir%2F%C3%A0dir%2Ffile.txt\"
    inlineObject9 := *openapiclient.NewInlineObject9() // InlineObject9 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.SetpropsUserFile(context.Background()).Path(path).InlineObject9(inlineObject9).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.SetpropsUserFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetpropsUserFile`: []ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.SetpropsUserFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetpropsUserFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Full file/directory path. It must be URL encoded, for example the path \&quot;my dir/àdir/file.txt\&quot; must be sent as \&quot;my%20dir%2F%C3%A0dir%2Ffile.txt\&quot; | 
 **inlineObject9** | [**InlineObject9**](InlineObject9.md) |  | 

### Return type

[**[]ApiResponse**](ApiResponse.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Streamzip

> *os.File Streamzip(ctx).RequestBody(requestBody).Execute()

Download multiple files and folders as a single zip file



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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.Streamzip(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.Streamzip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Streamzip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.Streamzip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStreamzipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserProfile

> ApiResponse UpdateUserProfile(ctx).UserProfile(userProfile).Execute()

Update user profile



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
    userProfile := *openapiclient.NewUserProfile() // UserProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.UpdateUserProfile(context.Background()).UserProfile(userProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.UpdateUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserProfile`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.UpdateUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userProfile** | [**UserProfile**](UserProfile.md) |  | 

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


## UpdateUserShare

> ApiResponse UpdateUserShare(ctx, id).Share(share).Execute()

Update share



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
    share := *openapiclient.NewShare() // Share | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.UpdateUserShare(context.Background(), id).Share(share).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.UpdateUserShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserShare`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.UpdateUserShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **share** | [**Share**](Share.md) |  | 

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


## ValidateUserTotpSecret

> ApiResponse ValidateUserTotpSecret(ctx).InlineObject7(inlineObject7).Execute()

Validate a one time authentication code



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
    inlineObject7 := *openapiclient.NewInlineObject7() // InlineObject7 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPIsApi.ValidateUserTotpSecret(context.Background()).InlineObject7(inlineObject7).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPIsApi.ValidateUserTotpSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateUserTotpSecret`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPIsApi.ValidateUserTotpSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateUserTotpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject7** | [**InlineObject7**](InlineObject7.md) |  | 

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

