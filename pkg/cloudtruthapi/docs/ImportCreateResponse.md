# ImportCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameter** | [**[]ImportParameter**](ImportParameter.md) | Project parameter values after import | 

## Methods

### NewImportCreateResponse

`func NewImportCreateResponse(parameter []ImportParameter, ) *ImportCreateResponse`

NewImportCreateResponse instantiates a new ImportCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCreateResponseWithDefaults

`func NewImportCreateResponseWithDefaults() *ImportCreateResponse`

NewImportCreateResponseWithDefaults instantiates a new ImportCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameter

`func (o *ImportCreateResponse) GetParameter() []ImportParameter`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ImportCreateResponse) GetParameterOk() (*[]ImportParameter, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ImportCreateResponse) SetParameter(v []ImportParameter)`

SetParameter sets Parameter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


