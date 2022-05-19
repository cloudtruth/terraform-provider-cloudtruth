# BackupProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | [**map[string]BackupParameter**](BackupParameter.md) |  | 
**Templates** | [**map[string]BackupTemplate**](BackupTemplate.md) |  | 
**Name** | **string** |  | 
**Parent** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupProject

`func NewBackupProject(parameters map[string]BackupParameter, templates map[string]BackupTemplate, name string, ) *BackupProject`

NewBackupProject instantiates a new BackupProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupProjectWithDefaults

`func NewBackupProjectWithDefaults() *BackupProject`

NewBackupProjectWithDefaults instantiates a new BackupProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *BackupProject) GetParameters() map[string]BackupParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BackupProject) GetParametersOk() (*map[string]BackupParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BackupProject) SetParameters(v map[string]BackupParameter)`

SetParameters sets Parameters field to given value.


### GetTemplates

`func (o *BackupProject) GetTemplates() map[string]BackupTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *BackupProject) GetTemplatesOk() (*map[string]BackupTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *BackupProject) SetTemplates(v map[string]BackupTemplate)`

SetTemplates sets Templates field to given value.


### GetName

`func (o *BackupProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupProject) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *BackupProject) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BackupProject) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BackupProject) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BackupProject) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *BackupProject) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *BackupProject) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDescription

`func (o *BackupProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BackupProject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BackupProject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


