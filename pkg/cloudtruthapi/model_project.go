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

// Project struct for Project
type Project struct {
	Url string `json:"url"`
	// A unique identifier for the project.
	Id string `json:"id"`
	// The project name.
	Name string `json:"name"`
	// A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content.
	Description *string `json:"description,omitempty"`
	// This is the opposite of `depends_on`, see that field for more details.
	Dependents []string `json:"dependents"`
	// Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project.
	DependsOn NullableString `json:"depends_on,omitempty"`
	// Indicates if access control is being enforced through grants.
	AccessControlled *bool `json:"access_controlled,omitempty"`
	// Your role in the project, if the project is access-controlled.
	Role NullableRoleEnum `json:"role"`
	// Deprecated. Only shows pushes for aws integrations in /api/v1/.
	Pushes []AwsPush `json:"pushes"`
	// Push actions associated with the project.
	PushUrls   []string  `json:"push_urls"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(url string, id string, name string, dependents []string, role NullableRoleEnum, pushes []AwsPush, pushUrls []string, createdAt time.Time, modifiedAt time.Time) *Project {
	this := Project{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Dependents = dependents
	this.Role = role
	this.Pushes = pushes
	this.PushUrls = pushUrls
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetUrl returns the Url field value
func (o *Project) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Project) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Project) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *Project) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Project) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Project) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Project) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Project) SetDescription(v string) {
	o.Description = &v
}

// GetDependents returns the Dependents field value
func (o *Project) GetDependents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dependents
}

// GetDependentsOk returns a tuple with the Dependents field value
// and a boolean to check if the value has been set.
func (o *Project) GetDependentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dependents, true
}

// SetDependents sets field value
func (o *Project) SetDependents(v []string) {
	o.Dependents = v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetDependsOn() string {
	if o == nil || o.DependsOn.Get() == nil {
		var ret string
		return ret
	}
	return *o.DependsOn.Get()
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetDependsOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn.Get(), o.DependsOn.IsSet()
}

// HasDependsOn returns a boolean if a field has been set.
func (o *Project) HasDependsOn() bool {
	if o != nil && o.DependsOn.IsSet() {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given NullableString and assigns it to the DependsOn field.
func (o *Project) SetDependsOn(v string) {
	o.DependsOn.Set(&v)
}

// SetDependsOnNil sets the value for DependsOn to be an explicit nil
func (o *Project) SetDependsOnNil() {
	o.DependsOn.Set(nil)
}

// UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
func (o *Project) UnsetDependsOn() {
	o.DependsOn.Unset()
}

// GetAccessControlled returns the AccessControlled field value if set, zero value otherwise.
func (o *Project) GetAccessControlled() bool {
	if o == nil || o.AccessControlled == nil {
		var ret bool
		return ret
	}
	return *o.AccessControlled
}

// GetAccessControlledOk returns a tuple with the AccessControlled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetAccessControlledOk() (*bool, bool) {
	if o == nil || o.AccessControlled == nil {
		return nil, false
	}
	return o.AccessControlled, true
}

// HasAccessControlled returns a boolean if a field has been set.
func (o *Project) HasAccessControlled() bool {
	if o != nil && o.AccessControlled != nil {
		return true
	}

	return false
}

// SetAccessControlled gets a reference to the given bool and assigns it to the AccessControlled field.
func (o *Project) SetAccessControlled(v bool) {
	o.AccessControlled = &v
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for RoleEnum will be returned
func (o *Project) GetRole() RoleEnum {
	if o == nil || o.Role.Get() == nil {
		var ret RoleEnum
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *Project) SetRole(v RoleEnum) {
	o.Role.Set(&v)
}

// GetPushes returns the Pushes field value
func (o *Project) GetPushes() []AwsPush {
	if o == nil {
		var ret []AwsPush
		return ret
	}

	return o.Pushes
}

// GetPushesOk returns a tuple with the Pushes field value
// and a boolean to check if the value has been set.
func (o *Project) GetPushesOk() ([]AwsPush, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pushes, true
}

// SetPushes sets field value
func (o *Project) SetPushes(v []AwsPush) {
	o.Pushes = v
}

// GetPushUrls returns the PushUrls field value
func (o *Project) GetPushUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PushUrls
}

// GetPushUrlsOk returns a tuple with the PushUrls field value
// and a boolean to check if the value has been set.
func (o *Project) GetPushUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushUrls, true
}

// SetPushUrls sets field value
func (o *Project) SetPushUrls(v []string) {
	o.PushUrls = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Project) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Project) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Project) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Project) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["dependents"] = o.Dependents
	}
	if o.DependsOn.IsSet() {
		toSerialize["depends_on"] = o.DependsOn.Get()
	}
	if o.AccessControlled != nil {
		toSerialize["access_controlled"] = o.AccessControlled
	}
	if true {
		toSerialize["role"] = o.Role.Get()
	}
	if true {
		toSerialize["pushes"] = o.Pushes
	}
	if true {
		toSerialize["push_urls"] = o.PushUrls
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
