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

// checks if the ImportParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportParameter{}

// ImportParameter Describes an imported parameter.
type ImportParameter struct {
	// Project name
	ProjectName string `json:"project_name"`
	// Project identifier
	ProjectId NullableString `json:"project_id,omitempty"`
	// Environment where parameter is defined
	EnvironmentName string `json:"environment_name"`
	// Environment identifier where value is set
	EnvironmentId NullableString `json:"environment_id,omitempty"`
	// Parameter name
	ParameterName string `json:"parameter_name"`
	// Parameter identifier
	ParameterId NullableString `json:"parameter_id,omitempty"`
	// Whether to treat the parameter as a secret
	Secret *bool `json:"secret,omitempty"`
	// Parameter value in the specified environment
	Value string `json:"value"`
	// Parameter value identifier in the environment
	ValueId NullableString `json:"value_id,omitempty"`
	// Date and time the parameter value was created
	CreatedAt NullableTime `json:"created_at"`
	// Date and time the parameter value was updated
	ModifiedAt NullableTime `json:"modified_at"`
	// Action taken on environment value for parameter
	Action string `json:"action"`
}

// NewImportParameter instantiates a new ImportParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportParameter(projectName string, environmentName string, parameterName string, value string, createdAt NullableTime, modifiedAt NullableTime, action string) *ImportParameter {
	this := ImportParameter{}
	this.ProjectName = projectName
	this.EnvironmentName = environmentName
	this.ParameterName = parameterName
	var secret bool = false
	this.Secret = &secret
	this.Value = value
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Action = action
	return &this
}

// NewImportParameterWithDefaults instantiates a new ImportParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportParameterWithDefaults() *ImportParameter {
	this := ImportParameter{}
	var secret bool = false
	this.Secret = &secret
	return &this
}

// GetProjectName returns the ProjectName field value
func (o *ImportParameter) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ImportParameter) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ImportParameter) SetProjectName(v string) {
	o.ProjectName = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportParameter) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportParameter) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *ImportParameter) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *ImportParameter) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *ImportParameter) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *ImportParameter) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetEnvironmentName returns the EnvironmentName field value
func (o *ImportParameter) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *ImportParameter) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value
func (o *ImportParameter) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportParameter) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportParameter) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ImportParameter) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *ImportParameter) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}
// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *ImportParameter) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *ImportParameter) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetParameterName returns the ParameterName field value
func (o *ImportParameter) GetParameterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value
// and a boolean to check if the value has been set.
func (o *ImportParameter) GetParameterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterName, true
}

// SetParameterName sets field value
func (o *ImportParameter) SetParameterName(v string) {
	o.ParameterName = v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportParameter) GetParameterId() string {
	if o == nil || IsNil(o.ParameterId.Get()) {
		var ret string
		return ret
	}
	return *o.ParameterId.Get()
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportParameter) GetParameterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParameterId.Get(), o.ParameterId.IsSet()
}

// HasParameterId returns a boolean if a field has been set.
func (o *ImportParameter) HasParameterId() bool {
	if o != nil && o.ParameterId.IsSet() {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given NullableString and assigns it to the ParameterId field.
func (o *ImportParameter) SetParameterId(v string) {
	o.ParameterId.Set(&v)
}
// SetParameterIdNil sets the value for ParameterId to be an explicit nil
func (o *ImportParameter) SetParameterIdNil() {
	o.ParameterId.Set(nil)
}

// UnsetParameterId ensures that no value is present for ParameterId, not even an explicit nil
func (o *ImportParameter) UnsetParameterId() {
	o.ParameterId.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ImportParameter) GetSecret() bool {
	if o == nil || IsNil(o.Secret) {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportParameter) GetSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ImportParameter) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *ImportParameter) SetSecret(v bool) {
	o.Secret = &v
}

// GetValue returns the Value field value
func (o *ImportParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ImportParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ImportParameter) SetValue(v string) {
	o.Value = v
}

// GetValueId returns the ValueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportParameter) GetValueId() string {
	if o == nil || IsNil(o.ValueId.Get()) {
		var ret string
		return ret
	}
	return *o.ValueId.Get()
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportParameter) GetValueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueId.Get(), o.ValueId.IsSet()
}

// HasValueId returns a boolean if a field has been set.
func (o *ImportParameter) HasValueId() bool {
	if o != nil && o.ValueId.IsSet() {
		return true
	}

	return false
}

// SetValueId gets a reference to the given NullableString and assigns it to the ValueId field.
func (o *ImportParameter) SetValueId(v string) {
	o.ValueId.Set(&v)
}
// SetValueIdNil sets the value for ValueId to be an explicit nil
func (o *ImportParameter) SetValueIdNil() {
	o.ValueId.Set(nil)
}

// UnsetValueId ensures that no value is present for ValueId, not even an explicit nil
func (o *ImportParameter) UnsetValueId() {
	o.ValueId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ImportParameter) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportParameter) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *ImportParameter) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ImportParameter) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportParameter) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *ImportParameter) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// GetAction returns the Action field value
func (o *ImportParameter) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ImportParameter) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ImportParameter) SetAction(v string) {
	o.Action = v
}

func (o ImportParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_name"] = o.ProjectName
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	toSerialize["environment_name"] = o.EnvironmentName
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	toSerialize["parameter_name"] = o.ParameterName
	if o.ParameterId.IsSet() {
		toSerialize["parameter_id"] = o.ParameterId.Get()
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	toSerialize["value"] = o.Value
	if o.ValueId.IsSet() {
		toSerialize["value_id"] = o.ValueId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableImportParameter struct {
	value *ImportParameter
	isSet bool
}

func (v NullableImportParameter) Get() *ImportParameter {
	return v.value
}

func (v *NullableImportParameter) Set(val *ImportParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableImportParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableImportParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportParameter(val *ImportParameter) *NullableImportParameter {
	return &NullableImportParameter{value: val, isSet: true}
}

func (v NullableImportParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


