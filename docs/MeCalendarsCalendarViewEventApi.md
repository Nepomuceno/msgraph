# \MeCalendarsCalendarViewEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarsCalendarViewCreateInstances**](MeCalendarsCalendarViewEventApi.md#MeCalendarsCalendarViewCreateInstances) | **Post** /me/calendars({calendar-id})/calendarView({event-id})/instances | Create new navigation property to instances for me
[**MeCalendarsCalendarViewGetInstances**](MeCalendarsCalendarViewEventApi.md#MeCalendarsCalendarViewGetInstances) | **Get** /me/calendars({calendar-id})/calendarView({event-id})/instances({event-id1}) | Get instances from me
[**MeCalendarsCalendarViewListInstances**](MeCalendarsCalendarViewEventApi.md#MeCalendarsCalendarViewListInstances) | **Get** /me/calendars({calendar-id})/calendarView({event-id})/instances | Get instances from me
[**MeCalendarsCalendarViewUpdateInstances**](MeCalendarsCalendarViewEventApi.md#MeCalendarsCalendarViewUpdateInstances) | **Patch** /me/calendars({calendar-id})/calendarView({event-id})/instances({event-id1}) | Update the navigation property instances in me



## MeCalendarsCalendarViewCreateInstances

> MicrosoftGraphEvent MeCalendarsCalendarViewCreateInstances(ctx, calendarId, eventId, microsoftGraphEvent)
Create new navigation property to instances for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
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


## MeCalendarsCalendarViewGetInstances

> MicrosoftGraphEvent MeCalendarsCalendarViewGetInstances(ctx, calendarId, eventId, eventId1, optional)
Get instances from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarsCalendarViewGetInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarsCalendarViewGetInstancesOpts struct


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


## MeCalendarsCalendarViewListInstances

> CollectionOfEvent MeCalendarsCalendarViewListInstances(ctx, calendarId, eventId, optional)
Get instances from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarsCalendarViewListInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarsCalendarViewListInstancesOpts struct


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


## MeCalendarsCalendarViewUpdateInstances

> MeCalendarsCalendarViewUpdateInstances(ctx, calendarId, eventId, eventId1, microsoftGraphEvent)
Update the navigation property instances in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
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

