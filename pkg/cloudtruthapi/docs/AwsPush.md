# AwsPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | Unique identifier for the action. | [readonly] 
**Name** | **string** | The action name. | 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**LatestTask** | [**NullableAwsPushLatestTask**](AwsPushLatestTask.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**CoerceParameters** | Pointer to **bool** | This setting allows parameters (non-secrets) to be pushed to a destination that only supports storing secrets.  This may increase your overall cost from the cloud provider as some cloud providers charge a premium for secrets-only storage. | [optional] 
**IncludeParameters** | Pointer to **bool** | Include parameters (non-secrets) in the values being pushed.  This setting requires the destination to support parameters or for the &#x60;coerce_parameters&#x60; flag to be enabled, otherwise the push will fail. | [optional] 
**IncludeSecrets** | Pointer to **bool** | Include secrets in the values being pushed.  This setting requires the destination to support secrets, otherwise the push will fail. | [optional] 
**IncludeTemplates** | Pointer to **bool** | Include templates in the values being pushed. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**Force** | Pointer to **bool** | Normally, push will check to see if it originated the values in the destination before making changes to them.  Forcing a push disables the ownership check. | [optional] 
**Local** | Pointer to **bool** | Normally, push will process all parameters including those that flow in from project dependencies.  Declaring a push as &#x60;local&#x60; will cause it to only process the parameters defined in the selected projects. | [optional] 
**Projects** | **[]string** | Projects that are included in the push. | 
**Tags** | **[]string** | Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push. | 
**Region** | [**NullableAwsRegionEnum**](AwsRegionEnum.md) | The AWS region this push targets.  This region must be enabled in the integration. | 
**Service** | [**NullableAwsServiceEnum**](AwsServiceEnum.md) | The AWS service this push targets.  This service must be enabled in the integration. | 
**Resource** | **NullableString** | Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to insert the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to insert the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to insert the project name   - &#x60;&#x60;{{ push }}&#x60;&#x60; to insert the push name   - &#x60;&#x60;{{ tag }}&#x60;&#x60; to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the &#x60;project&#x60; substitution is required.  If you include multiple tags from different environments in the push, the &#x60;environment&#x60; substitution is required.  If you include multiple tags from the same environment in the push, the &#x60;tag&#x60; substitution is required.  In all cases, the &#x60;parameter&#x60; substitution is always required. | 

## Methods

### NewAwsPush

`func NewAwsPush(url string, id string, name string, latestTask NullableAwsPushLatestTask, createdAt time.Time, modifiedAt time.Time, projects []string, tags []string, region NullableAwsRegionEnum, service NullableAwsServiceEnum, resource NullableString, ) *AwsPush`

NewAwsPush instantiates a new AwsPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPushWithDefaults

`func NewAwsPushWithDefaults() *AwsPush`

NewAwsPushWithDefaults instantiates a new AwsPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AwsPush) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AwsPush) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AwsPush) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AwsPush) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsPush) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsPush) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AwsPush) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsPush) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsPush) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AwsPush) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsPush) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsPush) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsPush) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLatestTask

`func (o *AwsPush) GetLatestTask() AwsPushLatestTask`

GetLatestTask returns the LatestTask field if non-nil, zero value otherwise.

### GetLatestTaskOk

`func (o *AwsPush) GetLatestTaskOk() (*AwsPushLatestTask, bool)`

GetLatestTaskOk returns a tuple with the LatestTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTask

`func (o *AwsPush) SetLatestTask(v AwsPushLatestTask)`

SetLatestTask sets LatestTask field to given value.


### SetLatestTaskNil

`func (o *AwsPush) SetLatestTaskNil(b bool)`

 SetLatestTaskNil sets the value for LatestTask to be an explicit nil

### UnsetLatestTask
`func (o *AwsPush) UnsetLatestTask()`

UnsetLatestTask ensures that no value is present for LatestTask, not even an explicit nil
### GetCreatedAt

