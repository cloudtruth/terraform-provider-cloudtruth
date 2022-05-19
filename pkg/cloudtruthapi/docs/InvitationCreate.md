# InvitationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the user to be invited. | 
**Role** | [**NullableRoleEnum**](RoleEnum.md) | The role that the user will have in the organization, should the user accept. | 

## Methods

### NewInvitationCreate

`func NewInvitationCreate(email string, role NullableRoleEnum, ) *InvitationCreate`

NewInvitationCreate instantiates a new InvitationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationCreateWithDefaults

`func NewInvitationCreateWithDefaults() *InvitationCreate`

NewInvitationCreateWithDefaults instantiates a new InvitationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InvitationCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InvitationCreate) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InvitationCreate) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InvitationCreate) SetRole(v RoleEnum)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *InvitationCreate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *InvitationCreate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


