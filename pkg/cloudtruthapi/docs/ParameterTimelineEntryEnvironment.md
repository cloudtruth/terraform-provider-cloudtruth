# ParameterTimelineEntryEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the environment. | [readonly] 
**Name** | **string** | The environment name. | 
**Override** | **bool** | Indicates if the value change was direct or if it flowed into the environment. If &#x60;true&#x60; then the value was actually set directly into this environment. If &#x60;false&#x60; then the environment has no value set directly so it inherited the value from its parent. | [readonly] 

## Methods

### NewParameterTimelineEntryEnvironment

`func NewParameterTimelineEntryEnvironment(id string, name string, override bool, ) *ParameterTimelineEntryEnvironment`

NewParameterTimelineEntryEnvironment instantiates a new ParameterTimelineEntryEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterTimelineEntryEnvironmentWithDefaults

`func NewParameterTimelineEntryEnvironmentWithDefaults() *ParameterTimelineEntryEnvironment`

NewParameterTimelineEntryEnvironmentWithDefaults instantiates a new ParameterTimelineEntryEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterTimelineEntryEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterTimelineEntryEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterTimelineEntryEnvironment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ParameterTimelineEntryEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterTimelineEntryEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterTimelineEntryEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetOverride

`func (o *ParameterTimelineEntryEnvironment) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *ParameterTimelineEntryEnvironment) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *ParameterTimelineEntryEnvironment) SetOverride(v bool)`

SetOverride sets Override field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


