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

// checks if the ValueUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueUpdate{}

// ValueUpdate Unlike other UpdateSerializers, we do not inherit from the CreateSerializer here because `environment` is not a required field for updates.
type ValueUpdate struct {
	Id string `json:"id"`
	// An external parameter leverages a CloudTruth integration to retrieve content on-demand from an external source.  When this is `false` the value is stored by CloudTruth and considered to be _internal_.  When this is `true`, the `external_fqn` field must be set.
	External *bool `json:"external,omitempty"`
	// The FQN, or Fully-Qualified Name, is the path through the integration to get to the desired content.  This must be present and reference a valid integration when the value is `external`.
	ExternalFqn *string `json:"external_fqn,omitempty"`
	// If the value is `external`, the content returned by the integration can be reduced by applying a JMESpath expression.  This is valid as long as the content is structured and of a supported format.  JMESpath expressions are supported on `json`, `yaml`, and `dotenv` content.
	ExternalFilter *string `json:"external_filter,omitempty"`
	// This is the content to use when resolving the Value for an internal non-secret, or when storing a secret.  This field cannot be specified when creating or updating an `external` value.
	InternalValue NullableString `json:"internal_value,omitempty"`
	// If `true`, apply template substitution rules to this value.  If `false`, this value is a literal value.  Note: secrets cannot be interpolated.
	Interpolated *bool `json:"interpolated,omitempty"`
	// Indicates the value content is a secret.  Normally this is `true` when the parameter is a secret. It is possible for a parameter to be a secret with a external value that is not a secret.  It is not possible to convert a parameter from a secret to a non-secret if any of the values are external and a secret.  Clients can check this condition by leveraging this field.  It is also possible for a parameter to not be a secret but for this value to be dynamic and reference a Parameter that is a secret.  In this case, we indicate the value is a secret.
	Secret NullableBool `json:"secret"`
	// This is the actual content of the Value for the given parameter in the given environment.  If you request secret masking, no secret content will be included in the result and instead a series of asterisks will be used instead for the value.  Clients applying this value to a shell environment should set `<parameter_name>=<value>` even if `value` is the empty string.  If `value` is `null`, the client should unset that shell environment variable.
	Value NullableString `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

type _ValueUpdate ValueUpdate

// NewValueUpdate instantiates a new ValueUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueUpdate(id string, secret NullableBool, value NullableString, createdAt time.Time, modifiedAt NullableTime) *ValueUpdate {
	this := ValueUpdate{}
	this.Id = id
	this.Secret = secret
	this.Value = value
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewValueUpdateWithDefaults instantiates a new ValueUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueUpdateWithDefaults() *ValueUpdate {
	this := ValueUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *ValueUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ValueUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ValueUpdate) SetId(v string) {
	o.Id = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ValueUpdate) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueUpdate) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *ValueUpdate) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *ValueUpdate) SetExternal(v bool) {
	o.External = &v
}

// GetExternalFqn returns the ExternalFqn field value if set, zero value otherwise.
func (o *ValueUpdate) GetExternalFqn() string {
	if o == nil || IsNil(o.ExternalFqn) {
		var ret string
		return ret
	}
	return *o.ExternalFqn
}

// GetExternalFqnOk returns a tuple with the ExternalFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueUpdate) GetExternalFqnOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalFqn) {
		return nil, false
	}
	return o.ExternalFqn, true
}

// HasExternalFqn returns a boolean if a field has been set.
func (o *ValueUpdate) HasExternalFqn() bool {
	if o != nil && !IsNil(o.ExternalFqn) {
		return true
	}

	return false
}

// SetExternalFqn gets a reference to the given string and assigns it to the ExternalFqn field.
func (o *ValueUpdate) SetExternalFqn(v string) {
	o.ExternalFqn = &v
}

// GetExternalFilter returns the ExternalFilter field value if set, zero value otherwise.
func (o *ValueUpdate) GetExternalFilter() string {
	if o == nil || IsNil(o.ExternalFilter) {
		var ret string
		return ret
	}
	return *o.ExternalFilter
}

// GetExternalFilterOk returns a tuple with the ExternalFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueUpdate) GetExternalFilterOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalFilter) {
		return nil, false
	}
	return o.ExternalFilter, true
}

// HasExternalFilter returns a boolean if a field has been set.
func (o *ValueUpdate) HasExternalFilter() bool {
	if o != nil && !IsNil(o.ExternalFilter) {
		return true
	}

	return false
}

// SetExternalFilter gets a reference to the given string and assigns it to the ExternalFilter field.
func (o *ValueUpdate) SetExternalFilter(v string) {
	o.ExternalFilter = &v
}

// GetInternalValue returns the InternalValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValueUpdate) GetInternalValue() string {
	if o == nil || IsNil(o.InternalValue.Get()) {
		var ret string
		return ret
	}
	return *o.InternalValue.Get()
}

// GetInternalValueOk returns a tuple with the InternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueUpdate) GetInternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalValue.Get(), o.InternalValue.IsSet()
}

// HasInternalValue returns a boolean if a field has been set.
func (o *ValueUpdate) HasInternalValue() bool {
	if o != nil && o.InternalValue.IsSet() {
		return true
	}

	return false
}

// SetInternalValue gets a reference to the given NullableString and assigns it to the InternalValue field.
func (o *ValueUpdate) SetInternalValue(v string) {
	o.InternalValue.Set(&v)
}
// SetInternalValueNil sets the value for InternalValue to be an explicit nil
func (o *ValueUpdate) SetInternalValueNil() {
	o.InternalValue.Set(nil)
}

// UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
func (o *ValueUpdate) UnsetInternalValue() {
	o.InternalValue.Unset()
}

// GetInterpolated returns the Interpolated field value if set, zero value otherwise.
func (o *ValueUpdate) GetInterpolated() bool {
	if o == nil || IsNil(o.Interpolated) {
		var ret bool
		return ret
	}
	return *o.Interpolated
}

// GetInterpolatedOk returns a tuple with the Interpolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueUpdate) GetInterpolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Interpolated) {
		return nil, false
	}
	return o.Interpolated, true
}

// HasInterpolated returns a boolean if a field has been set.
func (o *ValueUpdate) HasInterpolated() bool {
	if o != nil && !IsNil(o.Interpolated) {
		return true
	}

	return false
}

// SetInterpolated gets a reference to the given bool and assigns it to the Interpolated field.
func (o *ValueUpdate) SetInterpolated(v bool) {
	o.Interpolated = &v
}

// GetSecret returns the Secret field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ValueUpdate) GetSecret() bool {
	if o == nil || o.Secret.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueUpdate) GetSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// SetSecret sets field value
func (o *ValueUpdate) SetSecret(v bool) {
	o.Secret.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ValueUpdate) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueUpdate) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *ValueUpdate) SetValue(v string) {
	o.Value.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ValueUpdate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ValueUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ValueUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ValueUpdate) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueUpdate) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *ValueUpdate) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o ValueUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.ExternalFqn) {
		toSerialize["external_fqn"] = o.ExternalFqn
	}
	if !IsNil(o.ExternalFilter) {
		toSerialize["external_filter"] = o.ExternalFilter
	}
	if o.InternalValue.IsSet() {
		toSerialize["internal_value"] = o.InternalValue.Get()
	}
	if !IsNil(o.Interpolated) {
		toSerialize["interpolated"] = o.Interpolated
	}
	toSerialize["secret"] = o.Secret.Get()
	toSerialize["value"] = o.Value.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	return toSerialize, nil
}

func (o *ValueUpdate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"secret",
		"value",
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

	varValueUpdate := _ValueUpdate{}

	err = json.Unmarshal(bytes, &varValueUpdate)

	if err != nil {
		return err
	}

	*o = ValueUpdate(varValueUpdate)

	return err
}

type NullableValueUpdate struct {
	value *ValueUpdate
	isSet bool
}

func (v NullableValueUpdate) Get() *ValueUpdate {
	return v.value
}

func (v *NullableValueUpdate) Set(val *ValueUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableValueUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableValueUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueUpdate(val *ValueUpdate) *NullableValueUpdate {
	return &NullableValueUpdate{value: val, isSet: true}
}

func (v NullableValueUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


