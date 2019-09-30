# \UsersCalendarViewCalendarEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarViewCalendarCreateCalendarView**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarCreateCalendarView) | **Post** /users({user-id})/calendarView({event-id})/calendar/calendarView | Create new navigation property to calendarView for users
[**UsersCalendarViewCalendarCreateEvents**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarCreateEvents) | **Post** /users({user-id})/calendarView({event-id})/calendar/events | Create new navigation property to events for users
[**UsersCalendarViewCalendarGetCalendarView**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarGetCalendarView) | **Get** /users({user-id})/calendarView({event-id})/calendar/calendarView({event-id1}) | Get calendarView from users
[**UsersCalendarViewCalendarGetEvents**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarGetEvents) | **Get** /users({user-id})/calendarView({event-id})/calendar/events({event-id1}) | Get events from users
[**UsersCalendarViewCalendarListCalendarView**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarListCalendarView) | **Get** /users({user-id})/calendarView({event-id})/calendar/calendarView | Get calendarView from users
[**UsersCalendarViewCalendarListEvents**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarListEvents) | **Get** /users({user-id})/calendarView({event-id})/calendar/events | Get events from users
[**UsersCalendarViewCalendarUpdateCalendarView**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarUpdateCalendarView) | **Patch** /users({user-id})/calendarView({event-id})/calendar/calendarView({event-id1}) | Update the navigation property calendarView in users
[**UsersCalendarViewCalendarUpdateEvents**](UsersCalendarViewCalendarEventApi.md#UsersCalendarViewCalendarUpdateEvents) | **Patch** /users({user-id})/calendarView({event-id})/calendar/events({event-id1}) | Update the navigation property events in users



## UsersCalendarViewCalendarCreateCalendarView

> MicrosoftGraphEvent UsersCalendarViewCalendarCreateCalendarView(ctx, userId, eventId, microsoftGraphEvent)
Create new navigation property to calendarView for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersCalendarViewCalendarCreateEvents

> MicrosoftGraphEvent UsersCalendarViewCalendarCreateEvents(ctx, userId, eventId, microsoftGraphEvent)
Create new navigation property to events for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersCalendarViewCalendarGetCalendarView

> MicrosoftGraphEvent UsersCalendarViewCalendarGetCalendarView(ctx, userId, eventId, eventId1, optional)
Get calendarView from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarViewCalendarGetCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewCalendarGetCalendarViewOpts struct


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


## UsersCalendarViewCalendarGetEvents

> MicrosoftGraphEvent UsersCalendarViewCalendarGetEvents(ctx, userId, eventId, eventId1, optional)
Get events from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarViewCalendarGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewCalendarGetEventsOpts struct


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


## UsersCalendarViewCalendarListCalendarView

> CollectionOfEvent UsersCalendarViewCalendarListCalendarView(ctx, userId, eventId, optional)
Get calendarView from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarViewCalendarListCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewCalendarListCalendarViewOpts struct


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


## UsersCalendarViewCalendarListEvents

> CollectionOfEvent UsersCalendarViewCalendarListEvents(ctx, userId, eventId, optional)
Get events from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarViewCalendarListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewCalendarListEventsOpts struct


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


## UsersCalendarViewCalendarUpdateCalendarView

> UsersCalendarViewCalendarUpdateCalendarView(ctx, userId, eventId, eventId1, microsoftGraphEvent)
Update the navigation property calendarView in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersCalendarViewCalendarUpdateEvents

> UsersCalendarViewCalendarUpdateEvents(ctx, userId, eventId, eventId1, microsoftGraphEvent)
Update the navigation property events in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

