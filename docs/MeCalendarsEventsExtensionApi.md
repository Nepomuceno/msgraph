# \MeCalendarsEventsExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarsEventsCreateExtensions**](MeCalendarsEventsExtensionApi.md#MeCalendarsEventsCreateExtensions) | **Post** /me/calendars({calendar-id})/events({event-id})/extensions | Create new navigation property to extensions for me
[**MeCalendarsEventsGetExtensions**](MeCalendarsEventsExtensionApi.md#MeCalendarsEventsGetExtensions) | **Get** /me/calendars({calendar-id})/events({event-id})/extensions({extension-id}) | Get extensions from me
[**MeCalendarsEventsListExtensions**](MeCalendarsEventsExtensionApi.md#MeCalendarsEventsListExtensions) | **Get** /me/calendars({calendar-id})/events({event-id})/extensions | Get extensions from me
[**MeCalendarsEventsUpdateExtensions**](MeCalendarsEventsExtensionApi.md#MeCalendarsEventsUpdateExtensions) | **Patch** /me/calendars({calendar-id})/events({event-id})/extensions({extension-id}) | Update the navigation property extensions in me



## MeCalendarsEventsCreateExtensions

> MicrosoftGraphExtension MeCalendarsEventsCreateExtensions(ctx, calendarId, eventId, microsoftGraphExtension)
Create new navigation property to extensions for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeCalendarsEventsGetExtensions

> MicrosoftGraphExtension MeCalendarsEventsGetExtensions(ctx, calendarId, eventId, extensionId, optional)
Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***MeCalendarsEventsGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarsEventsGetExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeCalendarsEventsListExtensions

> CollectionOfExtension MeCalendarsEventsListExtensions(ctx, calendarId, eventId, optional)
Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarsEventsListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarsEventsListExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfExtension**](Collection of extension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeCalendarsEventsUpdateExtensions

> MeCalendarsEventsUpdateExtensions(ctx, calendarId, eventId, extensionId, microsoftGraphExtension)
Update the navigation property extensions in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**extensionId** | **string**| key: extension-id of extension | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property values | 

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

