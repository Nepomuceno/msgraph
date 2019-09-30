# \UsersEventsCalendarEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersEventsCalendarCreateCalendarView**](UsersEventsCalendarEventApi.md#UsersEventsCalendarCreateCalendarView) | **Post** /users({user-id})/events({event-id})/calendar/calendarView | Create new navigation property to calendarView for users
[**UsersEventsCalendarCreateEvents**](UsersEventsCalendarEventApi.md#UsersEventsCalendarCreateEvents) | **Post** /users({user-id})/events({event-id})/calendar/events | Create new navigation property to events for users
[**UsersEventsCalendarGetCalendarView**](UsersEventsCalendarEventApi.md#UsersEventsCalendarGetCalendarView) | **Get** /users({user-id})/events({event-id})/calendar/calendarView({event-id1}) | Get calendarView from users
[**UsersEventsCalendarGetEvents**](UsersEventsCalendarEventApi.md#UsersEventsCalendarGetEvents) | **Get** /users({user-id})/events({event-id})/calendar/events({event-id1}) | Get events from users
[**UsersEventsCalendarListCalendarView**](UsersEventsCalendarEventApi.md#UsersEventsCalendarListCalendarView) | **Get** /users({user-id})/events({event-id})/calendar/calendarView | Get calendarView from users
[**UsersEventsCalendarListEvents**](UsersEventsCalendarEventApi.md#UsersEventsCalendarListEvents) | **Get** /users({user-id})/events({event-id})/calendar/events | Get events from users
[**UsersEventsCalendarUpdateCalendarView**](UsersEventsCalendarEventApi.md#UsersEventsCalendarUpdateCalendarView) | **Patch** /users({user-id})/events({event-id})/calendar/calendarView({event-id1}) | Update the navigation property calendarView in users
[**UsersEventsCalendarUpdateEvents**](UsersEventsCalendarEventApi.md#UsersEventsCalendarUpdateEvents) | **Patch** /users({user-id})/events({event-id})/calendar/events({event-id1}) | Update the navigation property events in users



## UsersEventsCalendarCreateCalendarView

> MicrosoftGraphEvent UsersEventsCalendarCreateCalendarView(ctx, userId, eventId, microsoftGraphEvent)
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


## UsersEventsCalendarCreateEvents

> MicrosoftGraphEvent UsersEventsCalendarCreateEvents(ctx, userId, eventId, microsoftGraphEvent)
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


## UsersEventsCalendarGetCalendarView

> MicrosoftGraphEvent UsersEventsCalendarGetCalendarView(ctx, userId, eventId, eventId1, optional)
Get calendarView from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***UsersEventsCalendarGetCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersEventsCalendarGetCalendarViewOpts struct


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


## UsersEventsCalendarGetEvents

> MicrosoftGraphEvent UsersEventsCalendarGetEvents(ctx, userId, eventId, eventId1, optional)
Get events from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***UsersEventsCalendarGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersEventsCalendarGetEventsOpts struct


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


## UsersEventsCalendarListCalendarView

> CollectionOfEvent UsersEventsCalendarListCalendarView(ctx, userId, eventId, optional)
Get calendarView from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersEventsCalendarListCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersEventsCalendarListCalendarViewOpts struct


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


## UsersEventsCalendarListEvents

> CollectionOfEvent UsersEventsCalendarListEvents(ctx, userId, eventId, optional)
Get events from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersEventsCalendarListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersEventsCalendarListEventsOpts struct


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


## UsersEventsCalendarUpdateCalendarView

> UsersEventsCalendarUpdateCalendarView(ctx, userId, eventId, eventId1, microsoftGraphEvent)
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


## UsersEventsCalendarUpdateEvents

> UsersEventsCalendarUpdateEvents(ctx, userId, eventId, eventId1, microsoftGraphEvent)
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

