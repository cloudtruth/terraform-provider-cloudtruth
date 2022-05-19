# GitHubPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | Unique identifier for the action. | [readonly] 
**Name** | **string** | The action name. | 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | [**NullableGitHubPullLatestTask**](GitHubPullLatestTask.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**CreateEnvironments** | Pointer to **bool** | Allow the pull to create environments.  Any automatically created environments will be children of the &#x60;default&#x60; environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**CreateProjects** | Pointer to **bool** | Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**MappedValues** | [**[]Value**](Value.md) | Values being managed by a mapped pull. | [readonly] 
**Mode** | [**NullableModeEnum**](ModeEnum.md) | The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations. | [readonly] 

## Methods

### NewGitHubPull

`func NewGitHubPull(url string, id string, name string, latestTask NullableGitHubPullLatestTask, createdAt time.Time, modifiedAt time.Time, mappedValues []Value, mode NullableModeEnum, ) *GitHubPull`

NewGitHubPull instantiates a new GitHubPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubPullWithDefaults

`func NewGitHubPullWithDefaults() *GitHubPull`

NewGitHubPullWithDefaults instantiates a new GitHubPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GitHubPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitHubPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitHubPull) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *GitHubPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitHubPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitHubPull) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GitHubPull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitHubPull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitHubPull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GitHubPull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitHubPull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitHubPull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitHubPull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *GitHubPull) GetLatestTask() GitHubPullLatestTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *GitHubPull) GetLatestTaskOk() (*GitHubPullLatestTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *GitHubPull) SetLatestTask(v GitHubPullLatestTask)`

SetLatestTask sets LatestTask field to given value.


### SetLatestTaskNil

`func (o *GitHubPull) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *GitHubPull) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *GitHubPull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitHubPull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitHubPull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *GitHubPull) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GitHubPull) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GitHubPull) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreateEnvironments

`func (o *GitHubPull) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *GitHubPull) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *GitHubPull) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.

### HasCreateEnvironments

`func (o *GitHubPull) HasCreateEnvironments() bool`

HasCreateEnvironments returns a boolean if a field has been set.

### GetCreateProjects

`func (o *GitHubPull) GetCreateProjects() bool`

GetCreateProjects returns the CreateProjects field if non-nil, zero value otherwise.

### GetCreateProjectsOk

`func (o *GitHubPull) GetCreateProjectsOk() (*bool, bool)`

GetCreateProjectsOk returns a tuple with the CreateProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProjects

`func (o *GitHubPull) SetCreateProjects(v bool)`

SetCreateProjects sets CreateProjects field to given value.

### HasCreateProjects

`func (o *GitHubPull) HasCreateProjects() bool`

HasCreateProjects returns a boolean if a field has been set.

### GetDryRun

`func (o *GitHubPull) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *GitHubPull) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *GitHubPull) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *GitHubPull) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMappedValues

`func (o *GitHubPull) GetMappedValues() []Value`

GetMappedValues returns the MappedValues field if non-nil, zero value otherwise.

### GetMappedValuesOk

`func (o *GitHubPull) GetMappedValuesOk() (*[]Value, bool)`

GetMappedValuesOk returns a tuple with the MappedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedValues

`func (o *GitHubPull) SetMappedValues(v []Value)`

SetMappedValues sets MappedValues field to given value.


### GetMode

`func (o *GitHubPull) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GitHubPull) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GitHubPull) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *GitHubPull) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *GitHubPull) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


