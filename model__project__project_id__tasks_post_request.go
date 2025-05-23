/*
Semaphore API

Semaphore API provides endpoints for managing and interacting with the Semaphore UI. This documentation outlines the available operations and data models. 

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package semaphore

import (
	"encoding/json"
)

// checks if the ProjectProjectIdTasksPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectProjectIdTasksPostRequest{}

// ProjectProjectIdTasksPostRequest struct for ProjectProjectIdTasksPostRequest
type ProjectProjectIdTasksPostRequest struct {
	TemplateId *int32 `json:"template_id,omitempty"`
	Debug *bool `json:"debug,omitempty"`
	DryRun *bool `json:"dry_run,omitempty"`
	Diff *bool `json:"diff,omitempty"`
	Playbook *string `json:"playbook,omitempty"`
	Environment *string `json:"environment,omitempty"`
	Limit *string `json:"limit,omitempty"`
	GitBranch *string `json:"git_branch,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewProjectProjectIdTasksPostRequest instantiates a new ProjectProjectIdTasksPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectIdTasksPostRequest() *ProjectProjectIdTasksPostRequest {
	this := ProjectProjectIdTasksPostRequest{}
	return &this
}

// NewProjectProjectIdTasksPostRequestWithDefaults instantiates a new ProjectProjectIdTasksPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectIdTasksPostRequestWithDefaults() *ProjectProjectIdTasksPostRequest {
	this := ProjectProjectIdTasksPostRequest{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ProjectProjectIdTasksPostRequest) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetDebug() bool {
	if o == nil || IsNil(o.Debug) {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *ProjectProjectIdTasksPostRequest) SetDebug(v bool) {
	o.Debug = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ProjectProjectIdTasksPostRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetDiff() bool {
	if o == nil || IsNil(o.Diff) {
		var ret bool
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetDiffOk() (*bool, bool) {
	if o == nil || IsNil(o.Diff) {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasDiff() bool {
	if o != nil && !IsNil(o.Diff) {
		return true
	}

	return false
}

// SetDiff gets a reference to the given bool and assigns it to the Diff field.
func (o *ProjectProjectIdTasksPostRequest) SetDiff(v bool) {
	o.Diff = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetPlaybook() string {
	if o == nil || IsNil(o.Playbook) {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetPlaybookOk() (*string, bool) {
	if o == nil || IsNil(o.Playbook) {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasPlaybook() bool {
	if o != nil && !IsNil(o.Playbook) {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *ProjectProjectIdTasksPostRequest) SetPlaybook(v string) {
	o.Playbook = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ProjectProjectIdTasksPostRequest) SetEnvironment(v string) {
	o.Environment = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetLimit() string {
	if o == nil || IsNil(o.Limit) {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetLimitOk() (*string, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *ProjectProjectIdTasksPostRequest) SetLimit(v string) {
	o.Limit = &v
}

// GetGitBranch returns the GitBranch field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetGitBranch() string {
	if o == nil || IsNil(o.GitBranch) {
		var ret string
		return ret
	}
	return *o.GitBranch
}

// GetGitBranchOk returns a tuple with the GitBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetGitBranchOk() (*string, bool) {
	if o == nil || IsNil(o.GitBranch) {
		return nil, false
	}
	return o.GitBranch, true
}

// HasGitBranch returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasGitBranch() bool {
	if o != nil && !IsNil(o.GitBranch) {
		return true
	}

	return false
}

// SetGitBranch gets a reference to the given string and assigns it to the GitBranch field.
func (o *ProjectProjectIdTasksPostRequest) SetGitBranch(v string) {
	o.GitBranch = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProjectProjectIdTasksPostRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdTasksPostRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProjectProjectIdTasksPostRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProjectProjectIdTasksPostRequest) SetMessage(v string) {
	o.Message = &v
}

func (o ProjectProjectIdTasksPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectProjectIdTasksPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	if !IsNil(o.DryRun) {
		toSerialize["dry_run"] = o.DryRun
	}
	if !IsNil(o.Diff) {
		toSerialize["diff"] = o.Diff
	}
	if !IsNil(o.Playbook) {
		toSerialize["playbook"] = o.Playbook
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.GitBranch) {
		toSerialize["git_branch"] = o.GitBranch
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableProjectProjectIdTasksPostRequest struct {
	value *ProjectProjectIdTasksPostRequest
	isSet bool
}

func (v NullableProjectProjectIdTasksPostRequest) Get() *ProjectProjectIdTasksPostRequest {
	return v.value
}

func (v *NullableProjectProjectIdTasksPostRequest) Set(val *ProjectProjectIdTasksPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectIdTasksPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectIdTasksPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectIdTasksPostRequest(val *ProjectProjectIdTasksPostRequest) *NullableProjectProjectIdTasksPostRequest {
	return &NullableProjectProjectIdTasksPostRequest{value: val, isSet: true}
}

func (v NullableProjectProjectIdTasksPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectIdTasksPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


