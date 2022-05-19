# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The templates this value references, if interpolated. | [readonly] 
**Id** | **string** | A unique identifier for the template. | [readonly] 
**Name** | **string** | The template name. | 
**Description** | Pointer to **string** | (&#39;A description of the template.  You may find it helpful to document how this template is used to assist others when they need to maintain software that uses this content.&#39;,) | [optional] 
**Evaluated** | **bool** | If true, the &#x60;body&#x60; field has undergone evaluation. | [readonly] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 
**ReferencedParameters** | **[]string** | Parameters that this template references. | [readonly] 
**ReferencedTemplates** | **[]string** | Other templates that this template references. | [readonly] 
**ReferencingTemplates** | **[]string** | Other templates that reference this template. | [readonly] 
**ReferencingValues** | **[]string** | The dynamic values that reference this template. | [readonly] 
**HasSecret** | **bool** | If True, this template contains secrets. | [readonly] [default to false]
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewTemplate

`func NewTemplate(url string, id string, name string, evaluated bool, referencedParameters []string, referencedTemplates []string, referencingTemplates []string, referencingValues []string, hasSecret bool, createdAt time.Time, modifiedAt time.Time, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Template) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Template) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Template) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Template) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluated

`func (o *Template) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *Template) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *Template) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.


### GetBody

`func (o *Template) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Template) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Template) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Template) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetReferencedParameters

`func (o *Template) GetReferencedParameters() []string`

GetReferencedParameters returns the ReferencedParameters field if non-nil, zero value otherwise.

### GetReferencedParametersOk

`func (o *Template) GetReferencedParametersOk() (*[]string, bool)`

GetReferencedParametersOk returns a tuple with the ReferencedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedParameters

`func (o *Template) SetReferencedParameters(v []string)`

SetReferencedParameters sets ReferencedParameters field to given value.


### GetReferencedTemplates

`func (o *Template) GetReferencedTemplates() []string`

GetReferencedTemplates returns the ReferencedTemplates field if non-nil, zero value otherwise.

### GetReferencedTemplatesOk

`func (o *Template) GetReferencedTemplatesOk() (*[]string, bool)`

GetReferencedTemplatesOk returns a tuple with the ReferencedTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedTemplates

`func (o *Template) SetReferencedTemplates(v []string)`

SetReferencedTemplates sets ReferencedTemplates field to given value.


### GetReferencingTemplates

`func (o *Template) GetReferencingTemplates() []string`

GetReferencingTemplates returns the ReferencingTemplates field if non-nil, zero value otherwise.

### GetReferencingTemplatesOk

`func (o *Template) GetReferencingTemplatesOk() (*[]string, bool)`

GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingTemplates

`func (o *Template) SetReferencingTemplates(v []string)`

SetReferencingTemplates sets ReferencingTemplates field to given value.


### GetReferencingValues

`func (o *Template) GetReferencingValues() []string`

GetReferencingValues returns the ReferencingValues field if non-nil, zero value otherwise.

### GetReferencingValuesOk

`func (o *Template) GetReferencingValuesOk() (*[]string, bool)`

GetReferencingValuesOk returns a tuple with the ReferencingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingValues

`func (o *Template) SetReferencingValues(v []string)`

SetReferencingValues sets ReferencingValues field to given value.


### GetHasSecret

`func (o *Template) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *Template) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *Template) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.


### GetCreatedAt

`func (o *Template) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Template) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Template) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Template) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Template) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Template) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


