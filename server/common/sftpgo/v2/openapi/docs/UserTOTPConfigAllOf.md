# UserTOTPConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocols** | Pointer to [**[]MFAProtocols**](MFAProtocols.md) | TOTP will be required for the specified protocols. SSH protocol (SFTP/SCP/SSH commands) will ask for the TOTP passcode if the client uses keyboard interactive authentication. FTP has no standard way to support two factor authentication, if you enable the FTP support, you have to add the TOTP passcode after the password. For example if your password is \&quot;password\&quot; and your one time passcode is \&quot;123456\&quot; you have to use \&quot;password123456\&quot; as password. WebDAV is not supported since each single request must be authenticated and a passcode cannot be reused. | [optional] 

## Methods

### NewUserTOTPConfigAllOf

`func NewUserTOTPConfigAllOf() *UserTOTPConfigAllOf`

NewUserTOTPConfigAllOf instantiates a new UserTOTPConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTOTPConfigAllOfWithDefaults

`func NewUserTOTPConfigAllOfWithDefaults() *UserTOTPConfigAllOf`

NewUserTOTPConfigAllOfWithDefaults instantiates a new UserTOTPConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocols

`func (o *UserTOTPConfigAllOf) GetProtocols() []MFAProtocols`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *UserTOTPConfigAllOf) GetProtocolsOk() (*[]MFAProtocols, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *UserTOTPConfigAllOf) SetProtocols(v []MFAProtocols)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *UserTOTPConfigAllOf) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


