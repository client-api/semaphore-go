# ProjectBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ProjectBackupMeta**](ProjectBackupMeta.md) |  | [optional] 
**Templates** | Pointer to [**[]ProjectBackupTemplatesInner**](ProjectBackupTemplatesInner.md) |  | [optional] 
**Repositories** | Pointer to [**[]ProjectBackupRepositoriesInner**](ProjectBackupRepositoriesInner.md) |  | [optional] 
**Keys** | Pointer to [**[]ProjectBackupKeysInner**](ProjectBackupKeysInner.md) |  | [optional] 
**Views** | Pointer to [**[]ProjectBackupViewsInner**](ProjectBackupViewsInner.md) |  | [optional] 
**Inventories** | Pointer to [**[]ProjectBackupInventoriesInner**](ProjectBackupInventoriesInner.md) |  | [optional] 
**Environments** | Pointer to [**[]ProjectBackupEnvironmentsInner**](ProjectBackupEnvironmentsInner.md) |  | [optional] 

## Methods

### NewProjectBackup

`func NewProjectBackup() *ProjectBackup`

NewProjectBackup instantiates a new ProjectBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectBackupWithDefaults

`func NewProjectBackupWithDefaults() *ProjectBackup`

NewProjectBackupWithDefaults instantiates a new ProjectBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ProjectBackup) GetMeta() ProjectBackupMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectBackup) GetMetaOk() (*ProjectBackupMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectBackup) SetMeta(v ProjectBackupMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProjectBackup) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTemplates

`func (o *ProjectBackup) GetTemplates() []ProjectBackupTemplatesInner`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ProjectBackup) GetTemplatesOk() (*[]ProjectBackupTemplatesInner, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ProjectBackup) SetTemplates(v []ProjectBackupTemplatesInner)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *ProjectBackup) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetRepositories

`func (o *ProjectBackup) GetRepositories() []ProjectBackupRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ProjectBackup) GetRepositoriesOk() (*[]ProjectBackupRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ProjectBackup) SetRepositories(v []ProjectBackupRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ProjectBackup) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetKeys

`func (o *ProjectBackup) GetKeys() []ProjectBackupKeysInner`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ProjectBackup) GetKeysOk() (*[]ProjectBackupKeysInner, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ProjectBackup) SetKeys(v []ProjectBackupKeysInner)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ProjectBackup) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetViews

`func (o *ProjectBackup) GetViews() []ProjectBackupViewsInner`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ProjectBackup) GetViewsOk() (*[]ProjectBackupViewsInner, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ProjectBackup) SetViews(v []ProjectBackupViewsInner)`

SetViews sets Views field to given value.

### HasViews

`func (o *ProjectBackup) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetInventories

`func (o *ProjectBackup) GetInventories() []ProjectBackupInventoriesInner`

GetInventories returns the Inventories field if non-nil, zero value otherwise.

### GetInventoriesOk

`func (o *ProjectBackup) GetInventoriesOk() (*[]ProjectBackupInventoriesInner, bool)`

GetInventoriesOk returns a tuple with the Inventories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventories

`func (o *ProjectBackup) SetInventories(v []ProjectBackupInventoriesInner)`

SetInventories sets Inventories field to given value.

### HasInventories

`func (o *ProjectBackup) HasInventories() bool`

HasInventories returns a boolean if a field has been set.

### GetEnvironments

`func (o *ProjectBackup) GetEnvironments() []ProjectBackupEnvironmentsInner`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectBackup) GetEnvironmentsOk() (*[]ProjectBackupEnvironmentsInner, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectBackup) SetEnvironments(v []ProjectBackupEnvironmentsInner)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectBackup) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


