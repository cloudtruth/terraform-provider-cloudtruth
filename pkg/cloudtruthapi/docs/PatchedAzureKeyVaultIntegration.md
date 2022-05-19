# PatchedAzureKeyVaultIntegration

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
**VaultName** | Pointer to **string** | The Azure Key Vault name. | [optional] [readonly] 
**TenantId** | Pointer to **string** | The Azure Tenant ID. | [optional] [readonly] 

## Methods

### NewPatchedAzureKeyVaultIntegration

`func NewPatchedAzureKeyVaultIntegration() *PatchedAzureKeyVaultIntegration`

NewPatchedAzureKeyVaultIntegration instantiates a new PatchedAzureKeyVaultIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAzureKeyVaultIntegrationWithDefaults

`func NewPatchedAzureKeyVaultIntegrationWithDefaults() *PatchedAzureKeyVaultIntegration`

NewPatchedAzureKeyVaultIntegrationWithDefaults instantiates a new PatchedAzureKeyVaultIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedAzureKeyVaultIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedAzureKeyVaultIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedAzureKeyVaultIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedAzureKeyVaultIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedAzureKeyVaultIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAzureKeyVaultIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAzureKeyVaultIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAzureKeyVaultIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedAzureKeyVaultIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAzureKeyVaultIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAzureKeyVaultIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAzureKeyVaultIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedAzureKeyVaultIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedAzureKeyVaultIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedAzureKeyVaultIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedAzureKeyVaultIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedAzureKeyVaultIntegration) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedAzureKeyVaultIntegration) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedAzureKeyVaultIntegration) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedAzureKeyVaultIntegration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedAzureKeyVaultIntegration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedAzureKeyVaultIntegration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusDetail

`func (o *PatchedAzureKeyVaultIntegration) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *PatchedAzureKeyVaultIntegration) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *PatchedAzureKeyVaultIntegration) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *PatchedAzureKeyVaultIntegration) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetStatusLastCheckedAt

`func (o *PatchedAzureKeyVaultIntegration) GetStatusLastCheckedAt() time.Time`

GetStatusLastCheckedAt returns the StatusLastCheckedAt field if non-nil, zero value otherwise.

### GetStatusLastCheckedAtOk

`func (o *PatchedAzureKeyVaultIntegration) GetStatusLastCheckedAtOk() (*time.Time, bool)`

GetStatusLastCheckedAtOk returns a tuple with the StatusLastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLastCheckedAt

`func (o *PatchedAzureKeyVaultIntegration) SetStatusLastCheckedAt(v time.Time)`

SetStatusLastCheckedAt sets StatusLastCheckedAt field to given value.

### HasStatusLastCheckedAt

`func (o *PatchedAzureKeyVaultIntegration) HasStatusLastCheckedAt() bool`

HasStatusLastCheckedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedAzureKeyVaultIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedAzureKeyVaultIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedAzureKeyVaultIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedAzureKeyVaultIntegration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedAzureKeyVaultIntegration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedAzureKeyVaultIntegration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedAzureKeyVaultIntegration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedAzureKeyVaultIntegration) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetFqn

`func (o *PatchedAzureKeyVaultIntegration) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *PatchedAzureKeyVaultIntegration) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *PatchedAzureKeyVaultIntegration) SetFqn(v string)`

SetFqn sets Fqn field to given value.

### HasFqn

`func (o *PatchedAzureKeyVaultIntegration) HasFqn() bool`

HasFqn returns a boolean if a field has been set.

### GetType

`func (o *PatchedAzureKeyVaultIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedAzureKeyVaultIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedAzureKeyVaultIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedAzureKeyVaultIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWritable

`func (o *PatchedAzureKeyVaultIntegration) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *PatchedAzureKeyVaultIntegration) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *PatchedAzureKeyVaultIntegration) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *PatchedAzureKeyVaultIntegration) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetVaultName

`func (o *PatchedAzureKeyVaultIntegration) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *PatchedAzureKeyVaultIntegration) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *PatchedAzureKeyVaultIntegration) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *PatchedAzureKeyVaultIntegration) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedAzureKeyVaultIntegration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedAzureKeyVaultIntegration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedAzureKeyVaultIntegration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedAzureKeyVaultIntegration) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


