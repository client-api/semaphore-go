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

// checks if the ProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectRequest{}

// ProjectRequest struct for ProjectRequest
type ProjectRequest struct {
	Name *string `json:"name,omitempty"`
	Alert *bool `json:"alert,omitempty"`
	MaxParallelTasks *int32 `json:"max_parallel_tasks,omitempty"`
	// Create Demo project resources?
	Demo *bool `json:"demo,omitempty"`
}

// NewProjectRequest instantiates a new ProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRequest() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// NewProjectRequestWithDefaults instantiates a new ProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRequestWithDefaults() *ProjectRequest {
	this := ProjectRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectRequest) SetName(v string) {
	o.Name = &v
}

// GetAlert returns the Alert field value if set, zero value otherwise.
func (o *ProjectRequest) GetAlert() bool {
	if o == nil || IsNil(o.Alert) {
		var ret bool
		return ret
	}
	return *o.Alert
}

// GetAlertOk returns a tuple with the Alert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetAlertOk() (*bool, bool) {
	if o == nil || IsNil(o.Alert) {
		return nil, false
	}
	return o.Alert, true
}

// HasAlert returns a boolean if a field has been set.
func (o *ProjectRequest) HasAlert() bool {
	if o != nil && !IsNil(o.Alert) {
		return true
	}

	return false
}

// SetAlert gets a reference to the given bool and assigns it to the Alert field.
func (o *ProjectRequest) SetAlert(v bool) {
	o.Alert = &v
}

// GetMaxParallelTasks returns the MaxParallelTasks field value if set, zero value otherwise.
func (o *ProjectRequest) GetMaxParallelTasks() int32 {
	if o == nil || IsNil(o.MaxParallelTasks) {
		var ret int32
		return ret
	}
	return *o.MaxParallelTasks
}

// GetMaxParallelTasksOk returns a tuple with the MaxParallelTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetMaxParallelTasksOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxParallelTasks) {
		return nil, false
	}
	return o.MaxParallelTasks, true
}

// HasMaxParallelTasks returns a boolean if a field has been set.
func (o *ProjectRequest) HasMaxParallelTasks() bool {
	if o != nil && !IsNil(o.MaxParallelTasks) {
		return true
	}

	return false
}

// SetMaxParallelTasks gets a reference to the given int32 and assigns it to the MaxParallelTasks field.
func (o *ProjectRequest) SetMaxParallelTasks(v int32) {
	o.MaxParallelTasks = &v
}

// GetDemo returns the Demo field value if set, zero value otherwise.
func (o *ProjectRequest) GetDemo() bool {
	if o == nil || IsNil(o.Demo) {
		var ret bool
		return ret
	}
	return *o.Demo
}

// GetDemoOk returns a tuple with the Demo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRequest) GetDemoOk() (*bool, bool) {
	if o == nil || IsNil(o.Demo) {
		return nil, false
	}
	return o.Demo, true
}

// HasDemo returns a boolean if a field has been set.
func (o *ProjectRequest) HasDemo() bool {
	if o != nil && !IsNil(o.Demo) {
		return true
	}

	return false
}

// SetDemo gets a reference to the given bool and assigns it to the Demo field.
func (o *ProjectRequest) SetDemo(v bool) {
	o.Demo = &v
}

func (o ProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Alert) {
		toSerialize["alert"] = o.Alert
	}
	if !IsNil(o.MaxParallelTasks) {
		toSerialize["max_parallel_tasks"] = o.MaxParallelTasks
	}
	if !IsNil(o.Demo) {
		toSerialize["demo"] = o.Demo
	}
	return toSerialize, nil
}

type NullableProjectRequest struct {
	value *ProjectRequest
	isSet bool
}

func (v NullableProjectRequest) Get() *ProjectRequest {
	return v.value
}

func (v *NullableProjectRequest) Set(val *ProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRequest(val *ProjectRequest) *NullableProjectRequest {
	return &NullableProjectRequest{value: val, isSet: true}
}

func (v NullableProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


