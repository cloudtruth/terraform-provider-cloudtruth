# ParameterTypeRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | 
**Constraint** | **string** |  | 

## Methods

### NewParameterTypeRuleCreate

`func NewParameterTypeRuleCreate(type_ ParameterRuleTypeEnum, constraint string, ) *ParameterTypeRuleCreate`

NewParameterTypeRuleCreate instantiates a new ParameterTypeRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTypeRuleCreateWithDefaults

`func NewParameterTypeRuleCreateWithDefaults() *ParameterTypeRuleCreate`

NewParameterTypeRuleCreateWithDefaults instantiates a new ParameterTypeRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParameterTypeRuleCreate) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterTypeRuleCreate) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterTypeRuleCreate) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.


### GetConstraint

`func (o *ParameterTypeRuleCreate) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *ParameterTypeRuleCreate) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *ParameterTypeRuleCreate) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


