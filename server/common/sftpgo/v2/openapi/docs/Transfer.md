# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | Pointer to **string** | Operations:   * &#x60;upload&#x60;   * &#x60;download&#x60;  | [optional] 
**Path** | Pointer to **string** | file path for the upload/download | [optional] 
**StartTime** | Pointer to **int64** | start time as unix timestamp in milliseconds | [optional] 
**Size** | Pointer to **int64** | bytes transferred | [optional] 

## Methods

### NewTransfer

`func NewTransfer() *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *Transfer) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *Transfer) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *Transfer) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *Transfer) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetPath

`func (o *Transfer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Transfer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Transfer) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Transfer) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStartTime

`func (o *Transfer) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Transfer) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Transfer) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Transfer) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSize

`func (o *Transfer) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Transfer) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Transfer) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Transfer) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


