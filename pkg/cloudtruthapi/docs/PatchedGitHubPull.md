# PatchedGitHubPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the action. | [optional] [readonly] 
**Name** | Pointer to **string** | The action name. | [optional] 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | Pointer to [**NullableGitHubPullLatestTask**](GitHubPullLatestTask.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreateEnvironments** | Pointer to **bool** | Allow the pull to create environments.  Any automatically created environments will be children of the &#x60;default&#x60; environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**CreateProjects** | Pointer to **bool** | Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**MappedValues** | Pointer to [**[]Value**](Value.md) | Values being managed by a mapped pull. | [optional] [readonly] 
**Mode** | Pointer to [**NullableModeEnum**](ModeEnum.md) | The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations. | [optional] [readonly] 

## Methods

### NewPatchedGitHubPull

`func NewPatchedGitHubPull() *PatchedGitHubPull`

NewPatchedGitHubPull instantiates a new PatchedGitHubPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGitHubPullWithDefaults

`func NewPatchedGitHubPullWithDefaults() *PatchedGitHubPull`

NewPatchedGitHubPullWithDefaults instantiates a new PatchedGitHubPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedGitHubPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedGitHubPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedGitHubPull) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedGitHubPull) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedGitHubPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedGitHubPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedGitHubPull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedGitHubPull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedGitHubPull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGitHubPull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGitHubPull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGitHubPull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedGitHubPull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedGitHubPull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedGitHubPull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedGitHubPull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *PatchedGitHubPull) GetLatestTask() GitHubPullLatestTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *PatchedGitHubPull) GetLatestTaskOk() (*GitHubPullLatestTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *PatchedGitHubPull) SetLatestTask(v GitHubPullLatestTask)`

SetLatestTask sets LatestTask field to given value.

### HasLatestTask

`func (o *PatchedGitHubPull) HasLatestTask() bool`

HasLatestTask returns a boolean if a field has been set.

### SetLatestTaskNil

`func (o *PatchedGitHubPull) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *PatchedGitHubPull) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *PatchedGitHubPull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedGitHubPull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedGitHubPull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedGitHubPull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedGitHubPull) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedGitHubPull) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedGitHubPull) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedGitHubPull) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreateEnvironments

`func (o *PatchedGitHubPull) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *PatchedGitHubPull) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *PatchedGitHubPull) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.

### HasCreateEnvironments

`func (o *PatchedGitHubPull) HasCreateEnvironments() bool`

HasCreateEnvironments returns a boolean if a field has been set.

### GetCreateProjects

`func (o *PatchedGitHubPull) GetCreateProjects() bool`

GetCreateProjects returns the CreateProjects field if non-nil, zero value otherwise.

### GetCreateProjectsOk

`func (o *PatchedGitHubPull) GetCreateProjectsOk() (*bool, bool)`

GetCreateProjectsOk returns a tuple with the CreateProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProjects

`func (o *PatchedGitHubPull) SetCreateProjects(v bool)`

SetCreateProjects sets CreateProjects field to given value.

### HasCreateProjects

`func (o *PatchedGitHubPull) HasCreateProjects() bool`

HasCreateProjects returns a boolean if a field has been set.

### GetDryRun

`func (o *PatchedGitHubPull) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PatchedGitHubPull) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PatchedGitHubPull) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PatchedGitHubPull) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMappedValues

`func (o *PatchedGitHubPull) GetMappedValues() []Value`

GetMappedValues returns the MappedValues field if non-nil, zero value otherwise.

### GetMappedValuesOk

`func (o *PatchedGitHubPull) GetMappedValuesOk() (*[]Value, bool)`

GetMappedValuesOk returns a tuple with the MappedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedValues

`func (o *PatchedGitHubPull) SetMappedValues(v []Value)`

SetMappedValues sets MappedValues field to given value.

### HasMappedValues

`func (o *PatchedGitHubPull) HasMappedValues() bool`

HasMappedValues returns a boolean if a field has been set.

### GetMode

`func (o *PatchedGitHubPull) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedGitHubPull) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedGitHubPull) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedGitHubPull) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *PatchedGitHubPull) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *PatchedGitHubPull) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


