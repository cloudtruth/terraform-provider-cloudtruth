# ParameterExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The exported parameter body. | [readonly] 
**HasSecret** | **bool** | If True, the exported parameters include one or more secrets. | [readonly] 

## Methods

### NewParameterExport

`func NewParameterExport(body string, hasSecret bool, ) *ParameterExport`

NewParameterExport instantiates a new ParameterExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterExportWithDefaults

`func NewParameterExportWithDefaults() *ParameterExport`

NewParameterExportWithDefaults instantiates a new ParameterExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ParameterExport) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ParameterExport) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ParameterExport) SetBody(v string)`

SetBody sets Body field to given value.


### GetHasSecret

`func (o *ParameterExport) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *ParameterExport) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *ParameterExport) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


