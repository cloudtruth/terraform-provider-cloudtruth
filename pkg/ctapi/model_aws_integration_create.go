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
)

// AwsIntegrationCreate struct for AwsIntegrationCreate
type AwsIntegrationCreate struct {
	// An optional description for the integration.
	Description *string `json:"description,omitempty"`
	// Allow actions to write to the integration.
	Writable *bool `json:"writable,omitempty"`
	// The AWS Account ID.
	AwsAccountId string `json:"aws_account_id"`
	// The AWS regions to integrate with.
	AwsEnabledRegions []AwsRegionEnum `json:"aws_enabled_regions"`
	// The AWS services to integrate with.
	AwsEnabledServices []AwsServiceEnum `json:"aws_enabled_services"`
	// This is a shared secret between the AWS Administrator who set up your IAM trust relationship and your CloudTruth AWS Integration.  If your AWS Administrator provided you with a value use it, otherwise we will generate a random value for you to give to your AWS Administrator.
	AwsExternalId *string `json:"aws_external_id,omitempty"`
	// The role that CloudTruth will assume when interacting with your AWS Account through this integration.  The role is configured by your AWS Account Administrator.  If your AWS Administrator provided you with a value use it, otherwise make your own role name and give it to your AWS Administrator.
	AwsRoleName string `json:"aws_role_name"`
}

// NewAwsIntegrationCreate instantiates a new AwsIntegrationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsIntegrationCreate(awsAccountId string, awsEnabledRegions []AwsRegionEnum, awsEnabledServices []AwsServiceEnum, awsRoleName string) *AwsIntegrationCreate {
	this := AwsIntegrationCreate{}
	this.AwsAccountId = awsAccountId
	this.AwsEnabledRegions = awsEnabledRegions
	this.AwsEnabledServices = awsEnabledServices
	this.AwsRoleName = awsRoleName
	return &this
}

// NewAwsIntegrationCreateWithDefaults instantiates a new AwsIntegrationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsIntegrationCreateWithDefaults() *AwsIntegrationCreate {
	this := AwsIntegrationCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsIntegrationCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsIntegrationCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AwsIntegrationCreate) SetDescription(v string) {
	o.Description = &v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *AwsIntegrationCreate) GetWritable() bool {
	if o == nil || o.Writable == nil {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetWritableOk() (*bool, bool) {
	if o == nil || o.Writable == nil {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *AwsIntegrationCreate) HasWritable() bool {
	if o != nil && o.Writable != nil {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *AwsIntegrationCreate) SetWritable(v bool) {
	o.Writable = &v
}

// GetAwsAccountId returns the AwsAccountId field value
func (o *AwsIntegrationCreate) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value
func (o *AwsIntegrationCreate) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsEnabledRegions returns the AwsEnabledRegions field value
func (o *AwsIntegrationCreate) GetAwsEnabledRegions() []AwsRegionEnum {
	if o == nil {
		var ret []AwsRegionEnum
		return ret
	}

	return o.AwsEnabledRegions
}

// GetAwsEnabledRegionsOk returns a tuple with the AwsEnabledRegions field value
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetAwsEnabledRegionsOk() (*[]AwsRegionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsEnabledRegions, true
}

// SetAwsEnabledRegions sets field value
func (o *AwsIntegrationCreate) SetAwsEnabledRegions(v []AwsRegionEnum) {
	o.AwsEnabledRegions = v
}

// GetAwsEnabledServices returns the AwsEnabledServices field value
func (o *AwsIntegrationCreate) GetAwsEnabledServices() []AwsServiceEnum {
	if o == nil {
		var ret []AwsServiceEnum
		return ret
	}

	return o.AwsEnabledServices
}

// GetAwsEnabledServicesOk returns a tuple with the AwsEnabledServices field value
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetAwsEnabledServicesOk() (*[]AwsServiceEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsEnabledServices, true
}

// SetAwsEnabledServices sets field value
func (o *AwsIntegrationCreate) SetAwsEnabledServices(v []AwsServiceEnum) {
	o.AwsEnabledServices = v
}

// GetAwsExternalId returns the AwsExternalId field value if set, zero value otherwise.
func (o *AwsIntegrationCreate) GetAwsExternalId() string {
	if o == nil || o.AwsExternalId == nil {
		var ret string
		return ret
	}
	return *o.AwsExternalId
}

// GetAwsExternalIdOk returns a tuple with the AwsExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetAwsExternalIdOk() (*string, bool) {
	if o == nil || o.AwsExternalId == nil {
		return nil, false
	}
	return o.AwsExternalId, true
}

// HasAwsExternalId returns a boolean if a field has been set.
func (o *AwsIntegrationCreate) HasAwsExternalId() bool {
	if o != nil && o.AwsExternalId != nil {
		return true
	}

	return false
}

// SetAwsExternalId gets a reference to the given string and assigns it to the AwsExternalId field.
func (o *AwsIntegrationCreate) SetAwsExternalId(v string) {
	o.AwsExternalId = &v
}

// GetAwsRoleName returns the AwsRoleName field value
func (o *AwsIntegrationCreate) GetAwsRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsRoleName
}

// GetAwsRoleNameOk returns a tuple with the AwsRoleName field value
// and a boolean to check if the value has been set.
func (o *AwsIntegrationCreate) GetAwsRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRoleName, true
}

// SetAwsRoleName sets field value
func (o *AwsIntegrationCreate) SetAwsRoleName(v string) {
	o.AwsRoleName = v
}

func (o AwsIntegrationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Writable != nil {
		toSerialize["writable"] = o.Writable
	}
	if true {
		toSerialize["aws_account_id"] = o.AwsAccountId
	}
	if true {
		toSerialize["aws_enabled_regions"] = o.AwsEnabledRegions
	}
	if true {
		toSerialize["aws_enabled_services"] = o.AwsEnabledServices
	}
	if o.AwsExternalId != nil {
		toSerialize["aws_external_id"] = o.AwsExternalId
	}
	if true {
		toSerialize["aws_role_name"] = o.AwsRoleName
	}
	return json.Marshal(toSerialize)
}

type NullableAwsIntegrationCreate struct {
	value *AwsIntegrationCreate
	isSet bool
}

func (v NullableAwsIntegrationCreate) Get() *AwsIntegrationCreate {
	return v.value
}

func (v *NullableAwsIntegrationCreate) Set(val *AwsIntegrationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsIntegrationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsIntegrationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsIntegrationCreate(val *AwsIntegrationCreate) *NullableAwsIntegrationCreate {
	return &NullableAwsIntegrationCreate{value: val, isSet: true}
}

func (v NullableAwsIntegrationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsIntegrationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
