# FilesystemConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**FsProviders**](FsProviders.md) |  | [optional] 
**S3config** | Pointer to [**S3Config**](S3Config.md) |  | [optional] 
**Gcsconfig** | Pointer to [**GCSConfig**](GCSConfig.md) |  | [optional] 
**Azblobconfig** | Pointer to [**AzureBlobFsConfig**](AzureBlobFsConfig.md) |  | [optional] 
**Cryptconfig** | Pointer to [**CryptFsConfig**](CryptFsConfig.md) |  | [optional] 
**Sftpconfig** | Pointer to [**SFTPFsConfig**](SFTPFsConfig.md) |  | [optional] 

## Methods

### NewFilesystemConfig

`func NewFilesystemConfig() *FilesystemConfig`

NewFilesystemConfig instantiates a new FilesystemConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemConfigWithDefaults

`func NewFilesystemConfigWithDefaults() *FilesystemConfig`

NewFilesystemConfigWithDefaults instantiates a new FilesystemConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *FilesystemConfig) GetProvider() FsProviders`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FilesystemConfig) GetProviderOk() (*FsProviders, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FilesystemConfig) SetProvider(v FsProviders)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *FilesystemConfig) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetS3config

`func (o *FilesystemConfig) GetS3config() S3Config`

GetS3config returns the S3config field if non-nil, zero value otherwise.

### GetS3configOk

`func (o *FilesystemConfig) GetS3configOk() (*S3Config, bool)`

GetS3configOk returns a tuple with the S3config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3config

`func (o *FilesystemConfig) SetS3config(v S3Config)`

SetS3config sets S3config field to given value.

### HasS3config

`func (o *FilesystemConfig) HasS3config() bool`

HasS3config returns a boolean if a field has been set.

### GetGcsconfig

`func (o *FilesystemConfig) GetGcsconfig() GCSConfig`

GetGcsconfig returns the Gcsconfig field if non-nil, zero value otherwise.

### GetGcsconfigOk

`func (o *FilesystemConfig) GetGcsconfigOk() (*GCSConfig, bool)`

GetGcsconfigOk returns a tuple with the Gcsconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsconfig

`func (o *FilesystemConfig) SetGcsconfig(v GCSConfig)`

SetGcsconfig sets Gcsconfig field to given value.

### HasGcsconfig

`func (o *FilesystemConfig) HasGcsconfig() bool`

HasGcsconfig returns a boolean if a field has been set.

### GetAzblobconfig

`func (o *FilesystemConfig) GetAzblobconfig() AzureBlobFsConfig`

GetAzblobconfig returns the Azblobconfig field if non-nil, zero value otherwise.

### GetAzblobconfigOk

`func (o *FilesystemConfig) GetAzblobconfigOk() (*AzureBlobFsConfig, bool)`

GetAzblobconfigOk returns a tuple with the Azblobconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzblobconfig

`func (o *FilesystemConfig) SetAzblobconfig(v AzureBlobFsConfig)`

SetAzblobconfig sets Azblobconfig field to given value.

### HasAzblobconfig

`func (o *FilesystemConfig) HasAzblobconfig() bool`

HasAzblobconfig returns a boolean if a field has been set.

### GetCryptconfig

`func (o *FilesystemConfig) GetCryptconfig() CryptFsConfig`

GetCryptconfig returns the Cryptconfig field if non-nil, zero value otherwise.

### GetCryptconfigOk

`func (o *FilesystemConfig) GetCryptconfigOk() (*CryptFsConfig, bool)`

GetCryptconfigOk returns a tuple with the Cryptconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptconfig

`func (o *FilesystemConfig) SetCryptconfig(v CryptFsConfig)`

SetCryptconfig sets Cryptconfig field to given value.

### HasCryptconfig

`func (o *FilesystemConfig) HasCryptconfig() bool`

HasCryptconfig returns a boolean if a field has been set.

### GetSftpconfig

`func (o *FilesystemConfig) GetSftpconfig() SFTPFsConfig`

GetSftpconfig returns the Sftpconfig field if non-nil, zero value otherwise.

### GetSftpconfigOk

`func (o *FilesystemConfig) GetSftpconfigOk() (*SFTPFsConfig, bool)`

GetSftpconfigOk returns a tuple with the Sftpconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSftpconfig

`func (o *FilesystemConfig) SetSftpconfig(v SFTPFsConfig)`

SetSftpconfig sets Sftpconfig field to given value.

### HasSftpconfig

`func (o *FilesystemConfig) HasSftpconfig() bool`

HasSftpconfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


