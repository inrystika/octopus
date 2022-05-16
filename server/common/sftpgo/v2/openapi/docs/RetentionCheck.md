# RetentionCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | username to which the retention check refers | [optional] 
**Folders** | Pointer to [**[]FolderRetention**](FolderRetention.md) |  | [optional] 
**StartTime** | Pointer to **int64** | check start time as unix timestamp in milliseconds | [optional] 
**Notifications** | Pointer to [**[]RetentionCheckNotification**](RetentionCheckNotification.md) |  | [optional] 
**Email** | Pointer to **string** | if the notification method is set to \&quot;Email\&quot;, this is the e-mail address that receives the retention check report. This field is automatically set to the email address associated with the administrator starting the check | [optional] 

## Methods

### NewRetentionCheck

`func NewRetentionCheck() *RetentionCheck`

NewRetentionCheck instantiates a new RetentionCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionCheckWithDefaults

`func NewRetentionCheckWithDefaults() *RetentionCheck`

NewRetentionCheckWithDefaults instantiates a new RetentionCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RetentionCheck) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RetentionCheck) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RetentionCheck) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RetentionCheck) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFolders

`func (o *RetentionCheck) GetFolders() []FolderRetention`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *RetentionCheck) GetFoldersOk() (*[]FolderRetention, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *RetentionCheck) SetFolders(v []FolderRetention)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *RetentionCheck) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetStartTime

`func (o *RetentionCheck) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RetentionCheck) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RetentionCheck) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RetentionCheck) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetNotifications

`func (o *RetentionCheck) GetNotifications() []RetentionCheckNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *RetentionCheck) GetNotificationsOk() (*[]RetentionCheckNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *RetentionCheck) SetNotifications(v []RetentionCheckNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *RetentionCheck) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetEmail

`func (o *RetentionCheck) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RetentionCheck) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RetentionCheck) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RetentionCheck) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


