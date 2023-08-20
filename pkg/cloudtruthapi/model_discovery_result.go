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

// checks if the DiscoveryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryResult{}

// DiscoveryResult struct for DiscoveryResult
type DiscoveryResult struct {
	Matched map[string]DiscoveredContent `json:"matched"`
	Skipped map[string]string `json:"skipped"`
}

// NewDiscoveryResult instantiates a new DiscoveryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryResult(matched map[string]DiscoveredContent, skipped map[string]string) *DiscoveryResult {
	this := DiscoveryResult{}
	this.Matched = matched
	this.Skipped = skipped
	return &this
}

// NewDiscoveryResultWithDefaults instantiates a new DiscoveryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryResultWithDefaults() *DiscoveryResult {
	this := DiscoveryResult{}
	return &this
}

// GetMatched returns the Matched field value
func (o *DiscoveryResult) GetMatched() map[string]DiscoveredContent {
	if o == nil {
		var ret map[string]DiscoveredContent
		return ret
	}

	return o.Matched
}

// GetMatchedOk returns a tuple with the Matched field value
// and a boolean to check if the value has been set.
func (o *DiscoveryResult) GetMatchedOk() (*map[string]DiscoveredContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matched, true
}

// SetMatched sets field value
func (o *DiscoveryResult) SetMatched(v map[string]DiscoveredContent) {
	o.Matched = v
}

// GetSkipped returns the Skipped field value
func (o *DiscoveryResult) GetSkipped() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value
// and a boolean to check if the value has been set.
func (o *DiscoveryResult) GetSkippedOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Skipped, true
}

// SetSkipped sets field value
func (o *DiscoveryResult) SetSkipped(v map[string]string) {
	o.Skipped = v
}

func (o DiscoveryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matched"] = o.Matched
	toSerialize["skipped"] = o.Skipped
	return toSerialize, nil
}

type NullableDiscoveryResult struct {
	value *DiscoveryResult
	isSet bool
}

func (v NullableDiscoveryResult) Get() *DiscoveryResult {
	return v.value
}

func (v *NullableDiscoveryResult) Set(val *DiscoveryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryResult(val *DiscoveryResult) *NullableDiscoveryResult {
	return &NullableDiscoveryResult{value: val, isSet: true}
}

func (v NullableDiscoveryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


