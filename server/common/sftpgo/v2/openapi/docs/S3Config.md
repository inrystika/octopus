# S3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessSecret** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**Endpoint** | Pointer to **string** | optional endpoint | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**Acl** | Pointer to **string** | The canned ACL to apply to uploaded objects. Leave empty to use the default ACL. For more information and available ACLs, see here: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl | [optional] 
**UploadPartSize** | Pointer to **int32** | the buffer size (in MB) to use for multipart uploads. The minimum allowed part size is 5MB, and if this value is set to zero, the default value (5MB) for the AWS SDK will be used. The minimum allowed value is 5. | [optional] 
**UploadConcurrency** | Pointer to **int32** | the number of parts to upload in parallel. If this value is set to zero, the default value (5) will be used | [optional] 
**UploadPartMaxTime** | Pointer to **int32** | the maximum time allowed, in seconds, to upload a single chunk (the chunk size is defined via \&quot;upload_part_size\&quot;). 0 means no timeout | [optional] 
**DownloadPartSize** | Pointer to **int32** | the buffer size (in MB) to use for multipart downloads. The minimum allowed part size is 5MB, and if this value is set to zero, the default value (5MB) for the AWS SDK will be used. The minimum allowed value is 5. Ignored for partial downloads | [optional] 
**DownloadConcurrency** | Pointer to **int32** | the number of parts to download in parallel. If this value is set to zero, the default value (5) will be used. Ignored for partial downloads | [optional] 
**DownloadPartMaxTime** | Pointer to **int32** | the maximum time allowed, in seconds, to download a single chunk (the chunk size is defined via \&quot;download_part_size\&quot;). 0 means no timeout. Ignored for partial downloads. | [optional] 
**ForcePathStyle** | Pointer to **bool** | Set this to \&quot;true\&quot; to force the request to use path-style addressing, i.e., \&quot;http://s3.amazonaws.com/BUCKET/KEY\&quot;. By default, the S3 client will use virtual hosted bucket addressing when possible (\&quot;http://BUCKET.s3.amazonaws.com/KEY\&quot;) | [optional] 
**KeyPrefix** | Pointer to **string** | key_prefix is similar to a chroot directory for a local filesystem. If specified the user will only see contents that starts with this prefix and so you can restrict access to a specific virtual folder. The prefix, if not empty, must not start with \&quot;/\&quot; and must end with \&quot;/\&quot;. If empty the whole bucket contents will be available | [optional] 

## Methods

### NewS3Config

`func NewS3Config() *S3Config`

NewS3Config instantiates a new S3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigWithDefaults

`func NewS3ConfigWithDefaults() *S3Config`

NewS3ConfigWithDefaults instantiates a new S3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *S3Config) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3Config) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3Config) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *S3Config) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetRegion

`func (o *S3Config) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3Config) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3Config) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *S3Config) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAccessKey

`func (o *S3Config) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3Config) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3Config) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *S3Config) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessSecret

`func (o *S3Config) GetAccessSecret() Secret`

GetAccessSecret returns the AccessSecret field if non-nil, zero value otherwise.

### GetAccessSecretOk

`func (o *S3Config) GetAccessSecretOk() (*Secret, bool)`

GetAccessSecretOk returns a tuple with the AccessSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecret

`func (o *S3Config) SetAccessSecret(v Secret)`

SetAccessSecret sets AccessSecret field to given value.

### HasAccessSecret

`func (o *S3Config) HasAccessSecret() bool`

HasAccessSecret returns a boolean if a field has been set.

### GetEndpoint

`func (o *S3Config) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *S3Config) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *S3Config) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *S3Config) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetStorageClass

`func (o *S3Config) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *S3Config) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *S3Config) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *S3Config) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetAcl

