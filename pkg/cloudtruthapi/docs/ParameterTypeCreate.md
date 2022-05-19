# ParameterTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The parameter type name. | 
**Description** | Pointer to **string** | A description of the parameter type, provide documentation on how to use this type here. | [optional] 
**Parent** | Pointer to **NullableString** | All types must derive, either directly or indirectly, from one of the CloudTruth built-in types.   This is the ParameterType that this type is derived from. | [optional] 

## Methods

### NewParameterTypeCreate

`func NewParameterTypeCreate(name string, ) *ParameterTypeCreate`

NewParameterTypeCreate instantiates a new ParameterTypeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTypeCreateWithDefaults

`func NewParameterTypeCreateWithDefaults() *ParameterTypeCreate`

NewParameterTypeCreateWithDefaults instantiates a new ParameterTypeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParameterTypeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterTypeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterTypeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ParameterTypeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterTypeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterTypeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterTypeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *ParameterTypeCreate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ParameterTypeCreate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ParameterTypeCreate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ParameterTypeCreate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ParameterTypeCreate) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ParameterTypeCreate) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


