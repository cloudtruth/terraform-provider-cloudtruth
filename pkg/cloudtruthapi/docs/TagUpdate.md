# TagUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the tag. | [readonly] 
**Name** | **string** | The tag name. Tag names may contain alphanumeric, hyphen, underscore, or period characters. Tag names are case sensitive. The name cannot be modified. | 
**Description** | Pointer to **string** | A description of the tag.  You may find it helpful to document how this tag is used to assist others when they need to maintain software that uses this content. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The point in time this tag represents.  If explicitly set to &#x60;null&#x60; then the current time will be used. | [optional] 

## Methods

### NewTagUpdate

`func NewTagUpdate(id string, name string, ) *TagUpdate`

NewTagUpdate instantiates a new TagUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateWithDefaults

`func NewTagUpdateWithDefaults() *TagUpdate`

NewTagUpdateWithDefaults instantiates a new TagUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TagUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimestamp

`func (o *TagUpdate) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TagUpdate) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TagUpdate) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TagUpdate) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TagUpdate) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TagUpdate) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


