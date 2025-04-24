# ProjectBackupMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 
**MaxParallelTasks** | Pointer to **int32** |  | [optional] 

## Methods

### NewProjectBackupMeta

`func NewProjectBackupMeta() *ProjectBackupMeta`

NewProjectBackupMeta instantiates a new ProjectBackupMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectBackupMetaWithDefaults

`func NewProjectBackupMetaWithDefaults() *ProjectBackupMeta`

NewProjectBackupMetaWithDefaults instantiates a new ProjectBackupMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectBackupMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectBackupMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectBackupMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectBackupMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlert

`func (o *ProjectBackupMeta) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *ProjectBackupMeta) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *ProjectBackupMeta) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *ProjectBackupMeta) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetMaxParallelTasks

`func (o *ProjectBackupMeta) GetMaxParallelTasks() int32`

GetMaxParallelTasks returns the MaxParallelTasks field if non-nil, zero value otherwise.

### GetMaxParallelTasksOk

`func (o *ProjectBackupMeta) GetMaxParallelTasksOk() (*int32, bool)`

GetMaxParallelTasksOk returns a tuple with the MaxParallelTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelTasks

`func (o *ProjectBackupMeta) SetMaxParallelTasks(v int32)`

SetMaxParallelTasks sets MaxParallelTasks field to given value.

### HasMaxParallelTasks

`func (o *ProjectBackupMeta) HasMaxParallelTasks() bool`

HasMaxParallelTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


