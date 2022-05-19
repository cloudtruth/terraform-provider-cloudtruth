# MembershipCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** | The user of the membership. | 
**Role** | [**NullableRoleEnum**](RoleEnum.md) | The role that the user has in the organization. | 

## Methods

### NewMembershipCreate

`func NewMembershipCreate(user string, role NullableRoleEnum, ) *MembershipCreate`

NewMembershipCreate instantiates a new MembershipCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipCreateWithDefaults

`func NewMembershipCreateWithDefaults() *MembershipCreate`

NewMembershipCreateWithDefaults instantiates a new MembershipCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *MembershipCreate) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MembershipCreate) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MembershipCreate) SetUser(v string)`

SetUser sets User field to given value.


### GetRole

`func (o *MembershipCreate) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MembershipCreate) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MembershipCreate) SetRole(v RoleEnum)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *MembershipCreate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *MembershipCreate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


