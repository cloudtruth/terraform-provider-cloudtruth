# PatchedAwsIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the integration. | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** | An optional description for the integration. | [optional] 
**Status** | Pointer to [**NullableStatusEnum**](StatusEnum.md) | The status of the integration connection with the third-party provider as of the &#x60;status_last_checked_at&#x60; field.  The status is updated automatically by the server when the integration is modified. | [optional] [readonly] 
**StatusDetail** | Pointer to **string** | If an error occurs, more details will be available in this field. | [optional] [readonly] 
**StatusLastCheckedAt** | Pointer to **time.Time** | The last time the status was evaluated. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Fqn** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** | The type of integration. | [optional] [readonly] 
**Writable** | Pointer to **bool** | Allow actions to write to the integration. | [optional] 
**AwsAccountId** | Pointer to **string** | The AWS Account ID. | [optional] 
**AwsEnabledRegions** | Pointer to [**[]AwsRegionEnum**](AwsRegionEnum.md) | The AWS regions to integrate with. | [optional] 
**AwsEnabledServices** | Pointer to [**[]AwsServiceEnum**](AwsServiceEnum.md) | The AWS services to integrate with. | [optional] 
**AwsExternalId** | Pointer to **string** | This is a shared secret between the AWS Administrator who set up your IAM trust relationship and your CloudTruth AWS Integration.  CloudTruth will generate a random value for you to give to your AWS Administrator in order to create the necessary IAM role for proper access. | [optional] [readonly] 
**AwsKmsKeyId** | Pointer to **NullableString** | If present, this is the KMS Key Id that is used to push values.  This key must be accessible in the AWS account (it cannot be an ARN to a key in another AWS account).  | [optional] 
**AwsRoleName** | Pointer to **string** | The role that CloudTruth will assume when interacting with your AWS Account through this integration.  The role is configured by your AWS Account Administrator.  If your AWS Administrator provided you with a value use it, otherwise make your own role name and give it to your AWS Administrator. | [optional] 

## Methods

### NewPatchedAwsIntegration

`func NewPatchedAwsIntegration() *PatchedAwsIntegration`

NewPatchedAwsIntegration instantiates a new PatchedAwsIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAwsIntegrationWithDefaults

`func NewPatchedAwsIntegrationWithDefaults() *PatchedAwsIntegration`

NewPatchedAwsIntegrationWithDefaults instantiates a new PatchedAwsIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedAwsIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedAwsIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedAwsIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedAwsIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedAwsIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAwsIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAwsIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAwsIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedAwsIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAwsIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAwsIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAwsIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedAwsIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedAwsIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedAwsIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedAwsIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedAwsIntegration) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedAwsIntegration) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedAwsIntegration) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedAwsIntegration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedAwsIntegration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedAwsIntegration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusDetail

`func (o *PatchedAwsIntegration) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *PatchedAwsIntegration) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *PatchedAwsIntegration) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *PatchedAwsIntegration) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetStatusLastCheckedAt

`func (o *PatchedAwsIntegration) GetStatusLastCheckedAt() time.Time`

GetStatusLastCheckedAt returns the StatusLastCheckedAt field if non-nil, zero value otherwise.

### GetStatusLastCheckedAtOk

`func (o *PatchedAwsIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool)`

GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastCheckedAt

`func (o *PatchedAwsIntegration) SetStatusLastCheckedAt(v time.Time)`

SetStatusLastCheckedAt sets StatusLastCheckedAt field to given value.

### HasStatusLastCheckedAt

`func (o *PatchedAwsIntegration) HasStatusLastCheckedAt() bool`

HasStatusLastCheckedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedAwsIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedAwsIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedAwsIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedAwsIntegration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedAwsIntegration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedAwsIntegration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedAwsIntegration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedAwsIntegration) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetFqn

`func (o *PatchedAwsIntegration) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *PatchedAwsIntegration) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *PatchedAwsIntegration) SetFqn(v string)`

SetFqn sets Fqn field to given value.

### HasFqn

`func (o *PatchedAwsIntegration) HasFqn() bool`

HasFqn returns a boolean if a field has been set.

### GetType

`func (o *PatchedAwsIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedAwsIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedAwsIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedAwsIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWritable

`func (o *PatchedAwsIntegration) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *PatchedAwsIntegration) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *PatchedAwsIntegration) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *PatchedAwsIntegration) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetAwsAccountId

`func (o *PatchedAwsIntegration) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *PatchedAwsIntegration) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *PatchedAwsIntegration) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.

