# SSHServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**Bindings** | Pointer to [**[]SSHBinding**](SSHBinding.md) |  | [optional] 
**HostKeys** | Pointer to [**[]SSHHostKey**](SSHHostKey.md) |  | [optional] 
**SshCommands** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSSHServiceStatus

`func NewSSHServiceStatus() *SSHServiceStatus`

NewSSHServiceStatus instantiates a new SSHServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHServiceStatusWithDefaults

`func NewSSHServiceStatusWithDefaults() *SSHServiceStatus`

NewSSHServiceStatusWithDefaults instantiates a new SSHServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *SSHServiceStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SSHServiceStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SSHServiceStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SSHServiceStatus) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetBindings

`func (o *SSHServiceStatus) GetBindings() []SSHBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *SSHServiceStatus) GetBindingsOk() (*[]SSHBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *SSHServiceStatus) SetBindings(v []SSHBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *SSHServiceStatus) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindingsNil

`func (o *SSHServiceStatus) SetBindingsNil(b bool)`

 SetBindingsNil sets the value for Bindings to be an explicit nil

### UnsetBindings
`func (o *SSHServiceStatus) UnsetBindings()`

UnsetBindings ensures that no value is present for Bindings, not even an explicit nil
### GetHostKeys

`func (o *SSHServiceStatus) GetHostKeys() []SSHHostKey`

GetHostKeys returns the HostKeys field if non-nil, zero value otherwise.

### GetHostKeysOk

`func (o *SSHServiceStatus) GetHostKeysOk() (*[]SSHHostKey, bool)`

GetHostKeysOk returns a tuple with the HostKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostKeys

`func (o *SSHServiceStatus) SetHostKeys(v []SSHHostKey)`

SetHostKeys sets HostKeys field to given value.

### HasHostKeys

`func (o *SSHServiceStatus) HasHostKeys() bool`

HasHostKeys returns a boolean if a field has been set.

### SetHostKeysNil

`func (o *SSHServiceStatus) SetHostKeysNil(b bool)`

 SetHostKeysNil sets the value for HostKeys to be an explicit nil

### UnsetHostKeys
`func (o *SSHServiceStatus) UnsetHostKeys()`

UnsetHostKeys ensures that no value is present for HostKeys, not even an explicit nil
### GetSshCommands

`func (o *SSHServiceStatus) GetSshCommands() []string`

GetSshCommands returns the SshCommands field if non-nil, zero value otherwise.

### GetSshCommandsOk

`func (o *SSHServiceStatus) GetSshCommandsOk() (*[]string, bool)`

GetSshCommandsOk returns a tuple with the SshCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCommands

`func (o *SSHServiceStatus) SetSshCommands(v []string)`

SetSshCommands sets SshCommands field to given value.

### HasSshCommands

`func (o *SSHServiceStatus) HasSshCommands() bool`

HasSshCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


