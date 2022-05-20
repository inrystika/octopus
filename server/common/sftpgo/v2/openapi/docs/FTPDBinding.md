# FTPDBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | TCP address the server listen on | [optional] 
**Port** | Pointer to **int32** | the port used for serving requests | [optional] 
**ApplyProxyConfig** | Pointer to **bool** | apply the proxy configuration, if any | [optional] 
**TlsMode** | Pointer to **int32** | TLS mode:   * &#x60;0&#x60; - clear or explicit TLS   * &#x60;1&#x60; - explicit TLS required   * &#x60;2&#x60; - implicit TLS  | [optional] 
**ForcePassiveIp** | Pointer to **string** | External IP address to expose for passive connections | [optional] 
**ClientAuthType** | Pointer to **int32** | 1 means that client certificate authentication is required in addition to FTP authentication | [optional] 

## Methods

### NewFTPDBinding

`func NewFTPDBinding() *FTPDBinding`

NewFTPDBinding instantiates a new FTPDBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTPDBindingWithDefaults

`func NewFTPDBindingWithDefaults() *FTPDBinding`

NewFTPDBindingWithDefaults instantiates a new FTPDBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FTPDBinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FTPDBinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FTPDBinding) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FTPDBinding) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *FTPDBinding) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FTPDBinding) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FTPDBinding) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FTPDBinding) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetApplyProxyConfig

`func (o *FTPDBinding) GetApplyProxyConfig() bool`

GetApplyProxyConfig returns the ApplyProxyConfig field if non-nil, zero value otherwise.

### GetApplyProxyConfigOk

`func (o *FTPDBinding) GetApplyProxyConfigOk() (*bool, bool)`

GetApplyProxyConfigOk returns a tuple with the ApplyProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyProxyConfig

`func (o *FTPDBinding) SetApplyProxyConfig(v bool)`

SetApplyProxyConfig sets ApplyProxyConfig field to given value.

### HasApplyProxyConfig

`func (o *FTPDBinding) HasApplyProxyConfig() bool`

HasApplyProxyConfig returns a boolean if a field has been set.

### GetTlsMode

`func (o *FTPDBinding) GetTlsMode() int32`

GetTlsMode returns the TlsMode field if non-nil, zero value otherwise.

### GetTlsModeOk

`func (o *FTPDBinding) GetTlsModeOk() (*int32, bool)`

GetTlsModeOk returns a tuple with the TlsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMode

`func (o *FTPDBinding) SetTlsMode(v int32)`

SetTlsMode sets TlsMode field to given value.

### HasTlsMode

`func (o *FTPDBinding) HasTlsMode() bool`

HasTlsMode returns a boolean if a field has been set.

### GetForcePassiveIp

`func (o *FTPDBinding) GetForcePassiveIp() string`

GetForcePassiveIp returns the ForcePassiveIp field if non-nil, zero value otherwise.

### GetForcePassiveIpOk

`func (o *FTPDBinding) GetForcePassiveIpOk() (*string, bool)`

GetForcePassiveIpOk returns a tuple with the ForcePassiveIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePassiveIp

`func (o *FTPDBinding) SetForcePassiveIp(v string)`

SetForcePassiveIp sets ForcePassiveIp field to given value.

### HasForcePassiveIp

`func (o *FTPDBinding) HasForcePassiveIp() bool`

HasForcePassiveIp returns a boolean if a field has been set.

### GetClientAuthType

`func (o *FTPDBinding) GetClientAuthType() int32`

GetClientAuthType returns the ClientAuthType field if non-nil, zero value otherwise.

### GetClientAuthTypeOk

`func (o *FTPDBinding) GetClientAuthTypeOk() (*int32, bool)`

GetClientAuthTypeOk returns a tuple with the ClientAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthType

`func (o *FTPDBinding) SetClientAuthType(v int32)`

SetClientAuthType sets ClientAuthType field to given value.

### HasClientAuthType

`func (o *FTPDBinding) HasClientAuthType() bool`

HasClientAuthType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


