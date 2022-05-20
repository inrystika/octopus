# DirEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the file (or subdirectory) described by the entry. This name is the final element of the path (the base name), not the entire path | [optional] 
**Size** | Pointer to **int64** | file size, omitted for folders and non regular files | [optional] 
**Mode** | Pointer to **int32** | File mode and permission bits. More details here: https://golang.org/pkg/io/fs/#FileMode. Let&#39;s see some examples: - for a directory mode&amp;2147483648 !&#x3D; 0 - for a symlink mode&amp;134217728 !&#x3D; 0 - for a regular file mode&amp;2401763328 &#x3D;&#x3D; 0  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDirEntry

`func NewDirEntry() *DirEntry`

NewDirEntry instantiates a new DirEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirEntryWithDefaults

`func NewDirEntryWithDefaults() *DirEntry`

NewDirEntryWithDefaults instantiates a new DirEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DirEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DirEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *DirEntry) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DirEntry) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DirEntry) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DirEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMode

`func (o *DirEntry) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DirEntry) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DirEntry) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DirEntry) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetLastModified

`func (o *DirEntry) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DirEntry) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DirEntry) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DirEntry) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


