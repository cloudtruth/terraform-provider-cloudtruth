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

// TemplateTimelineEntryTemplate Helper methods for all views or serializers that expose template concepts.
type TemplateTimelineEntryTemplate struct {
	// A unique identifier for the parameter.
	Id string `json:"id"`
	// The parameter name.
	Name string `json:"name"`
	// A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// The content of the template.  Use mustache-style templating delimiters of `{{` and `}}` to reference parameter values by name for substitution into the template result.
	Body *string `json:"body,omitempty"`
}

// NewTemplateTimelineEntryTemplate instantiates a new TemplateTimelineEntryTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateTimelineEntryTemplate(id string, name string) *TemplateTimelineEntryTemplate {
	this := TemplateTimelineEntryTemplate{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTemplateTimelineEntryTemplateWithDefaults instantiates a new TemplateTimelineEntryTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateTimelineEntryTemplateWithDefaults() *TemplateTimelineEntryTemplate {
	this := TemplateTimelineEntryTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateTimelineEntryTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateTimelineEntryTemplate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateTimelineEntryTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateTimelineEntryTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateTimelineEntryTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateTimelineEntryTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateTimelineEntryTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *TemplateTimelineEntryTemplate) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryTemplate) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *TemplateTimelineEntryTemplate) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *TemplateTimelineEntryTemplate) SetBody(v string) {
	o.Body = &v
}

func (o TemplateTimelineEntryTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateTimelineEntryTemplate struct {
	value *TemplateTimelineEntryTemplate
	isSet bool
}

func (v NullableTemplateTimelineEntryTemplate) Get() *TemplateTimelineEntryTemplate {
	return v.value
}

func (v *NullableTemplateTimelineEntryTemplate) Set(val *TemplateTimelineEntryTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateTimelineEntryTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateTimelineEntryTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateTimelineEntryTemplate(val *TemplateTimelineEntryTemplate) *NullableTemplateTimelineEntryTemplate {
	return &NullableTemplateTimelineEntryTemplate{value: val, isSet: true}
}

func (v NullableTemplateTimelineEntryTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateTimelineEntryTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
