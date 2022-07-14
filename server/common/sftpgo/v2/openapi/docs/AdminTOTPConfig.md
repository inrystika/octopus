# AdminTOTPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigName** | Pointer to **string** | This name must be defined within the \&quot;totp\&quot; section of the SFTPGo configuration file. You will be unable to save a user/admin referencing a missing config_name | [optional] 
**Secret** | Pointer to [**Secret**](Secret.md) |  | [optional] 

## Methods

### NewAdminTOTPConfig

`func NewAdminTOTPConfig() *AdminTOTPConfig`

NewAdminTOTPConfig instantiates a new AdminTOTPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminTOTPConfigWithDefaults

`func NewAdminTOTPConfigWithDefaults() *AdminTOTPConfig`

NewAdminTOTPConfigWithDefaults instantiates a new AdminTOTPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AdminTOTPConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdminTOTPConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdminTOTPConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AdminTOTPConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigName

`func (o *AdminTOTPConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *AdminTOTPConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *AdminTOTPConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *AdminTOTPConfig) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetSecret

`func (o *AdminTOTPConfig) GetSecret() Secret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AdminTOTPConfig) GetSecretOk() (*Secret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AdminTOTPConfig) SetSecret(v Secret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AdminTOTPConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


