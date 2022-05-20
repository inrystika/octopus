# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | status:   * &#x60;0&#x60; user is disabled, login is not allowed   * &#x60;1&#x60; user is enabled  | [optional] 
**Username** | Pointer to **string** | username is unique | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | optional description, for example the user full name | [optional] 
**ExpirationDate** | Pointer to **int64** | expiration date as unix timestamp in milliseconds. An expired account cannot login. 0 means no expiration | [optional] 
**Password** | Pointer to **string** | password or public key/SSH user certificate are mandatory. If the password has no known hashing algo prefix it will be stored, by default, using bcrypt, argon2id is supported too. You can send a password hashed as bcrypt ($2a$ prefix), argon2id, pbkdf2 or unix crypt and it will be stored as is. For security reasons this field is omitted when you search/get users | [optional] 
**PublicKeys** | Pointer to **[]string** | Public keys in OpenSSH format. A password or at least one public key/SSH user certificate are mandatory. | [optional] 
**HomeDir** | Pointer to **string** | path to the user home directory. The user cannot upload or download files outside this directory. SFTPGo tries to automatically create this folder if missing. Must be an absolute path | [optional] 
**VirtualFolders** | Pointer to [**[]VirtualFolder**](VirtualFolder.md) | mapping between virtual SFTPGo paths and filesystem paths outside the user home directory. Supported for local filesystem only. If one or more of the specified folders are not inside the dataprovider they will be automatically created. You have to create the folder on the filesystem yourself | [optional] 
**Uid** | Pointer to **int32** | if you run SFTPGo as root user, the created files and directories will be assigned to this uid. 0 means no change, the owner will be the user that runs SFTPGo. Ignored on windows | [optional] 
**Gid** | Pointer to **int32** | if you run SFTPGo as root user, the created files and directories will be assigned to this gid. 0 means no change, the group will be the one of the user that runs SFTPGo. Ignored on windows | [optional] 
**MaxSessions** | Pointer to **int32** | Limit the sessions that a user can open. 0 means unlimited | [optional] 
**QuotaSize** | Pointer to **int64** | Quota as size in bytes. 0 menas unlimited. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed | [optional] 
**QuotaFiles** | Pointer to **int32** | Quota as number of files. 0 menas unlimited. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed | [optional] 
**Permissions** | Pointer to [**[]map[string][]Permission**](map[string][]Permission.md) |  | [optional] 
**UsedQuotaSize** | Pointer to **int64** |  | [optional] 
**UsedQuotaFiles** | Pointer to **int32** |  | [optional] 
**LastQuotaUpdate** | Pointer to **int64** | Last quota update as unix timestamp in milliseconds | [optional] 
**UploadBandwidth** | Pointer to **int32** | Maximum upload bandwidth as KB/s, 0 means unlimited | [optional] 
**DownloadBandwidth** | Pointer to **int32** | Maximum download bandwidth as KB/s, 0 means unlimited | [optional] 
**UploadDataTransfer** | Pointer to **int32** | Maximum data transfer allowed for uploads as MB. 0 means no limit | [optional] 
**DownloadDataTransfer** | Pointer to **int32** | Maximum data transfer allowed for downloads as MB. 0 means no limit | [optional] 
**TotalDataTransfer** | Pointer to **int32** | Maximum total data transfer as MB. 0 means unlimited. You can set a total data transfer instead of the individual values for uploads and downloads | [optional] 
**UsedUploadDataTransfer** | Pointer to **int32** | Uploaded size, as bytes, since the last reset | [optional] 
**UsedDownloadDataTransfer** | Pointer to **int32** | Downloaded size, as bytes, since the last reset | [optional] 
**CreatedAt** | Pointer to **int64** | creation time as unix timestamp in milliseconds. It will be 0 for users created before v2.2.0 | [optional] 
**UpdatedAt** | Pointer to **int64** | last update time as unix timestamp in milliseconds | [optional] 
**LastLogin** | Pointer to **int64** | Last user login as unix timestamp in milliseconds. It is saved at most once every 10 minutes | [optional] 
**Filters** | Pointer to [**UserFilters**](UserFilters.md) |  | [optional] 
**Filesystem** | Pointer to [**FilesystemConfig**](FilesystemConfig.md) |  | [optional] 
**AdditionalInfo** | Pointer to **string** | Free form text field for external systems | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationDate

`func (o *User) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *User) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *User) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *User) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPublicKeys

`func (o *User) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *User) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *User) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *User) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.

### GetHomeDir

`func (o *User) GetHomeDir() string`

GetHomeDir returns the HomeDir field if non-nil, zero value otherwise.

### GetHomeDirOk

`func (o *User) GetHomeDirOk() (*string, bool)`

GetHomeDirOk returns a tuple with the HomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDir

`func (o *User) SetHomeDir(v string)`

SetHomeDir sets HomeDir field to given value.

### HasHomeDir

`func (o *User) HasHomeDir() bool`

HasHomeDir returns a boolean if a field has been set.

