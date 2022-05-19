# ValueCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment this value is set in. | 
**External** | Pointer to **bool** | An external parameter leverages a CloudTruth integration to retrieve content on-demand from an external source.  When this is &#x60;false&#x60; the value is stored by CloudTruth and considered to be _internal_.  When this is &#x60;true&#x60;, the &#x60;external_fqn&#x60; field must be set. | [optional] 
**ExternalFqn** | Pointer to **string** | The FQN, or Fully-Qualified Name, is the path through the integration to get to the desired content.  This must be present and reference a valid integration when the value is &#x60;external&#x60;. | [optional] 
**ExternalFilter** | Pointer to **string** | If the value is &#x60;external&#x60;, the content returned by the integration can be reduced by applying a JMESpath expression.  This is valid as long as the content is structured and of a supported format.  JMESpath expressions are supported on &#x60;json&#x60;, &#x60;yaml&#x60;, and &#x60;dotenv&#x60; content. | [optional] 
**InternalValue** | Pointer to **NullableString** | This is the content to use when resolving the Value for an internal non-secret, or when storing a secret.  When storing a secret, this content is stored in your organization&#39;s dedicated vault and this field is cleared.  This field is required if the value is being created or updated and is &#x60;internal&#x60;.  This field cannot be specified when creating or updating an &#x60;external&#x60; value. | [optional] 
**Interpolated** | Pointer to **bool** | If &#x60;true&#x60;, apply template substitution rules to this value.  If &#x60;false&#x60;, this value is a literal value.  Note: secrets cannot be interpolated. | [optional] 

## Methods

### NewValueCreate

`func NewValueCreate(environment string, ) *ValueCreate`

NewValueCreate instantiates a new ValueCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueCreateWithDefaults

`func NewValueCreateWithDefaults() *ValueCreate`

NewValueCreateWithDefaults instantiates a new ValueCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ValueCreate) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ValueCreate) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ValueCreate) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetExternal

`func (o *ValueCreate) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ValueCreate) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ValueCreate) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ValueCreate) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalFqn

`func (o *ValueCreate) GetExternalFqn() string`

GetExternalFqn returns the ExternalFqn field if non-nil, zero value otherwise.

### GetExternalFqnOk

`func (o *ValueCreate) GetExternalFqnOk() (*string, bool)`

GetExternalFqnOk returns a tuple with the ExternalFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqn

`func (o *ValueCreate) SetExternalFqn(v string)`

SetExternalFqn sets ExternalFqn field to given value.

### HasExternalFqn

`func (o *ValueCreate) HasExternalFqn() bool`

HasExternalFqn returns a boolean if a field has been set.

### GetExternalFilter

`func (o *ValueCreate) GetExternalFilter() string`

GetExternalFilter returns the ExternalFilter field if non-nil, zero value otherwise.

### GetExternalFilterOk

`func (o *ValueCreate) GetExternalFilterOk() (*string, bool)`

GetExternalFilterOk returns a tuple with the ExternalFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFilter

`func (o *ValueCreate) SetExternalFilter(v string)`

SetExternalFilter sets ExternalFilter field to given value.

### HasExternalFilter

`func (o *ValueCreate) HasExternalFilter() bool`

HasExternalFilter returns a boolean if a field has been set.

### GetInternalValue

`func (o *ValueCreate) GetInternalValue() string`

GetInternalValue returns the InternalValue field if non-nil, zero value otherwise.

### GetInternalValueOk

`func (o *ValueCreate) GetInternalValueOk() (*string, bool)`

GetInternalValueOk returns a tuple with the InternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalValue

`func (o *ValueCreate) SetInternalValue(v string)`

SetInternalValue sets InternalValue field to given value.

### HasInternalValue

`func (o *ValueCreate) HasInternalValue() bool`

HasInternalValue returns a boolean if a field has been set.

### SetInternalValueNil

`func (o *ValueCreate) SetInternalValueNil(b bool)`

 SetInternalValueNil sets the value for InternalValue to be an explicit nil

### UnsetInternalValue
`func (o *ValueCreate) UnsetInternalValue()`

UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
### GetInterpolated

`func (o *ValueCreate) GetInterpolated() bool`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *ValueCreate) GetInterpolatedOk() (*bool, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *ValueCreate) SetInterpolated(v bool)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *ValueCreate) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


