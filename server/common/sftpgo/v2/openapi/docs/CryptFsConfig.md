# CryptFsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passphrase** | Pointer to [**Secret**](Secret.md) |  | [optional] 

## Methods

### NewCryptFsConfig

`func NewCryptFsConfig() *CryptFsConfig`

NewCryptFsConfig instantiates a new CryptFsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptFsConfigWithDefaults

`func NewCryptFsConfigWithDefaults() *CryptFsConfig`

NewCryptFsConfigWithDefaults instantiates a new CryptFsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassphrase

`func (o *CryptFsConfig) GetPassphrase() Secret`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CryptFsConfig) GetPassphraseOk() (*Secret, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CryptFsConfig) SetPassphrase(v Secret)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CryptFsConfig) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


