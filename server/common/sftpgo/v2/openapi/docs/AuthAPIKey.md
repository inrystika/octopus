# AuthAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | unique key identifier | [optional] 
**Name** | Pointer to **string** | User friendly key name | [optional] 
**Key** | Pointer to **string** | We store the hash of the key. This is just like a password. For security reasons this field is omitted when you search/get API keys | [optional] 
**Scope** | Pointer to [**AuthAPIKeyScope**](AuthAPIKeyScope.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** | creation time as unix timestamp in milliseconds | [optional] 
**UpdatedAt** | Pointer to **int64** | last update time as unix timestamp in milliseconds | [optional] 
**LastUseAt** | Pointer to **int64** | last use time as unix timestamp in milliseconds. It is saved at most once every 10 minutes | [optional] 
**ExpiresAt** | Pointer to **int64** | expiration time as unix timestamp in milliseconds | [optional] 
**Description** | Pointer to **string** | optional description | [optional] 
**User** | Pointer to **string** | username associated with this API key. If empty and the scope is \&quot;user scope\&quot; the key can impersonate any user | [optional] 
**Admin** | Pointer to **string** | admin associated with this API key. If empty and the scope is \&quot;admin scope\&quot; the key can impersonate any admin | [optional] 

## Methods

### NewAuthAPIKey

`func NewAuthAPIKey() *AuthAPIKey`

NewAuthAPIKey instantiates a new AuthAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAPIKeyWithDefaults

`func NewAuthAPIKeyWithDefaults() *AuthAPIKey`

NewAuthAPIKeyWithDefaults instantiates a new AuthAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthAPIKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthAPIKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthAPIKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthAPIKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *AuthAPIKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthAPIKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthAPIKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuthAPIKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetScope

`func (o *AuthAPIKey) GetScope() AuthAPIKeyScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AuthAPIKey) GetScopeOk() (*AuthAPIKeyScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AuthAPIKey) SetScope(v AuthAPIKeyScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AuthAPIKey) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthAPIKey) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthAPIKey) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthAPIKey) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthAPIKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthAPIKey) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthAPIKey) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthAPIKey) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthAPIKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastUseAt

`func (o *AuthAPIKey) GetLastUseAt() int64`

GetLastUseAt returns the LastUseAt field if non-nil, zero value otherwise.

### GetLastUseAtOk

`func (o *AuthAPIKey) GetLastUseAtOk() (*int64, bool)`

GetLastUseAtOk returns a tuple with the LastUseAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUseAt

`func (o *AuthAPIKey) SetLastUseAt(v int64)`

SetLastUseAt sets LastUseAt field to given value.

### HasLastUseAt

`func (o *AuthAPIKey) HasLastUseAt() bool`

HasLastUseAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AuthAPIKey) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthAPIKey) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthAPIKey) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AuthAPIKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDescription

`func (o *AuthAPIKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthAPIKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthAPIKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthAPIKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *AuthAPIKey) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthAPIKey) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthAPIKey) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthAPIKey) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAdmin

`func (o *AuthAPIKey) GetAdmin() string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AuthAPIKey) GetAdminOk() (*string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AuthAPIKey) SetAdmin(v string)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *AuthAPIKey) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


