# ConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | connected username | [optional] 
**ConnectionId** | Pointer to **string** | unique connection identifier | [optional] 
**ClientVersion** | Pointer to **string** | client version | [optional] 
**RemoteAddress** | Pointer to **string** | Remote address for the connected client | [optional] 
**ConnectionTime** | Pointer to **int64** | connection time as unix timestamp in milliseconds | [optional] 
**Command** | Pointer to **string** | Last SSH/FTP command or WebDAV method | [optional] 
**LastActivity** | Pointer to **int64** | last client activity as unix timestamp in milliseconds | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**ActiveTransfers** | Pointer to [**[]Transfer**](Transfer.md) |  | [optional] 

## Methods

### NewConnectionStatus

`func NewConnectionStatus() *ConnectionStatus`

NewConnectionStatus instantiates a new ConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusWithDefaults

`func NewConnectionStatusWithDefaults() *ConnectionStatus`

NewConnectionStatusWithDefaults instantiates a new ConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ConnectionStatus) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectionStatus) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectionStatus) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectionStatus) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetConnectionId

`func (o *ConnectionStatus) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionStatus) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionStatus) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConnectionStatus) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetClientVersion

`func (o *ConnectionStatus) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *ConnectionStatus) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *ConnectionStatus) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *ConnectionStatus) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *ConnectionStatus) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *ConnectionStatus) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *ConnectionStatus) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *ConnectionStatus) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetConnectionTime

`func (o *ConnectionStatus) GetConnectionTime() int64`

GetConnectionTime returns the ConnectionTime field if non-nil, zero value otherwise.

### GetConnectionTimeOk

`func (o *ConnectionStatus) GetConnectionTimeOk() (*int64, bool)`

GetConnectionTimeOk returns a tuple with the ConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTime

`func (o *ConnectionStatus) SetConnectionTime(v int64)`

SetConnectionTime sets ConnectionTime field to given value.

### HasConnectionTime

`func (o *ConnectionStatus) HasConnectionTime() bool`

HasConnectionTime returns a boolean if a field has been set.

### GetCommand

`func (o *ConnectionStatus) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ConnectionStatus) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ConnectionStatus) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ConnectionStatus) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetLastActivity

`func (o *ConnectionStatus) GetLastActivity() int64`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *ConnectionStatus) GetLastActivityOk() (*int64, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *ConnectionStatus) SetLastActivity(v int64)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *ConnectionStatus) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetProtocol

`func (o *ConnectionStatus) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ConnectionStatus) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ConnectionStatus) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ConnectionStatus) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetActiveTransfers

`func (o *ConnectionStatus) GetActiveTransfers() []Transfer`

GetActiveTransfers returns the ActiveTransfers field if non-nil, zero value otherwise.

### GetActiveTransfersOk

`func (o *ConnectionStatus) GetActiveTransfersOk() (*[]Transfer, bool)`

GetActiveTransfersOk returns a tuple with the ActiveTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTransfers

`func (o *ConnectionStatus) SetActiveTransfers(v []Transfer)`

SetActiveTransfers sets ActiveTransfers field to given value.

### HasActiveTransfers

`func (o *ConnectionStatus) HasActiveTransfers() bool`

HasActiveTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


