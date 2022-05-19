/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudtruthapi

import (
	"encoding/json"
	"time"
)

// BackupDataSnapshot Environment, parameter-type, and project (including parameters and values) data at a point in time.
type BackupDataSnapshot struct {
	Environments map[string]BackupEnvironment `json:"environments"`
	Types map[string]BackupParameterType `json:"types"`
	Projects map[string]BackupProject `json:"projects"`
	Timestamp time.Time `json:"timestamp"`
}

// NewBackupDataSnapshot instantiates a new BackupDataSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupDataSnapshot(environments map[string]BackupEnvironment, types map[string]BackupParameterType, projects map[string]BackupProject, timestamp time.Time) *BackupDataSnapshot {
	this := BackupDataSnapshot{}
	this.Environments = environments
	this.Types = types
	this.Projects = projects
	this.Timestamp = timestamp
	return &this
}

// NewBackupDataSnapshotWithDefaults instantiates a new BackupDataSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupDataSnapshotWithDefaults() *BackupDataSnapshot {
	this := BackupDataSnapshot{}
	return &this
}

// GetEnvironments returns the Environments field value
func (o *BackupDataSnapshot) GetEnvironments() map[string]BackupEnvironment {
	if o == nil {
		var ret map[string]BackupEnvironment
		return ret
	}

	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *BackupDataSnapshot) GetEnvironmentsOk() (*map[string]BackupEnvironment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environments, true
}

// SetEnvironments sets field value
func (o *BackupDataSnapshot) SetEnvironments(v map[string]BackupEnvironment) {
	o.Environments = v
}

// GetTypes returns the Types field value
func (o *BackupDataSnapshot) GetTypes() map[string]BackupParameterType {
	if o == nil {
		var ret map[string]BackupParameterType
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *BackupDataSnapshot) GetTypesOk() (*map[string]BackupParameterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Types, true
}

// SetTypes sets field value
func (o *BackupDataSnapshot) SetTypes(v map[string]BackupParameterType) {
	o.Types = v
}

// GetProjects returns the Projects field value
func (o *BackupDataSnapshot) GetProjects() map[string]BackupProject {
	if o == nil {
		var ret map[string]BackupProject
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *BackupDataSnapshot) GetProjectsOk() (*map[string]BackupProject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projects, true
}

// SetProjects sets field value
func (o *BackupDataSnapshot) SetProjects(v map[string]BackupProject) {
	o.Projects = v
}

// GetTimestamp returns the Timestamp field value
func (o *BackupDataSnapshot) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BackupDataSnapshot) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BackupDataSnapshot) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o BackupDataSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environments"] = o.Environments
	}
	if true {
		toSerialize["types"] = o.Types
	}
	if true {
		toSerialize["projects"] = o.Projects
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableBackupDataSnapshot struct {
	value *BackupDataSnapshot
	isSet bool
}

func (v NullableBackupDataSnapshot) Get() *BackupDataSnapshot {
	return v.value
}

func (v *NullableBackupDataSnapshot) Set(val *BackupDataSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupDataSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupDataSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupDataSnapshot(val *BackupDataSnapshot) *NullableBackupDataSnapshot {
	return &NullableBackupDataSnapshot{value: val, isSet: true}
}

func (v NullableBackupDataSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupDataSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


