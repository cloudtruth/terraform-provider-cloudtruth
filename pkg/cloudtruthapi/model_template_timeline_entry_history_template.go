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
	"fmt"
)

// checks if the TemplateTimelineEntryHistoryTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateTimelineEntryHistoryTemplate{}

// TemplateTimelineEntryHistoryTemplate The template record as it was when archived for history.
type TemplateTimelineEntryHistoryTemplate struct {
	Id string `json:"id"`
	LedgerId string `json:"ledger_id"`
	// The parameter name.
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// The content of the template.  Use mustache-style templating delimiters of `{{` and `}}` to reference parameter values by name for substitution into the template result.
	Body *string `json:"body,omitempty"`
}

type _TemplateTimelineEntryHistoryTemplate TemplateTimelineEntryHistoryTemplate

// NewTemplateTimelineEntryHistoryTemplate instantiates a new TemplateTimelineEntryHistoryTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateTimelineEntryHistoryTemplate(id string, ledgerId string, name string) *TemplateTimelineEntryHistoryTemplate {
	this := TemplateTimelineEntryHistoryTemplate{}
	this.Id = id
	this.LedgerId = ledgerId
	this.Name = name
	return &this
}

// NewTemplateTimelineEntryHistoryTemplateWithDefaults instantiates a new TemplateTimelineEntryHistoryTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateTimelineEntryHistoryTemplateWithDefaults() *TemplateTimelineEntryHistoryTemplate {
	this := TemplateTimelineEntryHistoryTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateTimelineEntryHistoryTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryHistoryTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateTimelineEntryHistoryTemplate) SetId(v string) {
	o.Id = v
}

// GetLedgerId returns the LedgerId field value
func (o *TemplateTimelineEntryHistoryTemplate) GetLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryHistoryTemplate) GetLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerId, true
}

// SetLedgerId sets field value
func (o *TemplateTimelineEntryHistoryTemplate) SetLedgerId(v string) {
	o.LedgerId = v
}

// GetName returns the Name field value
func (o *TemplateTimelineEntryHistoryTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryHistoryTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateTimelineEntryHistoryTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateTimelineEntryHistoryTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryHistoryTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateTimelineEntryHistoryTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateTimelineEntryHistoryTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *TemplateTimelineEntryHistoryTemplate) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTimelineEntryHistoryTemplate) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *TemplateTimelineEntryHistoryTemplate) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *TemplateTimelineEntryHistoryTemplate) SetBody(v string) {
	o.Body = &v
}

func (o TemplateTimelineEntryHistoryTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateTimelineEntryHistoryTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ledger_id"] = o.LedgerId
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

func (o *TemplateTimelineEntryHistoryTemplate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"ledger_id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplateTimelineEntryHistoryTemplate := _TemplateTimelineEntryHistoryTemplate{}

	err = json.Unmarshal(bytes, &varTemplateTimelineEntryHistoryTemplate)

	if err != nil {
		return err
	}

	*o = TemplateTimelineEntryHistoryTemplate(varTemplateTimelineEntryHistoryTemplate)

	return err
}

type NullableTemplateTimelineEntryHistoryTemplate struct {
	value *TemplateTimelineEntryHistoryTemplate
	isSet bool
}

func (v NullableTemplateTimelineEntryHistoryTemplate) Get() *TemplateTimelineEntryHistoryTemplate {
	return v.value
}

func (v *NullableTemplateTimelineEntryHistoryTemplate) Set(val *TemplateTimelineEntryHistoryTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateTimelineEntryHistoryTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateTimelineEntryHistoryTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateTimelineEntryHistoryTemplate(val *TemplateTimelineEntryHistoryTemplate) *NullableTemplateTimelineEntryHistoryTemplate {
	return &NullableTemplateTimelineEntryHistoryTemplate{value: val, isSet: true}
}

func (v NullableTemplateTimelineEntryHistoryTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateTimelineEntryHistoryTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


