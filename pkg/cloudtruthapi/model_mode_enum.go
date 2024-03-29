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

// ModeEnum the model 'ModeEnum'
type ModeEnum string

// List of ModeEnum
const (
	MODEENUM_MAPPED ModeEnum = "mapped"
	MODEENUM_PATTERN ModeEnum = "pattern"
)

// All allowed values of ModeEnum enum
var AllowedModeEnumEnumValues = []ModeEnum{
	"mapped",
	"pattern",
}

func (v *ModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModeEnum(value)
	for _, existing := range AllowedModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModeEnum", value)
}

// NewModeEnumFromValue returns a pointer to a valid ModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModeEnumFromValue(v string) (*ModeEnum, error) {
	ev := ModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModeEnum: valid values are %v", v, AllowedModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModeEnum) IsValid() bool {
	for _, existing := range AllowedModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModeEnum value
func (v ModeEnum) Ptr() *ModeEnum {
	return &v
}

type NullableModeEnum struct {
	value *ModeEnum
	isSet bool
}

func (v NullableModeEnum) Get() *ModeEnum {
	return v.value
}

func (v *NullableModeEnum) Set(val *ModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModeEnum(val *ModeEnum) *NullableModeEnum {
	return &NullableModeEnum{value: val, isSet: true}
}

func (v NullableModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

