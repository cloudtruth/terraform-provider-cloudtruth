# GitHubPullTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | The unique identifier for the task. | [readonly] 
**Reason** | Pointer to **NullableString** | The reason why this task was triggered. | [optional] 
**DryRun** | Pointer to **bool** | Indicates task steps were only simulated, not actually performed. | [optional] 
**State** | Pointer to [**NullableStateEnum**](StateEnum.md) | The current state of this task. | [optional] 
**ErrorCode** | Pointer to **NullableString** | If an error occurs early during processing, before attempting to process values, this code may be helpful in determining the problem. | [optional] 
**ErrorDetail** | Pointer to **NullableString** | If an error occurs early during processing, before attempting to process values, this detail may be helpful in determining the problem. | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewGitHubPullTask

`func NewGitHubPullTask(url string, id string, createdAt time.Time, modifiedAt time.Time, ) *GitHubPullTask`

NewGitHubPullTask instantiates a new GitHubPullTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitHubPullTaskWithDefaults

`func NewGitHubPullTaskWithDefaults() *GitHubPullTask`

NewGitHubPullTaskWithDefaults instantiates a new GitHubPullTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GitHubPullTask) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitHubPullTask) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitHubPullTask) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *GitHubPullTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitHubPullTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitHubPullTask) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *GitHubPullTask) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GitHubPullTask) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GitHubPullTask) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GitHubPullTask) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GitHubPullTask) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GitHubPullTask) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetDryRun

`func (o *GitHubPullTask) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *GitHubPullTask) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *GitHubPullTask) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *GitHubPullTask) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetState

`func (o *GitHubPullTask) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GitHubPullTask) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GitHubPullTask) SetState(v StateEnum)`

SetState sets State field to given value.

### HasState

`func (o *GitHubPullTask) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *GitHubPullTask) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GitHubPullTask) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetErrorCode

`func (o *GitHubPullTask) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GitHubPullTask) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GitHubPullTask) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GitHubPullTask) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GitHubPullTask) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GitHubPullTask) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorDetail

`func (o *GitHubPullTask) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *GitHubPullTask) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *GitHubPullTask) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *GitHubPullTask) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### SetErrorDetailNil

`func (o *GitHubPullTask) SetErrorDetailNil(b bool)`

 SetErrorDetailNil sets the value for ErrorDetail to be an explicit nil

### UnsetErrorDetail
`func (o *GitHubPullTask) UnsetErrorDetail()`

UnsetErrorDetail ensures that no value is present for ErrorDetail, not even an explicit nil
### GetCreatedAt

`func (o *GitHubPullTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitHubPullTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitHubPullTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *GitHubPullTask) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GitHubPullTask) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GitHubPullTask) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


