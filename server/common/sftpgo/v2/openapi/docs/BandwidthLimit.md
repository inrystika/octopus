# BandwidthLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Source networks in CIDR notation as defined in RFC 4632 and RFC 4291 for example &#x60;192.0.2.0/24&#x60; or &#x60;2001:db8::/32&#x60;. The limit applies if the defined networks contain the client IP | [optional] 
**UploadBandwidth** | Pointer to **int32** | Maximum upload bandwidth as KB/s, 0 means unlimited | [optional] 
**DownloadBandwidth** | Pointer to **int32** | Maximum download bandwidth as KB/s, 0 means unlimited | [optional] 

## Methods

### NewBandwidthLimit

`func NewBandwidthLimit() *BandwidthLimit`

NewBandwidthLimit instantiates a new BandwidthLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthLimitWithDefaults

`func NewBandwidthLimitWithDefaults() *BandwidthLimit`

NewBandwidthLimitWithDefaults instantiates a new BandwidthLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *BandwidthLimit) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BandwidthLimit) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BandwidthLimit) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BandwidthLimit) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetUploadBandwidth

`func (o *BandwidthLimit) GetUploadBandwidth() int32`

GetUploadBandwidth returns the UploadBandwidth field if non-nil, zero value otherwise.

### GetUploadBandwidthOk

`func (o *BandwidthLimit) GetUploadBandwidthOk() (*int32, bool)`

GetUploadBandwidthOk returns a tuple with the UploadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadBandwidth

`func (o *BandwidthLimit) SetUploadBandwidth(v int32)`

SetUploadBandwidth sets UploadBandwidth field to given value.

### HasUploadBandwidth

`func (o *BandwidthLimit) HasUploadBandwidth() bool`

HasUploadBandwidth returns a boolean if a field has been set.

### GetDownloadBandwidth

`func (o *BandwidthLimit) GetDownloadBandwidth() int32`

GetDownloadBandwidth returns the DownloadBandwidth field if non-nil, zero value otherwise.

### GetDownloadBandwidthOk

`func (o *BandwidthLimit) GetDownloadBandwidthOk() (*int32, bool)`

GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadBandwidth

`func (o *BandwidthLimit) SetDownloadBandwidth(v int32)`

SetDownloadBandwidth sets DownloadBandwidth field to given value.

### HasDownloadBandwidth

`func (o *BandwidthLimit) HasDownloadBandwidth() bool`

HasDownloadBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


