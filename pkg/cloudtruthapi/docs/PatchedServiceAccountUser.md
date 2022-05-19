# PatchedServiceAccountUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | The unique identifier of a user. | 
**Type** | Pointer to **string** | The type of user record. | [optional] 
**Name** | **NullableString** |  | [readonly] 
**OrganizationName** | **NullableString** | The user&#39;s organization name. | [readonly] 
**MembershipId** | **NullableString** | Membership identifier for user. | [readonly] 
**Role** | **NullableString** | The user&#39;s role in the current organization (defined by the request authorization header). | [readonly] 
**Email** | **NullableString** |  | [readonly] 
**PictureUrl** | **NullableString** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewPatchedServiceAccountUser

`func NewPatchedServiceAccountUser(url string, id string, name NullableString, organizationName NullableString, membershipId NullableString, role NullableString, email NullableString, pictureUrl NullableString, createdAt time.Time, modifiedAt time.Time, ) *PatchedServiceAccountUser`

NewPatchedServiceAccountUser instantiates a new PatchedServiceAccountUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServiceAccountUserWithDefaults

`func NewPatchedServiceAccountUserWithDefaults() *PatchedServiceAccountUser`

NewPatchedServiceAccountUserWithDefaults instantiates a new PatchedServiceAccountUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedServiceAccountUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedServiceAccountUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedServiceAccountUser) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *PatchedServiceAccountUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedServiceAccountUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedServiceAccountUser) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PatchedServiceAccountUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedServiceAccountUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedServiceAccountUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedServiceAccountUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PatchedServiceAccountUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedServiceAccountUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedServiceAccountUser) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PatchedServiceAccountUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedServiceAccountUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationName

`func (o *PatchedServiceAccountUser) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *PatchedServiceAccountUser) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *PatchedServiceAccountUser) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *PatchedServiceAccountUser) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *PatchedServiceAccountUser) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetMembershipId

`func (o *PatchedServiceAccountUser) GetMembershipId() string`

GetMembershipId returns the MembershipId field if non-nil, zero value otherwise.

### GetMembershipIdOk

`func (o *PatchedServiceAccountUser) GetMembershipIdOk() (*string, bool)`

GetMembershipIdOk returns a tuple with the MembershipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipId

`func (o *PatchedServiceAccountUser) SetMembershipId(v string)`

SetMembershipId sets MembershipId field to given value.


### SetMembershipIdNil

`func (o *PatchedServiceAccountUser) SetMembershipIdNil(b bool)`

 SetMembershipIdNil sets the value for MembershipId to be an explicit nil

### UnsetMembershipId
`func (o *PatchedServiceAccountUser) UnsetMembershipId()`

UnsetMembershipId ensures that no value is present for MembershipId, not even an explicit nil
### GetRole

`func (o *PatchedServiceAccountUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedServiceAccountUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedServiceAccountUser) SetRole(v string)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *PatchedServiceAccountUser) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedServiceAccountUser) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetEmail

`func (o *PatchedServiceAccountUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedServiceAccountUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedServiceAccountUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *PatchedServiceAccountUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PatchedServiceAccountUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPictureUrl

`func (o *PatchedServiceAccountUser) GetPictureUrl() string`

GetPictureUrl returns the PictureUrl field if non-nil, zero value otherwise.

### GetPictureUrlOk

`func (o *PatchedServiceAccountUser) GetPictureUrlOk() (*string, bool)`

GetPictureUrlOk returns a tuple with the PictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUrl

`func (o *PatchedServiceAccountUser) SetPictureUrl(v string)`

SetPictureUrl sets PictureUrl field to given value.


### SetPictureUrlNil

`func (o *PatchedServiceAccountUser) SetPictureUrlNil(b bool)`

 SetPictureUrlNil sets the value for PictureUrl to be an explicit nil

### UnsetPictureUrl
`func (o *PatchedServiceAccountUser) UnsetPictureUrl()`

UnsetPictureUrl ensures that no value is present for PictureUrl, not even an explicit nil
### GetCreatedAt

`func (o *PatchedServiceAccountUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedServiceAccountUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedServiceAccountUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *PatchedServiceAccountUser) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedServiceAccountUser) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedServiceAccountUser) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


