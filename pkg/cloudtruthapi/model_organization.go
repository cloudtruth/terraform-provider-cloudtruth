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

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	Url string `json:"url"`
	// A unique identifier for the organization.
	Id string `json:"id"`
	// The organization name.
	Name string `json:"name"`
	// Indicates if this Organization is the one currently targeted by the Bearer token used by the client to authorize.
	Current bool `json:"current"`
	Role RoleEnum `json:"role"`
	SubscriptionExpiresAt NullableTime `json:"subscription_expires_at"`
	SubscriptionFeatures []string `json:"subscription_features"`
	SubscriptionId NullableString `json:"subscription_id"`
	SubscriptionPlanId NullableString `json:"subscription_plan_id"`
	SubscriptionPlanName NullableString `json:"subscription_plan_name"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(url string, id string, name string, current bool, role RoleEnum, subscriptionExpiresAt NullableTime, subscriptionFeatures []string, subscriptionId NullableString, subscriptionPlanId NullableString, subscriptionPlanName NullableString, createdAt time.Time, modifiedAt time.Time) *Organization {
	this := Organization{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Current = current
	this.Role = role
	this.SubscriptionExpiresAt = subscriptionExpiresAt
	this.SubscriptionFeatures = subscriptionFeatures
	this.SubscriptionId = subscriptionId
	this.SubscriptionPlanId = subscriptionPlanId
	this.SubscriptionPlanName = subscriptionPlanName
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetUrl returns the Url field value
func (o *Organization) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Organization) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Organization) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Organization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Organization) SetName(v string) {
	o.Name = v
}

// GetCurrent returns the Current field value
func (o *Organization) GetCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *Organization) GetCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *Organization) SetCurrent(v bool) {
	o.Current = v
}

// GetRole returns the Role field value
func (o *Organization) GetRole() RoleEnum {
	if o == nil {
		var ret RoleEnum
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Organization) GetRoleOk() (*RoleEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Organization) SetRole(v RoleEnum) {
	o.Role = v
}

// GetSubscriptionExpiresAt returns the SubscriptionExpiresAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Organization) GetSubscriptionExpiresAt() time.Time {
	if o == nil || o.SubscriptionExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.SubscriptionExpiresAt.Get()
}

// GetSubscriptionExpiresAtOk returns a tuple with the SubscriptionExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetSubscriptionExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionExpiresAt.Get(), o.SubscriptionExpiresAt.IsSet()
}

// SetSubscriptionExpiresAt sets field value
func (o *Organization) SetSubscriptionExpiresAt(v time.Time) {
	o.SubscriptionExpiresAt.Set(&v)
}

// GetSubscriptionFeatures returns the SubscriptionFeatures field value
func (o *Organization) GetSubscriptionFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubscriptionFeatures
}

// GetSubscriptionFeaturesOk returns a tuple with the SubscriptionFeatures field value
// and a boolean to check if the value has been set.
func (o *Organization) GetSubscriptionFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionFeatures, true
}

// SetSubscriptionFeatures sets field value
func (o *Organization) SetSubscriptionFeatures(v []string) {
	o.SubscriptionFeatures = v
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Organization) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *Organization) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetSubscriptionPlanId returns the SubscriptionPlanId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Organization) GetSubscriptionPlanId() string {
	if o == nil || o.SubscriptionPlanId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionPlanId.Get()
}

// GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetSubscriptionPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionPlanId.Get(), o.SubscriptionPlanId.IsSet()
}

// SetSubscriptionPlanId sets field value
func (o *Organization) SetSubscriptionPlanId(v string) {
	o.SubscriptionPlanId.Set(&v)
}

// GetSubscriptionPlanName returns the SubscriptionPlanName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Organization) GetSubscriptionPlanName() string {
	if o == nil || o.SubscriptionPlanName.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionPlanName.Get()
}

// GetSubscriptionPlanNameOk returns a tuple with the SubscriptionPlanName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetSubscriptionPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionPlanName.Get(), o.SubscriptionPlanName.IsSet()
}

// SetSubscriptionPlanName sets field value
func (o *Organization) SetSubscriptionPlanName(v string) {
	o.SubscriptionPlanName.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *Organization) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Organization) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Organization) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Organization) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Organization) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["current"] = o.Current
	toSerialize["role"] = o.Role
	toSerialize["subscription_expires_at"] = o.SubscriptionExpiresAt.Get()
	toSerialize["subscription_features"] = o.SubscriptionFeatures
	toSerialize["subscription_id"] = o.SubscriptionId.Get()
	toSerialize["subscription_plan_id"] = o.SubscriptionPlanId.Get()
	toSerialize["subscription_plan_name"] = o.SubscriptionPlanName.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt
	return toSerialize, nil
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


