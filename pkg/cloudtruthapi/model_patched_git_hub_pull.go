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

// PatchedGitHubPull Pull actions can be configured to get configuration and secrets from integrations on demand.
type PatchedGitHubPull struct {
	Url *string `json:"url,omitempty"`
	// Unique identifier for the action.
	Id *string `json:"id,omitempty"`
	// The action name.
	Name *string `json:"name,omitempty"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	LatestTask NullableGitHubPullLatestTask `json:"latest_task,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Allow the pull to create environments.  Any automatically created environments will be children of the `default` environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it.
	CreateEnvironments *bool `json:"create_environments,omitempty"`
	// Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it.
	CreateProjects *bool `json:"create_projects,omitempty"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// Values being managed by a mapped pull.
	MappedValues []Value `json:"mapped_values,omitempty"`
	// The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations.
	Mode NullableModeEnum `json:"mode,omitempty"`
}

// NewPatchedGitHubPull instantiates a new PatchedGitHubPull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedGitHubPull() *PatchedGitHubPull {
	this := PatchedGitHubPull{}
	return &this
}

// NewPatchedGitHubPullWithDefaults instantiates a new PatchedGitHubPull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedGitHubPullWithDefaults() *PatchedGitHubPull {
	this := PatchedGitHubPull{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedGitHubPull) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedGitHubPull) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedGitHubPull) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedGitHubPull) SetDescription(v string) {
	o.Description = &v
}

// GetLatestTask returns the LatestTask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedGitHubPull) GetLatestTask() GitHubPullLatestTask {
	if o == nil || o.LatestTask.Get() == nil {
		var ret GitHubPullLatestTask
		return ret
	}
	return *o.LatestTask.Get()
}

// GetLatestTaskOk returns a tuple with the LatestTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedGitHubPull) GetLatestTaskOk() (*GitHubPullLatestTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestTask.Get(), o.LatestTask.IsSet()
}

// HasLatestTask returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasLatestTask() bool {
	if o != nil && o.LatestTask.IsSet() {
		return true
	}

	return false
}

// SetLatestTask gets a reference to the given NullableGitHubPullLatestTask and assigns it to the LatestTask field.
func (o *PatchedGitHubPull) SetLatestTask(v GitHubPullLatestTask) {
	o.LatestTask.Set(&v)
}
// SetLatestTaskNil sets the value for LatestTask to be an explicit nil
func (o *PatchedGitHubPull) SetLatestTaskNil() {
	o.LatestTask.Set(nil)
}

// UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
func (o *PatchedGitHubPull) UnsetLatestTask() {
	o.LatestTask.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedGitHubPull) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedGitHubPull) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetCreateEnvironments returns the CreateEnvironments field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetCreateEnvironments() bool {
	if o == nil || o.CreateEnvironments == nil {
		var ret bool
		return ret
	}
	return *o.CreateEnvironments
}

// GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetCreateEnvironmentsOk() (*bool, bool) {
	if o == nil || o.CreateEnvironments == nil {
		return nil, false
	}
	return o.CreateEnvironments, true
}

// HasCreateEnvironments returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasCreateEnvironments() bool {
	if o != nil && o.CreateEnvironments != nil {
		return true
	}

	return false
}

// SetCreateEnvironments gets a reference to the given bool and assigns it to the CreateEnvironments field.
func (o *PatchedGitHubPull) SetCreateEnvironments(v bool) {
	o.CreateEnvironments = &v
}

// GetCreateProjects returns the CreateProjects field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetCreateProjects() bool {
	if o == nil || o.CreateProjects == nil {
		var ret bool
		return ret
	}
	return *o.CreateProjects
}

// GetCreateProjectsOk returns a tuple with the CreateProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetCreateProjectsOk() (*bool, bool) {
	if o == nil || o.CreateProjects == nil {
		return nil, false
	}
	return o.CreateProjects, true
}

// HasCreateProjects returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasCreateProjects() bool {
	if o != nil && o.CreateProjects != nil {
		return true
	}

	return false
}

// SetCreateProjects gets a reference to the given bool and assigns it to the CreateProjects field.
func (o *PatchedGitHubPull) SetCreateProjects(v bool) {
	o.CreateProjects = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *PatchedGitHubPull) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMappedValues returns the MappedValues field value if set, zero value otherwise.
func (o *PatchedGitHubPull) GetMappedValues() []Value {
	if o == nil || o.MappedValues == nil {
		var ret []Value
		return ret
	}
	return o.MappedValues
}

// GetMappedValuesOk returns a tuple with the MappedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGitHubPull) GetMappedValuesOk() ([]Value, bool) {
	if o == nil || o.MappedValues == nil {
		return nil, false
	}
	return o.MappedValues, true
}

// HasMappedValues returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasMappedValues() bool {
	if o != nil && o.MappedValues != nil {
		return true
	}

	return false
}

// SetMappedValues gets a reference to the given []Value and assigns it to the MappedValues field.
func (o *PatchedGitHubPull) SetMappedValues(v []Value) {
	o.MappedValues = v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedGitHubPull) GetMode() ModeEnum {
	if o == nil || o.Mode.Get() == nil {
		var ret ModeEnum
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedGitHubPull) GetModeOk() (*ModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedGitHubPull) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableModeEnum and assigns it to the Mode field.
func (o *PatchedGitHubPull) SetMode(v ModeEnum) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *PatchedGitHubPull) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *PatchedGitHubPull) UnsetMode() {
	o.Mode.Unset()
}

func (o PatchedGitHubPull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.LatestTask.IsSet() {
		toSerialize["latest_task"] = o.LatestTask.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.CreateEnvironments != nil {
		toSerialize["create_environments"] = o.CreateEnvironments
	}
	if o.CreateProjects != nil {
		toSerialize["create_projects"] = o.CreateProjects
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.MappedValues != nil {
		toSerialize["mapped_values"] = o.MappedValues
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedGitHubPull struct {
	value *PatchedGitHubPull
	isSet bool
}

func (v NullablePatchedGitHubPull) Get() *PatchedGitHubPull {
	return v.value
}

func (v *NullablePatchedGitHubPull) Set(val *PatchedGitHubPull) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedGitHubPull) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedGitHubPull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedGitHubPull(val *PatchedGitHubPull) *NullablePatchedGitHubPull {
	return &NullablePatchedGitHubPull{value: val, isSet: true}
}

func (v NullablePatchedGitHubPull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedGitHubPull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


