# AzureKeyVaultIntegrationScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | **NullableString** | Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to identify the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to identify the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the &#x60;environment&#x60;, &#x60;project&#x60;, and &#x60;parameter&#x60;. | 

## Methods

### NewAzureKeyVaultIntegrationScan

`func NewAzureKeyVaultIntegrationScan(resource NullableString, ) *AzureKeyVaultIntegrationScan`

NewAzureKeyVaultIntegrationScan instantiates a new AzureKeyVaultIntegrationScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultIntegrationScanWithDefaults

`func NewAzureKeyVaultIntegrationScanWithDefaults() *AzureKeyVaultIntegrationScan`

NewAzureKeyVaultIntegrationScanWithDefaults instantiates a new AzureKeyVaultIntegrationScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *AzureKeyVaultIntegrationScan) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AzureKeyVaultIntegrationScan) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AzureKeyVaultIntegrationScan) SetResource(v string)`

SetResource sets Resource field to given value.


### SetResourceNil

`func (o *AzureKeyVaultIntegrationScan) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AzureKeyVaultIntegrationScan) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


