# BackupParameterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | [**NullableBackupParameterValueExternal**](BackupParameterValueExternal.md) |  | 
**Environment** | **string** |  | 
**Evaluated** | **bool** |  | 
**Source** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Raw** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupParameterValue

`func NewBackupParameterValue(external NullableBackupParameterValueExternal, environment string, evaluated bool, ) *BackupParameterValue`

NewBackupParameterValue instantiates a new BackupParameterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupParameterValueWithDefaults

`func NewBackupParameterValueWithDefaults() *BackupParameterValue`

NewBackupParameterValueWithDefaults instantiates a new BackupParameterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *BackupParameterValue) GetExternal() BackupParameterValueExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *BackupParameterValue) GetExternalOk() (*BackupParameterValueExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *BackupParameterValue) SetExternal(v BackupParameterValueExternal)`

SetExternal sets External field to given value.


### SetExternalNil

`func (o *BackupParameterValue) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *BackupParameterValue) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetEnvironment

`func (o *BackupParameterValue) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BackupParameterValue) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BackupParameterValue) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetEvaluated

`func (o *BackupParameterValue) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *BackupParameterValue) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *BackupParameterValue) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.


### GetSource

`func (o *BackupParameterValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BackupParameterValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BackupParameterValue) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BackupParameterValue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *BackupParameterValue) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *BackupParameterValue) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetProject

`func (o *BackupParameterValue) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BackupParameterValue) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BackupParameterValue) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BackupParameterValue) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *BackupParameterValue) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *BackupParameterValue) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetValue

`func (o *BackupParameterValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BackupParameterValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BackupParameterValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BackupParameterValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BackupParameterValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BackupParameterValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRaw

`func (o *BackupParameterValue) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *BackupParameterValue) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *BackupParameterValue) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *BackupParameterValue) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *BackupParameterValue) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *BackupParameterValue) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


