# \GroupsCalendarViewCalendarEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarViewCalendarCreateCalendarView**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarCreateCalendarView) | **Post** /groups({group-id})/calendarView({event-id})/calendar/calendarView | Create new navigation property to calendarView for groups
[**GroupsCalendarViewCalendarCreateEvents**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarCreateEvents) | **Post** /groups({group-id})/calendarView({event-id})/calendar/events | Create new navigation property to events for groups
[**GroupsCalendarViewCalendarGetCalendarView**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarGetCalendarView) | **Get** /groups({group-id})/calendarView({event-id})/calendar/calendarView({event-id1}) | Get calendarView from groups
[**GroupsCalendarViewCalendarGetEvents**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarGetEvents) | **Get** /groups({group-id})/calendarView({event-id})/calendar/events({event-id1}) | Get events from groups
[**GroupsCalendarViewCalendarListCalendarView**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarListCalendarView) | **Get** /groups({group-id})/calendarView({event-id})/calendar/calendarView | Get calendarView from groups
[**GroupsCalendarViewCalendarListEvents**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarListEvents) | **Get** /groups({group-id})/calendarView({event-id})/calendar/events | Get events from groups
[**GroupsCalendarViewCalendarUpdateCalendarView**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarUpdateCalendarView) | **Patch** /groups({group-id})/calendarView({event-id})/calendar/calendarView({event-id1}) | Update the navigation property calendarView in groups
[**GroupsCalendarViewCalendarUpdateEvents**](GroupsCalendarViewCalendarEventApi.md#GroupsCalendarViewCalendarUpdateEvents) | **Patch** /groups({group-id})/calendarView({event-id})/calendar/events({event-id1}) | Update the navigation property events in groups



## GroupsCalendarViewCalendarCreateCalendarView

> MicrosoftGraphEvent GroupsCalendarViewCalendarCreateCalendarView(ctx, groupId, eventId, microsoftGraphEvent)
Create new navigation property to calendarView for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsCalendarViewCalendarCreateEvents

> MicrosoftGraphEvent GroupsCalendarViewCalendarCreateEvents(ctx, groupId, eventId, microsoftGraphEvent)
Create new navigation property to events for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsCalendarViewCalendarGetCalendarView

> MicrosoftGraphEvent GroupsCalendarViewCalendarGetCalendarView(ctx, groupId, eventId, eventId1, optional)
Get calendarView from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarViewCalendarGetCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarViewCalendarGetCalendarViewOpts struct


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


## GroupsCalendarViewCalendarGetEvents

> MicrosoftGraphEvent GroupsCalendarViewCalendarGetEvents(ctx, groupId, eventId, eventId1, optional)
Get events from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarViewCalendarGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarViewCalendarGetEventsOpts struct


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


## GroupsCalendarViewCalendarListCalendarView

> CollectionOfEvent GroupsCalendarViewCalendarListCalendarView(ctx, groupId, eventId, optional)
Get calendarView from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarViewCalendarListCalendarViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarViewCalendarListCalendarViewOpts struct


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


## GroupsCalendarViewCalendarListEvents

> CollectionOfEvent GroupsCalendarViewCalendarListEvents(ctx, groupId, eventId, optional)
Get events from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarViewCalendarListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarViewCalendarListEventsOpts struct


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


## GroupsCalendarViewCalendarUpdateCalendarView

> GroupsCalendarViewCalendarUpdateCalendarView(ctx, groupId, eventId, eventId1, microsoftGraphEvent)
Update the navigation property calendarView in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsCalendarViewCalendarUpdateEvents

> GroupsCalendarViewCalendarUpdateEvents(ctx, groupId, eventId, eventId1, microsoftGraphEvent)
Update the navigation property events in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

