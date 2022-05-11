# VirtualFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | unique name for this virtual folder | [optional] 
**MappedPath** | Pointer to **string** | absolute filesystem path to use as virtual folder | [optional] 
**Description** | Pointer to **string** | optional description | [optional] 
**UsedQuotaSize** | Pointer to **int64** |  | [optional] 
**UsedQuotaFiles** | Pointer to **int32** |  | [optional] 
**LastQuotaUpdate** | Pointer to **int64** | Last quota update as unix timestamp in milliseconds | [optional] 
**Users** | Pointer to **[]string** | list of usernames associated with this virtual folder | [optional] 
**Filesystem** | Pointer to [**FilesystemConfig**](FilesystemConfig.md) |  | [optional] 
**VirtualPath** | **string** |  | 
**QuotaSize** | Pointer to **int64** | Quota as size in bytes. 0 menas unlimited, -1 means included in user quota. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed | [optional] 
**QuotaFiles** | Pointer to **int32** | Quota as number of files. 0 menas unlimited, , -1 means included in user quota. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed | [optional] 

## Methods

### NewVirtualFolder

`func NewVirtualFolder(virtualPath string, ) *VirtualFolder`

NewVirtualFolder instantiates a new VirtualFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualFolderWithDefaults

`func NewVirtualFolderWithDefaults() *VirtualFolder`

NewVirtualFolderWithDefaults instantiates a new VirtualFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualFolder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualFolder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualFolder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VirtualFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMappedPath

`func (o *VirtualFolder) GetMappedPath() string`

GetMappedPath returns the MappedPath field if non-nil, zero value otherwise.

### GetMappedPathOk

`func (o *VirtualFolder) GetMappedPathOk() (*string, bool)`

GetMappedPathOk returns a tuple with the MappedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedPath

`func (o *VirtualFolder) SetMappedPath(v string)`

SetMappedPath sets MappedPath field to given value.

### HasMappedPath

`func (o *VirtualFolder) HasMappedPath() bool`

HasMappedPath returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualFolder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualFolder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualFolder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualFolder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsedQuotaSize

`func (o *VirtualFolder) GetUsedQuotaSize() int64`

GetUsedQuotaSize returns the UsedQuotaSize field if non-nil, zero value otherwise.

### GetUsedQuotaSizeOk

`func (o *VirtualFolder) GetUsedQuotaSizeOk() (*int64, bool)`

GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaSize

`func (o *VirtualFolder) SetUsedQuotaSize(v int64)`

SetUsedQuotaSize sets UsedQuotaSize field to given value.

### HasUsedQuotaSize

`func (o *VirtualFolder) HasUsedQuotaSize() bool`

HasUsedQuotaSize returns a boolean if a field has been set.

### GetUsedQuotaFiles

`func (o *VirtualFolder) GetUsedQuotaFiles() int32`

GetUsedQuotaFiles returns the UsedQuotaFiles field if non-nil, zero value otherwise.

### GetUsedQuotaFilesOk

`func (o *VirtualFolder) GetUsedQuotaFilesOk() (*int32, bool)`

GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaFiles

`func (o *VirtualFolder) SetUsedQuotaFiles(v int32)`

SetUsedQuotaFiles sets UsedQuotaFiles field to given value.

### HasUsedQuotaFiles

`func (o *VirtualFolder) HasUsedQuotaFiles() bool`

HasUsedQuotaFiles returns a boolean if a field has been set.

### GetLastQuotaUpdate

`func (o *VirtualFolder) GetLastQuotaUpdate() int64`

GetLastQuotaUpdate returns the LastQuotaUpdate field if non-nil, zero value otherwise.

### GetLastQuotaUpdateOk

`func (o *VirtualFolder) GetLastQuotaUpdateOk() (*int64, bool)`

GetLastQuotaUpdateOk returns a tuple with the LastQuotaUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuotaUpdate

`func (o *VirtualFolder) SetLastQuotaUpdate(v int64)`

SetLastQuotaUpdate sets LastQuotaUpdate field to given value.

### HasLastQuotaUpdate

`func (o *VirtualFolder) HasLastQuotaUpdate() bool`

HasLastQuotaUpdate returns a boolean if a field has been set.

### GetUsers

`func (o *VirtualFolder) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *VirtualFolder) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *VirtualFolder) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *VirtualFolder) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetFilesystem

`func (o *VirtualFolder) GetFilesystem() FilesystemConfig`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *VirtualFolder) GetFilesystemOk() (*FilesystemConfig, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *VirtualFolder) SetFilesystem(v FilesystemConfig)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *VirtualFolder) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetVirtualPath

`func (o *VirtualFolder) GetVirtualPath() string`

GetVirtualPath returns the VirtualPath field if non-nil, zero value otherwise.

### GetVirtualPathOk

`func (o *VirtualFolder) GetVirtualPathOk() (*string, bool)`

GetVirtualPathOk returns a tuple with the VirtualPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPath

`func (o *VirtualFolder) SetVirtualPath(v string)`

SetVirtualPath sets VirtualPath field to given value.


### GetQuotaSize

`func (o *VirtualFolder) GetQuotaSize() int64`

GetQuotaSize returns the QuotaSize field if non-nil, zero value otherwise.

### GetQuotaSizeOk

`func (o *VirtualFolder) GetQuotaSizeOk() (*int64, bool)`

GetQuotaSizeOk returns a tuple with the QuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaSize

`func (o *VirtualFolder) SetQuotaSize(v int64)`

SetQuotaSize sets QuotaSize field to given value.

### HasQuotaSize

`func (o *VirtualFolder) HasQuotaSize() bool`

HasQuotaSize returns a boolean if a field has been set.

### GetQuotaFiles

`func (o *VirtualFolder) GetQuotaFiles() int32`

GetQuotaFiles returns the QuotaFiles field if non-nil, zero value otherwise.

### GetQuotaFilesOk

`func (o *VirtualFolder) GetQuotaFilesOk() (*int32, bool)`

GetQuotaFilesOk returns a tuple with the QuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaFiles

`func (o *VirtualFolder) SetQuotaFiles(v int32)`

SetQuotaFiles sets QuotaFiles field to given value.

### HasQuotaFiles

`func (o *VirtualFolder) HasQuotaFiles() bool`

HasQuotaFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


