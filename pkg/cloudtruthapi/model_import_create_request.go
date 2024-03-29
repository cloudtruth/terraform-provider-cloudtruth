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
)

// checks if the ImportCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportCreateRequest{}

// ImportCreateRequest struct for ImportCreateRequest
type ImportCreateRequest struct {
	// Project name or identifier
	Project string `json:"project"`
	// Environment name or identifier
	Environment NullableString `json:"environment,omitempty"`
	// Text to turn into variables
	Body string `json:"body"`
	// Parameter names that should be secrets
	Secrets []string `json:"secrets"`
	// Parameter names to be ignored
	Ignore []string `json:"ignore"`
	// Create the project (if it does not exist)
	AddProject *bool `json:"add_project,omitempty"`
	// Create the environment (if it does not exist)
	AddEnvironment *bool `json:"add_environment,omitempty"`
	// Create the parameters (if they do not exist)
	AddParameters *bool `json:"add_parameters,omitempty"`
	// Create the environment value override (if they do not exist)
	AddOverrides *bool `json:"add_overrides,omitempty"`
	// Skip adding a parameter override if it is the same
	InheritOnSame *bool `json:"inherit_on_same,omitempty"`
}

// NewImportCreateRequest instantiates a new ImportCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportCreateRequest(project string, body string, secrets []string, ignore []string) *ImportCreateRequest {
	this := ImportCreateRequest{}
	this.Project = project
	this.Body = body
	this.Secrets = secrets
	this.Ignore = ignore
	var addProject bool = true
	this.AddProject = &addProject
	var addEnvironment bool = true
	this.AddEnvironment = &addEnvironment
	var addParameters bool = true
	this.AddParameters = &addParameters
	var addOverrides bool = true
	this.AddOverrides = &addOverrides
	var inheritOnSame bool = true
	this.InheritOnSame = &inheritOnSame
	return &this
}

// NewImportCreateRequestWithDefaults instantiates a new ImportCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportCreateRequestWithDefaults() *ImportCreateRequest {
	this := ImportCreateRequest{}
	var addProject bool = true
	this.AddProject = &addProject
	var addEnvironment bool = true
	this.AddEnvironment = &addEnvironment
	var addParameters bool = true
	this.AddParameters = &addParameters
	var addOverrides bool = true
	this.AddOverrides = &addOverrides
	var inheritOnSame bool = true
	this.InheritOnSame = &inheritOnSame
	return &this
}

// GetProject returns the Project field value
func (o *ImportCreateRequest) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ImportCreateRequest) SetProject(v string) {
	o.Project = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportCreateRequest) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportCreateRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ImportCreateRequest) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *ImportCreateRequest) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *ImportCreateRequest) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *ImportCreateRequest) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetBody returns the Body field value
func (o *ImportCreateRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ImportCreateRequest) SetBody(v string) {
	o.Body = v
}

// GetSecrets returns the Secrets field value
func (o *ImportCreateRequest) GetSecrets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetSecretsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *ImportCreateRequest) SetSecrets(v []string) {
	o.Secrets = v
}

// GetIgnore returns the Ignore field value
func (o *ImportCreateRequest) GetIgnore() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ignore
}

// GetIgnoreOk returns a tuple with the Ignore field value
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetIgnoreOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ignore, true
}

// SetIgnore sets field value
func (o *ImportCreateRequest) SetIgnore(v []string) {
	o.Ignore = v
}

// GetAddProject returns the AddProject field value if set, zero value otherwise.
func (o *ImportCreateRequest) GetAddProject() bool {
	if o == nil || IsNil(o.AddProject) {
		var ret bool
		return ret
	}
	return *o.AddProject
}

// GetAddProjectOk returns a tuple with the AddProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetAddProjectOk() (*bool, bool) {
	if o == nil || IsNil(o.AddProject) {
		return nil, false
	}
	return o.AddProject, true
}

// HasAddProject returns a boolean if a field has been set.
func (o *ImportCreateRequest) HasAddProject() bool {
	if o != nil && !IsNil(o.AddProject) {
		return true
	}

	return false
}

// SetAddProject gets a reference to the given bool and assigns it to the AddProject field.
func (o *ImportCreateRequest) SetAddProject(v bool) {
	o.AddProject = &v
}

// GetAddEnvironment returns the AddEnvironment field value if set, zero value otherwise.
func (o *ImportCreateRequest) GetAddEnvironment() bool {
	if o == nil || IsNil(o.AddEnvironment) {
		var ret bool
		return ret
	}
	return *o.AddEnvironment
}

