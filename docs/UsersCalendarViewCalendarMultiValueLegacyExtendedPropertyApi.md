# \UsersCalendarViewCalendarMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarViewCalendarCreateMultiValueExtendedProperties**](UsersCalendarViewCalendarMultiValueLegacyExtendedPropertyApi.md#UsersCalendarViewCalendarCreateMultiValueExtendedProperties) | **Post** /users({user-id})/calendarView({event-id})/calendar/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersCalendarViewCalendarGetMultiValueExtendedProperties**](UsersCalendarViewCalendarMultiValueLegacyExtendedPropertyApi.md#UsersCalendarViewCalendarGetMultiValueExtendedProperties) | **Get** /users({user-id})/calendarView({event-id})/calendar/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from users
[**UsersCalendarViewCalendarListMultiValueExtendedProperties**](UsersCalendarViewCalendarMultiValueLegacyExtendedPropertyApi.md#UsersCalendarViewCalendarListMultiValueExtendedProperties) | **Get** /users({user-id})/calendarView({event-id})/calendar/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersCalendarViewCalendarUpdateMultiValueExtendedProperties**](UsersCalendarViewCalendarMultiValueLegacyExtendedPropertyApi.md#UsersCalendarViewCalendarUpdateMultiValueExtendedProperties) | **Patch** /users({user-id})/calendarView({event-id})/calendar/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in users



## UsersCalendarViewCalendarCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersCalendarViewCalendarCreateMultiValueExtendedProperties(ctx, userId, eventId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarViewCalendarGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersCalendarViewCalendarGetMultiValueExtendedProperties(ctx, userId, eventId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersCalendarViewCalendarGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewCalendarGetMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarViewCalendarListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersCalendarViewCalendarListMultiValueExtendedProperties(ctx, userId, eventId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarViewCalendarListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarViewCalendarListMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfMultiValueLegacyExtendedProperty**](Collection of multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarViewCalendarUpdateMultiValueExtendedProperties

> UsersCalendarViewCalendarUpdateMultiValueExtendedProperties(ctx, userId, eventId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property values | 

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

