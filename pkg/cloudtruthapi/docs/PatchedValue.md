# PatchedValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the value. | [optional] [readonly] 
**Environment** | Pointer to **string** | The environment this value is set in. | [optional] [readonly] 
**EnvironmentName** | Pointer to **string** | The environment name for this value.  This is a convenience to avoid another query against the server to resolve the environment url into a name. | [optional] [readonly] 
**EarliestTag** | Pointer to **NullableString** | The earliest tag name this value appears in (within the value&#39;s environment). | [optional] [readonly] 
**Parameter** | Pointer to **string** | The parameter this value is for. | [optional] [readonly] 
**External** | Pointer to **bool** | An external parameter leverages a CloudTruth integration to retrieve content on-demand from an external source.  When this is &#x60;false&#x60; the value is stored by CloudTruth and considered to be _internal_.  When this is &#x60;true&#x60;, the &#x60;external_fqn&#x60; field must be set. | [optional] 
**ExternalFqn** | Pointer to **string** | The FQN, or Fully-Qualified Name, is the path through the integration to get to the desired content.  This must be present and reference a valid integration when the value is &#x60;external&#x60;. | [optional] 
**ExternalFilter** | Pointer to **string** | If the value is &#x60;external&#x60;, the content returned by the integration can be reduced by applying a JMESpath expression.  This is valid as long as the content is structured and of a supported format.  JMESpath expressions are supported on &#x60;json&#x60;, &#x60;yaml&#x60;, and &#x60;dotenv&#x60; content. | [optional] 
**ExternalError** | Pointer to **NullableString** | This field is deprecated and unused. | [optional] [readonly] 
**ExternalStatus** | Pointer to [**NullablePatchedValueExternalStatus**](PatchedValueExternalStatus.md) |  | [optional] 
**InternalValue** | Pointer to **NullableString** | This is the content to use when resolving the Value for an internal non-secret, or when storing a secret.  When storing a secret, this content is stored in your organization&#39;s dedicated vault and this field is cleared.  This field is required if the value is being created or updated and is &#x60;internal&#x60;.  This field cannot be specified when creating or updating an &#x60;external&#x60; value. | [optional] 
**Interpolated** | Pointer to **bool** | If &#x60;true&#x60;, apply template substitution rules to this value.  If &#x60;false&#x60;, this value is a literal value.  Note: secrets cannot be interpolated. | [optional] 
**Value** | Pointer to **NullableString** | This is the actual content of the Value for the given parameter in the given environment.  If you request secret masking, no secret content will be included in the result and instead a series of asterisks will be used instead for the value.  If you request wrapping, the secret content will be wrapped in an envelope that is bound to your JWT token.  For more information about secret wrapping, see the docs.  Clients applying this value to a shell environment should set &#x60;&lt;parameter_name&gt;&#x3D;&lt;value&gt;&#x60; even if &#x60;value&#x60; is the empty string.  If &#x60;value&#x60; is &#x60;null&#x60;, the client should unset that shell environment variable. | [optional] [readonly] 
**Evaluated** | Pointer to **bool** | If true, the &#x60;value&#x60; field has undergone template evaluation. | [optional] [readonly] 
**Secret** | Pointer to **NullableBool** | Indicates the value content is a secret.  Normally this is &#x60;true&#x60; when the parameter is a secret. It is possible for a parameter to be a secret with a external value that is not a secret.  It is not possible to convert a parameter from a secret to a non-secret if any of the values are external and a secret.  Clients can check this condition by leveraging this field.  It is also possible for a parameter to not be a secret but for this value to be dynamic and reference a Parameter that is a secret.  In this case, we indicate the value is a secret. | [optional] [readonly] 
**ReferencedParameters** | Pointer to **[]string** | The parameters this value references, if interpolated. | [optional] [readonly] 
**ReferencedTemplates** | Pointer to **[]string** | The templates this value references, if interpolated. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedValue

`func NewPatchedValue() *PatchedValue`

NewPatchedValue instantiates a new PatchedValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedValueWithDefaults

`func NewPatchedValueWithDefaults() *PatchedValue`

NewPatchedValueWithDefaults instantiates a new PatchedValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedValue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedValue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedValue) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedValue) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedValue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *PatchedValue) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PatchedValue) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PatchedValue) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PatchedValue) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *PatchedValue) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *PatchedValue) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *PatchedValue) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *PatchedValue) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetEarliestTag

`func (o *PatchedValue) GetEarliestTag() string`

GetEarliestTag returns the EarliestTag field if non-nil, zero value otherwise.

### GetEarliestTagOk

`func (o *PatchedValue) GetEarliestTagOk() (*string, bool)`

GetEarliestTagOk returns a tuple with the EarliestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestTag

`func (o *PatchedValue) SetEarliestTag(v string)`

SetEarliestTag sets EarliestTag field to given value.

### HasEarliestTag

`func (o *PatchedValue) HasEarliestTag() bool`

HasEarliestTag returns a boolean if a field has been set.

### SetEarliestTagNil

`func (o *PatchedValue) SetEarliestTagNil(b bool)`

 SetEarliestTagNil sets the value for EarliestTag to be an explicit nil

### UnsetEarliestTag
`func (o *PatchedValue) UnsetEarliestTag()`

UnsetEarliestTag ensures that no value is present for EarliestTag, not even an explicit nil
### GetParameter

`func (o *PatchedValue) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *PatchedValue) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *PatchedValue) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *PatchedValue) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetExternal

`func (o *PatchedValue) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *PatchedValue) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *PatchedValue) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *PatchedValue) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalFqn

`func (o *PatchedValue) GetExternalFqn() string`

GetExternalFqn returns the ExternalFqn field if non-nil, zero value otherwise.

### GetExternalFqnOk

