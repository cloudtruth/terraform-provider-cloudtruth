# GitHubIntegrationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional description for the integration. | [optional] 
**Writable** | Pointer to **bool** | Allow actions to write to the integration. | [optional] 
**GhInstallationId** | **int32** |  | 

## Methods

### NewGitHubIntegrationCreate

`func NewGitHubIntegrationCreate(ghInstallationId int32, ) *GitHubIntegrationCreate`

NewGitHubIntegrationCreate instantiates a new GitHubIntegrationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubIntegrationCreateWithDefaults

`func NewGitHubIntegrationCreateWithDefaults() *GitHubIntegrationCreate`

NewGitHubIntegrationCreateWithDefaults instantiates a new GitHubIntegrationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GitHubIntegrationCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitHubIntegrationCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitHubIntegrationCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitHubIntegrationCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWritable

`func (o *GitHubIntegrationCreate) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *GitHubIntegrationCreate) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *GitHubIntegrationCreate) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *GitHubIntegrationCreate) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetGhInstallationId

`func (o *GitHubIntegrationCreate) GetGhInstallationId() int32`

GetGhInstallationId returns the GhInstallationId field if non-nil, zero value otherwise.

### GetGhInstallationIdOk

`func (o *GitHubIntegrationCreate) GetGhInstallationIdOk() (*int32, bool)`

GetGhInstallationIdOk returns a tuple with the GhInstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhInstallationId

`func (o *GitHubIntegrationCreate) SetGhInstallationId(v int32)`

SetGhInstallationId sets GhInstallationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


