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

// checks if the PatchedAwsPushUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedAwsPushUpdate{}

// PatchedAwsPushUpdate Update a push.  The `region` and `service` cannot be changed on an existing push.
type PatchedAwsPushUpdate struct {
	// The action name.
	Name *string `json:"name,omitempty"`
	// The optional description for the action.
	Description *string `json:"description,omitempty"`
	// Projects that are included in the push.
	Projects []string `json:"projects,omitempty"`
	// Tags are used to select parameters by environment from the projects included in the push.  You cannot have two tags from the same environment in the same push.
	Tags []string `json:"tags,omitempty"`
	// Defines a path through the integration to the location where values will be pushed.  The following mustache-style substitutions can be used in the string:    - ``{{ environment }}`` to insert the environment name   - ``{{ parameter }}`` to insert the parameter name   - ``{{ project }}`` to insert the project name   - ``{{ push }}`` to insert the push name   - ``{{ tag }}`` to insert the tag name  We recommend that you use project, environment, and parameter at a minimum to disambiguate your pushed resource identifiers.  If you include multiple projects in the push, the `project` substitution is required.  If you include multiple tags from different environments in the push, the `environment` substitution is required.  If you include multiple tags from the same environment in the push, the `tag` substitution is required.  In all cases, the `parameter` substitution is always required.
	Resource NullableString `json:"resource,omitempty"`
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
	// Include templates in the values being pushed.
	IncludeTemplates *bool `json:"include_templates,omitempty"`
}

// NewPatchedAwsPushUpdate instantiates a new PatchedAwsPushUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAwsPushUpdate() *PatchedAwsPushUpdate {
	this := PatchedAwsPushUpdate{}
	return &this
}

// NewPatchedAwsPushUpdateWithDefaults instantiates a new PatchedAwsPushUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAwsPushUpdateWithDefaults() *PatchedAwsPushUpdate {
	this := PatchedAwsPushUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAwsPushUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedAwsPushUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetProjects() []string {
	if o == nil || IsNil(o.Projects) {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *PatchedAwsPushUpdate) SetProjects(v []string) {
	o.Projects = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PatchedAwsPushUpdate) SetTags(v []string) {
	o.Tags = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAwsPushUpdate) GetResource() string {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret string
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAwsPushUpdate) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableString and assigns it to the Resource field.
func (o *PatchedAwsPushUpdate) SetResource(v string) {
	o.Resource.Set(&v)
}
// SetResourceNil sets the value for Resource to be an explicit nil
func (o *PatchedAwsPushUpdate) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *PatchedAwsPushUpdate) UnsetResource() {
	o.Resource.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *PatchedAwsPushUpdate) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *PatchedAwsPushUpdate) SetForce(v bool) {
	o.Force = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetLocal() bool {
	if o == nil || IsNil(o.Local) {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.Local) {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasLocal() bool {
	if o != nil && !IsNil(o.Local) {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *PatchedAwsPushUpdate) SetLocal(v bool) {
	o.Local = &v
}

// GetCoerceParameters returns the CoerceParameters field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetCoerceParameters() bool {
	if o == nil || IsNil(o.CoerceParameters) {
		var ret bool
		return ret
	}
	return *o.CoerceParameters
}

// GetCoerceParametersOk returns a tuple with the CoerceParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetCoerceParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.CoerceParameters) {
		return nil, false
	}
	return o.CoerceParameters, true
}

// HasCoerceParameters returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasCoerceParameters() bool {
	if o != nil && !IsNil(o.CoerceParameters) {
		return true
	}

	return false
}

// SetCoerceParameters gets a reference to the given bool and assigns it to the CoerceParameters field.
func (o *PatchedAwsPushUpdate) SetCoerceParameters(v bool) {
	o.CoerceParameters = &v
}

// GetIncludeParameters returns the IncludeParameters field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetIncludeParameters() bool {
	if o == nil || IsNil(o.IncludeParameters) {
		var ret bool
		return ret
	}
	return *o.IncludeParameters
}

// GetIncludeParametersOk returns a tuple with the IncludeParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetIncludeParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeParameters) {
		return nil, false
	}
	return o.IncludeParameters, true
}

// HasIncludeParameters returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasIncludeParameters() bool {
	if o != nil && !IsNil(o.IncludeParameters) {
		return true
	}

	return false
}

// SetIncludeParameters gets a reference to the given bool and assigns it to the IncludeParameters field.
func (o *PatchedAwsPushUpdate) SetIncludeParameters(v bool) {
	o.IncludeParameters = &v
}

// GetIncludeSecrets returns the IncludeSecrets field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetIncludeSecrets() bool {
	if o == nil || IsNil(o.IncludeSecrets) {
		var ret bool
		return ret
	}
	return *o.IncludeSecrets
}

// GetIncludeSecretsOk returns a tuple with the IncludeSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetIncludeSecretsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSecrets) {
		return nil, false
	}
	return o.IncludeSecrets, true
}

// HasIncludeSecrets returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasIncludeSecrets() bool {
	if o != nil && !IsNil(o.IncludeSecrets) {
		return true
	}

	return false
}

// SetIncludeSecrets gets a reference to the given bool and assigns it to the IncludeSecrets field.
func (o *PatchedAwsPushUpdate) SetIncludeSecrets(v bool) {
	o.IncludeSecrets = &v
}

// GetIncludeTemplates returns the IncludeTemplates field value if set, zero value otherwise.
func (o *PatchedAwsPushUpdate) GetIncludeTemplates() bool {
	if o == nil || IsNil(o.IncludeTemplates) {
		var ret bool
		return ret
	}
	return *o.IncludeTemplates
}

// GetIncludeTemplatesOk returns a tuple with the IncludeTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAwsPushUpdate) GetIncludeTemplatesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTemplates) {
		return nil, false
	}
	return o.IncludeTemplates, true
}

// HasIncludeTemplates returns a boolean if a field has been set.
func (o *PatchedAwsPushUpdate) HasIncludeTemplates() bool {
	if o != nil && !IsNil(o.IncludeTemplates) {
		return true
	}

	return false
}

// SetIncludeTemplates gets a reference to the given bool and assigns it to the IncludeTemplates field.
func (o *PatchedAwsPushUpdate) SetIncludeTemplates(v bool) {
	o.IncludeTemplates = &v
}

func (o PatchedAwsPushUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedAwsPushUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.Resource.IsSet() {
		toSerialize["resource"] = o.Resource.Get()
	}
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
	if !IsNil(o.IncludeTemplates) {
		toSerialize["include_templates"] = o.IncludeTemplates
	}
	return toSerialize, nil
}

type NullablePatchedAwsPushUpdate struct {
	value *PatchedAwsPushUpdate
	isSet bool
}

func (v NullablePatchedAwsPushUpdate) Get() *PatchedAwsPushUpdate {
	return v.value
}

func (v *NullablePatchedAwsPushUpdate) Set(val *PatchedAwsPushUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAwsPushUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAwsPushUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAwsPushUpdate(val *PatchedAwsPushUpdate) *NullablePatchedAwsPushUpdate {
	return &NullablePatchedAwsPushUpdate{value: val, isSet: true}
}

func (v NullablePatchedAwsPushUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAwsPushUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


