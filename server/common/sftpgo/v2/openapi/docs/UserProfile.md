# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AllowApiKeyAuth** | Pointer to **bool** | If enabled, you can impersonate this user, in REST API, using an API key. If disabled user credentials are required for impersonation | [optional] 
**PublicKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserProfile

`func NewUserProfile() *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *UserProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowApiKeyAuth

`func (o *UserProfile) GetAllowApiKeyAuth() bool`

GetAllowApiKeyAuth returns the AllowApiKeyAuth field if non-nil, zero value otherwise.

### GetAllowApiKeyAuthOk

`func (o *UserProfile) GetAllowApiKeyAuthOk() (*bool, bool)`

GetAllowApiKeyAuthOk returns a tuple with the AllowApiKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiKeyAuth

`func (o *UserProfile) SetAllowApiKeyAuth(v bool)`

SetAllowApiKeyAuth sets AllowApiKeyAuth field to given value.

### HasAllowApiKeyAuth

`func (o *UserProfile) HasAllowApiKeyAuth() bool`

HasAllowApiKeyAuth returns a boolean if a field has been set.

### GetPublicKeys

`func (o *UserProfile) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *UserProfile) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *UserProfile) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *UserProfile) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


