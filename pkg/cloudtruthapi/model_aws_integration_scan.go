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

// AwsIntegrationScan struct for AwsIntegrationScan
type AwsIntegrationScan struct {
	// The AWS region to use.  This region must be enabled in the integration.
	Region NullableAwsRegionEnum `json:"region"`
	// The AWS service to use.  This service must be enabled in the integration.
	Service NullableAwsServiceEnum `json:"service"`
	// Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - ``{{ environment }}`` to identify the environment name   - ``{{ parameter }}`` to identify the parameter name   - ``{{ project }}`` to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the `environment`, `project`, and `parameter`.
	Resource NullableString `json:"resource"`
}

// NewAwsIntegrationScan instantiates a new AwsIntegrationScan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsIntegrationScan(region NullableAwsRegionEnum, service NullableAwsServiceEnum, resource NullableString) *AwsIntegrationScan {
	this := AwsIntegrationScan{}
	this.Region = region
	this.Service = service
	this.Resource = resource
	return &this
}

// NewAwsIntegrationScanWithDefaults instantiates a new AwsIntegrationScan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsIntegrationScanWithDefaults() *AwsIntegrationScan {
	this := AwsIntegrationScan{}
	return &this
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for AwsRegionEnum will be returned
func (o *AwsIntegrationScan) GetRegion() AwsRegionEnum {
	if o == nil || o.Region.Get() == nil {
		var ret AwsRegionEnum
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIntegrationScan) GetRegionOk() (*AwsRegionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *AwsIntegrationScan) SetRegion(v AwsRegionEnum) {
	o.Region.Set(&v)
}

// GetService returns the Service field value
// If the value is explicit nil, the zero value for AwsServiceEnum will be returned
func (o *AwsIntegrationScan) GetService() AwsServiceEnum {
	if o == nil || o.Service.Get() == nil {
		var ret AwsServiceEnum
		return ret
	}

	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIntegrationScan) GetServiceOk() (*AwsServiceEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// SetService sets field value
func (o *AwsIntegrationScan) SetService(v AwsServiceEnum) {
	o.Service.Set(&v)
}

// GetResource returns the Resource field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsIntegrationScan) GetResource() string {
	if o == nil || o.Resource.Get() == nil {
		var ret string
		return ret
	}

	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIntegrationScan) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// SetResource sets field value
func (o *AwsIntegrationScan) SetResource(v string) {
	o.Resource.Set(&v)
}

func (o AwsIntegrationScan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["service"] = o.Service.Get()
	}
	if true {
		toSerialize["resource"] = o.Resource.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsIntegrationScan struct {
	value *AwsIntegrationScan
	isSet bool
}

func (v NullableAwsIntegrationScan) Get() *AwsIntegrationScan {
	return v.value
}

func (v *NullableAwsIntegrationScan) Set(val *AwsIntegrationScan) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsIntegrationScan) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsIntegrationScan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsIntegrationScan(val *AwsIntegrationScan) *NullableAwsIntegrationScan {
	return &NullableAwsIntegrationScan{value: val, isSet: true}
}

func (v NullableAwsIntegrationScan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsIntegrationScan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
