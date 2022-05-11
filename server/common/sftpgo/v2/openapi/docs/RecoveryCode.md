# RecoveryCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**Used** | Pointer to **bool** |  | [optional] 

## Methods

### NewRecoveryCode

`func NewRecoveryCode() *RecoveryCode`

NewRecoveryCode instantiates a new RecoveryCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryCodeWithDefaults

`func NewRecoveryCodeWithDefaults() *RecoveryCode`

NewRecoveryCodeWithDefaults instantiates a new RecoveryCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *RecoveryCode) GetSecret() Secret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RecoveryCode) GetSecretOk() (*Secret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RecoveryCode) SetSecret(v Secret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *RecoveryCode) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUsed

`func (o *RecoveryCode) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *RecoveryCode) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *RecoveryCode) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *RecoveryCode) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


