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

// checks if the TemplatePreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplatePreview{}

// TemplatePreview struct for TemplatePreview
type TemplatePreview struct {
	// The template body to instantiate on request, instantiated on response.
	Body string `json:"body"`
	// If True, the instantiated template contains secrets.
	HasSecret bool `json:"has_secret"`
}

// NewTemplatePreview instantiates a new TemplatePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatePreview(body string, hasSecret bool) *TemplatePreview {
	this := TemplatePreview{}
	this.Body = body
	this.HasSecret = hasSecret
	return &this
}

// NewTemplatePreviewWithDefaults instantiates a new TemplatePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePreviewWithDefaults() *TemplatePreview {
	this := TemplatePreview{}
	return &this
}

// GetBody returns the Body field value
func (o *TemplatePreview) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *TemplatePreview) SetBody(v string) {
	o.Body = v
}

// GetHasSecret returns the HasSecret field value
func (o *TemplatePreview) GetHasSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasSecret
}

// GetHasSecretOk returns a tuple with the HasSecret field value
// and a boolean to check if the value has been set.
func (o *TemplatePreview) GetHasSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasSecret, true
}

// SetHasSecret sets field value
func (o *TemplatePreview) SetHasSecret(v bool) {
	o.HasSecret = v
}

func (o TemplatePreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplatePreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["has_secret"] = o.HasSecret
	return toSerialize, nil
}

type NullableTemplatePreview struct {
	value *TemplatePreview
	isSet bool
}

func (v NullableTemplatePreview) Get() *TemplatePreview {
	return v.value
}

func (v *NullableTemplatePreview) Set(val *TemplatePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatePreview(val *TemplatePreview) *NullableTemplatePreview {
	return &NullableTemplatePreview{value: val, isSet: true}
}

func (v NullableTemplatePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


