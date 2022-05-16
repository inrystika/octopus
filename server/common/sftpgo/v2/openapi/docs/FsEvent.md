# FsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** | unix timestamp in nanoseconds | [optional] 
**Action** | Pointer to [**FsEventAction**](FsEventAction.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**FsPath** | Pointer to **string** |  | [optional] 
**FsTargetPath** | Pointer to **string** |  | [optional] 
**VirtualPath** | Pointer to **string** |  | [optional] 
**VirtualTargetPath** | Pointer to **string** |  | [optional] 
**SshCmd** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**FsEventStatus**](FsEventStatus.md) |  | [optional] 
**Protocol** | Pointer to [**EventProtocols**](EventProtocols.md) |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**FsProvider** | Pointer to [**FsProviders**](FsProviders.md) |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**OpenFlags** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 

## Methods

### NewFsEvent

`func NewFsEvent() *FsEvent`

NewFsEvent instantiates a new FsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFsEventWithDefaults

`func NewFsEventWithDefaults() *FsEvent`

NewFsEventWithDefaults instantiates a new FsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *FsEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FsEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FsEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FsEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAction

`func (o *FsEvent) GetAction() FsEventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FsEvent) GetActionOk() (*FsEventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FsEvent) SetAction(v FsEventAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *FsEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUsername

`func (o *FsEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FsEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FsEvent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FsEvent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFsPath

`func (o *FsEvent) GetFsPath() string`

GetFsPath returns the FsPath field if non-nil, zero value otherwise.

### GetFsPathOk

`func (o *FsEvent) GetFsPathOk() (*string, bool)`

GetFsPathOk returns a tuple with the FsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsPath

`func (o *FsEvent) SetFsPath(v string)`

SetFsPath sets FsPath field to given value.

### HasFsPath

`func (o *FsEvent) HasFsPath() bool`

HasFsPath returns a boolean if a field has been set.

### GetFsTargetPath

`func (o *FsEvent) GetFsTargetPath() string`

GetFsTargetPath returns the FsTargetPath field if non-nil, zero value otherwise.

### GetFsTargetPathOk

`func (o *FsEvent) GetFsTargetPathOk() (*string, bool)`

GetFsTargetPathOk returns a tuple with the FsTargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsTargetPath

`func (o *FsEvent) SetFsTargetPath(v string)`

SetFsTargetPath sets FsTargetPath field to given value.

### HasFsTargetPath

`func (o *FsEvent) HasFsTargetPath() bool`

HasFsTargetPath returns a boolean if a field has been set.

### GetVirtualPath

`func (o *FsEvent) GetVirtualPath() string`

GetVirtualPath returns the VirtualPath field if non-nil, zero value otherwise.

### GetVirtualPathOk

`func (o *FsEvent) GetVirtualPathOk() (*string, bool)`

GetVirtualPathOk returns a tuple with the VirtualPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPath

`func (o *FsEvent) SetVirtualPath(v string)`

SetVirtualPath sets VirtualPath field to given value.

### HasVirtualPath

`func (o *FsEvent) HasVirtualPath() bool`

HasVirtualPath returns a boolean if a field has been set.

### GetVirtualTargetPath

`func (o *FsEvent) GetVirtualTargetPath() string`

GetVirtualTargetPath returns the VirtualTargetPath field if non-nil, zero value otherwise.

### GetVirtualTargetPathOk

`func (o *FsEvent) GetVirtualTargetPathOk() (*string, bool)`

GetVirtualTargetPathOk returns a tuple with the VirtualTargetPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualTargetPath

`func (o *FsEvent) SetVirtualTargetPath(v string)`

SetVirtualTargetPath sets VirtualTargetPath field to given value.

### HasVirtualTargetPath

`func (o *FsEvent) HasVirtualTargetPath() bool`

HasVirtualTargetPath returns a boolean if a field has been set.

### GetSshCmd

`func (o *FsEvent) GetSshCmd() string`

GetSshCmd returns the SshCmd field if non-nil, zero value otherwise.

### GetSshCmdOk

`func (o *FsEvent) GetSshCmdOk() (*string, bool)`

GetSshCmdOk returns a tuple with the SshCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCmd

`func (o *FsEvent) SetSshCmd(v string)`

SetSshCmd sets SshCmd field to given value.

### HasSshCmd

`func (o *FsEvent) HasSshCmd() bool`

HasSshCmd returns a boolean if a field has been set.

### GetFileSize

`func (o *FsEvent) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FsEvent) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FsEvent) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FsEvent) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetStatus

`func (o *FsEvent) GetStatus() FsEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FsEvent) GetStatusOk() (*FsEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FsEvent) SetStatus(v FsEventStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FsEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProtocol

`func (o *FsEvent) GetProtocol() EventProtocols`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FsEvent) GetProtocolOk() (*EventProtocols, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FsEvent) SetProtocol(v EventProtocols)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FsEvent) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetIp

`func (o *FsEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FsEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FsEvent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FsEvent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSessionId

`func (o *FsEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *FsEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *FsEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *FsEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetFsProvider

`func (o *FsEvent) GetFsProvider() FsProviders`

GetFsProvider returns the FsProvider field if non-nil, zero value otherwise.

### GetFsProviderOk

`func (o *FsEvent) GetFsProviderOk() (*FsProviders, bool)`

GetFsProviderOk returns a tuple with the FsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsProvider

`func (o *FsEvent) SetFsProvider(v FsProviders)`

SetFsProvider sets FsProvider field to given value.

### HasFsProvider

`func (o *FsEvent) HasFsProvider() bool`

HasFsProvider returns a boolean if a field has been set.

### GetBucket

`func (o *FsEvent) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *FsEvent) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *FsEvent) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *FsEvent) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *FsEvent) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FsEvent) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FsEvent) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FsEvent) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetOpenFlags

`func (o *FsEvent) GetOpenFlags() string`

GetOpenFlags returns the OpenFlags field if non-nil, zero value otherwise.

### GetOpenFlagsOk

`func (o *FsEvent) GetOpenFlagsOk() (*string, bool)`

GetOpenFlagsOk returns a tuple with the OpenFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenFlags

`func (o *FsEvent) SetOpenFlags(v string)`

SetOpenFlags sets OpenFlags field to given value.

### HasOpenFlags

`func (o *FsEvent) HasOpenFlags() bool`

HasOpenFlags returns a boolean if a field has been set.

### GetInstanceId

`func (o *FsEvent) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FsEvent) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FsEvent) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *FsEvent) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