`func (o *S3Config) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *S3Config) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *S3Config) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *S3Config) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetUploadPartSize

`func (o *S3Config) GetUploadPartSize() int32`

GetUploadPartSize returns the UploadPartSize field if non-nil, zero value otherwise.

### GetUploadPartSizeOk

`func (o *S3Config) GetUploadPartSizeOk() (*int32, bool)`

GetUploadPartSizeOk returns a tuple with the UploadPartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPartSize

`func (o *S3Config) SetUploadPartSize(v int32)`

SetUploadPartSize sets UploadPartSize field to given value.

### HasUploadPartSize

`func (o *S3Config) HasUploadPartSize() bool`

HasUploadPartSize returns a boolean if a field has been set.

### GetUploadConcurrency

`func (o *S3Config) GetUploadConcurrency() int32`

GetUploadConcurrency returns the UploadConcurrency field if non-nil, zero value otherwise.

### GetUploadConcurrencyOk

`func (o *S3Config) GetUploadConcurrencyOk() (*int32, bool)`

GetUploadConcurrencyOk returns a tuple with the UploadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadConcurrency

`func (o *S3Config) SetUploadConcurrency(v int32)`

SetUploadConcurrency sets UploadConcurrency field to given value.

### HasUploadConcurrency

`func (o *S3Config) HasUploadConcurrency() bool`

HasUploadConcurrency returns a boolean if a field has been set.

### GetUploadPartMaxTime

`func (o *S3Config) GetUploadPartMaxTime() int32`

GetUploadPartMaxTime returns the UploadPartMaxTime field if non-nil, zero value otherwise.

### GetUploadPartMaxTimeOk

`func (o *S3Config) GetUploadPartMaxTimeOk() (*int32, bool)`

GetUploadPartMaxTimeOk returns a tuple with the UploadPartMaxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPartMaxTime

`func (o *S3Config) SetUploadPartMaxTime(v int32)`

SetUploadPartMaxTime sets UploadPartMaxTime field to given value.

### HasUploadPartMaxTime

`func (o *S3Config) HasUploadPartMaxTime() bool`

HasUploadPartMaxTime returns a boolean if a field has been set.

### GetDownloadPartSize

`func (o *S3Config) GetDownloadPartSize() int32`

GetDownloadPartSize returns the DownloadPartSize field if non-nil, zero value otherwise.

### GetDownloadPartSizeOk

`func (o *S3Config) GetDownloadPartSizeOk() (*int32, bool)`

GetDownloadPartSizeOk returns a tuple with the DownloadPartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPartSize

`func (o *S3Config) SetDownloadPartSize(v int32)`

SetDownloadPartSize sets DownloadPartSize field to given value.

### HasDownloadPartSize

`func (o *S3Config) HasDownloadPartSize() bool`

HasDownloadPartSize returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *S3Config) GetDownloadConcurrency() int32`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *S3Config) GetDownloadConcurrencyOk() (*int32, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *S3Config) SetDownloadConcurrency(v int32)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *S3Config) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### GetDownloadPartMaxTime

`func (o *S3Config) GetDownloadPartMaxTime() int32`

GetDownloadPartMaxTime returns the DownloadPartMaxTime field if non-nil, zero value otherwise.

### GetDownloadPartMaxTimeOk

`func (o *S3Config) GetDownloadPartMaxTimeOk() (*int32, bool)`

GetDownloadPartMaxTimeOk returns a tuple with the DownloadPartMaxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPartMaxTime

`func (o *S3Config) SetDownloadPartMaxTime(v int32)`

SetDownloadPartMaxTime sets DownloadPartMaxTime field to given value.

### HasDownloadPartMaxTime

`func (o *S3Config) HasDownloadPartMaxTime() bool`

HasDownloadPartMaxTime returns a boolean if a field has been set.

### GetForcePathStyle

`func (o *S3Config) GetForcePathStyle() bool`

GetForcePathStyle returns the ForcePathStyle field if non-nil, zero value otherwise.

### GetForcePathStyleOk

`func (o *S3Config) GetForcePathStyleOk() (*bool, bool)`

GetForcePathStyleOk returns a tuple with the ForcePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePathStyle

`func (o *S3Config) SetForcePathStyle(v bool)`

SetForcePathStyle sets ForcePathStyle field to given value.

### HasForcePathStyle

`func (o *S3Config) HasForcePathStyle() bool`

HasForcePathStyle returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *S3Config) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *S3Config) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *S3Config) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *S3Config) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


