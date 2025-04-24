# AccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**OverrideSecret** | Pointer to **bool** |  | [optional] 
**LoginPassword** | Pointer to [**AccessKeyRequestLoginPassword**](AccessKeyRequestLoginPassword.md) |  | [optional] 
**Ssh** | Pointer to [**AccessKeyRequestSsh**](AccessKeyRequestSsh.md) |  | [optional] 

## Methods

### NewAccessKeyRequest

`func NewAccessKeyRequest() *AccessKeyRequest`

NewAccessKeyRequest instantiates a new AccessKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyRequestWithDefaults

`func NewAccessKeyRequestWithDefaults() *AccessKeyRequest`

NewAccessKeyRequestWithDefaults instantiates a new AccessKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessKeyRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKeyRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKeyRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessKeyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AccessKeyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessKeyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessKeyRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessKeyRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProjectId

`func (o *AccessKeyRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AccessKeyRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AccessKeyRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AccessKeyRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetOverrideSecret

`func (o *AccessKeyRequest) GetOverrideSecret() bool`

GetOverrideSecret returns the OverrideSecret field if non-nil, zero value otherwise.

### GetOverrideSecretOk

`func (o *AccessKeyRequest) GetOverrideSecretOk() (*bool, bool)`

GetOverrideSecretOk returns a tuple with the OverrideSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSecret

`func (o *AccessKeyRequest) SetOverrideSecret(v bool)`

SetOverrideSecret sets OverrideSecret field to given value.

### HasOverrideSecret

`func (o *AccessKeyRequest) HasOverrideSecret() bool`

HasOverrideSecret returns a boolean if a field has been set.

### GetLoginPassword

`func (o *AccessKeyRequest) GetLoginPassword() AccessKeyRequestLoginPassword`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *AccessKeyRequest) GetLoginPasswordOk() (*AccessKeyRequestLoginPassword, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *AccessKeyRequest) SetLoginPassword(v AccessKeyRequestLoginPassword)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *AccessKeyRequest) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetSsh

`func (o *AccessKeyRequest) GetSsh() AccessKeyRequestSsh`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *AccessKeyRequest) GetSshOk() (*AccessKeyRequestSsh, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *AccessKeyRequest) SetSsh(v AccessKeyRequestSsh)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *AccessKeyRequest) HasSsh() bool`

HasSsh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


