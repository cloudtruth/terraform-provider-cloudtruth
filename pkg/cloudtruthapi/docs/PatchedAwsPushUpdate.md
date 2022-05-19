# PatchedAwsPushUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The action name. | [optional] 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**Projects** | Pointer to **[]string** | Projects that are included in the push. | [optional] 
**Tags** | Pointer to **[]string** | Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push. | [optional] 
**Resource** | Pointer to **NullableString** | Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to insert the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to insert the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to insert the project name   - &#x60;&#x60;{{ push }}&#x60;&#x60; to insert the push name   - &#x60;&#x60;{{ tag }}&#x60;&#x60; to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the &#x60;project&#x60; substitution is required.  If you include multiple tags from different environments in the push, the &#x60;environment&#x60; substitution is required.  If you include multiple tags from the same environment in the push, the &#x60;tag&#x60; substitution is required.  In all cases, the &#x60;parameter&#x60; substitution is always required. | [optional] 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**Force** | Pointer to **bool** | Normally, push will check to see if it originated the values in the destination before making changes to them.  Forcing a push disables the ownership check. | [optional] 
**Local** | Pointer to **bool** | Normally, push will process all parameters including those that flow in from project dependencies.  Declaring a push as &#x60;local&#x60; will cause it to only process the parameters defined in the selected projects. | [optional] 
**CoerceParameters** | Pointer to **bool** | This setting allows parameters (non-secrets) to be pushed to a destination that only supports storing secrets.  This may increase your overall cost from the cloud provider as some cloud providers charge a premium for secrets-only storage. | [optional] 
**IncludeParameters** | Pointer to **bool** | Include parameters (non-secrets) in the values being pushed.  This setting requires the destination to support parameters or for the &#x60;coerce_parameters&#x60; flag to be enabled, otherwise the push will fail. | [optional] 
**IncludeSecrets** | Pointer to **bool** | Include secrets in the values being pushed.  This setting requires the destination to support secrets, otherwise the push will fail. | [optional] 
**IncludeTemplates** | Pointer to **bool** | Include templates in the values being pushed. | [optional] 

## Methods

### NewPatchedAwsPushUpdate

`func NewPatchedAwsPushUpdate() *PatchedAwsPushUpdate`

NewPatchedAwsPushUpdate instantiates a new PatchedAwsPushUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAwsPushUpdateWithDefaults

`func NewPatchedAwsPushUpdateWithDefaults() *PatchedAwsPushUpdate`

NewPatchedAwsPushUpdateWithDefaults instantiates a new PatchedAwsPushUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAwsPushUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAwsPushUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAwsPushUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAwsPushUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedAwsPushUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedAwsPushUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedAwsPushUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedAwsPushUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjects

`func (o *PatchedAwsPushUpdate) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PatchedAwsPushUpdate) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PatchedAwsPushUpdate) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PatchedAwsPushUpdate) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTags

`func (o *PatchedAwsPushUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedAwsPushUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedAwsPushUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedAwsPushUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetResource

`func (o *PatchedAwsPushUpdate) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PatchedAwsPushUpdate) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PatchedAwsPushUpdate) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PatchedAwsPushUpdate) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *PatchedAwsPushUpdate) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *PatchedAwsPushUpdate) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDryRun

`func (o *PatchedAwsPushUpdate) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PatchedAwsPushUpdate) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PatchedAwsPushUpdate) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PatchedAwsPushUpdate) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForce

`func (o *PatchedAwsPushUpdate) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *PatchedAwsPushUpdate) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *PatchedAwsPushUpdate) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *PatchedAwsPushUpdate) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetLocal

`func (o *PatchedAwsPushUpdate) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PatchedAwsPushUpdate) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PatchedAwsPushUpdate) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *PatchedAwsPushUpdate) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetCoerceParameters

`func (o *PatchedAwsPushUpdate) GetCoerceParameters() bool`

GetCoerceParameters returns the CoerceParameters field if non-nil, zero value otherwise.

### GetCoerceParametersOk

`func (o *PatchedAwsPushUpdate) GetCoerceParametersOk() (*bool, bool)`

GetCoerceParametersOk returns a tuple with the CoerceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoerceParameters

`func (o *PatchedAwsPushUpdate) SetCoerceParameters(v bool)`

SetCoerceParameters sets CoerceParameters field to given value.

### HasCoerceParameters

`func (o *PatchedAwsPushUpdate) HasCoerceParameters() bool`

HasCoerceParameters returns a boolean if a field has been set.

### GetIncludeParameters

`func (o *PatchedAwsPushUpdate) GetIncludeParameters() bool`

GetIncludeParameters returns the IncludeParameters field if non-nil, zero value otherwise.

### GetIncludeParametersOk

`func (o *PatchedAwsPushUpdate) GetIncludeParametersOk() (*bool, bool)`

GetIncludeParametersOk returns a tuple with the IncludeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParameters

`func (o *PatchedAwsPushUpdate) SetIncludeParameters(v bool)`

SetIncludeParameters sets IncludeParameters field to given value.

### HasIncludeParameters

`func (o *PatchedAwsPushUpdate) HasIncludeParameters() bool`

HasIncludeParameters returns a boolean if a field has been set.

### GetIncludeSecrets

`func (o *PatchedAwsPushUpdate) GetIncludeSecrets() bool`

GetIncludeSecrets returns the IncludeSecrets field if non-nil, zero value otherwise.

### GetIncludeSecretsOk

`func (o *PatchedAwsPushUpdate) GetIncludeSecretsOk() (*bool, bool)`

GetIncludeSecretsOk returns a tuple with the IncludeSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSecrets

`func (o *PatchedAwsPushUpdate) SetIncludeSecrets(v bool)`

SetIncludeSecrets sets IncludeSecrets field to given value.

### HasIncludeSecrets

`func (o *PatchedAwsPushUpdate) HasIncludeSecrets() bool`

HasIncludeSecrets returns a boolean if a field has been set.

### GetIncludeTemplates

`func (o *PatchedAwsPushUpdate) GetIncludeTemplates() bool`

GetIncludeTemplates returns the IncludeTemplates field if non-nil, zero value otherwise.

### GetIncludeTemplatesOk

`func (o *PatchedAwsPushUpdate) GetIncludeTemplatesOk() (*bool, bool)`

GetIncludeTemplatesOk returns a tuple with the IncludeTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTemplates

`func (o *PatchedAwsPushUpdate) SetIncludeTemplates(v bool)`

SetIncludeTemplates sets IncludeTemplates field to given value.

### HasIncludeTemplates

`func (o *PatchedAwsPushUpdate) HasIncludeTemplates() bool`

HasIncludeTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


