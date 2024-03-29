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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Url string `json:"url"`
	// The unique identifier of a user.
	Id string `json:"id"`
	// The type of user record.
	Type *string `json:"type,omitempty"`
	Name NullableString `json:"name"`
	// The user's organization name.
	OrganizationName NullableString `json:"organization_name"`
	// Membership identifier for user.
	MembershipId NullableString `json:"membership_id"`
	// The user's role in the current organization (defined by the request authorization header).
	Role NullableString `json:"role"`
	Email NullableString `json:"email"`
	PictureUrl NullableString `json:"picture_url"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(url string, id string, name NullableString, organizationName NullableString, membershipId NullableString, role NullableString, email NullableString, pictureUrl NullableString, createdAt time.Time, modifiedAt NullableTime) *User {
	this := User{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.OrganizationName = organizationName
	this.MembershipId = membershipId
	this.Role = role
	this.Email = email
	this.PictureUrl = pictureUrl
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetUrl returns the Url field value
func (o *User) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *User) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *User) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *User) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *User) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *User) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *User) SetName(v string) {
	o.Name.Set(&v)
}

// GetOrganizationName returns the OrganizationName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetOrganizationName() string {
	if o == nil || o.OrganizationName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// SetOrganizationName sets field value
func (o *User) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// GetMembershipId returns the MembershipId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetMembershipId() string {
	if o == nil || o.MembershipId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MembershipId.Get()
}

// GetMembershipIdOk returns a tuple with the MembershipId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMembershipIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MembershipId.Get(), o.MembershipId.IsSet()
}

// SetMembershipId sets field value
func (o *User) SetMembershipId(v string) {
	o.MembershipId.Set(&v)
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetRole() string {
	if o == nil || o.Role.Get() == nil {
		var ret string
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *User) SetRole(v string) {
	o.Role.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetPictureUrl returns the PictureUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetPictureUrl() string {
	if o == nil || o.PictureUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PictureUrl.Get()
}

// GetPictureUrlOk returns a tuple with the PictureUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPictureUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PictureUrl.Get(), o.PictureUrl.IsSet()
}

// SetPictureUrl sets field value
func (o *User) SetPictureUrl(v string) {
	o.PictureUrl.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *User) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *User) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *User) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name.Get()
	toSerialize["organization_name"] = o.OrganizationName.Get()
	toSerialize["membership_id"] = o.MembershipId.Get()
	toSerialize["role"] = o.Role.Get()
	toSerialize["email"] = o.Email.Get()
	toSerialize["picture_url"] = o.PictureUrl.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


