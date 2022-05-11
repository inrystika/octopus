# DataProviderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewDataProviderStatus

`func NewDataProviderStatus() *DataProviderStatus`

NewDataProviderStatus instantiates a new DataProviderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProviderStatusWithDefaults

`func NewDataProviderStatusWithDefaults() *DataProviderStatus`

NewDataProviderStatusWithDefaults instantiates a new DataProviderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *DataProviderStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DataProviderStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DataProviderStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DataProviderStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDriver

`func (o *DataProviderStatus) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *DataProviderStatus) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *DataProviderStatus) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *DataProviderStatus) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetError

`func (o *DataProviderStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataProviderStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataProviderStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DataProviderStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


