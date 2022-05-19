# DiscoveryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matched** | [**map[string]DiscoveredContent**](DiscoveredContent.md) |  | 
**Skipped** | **map[string]string** |  | 

## Methods

### NewDiscoveryResult

`func NewDiscoveryResult(matched map[string]DiscoveredContent, skipped map[string]string, ) *DiscoveryResult`

NewDiscoveryResult instantiates a new DiscoveryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryResultWithDefaults

`func NewDiscoveryResultWithDefaults() *DiscoveryResult`

NewDiscoveryResultWithDefaults instantiates a new DiscoveryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatched

`func (o *DiscoveryResult) GetMatched() map[string]DiscoveredContent`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *DiscoveryResult) GetMatchedOk() (*map[string]DiscoveredContent, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *DiscoveryResult) SetMatched(v map[string]DiscoveredContent)`

SetMatched sets Matched field to given value.


### GetSkipped

`func (o *DiscoveryResult) GetSkipped() map[string]string`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *DiscoveryResult) GetSkippedOk() (*map[string]string, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *DiscoveryResult) SetSkipped(v map[string]string)`

SetSkipped sets Skipped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


