# AuditTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the audit record. | [readonly] 
**Action** | **string** | The action that was taken. | [readonly] 
**ObjectId** | **string** | The id of the object associated with the action. | [readonly] 
**ObjectName** | **string** | The name of the object associated with the action, if applicable. | [readonly] 
**ObjectType** | [**NullableObjectTypeEnum**](ObjectTypeEnum.md) | The type of object associated with the action. | [readonly] 
**Timestamp** | **time.Time** | The timestamp of the activity that was audited. | [readonly] 
**User** | [**AuditTrailUser**](AuditTrailUser.md) |  | 

## Methods

### NewAuditTrail

`func NewAuditTrail(url string, id string, action string, objectId string, objectName string, objectType NullableObjectTypeEnum, timestamp time.Time, user AuditTrailUser, ) *AuditTrail`

NewAuditTrail instantiates a new AuditTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailWithDefaults

`func NewAuditTrailWithDefaults() *AuditTrail`

NewAuditTrailWithDefaults instantiates a new AuditTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AuditTrail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuditTrail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuditTrail) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AuditTrail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditTrail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditTrail) SetId(v string)`

SetId sets Id field to given value.


### GetAction

`func (o *AuditTrail) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditTrail) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditTrail) SetAction(v string)`

SetAction sets Action field to given value.


### GetObjectId

`func (o *AuditTrail) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AuditTrail) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AuditTrail) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectName

`func (o *AuditTrail) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *AuditTrail) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *AuditTrail) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetObjectType

`func (o *AuditTrail) GetObjectType() ObjectTypeEnum`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AuditTrail) GetObjectTypeOk() (*ObjectTypeEnum, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AuditTrail) SetObjectType(v ObjectTypeEnum)`

SetObjectType sets ObjectType field to given value.


### SetObjectTypeNil

`func (o *AuditTrail) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *AuditTrail) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetTimestamp

`func (o *AuditTrail) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditTrail) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditTrail) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUser

`func (o *AuditTrail) GetUser() AuditTrailUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditTrail) GetUserOk() (*AuditTrailUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditTrail) SetUser(v AuditTrailUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


