# PatchedParameterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Parameter** | Pointer to **string** | The parameter this rule is for. | [optional] [readonly] 
**Type** | Pointer to [**ParameterRuleTypeEnum**](ParameterRuleTypeEnum.md) |  | [optional] 
**Constraint** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedParameterRule

`func NewPatchedParameterRule() *PatchedParameterRule`

NewPatchedParameterRule instantiates a new PatchedParameterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterRuleWithDefaults

`func NewPatchedParameterRuleWithDefaults() *PatchedParameterRule`

NewPatchedParameterRuleWithDefaults instantiates a new PatchedParameterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedParameterRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedParameterRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedParameterRule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedParameterRule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedParameterRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameterRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameterRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameterRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameter

`func (o *PatchedParameterRule) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *PatchedParameterRule) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *PatchedParameterRule) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *PatchedParameterRule) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetType

`func (o *PatchedParameterRule) GetType() ParameterRuleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedParameterRule) GetTypeOk() (*ParameterRuleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedParameterRule) SetType(v ParameterRuleTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedParameterRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConstraint

`func (o *PatchedParameterRule) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *PatchedParameterRule) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *PatchedParameterRule) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *PatchedParameterRule) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedParameterRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameterRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameterRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameterRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameterRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameterRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameterRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameterRule) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


