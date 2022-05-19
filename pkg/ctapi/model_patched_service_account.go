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
	"time"
)

// PatchedServiceAccount struct for PatchedServiceAccount
type PatchedServiceAccount struct {
	Url  *string `json:"url,omitempty"`
	Id   *string `json:"id,omitempty"`
	User *User   `json:"user,omitempty"`
	// An optional description of the process or system using the service account.
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ModifiedAt  *time.Time `json:"modified_at,omitempty"`
	// The most recent date and time the service account was used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
}

// NewPatchedServiceAccount instantiates a new PatchedServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedServiceAccount() *PatchedServiceAccount {
	this := PatchedServiceAccount{}
	return &this
}

// NewPatchedServiceAccountWithDefaults instantiates a new PatchedServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedServiceAccountWithDefaults() *PatchedServiceAccount {
	this := PatchedServiceAccount{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedServiceAccount) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedServiceAccount) SetId(v string) {
	o.Id = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *PatchedServiceAccount) SetUser(v User) {
	o.User = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedServiceAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedServiceAccount) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *PatchedServiceAccount) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceAccount) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *PatchedServiceAccount) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *PatchedServiceAccount) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

func (o PatchedServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.LastUsedAt != nil {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedServiceAccount struct {
	value *PatchedServiceAccount
	isSet bool
}

func (v NullablePatchedServiceAccount) Get() *PatchedServiceAccount {
	return v.value
}

func (v *NullablePatchedServiceAccount) Set(val *PatchedServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedServiceAccount(val *PatchedServiceAccount) *NullablePatchedServiceAccount {
	return &NullablePatchedServiceAccount{value: val, isSet: true}
}

func (v NullablePatchedServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
