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

// TemplatePreviewCreateRequest struct for TemplatePreviewCreateRequest
type TemplatePreviewCreateRequest struct {
	// The template body to instantiate on request, instantiated on response.
	Body string `json:"body"`
}

// NewTemplatePreviewCreateRequest instantiates a new TemplatePreviewCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatePreviewCreateRequest(body string) *TemplatePreviewCreateRequest {
	this := TemplatePreviewCreateRequest{}
	this.Body = body
	return &this
}

// NewTemplatePreviewCreateRequestWithDefaults instantiates a new TemplatePreviewCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePreviewCreateRequestWithDefaults() *TemplatePreviewCreateRequest {
	this := TemplatePreviewCreateRequest{}
	return &this
}

// GetBody returns the Body field value
func (o *TemplatePreviewCreateRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *TemplatePreviewCreateRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *TemplatePreviewCreateRequest) SetBody(v string) {
	o.Body = v
}

func (o TemplatePreviewCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableTemplatePreviewCreateRequest struct {
	value *TemplatePreviewCreateRequest
	isSet bool
}

func (v NullableTemplatePreviewCreateRequest) Get() *TemplatePreviewCreateRequest {
	return v.value
}

func (v *NullableTemplatePreviewCreateRequest) Set(val *TemplatePreviewCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatePreviewCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatePreviewCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatePreviewCreateRequest(val *TemplatePreviewCreateRequest) *NullableTemplatePreviewCreateRequest {
	return &NullableTemplatePreviewCreateRequest{value: val, isSet: true}
}

func (v NullableTemplatePreviewCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatePreviewCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
