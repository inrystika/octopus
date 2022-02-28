# AdminProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AllowApiKeyAuth** | Pointer to **bool** | If enabled, you can impersonate this admin, in REST API, using an API key. If disabled admin credentials are required for impersonation | [optional] 

## Methods

### NewAdminProfile

`func NewAdminProfile() *AdminProfile`

NewAdminProfile instantiates a new AdminProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminProfileWithDefaults

`func NewAdminProfileWithDefaults() *AdminProfile`

NewAdminProfileWithDefaults instantiates a new AdminProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AdminProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdminProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *AdminProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowApiKeyAuth

`func (o *AdminProfile) GetAllowApiKeyAuth() bool`

GetAllowApiKeyAuth returns the AllowApiKeyAuth field if non-nil, zero value otherwise.

### GetAllowApiKeyAuthOk

`func (o *AdminProfile) GetAllowApiKeyAuthOk() (*bool, bool)`

GetAllowApiKeyAuthOk returns a tuple with the AllowApiKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiKeyAuth

`func (o *AdminProfile) SetAllowApiKeyAuth(v bool)`

SetAllowApiKeyAuth sets AllowApiKeyAuth field to given value.

### HasAllowApiKeyAuth

`func (o *AdminProfile) HasAllowApiKeyAuth() bool`

HasAllowApiKeyAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


