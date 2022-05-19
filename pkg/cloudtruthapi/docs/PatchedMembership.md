# PatchedMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the membership. | [optional] [readonly] 
**User** | Pointer to **string** | The user of the membership. | [optional] 
**Organization** | Pointer to **string** | The organization that the user is a member of. | [optional] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | The role that the user has in the organization. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedMembership

`func NewPatchedMembership() *PatchedMembership`

NewPatchedMembership instantiates a new PatchedMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedMembershipWithDefaults

`func NewPatchedMembershipWithDefaults() *PatchedMembership`

NewPatchedMembershipWithDefaults instantiates a new PatchedMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedMembership) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedMembership) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedMembership) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedMembership) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedMembership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedMembership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedMembership) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedMembership) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedMembership) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedMembership) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedMembership) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedMembership) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganization

`func (o *PatchedMembership) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PatchedMembership) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PatchedMembership) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PatchedMembership) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRole

`func (o *PatchedMembership) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedMembership) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedMembership) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedMembership) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedMembership) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedMembership) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetCreatedAt

`func (o *PatchedMembership) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedMembership) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedMembership) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedMembership) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedMembership) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedMembership) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedMembership) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedMembership) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


