# Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the value. | [readonly] 
**Environment** | **string** | The environment this value is set in. | [readonly] 
**EnvironmentName** | **string** | The environment name for this value.  This is a convenience to avoid another query against the server to resolve the environment url into a name. | [readonly] 
**EarliestTag** | **NullableString** | The earliest tag name this value appears in (within the value&#39;s environment). | [readonly] 
**Parameter** | **string** | The parameter this value is for. | [readonly] 
**External** | Pointer to **bool** | An external parameter leverages a CloudTruth integration to retrieve content on-demand from an external source.  When this is &#x60;false&#x60; the value is stored by CloudTruth and considered to be _internal_.  When this is &#x60;true&#x60;, the &#x60;external_fqn&#x60; field must be set. | [optional] 
**ExternalFqn** | Pointer to **string** | The FQN, or Fully-Qualified Name, is the path through the integration to get to the desired content.  This must be present and reference a valid integration when the value is &#x60;external&#x60;. | [optional] 
**ExternalFilter** | Pointer to **string** | If the value is &#x60;external&#x60;, the content returned by the integration can be reduced by applying a JMESpath expression.  This is valid as long as the content is structured and of a supported format.  JMESpath expressions are supported on &#x60;json&#x60;, &#x60;yaml&#x60;, and &#x60;dotenv&#x60; content. | [optional] 
**ExternalError** | **NullableString** | This field is deprecated and unused. | [readonly] 
**ExternalStatus** | [**NullablePatchedValueExternalStatus**](PatchedValueExternalStatus.md) |  | 
**InternalValue** | Pointer to **NullableString** | This is the content to use when resolving the Value for an internal non-secret, or when storing a secret.  When storing a secret, this content is stored in your organization&#39;s dedicated vault and this field is cleared.  This field is required if the value is being created or updated and is &#x60;internal&#x60;.  This field cannot be specified when creating or updating an &#x60;external&#x60; value. | [optional] 
**Interpolated** | Pointer to **bool** | If &#x60;true&#x60;, apply template substitution rules to this value.  If &#x60;false&#x60;, this value is a literal value.  Note: secrets cannot be interpolated. | [optional] 
**Value** | **NullableString** | This is the actual content of the Value for the given parameter in the given environment.  If you request secret masking, no secret content will be included in the result and instead a series of asterisks will be used instead for the value.  If you request wrapping, the secret content will be wrapped in an envelope that is bound to your JWT token.  For more information about secret wrapping, see the docs.  Clients applying this value to a shell environment should set &#x60;&lt;parameter_name&gt;&#x3D;&lt;value&gt;&#x60; even if &#x60;value&#x60; is the empty string.  If &#x60;value&#x60; is &#x60;null&#x60;, the client should unset that shell environment variable. | [readonly] 
**Evaluated** | **bool** | If true, the &#x60;value&#x60; field has undergone template evaluation. | [readonly] 
**Secret** | **NullableBool** | Indicates the value content is a secret.  Normally this is &#x60;true&#x60; when the parameter is a secret. It is possible for a parameter to be a secret with a external value that is not a secret.  It is not possible to convert a parameter from a secret to a non-secret if any of the values are external and a secret.  Clients can check this condition by leveraging this field.  It is also possible for a parameter to not be a secret but for this value to be dynamic and reference a Parameter that is a secret.  In this case, we indicate the value is a secret. | [readonly] 
**ReferencedParameters** | **[]string** | The parameters this value references, if interpolated. | [readonly] 
**ReferencedTemplates** | **[]string** | The templates this value references, if interpolated. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewValue

`func NewValue(url string, id string, environment string, environmentName string, earliestTag NullableString, parameter string, externalError NullableString, externalStatus NullablePatchedValueExternalStatus, value NullableString, evaluated bool, secret NullableBool, referencedParameters []string, referencedTemplates []string, createdAt time.Time, modifiedAt time.Time, ) *Value`

NewValue instantiates a new Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithDefaults

`func NewValueWithDefaults() *Value`

NewValueWithDefaults instantiates a new Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Value) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Value) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Value) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Value) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Value) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Value) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *Value) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Value) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Value) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetEnvironmentName

`func (o *Value) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *Value) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *Value) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetEarliestTag

`func (o *Value) GetEarliestTag() string`

GetEarliestTag returns the EarliestTag field if non-nil, zero value otherwise.

### GetEarliestTagOk

`func (o *Value) GetEarliestTagOk() (*string, bool)`

GetEarliestTagOk returns a tuple with the EarliestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestTag

`func (o *Value) SetEarliestTag(v string)`

SetEarliestTag sets EarliestTag field to given value.


### SetEarliestTagNil

`func (o *Value) SetEarliestTagNil(b bool)`

 SetEarliestTagNil sets the value for EarliestTag to be an explicit nil

### UnsetEarliestTag
`func (o *Value) UnsetEarliestTag()`

UnsetEarliestTag ensures that no value is present for EarliestTag, not even an explicit nil
### GetParameter

`func (o *Value) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *Value) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *Value) SetParameter(v string)`

SetParameter sets Parameter field to given value.


### GetExternal

`func (o *Value) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *Value) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *Value) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *Value) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalFqn

