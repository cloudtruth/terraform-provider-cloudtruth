# ParameterTypeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** |  | [readonly] 
**ParameterType** | **string** | The type this rule is for. | [readonly] 
**Type** | [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | 
**Constraint** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewParameterTypeRule

`func NewParameterTypeRule(url string, id string, parameterType string, type_ ParameterRuleTypeEnum, constraint string, createdAt time.Time, modifiedAt time.Time, ) *ParameterTypeRule`

NewParameterTypeRule instantiates a new ParameterTypeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTypeRuleWithDefaults

`func NewParameterTypeRuleWithDefaults() *ParameterTypeRule`

NewParameterTypeRuleWithDefaults instantiates a new ParameterTypeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ParameterTypeRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ParameterTypeRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ParameterTypeRule) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *ParameterTypeRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterTypeRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterTypeRule) SetId(v string)`

SetId sets Id field to given value.


### GetParameterType

`func (o *ParameterTypeRule) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *ParameterTypeRule) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *ParameterTypeRule) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.


### GetType

`func (o *ParameterTypeRule) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterTypeRule) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterTypeRule) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.


### GetConstraint

`func (o *ParameterTypeRule) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *ParameterTypeRule) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *ParameterTypeRule) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.


### GetCreatedAt

`func (o *ParameterTypeRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParameterTypeRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParameterTypeRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ParameterTypeRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParameterTypeRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParameterTypeRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


