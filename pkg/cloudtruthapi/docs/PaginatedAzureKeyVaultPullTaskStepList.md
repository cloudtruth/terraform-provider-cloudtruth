# PaginatedAzureKeyVaultPullTaskStepList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]AzureKeyVaultPullTaskStep**](AzureKeyVaultPullTaskStep.md) |  | [optional] 

## Methods

### NewPaginatedAzureKeyVaultPullTaskStepList

`func NewPaginatedAzureKeyVaultPullTaskStepList() *PaginatedAzureKeyVaultPullTaskStepList`

NewPaginatedAzureKeyVaultPullTaskStepList instantiates a new PaginatedAzureKeyVaultPullTaskStepList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAzureKeyVaultPullTaskStepListWithDefaults

`func NewPaginatedAzureKeyVaultPullTaskStepListWithDefaults() *PaginatedAzureKeyVaultPullTaskStepList`

NewPaginatedAzureKeyVaultPullTaskStepListWithDefaults instantiates a new PaginatedAzureKeyVaultPullTaskStepList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedAzureKeyVaultPullTaskStepList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedAzureKeyVaultPullTaskStepList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedAzureKeyVaultPullTaskStepList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedAzureKeyVaultPullTaskStepList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedAzureKeyVaultPullTaskStepList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedAzureKeyVaultPullTaskStepList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedAzureKeyVaultPullTaskStepList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedAzureKeyVaultPullTaskStepList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedAzureKeyVaultPullTaskStepList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedAzureKeyVaultPullTaskStepList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetResults() []AzureKeyVaultPullTaskStep`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAzureKeyVaultPullTaskStepList) GetResultsOk() (*[]AzureKeyVaultPullTaskStep, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAzureKeyVaultPullTaskStepList) SetResults(v []AzureKeyVaultPullTaskStep)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedAzureKeyVaultPullTaskStepList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


