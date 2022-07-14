/*
SFTPGo

SFTPGo allows to securely share your files over SFTP, HTTP and optionally FTP/S and WebDAV as well. Several storage backends are supported and they are configurable per user, so you can serve a local directory for a user and an S3 bucket (or part of it) for another one. SFTPGo also supports virtual folders, a virtual folder can use any of the supported storage backends. So you can have, for example, an S3 user that exposes a GCS bucket (or part of it) on a specified path and an encrypted local filesystem on another one. Virtual folders can be private or shared among multiple users, for shared virtual folders you can define different quota limits for each user. SFTPGo allows to create HTTP/S links to externally share files and folders securely, by setting limits to the number of downloads/uploads, protecting the share with a password, limiting access by source IP address, setting an automatic expiration date. 

API version: 2.2.2-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UserProfile struct for UserProfile
type UserProfile struct {
	Email *string `json:"email,omitempty"`
	Description *string `json:"description,omitempty"`
	// If enabled, you can impersonate this user, in REST API, using an API key. If disabled user credentials are required for impersonation
	AllowApiKeyAuth *bool `json:"allow_api_key_auth,omitempty"`
	PublicKeys []string `json:"public_keys,omitempty"`
}

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile() *UserProfile {
	this := UserProfile{}
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserProfile) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserProfile) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserProfile) SetEmail(v string) {
	o.Email = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserProfile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserProfile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserProfile) SetDescription(v string) {
	o.Description = &v
}

// GetAllowApiKeyAuth returns the AllowApiKeyAuth field value if set, zero value otherwise.
func (o *UserProfile) GetAllowApiKeyAuth() bool {
	if o == nil || o.AllowApiKeyAuth == nil {
		var ret bool
		return ret
	}
	return *o.AllowApiKeyAuth
}

// GetAllowApiKeyAuthOk returns a tuple with the AllowApiKeyAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetAllowApiKeyAuthOk() (*bool, bool) {
	if o == nil || o.AllowApiKeyAuth == nil {
		return nil, false
	}
	return o.AllowApiKeyAuth, true
}

// HasAllowApiKeyAuth returns a boolean if a field has been set.
func (o *UserProfile) HasAllowApiKeyAuth() bool {
	if o != nil && o.AllowApiKeyAuth != nil {
		return true
	}

	return false
}

// SetAllowApiKeyAuth gets a reference to the given bool and assigns it to the AllowApiKeyAuth field.
func (o *UserProfile) SetAllowApiKeyAuth(v bool) {
	o.AllowApiKeyAuth = &v
}

// GetPublicKeys returns the PublicKeys field value if set, zero value otherwise.
func (o *UserProfile) GetPublicKeys() []string {
	if o == nil || o.PublicKeys == nil {
		var ret []string
		return ret
	}
	return o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetPublicKeysOk() ([]string, bool) {
	if o == nil || o.PublicKeys == nil {
		return nil, false
	}
	return o.PublicKeys, true
}

// HasPublicKeys returns a boolean if a field has been set.
func (o *UserProfile) HasPublicKeys() bool {
	if o != nil && o.PublicKeys != nil {
		return true
	}

	return false
}

// SetPublicKeys gets a reference to the given []string and assigns it to the PublicKeys field.
func (o *UserProfile) SetPublicKeys(v []string) {
	o.PublicKeys = v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AllowApiKeyAuth != nil {
		toSerialize["allow_api_key_auth"] = o.AllowApiKeyAuth
	}
	if o.PublicKeys != nil {
		toSerialize["public_keys"] = o.PublicKeys
	}
	return json.Marshal(toSerialize)
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


