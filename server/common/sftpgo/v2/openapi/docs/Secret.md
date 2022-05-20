# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Set to \&quot;Plain\&quot; to add or update an existing secret, set to \&quot;Redacted\&quot; to preserve the existing value | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**AdditionalData** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **int32** | 1 means encrypted using a master key | [optional] 

## Methods

### NewSecret

`func NewSecret() *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Secret) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Secret) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Secret) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Secret) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPayload

`func (o *Secret) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Secret) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Secret) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Secret) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetKey

`func (o *Secret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Secret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Secret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Secret) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetAdditionalData

`func (o *Secret) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *Secret) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *Secret) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *Secret) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetMode

`func (o *Secret) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Secret) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Secret) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Secret) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


