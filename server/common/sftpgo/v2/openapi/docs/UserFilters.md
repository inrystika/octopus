# UserFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedIp** | Pointer to **[]string** | only clients connecting from these IP/Mask are allowed. IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291, for example \&quot;192.0.2.0/24\&quot; or \&quot;2001:db8::/32\&quot; | [optional] 
**DeniedIp** | Pointer to **[]string** | clients connecting from these IP/Mask are not allowed. Denied rules are evaluated before allowed ones | [optional] 
**DeniedLoginMethods** | Pointer to [**[]LoginMethods**](LoginMethods.md) | if null or empty any available login method is allowed | [optional] 
**DeniedProtocols** | Pointer to [**[]SupportedProtocols**](SupportedProtocols.md) | if null or empty any available protocol is allowed | [optional] 
**FilePatterns** | Pointer to [**[]PatternsFilter**](PatternsFilter.md) | filters based on shell like file patterns. These restrictions do not apply to files listing for performance reasons, so a denied file cannot be downloaded/overwritten/renamed but it will still be in the list of files. Please note that these restrictions can be easily bypassed | [optional] 
**MaxUploadFileSize** | Pointer to **int64** | maximum allowed size, as bytes, for a single file upload. The upload will be aborted if/when the size of the file being sent exceeds this limit. 0 means unlimited. This restriction does not apply for SSH system commands such as &#x60;git&#x60; and &#x60;rsync&#x60; | [optional] 
**TlsUsername** | Pointer to **string** | defines the TLS certificate field to use as username. For FTP clients it must match the name provided using the \&quot;USER\&quot; command. For WebDAV, if no username is provided, the CN will be used as username. For WebDAV clients it must match the implicit or provided username. Ignored if mutual TLS is disabled | [optional] 
**Hooks** | Pointer to [**HooksFilter**](HooksFilter.md) |  | [optional] 
**DisableFsChecks** | Pointer to **bool** | Disable checks for existence and automatic creation of home directory and virtual folders. SFTPGo requires that the user&#39;s home directory, virtual folder root, and intermediate paths to virtual folders exist to work properly. If you already know that the required directories exist, disabling these checks will speed up login. You could, for example, disable these checks after the first login | [optional] 
**WebClient** | Pointer to [**[]WebClientOptions**](WebClientOptions.md) | WebClient/user REST API related configuration options | [optional] 
**AllowApiKeyAuth** | Pointer to **bool** | API key authentication allows to impersonate this user with an API key | [optional] 
**UserType** | Pointer to [**UserType**](UserType.md) |  | [optional] 
**TotpConfig** | Pointer to [**UserTOTPConfig**](UserTOTPConfig.md) |  | [optional] 
**RecoveryCodes** | Pointer to [**[]RecoveryCode**](RecoveryCode.md) |  | [optional] 
**BandwidthLimits** | Pointer to [**[]BandwidthLimit**](BandwidthLimit.md) |  | [optional] 
**DataTransferLimits** | Pointer to [**[]DataTransferLimit**](DataTransferLimit.md) |  | [optional] 

## Methods

### NewUserFilters

`func NewUserFilters() *UserFilters`

NewUserFilters instantiates a new UserFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFiltersWithDefaults

`func NewUserFiltersWithDefaults() *UserFilters`

NewUserFiltersWithDefaults instantiates a new UserFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedIp

`func (o *UserFilters) GetAllowedIp() []string`

GetAllowedIp returns the AllowedIp field if non-nil, zero value otherwise.

### GetAllowedIpOk

`func (o *UserFilters) GetAllowedIpOk() (*[]string, bool)`

GetAllowedIpOk returns a tuple with the AllowedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIp

`func (o *UserFilters) SetAllowedIp(v []string)`

SetAllowedIp sets AllowedIp field to given value.

### HasAllowedIp

`func (o *UserFilters) HasAllowedIp() bool`

HasAllowedIp returns a boolean if a field has been set.

### GetDeniedIp

`func (o *UserFilters) GetDeniedIp() []string`

GetDeniedIp returns the DeniedIp field if non-nil, zero value otherwise.

### GetDeniedIpOk

`func (o *UserFilters) GetDeniedIpOk() (*[]string, bool)`

GetDeniedIpOk returns a tuple with the DeniedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedIp

`func (o *UserFilters) SetDeniedIp(v []string)`

SetDeniedIp sets DeniedIp field to given value.

### HasDeniedIp

`func (o *UserFilters) HasDeniedIp() bool`

HasDeniedIp returns a boolean if a field has been set.

### GetDeniedLoginMethods

`func (o *UserFilters) GetDeniedLoginMethods() []LoginMethods`

GetDeniedLoginMethods returns the DeniedLoginMethods field if non-nil, zero value otherwise.

### GetDeniedLoginMethodsOk

`func (o *UserFilters) GetDeniedLoginMethodsOk() (*[]LoginMethods, bool)`

GetDeniedLoginMethodsOk returns a tuple with the DeniedLoginMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedLoginMethods

`func (o *UserFilters) SetDeniedLoginMethods(v []LoginMethods)`

SetDeniedLoginMethods sets DeniedLoginMethods field to given value.

### HasDeniedLoginMethods

`func (o *UserFilters) HasDeniedLoginMethods() bool`

HasDeniedLoginMethods returns a boolean if a field has been set.

### GetDeniedProtocols

`func (o *UserFilters) GetDeniedProtocols() []SupportedProtocols`

GetDeniedProtocols returns the DeniedProtocols field if non-nil, zero value otherwise.

### GetDeniedProtocolsOk

`func (o *UserFilters) GetDeniedProtocolsOk() (*[]SupportedProtocols, bool)`

GetDeniedProtocolsOk returns a tuple with the DeniedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedProtocols

`func (o *UserFilters) SetDeniedProtocols(v []SupportedProtocols)`

SetDeniedProtocols sets DeniedProtocols field to given value.

### HasDeniedProtocols

`func (o *UserFilters) HasDeniedProtocols() bool`

HasDeniedProtocols returns a boolean if a field has been set.

### GetFilePatterns

`func (o *UserFilters) GetFilePatterns() []PatternsFilter`

GetFilePatterns returns the FilePatterns field if non-nil, zero value otherwise.

### GetFilePatternsOk

`func (o *UserFilters) GetFilePatternsOk() (*[]PatternsFilter, bool)`

GetFilePatternsOk returns a tuple with the FilePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePatterns

`func (o *UserFilters) SetFilePatterns(v []PatternsFilter)`

SetFilePatterns sets FilePatterns field to given value.

### HasFilePatterns

`func (o *UserFilters) HasFilePatterns() bool`

HasFilePatterns returns a boolean if a field has been set.

### GetMaxUploadFileSize

`func (o *UserFilters) GetMaxUploadFileSize() int64`

GetMaxUploadFileSize returns the MaxUploadFileSize field if non-nil, zero value otherwise.

### GetMaxUploadFileSizeOk

`func (o *UserFilters) GetMaxUploadFileSizeOk() (*int64, bool)`

GetMaxUploadFileSizeOk returns a tuple with the MaxUploadFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUploadFileSize

`func (o *UserFilters) SetMaxUploadFileSize(v int64)`

SetMaxUploadFileSize sets MaxUploadFileSize field to given value.

### HasMaxUploadFileSize

`func (o *UserFilters) HasMaxUploadFileSize() bool`

HasMaxUploadFileSize returns a boolean if a field has been set.

### GetTlsUsername

`func (o *UserFilters) GetTlsUsername() string`

GetTlsUsername returns the TlsUsername field if non-nil, zero value otherwise.

### GetTlsUsernameOk

`func (o *UserFilters) GetTlsUsernameOk() (*string, bool)`

GetTlsUsernameOk returns a tuple with the TlsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsUsername

`func (o *UserFilters) SetTlsUsername(v string)`

SetTlsUsername sets TlsUsername field to given value.

### HasTlsUsername

`func (o *UserFilters) HasTlsUsername() bool`

HasTlsUsername returns a boolean if a field has been set.

### GetHooks

`func (o *UserFilters) GetHooks() HooksFilter`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *UserFilters) GetHooksOk() (*HooksFilter, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *UserFilters) SetHooks(v HooksFilter)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *UserFilters) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetDisableFsChecks

`func (o *UserFilters) GetDisableFsChecks() bool`

GetDisableFsChecks returns the DisableFsChecks field if non-nil, zero value otherwise.

### GetDisableFsChecksOk

`func (o *UserFilters) GetDisableFsChecksOk() (*bool, bool)`

