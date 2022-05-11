# GCSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**AutomaticCredentials** | Pointer to **int32** | Automatic credentials:   * &#x60;0&#x60; - disabled, explicit credentials, using a JSON credentials file, must be provided. This is the default value if the field is null   * &#x60;1&#x60; - enabled, we try to use the Application Default Credentials (ADC) strategy to find your application&#39;s credentials  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**Acl** | Pointer to **string** | The ACL to apply to uploaded objects. Leave empty to use the default ACL. For more information and available ACLs, refer to the JSON API here: https://cloud.google.com/storage/docs/access-control/lists#predefined-acl | [optional] 
**KeyPrefix** | Pointer to **string** | key_prefix is similar to a chroot directory for a local filesystem. If specified the user will only see contents that starts with this prefix and so you can restrict access to a specific virtual folder. The prefix, if not empty, must not start with \&quot;/\&quot; and must end with \&quot;/\&quot;. If empty the whole bucket contents will be available | [optional] 

## Methods

### NewGCSConfig

`func NewGCSConfig() *GCSConfig`

NewGCSConfig instantiates a new GCSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSConfigWithDefaults

`func NewGCSConfigWithDefaults() *GCSConfig`

NewGCSConfigWithDefaults instantiates a new GCSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *GCSConfig) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *GCSConfig) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *GCSConfig) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *GCSConfig) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCredentials

`func (o *GCSConfig) GetCredentials() Secret`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GCSConfig) GetCredentialsOk() (*Secret, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GCSConfig) SetCredentials(v Secret)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GCSConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetAutomaticCredentials

`func (o *GCSConfig) GetAutomaticCredentials() int32`

GetAutomaticCredentials returns the AutomaticCredentials field if non-nil, zero value otherwise.

### GetAutomaticCredentialsOk

`func (o *GCSConfig) GetAutomaticCredentialsOk() (*int32, bool)`

GetAutomaticCredentialsOk returns a tuple with the AutomaticCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticCredentials

`func (o *GCSConfig) SetAutomaticCredentials(v int32)`

SetAutomaticCredentials sets AutomaticCredentials field to given value.

### HasAutomaticCredentials

`func (o *GCSConfig) HasAutomaticCredentials() bool`

HasAutomaticCredentials returns a boolean if a field has been set.

### GetStorageClass

`func (o *GCSConfig) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *GCSConfig) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *GCSConfig) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *GCSConfig) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetAcl

`func (o *GCSConfig) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *GCSConfig) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *GCSConfig) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *GCSConfig) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *GCSConfig) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *GCSConfig) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *GCSConfig) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *GCSConfig) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


