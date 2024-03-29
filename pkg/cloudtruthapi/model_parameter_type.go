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

// checks if the ParameterType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterType{}

// ParameterType struct for ParameterType
type ParameterType struct {
	// The URL for the parameter type.
	Url string `json:"url"`
	Id string `json:"id"`
	LedgerId string `json:"ledger_id"`
	// The parameter type name.
	Name string `json:"name"`
	// A description of the parameter type, provide documentation on how to use this type here.
	Description *string `json:"description,omitempty"`
	// Rules applied to this parameter.
	Rules []ParameterTypeRule `json:"rules"`
	// The URL for this parameter type's parent
	Parent NullableString `json:"parent"`
	// Name of the parent ParameterType (if any).
	ParentName NullableString `json:"parent_name"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewParameterType instantiates a new ParameterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterType(url string, id string, ledgerId string, name string, rules []ParameterTypeRule, parent NullableString, parentName NullableString, createdAt time.Time, modifiedAt NullableTime) *ParameterType {
	this := ParameterType{}
	this.Url = url
	this.Id = id
	this.LedgerId = ledgerId
	this.Name = name
	this.Rules = rules
	this.Parent = parent
	this.ParentName = parentName
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewParameterTypeWithDefaults instantiates a new ParameterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterTypeWithDefaults() *ParameterType {
	this := ParameterType{}
	return &this
}

// GetUrl returns the Url field value
func (o *ParameterType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ParameterType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ParameterType) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *ParameterType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterType) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterType) SetId(v string) {
	o.Id = v
}

// GetLedgerId returns the LedgerId field value
func (o *ParameterType) GetLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value
// and a boolean to check if the value has been set.
func (o *ParameterType) GetLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerId, true
}

// SetLedgerId sets field value
func (o *ParameterType) SetLedgerId(v string) {
	o.LedgerId = v
}

// GetName returns the Name field value
func (o *ParameterType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterType) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterType) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value
func (o *ParameterType) GetRules() []ParameterTypeRule {
	if o == nil {
		var ret []ParameterTypeRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ParameterType) GetRulesOk() ([]ParameterTypeRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *ParameterType) SetRules(v []ParameterTypeRule) {
	o.Rules = v
}

// GetParent returns the Parent field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ParameterType) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}

	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterType) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// SetParent sets field value
func (o *ParameterType) SetParent(v string) {
	o.Parent.Set(&v)
}

// GetParentName returns the ParentName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ParameterType) GetParentName() string {
	if o == nil || o.ParentName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentName.Get()
}

// GetParentNameOk returns a tuple with the ParentName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterType) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentName.Get(), o.ParentName.IsSet()
}

// SetParentName sets field value
func (o *ParameterType) SetParentName(v string) {
	o.ParentName.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ParameterType) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ParameterType) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ParameterType) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParameterType) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *ParameterType) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o ParameterType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["ledger_id"] = o.LedgerId
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["rules"] = o.Rules
	toSerialize["parent"] = o.Parent.Get()
	toSerialize["parent_name"] = o.ParentName.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	return toSerialize, nil
}

type NullableParameterType struct {
	value *ParameterType
	isSet bool
}

func (v NullableParameterType) Get() *ParameterType {
	return v.value
}

func (v *NullableParameterType) Set(val *ParameterType) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterType) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterType(val *ParameterType) *NullableParameterType {
	return &NullableParameterType{value: val, isSet: true}
}

func (v NullableParameterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


