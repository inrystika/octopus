# PatternsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | exposed virtual path, if no other specific filter is defined, the filter applies for sub directories too. For example if filters are defined for the paths \&quot;/\&quot; and \&quot;/sub\&quot; then the filters for \&quot;/\&quot; are applied for any file outside the \&quot;/sub\&quot; directory | [optional] 
**AllowedPatterns** | Pointer to **[]string** | list of, case insensitive, allowed shell like patterns. | [optional] 
**DeniedPatterns** | Pointer to **[]string** | list of, case insensitive, denied shell like patterns. Denied patterns are evaluated before the allowed ones | [optional] 
**DenyPolicy** | Pointer to **int32** | Deny policies   * &#x60;0&#x60; - default policy. Denied files/directories matching the filters are visible in directory listing but cannot be uploaded/downloaded/overwritten/renamed   * &#x60;1&#x60; - deny policy hide. This policy applies the same restrictions as the default one and denied files/directories matching the filters will also be hidden in directory listing. This mode may cause performance issues for large directories  | [optional] 

## Methods

### NewPatternsFilter

`func NewPatternsFilter() *PatternsFilter`

NewPatternsFilter instantiates a new PatternsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatternsFilterWithDefaults

`func NewPatternsFilterWithDefaults() *PatternsFilter`

NewPatternsFilterWithDefaults instantiates a new PatternsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PatternsFilter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatternsFilter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatternsFilter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatternsFilter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetAllowedPatterns

`func (o *PatternsFilter) GetAllowedPatterns() []string`

GetAllowedPatterns returns the AllowedPatterns field if non-nil, zero value otherwise.

### GetAllowedPatternsOk

`func (o *PatternsFilter) GetAllowedPatternsOk() (*[]string, bool)`

GetAllowedPatternsOk returns a tuple with the AllowedPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPatterns

`func (o *PatternsFilter) SetAllowedPatterns(v []string)`

SetAllowedPatterns sets AllowedPatterns field to given value.

### HasAllowedPatterns

`func (o *PatternsFilter) HasAllowedPatterns() bool`

HasAllowedPatterns returns a boolean if a field has been set.

### GetDeniedPatterns

`func (o *PatternsFilter) GetDeniedPatterns() []string`

GetDeniedPatterns returns the DeniedPatterns field if non-nil, zero value otherwise.

### GetDeniedPatternsOk

`func (o *PatternsFilter) GetDeniedPatternsOk() (*[]string, bool)`

GetDeniedPatternsOk returns a tuple with the DeniedPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedPatterns

`func (o *PatternsFilter) SetDeniedPatterns(v []string)`

SetDeniedPatterns sets DeniedPatterns field to given value.

### HasDeniedPatterns

`func (o *PatternsFilter) HasDeniedPatterns() bool`

HasDeniedPatterns returns a boolean if a field has been set.

### GetDenyPolicy

`func (o *PatternsFilter) GetDenyPolicy() int32`

GetDenyPolicy returns the DenyPolicy field if non-nil, zero value otherwise.

### GetDenyPolicyOk

`func (o *PatternsFilter) GetDenyPolicyOk() (*int32, bool)`

GetDenyPolicyOk returns a tuple with the DenyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPolicy

`func (o *PatternsFilter) SetDenyPolicy(v int32)`

SetDenyPolicy sets DenyPolicy field to given value.

### HasDenyPolicy

`func (o *PatternsFilter) HasDenyPolicy() bool`

HasDenyPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


