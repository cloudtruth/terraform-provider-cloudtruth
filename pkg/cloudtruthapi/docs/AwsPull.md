# AwsPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | Unique identifier for the action. | [readonly] 
**Name** | **string** | The action name. | 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | [**NullableAwsPullLatestTask**](AwsPullLatestTask.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**CreateEnvironments** | Pointer to **bool** | Allow the pull to create environments.  Any automatically created environments will be children of the &#x60;default&#x60; environment.  If an environment needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**CreateProjects** | Pointer to **bool** | Allow the pull to create projects.  If a project needs to be created but the action does not allow it, a task step will be added with a null operation, and success_detail will indicate the action did not allow it. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**MappedValues** | [**[]Value**](Value.md) | Values being managed by a mapped pull. | [readonly] 
**Mode** | [**NullableModeEnum**](ModeEnum.md) | The pull mode used.  A pattern pull uses a pattern-matching resource string with mustache-style markers to identify the project, parameter, and environment names, or with a Python regular expression that uses named capture groups that define the same three concepts.  A mapped pull uses a specific resource and JMESpath expression to deliver a value to a specific project, parameter, and environment.  This leverages external value linkages made in the value editor, and there is one mapped pull per integration provided by the system so that you can trigger external value pull synchronizations. | [readonly] 
**Region** | [**NullableAwsRegionEnum**](AwsRegionEnum.md) | The AWS region to use.  This region must be enabled in the integration. | 
**Service** | [**NullableAwsServiceEnum**](AwsServiceEnum.md) | The AWS service to use.  This service must be enabled in the integration. | 
**Resource** | **NullableString** | Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to identify the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to identify the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the &#x60;environment&#x60;, &#x60;project&#x60;, and &#x60;parameter&#x60;. | 

## Methods

### NewAwsPull

`func NewAwsPull(url string, id string, name string, latestTask NullableAwsPullLatestTask, createdAt time.Time, modifiedAt time.Time, mappedValues []Value, mode NullableModeEnum, region NullableAwsRegionEnum, service NullableAwsServiceEnum, resource NullableString, ) *AwsPull`

NewAwsPull instantiates a new AwsPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPullWithDefaults

`func NewAwsPullWithDefaults() *AwsPull`

NewAwsPullWithDefaults instantiates a new AwsPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AwsPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AwsPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AwsPull) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AwsPull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsPull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsPull) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AwsPull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsPull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsPull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AwsPull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsPull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsPull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsPull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *AwsPull) GetLatestTask() AwsPullLatestTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *AwsPull) GetLatestTaskOk() (*AwsPullLatestTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *AwsPull) SetLatestTask(v AwsPullLatestTask)`

SetLatestTask sets LatestTask field to given value.


### SetLatestTaskNil

`func (o *AwsPull) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *AwsPull) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *AwsPull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AwsPull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AwsPull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AwsPull) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AwsPull) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AwsPull) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreateEnvironments

`func (o *AwsPull) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *AwsPull) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *AwsPull) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.

### HasCreateEnvironments

`func (o *AwsPull) HasCreateEnvironments() bool`

HasCreateEnvironments returns a boolean if a field has been set.

### GetCreateProjects

`func (o *AwsPull) GetCreateProjects() bool`

GetCreateProjects returns the CreateProjects field if non-nil, zero value otherwise.

### GetCreateProjectsOk

`func (o *AwsPull) GetCreateProjectsOk() (*bool, bool)`

GetCreateProjectsOk returns a tuple with the CreateProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProjects

`func (o *AwsPull) SetCreateProjects(v bool)`

SetCreateProjects sets CreateProjects field to given value.

### HasCreateProjects

`func (o *AwsPull) HasCreateProjects() bool`

HasCreateProjects returns a boolean if a field has been set.

### GetDryRun

`func (o *AwsPull) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AwsPull) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AwsPull) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AwsPull) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMappedValues

`func (o *AwsPull) GetMappedValues() []Value`

GetMappedValues returns the MappedValues field if non-nil, zero value otherwise.

### GetMappedValuesOk

`func (o *AwsPull) GetMappedValuesOk() (*[]Value, bool)`

GetMappedValuesOk returns a tuple with the MappedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedValues

`func (o *AwsPull) SetMappedValues(v []Value)`

SetMappedValues sets MappedValues field to given value.


### GetMode

`func (o *AwsPull) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AwsPull) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AwsPull) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *AwsPull) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *AwsPull) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetRegion

`func (o *AwsPull) GetRegion() AwsRegionEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsPull) GetRegionOk() (*AwsRegionEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsPull) SetRegion(v AwsRegionEnum)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsPull) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsPull) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetService

`func (o *AwsPull) GetService() AwsServiceEnum`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AwsPull) GetServiceOk() (*AwsServiceEnum, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AwsPull) SetService(v AwsServiceEnum)`

SetService sets Service field to given value.


### SetServiceNil

`func (o *AwsPull) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *AwsPull) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetResource

`func (o *AwsPull) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AwsPull) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AwsPull) SetResource(v string)`

SetResource sets Resource field to given value.


### SetResourceNil

`func (o *AwsPull) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AwsPull) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


