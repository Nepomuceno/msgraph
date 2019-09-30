# \UsersCalendarGroupsCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarGroupsCreateCalendars**](UsersCalendarGroupsCalendarApi.md#UsersCalendarGroupsCreateCalendars) | **Post** /users({user-id})/calendarGroups({calendarGroup-id})/calendars | Create new navigation property to calendars for users
[**UsersCalendarGroupsGetCalendars**](UsersCalendarGroupsCalendarApi.md#UsersCalendarGroupsGetCalendars) | **Get** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id}) | Get calendars from users
[**UsersCalendarGroupsListCalendars**](UsersCalendarGroupsCalendarApi.md#UsersCalendarGroupsListCalendars) | **Get** /users({user-id})/calendarGroups({calendarGroup-id})/calendars | Get calendars from users
[**UsersCalendarGroupsUpdateCalendars**](UsersCalendarGroupsCalendarApi.md#UsersCalendarGroupsUpdateCalendars) | **Patch** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id}) | Update the navigation property calendars in users



## UsersCalendarGroupsCreateCalendars

> MicrosoftGraphCalendar UsersCalendarGroupsCreateCalendars(ctx, userId, calendarGroupId, microsoftGraphCalendar)
Create new navigation property to calendars for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**microsoftGraphCalendar** | [**MicrosoftGraphCalendar**](MicrosoftGraphCalendar.md)| New navigation property | 

### Return type

[**MicrosoftGraphCalendar**](microsoft.graph.calendar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarGroupsGetCalendars

> MicrosoftGraphCalendar UsersCalendarGroupsGetCalendars(ctx, userId, calendarGroupId, calendarId, optional)
Get calendars from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
 **optional** | ***UsersCalendarGroupsGetCalendarsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarGroupsGetCalendarsOpts struct


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


## UsersCalendarGroupsListCalendars

> CollectionOfCalendar UsersCalendarGroupsListCalendars(ctx, userId, calendarGroupId, optional)
Get calendars from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
 **optional** | ***UsersCalendarGroupsListCalendarsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarGroupsListCalendarsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfCalendar**](Collection of calendar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarGroupsUpdateCalendars

> UsersCalendarGroupsUpdateCalendars(ctx, userId, calendarGroupId, calendarId, microsoftGraphCalendar)
Update the navigation property calendars in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
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

