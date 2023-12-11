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

// checks if the PatchedProjectUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedProjectUpdate{}

// PatchedProjectUpdate struct for PatchedProjectUpdate
type PatchedProjectUpdate struct {
	Id *string `json:"id,omitempty"`
	// The project name.
	Name *string `json:"name,omitempty"`
	// A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// A regular expression parameter names must match
	ParameterNamePattern *string `json:"parameter_name_pattern,omitempty"`
	// Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project.
	DependsOn NullableString `json:"depends_on,omitempty"`
	// Indicates if access control is being enforced through grants.
	AccessControlled *bool `json:"access_controlled,omitempty"`
	Role NullableRoleEnum `json:"role,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt NullableTime `json:"modified_at,omitempty"`
}

// NewPatchedProjectUpdate instantiates a new PatchedProjectUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedProjectUpdate() *PatchedProjectUpdate {
	this := PatchedProjectUpdate{}
	return &this
}

// NewPatchedProjectUpdateWithDefaults instantiates a new PatchedProjectUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedProjectUpdateWithDefaults() *PatchedProjectUpdate {
	this := PatchedProjectUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedProjectUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProjectUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedProjectUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedProjectUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProjectUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedProjectUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedProjectUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProjectUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedProjectUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetParameterNamePattern returns the ParameterNamePattern field value if set, zero value otherwise.
func (o *PatchedProjectUpdate) GetParameterNamePattern() string {
	if o == nil || IsNil(o.ParameterNamePattern) {
		var ret string
		return ret
	}
	return *o.ParameterNamePattern
}

// GetParameterNamePatternOk returns a tuple with the ParameterNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProjectUpdate) GetParameterNamePatternOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterNamePattern) {
		return nil, false
	}
	return o.ParameterNamePattern, true
}

// HasParameterNamePattern returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasParameterNamePattern() bool {
	if o != nil && !IsNil(o.ParameterNamePattern) {
		return true
	}

	return false
}

// SetParameterNamePattern gets a reference to the given string and assigns it to the ParameterNamePattern field.
func (o *PatchedProjectUpdate) SetParameterNamePattern(v string) {
	o.ParameterNamePattern = &v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProjectUpdate) GetDependsOn() string {
	if o == nil || IsNil(o.DependsOn.Get()) {
		var ret string
		return ret
	}
	return *o.DependsOn.Get()
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProjectUpdate) GetDependsOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn.Get(), o.DependsOn.IsSet()
}

// HasDependsOn returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasDependsOn() bool {
	if o != nil && o.DependsOn.IsSet() {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given NullableString and assigns it to the DependsOn field.
func (o *PatchedProjectUpdate) SetDependsOn(v string) {
	o.DependsOn.Set(&v)
}
// SetDependsOnNil sets the value for DependsOn to be an explicit nil
func (o *PatchedProjectUpdate) SetDependsOnNil() {
	o.DependsOn.Set(nil)
}

// UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
func (o *PatchedProjectUpdate) UnsetDependsOn() {
	o.DependsOn.Unset()
}

// GetAccessControlled returns the AccessControlled field value if set, zero value otherwise.
func (o *PatchedProjectUpdate) GetAccessControlled() bool {
	if o == nil || IsNil(o.AccessControlled) {
		var ret bool
		return ret
	}
	return *o.AccessControlled
}

// GetAccessControlledOk returns a tuple with the AccessControlled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProjectUpdate) GetAccessControlledOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessControlled) {
		return nil, false
	}
	return o.AccessControlled, true
}

// HasAccessControlled returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasAccessControlled() bool {
	if o != nil && !IsNil(o.AccessControlled) {
		return true
	}

	return false
}

// SetAccessControlled gets a reference to the given bool and assigns it to the AccessControlled field.
func (o *PatchedProjectUpdate) SetAccessControlled(v bool) {
	o.AccessControlled = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProjectUpdate) GetRole() RoleEnum {
	if o == nil || IsNil(o.Role.Get()) {
		var ret RoleEnum
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProjectUpdate) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableRoleEnum and assigns it to the Role field.
func (o *PatchedProjectUpdate) SetRole(v RoleEnum) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedProjectUpdate) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedProjectUpdate) UnsetRole() {
	o.Role.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedProjectUpdate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProjectUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedProjectUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProjectUpdate) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProjectUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedProjectUpdate) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *PatchedProjectUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}
// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *PatchedProjectUpdate) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *PatchedProjectUpdate) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

func (o PatchedProjectUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedProjectUpdate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ParameterNamePattern) {
		toSerialize["parameter_name_pattern"] = o.ParameterNamePattern
	}
	if o.DependsOn.IsSet() {
		toSerialize["depends_on"] = o.DependsOn.Get()
	}
	if !IsNil(o.AccessControlled) {
		toSerialize["access_controlled"] = o.AccessControlled
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return toSerialize, nil
}

type NullablePatchedProjectUpdate struct {
	value *PatchedProjectUpdate
	isSet bool
}

func (v NullablePatchedProjectUpdate) Get() *PatchedProjectUpdate {
	return v.value
}

func (v *NullablePatchedProjectUpdate) Set(val *PatchedProjectUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedProjectUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedProjectUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedProjectUpdate(val *PatchedProjectUpdate) *NullablePatchedProjectUpdate {
	return &NullablePatchedProjectUpdate{value: val, isSet: true}
}

func (v NullablePatchedProjectUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedProjectUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

