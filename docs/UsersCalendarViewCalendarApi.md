# \UsersCalendarViewCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarViewGetCalendar**](UsersCalendarViewCalendarApi.md#UsersCalendarViewGetCalendar) | **Get** /users({user-id})/calendarView({event-id})/calendar | Get calendar from users
[**UsersCalendarViewUpdateCalendar**](UsersCalendarViewCalendarApi.md#UsersCalendarViewUpdateCalendar) | **Patch** /users({user-id})/calendarView({event-id})/calendar | Update the navigation property calendar in users



## UsersCalendarViewGetCalendar

> MicrosoftGraphCalendar UsersCalendarViewGetCalendar(ctx, userId, eventId, optional)
Get calendar from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarViewGetCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewGetCalendarOpts struct


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


## UsersCalendarViewUpdateCalendar

> UsersCalendarViewUpdateCalendar(ctx, userId, eventId, microsoftGraphCalendar)
Update the navigation property calendar in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