`func (o *PatchedValue) GetExternalFqnOk() (*string, bool)`

GetExternalFqnOk returns a tuple with the ExternalFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqn

`func (o *PatchedValue) SetExternalFqn(v string)`

SetExternalFqn sets ExternalFqn field to given value.

### HasExternalFqn

`func (o *PatchedValue) HasExternalFqn() bool`

HasExternalFqn returns a boolean if a field has been set.

### GetExternalFilter

`func (o *PatchedValue) GetExternalFilter() string`

GetExternalFilter returns the ExternalFilter field if non-nil, zero value otherwise.

### GetExternalFilterOk

`func (o *PatchedValue) GetExternalFilterOk() (*string, bool)`

GetExternalFilterOk returns a tuple with the ExternalFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilter

`func (o *PatchedValue) SetExternalFilter(v string)`

SetExternalFilter sets ExternalFilter field to given value.

### HasExternalFilter

`func (o *PatchedValue) HasExternalFilter() bool`

HasExternalFilter returns a boolean if a field has been set.

### GetExternalError

`func (o *PatchedValue) GetExternalError() string`

GetExternalError returns the ExternalError field if non-nil, zero value otherwise.

### GetExternalErrorOk

`func (o *PatchedValue) GetExternalErrorOk() (*string, bool)`

GetExternalErrorOk returns a tuple with the ExternalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalError

`func (o *PatchedValue) SetExternalError(v string)`

SetExternalError sets ExternalError field to given value.

### HasExternalError

`func (o *PatchedValue) HasExternalError() bool`

HasExternalError returns a boolean if a field has been set.

### SetExternalErrorNil

`func (o *PatchedValue) SetExternalErrorNil(b bool)`

 SetExternalErrorNil sets the value for ExternalError to be an explicit nil

### UnsetExternalError
`func (o *PatchedValue) UnsetExternalError()`

UnsetExternalError ensures that no value is present for ExternalError, not even an explicit nil
### GetExternalStatus

`func (o *PatchedValue) GetExternalStatus() PatchedValueExternalStatus`

GetExternalStatus returns the ExternalStatus field if non-nil, zero value otherwise.

### GetExternalStatusOk

`func (o *PatchedValue) GetExternalStatusOk() (*PatchedValueExternalStatus, bool)`

GetExternalStatusOk returns a tuple with the ExternalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatus

`func (o *PatchedValue) SetExternalStatus(v PatchedValueExternalStatus)`

SetExternalStatus sets ExternalStatus field to given value.

### HasExternalStatus

`func (o *PatchedValue) HasExternalStatus() bool`

HasExternalStatus returns a boolean if a field has been set.

### SetExternalStatusNil

`func (o *PatchedValue) SetExternalStatusNil(b bool)`

 SetExternalStatusNil sets the value for ExternalStatus to be an explicit nil

### UnsetExternalStatus
`func (o *PatchedValue) UnsetExternalStatus()`

UnsetExternalStatus ensures that no value is present for ExternalStatus, not even an explicit nil
### GetInternalValue

`func (o *PatchedValue) GetInternalValue() string`

GetInternalValue returns the InternalValue field if non-nil, zero value otherwise.

### GetInternalValueOk

`func (o *PatchedValue) GetInternalValueOk() (*string, bool)`

GetInternalValueOk returns a tuple with the InternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalValue

`func (o *PatchedValue) SetInternalValue(v string)`

SetInternalValue sets InternalValue field to given value.

### HasInternalValue

`func (o *PatchedValue) HasInternalValue() bool`

HasInternalValue returns a boolean if a field has been set.

### SetInternalValueNil

`func (o *PatchedValue) SetInternalValueNil(b bool)`

 SetInternalValueNil sets the value for InternalValue to be an explicit nil

### UnsetInternalValue
`func (o *PatchedValue) UnsetInternalValue()`

UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
### GetInterpolated

`func (o *PatchedValue) GetInterpolated() bool`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *PatchedValue) GetInterpolatedOk() (*bool, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *PatchedValue) SetInterpolated(v bool)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *PatchedValue) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.

### GetValue

`func (o *PatchedValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchedValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchedValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEvaluated

`func (o *PatchedValue) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *PatchedValue) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *PatchedValue) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.

### HasEvaluated

`func (o *PatchedValue) HasEvaluated() bool`

HasEvaluated returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedValue) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedValue) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedValue) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedValue) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *PatchedValue) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *PatchedValue) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetReferencedParameters

`func (o *PatchedValue) GetReferencedParameters() []string`

GetReferencedParameters returns the ReferencedParameters field if non-nil, zero value otherwise.

### GetReferencedParametersOk

`func (o *PatchedValue) GetReferencedParametersOk() (*[]string, bool)`

GetReferencedParametersOk returns a tuple with the ReferencedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedParameters

`func (o *PatchedValue) SetReferencedParameters(v []string)`

SetReferencedParameters sets ReferencedParameters field to given value.

### HasReferencedParameters

`func (o *PatchedValue) HasReferencedParameters() bool`

HasReferencedParameters returns a boolean if a field has been set.

### GetReferencedTemplates

`func (o *PatchedValue) GetReferencedTemplates() []string`

GetReferencedTemplates returns the ReferencedTemplates field if non-nil, zero value otherwise.

### GetReferencedTemplatesOk

`func (o *PatchedValue) GetReferencedTemplatesOk() (*[]string, bool)`

GetReferencedTemplatesOk returns a tuple with the ReferencedTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedTemplates

`func (o *PatchedValue) SetReferencedTemplates(v []string)`

SetReferencedTemplates sets ReferencedTemplates field to given value.

### HasReferencedTemplates

`func (o *PatchedValue) HasReferencedTemplates() bool`

HasReferencedTemplates returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedValue) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedValue) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedValue) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedValue) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


