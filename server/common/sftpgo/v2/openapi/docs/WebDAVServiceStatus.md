# WebDAVServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**Bindings** | Pointer to [**[]WebDAVBinding**](WebDAVBinding.md) |  | [optional] 

## Methods

### NewWebDAVServiceStatus

`func NewWebDAVServiceStatus() *WebDAVServiceStatus`

NewWebDAVServiceStatus instantiates a new WebDAVServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebDAVServiceStatusWithDefaults

`func NewWebDAVServiceStatusWithDefaults() *WebDAVServiceStatus`

NewWebDAVServiceStatusWithDefaults instantiates a new WebDAVServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *WebDAVServiceStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WebDAVServiceStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WebDAVServiceStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WebDAVServiceStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetBindings

`func (o *WebDAVServiceStatus) GetBindings() []WebDAVBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *WebDAVServiceStatus) GetBindingsOk() (*[]WebDAVBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *WebDAVServiceStatus) SetBindings(v []WebDAVBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *WebDAVServiceStatus) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindingsNil

`func (o *WebDAVServiceStatus) SetBindingsNil(b bool)`

 SetBindingsNil sets the value for Bindings to be an explicit nil

### UnsetBindings
`func (o *WebDAVServiceStatus) UnsetBindings()`

UnsetBindings ensures that no value is present for Bindings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


