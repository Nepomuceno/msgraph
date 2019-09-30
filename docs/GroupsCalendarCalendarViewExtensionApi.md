# \GroupsCalendarCalendarViewExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarCalendarViewCreateExtensions**](GroupsCalendarCalendarViewExtensionApi.md#GroupsCalendarCalendarViewCreateExtensions) | **Post** /groups({group-id})/calendar/calendarView({event-id})/extensions | Create new navigation property to extensions for groups
[**GroupsCalendarCalendarViewGetExtensions**](GroupsCalendarCalendarViewExtensionApi.md#GroupsCalendarCalendarViewGetExtensions) | **Get** /groups({group-id})/calendar/calendarView({event-id})/extensions({extension-id}) | Get extensions from groups
[**GroupsCalendarCalendarViewListExtensions**](GroupsCalendarCalendarViewExtensionApi.md#GroupsCalendarCalendarViewListExtensions) | **Get** /groups({group-id})/calendar/calendarView({event-id})/extensions | Get extensions from groups
[**GroupsCalendarCalendarViewUpdateExtensions**](GroupsCalendarCalendarViewExtensionApi.md#GroupsCalendarCalendarViewUpdateExtensions) | **Patch** /groups({group-id})/calendar/calendarView({event-id})/extensions({extension-id}) | Update the navigation property extensions in groups



## GroupsCalendarCalendarViewCreateExtensions

> MicrosoftGraphExtension GroupsCalendarCalendarViewCreateExtensions(ctx, groupId, eventId, microsoftGraphExtension)
Create new navigation property to extensions for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsCalendarCalendarViewGetExtensions

> MicrosoftGraphExtension GroupsCalendarCalendarViewGetExtensions(ctx, groupId, eventId, extensionId, optional)
Get extensions from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***GroupsCalendarCalendarViewGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarCalendarViewGetExtensionsOpts struct


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


## GroupsCalendarCalendarViewListExtensions

> CollectionOfExtension GroupsCalendarCalendarViewListExtensions(ctx, groupId, eventId, optional)
Get extensions from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarCalendarViewListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarCalendarViewListExtensionsOpts struct


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


## GroupsCalendarCalendarViewUpdateExtensions

> GroupsCalendarCalendarViewUpdateExtensions(ctx, groupId, eventId, extensionId, microsoftGraphExtension)
Update the navigation property extensions in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

