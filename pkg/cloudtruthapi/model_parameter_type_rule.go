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

// checks if the ParameterTypeRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterTypeRule{}

// ParameterTypeRule A type of `ModelSerializer` that uses hyperlinked relationships with compound keys instead of primary key relationships.  Specifically:  * A 'url' field is included instead of the 'id' field. * Relationships to other instances are hyperlinks, instead of primary keys.  NOTE: this only works with DRF 3.1.0 and above.
type ParameterTypeRule struct {
	// The URL for the project.
	Url string `json:"url"`
	Id string `json:"id"`
	LedgerId string `json:"ledger_id"`
	// The type this rule is for.
	ParameterType string `json:"parameter_type"`
	Type ParameterRuleTypeEnum `json:"type"`
	Constraint string `json:"constraint"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewParameterTypeRule instantiates a new ParameterTypeRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterTypeRule(url string, id string, ledgerId string, parameterType string, type_ ParameterRuleTypeEnum, constraint string, createdAt time.Time, modifiedAt NullableTime) *ParameterTypeRule {
	this := ParameterTypeRule{}
	this.Url = url
	this.Id = id
	this.LedgerId = ledgerId
	this.ParameterType = parameterType
	this.Type = type_
	this.Constraint = constraint
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewParameterTypeRuleWithDefaults instantiates a new ParameterTypeRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterTypeRuleWithDefaults() *ParameterTypeRule {
	this := ParameterTypeRule{}
	return &this
}

// GetUrl returns the Url field value
func (o *ParameterTypeRule) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ParameterTypeRule) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *ParameterTypeRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterTypeRule) SetId(v string) {
	o.Id = v
}

// GetLedgerId returns the LedgerId field value
func (o *ParameterTypeRule) GetLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerId, true
}

// SetLedgerId sets field value
func (o *ParameterTypeRule) SetLedgerId(v string) {
	o.LedgerId = v
}

// GetParameterType returns the ParameterType field value
func (o *ParameterTypeRule) GetParameterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetParameterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterType, true
}

// SetParameterType sets field value
func (o *ParameterTypeRule) SetParameterType(v string) {
	o.ParameterType = v
}

// GetType returns the Type field value
func (o *ParameterTypeRule) GetType() ParameterRuleTypeEnum {
	if o == nil {
		var ret ParameterRuleTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetTypeOk() (*ParameterRuleTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParameterTypeRule) SetType(v ParameterRuleTypeEnum) {
	o.Type = v
}

// GetConstraint returns the Constraint field value
func (o *ParameterTypeRule) GetConstraint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraint, true
}

// SetConstraint sets field value
func (o *ParameterTypeRule) SetConstraint(v string) {
	o.Constraint = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ParameterTypeRule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterTypeRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ParameterTypeRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ParameterTypeRule) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterTypeRule) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *ParameterTypeRule) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o ParameterTypeRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterTypeRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["ledger_id"] = o.LedgerId
	toSerialize["parameter_type"] = o.ParameterType
	toSerialize["type"] = o.Type
	toSerialize["constraint"] = o.Constraint
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	return toSerialize, nil
}

type NullableParameterTypeRule struct {
	value *ParameterTypeRule
	isSet bool
}

func (v NullableParameterTypeRule) Get() *ParameterTypeRule {
	return v.value
}

func (v *NullableParameterTypeRule) Set(val *ParameterTypeRule) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterTypeRule) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterTypeRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterTypeRule(val *ParameterTypeRule) *NullableParameterTypeRule {
	return &NullableParameterTypeRule{value: val, isSet: true}
}

func (v NullableParameterTypeRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterTypeRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


