# \MeEventsCalendarEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeEventsCalendarCreateCalendarView**](MeEventsCalendarEventApi.md#MeEventsCalendarCreateCalendarView) | **Post** /me/events({event-id})/calendar/calendarView | Create new navigation property to calendarView for me
[**MeEventsCalendarCreateEvents**](MeEventsCalendarEventApi.md#MeEventsCalendarCreateEvents) | **Post** /me/events({event-id})/calendar/events | Create new navigation property to events for me
[**MeEventsCalendarGetCalendarView**](MeEventsCalendarEventApi.md#MeEventsCalendarGetCalendarView) | **Get** /me/events({event-id})/calendar/calendarView({event-id1}) | Get calendarView from me
[**MeEventsCalendarGetEvents**](MeEventsCalendarEventApi.md#MeEventsCalendarGetEvents) | **Get** /me/events({event-id})/calendar/events({event-id1}) | Get events from me
[**MeEventsCalendarListCalendarView**](MeEventsCalendarEventApi.md#MeEventsCalendarListCalendarView) | **Get** /me/events({event-id})/calendar/calendarView | Get calendarView from me
[**MeEventsCalendarListEvents**](MeEventsCalendarEventApi.md#MeEventsCalendarListEvents) | **Get** /me/events({event-id})/calendar/events | Get events from me
[**MeEventsCalendarUpdateCalendarView**](MeEventsCalendarEventApi.md#MeEventsCalendarUpdateCalendarView) | **Patch** /me/events({event-id})/calendar/calendarView({event-id1}) | Update the navigation property calendarView in me
[**MeEventsCalendarUpdateEvents**](MeEventsCalendarEventApi.md#MeEventsCalendarUpdateEvents) | **Patch** /me/events({event-id})/calendar/events({event-id1}) | Update the navigation property events in me



## MeEventsCalendarCreateCalendarView

> MicrosoftGraphEvent MeEventsCalendarCreateCalendarView(ctx, eventId, microsoftGraphEvent)
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


## MeEventsCalendarCreateEvents

> MicrosoftGraphEvent MeEventsCalendarCreateEvents(ctx, eventId, microsoftGraphEvent)
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


## MeEventsCalendarGetCalendarView

> MicrosoftGraphEvent MeEventsCalendarGetCalendarView(ctx, eventId, eventId1, optional)
Get calendarView from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***MeEventsCalendarGetCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsCalendarGetCalendarViewOpts struct


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


## MeEventsCalendarGetEvents

> MicrosoftGraphEvent MeEventsCalendarGetEvents(ctx, eventId, eventId1, optional)
Get events from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***MeEventsCalendarGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsCalendarGetEventsOpts struct


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


## MeEventsCalendarListCalendarView

> CollectionOfEvent MeEventsCalendarListCalendarView(ctx, eventId, optional)
Get calendarView from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeEventsCalendarListCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsCalendarListCalendarViewOpts struct


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


## MeEventsCalendarListEvents

> CollectionOfEvent MeEventsCalendarListEvents(ctx, eventId, optional)
Get events from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeEventsCalendarListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsCalendarListEventsOpts struct


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


## MeEventsCalendarUpdateCalendarView

> MeEventsCalendarUpdateCalendarView(ctx, eventId, eventId1, microsoftGraphEvent)
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


## MeEventsCalendarUpdateEvents

> MeEventsCalendarUpdateEvents(ctx, eventId, eventId1, microsoftGraphEvent)
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

