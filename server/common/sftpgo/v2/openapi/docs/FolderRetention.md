# FolderRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | exposed virtual directory path, if no other specific retention is defined, the retention applies for sub directories too. For example if retention is defined for the paths \&quot;/\&quot; and \&quot;/sub\&quot; then the retention for \&quot;/\&quot; is applied for any file outside the \&quot;/sub\&quot; directory | [optional] 
**Retention** | Pointer to **int32** | retention time in hours. All the files with a modification time older than the defined value will be deleted. 0 means exclude this path | [optional] 
**DeleteEmptyDirs** | Pointer to **bool** | if enabled, empty directories will be deleted | [optional] 
**IgnoreUserPermissions** | Pointer to **bool** | if enabled, files will be deleted even if the user does not have the delete permission. The default is \&quot;false\&quot; which means that files will be skipped if the user does not have permission to delete them. File patterns filters will always be silently ignored | [optional] 

## Methods

### NewFolderRetention

`func NewFolderRetention() *FolderRetention`

NewFolderRetention instantiates a new FolderRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderRetentionWithDefaults

`func NewFolderRetentionWithDefaults() *FolderRetention`

NewFolderRetentionWithDefaults instantiates a new FolderRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FolderRetention) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FolderRetention) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FolderRetention) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FolderRetention) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRetention

`func (o *FolderRetention) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *FolderRetention) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *FolderRetention) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *FolderRetention) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetDeleteEmptyDirs

`func (o *FolderRetention) GetDeleteEmptyDirs() bool`

GetDeleteEmptyDirs returns the DeleteEmptyDirs field if non-nil, zero value otherwise.

### GetDeleteEmptyDirsOk

`func (o *FolderRetention) GetDeleteEmptyDirsOk() (*bool, bool)`

GetDeleteEmptyDirsOk returns a tuple with the DeleteEmptyDirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEmptyDirs

`func (o *FolderRetention) SetDeleteEmptyDirs(v bool)`

SetDeleteEmptyDirs sets DeleteEmptyDirs field to given value.

### HasDeleteEmptyDirs

`func (o *FolderRetention) HasDeleteEmptyDirs() bool`

HasDeleteEmptyDirs returns a boolean if a field has been set.

### GetIgnoreUserPermissions

`func (o *FolderRetention) GetIgnoreUserPermissions() bool`

GetIgnoreUserPermissions returns the IgnoreUserPermissions field if non-nil, zero value otherwise.

### GetIgnoreUserPermissionsOk

`func (o *FolderRetention) GetIgnoreUserPermissionsOk() (*bool, bool)`

GetIgnoreUserPermissionsOk returns a tuple with the IgnoreUserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUserPermissions

`func (o *FolderRetention) SetIgnoreUserPermissions(v bool)`

SetIgnoreUserPermissions sets IgnoreUserPermissions field to given value.

### HasIgnoreUserPermissions

`func (o *FolderRetention) HasIgnoreUserPermissions() bool`

HasIgnoreUserPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


