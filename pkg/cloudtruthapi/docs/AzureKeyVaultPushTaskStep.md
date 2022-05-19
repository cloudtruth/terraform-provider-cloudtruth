# AzureKeyVaultPushTaskStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | Unique identifier for a task step. | [readonly] 
**Operation** | Pointer to [**NullableOperationEnum**](OperationEnum.md) | The operation performed, if any.  When the operation is an update, there may be additional details in the success_detail field to describe the change.  When the project is filled in but the environment and parameterare not, the operation is on the project.  When the environmentis filled in but the project and parameter are not, the operationis on the environment.  When the project and parameter are filledin but the environment is not, the operation is on the parameter.When all three are filled in, the operation is on the value ofthe parameter of the project in the specified environment. | [optional] 
**Success** | **bool** | Indicates if the operation was successful. | 
**SuccessDetail** | Pointer to **NullableString** | Additional details about the successful operation, if any. | [optional] 
**Fqn** | Pointer to **NullableString** | The fully-qualified name (FQN) this of the value that was changed. | [optional] 
**Environment** | **NullableString** | The environment affected by this step. | [readonly] 
**EnvironmentId** | Pointer to **NullableString** | The environment id involved in the operation. | [optional] 
**EnvironmentName** | Pointer to **NullableString** | The environment name involved in the operation. | [optional] 
**Project** | **NullableString** | The project affected by this step. | [readonly] 
**ProjectId** | Pointer to **NullableString** | The project id involved in the operation. | [optional] 
**ProjectName** | Pointer to **NullableString** | The project name involved in the operation. | [optional] 
**Parameter** | **NullableString** | The parameter affected by this step. | [readonly] 
**ParameterId** | Pointer to **NullableString** | The parameter id involved in the operation. | [optional] 
**ParameterName** | Pointer to **NullableString** | The parameter name involved in the operation. | [optional] 
**VenueId** | Pointer to **NullableString** | The integration-native id for the resource. | [optional] 
**VenueName** | Pointer to **NullableString** | The name of the item or resource as known by the integration. | [optional] 
**Summary** | Pointer to **NullableString** | The summary of this step (what it was trying to do). | [optional] 
**ErrorCode** | Pointer to **NullableString** | An error code, if not successful. | [optional] 
**ErrorDetail** | Pointer to **NullableString** | Details on the error that occurred during processing. | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewAzureKeyVaultPushTaskStep

`func NewAzureKeyVaultPushTaskStep(url string, id string, success bool, environment NullableString, project NullableString, parameter NullableString, createdAt time.Time, modifiedAt time.Time, ) *AzureKeyVaultPushTaskStep`

NewAzureKeyVaultPushTaskStep instantiates a new AzureKeyVaultPushTaskStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultPushTaskStepWithDefaults

`func NewAzureKeyVaultPushTaskStepWithDefaults() *AzureKeyVaultPushTaskStep`

NewAzureKeyVaultPushTaskStepWithDefaults instantiates a new AzureKeyVaultPushTaskStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AzureKeyVaultPushTaskStep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AzureKeyVaultPushTaskStep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AzureKeyVaultPushTaskStep) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AzureKeyVaultPushTaskStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureKeyVaultPushTaskStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureKeyVaultPushTaskStep) SetId(v string)`

SetId sets Id field to given value.


### GetOperation

`func (o *AzureKeyVaultPushTaskStep) GetOperation() OperationEnum`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AzureKeyVaultPushTaskStep) GetOperationOk() (*OperationEnum, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AzureKeyVaultPushTaskStep) SetOperation(v OperationEnum)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *AzureKeyVaultPushTaskStep) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *AzureKeyVaultPushTaskStep) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *AzureKeyVaultPushTaskStep) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetSuccess

`func (o *AzureKeyVaultPushTaskStep) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AzureKeyVaultPushTaskStep) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AzureKeyVaultPushTaskStep) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetSuccessDetail

`func (o *AzureKeyVaultPushTaskStep) GetSuccessDetail() string`

GetSuccessDetail returns the SuccessDetail field if non-nil, zero value otherwise.

### GetSuccessDetailOk

`func (o *AzureKeyVaultPushTaskStep) GetSuccessDetailOk() (*string, bool)`

