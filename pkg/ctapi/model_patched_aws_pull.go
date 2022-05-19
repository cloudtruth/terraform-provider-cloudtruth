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
	"time"
)

// PatchedAwsPull Pull actions can be configured to get configuration and secrets from integrations on demand.
type PatchedAwsPull struct {
	Url *string `json:"url,omitempty"`
	// Unique identifier for the action.
	Id *string `json:"id,omitempty"`
	// The action name.
	Name *string `json:"name,omitempty"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	// The most recent task run for this action.
	LatestTask NullableAwsPullTask `json:"latest_task,omitempty"`
	CreatedAt  *time.Time          `json:"created_at,omitempty"`
	ModifiedAt *time.Time          `json:"modified_at,omitempty"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// The AWS region this pull uses.  This region must be enabled in the integration.
	Region *AwsRegionEnum `json:"region,omitempty"`
	// The AWS service this pull uses.  This service must be enabled in the integration.
	Service *AwsServiceEnum `json:"service,omitempty"`
	// Defines a path through the integration to the location where values will be pulled.  The following mustache-style substitutions must be used in the resource locator string:    - ``{{ environment }}`` to identify the environment name   - ``{{ parameter }}`` to identify the parameter name   - ``{{ project }}`` to identify the project name
	Resource *string `json:"resource,omitempty"`
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
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasUrl() bool {
	if o != nil && o.Url != nil {
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
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasId() bool {
	if o != nil && o.Id != nil {
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedAwsPull) SetDescription(v string) {
	o.Description = &v
}

// GetLatestTask returns the LatestTask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsPull) GetLatestTask() AwsPullTask {
	if o == nil || o.LatestTask.Get() == nil {
		var ret AwsPullTask
		return ret
	}
	return *o.LatestTask.Get()
}

// GetLatestTaskOk returns a tuple with the LatestTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsPull) GetLatestTaskOk() (*AwsPullTask, bool) {
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

// SetLatestTask gets a reference to the given NullableAwsPullTask and assigns it to the LatestTask field.
func (o *PatchedAwsPull) SetLatestTask(v AwsPullTask) {
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
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedAwsPull) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedAwsPull) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *PatchedAwsPull) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetRegion() AwsRegionEnum {
	if o == nil || o.Region == nil {
		var ret AwsRegionEnum
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetRegionOk() (*AwsRegionEnum, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasRegion() bool {
	if o != nil && o.Region != nil {
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
	if o == nil || o.Service == nil {
		var ret AwsServiceEnum
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetServiceOk() (*AwsServiceEnum, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given AwsServiceEnum and assigns it to the Service field.
func (o *PatchedAwsPull) SetService(v AwsServiceEnum) {
	o.Service = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PatchedAwsPull) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPull) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PatchedAwsPull) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *PatchedAwsPull) SetResource(v string) {
	o.Resource = &v
}

func (o PatchedAwsPull) MarshalJSON() ([]byte, error) {
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
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
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
