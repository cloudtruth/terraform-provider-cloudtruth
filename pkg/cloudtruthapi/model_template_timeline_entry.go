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

// checks if the TemplateTimelineEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateTimelineEntry{}

// TemplateTimelineEntry Details about a single change.
type TemplateTimelineEntry struct {
	HistoryType HistoryTypeEnum `json:"history_type"`
	ModifiedAt NullableTime `json:"modified_at"`
	ModifiedBy *string `json:"modified_by,omitempty"`
	HistoryTemplate TemplateTimelineEntryHistoryTemplate `json:"history_template"`
}

// NewTemplateTimelineEntry instantiates a new TemplateTimelineEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateTimelineEntry(historyType HistoryTypeEnum, modifiedAt NullableTime, historyTemplate TemplateTimelineEntryHistoryTemplate) *TemplateTimelineEntry {
	this := TemplateTimelineEntry{}
	this.HistoryType = historyType
	this.ModifiedAt = modifiedAt
	this.HistoryTemplate = historyTemplate
	return &this
}

// NewTemplateTimelineEntryWithDefaults instantiates a new TemplateTimelineEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateTimelineEntryWithDefaults() *TemplateTimelineEntry {
	this := TemplateTimelineEntry{}
	return &this
}

// GetHistoryType returns the HistoryType field value
func (o *TemplateTimelineEntry) GetHistoryType() HistoryTypeEnum {
	if o == nil {
		var ret HistoryTypeEnum
		return ret
	}

	return o.HistoryType
}

// GetHistoryTypeOk returns a tuple with the HistoryType field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntry) GetHistoryTypeOk() (*HistoryTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoryType, true
}

// SetHistoryType sets field value
func (o *TemplateTimelineEntry) SetHistoryType(v HistoryTypeEnum) {
	o.HistoryType = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TemplateTimelineEntry) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateTimelineEntry) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *TemplateTimelineEntry) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *TemplateTimelineEntry) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntry) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *TemplateTimelineEntry) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *TemplateTimelineEntry) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetHistoryTemplate returns the HistoryTemplate field value
func (o *TemplateTimelineEntry) GetHistoryTemplate() TemplateTimelineEntryHistoryTemplate {
	if o == nil {
		var ret TemplateTimelineEntryHistoryTemplate
		return ret
	}

	return o.HistoryTemplate
}

// GetHistoryTemplateOk returns a tuple with the HistoryTemplate field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntry) GetHistoryTemplateOk() (*TemplateTimelineEntryHistoryTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoryTemplate, true
}

// SetHistoryTemplate sets field value
func (o *TemplateTimelineEntry) SetHistoryTemplate(v TemplateTimelineEntryHistoryTemplate) {
	o.HistoryTemplate = v
}

func (o TemplateTimelineEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateTimelineEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["history_type"] = o.HistoryType
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	if !IsNil(o.ModifiedBy) {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	toSerialize["history_template"] = o.HistoryTemplate
	return toSerialize, nil
}

type NullableTemplateTimelineEntry struct {
	value *TemplateTimelineEntry
	isSet bool
}

func (v NullableTemplateTimelineEntry) Get() *TemplateTimelineEntry {
	return v.value
}

func (v *NullableTemplateTimelineEntry) Set(val *TemplateTimelineEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateTimelineEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateTimelineEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateTimelineEntry(val *TemplateTimelineEntry) *NullableTemplateTimelineEntry {
	return &NullableTemplateTimelineEntry{value: val, isSet: true}
}

func (v NullableTemplateTimelineEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateTimelineEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


