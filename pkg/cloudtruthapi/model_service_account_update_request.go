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

// checks if the ServiceAccountUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountUpdateRequest{}

// ServiceAccountUpdateRequest struct for ServiceAccountUpdateRequest
type ServiceAccountUpdateRequest struct {
	// An optional description of the process or system using the service account.
	Description *string `json:"description,omitempty"`
	// The role for the service acount
	Role string `json:"role"`
	// The owner of the service account.
	Owner *string `json:"owner,omitempty"`
}

// NewServiceAccountUpdateRequest instantiates a new ServiceAccountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountUpdateRequest(role string) *ServiceAccountUpdateRequest {
	this := ServiceAccountUpdateRequest{}
	this.Role = role
	return &this
}

// NewServiceAccountUpdateRequestWithDefaults instantiates a new ServiceAccountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountUpdateRequestWithDefaults() *ServiceAccountUpdateRequest {
	this := ServiceAccountUpdateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRole returns the Role field value
func (o *ServiceAccountUpdateRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ServiceAccountUpdateRequest) SetRole(v string) {
	o.Role = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ServiceAccountUpdateRequest) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServiceAccountUpdateRequest) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ServiceAccountUpdateRequest) SetOwner(v string) {
	o.Owner = &v
}

func (o ServiceAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["role"] = o.Role
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableServiceAccountUpdateRequest struct {
	value *ServiceAccountUpdateRequest
	isSet bool
}

func (v NullableServiceAccountUpdateRequest) Get() *ServiceAccountUpdateRequest {
	return v.value
}

func (v *NullableServiceAccountUpdateRequest) Set(val *ServiceAccountUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountUpdateRequest(val *ServiceAccountUpdateRequest) *NullableServiceAccountUpdateRequest {
	return &NullableServiceAccountUpdateRequest{value: val, isSet: true}
}

func (v NullableServiceAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

