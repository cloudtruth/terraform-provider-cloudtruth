# TemplateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The template name. | 
**Description** | Pointer to **string** | (&#39;A description of the template.  You may find it helpful to document how this template is used to assist others when they need to maintain software that uses this content.&#39;,) | [optional] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 

## Methods

### NewTemplateCreate

`func NewTemplateCreate(name string, ) *TemplateCreate`

NewTemplateCreate instantiates a new TemplateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateCreateWithDefaults

`func NewTemplateCreateWithDefaults() *TemplateCreate`

NewTemplateCreateWithDefaults instantiates a new TemplateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TemplateCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBody

`func (o *TemplateCreate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateCreate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateCreate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TemplateCreate) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


