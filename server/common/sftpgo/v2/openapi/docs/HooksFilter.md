# HooksFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAuthDisabled** | Pointer to **bool** | If true, the external auth hook, if defined, will not be executed | [optional] 
**PreLoginDisabled** | Pointer to **bool** | If true, the pre-login hook, if defined, will not be executed | [optional] 
**CheckPasswordDisabled** | Pointer to **bool** | If true, the check password hook, if defined, will not be executed | [optional] 

## Methods

### NewHooksFilter

`func NewHooksFilter() *HooksFilter`

NewHooksFilter instantiates a new HooksFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHooksFilterWithDefaults

`func NewHooksFilterWithDefaults() *HooksFilter`

NewHooksFilterWithDefaults instantiates a new HooksFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAuthDisabled

`func (o *HooksFilter) GetExternalAuthDisabled() bool`

GetExternalAuthDisabled returns the ExternalAuthDisabled field if non-nil, zero value otherwise.

### GetExternalAuthDisabledOk

`func (o *HooksFilter) GetExternalAuthDisabledOk() (*bool, bool)`

GetExternalAuthDisabledOk returns a tuple with the ExternalAuthDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuthDisabled

`func (o *HooksFilter) SetExternalAuthDisabled(v bool)`

SetExternalAuthDisabled sets ExternalAuthDisabled field to given value.

### HasExternalAuthDisabled

`func (o *HooksFilter) HasExternalAuthDisabled() bool`

HasExternalAuthDisabled returns a boolean if a field has been set.

### GetPreLoginDisabled

`func (o *HooksFilter) GetPreLoginDisabled() bool`

GetPreLoginDisabled returns the PreLoginDisabled field if non-nil, zero value otherwise.

### GetPreLoginDisabledOk

`func (o *HooksFilter) GetPreLoginDisabledOk() (*bool, bool)`

GetPreLoginDisabledOk returns a tuple with the PreLoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLoginDisabled

`func (o *HooksFilter) SetPreLoginDisabled(v bool)`

SetPreLoginDisabled sets PreLoginDisabled field to given value.

### HasPreLoginDisabled

`func (o *HooksFilter) HasPreLoginDisabled() bool`

HasPreLoginDisabled returns a boolean if a field has been set.

### GetCheckPasswordDisabled

`func (o *HooksFilter) GetCheckPasswordDisabled() bool`

GetCheckPasswordDisabled returns the CheckPasswordDisabled field if non-nil, zero value otherwise.

### GetCheckPasswordDisabledOk

`func (o *HooksFilter) GetCheckPasswordDisabledOk() (*bool, bool)`

GetCheckPasswordDisabledOk returns a tuple with the CheckPasswordDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordDisabled

`func (o *HooksFilter) SetCheckPasswordDisabled(v bool)`

SetCheckPasswordDisabled sets CheckPasswordDisabled field to given value.

### HasCheckPasswordDisabled

`func (o *HooksFilter) HasCheckPasswordDisabled() bool`

HasCheckPasswordDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


