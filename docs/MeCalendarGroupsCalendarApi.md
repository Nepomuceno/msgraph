# \MeCalendarGroupsCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarGroupsCreateCalendars**](MeCalendarGroupsCalendarApi.md#MeCalendarGroupsCreateCalendars) | **Post** /me/calendarGroups({calendarGroup-id})/calendars | Create new navigation property to calendars for me
[**MeCalendarGroupsGetCalendars**](MeCalendarGroupsCalendarApi.md#MeCalendarGroupsGetCalendars) | **Get** /me/calendarGroups({calendarGroup-id})/calendars({calendar-id}) | Get calendars from me
[**MeCalendarGroupsListCalendars**](MeCalendarGroupsCalendarApi.md#MeCalendarGroupsListCalendars) | **Get** /me/calendarGroups({calendarGroup-id})/calendars | Get calendars from me
[**MeCalendarGroupsUpdateCalendars**](MeCalendarGroupsCalendarApi.md#MeCalendarGroupsUpdateCalendars) | **Patch** /me/calendarGroups({calendarGroup-id})/calendars({calendar-id}) | Update the navigation property calendars in me



## MeCalendarGroupsCreateCalendars

> MicrosoftGraphCalendar MeCalendarGroupsCreateCalendars(ctx, calendarGroupId, microsoftGraphCalendar)
Create new navigation property to calendars for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeCalendarGroupsGetCalendars

> MicrosoftGraphCalendar MeCalendarGroupsGetCalendars(ctx, calendarGroupId, calendarId, optional)
Get calendars from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
 **optional** | ***MeCalendarGroupsGetCalendarsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarGroupsGetCalendarsOpts struct


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


## MeCalendarGroupsListCalendars

> CollectionOfCalendar MeCalendarGroupsListCalendars(ctx, calendarGroupId, optional)
Get calendars from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
 **optional** | ***MeCalendarGroupsListCalendarsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarGroupsListCalendarsOpts struct


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


## MeCalendarGroupsUpdateCalendars

> MeCalendarGroupsUpdateCalendars(ctx, calendarGroupId, calendarId, microsoftGraphCalendar)
Update the navigation property calendars in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

