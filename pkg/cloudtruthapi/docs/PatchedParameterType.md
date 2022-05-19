# PatchedParameterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the parameter type. | [optional] [readonly] 
**Name** | Pointer to **string** | The parameter type name. | [optional] 
**Description** | Pointer to **string** | A description of the parameter type, provide documentation on how to use this type here. | [optional] 
**Rules** | Pointer to [**[]ParameterTypeRule**](ParameterTypeRule.md) | Rules applied to this parameter. | [optional] [readonly] 
**Parent** | Pointer to **NullableString** | All types must derive, either directly or indirectly, from one of the CloudTruth built-in types.   This is the ParameterType that this type is derived from. | [optional] 
**ParentName** | Pointer to **NullableString** | Name of the parent ParameterType (if any). | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterType

`func NewPatchedParameterType() *PatchedParameterType`

NewPatchedParameterType instantiates a new PatchedParameterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterTypeWithDefaults

`func NewPatchedParameterTypeWithDefaults() *PatchedParameterType`

NewPatchedParameterTypeWithDefaults instantiates a new PatchedParameterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedParameterType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedParameterType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedParameterType) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedParameterType) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedParameterType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedParameterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedParameterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedParameterType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedParameterType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedParameterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedParameterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedParameterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedParameterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *PatchedParameterType) GetRules() []ParameterTypeRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PatchedParameterType) GetRulesOk() (*[]ParameterTypeRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PatchedParameterType) SetRules(v []ParameterTypeRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PatchedParameterType) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetParent

`func (o *PatchedParameterType) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedParameterType) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedParameterType) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedParameterType) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedParameterType) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedParameterType) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetParentName

`func (o *PatchedParameterType) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *PatchedParameterType) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *PatchedParameterType) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *PatchedParameterType) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### SetParentNameNil

`func (o *PatchedParameterType) SetParentNameNil(b bool)`

 SetParentNameNil sets the value for ParentName to be an explicit nil

### UnsetParentName
`func (o *PatchedParameterType) UnsetParentName()`

UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
### GetCreatedAt

`func (o *PatchedParameterType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


