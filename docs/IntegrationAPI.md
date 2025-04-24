# \IntegrationAPI

All URIs are relative to *http://localhost:3000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectProjectIdIntegrationsIntegrationIdMatchersGet**](IntegrationAPI.md#ProjectProjectIdIntegrationsIntegrationIdMatchersGet) | **Get** /project/{project_id}/integrations/{integration_id}/matchers | Get Integration Matcher linked to integration extractor
[**ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDelete**](IntegrationAPI.md#ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDelete) | **Delete** /project/{project_id}/integrations/{integration_id}/matchers/{matcher_id} | Removes integration matcher
[**ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPut**](IntegrationAPI.md#ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPut) | **Put** /project/{project_id}/integrations/{integration_id}/matchers/{matcher_id} | Updates Integration Matcher
[**ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDelete**](IntegrationAPI.md#ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDelete) | **Delete** /project/{project_id}/integrations/{integration_id}/values/{extractvalue_id} | Removes integration extract value
[**ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPut**](IntegrationAPI.md#ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPut) | **Put** /project/{project_id}/integrations/{integration_id}/values/{extractvalue_id} | Updates Integration ExtractValue
[**ProjectProjectIdIntegrationsIntegrationIdValuesGet**](IntegrationAPI.md#ProjectProjectIdIntegrationsIntegrationIdValuesGet) | **Get** /project/{project_id}/integrations/{integration_id}/values | Get Integration Extracted Values linked to integration extractor



## ProjectProjectIdIntegrationsIntegrationIdMatchersGet

> []IntegrationMatcher ProjectProjectIdIntegrationsIntegrationIdMatchersGet(ctx, projectId, integrationId).Execute()

Get Integration Matcher linked to integration extractor

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
	resp, r, err := apiClient.IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersGet(context.Background(), projectId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdIntegrationsIntegrationIdMatchersGet`: []IntegrationMatcher
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**integrationId** | **int32** | integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdMatchersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IntegrationMatcher**](IntegrationMatcher.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDelete

> ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDelete(ctx, projectId, integrationId, matcherId).Execute()

Removes integration matcher

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
	matcherId := int32(56) // int32 | matcher ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDelete(context.Background(), projectId, integrationId, matcherId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDelete``: %v\n", err)
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
**matcherId** | **int32** | matcher ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPut

> ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPut(ctx, projectId, integrationId, matcherId).IntegrationMatcher(integrationMatcher).Execute()

Updates Integration Matcher

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
	matcherId := int32(56) // int32 | matcher ID
	integrationMatcher := *openapiclient.NewIntegrationMatcherRequest() // IntegrationMatcherRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPut(context.Background(), projectId, integrationId, matcherId).IntegrationMatcher(integrationMatcher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPut``: %v\n", err)
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
**matcherId** | **int32** | matcher ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdMatchersMatcherIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **integrationMatcher** | [**IntegrationMatcherRequest**](IntegrationMatcherRequest.md) |  | 

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


## ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDelete

> ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDelete(ctx, projectId, integrationId, extractvalueId).Execute()

Removes integration extract value

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
	extractvalueId := int32(56) // int32 | extractValue ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDelete(context.Background(), projectId, integrationId, extractvalueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDelete``: %v\n", err)
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
**extractvalueId** | **int32** | extractValue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdDeleteRequest struct via the builder pattern


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


## ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPut

> ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPut(ctx, projectId, integrationId, extractvalueId).IntegrationExtractValue(integrationExtractValue).Execute()

Updates Integration ExtractValue

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
	extractvalueId := int32(56) // int32 | extractValue ID
	integrationExtractValue := *openapiclient.NewIntegrationExtractValueRequest() // IntegrationExtractValueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPut(context.Background(), projectId, integrationId, extractvalueId).IntegrationExtractValue(integrationExtractValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPut``: %v\n", err)
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
**extractvalueId** | **int32** | extractValue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **integrationExtractValue** | [**IntegrationExtractValueRequest**](IntegrationExtractValueRequest.md) |  | 

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


## ProjectProjectIdIntegrationsIntegrationIdValuesGet

> []IntegrationExtractValue ProjectProjectIdIntegrationsIntegrationIdValuesGet(ctx, projectId, integrationId).Execute()

Get Integration Extracted Values linked to integration extractor

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
	resp, r, err := apiClient.IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesGet(context.Background(), projectId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectProjectIdIntegrationsIntegrationIdValuesGet`: []IntegrationExtractValue
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ProjectProjectIdIntegrationsIntegrationIdValuesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | Project ID | 
**integrationId** | **int32** | integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdIntegrationsIntegrationIdValuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IntegrationExtractValue**](IntegrationExtractValue.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

