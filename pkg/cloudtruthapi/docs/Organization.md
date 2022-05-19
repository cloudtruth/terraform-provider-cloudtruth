# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the organization. | [readonly] 
**Name** | **string** | The organization name. | 
**Current** | **bool** | Indicates if this Organization is the one currently targeted by the Bearer token used by the client to authorize. | [readonly] 
**Role** | [**NullableRoleEnum**](RoleEnum.md) | Your role in the organization. | [readonly] 
**SubscriptionExpiresAt** | **NullableTime** |  | [readonly] 
**SubscriptionFeatures** | **[]string** |  | [readonly] 
**SubscriptionId** | **NullableString** |  | [readonly] 
**SubscriptionPlanId** | **NullableString** |  | [readonly] 
**SubscriptionPlanName** | **NullableString** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewOrganization

`func NewOrganization(url string, id string, name string, current bool, role NullableRoleEnum, subscriptionExpiresAt NullableTime, subscriptionFeatures []string, subscriptionId NullableString, subscriptionPlanId NullableString, subscriptionPlanName NullableString, createdAt time.Time, modifiedAt time.Time, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Organization) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Organization) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Organization) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


### GetCurrent

`func (o *Organization) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Organization) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Organization) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetRole

`func (o *Organization) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Organization) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Organization) SetRole(v RoleEnum)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *Organization) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Organization) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSubscriptionExpiresAt

`func (o *Organization) GetSubscriptionExpiresAt() time.Time`

GetSubscriptionExpiresAt returns the SubscriptionExpiresAt field if non-nil, zero value otherwise.

### GetSubscriptionExpiresAtOk

`func (o *Organization) GetSubscriptionExpiresAtOk() (*time.Time, bool)`

GetSubscriptionExpiresAtOk returns a tuple with the SubscriptionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpiresAt

`func (o *Organization) SetSubscriptionExpiresAt(v time.Time)`

SetSubscriptionExpiresAt sets SubscriptionExpiresAt field to given value.


### SetSubscriptionExpiresAtNil

`func (o *Organization) SetSubscriptionExpiresAtNil(b bool)`

 SetSubscriptionExpiresAtNil sets the value for SubscriptionExpiresAt to be an explicit nil

### UnsetSubscriptionExpiresAt
`func (o *Organization) UnsetSubscriptionExpiresAt()`

UnsetSubscriptionExpiresAt ensures that no value is present for SubscriptionExpiresAt, not even an explicit nil
### GetSubscriptionFeatures

`func (o *Organization) GetSubscriptionFeatures() []string`

GetSubscriptionFeatures returns the SubscriptionFeatures field if non-nil, zero value otherwise.

### GetSubscriptionFeaturesOk

`func (o *Organization) GetSubscriptionFeaturesOk() (*[]string, bool)`

GetSubscriptionFeaturesOk returns a tuple with the SubscriptionFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFeatures

`func (o *Organization) SetSubscriptionFeatures(v []string)`

SetSubscriptionFeatures sets SubscriptionFeatures field to given value.


### GetSubscriptionId

`func (o *Organization) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Organization) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Organization) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *Organization) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *Organization) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionPlanId

`func (o *Organization) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *Organization) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *Organization) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.


### SetSubscriptionPlanIdNil

`func (o *Organization) SetSubscriptionPlanIdNil(b bool)`

 SetSubscriptionPlanIdNil sets the value for SubscriptionPlanId to be an explicit nil

### UnsetSubscriptionPlanId
`func (o *Organization) UnsetSubscriptionPlanId()`

UnsetSubscriptionPlanId ensures that no value is present for SubscriptionPlanId, not even an explicit nil
### GetSubscriptionPlanName

`func (o *Organization) GetSubscriptionPlanName() string`

GetSubscriptionPlanName returns the SubscriptionPlanName field if non-nil, zero value otherwise.

### GetSubscriptionPlanNameOk

`func (o *Organization) GetSubscriptionPlanNameOk() (*string, bool)`

GetSubscriptionPlanNameOk returns a tuple with the SubscriptionPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanName

`func (o *Organization) SetSubscriptionPlanName(v string)`

SetSubscriptionPlanName sets SubscriptionPlanName field to given value.


### SetSubscriptionPlanNameNil

`func (o *Organization) SetSubscriptionPlanNameNil(b bool)`

 SetSubscriptionPlanNameNil sets the value for SubscriptionPlanName to be an explicit nil

### UnsetSubscriptionPlanName
`func (o *Organization) UnsetSubscriptionPlanName()`

UnsetSubscriptionPlanName ensures that no value is present for SubscriptionPlanName, not even an explicit nil
### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Organization) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Organization) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Organization) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


