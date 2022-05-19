# PaginatedAzureKeyVaultPushTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]AzureKeyVaultPushTask**](AzureKeyVaultPushTask.md) |  | [optional] 

## Methods

### NewPaginatedAzureKeyVaultPushTaskList

`func NewPaginatedAzureKeyVaultPushTaskList() *PaginatedAzureKeyVaultPushTaskList`

NewPaginatedAzureKeyVaultPushTaskList instantiates a new PaginatedAzureKeyVaultPushTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAzureKeyVaultPushTaskListWithDefaults

`func NewPaginatedAzureKeyVaultPushTaskListWithDefaults() *PaginatedAzureKeyVaultPushTaskList`

NewPaginatedAzureKeyVaultPushTaskListWithDefaults instantiates a new PaginatedAzureKeyVaultPushTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedAzureKeyVaultPushTaskList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedAzureKeyVaultPushTaskList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedAzureKeyVaultPushTaskList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedAzureKeyVaultPushTaskList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedAzureKeyVaultPushTaskList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedAzureKeyVaultPushTaskList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedAzureKeyVaultPushTaskList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedAzureKeyVaultPushTaskList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedAzureKeyVaultPushTaskList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedAzureKeyVaultPushTaskList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedAzureKeyVaultPushTaskList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedAzureKeyVaultPushTaskList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedAzureKeyVaultPushTaskList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedAzureKeyVaultPushTaskList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedAzureKeyVaultPushTaskList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedAzureKeyVaultPushTaskList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedAzureKeyVaultPushTaskList) GetResults() []AzureKeyVaultPushTask`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAzureKeyVaultPushTaskList) GetResultsOk() (*[]AzureKeyVaultPushTask, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAzureKeyVaultPushTaskList) SetResults(v []AzureKeyVaultPushTask)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedAzureKeyVaultPushTaskList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


