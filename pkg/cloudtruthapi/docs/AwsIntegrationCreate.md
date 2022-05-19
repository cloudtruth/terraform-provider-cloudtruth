# AwsIntegrationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional description for the integration. | [optional] 
**Writable** | Pointer to **bool** | Allow actions to write to the integration. | [optional] 
**AwsAccountId** | **string** | The AWS Account ID. | 
**AwsEnabledRegions** | [**[]AwsRegionEnum**](AwsRegionEnum.md) | The AWS regions to integrate with. | 
**AwsEnabledServices** | [**[]AwsServiceEnum**](AwsServiceEnum.md) | The AWS services to integrate with. | 
**AwsExternalId** | Pointer to **string** | This is a shared secret between the AWS Administrator who set up your IAM trust relationship and your CloudTruth AWS Integration.  CloudTruth will generate a random value for you to give to your AWS Administrator in order to create the necessary IAM role for proper access. | [optional] 
**AwsKmsKeyId** | Pointer to **NullableString** | If present, this is the KMS Key Id that is used to push values.  This key must be accessible in the AWS account (it cannot be an ARN to a key in another AWS account).  | [optional] 
**AwsRoleName** | **string** | The role that CloudTruth will assume when interacting with your AWS Account through this integration.  The role is configured by your AWS Account Administrator.  If your AWS Administrator provided you with a value use it, otherwise make your own role name and give it to your AWS Administrator. | 

## Methods

### NewAwsIntegrationCreate

`func NewAwsIntegrationCreate(awsAccountId string, awsEnabledRegions []AwsRegionEnum, awsEnabledServices []AwsServiceEnum, awsRoleName string, ) *AwsIntegrationCreate`

NewAwsIntegrationCreate instantiates a new AwsIntegrationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsIntegrationCreateWithDefaults

`func NewAwsIntegrationCreateWithDefaults() *AwsIntegrationCreate`

NewAwsIntegrationCreateWithDefaults instantiates a new AwsIntegrationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AwsIntegrationCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsIntegrationCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsIntegrationCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsIntegrationCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWritable

`func (o *AwsIntegrationCreate) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *AwsIntegrationCreate) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *AwsIntegrationCreate) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *AwsIntegrationCreate) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetAwsAccountId

`func (o *AwsIntegrationCreate) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *AwsIntegrationCreate) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *AwsIntegrationCreate) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetAwsEnabledRegions

`func (o *AwsIntegrationCreate) GetAwsEnabledRegions() []AwsRegionEnum`

GetAwsEnabledRegions returns the AwsEnabledRegions field if non-nil, zero value otherwise.

### GetAwsEnabledRegionsOk

`func (o *AwsIntegrationCreate) GetAwsEnabledRegionsOk() (*[]AwsRegionEnum, bool)`

GetAwsEnabledRegionsOk returns a tuple with the AwsEnabledRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEnabledRegions

`func (o *AwsIntegrationCreate) SetAwsEnabledRegions(v []AwsRegionEnum)`

SetAwsEnabledRegions sets AwsEnabledRegions field to given value.


### GetAwsEnabledServices

`func (o *AwsIntegrationCreate) GetAwsEnabledServices() []AwsServiceEnum`

GetAwsEnabledServices returns the AwsEnabledServices field if non-nil, zero value otherwise.

### GetAwsEnabledServicesOk

`func (o *AwsIntegrationCreate) GetAwsEnabledServicesOk() (*[]AwsServiceEnum, bool)`

GetAwsEnabledServicesOk returns a tuple with the AwsEnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEnabledServices

`func (o *AwsIntegrationCreate) SetAwsEnabledServices(v []AwsServiceEnum)`

SetAwsEnabledServices sets AwsEnabledServices field to given value.


### GetAwsExternalId

`func (o *AwsIntegrationCreate) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### GetAwsExternalIdOk

`func (o *AwsIntegrationCreate) GetAwsExternalIdOk() (*string, bool)`

GetAwsExternalIdOk returns a tuple with the AwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalId

`func (o *AwsIntegrationCreate) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### HasAwsExternalId

`func (o *AwsIntegrationCreate) HasAwsExternalId() bool`

HasAwsExternalId returns a boolean if a field has been set.

### GetAwsKmsKeyId

`func (o *AwsIntegrationCreate) GetAwsKmsKeyId() string`

GetAwsKmsKeyId returns the AwsKmsKeyId field if non-nil, zero value otherwise.

### GetAwsKmsKeyIdOk

`func (o *AwsIntegrationCreate) GetAwsKmsKeyIdOk() (*string, bool)`

GetAwsKmsKeyIdOk returns a tuple with the AwsKmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsKeyId

`func (o *AwsIntegrationCreate) SetAwsKmsKeyId(v string)`

SetAwsKmsKeyId sets AwsKmsKeyId field to given value.

### HasAwsKmsKeyId

`func (o *AwsIntegrationCreate) HasAwsKmsKeyId() bool`

HasAwsKmsKeyId returns a boolean if a field has been set.

### SetAwsKmsKeyIdNil

`func (o *AwsIntegrationCreate) SetAwsKmsKeyIdNil(b bool)`

 SetAwsKmsKeyIdNil sets the value for AwsKmsKeyId to be an explicit nil

### UnsetAwsKmsKeyId
`func (o *AwsIntegrationCreate) UnsetAwsKmsKeyId()`

UnsetAwsKmsKeyId ensures that no value is present for AwsKmsKeyId, not even an explicit nil
### GetAwsRoleName

`func (o *AwsIntegrationCreate) GetAwsRoleName() string`

GetAwsRoleName returns the AwsRoleName field if non-nil, zero value otherwise.

### GetAwsRoleNameOk

`func (o *AwsIntegrationCreate) GetAwsRoleNameOk() (*string, bool)`

GetAwsRoleNameOk returns a tuple with the AwsRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleName

`func (o *AwsIntegrationCreate) SetAwsRoleName(v string)`

SetAwsRoleName sets AwsRoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


