# PatchedInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier of an invitation. | [optional] [readonly] 
**Email** | Pointer to **string** | The email address of the user to be invited. | [optional] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | The role that the user will have in the organization, should the user accept. | [optional] 
**Inviter** | Pointer to **string** | The user that created the invitation. | [optional] [readonly] 
**InviterName** | Pointer to **string** | The name of the user that created the invitation. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the invitation. | [optional] [readonly] 
**StateDetail** | Pointer to **string** | Additional details about the state of the invitation. | [optional] [readonly] 
**Membership** | Pointer to **NullableString** | The resulting membership, should the user accept. | [optional] [readonly] 
**Organization** | Pointer to **string** | The organization that the user will become a member of, should the user accept. | [optional] [readonly] 

## Methods

### NewPatchedInvitation

`func NewPatchedInvitation() *PatchedInvitation`

NewPatchedInvitation instantiates a new PatchedInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInvitationWithDefaults

`func NewPatchedInvitationWithDefaults() *PatchedInvitation`

NewPatchedInvitationWithDefaults instantiates a new PatchedInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedInvitation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedInvitation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedInvitation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedInvitation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRole

`func (o *PatchedInvitation) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedInvitation) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedInvitation) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedInvitation) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedInvitation) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedInvitation) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetInviter

`func (o *PatchedInvitation) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *PatchedInvitation) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *PatchedInvitation) SetInviter(v string)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *PatchedInvitation) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetInviterName

`func (o *PatchedInvitation) GetInviterName() string`

GetInviterName returns the InviterName field if non-nil, zero value otherwise.

### GetInviterNameOk

`func (o *PatchedInvitation) GetInviterNameOk() (*string, bool)`

GetInviterNameOk returns a tuple with the InviterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterName

`func (o *PatchedInvitation) SetInviterName(v string)`

SetInviterName sets InviterName field to given value.

### HasInviterName

`func (o *PatchedInvitation) HasInviterName() bool`

HasInviterName returns a boolean if a field has been set.

### GetState

`func (o *PatchedInvitation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchedInvitation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchedInvitation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatchedInvitation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDetail

`func (o *PatchedInvitation) GetStateDetail() string`

GetStateDetail returns the StateDetail field if non-nil, zero value otherwise.

### GetStateDetailOk

`func (o *PatchedInvitation) GetStateDetailOk() (*string, bool)`

GetStateDetailOk returns a tuple with the StateDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDetail

`func (o *PatchedInvitation) SetStateDetail(v string)`

SetStateDetail sets StateDetail field to given value.

### HasStateDetail

`func (o *PatchedInvitation) HasStateDetail() bool`

HasStateDetail returns a boolean if a field has been set.

### GetMembership

`func (o *PatchedInvitation) GetMembership() string`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *PatchedInvitation) GetMembershipOk() (*string, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *PatchedInvitation) SetMembership(v string)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *PatchedInvitation) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### SetMembershipNil

`func (o *PatchedInvitation) SetMembershipNil(b bool)`

 SetMembershipNil sets the value for Membership to be an explicit nil

### UnsetMembership
`func (o *PatchedInvitation) UnsetMembership()`

UnsetMembership ensures that no value is present for Membership, not even an explicit nil
### GetOrganization

`func (o *PatchedInvitation) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PatchedInvitation) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PatchedInvitation) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PatchedInvitation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


