# \GroupsCalendarApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGetCalendar**](GroupsCalendarApi.md#GroupsGetCalendar) | **Get** /groups({group-id})/calendar | Get calendar from groups
[**GroupsUpdateCalendar**](GroupsCalendarApi.md#GroupsUpdateCalendar) | **Patch** /groups({group-id})/calendar | Update the navigation property calendar in groups



## GroupsGetCalendar

> MicrosoftGraphCalendar GroupsGetCalendar(ctx, groupId, optional)
Get calendar from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsGetCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetCalendarOpts struct


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


## GroupsUpdateCalendar

> GroupsUpdateCalendar(ctx, groupId, microsoftGraphCalendar)
Update the navigation property calendar in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

