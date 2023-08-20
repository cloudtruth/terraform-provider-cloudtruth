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

// checks if the OrganizationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCreate{}

// OrganizationCreate struct for OrganizationCreate
type OrganizationCreate struct {
	// The organization name.
	Name string `json:"name"`
}

// NewOrganizationCreate instantiates a new OrganizationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCreate(name string) *OrganizationCreate {
	this := OrganizationCreate{}
	this.Name = name
	return &this
}

// NewOrganizationCreateWithDefaults instantiates a new OrganizationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCreateWithDefaults() *OrganizationCreate {
	this := OrganizationCreate{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationCreate) SetName(v string) {
	o.Name = v
}

func (o OrganizationCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableOrganizationCreate struct {
	value *OrganizationCreate
	isSet bool
}

func (v NullableOrganizationCreate) Get() *OrganizationCreate {
	return v.value
}

func (v *NullableOrganizationCreate) Set(val *OrganizationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCreate(val *OrganizationCreate) *NullableOrganizationCreate {
	return &NullableOrganizationCreate{value: val, isSet: true}
}

func (v NullableOrganizationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


