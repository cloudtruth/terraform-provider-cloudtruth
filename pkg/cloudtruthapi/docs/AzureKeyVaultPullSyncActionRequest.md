# AzureKeyVaultPullSyncActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | Allows you to set the dry_run flag on the pull action before triggering a sync. | [optional] 

## Methods

### NewAzureKeyVaultPullSyncActionRequest

`func NewAzureKeyVaultPullSyncActionRequest() *AzureKeyVaultPullSyncActionRequest`

NewAzureKeyVaultPullSyncActionRequest instantiates a new AzureKeyVaultPullSyncActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultPullSyncActionRequestWithDefaults

`func NewAzureKeyVaultPullSyncActionRequestWithDefaults() *AzureKeyVaultPullSyncActionRequest`

NewAzureKeyVaultPullSyncActionRequestWithDefaults instantiates a new AzureKeyVaultPullSyncActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *AzureKeyVaultPullSyncActionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AzureKeyVaultPullSyncActionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AzureKeyVaultPullSyncActionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AzureKeyVaultPullSyncActionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


