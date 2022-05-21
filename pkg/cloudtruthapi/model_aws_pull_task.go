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

// AwsPullTask Pull task for an AWS integration.
type AwsPullTask struct {
	Url string `json:"url"`
	// The unique identifier for the task.
	Id string `json:"id"`
	// The reason why this task was triggered.
	Reason NullableString `json:"reason,omitempty"`
	// Indicates task steps were only simulated, not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// The current state of this task.
	State NullableStateEnum `json:"state,omitempty"`
	// If an error occurs early during processing, before attempting to process values, this code may be helpful in determining the problem.
	ErrorCode NullableString `json:"error_code,omitempty"`
	// If an error occurs early during processing, before attempting to process values, this detail may be helpful in determining the problem.
	ErrorDetail NullableString `json:"error_detail,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	ModifiedAt  time.Time      `json:"modified_at"`
}

// NewAwsPullTask instantiates a new AwsPullTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsPullTask(url string, id string, createdAt time.Time, modifiedAt time.Time) *AwsPullTask {
	this := AwsPullTask{}
	this.Url = url
	this.Id = id
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAwsPullTaskWithDefaults instantiates a new AwsPullTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsPullTaskWithDefaults() *AwsPullTask {
	this := AwsPullTask{}
	return &this
}

// GetUrl returns the Url field value
func (o *AwsPullTask) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AwsPullTask) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AwsPullTask) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AwsPullTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AwsPullTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AwsPullTask) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPullTask) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPullTask) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *AwsPullTask) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *AwsPullTask) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *AwsPullTask) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *AwsPullTask) UnsetReason() {
	o.Reason.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AwsPullTask) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPullTask) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AwsPullTask) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AwsPullTask) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPullTask) GetState() StateEnum {
	if o == nil || o.State.Get() == nil {
		var ret StateEnum
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPullTask) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *AwsPullTask) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableStateEnum and assigns it to the State field.
func (o *AwsPullTask) SetState(v StateEnum) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil
func (o *AwsPullTask) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *AwsPullTask) UnsetState() {
	o.State.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPullTask) GetErrorCode() string {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPullTask) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AwsPullTask) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *AwsPullTask) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}

// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AwsPullTask) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AwsPullTask) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPullTask) GetErrorDetail() string {
	if o == nil || o.ErrorDetail.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorDetail.Get()
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPullTask) GetErrorDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDetail.Get(), o.ErrorDetail.IsSet()
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *AwsPullTask) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail.IsSet() {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given NullableString and assigns it to the ErrorDetail field.
func (o *AwsPullTask) SetErrorDetail(v string) {
	o.ErrorDetail.Set(&v)
}

// SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil
func (o *AwsPullTask) SetErrorDetailNil() {
	o.ErrorDetail.Set(nil)
}

// UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
func (o *AwsPullTask) UnsetErrorDetail() {
	o.ErrorDetail.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AwsPullTask) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AwsPullTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AwsPullTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AwsPullTask) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AwsPullTask) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AwsPullTask) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o AwsPullTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorDetail.IsSet() {
		toSerialize["error_detail"] = o.ErrorDetail.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAwsPullTask struct {
	value *AwsPullTask
	isSet bool
}

func (v NullableAwsPullTask) Get() *AwsPullTask {
	return v.value
}

func (v *NullableAwsPullTask) Set(val *AwsPullTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsPullTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsPullTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsPullTask(val *AwsPullTask) *NullableAwsPullTask {
	return &NullableAwsPullTask{value: val, isSet: true}
}

func (v NullableAwsPullTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsPullTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
