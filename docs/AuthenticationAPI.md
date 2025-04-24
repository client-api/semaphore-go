# \AuthenticationAPI

All URIs are relative to *http://localhost:3000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthLoginGet**](AuthenticationAPI.md#AuthLoginGet) | **Get** /auth/login | Fetches login metadata
[**AuthLoginPost**](AuthenticationAPI.md#AuthLoginPost) | **Post** /auth/login | Performs Login
[**AuthLogoutPost**](AuthenticationAPI.md#AuthLogoutPost) | **Post** /auth/logout | Destroys current session
[**AuthOidcProviderIdLoginGet**](AuthenticationAPI.md#AuthOidcProviderIdLoginGet) | **Get** /auth/oidc/{provider_id}/login | Begin OIDC authentication flow and redirect to OIDC provider
[**AuthOidcProviderIdRedirectGet**](AuthenticationAPI.md#AuthOidcProviderIdRedirectGet) | **Get** /auth/oidc/{provider_id}/redirect | Finish OIDC authentication flow, upon succes you will be logged in
[**UserTokensApiTokenIdDelete**](AuthenticationAPI.md#UserTokensApiTokenIdDelete) | **Delete** /user/tokens/{api_token_id} | Expires API token
[**UserTokensGet**](AuthenticationAPI.md#UserTokensGet) | **Get** /user/tokens | Fetch API tokens for user
[**UserTokensPost**](AuthenticationAPI.md#UserTokensPost) | **Post** /user/tokens | Create an API token



## AuthLoginGet

> LoginMetadata AuthLoginGet(ctx).Execute()

Fetches login metadata



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthLoginGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthLoginGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLoginGet`: LoginMetadata
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthLoginGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginGetRequest struct via the builder pattern


### Return type

[**LoginMetadata**](LoginMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLoginPost

> AuthLoginPost(ctx).LoginBody(loginBody).Execute()

Performs Login



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
	loginBody := *openapiclient.NewLogin() // Login | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AuthLoginPost(context.Background()).LoginBody(loginBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginBody** | [**Login**](Login.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLogoutPost

> AuthLogoutPost(ctx).Execute()

Destroys current session

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AuthLogoutPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLogoutPostRequest struct via the builder pattern


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


## AuthOidcProviderIdLoginGet

> AuthOidcProviderIdLoginGet(ctx, providerId).Execute()

Begin OIDC authentication flow and redirect to OIDC provider



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
	providerId := "mysso" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AuthOidcProviderIdLoginGet(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthOidcProviderIdLoginGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthOidcProviderIdLoginGetRequest struct via the builder pattern


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


## AuthOidcProviderIdRedirectGet

> AuthOidcProviderIdRedirectGet(ctx, providerId).Execute()

Finish OIDC authentication flow, upon succes you will be logged in



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
	providerId := "mysso" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.AuthOidcProviderIdRedirectGet(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthOidcProviderIdRedirectGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthOidcProviderIdRedirectGetRequest struct via the builder pattern


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


## UserTokensApiTokenIdDelete

> UserTokensApiTokenIdDelete(ctx, apiTokenId).Execute()

Expires API token

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
	apiTokenId := "kwofd61g93-yuqvex8efmhjkgnbxlo8mp1tin6spyhu=" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.UserTokensApiTokenIdDelete(context.Background(), apiTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UserTokensApiTokenIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokensApiTokenIdDeleteRequest struct via the builder pattern


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


## UserTokensGet

> []APIToken UserTokensGet(ctx).Execute()

Fetch API tokens for user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UserTokensGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UserTokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokensGet`: []APIToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UserTokensGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokensGetRequest struct via the builder pattern


### Return type

[**[]APIToken**](APIToken.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTokensPost

> APIToken UserTokensPost(ctx).Execute()

Create an API token

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UserTokensPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UserTokensPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTokensPost`: APIToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UserTokensPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserTokensPostRequest struct via the builder pattern


### Return type

[**APIToken**](APIToken.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

