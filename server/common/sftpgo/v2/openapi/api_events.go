/*
SFTPGo

SFTPGo allows to securely share your files over SFTP, HTTP and optionally FTP/S and WebDAV as well. Several storage backends are supported and they are configurable per user, so you can serve a local directory for a user and an S3 bucket (or part of it) for another one. SFTPGo also supports virtual folders, a virtual folder can use any of the supported storage backends. So you can have, for example, an S3 user that exposes a GCS bucket (or part of it) on a specified path and an encrypted local filesystem on another one. Virtual folders can be private or shared among multiple users, for shared virtual folders you can define different quota limits for each user. SFTPGo allows to create HTTP/S links to externally share files and folders securely, by setting limits to the number of downloads/uploads, protecting the share with a password, limiting access by source IP address, setting an automatic expiration date. 

API version: 2.2.2-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// EventsApiService EventsApi service
type EventsApiService service

type EventsApiApiGetFsEventsRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	startTimestamp *int64
	endTimestamp *int64
	actions *[]FsEventAction
	username *string
	ip *string
	sshCmd *string
	fsProvider *FsProviders
	bucket *string
	endpoint *string
	protocols *[]EventProtocols
	statuses *[]FsEventStatus
	instanceIds *[]string
	excludeIds *[]string
	limit *int32
	order *string
}

// the event timestamp, unix timestamp in nanoseconds, must be greater than or equal to the specified one. 0 or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) StartTimestamp(startTimestamp int64) EventsApiApiGetFsEventsRequest {
	r.startTimestamp = &startTimestamp
	return r
}
// the event timestamp, unix timestamp in nanoseconds, must be less than or equal to the specified one. 0 or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) EndTimestamp(endTimestamp int64) EventsApiApiGetFsEventsRequest {
	r.endTimestamp = &endTimestamp
	return r
}
// the event action must be included among those specified. Empty or missing means omit this filter. Actions must be specified comma separated
func (r EventsApiApiGetFsEventsRequest) Actions(actions []FsEventAction) EventsApiApiGetFsEventsRequest {
	r.actions = &actions
	return r
}
// the event username must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) Username(username string) EventsApiApiGetFsEventsRequest {
	r.username = &username
	return r
}
// the event IP must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) Ip(ip string) EventsApiApiGetFsEventsRequest {
	r.ip = &ip
	return r
}
// the event SSH command must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) SshCmd(sshCmd string) EventsApiApiGetFsEventsRequest {
	r.sshCmd = &sshCmd
	return r
}
// the event filesystem provider must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) FsProvider(fsProvider FsProviders) EventsApiApiGetFsEventsRequest {
	r.fsProvider = &fsProvider
	return r
}
// the bucket must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) Bucket(bucket string) EventsApiApiGetFsEventsRequest {
	r.bucket = &bucket
	return r
}
// the endpoint must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetFsEventsRequest) Endpoint(endpoint string) EventsApiApiGetFsEventsRequest {
	r.endpoint = &endpoint
	return r
}
// the event protocol must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetFsEventsRequest) Protocols(protocols []EventProtocols) EventsApiApiGetFsEventsRequest {
	r.protocols = &protocols
	return r
}
// the event status must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetFsEventsRequest) Statuses(statuses []FsEventStatus) EventsApiApiGetFsEventsRequest {
	r.statuses = &statuses
	return r
}
// the event instance id must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetFsEventsRequest) InstanceIds(instanceIds []string) EventsApiApiGetFsEventsRequest {
	r.instanceIds = &instanceIds
	return r
}
// the event id must not be included among those specified. This is useful for cursor based pagination. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetFsEventsRequest) ExcludeIds(excludeIds []string) EventsApiApiGetFsEventsRequest {
	r.excludeIds = &excludeIds
	return r
}
// The maximum number of items to return. Max value is 500, default is 100
func (r EventsApiApiGetFsEventsRequest) Limit(limit int32) EventsApiApiGetFsEventsRequest {
	r.limit = &limit
	return r
}
// Ordering events by timestamp. Default DESC
func (r EventsApiApiGetFsEventsRequest) Order(order string) EventsApiApiGetFsEventsRequest {
	r.order = &order
	return r
}

func (r EventsApiApiGetFsEventsRequest) Execute() ([]FsEvent, *http.Response, error) {
	return r.ApiService.GetFsEventsExecute(r)
}

/*
GetFsEvents Get filesystem events

Returns an array with one or more filesystem events applying the specified filters. This API is only available if you configure an "eventsearcher" plugin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return EventsApiApiGetFsEventsRequest
*/
func (a *EventsApiService) GetFsEvents(ctx context.Context) EventsApiApiGetFsEventsRequest {
	return EventsApiApiGetFsEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []FsEvent
func (a *EventsApiService) GetFsEventsExecute(r EventsApiApiGetFsEventsRequest) ([]FsEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []FsEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetFsEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/fs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTimestamp != nil {
		localVarQueryParams.Add("start_timestamp", parameterToString(*r.startTimestamp, ""))
	}
	if r.endTimestamp != nil {
		localVarQueryParams.Add("end_timestamp", parameterToString(*r.endTimestamp, ""))
	}
	if r.actions != nil {
		localVarQueryParams.Add("actions", parameterToString(*r.actions, "csv"))
	}
	if r.username != nil {
		localVarQueryParams.Add("username", parameterToString(*r.username, ""))
	}
	if r.ip != nil {
		localVarQueryParams.Add("ip", parameterToString(*r.ip, ""))
	}
	if r.sshCmd != nil {
		localVarQueryParams.Add("ssh_cmd", parameterToString(*r.sshCmd, ""))
	}
	if r.fsProvider != nil {
		localVarQueryParams.Add("fs_provider", parameterToString(*r.fsProvider, ""))
	}
	if r.bucket != nil {
		localVarQueryParams.Add("bucket", parameterToString(*r.bucket, ""))
	}
	if r.endpoint != nil {
		localVarQueryParams.Add("endpoint", parameterToString(*r.endpoint, ""))
	}
	if r.protocols != nil {
		localVarQueryParams.Add("protocols", parameterToString(*r.protocols, "csv"))
	}
	if r.statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.statuses, "csv"))
	}
	if r.instanceIds != nil {
		localVarQueryParams.Add("instance_ids", parameterToString(*r.instanceIds, "csv"))
	}
	if r.excludeIds != nil {
		localVarQueryParams.Add("exclude_ids", parameterToString(*r.excludeIds, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-SFTPGO-API-KEY"] = key
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type EventsApiApiGetProviderEventsRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	startTimestamp *int64
	endTimestamp *int64
	actions *[]ProviderEventAction
	username *string
	ip *string
	objectName *string
	objectTypes *[]ProviderEventObjectType
	instanceIds *[]string
	excludeIds *[]string
	limit *int32
	order *string
}

// the event timestamp, unix timestamp in nanoseconds, must be greater than or equal to the specified one. 0 or missing means omit this filter
func (r EventsApiApiGetProviderEventsRequest) StartTimestamp(startTimestamp int64) EventsApiApiGetProviderEventsRequest {
	r.startTimestamp = &startTimestamp
	return r
}
// the event timestamp, unix timestamp in nanoseconds, must be less than or equal to the specified one. 0 or missing means omit this filter
func (r EventsApiApiGetProviderEventsRequest) EndTimestamp(endTimestamp int64) EventsApiApiGetProviderEventsRequest {
	r.endTimestamp = &endTimestamp
	return r
}
// the event action must be included among those specified. Empty or missing means omit this filter. Actions must be specified comma separated
func (r EventsApiApiGetProviderEventsRequest) Actions(actions []ProviderEventAction) EventsApiApiGetProviderEventsRequest {
	r.actions = &actions
	return r
}
// the event username must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetProviderEventsRequest) Username(username string) EventsApiApiGetProviderEventsRequest {
	r.username = &username
	return r
}
// the event IP must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetProviderEventsRequest) Ip(ip string) EventsApiApiGetProviderEventsRequest {
	r.ip = &ip
	return r
}
// the event object name must be the same as the one specified. Empty or missing means omit this filter
func (r EventsApiApiGetProviderEventsRequest) ObjectName(objectName string) EventsApiApiGetProviderEventsRequest {
	r.objectName = &objectName
	return r
}
// the event object type must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetProviderEventsRequest) ObjectTypes(objectTypes []ProviderEventObjectType) EventsApiApiGetProviderEventsRequest {
	r.objectTypes = &objectTypes
	return r
}
// the event instance id must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetProviderEventsRequest) InstanceIds(instanceIds []string) EventsApiApiGetProviderEventsRequest {
	r.instanceIds = &instanceIds
	return r
}
// the event id must not be included among those specified. This is useful for cursor based pagination. Empty or missing means omit this filter. Values must be specified comma separated
func (r EventsApiApiGetProviderEventsRequest) ExcludeIds(excludeIds []string) EventsApiApiGetProviderEventsRequest {
	r.excludeIds = &excludeIds
	return r
}
// The maximum number of items to return. Max value is 500, default is 100
func (r EventsApiApiGetProviderEventsRequest) Limit(limit int32) EventsApiApiGetProviderEventsRequest {
	r.limit = &limit
	return r
}
// Ordering events by timestamp. Default DESC
func (r EventsApiApiGetProviderEventsRequest) Order(order string) EventsApiApiGetProviderEventsRequest {
	r.order = &order
	return r
}