GetDisableFsChecksOk returns a tuple with the DisableFsChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFsChecks

`func (o *UserFilters) SetDisableFsChecks(v bool)`

SetDisableFsChecks sets DisableFsChecks field to given value.

### HasDisableFsChecks

`func (o *UserFilters) HasDisableFsChecks() bool`

HasDisableFsChecks returns a boolean if a field has been set.

### GetWebClient

`func (o *UserFilters) GetWebClient() []WebClientOptions`

GetWebClient returns the WebClient field if non-nil, zero value otherwise.

### GetWebClientOk

`func (o *UserFilters) GetWebClientOk() (*[]WebClientOptions, bool)`

GetWebClientOk returns a tuple with the WebClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebClient

`func (o *UserFilters) SetWebClient(v []WebClientOptions)`

SetWebClient sets WebClient field to given value.

### HasWebClient

`func (o *UserFilters) HasWebClient() bool`

HasWebClient returns a boolean if a field has been set.

### GetAllowApiKeyAuth

`func (o *UserFilters) GetAllowApiKeyAuth() bool`

GetAllowApiKeyAuth returns the AllowApiKeyAuth field if non-nil, zero value otherwise.

### GetAllowApiKeyAuthOk

`func (o *UserFilters) GetAllowApiKeyAuthOk() (*bool, bool)`

GetAllowApiKeyAuthOk returns a tuple with the AllowApiKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiKeyAuth

`func (o *UserFilters) SetAllowApiKeyAuth(v bool)`

SetAllowApiKeyAuth sets AllowApiKeyAuth field to given value.

### HasAllowApiKeyAuth

`func (o *UserFilters) HasAllowApiKeyAuth() bool`

HasAllowApiKeyAuth returns a boolean if a field has been set.

### GetUserType

`func (o *UserFilters) GetUserType() UserType`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserFilters) GetUserTypeOk() (*UserType, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserFilters) SetUserType(v UserType)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserFilters) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetTotpConfig

`func (o *UserFilters) GetTotpConfig() UserTOTPConfig`

GetTotpConfig returns the TotpConfig field if non-nil, zero value otherwise.

### GetTotpConfigOk

`func (o *UserFilters) GetTotpConfigOk() (*UserTOTPConfig, bool)`

GetTotpConfigOk returns a tuple with the TotpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpConfig

`func (o *UserFilters) SetTotpConfig(v UserTOTPConfig)`

SetTotpConfig sets TotpConfig field to given value.

### HasTotpConfig

`func (o *UserFilters) HasTotpConfig() bool`

HasTotpConfig returns a boolean if a field has been set.

### GetRecoveryCodes

`func (o *UserFilters) GetRecoveryCodes() []RecoveryCode`

GetRecoveryCodes returns the RecoveryCodes field if non-nil, zero value otherwise.

### GetRecoveryCodesOk

`func (o *UserFilters) GetRecoveryCodesOk() (*[]RecoveryCode, bool)`

GetRecoveryCodesOk returns a tuple with the RecoveryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodes

`func (o *UserFilters) SetRecoveryCodes(v []RecoveryCode)`

SetRecoveryCodes sets RecoveryCodes field to given value.

### HasRecoveryCodes

`func (o *UserFilters) HasRecoveryCodes() bool`

HasRecoveryCodes returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *UserFilters) GetBandwidthLimits() []BandwidthLimit`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *UserFilters) GetBandwidthLimitsOk() (*[]BandwidthLimit, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *UserFilters) SetBandwidthLimits(v []BandwidthLimit)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *UserFilters) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.

### GetDataTransferLimits

`func (o *UserFilters) GetDataTransferLimits() []DataTransferLimit`

GetDataTransferLimits returns the DataTransferLimits field if non-nil, zero value otherwise.

### GetDataTransferLimitsOk

`func (o *UserFilters) GetDataTransferLimitsOk() (*[]DataTransferLimit, bool)`

GetDataTransferLimitsOk returns a tuple with the DataTransferLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferLimits

`func (o *UserFilters) SetDataTransferLimits(v []DataTransferLimit)`

SetDataTransferLimits sets DataTransferLimits field to given value.

### HasDataTransferLimits

`func (o *UserFilters) HasDataTransferLimits() bool`

HasDataTransferLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


