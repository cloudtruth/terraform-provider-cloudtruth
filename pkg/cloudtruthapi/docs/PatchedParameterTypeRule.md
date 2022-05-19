# PatchedParameterTypeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ParameterType** | Pointer to **string** | The type this rule is for. | [optional] [readonly] 
**Type** | Pointer to [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | [optional] 
**Constraint** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterTypeRule

`func NewPatchedParameterTypeRule() *PatchedParameterTypeRule`

NewPatchedParameterTypeRule instantiates a new PatchedParameterTypeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterTypeRuleWithDefaults

`func NewPatchedParameterTypeRuleWithDefaults() *PatchedParameterTypeRule`

NewPatchedParameterTypeRuleWithDefaults instantiates a new PatchedParameterTypeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedParameterTypeRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedParameterTypeRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedParameterTypeRule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedParameterTypeRule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedParameterTypeRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterTypeRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterTypeRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterTypeRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameterType

`func (o *PatchedParameterTypeRule) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *PatchedParameterTypeRule) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *PatchedParameterTypeRule) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *PatchedParameterTypeRule) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetType

`func (o *PatchedParameterTypeRule) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedParameterTypeRule) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedParameterTypeRule) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedParameterTypeRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConstraint

`func (o *PatchedParameterTypeRule) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *PatchedParameterTypeRule) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *PatchedParameterTypeRule) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *PatchedParameterTypeRule) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedParameterTypeRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterTypeRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterTypeRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterTypeRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterTypeRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterTypeRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterTypeRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterTypeRule) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


