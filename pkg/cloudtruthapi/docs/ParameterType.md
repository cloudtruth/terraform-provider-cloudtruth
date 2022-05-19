# ParameterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the parameter type. | [readonly] 
**Name** | **string** | The parameter type name. | 
**Description** | Pointer to **string** | A description of the parameter type, provide documentation on how to use this type here. | [optional] 
**Rules** | [**[]ParameterTypeRule**](ParameterTypeRule.md) | Rules applied to this parameter. | [readonly] 
**Parent** | Pointer to **NullableString** | All types must derive, either directly or indirectly, from one of the CloudTruth built-in types.   This is the ParameterType that this type is derived from. | [optional] 
**ParentName** | **NullableString** | Name of the parent ParameterType (if any). | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewParameterType

`func NewParameterType(url string, id string, name string, rules []ParameterTypeRule, parentName NullableString, createdAt time.Time, modifiedAt time.Time, ) *ParameterType`

NewParameterType instantiates a new ParameterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTypeWithDefaults

`func NewParameterTypeWithDefaults() *ParameterType`

NewParameterTypeWithDefaults instantiates a new ParameterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ParameterType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ParameterType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ParameterType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *ParameterType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterType) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ParameterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ParameterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *ParameterType) GetRules() []ParameterTypeRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ParameterType) GetRulesOk() (*[]ParameterTypeRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ParameterType) SetRules(v []ParameterTypeRule)`

SetRules sets Rules field to given value.


### GetParent

`func (o *ParameterType) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ParameterType) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ParameterType) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ParameterType) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ParameterType) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ParameterType) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetParentName

`func (o *ParameterType) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *ParameterType) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *ParameterType) SetParentName(v string)`

SetParentName sets ParentName field to given value.


### SetParentNameNil

`func (o *ParameterType) SetParentNameNil(b bool)`

 SetParentNameNil sets the value for ParentName to be an explicit nil

### UnsetParentName
`func (o *ParameterType) UnsetParentName()`

UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
### GetCreatedAt

`func (o *ParameterType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParameterType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParameterType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ParameterType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


