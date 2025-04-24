# EnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**Secrets** | Pointer to [**[]EnvironmentSecretRequest**](EnvironmentSecretRequest.md) |  | [optional] 

## Methods

### NewEnvironmentRequest

`func NewEnvironmentRequest() *EnvironmentRequest`

NewEnvironmentRequest instantiates a new EnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentRequestWithDefaults

`func NewEnvironmentRequestWithDefaults() *EnvironmentRequest`

NewEnvironmentRequestWithDefaults instantiates a new EnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *EnvironmentRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnvironmentRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnvironmentRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *EnvironmentRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPassword

`func (o *EnvironmentRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnvironmentRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnvironmentRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnvironmentRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetJson

`func (o *EnvironmentRequest) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EnvironmentRequest) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EnvironmentRequest) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *EnvironmentRequest) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetEnv

`func (o *EnvironmentRequest) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EnvironmentRequest) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EnvironmentRequest) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EnvironmentRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetSecrets

`func (o *EnvironmentRequest) GetSecrets() []EnvironmentSecretRequest`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *EnvironmentRequest) GetSecretsOk() (*[]EnvironmentSecretRequest, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *EnvironmentRequest) SetSecrets(v []EnvironmentSecretRequest)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *EnvironmentRequest) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


