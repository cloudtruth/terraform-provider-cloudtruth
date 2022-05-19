# BackupParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**map[string]BackupParameterRule**](BackupParameterRule.md) |  | 
**Values** | [**map[string]BackupParameterValuesValue**](BackupParameterValuesValue.md) |  | 
**Name** | **string** |  | 
**ParamType** | **string** |  | 
**Project** | **string** |  | 
**Secret** | **bool** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupParameter

`func NewBackupParameter(rules map[string]BackupParameterRule, values map[string]BackupParameterValuesValue, name string, paramType string, project string, secret bool, ) *BackupParameter`

NewBackupParameter instantiates a new BackupParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupParameterWithDefaults

`func NewBackupParameterWithDefaults() *BackupParameter`

NewBackupParameterWithDefaults instantiates a new BackupParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *BackupParameter) GetRules() map[string]BackupParameterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BackupParameter) GetRulesOk() (*map[string]BackupParameterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BackupParameter) SetRules(v map[string]BackupParameterRule)`

SetRules sets Rules field to given value.


### GetValues

`func (o *BackupParameter) GetValues() map[string]BackupParameterValuesValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *BackupParameter) GetValuesOk() (*map[string]BackupParameterValuesValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *BackupParameter) SetValues(v map[string]BackupParameterValuesValue)`

SetValues sets Values field to given value.


### GetName

`func (o *BackupParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupParameter) SetName(v string)`

SetName sets Name field to given value.


### GetParamType

`func (o *BackupParameter) GetParamType() string`

GetParamType returns the ParamType field if non-nil, zero value otherwise.

### GetParamTypeOk

`func (o *BackupParameter) GetParamTypeOk() (*string, bool)`

GetParamTypeOk returns a tuple with the ParamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamType

`func (o *BackupParameter) SetParamType(v string)`

SetParamType sets ParamType field to given value.


### GetProject

`func (o *BackupParameter) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BackupParameter) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BackupParameter) SetProject(v string)`

SetProject sets Project field to given value.


### GetSecret

`func (o *BackupParameter) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BackupParameter) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BackupParameter) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetDescription

`func (o *BackupParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BackupParameter) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BackupParameter) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


