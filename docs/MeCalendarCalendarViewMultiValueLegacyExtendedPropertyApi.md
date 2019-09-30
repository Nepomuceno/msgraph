# \MeCalendarCalendarViewMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCalendarCalendarViewCreateMultiValueExtendedProperties**](MeCalendarCalendarViewMultiValueLegacyExtendedPropertyApi.md#MeCalendarCalendarViewCreateMultiValueExtendedProperties) | **Post** /me/calendar/calendarView({event-id})/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeCalendarCalendarViewGetMultiValueExtendedProperties**](MeCalendarCalendarViewMultiValueLegacyExtendedPropertyApi.md#MeCalendarCalendarViewGetMultiValueExtendedProperties) | **Get** /me/calendar/calendarView({event-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from me
[**MeCalendarCalendarViewListMultiValueExtendedProperties**](MeCalendarCalendarViewMultiValueLegacyExtendedPropertyApi.md#MeCalendarCalendarViewListMultiValueExtendedProperties) | **Get** /me/calendar/calendarView({event-id})/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeCalendarCalendarViewUpdateMultiValueExtendedProperties**](MeCalendarCalendarViewMultiValueLegacyExtendedPropertyApi.md#MeCalendarCalendarViewUpdateMultiValueExtendedProperties) | **Patch** /me/calendar/calendarView({event-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in me



## MeCalendarCalendarViewCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeCalendarCalendarViewCreateMultiValueExtendedProperties(ctx, eventId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeCalendarCalendarViewGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeCalendarCalendarViewGetMultiValueExtendedProperties(ctx, eventId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***MeCalendarCalendarViewGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarCalendarViewGetMultiValueExtendedPropertiesOpts struct


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


## MeCalendarCalendarViewListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeCalendarCalendarViewListMultiValueExtendedProperties(ctx, eventId, optional)
Get multiValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeCalendarCalendarViewListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeCalendarCalendarViewListMultiValueExtendedPropertiesOpts struct


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


## MeCalendarCalendarViewUpdateMultiValueExtendedProperties

> MeCalendarCalendarViewUpdateMultiValueExtendedProperties(ctx, eventId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

