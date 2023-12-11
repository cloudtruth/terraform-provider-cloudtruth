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

// checks if the PatchedAwsPull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedAwsPull{}

// PatchedAwsPull Pull actions can be configured to get configuration and secrets from integrations on demand.
type PatchedAwsPull struct {
	Url *string `json:"url,omitempty"`
	// Unique identifier for the action.
	Id *string `json:"id,omitempty"`
	// The action name.
	Name *string `json:"name,omitempty"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	LatestTask NullableAwsPullLatestTask `json:"latest_task,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt NullableTime `json:"modified_at,omitempty"`
	// Allow the pull to create environments.  Any automatically created environments will be children of the `default` environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it.
	CreateEnvironments *bool `json:"create_environments,omitempty"`
	// Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it.
	CreateProjects *bool `json:"create_projects,omitempty"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// Values being managed by a mapped pull.
	MappedValues []ValueCreate `json:"mapped_values,omitempty"`
	Mode *ModeEnum `json:"mode,omitempty"`
	Region *AwsRegionEnum `json:"region,omitempty"`
	Service *AwsServiceEnum `json:"service,omitempty"`
	// Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - ``{{ environment }}`` to identify the environment name   - ``{{ parameter }}`` to identify the parameter name   - ``{{ project }}`` to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the `environment`, `project`, and `parameter`.
	Resource NullableString `json:"resource,omitempty"`
}

// NewPatchedAwsPull instantiates a new PatchedAwsPull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAwsPull() *PatchedAwsPull {
	this := PatchedAwsPull{}
	return &this
}

// NewPatchedAwsPullWithDefaults instantiates a new PatchedAwsPull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAwsPullWithDefaults() *PatchedAwsPull {
	this := PatchedAwsPull{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedAwsPull) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedAwsPull) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAwsPull) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedAwsPull) SetDescription(v string) {
	o.Description = &v
}

// GetLatestTask returns the LatestTask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsPull) GetLatestTask() AwsPullLatestTask {
	if o == nil || IsNil(o.LatestTask.Get()) {
		var ret AwsPullLatestTask
		return ret
	}
	return *o.LatestTask.Get()
}

// GetLatestTaskOk returns a tuple with the LatestTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsPull) GetLatestTaskOk() (*AwsPullLatestTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestTask.Get(), o.LatestTask.IsSet()
}

// HasLatestTask returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasLatestTask() bool {
	if o != nil && o.LatestTask.IsSet() {
		return true
	}

	return false
}

// SetLatestTask gets a reference to the given NullableAwsPullLatestTask and assigns it to the LatestTask field.
func (o *PatchedAwsPull) SetLatestTask(v AwsPullLatestTask) {
	o.LatestTask.Set(&v)
}
// SetLatestTaskNil sets the value for LatestTask to be an explicit nil
func (o *PatchedAwsPull) SetLatestTaskNil() {
	o.LatestTask.Set(nil)
}

// UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
func (o *PatchedAwsPull) UnsetLatestTask() {
	o.LatestTask.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedAwsPull) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsPull) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsPull) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *PatchedAwsPull) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}
// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *PatchedAwsPull) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *PatchedAwsPull) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetCreateEnvironments returns the CreateEnvironments field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetCreateEnvironments() bool {
	if o == nil || IsNil(o.CreateEnvironments) {
		var ret bool
		return ret
	}
	return *o.CreateEnvironments
}

// GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetCreateEnvironmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateEnvironments) {
		return nil, false
	}
	return o.CreateEnvironments, true
}

// HasCreateEnvironments returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasCreateEnvironments() bool {
	if o != nil && !IsNil(o.CreateEnvironments) {
		return true
	}

	return false
}

// SetCreateEnvironments gets a reference to the given bool and assigns it to the CreateEnvironments field.
func (o *PatchedAwsPull) SetCreateEnvironments(v bool) {
	o.CreateEnvironments = &v
}

// GetCreateProjects returns the CreateProjects field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetCreateProjects() bool {
	if o == nil || IsNil(o.CreateProjects) {
		var ret bool
		return ret
	}
	return *o.CreateProjects
}

// GetCreateProjectsOk returns a tuple with the CreateProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetCreateProjectsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateProjects) {
		return nil, false
	}
	return o.CreateProjects, true
}

// HasCreateProjects returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasCreateProjects() bool {
	if o != nil && !IsNil(o.CreateProjects) {
		return true
	}

	return false
}

// SetCreateProjects gets a reference to the given bool and assigns it to the CreateProjects field.
func (o *PatchedAwsPull) SetCreateProjects(v bool) {
	o.CreateProjects = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *PatchedAwsPull) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMappedValues returns the MappedValues field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetMappedValues() []ValueCreate {
	if o == nil || IsNil(o.MappedValues) {
		var ret []ValueCreate
		return ret
	}
	return o.MappedValues
}

// GetMappedValuesOk returns a tuple with the MappedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetMappedValuesOk() ([]ValueCreate, bool) {
	if o == nil || IsNil(o.MappedValues) {
		return nil, false
	}
	return o.MappedValues, true
}

// HasMappedValues returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasMappedValues() bool {
	if o != nil && !IsNil(o.MappedValues) {
		return true
	}

	return false
}

// SetMappedValues gets a reference to the given []ValueCreate and assigns it to the MappedValues field.
func (o *PatchedAwsPull) SetMappedValues(v []ValueCreate) {
	o.MappedValues = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetMode() ModeEnum {
	if o == nil || IsNil(o.Mode) {
		var ret ModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetModeOk() (*ModeEnum, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given ModeEnum and assigns it to the Mode field.
func (o *PatchedAwsPull) SetMode(v ModeEnum) {
	o.Mode = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetRegion() AwsRegionEnum {
	if o == nil || IsNil(o.Region) {
		var ret AwsRegionEnum
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetRegionOk() (*AwsRegionEnum, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given AwsRegionEnum and assigns it to the Region field.
func (o *PatchedAwsPull) SetRegion(v AwsRegionEnum) {
	o.Region = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetService() AwsServiceEnum {
	if o == nil || IsNil(o.Service) {
		var ret AwsServiceEnum
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetServiceOk() (*AwsServiceEnum, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given AwsServiceEnum and assigns it to the Service field.
func (o *PatchedAwsPull) SetService(v AwsServiceEnum) {
	o.Service = &v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsPull) GetResource() string {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret string
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsPull) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableString and assigns it to the Resource field.
func (o *PatchedAwsPull) SetResource(v string) {
	o.Resource.Set(&v)
}
// SetResourceNil sets the value for Resource to be an explicit nil
func (o *PatchedAwsPull) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *PatchedAwsPull) UnsetResource() {
	o.Resource.Unset()
}

func (o PatchedAwsPull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedAwsPull) ToMap() (map[string]interface{}, error) {
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
	if o.LatestTask.IsSet() {
		toSerialize["latest_task"] = o.LatestTask.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if !IsNil(o.CreateEnvironments) {
		toSerialize["create_environments"] = o.CreateEnvironments
	}
	if !IsNil(o.CreateProjects) {
		toSerialize["create_projects"] = o.CreateProjects
	}
	if !IsNil(o.DryRun) {
		toSerialize["dry_run"] = o.DryRun
	}
	if !IsNil(o.MappedValues) {
		toSerialize["mapped_values"] = o.MappedValues
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if o.Resource.IsSet() {
		toSerialize["resource"] = o.Resource.Get()
	}
	return toSerialize, nil
}

type NullablePatchedAwsPull struct {
	value *PatchedAwsPull
	isSet bool
}

func (v NullablePatchedAwsPull) Get() *PatchedAwsPull {
	return v.value
}

func (v *NullablePatchedAwsPull) Set(val *PatchedAwsPull) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAwsPull) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAwsPull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAwsPull(val *PatchedAwsPull) *NullablePatchedAwsPull {
	return &NullablePatchedAwsPull{value: val, isSet: true}
}

func (v NullablePatchedAwsPull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAwsPull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


