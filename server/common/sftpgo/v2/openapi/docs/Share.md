# Share

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | auto-generated unique share identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | optional description | [optional] 
**Scope** | Pointer to [**ShareScope**](ShareScope.md) |  | [optional] 
**Paths** | Pointer to **[]string** | paths to files or directories, for share scope write this array must contain exactly one directory. Paths will not be validated on save so you can also create them after creating the share | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | creation time as unix timestamp in milliseconds | [optional] 
**UpdatedAt** | Pointer to **int64** | last update time as unix timestamp in milliseconds | [optional] 
**LastUseAt** | Pointer to **int64** | last use time as unix timestamp in milliseconds | [optional] 
**ExpiresAt** | Pointer to **int64** | optional share expiration, as unix timestamp in milliseconds. 0 means no expiration | [optional] 
**Password** | Pointer to **string** | optional password to protect the share. The special value \&quot;[**redacted**]\&quot; means that a password has been set, you can use this value if you want to preserve the current password when you update a share | [optional] 
**MaxTokens** | Pointer to **int32** | maximum allowed access tokens. 0 means no limit | [optional] 
**UsedTokens** | Pointer to **int32** |  | [optional] 
**AllowFrom** | Pointer to **[]string** | Limit the share availability to these IP/Mask. IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291, for example \&quot;192.0.2.0/24\&quot; or \&quot;2001:db8::/32\&quot;. An empty list means no restrictions | [optional] 

## Methods

### NewShare

`func NewShare() *Share`

NewShare instantiates a new Share object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareWithDefaults

`func NewShareWithDefaults() *Share`

NewShareWithDefaults instantiates a new Share object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Share) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Share) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Share) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Share) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Share) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Share) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Share) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Share) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Share) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Share) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Share) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Share) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *Share) GetScope() ShareScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Share) GetScopeOk() (*ShareScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Share) SetScope(v ShareScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Share) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetPaths

`func (o *Share) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *Share) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *Share) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *Share) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetUsername

`func (o *Share) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Share) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Share) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Share) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Share) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Share) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Share) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Share) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Share) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Share) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Share) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Share) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastUseAt

`func (o *Share) GetLastUseAt() int64`

GetLastUseAt returns the LastUseAt field if non-nil, zero value otherwise.

### GetLastUseAtOk

`func (o *Share) GetLastUseAtOk() (*int64, bool)`

GetLastUseAtOk returns a tuple with the LastUseAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUseAt

`func (o *Share) SetLastUseAt(v int64)`

SetLastUseAt sets LastUseAt field to given value.

### HasLastUseAt

`func (o *Share) HasLastUseAt() bool`

HasLastUseAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Share) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Share) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Share) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Share) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetPassword

`func (o *Share) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Share) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Share) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Share) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMaxTokens

`func (o *Share) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *Share) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *Share) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *Share) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetUsedTokens

`func (o *Share) GetUsedTokens() int32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *Share) GetUsedTokensOk() (*int32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *Share) SetUsedTokens(v int32)`

SetUsedTokens sets UsedTokens field to given value.

### HasUsedTokens

`func (o *Share) HasUsedTokens() bool`

HasUsedTokens returns a boolean if a field has been set.

### GetAllowFrom

`func (o *Share) GetAllowFrom() []string`

GetAllowFrom returns the AllowFrom field if non-nil, zero value otherwise.

### GetAllowFromOk

`func (o *Share) GetAllowFromOk() (*[]string, bool)`

GetAllowFromOk returns a tuple with the AllowFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFrom

`func (o *Share) SetAllowFrom(v []string)`

SetAllowFrom sets AllowFrom field to given value.

### HasAllowFrom

`func (o *Share) HasAllowFrom() bool`

HasAllowFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


