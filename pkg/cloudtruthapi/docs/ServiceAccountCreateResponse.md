# ServiceAccountCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** |  | [readonly] 
**User** | [**PatchedServiceAccountUser**](PatchedServiceAccountUser.md) |  | 
**Description** | Pointer to **string** | An optional description of the process or system using the service account. | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**LastUsedAt** | **NullableTime** | The most recent date and time the service account was used.  It will be null if the service account has not been used. | [readonly] 
**Apikey** | **string** | The API Key to use as a Bearer token for the service account. | [readonly] 

## Methods

### NewServiceAccountCreateResponse

`func NewServiceAccountCreateResponse(url string, id string, user PatchedServiceAccountUser, createdAt time.Time, modifiedAt time.Time, lastUsedAt NullableTime, apikey string, ) *ServiceAccountCreateResponse`

NewServiceAccountCreateResponse instantiates a new ServiceAccountCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountCreateResponseWithDefaults

`func NewServiceAccountCreateResponseWithDefaults() *ServiceAccountCreateResponse`

NewServiceAccountCreateResponseWithDefaults instantiates a new ServiceAccountCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ServiceAccountCreateResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceAccountCreateResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceAccountCreateResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *ServiceAccountCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountCreateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *ServiceAccountCreateResponse) GetUser() PatchedServiceAccountUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServiceAccountCreateResponse) GetUserOk() (*PatchedServiceAccountUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServiceAccountCreateResponse) SetUser(v PatchedServiceAccountUser)`

SetUser sets User field to given value.


### GetDescription

`func (o *ServiceAccountCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceAccountCreateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountCreateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountCreateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ServiceAccountCreateResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ServiceAccountCreateResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ServiceAccountCreateResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetLastUsedAt

`func (o *ServiceAccountCreateResponse) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ServiceAccountCreateResponse) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ServiceAccountCreateResponse) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *ServiceAccountCreateResponse) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ServiceAccountCreateResponse) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetApikey

`func (o *ServiceAccountCreateResponse) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *ServiceAccountCreateResponse) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *ServiceAccountCreateResponse) SetApikey(v string)`

SetApikey sets Apikey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


