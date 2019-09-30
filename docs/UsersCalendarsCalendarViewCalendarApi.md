# \UsersCalendarsCalendarViewCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarsCalendarViewGetCalendar**](UsersCalendarsCalendarViewCalendarApi.md#UsersCalendarsCalendarViewGetCalendar) | **Get** /users({user-id})/calendars({calendar-id})/calendarView({event-id})/calendar | Get calendar from users
[**UsersCalendarsCalendarViewUpdateCalendar**](UsersCalendarsCalendarViewCalendarApi.md#UsersCalendarsCalendarViewUpdateCalendar) | **Patch** /users({user-id})/calendars({calendar-id})/calendarView({event-id})/calendar | Update the navigation property calendar in users



## UsersCalendarsCalendarViewGetCalendar

> MicrosoftGraphCalendar UsersCalendarsCalendarViewGetCalendar(ctx, userId, calendarId, eventId, optional)
Get calendar from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarsCalendarViewGetCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarsCalendarViewGetCalendarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphCalendar**](microsoft.graph.calendar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarsCalendarViewUpdateCalendar

> UsersCalendarsCalendarViewUpdateCalendar(ctx, userId, calendarId, eventId, microsoftGraphCalendar)
Update the navigation property calendar in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**microsoftGraphCalendar** | [**MicrosoftGraphCalendar**](MicrosoftGraphCalendar.md)| New navigation property values | 

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

