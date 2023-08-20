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

// checks if the PatchedTagUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedTagUpdate{}

// PatchedTagUpdate Details for updating a tag.
type PatchedTagUpdate struct {
	// A unique identifier for the tag.
	Id *string `json:"id,omitempty"`
	// The tag name. Tag names may contain alphanumeric, hyphen, underscore, or period characters. Tag names are case sensitive. The name cannot be modified.
	Name *string `json:"name,omitempty"`
	// A description of the tag.  You may find it helpful to document how this tag is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// The point in time this tag represents.  If explicitly set to `null` then the current time will be used.
	Timestamp NullableTime `json:"timestamp,omitempty"`
}

// NewPatchedTagUpdate instantiates a new PatchedTagUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTagUpdate() *PatchedTagUpdate {
	this := PatchedTagUpdate{}
	return &this
}

// NewPatchedTagUpdateWithDefaults instantiates a new PatchedTagUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTagUpdateWithDefaults() *PatchedTagUpdate {
	this := PatchedTagUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedTagUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedTagUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedTagUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedTagUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedTagUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedTagUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedTagUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedTagUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedTagUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedTagUpdate) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedTagUpdate) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PatchedTagUpdate) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *PatchedTagUpdate) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *PatchedTagUpdate) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *PatchedTagUpdate) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o PatchedTagUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedTagUpdate) ToMap() (map[string]interface{}, error) {
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
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	return toSerialize, nil
}

type NullablePatchedTagUpdate struct {
	value *PatchedTagUpdate
	isSet bool
}

func (v NullablePatchedTagUpdate) Get() *PatchedTagUpdate {
	return v.value
}

func (v *NullablePatchedTagUpdate) Set(val *PatchedTagUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTagUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTagUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTagUpdate(val *PatchedTagUpdate) *NullablePatchedTagUpdate {
	return &NullablePatchedTagUpdate{value: val, isSet: true}
}

func (v NullablePatchedTagUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTagUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


