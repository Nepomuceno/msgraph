# \MeEventsCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeEventsGetCalendar**](MeEventsCalendarApi.md#MeEventsGetCalendar) | **Get** /me/events({event-id})/calendar | Get calendar from me
[**MeEventsUpdateCalendar**](MeEventsCalendarApi.md#MeEventsUpdateCalendar) | **Patch** /me/events({event-id})/calendar | Update the navigation property calendar in me



## MeEventsGetCalendar

> MicrosoftGraphCalendar MeEventsGetCalendar(ctx, eventId, optional)
Get calendar from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeEventsGetCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsGetCalendarOpts struct


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


## MeEventsUpdateCalendar

> MeEventsUpdateCalendar(ctx, eventId, microsoftGraphCalendar)
Update the navigation property calendar in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

