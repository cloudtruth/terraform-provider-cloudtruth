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

// PatchedAwsIntegration struct for PatchedAwsIntegration
type PatchedAwsIntegration struct {
	Url *string `json:"url,omitempty"`
	// The unique identifier for the integration.
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// An optional description for the integration.
	Description *string `json:"description,omitempty"`
	// The status of the integration connection with the third-party provider as of the `status_last_checked_at` field.  The status is updated automatically by the server when the integration is modified.
	Status NullableStatusEnum `json:"status,omitempty"`
	// If an error occurs, more details will be available in this field.
	StatusDetail *string `json:"status_detail,omitempty"`
	// The last time the status was evaluated.
	StatusLastCheckedAt *time.Time `json:"status_last_checked_at,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	ModifiedAt          *time.Time `json:"modified_at,omitempty"`
	Fqn                 *string    `json:"fqn,omitempty"`
	// The type of integration.
	Type *string `json:"type,omitempty"`
	// Allow actions to write to the integration.
	Writable *bool `json:"writable,omitempty"`
	// The AWS Account ID.
	AwsAccountId *string `json:"aws_account_id,omitempty"`
	// The AWS regions to integrate with.
	AwsEnabledRegions []AwsRegionEnum `json:"aws_enabled_regions,omitempty"`
	// The AWS services to integrate with.
	AwsEnabledServices []AwsServiceEnum `json:"aws_enabled_services,omitempty"`
	// This is a shared secret between the AWS Administrator who set up your IAM trust relationship and your CloudTruth AWS Integration.  CloudTruth will generate a random value for you to give to your AWS Administrator in order to create the necessary IAM role for proper access.
	AwsExternalId *string `json:"aws_external_id,omitempty"`
	// If present, this is the KMS Key Id that is used to push values.  This key must be accessible in the AWS account (it cannot be an ARN to a key in another AWS account).
	AwsKmsKeyId NullableString `json:"aws_kms_key_id,omitempty"`
	// The role that CloudTruth will assume when interacting with your AWS Account through this integration.  The role is configured by your AWS Account Administrator.  If your AWS Administrator provided you with a value use it, otherwise make your own role name and give it to your AWS Administrator.
	AwsRoleName *string `json:"aws_role_name,omitempty"`
}

// NewPatchedAwsIntegration instantiates a new PatchedAwsIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAwsIntegration() *PatchedAwsIntegration {
	this := PatchedAwsIntegration{}
	return &this
}

// NewPatchedAwsIntegrationWithDefaults instantiates a new PatchedAwsIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAwsIntegrationWithDefaults() *PatchedAwsIntegration {
	this := PatchedAwsIntegration{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedAwsIntegration) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedAwsIntegration) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAwsIntegration) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedAwsIntegration) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsIntegration) GetStatus() StatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret StatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsIntegration) GetStatusOk() (*StatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableStatusEnum and assigns it to the Status field.
func (o *PatchedAwsIntegration) SetStatus(v StatusEnum) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *PatchedAwsIntegration) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PatchedAwsIntegration) UnsetStatus() {
	o.Status.Unset()
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetStatusDetail() string {
	if o == nil || o.StatusDetail == nil {
		var ret string
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetStatusDetailOk() (*string, bool) {
	if o == nil || o.StatusDetail == nil {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasStatusDetail() bool {
	if o != nil && o.StatusDetail != nil {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given string and assigns it to the StatusDetail field.
func (o *PatchedAwsIntegration) SetStatusDetail(v string) {
	o.StatusDetail = &v
}

// GetStatusLastCheckedAt returns the StatusLastCheckedAt field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetStatusLastCheckedAt() time.Time {
	if o == nil || o.StatusLastCheckedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusLastCheckedAt
}

// GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool) {
	if o == nil || o.StatusLastCheckedAt == nil {
		return nil, false
	}
	return o.StatusLastCheckedAt, true
}

// HasStatusLastCheckedAt returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasStatusLastCheckedAt() bool {
	if o != nil && o.StatusLastCheckedAt != nil {
		return true
	}

	return false
}

// SetStatusLastCheckedAt gets a reference to the given time.Time and assigns it to the StatusLastCheckedAt field.
func (o *PatchedAwsIntegration) SetStatusLastCheckedAt(v time.Time) {
	o.StatusLastCheckedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedAwsIntegration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedAwsIntegration) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetFqn returns the Fqn field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetFqn() string {
	if o == nil || o.Fqn == nil {
		var ret string
		return ret
	}
	return *o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetFqnOk() (*string, bool) {
	if o == nil || o.Fqn == nil {
		return nil, false
	}
	return o.Fqn, true
}

// HasFqn returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasFqn() bool {
	if o != nil && o.Fqn != nil {
		return true
	}

	return false
}

// SetFqn gets a reference to the given string and assigns it to the Fqn field.
func (o *PatchedAwsIntegration) SetFqn(v string) {
	o.Fqn = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedAwsIntegration) SetType(v string) {
	o.Type = &v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetWritable() bool {
	if o == nil || o.Writable == nil {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetWritableOk() (*bool, bool) {
	if o == nil || o.Writable == nil {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasWritable() bool {
	if o != nil && o.Writable != nil {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *PatchedAwsIntegration) SetWritable(v bool) {
	o.Writable = &v
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetAwsAccountId() string {
	if o == nil || o.AwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || o.AwsAccountId == nil {
		return nil, false
	}
	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasAwsAccountId() bool {
	if o != nil && o.AwsAccountId != nil {
		return true
	}

	return false
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *PatchedAwsIntegration) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

// GetAwsEnabledRegions returns the AwsEnabledRegions field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetAwsEnabledRegions() []AwsRegionEnum {
	if o == nil || o.AwsEnabledRegions == nil {
		var ret []AwsRegionEnum
		return ret
	}
	return o.AwsEnabledRegions
}

// GetAwsEnabledRegionsOk returns a tuple with the AwsEnabledRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetAwsEnabledRegionsOk() ([]AwsRegionEnum, bool) {
	if o == nil || o.AwsEnabledRegions == nil {
		return nil, false
	}
	return o.AwsEnabledRegions, true
}

// HasAwsEnabledRegions returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasAwsEnabledRegions() bool {
	if o != nil && o.AwsEnabledRegions != nil {
		return true
	}

	return false
}

// SetAwsEnabledRegions gets a reference to the given []AwsRegionEnum and assigns it to the AwsEnabledRegions field.
func (o *PatchedAwsIntegration) SetAwsEnabledRegions(v []AwsRegionEnum) {
	o.AwsEnabledRegions = v
}

// GetAwsEnabledServices returns the AwsEnabledServices field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetAwsEnabledServices() []AwsServiceEnum {
	if o == nil || o.AwsEnabledServices == nil {
		var ret []AwsServiceEnum
		return ret
	}
	return o.AwsEnabledServices
}

// GetAwsEnabledServicesOk returns a tuple with the AwsEnabledServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetAwsEnabledServicesOk() ([]AwsServiceEnum, bool) {
	if o == nil || o.AwsEnabledServices == nil {
		return nil, false
	}
	return o.AwsEnabledServices, true
}

// HasAwsEnabledServices returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasAwsEnabledServices() bool {
	if o != nil && o.AwsEnabledServices != nil {
		return true
	}

	return false
}

// SetAwsEnabledServices gets a reference to the given []AwsServiceEnum and assigns it to the AwsEnabledServices field.
func (o *PatchedAwsIntegration) SetAwsEnabledServices(v []AwsServiceEnum) {
	o.AwsEnabledServices = v
}

// GetAwsExternalId returns the AwsExternalId field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetAwsExternalId() string {
	if o == nil || o.AwsExternalId == nil {
		var ret string
		return ret
	}
	return *o.AwsExternalId
}

// GetAwsExternalIdOk returns a tuple with the AwsExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetAwsExternalIdOk() (*string, bool) {
	if o == nil || o.AwsExternalId == nil {
		return nil, false
	}
	return o.AwsExternalId, true
}

// HasAwsExternalId returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasAwsExternalId() bool {
	if o != nil && o.AwsExternalId != nil {
		return true
	}

	return false
}

// SetAwsExternalId gets a reference to the given string and assigns it to the AwsExternalId field.
func (o *PatchedAwsIntegration) SetAwsExternalId(v string) {
	o.AwsExternalId = &v
}

// GetAwsKmsKeyId returns the AwsKmsKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsIntegration) GetAwsKmsKeyId() string {
	if o == nil || o.AwsKmsKeyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AwsKmsKeyId.Get()
}

// GetAwsKmsKeyIdOk returns a tuple with the AwsKmsKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsIntegration) GetAwsKmsKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsKmsKeyId.Get(), o.AwsKmsKeyId.IsSet()
}

// HasAwsKmsKeyId returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasAwsKmsKeyId() bool {
	if o != nil && o.AwsKmsKeyId.IsSet() {
		return true
	}

	return false
}

// SetAwsKmsKeyId gets a reference to the given NullableString and assigns it to the AwsKmsKeyId field.
func (o *PatchedAwsIntegration) SetAwsKmsKeyId(v string) {
	o.AwsKmsKeyId.Set(&v)
}

// SetAwsKmsKeyIdNil sets the value for AwsKmsKeyId to be an explicit nil
func (o *PatchedAwsIntegration) SetAwsKmsKeyIdNil() {
	o.AwsKmsKeyId.Set(nil)
}

// UnsetAwsKmsKeyId ensures that no value is present for AwsKmsKeyId, not even an explicit nil
func (o *PatchedAwsIntegration) UnsetAwsKmsKeyId() {
	o.AwsKmsKeyId.Unset()
}

// GetAwsRoleName returns the AwsRoleName field value if set, zero value otherwise.
func (o *PatchedAwsIntegration) GetAwsRoleName() string {
	if o == nil || o.AwsRoleName == nil {
		var ret string
		return ret
	}
	return *o.AwsRoleName
}

// GetAwsRoleNameOk returns a tuple with the AwsRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsIntegration) GetAwsRoleNameOk() (*string, bool) {
	if o == nil || o.AwsRoleName == nil {
		return nil, false
	}
	return o.AwsRoleName, true
}

// HasAwsRoleName returns a boolean if a field has been set.
func (o *PatchedAwsIntegration) HasAwsRoleName() bool {
	if o != nil && o.AwsRoleName != nil {
		return true
	}

	return false
}

// SetAwsRoleName gets a reference to the given string and assigns it to the AwsRoleName field.
func (o *PatchedAwsIntegration) SetAwsRoleName(v string) {
	o.AwsRoleName = &v
}

func (o PatchedAwsIntegration) MarshalJSON() ([]byte, error) {
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
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StatusDetail != nil {
		toSerialize["status_detail"] = o.StatusDetail
	}
	if o.StatusLastCheckedAt != nil {
		toSerialize["status_last_checked_at"] = o.StatusLastCheckedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Fqn != nil {
		toSerialize["fqn"] = o.Fqn
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Writable != nil {
		toSerialize["writable"] = o.Writable
	}
	if o.AwsAccountId != nil {
		toSerialize["aws_account_id"] = o.AwsAccountId
	}
	if o.AwsEnabledRegions != nil {
		toSerialize["aws_enabled_regions"] = o.AwsEnabledRegions
	}
	if o.AwsEnabledServices != nil {
		toSerialize["aws_enabled_services"] = o.AwsEnabledServices
	}
	if o.AwsExternalId != nil {
		toSerialize["aws_external_id"] = o.AwsExternalId
	}
	if o.AwsKmsKeyId.IsSet() {
		toSerialize["aws_kms_key_id"] = o.AwsKmsKeyId.Get()
	}
	if o.AwsRoleName != nil {
		toSerialize["aws_role_name"] = o.AwsRoleName
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAwsIntegration struct {
	value *PatchedAwsIntegration
	isSet bool
}

func (v NullablePatchedAwsIntegration) Get() *PatchedAwsIntegration {
	return v.value
}

func (v *NullablePatchedAwsIntegration) Set(val *PatchedAwsIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAwsIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAwsIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAwsIntegration(val *PatchedAwsIntegration) *NullablePatchedAwsIntegration {
	return &NullablePatchedAwsIntegration{value: val, isSet: true}
}

func (v NullablePatchedAwsIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAwsIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
