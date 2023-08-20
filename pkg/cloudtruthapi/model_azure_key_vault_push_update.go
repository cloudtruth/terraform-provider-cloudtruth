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

// checks if the AzureKeyVaultPushUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureKeyVaultPushUpdate{}

// AzureKeyVaultPushUpdate Update a push.
type AzureKeyVaultPushUpdate struct {
	// The action name.
	Name string `json:"name"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	// Projects that are included in the push.
	Projects []string `json:"projects"`
	// Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push.
	Tags []string `json:"tags"`
	// Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - ``{{ environment }}`` to insert the environment name   - ``{{ parameter }}`` to insert the parameter name   - ``{{ project }}`` to insert the project name   - ``{{ push }}`` to insert the push name   - ``{{ tag }}`` to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the `project` substitution is required.  If you include multiple tags from different environments in the push, the `environment` substitution is required.  If you include multiple tags from the same environment in the push, the `tag` substitution is required.  In all cases, the `parameter` substitution is always required.
	Resource NullableString `json:"resource"`
	// When set to dry-run mode an action will report the changes that it would have made in task steps, however those changes are not actually performed.
	DryRun *bool `json:"dry_run,omitempty"`
	// Normally, push will check to see if it originated the values in the destination before making changes to them.  Forcing a push disables the ownership check.
	Force *bool `json:"force,omitempty"`
	// Normally, push will process all parameters including those that flow in from project dependencies.  Declaring a push as `local` will cause it to only process the parameters defined in the selected projects.
	Local *bool `json:"local,omitempty"`
	// This setting allows parameters (non-secrets) to be pushed to a destination that only supports storing secrets.  This may increase your overall cost from the cloud provider as some cloud providers charge a premium for secrets-only storage.
	CoerceParameters *bool `json:"coerce_parameters,omitempty"`
	// Include parameters (non-secrets) in the values being pushed.  This setting requires the destination to support parameters or for the `coerce_parameters` flag to be enabled, otherwise the push will fail.
	IncludeParameters *bool `json:"include_parameters,omitempty"`
	// Include secrets in the values being pushed.  This setting requires the destination to support secrets, otherwise the push will fail.
	IncludeSecrets *bool `json:"include_secrets,omitempty"`
}

// NewAzureKeyVaultPushUpdate instantiates a new AzureKeyVaultPushUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultPushUpdate(name string, projects []string, tags []string, resource NullableString) *AzureKeyVaultPushUpdate {
	this := AzureKeyVaultPushUpdate{}
	this.Name = name
	this.Projects = projects
	this.Tags = tags
	this.Resource = resource
	return &this
}

// NewAzureKeyVaultPushUpdateWithDefaults instantiates a new AzureKeyVaultPushUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultPushUpdateWithDefaults() *AzureKeyVaultPushUpdate {
	this := AzureKeyVaultPushUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *AzureKeyVaultPushUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureKeyVaultPushUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureKeyVaultPushUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetProjects returns the Projects field value
func (o *AzureKeyVaultPushUpdate) GetProjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetProjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *AzureKeyVaultPushUpdate) SetProjects(v []string) {
	o.Projects = v
}

// GetTags returns the Tags field value
func (o *AzureKeyVaultPushUpdate) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *AzureKeyVaultPushUpdate) SetTags(v []string) {
	o.Tags = v
}

// GetResource returns the Resource field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureKeyVaultPushUpdate) GetResource() string {
	if o == nil || o.Resource.Get() == nil {
		var ret string
		return ret
	}

	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultPushUpdate) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// SetResource sets field value
func (o *AzureKeyVaultPushUpdate) SetResource(v string) {
	o.Resource.Set(&v)
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *AzureKeyVaultPushUpdate) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *AzureKeyVaultPushUpdate) SetForce(v bool) {
	o.Force = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetLocal() bool {
	if o == nil || IsNil(o.Local) {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.Local) {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasLocal() bool {
	if o != nil && !IsNil(o.Local) {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *AzureKeyVaultPushUpdate) SetLocal(v bool) {
	o.Local = &v
}

// GetCoerceParameters returns the CoerceParameters field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetCoerceParameters() bool {
	if o == nil || IsNil(o.CoerceParameters) {
		var ret bool
		return ret
	}
	return *o.CoerceParameters
}

// GetCoerceParametersOk returns a tuple with the CoerceParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetCoerceParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.CoerceParameters) {
		return nil, false
	}
	return o.CoerceParameters, true
}

// HasCoerceParameters returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasCoerceParameters() bool {
	if o != nil && !IsNil(o.CoerceParameters) {
		return true
	}

	return false
}

// SetCoerceParameters gets a reference to the given bool and assigns it to the CoerceParameters field.
func (o *AzureKeyVaultPushUpdate) SetCoerceParameters(v bool) {
	o.CoerceParameters = &v
}

// GetIncludeParameters returns the IncludeParameters field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetIncludeParameters() bool {
	if o == nil || IsNil(o.IncludeParameters) {
		var ret bool
		return ret
	}
	return *o.IncludeParameters
}

// GetIncludeParametersOk returns a tuple with the IncludeParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetIncludeParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeParameters) {
		return nil, false
	}
	return o.IncludeParameters, true
}

// HasIncludeParameters returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasIncludeParameters() bool {
	if o != nil && !IsNil(o.IncludeParameters) {
		return true
	}

	return false
}

// SetIncludeParameters gets a reference to the given bool and assigns it to the IncludeParameters field.
func (o *AzureKeyVaultPushUpdate) SetIncludeParameters(v bool) {
	o.IncludeParameters = &v
}

// GetIncludeSecrets returns the IncludeSecrets field value if set, zero value otherwise.
func (o *AzureKeyVaultPushUpdate) GetIncludeSecrets() bool {
	if o == nil || IsNil(o.IncludeSecrets) {
		var ret bool
		return ret
	}
	return *o.IncludeSecrets
}

// GetIncludeSecretsOk returns a tuple with the IncludeSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultPushUpdate) GetIncludeSecretsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSecrets) {
		return nil, false
	}
	return o.IncludeSecrets, true
}

// HasIncludeSecrets returns a boolean if a field has been set.
func (o *AzureKeyVaultPushUpdate) HasIncludeSecrets() bool {
	if o != nil && !IsNil(o.IncludeSecrets) {
		return true
	}

	return false
}

// SetIncludeSecrets gets a reference to the given bool and assigns it to the IncludeSecrets field.
func (o *AzureKeyVaultPushUpdate) SetIncludeSecrets(v bool) {
	o.IncludeSecrets = &v
}

func (o AzureKeyVaultPushUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureKeyVaultPushUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["projects"] = o.Projects
	toSerialize["tags"] = o.Tags
	toSerialize["resource"] = o.Resource.Get()
	if !IsNil(o.DryRun) {
		toSerialize["dry_run"] = o.DryRun
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.Local) {
		toSerialize["local"] = o.Local
	}
	if !IsNil(o.CoerceParameters) {
		toSerialize["coerce_parameters"] = o.CoerceParameters
	}
	if !IsNil(o.IncludeParameters) {
		toSerialize["include_parameters"] = o.IncludeParameters
	}
	if !IsNil(o.IncludeSecrets) {
		toSerialize["include_secrets"] = o.IncludeSecrets
	}
	return toSerialize, nil
}

type NullableAzureKeyVaultPushUpdate struct {
	value *AzureKeyVaultPushUpdate
	isSet bool
}

func (v NullableAzureKeyVaultPushUpdate) Get() *AzureKeyVaultPushUpdate {
	return v.value
}

func (v *NullableAzureKeyVaultPushUpdate) Set(val *AzureKeyVaultPushUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultPushUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultPushUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultPushUpdate(val *AzureKeyVaultPushUpdate) *NullableAzureKeyVaultPushUpdate {
	return &NullableAzureKeyVaultPushUpdate{value: val, isSet: true}
}

func (v NullableAzureKeyVaultPushUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultPushUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


