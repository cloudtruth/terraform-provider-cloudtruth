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

// checks if the AzureKeyVaultIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureKeyVaultIntegration{}

// AzureKeyVaultIntegration struct for AzureKeyVaultIntegration
type AzureKeyVaultIntegration struct {
	Url string `json:"url"`
	// The unique identifier for the integration.
	Id string `json:"id"`
	Name string `json:"name"`
	// An optional description for the integration.
	Description *string `json:"description,omitempty"`
	Status StatusEnum `json:"status"`
	// If an error occurs, more details will be available in this field.
	StatusDetail string `json:"status_detail"`
	// The last time the status was evaluated.
	StatusLastCheckedAt NullableTime `json:"status_last_checked_at"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
	Fqn string `json:"fqn"`
	// The type of integration.
	Type string `json:"type"`
	// Allow actions to write to the integration.
	Writable *bool `json:"writable,omitempty"`
	// A list of tags to be set on all integration resources.
	ResourceTags map[string]interface{} `json:"resource_tags,omitempty"`
	// The Azure Key Vault name.
	VaultName string `json:"vault_name"`
	// The Azure Tenant ID.
	TenantId string `json:"tenant_id"`
}

// NewAzureKeyVaultIntegration instantiates a new AzureKeyVaultIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultIntegration(url string, id string, name string, status StatusEnum, statusDetail string, statusLastCheckedAt NullableTime, createdAt time.Time, modifiedAt NullableTime, fqn string, type_ string, vaultName string, tenantId string) *AzureKeyVaultIntegration {
	this := AzureKeyVaultIntegration{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusDetail = statusDetail
	this.StatusLastCheckedAt = statusLastCheckedAt
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Fqn = fqn
	this.Type = type_
	this.VaultName = vaultName
	this.TenantId = tenantId
	return &this
}

// NewAzureKeyVaultIntegrationWithDefaults instantiates a new AzureKeyVaultIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultIntegrationWithDefaults() *AzureKeyVaultIntegration {
	this := AzureKeyVaultIntegration{}
	return &this
}

// GetUrl returns the Url field value
func (o *AzureKeyVaultIntegration) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AzureKeyVaultIntegration) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AzureKeyVaultIntegration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzureKeyVaultIntegration) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AzureKeyVaultIntegration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureKeyVaultIntegration) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureKeyVaultIntegration) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureKeyVaultIntegration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureKeyVaultIntegration) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *AzureKeyVaultIntegration) GetStatus() StatusEnum {
	if o == nil {
		var ret StatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetStatusOk() (*StatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AzureKeyVaultIntegration) SetStatus(v StatusEnum) {
	o.Status = v
}

// GetStatusDetail returns the StatusDetail field value
func (o *AzureKeyVaultIntegration) GetStatusDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetStatusDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDetail, true
}

// SetStatusDetail sets field value
func (o *AzureKeyVaultIntegration) SetStatusDetail(v string) {
	o.StatusDetail = v
}

// GetStatusLastCheckedAt returns the StatusLastCheckedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AzureKeyVaultIntegration) GetStatusLastCheckedAt() time.Time {
	if o == nil || o.StatusLastCheckedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StatusLastCheckedAt.Get()
}

// GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusLastCheckedAt.Get(), o.StatusLastCheckedAt.IsSet()
}

// SetStatusLastCheckedAt sets field value
func (o *AzureKeyVaultIntegration) SetStatusLastCheckedAt(v time.Time) {
	o.StatusLastCheckedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *AzureKeyVaultIntegration) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AzureKeyVaultIntegration) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AzureKeyVaultIntegration) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultIntegration) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AzureKeyVaultIntegration) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// GetFqn returns the Fqn field value
func (o *AzureKeyVaultIntegration) GetFqn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqn, true
}

// SetFqn sets field value
func (o *AzureKeyVaultIntegration) SetFqn(v string) {
	o.Fqn = v
}

// GetType returns the Type field value
func (o *AzureKeyVaultIntegration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AzureKeyVaultIntegration) SetType(v string) {
	o.Type = v
}

// GetWritable returns the Writable field value if set, zero value otherwise.
func (o *AzureKeyVaultIntegration) GetWritable() bool {
	if o == nil || IsNil(o.Writable) {
		var ret bool
		return ret
	}
	return *o.Writable
}

// GetWritableOk returns a tuple with the Writable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetWritableOk() (*bool, bool) {
	if o == nil || IsNil(o.Writable) {
		return nil, false
	}
	return o.Writable, true
}

// HasWritable returns a boolean if a field has been set.
func (o *AzureKeyVaultIntegration) HasWritable() bool {
	if o != nil && !IsNil(o.Writable) {
		return true
	}

	return false
}

// SetWritable gets a reference to the given bool and assigns it to the Writable field.
func (o *AzureKeyVaultIntegration) SetWritable(v bool) {
	o.Writable = &v
}

// GetResourceTags returns the ResourceTags field value if set, zero value otherwise.
func (o *AzureKeyVaultIntegration) GetResourceTags() map[string]interface{} {
	if o == nil || IsNil(o.ResourceTags) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResourceTags
}

// GetResourceTagsOk returns a tuple with the ResourceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetResourceTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResourceTags) {
		return map[string]interface{}{}, false
	}
	return o.ResourceTags, true
}

// HasResourceTags returns a boolean if a field has been set.
func (o *AzureKeyVaultIntegration) HasResourceTags() bool {
	if o != nil && !IsNil(o.ResourceTags) {
		return true
	}

	return false
}

// SetResourceTags gets a reference to the given map[string]interface{} and assigns it to the ResourceTags field.
func (o *AzureKeyVaultIntegration) SetResourceTags(v map[string]interface{}) {
	o.ResourceTags = v
}

// GetVaultName returns the VaultName field value
func (o *AzureKeyVaultIntegration) GetVaultName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultName
}

// GetVaultNameOk returns a tuple with the VaultName field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetVaultNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultName, true
}

// SetVaultName sets field value
func (o *AzureKeyVaultIntegration) SetVaultName(v string) {
	o.VaultName = v
}

// GetTenantId returns the TenantId field value
func (o *AzureKeyVaultIntegration) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultIntegration) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *AzureKeyVaultIntegration) SetTenantId(v string) {
	o.TenantId = v
}

func (o AzureKeyVaultIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureKeyVaultIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["status_detail"] = o.StatusDetail
	toSerialize["status_last_checked_at"] = o.StatusLastCheckedAt.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["modified_at"] = o.ModifiedAt.Get()
	toSerialize["fqn"] = o.Fqn
	toSerialize["type"] = o.Type
	if !IsNil(o.Writable) {
		toSerialize["writable"] = o.Writable
	}
	if !IsNil(o.ResourceTags) {
		toSerialize["resource_tags"] = o.ResourceTags
	}
	toSerialize["vault_name"] = o.VaultName
	toSerialize["tenant_id"] = o.TenantId
	return toSerialize, nil
}

type NullableAzureKeyVaultIntegration struct {
	value *AzureKeyVaultIntegration
	isSet bool
}

func (v NullableAzureKeyVaultIntegration) Get() *AzureKeyVaultIntegration {
	return v.value
}

func (v *NullableAzureKeyVaultIntegration) Set(val *AzureKeyVaultIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultIntegration(val *AzureKeyVaultIntegration) *NullableAzureKeyVaultIntegration {
	return &NullableAzureKeyVaultIntegration{value: val, isSet: true}
}

func (v NullableAzureKeyVaultIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


