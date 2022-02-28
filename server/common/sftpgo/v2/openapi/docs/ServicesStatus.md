# ServicesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssh** | Pointer to [**SSHServiceStatus**](SSHServiceStatus.md) |  | [optional] 
**Ftp** | Pointer to [**FTPServiceStatus**](FTPServiceStatus.md) |  | [optional] 
**Webdav** | Pointer to [**WebDAVServiceStatus**](WebDAVServiceStatus.md) |  | [optional] 
**DataProvider** | Pointer to [**DataProviderStatus**](DataProviderStatus.md) |  | [optional] 
**Defender** | Pointer to [**ServicesStatusDefender**](ServicesStatusDefender.md) |  | [optional] 
**Mfa** | Pointer to [**MFAStatus**](MFAStatus.md) |  | [optional] 

## Methods

### NewServicesStatus

`func NewServicesStatus() *ServicesStatus`

NewServicesStatus instantiates a new ServicesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesStatusWithDefaults

`func NewServicesStatusWithDefaults() *ServicesStatus`

NewServicesStatusWithDefaults instantiates a new ServicesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsh

`func (o *ServicesStatus) GetSsh() SSHServiceStatus`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *ServicesStatus) GetSshOk() (*SSHServiceStatus, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *ServicesStatus) SetSsh(v SSHServiceStatus)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *ServicesStatus) HasSsh() bool`

HasSsh returns a boolean if a field has been set.

### GetFtp

`func (o *ServicesStatus) GetFtp() FTPServiceStatus`

GetFtp returns the Ftp field if non-nil, zero value otherwise.

### GetFtpOk

`func (o *ServicesStatus) GetFtpOk() (*FTPServiceStatus, bool)`

GetFtpOk returns a tuple with the Ftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtp

`func (o *ServicesStatus) SetFtp(v FTPServiceStatus)`

SetFtp sets Ftp field to given value.

### HasFtp

`func (o *ServicesStatus) HasFtp() bool`

HasFtp returns a boolean if a field has been set.

### GetWebdav

`func (o *ServicesStatus) GetWebdav() WebDAVServiceStatus`

GetWebdav returns the Webdav field if non-nil, zero value otherwise.

### GetWebdavOk

`func (o *ServicesStatus) GetWebdavOk() (*WebDAVServiceStatus, bool)`

GetWebdavOk returns a tuple with the Webdav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdav

`func (o *ServicesStatus) SetWebdav(v WebDAVServiceStatus)`

SetWebdav sets Webdav field to given value.

### HasWebdav

`func (o *ServicesStatus) HasWebdav() bool`

HasWebdav returns a boolean if a field has been set.

### GetDataProvider

`func (o *ServicesStatus) GetDataProvider() DataProviderStatus`

GetDataProvider returns the DataProvider field if non-nil, zero value otherwise.

### GetDataProviderOk

`func (o *ServicesStatus) GetDataProviderOk() (*DataProviderStatus, bool)`

GetDataProviderOk returns a tuple with the DataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProvider

`func (o *ServicesStatus) SetDataProvider(v DataProviderStatus)`

SetDataProvider sets DataProvider field to given value.

### HasDataProvider

`func (o *ServicesStatus) HasDataProvider() bool`

HasDataProvider returns a boolean if a field has been set.

### GetDefender

`func (o *ServicesStatus) GetDefender() ServicesStatusDefender`

GetDefender returns the Defender field if non-nil, zero value otherwise.

### GetDefenderOk

`func (o *ServicesStatus) GetDefenderOk() (*ServicesStatusDefender, bool)`

GetDefenderOk returns a tuple with the Defender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefender

`func (o *ServicesStatus) SetDefender(v ServicesStatusDefender)`

SetDefender sets Defender field to given value.

### HasDefender

`func (o *ServicesStatus) HasDefender() bool`

HasDefender returns a boolean if a field has been set.

### GetMfa

`func (o *ServicesStatus) GetMfa() MFAStatus`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *ServicesStatus) GetMfaOk() (*MFAStatus, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *ServicesStatus) SetMfa(v MFAStatus)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *ServicesStatus) HasMfa() bool`

HasMfa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


