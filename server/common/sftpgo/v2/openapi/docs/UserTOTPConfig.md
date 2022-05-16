# UserTOTPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigName** | Pointer to **string** | This name must be defined within the \&quot;totp\&quot; section of the SFTPGo configuration file. You will be unable to save a user/admin referencing a missing config_name | [optional] 
**Secret** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**Protocols** | Pointer to [**[]MFAProtocols**](MFAProtocols.md) | TOTP will be required for the specified protocols. SSH protocol (SFTP/SCP/SSH commands) will ask for the TOTP passcode if the client uses keyboard interactive authentication. FTP has no standard way to support two factor authentication, if you enable the FTP support, you have to add the TOTP passcode after the password. For example if your password is \&quot;password\&quot; and your one time passcode is \&quot;123456\&quot; you have to use \&quot;password123456\&quot; as password. WebDAV is not supported since each single request must be authenticated and a passcode cannot be reused. | [optional] 

## Methods

### NewUserTOTPConfig

`func NewUserTOTPConfig() *UserTOTPConfig`

NewUserTOTPConfig instantiates a new UserTOTPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTOTPConfigWithDefaults

`func NewUserTOTPConfigWithDefaults() *UserTOTPConfig`

NewUserTOTPConfigWithDefaults instantiates a new UserTOTPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UserTOTPConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserTOTPConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserTOTPConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserTOTPConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigName

`func (o *UserTOTPConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *UserTOTPConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *UserTOTPConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *UserTOTPConfig) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetSecret

`func (o *UserTOTPConfig) GetSecret() Secret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UserTOTPConfig) GetSecretOk() (*Secret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UserTOTPConfig) SetSecret(v Secret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UserTOTPConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetProtocols

`func (o *UserTOTPConfig) GetProtocols() []MFAProtocols`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *UserTOTPConfig) GetProtocolsOk() (*[]MFAProtocols, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *UserTOTPConfig) SetProtocols(v []MFAProtocols)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *UserTOTPConfig) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


