# EnvironmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The environment name. | 
**Description** | Pointer to **string** | A description of the environment.  You may find it helpful to document how this environment is used to assist others when they need to maintain software that uses this content. | [optional] 
**Parent** | Pointer to **NullableString** | Environments can inherit from a single parent environment which provides values for parameters when specific environments do not have a value set.  Every organization has one default environment that cannot be removed. | [optional] 

## Methods

### NewEnvironmentCreate

`func NewEnvironmentCreate(name string, ) *EnvironmentCreate`

NewEnvironmentCreate instantiates a new EnvironmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateWithDefaults

`func NewEnvironmentCreateWithDefaults() *EnvironmentCreate`

NewEnvironmentCreateWithDefaults instantiates a new EnvironmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnvironmentCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *EnvironmentCreate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EnvironmentCreate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EnvironmentCreate) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EnvironmentCreate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *EnvironmentCreate) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *EnvironmentCreate) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


