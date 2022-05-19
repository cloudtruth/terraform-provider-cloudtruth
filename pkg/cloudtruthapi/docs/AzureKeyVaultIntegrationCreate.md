# AzureKeyVaultIntegrationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional description for the integration. | [optional] 
**Writable** | Pointer to **bool** | Allow actions to write to the integration. | [optional] 
**VaultName** | **string** | The Azure Key Vault name. | 
**TenantId** | **string** | The Azure Tenant ID. | 

## Methods

### NewAzureKeyVaultIntegrationCreate

`func NewAzureKeyVaultIntegrationCreate(vaultName string, tenantId string, ) *AzureKeyVaultIntegrationCreate`

NewAzureKeyVaultIntegrationCreate instantiates a new AzureKeyVaultIntegrationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultIntegrationCreateWithDefaults

`func NewAzureKeyVaultIntegrationCreateWithDefaults() *AzureKeyVaultIntegrationCreate`

NewAzureKeyVaultIntegrationCreateWithDefaults instantiates a new AzureKeyVaultIntegrationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AzureKeyVaultIntegrationCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureKeyVaultIntegrationCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureKeyVaultIntegrationCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureKeyVaultIntegrationCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWritable

`func (o *AzureKeyVaultIntegrationCreate) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *AzureKeyVaultIntegrationCreate) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *AzureKeyVaultIntegrationCreate) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *AzureKeyVaultIntegrationCreate) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetVaultName

`func (o *AzureKeyVaultIntegrationCreate) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *AzureKeyVaultIntegrationCreate) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *AzureKeyVaultIntegrationCreate) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.


### GetTenantId

`func (o *AzureKeyVaultIntegrationCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureKeyVaultIntegrationCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureKeyVaultIntegrationCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


