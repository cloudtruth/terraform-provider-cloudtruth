# AzureKeyVaultPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | Unique identifier for the action. | [readonly] 
**Name** | **string** | The action name. | 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | [**NullableAzureKeyVaultPullLatestTask**](AzureKeyVaultPullLatestTask.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**CreateEnvironments** | Pointer to **bool** | Allow the pull to create environments.  Any automatically created environments will be children of the &#x60;default&#x60; environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**CreateProjects** | Pointer to **bool** | Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**MappedValues** | [**[]Value**](Value.md) | Values being managed by a mapped pull. | [readonly] 
**Mode** | [**NullableModeEnum**](ModeEnum.md) | The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations. | [readonly] 
**Resource** | **NullableString** | Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to identify the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to identify the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the &#x60;environment&#x60;, &#x60;project&#x60;, and &#x60;parameter&#x60;. | 

## Methods

### NewAzureKeyVaultPull

`func NewAzureKeyVaultPull(url string, id string, name string, latestTask NullableAzureKeyVaultPullLatestTask, createdAt time.Time, modifiedAt time.Time, mappedValues []Value, mode NullableModeEnum, resource NullableString, ) *AzureKeyVaultPull`

NewAzureKeyVaultPull instantiates a new AzureKeyVaultPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultPullWithDefaults

`func NewAzureKeyVaultPullWithDefaults() *AzureKeyVaultPull`

NewAzureKeyVaultPullWithDefaults instantiates a new AzureKeyVaultPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AzureKeyVaultPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AzureKeyVaultPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AzureKeyVaultPull) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AzureKeyVaultPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureKeyVaultPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureKeyVaultPull) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AzureKeyVaultPull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureKeyVaultPull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureKeyVaultPull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AzureKeyVaultPull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureKeyVaultPull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureKeyVaultPull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureKeyVaultPull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *AzureKeyVaultPull) GetLatestTask() AzureKeyVaultPullLatestTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *AzureKeyVaultPull) GetLatestTaskOk() (*AzureKeyVaultPullLatestTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *AzureKeyVaultPull) SetLatestTask(v AzureKeyVaultPullLatestTask)`

SetLatestTask sets LatestTask field to given value.


### SetLatestTaskNil

`func (o *AzureKeyVaultPull) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *AzureKeyVaultPull) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *AzureKeyVaultPull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AzureKeyVaultPull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AzureKeyVaultPull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AzureKeyVaultPull) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AzureKeyVaultPull) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AzureKeyVaultPull) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreateEnvironments

`func (o *AzureKeyVaultPull) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *AzureKeyVaultPull) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *AzureKeyVaultPull) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.

### HasCreateEnvironments

`func (o *AzureKeyVaultPull) HasCreateEnvironments() bool`

HasCreateEnvironments returns a boolean if a field has been set.

### GetCreateProjects

`func (o *AzureKeyVaultPull) GetCreateProjects() bool`

GetCreateProjects returns the CreateProjects field if non-nil, zero value otherwise.

### GetCreateProjectsOk

`func (o *AzureKeyVaultPull) GetCreateProjectsOk() (*bool, bool)`

GetCreateProjectsOk returns a tuple with the CreateProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProjects

`func (o *AzureKeyVaultPull) SetCreateProjects(v bool)`

SetCreateProjects sets CreateProjects field to given value.

### HasCreateProjects

`func (o *AzureKeyVaultPull) HasCreateProjects() bool`

HasCreateProjects returns a boolean if a field has been set.

### GetDryRun

`func (o *AzureKeyVaultPull) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AzureKeyVaultPull) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AzureKeyVaultPull) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AzureKeyVaultPull) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMappedValues

`func (o *AzureKeyVaultPull) GetMappedValues() []Value`

GetMappedValues returns the MappedValues field if non-nil, zero value otherwise.

### GetMappedValuesOk

`func (o *AzureKeyVaultPull) GetMappedValuesOk() (*[]Value, bool)`

GetMappedValuesOk returns a tuple with the MappedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedValues

`func (o *AzureKeyVaultPull) SetMappedValues(v []Value)`

SetMappedValues sets MappedValues field to given value.


### GetMode

`func (o *AzureKeyVaultPull) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AzureKeyVaultPull) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AzureKeyVaultPull) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *AzureKeyVaultPull) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *AzureKeyVaultPull) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetResource

`func (o *AzureKeyVaultPull) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AzureKeyVaultPull) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AzureKeyVaultPull) SetResource(v string)`

SetResource sets Resource field to given value.


### SetResourceNil

`func (o *AzureKeyVaultPull) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AzureKeyVaultPull) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


