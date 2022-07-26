# WebDAVBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | TCP address the server listen on | [optional] 
**Port** | Pointer to **int32** | the port used for serving requests | [optional] 
**EnableHttps** | Pointer to **bool** |  | [optional] 
**ClientAuthType** | Pointer to **int32** | 1 means that client certificate authentication is required in addition to HTTP basic authentication | [optional] 

## Methods

### NewWebDAVBinding

`func NewWebDAVBinding() *WebDAVBinding`

NewWebDAVBinding instantiates a new WebDAVBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebDAVBindingWithDefaults

`func NewWebDAVBindingWithDefaults() *WebDAVBinding`

NewWebDAVBindingWithDefaults instantiates a new WebDAVBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *WebDAVBinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WebDAVBinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WebDAVBinding) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WebDAVBinding) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *WebDAVBinding) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WebDAVBinding) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WebDAVBinding) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WebDAVBinding) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEnableHttps

`func (o *WebDAVBinding) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *WebDAVBinding) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *WebDAVBinding) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *WebDAVBinding) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetClientAuthType

`func (o *WebDAVBinding) GetClientAuthType() int32`

GetClientAuthType returns the ClientAuthType field if non-nil, zero value otherwise.

### GetClientAuthTypeOk

`func (o *WebDAVBinding) GetClientAuthTypeOk() (*int32, bool)`

GetClientAuthTypeOk returns a tuple with the ClientAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthType

`func (o *WebDAVBinding) SetClientAuthType(v int32)`

SetClientAuthType sets ClientAuthType field to given value.

### HasClientAuthType

`func (o *WebDAVBinding) HasClientAuthType() bool`

HasClientAuthType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


