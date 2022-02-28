# MFAStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**TotpConfigs** | Pointer to [**[]TOTPConfig**](TOTPConfig.md) |  | [optional] 

## Methods

### NewMFAStatus

`func NewMFAStatus() *MFAStatus`

NewMFAStatus instantiates a new MFAStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAStatusWithDefaults

`func NewMFAStatusWithDefaults() *MFAStatus`

NewMFAStatusWithDefaults instantiates a new MFAStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *MFAStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MFAStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MFAStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MFAStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTotpConfigs

`func (o *MFAStatus) GetTotpConfigs() []TOTPConfig`

GetTotpConfigs returns the TotpConfigs field if non-nil, zero value otherwise.

### GetTotpConfigsOk

`func (o *MFAStatus) GetTotpConfigsOk() (*[]TOTPConfig, bool)`

GetTotpConfigsOk returns a tuple with the TotpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpConfigs

`func (o *MFAStatus) SetTotpConfigs(v []TOTPConfig)`

SetTotpConfigs sets TotpConfigs field to given value.

### HasTotpConfigs

`func (o *MFAStatus) HasTotpConfigs() bool`

HasTotpConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


