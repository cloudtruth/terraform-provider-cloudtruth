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

// PatchedAwsPushUpdate Update a push.  The `region` and `service` cannot be changed on an existing push.
type PatchedAwsPushUpdate struct {
	// The action name.
	Name *string `json:"name,omitempty"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	// Projects that are included in the push.
	Projects *[]string `json:"projects,omitempty"`
	// Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push.
	Tags *[]string `json:"tags,omitempty"`
	// Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - ``{{ environment }}`` to insert the environment name   - ``{{ parameter }}`` to insert the parameter name   - ``{{ project }}`` to insert the project name   - ``{{ push }}`` to insert the push name   - ``{{ tag }}`` to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the `project` substitution is required.  If you include multiple tags from different environments in the push, the `environment` substitution is required.  If you include multiple tags from the same environment in the push, the `tag` substitution is required.  In all cases, the `parameter` substitution is always required.
	Resource *string `json:"resource,omitempty"`
}

// NewPatchedAwsPushUpdate instantiates a new PatchedAwsPushUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAwsPushUpdate() *PatchedAwsPushUpdate {
	this := PatchedAwsPushUpdate{}
	return &this
}

// NewPatchedAwsPushUpdateWithDefaults instantiates a new PatchedAwsPushUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAwsPushUpdateWithDefaults() *PatchedAwsPushUpdate {
	this := PatchedAwsPushUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAwsPushUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedAwsPushUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetProjects() []string {
	if o == nil || o.Projects == nil {
		var ret []string
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetProjectsOk() (*[]string, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *PatchedAwsPushUpdate) SetProjects(v []string) {
	o.Projects = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PatchedAwsPushUpdate) SetTags(v []string) {
	o.Tags = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *PatchedAwsPushUpdate) SetResource(v string) {
	o.Resource = &v
}

func (o PatchedAwsPushUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAwsPushUpdate struct {
	value *PatchedAwsPushUpdate
	isSet bool
}

func (v NullablePatchedAwsPushUpdate) Get() *PatchedAwsPushUpdate {
	return v.value
}

func (v *NullablePatchedAwsPushUpdate) Set(val *PatchedAwsPushUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAwsPushUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAwsPushUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAwsPushUpdate(val *PatchedAwsPushUpdate) *NullablePatchedAwsPushUpdate {
	return &NullablePatchedAwsPushUpdate{value: val, isSet: true}
}

func (v NullablePatchedAwsPushUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAwsPushUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
