# \AdminsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdmin**](AdminsApi.md#AddAdmin) | **Post** /admins | Add admin
[**AdminForgotPassword**](AdminsApi.md#AdminForgotPassword) | **Post** /admins/{username}/forgot-password | Send a password reset code by email
[**AdminResetPassword**](AdminsApi.md#AdminResetPassword) | **Post** /admins/{username}/reset-password | Reset the password
[**ChangeAdminPassword**](AdminsApi.md#ChangeAdminPassword) | **Put** /admin/changepwd | Change admin password
[**ChangeAdminPasswordDeprecated**](AdminsApi.md#ChangeAdminPasswordDeprecated) | **Put** /changepwd/admin | Change admin password
[**DeleteAdmin**](AdminsApi.md#DeleteAdmin) | **Delete** /admins/{username} | Delete admin
[**DisableAdmin2fa**](AdminsApi.md#DisableAdmin2fa) | **Put** /admins/{username}/2fa/disable | Disable second factor authentication
[**GenerateAdminRecoveryCodes**](AdminsApi.md#GenerateAdminRecoveryCodes) | **Post** /admin/2fa/recoverycodes | Generate recovery codes
[**GenerateAdminTotpSecret**](AdminsApi.md#GenerateAdminTotpSecret) | **Post** /admin/totp/generate | Generate a new TOTP secret
[**GetAdminByUsername**](AdminsApi.md#GetAdminByUsername) | **Get** /admins/{username} | Find admins by username
[**GetAdminProfile**](AdminsApi.md#GetAdminProfile) | **Get** /admin/profile | Get admin profile
[**GetAdminRecoveryCodes**](AdminsApi.md#GetAdminRecoveryCodes) | **Get** /admin/2fa/recoverycodes | Get recovery codes
[**GetAdminTotpConfigs**](AdminsApi.md#GetAdminTotpConfigs) | **Get** /admin/totp/configs | Get available TOTP configuration
[**GetAdmins**](AdminsApi.md#GetAdmins) | **Get** /admins | Get admins
[**SaveAdminTotpConfig**](AdminsApi.md#SaveAdminTotpConfig) | **Post** /admin/totp/save | Save a TOTP config
[**UpdateAdmin**](AdminsApi.md#UpdateAdmin) | **Put** /admins/{username} | Update admin
[**UpdateAdminProfile**](AdminsApi.md#UpdateAdminProfile) | **Put** /admin/profile | Update admin profile
[**ValidateAdminTotpSecret**](AdminsApi.md#ValidateAdminTotpSecret) | **Post** /admin/totp/validate | Validate a one time authentication code



## AddAdmin

> Admin AddAdmin(ctx).Admin(admin).Execute()

Add admin



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
    admin := *openapiclient.NewAdmin() // Admin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.AddAdmin(context.Background()).Admin(admin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.AddAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAdmin`: Admin
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.AddAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **admin** | [**Admin**](Admin.md) |  | 

### Return type

[**Admin**](Admin.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminForgotPassword

> ApiResponse AdminForgotPassword(ctx, username).Execute()

Send a password reset code by email



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
    username := "username_example" // string | the admin username

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.AdminForgotPassword(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.AdminForgotPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminForgotPassword`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.AdminForgotPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the admin username | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminResetPassword

> ApiResponse AdminResetPassword(ctx, username).InlineObject4(inlineObject4).Execute()

Reset the password



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
    username := "username_example" // string | the admin username
    inlineObject4 := *openapiclient.NewInlineObject4() // InlineObject4 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.AdminResetPassword(context.Background(), username).InlineObject4(inlineObject4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.AdminResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminResetPassword`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.AdminResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the admin username | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject4** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeAdminPassword

> ApiResponse ChangeAdminPassword(ctx).PwdChange(pwdChange).Execute()

Change admin password



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
    resp, r, err := apiClient.AdminsApi.ChangeAdminPassword(context.Background()).PwdChange(pwdChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.ChangeAdminPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeAdminPassword`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.ChangeAdminPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeAdminPasswordRequest struct via the builder pattern


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


## ChangeAdminPasswordDeprecated

> ApiResponse ChangeAdminPasswordDeprecated(ctx).PwdChange(pwdChange).Execute()

Change admin password



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
    resp, r, err := apiClient.AdminsApi.ChangeAdminPasswordDeprecated(context.Background()).PwdChange(pwdChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.ChangeAdminPasswordDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeAdminPasswordDeprecated`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.ChangeAdminPasswordDeprecated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeAdminPasswordDeprecatedRequest struct via the builder pattern


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


## DeleteAdmin

> ApiResponse DeleteAdmin(ctx, username).Execute()

Delete admin



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
    username := "username_example" // string | the admin username

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.DeleteAdmin(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.DeleteAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAdmin`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.DeleteAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the admin username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminRequest struct via the builder pattern


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


## DisableAdmin2fa

> ApiResponse DisableAdmin2fa(ctx, username).Execute()

Disable second factor authentication



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
    username := "username_example" // string | the admin username

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.DisableAdmin2fa(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.DisableAdmin2fa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableAdmin2fa`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.DisableAdmin2fa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the admin username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableAdmin2faRequest struct via the builder pattern


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


## GenerateAdminRecoveryCodes

> []string GenerateAdminRecoveryCodes(ctx).Execute()

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
    resp, r, err := apiClient.AdminsApi.GenerateAdminRecoveryCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GenerateAdminRecoveryCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAdminRecoveryCodes`: []string
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GenerateAdminRecoveryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAdminRecoveryCodesRequest struct via the builder pattern


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


## GenerateAdminTotpSecret

> InlineResponse200 GenerateAdminTotpSecret(ctx).InlineObject1(inlineObject1).Execute()

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
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.GenerateAdminTotpSecret(context.Background()).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GenerateAdminTotpSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAdminTotpSecret`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GenerateAdminTotpSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAdminTotpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

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


## GetAdminByUsername

> Admin GetAdminByUsername(ctx, username).Execute()

Find admins by username



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
    username := "username_example" // string | the admin username

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.GetAdminByUsername(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetAdminByUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminByUsername`: Admin
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetAdminByUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the admin username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminByUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Admin**](Admin.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminProfile

> AdminProfile GetAdminProfile(ctx).Execute()

Get admin profile



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
    resp, r, err := apiClient.AdminsApi.GetAdminProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetAdminProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminProfile`: AdminProfile
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetAdminProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminProfileRequest struct via the builder pattern


### Return type

[**AdminProfile**](AdminProfile.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecoveryCodes

> []RecoveryCode GetAdminRecoveryCodes(ctx).Execute()

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
    resp, r, err := apiClient.AdminsApi.GetAdminRecoveryCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetAdminRecoveryCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecoveryCodes`: []RecoveryCode
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetAdminRecoveryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecoveryCodesRequest struct via the builder pattern


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


## GetAdminTotpConfigs

> []TOTPConfig GetAdminTotpConfigs(ctx).Execute()

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
    resp, r, err := apiClient.AdminsApi.GetAdminTotpConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetAdminTotpConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminTotpConfigs`: []TOTPConfig
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetAdminTotpConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminTotpConfigsRequest struct via the builder pattern


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


## GetAdmins

> []Admin GetAdmins(ctx).Offset(offset).Limit(limit).Order(order).Execute()

Get admins



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
    order := "ASC" // string | Ordering admins by username. Default ASC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.GetAdmins(context.Background()).Offset(offset).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdmins`: []Admin
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetAdmins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** | The maximum number of items to return. Max value is 500, default is 100 | [default to 100]
 **order** | **string** | Ordering admins by username. Default ASC | 

### Return type

[**[]Admin**](Admin.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAdminTotpConfig

> ApiResponse SaveAdminTotpConfig(ctx).AdminTOTPConfig(adminTOTPConfig).Execute()

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
    adminTOTPConfig := *openapiclient.NewAdminTOTPConfig() // AdminTOTPConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.SaveAdminTotpConfig(context.Background()).AdminTOTPConfig(adminTOTPConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.SaveAdminTotpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveAdminTotpConfig`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.SaveAdminTotpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAdminTotpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminTOTPConfig** | [**AdminTOTPConfig**](AdminTOTPConfig.md) |  | 

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


## UpdateAdmin

> ApiResponse UpdateAdmin(ctx, username).Admin(admin).Execute()

Update admin



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
    username := "username_example" // string | the admin username
    admin := *openapiclient.NewAdmin() // Admin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.UpdateAdmin(context.Background(), username).Admin(admin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.UpdateAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdmin`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.UpdateAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the admin username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **admin** | [**Admin**](Admin.md) |  | 

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


## UpdateAdminProfile

> ApiResponse UpdateAdminProfile(ctx).AdminProfile(adminProfile).Execute()

Update admin profile



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
    adminProfile := *openapiclient.NewAdminProfile() // AdminProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.UpdateAdminProfile(context.Background()).AdminProfile(adminProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.UpdateAdminProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdminProfile`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.UpdateAdminProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminProfile** | [**AdminProfile**](AdminProfile.md) |  | 

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


## ValidateAdminTotpSecret

> ApiResponse ValidateAdminTotpSecret(ctx).InlineObject2(inlineObject2).Execute()

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
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.ValidateAdminTotpSecret(context.Background()).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.ValidateAdminTotpSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAdminTotpSecret`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.ValidateAdminTotpSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAdminTotpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

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

