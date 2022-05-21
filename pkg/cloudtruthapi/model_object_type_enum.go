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

// ObjectTypeEnum the model 'ObjectTypeEnum'
type ObjectTypeEnum string

// List of ObjectTypeEnum
const (
	OBJECTTYPEENUM_AWS_INTEGRATION     ObjectTypeEnum = "AwsIntegration"
	OBJECTTYPEENUM_ENVIRONMENT         ObjectTypeEnum = "Environment"
	OBJECTTYPEENUM_GIT_HUB_INTEGRATION ObjectTypeEnum = "GitHubIntegration"
	OBJECTTYPEENUM_GRANT               ObjectTypeEnum = "Grant"
	OBJECTTYPEENUM_INVITATION          ObjectTypeEnum = "Invitation"
	OBJECTTYPEENUM_MEMBERSHIP          ObjectTypeEnum = "Membership"
	OBJECTTYPEENUM_ORGANIZATION        ObjectTypeEnum = "Organization"
	OBJECTTYPEENUM_PARAMETER           ObjectTypeEnum = "Parameter"
	OBJECTTYPEENUM_PARAMETER_RULE      ObjectTypeEnum = "ParameterRule"
	OBJECTTYPEENUM_PARAMETER_TYPE      ObjectTypeEnum = "ParameterType"
	OBJECTTYPEENUM_PARAMETER_TYPE_RULE ObjectTypeEnum = "ParameterTypeRule"
	OBJECTTYPEENUM_PROJECT             ObjectTypeEnum = "Project"
	OBJECTTYPEENUM_PULL                ObjectTypeEnum = "Pull"
	OBJECTTYPEENUM_PUSH                ObjectTypeEnum = "Push"
	OBJECTTYPEENUM_SERVICE_ACCOUNT     ObjectTypeEnum = "ServiceAccount"
	OBJECTTYPEENUM_TAG                 ObjectTypeEnum = "Tag"
	OBJECTTYPEENUM_TASK                ObjectTypeEnum = "Task"
	OBJECTTYPEENUM_TEMPLATE            ObjectTypeEnum = "Template"
	OBJECTTYPEENUM_VALUE               ObjectTypeEnum = "Value"
)

// All allowed values of ObjectTypeEnum enum
var AllowedObjectTypeEnumEnumValues = []ObjectTypeEnum{
	"AwsIntegration",
	"Environment",
	"GitHubIntegration",
	"Grant",
	"Invitation",
	"Membership",
	"Organization",
	"Parameter",
	"ParameterRule",
	"ParameterType",
	"ParameterTypeRule",
	"Project",
	"Pull",
	"Push",
	"ServiceAccount",
	"Tag",
	"Task",
	"Template",
	"Value",
}

func (v *ObjectTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjectTypeEnum(value)
	for _, existing := range AllowedObjectTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjectTypeEnum", value)
}

// NewObjectTypeEnumFromValue returns a pointer to a valid ObjectTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjectTypeEnumFromValue(v string) (*ObjectTypeEnum, error) {
	ev := ObjectTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjectTypeEnum: valid values are %v", v, AllowedObjectTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjectTypeEnum) IsValid() bool {
	for _, existing := range AllowedObjectTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectTypeEnum value
func (v ObjectTypeEnum) Ptr() *ObjectTypeEnum {
	return &v
}

type NullableObjectTypeEnum struct {
	value *ObjectTypeEnum
	isSet bool
}

func (v NullableObjectTypeEnum) Get() *ObjectTypeEnum {
	return v.value
}

func (v *NullableObjectTypeEnum) Set(val *ObjectTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectTypeEnum(val *ObjectTypeEnum) *NullableObjectTypeEnum {
	return &NullableObjectTypeEnum{value: val, isSet: true}
}

func (v NullableObjectTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
