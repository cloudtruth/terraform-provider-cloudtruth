# TemplateTimelineEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryDate** | **time.Time** |  | 
**HistoryType** | [**NullableHistoryTypeEnum**](HistoryTypeEnum.md) |  | [readonly] 
**HistoryUser** | Pointer to **NullableString** | The unique identifier of a user. | [optional] 
**HistoryTemplate** | [**TemplateTimelineEntryHistoryTemplate**](TemplateTimelineEntryHistoryTemplate.md) |  | 

## Methods

### NewTemplateTimelineEntry

`func NewTemplateTimelineEntry(historyDate time.Time, historyType NullableHistoryTypeEnum, historyTemplate TemplateTimelineEntryHistoryTemplate, ) *TemplateTimelineEntry`

NewTemplateTimelineEntry instantiates a new TemplateTimelineEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateTimelineEntryWithDefaults

`func NewTemplateTimelineEntryWithDefaults() *TemplateTimelineEntry`

NewTemplateTimelineEntryWithDefaults instantiates a new TemplateTimelineEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryDate

`func (o *TemplateTimelineEntry) GetHistoryDate() time.Time`

GetHistoryDate returns the HistoryDate field if non-nil, zero value otherwise.

### GetHistoryDateOk

`func (o *TemplateTimelineEntry) GetHistoryDateOk() (*time.Time, bool)`

GetHistoryDateOk returns a tuple with the HistoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryDate

`func (o *TemplateTimelineEntry) SetHistoryDate(v time.Time)`

SetHistoryDate sets HistoryDate field to given value.


### GetHistoryType

`func (o *TemplateTimelineEntry) GetHistoryType() HistoryTypeEnum`

GetHistoryType returns the HistoryType field if non-nil, zero value otherwise.

### GetHistoryTypeOk

`func (o *TemplateTimelineEntry) GetHistoryTypeOk() (*HistoryTypeEnum, bool)`

GetHistoryTypeOk returns a tuple with the HistoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryType

`func (o *TemplateTimelineEntry) SetHistoryType(v HistoryTypeEnum)`

SetHistoryType sets HistoryType field to given value.


### SetHistoryTypeNil

`func (o *TemplateTimelineEntry) SetHistoryTypeNil(b bool)`

 SetHistoryTypeNil sets the value for HistoryType to be an explicit nil

### UnsetHistoryType
`func (o *TemplateTimelineEntry) UnsetHistoryType()`

UnsetHistoryType ensures that no value is present for HistoryType, not even an explicit nil
### GetHistoryUser

`func (o *TemplateTimelineEntry) GetHistoryUser() string`

GetHistoryUser returns the HistoryUser field if non-nil, zero value otherwise.

### GetHistoryUserOk

`func (o *TemplateTimelineEntry) GetHistoryUserOk() (*string, bool)`

GetHistoryUserOk returns a tuple with the HistoryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryUser

`func (o *TemplateTimelineEntry) SetHistoryUser(v string)`

SetHistoryUser sets HistoryUser field to given value.

### HasHistoryUser

`func (o *TemplateTimelineEntry) HasHistoryUser() bool`

HasHistoryUser returns a boolean if a field has been set.

### SetHistoryUserNil

`func (o *TemplateTimelineEntry) SetHistoryUserNil(b bool)`

 SetHistoryUserNil sets the value for HistoryUser to be an explicit nil

### UnsetHistoryUser
`func (o *TemplateTimelineEntry) UnsetHistoryUser()`

UnsetHistoryUser ensures that no value is present for HistoryUser, not even an explicit nil
### GetHistoryTemplate

`func (o *TemplateTimelineEntry) GetHistoryTemplate() TemplateTimelineEntryHistoryTemplate`

GetHistoryTemplate returns the HistoryTemplate field if non-nil, zero value otherwise.

### GetHistoryTemplateOk

`func (o *TemplateTimelineEntry) GetHistoryTemplateOk() (*TemplateTimelineEntryHistoryTemplate, bool)`

GetHistoryTemplateOk returns a tuple with the HistoryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTemplate

`func (o *TemplateTimelineEntry) SetHistoryTemplate(v TemplateTimelineEntryHistoryTemplate)`

SetHistoryTemplate sets HistoryTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


