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

// checks if the ParameterCopy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterCopy{}

// ParameterCopy A single parameter inside of a project.
type ParameterCopy struct {
	// The parameter name.
	Name string `json:"name"`
	// A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// The project url.
	Project string `json:"project"`
}

// NewParameterCopy instantiates a new ParameterCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterCopy(name string, project string) *ParameterCopy {
	this := ParameterCopy{}
	this.Name = name
	this.Project = project
	return &this
}

// NewParameterCopyWithDefaults instantiates a new ParameterCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterCopyWithDefaults() *ParameterCopy {
	this := ParameterCopy{}
	return &this
}

// GetName returns the Name field value
func (o *ParameterCopy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterCopy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterCopy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterCopy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterCopy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterCopy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterCopy) SetDescription(v string) {
	o.Description = &v
}

// GetProject returns the Project field value
func (o *ParameterCopy) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ParameterCopy) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ParameterCopy) SetProject(v string) {
	o.Project = v
}

func (o ParameterCopy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterCopy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["project"] = o.Project
	return toSerialize, nil
}

type NullableParameterCopy struct {
	value *ParameterCopy
	isSet bool
}

func (v NullableParameterCopy) Get() *ParameterCopy {
	return v.value
}

func (v *NullableParameterCopy) Set(val *ParameterCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterCopy(val *ParameterCopy) *NullableParameterCopy {
	return &NullableParameterCopy{value: val, isSet: true}
}

func (v NullableParameterCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


