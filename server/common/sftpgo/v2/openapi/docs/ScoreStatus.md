# ScoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int32** | if 0 the host is not listed | [optional] 

## Methods

### NewScoreStatus

`func NewScoreStatus() *ScoreStatus`

NewScoreStatus instantiates a new ScoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreStatusWithDefaults

`func NewScoreStatusWithDefaults() *ScoreStatus`

NewScoreStatusWithDefaults instantiates a new ScoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ScoreStatus) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScoreStatus) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScoreStatus) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ScoreStatus) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


