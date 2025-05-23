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

// checks if the ProjectProjectIdUsersPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectProjectIdUsersPostRequest{}

// ProjectProjectIdUsersPostRequest struct for ProjectProjectIdUsersPostRequest
type ProjectProjectIdUsersPostRequest struct {
	UserId *int32 `json:"user_id,omitempty"`
	Role *string `json:"role,omitempty"`
}

// NewProjectProjectIdUsersPostRequest instantiates a new ProjectProjectIdUsersPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectIdUsersPostRequest() *ProjectProjectIdUsersPostRequest {
	this := ProjectProjectIdUsersPostRequest{}
	return &this
}

// NewProjectProjectIdUsersPostRequestWithDefaults instantiates a new ProjectProjectIdUsersPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectIdUsersPostRequestWithDefaults() *ProjectProjectIdUsersPostRequest {
	this := ProjectProjectIdUsersPostRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ProjectProjectIdUsersPostRequest) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdUsersPostRequest) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ProjectProjectIdUsersPostRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ProjectProjectIdUsersPostRequest) SetUserId(v int32) {
	o.UserId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ProjectProjectIdUsersPostRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdUsersPostRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ProjectProjectIdUsersPostRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ProjectProjectIdUsersPostRequest) SetRole(v string) {
	o.Role = &v
}

func (o ProjectProjectIdUsersPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectProjectIdUsersPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableProjectProjectIdUsersPostRequest struct {
	value *ProjectProjectIdUsersPostRequest
	isSet bool
}

func (v NullableProjectProjectIdUsersPostRequest) Get() *ProjectProjectIdUsersPostRequest {
	return v.value
}

func (v *NullableProjectProjectIdUsersPostRequest) Set(val *ProjectProjectIdUsersPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectIdUsersPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectIdUsersPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectIdUsersPostRequest(val *ProjectProjectIdUsersPostRequest) *NullableProjectProjectIdUsersPostRequest {
	return &NullableProjectProjectIdUsersPostRequest{value: val, isSet: true}
}

func (v NullableProjectProjectIdUsersPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectIdUsersPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


