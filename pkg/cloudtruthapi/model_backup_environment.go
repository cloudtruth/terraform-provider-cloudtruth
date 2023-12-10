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
	"fmt"
)

// checks if the BackupEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupEnvironment{}

// BackupEnvironment Basic environment data at a point in time.
type BackupEnvironment struct {
	Name string `json:"name"`
	Parent NullableString `json:"parent,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

type _BackupEnvironment BackupEnvironment

// NewBackupEnvironment instantiates a new BackupEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupEnvironment(name string) *BackupEnvironment {
	this := BackupEnvironment{}
	this.Name = name
	return &this
}

// NewBackupEnvironmentWithDefaults instantiates a new BackupEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupEnvironmentWithDefaults() *BackupEnvironment {
	this := BackupEnvironment{}
	return &this
}

// GetName returns the Name field value
func (o *BackupEnvironment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupEnvironment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BackupEnvironment) SetName(v string) {
	o.Name = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupEnvironment) GetParent() string {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupEnvironment) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *BackupEnvironment) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *BackupEnvironment) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *BackupEnvironment) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *BackupEnvironment) UnsetParent() {
	o.Parent.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupEnvironment) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupEnvironment) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *BackupEnvironment) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *BackupEnvironment) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *BackupEnvironment) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *BackupEnvironment) UnsetDescription() {
	o.Description.Unset()
}

func (o BackupEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *BackupEnvironment) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBackupEnvironment := _BackupEnvironment{}

	err = json.Unmarshal(bytes, &varBackupEnvironment)

	if err != nil {
		return err
	}

	*o = BackupEnvironment(varBackupEnvironment)

	return err
}

type NullableBackupEnvironment struct {
	value *BackupEnvironment
	isSet bool
}

func (v NullableBackupEnvironment) Get() *BackupEnvironment {
	return v.value
}

func (v *NullableBackupEnvironment) Set(val *BackupEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupEnvironment(val *BackupEnvironment) *NullableBackupEnvironment {
	return &NullableBackupEnvironment{value: val, isSet: true}
}

func (v NullableBackupEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


