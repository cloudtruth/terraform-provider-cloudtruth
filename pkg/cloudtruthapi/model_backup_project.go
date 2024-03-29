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

// checks if the BackupProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupProject{}

// BackupProject Environment, parameter-type, and project (including parameters and values) data at a point in time.
type BackupProject struct {
	Parameters map[string]BackupParameter `json:"parameters"`
	Templates map[string]BackupTemplate `json:"templates"`
	Name string `json:"name"`
	Parent NullableString `json:"parent,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewBackupProject instantiates a new BackupProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupProject(parameters map[string]BackupParameter, templates map[string]BackupTemplate, name string) *BackupProject {
	this := BackupProject{}
	this.Parameters = parameters
	this.Templates = templates
	this.Name = name
	return &this
}

// NewBackupProjectWithDefaults instantiates a new BackupProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupProjectWithDefaults() *BackupProject {
	this := BackupProject{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *BackupProject) GetParameters() map[string]BackupParameter {
	if o == nil {
		var ret map[string]BackupParameter
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *BackupProject) GetParametersOk() (*map[string]BackupParameter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *BackupProject) SetParameters(v map[string]BackupParameter) {
	o.Parameters = v
}

// GetTemplates returns the Templates field value
func (o *BackupProject) GetTemplates() map[string]BackupTemplate {
	if o == nil {
		var ret map[string]BackupTemplate
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *BackupProject) GetTemplatesOk() (*map[string]BackupTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Templates, true
}

// SetTemplates sets field value
func (o *BackupProject) SetTemplates(v map[string]BackupTemplate) {
	o.Templates = v
}

// GetName returns the Name field value
func (o *BackupProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BackupProject) SetName(v string) {
	o.Name = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupProject) GetParent() string {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProject) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *BackupProject) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *BackupProject) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *BackupProject) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *BackupProject) UnsetParent() {
	o.Parent.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupProject) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *BackupProject) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *BackupProject) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *BackupProject) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *BackupProject) UnsetDescription() {
	o.Description.Unset()
}

func (o BackupProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	toSerialize["templates"] = o.Templates
	toSerialize["name"] = o.Name
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableBackupProject struct {
	value *BackupProject
	isSet bool
}

func (v NullableBackupProject) Get() *BackupProject {
	return v.value
}

func (v *NullableBackupProject) Set(val *BackupProject) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupProject) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupProject(val *BackupProject) *NullableBackupProject {
	return &NullableBackupProject{value: val, isSet: true}
}

func (v NullableBackupProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


