# SFTPFsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | remote SFTP endpoint as host:port | [optional] 
**Username** | Pointer to **string** | you can specify a password or private key or both. In the latter case the private key will be tried first. | [optional] 
**Password** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**PrivateKey** | Pointer to [**Secret**](Secret.md) |  | [optional] 
**Fingerprints** | Pointer to **[]string** | SHA256 fingerprints to use for host key verification. If you don&#39;t provide any fingerprint the remote host key will not be verified, this is a security risk | [optional] 
**Prefix** | Pointer to **string** | Specifying a prefix you can restrict all operations to a given path within the remote SFTP server. | [optional] 
**DisableConcurrentReads** | Pointer to **bool** | Concurrent reads are safe to use and disabling them will degrade performance. Some servers automatically delete files once they are downloaded. Using concurrent reads is problematic with such servers. | [optional] 
**BufferSize** | Pointer to **int32** | The size of the buffer (in MB) to use for transfers. By enabling buffering, the reads and writes, from/to the remote SFTP server, are split in multiple concurrent requests and this allows data to be transferred at a faster rate, over high latency networks, by overlapping round-trip times. With buffering enabled, resuming uploads is not supported and a file cannot be opened for both reading and writing at the same time. 0 means disabled. | [optional] 

## Methods

### NewSFTPFsConfig

`func NewSFTPFsConfig() *SFTPFsConfig`

NewSFTPFsConfig instantiates a new SFTPFsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFTPFsConfigWithDefaults

`func NewSFTPFsConfigWithDefaults() *SFTPFsConfig`

NewSFTPFsConfigWithDefaults instantiates a new SFTPFsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *SFTPFsConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SFTPFsConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SFTPFsConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *SFTPFsConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetUsername

`func (o *SFTPFsConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SFTPFsConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SFTPFsConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SFTPFsConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *SFTPFsConfig) GetPassword() Secret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SFTPFsConfig) GetPasswordOk() (*Secret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SFTPFsConfig) SetPassword(v Secret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SFTPFsConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKey

`func (o *SFTPFsConfig) GetPrivateKey() Secret`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SFTPFsConfig) GetPrivateKeyOk() (*Secret, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SFTPFsConfig) SetPrivateKey(v Secret)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SFTPFsConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetFingerprints

`func (o *SFTPFsConfig) GetFingerprints() []string`

GetFingerprints returns the Fingerprints field if non-nil, zero value otherwise.

### GetFingerprintsOk

`func (o *SFTPFsConfig) GetFingerprintsOk() (*[]string, bool)`

GetFingerprintsOk returns a tuple with the Fingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprints

`func (o *SFTPFsConfig) SetFingerprints(v []string)`

SetFingerprints sets Fingerprints field to given value.

### HasFingerprints

`func (o *SFTPFsConfig) HasFingerprints() bool`

HasFingerprints returns a boolean if a field has been set.

### GetPrefix

`func (o *SFTPFsConfig) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SFTPFsConfig) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SFTPFsConfig) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SFTPFsConfig) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetDisableConcurrentReads

`func (o *SFTPFsConfig) GetDisableConcurrentReads() bool`

GetDisableConcurrentReads returns the DisableConcurrentReads field if non-nil, zero value otherwise.

### GetDisableConcurrentReadsOk

`func (o *SFTPFsConfig) GetDisableConcurrentReadsOk() (*bool, bool)`

GetDisableConcurrentReadsOk returns a tuple with the DisableConcurrentReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConcurrentReads

`func (o *SFTPFsConfig) SetDisableConcurrentReads(v bool)`

SetDisableConcurrentReads sets DisableConcurrentReads field to given value.

### HasDisableConcurrentReads

`func (o *SFTPFsConfig) HasDisableConcurrentReads() bool`

HasDisableConcurrentReads returns a boolean if a field has been set.

### GetBufferSize

`func (o *SFTPFsConfig) GetBufferSize() int32`

GetBufferSize returns the BufferSize field if non-nil, zero value otherwise.

### GetBufferSizeOk

`func (o *SFTPFsConfig) GetBufferSizeOk() (*int32, bool)`

GetBufferSizeOk returns a tuple with the BufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferSize

`func (o *SFTPFsConfig) SetBufferSize(v int32)`

SetBufferSize sets BufferSize field to given value.

### HasBufferSize

`func (o *SFTPFsConfig) HasBufferSize() bool`

HasBufferSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


