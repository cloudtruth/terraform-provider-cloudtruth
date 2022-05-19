# BackupParameterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**map[string]BackupParameterRule**](BackupParameterRule.md) |  | 
**Name** | **string** |  | 
**Parent** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupParameterType

`func NewBackupParameterType(rules map[string]BackupParameterRule, name string, parent string, ) *BackupParameterType`

NewBackupParameterType instantiates a new BackupParameterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupParameterTypeWithDefaults

`func NewBackupParameterTypeWithDefaults() *BackupParameterType`

NewBackupParameterTypeWithDefaults instantiates a new BackupParameterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *BackupParameterType) GetRules() map[string]BackupParameterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BackupParameterType) GetRulesOk() (*map[string]BackupParameterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BackupParameterType) SetRules(v map[string]BackupParameterRule)`

SetRules sets Rules field to given value.


### GetName

`func (o *BackupParameterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupParameterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupParameterType) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *BackupParameterType) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BackupParameterType) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BackupParameterType) SetParent(v string)`

SetParent sets Parent field to given value.


### GetDescription

`func (o *BackupParameterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupParameterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupParameterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupParameterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BackupParameterType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BackupParameterType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