`func (o *AwsPush) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AwsPush) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AwsPush) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AwsPush) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AwsPush) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AwsPush) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCoerceParameters

`func (o *AwsPush) GetCoerceParameters() bool`

GetCoerceParameters returns the CoerceParameters field if non-nil, zero value otherwise.

### GetCoerceParametersOk

`func (o *AwsPush) GetCoerceParametersOk() (*bool, bool)`

GetCoerceParametersOk returns a tuple with the CoerceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoerceParameters

`func (o *AwsPush) SetCoerceParameters(v bool)`

SetCoerceParameters sets CoerceParameters field to given value.

### HasCoerceParameters

`func (o *AwsPush) HasCoerceParameters() bool`

HasCoerceParameters returns a boolean if a field has been set.

### GetIncludeParameters

`func (o *AwsPush) GetIncludeParameters() bool`

GetIncludeParameters returns the IncludeParameters field if non-nil, zero value otherwise.

### GetIncludeParametersOk

`func (o *AwsPush) GetIncludeParametersOk() (*bool, bool)`

GetIncludeParametersOk returns a tuple with the IncludeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParameters

`func (o *AwsPush) SetIncludeParameters(v bool)`

SetIncludeParameters sets IncludeParameters field to given value.

### HasIncludeParameters

`func (o *AwsPush) HasIncludeParameters() bool`

HasIncludeParameters returns a boolean if a field has been set.

### GetIncludeSecrets

`func (o *AwsPush) GetIncludeSecrets() bool`

GetIncludeSecrets returns the IncludeSecrets field if non-nil, zero value otherwise.

### GetIncludeSecretsOk

`func (o *AwsPush) GetIncludeSecretsOk() (*bool, bool)`

GetIncludeSecretsOk returns a tuple with the IncludeSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSecrets

`func (o *AwsPush) SetIncludeSecrets(v bool)`

SetIncludeSecrets sets IncludeSecrets field to given value.

### HasIncludeSecrets

`func (o *AwsPush) HasIncludeSecrets() bool`

HasIncludeSecrets returns a boolean if a field has been set.

### GetIncludeTemplates

`func (o *AwsPush) GetIncludeTemplates() bool`

GetIncludeTemplates returns the IncludeTemplates field if non-nil, zero value otherwise.

### GetIncludeTemplatesOk

`func (o *AwsPush) GetIncludeTemplatesOk() (*bool, bool)`

GetIncludeTemplatesOk returns a tuple with the IncludeTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTemplates

`func (o *AwsPush) SetIncludeTemplates(v bool)`

SetIncludeTemplates sets IncludeTemplates field to given value.

### HasIncludeTemplates

`func (o *AwsPush) HasIncludeTemplates() bool`

HasIncludeTemplates returns a boolean if a field has been set.

### GetDryRun

`func (o *AwsPush) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AwsPush) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AwsPush) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AwsPush) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForce

`func (o *AwsPush) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AwsPush) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AwsPush) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *AwsPush) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetLocal

`func (o *AwsPush) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AwsPush) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AwsPush) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *AwsPush) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetProjects

`func (o *AwsPush) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AwsPush) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AwsPush) SetProjects(v []string)`

SetProjects sets Projects field to given value.


### GetTags

`func (o *AwsPush) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsPush) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsPush) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetRegion

`func (o *AwsPush) GetRegion() AwsRegionEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsPush) GetRegionOk() (*AwsRegionEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsPush) SetRegion(v AwsRegionEnum)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsPush) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsPush) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetService

`func (o *AwsPush) GetService() AwsServiceEnum`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AwsPush) GetServiceOk() (*AwsServiceEnum, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AwsPush) SetService(v AwsServiceEnum)`

SetService sets Service field to given value.


### SetServiceNil

`func (o *AwsPush) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *AwsPush) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetResource

`func (o *AwsPush) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AwsPush) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AwsPush) SetResource(v string)`

SetResource sets Resource field to given value.


### SetResourceNil

`func (o *AwsPush) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AwsPush) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


