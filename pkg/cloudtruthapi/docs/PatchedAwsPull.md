# PatchedAwsPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the action. | [optional] [readonly] 
**Name** | Pointer to **string** | The action name. | [optional] 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | Pointer to [**NullableAwsPullLatestTask**](AwsPullLatestTask.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreateEnvironments** | Pointer to **bool** | Allow the pull to create environments.  Any automatically created environments will be children of the &#x60;default&#x60; environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**CreateProjects** | Pointer to **bool** | Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**MappedValues** | Pointer to [**[]Value**](Value.md) | Values being managed by a mapped pull. | [optional] [readonly] 
**Mode** | Pointer to [**NullableModeEnum**](ModeEnum.md) | The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations. | [optional] [readonly] 
**Region** | Pointer to [**NullableAwsRegionEnum**](AwsRegionEnum.md) | The AWS region to use.  This region must be enabled in the integration. | [optional] 
**Service** | Pointer to [**NullableAwsServiceEnum**](AwsServiceEnum.md) | The AWS service to use.  This service must be enabled in the integration. | [optional] 
**Resource** | Pointer to **NullableString** | Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to identify the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to identify the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the &#x60;environment&#x60;, &#x60;project&#x60;, and &#x60;parameter&#x60;. | [optional] 

## Methods

### NewPatchedAwsPull

`func NewPatchedAwsPull() *PatchedAwsPull`

NewPatchedAwsPull instantiates a new PatchedAwsPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAwsPullWithDefaults

`func NewPatchedAwsPullWithDefaults() *PatchedAwsPull`

NewPatchedAwsPullWithDefaults instantiates a new PatchedAwsPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedAwsPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedAwsPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedAwsPull) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedAwsPull) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedAwsPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAwsPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAwsPull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAwsPull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedAwsPull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAwsPull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAwsPull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAwsPull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedAwsPull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedAwsPull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedAwsPull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedAwsPull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *PatchedAwsPull) GetLatestTask() AwsPullLatestTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *PatchedAwsPull) GetLatestTaskOk() (*AwsPullLatestTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *PatchedAwsPull) SetLatestTask(v AwsPullLatestTask)`

SetLatestTask sets LatestTask field to given value.

### HasLatestTask

`func (o *PatchedAwsPull) HasLatestTask() bool`

HasLatestTask returns a boolean if a field has been set.

### SetLatestTaskNil

`func (o *PatchedAwsPull) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *PatchedAwsPull) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *PatchedAwsPull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedAwsPull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedAwsPull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedAwsPull) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedAwsPull) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedAwsPull) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedAwsPull) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedAwsPull) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreateEnvironments

`func (o *PatchedAwsPull) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *PatchedAwsPull) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *PatchedAwsPull) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.

### HasCreateEnvironments

`func (o *PatchedAwsPull) HasCreateEnvironments() bool`

HasCreateEnvironments returns a boolean if a field has been set.

### GetCreateProjects

`func (o *PatchedAwsPull) GetCreateProjects() bool`

GetCreateProjects returns the CreateProjects field if non-nil, zero value otherwise.

### GetCreateProjectsOk

`func (o *PatchedAwsPull) GetCreateProjectsOk() (*bool, bool)`

GetCreateProjectsOk returns a tuple with the CreateProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProjects

`func (o *PatchedAwsPull) SetCreateProjects(v bool)`

SetCreateProjects sets CreateProjects field to given value.

### HasCreateProjects

`func (o *PatchedAwsPull) HasCreateProjects() bool`

HasCreateProjects returns a boolean if a field has been set.

### GetDryRun

`func (o *PatchedAwsPull) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PatchedAwsPull) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PatchedAwsPull) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PatchedAwsPull) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMappedValues

`func (o *PatchedAwsPull) GetMappedValues() []Value`

GetMappedValues returns the MappedValues field if non-nil, zero value otherwise.

### GetMappedValuesOk

`func (o *PatchedAwsPull) GetMappedValuesOk() (*[]Value, bool)`

GetMappedValuesOk returns a tuple with the MappedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedValues

`func (o *PatchedAwsPull) SetMappedValues(v []Value)`

SetMappedValues sets MappedValues field to given value.

### HasMappedValues

`func (o *PatchedAwsPull) HasMappedValues() bool`

HasMappedValues returns a boolean if a field has been set.

### GetMode

`func (o *PatchedAwsPull) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedAwsPull) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedAwsPull) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedAwsPull) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *PatchedAwsPull) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *PatchedAwsPull) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetRegion

`func (o *PatchedAwsPull) GetRegion() AwsRegionEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatchedAwsPull) GetRegionOk() (*AwsRegionEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatchedAwsPull) SetRegion(v AwsRegionEnum)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PatchedAwsPull) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *PatchedAwsPull) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *PatchedAwsPull) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetService

`func (o *PatchedAwsPull) GetService() AwsServiceEnum`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchedAwsPull) GetServiceOk() (*AwsServiceEnum, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchedAwsPull) SetService(v AwsServiceEnum)`

SetService sets Service field to given value.

### HasService

`func (o *PatchedAwsPull) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *PatchedAwsPull) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *PatchedAwsPull) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetResource

`func (o *PatchedAwsPull) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PatchedAwsPull) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PatchedAwsPull) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PatchedAwsPull) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *PatchedAwsPull) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *PatchedAwsPull) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


