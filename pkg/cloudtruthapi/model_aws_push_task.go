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
	"fmt"
)

// checks if the AwsPushTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsPushTask{}

// AwsPushTask Push task for an AWS integration.
type AwsPushTask struct {
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
	ModifiedAt NullableTime `json:"modified_at"`
}

type _AwsPushTask AwsPushTask

// NewAwsPushTask instantiates a new AwsPushTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsPushTask(url string, id string, createdAt time.Time, modifiedAt NullableTime) *AwsPushTask {
	this := AwsPushTask{}
	this.Url = url
	this.Id = id
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAwsPushTaskWithDefaults instantiates a new AwsPushTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsPushTaskWithDefaults() *AwsPushTask {
	this := AwsPushTask{}
	return &this
}

// GetUrl returns the Url field value
func (o *AwsPushTask) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AwsPushTask) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AwsPushTask) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AwsPushTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AwsPushTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AwsPushTask) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPushTask) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushTask) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *AwsPushTask) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *AwsPushTask) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *AwsPushTask) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *AwsPushTask) UnsetReason() {
	o.Reason.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AwsPushTask) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPushTask) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AwsPushTask) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AwsPushTask) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AwsPushTask) GetState() StateEnum {
	if o == nil || IsNil(o.State) {
		var ret StateEnum
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPushTask) GetStateOk() (*StateEnum, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AwsPushTask) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given StateEnum and assigns it to the State field.
func (o *AwsPushTask) SetState(v StateEnum) {
	o.State = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPushTask) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushTask) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AwsPushTask) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *AwsPushTask) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AwsPushTask) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AwsPushTask) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsPushTask) GetErrorDetail() string {
	if o == nil || IsNil(o.ErrorDetail.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorDetail.Get()
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushTask) GetErrorDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDetail.Get(), o.ErrorDetail.IsSet()
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *AwsPushTask) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail.IsSet() {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given NullableString and assigns it to the ErrorDetail field.
func (o *AwsPushTask) SetErrorDetail(v string) {
	o.ErrorDetail.Set(&v)
}
// SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil
func (o *AwsPushTask) SetErrorDetailNil() {
	o.ErrorDetail.Set(nil)
}

// UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
func (o *AwsPushTask) UnsetErrorDetail() {
	o.ErrorDetail.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AwsPushTask) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AwsPushTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AwsPushTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AwsPushTask) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsPushTask) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AwsPushTask) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o AwsPushTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsPushTask) ToMap() (map[string]interface{}, error) {
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
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	return toSerialize, nil
}

func (o *AwsPushTask) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"id",
		"created_at",
		"modified_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAwsPushTask := _AwsPushTask{}

	err = json.Unmarshal(bytes, &varAwsPushTask)

	if err != nil {
		return err
	}

	*o = AwsPushTask(varAwsPushTask)

	return err
}

type NullableAwsPushTask struct {
	value *AwsPushTask
	isSet bool
}

func (v NullableAwsPushTask) Get() *AwsPushTask {
	return v.value
}

func (v *NullableAwsPushTask) Set(val *AwsPushTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsPushTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsPushTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsPushTask(val *AwsPushTask) *NullableAwsPushTask {
	return &NullableAwsPushTask{value: val, isSet: true}
}

func (v NullableAwsPushTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsPushTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


