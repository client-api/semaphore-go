/*
Semaphore API

Semaphore API provides endpoints for managing and interacting with the Semaphore UI. This documentation outlines the available operations and data models. 

API version: 2.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package semaphore

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type ScheduleAPI interface {

	/*
	ProjectProjectIdSchedulesPost create schedule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@return ApiProjectProjectIdSchedulesPostRequest
	*/
	ProjectProjectIdSchedulesPost(ctx context.Context, projectId int32) ApiProjectProjectIdSchedulesPostRequest

	// ProjectProjectIdSchedulesPostExecute executes the request
	//  @return Schedule
	ProjectProjectIdSchedulesPostExecute(r ApiProjectProjectIdSchedulesPostRequest) (*Schedule, *http.Response, error)

	/*
	ProjectProjectIdSchedulesScheduleIdDelete Deletes schedule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param scheduleId schedule ID
	@return ApiProjectProjectIdSchedulesScheduleIdDeleteRequest
	*/
	ProjectProjectIdSchedulesScheduleIdDelete(ctx context.Context, projectId int32, scheduleId int32) ApiProjectProjectIdSchedulesScheduleIdDeleteRequest

	// ProjectProjectIdSchedulesScheduleIdDeleteExecute executes the request
	ProjectProjectIdSchedulesScheduleIdDeleteExecute(r ApiProjectProjectIdSchedulesScheduleIdDeleteRequest) (*http.Response, error)

	/*
	ProjectProjectIdSchedulesScheduleIdGet Get schedule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param scheduleId schedule ID
	@return ApiProjectProjectIdSchedulesScheduleIdGetRequest
	*/
	ProjectProjectIdSchedulesScheduleIdGet(ctx context.Context, projectId int32, scheduleId int32) ApiProjectProjectIdSchedulesScheduleIdGetRequest

	// ProjectProjectIdSchedulesScheduleIdGetExecute executes the request
	//  @return Schedule
	ProjectProjectIdSchedulesScheduleIdGetExecute(r ApiProjectProjectIdSchedulesScheduleIdGetRequest) (*Schedule, *http.Response, error)

	/*
	ProjectProjectIdSchedulesScheduleIdPut Updates schedule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param scheduleId schedule ID
	@return ApiProjectProjectIdSchedulesScheduleIdPutRequest
	*/
	ProjectProjectIdSchedulesScheduleIdPut(ctx context.Context, projectId int32, scheduleId int32) ApiProjectProjectIdSchedulesScheduleIdPutRequest

	// ProjectProjectIdSchedulesScheduleIdPutExecute executes the request
	ProjectProjectIdSchedulesScheduleIdPutExecute(r ApiProjectProjectIdSchedulesScheduleIdPutRequest) (*http.Response, error)
}

// ScheduleAPIService ScheduleAPI service
type ScheduleAPIService service

type ApiProjectProjectIdSchedulesPostRequest struct {
	ctx context.Context
	ApiService ScheduleAPI
	projectId int32
	schedule *ScheduleRequest
}

func (r ApiProjectProjectIdSchedulesPostRequest) Schedule(schedule ScheduleRequest) ApiProjectProjectIdSchedulesPostRequest {
	r.schedule = &schedule
	return r
}

func (r ApiProjectProjectIdSchedulesPostRequest) Execute() (*Schedule, *http.Response, error) {
	return r.ApiService.ProjectProjectIdSchedulesPostExecute(r)
}

