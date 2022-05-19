# AwsIntegrationScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**NullableAwsRegionEnum**](AwsRegionEnum.md) | The AWS region to use.  This region must be enabled in the integration. | 
**Service** | [**NullableAwsServiceEnum**](AwsServiceEnum.md) | The AWS service to use.  This service must be enabled in the integration. | 
**Resource** | **NullableString** | Defines a pattern matching string that contains either mustache or regular expression syntax (with named capture groups) that locate the environment, project, and parameter name of the content you are looking for.  If you are using mustache pattern matching, use:    - &#x60;&#x60;{{ environment }}&#x60;&#x60; to identify the environment name   - &#x60;&#x60;{{ parameter }}&#x60;&#x60; to identify the parameter name   - &#x60;&#x60;{{ project }}&#x60;&#x60; to identify the project name  If you are using a regular expression, use Python syntax with named capture groups that locate the &#x60;environment&#x60;, &#x60;project&#x60;, and &#x60;parameter&#x60;. | 

## Methods

### NewAwsIntegrationScan

`func NewAwsIntegrationScan(region NullableAwsRegionEnum, service NullableAwsServiceEnum, resource NullableString, ) *AwsIntegrationScan`

NewAwsIntegrationScan instantiates a new AwsIntegrationScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsIntegrationScanWithDefaults

`func NewAwsIntegrationScanWithDefaults() *AwsIntegrationScan`

NewAwsIntegrationScanWithDefaults instantiates a new AwsIntegrationScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AwsIntegrationScan) GetRegion() AwsRegionEnum`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsIntegrationScan) GetRegionOk() (*AwsRegionEnum, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsIntegrationScan) SetRegion(v AwsRegionEnum)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsIntegrationScan) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsIntegrationScan) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetService

`func (o *AwsIntegrationScan) GetService() AwsServiceEnum`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AwsIntegrationScan) GetServiceOk() (*AwsServiceEnum, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AwsIntegrationScan) SetService(v AwsServiceEnum)`

SetService sets Service field to given value.


### SetServiceNil

`func (o *AwsIntegrationScan) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *AwsIntegrationScan) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetResource

`func (o *AwsIntegrationScan) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AwsIntegrationScan) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AwsIntegrationScan) SetResource(v string)`

SetResource sets Resource field to given value.


### SetResourceNil

`func (o *AwsIntegrationScan) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AwsIntegrationScan) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


