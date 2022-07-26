# Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | status:   * &#x60;0&#x60; user is disabled, login is not allowed   * &#x60;1&#x60; user is enabled  | [optional] 
**Username** | Pointer to **string** | username is unique | [optional] 
**Description** | Pointer to **string** | optional description, for example the admin full name | [optional] 
**Password** | Pointer to **string** | Admin password. For security reasons this field is omitted when you search/get admins | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]AdminPermissions**](AdminPermissions.md) |  | [optional] 
**Filters** | Pointer to [**AdminFilters**](AdminFilters.md) |  | [optional] 
**AdditionalInfo** | Pointer to **string** | Free form text field | [optional] 
**CreatedAt** | Pointer to **int64** | creation time as unix timestamp in milliseconds. It will be 0 for admins created before v2.2.0 | [optional] 
**UpdatedAt** | Pointer to **int64** | last update time as unix timestamp in milliseconds | [optional] 
**LastLogin** | Pointer to **int64** | Last user login as unix timestamp in milliseconds. It is saved at most once every 10 minutes | [optional] 

## Methods

### NewAdmin

`func NewAdmin() *Admin`

NewAdmin instantiates a new Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWithDefaults

`func NewAdminWithDefaults() *Admin`

NewAdminWithDefaults instantiates a new Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Admin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Admin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Admin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Admin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Admin) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Admin) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Admin) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Admin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *Admin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Admin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Admin) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Admin) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDescription

`func (o *Admin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Admin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Admin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Admin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPassword

`func (o *Admin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Admin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Admin) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Admin) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *Admin) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Admin) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Admin) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Admin) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPermissions

`func (o *Admin) GetPermissions() []AdminPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Admin) GetPermissionsOk() (*[]AdminPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Admin) SetPermissions(v []AdminPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Admin) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetFilters

`func (o *Admin) GetFilters() AdminFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Admin) GetFiltersOk() (*AdminFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Admin) SetFilters(v AdminFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Admin) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Admin) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Admin) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Admin) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Admin) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Admin) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Admin) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Admin) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Admin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Admin) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Admin) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Admin) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Admin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *Admin) GetLastLogin() int64`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Admin) GetLastLoginOk() (*int64, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Admin) SetLastLogin(v int64)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Admin) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


