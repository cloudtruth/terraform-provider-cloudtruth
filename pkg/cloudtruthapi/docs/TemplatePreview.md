# TemplatePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The template body to instantiate on request, instantiated on response. | 
**HasSecret** | **bool** | If True, the instantiated template contains secrets. | [readonly] 

## Methods

### NewTemplatePreview

`func NewTemplatePreview(body string, hasSecret bool, ) *TemplatePreview`

NewTemplatePreview instantiates a new TemplatePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePreviewWithDefaults

`func NewTemplatePreviewWithDefaults() *TemplatePreview`

NewTemplatePreviewWithDefaults instantiates a new TemplatePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TemplatePreview) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplatePreview) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplatePreview) SetBody(v string)`

SetBody sets Body field to given value.


### GetHasSecret

`func (o *TemplatePreview) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *TemplatePreview) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *TemplatePreview) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


