# TemplateSurveyVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to [**[]TemplateSurveyVarValue**](TemplateSurveyVarValue.md) |  | [optional] 

## Methods

### NewTemplateSurveyVar

`func NewTemplateSurveyVar() *TemplateSurveyVar`

NewTemplateSurveyVar instantiates a new TemplateSurveyVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSurveyVarWithDefaults

`func NewTemplateSurveyVarWithDefaults() *TemplateSurveyVar`

NewTemplateSurveyVarWithDefaults instantiates a new TemplateSurveyVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateSurveyVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateSurveyVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateSurveyVar) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateSurveyVar) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *TemplateSurveyVar) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateSurveyVar) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateSurveyVar) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TemplateSurveyVar) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateSurveyVar) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateSurveyVar) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateSurveyVar) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateSurveyVar) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *TemplateSurveyVar) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateSurveyVar) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateSurveyVar) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateSurveyVar) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequired

`func (o *TemplateSurveyVar) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TemplateSurveyVar) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TemplateSurveyVar) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *TemplateSurveyVar) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValues

`func (o *TemplateSurveyVar) GetValues() []TemplateSurveyVarValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TemplateSurveyVar) GetValuesOk() (*[]TemplateSurveyVarValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TemplateSurveyVar) SetValues(v []TemplateSurveyVarValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *TemplateSurveyVar) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


