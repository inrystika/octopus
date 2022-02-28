# \EventsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFsEvents**](EventsApi.md#GetFsEvents) | **Get** /events/fs | Get filesystem events
[**GetProviderEvents**](EventsApi.md#GetProviderEvents) | **Get** /events/provider | Get provider events



## GetFsEvents

> []FsEvent GetFsEvents(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Actions(actions).Username(username).Ip(ip).SshCmd(sshCmd).FsProvider(fsProvider).Bucket(bucket).Endpoint(endpoint).Protocols(protocols).Statuses(statuses).InstanceIds(instanceIds).ExcludeIds(excludeIds).Limit(limit).Order(order).Execute()

Get filesystem events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startTimestamp := int64(789) // int64 | the event timestamp, unix timestamp in nanoseconds, must be greater than or equal to the specified one. 0 or missing means omit this filter (optional) (default to 0)
    endTimestamp := int64(789) // int64 | the event timestamp, unix timestamp in nanoseconds, must be less than or equal to the specified one. 0 or missing means omit this filter (optional) (default to 0)
    actions := []openapiclient.FsEventAction{openapiclient.FsEventAction("download")} // []FsEventAction | the event action must be included among those specified. Empty or missing means omit this filter. Actions must be specified comma separated (optional)
    username := "username_example" // string | the event username must be the same as the one specified. Empty or missing means omit this filter (optional)
    ip := "ip_example" // string | the event IP must be the same as the one specified. Empty or missing means omit this filter (optional)
    sshCmd := "sshCmd_example" // string | the event SSH command must be the same as the one specified. Empty or missing means omit this filter (optional)
    fsProvider := openapiclient.FsProviders(0) // FsProviders | the event filesystem provider must be the same as the one specified. Empty or missing means omit this filter (optional)
    bucket := "bucket_example" // string | the bucket must be the same as the one specified. Empty or missing means omit this filter (optional)
    endpoint := "endpoint_example" // string | the endpoint must be the same as the one specified. Empty or missing means omit this filter (optional)
    protocols := []openapiclient.EventProtocols{openapiclient.EventProtocols("SSH")} // []EventProtocols | the event protocol must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    statuses := []openapiclient.FsEventStatus{openapiclient.FsEventStatus(1)} // []FsEventStatus | the event status must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    instanceIds := []string{"Inner_example"} // []string | the event instance id must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    excludeIds := []string{"Inner_example"} // []string | the event id must not be included among those specified. This is useful for cursor based pagination. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    limit := int32(56) // int32 | The maximum number of items to return. Max value is 500, default is 100 (optional) (default to 100)
    order := "DESC" // string | Ordering events by timestamp. Default DESC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetFsEvents(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Actions(actions).Username(username).Ip(ip).SshCmd(sshCmd).FsProvider(fsProvider).Bucket(bucket).Endpoint(endpoint).Protocols(protocols).Statuses(statuses).InstanceIds(instanceIds).ExcludeIds(excludeIds).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetFsEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFsEvents`: []FsEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetFsEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFsEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | the event timestamp, unix timestamp in nanoseconds, must be greater than or equal to the specified one. 0 or missing means omit this filter | [default to 0]
 **endTimestamp** | **int64** | the event timestamp, unix timestamp in nanoseconds, must be less than or equal to the specified one. 0 or missing means omit this filter | [default to 0]
 **actions** | [**[]FsEventAction**](FsEventAction.md) | the event action must be included among those specified. Empty or missing means omit this filter. Actions must be specified comma separated | 
 **username** | **string** | the event username must be the same as the one specified. Empty or missing means omit this filter | 
 **ip** | **string** | the event IP must be the same as the one specified. Empty or missing means omit this filter | 
 **sshCmd** | **string** | the event SSH command must be the same as the one specified. Empty or missing means omit this filter | 
 **fsProvider** | [**FsProviders**](FsProviders.md) | the event filesystem provider must be the same as the one specified. Empty or missing means omit this filter | 
 **bucket** | **string** | the bucket must be the same as the one specified. Empty or missing means omit this filter | 
 **endpoint** | **string** | the endpoint must be the same as the one specified. Empty or missing means omit this filter | 
 **protocols** | [**[]EventProtocols**](EventProtocols.md) | the event protocol must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated | 
 **statuses** | [**[]FsEventStatus**](FsEventStatus.md) | the event status must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated | 
 **instanceIds** | **[]string** | the event instance id must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated | 
 **excludeIds** | **[]string** | the event id must not be included among those specified. This is useful for cursor based pagination. Empty or missing means omit this filter. Values must be specified comma separated | 
 **limit** | **int32** | The maximum number of items to return. Max value is 500, default is 100 | [default to 100]
 **order** | **string** | Ordering events by timestamp. Default DESC | 

### Return type

[**[]FsEvent**](FsEvent.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderEvents

> []ProviderEvent GetProviderEvents(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Actions(actions).Username(username).Ip(ip).ObjectName(objectName).ObjectTypes(objectTypes).InstanceIds(instanceIds).ExcludeIds(excludeIds).Limit(limit).Order(order).Execute()

Get provider events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startTimestamp := int64(789) // int64 | the event timestamp, unix timestamp in nanoseconds, must be greater than or equal to the specified one. 0 or missing means omit this filter (optional) (default to 0)
    endTimestamp := int64(789) // int64 | the event timestamp, unix timestamp in nanoseconds, must be less than or equal to the specified one. 0 or missing means omit this filter (optional) (default to 0)
    actions := []openapiclient.ProviderEventAction{openapiclient.ProviderEventAction("add")} // []ProviderEventAction | the event action must be included among those specified. Empty or missing means omit this filter. Actions must be specified comma separated (optional)
    username := "username_example" // string | the event username must be the same as the one specified. Empty or missing means omit this filter (optional)
    ip := "ip_example" // string | the event IP must be the same as the one specified. Empty or missing means omit this filter (optional)
    objectName := "objectName_example" // string | the event object name must be the same as the one specified. Empty or missing means omit this filter (optional)
    objectTypes := []openapiclient.ProviderEventObjectType{openapiclient.ProviderEventObjectType("user")} // []ProviderEventObjectType | the event object type must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    instanceIds := []string{"Inner_example"} // []string | the event instance id must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    excludeIds := []string{"Inner_example"} // []string | the event id must not be included among those specified. This is useful for cursor based pagination. Empty or missing means omit this filter. Values must be specified comma separated (optional)
    limit := int32(56) // int32 | The maximum number of items to return. Max value is 500, default is 100 (optional) (default to 100)
    order := "DESC" // string | Ordering events by timestamp. Default DESC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetProviderEvents(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Actions(actions).Username(username).Ip(ip).ObjectName(objectName).ObjectTypes(objectTypes).InstanceIds(instanceIds).ExcludeIds(excludeIds).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetProviderEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviderEvents`: []ProviderEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetProviderEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | the event timestamp, unix timestamp in nanoseconds, must be greater than or equal to the specified one. 0 or missing means omit this filter | [default to 0]
 **endTimestamp** | **int64** | the event timestamp, unix timestamp in nanoseconds, must be less than or equal to the specified one. 0 or missing means omit this filter | [default to 0]
 **actions** | [**[]ProviderEventAction**](ProviderEventAction.md) | the event action must be included among those specified. Empty or missing means omit this filter. Actions must be specified comma separated | 
 **username** | **string** | the event username must be the same as the one specified. Empty or missing means omit this filter | 
 **ip** | **string** | the event IP must be the same as the one specified. Empty or missing means omit this filter | 
 **objectName** | **string** | the event object name must be the same as the one specified. Empty or missing means omit this filter | 
 **objectTypes** | [**[]ProviderEventObjectType**](ProviderEventObjectType.md) | the event object type must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated | 
 **instanceIds** | **[]string** | the event instance id must be included among those specified. Empty or missing means omit this filter. Values must be specified comma separated | 
 **excludeIds** | **[]string** | the event id must not be included among those specified. This is useful for cursor based pagination. Empty or missing means omit this filter. Values must be specified comma separated | 
 **limit** | **int32** | The maximum number of items to return. Max value is 500, default is 100 | [default to 100]
 **order** | **string** | Ordering events by timestamp. Default DESC | 

### Return type

[**[]ProviderEvent**](ProviderEvent.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

