# BaseTOTPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigName** | Pointer to **string** | This name must be defined within the \&quot;totp\&quot; section of the SFTPGo configuration file. You will be unable to save a user/admin referencing a missing config_name | [optional] 
**Secret** | Pointer to [**Secret**](Secret.md) |  | [optional] 

## Methods

### NewBaseTOTPConfig

`func NewBaseTOTPConfig() *BaseTOTPConfig`

NewBaseTOTPConfig instantiates a new BaseTOTPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTOTPConfigWithDefaults

`func NewBaseTOTPConfigWithDefaults() *BaseTOTPConfig`

NewBaseTOTPConfigWithDefaults instantiates a new BaseTOTPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BaseTOTPConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseTOTPConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseTOTPConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseTOTPConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigName

`func (o *BaseTOTPConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *BaseTOTPConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *BaseTOTPConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *BaseTOTPConfig) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetSecret

`func (o *BaseTOTPConfig) GetSecret() Secret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BaseTOTPConfig) GetSecretOk() (*Secret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BaseTOTPConfig) SetSecret(v Secret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BaseTOTPConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


