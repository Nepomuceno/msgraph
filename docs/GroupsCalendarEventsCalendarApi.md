# \GroupsCalendarEventsCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarEventsGetCalendar**](GroupsCalendarEventsCalendarApi.md#GroupsCalendarEventsGetCalendar) | **Get** /groups({group-id})/calendar/events({event-id})/calendar | Get calendar from groups
[**GroupsCalendarEventsUpdateCalendar**](GroupsCalendarEventsCalendarApi.md#GroupsCalendarEventsUpdateCalendar) | **Patch** /groups({group-id})/calendar/events({event-id})/calendar | Update the navigation property calendar in groups



## GroupsCalendarEventsGetCalendar

> MicrosoftGraphCalendar GroupsCalendarEventsGetCalendar(ctx, groupId, eventId, optional)
Get calendar from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarEventsGetCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarEventsGetCalendarOpts struct


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


## GroupsCalendarEventsUpdateCalendar

> GroupsCalendarEventsUpdateCalendar(ctx, groupId, eventId, microsoftGraphCalendar)
Update the navigation property calendar in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

