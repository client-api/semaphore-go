# \ScheduleAPI

All URIs are relative to *http://localhost:3000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectProjectIdSchedulesPost**](ScheduleAPI.md#ProjectProjectIdSchedulesPost) | **Post** /project/{project_id}/schedules | create schedule
[**ProjectProjectIdSchedulesScheduleIdDelete**](ScheduleAPI.md#ProjectProjectIdSchedulesScheduleIdDelete) | **Delete** /project/{project_id}/schedules/{schedule_id} | Deletes schedule
[**ProjectProjectIdSchedulesScheduleIdGet**](ScheduleAPI.md#ProjectProjectIdSchedulesScheduleIdGet) | **Get** /project/{project_id}/schedules/{schedule_id} | Get schedule
[**ProjectProjectIdSchedulesScheduleIdPut**](ScheduleAPI.md#ProjectProjectIdSchedulesScheduleIdPut) | **Put** /project/{project_id}/schedules/{schedule_id} | Updates schedule



## ProjectProjectIdSchedulesPost

> Schedule ProjectProjectIdSchedulesPost(ctx, projectId).Schedule(schedule).Execute()

create schedule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/client-api/semaphore-go"
)

func main() {
	projectId := int32(56) // int32 | Project ID
	schedule := *openapiclient.NewScheduleRequest() // ScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.ProjectProjectIdSchedulesPost(context.Background(), projectId).Schedule(schedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ProjectProjectIdSchedulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdSchedulesPost`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.ProjectProjectIdSchedulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedule** | [**ScheduleRequest**](ScheduleRequest.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdSchedulesScheduleIdDelete

> ProjectProjectIdSchedulesScheduleIdDelete(ctx, projectId, scheduleId).Execute()

Deletes schedule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/client-api/semaphore-go"
)

func main() {
	projectId := int32(56) // int32 | Project ID
	scheduleId := int32(56) // int32 | schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduleAPI.ProjectProjectIdSchedulesScheduleIdDelete(context.Background(), projectId, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ProjectProjectIdSchedulesScheduleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**scheduleId** | **int32** | schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesScheduleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdSchedulesScheduleIdGet

> Schedule ProjectProjectIdSchedulesScheduleIdGet(ctx, projectId, scheduleId).Execute()

Get schedule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/client-api/semaphore-go"
)

func main() {
	projectId := int32(56) // int32 | Project ID
	scheduleId := int32(56) // int32 | schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleAPI.ProjectProjectIdSchedulesScheduleIdGet(context.Background(), projectId, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ProjectProjectIdSchedulesScheduleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdSchedulesScheduleIdGet`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleAPI.ProjectProjectIdSchedulesScheduleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**scheduleId** | **int32** | schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesScheduleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schedule**](Schedule.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdSchedulesScheduleIdPut

> ProjectProjectIdSchedulesScheduleIdPut(ctx, projectId, scheduleId).Schedule(schedule).Execute()

Updates schedule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/client-api/semaphore-go"
)

func main() {
	projectId := int32(56) // int32 | Project ID
	scheduleId := int32(56) // int32 | schedule ID
	schedule := *openapiclient.NewScheduleRequest() // ScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduleAPI.ProjectProjectIdSchedulesScheduleIdPut(context.Background(), projectId, scheduleId).Schedule(schedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleAPI.ProjectProjectIdSchedulesScheduleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**scheduleId** | **int32** | schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSchedulesScheduleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schedule** | [**ScheduleRequest**](ScheduleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

