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

// checks if the ProjectBackupEnvironmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectBackupEnvironmentsInner{}

// ProjectBackupEnvironmentsInner struct for ProjectBackupEnvironmentsInner
type ProjectBackupEnvironmentsInner struct {
	Name *string `json:"name,omitempty"`
	Json *string `json:"json,omitempty"`
}

// NewProjectBackupEnvironmentsInner instantiates a new ProjectBackupEnvironmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectBackupEnvironmentsInner() *ProjectBackupEnvironmentsInner {
	this := ProjectBackupEnvironmentsInner{}
	return &this
}

// NewProjectBackupEnvironmentsInnerWithDefaults instantiates a new ProjectBackupEnvironmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectBackupEnvironmentsInnerWithDefaults() *ProjectBackupEnvironmentsInner {
	this := ProjectBackupEnvironmentsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectBackupEnvironmentsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBackupEnvironmentsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectBackupEnvironmentsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectBackupEnvironmentsInner) SetName(v string) {
	o.Name = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ProjectBackupEnvironmentsInner) GetJson() string {
	if o == nil || IsNil(o.Json) {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBackupEnvironmentsInner) GetJsonOk() (*string, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ProjectBackupEnvironmentsInner) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *ProjectBackupEnvironmentsInner) SetJson(v string) {
	o.Json = &v
}

func (o ProjectBackupEnvironmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectBackupEnvironmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	return toSerialize, nil
}

type NullableProjectBackupEnvironmentsInner struct {
	value *ProjectBackupEnvironmentsInner
	isSet bool
}

func (v NullableProjectBackupEnvironmentsInner) Get() *ProjectBackupEnvironmentsInner {
	return v.value
}

func (v *NullableProjectBackupEnvironmentsInner) Set(val *ProjectBackupEnvironmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectBackupEnvironmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectBackupEnvironmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectBackupEnvironmentsInner(val *ProjectBackupEnvironmentsInner) *NullableProjectBackupEnvironmentsInner {
	return &NullableProjectBackupEnvironmentsInner{value: val, isSet: true}
}

func (v NullableProjectBackupEnvironmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectBackupEnvironmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


