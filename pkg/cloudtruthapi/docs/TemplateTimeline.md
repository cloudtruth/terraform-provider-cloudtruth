# TemplateTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The number of records in this response. | 
**NextAsOf** | Pointer to **time.Time** | If present, additional history can be retrieved using this timestamp in the next call for the as_of query parameter value. | [optional] 
**Results** | [**[]TemplateTimelineEntry**](TemplateTimelineEntry.md) |  | 

## Methods

### NewTemplateTimeline

`func NewTemplateTimeline(count int32, results []TemplateTimelineEntry, ) *TemplateTimeline`

NewTemplateTimeline instantiates a new TemplateTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateTimelineWithDefaults

`func NewTemplateTimelineWithDefaults() *TemplateTimeline`

NewTemplateTimelineWithDefaults instantiates a new TemplateTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TemplateTimeline) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TemplateTimeline) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TemplateTimeline) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNextAsOf

`func (o *TemplateTimeline) GetNextAsOf() time.Time`

GetNextAsOf returns the NextAsOf field if non-nil, zero value otherwise.

### GetNextAsOfOk

`func (o *TemplateTimeline) GetNextAsOfOk() (*time.Time, bool)`

GetNextAsOfOk returns a tuple with the NextAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAsOf

`func (o *TemplateTimeline) SetNextAsOf(v time.Time)`

SetNextAsOf sets NextAsOf field to given value.

### HasNextAsOf

`func (o *TemplateTimeline) HasNextAsOf() bool`

HasNextAsOf returns a boolean if a field has been set.

### GetResults

`func (o *TemplateTimeline) GetResults() []TemplateTimelineEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TemplateTimeline) GetResultsOk() (*[]TemplateTimelineEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TemplateTimeline) SetResults(v []TemplateTimelineEntry)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


