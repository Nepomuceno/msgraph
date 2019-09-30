# \MeEventsCalendarSingleValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeEventsCalendarCreateSingleValueExtendedProperties**](MeEventsCalendarSingleValueLegacyExtendedPropertyApi.md#MeEventsCalendarCreateSingleValueExtendedProperties) | **Post** /me/events({event-id})/calendar/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeEventsCalendarGetSingleValueExtendedProperties**](MeEventsCalendarSingleValueLegacyExtendedPropertyApi.md#MeEventsCalendarGetSingleValueExtendedProperties) | **Get** /me/events({event-id})/calendar/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Get singleValueExtendedProperties from me
[**MeEventsCalendarListSingleValueExtendedProperties**](MeEventsCalendarSingleValueLegacyExtendedPropertyApi.md#MeEventsCalendarListSingleValueExtendedProperties) | **Get** /me/events({event-id})/calendar/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeEventsCalendarUpdateSingleValueExtendedProperties**](MeEventsCalendarSingleValueLegacyExtendedPropertyApi.md#MeEventsCalendarUpdateSingleValueExtendedProperties) | **Patch** /me/events({event-id})/calendar/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Update the navigation property singleValueExtendedProperties in me



## MeEventsCalendarCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeEventsCalendarCreateSingleValueExtendedProperties(ctx, eventId, microsoftGraphSingleValueLegacyExtendedProperty)
Create new navigation property to singleValueExtendedProperties for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeEventsCalendarGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeEventsCalendarGetSingleValueExtendedProperties(ctx, eventId, singleValueLegacyExtendedPropertyId, optional)
Get singleValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***MeEventsCalendarGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsCalendarGetSingleValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeEventsCalendarListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeEventsCalendarListSingleValueExtendedProperties(ctx, eventId, optional)
Get singleValueExtendedProperties from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
 **optional** | ***MeEventsCalendarListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeEventsCalendarListSingleValueExtendedPropertiesOpts struct


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

[**CollectionOfSingleValueLegacyExtendedProperty**](Collection of singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeEventsCalendarUpdateSingleValueExtendedProperties

> MeEventsCalendarUpdateSingleValueExtendedProperties(ctx, eventId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)
Update the navigation property singleValueExtendedProperties in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| key: event-id of event | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property values | 

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

