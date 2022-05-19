# Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | The unique identifier of an invitation. | [readonly] 
**Email** | **string** | The email address of the user to be invited. | 
**Role** | [**NullableRoleEnum**](RoleEnum.md) | The role that the user will have in the organization, should the user accept. | 
**Inviter** | **string** | The user that created the invitation. | [readonly] 
**InviterName** | **string** | The name of the user that created the invitation. | [readonly] 
**State** | **string** | The current state of the invitation. | [readonly] 
**StateDetail** | **string** | Additional details about the state of the invitation. | [readonly] 
**Membership** | **NullableString** | The resulting membership, should the user accept. | [readonly] 
**Organization** | **string** | The organization that the user will become a member of, should the user accept. | [readonly] 

## Methods

### NewInvitation

`func NewInvitation(url string, id string, email string, role NullableRoleEnum, inviter string, inviterName string, state string, stateDetail string, membership NullableString, organization string, ) *Invitation`

NewInvitation instantiates a new Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationWithDefaults

`func NewInvitationWithDefaults() *Invitation`

NewInvitationWithDefaults instantiates a new Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Invitation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Invitation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Invitation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Invitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invitation) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *Invitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Invitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Invitation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *Invitation) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Invitation) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Invitation) SetRole(v RoleEnum)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *Invitation) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Invitation) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetInviter

`func (o *Invitation) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *Invitation) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *Invitation) SetInviter(v string)`

SetInviter sets Inviter field to given value.


### GetInviterName

`func (o *Invitation) GetInviterName() string`

GetInviterName returns the InviterName field if non-nil, zero value otherwise.

### GetInviterNameOk

`func (o *Invitation) GetInviterNameOk() (*string, bool)`

GetInviterNameOk returns a tuple with the InviterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterName

`func (o *Invitation) SetInviterName(v string)`

SetInviterName sets InviterName field to given value.


### GetState

`func (o *Invitation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Invitation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Invitation) SetState(v string)`

SetState sets State field to given value.


### GetStateDetail

`func (o *Invitation) GetStateDetail() string`

GetStateDetail returns the StateDetail field if non-nil, zero value otherwise.

### GetStateDetailOk

`func (o *Invitation) GetStateDetailOk() (*string, bool)`

GetStateDetailOk returns a tuple with the StateDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDetail

`func (o *Invitation) SetStateDetail(v string)`

SetStateDetail sets StateDetail field to given value.


### GetMembership

`func (o *Invitation) GetMembership() string`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *Invitation) GetMembershipOk() (*string, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *Invitation) SetMembership(v string)`

SetMembership sets Membership field to given value.


### SetMembershipNil

`func (o *Invitation) SetMembershipNil(b bool)`

 SetMembershipNil sets the value for Membership to be an explicit nil

### UnsetMembership
`func (o *Invitation) UnsetMembership()`

UnsetMembership ensures that no value is present for Membership, not even an explicit nil
### GetOrganization

`func (o *Invitation) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Invitation) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Invitation) SetOrganization(v string)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


