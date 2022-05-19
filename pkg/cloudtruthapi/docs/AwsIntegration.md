# AwsIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | The unique identifier for the integration. | [readonly] 
**Name** | **string** |  | [readonly] 
**Description** | Pointer to **string** | An optional description for the integration. | [optional] 
**Status** | [**NullableStatusEnum**](StatusEnum.md) | The status of the integration connection with the third-party provider as of the &#x60;status_last_checked_at&#x60; field.  The status is updated automatically by the server when the integration is modified. | [readonly] 
**StatusDetail** | **string** | If an error occurs, more details will be available in this field. | [readonly] 
**StatusLastCheckedAt** | **time.Time** | The last time the status was evaluated. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 
**Fqn** | **string** |  | [readonly] 
**Type** | **string** | The type of integration. | [readonly] 
**Writable** | Pointer to **bool** | Allow actions to write to the integration. | [optional] 
**AwsAccountId** | **string** | The AWS Account ID. | 
**AwsEnabledRegions** | [**[]AwsRegionEnum**](AwsRegionEnum.md) | The AWS regions to integrate with. | 
**AwsEnabledServices** | [**[]AwsServiceEnum**](AwsServiceEnum.md) | The AWS services to integrate with. | 
**AwsExternalId** | **string** | This is a shared secret between the AWS Administrator who set up your IAM trust relationship and your CloudTruth AWS Integration.  CloudTruth will generate a random value for you to give to your AWS Administrator in order to create the necessary IAM role for proper access. | [readonly] 
**AwsKmsKeyId** | Pointer to **NullableString** | If present, this is the KMS Key Id that is used to push values.  This key must be accessible in the AWS account (it cannot be an ARN to a key in another AWS account).  | [optional] 
**AwsRoleName** | **string** | The role that CloudTruth will assume when interacting with your AWS Account through this integration.  The role is configured by your AWS Account Administrator.  If your AWS Administrator provided you with a value use it, otherwise make your own role name and give it to your AWS Administrator. | 

## Methods

### NewAwsIntegration

`func NewAwsIntegration(url string, id string, name string, status NullableStatusEnum, statusDetail string, statusLastCheckedAt time.Time, createdAt time.Time, modifiedAt time.Time, fqn string, type_ string, awsAccountId string, awsEnabledRegions []AwsRegionEnum, awsEnabledServices []AwsServiceEnum, awsExternalId string, awsRoleName string, ) *AwsIntegration`

NewAwsIntegration instantiates a new AwsIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsIntegrationWithDefaults

`func NewAwsIntegrationWithDefaults() *AwsIntegration`

NewAwsIntegrationWithDefaults instantiates a new AwsIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AwsIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AwsIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AwsIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AwsIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AwsIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AwsIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *AwsIntegration) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsIntegration) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsIntegration) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *AwsIntegration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsIntegration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusDetail

`func (o *AwsIntegration) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *AwsIntegration) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *AwsIntegration) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.


### GetStatusLastCheckedAt

`func (o *AwsIntegration) GetStatusLastCheckedAt() time.Time`

GetStatusLastCheckedAt returns the StatusLastCheckedAt field if non-nil, zero value otherwise.

### GetStatusLastCheckedAtOk

`func (o *AwsIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool)`

GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastCheckedAt

`func (o *AwsIntegration) SetStatusLastCheckedAt(v time.Time)`

SetStatusLastCheckedAt sets StatusLastCheckedAt field to given value.


### GetCreatedAt

`func (o *AwsIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AwsIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AwsIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AwsIntegration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AwsIntegration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AwsIntegration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetFqn

`func (o *AwsIntegration) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *AwsIntegration) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *AwsIntegration) SetFqn(v string)`

SetFqn sets Fqn field to given value.


### GetType

`func (o *AwsIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetWritable

`func (o *AwsIntegration) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *AwsIntegration) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *AwsIntegration) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *AwsIntegration) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetAwsAccountId

`func (o *AwsIntegration) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *AwsIntegration) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *AwsIntegration) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetAwsEnabledRegions

`func (o *AwsIntegration) GetAwsEnabledRegions() []AwsRegionEnum`

GetAwsEnabledRegions returns the AwsEnabledRegions field if non-nil, zero value otherwise.

### GetAwsEnabledRegionsOk

`func (o *AwsIntegration) GetAwsEnabledRegionsOk() (*[]AwsRegionEnum, bool)`

GetAwsEnabledRegionsOk returns a tuple with the AwsEnabledRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEnabledRegions

`func (o *AwsIntegration) SetAwsEnabledRegions(v []AwsRegionEnum)`

SetAwsEnabledRegions sets AwsEnabledRegions field to given value.


### GetAwsEnabledServices

`func (o *AwsIntegration) GetAwsEnabledServices() []AwsServiceEnum`

GetAwsEnabledServices returns the AwsEnabledServices field if non-nil, zero value otherwise.

### GetAwsEnabledServicesOk

`func (o *AwsIntegration) GetAwsEnabledServicesOk() (*[]AwsServiceEnum, bool)`

GetAwsEnabledServicesOk returns a tuple with the AwsEnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEnabledServices

`func (o *AwsIntegration) SetAwsEnabledServices(v []AwsServiceEnum)`

SetAwsEnabledServices sets AwsEnabledServices field to given value.


### GetAwsExternalId

`func (o *AwsIntegration) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### GetAwsExternalIdOk

`func (o *AwsIntegration) GetAwsExternalIdOk() (*string, bool)`

GetAwsExternalIdOk returns a tuple with the AwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalId

`func (o *AwsIntegration) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.


### GetAwsKmsKeyId

`func (o *AwsIntegration) GetAwsKmsKeyId() string`

GetAwsKmsKeyId returns the AwsKmsKeyId field if non-nil, zero value otherwise.

### GetAwsKmsKeyIdOk

`func (o *AwsIntegration) GetAwsKmsKeyIdOk() (*string, bool)`

GetAwsKmsKeyIdOk returns a tuple with the AwsKmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsKeyId

`func (o *AwsIntegration) SetAwsKmsKeyId(v string)`

SetAwsKmsKeyId sets AwsKmsKeyId field to given value.

### HasAwsKmsKeyId

`func (o *AwsIntegration) HasAwsKmsKeyId() bool`

HasAwsKmsKeyId returns a boolean if a field has been set.

### SetAwsKmsKeyIdNil

`func (o *AwsIntegration) SetAwsKmsKeyIdNil(b bool)`

 SetAwsKmsKeyIdNil sets the value for AwsKmsKeyId to be an explicit nil

### UnsetAwsKmsKeyId
`func (o *AwsIntegration) UnsetAwsKmsKeyId()`

UnsetAwsKmsKeyId ensures that no value is present for AwsKmsKeyId, not even an explicit nil
### GetAwsRoleName

`func (o *AwsIntegration) GetAwsRoleName() string`

GetAwsRoleName returns the AwsRoleName field if non-nil, zero value otherwise.

### GetAwsRoleNameOk

`func (o *AwsIntegration) GetAwsRoleNameOk() (*string, bool)`

GetAwsRoleNameOk returns a tuple with the AwsRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleName

`func (o *AwsIntegration) SetAwsRoleName(v string)`

SetAwsRoleName sets AwsRoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