### GetVirtualFolders

`func (o *User) GetVirtualFolders() []VirtualFolder`

GetVirtualFolders returns the VirtualFolders field if non-nil, zero value otherwise.

### GetVirtualFoldersOk

`func (o *User) GetVirtualFoldersOk() (*[]VirtualFolder, bool)`

GetVirtualFoldersOk returns a tuple with the VirtualFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualFolders

`func (o *User) SetVirtualFolders(v []VirtualFolder)`

SetVirtualFolders sets VirtualFolders field to given value.

### HasVirtualFolders

`func (o *User) HasVirtualFolders() bool`

HasVirtualFolders returns a boolean if a field has been set.

### GetUid

`func (o *User) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *User) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *User) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *User) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetGid

`func (o *User) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *User) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *User) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *User) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetMaxSessions

`func (o *User) GetMaxSessions() int32`

GetMaxSessions returns the MaxSessions field if non-nil, zero value otherwise.

### GetMaxSessionsOk

`func (o *User) GetMaxSessionsOk() (*int32, bool)`

GetMaxSessionsOk returns a tuple with the MaxSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessions

`func (o *User) SetMaxSessions(v int32)`

SetMaxSessions sets MaxSessions field to given value.

### HasMaxSessions

`func (o *User) HasMaxSessions() bool`

HasMaxSessions returns a boolean if a field has been set.

### GetQuotaSize

`func (o *User) GetQuotaSize() int64`

GetQuotaSize returns the QuotaSize field if non-nil, zero value otherwise.

### GetQuotaSizeOk

`func (o *User) GetQuotaSizeOk() (*int64, bool)`

GetQuotaSizeOk returns a tuple with the QuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaSize

`func (o *User) SetQuotaSize(v int64)`

SetQuotaSize sets QuotaSize field to given value.

### HasQuotaSize

`func (o *User) HasQuotaSize() bool`

HasQuotaSize returns a boolean if a field has been set.

### GetQuotaFiles

`func (o *User) GetQuotaFiles() int32`

GetQuotaFiles returns the QuotaFiles field if non-nil, zero value otherwise.

### GetQuotaFilesOk

`func (o *User) GetQuotaFilesOk() (*int32, bool)`

GetQuotaFilesOk returns a tuple with the QuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaFiles

`func (o *User) SetQuotaFiles(v int32)`

SetQuotaFiles sets QuotaFiles field to given value.

### HasQuotaFiles

`func (o *User) HasQuotaFiles() bool`

HasQuotaFiles returns a boolean if a field has been set.

### GetPermissions

