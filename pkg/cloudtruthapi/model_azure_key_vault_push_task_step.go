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

// checks if the AzureKeyVaultPushTaskStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureKeyVaultPushTaskStep{}

// AzureKeyVaultPushTaskStep Push task step for an Azure Key Vault integration.
type AzureKeyVaultPushTaskStep struct {
	Url string `json:"url"`
	// Unique identifier for a task step.
	Id string `json:"id"`
	Operation NullableOperationEnum `json:"operation,omitempty"`
	// Indicates if the operation was successful.
	Success bool `json:"success"`
	// Additional details about the successful operation, if any.
	SuccessDetail NullableString `json:"success_detail,omitempty"`
	// The fully-qualified name (FQN) this of the value that was changed.
	Fqn NullableString `json:"fqn,omitempty"`
	// The environment affected by this step.
	Environment NullableString `json:"environment"`
	// The environment id involved in the operation.
	EnvironmentId NullableString `json:"environment_id,omitempty"`
	// The environment name involved in the operation.
	EnvironmentName NullableString `json:"environment_name,omitempty"`
	// The project affected by this step.
	Project NullableString `json:"project"`
	// The project id involved in the operation.
	ProjectId NullableString `json:"project_id,omitempty"`
	// The project name involved in the operation.
	ProjectName NullableString `json:"project_name,omitempty"`
	// The parameter affected by this step.
	Parameter NullableString `json:"parameter"`
	// The parameter id involved in the operation.
	ParameterId NullableString `json:"parameter_id,omitempty"`
	// The parameter name involved in the operation.
	ParameterName NullableString `json:"parameter_name,omitempty"`
	// The integration-native id for the resource.
	VenueId NullableString `json:"venue_id,omitempty"`
	// The name of the item or resource as known by the integration.
	VenueName NullableString `json:"venue_name,omitempty"`
	// The summary of this step (what it was trying to do).
	Summary NullableString `json:"summary,omitempty"`
	// An error code, if not successful.
	ErrorCode NullableString `json:"error_code,omitempty"`
	// Details on the error that occurred during processing.
	ErrorDetail NullableString `json:"error_detail,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewAzureKeyVaultPushTaskStep instantiates a new AzureKeyVaultPushTaskStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPushTaskStep(url string, id string, success bool, environment NullableString, project NullableString, parameter NullableString, createdAt time.Time, modifiedAt NullableTime) *AzureKeyVaultPushTaskStep {
	this := AzureKeyVaultPushTaskStep{}
	this.Url = url
	this.Id = id
	this.Success = success
	this.Environment = environment
	this.Project = project
	this.Parameter = parameter
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAzureKeyVaultPushTaskStepWithDefaults instantiates a new AzureKeyVaultPushTaskStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPushTaskStepWithDefaults() *AzureKeyVaultPushTaskStep {
	this := AzureKeyVaultPushTaskStep{}
	return &this
}

// GetUrl returns the Url field value
func (o *AzureKeyVaultPushTaskStep) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushTaskStep) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AzureKeyVaultPushTaskStep) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AzureKeyVaultPushTaskStep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushTaskStep) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultPushTaskStep) SetId(v string) {
	o.Id = v
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetOperation() OperationEnum {
	if o == nil || IsNil(o.Operation.Get()) {
		var ret OperationEnum
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetOperationOk() (*OperationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableOperationEnum and assigns it to the Operation field.
func (o *AzureKeyVaultPushTaskStep) SetOperation(v OperationEnum) {
	o.Operation.Set(&v)
}
// SetOperationNil sets the value for Operation to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetOperation() {
	o.Operation.Unset()
}

// GetSuccess returns the Success field value
func (o *AzureKeyVaultPushTaskStep) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushTaskStep) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AzureKeyVaultPushTaskStep) SetSuccess(v bool) {
	o.Success = v
}

// GetSuccessDetail returns the SuccessDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetSuccessDetail() string {
	if o == nil || IsNil(o.SuccessDetail.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessDetail.Get()
}

// GetSuccessDetailOk returns a tuple with the SuccessDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetSuccessDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessDetail.Get(), o.SuccessDetail.IsSet()
}

// HasSuccessDetail returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasSuccessDetail() bool {
	if o != nil && o.SuccessDetail.IsSet() {
		return true
	}

	return false
}

// SetSuccessDetail gets a reference to the given NullableString and assigns it to the SuccessDetail field.
func (o *AzureKeyVaultPushTaskStep) SetSuccessDetail(v string) {
	o.SuccessDetail.Set(&v)
}
// SetSuccessDetailNil sets the value for SuccessDetail to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetSuccessDetailNil() {
	o.SuccessDetail.Set(nil)
}

// UnsetSuccessDetail ensures that no value is present for SuccessDetail, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetSuccessDetail() {
	o.SuccessDetail.Unset()
}

// GetFqn returns the Fqn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetFqn() string {
	if o == nil || IsNil(o.Fqn.Get()) {
		var ret string
		return ret
	}
	return *o.Fqn.Get()
}

// GetFqnOk returns a tuple with the Fqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fqn.Get(), o.Fqn.IsSet()
}

// HasFqn returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasFqn() bool {
	if o != nil && o.Fqn.IsSet() {
		return true
	}

	return false
}

// SetFqn gets a reference to the given NullableString and assigns it to the Fqn field.
func (o *AzureKeyVaultPushTaskStep) SetFqn(v string) {
	o.Fqn.Set(&v)
}
// SetFqnNil sets the value for Fqn to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetFqnNil() {
	o.Fqn.Set(nil)
}

// UnsetFqn ensures that no value is present for Fqn, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetFqn() {
	o.Fqn.Unset()
}

// GetEnvironment returns the Environment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureKeyVaultPushTaskStep) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// SetEnvironment sets field value
func (o *AzureKeyVaultPushTaskStep) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *AzureKeyVaultPushTaskStep) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}
// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetEnvironmentName() string {
	if o == nil || IsNil(o.EnvironmentName.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentName.Get()
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentName.Get(), o.EnvironmentName.IsSet()
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasEnvironmentName() bool {
	if o != nil && o.EnvironmentName.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentName gets a reference to the given NullableString and assigns it to the EnvironmentName field.
func (o *AzureKeyVaultPushTaskStep) SetEnvironmentName(v string) {
	o.EnvironmentName.Set(&v)
}
// SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetEnvironmentNameNil() {
	o.EnvironmentName.Set(nil)
}

// UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetEnvironmentName() {
	o.EnvironmentName.Unset()
}

// GetProject returns the Project field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureKeyVaultPushTaskStep) GetProject() string {
	if o == nil || o.Project.Get() == nil {
		var ret string
		return ret
	}

	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// SetProject sets field value
func (o *AzureKeyVaultPushTaskStep) SetProject(v string) {
	o.Project.Set(&v)
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *AzureKeyVaultPushTaskStep) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *AzureKeyVaultPushTaskStep) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetParameter returns the Parameter field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureKeyVaultPushTaskStep) GetParameter() string {
	if o == nil || o.Parameter.Get() == nil {
		var ret string
		return ret
	}

	return *o.Parameter.Get()
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameter.Get(), o.Parameter.IsSet()
}

// SetParameter sets field value
func (o *AzureKeyVaultPushTaskStep) SetParameter(v string) {
	o.Parameter.Set(&v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetParameterId() string {
	if o == nil || IsNil(o.ParameterId.Get()) {
		var ret string
		return ret
	}
	return *o.ParameterId.Get()
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetParameterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParameterId.Get(), o.ParameterId.IsSet()
}

// HasParameterId returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasParameterId() bool {
	if o != nil && o.ParameterId.IsSet() {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given NullableString and assigns it to the ParameterId field.
func (o *AzureKeyVaultPushTaskStep) SetParameterId(v string) {
	o.ParameterId.Set(&v)
}
// SetParameterIdNil sets the value for ParameterId to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetParameterIdNil() {
	o.ParameterId.Set(nil)
}

// UnsetParameterId ensures that no value is present for ParameterId, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetParameterId() {
	o.ParameterId.Unset()
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetParameterName() string {
	if o == nil || IsNil(o.ParameterName.Get()) {
		var ret string
		return ret
	}
	return *o.ParameterName.Get()
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetParameterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParameterName.Get(), o.ParameterName.IsSet()
}

// HasParameterName returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasParameterName() bool {
	if o != nil && o.ParameterName.IsSet() {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given NullableString and assigns it to the ParameterName field.
func (o *AzureKeyVaultPushTaskStep) SetParameterName(v string) {
	o.ParameterName.Set(&v)
}
// SetParameterNameNil sets the value for ParameterName to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetParameterNameNil() {
	o.ParameterName.Set(nil)
}

// UnsetParameterName ensures that no value is present for ParameterName, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetParameterName() {
	o.ParameterName.Unset()
}

// GetVenueId returns the VenueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetVenueId() string {
	if o == nil || IsNil(o.VenueId.Get()) {
		var ret string
		return ret
	}
	return *o.VenueId.Get()
}

// GetVenueIdOk returns a tuple with the VenueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetVenueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VenueId.Get(), o.VenueId.IsSet()
}

// HasVenueId returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasVenueId() bool {
	if o != nil && o.VenueId.IsSet() {
		return true
	}

	return false
}

// SetVenueId gets a reference to the given NullableString and assigns it to the VenueId field.
func (o *AzureKeyVaultPushTaskStep) SetVenueId(v string) {
	o.VenueId.Set(&v)
}
// SetVenueIdNil sets the value for VenueId to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetVenueIdNil() {
	o.VenueId.Set(nil)
}

// UnsetVenueId ensures that no value is present for VenueId, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetVenueId() {
	o.VenueId.Unset()
}

// GetVenueName returns the VenueName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetVenueName() string {
	if o == nil || IsNil(o.VenueName.Get()) {
		var ret string
		return ret
	}
	return *o.VenueName.Get()
}

// GetVenueNameOk returns a tuple with the VenueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetVenueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VenueName.Get(), o.VenueName.IsSet()
}

// HasVenueName returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasVenueName() bool {
	if o != nil && o.VenueName.IsSet() {
		return true
	}

	return false
}

// SetVenueName gets a reference to the given NullableString and assigns it to the VenueName field.
func (o *AzureKeyVaultPushTaskStep) SetVenueName(v string) {
	o.VenueName.Set(&v)
}
// SetVenueNameNil sets the value for VenueName to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetVenueNameNil() {
	o.VenueName.Set(nil)
}

// UnsetVenueName ensures that no value is present for VenueName, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetVenueName() {
	o.VenueName.Unset()
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetSummary() string {
	if o == nil || IsNil(o.Summary.Get()) {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasSummary() bool {
	if o != nil && o.Summary.IsSet() {
		return true
	}

	return false
}

// SetSummary gets a reference to the given NullableString and assigns it to the Summary field.
func (o *AzureKeyVaultPushTaskStep) SetSummary(v string) {
	o.Summary.Set(&v)
}
// SetSummaryNil sets the value for Summary to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetSummary() {
	o.Summary.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *AzureKeyVaultPushTaskStep) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultPushTaskStep) GetErrorDetail() string {
	if o == nil || IsNil(o.ErrorDetail.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorDetail.Get()
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetErrorDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDetail.Get(), o.ErrorDetail.IsSet()
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *AzureKeyVaultPushTaskStep) HasErrorDetail() bool {
	if o != nil && o.ErrorDetail.IsSet() {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given NullableString and assigns it to the ErrorDetail field.
func (o *AzureKeyVaultPushTaskStep) SetErrorDetail(v string) {
	o.ErrorDetail.Set(&v)
}
// SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil
func (o *AzureKeyVaultPushTaskStep) SetErrorDetailNil() {
	o.ErrorDetail.Set(nil)
}

// UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
func (o *AzureKeyVaultPushTaskStep) UnsetErrorDetail() {
	o.ErrorDetail.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AzureKeyVaultPushTaskStep) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushTaskStep) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AzureKeyVaultPushTaskStep) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AzureKeyVaultPushTaskStep) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushTaskStep) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AzureKeyVaultPushTaskStep) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o AzureKeyVaultPushTaskStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureKeyVaultPushTaskStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	toSerialize["success"] = o.Success
	if o.SuccessDetail.IsSet() {
		toSerialize["success_detail"] = o.SuccessDetail.Get()
	}
	if o.Fqn.IsSet() {
		toSerialize["fqn"] = o.Fqn.Get()
	}
	toSerialize["environment"] = o.Environment.Get()
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if o.EnvironmentName.IsSet() {
		toSerialize["environment_name"] = o.EnvironmentName.Get()
	}
	toSerialize["project"] = o.Project.Get()
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if o.ProjectName.IsSet() {
		toSerialize["project_name"] = o.ProjectName.Get()
	}
	toSerialize["parameter"] = o.Parameter.Get()
	if o.ParameterId.IsSet() {
		toSerialize["parameter_id"] = o.ParameterId.Get()
	}
	if o.ParameterName.IsSet() {
		toSerialize["parameter_name"] = o.ParameterName.Get()
	}
	if o.VenueId.IsSet() {
		toSerialize["venue_id"] = o.VenueId.Get()
	}
	if o.VenueName.IsSet() {
		toSerialize["venue_name"] = o.VenueName.Get()
	}
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorDetail.IsSet() {
		toSerialize["error_detail"] = o.ErrorDetail.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	return toSerialize, nil
}

type NullableAzureKeyVaultPushTaskStep struct {
	value *AzureKeyVaultPushTaskStep
	isSet bool
}

func (v NullableAzureKeyVaultPushTaskStep) Get() *AzureKeyVaultPushTaskStep {
	return v.value
}

func (v *NullableAzureKeyVaultPushTaskStep) Set(val *AzureKeyVaultPushTaskStep) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPushTaskStep) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPushTaskStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPushTaskStep(val *AzureKeyVaultPushTaskStep) *NullableAzureKeyVaultPushTaskStep {
	return &NullableAzureKeyVaultPushTaskStep{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPushTaskStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPushTaskStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


