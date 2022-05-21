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

// HistoryModelEnum the model 'HistoryModelEnum'
type HistoryModelEnum string

// List of HistoryModelEnum
const (
	HISTORYMODELENUM_PARAMETER      HistoryModelEnum = "Parameter"
	HISTORYMODELENUM_PARAMETER_RULE HistoryModelEnum = "ParameterRule"
	HISTORYMODELENUM_VALUE          HistoryModelEnum = "Value"
)

// All allowed values of HistoryModelEnum enum
var AllowedHistoryModelEnumEnumValues = []HistoryModelEnum{
	"Parameter",
	"ParameterRule",
	"Value",
}

func (v *HistoryModelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HistoryModelEnum(value)
	for _, existing := range AllowedHistoryModelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HistoryModelEnum", value)
}

// NewHistoryModelEnumFromValue returns a pointer to a valid HistoryModelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHistoryModelEnumFromValue(v string) (*HistoryModelEnum, error) {
	ev := HistoryModelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HistoryModelEnum: valid values are %v", v, AllowedHistoryModelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HistoryModelEnum) IsValid() bool {
	for _, existing := range AllowedHistoryModelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistoryModelEnum value
func (v HistoryModelEnum) Ptr() *HistoryModelEnum {
	return &v
}

type NullableHistoryModelEnum struct {
	value *HistoryModelEnum
	isSet bool
}

func (v NullableHistoryModelEnum) Get() *HistoryModelEnum {
	return v.value
}

func (v *NullableHistoryModelEnum) Set(val *HistoryModelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryModelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryModelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryModelEnum(val *HistoryModelEnum) *NullableHistoryModelEnum {
	return &NullableHistoryModelEnum{value: val, isSet: true}
}

func (v NullableHistoryModelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryModelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
