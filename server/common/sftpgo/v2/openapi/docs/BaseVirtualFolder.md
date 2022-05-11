# BaseVirtualFolder

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

## Methods

### NewBaseVirtualFolder

`func NewBaseVirtualFolder() *BaseVirtualFolder`

NewBaseVirtualFolder instantiates a new BaseVirtualFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseVirtualFolderWithDefaults

`func NewBaseVirtualFolderWithDefaults() *BaseVirtualFolder`

NewBaseVirtualFolderWithDefaults instantiates a new BaseVirtualFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseVirtualFolder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseVirtualFolder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseVirtualFolder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BaseVirtualFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseVirtualFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseVirtualFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseVirtualFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseVirtualFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMappedPath

`func (o *BaseVirtualFolder) GetMappedPath() string`

GetMappedPath returns the MappedPath field if non-nil, zero value otherwise.

### GetMappedPathOk

`func (o *BaseVirtualFolder) GetMappedPathOk() (*string, bool)`

GetMappedPathOk returns a tuple with the MappedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedPath

`func (o *BaseVirtualFolder) SetMappedPath(v string)`

SetMappedPath sets MappedPath field to given value.

### HasMappedPath

`func (o *BaseVirtualFolder) HasMappedPath() bool`

HasMappedPath returns a boolean if a field has been set.

### GetDescription

`func (o *BaseVirtualFolder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseVirtualFolder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseVirtualFolder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseVirtualFolder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsedQuotaSize

`func (o *BaseVirtualFolder) GetUsedQuotaSize() int64`

GetUsedQuotaSize returns the UsedQuotaSize field if non-nil, zero value otherwise.

### GetUsedQuotaSizeOk

`func (o *BaseVirtualFolder) GetUsedQuotaSizeOk() (*int64, bool)`

GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaSize

`func (o *BaseVirtualFolder) SetUsedQuotaSize(v int64)`

SetUsedQuotaSize sets UsedQuotaSize field to given value.

### HasUsedQuotaSize

`func (o *BaseVirtualFolder) HasUsedQuotaSize() bool`

HasUsedQuotaSize returns a boolean if a field has been set.

### GetUsedQuotaFiles

`func (o *BaseVirtualFolder) GetUsedQuotaFiles() int32`

GetUsedQuotaFiles returns the UsedQuotaFiles field if non-nil, zero value otherwise.

### GetUsedQuotaFilesOk

`func (o *BaseVirtualFolder) GetUsedQuotaFilesOk() (*int32, bool)`

GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuotaFiles

`func (o *BaseVirtualFolder) SetUsedQuotaFiles(v int32)`

SetUsedQuotaFiles sets UsedQuotaFiles field to given value.

### HasUsedQuotaFiles

`func (o *BaseVirtualFolder) HasUsedQuotaFiles() bool`

HasUsedQuotaFiles returns a boolean if a field has been set.

### GetLastQuotaUpdate

`func (o *BaseVirtualFolder) GetLastQuotaUpdate() int64`

GetLastQuotaUpdate returns the LastQuotaUpdate field if non-nil, zero value otherwise.

### GetLastQuotaUpdateOk

`func (o *BaseVirtualFolder) GetLastQuotaUpdateOk() (*int64, bool)`

GetLastQuotaUpdateOk returns a tuple with the LastQuotaUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuotaUpdate

`func (o *BaseVirtualFolder) SetLastQuotaUpdate(v int64)`

SetLastQuotaUpdate sets LastQuotaUpdate field to given value.

### HasLastQuotaUpdate

`func (o *BaseVirtualFolder) HasLastQuotaUpdate() bool`

HasLastQuotaUpdate returns a boolean if a field has been set.

### GetUsers

`func (o *BaseVirtualFolder) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BaseVirtualFolder) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BaseVirtualFolder) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BaseVirtualFolder) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetFilesystem

`func (o *BaseVirtualFolder) GetFilesystem() FilesystemConfig`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *BaseVirtualFolder) GetFilesystemOk() (*FilesystemConfig, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *BaseVirtualFolder) SetFilesystem(v FilesystemConfig)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *BaseVirtualFolder) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


