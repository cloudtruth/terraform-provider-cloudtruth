# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the tag. | [readonly] 
**Name** | **string** | The tag name. Tag names may contain alphanumeric, hyphen, underscore, or period characters. Tag names are case sensitive. The name cannot be modified. | 
**Description** | Pointer to **string** | A description of the tag.  You may find it helpful to document how this tag is used to assist others when they need to maintain software that uses this content. | [optional] 
**Timestamp** | **time.Time** | The point in time this tag represents. | 
**Pushes** | [**[]AwsPush**](AwsPush.md) | Deprecated. Only shows pushes for aws integrations in /api/v1/. | [readonly] 
**PushUrls** | **[]string** | Push actions associated with the tag. | [readonly] 
**Usage** | [**TagUsage**](TagUsage.md) |  | 

## Methods

### NewTag

`func NewTag(url string, id string, name string, timestamp time.Time, pushes []AwsPush, pushUrls []string, usage TagUsage, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Tag) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Tag) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Tag) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Tag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Tag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimestamp

`func (o *Tag) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Tag) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Tag) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetPushes

`func (o *Tag) GetPushes() []AwsPush`

GetPushes returns the Pushes field if non-nil, zero value otherwise.

### GetPushesOk

`func (o *Tag) GetPushesOk() (*[]AwsPush, bool)`

GetPushesOk returns a tuple with the Pushes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushes

`func (o *Tag) SetPushes(v []AwsPush)`

SetPushes sets Pushes field to given value.


### GetPushUrls

`func (o *Tag) GetPushUrls() []string`

GetPushUrls returns the PushUrls field if non-nil, zero value otherwise.

### GetPushUrlsOk

`func (o *Tag) GetPushUrlsOk() (*[]string, bool)`

GetPushUrlsOk returns a tuple with the PushUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushUrls

`func (o *Tag) SetPushUrls(v []string)`

SetPushUrls sets PushUrls field to given value.


### GetUsage

`func (o *Tag) GetUsage() TagUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Tag) GetUsageOk() (*TagUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Tag) SetUsage(v TagUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


