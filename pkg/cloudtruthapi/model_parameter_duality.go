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

// ParameterDuality Details about a parameter at two timepoints.  If t1 or t2 is null then the parameter did not exist at that timepoint.
type ParameterDuality struct {
	T1 NullableParameterDualityT1 `json:"t1"`
	T2 NullableParameterDualityT1 `json:"t2"`
}

// NewParameterDuality instantiates a new ParameterDuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterDuality(t1 NullableParameterDualityT1, t2 NullableParameterDualityT1) *ParameterDuality {
	this := ParameterDuality{}
	this.T1 = t1
	this.T2 = t2
	return &this
}

// NewParameterDualityWithDefaults instantiates a new ParameterDuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterDualityWithDefaults() *ParameterDuality {
	this := ParameterDuality{}
	return &this
}

// GetT1 returns the T1 field value
// If the value is explicit nil, the zero value for ParameterDualityT1 will be returned
func (o *ParameterDuality) GetT1() ParameterDualityT1 {
	if o == nil || o.T1.Get() == nil {
		var ret ParameterDualityT1
		return ret
	}

	return *o.T1.Get()
}

// GetT1Ok returns a tuple with the T1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterDuality) GetT1Ok() (*ParameterDualityT1, bool) {
	if o == nil {
		return nil, false
	}
	return o.T1.Get(), o.T1.IsSet()
}

// SetT1 sets field value
func (o *ParameterDuality) SetT1(v ParameterDualityT1) {
	o.T1.Set(&v)
}

// GetT2 returns the T2 field value
// If the value is explicit nil, the zero value for ParameterDualityT1 will be returned
func (o *ParameterDuality) GetT2() ParameterDualityT1 {
	if o == nil || o.T2.Get() == nil {
		var ret ParameterDualityT1
		return ret
	}

	return *o.T2.Get()
}

// GetT2Ok returns a tuple with the T2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterDuality) GetT2Ok() (*ParameterDualityT1, bool) {
	if o == nil {
		return nil, false
	}
	return o.T2.Get(), o.T2.IsSet()
}

// SetT2 sets field value
func (o *ParameterDuality) SetT2(v ParameterDualityT1) {
	o.T2.Set(&v)
}

func (o ParameterDuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["t1"] = o.T1.Get()
	}
	if true {
		toSerialize["t2"] = o.T2.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableParameterDuality struct {
	value *ParameterDuality
	isSet bool
}

func (v NullableParameterDuality) Get() *ParameterDuality {
	return v.value
}

func (v *NullableParameterDuality) Set(val *ParameterDuality) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterDuality) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterDuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterDuality(val *ParameterDuality) *NullableParameterDuality {
	return &NullableParameterDuality{value: val, isSet: true}
}

func (v NullableParameterDuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterDuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
