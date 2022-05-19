# AzureKeyVaultIntegration

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
**VaultName** | **string** | The Azure Key Vault name. | [readonly] 
**TenantId** | **string** | The Azure Tenant ID. | [readonly] 

## Methods

### NewAzureKeyVaultIntegration

`func NewAzureKeyVaultIntegration(url string, id string, name string, status NullableStatusEnum, statusDetail string, statusLastCheckedAt time.Time, createdAt time.Time, modifiedAt time.Time, fqn string, type_ string, vaultName string, tenantId string, ) *AzureKeyVaultIntegration`

NewAzureKeyVaultIntegration instantiates a new AzureKeyVaultIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultIntegrationWithDefaults

`func NewAzureKeyVaultIntegrationWithDefaults() *AzureKeyVaultIntegration`

NewAzureKeyVaultIntegrationWithDefaults instantiates a new AzureKeyVaultIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AzureKeyVaultIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AzureKeyVaultIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AzureKeyVaultIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *AzureKeyVaultIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureKeyVaultIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureKeyVaultIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AzureKeyVaultIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureKeyVaultIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureKeyVaultIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AzureKeyVaultIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureKeyVaultIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureKeyVaultIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureKeyVaultIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *AzureKeyVaultIntegration) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AzureKeyVaultIntegration) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AzureKeyVaultIntegration) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *AzureKeyVaultIntegration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AzureKeyVaultIntegration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusDetail

`func (o *AzureKeyVaultIntegration) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *AzureKeyVaultIntegration) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *AzureKeyVaultIntegration) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.


### GetStatusLastCheckedAt

`func (o *AzureKeyVaultIntegration) GetStatusLastCheckedAt() time.Time`

GetStatusLastCheckedAt returns the StatusLastCheckedAt field if non-nil, zero value otherwise.

### GetStatusLastCheckedAtOk

`func (o *AzureKeyVaultIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool)`

GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastCheckedAt

`func (o *AzureKeyVaultIntegration) SetStatusLastCheckedAt(v time.Time)`

SetStatusLastCheckedAt sets StatusLastCheckedAt field to given value.


### GetCreatedAt

`func (o *AzureKeyVaultIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AzureKeyVaultIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AzureKeyVaultIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *AzureKeyVaultIntegration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AzureKeyVaultIntegration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AzureKeyVaultIntegration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetFqn

`func (o *AzureKeyVaultIntegration) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *AzureKeyVaultIntegration) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *AzureKeyVaultIntegration) SetFqn(v string)`

SetFqn sets Fqn field to given value.


### GetType

`func (o *AzureKeyVaultIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureKeyVaultIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureKeyVaultIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetWritable

`func (o *AzureKeyVaultIntegration) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *AzureKeyVaultIntegration) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *AzureKeyVaultIntegration) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *AzureKeyVaultIntegration) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetVaultName

`func (o *AzureKeyVaultIntegration) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *AzureKeyVaultIntegration) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *AzureKeyVaultIntegration) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.


### GetTenantId

`func (o *AzureKeyVaultIntegration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureKeyVaultIntegration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureKeyVaultIntegration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


