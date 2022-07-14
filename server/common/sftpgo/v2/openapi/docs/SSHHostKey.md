# SSHHostKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 

## Methods

### NewSSHHostKey

`func NewSSHHostKey() *SSHHostKey`

NewSSHHostKey instantiates a new SSHHostKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHHostKeyWithDefaults

`func NewSSHHostKeyWithDefaults() *SSHHostKey`

NewSSHHostKeyWithDefaults instantiates a new SSHHostKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SSHHostKey) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SSHHostKey) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SSHHostKey) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SSHHostKey) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetFingerprint

`func (o *SSHHostKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SSHHostKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SSHHostKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SSHHostKey) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


