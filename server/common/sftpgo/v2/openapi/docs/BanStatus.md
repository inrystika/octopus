# BanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTime** | Pointer to **NullableTime** | if null the host is not banned | [optional] 

## Methods

### NewBanStatus

`func NewBanStatus() *BanStatus`

NewBanStatus instantiates a new BanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanStatusWithDefaults

`func NewBanStatusWithDefaults() *BanStatus`

NewBanStatusWithDefaults instantiates a new BanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTime

`func (o *BanStatus) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *BanStatus) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *BanStatus) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *BanStatus) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *BanStatus) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *BanStatus) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


