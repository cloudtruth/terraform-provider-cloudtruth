/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ctapi

import (
	"encoding/json"
)

// MembershipCreate struct for MembershipCreate
type MembershipCreate struct {
	// The user of the membership.
	User string `json:"user"`
	// The role that the user has in the organization.
	Role RoleEnum `json:"role"`
}

// NewMembershipCreate instantiates a new MembershipCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipCreate(user string, role RoleEnum) *MembershipCreate {
	this := MembershipCreate{}
	this.User = user
	this.Role = role
	return &this
}

// NewMembershipCreateWithDefaults instantiates a new MembershipCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipCreateWithDefaults() *MembershipCreate {
	this := MembershipCreate{}
	return &this
}

// GetUser returns the User field value
func (o *MembershipCreate) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *MembershipCreate) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *MembershipCreate) SetUser(v string) {
	o.User = v
}

// GetRole returns the Role field value
func (o *MembershipCreate) GetRole() RoleEnum {
	if o == nil {
		var ret RoleEnum
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *MembershipCreate) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *MembershipCreate) SetRole(v RoleEnum) {
	o.Role = v
}

func (o MembershipCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableMembershipCreate struct {
	value *MembershipCreate
	isSet bool
}

func (v NullableMembershipCreate) Get() *MembershipCreate {
	return v.value
}

func (v *NullableMembershipCreate) Set(val *MembershipCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipCreate(val *MembershipCreate) *NullableMembershipCreate {
	return &NullableMembershipCreate{value: val, isSet: true}
}

func (v NullableMembershipCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
