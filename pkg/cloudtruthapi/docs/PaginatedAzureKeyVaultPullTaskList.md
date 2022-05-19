# PaginatedAzureKeyVaultPullTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]AzureKeyVaultPullTask**](AzureKeyVaultPullTask.md) |  | [optional] 

## Methods

### NewPaginatedAzureKeyVaultPullTaskList

`func NewPaginatedAzureKeyVaultPullTaskList() *PaginatedAzureKeyVaultPullTaskList`

NewPaginatedAzureKeyVaultPullTaskList instantiates a new PaginatedAzureKeyVaultPullTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAzureKeyVaultPullTaskListWithDefaults

`func NewPaginatedAzureKeyVaultPullTaskListWithDefaults() *PaginatedAzureKeyVaultPullTaskList`

NewPaginatedAzureKeyVaultPullTaskListWithDefaults instantiates a new PaginatedAzureKeyVaultPullTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedAzureKeyVaultPullTaskList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedAzureKeyVaultPullTaskList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedAzureKeyVaultPullTaskList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedAzureKeyVaultPullTaskList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedAzureKeyVaultPullTaskList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedAzureKeyVaultPullTaskList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedAzureKeyVaultPullTaskList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedAzureKeyVaultPullTaskList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedAzureKeyVaultPullTaskList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedAzureKeyVaultPullTaskList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedAzureKeyVaultPullTaskList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedAzureKeyVaultPullTaskList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedAzureKeyVaultPullTaskList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedAzureKeyVaultPullTaskList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedAzureKeyVaultPullTaskList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedAzureKeyVaultPullTaskList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedAzureKeyVaultPullTaskList) GetResults() []AzureKeyVaultPullTask`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAzureKeyVaultPullTaskList) GetResultsOk() (*[]AzureKeyVaultPullTask, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAzureKeyVaultPullTaskList) SetResults(v []AzureKeyVaultPullTask)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedAzureKeyVaultPullTaskList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