### HasAwsAccountId

`func (o *PatchedAwsIntegration) HasAwsAccountId() bool`

HasAwsAccountId returns a boolean if a field has been set.

### GetAwsEnabledRegions

`func (o *PatchedAwsIntegration) GetAwsEnabledRegions() []AwsRegionEnum`

GetAwsEnabledRegions returns the AwsEnabledRegions field if non-nil, zero value otherwise.

### GetAwsEnabledRegionsOk

`func (o *PatchedAwsIntegration) GetAwsEnabledRegionsOk() (*[]AwsRegionEnum, bool)`

GetAwsEnabledRegionsOk returns a tuple with the AwsEnabledRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEnabledRegions

`func (o *PatchedAwsIntegration) SetAwsEnabledRegions(v []AwsRegionEnum)`

SetAwsEnabledRegions sets AwsEnabledRegions field to given value.

### HasAwsEnabledRegions

`func (o *PatchedAwsIntegration) HasAwsEnabledRegions() bool`

HasAwsEnabledRegions returns a boolean if a field has been set.

### GetAwsEnabledServices

`func (o *PatchedAwsIntegration) GetAwsEnabledServices() []AwsServiceEnum`

GetAwsEnabledServices returns the AwsEnabledServices field if non-nil, zero value otherwise.

### GetAwsEnabledServicesOk

`func (o *PatchedAwsIntegration) GetAwsEnabledServicesOk() (*[]AwsServiceEnum, bool)`

GetAwsEnabledServicesOk returns a tuple with the AwsEnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEnabledServices

`func (o *PatchedAwsIntegration) SetAwsEnabledServices(v []AwsServiceEnum)`

SetAwsEnabledServices sets AwsEnabledServices field to given value.

### HasAwsEnabledServices

`func (o *PatchedAwsIntegration) HasAwsEnabledServices() bool`

HasAwsEnabledServices returns a boolean if a field has been set.

### GetAwsExternalId

`func (o *PatchedAwsIntegration) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### GetAwsExternalIdOk

`func (o *PatchedAwsIntegration) GetAwsExternalIdOk() (*string, bool)`

GetAwsExternalIdOk returns a tuple with the AwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalId

`func (o *PatchedAwsIntegration) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### HasAwsExternalId

`func (o *PatchedAwsIntegration) HasAwsExternalId() bool`

HasAwsExternalId returns a boolean if a field has been set.

### GetAwsKmsKeyId

`func (o *PatchedAwsIntegration) GetAwsKmsKeyId() string`

GetAwsKmsKeyId returns the AwsKmsKeyId field if non-nil, zero value otherwise.

### GetAwsKmsKeyIdOk

`func (o *PatchedAwsIntegration) GetAwsKmsKeyIdOk() (*string, bool)`

GetAwsKmsKeyIdOk returns a tuple with the AwsKmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsKeyId

`func (o *PatchedAwsIntegration) SetAwsKmsKeyId(v string)`

SetAwsKmsKeyId sets AwsKmsKeyId field to given value.

### HasAwsKmsKeyId

`func (o *PatchedAwsIntegration) HasAwsKmsKeyId() bool`

HasAwsKmsKeyId returns a boolean if a field has been set.

### SetAwsKmsKeyIdNil

`func (o *PatchedAwsIntegration) SetAwsKmsKeyIdNil(b bool)`

 SetAwsKmsKeyIdNil sets the value for AwsKmsKeyId to be an explicit nil

### UnsetAwsKmsKeyId
`func (o *PatchedAwsIntegration) UnsetAwsKmsKeyId()`

UnsetAwsKmsKeyId ensures that no value is present for AwsKmsKeyId, not even an explicit nil
### GetAwsRoleName

`func (o *PatchedAwsIntegration) GetAwsRoleName() string`

GetAwsRoleName returns the AwsRoleName field if non-nil, zero value otherwise.

### GetAwsRoleNameOk

`func (o *PatchedAwsIntegration) GetAwsRoleNameOk() (*string, bool)`

GetAwsRoleNameOk returns a tuple with the AwsRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleName

`func (o *PatchedAwsIntegration) SetAwsRoleName(v string)`

SetAwsRoleName sets AwsRoleName field to given value.

### HasAwsRoleName

`func (o *PatchedAwsIntegration) HasAwsRoleName() bool`

HasAwsRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


