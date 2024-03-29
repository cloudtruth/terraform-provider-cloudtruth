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

// checks if the PatchedParameterTypeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedParameterTypeUpdate{}

// PatchedParameterTypeUpdate struct for PatchedParameterTypeUpdate
type PatchedParameterTypeUpdate struct {
	Id *string `json:"id,omitempty"`
	// The parameter type name.
	Name *string `json:"name,omitempty"`
	// A description of the parameter type, provide documentation on how to use this type here.
	Description *string `json:"description,omitempty"`
	// The URL for this parameter type's parent
	Parent *string `json:"parent,omitempty"`
	// Rules applied to this parameter.
	Rules []ParameterTypeRule `json:"rules,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt NullableTime `json:"modified_at,omitempty"`
}

// NewPatchedParameterTypeUpdate instantiates a new PatchedParameterTypeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedParameterTypeUpdate() *PatchedParameterTypeUpdate {
	this := PatchedParameterTypeUpdate{}
	return &this
}

// NewPatchedParameterTypeUpdateWithDefaults instantiates a new PatchedParameterTypeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedParameterTypeUpdateWithDefaults() *PatchedParameterTypeUpdate {
	this := PatchedParameterTypeUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedParameterTypeUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterTypeUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedParameterTypeUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedParameterTypeUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterTypeUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedParameterTypeUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedParameterTypeUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterTypeUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedParameterTypeUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *PatchedParameterTypeUpdate) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterTypeUpdate) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *PatchedParameterTypeUpdate) SetParent(v string) {
	o.Parent = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *PatchedParameterTypeUpdate) GetRules() []ParameterTypeRule {
	if o == nil || IsNil(o.Rules) {
		var ret []ParameterTypeRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterTypeUpdate) GetRulesOk() ([]ParameterTypeRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ParameterTypeRule and assigns it to the Rules field.
func (o *PatchedParameterTypeUpdate) SetRules(v []ParameterTypeRule) {
	o.Rules = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedParameterTypeUpdate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedParameterTypeUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedParameterTypeUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedParameterTypeUpdate) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedParameterTypeUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedParameterTypeUpdate) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *PatchedParameterTypeUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}
// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *PatchedParameterTypeUpdate) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *PatchedParameterTypeUpdate) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

func (o PatchedParameterTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedParameterTypeUpdate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return toSerialize, nil
}

type NullablePatchedParameterTypeUpdate struct {
	value *PatchedParameterTypeUpdate
	isSet bool
}

func (v NullablePatchedParameterTypeUpdate) Get() *PatchedParameterTypeUpdate {
	return v.value
}

func (v *NullablePatchedParameterTypeUpdate) Set(val *PatchedParameterTypeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedParameterTypeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedParameterTypeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedParameterTypeUpdate(val *PatchedParameterTypeUpdate) *NullablePatchedParameterTypeUpdate {
	return &NullablePatchedParameterTypeUpdate{value: val, isSet: true}
}

func (v NullablePatchedParameterTypeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedParameterTypeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


