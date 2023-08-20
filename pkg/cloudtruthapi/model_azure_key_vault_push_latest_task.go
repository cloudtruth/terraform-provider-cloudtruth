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

// checks if the AzureKeyVaultPushLatestTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureKeyVaultPushLatestTask{}

// AzureKeyVaultPushLatestTask The most recent task run for this action.
type AzureKeyVaultPushLatestTask struct {
	Url string `json:"url"`
	// The unique identifier for the task.
	Id string `json:"id"`
	// The reason why this task was triggered.
	Reason NullableString `json:"reason,omitempty"`
	// Indicates task steps were only simulated, not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	State *StateEnum `json:"state,omitempty"`
	// If an error occurs early during processing, before attempting to process values, this code may be helpful in determining the problem.
	ErrorCode NullableString `json:"error_code,omitempty"`
	// If an error occurs early during processing, before attempting to process values, this detail may be helpful in determining the problem.
	ErrorDetail NullableString `json:"error_detail,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewAzureKeyVaultPushLatestTask instantiates a new AzureKeyVaultPushLatestTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPushLatestTask(url string, id string, createdAt time.Time, modifiedAt time.Time) *AzureKeyVaultPushLatestTask {
	this := AzureKeyVaultPushLatestTask{}
	this.Url = url
	this.Id = id
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAzureKeyVaultPushLatestTaskWithDefaults instantiates a new AzureKeyVaultPushLatestTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPushLatestTaskWithDefaults() *AzureKeyVaultPushLatestTask {
	this := AzureKeyVaultPushLatestTask{}
	return &this
}

// GetUrl returns the Url field value
func (o *AzureKeyVaultPushLatestTask) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushLatestTask) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AzureKeyVaultPushLatestTask) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AzureKeyVaultPushLatestTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushLatestTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultPushLatestTask) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushLatestTask) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushLatestTask) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *AzureKeyVaultPushLatestTask) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *AzureKeyVaultPushLatestTask) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *AzureKeyVaultPushLatestTask) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *AzureKeyVaultPushLatestTask) UnsetReason() {
	o.Reason.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AzureKeyVaultPushLatestTask) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushLatestTask) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AzureKeyVaultPushLatestTask) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AzureKeyVaultPushLatestTask) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AzureKeyVaultPushLatestTask) GetState() StateEnum {
	if o == nil || IsNil(o.State) {
		var ret StateEnum
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushLatestTask) GetStateOk() (*StateEnum, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AzureKeyVaultPushLatestTask) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given StateEnum and assigns it to the State field.
func (o *AzureKeyVaultPushLatestTask) SetState(v StateEnum) {
	o.State = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushLatestTask) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushLatestTask) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AzureKeyVaultPushLatestTask) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *AzureKeyVaultPushLatestTask) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AzureKeyVaultPushLatestTask) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AzureKeyVaultPushLatestTask) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushLatestTask) GetErrorDetail() string {
	if o == nil || IsNil(o.ErrorDetail.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorDetail.Get()
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushLatestTask) GetErrorDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDetail.Get(), o.ErrorDetail.IsSet()
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *AzureKeyVaultPushLatestTask) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail.IsSet() {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given NullableString and assigns it to the ErrorDetail field.
func (o *AzureKeyVaultPushLatestTask) SetErrorDetail(v string) {
	o.ErrorDetail.Set(&v)
}
// SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil
func (o *AzureKeyVaultPushLatestTask) SetErrorDetailNil() {
	o.ErrorDetail.Set(nil)
}

// UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
func (o *AzureKeyVaultPushLatestTask) UnsetErrorDetail() {
	o.ErrorDetail.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AzureKeyVaultPushLatestTask) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushLatestTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AzureKeyVaultPushLatestTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AzureKeyVaultPushLatestTask) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushLatestTask) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AzureKeyVaultPushLatestTask) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o AzureKeyVaultPushLatestTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureKeyVaultPushLatestTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if !IsNil(o.DryRun) {
		toSerialize["dry_run"] = o.DryRun
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorDetail.IsSet() {
		toSerialize["error_detail"] = o.ErrorDetail.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt
	return toSerialize, nil
}

type NullableAzureKeyVaultPushLatestTask struct {
	value *AzureKeyVaultPushLatestTask
	isSet bool
}

func (v NullableAzureKeyVaultPushLatestTask) Get() *AzureKeyVaultPushLatestTask {
	return v.value
}

func (v *NullableAzureKeyVaultPushLatestTask) Set(val *AzureKeyVaultPushLatestTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPushLatestTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPushLatestTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPushLatestTask(val *AzureKeyVaultPushLatestTask) *NullableAzureKeyVaultPushLatestTask {
	return &NullableAzureKeyVaultPushLatestTask{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPushLatestTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPushLatestTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


