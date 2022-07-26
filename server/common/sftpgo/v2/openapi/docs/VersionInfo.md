# VersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**BuildDate** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** | Features for the current build. Available features are \&quot;portable\&quot;, \&quot;bolt\&quot;, \&quot;mysql\&quot;, \&quot;sqlite\&quot;, \&quot;pgsql\&quot;, \&quot;s3\&quot;, \&quot;gcs\&quot;, \&quot;metrics\&quot;. If a feature is available it has a \&quot;+\&quot; prefix, otherwise a \&quot;-\&quot; prefix | [optional] 

## Methods

### NewVersionInfo

`func NewVersionInfo() *VersionInfo`

NewVersionInfo instantiates a new VersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoWithDefaults

`func NewVersionInfoWithDefaults() *VersionInfo`

NewVersionInfoWithDefaults instantiates a new VersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VersionInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuildDate

`func (o *VersionInfo) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *VersionInfo) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *VersionInfo) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *VersionInfo) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetCommitHash

`func (o *VersionInfo) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *VersionInfo) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *VersionInfo) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *VersionInfo) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetFeatures

`func (o *VersionInfo) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *VersionInfo) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *VersionInfo) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *VersionInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


