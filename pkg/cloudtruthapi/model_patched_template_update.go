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
	"time"
)

// checks if the PatchedTemplateUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedTemplateUpdate{}

// PatchedTemplateUpdate A parameter template in a given project, optionally instantiated against an environment.
type PatchedTemplateUpdate struct {
	Id *string `json:"id,omitempty"`
	// The template name.
	Name *string `json:"name,omitempty"`
	// ('A description of the template.  You may find it helpful to document how this template is used to assist others when they need to maintain software that uses this content.',)
	Description *string `json:"description,omitempty"`
	// If true, the `body` field has undergone evaluation.
	Evaluated *bool `json:"evaluated,omitempty"`
	// The content of the template.  Use mustache-style templating delimiters of `{{` and `}}` to reference parameter values by name for substitution into the template result.
	Body *string `json:"body,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt NullableTime `json:"modified_at,omitempty"`
}

// NewPatchedTemplateUpdate instantiates a new PatchedTemplateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTemplateUpdate() *PatchedTemplateUpdate {
	this := PatchedTemplateUpdate{}
	return &this
}

// NewPatchedTemplateUpdateWithDefaults instantiates a new PatchedTemplateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTemplateUpdateWithDefaults() *PatchedTemplateUpdate {
	this := PatchedTemplateUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedTemplateUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTemplateUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedTemplateUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedTemplateUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTemplateUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedTemplateUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedTemplateUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTemplateUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedTemplateUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetEvaluated returns the Evaluated field value if set, zero value otherwise.
func (o *PatchedTemplateUpdate) GetEvaluated() bool {
	if o == nil || IsNil(o.Evaluated) {
		var ret bool
		return ret
	}
	return *o.Evaluated
}

// GetEvaluatedOk returns a tuple with the Evaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTemplateUpdate) GetEvaluatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Evaluated) {
		return nil, false
	}
	return o.Evaluated, true
}

// HasEvaluated returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasEvaluated() bool {
	if o != nil && !IsNil(o.Evaluated) {
		return true
	}

	return false
}

// SetEvaluated gets a reference to the given bool and assigns it to the Evaluated field.
func (o *PatchedTemplateUpdate) SetEvaluated(v bool) {
	o.Evaluated = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PatchedTemplateUpdate) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTemplateUpdate) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *PatchedTemplateUpdate) SetBody(v string) {
	o.Body = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedTemplateUpdate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTemplateUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedTemplateUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedTemplateUpdate) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedTemplateUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedTemplateUpdate) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *PatchedTemplateUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}
// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *PatchedTemplateUpdate) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *PatchedTemplateUpdate) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

func (o PatchedTemplateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedTemplateUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Evaluated) {
		toSerialize["evaluated"] = o.Evaluated
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return toSerialize, nil
}

type NullablePatchedTemplateUpdate struct {
	value *PatchedTemplateUpdate
	isSet bool
}

func (v NullablePatchedTemplateUpdate) Get() *PatchedTemplateUpdate {
	return v.value
}

func (v *NullablePatchedTemplateUpdate) Set(val *PatchedTemplateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTemplateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTemplateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTemplateUpdate(val *PatchedTemplateUpdate) *NullablePatchedTemplateUpdate {
	return &NullablePatchedTemplateUpdate{value: val, isSet: true}
}

func (v NullablePatchedTemplateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTemplateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


