# ImportCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | Project name or identifier | 
**Environment** | Pointer to **NullableString** | Environment name or identifier | [optional] 
**Body** | **string** | Text to turn into variables | 
**Secrets** | **[]string** | Parameter names that should be secrets | 
**Ignore** | **[]string** | Parameter names to be ignored | 
**AddProject** | Pointer to **bool** | Create the project (if it does not exist) | [optional] [default to true]
**AddEnvironment** | Pointer to **bool** | Create the environment (if it does not exist) | [optional] [default to true]
**AddParameters** | Pointer to **bool** | Create the parameters (if they do not exist) | [optional] [default to true]
**AddOverrides** | Pointer to **bool** | Create the environment value override (if they do not exist) | [optional] [default to true]
**InheritOnSame** | Pointer to **bool** | Skip adding a parameter override if it is the same | [optional] [default to true]

## Methods

### NewImportCreateRequest

`func NewImportCreateRequest(project string, body string, secrets []string, ignore []string, ) *ImportCreateRequest`

NewImportCreateRequest instantiates a new ImportCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCreateRequestWithDefaults

`func NewImportCreateRequestWithDefaults() *ImportCreateRequest`

NewImportCreateRequestWithDefaults instantiates a new ImportCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ImportCreateRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ImportCreateRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ImportCreateRequest) SetProject(v string)`

SetProject sets Project field to given value.


### GetEnvironment

`func (o *ImportCreateRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ImportCreateRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ImportCreateRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ImportCreateRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ImportCreateRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ImportCreateRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetBody

`func (o *ImportCreateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ImportCreateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ImportCreateRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetSecrets

`func (o *ImportCreateRequest) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ImportCreateRequest) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ImportCreateRequest) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.


### GetIgnore

`func (o *ImportCreateRequest) GetIgnore() []string`

GetIgnore returns the Ignore field if non-nil, zero value otherwise.

### GetIgnoreOk

`func (o *ImportCreateRequest) GetIgnoreOk() (*[]string, bool)`

GetIgnoreOk returns a tuple with the Ignore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnore

`func (o *ImportCreateRequest) SetIgnore(v []string)`

SetIgnore sets Ignore field to given value.


### GetAddProject

`func (o *ImportCreateRequest) GetAddProject() bool`

GetAddProject returns the AddProject field if non-nil, zero value otherwise.

### GetAddProjectOk

`func (o *ImportCreateRequest) GetAddProjectOk() (*bool, bool)`

GetAddProjectOk returns a tuple with the AddProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddProject

`func (o *ImportCreateRequest) SetAddProject(v bool)`

SetAddProject sets AddProject field to given value.

### HasAddProject

`func (o *ImportCreateRequest) HasAddProject() bool`

HasAddProject returns a boolean if a field has been set.

### GetAddEnvironment

`func (o *ImportCreateRequest) GetAddEnvironment() bool`

GetAddEnvironment returns the AddEnvironment field if non-nil, zero value otherwise.

### GetAddEnvironmentOk

`func (o *ImportCreateRequest) GetAddEnvironmentOk() (*bool, bool)`

GetAddEnvironmentOk returns a tuple with the AddEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEnvironment

`func (o *ImportCreateRequest) SetAddEnvironment(v bool)`

SetAddEnvironment sets AddEnvironment field to given value.

### HasAddEnvironment

`func (o *ImportCreateRequest) HasAddEnvironment() bool`

HasAddEnvironment returns a boolean if a field has been set.

### GetAddParameters

`func (o *ImportCreateRequest) GetAddParameters() bool`

GetAddParameters returns the AddParameters field if non-nil, zero value otherwise.

### GetAddParametersOk

`func (o *ImportCreateRequest) GetAddParametersOk() (*bool, bool)`

GetAddParametersOk returns a tuple with the AddParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddParameters

`func (o *ImportCreateRequest) SetAddParameters(v bool)`

SetAddParameters sets AddParameters field to given value.

### HasAddParameters

`func (o *ImportCreateRequest) HasAddParameters() bool`

HasAddParameters returns a boolean if a field has been set.

### GetAddOverrides

`func (o *ImportCreateRequest) GetAddOverrides() bool`

GetAddOverrides returns the AddOverrides field if non-nil, zero value otherwise.

### GetAddOverridesOk

`func (o *ImportCreateRequest) GetAddOverridesOk() (*bool, bool)`

GetAddOverridesOk returns a tuple with the AddOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOverrides

`func (o *ImportCreateRequest) SetAddOverrides(v bool)`

SetAddOverrides sets AddOverrides field to given value.

### HasAddOverrides

`func (o *ImportCreateRequest) HasAddOverrides() bool`

HasAddOverrides returns a boolean if a field has been set.

### GetInheritOnSame

`func (o *ImportCreateRequest) GetInheritOnSame() bool`

GetInheritOnSame returns the InheritOnSame field if non-nil, zero value otherwise.

### GetInheritOnSameOk

`func (o *ImportCreateRequest) GetInheritOnSameOk() (*bool, bool)`

GetInheritOnSameOk returns a tuple with the InheritOnSame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritOnSame

`func (o *ImportCreateRequest) SetInheritOnSame(v bool)`

SetInheritOnSame sets InheritOnSame field to given value.

### HasInheritOnSame

`func (o *ImportCreateRequest) HasInheritOnSame() bool`

HasInheritOnSame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


