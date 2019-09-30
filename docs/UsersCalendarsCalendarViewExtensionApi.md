# \UsersCalendarsCalendarViewExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarsCalendarViewCreateExtensions**](UsersCalendarsCalendarViewExtensionApi.md#UsersCalendarsCalendarViewCreateExtensions) | **Post** /users({user-id})/calendars({calendar-id})/calendarView({event-id})/extensions | Create new navigation property to extensions for users
[**UsersCalendarsCalendarViewGetExtensions**](UsersCalendarsCalendarViewExtensionApi.md#UsersCalendarsCalendarViewGetExtensions) | **Get** /users({user-id})/calendars({calendar-id})/calendarView({event-id})/extensions({extension-id}) | Get extensions from users
[**UsersCalendarsCalendarViewListExtensions**](UsersCalendarsCalendarViewExtensionApi.md#UsersCalendarsCalendarViewListExtensions) | **Get** /users({user-id})/calendars({calendar-id})/calendarView({event-id})/extensions | Get extensions from users
[**UsersCalendarsCalendarViewUpdateExtensions**](UsersCalendarsCalendarViewExtensionApi.md#UsersCalendarsCalendarViewUpdateExtensions) | **Patch** /users({user-id})/calendars({calendar-id})/calendarView({event-id})/extensions({extension-id}) | Update the navigation property extensions in users



## UsersCalendarsCalendarViewCreateExtensions

> MicrosoftGraphExtension UsersCalendarsCalendarViewCreateExtensions(ctx, userId, calendarId, eventId, microsoftGraphExtension)
Create new navigation property to extensions for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersCalendarsCalendarViewGetExtensions

> MicrosoftGraphExtension UsersCalendarsCalendarViewGetExtensions(ctx, userId, calendarId, eventId, extensionId, optional)
Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***UsersCalendarsCalendarViewGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarsCalendarViewGetExtensionsOpts struct


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


## UsersCalendarsCalendarViewListExtensions

> CollectionOfExtension UsersCalendarsCalendarViewListExtensions(ctx, userId, calendarId, eventId, optional)
Get extensions from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarsCalendarViewListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarsCalendarViewListExtensionsOpts struct


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


## UsersCalendarsCalendarViewUpdateExtensions

> UsersCalendarsCalendarViewUpdateExtensions(ctx, userId, calendarId, eventId, extensionId, microsoftGraphExtension)
Update the navigation property extensions in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