func (r EventsApiApiGetProviderEventsRequest) Execute() ([]ProviderEvent, *http.Response, error) {
	return r.ApiService.GetProviderEventsExecute(r)
}

/*
GetProviderEvents Get provider events

Returns an array with one or more provider events applying the specified filters. This API is only available if you configure an "eventsearcher" plugin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return EventsApiApiGetProviderEventsRequest
*/
func (a *EventsApiService) GetProviderEvents(ctx context.Context) EventsApiApiGetProviderEventsRequest {
	return EventsApiApiGetProviderEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ProviderEvent
func (a *EventsApiService) GetProviderEventsExecute(r EventsApiApiGetProviderEventsRequest) ([]ProviderEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ProviderEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetProviderEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/provider"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTimestamp != nil {
		localVarQueryParams.Add("start_timestamp", parameterToString(*r.startTimestamp, ""))
	}
	if r.endTimestamp != nil {
		localVarQueryParams.Add("end_timestamp", parameterToString(*r.endTimestamp, ""))
	}
	if r.actions != nil {
		localVarQueryParams.Add("actions", parameterToString(*r.actions, "csv"))
	}
	if r.username != nil {
		localVarQueryParams.Add("username", parameterToString(*r.username, ""))
	}
	if r.ip != nil {
		localVarQueryParams.Add("ip", parameterToString(*r.ip, ""))
	}
	if r.objectName != nil {
		localVarQueryParams.Add("object_name", parameterToString(*r.objectName, ""))
	}
	if r.objectTypes != nil {
		localVarQueryParams.Add("object_types", parameterToString(*r.objectTypes, "csv"))
	}
	if r.instanceIds != nil {
		localVarQueryParams.Add("instance_ids", parameterToString(*r.instanceIds, "csv"))
	}
	if r.excludeIds != nil {
		localVarQueryParams.Add("exclude_ids", parameterToString(*r.excludeIds, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-SFTPGO-API-KEY"] = key
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
