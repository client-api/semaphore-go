# \ProjectAPI

All URIs are relative to *http://localhost:3000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectProjectIdBackupGet**](ProjectAPI.md#ProjectProjectIdBackupGet) | **Get** /project/{project_id}/backup | Backup A Project
[**ProjectProjectIdDelete**](ProjectAPI.md#ProjectProjectIdDelete) | **Delete** /project/{project_id}/ | Delete project
[**ProjectProjectIdEnvironmentEnvironmentIdDelete**](ProjectAPI.md#ProjectProjectIdEnvironmentEnvironmentIdDelete) | **Delete** /project/{project_id}/environment/{environment_id} | Removes environment
[**ProjectProjectIdEnvironmentEnvironmentIdGet**](ProjectAPI.md#ProjectProjectIdEnvironmentEnvironmentIdGet) | **Get** /project/{project_id}/environment/{environment_id} | Get environment
[**ProjectProjectIdEnvironmentEnvironmentIdPut**](ProjectAPI.md#ProjectProjectIdEnvironmentEnvironmentIdPut) | **Put** /project/{project_id}/environment/{environment_id} | Update environment
[**ProjectProjectIdEnvironmentGet**](ProjectAPI.md#ProjectProjectIdEnvironmentGet) | **Get** /project/{project_id}/environment | Get environment
[**ProjectProjectIdEnvironmentPost**](ProjectAPI.md#ProjectProjectIdEnvironmentPost) | **Post** /project/{project_id}/environment | Add environment
[**ProjectProjectIdEventsGet**](ProjectAPI.md#ProjectProjectIdEventsGet) | **Get** /project/{project_id}/events | Get Events related to this project
[**ProjectProjectIdGet**](ProjectAPI.md#ProjectProjectIdGet) | **Get** /project/{project_id}/ | Fetch project
[**ProjectProjectIdIntegrationsGet**](ProjectAPI.md#ProjectProjectIdIntegrationsGet) | **Get** /project/{project_id}/integrations | get all integrations
[**ProjectProjectIdIntegrationsIntegrationIdDelete**](ProjectAPI.md#ProjectProjectIdIntegrationsIntegrationIdDelete) | **Delete** /project/{project_id}/integrations/{integration_id} | Remove integration
[**ProjectProjectIdIntegrationsIntegrationIdMatchersPost**](ProjectAPI.md#ProjectProjectIdIntegrationsIntegrationIdMatchersPost) | **Post** /project/{project_id}/integrations/{integration_id}/matchers | Add Integration Matcher
[**ProjectProjectIdIntegrationsIntegrationIdPut**](ProjectAPI.md#ProjectProjectIdIntegrationsIntegrationIdPut) | **Put** /project/{project_id}/integrations/{integration_id} | Update Integration
[**ProjectProjectIdIntegrationsIntegrationIdValuesPost**](ProjectAPI.md#ProjectProjectIdIntegrationsIntegrationIdValuesPost) | **Post** /project/{project_id}/integrations/{integration_id}/values | Add Integration Extracted Value
[**ProjectProjectIdIntegrationsPost**](ProjectAPI.md#ProjectProjectIdIntegrationsPost) | **Post** /project/{project_id}/integrations | create a new integration
[**ProjectProjectIdInventoryGet**](ProjectAPI.md#ProjectProjectIdInventoryGet) | **Get** /project/{project_id}/inventory | Get inventory
[**ProjectProjectIdInventoryInventoryIdDelete**](ProjectAPI.md#ProjectProjectIdInventoryInventoryIdDelete) | **Delete** /project/{project_id}/inventory/{inventory_id} | Removes inventory
[**ProjectProjectIdInventoryInventoryIdGet**](ProjectAPI.md#ProjectProjectIdInventoryInventoryIdGet) | **Get** /project/{project_id}/inventory/{inventory_id} | Get inventory
[**ProjectProjectIdInventoryInventoryIdPut**](ProjectAPI.md#ProjectProjectIdInventoryInventoryIdPut) | **Put** /project/{project_id}/inventory/{inventory_id} | Updates inventory
[**ProjectProjectIdInventoryPost**](ProjectAPI.md#ProjectProjectIdInventoryPost) | **Post** /project/{project_id}/inventory | create inventory
[**ProjectProjectIdKeysGet**](ProjectAPI.md#ProjectProjectIdKeysGet) | **Get** /project/{project_id}/keys | Get access keys linked to project
[**ProjectProjectIdKeysKeyIdDelete**](ProjectAPI.md#ProjectProjectIdKeysKeyIdDelete) | **Delete** /project/{project_id}/keys/{key_id} | Removes access key
[**ProjectProjectIdKeysKeyIdPut**](ProjectAPI.md#ProjectProjectIdKeysKeyIdPut) | **Put** /project/{project_id}/keys/{key_id} | Updates access key
[**ProjectProjectIdKeysPost**](ProjectAPI.md#ProjectProjectIdKeysPost) | **Post** /project/{project_id}/keys | Add access key
[**ProjectProjectIdPut**](ProjectAPI.md#ProjectProjectIdPut) | **Put** /project/{project_id}/ | Update project
[**ProjectProjectIdRepositoriesGet**](ProjectAPI.md#ProjectProjectIdRepositoriesGet) | **Get** /project/{project_id}/repositories | Get repositories
[**ProjectProjectIdRepositoriesPost**](ProjectAPI.md#ProjectProjectIdRepositoriesPost) | **Post** /project/{project_id}/repositories | Add repository
[**ProjectProjectIdRepositoriesRepositoryIdDelete**](ProjectAPI.md#ProjectProjectIdRepositoriesRepositoryIdDelete) | **Delete** /project/{project_id}/repositories/{repository_id} | Removes repository
[**ProjectProjectIdRepositoriesRepositoryIdGet**](ProjectAPI.md#ProjectProjectIdRepositoriesRepositoryIdGet) | **Get** /project/{project_id}/repositories/{repository_id} | Get repository
[**ProjectProjectIdRepositoriesRepositoryIdPut**](ProjectAPI.md#ProjectProjectIdRepositoriesRepositoryIdPut) | **Put** /project/{project_id}/repositories/{repository_id} | Updates repository
[**ProjectProjectIdRoleGet**](ProjectAPI.md#ProjectProjectIdRoleGet) | **Get** /project/{project_id}/role | Fetch permissions of the current user for project
[**ProjectProjectIdTasksGet**](ProjectAPI.md#ProjectProjectIdTasksGet) | **Get** /project/{project_id}/tasks | Get Tasks related to current project
[**ProjectProjectIdTasksLastGet**](ProjectAPI.md#ProjectProjectIdTasksLastGet) | **Get** /project/{project_id}/tasks/last | Get last 200 Tasks related to current project
[**ProjectProjectIdTasksPost**](ProjectAPI.md#ProjectProjectIdTasksPost) | **Post** /project/{project_id}/tasks | Starts a job
[**ProjectProjectIdTasksTaskIdDelete**](ProjectAPI.md#ProjectProjectIdTasksTaskIdDelete) | **Delete** /project/{project_id}/tasks/{task_id} | Deletes task (including output)
[**ProjectProjectIdTasksTaskIdGet**](ProjectAPI.md#ProjectProjectIdTasksTaskIdGet) | **Get** /project/{project_id}/tasks/{task_id} | Get a single task
[**ProjectProjectIdTasksTaskIdOutputGet**](ProjectAPI.md#ProjectProjectIdTasksTaskIdOutputGet) | **Get** /project/{project_id}/tasks/{task_id}/output | Get task output
[**ProjectProjectIdTasksTaskIdStopPost**](ProjectAPI.md#ProjectProjectIdTasksTaskIdStopPost) | **Post** /project/{project_id}/tasks/{task_id}/stop | Stop a job
[**ProjectProjectIdTemplatesGet**](ProjectAPI.md#ProjectProjectIdTemplatesGet) | **Get** /project/{project_id}/templates | Get template
[**ProjectProjectIdTemplatesPost**](ProjectAPI.md#ProjectProjectIdTemplatesPost) | **Post** /project/{project_id}/templates | create template
[**ProjectProjectIdTemplatesTemplateIdDelete**](ProjectAPI.md#ProjectProjectIdTemplatesTemplateIdDelete) | **Delete** /project/{project_id}/templates/{template_id} | Removes template
[**ProjectProjectIdTemplatesTemplateIdGet**](ProjectAPI.md#ProjectProjectIdTemplatesTemplateIdGet) | **Get** /project/{project_id}/templates/{template_id} | Get template
[**ProjectProjectIdTemplatesTemplateIdPut**](ProjectAPI.md#ProjectProjectIdTemplatesTemplateIdPut) | **Put** /project/{project_id}/templates/{template_id} | Updates template
[**ProjectProjectIdUsersGet**](ProjectAPI.md#ProjectProjectIdUsersGet) | **Get** /project/{project_id}/users | Get users linked to project
[**ProjectProjectIdUsersPost**](ProjectAPI.md#ProjectProjectIdUsersPost) | **Post** /project/{project_id}/users | Link user to project
[**ProjectProjectIdUsersUserIdDelete**](ProjectAPI.md#ProjectProjectIdUsersUserIdDelete) | **Delete** /project/{project_id}/users/{user_id} | Removes user from project
[**ProjectProjectIdUsersUserIdPut**](ProjectAPI.md#ProjectProjectIdUsersUserIdPut) | **Put** /project/{project_id}/users/{user_id} | Update user role
[**ProjectProjectIdViewsGet**](ProjectAPI.md#ProjectProjectIdViewsGet) | **Get** /project/{project_id}/views | Get view
[**ProjectProjectIdViewsPost**](ProjectAPI.md#ProjectProjectIdViewsPost) | **Post** /project/{project_id}/views | create view
[**ProjectProjectIdViewsViewIdDelete**](ProjectAPI.md#ProjectProjectIdViewsViewIdDelete) | **Delete** /project/{project_id}/views/{view_id} | Removes view
[**ProjectProjectIdViewsViewIdGet**](ProjectAPI.md#ProjectProjectIdViewsViewIdGet) | **Get** /project/{project_id}/views/{view_id} | Get view
[**ProjectProjectIdViewsViewIdPut**](ProjectAPI.md#ProjectProjectIdViewsViewIdPut) | **Put** /project/{project_id}/views/{view_id} | Updates view



## ProjectProjectIdBackupGet

> ProjectBackup ProjectProjectIdBackupGet(ctx, projectId).Execute()

Backup A Project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdBackupGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdBackupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdBackupGet`: ProjectBackup
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdBackupGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdBackupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectBackup**](ProjectBackup.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdDelete

> ProjectProjectIdDelete(ctx, projectId).Execute()

Delete project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdDelete(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdEnvironmentEnvironmentIdDelete

> ProjectProjectIdEnvironmentEnvironmentIdDelete(ctx, projectId, environmentId).Execute()

Removes environment

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
	environmentId := int32(56) // int32 | environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdDelete(context.Background(), projectId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**environmentId** | **int32** | environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentEnvironmentIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdEnvironmentEnvironmentIdGet

> Environment ProjectProjectIdEnvironmentEnvironmentIdGet(ctx, projectId, environmentId).Execute()

Get environment

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
	environmentId := int32(56) // int32 | environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdGet(context.Background(), projectId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdEnvironmentEnvironmentIdGet`: Environment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**environmentId** | **int32** | environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentEnvironmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Environment**](Environment.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdEnvironmentEnvironmentIdPut

> ProjectProjectIdEnvironmentEnvironmentIdPut(ctx, projectId, environmentId).Environment(environment).Execute()

Update environment

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
	environmentId := int32(56) // int32 | environment ID
	environment := *openapiclient.NewEnvironmentRequest() // EnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdPut(context.Background(), projectId, environmentId).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdEnvironmentEnvironmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**environmentId** | **int32** | environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentEnvironmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environment** | [**EnvironmentRequest**](EnvironmentRequest.md) |  | 

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


## ProjectProjectIdEnvironmentGet

> []Environment ProjectProjectIdEnvironmentGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get environment

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
	sort := "name" // string | sorting name
	order := "desc" // string | ordering manner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdEnvironmentGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdEnvironmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdEnvironmentGet`: []Environment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdEnvironmentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Environment**](Environment.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdEnvironmentPost

> Environment ProjectProjectIdEnvironmentPost(ctx, projectId).Environment(environment).Execute()

Add environment

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
	environment := *openapiclient.NewEnvironmentRequest() // EnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdEnvironmentPost(context.Background(), projectId).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdEnvironmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdEnvironmentPost`: Environment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdEnvironmentPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEnvironmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**EnvironmentRequest**](EnvironmentRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdEventsGet

> []Event ProjectProjectIdEventsGet(ctx, projectId).Execute()

Get Events related to this project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdEventsGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdEventsGet`: []Event
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Event**](Event.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdGet

> Project ProjectProjectIdGet(ctx, projectId).Execute()

Fetch project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdGet`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdIntegrationsGet

> []Integration ProjectProjectIdIntegrationsGet(ctx, projectId).Execute()

get all integrations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdIntegrationsGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdIntegrationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdIntegrationsGet`: []Integration
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdIntegrationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Integration**](Integration.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdIntegrationsIntegrationIdDelete

> ProjectProjectIdIntegrationsIntegrationIdDelete(ctx, projectId, integrationId).Execute()

Remove integration

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
	integrationId := int32(56) // int32 | integration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdDelete(context.Background(), projectId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**integrationId** | **int32** | integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdIntegrationsIntegrationIdMatchersPost

> ProjectProjectIdIntegrationsIntegrationIdMatchersPost(ctx, projectId, integrationId).IntegrationMatcher(integrationMatcher).Execute()

Add Integration Matcher

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
	integrationId := int32(56) // int32 | integration ID
	integrationMatcher := *openapiclient.NewIntegrationMatcher() // IntegrationMatcher | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersPost(context.Background(), projectId, integrationId).IntegrationMatcher(integrationMatcher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**integrationId** | **int32** | integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdMatchersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **integrationMatcher** | [**IntegrationMatcher**](IntegrationMatcher.md) |  | 

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


## ProjectProjectIdIntegrationsIntegrationIdPut

> ProjectProjectIdIntegrationsIntegrationIdPut(ctx, projectId, integrationId).Integration(integration).Execute()

Update Integration

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
	integrationId := int32(56) // int32 | integration ID
	integration := *openapiclient.NewIntegrationRequest() // IntegrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdPut(context.Background(), projectId, integrationId).Integration(integration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**integrationId** | **int32** | integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **integration** | [**IntegrationRequest**](IntegrationRequest.md) |  | 

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


## ProjectProjectIdIntegrationsIntegrationIdValuesPost

> ProjectProjectIdIntegrationsIntegrationIdValuesPost(ctx, projectId, integrationId).IntegrationExtractedValue(integrationExtractedValue).Execute()

Add Integration Extracted Value

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
	integrationId := int32(56) // int32 | integration ID
	integrationExtractedValue := *openapiclient.NewIntegrationExtractValue() // IntegrationExtractValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdValuesPost(context.Background(), projectId, integrationId).IntegrationExtractedValue(integrationExtractedValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdIntegrationsIntegrationIdValuesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**integrationId** | **int32** | integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdValuesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **integrationExtractedValue** | [**IntegrationExtractValue**](IntegrationExtractValue.md) |  | 

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


## ProjectProjectIdIntegrationsPost

> Integration ProjectProjectIdIntegrationsPost(ctx, projectId).Integration(integration).Execute()

create a new integration

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
	integration := *openapiclient.NewIntegrationRequest() // IntegrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdIntegrationsPost(context.Background(), projectId).Integration(integration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdIntegrationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdIntegrationsPost`: Integration
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdIntegrationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integration** | [**IntegrationRequest**](IntegrationRequest.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdInventoryGet

> []Inventory ProjectProjectIdInventoryGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get inventory

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
	sort := "sort_example" // string | sorting name
	order := "order_example" // string | ordering manner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdInventoryGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdInventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdInventoryGet`: []Inventory
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdInventoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Inventory**](Inventory.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdInventoryInventoryIdDelete

> ProjectProjectIdInventoryInventoryIdDelete(ctx, projectId, inventoryId).Execute()

Removes inventory

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
	inventoryId := int32(56) // int32 | inventory ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdInventoryInventoryIdDelete(context.Background(), projectId, inventoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdInventoryInventoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**inventoryId** | **int32** | inventory ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryInventoryIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdInventoryInventoryIdGet

> Inventory ProjectProjectIdInventoryInventoryIdGet(ctx, projectId, inventoryId).Execute()

Get inventory

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
	inventoryId := int32(56) // int32 | inventory ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdInventoryInventoryIdGet(context.Background(), projectId, inventoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdInventoryInventoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdInventoryInventoryIdGet`: Inventory
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdInventoryInventoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**inventoryId** | **int32** | inventory ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryInventoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Inventory**](Inventory.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdInventoryInventoryIdPut

> ProjectProjectIdInventoryInventoryIdPut(ctx, projectId, inventoryId).Inventory(inventory).Execute()

Updates inventory

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
	inventoryId := int32(56) // int32 | inventory ID
	inventory := *openapiclient.NewInventoryRequest() // InventoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdInventoryInventoryIdPut(context.Background(), projectId, inventoryId).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdInventoryInventoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**inventoryId** | **int32** | inventory ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryInventoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inventory** | [**InventoryRequest**](InventoryRequest.md) |  | 

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


## ProjectProjectIdInventoryPost

> Inventory ProjectProjectIdInventoryPost(ctx, projectId).Inventory(inventory).Execute()

create inventory

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
	inventory := *openapiclient.NewInventoryRequest() // InventoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdInventoryPost(context.Background(), projectId).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdInventoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdInventoryPost`: Inventory
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdInventoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdInventoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventory** | [**InventoryRequest**](InventoryRequest.md) |  | 

### Return type

[**Inventory**](Inventory.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdKeysGet

> []AccessKey ProjectProjectIdKeysGet(ctx, projectId).Sort(sort).Order(order).KeyType(keyType).Execute()

Get access keys linked to project

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
	sort := "type" // string | sorting name
	order := "asc" // string | ordering manner
	keyType := "none" // string | Filter by key type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdKeysGet(context.Background(), projectId).Sort(sort).Order(order).KeyType(keyType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdKeysGet`: []AccessKey
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 
 **keyType** | **string** | Filter by key type | 

### Return type

[**[]AccessKey**](AccessKey.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdKeysKeyIdDelete

> ProjectProjectIdKeysKeyIdDelete(ctx, projectId, keyId).Execute()

Removes access key

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
	keyId := int32(56) // int32 | key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdKeysKeyIdDelete(context.Background(), projectId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdKeysKeyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**keyId** | **int32** | key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysKeyIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdKeysKeyIdPut

> ProjectProjectIdKeysKeyIdPut(ctx, projectId, keyId).AccessKey(accessKey).Execute()

Updates access key

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
	keyId := int32(56) // int32 | key ID
	accessKey := *openapiclient.NewAccessKeyRequest() // AccessKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdKeysKeyIdPut(context.Background(), projectId, keyId).AccessKey(accessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdKeysKeyIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**keyId** | **int32** | key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysKeyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessKey** | [**AccessKeyRequest**](AccessKeyRequest.md) |  | 

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


## ProjectProjectIdKeysPost

> AccessKey ProjectProjectIdKeysPost(ctx, projectId).AccessKey(accessKey).Execute()

Add access key

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
	accessKey := *openapiclient.NewAccessKeyRequest() // AccessKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdKeysPost(context.Background(), projectId).AccessKey(accessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdKeysPost`: AccessKey
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdKeysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessKey** | [**AccessKeyRequest**](AccessKeyRequest.md) |  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdPut

> ProjectProjectIdPut(ctx, projectId).Project(project).Execute()

Update project

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
	project := *openapiclient.NewProjectProjectIdPutRequest() // ProjectProjectIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdPut(context.Background(), projectId).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ProjectProjectIdPutRequest**](ProjectProjectIdPutRequest.md) |  | 

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


## ProjectProjectIdRepositoriesGet

> []Repository ProjectProjectIdRepositoriesGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get repositories

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
	sort := "sort_example" // string | sorting name
	order := "order_example" // string | ordering manner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdRepositoriesGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdRepositoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdRepositoriesGet`: []Repository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdRepositoriesPost

> Repository ProjectProjectIdRepositoriesPost(ctx, projectId).Repository(repository).Execute()

Add repository

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
	repository := *openapiclient.NewRepositoryRequest() // RepositoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdRepositoriesPost(context.Background(), projectId).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdRepositoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdRepositoriesPost`: Repository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdRepositoriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**RepositoryRequest**](RepositoryRequest.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdRepositoriesRepositoryIdDelete

> ProjectProjectIdRepositoriesRepositoryIdDelete(ctx, projectId, repositoryId).Execute()

Removes repository

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
	repositoryId := int32(56) // int32 | repository ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdDelete(context.Background(), projectId, repositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**repositoryId** | **int32** | repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesRepositoryIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdRepositoriesRepositoryIdGet

> Repository ProjectProjectIdRepositoriesRepositoryIdGet(ctx, projectId, repositoryId).Execute()

Get repository

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
	repositoryId := int32(56) // int32 | repository ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdGet(context.Background(), projectId, repositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdRepositoriesRepositoryIdGet`: Repository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**repositoryId** | **int32** | repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesRepositoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdRepositoriesRepositoryIdPut

> ProjectProjectIdRepositoriesRepositoryIdPut(ctx, projectId, repositoryId).Repository(repository).Execute()

Updates repository

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
	repositoryId := int32(56) // int32 | repository ID
	repository := *openapiclient.NewRepositoryRequest() // RepositoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdPut(context.Background(), projectId, repositoryId).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdRepositoriesRepositoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**repositoryId** | **int32** | repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRepositoriesRepositoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repository** | [**RepositoryRequest**](RepositoryRequest.md) |  | 

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


## ProjectProjectIdRoleGet

> ProjectProjectIdRoleGet200Response ProjectProjectIdRoleGet(ctx, projectId).Execute()

Fetch permissions of the current user for project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdRoleGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdRoleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdRoleGet`: ProjectProjectIdRoleGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdRoleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdRoleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectIdRoleGet200Response**](ProjectProjectIdRoleGet200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksGet

> []Task ProjectProjectIdTasksGet(ctx, projectId).Execute()

Get Tasks related to current project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTasksGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTasksGet`: []Task
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Task**](Task.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksLastGet

> []Task ProjectProjectIdTasksLastGet(ctx, projectId).Execute()

Get last 200 Tasks related to current project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTasksLastGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksLastGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTasksLastGet`: []Task
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTasksLastGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksLastGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Task**](Task.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksPost

> Task ProjectProjectIdTasksPost(ctx, projectId).Task(task).Execute()

Starts a job

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
	task := *openapiclient.NewProjectProjectIdTasksPostRequest() // ProjectProjectIdTasksPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTasksPost(context.Background(), projectId).Task(task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTasksPost`: Task
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**ProjectProjectIdTasksPostRequest**](ProjectProjectIdTasksPostRequest.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksTaskIdDelete

> ProjectProjectIdTasksTaskIdDelete(ctx, projectId, taskId).Execute()

Deletes task (including output)

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
	taskId := int32(56) // int32 | task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdTasksTaskIdDelete(context.Background(), projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksTaskIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdTasksTaskIdGet

> Task ProjectProjectIdTasksTaskIdGet(ctx, projectId, taskId).Execute()

Get a single task

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
	taskId := int32(56) // int32 | task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTasksTaskIdGet(context.Background(), projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTasksTaskIdGet`: Task
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksTaskIdOutputGet

> []TaskOutput ProjectProjectIdTasksTaskIdOutputGet(ctx, projectId, taskId).Execute()

Get task output

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
	taskId := int32(56) // int32 | task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTasksTaskIdOutputGet(context.Background(), projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksTaskIdOutputGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTasksTaskIdOutputGet`: []TaskOutput
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTasksTaskIdOutputGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdOutputGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TaskOutput**](TaskOutput.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTasksTaskIdStopPost

> ProjectProjectIdTasksTaskIdStopPost(ctx, projectId, taskId).Execute()

Stop a job

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
	taskId := int32(56) // int32 | task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdTasksTaskIdStopPost(context.Background(), projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTasksTaskIdStopPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**taskId** | **int32** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTasksTaskIdStopPostRequest struct via the builder pattern


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


## ProjectProjectIdTemplatesGet

> []Template ProjectProjectIdTemplatesGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get template

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
	sort := "sort_example" // string | sorting name
	order := "order_example" // string | ordering manner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTemplatesGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTemplatesGet`: []Template
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]Template**](Template.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesPost

> Template ProjectProjectIdTemplatesPost(ctx, projectId).Template(template).Execute()

create template

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
	template := *openapiclient.NewTemplateRequest() // TemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTemplatesPost(context.Background(), projectId).Template(template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTemplatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTemplatesPost`: Template
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | [**TemplateRequest**](TemplateRequest.md) |  | 

### Return type

[**Template**](Template.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesTemplateIdDelete

> ProjectProjectIdTemplatesTemplateIdDelete(ctx, projectId, templateId).Execute()

Removes template

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
	templateId := int32(56) // int32 | template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdTemplatesTemplateIdDelete(context.Background(), projectId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTemplatesTemplateIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**templateId** | **int32** | template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesTemplateIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdTemplatesTemplateIdGet

> Template ProjectProjectIdTemplatesTemplateIdGet(ctx, projectId, templateId).Execute()

Get template

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
	templateId := int32(56) // int32 | template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdTemplatesTemplateIdGet(context.Background(), projectId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTemplatesTemplateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdTemplatesTemplateIdGet`: Template
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdTemplatesTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**templateId** | **int32** | template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Template**](Template.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdTemplatesTemplateIdPut

> ProjectProjectIdTemplatesTemplateIdPut(ctx, projectId, templateId).Template(template).Execute()

Updates template

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
	templateId := int32(56) // int32 | template ID
	template := *openapiclient.NewTemplateRequest() // TemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdTemplatesTemplateIdPut(context.Background(), projectId, templateId).Template(template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdTemplatesTemplateIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**templateId** | **int32** | template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdTemplatesTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **template** | [**TemplateRequest**](TemplateRequest.md) |  | 

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


## ProjectProjectIdUsersGet

> []ProjectUser ProjectProjectIdUsersGet(ctx, projectId).Sort(sort).Order(order).Execute()

Get users linked to project

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
	sort := "email" // string | sorting name
	order := "desc" // string | ordering manner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdUsersGet(context.Background(), projectId).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdUsersGet`: []ProjectUser
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | sorting name | 
 **order** | **string** | ordering manner | 

### Return type

[**[]ProjectUser**](ProjectUser.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdUsersPost

> ProjectProjectIdUsersPost(ctx, projectId).User(user).Execute()

Link user to project

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
	user := *openapiclient.NewProjectProjectIdUsersPostRequest() // ProjectProjectIdUsersPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdUsersPost(context.Background(), projectId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**ProjectProjectIdUsersPostRequest**](ProjectProjectIdUsersPostRequest.md) |  | 

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


## ProjectProjectIdUsersUserIdDelete

> ProjectProjectIdUsersUserIdDelete(ctx, projectId, userId).Execute()

Removes user from project

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
	userId := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdUsersUserIdDelete(context.Background(), projectId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdUsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersUserIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdUsersUserIdPut

> ProjectProjectIdUsersUserIdPut(ctx, projectId, userId).ProjectUser(projectUser).Execute()

Update user role

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
	userId := int32(56) // int32 | User ID
	projectUser := *openapiclient.NewProjectProjectIdUsersUserIdPutRequest() // ProjectProjectIdUsersUserIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdUsersUserIdPut(context.Background(), projectId, userId).ProjectUser(projectUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdUsersUserIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdUsersUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectUser** | [**ProjectProjectIdUsersUserIdPutRequest**](ProjectProjectIdUsersUserIdPutRequest.md) |  | 

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


## ProjectProjectIdViewsGet

> []View ProjectProjectIdViewsGet(ctx, projectId).Execute()

Get view

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdViewsGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdViewsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdViewsGet`: []View
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdViewsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]View**](View.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdViewsPost

> View ProjectProjectIdViewsPost(ctx, projectId).View(view).Execute()

create view

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
	view := *openapiclient.NewViewRequest() // ViewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdViewsPost(context.Background(), projectId).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdViewsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdViewsPost`: View
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdViewsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | [**ViewRequest**](ViewRequest.md) |  | 

### Return type

[**View**](View.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdViewsViewIdDelete

> ProjectProjectIdViewsViewIdDelete(ctx, projectId, viewId).Execute()

Removes view

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
	viewId := int32(56) // int32 | view ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdViewsViewIdDelete(context.Background(), projectId, viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdViewsViewIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**viewId** | **int32** | view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsViewIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdViewsViewIdGet

> View ProjectProjectIdViewsViewIdGet(ctx, projectId, viewId).Execute()

Get view

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
	viewId := int32(56) // int32 | view ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ProjectProjectIdViewsViewIdGet(context.Background(), projectId, viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdViewsViewIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdViewsViewIdGet`: View
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectProjectIdViewsViewIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**viewId** | **int32** | view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsViewIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**View**](View.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdViewsViewIdPut

> ProjectProjectIdViewsViewIdPut(ctx, projectId, viewId).View(view).Execute()

Updates view

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
	viewId := int32(56) // int32 | view ID
	view := *openapiclient.NewViewRequest() // ViewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ProjectProjectIdViewsViewIdPut(context.Background(), projectId, viewId).View(view).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectProjectIdViewsViewIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**viewId** | **int32** | view ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdViewsViewIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | [**ViewRequest**](ViewRequest.md) |  | 

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

