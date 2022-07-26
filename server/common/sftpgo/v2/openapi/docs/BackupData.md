# BackupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 
**Folders** | Pointer to [**[]BaseVirtualFolder**](BaseVirtualFolder.md) |  | [optional] 
**Admins** | Pointer to [**[]Admin**](Admin.md) |  | [optional] 
**ApiKeys** | Pointer to [**[]AuthAPIKey**](AuthAPIKey.md) |  | [optional] 
**Shares** | Pointer to [**[]Share**](Share.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewBackupData

`func NewBackupData() *BackupData`

NewBackupData instantiates a new BackupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDataWithDefaults

`func NewBackupDataWithDefaults() *BackupData`

NewBackupDataWithDefaults instantiates a new BackupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *BackupData) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BackupData) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BackupData) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BackupData) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetFolders

`func (o *BackupData) GetFolders() []BaseVirtualFolder`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *BackupData) GetFoldersOk() (*[]BaseVirtualFolder, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *BackupData) SetFolders(v []BaseVirtualFolder)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *BackupData) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetAdmins

`func (o *BackupData) GetAdmins() []Admin`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *BackupData) GetAdminsOk() (*[]Admin, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *BackupData) SetAdmins(v []Admin)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *BackupData) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetApiKeys

`func (o *BackupData) GetApiKeys() []AuthAPIKey`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *BackupData) GetApiKeysOk() (*[]AuthAPIKey, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *BackupData) SetApiKeys(v []AuthAPIKey)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *BackupData) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetShares

`func (o *BackupData) GetShares() []Share`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *BackupData) GetSharesOk() (*[]Share, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *BackupData) SetShares(v []Share)`

SetShares sets Shares field to given value.

### HasShares

`func (o *BackupData) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetVersion

`func (o *BackupData) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BackupData) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BackupData) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BackupData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


