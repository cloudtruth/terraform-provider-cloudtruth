# GitHubIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | The unique identifier for the integration. | [readonly] 
**Name** | **string** |  | [readonly] 
**Description** | Pointer to **string** | An optional description for the integration. | [optional] 
**Status** | [**NullableStatusEnum**](StatusEnum.md) | The status of the integration connection with the third-party provider as of the &#x60;status_last_checked_at&#x60; field.  The status is updated automatically by the server when the integration is modified. | [readonly] 
**StatusDetail** | **string** | If an error occurs, more details will be available in this field. | [readonly] 
**StatusLastCheckedAt** | **time.Time** | The last time the status was evaluated. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**Fqn** | **string** |  | [readonly] 
**Type** | **string** | The type of integration. | [readonly] 
**Writable** | Pointer to **bool** | Allow actions to write to the integration. | [optional] 
**GhInstallationId** | **int32** |  | 
**GhOrganizationSlug** | **string** |  | [readonly] 

## Methods

### NewGitHubIntegration

`func NewGitHubIntegration(url string, id string, name string, status NullableStatusEnum, statusDetail string, statusLastCheckedAt time.Time, createdAt time.Time, modifiedAt time.Time, fqn string, type_ string, ghInstallationId int32, ghOrganizationSlug string, ) *GitHubIntegration`

NewGitHubIntegration instantiates a new GitHubIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubIntegrationWithDefaults

`func NewGitHubIntegrationWithDefaults() *GitHubIntegration`

NewGitHubIntegrationWithDefaults instantiates a new GitHubIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GitHubIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitHubIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitHubIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *GitHubIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitHubIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitHubIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GitHubIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitHubIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitHubIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GitHubIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitHubIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitHubIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitHubIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *GitHubIntegration) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GitHubIntegration) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GitHubIntegration) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *GitHubIntegration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GitHubIntegration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusDetail

`func (o *GitHubIntegration) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *GitHubIntegration) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *GitHubIntegration) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.


### GetStatusLastCheckedAt

`func (o *GitHubIntegration) GetStatusLastCheckedAt() time.Time`

GetStatusLastCheckedAt returns the StatusLastCheckedAt field if non-nil, zero value otherwise.

### GetStatusLastCheckedAtOk

`func (o *GitHubIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool)`

GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastCheckedAt

`func (o *GitHubIntegration) SetStatusLastCheckedAt(v time.Time)`

SetStatusLastCheckedAt sets StatusLastCheckedAt field to given value.


### GetCreatedAt

`func (o *GitHubIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitHubIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitHubIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *GitHubIntegration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GitHubIntegration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GitHubIntegration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetFqn

`func (o *GitHubIntegration) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *GitHubIntegration) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *GitHubIntegration) SetFqn(v string)`

SetFqn sets Fqn field to given value.


### GetType

`func (o *GitHubIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitHubIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitHubIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetWritable

`func (o *GitHubIntegration) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *GitHubIntegration) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *GitHubIntegration) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *GitHubIntegration) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetGhInstallationId

`func (o *GitHubIntegration) GetGhInstallationId() int32`

GetGhInstallationId returns the GhInstallationId field if non-nil, zero value otherwise.

### GetGhInstallationIdOk

`func (o *GitHubIntegration) GetGhInstallationIdOk() (*int32, bool)`

GetGhInstallationIdOk returns a tuple with the GhInstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhInstallationId

`func (o *GitHubIntegration) SetGhInstallationId(v int32)`

SetGhInstallationId sets GhInstallationId field to given value.


### GetGhOrganizationSlug

`func (o *GitHubIntegration) GetGhOrganizationSlug() string`

GetGhOrganizationSlug returns the GhOrganizationSlug field if non-nil, zero value otherwise.

### GetGhOrganizationSlugOk

`func (o *GitHubIntegration) GetGhOrganizationSlugOk() (*string, bool)`

GetGhOrganizationSlugOk returns a tuple with the GhOrganizationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhOrganizationSlug

`func (o *GitHubIntegration) SetGhOrganizationSlug(v string)`

SetGhOrganizationSlug sets GhOrganizationSlug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


