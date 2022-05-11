# QuotaUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedQuotaSize** | Pointer to **int64** |  | [optional] 
**UsedQuotaFiles** | Pointer to **int32** |  | [optional] 

## Methods

### NewQuotaUsage

`func NewQuotaUsage() *QuotaUsage`

NewQuotaUsage instantiates a new QuotaUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaUsageWithDefaults

`func NewQuotaUsageWithDefaults() *QuotaUsage`

NewQuotaUsageWithDefaults instantiates a new QuotaUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedQuotaSize

`func (o *QuotaUsage) GetUsedQuotaSize() int64`

GetUsedQuotaSize returns the UsedQuotaSize field if non-nil, zero value otherwise.

### GetUsedQuotaSizeOk

`func (o *QuotaUsage) GetUsedQuotaSizeOk() (*int64, bool)`

GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaSize

`func (o *QuotaUsage) SetUsedQuotaSize(v int64)`

SetUsedQuotaSize sets UsedQuotaSize field to given value.

### HasUsedQuotaSize

`func (o *QuotaUsage) HasUsedQuotaSize() bool`

HasUsedQuotaSize returns a boolean if a field has been set.

### GetUsedQuotaFiles

`func (o *QuotaUsage) GetUsedQuotaFiles() int32`

GetUsedQuotaFiles returns the UsedQuotaFiles field if non-nil, zero value otherwise.

### GetUsedQuotaFilesOk

`func (o *QuotaUsage) GetUsedQuotaFilesOk() (*int32, bool)`

GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaFiles

`func (o *QuotaUsage) SetUsedQuotaFiles(v int32)`

SetUsedQuotaFiles sets UsedQuotaFiles field to given value.

### HasUsedQuotaFiles

`func (o *QuotaUsage) HasUsedQuotaFiles() bool`

HasUsedQuotaFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


