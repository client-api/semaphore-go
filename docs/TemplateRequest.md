# TemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**InventoryId** | Pointer to **int32** |  | [optional] 
**RepositoryId** | Pointer to **int32** |  | [optional] 
**EnvironmentId** | Pointer to **int32** |  | [optional] 
**ViewId** | Pointer to **int32** |  | [optional] 
**Vaults** | Pointer to [**[]TemplateVault**](TemplateVault.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Playbook** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AllowOverrideArgsInTask** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**SuppressSuccessAlerts** | Pointer to **bool** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**SurveyVars** | Pointer to [**[]TemplateSurveyVar**](TemplateSurveyVar.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**StartVersion** | Pointer to **string** |  | [optional] 
**BuildTemplateId** | Pointer to **int32** |  | [optional] 
**Autorun** | Pointer to **bool** |  | [optional] 

## Methods

### NewTemplateRequest

`func NewTemplateRequest() *TemplateRequest`

NewTemplateRequest instantiates a new TemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateRequestWithDefaults

`func NewTemplateRequestWithDefaults() *TemplateRequest`

NewTemplateRequestWithDefaults instantiates a new TemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *TemplateRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TemplateRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TemplateRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TemplateRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetInventoryId

`func (o *TemplateRequest) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *TemplateRequest) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *TemplateRequest) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *TemplateRequest) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetRepositoryId

`func (o *TemplateRequest) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *TemplateRequest) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *TemplateRequest) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *TemplateRequest) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *TemplateRequest) GetEnvironmentId() int32`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *TemplateRequest) GetEnvironmentIdOk() (*int32, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *TemplateRequest) SetEnvironmentId(v int32)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *TemplateRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetViewId

`func (o *TemplateRequest) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *TemplateRequest) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *TemplateRequest) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *TemplateRequest) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetVaults

`func (o *TemplateRequest) GetVaults() []TemplateVault`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *TemplateRequest) GetVaultsOk() (*[]TemplateVault, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *TemplateRequest) SetVaults(v []TemplateVault)`

SetVaults sets Vaults field to given value.

### HasVaults

`func (o *TemplateRequest) HasVaults() bool`

HasVaults returns a boolean if a field has been set.

### GetName

`func (o *TemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaybook

`func (o *TemplateRequest) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *TemplateRequest) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *TemplateRequest) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *TemplateRequest) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetArguments

`func (o *TemplateRequest) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *TemplateRequest) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *TemplateRequest) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *TemplateRequest) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowOverrideArgsInTask

`func (o *TemplateRequest) GetAllowOverrideArgsInTask() bool`

GetAllowOverrideArgsInTask returns the AllowOverrideArgsInTask field if non-nil, zero value otherwise.

### GetAllowOverrideArgsInTaskOk

`func (o *TemplateRequest) GetAllowOverrideArgsInTaskOk() (*bool, bool)`

GetAllowOverrideArgsInTaskOk returns a tuple with the AllowOverrideArgsInTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverrideArgsInTask

`func (o *TemplateRequest) SetAllowOverrideArgsInTask(v bool)`

SetAllowOverrideArgsInTask sets AllowOverrideArgsInTask field to given value.

### HasAllowOverrideArgsInTask

`func (o *TemplateRequest) HasAllowOverrideArgsInTask() bool`

HasAllowOverrideArgsInTask returns a boolean if a field has been set.

### GetLimit

`func (o *TemplateRequest) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TemplateRequest) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TemplateRequest) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TemplateRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSuppressSuccessAlerts

`func (o *TemplateRequest) GetSuppressSuccessAlerts() bool`

GetSuppressSuccessAlerts returns the SuppressSuccessAlerts field if non-nil, zero value otherwise.

### GetSuppressSuccessAlertsOk

`func (o *TemplateRequest) GetSuppressSuccessAlertsOk() (*bool, bool)`

GetSuppressSuccessAlertsOk returns a tuple with the SuppressSuccessAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressSuccessAlerts

`func (o *TemplateRequest) SetSuppressSuccessAlerts(v bool)`

SetSuppressSuccessAlerts sets SuppressSuccessAlerts field to given value.

### HasSuppressSuccessAlerts

`func (o *TemplateRequest) HasSuppressSuccessAlerts() bool`

HasSuppressSuccessAlerts returns a boolean if a field has been set.

### GetApp

`func (o *TemplateRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *TemplateRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *TemplateRequest) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *TemplateRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetGitBranch

`func (o *TemplateRequest) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *TemplateRequest) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *TemplateRequest) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *TemplateRequest) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetSurveyVars

`func (o *TemplateRequest) GetSurveyVars() []TemplateSurveyVar`

GetSurveyVars returns the SurveyVars field if non-nil, zero value otherwise.

### GetSurveyVarsOk

`func (o *TemplateRequest) GetSurveyVarsOk() (*[]TemplateSurveyVar, bool)`

GetSurveyVarsOk returns a tuple with the SurveyVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyVars

`func (o *TemplateRequest) SetSurveyVars(v []TemplateSurveyVar)`

SetSurveyVars sets SurveyVars field to given value.

### HasSurveyVars

`func (o *TemplateRequest) HasSurveyVars() bool`

HasSurveyVars returns a boolean if a field has been set.

### GetType

`func (o *TemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartVersion

`func (o *TemplateRequest) GetStartVersion() string`

GetStartVersion returns the StartVersion field if non-nil, zero value otherwise.

### GetStartVersionOk

`func (o *TemplateRequest) GetStartVersionOk() (*string, bool)`

GetStartVersionOk returns a tuple with the StartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVersion

`func (o *TemplateRequest) SetStartVersion(v string)`

SetStartVersion sets StartVersion field to given value.

### HasStartVersion

`func (o *TemplateRequest) HasStartVersion() bool`

HasStartVersion returns a boolean if a field has been set.

### GetBuildTemplateId

`func (o *TemplateRequest) GetBuildTemplateId() int32`

GetBuildTemplateId returns the BuildTemplateId field if non-nil, zero value otherwise.

### GetBuildTemplateIdOk

`func (o *TemplateRequest) GetBuildTemplateIdOk() (*int32, bool)`

GetBuildTemplateIdOk returns a tuple with the BuildTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTemplateId

`func (o *TemplateRequest) SetBuildTemplateId(v int32)`

SetBuildTemplateId sets BuildTemplateId field to given value.

### HasBuildTemplateId

`func (o *TemplateRequest) HasBuildTemplateId() bool`

HasBuildTemplateId returns a boolean if a field has been set.

### GetAutorun

`func (o *TemplateRequest) GetAutorun() bool`

GetAutorun returns the Autorun field if non-nil, zero value otherwise.

### GetAutorunOk

`func (o *TemplateRequest) GetAutorunOk() (*bool, bool)`

GetAutorunOk returns a tuple with the Autorun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorun

`func (o *TemplateRequest) SetAutorun(v bool)`

SetAutorun sets Autorun field to given value.

### HasAutorun

`func (o *TemplateRequest) HasAutorun() bool`

HasAutorun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


