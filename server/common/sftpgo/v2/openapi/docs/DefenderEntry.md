# DefenderEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int32** | the score increases whenever a violation is detected, such as an attempt to log in using an incorrect password or invalid username. If the score exceeds the configured threshold, the IP is banned. Omitted for banned IPs | [optional] 
**BanTime** | Pointer to **time.Time** | date time until the IP is banned. For already banned hosts, the ban time is increased each time a new violation is detected. Omitted if the IP is not banned | [optional] 

## Methods

### NewDefenderEntry

`func NewDefenderEntry() *DefenderEntry`

NewDefenderEntry instantiates a new DefenderEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefenderEntryWithDefaults

`func NewDefenderEntryWithDefaults() *DefenderEntry`

NewDefenderEntryWithDefaults instantiates a new DefenderEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DefenderEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefenderEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefenderEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DefenderEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *DefenderEntry) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DefenderEntry) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DefenderEntry) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DefenderEntry) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetScore

`func (o *DefenderEntry) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DefenderEntry) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DefenderEntry) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *DefenderEntry) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetBanTime

`func (o *DefenderEntry) GetBanTime() time.Time`

GetBanTime returns the BanTime field if non-nil, zero value otherwise.

### GetBanTimeOk

`func (o *DefenderEntry) GetBanTimeOk() (*time.Time, bool)`

GetBanTimeOk returns a tuple with the BanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanTime

`func (o *DefenderEntry) SetBanTime(v time.Time)`

SetBanTime sets BanTime field to given value.

### HasBanTime

`func (o *DefenderEntry) HasBanTime() bool`

HasBanTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


