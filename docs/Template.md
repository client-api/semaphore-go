# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**InventoryId** | Pointer to **int32** |  | [optional] 
**RepositoryId** | Pointer to **int32** |  | [optional] 
**EnvironmentId** | Pointer to **int32** |  | [optional] 
**ViewId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Playbook** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AllowOverrideArgsInTask** | Pointer to **bool** |  | [optional] 
**SuppressSuccessAlerts** | Pointer to **bool** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Autorun** | Pointer to **bool** |  | [optional] 
**SurveyVars** | Pointer to [**[]TemplateSurveyVar**](TemplateSurveyVar.md) |  | [optional] 
**Vaults** | Pointer to [**[]TemplateVault**](TemplateVault.md) |  | [optional] 

## Methods

### NewTemplate

`func NewTemplate() *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Template) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *Template) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Template) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Template) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Template) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetInventoryId

`func (o *Template) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *Template) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *Template) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *Template) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *Template) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Template) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Template) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Template) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *Template) GetEnvironmentId() int32`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Template) GetEnvironmentIdOk() (*int32, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Template) SetEnvironmentId(v int32)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Template) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetViewId

`func (o *Template) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *Template) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *Template) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *Template) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaybook

`func (o *Template) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *Template) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *Template) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *Template) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetArguments

`func (o *Template) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Template) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Template) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *Template) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Template) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowOverrideArgsInTask

`func (o *Template) GetAllowOverrideArgsInTask() bool`

GetAllowOverrideArgsInTask returns the AllowOverrideArgsInTask field if non-nil, zero value otherwise.

### GetAllowOverrideArgsInTaskOk

`func (o *Template) GetAllowOverrideArgsInTaskOk() (*bool, bool)`

GetAllowOverrideArgsInTaskOk returns a tuple with the AllowOverrideArgsInTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverrideArgsInTask

`func (o *Template) SetAllowOverrideArgsInTask(v bool)`

SetAllowOverrideArgsInTask sets AllowOverrideArgsInTask field to given value.

### HasAllowOverrideArgsInTask

`func (o *Template) HasAllowOverrideArgsInTask() bool`

HasAllowOverrideArgsInTask returns a boolean if a field has been set.

### GetSuppressSuccessAlerts

`func (o *Template) GetSuppressSuccessAlerts() bool`

GetSuppressSuccessAlerts returns the SuppressSuccessAlerts field if non-nil, zero value otherwise.

### GetSuppressSuccessAlertsOk

`func (o *Template) GetSuppressSuccessAlertsOk() (*bool, bool)`

GetSuppressSuccessAlertsOk returns a tuple with the SuppressSuccessAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressSuccessAlerts

`func (o *Template) SetSuppressSuccessAlerts(v bool)`

SetSuppressSuccessAlerts sets SuppressSuccessAlerts field to given value.

### HasSuppressSuccessAlerts

`func (o *Template) HasSuppressSuccessAlerts() bool`

HasSuppressSuccessAlerts returns a boolean if a field has been set.

### GetApp

`func (o *Template) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Template) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Template) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *Template) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetGitBranch

`func (o *Template) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *Template) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *Template) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *Template) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetType

`func (o *Template) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Template) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Template) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Template) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutorun

`func (o *Template) GetAutorun() bool`

GetAutorun returns the Autorun field if non-nil, zero value otherwise.

### GetAutorunOk

`func (o *Template) GetAutorunOk() (*bool, bool)`

GetAutorunOk returns a tuple with the Autorun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorun

`func (o *Template) SetAutorun(v bool)`

SetAutorun sets Autorun field to given value.

### HasAutorun

`func (o *Template) HasAutorun() bool`

HasAutorun returns a boolean if a field has been set.

### GetSurveyVars

`func (o *Template) GetSurveyVars() []TemplateSurveyVar`

GetSurveyVars returns the SurveyVars field if non-nil, zero value otherwise.

### GetSurveyVarsOk

`func (o *Template) GetSurveyVarsOk() (*[]TemplateSurveyVar, bool)`

GetSurveyVarsOk returns a tuple with the SurveyVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyVars

`func (o *Template) SetSurveyVars(v []TemplateSurveyVar)`

SetSurveyVars sets SurveyVars field to given value.

### HasSurveyVars

`func (o *Template) HasSurveyVars() bool`

HasSurveyVars returns a boolean if a field has been set.

### GetVaults

`func (o *Template) GetVaults() []TemplateVault`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *Template) GetVaultsOk() (*[]TemplateVault, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *Template) SetVaults(v []TemplateVault)`

SetVaults sets Vaults field to given value.

### HasVaults

`func (o *Template) HasVaults() bool`

HasVaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


