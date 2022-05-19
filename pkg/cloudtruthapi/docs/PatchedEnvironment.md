# PatchedEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the environment. | [optional] [readonly] 
**Name** | Pointer to **string** | The environment name. | [optional] 
**Description** | Pointer to **string** | A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content. | [optional] 
**Parent** | Pointer to **NullableString** | Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed. | [optional] 
**Children** | Pointer to **[]string** | This is the opposite of &#x60;parent&#x60;, see that field for more details. | [optional] [readonly] 
**AccessControlled** | Pointer to **bool** | Indicates if access control is being enforced through grants. | [optional] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | Your role in the environment, if the environment is access-controlled. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedEnvironment

`func NewPatchedEnvironment() *PatchedEnvironment`

NewPatchedEnvironment instantiates a new PatchedEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEnvironmentWithDefaults

`func NewPatchedEnvironmentWithDefaults() *PatchedEnvironment`

NewPatchedEnvironmentWithDefaults instantiates a new PatchedEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedEnvironment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedEnvironment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedEnvironment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedEnvironment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedEnvironment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *PatchedEnvironment) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedEnvironment) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedEnvironment) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedEnvironment) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedEnvironment) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedEnvironment) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetChildren

`func (o *PatchedEnvironment) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PatchedEnvironment) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PatchedEnvironment) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PatchedEnvironment) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetAccessControlled

`func (o *PatchedEnvironment) GetAccessControlled() bool`

GetAccessControlled returns the AccessControlled field if non-nil, zero value otherwise.

### GetAccessControlledOk

`func (o *PatchedEnvironment) GetAccessControlledOk() (*bool, bool)`

GetAccessControlledOk returns a tuple with the AccessControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlled

`func (o *PatchedEnvironment) SetAccessControlled(v bool)`

SetAccessControlled sets AccessControlled field to given value.

### HasAccessControlled

`func (o *PatchedEnvironment) HasAccessControlled() bool`

HasAccessControlled returns a boolean if a field has been set.

### GetRole

`func (o *PatchedEnvironment) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedEnvironment) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedEnvironment) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedEnvironment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedEnvironment) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedEnvironment) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetCreatedAt

`func (o *PatchedEnvironment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedEnvironment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedEnvironment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedEnvironment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedEnvironment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedEnvironment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedEnvironment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedEnvironment) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


