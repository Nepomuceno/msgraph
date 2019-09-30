# \MeCalendarsMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarsCreateMultiValueExtendedProperties**](MeCalendarsMultiValueLegacyExtendedPropertyApi.md#MeCalendarsCreateMultiValueExtendedProperties) | **Post** /me/calendars({calendar-id})/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeCalendarsGetMultiValueExtendedProperties**](MeCalendarsMultiValueLegacyExtendedPropertyApi.md#MeCalendarsGetMultiValueExtendedProperties) | **Get** /me/calendars({calendar-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from me
[**MeCalendarsListMultiValueExtendedProperties**](MeCalendarsMultiValueLegacyExtendedPropertyApi.md#MeCalendarsListMultiValueExtendedProperties) | **Get** /me/calendars({calendar-id})/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeCalendarsUpdateMultiValueExtendedProperties**](MeCalendarsMultiValueLegacyExtendedPropertyApi.md#MeCalendarsUpdateMultiValueExtendedProperties) | **Patch** /me/calendars({calendar-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in me



## MeCalendarsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeCalendarsCreateMultiValueExtendedProperties(ctx, calendarId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
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


## MeCalendarsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeCalendarsGetMultiValueExtendedProperties(ctx, calendarId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***MeCalendarsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarsGetMultiValueExtendedPropertiesOpts struct


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


## MeCalendarsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeCalendarsListMultiValueExtendedProperties(ctx, calendarId, optional)
Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
 **optional** | ***MeCalendarsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarsListMultiValueExtendedPropertiesOpts struct


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


## MeCalendarsUpdateMultiValueExtendedProperties

> MeCalendarsUpdateMultiValueExtendedProperties(ctx, calendarId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarId** | **string**| key: calendar-id of calendar | 
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