GetSuccessDetailOk returns a tuple with the SuccessDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessDetail

`func (o *AzureKeyVaultPushTaskStep) SetSuccessDetail(v string)`

SetSuccessDetail sets SuccessDetail field to given value.

### HasSuccessDetail

`func (o *AzureKeyVaultPushTaskStep) HasSuccessDetail() bool`

HasSuccessDetail returns a boolean if a field has been set.

### SetSuccessDetailNil

`func (o *AzureKeyVaultPushTaskStep) SetSuccessDetailNil(b bool)`

 SetSuccessDetailNil sets the value for SuccessDetail to be an explicit nil

### UnsetSuccessDetail
`func (o *AzureKeyVaultPushTaskStep) UnsetSuccessDetail()`

UnsetSuccessDetail ensures that no value is present for SuccessDetail, not even an explicit nil
### GetFqn

`func (o *AzureKeyVaultPushTaskStep) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *AzureKeyVaultPushTaskStep) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *AzureKeyVaultPushTaskStep) SetFqn(v string)`

SetFqn sets Fqn field to given value.

### HasFqn

`func (o *AzureKeyVaultPushTaskStep) HasFqn() bool`

HasFqn returns a boolean if a field has been set.

### SetFqnNil

`func (o *AzureKeyVaultPushTaskStep) SetFqnNil(b bool)`

 SetFqnNil sets the value for Fqn to be an explicit nil

### UnsetFqn
`func (o *AzureKeyVaultPushTaskStep) UnsetFqn()`

UnsetFqn ensures that no value is present for Fqn, not even an explicit nil
### GetEnvironment

`func (o *AzureKeyVaultPushTaskStep) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AzureKeyVaultPushTaskStep) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AzureKeyVaultPushTaskStep) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *AzureKeyVaultPushTaskStep) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *AzureKeyVaultPushTaskStep) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEnvironmentId

`func (o *AzureKeyVaultPushTaskStep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *AzureKeyVaultPushTaskStep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *AzureKeyVaultPushTaskStep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *AzureKeyVaultPushTaskStep) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *AzureKeyVaultPushTaskStep) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *AzureKeyVaultPushTaskStep) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetEnvironmentName

`func (o *AzureKeyVaultPushTaskStep) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *AzureKeyVaultPushTaskStep) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *AzureKeyVaultPushTaskStep) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *AzureKeyVaultPushTaskStep) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### SetEnvironmentNameNil

`func (o *AzureKeyVaultPushTaskStep) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *AzureKeyVaultPushTaskStep) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil
### GetProject

`func (o *AzureKeyVaultPushTaskStep) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AzureKeyVaultPushTaskStep) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AzureKeyVaultPushTaskStep) SetProject(v string)`

SetProject sets Project field to given value.


### SetProjectNil

`func (o *AzureKeyVaultPushTaskStep) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *AzureKeyVaultPushTaskStep) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProjectId

`func (o *AzureKeyVaultPushTaskStep) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AzureKeyVaultPushTaskStep) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AzureKeyVaultPushTaskStep) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AzureKeyVaultPushTaskStep) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *AzureKeyVaultPushTaskStep) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AzureKeyVaultPushTaskStep) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *AzureKeyVaultPushTaskStep) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *AzureKeyVaultPushTaskStep) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *AzureKeyVaultPushTaskStep) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *AzureKeyVaultPushTaskStep) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *AzureKeyVaultPushTaskStep) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *AzureKeyVaultPushTaskStep) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetParameter

`func (o *AzureKeyVaultPushTaskStep) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *AzureKeyVaultPushTaskStep) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *AzureKeyVaultPushTaskStep) SetParameter(v string)`

SetParameter sets Parameter field to given value.


### SetParameterNil

`func (o *AzureKeyVaultPushTaskStep) SetParameterNil(b bool)`

 SetParameterNil sets the value for Parameter to be an explicit nil

### UnsetParameter
`func (o *AzureKeyVaultPushTaskStep) UnsetParameter()`

UnsetParameter ensures that no value is present for Parameter, not even an explicit nil
### GetParameterId

