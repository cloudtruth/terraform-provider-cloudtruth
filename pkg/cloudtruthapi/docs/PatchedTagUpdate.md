# PatchedTagUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the tag. | [optional] [readonly] 
**Name** | Pointer to **string** | The tag name. Tag names may contain alphanumeric, hyphen, underscore, or period characters. Tag names are case sensitive. The name cannot be modified. | [optional] 
**Description** | Pointer to **string** | A description of the tag.  You may find it helpful to document how this tag is used to assist others when they need to maintain software that uses this content. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The point in time this tag represents.  If explicitly set to &#x60;null&#x60; then the current time will be used. | [optional] 

## Methods

### NewPatchedTagUpdate

`func NewPatchedTagUpdate() *PatchedTagUpdate`

NewPatchedTagUpdate instantiates a new PatchedTagUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTagUpdateWithDefaults

`func NewPatchedTagUpdateWithDefaults() *PatchedTagUpdate`

NewPatchedTagUpdateWithDefaults instantiates a new PatchedTagUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTagUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTagUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTagUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTagUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedTagUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTagUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTagUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTagUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTagUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTagUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTagUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTagUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimestamp

`func (o *PatchedTagUpdate) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PatchedTagUpdate) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PatchedTagUpdate) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PatchedTagUpdate) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PatchedTagUpdate) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PatchedTagUpdate) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


