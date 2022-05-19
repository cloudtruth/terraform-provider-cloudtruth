# PatchedOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the organization. | [optional] [readonly] 
**Name** | Pointer to **string** | The organization name. | [optional] 
**Current** | Pointer to **bool** | Indicates if this Organization is the one currently targeted by the Bearer token used by the client to authorize. | [optional] [readonly] 
**Role** | Pointer to [**NullableRoleEnum**](RoleEnum.md) | Your role in the organization. | [optional] [readonly] 
**SubscriptionExpiresAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**SubscriptionFeatures** | Pointer to **[]string** |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] [readonly] 
**SubscriptionPlanId** | Pointer to **NullableString** |  | [optional] [readonly] 
**SubscriptionPlanName** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedOrganization

`func NewPatchedOrganization() *PatchedOrganization`

NewPatchedOrganization instantiates a new PatchedOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOrganizationWithDefaults

`func NewPatchedOrganizationWithDefaults() *PatchedOrganization`

NewPatchedOrganizationWithDefaults instantiates a new PatchedOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedOrganization) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedOrganization) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedOrganization) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedOrganization) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrent

`func (o *PatchedOrganization) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *PatchedOrganization) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *PatchedOrganization) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *PatchedOrganization) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetRole

`func (o *PatchedOrganization) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedOrganization) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedOrganization) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedOrganization) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedOrganization) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedOrganization) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSubscriptionExpiresAt

`func (o *PatchedOrganization) GetSubscriptionExpiresAt() time.Time`

GetSubscriptionExpiresAt returns the SubscriptionExpiresAt field if non-nil, zero value otherwise.

### GetSubscriptionExpiresAtOk

`func (o *PatchedOrganization) GetSubscriptionExpiresAtOk() (*time.Time, bool)`

GetSubscriptionExpiresAtOk returns a tuple with the SubscriptionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpiresAt

`func (o *PatchedOrganization) SetSubscriptionExpiresAt(v time.Time)`

SetSubscriptionExpiresAt sets SubscriptionExpiresAt field to given value.

### HasSubscriptionExpiresAt

`func (o *PatchedOrganization) HasSubscriptionExpiresAt() bool`

HasSubscriptionExpiresAt returns a boolean if a field has been set.

### SetSubscriptionExpiresAtNil

`func (o *PatchedOrganization) SetSubscriptionExpiresAtNil(b bool)`

 SetSubscriptionExpiresAtNil sets the value for SubscriptionExpiresAt to be an explicit nil

### UnsetSubscriptionExpiresAt
`func (o *PatchedOrganization) UnsetSubscriptionExpiresAt()`

UnsetSubscriptionExpiresAt ensures that no value is present for SubscriptionExpiresAt, not even an explicit nil
### GetSubscriptionFeatures

`func (o *PatchedOrganization) GetSubscriptionFeatures() []string`

GetSubscriptionFeatures returns the SubscriptionFeatures field if non-nil, zero value otherwise.

### GetSubscriptionFeaturesOk

`func (o *PatchedOrganization) GetSubscriptionFeaturesOk() (*[]string, bool)`

GetSubscriptionFeaturesOk returns a tuple with the SubscriptionFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFeatures

`func (o *PatchedOrganization) SetSubscriptionFeatures(v []string)`

SetSubscriptionFeatures sets SubscriptionFeatures field to given value.

### HasSubscriptionFeatures

`func (o *PatchedOrganization) HasSubscriptionFeatures() bool`

HasSubscriptionFeatures returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PatchedOrganization) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PatchedOrganization) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PatchedOrganization) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *PatchedOrganization) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *PatchedOrganization) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *PatchedOrganization) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionPlanId

`func (o *PatchedOrganization) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *PatchedOrganization) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *PatchedOrganization) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.

### HasSubscriptionPlanId

`func (o *PatchedOrganization) HasSubscriptionPlanId() bool`

HasSubscriptionPlanId returns a boolean if a field has been set.

### SetSubscriptionPlanIdNil

`func (o *PatchedOrganization) SetSubscriptionPlanIdNil(b bool)`

 SetSubscriptionPlanIdNil sets the value for SubscriptionPlanId to be an explicit nil

### UnsetSubscriptionPlanId
`func (o *PatchedOrganization) UnsetSubscriptionPlanId()`

UnsetSubscriptionPlanId ensures that no value is present for SubscriptionPlanId, not even an explicit nil
### GetSubscriptionPlanName

`func (o *PatchedOrganization) GetSubscriptionPlanName() string`

GetSubscriptionPlanName returns the SubscriptionPlanName field if non-nil, zero value otherwise.

### GetSubscriptionPlanNameOk

`func (o *PatchedOrganization) GetSubscriptionPlanNameOk() (*string, bool)`

GetSubscriptionPlanNameOk returns a tuple with the SubscriptionPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanName

`func (o *PatchedOrganization) SetSubscriptionPlanName(v string)`

SetSubscriptionPlanName sets SubscriptionPlanName field to given value.

### HasSubscriptionPlanName

`func (o *PatchedOrganization) HasSubscriptionPlanName() bool`

HasSubscriptionPlanName returns a boolean if a field has been set.

### SetSubscriptionPlanNameNil

`func (o *PatchedOrganization) SetSubscriptionPlanNameNil(b bool)`

 SetSubscriptionPlanNameNil sets the value for SubscriptionPlanName to be an explicit nil

### UnsetSubscriptionPlanName
`func (o *PatchedOrganization) UnsetSubscriptionPlanName()`

UnsetSubscriptionPlanName ensures that no value is present for SubscriptionPlanName, not even an explicit nil
### GetCreatedAt

`func (o *PatchedOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedOrganization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedOrganization) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedOrganization) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedOrganization) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedOrganization) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


