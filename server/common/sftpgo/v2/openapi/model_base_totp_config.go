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

// BaseTOTPConfig struct for BaseTOTPConfig
type BaseTOTPConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
	// This name must be defined within the \"totp\" section of the SFTPGo configuration file. You will be unable to save a user/admin referencing a missing config_name
	ConfigName *string `json:"config_name,omitempty"`
	Secret *Secret `json:"secret,omitempty"`
}

// NewBaseTOTPConfig instantiates a new BaseTOTPConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTOTPConfig() *BaseTOTPConfig {
	this := BaseTOTPConfig{}
	return &this
}

// NewBaseTOTPConfigWithDefaults instantiates a new BaseTOTPConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTOTPConfigWithDefaults() *BaseTOTPConfig {
	this := BaseTOTPConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BaseTOTPConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTOTPConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BaseTOTPConfig) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BaseTOTPConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetConfigName returns the ConfigName field value if set, zero value otherwise.
func (o *BaseTOTPConfig) GetConfigName() string {
	if o == nil || o.ConfigName == nil {
		var ret string
		return ret
	}
	return *o.ConfigName
}

// GetConfigNameOk returns a tuple with the ConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTOTPConfig) GetConfigNameOk() (*string, bool) {
	if o == nil || o.ConfigName == nil {
		return nil, false
	}
	return o.ConfigName, true
}

// HasConfigName returns a boolean if a field has been set.
func (o *BaseTOTPConfig) HasConfigName() bool {
	if o != nil && o.ConfigName != nil {
		return true
	}

	return false
}

// SetConfigName gets a reference to the given string and assigns it to the ConfigName field.
func (o *BaseTOTPConfig) SetConfigName(v string) {
	o.ConfigName = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BaseTOTPConfig) GetSecret() Secret {
	if o == nil || o.Secret == nil {
		var ret Secret
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTOTPConfig) GetSecretOk() (*Secret, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BaseTOTPConfig) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given Secret and assigns it to the Secret field.
func (o *BaseTOTPConfig) SetSecret(v Secret) {
	o.Secret = &v
}

func (o BaseTOTPConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ConfigName != nil {
		toSerialize["config_name"] = o.ConfigName
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableBaseTOTPConfig struct {
	value *BaseTOTPConfig
	isSet bool
}

func (v NullableBaseTOTPConfig) Get() *BaseTOTPConfig {
	return v.value
}

func (v *NullableBaseTOTPConfig) Set(val *BaseTOTPConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTOTPConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTOTPConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTOTPConfig(val *BaseTOTPConfig) *NullableBaseTOTPConfig {
	return &NullableBaseTOTPConfig{value: val, isSet: true}
}

func (v NullableBaseTOTPConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTOTPConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


