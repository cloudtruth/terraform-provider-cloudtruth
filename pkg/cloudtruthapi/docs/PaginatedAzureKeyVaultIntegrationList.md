# PaginatedAzureKeyVaultIntegrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]AzureKeyVaultIntegration**](AzureKeyVaultIntegration.md) |  | [optional] 

## Methods

### NewPaginatedAzureKeyVaultIntegrationList

`func NewPaginatedAzureKeyVaultIntegrationList() *PaginatedAzureKeyVaultIntegrationList`

NewPaginatedAzureKeyVaultIntegrationList instantiates a new PaginatedAzureKeyVaultIntegrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAzureKeyVaultIntegrationListWithDefaults

`func NewPaginatedAzureKeyVaultIntegrationListWithDefaults() *PaginatedAzureKeyVaultIntegrationList`

NewPaginatedAzureKeyVaultIntegrationListWithDefaults instantiates a new PaginatedAzureKeyVaultIntegrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedAzureKeyVaultIntegrationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedAzureKeyVaultIntegrationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedAzureKeyVaultIntegrationList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedAzureKeyVaultIntegrationList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedAzureKeyVaultIntegrationList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedAzureKeyVaultIntegrationList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedAzureKeyVaultIntegrationList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedAzureKeyVaultIntegrationList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedAzureKeyVaultIntegrationList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedAzureKeyVaultIntegrationList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedAzureKeyVaultIntegrationList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedAzureKeyVaultIntegrationList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedAzureKeyVaultIntegrationList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedAzureKeyVaultIntegrationList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedAzureKeyVaultIntegrationList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedAzureKeyVaultIntegrationList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedAzureKeyVaultIntegrationList) GetResults() []AzureKeyVaultIntegration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAzureKeyVaultIntegrationList) GetResultsOk() (*[]AzureKeyVaultIntegration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAzureKeyVaultIntegrationList) SetResults(v []AzureKeyVaultIntegration)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedAzureKeyVaultIntegrationList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


