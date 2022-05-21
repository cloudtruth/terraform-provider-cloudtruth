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

// ParameterTimelineEntryEnvironment struct for ParameterTimelineEntryEnvironment
type ParameterTimelineEntryEnvironment struct {
	// A unique identifier for the environment.
	Id string `json:"id"`
	// The environment name.
	Name string `json:"name"`
	// Indicates if the value change was direct or if it flowed into the environment. If `true` then the value was actually set directly into this environment. If `false` then the environment has no value set directly so it inherited the value from its parent.
	Override bool `json:"override"`
}

// NewParameterTimelineEntryEnvironment instantiates a new ParameterTimelineEntryEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterTimelineEntryEnvironment(id string, name string, override bool) *ParameterTimelineEntryEnvironment {
	this := ParameterTimelineEntryEnvironment{}
	this.Id = id
	this.Name = name
	this.Override = override
	return &this
}

// NewParameterTimelineEntryEnvironmentWithDefaults instantiates a new ParameterTimelineEntryEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterTimelineEntryEnvironmentWithDefaults() *ParameterTimelineEntryEnvironment {
	this := ParameterTimelineEntryEnvironment{}
	return &this
}

// GetId returns the Id field value
func (o *ParameterTimelineEntryEnvironment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntryEnvironment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterTimelineEntryEnvironment) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ParameterTimelineEntryEnvironment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntryEnvironment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterTimelineEntryEnvironment) SetName(v string) {
	o.Name = v
}

// GetOverride returns the Override field value
func (o *ParameterTimelineEntryEnvironment) GetOverride() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Override
}

// GetOverrideOk returns a tuple with the Override field value
// and a boolean to check if the value has been set.
func (o *ParameterTimelineEntryEnvironment) GetOverrideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Override, true
}

// SetOverride sets field value
func (o *ParameterTimelineEntryEnvironment) SetOverride(v bool) {
	o.Override = v
}

func (o ParameterTimelineEntryEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["override"] = o.Override
	}
	return json.Marshal(toSerialize)
}

type NullableParameterTimelineEntryEnvironment struct {
	value *ParameterTimelineEntryEnvironment
	isSet bool
}

func (v NullableParameterTimelineEntryEnvironment) Get() *ParameterTimelineEntryEnvironment {
	return v.value
}

func (v *NullableParameterTimelineEntryEnvironment) Set(val *ParameterTimelineEntryEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterTimelineEntryEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterTimelineEntryEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterTimelineEntryEnvironment(val *ParameterTimelineEntryEnvironment) *NullableParameterTimelineEntryEnvironment {
	return &NullableParameterTimelineEntryEnvironment{value: val, isSet: true}
}

func (v NullableParameterTimelineEntryEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterTimelineEntryEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