`func (o *Value) GetExternalFqn() string`

GetExternalFqn returns the ExternalFqn field if non-nil, zero value otherwise.

### GetExternalFqnOk

`func (o *Value) GetExternalFqnOk() (*string, bool)`

GetExternalFqnOk returns a tuple with the ExternalFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqn

`func (o *Value) SetExternalFqn(v string)`

SetExternalFqn sets ExternalFqn field to given value.

### HasExternalFqn

`func (o *Value) HasExternalFqn() bool`

HasExternalFqn returns a boolean if a field has been set.

### GetExternalFilter

`func (o *Value) GetExternalFilter() string`

GetExternalFilter returns the ExternalFilter field if non-nil, zero value otherwise.

### GetExternalFilterOk

`func (o *Value) GetExternalFilterOk() (*string, bool)`

GetExternalFilterOk returns a tuple with the ExternalFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilter

`func (o *Value) SetExternalFilter(v string)`

SetExternalFilter sets ExternalFilter field to given value.

### HasExternalFilter

`func (o *Value) HasExternalFilter() bool`

HasExternalFilter returns a boolean if a field has been set.

### GetExternalError

`func (o *Value) GetExternalError() string`

GetExternalError returns the ExternalError field if non-nil, zero value otherwise.

### GetExternalErrorOk

`func (o *Value) GetExternalErrorOk() (*string, bool)`

GetExternalErrorOk returns a tuple with the ExternalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalError

`func (o *Value) SetExternalError(v string)`

SetExternalError sets ExternalError field to given value.


### SetExternalErrorNil

`func (o *Value) SetExternalErrorNil(b bool)`

 SetExternalErrorNil sets the value for ExternalError to be an explicit nil

### UnsetExternalError
`func (o *Value) UnsetExternalError()`

UnsetExternalError ensures that no value is present for ExternalError, not even an explicit nil
### GetExternalStatus

`func (o *Value) GetExternalStatus() PatchedValueExternalStatus`

GetExternalStatus returns the ExternalStatus field if non-nil, zero value otherwise.

### GetExternalStatusOk

`func (o *Value) GetExternalStatusOk() (*PatchedValueExternalStatus, bool)`

GetExternalStatusOk returns a tuple with the ExternalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatus

`func (o *Value) SetExternalStatus(v PatchedValueExternalStatus)`

SetExternalStatus sets ExternalStatus field to given value.


### SetExternalStatusNil

`func (o *Value) SetExternalStatusNil(b bool)`

 SetExternalStatusNil sets the value for ExternalStatus to be an explicit nil

### UnsetExternalStatus
`func (o *Value) UnsetExternalStatus()`

UnsetExternalStatus ensures that no value is present for ExternalStatus, not even an explicit nil
### GetInternalValue

`func (o *Value) GetInternalValue() string`

GetInternalValue returns the InternalValue field if non-nil, zero value otherwise.

### GetInternalValueOk

`func (o *Value) GetInternalValueOk() (*string, bool)`

GetInternalValueOk returns a tuple with the InternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalValue

`func (o *Value) SetInternalValue(v string)`

SetInternalValue sets InternalValue field to given value.

### HasInternalValue

`func (o *Value) HasInternalValue() bool`

HasInternalValue returns a boolean if a field has been set.

### SetInternalValueNil

`func (o *Value) SetInternalValueNil(b bool)`

 SetInternalValueNil sets the value for InternalValue to be an explicit nil

### UnsetInternalValue
`func (o *Value) UnsetInternalValue()`

UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
### GetInterpolated

`func (o *Value) GetInterpolated() bool`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *Value) GetInterpolatedOk() (*bool, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *Value) SetInterpolated(v bool)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *Value) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.

### GetValue

`func (o *Value) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Value) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Value) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Value) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Value) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEvaluated

`func (o *Value) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *Value) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *Value) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.


### GetSecret

`func (o *Value) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Value) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Value) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### SetSecretNil

`func (o *Value) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *Value) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetReferencedParameters

`func (o *Value) GetReferencedParameters() []string`

GetReferencedParameters returns the ReferencedParameters field if non-nil, zero value otherwise.

### GetReferencedParametersOk

`func (o *Value) GetReferencedParametersOk() (*[]string, bool)`

GetReferencedParametersOk returns a tuple with the ReferencedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedParameters

`func (o *Value) SetReferencedParameters(v []string)`

SetReferencedParameters sets ReferencedParameters field to given value.


### GetReferencedTemplates

`func (o *Value) GetReferencedTemplates() []string`

GetReferencedTemplates returns the ReferencedTemplates field if non-nil, zero value otherwise.

### GetReferencedTemplatesOk

`func (o *Value) GetReferencedTemplatesOk() (*[]string, bool)`

GetReferencedTemplatesOk returns a tuple with the ReferencedTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedTemplates

`func (o *Value) SetReferencedTemplates(v []string)`

SetReferencedTemplates sets ReferencedTemplates field to given value.


### GetCreatedAt

`func (o *Value) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Value) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Value) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Value) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Value) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Value) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


