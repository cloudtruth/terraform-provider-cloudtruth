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

// checks if the Tag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tag{}

// Tag The details of a tag.
type Tag struct {
	Url string `json:"url"`
	// A unique identifier for the tag.
	Id string `json:"id"`
	// The tag name. Tag names may contain alphanumeric, hyphen, underscore, or period characters. Tag names are case sensitive. The name cannot be modified.
	Name string `json:"name"`
	// A description of the tag.  You may find it helpful to document how this tag is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// The point in time this tag represents.
	Timestamp time.Time `json:"timestamp"`
	// Deprecated. Only shows pushes for aws integrations in /api/v1/.
	Pushes []AwsPush `json:"pushes"`
	// Push actions associated with the tag.
	PushUrls []string `json:"push_urls"`
	Usage TagReadUsage `json:"usage"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(url string, id string, name string, timestamp time.Time, pushes []AwsPush, pushUrls []string, usage TagReadUsage) *Tag {
	this := Tag{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Timestamp = timestamp
	this.Pushes = pushes
	this.PushUrls = pushUrls
	this.Usage = usage
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetUrl returns the Url field value
func (o *Tag) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Tag) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Tag) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *Tag) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tag) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tag) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tag) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tag) SetDescription(v string) {
	o.Description = &v
}

// GetTimestamp returns the Timestamp field value
func (o *Tag) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Tag) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Tag) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetPushes returns the Pushes field value
func (o *Tag) GetPushes() []AwsPush {
	if o == nil {
		var ret []AwsPush
		return ret
	}

	return o.Pushes
}

// GetPushesOk returns a tuple with the Pushes field value
// and a boolean to check if the value has been set.
func (o *Tag) GetPushesOk() ([]AwsPush, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pushes, true
}

// SetPushes sets field value
func (o *Tag) SetPushes(v []AwsPush) {
	o.Pushes = v
}

// GetPushUrls returns the PushUrls field value
func (o *Tag) GetPushUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PushUrls
}

// GetPushUrlsOk returns a tuple with the PushUrls field value
// and a boolean to check if the value has been set.
func (o *Tag) GetPushUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushUrls, true
}

// SetPushUrls sets field value
func (o *Tag) SetPushUrls(v []string) {
	o.PushUrls = v
}

// GetUsage returns the Usage field value
func (o *Tag) GetUsage() TagReadUsage {
	if o == nil {
		var ret TagReadUsage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *Tag) GetUsageOk() (*TagReadUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *Tag) SetUsage(v TagReadUsage) {
	o.Usage = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["pushes"] = o.Pushes
	toSerialize["push_urls"] = o.PushUrls
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


