# TemplateTimelineEntryTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the parameter. | [readonly] 
**Name** | **string** | The parameter name. | 
**Description** | Pointer to **string** | A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content. | [optional] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 

## Methods

### NewTemplateTimelineEntryTemplate

`func NewTemplateTimelineEntryTemplate(id string, name string, ) *TemplateTimelineEntryTemplate`

NewTemplateTimelineEntryTemplate instantiates a new TemplateTimelineEntryTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateTimelineEntryTemplateWithDefaults

`func NewTemplateTimelineEntryTemplateWithDefaults() *TemplateTimelineEntryTemplate`

NewTemplateTimelineEntryTemplateWithDefaults instantiates a new TemplateTimelineEntryTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateTimelineEntryTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateTimelineEntryTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateTimelineEntryTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TemplateTimelineEntryTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateTimelineEntryTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateTimelineEntryTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TemplateTimelineEntryTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateTimelineEntryTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateTimelineEntryTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateTimelineEntryTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBody

`func (o *TemplateTimelineEntryTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateTimelineEntryTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateTimelineEntryTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TemplateTimelineEntryTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