/*
ProjectProjectIdSchedulesPost create schedule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @return ApiProjectProjectIdSchedulesPostRequest
*/
func (a *ScheduleAPIService) ProjectProjectIdSchedulesPost(ctx context.Context, projectId int32) ApiProjectProjectIdSchedulesPostRequest {
	return ApiProjectProjectIdSchedulesPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Schedule
func (a *ScheduleAPIService) ProjectProjectIdSchedulesPostExecute(r ApiProjectProjectIdSchedulesPostRequest) (*Schedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Schedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleAPIService.ProjectProjectIdSchedulesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/schedules"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.schedule == nil {
		return localVarReturnValue, nil, reportError("schedule is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.schedule
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["cookie"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Cookie"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProjectProjectIdSchedulesScheduleIdDeleteRequest struct {
	ctx context.Context
	ApiService ScheduleAPI
	projectId int32
	scheduleId int32
}

func (r ApiProjectProjectIdSchedulesScheduleIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProjectProjectIdSchedulesScheduleIdDeleteExecute(r)
}

/*
ProjectProjectIdSchedulesScheduleIdDelete Deletes schedule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param scheduleId schedule ID
 @return ApiProjectProjectIdSchedulesScheduleIdDeleteRequest
*/
func (a *ScheduleAPIService) ProjectProjectIdSchedulesScheduleIdDelete(ctx context.Context, projectId int32, scheduleId int32) ApiProjectProjectIdSchedulesScheduleIdDeleteRequest {
	return ApiProjectProjectIdSchedulesScheduleIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		scheduleId: scheduleId,
	}
}

// Execute executes the request
func (a *ScheduleAPIService) ProjectProjectIdSchedulesScheduleIdDeleteExecute(r ApiProjectProjectIdSchedulesScheduleIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleAPIService.ProjectProjectIdSchedulesScheduleIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/schedules/{schedule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schedule_id"+"}", url.PathEscape(parameterValueToString(r.scheduleId, "scheduleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["cookie"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Cookie"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiProjectProjectIdSchedulesScheduleIdGetRequest struct {
	ctx context.Context
	ApiService ScheduleAPI
	projectId int32
	scheduleId int32
}

func (r ApiProjectProjectIdSchedulesScheduleIdGetRequest) Execute() (*Schedule, *http.Response, error) {
	return r.ApiService.ProjectProjectIdSchedulesScheduleIdGetExecute(r)
}

/*
ProjectProjectIdSchedulesScheduleIdGet Get schedule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param scheduleId schedule ID
 @return ApiProjectProjectIdSchedulesScheduleIdGetRequest
*/
func (a *ScheduleAPIService) ProjectProjectIdSchedulesScheduleIdGet(ctx context.Context, projectId int32, scheduleId int32) ApiProjectProjectIdSchedulesScheduleIdGetRequest {
	return ApiProjectProjectIdSchedulesScheduleIdGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		scheduleId: scheduleId,
	}
}

// Execute executes the request
//  @return Schedule
func (a *ScheduleAPIService) ProjectProjectIdSchedulesScheduleIdGetExecute(r ApiProjectProjectIdSchedulesScheduleIdGetRequest) (*Schedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Schedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleAPIService.ProjectProjectIdSchedulesScheduleIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/schedules/{schedule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schedule_id"+"}", url.PathEscape(parameterValueToString(r.scheduleId, "scheduleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["cookie"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Cookie"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProjectProjectIdSchedulesScheduleIdPutRequest struct {
	ctx context.Context
	ApiService ScheduleAPI
	projectId int32
	scheduleId int32
	schedule *ScheduleRequest
}

func (r ApiProjectProjectIdSchedulesScheduleIdPutRequest) Schedule(schedule ScheduleRequest) ApiProjectProjectIdSchedulesScheduleIdPutRequest {
	r.schedule = &schedule
	return r
}

func (r ApiProjectProjectIdSchedulesScheduleIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProjectProjectIdSchedulesScheduleIdPutExecute(r)
}

/*
ProjectProjectIdSchedulesScheduleIdPut Updates schedule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project ID
 @param scheduleId schedule ID
 @return ApiProjectProjectIdSchedulesScheduleIdPutRequest
*/
func (a *ScheduleAPIService) ProjectProjectIdSchedulesScheduleIdPut(ctx context.Context, projectId int32, scheduleId int32) ApiProjectProjectIdSchedulesScheduleIdPutRequest {
	return ApiProjectProjectIdSchedulesScheduleIdPutRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		scheduleId: scheduleId,
	}
}

// Execute executes the request
func (a *ScheduleAPIService) ProjectProjectIdSchedulesScheduleIdPutExecute(r ApiProjectProjectIdSchedulesScheduleIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleAPIService.ProjectProjectIdSchedulesScheduleIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/schedules/{schedule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"schedule_id"+"}", url.PathEscape(parameterValueToString(r.scheduleId, "scheduleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.schedule == nil {
		return nil, reportError("schedule is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.schedule
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["cookie"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Cookie"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
