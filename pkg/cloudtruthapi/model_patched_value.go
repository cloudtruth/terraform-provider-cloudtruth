/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudtruthapi

import (
	"encoding/json"
	"time"
)

// checks if the PatchedValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedValue{}

// PatchedValue A value for a parameter in a given environment.
type PatchedValue struct {
	Url *string `json:"url,omitempty"`
	// A unique identifier for the value.
	Id *string `json:"id,omitempty"`
	// The environment this value is set in.
	Environment *string `json:"environment,omitempty"`
	// The environment name for this value.  This is a convenience to avoid another query against the server to resolve the environment url into a name.
	EnvironmentName *string `json:"environment_name,omitempty"`
	// The earliest tag name this value appears in (within the value's environment).
	EarliestTag NullableString `json:"earliest_tag,omitempty"`
	// The parameter this value is for.
	Parameter *string `json:"parameter,omitempty"`
	// An external parameter leverages a CloudTruth integration to retrieve content on-demand from an external source.  When this is `false` the value is stored by CloudTruth and considered to be _internal_.  When this is `true`, the `external_fqn` field must be set.
	External *bool `json:"external,omitempty"`
	// The FQN, or Fully-Qualified Name, is the path through the integration to get to the desired content.  This must be present and reference a valid integration when the value is `external`.
	ExternalFqn *string `json:"external_fqn,omitempty"`
	// If the value is `external`, the content returned by the integration can be reduced by applying a JMESpath expression.  This is valid as long as the content is structured and of a supported format.  JMESpath expressions are supported on `json`, `yaml`, and `dotenv` content.
	ExternalFilter *string `json:"external_filter,omitempty"`
	// This field is deprecated and unused.
	ExternalError NullableString `json:"external_error,omitempty"`
	ExternalStatus NullablePatchedValueExternalStatus `json:"external_status,omitempty"`
	// This is the content to use when resolving the Value for an internal non-secret, or when storing a secret.  When storing a secret, this content is stored in your organization's dedicated vault and this field is cleared.  This field is required if the value is being created or updated and is `internal`.  This field cannot be specified when creating or updating an `external` value.
	InternalValue NullableString `json:"internal_value,omitempty"`
	// If `true`, apply template substitution rules to this value.  If `false`, this value is a literal value.  Note: secrets cannot be interpolated.
	Interpolated *bool `json:"interpolated,omitempty"`
	// This is the actual content of the Value for the given parameter in the given environment.  If you request secret masking, no secret content will be included in the result and instead a series of asterisks will be used instead for the value.  If you request wrapping, the secret content will be wrapped in an envelope that is bound to your JWT token.  For more information about secret wrapping, see the docs.  Clients applying this value to a shell environment should set `<parameter_name>=<value>` even if `value` is the empty string.  If `value` is `null`, the client should unset that shell environment variable.
	Value NullableString `json:"value,omitempty"`
	// If true, the `value` field has undergone template evaluation.
	Evaluated *bool `json:"evaluated,omitempty"`
	// Indicates the value content is a secret.  Normally this is `true` when the parameter is a secret. It is possible for a parameter to be a secret with a external value that is not a secret.  It is not possible to convert a parameter from a secret to a non-secret if any of the values are external and a secret.  Clients can check this condition by leveraging this field.  It is also possible for a parameter to not be a secret but for this value to be dynamic and reference a Parameter that is a secret.  In this case, we indicate the value is a secret.
	Secret NullableBool `json:"secret,omitempty"`
	// The parameters this value references, if interpolated.
	ReferencedParameters []string `json:"referenced_parameters,omitempty"`
	// The templates this value references, if interpolated.
	ReferencedTemplates []string `json:"referenced_templates,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

// NewPatchedValue instantiates a new PatchedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedValue() *PatchedValue {
	this := PatchedValue{}
	return &this
}

// NewPatchedValueWithDefaults instantiates a new PatchedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedValueWithDefaults() *PatchedValue {
	this := PatchedValue{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedValue) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedValue) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedValue) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedValue) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedValue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedValue) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *PatchedValue) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PatchedValue) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *PatchedValue) SetEnvironment(v string) {
	o.Environment = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *PatchedValue) GetEnvironmentName() string {
	if o == nil || IsNil(o.EnvironmentName) {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentName) {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *PatchedValue) HasEnvironmentName() bool {
	if o != nil && !IsNil(o.EnvironmentName) {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *PatchedValue) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetEarliestTag returns the EarliestTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedValue) GetEarliestTag() string {
	if o == nil || IsNil(o.EarliestTag.Get()) {
		var ret string
		return ret
	}
	return *o.EarliestTag.Get()
}

// GetEarliestTagOk returns a tuple with the EarliestTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedValue) GetEarliestTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EarliestTag.Get(), o.EarliestTag.IsSet()
}

// HasEarliestTag returns a boolean if a field has been set.
func (o *PatchedValue) HasEarliestTag() bool {
	if o != nil && o.EarliestTag.IsSet() {
		return true
	}

	return false
}

// SetEarliestTag gets a reference to the given NullableString and assigns it to the EarliestTag field.
func (o *PatchedValue) SetEarliestTag(v string) {
	o.EarliestTag.Set(&v)
}
// SetEarliestTagNil sets the value for EarliestTag to be an explicit nil
func (o *PatchedValue) SetEarliestTagNil() {
	o.EarliestTag.Set(nil)
}

// UnsetEarliestTag ensures that no value is present for EarliestTag, not even an explicit nil
func (o *PatchedValue) UnsetEarliestTag() {
	o.EarliestTag.Unset()
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *PatchedValue) GetParameter() string {
	if o == nil || IsNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetParameterOk() (*string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *PatchedValue) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *PatchedValue) SetParameter(v string) {
	o.Parameter = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *PatchedValue) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *PatchedValue) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *PatchedValue) SetExternal(v bool) {
	o.External = &v
}

// GetExternalFqn returns the ExternalFqn field value if set, zero value otherwise.
func (o *PatchedValue) GetExternalFqn() string {
	if o == nil || IsNil(o.ExternalFqn) {
		var ret string
		return ret
	}
	return *o.ExternalFqn
}

// GetExternalFqnOk returns a tuple with the ExternalFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetExternalFqnOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalFqn) {
		return nil, false
	}
	return o.ExternalFqn, true
}

// HasExternalFqn returns a boolean if a field has been set.
func (o *PatchedValue) HasExternalFqn() bool {
	if o != nil && !IsNil(o.ExternalFqn) {
		return true
	}

	return false
}

// SetExternalFqn gets a reference to the given string and assigns it to the ExternalFqn field.
func (o *PatchedValue) SetExternalFqn(v string) {
	o.ExternalFqn = &v
}

// GetExternalFilter returns the ExternalFilter field value if set, zero value otherwise.
func (o *PatchedValue) GetExternalFilter() string {
	if o == nil || IsNil(o.ExternalFilter) {
		var ret string
		return ret
	}
	return *o.ExternalFilter
}

// GetExternalFilterOk returns a tuple with the ExternalFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetExternalFilterOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalFilter) {
		return nil, false
	}
	return o.ExternalFilter, true
}

// HasExternalFilter returns a boolean if a field has been set.
func (o *PatchedValue) HasExternalFilter() bool {
	if o != nil && !IsNil(o.ExternalFilter) {
		return true
	}

	return false
}

// SetExternalFilter gets a reference to the given string and assigns it to the ExternalFilter field.
func (o *PatchedValue) SetExternalFilter(v string) {
	o.ExternalFilter = &v
}

// GetExternalError returns the ExternalError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedValue) GetExternalError() string {
	if o == nil || IsNil(o.ExternalError.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalError.Get()
}

// GetExternalErrorOk returns a tuple with the ExternalError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedValue) GetExternalErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalError.Get(), o.ExternalError.IsSet()
}

// HasExternalError returns a boolean if a field has been set.
func (o *PatchedValue) HasExternalError() bool {
	if o != nil && o.ExternalError.IsSet() {
		return true
	}

	return false
}

// SetExternalError gets a reference to the given NullableString and assigns it to the ExternalError field.
func (o *PatchedValue) SetExternalError(v string) {
	o.ExternalError.Set(&v)
}
// SetExternalErrorNil sets the value for ExternalError to be an explicit nil
func (o *PatchedValue) SetExternalErrorNil() {
	o.ExternalError.Set(nil)
}

// UnsetExternalError ensures that no value is present for ExternalError, not even an explicit nil
func (o *PatchedValue) UnsetExternalError() {
	o.ExternalError.Unset()
}

// GetExternalStatus returns the ExternalStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedValue) GetExternalStatus() PatchedValueExternalStatus {
	if o == nil || IsNil(o.ExternalStatus.Get()) {
		var ret PatchedValueExternalStatus
		return ret
	}
	return *o.ExternalStatus.Get()
}

// GetExternalStatusOk returns a tuple with the ExternalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedValue) GetExternalStatusOk() (*PatchedValueExternalStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalStatus.Get(), o.ExternalStatus.IsSet()
}

// HasExternalStatus returns a boolean if a field has been set.
func (o *PatchedValue) HasExternalStatus() bool {
	if o != nil && o.ExternalStatus.IsSet() {
		return true
	}

	return false
}

// SetExternalStatus gets a reference to the given NullablePatchedValueExternalStatus and assigns it to the ExternalStatus field.
func (o *PatchedValue) SetExternalStatus(v PatchedValueExternalStatus) {
	o.ExternalStatus.Set(&v)
}
// SetExternalStatusNil sets the value for ExternalStatus to be an explicit nil
func (o *PatchedValue) SetExternalStatusNil() {
	o.ExternalStatus.Set(nil)
}

// UnsetExternalStatus ensures that no value is present for ExternalStatus, not even an explicit nil
func (o *PatchedValue) UnsetExternalStatus() {
	o.ExternalStatus.Unset()
}

// GetInternalValue returns the InternalValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedValue) GetInternalValue() string {
	if o == nil || IsNil(o.InternalValue.Get()) {
		var ret string
		return ret
	}
	return *o.InternalValue.Get()
}

// GetInternalValueOk returns a tuple with the InternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedValue) GetInternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalValue.Get(), o.InternalValue.IsSet()
}

// HasInternalValue returns a boolean if a field has been set.
func (o *PatchedValue) HasInternalValue() bool {
	if o != nil && o.InternalValue.IsSet() {
		return true
	}

	return false
}

// SetInternalValue gets a reference to the given NullableString and assigns it to the InternalValue field.
func (o *PatchedValue) SetInternalValue(v string) {
	o.InternalValue.Set(&v)
}
// SetInternalValueNil sets the value for InternalValue to be an explicit nil
func (o *PatchedValue) SetInternalValueNil() {
	o.InternalValue.Set(nil)
}

// UnsetInternalValue ensures that no value is present for InternalValue, not even an explicit nil
func (o *PatchedValue) UnsetInternalValue() {
	o.InternalValue.Unset()
}

// GetInterpolated returns the Interpolated field value if set, zero value otherwise.
func (o *PatchedValue) GetInterpolated() bool {
	if o == nil || IsNil(o.Interpolated) {
		var ret bool
		return ret
	}
	return *o.Interpolated
}

// GetInterpolatedOk returns a tuple with the Interpolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetInterpolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Interpolated) {
		return nil, false
	}
	return o.Interpolated, true
}

// HasInterpolated returns a boolean if a field has been set.
func (o *PatchedValue) HasInterpolated() bool {
	if o != nil && !IsNil(o.Interpolated) {
		return true
	}

	return false
}

// SetInterpolated gets a reference to the given bool and assigns it to the Interpolated field.
func (o *PatchedValue) SetInterpolated(v bool) {
	o.Interpolated = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedValue) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *PatchedValue) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *PatchedValue) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *PatchedValue) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *PatchedValue) UnsetValue() {
	o.Value.Unset()
}

// GetEvaluated returns the Evaluated field value if set, zero value otherwise.
func (o *PatchedValue) GetEvaluated() bool {
	if o == nil || IsNil(o.Evaluated) {
		var ret bool
		return ret
	}
	return *o.Evaluated
}

// GetEvaluatedOk returns a tuple with the Evaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetEvaluatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Evaluated) {
		return nil, false
	}
	return o.Evaluated, true
}

// HasEvaluated returns a boolean if a field has been set.
func (o *PatchedValue) HasEvaluated() bool {
	if o != nil && !IsNil(o.Evaluated) {
		return true
	}

	return false
}

// SetEvaluated gets a reference to the given bool and assigns it to the Evaluated field.
func (o *PatchedValue) SetEvaluated(v bool) {
	o.Evaluated = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedValue) GetSecret() bool {
	if o == nil || IsNil(o.Secret.Get()) {
		var ret bool
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedValue) GetSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *PatchedValue) HasSecret() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableBool and assigns it to the Secret field.
func (o *PatchedValue) SetSecret(v bool) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *PatchedValue) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *PatchedValue) UnsetSecret() {
	o.Secret.Unset()
}

// GetReferencedParameters returns the ReferencedParameters field value if set, zero value otherwise.
func (o *PatchedValue) GetReferencedParameters() []string {
	if o == nil || IsNil(o.ReferencedParameters) {
		var ret []string
		return ret
	}
	return o.ReferencedParameters
}

// GetReferencedParametersOk returns a tuple with the ReferencedParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetReferencedParametersOk() ([]string, bool) {
	if o == nil || IsNil(o.ReferencedParameters) {
		return nil, false
	}
	return o.ReferencedParameters, true
}

// HasReferencedParameters returns a boolean if a field has been set.
func (o *PatchedValue) HasReferencedParameters() bool {
	if o != nil && !IsNil(o.ReferencedParameters) {
		return true
	}

	return false
}

// SetReferencedParameters gets a reference to the given []string and assigns it to the ReferencedParameters field.
func (o *PatchedValue) SetReferencedParameters(v []string) {
	o.ReferencedParameters = v
}

// GetReferencedTemplates returns the ReferencedTemplates field value if set, zero value otherwise.
func (o *PatchedValue) GetReferencedTemplates() []string {
	if o == nil || IsNil(o.ReferencedTemplates) {
		var ret []string
		return ret
	}
	return o.ReferencedTemplates
}

// GetReferencedTemplatesOk returns a tuple with the ReferencedTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetReferencedTemplatesOk() ([]string, bool) {
	if o == nil || IsNil(o.ReferencedTemplates) {
		return nil, false
	}
	return o.ReferencedTemplates, true
}

// HasReferencedTemplates returns a boolean if a field has been set.
func (o *PatchedValue) HasReferencedTemplates() bool {
	if o != nil && !IsNil(o.ReferencedTemplates) {
		return true
	}

	return false
}

// SetReferencedTemplates gets a reference to the given []string and assigns it to the ReferencedTemplates field.
func (o *PatchedValue) SetReferencedTemplates(v []string) {
	o.ReferencedTemplates = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PatchedValue) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedValue) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PatchedValue) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PatchedValue) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedValue) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PatchedValue) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *PatchedValue) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o PatchedValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.EnvironmentName) {
		toSerialize["environment_name"] = o.EnvironmentName
	}
	if o.EarliestTag.IsSet() {
		toSerialize["earliest_tag"] = o.EarliestTag.Get()
	}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.ExternalFqn) {
		toSerialize["external_fqn"] = o.ExternalFqn
	}
	if !IsNil(o.ExternalFilter) {
		toSerialize["external_filter"] = o.ExternalFilter
	}
	if o.ExternalError.IsSet() {
		toSerialize["external_error"] = o.ExternalError.Get()
	}
	if o.ExternalStatus.IsSet() {
		toSerialize["external_status"] = o.ExternalStatus.Get()
	}
	if o.InternalValue.IsSet() {
		toSerialize["internal_value"] = o.InternalValue.Get()
	}
	if !IsNil(o.Interpolated) {
		toSerialize["interpolated"] = o.Interpolated
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.Evaluated) {
		toSerialize["evaluated"] = o.Evaluated
	}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}
	if !IsNil(o.ReferencedParameters) {
		toSerialize["referenced_parameters"] = o.ReferencedParameters
	}
	if !IsNil(o.ReferencedTemplates) {
		toSerialize["referenced_templates"] = o.ReferencedTemplates
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return toSerialize, nil
}

type NullablePatchedValue struct {
	value *PatchedValue
	isSet bool
}

func (v NullablePatchedValue) Get() *PatchedValue {
	return v.value
}

func (v *NullablePatchedValue) Set(val *PatchedValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedValue(val *PatchedValue) *NullablePatchedValue {
	return &NullablePatchedValue{value: val, isSet: true}
}

func (v NullablePatchedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


