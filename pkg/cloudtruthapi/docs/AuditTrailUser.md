# AuditTrailUser

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

### NewAuditTrailUser

`func NewAuditTrailUser(url string, id string, name NullableString, organizationName NullableString, membershipId NullableString, role NullableString, email NullableString, pictureUrl NullableString, createdAt time.Time, modifiedAt time.Time, ) *AuditTrailUser`

NewAuditTrailUser instantiates a new AuditTrailUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailUserWithDefaults

`func NewAuditTrailUserWithDefaults() *AuditTrailUser`

NewAuditTrailUserWithDefaults instantiates a new AuditTrailUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AuditTrailUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuditTrailUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuditTrailUser) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AuditTrailUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditTrailUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditTrailUser) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AuditTrailUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditTrailUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditTrailUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditTrailUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AuditTrailUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditTrailUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditTrailUser) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AuditTrailUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuditTrailUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationName

`func (o *AuditTrailUser) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AuditTrailUser) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AuditTrailUser) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *AuditTrailUser) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AuditTrailUser) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetMembershipId

`func (o *AuditTrailUser) GetMembershipId() string`

GetMembershipId returns the MembershipId field if non-nil, zero value otherwise.

### GetMembershipIdOk

`func (o *AuditTrailUser) GetMembershipIdOk() (*string, bool)`

GetMembershipIdOk returns a tuple with the MembershipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipId

`func (o *AuditTrailUser) SetMembershipId(v string)`

SetMembershipId sets MembershipId field to given value.


### SetMembershipIdNil

`func (o *AuditTrailUser) SetMembershipIdNil(b bool)`

 SetMembershipIdNil sets the value for MembershipId to be an explicit nil

### UnsetMembershipId
`func (o *AuditTrailUser) UnsetMembershipId()`

UnsetMembershipId ensures that no value is present for MembershipId, not even an explicit nil
### GetRole

`func (o *AuditTrailUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuditTrailUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuditTrailUser) SetRole(v string)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *AuditTrailUser) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AuditTrailUser) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetEmail

`func (o *AuditTrailUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuditTrailUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuditTrailUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AuditTrailUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AuditTrailUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPictureUrl

`func (o *AuditTrailUser) GetPictureUrl() string`

GetPictureUrl returns the PictureUrl field if non-nil, zero value otherwise.

### GetPictureUrlOk

`func (o *AuditTrailUser) GetPictureUrlOk() (*string, bool)`

GetPictureUrlOk returns a tuple with the PictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUrl

`func (o *AuditTrailUser) SetPictureUrl(v string)`

SetPictureUrl sets PictureUrl field to given value.


### SetPictureUrlNil

`func (o *AuditTrailUser) SetPictureUrlNil(b bool)`

 SetPictureUrlNil sets the value for PictureUrl to be an explicit nil

### UnsetPictureUrl
`func (o *AuditTrailUser) UnsetPictureUrl()`

UnsetPictureUrl ensures that no value is present for PictureUrl, not even an explicit nil
### GetCreatedAt

`func (o *AuditTrailUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditTrailUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditTrailUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AuditTrailUser) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AuditTrailUser) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AuditTrailUser) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


