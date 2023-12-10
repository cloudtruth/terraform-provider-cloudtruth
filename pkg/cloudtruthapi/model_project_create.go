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

// checks if the ProjectCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCreate{}

// ProjectCreate struct for ProjectCreate
type ProjectCreate struct {
	// The project name.
	Name string `json:"name"`
	// A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// A regular expression parameter names must match
	ParameterNamePattern *string `json:"parameter_name_pattern,omitempty"`
	// Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project.
	DependsOn NullableString `json:"depends_on,omitempty"`
}

type _ProjectCreate ProjectCreate

// NewProjectCreate instantiates a new ProjectCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreate(name string) *ProjectCreate {
	this := ProjectCreate{}
	this.Name = name
	return &this
}

// NewProjectCreateWithDefaults instantiates a new ProjectCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateWithDefaults() *ProjectCreate {
	this := ProjectCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectCreate) SetDescription(v string) {
	o.Description = &v
}

// GetParameterNamePattern returns the ParameterNamePattern field value if set, zero value otherwise.
func (o *ProjectCreate) GetParameterNamePattern() string {
	if o == nil || IsNil(o.ParameterNamePattern) {
		var ret string
		return ret
	}
	return *o.ParameterNamePattern
}

// GetParameterNamePatternOk returns a tuple with the ParameterNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetParameterNamePatternOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterNamePattern) {
		return nil, false
	}
	return o.ParameterNamePattern, true
}

// HasParameterNamePattern returns a boolean if a field has been set.
func (o *ProjectCreate) HasParameterNamePattern() bool {
	if o != nil && !IsNil(o.ParameterNamePattern) {
		return true
	}

	return false
}

// SetParameterNamePattern gets a reference to the given string and assigns it to the ParameterNamePattern field.
func (o *ProjectCreate) SetParameterNamePattern(v string) {
	o.ParameterNamePattern = &v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectCreate) GetDependsOn() string {
	if o == nil || IsNil(o.DependsOn.Get()) {
		var ret string
		return ret
	}
	return *o.DependsOn.Get()
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCreate) GetDependsOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn.Get(), o.DependsOn.IsSet()
}

// HasDependsOn returns a boolean if a field has been set.
func (o *ProjectCreate) HasDependsOn() bool {
	if o != nil && o.DependsOn.IsSet() {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given NullableString and assigns it to the DependsOn field.
func (o *ProjectCreate) SetDependsOn(v string) {
	o.DependsOn.Set(&v)
}
// SetDependsOnNil sets the value for DependsOn to be an explicit nil
func (o *ProjectCreate) SetDependsOnNil() {
	o.DependsOn.Set(nil)
}

// UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
func (o *ProjectCreate) UnsetDependsOn() {
	o.DependsOn.Unset()
}

func (o ProjectCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ParameterNamePattern) {
		toSerialize["parameter_name_pattern"] = o.ParameterNamePattern
	}
	if o.DependsOn.IsSet() {
		toSerialize["depends_on"] = o.DependsOn.Get()
	}
	return toSerialize, nil
}

func (o *ProjectCreate) UnmarshalJSON(bytes []byte) (err error) {
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

	varProjectCreate := _ProjectCreate{}

	err = json.Unmarshal(bytes, &varProjectCreate)

	if err != nil {
		return err
	}

	*o = ProjectCreate(varProjectCreate)

	return err
}

type NullableProjectCreate struct {
	value *ProjectCreate
	isSet bool
}

func (v NullableProjectCreate) Get() *ProjectCreate {
	return v.value
}

func (v *NullableProjectCreate) Set(val *ProjectCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreate(val *ProjectCreate) *NullableProjectCreate {
	return &NullableProjectCreate{value: val, isSet: true}
}

func (v NullableProjectCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


