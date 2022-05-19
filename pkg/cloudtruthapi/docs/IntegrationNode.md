# IntegrationNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqn** | **string** |  | 
**NodeType** | [**NodeTypeEnum**](NodeTypeEnum.md) |  | 
**Secret** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Jmespath** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ContentData** | Pointer to **NullableString** |  | [optional] 
**ContentSize** | Pointer to **NullableInt32** |  | [optional] 
**ContentKeys** | Pointer to **[]string** |  | [optional] 
**VenueId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIntegrationNode

`func NewIntegrationNode(fqn string, nodeType NodeTypeEnum, ) *IntegrationNode`

NewIntegrationNode instantiates a new IntegrationNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationNodeWithDefaults

`func NewIntegrationNodeWithDefaults() *IntegrationNode`

NewIntegrationNodeWithDefaults instantiates a new IntegrationNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqn

`func (o *IntegrationNode) GetFqn() string`

GetFqn returns the Fqn field if non-nil, zero value otherwise.

### GetFqnOk

`func (o *IntegrationNode) GetFqnOk() (*string, bool)`

GetFqnOk returns a tuple with the Fqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqn

`func (o *IntegrationNode) SetFqn(v string)`

SetFqn sets Fqn field to given value.


### GetNodeType

`func (o *IntegrationNode) GetNodeType() NodeTypeEnum`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *IntegrationNode) GetNodeTypeOk() (*NodeTypeEnum, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *IntegrationNode) SetNodeType(v NodeTypeEnum)`

SetNodeType sets NodeType field to given value.


### GetSecret

`func (o *IntegrationNode) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IntegrationNode) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IntegrationNode) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IntegrationNode) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetName

`func (o *IntegrationNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetJmespath

`func (o *IntegrationNode) GetJmespath() string`

GetJmespath returns the Jmespath field if non-nil, zero value otherwise.

### GetJmespathOk

`func (o *IntegrationNode) GetJmespathOk() (*string, bool)`

GetJmespathOk returns a tuple with the Jmespath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmespath

`func (o *IntegrationNode) SetJmespath(v string)`

SetJmespath sets Jmespath field to given value.

### HasJmespath

`func (o *IntegrationNode) HasJmespath() bool`

HasJmespath returns a boolean if a field has been set.

### SetJmespathNil

`func (o *IntegrationNode) SetJmespathNil(b bool)`

 SetJmespathNil sets the value for Jmespath to be an explicit nil

### UnsetJmespath
`func (o *IntegrationNode) UnsetJmespath()`

UnsetJmespath ensures that no value is present for Jmespath, not even an explicit nil
### GetContentType

`func (o *IntegrationNode) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *IntegrationNode) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *IntegrationNode) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *IntegrationNode) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *IntegrationNode) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *IntegrationNode) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentData

`func (o *IntegrationNode) GetContentData() string`

GetContentData returns the ContentData field if non-nil, zero value otherwise.

### GetContentDataOk

`func (o *IntegrationNode) GetContentDataOk() (*string, bool)`

GetContentDataOk returns a tuple with the ContentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentData

`func (o *IntegrationNode) SetContentData(v string)`

SetContentData sets ContentData field to given value.

### HasContentData

`func (o *IntegrationNode) HasContentData() bool`

HasContentData returns a boolean if a field has been set.

### SetContentDataNil

`func (o *IntegrationNode) SetContentDataNil(b bool)`

 SetContentDataNil sets the value for ContentData to be an explicit nil

### UnsetContentData
`func (o *IntegrationNode) UnsetContentData()`

UnsetContentData ensures that no value is present for ContentData, not even an explicit nil
### GetContentSize

`func (o *IntegrationNode) GetContentSize() int32`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *IntegrationNode) GetContentSizeOk() (*int32, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *IntegrationNode) SetContentSize(v int32)`

SetContentSize sets ContentSize field to given value.

### HasContentSize

`func (o *IntegrationNode) HasContentSize() bool`

HasContentSize returns a boolean if a field has been set.

### SetContentSizeNil

`func (o *IntegrationNode) SetContentSizeNil(b bool)`

 SetContentSizeNil sets the value for ContentSize to be an explicit nil

### UnsetContentSize
`func (o *IntegrationNode) UnsetContentSize()`

UnsetContentSize ensures that no value is present for ContentSize, not even an explicit nil
### GetContentKeys

`func (o *IntegrationNode) GetContentKeys() []string`

GetContentKeys returns the ContentKeys field if non-nil, zero value otherwise.

### GetContentKeysOk

`func (o *IntegrationNode) GetContentKeysOk() (*[]string, bool)`

GetContentKeysOk returns a tuple with the ContentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentKeys

`func (o *IntegrationNode) SetContentKeys(v []string)`

SetContentKeys sets ContentKeys field to given value.

### HasContentKeys

`func (o *IntegrationNode) HasContentKeys() bool`

HasContentKeys returns a boolean if a field has been set.

### SetContentKeysNil

`func (o *IntegrationNode) SetContentKeysNil(b bool)`

 SetContentKeysNil sets the value for ContentKeys to be an explicit nil

### UnsetContentKeys
`func (o *IntegrationNode) UnsetContentKeys()`

UnsetContentKeys ensures that no value is present for ContentKeys, not even an explicit nil
### GetVenueId

`func (o *IntegrationNode) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *IntegrationNode) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *IntegrationNode) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.

### HasVenueId

`func (o *IntegrationNode) HasVenueId() bool`

HasVenueId returns a boolean if a field has been set.

### SetVenueIdNil

`func (o *IntegrationNode) SetVenueIdNil(b bool)`

 SetVenueIdNil sets the value for VenueId to be an explicit nil

### UnsetVenueId
`func (o *IntegrationNode) UnsetVenueId()`

UnsetVenueId ensures that no value is present for VenueId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


