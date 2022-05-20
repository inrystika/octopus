# SSHBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | TCP address the server listen on | [optional] 
**Port** | Pointer to **int32** | the port used for serving requests | [optional] 
**ApplyProxyConfig** | Pointer to **bool** | apply the proxy configuration, if any | [optional] 

## Methods

### NewSSHBinding

`func NewSSHBinding() *SSHBinding`

NewSSHBinding instantiates a new SSHBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHBindingWithDefaults

`func NewSSHBindingWithDefaults() *SSHBinding`

NewSSHBindingWithDefaults instantiates a new SSHBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SSHBinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SSHBinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SSHBinding) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SSHBinding) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *SSHBinding) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SSHBinding) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SSHBinding) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SSHBinding) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetApplyProxyConfig

`func (o *SSHBinding) GetApplyProxyConfig() bool`

GetApplyProxyConfig returns the ApplyProxyConfig field if non-nil, zero value otherwise.

### GetApplyProxyConfigOk

`func (o *SSHBinding) GetApplyProxyConfigOk() (*bool, bool)`

GetApplyProxyConfigOk returns a tuple with the ApplyProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyProxyConfig

`func (o *SSHBinding) SetApplyProxyConfig(v bool)`

SetApplyProxyConfig sets ApplyProxyConfig field to given value.

### HasApplyProxyConfig

`func (o *SSHBinding) HasApplyProxyConfig() bool`

HasApplyProxyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


