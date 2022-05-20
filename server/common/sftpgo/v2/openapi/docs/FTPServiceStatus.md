# FTPServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**Bindings** | Pointer to [**[]FTPDBinding**](FTPDBinding.md) |  | [optional] 
**PassivePortRange** | Pointer to [**FTPPassivePortRange**](FTPPassivePortRange.md) |  | [optional] 

## Methods

### NewFTPServiceStatus

`func NewFTPServiceStatus() *FTPServiceStatus`

NewFTPServiceStatus instantiates a new FTPServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTPServiceStatusWithDefaults

`func NewFTPServiceStatusWithDefaults() *FTPServiceStatus`

NewFTPServiceStatusWithDefaults instantiates a new FTPServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *FTPServiceStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *FTPServiceStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *FTPServiceStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *FTPServiceStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetBindings

`func (o *FTPServiceStatus) GetBindings() []FTPDBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *FTPServiceStatus) GetBindingsOk() (*[]FTPDBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *FTPServiceStatus) SetBindings(v []FTPDBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *FTPServiceStatus) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindingsNil

`func (o *FTPServiceStatus) SetBindingsNil(b bool)`

 SetBindingsNil sets the value for Bindings to be an explicit nil

### UnsetBindings
`func (o *FTPServiceStatus) UnsetBindings()`

UnsetBindings ensures that no value is present for Bindings, not even an explicit nil
### GetPassivePortRange

`func (o *FTPServiceStatus) GetPassivePortRange() FTPPassivePortRange`

GetPassivePortRange returns the PassivePortRange field if non-nil, zero value otherwise.

### GetPassivePortRangeOk

`func (o *FTPServiceStatus) GetPassivePortRangeOk() (*FTPPassivePortRange, bool)`

GetPassivePortRangeOk returns a tuple with the PassivePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassivePortRange

`func (o *FTPServiceStatus) SetPassivePortRange(v FTPPassivePortRange)`

SetPassivePortRange sets PassivePortRange field to given value.

### HasPassivePortRange

`func (o *FTPServiceStatus) HasPassivePortRange() bool`

HasPassivePortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


