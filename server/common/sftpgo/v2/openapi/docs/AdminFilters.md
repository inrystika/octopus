# AdminFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowList** | Pointer to **[]string** | only clients connecting from these IP/Mask are allowed. IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291, for example \&quot;192.0.2.0/24\&quot; or \&quot;2001:db8::/32\&quot; | [optional] 
**AllowApiKeyAuth** | Pointer to **bool** | API key auth allows to impersonate this administrator with an API key | [optional] 
**TotpConfig** | Pointer to [**AdminTOTPConfig**](AdminTOTPConfig.md) |  | [optional] 
**RecoveryCodes** | Pointer to [**[]RecoveryCode**](RecoveryCode.md) |  | [optional] 

## Methods

### NewAdminFilters

`func NewAdminFilters() *AdminFilters`

NewAdminFilters instantiates a new AdminFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFiltersWithDefaults

`func NewAdminFiltersWithDefaults() *AdminFilters`

NewAdminFiltersWithDefaults instantiates a new AdminFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowList

`func (o *AdminFilters) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *AdminFilters) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *AdminFilters) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *AdminFilters) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetAllowApiKeyAuth

`func (o *AdminFilters) GetAllowApiKeyAuth() bool`

GetAllowApiKeyAuth returns the AllowApiKeyAuth field if non-nil, zero value otherwise.

### GetAllowApiKeyAuthOk

`func (o *AdminFilters) GetAllowApiKeyAuthOk() (*bool, bool)`

GetAllowApiKeyAuthOk returns a tuple with the AllowApiKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiKeyAuth

`func (o *AdminFilters) SetAllowApiKeyAuth(v bool)`

SetAllowApiKeyAuth sets AllowApiKeyAuth field to given value.

### HasAllowApiKeyAuth

`func (o *AdminFilters) HasAllowApiKeyAuth() bool`

HasAllowApiKeyAuth returns a boolean if a field has been set.

### GetTotpConfig

`func (o *AdminFilters) GetTotpConfig() AdminTOTPConfig`

GetTotpConfig returns the TotpConfig field if non-nil, zero value otherwise.

### GetTotpConfigOk

`func (o *AdminFilters) GetTotpConfigOk() (*AdminTOTPConfig, bool)`

GetTotpConfigOk returns a tuple with the TotpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpConfig

`func (o *AdminFilters) SetTotpConfig(v AdminTOTPConfig)`

SetTotpConfig sets TotpConfig field to given value.

### HasTotpConfig

`func (o *AdminFilters) HasTotpConfig() bool`

HasTotpConfig returns a boolean if a field has been set.

### GetRecoveryCodes

`func (o *AdminFilters) GetRecoveryCodes() []RecoveryCode`

GetRecoveryCodes returns the RecoveryCodes field if non-nil, zero value otherwise.

### GetRecoveryCodesOk

`func (o *AdminFilters) GetRecoveryCodesOk() (*[]RecoveryCode, bool)`

GetRecoveryCodesOk returns a tuple with the RecoveryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCodes

`func (o *AdminFilters) SetRecoveryCodes(v []RecoveryCode)`

SetRecoveryCodes sets RecoveryCodes field to given value.

### HasRecoveryCodes

`func (o *AdminFilters) HasRecoveryCodes() bool`

HasRecoveryCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


