# VirtualFolderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualPath** | **string** |  | 
**QuotaSize** | Pointer to **int64** | Quota as size in bytes. 0 menas unlimited, -1 means included in user quota. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed | [optional] 
**QuotaFiles** | Pointer to **int32** | Quota as number of files. 0 menas unlimited, , -1 means included in user quota. Please note that quota is updated if files are added/removed via SFTPGo otherwise a quota scan or a manual quota update is needed | [optional] 

## Methods

### NewVirtualFolderAllOf

`func NewVirtualFolderAllOf(virtualPath string, ) *VirtualFolderAllOf`

NewVirtualFolderAllOf instantiates a new VirtualFolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualFolderAllOfWithDefaults

`func NewVirtualFolderAllOfWithDefaults() *VirtualFolderAllOf`

NewVirtualFolderAllOfWithDefaults instantiates a new VirtualFolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualPath

`func (o *VirtualFolderAllOf) GetVirtualPath() string`

GetVirtualPath returns the VirtualPath field if non-nil, zero value otherwise.

### GetVirtualPathOk

`func (o *VirtualFolderAllOf) GetVirtualPathOk() (*string, bool)`

GetVirtualPathOk returns a tuple with the VirtualPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPath

`func (o *VirtualFolderAllOf) SetVirtualPath(v string)`

SetVirtualPath sets VirtualPath field to given value.


### GetQuotaSize

`func (o *VirtualFolderAllOf) GetQuotaSize() int64`

GetQuotaSize returns the QuotaSize field if non-nil, zero value otherwise.

### GetQuotaSizeOk

`func (o *VirtualFolderAllOf) GetQuotaSizeOk() (*int64, bool)`

GetQuotaSizeOk returns a tuple with the QuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaSize

`func (o *VirtualFolderAllOf) SetQuotaSize(v int64)`

SetQuotaSize sets QuotaSize field to given value.

### HasQuotaSize

`func (o *VirtualFolderAllOf) HasQuotaSize() bool`

HasQuotaSize returns a boolean if a field has been set.

### GetQuotaFiles

`func (o *VirtualFolderAllOf) GetQuotaFiles() int32`

GetQuotaFiles returns the QuotaFiles field if non-nil, zero value otherwise.

### GetQuotaFilesOk

`func (o *VirtualFolderAllOf) GetQuotaFilesOk() (*int32, bool)`

GetQuotaFilesOk returns a tuple with the QuotaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaFiles

`func (o *VirtualFolderAllOf) SetQuotaFiles(v int32)`

SetQuotaFiles sets QuotaFiles field to given value.

### HasQuotaFiles

`func (o *VirtualFolderAllOf) HasQuotaFiles() bool`

HasQuotaFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


