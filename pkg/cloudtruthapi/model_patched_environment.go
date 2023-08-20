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

// checks if the PatchedEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedEnvironment{}

// PatchedEnvironment struct for PatchedEnvironment
type PatchedEnvironment struct {
	Url *string `json:"url,omitempty"`
	// A unique identifier for the environment.
	Id *string `json:"id,omitempty"`
	// The environment name.
	Name *string `json:"name,omitempty"`
	// A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed.
	Parent NullableString `json:"parent,omitempty"`
	// This is the opposite of `parent`, see that field for more details.
	Children []string `json:"children,omitempty"`
	// Indicates if access control is being enforced through grants.
	AccessControlled *bool `json:"access_controlled,omitempty"`
	Role NullableRoleEnum `json:"role,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

// NewPatchedEnvironment instantiates a new PatchedEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEnvironment() *PatchedEnvironment {
	this := PatchedEnvironment{}
	return &this
}

// NewPatchedEnvironmentWithDefaults instantiates a new PatchedEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEnvironmentWithDefaults() *PatchedEnvironment {
	this := PatchedEnvironment{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedEnvironment) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedEnvironment) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEnvironment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedEnvironment) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEnvironment) GetParent() string {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEnvironment) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *PatchedEnvironment) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedEnvironment) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedEnvironment) UnsetParent() {
	o.Parent.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetChildren() []string {
	if o == nil || IsNil(o.Children) {
		var ret []string
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetChildrenOk() ([]string, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []string and assigns it to the Children field.
func (o *PatchedEnvironment) SetChildren(v []string) {
	o.Children = v
}

// GetAccessControlled returns the AccessControlled field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetAccessControlled() bool {
	if o == nil || IsNil(o.AccessControlled) {
		var ret bool
		return ret
	}
	return *o.AccessControlled
}

// GetAccessControlledOk returns a tuple with the AccessControlled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetAccessControlledOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessControlled) {
		return nil, false
	}
	return o.AccessControlled, true
}

// HasAccessControlled returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasAccessControlled() bool {
	if o != nil && !IsNil(o.AccessControlled) {
		return true
	}

	return false
}

// SetAccessControlled gets a reference to the given bool and assigns it to the AccessControlled field.
func (o *PatchedEnvironment) SetAccessControlled(v bool) {
	o.AccessControlled = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEnvironment) GetRole() RoleEnum {
	if o == nil || IsNil(o.Role.Get()) {
		var ret RoleEnum
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEnvironment) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableRoleEnum and assigns it to the Role field.
func (o *PatchedEnvironment) SetRole(v RoleEnum) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedEnvironment) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedEnvironment) UnsetRole() {
	o.Role.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedEnvironment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedEnvironment) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEnvironment) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedEnvironment) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedEnvironment) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o PatchedEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
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
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return toSerialize, nil
}

type NullablePatchedEnvironment struct {
	value *PatchedEnvironment
	isSet bool
}

func (v NullablePatchedEnvironment) Get() *PatchedEnvironment {
	return v.value
}

func (v *NullablePatchedEnvironment) Set(val *PatchedEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEnvironment(val *PatchedEnvironment) *NullablePatchedEnvironment {
	return &NullablePatchedEnvironment{value: val, isSet: true}
}

func (v NullablePatchedEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


