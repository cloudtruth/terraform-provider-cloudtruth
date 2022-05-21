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

// GitHubIntegration struct for GitHubIntegration
type GitHubIntegration struct {
	Url string `json:"url"`
	// The unique identifier for the integration.
	Id   string `json:"id"`
	Name string `json:"name"`
	// An optional description for the integration.
	Description *string `json:"description,omitempty"`
	// The status of the integration connection with the third-party provider as of the `status_last_checked_at` field.  The status is updated automatically by the server when the integration is modified.
	Status NullableStatusEnum `json:"status"`
	// If an error occurs, more details will be available in this field.
	StatusDetail string `json:"status_detail"`
	// The last time the status was evaluated.
	StatusLastCheckedAt time.Time `json:"status_last_checked_at"`
	CreatedAt           time.Time `json:"created_at"`
	ModifiedAt          time.Time `json:"modified_at"`
	Fqn                 string    `json:"fqn"`
	// The type of integration.
	Type string `json:"type"`
	// Allow actions to write to the integration.
	Writable           *bool  `json:"writable,omitempty"`
	GhInstallationId   int32  `json:"gh_installation_id"`
	GhOrganizationSlug string `json:"gh_organization_slug"`
}

// NewGitHubIntegration instantiates a new GitHubIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitHubIntegration(url string, id string, name string, status NullableStatusEnum, statusDetail string, statusLastCheckedAt time.Time, createdAt time.Time, modifiedAt time.Time, fqn string, type_ string, ghInstallationId int32, ghOrganizationSlug string) *GitHubIntegration {
	this := GitHubIntegration{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusDetail = statusDetail
	this.StatusLastCheckedAt = statusLastCheckedAt
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Fqn = fqn
	this.Type = type_
	this.GhInstallationId = ghInstallationId
	this.GhOrganizationSlug = ghOrganizationSlug
	return &this
}

// NewGitHubIntegrationWithDefaults instantiates a new GitHubIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubIntegrationWithDefaults() *GitHubIntegration {
	this := GitHubIntegration{}
	return &this
}

// GetUrl returns the Url field value
func (o *GitHubIntegration) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GitHubIntegration) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *GitHubIntegration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GitHubIntegration) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GitHubIntegration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitHubIntegration) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GitHubIntegration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GitHubIntegration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GitHubIntegration) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for StatusEnum will be returned
func (o *GitHubIntegration) GetStatus() StatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret StatusEnum
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitHubIntegration) GetStatusOk() (*StatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *GitHubIntegration) SetStatus(v StatusEnum) {
	o.Status.Set(&v)
}

// GetStatusDetail returns the StatusDetail field value
func (o *GitHubIntegration) GetStatusDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetStatusDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDetail, true
}

// SetStatusDetail sets field value
func (o *GitHubIntegration) SetStatusDetail(v string) {
	o.StatusDetail = v
}

// GetStatusLastCheckedAt returns the StatusLastCheckedAt field value
func (o *GitHubIntegration) GetStatusLastCheckedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StatusLastCheckedAt
}

// GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusLastCheckedAt, true
}

// SetStatusLastCheckedAt sets field value
func (o *GitHubIntegration) SetStatusLastCheckedAt(v time.Time) {
	o.StatusLastCheckedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GitHubIntegration) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GitHubIntegration) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *GitHubIntegration) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *GitHubIntegration) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetFqn returns the Fqn field value
func (o *GitHubIntegration) GetFqn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqn, true
}

// SetFqn sets field value
func (o *GitHubIntegration) SetFqn(v string) {
	o.Fqn = v
}

// GetType returns the Type field value
func (o *GitHubIntegration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GitHubIntegration) SetType(v string) {
	o.Type = v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *GitHubIntegration) GetWritable() bool {
	if o == nil || o.Writable == nil {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetWritableOk() (*bool, bool) {
	if o == nil || o.Writable == nil {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *GitHubIntegration) HasWritable() bool {
	if o != nil && o.Writable != nil {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *GitHubIntegration) SetWritable(v bool) {
	o.Writable = &v
}

// GetGhInstallationId returns the GhInstallationId field value
func (o *GitHubIntegration) GetGhInstallationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GhInstallationId
}

// GetGhInstallationIdOk returns a tuple with the GhInstallationId field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetGhInstallationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GhInstallationId, true
}

// SetGhInstallationId sets field value
func (o *GitHubIntegration) SetGhInstallationId(v int32) {
	o.GhInstallationId = v
}

// GetGhOrganizationSlug returns the GhOrganizationSlug field value
func (o *GitHubIntegration) GetGhOrganizationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GhOrganizationSlug
}

// GetGhOrganizationSlugOk returns a tuple with the GhOrganizationSlug field value
// and a boolean to check if the value has been set.
func (o *GitHubIntegration) GetGhOrganizationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GhOrganizationSlug, true
}

// SetGhOrganizationSlug sets field value
func (o *GitHubIntegration) SetGhOrganizationSlug(v string) {
	o.GhOrganizationSlug = v
}

func (o GitHubIntegration) MarshalJSON() ([]byte, error) {
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
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["status_detail"] = o.StatusDetail
	}
	if true {
		toSerialize["status_last_checked_at"] = o.StatusLastCheckedAt
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if true {
		toSerialize["fqn"] = o.Fqn
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Writable != nil {
		toSerialize["writable"] = o.Writable
	}
	if true {
		toSerialize["gh_installation_id"] = o.GhInstallationId
	}
	if true {
		toSerialize["gh_organization_slug"] = o.GhOrganizationSlug
	}
	return json.Marshal(toSerialize)
}

type NullableGitHubIntegration struct {
	value *GitHubIntegration
	isSet bool
}

func (v NullableGitHubIntegration) Get() *GitHubIntegration {
	return v.value
}

func (v *NullableGitHubIntegration) Set(val *GitHubIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableGitHubIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableGitHubIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitHubIntegration(val *GitHubIntegration) *NullableGitHubIntegration {
	return &NullableGitHubIntegration{value: val, isSet: true}
}

func (v NullableGitHubIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitHubIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
