# AzureKeyVaultPushUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The action name. | 
**Description** | Pointer to **string** | The optional description for the action. | [optional] 
**Projects** | **[]string** | Projects that are included in the push. | 
**Tags** | **[]string** | Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push. | 
**Resource** | **NullableString** | Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to insert the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to insert the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to insert the project name   - &#x60;&#x60;{{ push }}&#x60;&#x60; to insert the push name   - &#x60;&#x60;{{ tag }}&#x60;&#x60; to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the &#x60;project&#x60; substitution is required.  If you include multiple tags from different environments in the push, the &#x60;environment&#x60; substitution is required.  If you include multiple tags from the same environment in the push, the &#x60;tag&#x60; substitution is required.  In all cases, the &#x60;parameter&#x60; substitution is always required. | 
**DryRun** | Pointer to **bool** | When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed. | [optional] 
**Force** | Pointer to **bool** | Normally, push will check to see if it originated the values in the destination before making changes to them.  Forcing a push disables the ownership check. | [optional] 
**Local** | Pointer to **bool** | Normally, push will process all parameters including those that flow in from project dependencies.  Declaring a push as &#x60;local&#x60; will cause it to only process the parameters defined in the selected projects. | [optional] 
**CoerceParameters** | Pointer to **bool** | This setting allows parameters (non-secrets) to be pushed to a destination that only supports storing secrets.  This may increase your overall cost from the cloud provider as some cloud providers charge a premium for secrets-only storage. | [optional] 
**IncludeParameters** | Pointer to **bool** | Include parameters (non-secrets) in the values being pushed.  This setting requires the destination to support parameters or for the &#x60;coerce_parameters&#x60; flag to be enabled, otherwise the push will fail. | [optional] 
**IncludeSecrets** | Pointer to **bool** | Include secrets in the values being pushed.  This setting requires the destination to support secrets, otherwise the push will fail. | [optional] 

## Methods

### NewAzureKeyVaultPushUpdate

`func NewAzureKeyVaultPushUpdate(name string, projects []string, tags []string, resource NullableString, ) *AzureKeyVaultPushUpdate`

NewAzureKeyVaultPushUpdate instantiates a new AzureKeyVaultPushUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultPushUpdateWithDefaults

`func NewAzureKeyVaultPushUpdateWithDefaults() *AzureKeyVaultPushUpdate`

NewAzureKeyVaultPushUpdateWithDefaults instantiates a new AzureKeyVaultPushUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureKeyVaultPushUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureKeyVaultPushUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureKeyVaultPushUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AzureKeyVaultPushUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureKeyVaultPushUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureKeyVaultPushUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureKeyVaultPushUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjects

`func (o *AzureKeyVaultPushUpdate) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AzureKeyVaultPushUpdate) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AzureKeyVaultPushUpdate) SetProjects(v []string)`

SetProjects sets Projects field to given value.


### GetTags

`func (o *AzureKeyVaultPushUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureKeyVaultPushUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureKeyVaultPushUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetResource

`func (o *AzureKeyVaultPushUpdate) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AzureKeyVaultPushUpdate) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AzureKeyVaultPushUpdate) SetResource(v string)`

SetResource sets Resource field to given value.


### SetResourceNil

`func (o *AzureKeyVaultPushUpdate) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AzureKeyVaultPushUpdate) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetDryRun

`func (o *AzureKeyVaultPushUpdate) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AzureKeyVaultPushUpdate) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AzureKeyVaultPushUpdate) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AzureKeyVaultPushUpdate) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetForce

`func (o *AzureKeyVaultPushUpdate) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AzureKeyVaultPushUpdate) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AzureKeyVaultPushUpdate) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *AzureKeyVaultPushUpdate) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetLocal

`func (o *AzureKeyVaultPushUpdate) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AzureKeyVaultPushUpdate) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AzureKeyVaultPushUpdate) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *AzureKeyVaultPushUpdate) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetCoerceParameters

`func (o *AzureKeyVaultPushUpdate) GetCoerceParameters() bool`

GetCoerceParameters returns the CoerceParameters field if non-nil, zero value otherwise.

### GetCoerceParametersOk

`func (o *AzureKeyVaultPushUpdate) GetCoerceParametersOk() (*bool, bool)`

GetCoerceParametersOk returns a tuple with the CoerceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoerceParameters

`func (o *AzureKeyVaultPushUpdate) SetCoerceParameters(v bool)`

SetCoerceParameters sets CoerceParameters field to given value.

### HasCoerceParameters

`func (o *AzureKeyVaultPushUpdate) HasCoerceParameters() bool`

HasCoerceParameters returns a boolean if a field has been set.

### GetIncludeParameters

`func (o *AzureKeyVaultPushUpdate) GetIncludeParameters() bool`

GetIncludeParameters returns the IncludeParameters field if non-nil, zero value otherwise.

### GetIncludeParametersOk

`func (o *AzureKeyVaultPushUpdate) GetIncludeParametersOk() (*bool, bool)`

GetIncludeParametersOk returns a tuple with the IncludeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParameters

`func (o *AzureKeyVaultPushUpdate) SetIncludeParameters(v bool)`

SetIncludeParameters sets IncludeParameters field to given value.

### HasIncludeParameters

`func (o *AzureKeyVaultPushUpdate) HasIncludeParameters() bool`

HasIncludeParameters returns a boolean if a field has been set.

### GetIncludeSecrets

`func (o *AzureKeyVaultPushUpdate) GetIncludeSecrets() bool`

GetIncludeSecrets returns the IncludeSecrets field if non-nil, zero value otherwise.

### GetIncludeSecretsOk

`func (o *AzureKeyVaultPushUpdate) GetIncludeSecretsOk() (*bool, bool)`

GetIncludeSecretsOk returns a tuple with the IncludeSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSecrets

`func (o *AzureKeyVaultPushUpdate) SetIncludeSecrets(v bool)`

SetIncludeSecrets sets IncludeSecrets field to given value.

### HasIncludeSecrets

`func (o *AzureKeyVaultPushUpdate) HasIncludeSecrets() bool`

HasIncludeSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


