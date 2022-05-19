# AuditTrailSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Earliest** | **NullableTime** | The earliest audit record timestamp available. | [readonly] 
**MaxDays** | **int32** | The maximum number of days the system will save auditing records, based on your subscription. | [readonly] 
**MaxRecords** | **int32** | The maximum number of auditing records the system will save based on your subscription. | [readonly] 
**Total** | **int32** | The total number of audit records available. | [readonly] 

## Methods

### NewAuditTrailSummary

`func NewAuditTrailSummary(earliest NullableTime, maxDays int32, maxRecords int32, total int32, ) *AuditTrailSummary`

NewAuditTrailSummary instantiates a new AuditTrailSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailSummaryWithDefaults

`func NewAuditTrailSummaryWithDefaults() *AuditTrailSummary`

NewAuditTrailSummaryWithDefaults instantiates a new AuditTrailSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarliest

`func (o *AuditTrailSummary) GetEarliest() time.Time`

GetEarliest returns the Earliest field if non-nil, zero value otherwise.

### GetEarliestOk

`func (o *AuditTrailSummary) GetEarliestOk() (*time.Time, bool)`

GetEarliestOk returns a tuple with the Earliest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliest

`func (o *AuditTrailSummary) SetEarliest(v time.Time)`

SetEarliest sets Earliest field to given value.


### SetEarliestNil

`func (o *AuditTrailSummary) SetEarliestNil(b bool)`

 SetEarliestNil sets the value for Earliest to be an explicit nil

### UnsetEarliest
`func (o *AuditTrailSummary) UnsetEarliest()`

UnsetEarliest ensures that no value is present for Earliest, not even an explicit nil
### GetMaxDays

`func (o *AuditTrailSummary) GetMaxDays() int32`

GetMaxDays returns the MaxDays field if non-nil, zero value otherwise.

### GetMaxDaysOk

`func (o *AuditTrailSummary) GetMaxDaysOk() (*int32, bool)`

GetMaxDaysOk returns a tuple with the MaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDays

`func (o *AuditTrailSummary) SetMaxDays(v int32)`

SetMaxDays sets MaxDays field to given value.


### GetMaxRecords

`func (o *AuditTrailSummary) GetMaxRecords() int32`

GetMaxRecords returns the MaxRecords field if non-nil, zero value otherwise.

### GetMaxRecordsOk

`func (o *AuditTrailSummary) GetMaxRecordsOk() (*int32, bool)`

GetMaxRecordsOk returns a tuple with the MaxRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecords

`func (o *AuditTrailSummary) SetMaxRecords(v int32)`

SetMaxRecords sets MaxRecords field to given value.


### GetTotal

`func (o *AuditTrailSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuditTrailSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuditTrailSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