// GetAddEnvironmentOk returns a tuple with the AddEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetAddEnvironmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AddEnvironment) {
		return nil, false
	}
	return o.AddEnvironment, true
}

// HasAddEnvironment returns a boolean if a field has been set.
func (o *ImportCreateRequest) HasAddEnvironment() bool {
	if o != nil && !IsNil(o.AddEnvironment) {
		return true
	}

	return false
}

// SetAddEnvironment gets a reference to the given bool and assigns it to the AddEnvironment field.
func (o *ImportCreateRequest) SetAddEnvironment(v bool) {
	o.AddEnvironment = &v
}

// GetAddParameters returns the AddParameters field value if set, zero value otherwise.
func (o *ImportCreateRequest) GetAddParameters() bool {
	if o == nil || IsNil(o.AddParameters) {
		var ret bool
		return ret
	}
	return *o.AddParameters
}

// GetAddParametersOk returns a tuple with the AddParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetAddParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.AddParameters) {
		return nil, false
	}
	return o.AddParameters, true
}

// HasAddParameters returns a boolean if a field has been set.
func (o *ImportCreateRequest) HasAddParameters() bool {
	if o != nil && !IsNil(o.AddParameters) {
		return true
	}

	return false
}

// SetAddParameters gets a reference to the given bool and assigns it to the AddParameters field.
func (o *ImportCreateRequest) SetAddParameters(v bool) {
	o.AddParameters = &v
}

// GetAddOverrides returns the AddOverrides field value if set, zero value otherwise.
func (o *ImportCreateRequest) GetAddOverrides() bool {
	if o == nil || IsNil(o.AddOverrides) {
		var ret bool
		return ret
	}
	return *o.AddOverrides
}

// GetAddOverridesOk returns a tuple with the AddOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetAddOverridesOk() (*bool, bool) {
	if o == nil || IsNil(o.AddOverrides) {
		return nil, false
	}
	return o.AddOverrides, true
}

// HasAddOverrides returns a boolean if a field has been set.
func (o *ImportCreateRequest) HasAddOverrides() bool {
	if o != nil && !IsNil(o.AddOverrides) {
		return true
	}

	return false
}

// SetAddOverrides gets a reference to the given bool and assigns it to the AddOverrides field.
func (o *ImportCreateRequest) SetAddOverrides(v bool) {
	o.AddOverrides = &v
}

// GetInheritOnSame returns the InheritOnSame field value if set, zero value otherwise.
func (o *ImportCreateRequest) GetInheritOnSame() bool {
	if o == nil || IsNil(o.InheritOnSame) {
		var ret bool
		return ret
	}
	return *o.InheritOnSame
}

// GetInheritOnSameOk returns a tuple with the InheritOnSame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportCreateRequest) GetInheritOnSameOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritOnSame) {
		return nil, false
	}
	return o.InheritOnSame, true
}

// HasInheritOnSame returns a boolean if a field has been set.
func (o *ImportCreateRequest) HasInheritOnSame() bool {
	if o != nil && !IsNil(o.InheritOnSame) {
		return true
	}

	return false
}

// SetInheritOnSame gets a reference to the given bool and assigns it to the InheritOnSame field.
func (o *ImportCreateRequest) SetInheritOnSame(v bool) {
	o.InheritOnSame = &v
}

func (o ImportCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project"] = o.Project
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	toSerialize["body"] = o.Body
	toSerialize["secrets"] = o.Secrets
	toSerialize["ignore"] = o.Ignore
	if !IsNil(o.AddProject) {
		toSerialize["add_project"] = o.AddProject
	}
	if !IsNil(o.AddEnvironment) {
		toSerialize["add_environment"] = o.AddEnvironment
	}
	if !IsNil(o.AddParameters) {
		toSerialize["add_parameters"] = o.AddParameters
	}
	if !IsNil(o.AddOverrides) {
		toSerialize["add_overrides"] = o.AddOverrides
	}
	if !IsNil(o.InheritOnSame) {
		toSerialize["inherit_on_same"] = o.InheritOnSame
	}
	return toSerialize, nil
}

type NullableImportCreateRequest struct {
	value *ImportCreateRequest
	isSet bool
}

func (v NullableImportCreateRequest) Get() *ImportCreateRequest {
	return v.value
}

func (v *NullableImportCreateRequest) Set(val *ImportCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportCreateRequest(val *ImportCreateRequest) *NullableImportCreateRequest {
	return &NullableImportCreateRequest{value: val, isSet: true}
}

func (v NullableImportCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