`func (o *AzureKeyVaultPushTaskStep) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *AzureKeyVaultPushTaskStep) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *AzureKeyVaultPushTaskStep) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.

### HasParameterId

`func (o *AzureKeyVaultPushTaskStep) HasParameterId() bool`

HasParameterId returns a boolean if a field has been set.

### SetParameterIdNil

`func (o *AzureKeyVaultPushTaskStep) SetParameterIdNil(b bool)`

 SetParameterIdNil sets the value for ParameterId to be an explicit nil

### UnsetParameterId
`func (o *AzureKeyVaultPushTaskStep) UnsetParameterId()`

UnsetParameterId ensures that no value is present for ParameterId, not even an explicit nil
### GetParameterName

`func (o *AzureKeyVaultPushTaskStep) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *AzureKeyVaultPushTaskStep) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *AzureKeyVaultPushTaskStep) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *AzureKeyVaultPushTaskStep) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.

### SetParameterNameNil

`func (o *AzureKeyVaultPushTaskStep) SetParameterNameNil(b bool)`

 SetParameterNameNil sets the value for ParameterName to be an explicit nil

### UnsetParameterName
`func (o *AzureKeyVaultPushTaskStep) UnsetParameterName()`

UnsetParameterName ensures that no value is present for ParameterName, not even an explicit nil
### GetVenueId

`func (o *AzureKeyVaultPushTaskStep) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *AzureKeyVaultPushTaskStep) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *AzureKeyVaultPushTaskStep) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.

### HasVenueId

`func (o *AzureKeyVaultPushTaskStep) HasVenueId() bool`

HasVenueId returns a boolean if a field has been set.

### SetVenueIdNil

`func (o *AzureKeyVaultPushTaskStep) SetVenueIdNil(b bool)`

 SetVenueIdNil sets the value for VenueId to be an explicit nil

### UnsetVenueId
`func (o *AzureKeyVaultPushTaskStep) UnsetVenueId()`

UnsetVenueId ensures that no value is present for VenueId, not even an explicit nil
### GetVenueName

`func (o *AzureKeyVaultPushTaskStep) GetVenueName() string`

GetVenueName returns the VenueName field if non-nil, zero value otherwise.

### GetVenueNameOk

`func (o *AzureKeyVaultPushTaskStep) GetVenueNameOk() (*string, bool)`

GetVenueNameOk returns a tuple with the VenueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueName

`func (o *AzureKeyVaultPushTaskStep) SetVenueName(v string)`

SetVenueName sets VenueName field to given value.

### HasVenueName

`func (o *AzureKeyVaultPushTaskStep) HasVenueName() bool`

HasVenueName returns a boolean if a field has been set.

### SetVenueNameNil

`func (o *AzureKeyVaultPushTaskStep) SetVenueNameNil(b bool)`

 SetVenueNameNil sets the value for VenueName to be an explicit nil

### UnsetVenueName
`func (o *AzureKeyVaultPushTaskStep) UnsetVenueName()`

UnsetVenueName ensures that no value is present for VenueName, not even an explicit nil
### GetSummary

`func (o *AzureKeyVaultPushTaskStep) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AzureKeyVaultPushTaskStep) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AzureKeyVaultPushTaskStep) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AzureKeyVaultPushTaskStep) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *AzureKeyVaultPushTaskStep) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *AzureKeyVaultPushTaskStep) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetErrorCode

`func (o *AzureKeyVaultPushTaskStep) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AzureKeyVaultPushTaskStep) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AzureKeyVaultPushTaskStep) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AzureKeyVaultPushTaskStep) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AzureKeyVaultPushTaskStep) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AzureKeyVaultPushTaskStep) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *AzureKeyVaultPushTaskStep) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *AzureKeyVaultPushTaskStep) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *AzureKeyVaultPushTaskStep) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *AzureKeyVaultPushTaskStep) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *AzureKeyVaultPushTaskStep) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *AzureKeyVaultPushTaskStep) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetCreatedAt

`func (o *AzureKeyVaultPushTaskStep) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AzureKeyVaultPushTaskStep) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AzureKeyVaultPushTaskStep) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AzureKeyVaultPushTaskStep) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AzureKeyVaultPushTaskStep) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AzureKeyVaultPushTaskStep) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


