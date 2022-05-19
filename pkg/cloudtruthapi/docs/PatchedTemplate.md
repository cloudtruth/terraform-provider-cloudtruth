# PatchedTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The templates this value references, if interpolated. | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the template. | [optional] [readonly] 
**Name** | Pointer to **string** | The template name. | [optional] 
**Description** | Pointer to **string** | (&#39;A description of the template.  You may find it helpful to document how this template is used to assist others when they need to maintain software that uses this content.&#39;,) | [optional] 
**Evaluated** | Pointer to **bool** | If true, the &#x60;body&#x60; field has undergone evaluation. | [optional] [readonly] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 
**ReferencedParameters** | Pointer to **[]string** | Parameters that this template references. | [optional] [readonly] 
**ReferencedTemplates** | Pointer to **[]string** | Other templates that this template references. | [optional] [readonly] 
**ReferencingTemplates** | Pointer to **[]string** | Other templates that reference this template. | [optional] [readonly] 
**ReferencingValues** | Pointer to **[]string** | The dynamic values that reference this template. | [optional] [readonly] 
**HasSecret** | Pointer to **bool** | If True, this template contains secrets. | [optional] [readonly] [default to false]
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedTemplate

`func NewPatchedTemplate() *PatchedTemplate`

NewPatchedTemplate instantiates a new PatchedTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTemplateWithDefaults

`func NewPatchedTemplateWithDefaults() *PatchedTemplate`

NewPatchedTemplateWithDefaults instantiates a new PatchedTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedTemplate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluated

`func (o *PatchedTemplate) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *PatchedTemplate) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *PatchedTemplate) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.

### HasEvaluated

`func (o *PatchedTemplate) HasEvaluated() bool`

HasEvaluated returns a boolean if a field has been set.

### GetBody

`func (o *PatchedTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PatchedTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PatchedTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PatchedTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetReferencedParameters

`func (o *PatchedTemplate) GetReferencedParameters() []string`

GetReferencedParameters returns the ReferencedParameters field if non-nil, zero value otherwise.

### GetReferencedParametersOk

`func (o *PatchedTemplate) GetReferencedParametersOk() (*[]string, bool)`

GetReferencedParametersOk returns a tuple with the ReferencedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedParameters

`func (o *PatchedTemplate) SetReferencedParameters(v []string)`

SetReferencedParameters sets ReferencedParameters field to given value.

### HasReferencedParameters

`func (o *PatchedTemplate) HasReferencedParameters() bool`

HasReferencedParameters returns a boolean if a field has been set.

### GetReferencedTemplates

`func (o *PatchedTemplate) GetReferencedTemplates() []string`

GetReferencedTemplates returns the ReferencedTemplates field if non-nil, zero value otherwise.

### GetReferencedTemplatesOk

`func (o *PatchedTemplate) GetReferencedTemplatesOk() (*[]string, bool)`

GetReferencedTemplatesOk returns a tuple with the ReferencedTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedTemplates

`func (o *PatchedTemplate) SetReferencedTemplates(v []string)`

SetReferencedTemplates sets ReferencedTemplates field to given value.

### HasReferencedTemplates

`func (o *PatchedTemplate) HasReferencedTemplates() bool`

HasReferencedTemplates returns a boolean if a field has been set.

### GetReferencingTemplates

`func (o *PatchedTemplate) GetReferencingTemplates() []string`

GetReferencingTemplates returns the ReferencingTemplates field if non-nil, zero value otherwise.

### GetReferencingTemplatesOk

`func (o *PatchedTemplate) GetReferencingTemplatesOk() (*[]string, bool)`

GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingTemplates

`func (o *PatchedTemplate) SetReferencingTemplates(v []string)`

SetReferencingTemplates sets ReferencingTemplates field to given value.

### HasReferencingTemplates

`func (o *PatchedTemplate) HasReferencingTemplates() bool`

HasReferencingTemplates returns a boolean if a field has been set.

### GetReferencingValues

`func (o *PatchedTemplate) GetReferencingValues() []string`

GetReferencingValues returns the ReferencingValues field if non-nil, zero value otherwise.

### GetReferencingValuesOk

`func (o *PatchedTemplate) GetReferencingValuesOk() (*[]string, bool)`

GetReferencingValuesOk returns a tuple with the ReferencingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingValues

`func (o *PatchedTemplate) SetReferencingValues(v []string)`

SetReferencingValues sets ReferencingValues field to given value.

### HasReferencingValues

`func (o *PatchedTemplate) HasReferencingValues() bool`

HasReferencingValues returns a boolean if a field has been set.

### GetHasSecret

`func (o *PatchedTemplate) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *PatchedTemplate) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *PatchedTemplate) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.

### HasHasSecret

`func (o *PatchedTemplate) HasHasSecret() bool`

HasHasSecret returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedTemplate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedTemplate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedTemplate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedTemplate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


