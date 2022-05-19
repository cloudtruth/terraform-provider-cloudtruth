# TagUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastRead** | **NullableTime** | The last time a configuration was retrieved with this tag. | 
**LastReadBy** | Pointer to **string** | The last user (id) to use this tag to read configuration. | [optional] 
**TotalReads** | **int32** | The number of times the tag has been used to read configuration. | 

## Methods

### NewTagUsage

`func NewTagUsage(lastRead NullableTime, totalReads int32, ) *TagUsage`

NewTagUsage instantiates a new TagUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUsageWithDefaults

`func NewTagUsageWithDefaults() *TagUsage`

NewTagUsageWithDefaults instantiates a new TagUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastRead

`func (o *TagUsage) GetLastRead() time.Time`

GetLastRead returns the LastRead field if non-nil, zero value otherwise.

### GetLastReadOk

`func (o *TagUsage) GetLastReadOk() (*time.Time, bool)`

GetLastReadOk returns a tuple with the LastRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRead

`func (o *TagUsage) SetLastRead(v time.Time)`

SetLastRead sets LastRead field to given value.


### SetLastReadNil

`func (o *TagUsage) SetLastReadNil(b bool)`

 SetLastReadNil sets the value for LastRead to be an explicit nil

### UnsetLastRead
`func (o *TagUsage) UnsetLastRead()`

UnsetLastRead ensures that no value is present for LastRead, not even an explicit nil
### GetLastReadBy

`func (o *TagUsage) GetLastReadBy() string`

GetLastReadBy returns the LastReadBy field if non-nil, zero value otherwise.

### GetLastReadByOk

`func (o *TagUsage) GetLastReadByOk() (*string, bool)`

GetLastReadByOk returns a tuple with the LastReadBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadBy

`func (o *TagUsage) SetLastReadBy(v string)`

SetLastReadBy sets LastReadBy field to given value.

### HasLastReadBy

`func (o *TagUsage) HasLastReadBy() bool`

HasLastReadBy returns a boolean if a field has been set.

### GetTotalReads

`func (o *TagUsage) GetTotalReads() int32`

GetTotalReads returns the TotalReads field if non-nil, zero value otherwise.

### GetTotalReadsOk

`func (o *TagUsage) GetTotalReadsOk() (*int32, bool)`

GetTotalReadsOk returns a tuple with the TotalReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReads

`func (o *TagUsage) SetTotalReads(v int32)`

SetTotalReads sets TotalReads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


