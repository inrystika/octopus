# ProviderEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** | unix timestamp in nanoseconds | [optional] 
**Action** | Pointer to [**ProviderEventAction**](ProviderEventAction.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to [**ProviderEventObjectType**](ProviderEventObjectType.md) |  | [optional] 
**ObjectName** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 

## Methods

### NewProviderEvent

`func NewProviderEvent() *ProviderEvent`

NewProviderEvent instantiates a new ProviderEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderEventWithDefaults

`func NewProviderEventWithDefaults() *ProviderEvent`

NewProviderEventWithDefaults instantiates a new ProviderEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProviderEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProviderEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProviderEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProviderEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAction

`func (o *ProviderEvent) GetAction() ProviderEventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProviderEvent) GetActionOk() (*ProviderEventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProviderEvent) SetAction(v ProviderEventAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProviderEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUsername

`func (o *ProviderEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProviderEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProviderEvent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProviderEvent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIp

`func (o *ProviderEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ProviderEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ProviderEvent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ProviderEvent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetObjectType

`func (o *ProviderEvent) GetObjectType() ProviderEventObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProviderEvent) GetObjectTypeOk() (*ProviderEventObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProviderEvent) SetObjectType(v ProviderEventObjectType)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProviderEvent) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectName

`func (o *ProviderEvent) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ProviderEvent) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ProviderEvent) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ProviderEvent) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetInstanceId

`func (o *ProviderEvent) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ProviderEvent) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ProviderEvent) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ProviderEvent) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


