# AzureBlobFsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** | Storage Account Name, leave blank to use SAS URL | [optional] 
**AccountKey** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**SasUrl** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**Endpoint** | Pointer to **string** | optional endpoint. Default is \&quot;blob.core.windows.net\&quot;. If you use the emulator the endpoint must include the protocol, for example \&quot;http://127.0.0.1:10000\&quot; | [optional] 
**UploadPartSize** | Pointer to **int32** | the buffer size (in MB) to use for multipart uploads. If this value is set to zero, the default value (4MB) will be used. | [optional] 
**UploadConcurrency** | Pointer to **int32** | the number of parts to upload in parallel. If this value is set to zero, the default value (2) will be used | [optional] 
**AccessTier** | Pointer to **string** |  | [optional] 
**KeyPrefix** | Pointer to **string** | key_prefix is similar to a chroot directory for a local filesystem. If specified the user will only see contents that starts with this prefix and so you can restrict access to a specific virtual folder. The prefix, if not empty, must not start with \&quot;/\&quot; and must end with \&quot;/\&quot;. If empty the whole container contents will be available | [optional] 
**UseEmulator** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureBlobFsConfig

`func NewAzureBlobFsConfig() *AzureBlobFsConfig`

NewAzureBlobFsConfig instantiates a new AzureBlobFsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobFsConfigWithDefaults

`func NewAzureBlobFsConfigWithDefaults() *AzureBlobFsConfig`

NewAzureBlobFsConfigWithDefaults instantiates a new AzureBlobFsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *AzureBlobFsConfig) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureBlobFsConfig) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureBlobFsConfig) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *AzureBlobFsConfig) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAccountName

`func (o *AzureBlobFsConfig) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AzureBlobFsConfig) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AzureBlobFsConfig) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AzureBlobFsConfig) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountKey

`func (o *AzureBlobFsConfig) GetAccountKey() Secret`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AzureBlobFsConfig) GetAccountKeyOk() (*Secret, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AzureBlobFsConfig) SetAccountKey(v Secret)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *AzureBlobFsConfig) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetSasUrl

`func (o *AzureBlobFsConfig) GetSasUrl() Secret`

GetSasUrl returns the SasUrl field if non-nil, zero value otherwise.

### GetSasUrlOk

`func (o *AzureBlobFsConfig) GetSasUrlOk() (*Secret, bool)`

GetSasUrlOk returns a tuple with the SasUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasUrl

`func (o *AzureBlobFsConfig) SetSasUrl(v Secret)`

SetSasUrl sets SasUrl field to given value.

### HasSasUrl

`func (o *AzureBlobFsConfig) HasSasUrl() bool`

HasSasUrl returns a boolean if a field has been set.

### GetEndpoint

`func (o *AzureBlobFsConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AzureBlobFsConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AzureBlobFsConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AzureBlobFsConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetUploadPartSize

`func (o *AzureBlobFsConfig) GetUploadPartSize() int32`

GetUploadPartSize returns the UploadPartSize field if non-nil, zero value otherwise.

### GetUploadPartSizeOk

`func (o *AzureBlobFsConfig) GetUploadPartSizeOk() (*int32, bool)`

GetUploadPartSizeOk returns a tuple with the UploadPartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPartSize

`func (o *AzureBlobFsConfig) SetUploadPartSize(v int32)`

SetUploadPartSize sets UploadPartSize field to given value.

### HasUploadPartSize

`func (o *AzureBlobFsConfig) HasUploadPartSize() bool`

HasUploadPartSize returns a boolean if a field has been set.

### GetUploadConcurrency

`func (o *AzureBlobFsConfig) GetUploadConcurrency() int32`

GetUploadConcurrency returns the UploadConcurrency field if non-nil, zero value otherwise.

### GetUploadConcurrencyOk

`func (o *AzureBlobFsConfig) GetUploadConcurrencyOk() (*int32, bool)`

GetUploadConcurrencyOk returns a tuple with the UploadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadConcurrency

`func (o *AzureBlobFsConfig) SetUploadConcurrency(v int32)`

SetUploadConcurrency sets UploadConcurrency field to given value.

### HasUploadConcurrency

`func (o *AzureBlobFsConfig) HasUploadConcurrency() bool`

HasUploadConcurrency returns a boolean if a field has been set.

### GetAccessTier

`func (o *AzureBlobFsConfig) GetAccessTier() string`

GetAccessTier returns the AccessTier field if non-nil, zero value otherwise.

### GetAccessTierOk

`func (o *AzureBlobFsConfig) GetAccessTierOk() (*string, bool)`

GetAccessTierOk returns a tuple with the AccessTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTier

`func (o *AzureBlobFsConfig) SetAccessTier(v string)`

SetAccessTier sets AccessTier field to given value.

### HasAccessTier

`func (o *AzureBlobFsConfig) HasAccessTier() bool`

HasAccessTier returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *AzureBlobFsConfig) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *AzureBlobFsConfig) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *AzureBlobFsConfig) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *AzureBlobFsConfig) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetUseEmulator

`func (o *AzureBlobFsConfig) GetUseEmulator() bool`

GetUseEmulator returns the UseEmulator field if non-nil, zero value otherwise.

### GetUseEmulatorOk

`func (o *AzureBlobFsConfig) GetUseEmulatorOk() (*bool, bool)`

GetUseEmulatorOk returns a tuple with the UseEmulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmulator

`func (o *AzureBlobFsConfig) SetUseEmulator(v bool)`

SetUseEmulator sets UseEmulator field to given value.

### HasUseEmulator

`func (o *AzureBlobFsConfig) HasUseEmulator() bool`

HasUseEmulator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


