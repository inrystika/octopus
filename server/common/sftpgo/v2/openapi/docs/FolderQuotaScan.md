# FolderQuotaScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | folder name to which the quota scan refers | [optional] 
**StartTime** | Pointer to **int64** | scan start time as unix timestamp in milliseconds | [optional] 

## Methods

### NewFolderQuotaScan

`func NewFolderQuotaScan() *FolderQuotaScan`

NewFolderQuotaScan instantiates a new FolderQuotaScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderQuotaScanWithDefaults

`func NewFolderQuotaScanWithDefaults() *FolderQuotaScan`

NewFolderQuotaScanWithDefaults instantiates a new FolderQuotaScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FolderQuotaScan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderQuotaScan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderQuotaScan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FolderQuotaScan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *FolderQuotaScan) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FolderQuotaScan) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FolderQuotaScan) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FolderQuotaScan) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


