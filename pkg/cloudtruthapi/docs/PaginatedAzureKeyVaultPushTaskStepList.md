# PaginatedAzureKeyVaultPushTaskStepList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]AzureKeyVaultPushTaskStep**](AzureKeyVaultPushTaskStep.md) |  | [optional] 

## Methods

### NewPaginatedAzureKeyVaultPushTaskStepList

`func NewPaginatedAzureKeyVaultPushTaskStepList() *PaginatedAzureKeyVaultPushTaskStepList`

NewPaginatedAzureKeyVaultPushTaskStepList instantiates a new PaginatedAzureKeyVaultPushTaskStepList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAzureKeyVaultPushTaskStepListWithDefaults

`func NewPaginatedAzureKeyVaultPushTaskStepListWithDefaults() *PaginatedAzureKeyVaultPushTaskStepList`

NewPaginatedAzureKeyVaultPushTaskStepListWithDefaults instantiates a new PaginatedAzureKeyVaultPushTaskStepList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedAzureKeyVaultPushTaskStepList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedAzureKeyVaultPushTaskStepList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedAzureKeyVaultPushTaskStepList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedAzureKeyVaultPushTaskStepList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedAzureKeyVaultPushTaskStepList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedAzureKeyVaultPushTaskStepList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedAzureKeyVaultPushTaskStepList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedAzureKeyVaultPushTaskStepList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedAzureKeyVaultPushTaskStepList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedAzureKeyVaultPushTaskStepList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetResults() []AzureKeyVaultPushTaskStep`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAzureKeyVaultPushTaskStepList) GetResultsOk() (*[]AzureKeyVaultPushTaskStep, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAzureKeyVaultPushTaskStepList) SetResults(v []AzureKeyVaultPushTaskStep)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedAzureKeyVaultPushTaskStepList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


