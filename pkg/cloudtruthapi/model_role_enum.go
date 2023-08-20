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
	"fmt"
)

// RoleEnum the model 'RoleEnum'
type RoleEnum string

// List of RoleEnum
const (
	ROLEENUM_OWNER RoleEnum = "OWNER"
	ROLEENUM_ADMIN RoleEnum = "ADMIN"
	ROLEENUM_CONTRIB RoleEnum = "CONTRIB"
	ROLEENUM_VIEWER RoleEnum = "VIEWER"
)

// All allowed values of RoleEnum enum
var AllowedRoleEnumEnumValues = []RoleEnum{
	"OWNER",
	"ADMIN",
	"CONTRIB",
	"VIEWER",
}

func (v *RoleEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleEnum(value)
	for _, existing := range AllowedRoleEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleEnum", value)
}

// NewRoleEnumFromValue returns a pointer to a valid RoleEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleEnumFromValue(v string) (*RoleEnum, error) {
	ev := RoleEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoleEnum: valid values are %v", v, AllowedRoleEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleEnum) IsValid() bool {
	for _, existing := range AllowedRoleEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleEnum value
func (v RoleEnum) Ptr() *RoleEnum {
	return &v
}

type NullableRoleEnum struct {
	value *RoleEnum
	isSet bool
}

func (v NullableRoleEnum) Get() *RoleEnum {
	return v.value
}

func (v *NullableRoleEnum) Set(val *RoleEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleEnum(val *RoleEnum) *NullableRoleEnum {
	return &NullableRoleEnum{value: val, isSet: true}
}

func (v NullableRoleEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

