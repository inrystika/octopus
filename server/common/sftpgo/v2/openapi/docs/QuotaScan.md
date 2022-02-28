# QuotaScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | username to which the quota scan refers | [optional] 
**StartTime** | Pointer to **int64** | scan start time as unix timestamp in milliseconds | [optional] 

## Methods

### NewQuotaScan

`func NewQuotaScan() *QuotaScan`

NewQuotaScan instantiates a new QuotaScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaScanWithDefaults

`func NewQuotaScanWithDefaults() *QuotaScan`

NewQuotaScanWithDefaults instantiates a new QuotaScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *QuotaScan) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *QuotaScan) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *QuotaScan) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *QuotaScan) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStartTime

`func (o *QuotaScan) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QuotaScan) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QuotaScan) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *QuotaScan) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