`func (o *User) GetPermissions() []map[string][]Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *User) GetPermissionsOk() (*[]map[string][]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *User) SetPermissions(v []map[string][]Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *User) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUsedQuotaSize

`func (o *User) GetUsedQuotaSize() int64`

GetUsedQuotaSize returns the UsedQuotaSize field if non-nil, zero value otherwise.

### GetUsedQuotaSizeOk

`func (o *User) GetUsedQuotaSizeOk() (*int64, bool)`

GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaSize

`func (o *User) SetUsedQuotaSize(v int64)`

SetUsedQuotaSize sets UsedQuotaSize field to given value.

### HasUsedQuotaSize

`func (o *User) HasUsedQuotaSize() bool`

HasUsedQuotaSize returns a boolean if a field has been set.

### GetUsedQuotaFiles

`func (o *User) GetUsedQuotaFiles() int32`

GetUsedQuotaFiles returns the UsedQuotaFiles field if non-nil, zero value otherwise.

### GetUsedQuotaFilesOk

`func (o *User) GetUsedQuotaFilesOk() (*int32, bool)`

GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaFiles

`func (o *User) SetUsedQuotaFiles(v int32)`

SetUsedQuotaFiles sets UsedQuotaFiles field to given value.

### HasUsedQuotaFiles

`func (o *User) HasUsedQuotaFiles() bool`

HasUsedQuotaFiles returns a boolean if a field has been set.

### GetLastQuotaUpdate

`func (o *User) GetLastQuotaUpdate() int64`

GetLastQuotaUpdate returns the LastQuotaUpdate field if non-nil, zero value otherwise.

### GetLastQuotaUpdateOk

`func (o *User) GetLastQuotaUpdateOk() (*int64, bool)`

GetLastQuotaUpdateOk returns a tuple with the LastQuotaUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuotaUpdate

`func (o *User) SetLastQuotaUpdate(v int64)`

SetLastQuotaUpdate sets LastQuotaUpdate field to given value.

### HasLastQuotaUpdate

`func (o *User) HasLastQuotaUpdate() bool`

HasLastQuotaUpdate returns a boolean if a field has been set.

### GetUploadBandwidth

`func (o *User) GetUploadBandwidth() int32`

GetUploadBandwidth returns the UploadBandwidth field if non-nil, zero value otherwise.

### GetUploadBandwidthOk

`func (o *User) GetUploadBandwidthOk() (*int32, bool)`

GetUploadBandwidthOk returns a tuple with the UploadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadBandwidth

`func (o *User) SetUploadBandwidth(v int32)`

SetUploadBandwidth sets UploadBandwidth field to given value.

### HasUploadBandwidth

`func (o *User) HasUploadBandwidth() bool`

HasUploadBandwidth returns a boolean if a field has been set.

### GetDownloadBandwidth

`func (o *User) GetDownloadBandwidth() int32`

GetDownloadBandwidth returns the DownloadBandwidth field if non-nil, zero value otherwise.

### GetDownloadBandwidthOk

`func (o *User) GetDownloadBandwidthOk() (*int32, bool)`

GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadBandwidth

`func (o *User) SetDownloadBandwidth(v int32)`

SetDownloadBandwidth sets DownloadBandwidth field to given value.

### HasDownloadBandwidth

`func (o *User) HasDownloadBandwidth() bool`

HasDownloadBandwidth returns a boolean if a field has been set.

### GetUploadDataTransfer

`func (o *User) GetUploadDataTransfer() int32`

GetUploadDataTransfer returns the UploadDataTransfer field if non-nil, zero value otherwise.

### GetUploadDataTransferOk

`func (o *User) GetUploadDataTransferOk() (*int32, bool)`

GetUploadDataTransferOk returns a tuple with the UploadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDataTransfer

`func (o *User) SetUploadDataTransfer(v int32)`

SetUploadDataTransfer sets UploadDataTransfer field to given value.

### HasUploadDataTransfer

`func (o *User) HasUploadDataTransfer() bool`

HasUploadDataTransfer returns a boolean if a field has been set.

### GetDownloadDataTransfer

`func (o *User) GetDownloadDataTransfer() int32`

GetDownloadDataTransfer returns the DownloadDataTransfer field if non-nil, zero value otherwise.

### GetDownloadDataTransferOk

`func (o *User) GetDownloadDataTransferOk() (*int32, bool)`

GetDownloadDataTransferOk returns a tuple with the DownloadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDataTransfer

`func (o *User) SetDownloadDataTransfer(v int32)`

SetDownloadDataTransfer sets DownloadDataTransfer field to given value.

### HasDownloadDataTransfer

`func (o *User) HasDownloadDataTransfer() bool`

HasDownloadDataTransfer returns a boolean if a field has been set.

### GetTotalDataTransfer

`func (o *User) GetTotalDataTransfer() int32`

GetTotalDataTransfer returns the TotalDataTransfer field if non-nil, zero value otherwise.

### GetTotalDataTransferOk

`func (o *User) GetTotalDataTransferOk() (*int32, bool)`

GetTotalDataTransferOk returns a tuple with the TotalDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataTransfer

`func (o *User) SetTotalDataTransfer(v int32)`

SetTotalDataTransfer sets TotalDataTransfer field to given value.

### HasTotalDataTransfer

`func (o *User) HasTotalDataTransfer() bool`

HasTotalDataTransfer returns a boolean if a field has been set.

### GetUsedUploadDataTransfer

`func (o *User) GetUsedUploadDataTransfer() int32`

GetUsedUploadDataTransfer returns the UsedUploadDataTransfer field if non-nil, zero value otherwise.

### GetUsedUploadDataTransferOk

`func (o *User) GetUsedUploadDataTransferOk() (*int32, bool)`

GetUsedUploadDataTransferOk returns a tuple with the UsedUploadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedUploadDataTransfer

`func (o *User) SetUsedUploadDataTransfer(v int32)`

SetUsedUploadDataTransfer sets UsedUploadDataTransfer field to given value.

### HasUsedUploadDataTransfer

`func (o *User) HasUsedUploadDataTransfer() bool`

HasUsedUploadDataTransfer returns a boolean if a field has been set.

### GetUsedDownloadDataTransfer

`func (o *User) GetUsedDownloadDataTransfer() int32`

GetUsedDownloadDataTransfer returns the UsedDownloadDataTransfer field if non-nil, zero value otherwise.

### GetUsedDownloadDataTransferOk

`func (o *User) GetUsedDownloadDataTransferOk() (*int32, bool)`

GetUsedDownloadDataTransferOk returns a tuple with the UsedDownloadDataTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDownloadDataTransfer

`func (o *User) SetUsedDownloadDataTransfer(v int32)`

SetUsedDownloadDataTransfer sets UsedDownloadDataTransfer field to given value.

### HasUsedDownloadDataTransfer

`func (o *User) HasUsedDownloadDataTransfer() bool`

HasUsedDownloadDataTransfer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() int64`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*int64, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v int64)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetFilters

`func (o *User) GetFilters() UserFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *User) GetFiltersOk() (*UserFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *User) SetFilters(v UserFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *User) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFilesystem

`func (o *User) GetFilesystem() FilesystemConfig`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *User) GetFilesystemOk() (*FilesystemConfig, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *User) SetFilesystem(v FilesystemConfig)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *User) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *User) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *User) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *User) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *User) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


