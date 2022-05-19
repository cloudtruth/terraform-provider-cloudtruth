# BackupDataSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | [**map[string]BackupEnvironment**](BackupEnvironment.md) |  | 
**Types** | [**map[string]BackupParameterType**](BackupParameterType.md) |  | 
**Projects** | [**map[string]BackupProject**](BackupProject.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewBackupDataSnapshot

`func NewBackupDataSnapshot(environments map[string]BackupEnvironment, types map[string]BackupParameterType, projects map[string]BackupProject, timestamp time.Time, ) *BackupDataSnapshot`

NewBackupDataSnapshot instantiates a new BackupDataSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDataSnapshotWithDefaults

`func NewBackupDataSnapshotWithDefaults() *BackupDataSnapshot`

NewBackupDataSnapshotWithDefaults instantiates a new BackupDataSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *BackupDataSnapshot) GetEnvironments() map[string]BackupEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *BackupDataSnapshot) GetEnvironmentsOk() (*map[string]BackupEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *BackupDataSnapshot) SetEnvironments(v map[string]BackupEnvironment)`

SetEnvironments sets Environments field to given value.


### GetTypes

`func (o *BackupDataSnapshot) GetTypes() map[string]BackupParameterType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *BackupDataSnapshot) GetTypesOk() (*map[string]BackupParameterType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *BackupDataSnapshot) SetTypes(v map[string]BackupParameterType)`

SetTypes sets Types field to given value.


### GetProjects

`func (o *BackupDataSnapshot) GetProjects() map[string]BackupProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BackupDataSnapshot) GetProjectsOk() (*map[string]BackupProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BackupDataSnapshot) SetProjects(v map[string]BackupProject)`

SetProjects sets Projects field to given value.


### GetTimestamp

`func (o *BackupDataSnapshot) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BackupDataSnapshot) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BackupDataSnapshot) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


