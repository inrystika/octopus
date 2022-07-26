# DataTransferLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Source networks in CIDR notation as defined in RFC 4632 and RFC 4291 for example &#x60;192.0.2.0/24&#x60; or &#x60;2001:db8::/32&#x60;. The limit applies if the defined networks contain the client IP | [optional] 
**UploadDataTransfer** | Pointer to **int32** | Maximum data transfer allowed for uploads as MB. 0 means no limit | [optional] 
**DownloadDataTransfer** | Pointer to **int32** | Maximum data transfer allowed for downloads as MB. 0 means no limit | [optional] 
**TotalDataTransfer** | Pointer to **int32** | Maximum total data transfer as MB. 0 means unlimited. You can set a total data transfer instead of the individual values for uploads and downloads | [optional] 

## Methods

### NewDataTransferLimit

`func NewDataTransferLimit() *DataTransferLimit`

NewDataTransferLimit instantiates a new DataTransferLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferLimitWithDefaults

`func NewDataTransferLimitWithDefaults() *DataTransferLimit`

NewDataTransferLimitWithDefaults instantiates a new DataTransferLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *DataTransferLimit) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DataTransferLimit) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DataTransferLimit) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DataTransferLimit) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetUploadDataTransfer

`func (o *DataTransferLimit) GetUploadDataTransfer() int32`

GetUploadDataTransfer returns the UploadDataTransfer field if non-nil, zero value otherwise.

### GetUploadDataTransferOk

`func (o *DataTransferLimit) GetUploadDataTransferOk() (*int32, bool)`

GetUploadDataTransferOk returns a tuple with the UploadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDataTransfer

`func (o *DataTransferLimit) SetUploadDataTransfer(v int32)`

SetUploadDataTransfer sets UploadDataTransfer field to given value.

### HasUploadDataTransfer

`func (o *DataTransferLimit) HasUploadDataTransfer() bool`

HasUploadDataTransfer returns a boolean if a field has been set.

### GetDownloadDataTransfer

`func (o *DataTransferLimit) GetDownloadDataTransfer() int32`

GetDownloadDataTransfer returns the DownloadDataTransfer field if non-nil, zero value otherwise.

### GetDownloadDataTransferOk

`func (o *DataTransferLimit) GetDownloadDataTransferOk() (*int32, bool)`

GetDownloadDataTransferOk returns a tuple with the DownloadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDataTransfer

`func (o *DataTransferLimit) SetDownloadDataTransfer(v int32)`

SetDownloadDataTransfer sets DownloadDataTransfer field to given value.

### HasDownloadDataTransfer

`func (o *DataTransferLimit) HasDownloadDataTransfer() bool`

HasDownloadDataTransfer returns a boolean if a field has been set.

### GetTotalDataTransfer

`func (o *DataTransferLimit) GetTotalDataTransfer() int32`

GetTotalDataTransfer returns the TotalDataTransfer field if non-nil, zero value otherwise.

### GetTotalDataTransferOk

`func (o *DataTransferLimit) GetTotalDataTransferOk() (*int32, bool)`

GetTotalDataTransferOk returns a tuple with the TotalDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataTransfer

`func (o *DataTransferLimit) SetTotalDataTransfer(v int32)`

SetTotalDataTransfer sets TotalDataTransfer field to given value.

### HasTotalDataTransfer

`func (o *DataTransferLimit) HasTotalDataTransfer() bool`

HasTotalDataTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


