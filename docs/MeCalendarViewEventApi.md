# \MeCalendarViewEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarViewCreateInstances**](MeCalendarViewEventApi.md#MeCalendarViewCreateInstances) | **Post** /me/calendarView({event-id})/instances | Create new navigation property to instances for me
[**MeCalendarViewGetInstances**](MeCalendarViewEventApi.md#MeCalendarViewGetInstances) | **Get** /me/calendarView({event-id})/instances({event-id1}) | Get instances from me
[**MeCalendarViewListInstances**](MeCalendarViewEventApi.md#MeCalendarViewListInstances) | **Get** /me/calendarView({event-id})/instances | Get instances from me
[**MeCalendarViewUpdateInstances**](MeCalendarViewEventApi.md#MeCalendarViewUpdateInstances) | **Patch** /me/calendarView({event-id})/instances({event-id1}) | Update the navigation property instances in me



## MeCalendarViewCreateInstances

> MicrosoftGraphEvent MeCalendarViewCreateInstances(ctx, eventId, microsoftGraphEvent)
Create new navigation property to instances for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**microsoftGraphEvent** | [**MicrosoftGraphEvent**](MicrosoftGraphEvent.md)| New navigation property | 

### Return type

[**MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeCalendarViewGetInstances

> MicrosoftGraphEvent MeCalendarViewGetInstances(ctx, eventId, eventId1, optional)
Get instances from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarViewGetInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarViewGetInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeCalendarViewListInstances

> CollectionOfEvent MeCalendarViewListInstances(ctx, eventId, optional)
Get instances from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarViewListInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarViewListInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfEvent**](Collection of event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeCalendarViewUpdateInstances

> MeCalendarViewUpdateInstances(ctx, eventId, eventId1, microsoftGraphEvent)
Update the navigation property instances in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
**microsoftGraphEvent** | [**MicrosoftGraphEvent**](MicrosoftGraphEvent.md)| New navigation property values | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
