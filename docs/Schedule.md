# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CronFormat** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Schedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCronFormat

`func (o *Schedule) GetCronFormat() string`

GetCronFormat returns the CronFormat field if non-nil, zero value otherwise.

### GetCronFormatOk

`func (o *Schedule) GetCronFormatOk() (*string, bool)`

GetCronFormatOk returns a tuple with the CronFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronFormat

`func (o *Schedule) SetCronFormat(v string)`

SetCronFormat sets CronFormat field to given value.

### HasCronFormat

`func (o *Schedule) HasCronFormat() bool`

HasCronFormat returns a boolean if a field has been set.

### GetProjectId

`func (o *Schedule) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Schedule) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Schedule) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Schedule) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTemplateId

`func (o *Schedule) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Schedule) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Schedule) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Schedule) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetName

`func (o *Schedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Schedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *Schedule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Schedule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Schedule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Schedule) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


