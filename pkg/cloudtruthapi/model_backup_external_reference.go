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
)

// checks if the BackupExternalReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupExternalReference{}

// BackupExternalReference External reference data at a point in time.
type BackupExternalReference struct {
	Fqn string `json:"fqn"`
	JmesPath NullableString `json:"jmes_path,omitempty"`
}

// NewBackupExternalReference instantiates a new BackupExternalReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupExternalReference(fqn string) *BackupExternalReference {
	this := BackupExternalReference{}
	this.Fqn = fqn
	return &this
}

// NewBackupExternalReferenceWithDefaults instantiates a new BackupExternalReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupExternalReferenceWithDefaults() *BackupExternalReference {
	this := BackupExternalReference{}
	return &this
}

// GetFqn returns the Fqn field value
func (o *BackupExternalReference) GetFqn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value
// and a boolean to check if the value has been set.
func (o *BackupExternalReference) GetFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqn, true
}

// SetFqn sets field value
func (o *BackupExternalReference) SetFqn(v string) {
	o.Fqn = v
}

// GetJmesPath returns the JmesPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupExternalReference) GetJmesPath() string {
	if o == nil || IsNil(o.JmesPath.Get()) {
		var ret string
		return ret
	}
	return *o.JmesPath.Get()
}

// GetJmesPathOk returns a tuple with the JmesPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupExternalReference) GetJmesPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JmesPath.Get(), o.JmesPath.IsSet()
}

// HasJmesPath returns a boolean if a field has been set.
func (o *BackupExternalReference) HasJmesPath() bool {
	if o != nil && o.JmesPath.IsSet() {
		return true
	}

	return false
}

// SetJmesPath gets a reference to the given NullableString and assigns it to the JmesPath field.
func (o *BackupExternalReference) SetJmesPath(v string) {
	o.JmesPath.Set(&v)
}
// SetJmesPathNil sets the value for JmesPath to be an explicit nil
func (o *BackupExternalReference) SetJmesPathNil() {
	o.JmesPath.Set(nil)
}

// UnsetJmesPath ensures that no value is present for JmesPath, not even an explicit nil
func (o *BackupExternalReference) UnsetJmesPath() {
	o.JmesPath.Unset()
}

func (o BackupExternalReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupExternalReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fqn"] = o.Fqn
	if o.JmesPath.IsSet() {
		toSerialize["jmes_path"] = o.JmesPath.Get()
	}
	return toSerialize, nil
}

type NullableBackupExternalReference struct {
	value *BackupExternalReference
	isSet bool
}

func (v NullableBackupExternalReference) Get() *BackupExternalReference {
	return v.value
}

func (v *NullableBackupExternalReference) Set(val *BackupExternalReference) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupExternalReference) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupExternalReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupExternalReference(val *BackupExternalReference) *NullableBackupExternalReference {
	return &NullableBackupExternalReference{value: val, isSet: true}
}

func (v NullableBackupExternalReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupExternalReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


