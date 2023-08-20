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
)

// checks if the AzureKeyVaultPullSyncActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureKeyVaultPullSyncActionRequest{}

// AzureKeyVaultPullSyncActionRequest struct for AzureKeyVaultPullSyncActionRequest
type AzureKeyVaultPullSyncActionRequest struct {
	// Allows you to set the dry_run flag on the pull action before triggering a sync.
	DryRun *bool `json:"dry_run,omitempty"`
}

// NewAzureKeyVaultPullSyncActionRequest instantiates a new AzureKeyVaultPullSyncActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPullSyncActionRequest() *AzureKeyVaultPullSyncActionRequest {
	this := AzureKeyVaultPullSyncActionRequest{}
	return &this
}

// NewAzureKeyVaultPullSyncActionRequestWithDefaults instantiates a new AzureKeyVaultPullSyncActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPullSyncActionRequestWithDefaults() *AzureKeyVaultPullSyncActionRequest {
	this := AzureKeyVaultPullSyncActionRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AzureKeyVaultPullSyncActionRequest) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPullSyncActionRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AzureKeyVaultPullSyncActionRequest) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AzureKeyVaultPullSyncActionRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o AzureKeyVaultPullSyncActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureKeyVaultPullSyncActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DryRun) {
		toSerialize["dry_run"] = o.DryRun
	}
	return toSerialize, nil
}

type NullableAzureKeyVaultPullSyncActionRequest struct {
	value *AzureKeyVaultPullSyncActionRequest
	isSet bool
}

func (v NullableAzureKeyVaultPullSyncActionRequest) Get() *AzureKeyVaultPullSyncActionRequest {
	return v.value
}

func (v *NullableAzureKeyVaultPullSyncActionRequest) Set(val *AzureKeyVaultPullSyncActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPullSyncActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPullSyncActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPullSyncActionRequest(val *AzureKeyVaultPullSyncActionRequest) *NullableAzureKeyVaultPullSyncActionRequest {
	return &NullableAzureKeyVaultPullSyncActionRequest{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPullSyncActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPullSyncActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


