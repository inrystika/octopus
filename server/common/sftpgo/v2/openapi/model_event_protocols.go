/*
SFTPGo

SFTPGo allows to securely share your files over SFTP, HTTP and optionally FTP/S and WebDAV as well. Several storage backends are supported and they are configurable per user, so you can serve a local directory for a user and an S3 bucket (or part of it) for another one. SFTPGo also supports virtual folders, a virtual folder can use any of the supported storage backends. So you can have, for example, an S3 user that exposes a GCS bucket (or part of it) on a specified path and an encrypted local filesystem on another one. Virtual folders can be private or shared among multiple users, for shared virtual folders you can define different quota limits for each user. SFTPGo allows to create HTTP/S links to externally share files and folders securely, by setting limits to the number of downloads/uploads, protecting the share with a password, limiting access by source IP address, setting an automatic expiration date. 

API version: 2.2.2-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EventProtocols Protocols:   * `SSH` - SSH commands   * `SFTP` - SFTP protocol   * `FTP` - plain FTP and FTPES/FTPS   * `DAV` - WebDAV   * `HTTP` - WebClient/REST API   * `DataRetention` - the event is generated by a data retention check 
type EventProtocols string

// List of EventProtocols
const (
	EVENTPROTOCOLS_SSH EventProtocols = "SSH"
	EVENTPROTOCOLS_SFTP EventProtocols = "SFTP"
	EVENTPROTOCOLS_SCP EventProtocols = "SCP"
	EVENTPROTOCOLS_FTP EventProtocols = "FTP"
	EVENTPROTOCOLS_DAV EventProtocols = "DAV"
	EVENTPROTOCOLS_HTTP EventProtocols = "HTTP"
	EVENTPROTOCOLS_DATA_RETENTION EventProtocols = "DataRetention"
)

// All allowed values of EventProtocols enum
var AllowedEventProtocolsEnumValues = []EventProtocols{
	"SSH",
	"SFTP",
	"SCP",
	"FTP",
	"DAV",
	"HTTP",
	"DataRetention",
}

func (v *EventProtocols) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventProtocols(value)
	for _, existing := range AllowedEventProtocolsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventProtocols", value)
}

// NewEventProtocolsFromValue returns a pointer to a valid EventProtocols
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventProtocolsFromValue(v string) (*EventProtocols, error) {
	ev := EventProtocols(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventProtocols: valid values are %v", v, AllowedEventProtocolsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventProtocols) IsValid() bool {
	for _, existing := range AllowedEventProtocolsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventProtocols value
func (v EventProtocols) Ptr() *EventProtocols {
	return &v
}

type NullableEventProtocols struct {
	value *EventProtocols
	isSet bool
}

func (v NullableEventProtocols) Get() *EventProtocols {
	return v.value
}

func (v *NullableEventProtocols) Set(val *EventProtocols) {
	v.value = val
	v.isSet = true
}

func (v NullableEventProtocols) IsSet() bool {
	return v.isSet
}

func (v *NullableEventProtocols) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventProtocols(val *EventProtocols) *NullableEventProtocols {
	return &NullableEventProtocols{value: val, isSet: true}
}

func (v NullableEventProtocols) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventProtocols) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

