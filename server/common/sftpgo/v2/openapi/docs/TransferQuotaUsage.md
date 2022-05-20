# TransferQuotaUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedUploadDataTransfer** | Pointer to **int64** | The value must be specified as bytes | [optional] 
**UsedDownloadDataTransfer** | Pointer to **int64** | The value must be specified as bytes | [optional] 

## Methods

### NewTransferQuotaUsage

`func NewTransferQuotaUsage() *TransferQuotaUsage`

NewTransferQuotaUsage instantiates a new TransferQuotaUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferQuotaUsageWithDefaults

`func NewTransferQuotaUsageWithDefaults() *TransferQuotaUsage`

NewTransferQuotaUsageWithDefaults instantiates a new TransferQuotaUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedUploadDataTransfer

`func (o *TransferQuotaUsage) GetUsedUploadDataTransfer() int64`

GetUsedUploadDataTransfer returns the UsedUploadDataTransfer field if non-nil, zero value otherwise.

### GetUsedUploadDataTransferOk

`func (o *TransferQuotaUsage) GetUsedUploadDataTransferOk() (*int64, bool)`

GetUsedUploadDataTransferOk returns a tuple with the UsedUploadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedUploadDataTransfer

`func (o *TransferQuotaUsage) SetUsedUploadDataTransfer(v int64)`

SetUsedUploadDataTransfer sets UsedUploadDataTransfer field to given value.

### HasUsedUploadDataTransfer

`func (o *TransferQuotaUsage) HasUsedUploadDataTransfer() bool`

HasUsedUploadDataTransfer returns a boolean if a field has been set.

### GetUsedDownloadDataTransfer

`func (o *TransferQuotaUsage) GetUsedDownloadDataTransfer() int64`

GetUsedDownloadDataTransfer returns the UsedDownloadDataTransfer field if non-nil, zero value otherwise.

### GetUsedDownloadDataTransferOk

`func (o *TransferQuotaUsage) GetUsedDownloadDataTransferOk() (*int64, bool)`

GetUsedDownloadDataTransferOk returns a tuple with the UsedDownloadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDownloadDataTransfer

`func (o *TransferQuotaUsage) SetUsedDownloadDataTransfer(v int64)`

SetUsedDownloadDataTransfer sets UsedDownloadDataTransfer field to given value.

### HasUsedDownloadDataTransfer

`func (o *TransferQuotaUsage) HasUsedDownloadDataTransfer() bool`

HasUsedDownloadDataTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


