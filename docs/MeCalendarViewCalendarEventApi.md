# \MeCalendarViewCalendarEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarViewCalendarCreateCalendarView**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarCreateCalendarView) | **Post** /me/calendarView({event-id})/calendar/calendarView | Create new navigation property to calendarView for me
[**MeCalendarViewCalendarCreateEvents**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarCreateEvents) | **Post** /me/calendarView({event-id})/calendar/events | Create new navigation property to events for me
[**MeCalendarViewCalendarGetCalendarView**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarGetCalendarView) | **Get** /me/calendarView({event-id})/calendar/calendarView({event-id1}) | Get calendarView from me
[**MeCalendarViewCalendarGetEvents**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarGetEvents) | **Get** /me/calendarView({event-id})/calendar/events({event-id1}) | Get events from me
[**MeCalendarViewCalendarListCalendarView**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarListCalendarView) | **Get** /me/calendarView({event-id})/calendar/calendarView | Get calendarView from me
[**MeCalendarViewCalendarListEvents**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarListEvents) | **Get** /me/calendarView({event-id})/calendar/events | Get events from me
[**MeCalendarViewCalendarUpdateCalendarView**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarUpdateCalendarView) | **Patch** /me/calendarView({event-id})/calendar/calendarView({event-id1}) | Update the navigation property calendarView in me
[**MeCalendarViewCalendarUpdateEvents**](MeCalendarViewCalendarEventApi.md#MeCalendarViewCalendarUpdateEvents) | **Patch** /me/calendarView({event-id})/calendar/events({event-id1}) | Update the navigation property events in me



## MeCalendarViewCalendarCreateCalendarView

> MicrosoftGraphEvent MeCalendarViewCalendarCreateCalendarView(ctx, eventId, microsoftGraphEvent)
Create new navigation property to calendarView for me

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


## MeCalendarViewCalendarCreateEvents

> MicrosoftGraphEvent MeCalendarViewCalendarCreateEvents(ctx, eventId, microsoftGraphEvent)
Create new navigation property to events for me

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


## MeCalendarViewCalendarGetCalendarView

> MicrosoftGraphEvent MeCalendarViewCalendarGetCalendarView(ctx, eventId, eventId1, optional)
Get calendarView from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarViewCalendarGetCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarViewCalendarGetCalendarViewOpts struct


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


## MeCalendarViewCalendarGetEvents

> MicrosoftGraphEvent MeCalendarViewCalendarGetEvents(ctx, eventId, eventId1, optional)
Get events from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarViewCalendarGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarViewCalendarGetEventsOpts struct


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


## MeCalendarViewCalendarListCalendarView

> CollectionOfEvent MeCalendarViewCalendarListCalendarView(ctx, eventId, optional)
Get calendarView from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarViewCalendarListCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarViewCalendarListCalendarViewOpts struct


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


## MeCalendarViewCalendarListEvents

> CollectionOfEvent MeCalendarViewCalendarListEvents(ctx, eventId, optional)
Get events from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarViewCalendarListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarViewCalendarListEventsOpts struct


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


## MeCalendarViewCalendarUpdateCalendarView

> MeCalendarViewCalendarUpdateCalendarView(ctx, eventId, eventId1, microsoftGraphEvent)
Update the navigation property calendarView in me

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


## MeCalendarViewCalendarUpdateEvents

> MeCalendarViewCalendarUpdateEvents(ctx, eventId, eventId1, microsoftGraphEvent)
Update the navigation property events in me

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

