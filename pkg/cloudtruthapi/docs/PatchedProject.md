# PatchedProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the project. | [optional] [readonly] 
**Name** | Pointer to **string** | The project name. | [optional] 
**Description** | Pointer to **string** | A description of the project.  You may find it helpful to document how this project is used to assist others when they need to maintain software that uses this content. | [optional] 
**Dependents** | Pointer to **[]string** | This is the opposite of &#x60;depends_on&#x60;, see that field for more details. | [optional] [readonly] 
**DependsOn** | Pointer to **NullableString** | Project dependencies allow projects to be used for shared configuration, for example a database used by many applications needs to advertise its port number.  Projects can depend on another project which will add the parameters from the parent project into the current project.  All of the parameter names between the two projects must be unique.  When retrieving values or rendering templates, all of the parameters from the parent project will also be available in the current project. | [optional] 
**AccessControlled** | Pointer to **bool** | Indicates if access control is being enforced through grants. | [optional] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | Your role in the project, if the project is access-controlled. | [optional] [readonly] 
**Pushes** | Pointer to [**[]AwsPush**](AwsPush.md) | Deprecated. Only shows pushes for aws integrations in /api/v1/. | [optional] [readonly] 
**PushUrls** | Pointer to **[]string** | Push actions associated with the project. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedProject

`func NewPatchedProject() *PatchedProject`

NewPatchedProject instantiates a new PatchedProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProjectWithDefaults

`func NewPatchedProjectWithDefaults() *PatchedProject`

NewPatchedProjectWithDefaults instantiates a new PatchedProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedProject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedProject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedProject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedProject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDependents

`func (o *PatchedProject) GetDependents() []string`

GetDependents returns the Dependents field if non-nil, zero value otherwise.

### GetDependentsOk

`func (o *PatchedProject) GetDependentsOk() (*[]string, bool)`

GetDependentsOk returns a tuple with the Dependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependents

`func (o *PatchedProject) SetDependents(v []string)`

SetDependents sets Dependents field to given value.

### HasDependents

`func (o *PatchedProject) HasDependents() bool`

HasDependents returns a boolean if a field has been set.

### GetDependsOn

`func (o *PatchedProject) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *PatchedProject) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *PatchedProject) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *PatchedProject) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *PatchedProject) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *PatchedProject) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetAccessControlled

`func (o *PatchedProject) GetAccessControlled() bool`

GetAccessControlled returns the AccessControlled field if non-nil, zero value otherwise.

### GetAccessControlledOk

`func (o *PatchedProject) GetAccessControlledOk() (*bool, bool)`

GetAccessControlledOk returns a tuple with the AccessControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlled

`func (o *PatchedProject) SetAccessControlled(v bool)`

SetAccessControlled sets AccessControlled field to given value.

### HasAccessControlled

`func (o *PatchedProject) HasAccessControlled() bool`

HasAccessControlled returns a boolean if a field has been set.

### GetRole

`func (o *PatchedProject) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedProject) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedProject) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedProject) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedProject) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedProject) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPushes

`func (o *PatchedProject) GetPushes() []AwsPush`

GetPushes returns the Pushes field if non-nil, zero value otherwise.

### GetPushesOk

`func (o *PatchedProject) GetPushesOk() (*[]AwsPush, bool)`

GetPushesOk returns a tuple with the Pushes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushes

`func (o *PatchedProject) SetPushes(v []AwsPush)`

SetPushes sets Pushes field to given value.

### HasPushes

`func (o *PatchedProject) HasPushes() bool`

HasPushes returns a boolean if a field has been set.

### GetPushUrls

`func (o *PatchedProject) GetPushUrls() []string`

GetPushUrls returns the PushUrls field if non-nil, zero value otherwise.

### GetPushUrlsOk

`func (o *PatchedProject) GetPushUrlsOk() (*[]string, bool)`

GetPushUrlsOk returns a tuple with the PushUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushUrls

`func (o *PatchedProject) SetPushUrls(v []string)`

SetPushUrls sets PushUrls field to given value.

### HasPushUrls

`func (o *PatchedProject) HasPushUrls() bool`

HasPushUrls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedProject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedProject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedProject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedProject) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


