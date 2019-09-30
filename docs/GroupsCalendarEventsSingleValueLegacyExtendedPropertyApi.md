# \GroupsCalendarEventsSingleValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarEventsCreateSingleValueExtendedProperties**](GroupsCalendarEventsSingleValueLegacyExtendedPropertyApi.md#GroupsCalendarEventsCreateSingleValueExtendedProperties) | **Post** /groups({group-id})/calendar/events({event-id})/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for groups
[**GroupsCalendarEventsGetSingleValueExtendedProperties**](GroupsCalendarEventsSingleValueLegacyExtendedPropertyApi.md#GroupsCalendarEventsGetSingleValueExtendedProperties) | **Get** /groups({group-id})/calendar/events({event-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Get singleValueExtendedProperties from groups
[**GroupsCalendarEventsListSingleValueExtendedProperties**](GroupsCalendarEventsSingleValueLegacyExtendedPropertyApi.md#GroupsCalendarEventsListSingleValueExtendedProperties) | **Get** /groups({group-id})/calendar/events({event-id})/singleValueExtendedProperties | Get singleValueExtendedProperties from groups
[**GroupsCalendarEventsUpdateSingleValueExtendedProperties**](GroupsCalendarEventsSingleValueLegacyExtendedPropertyApi.md#GroupsCalendarEventsUpdateSingleValueExtendedProperties) | **Patch** /groups({group-id})/calendar/events({event-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Update the navigation property singleValueExtendedProperties in groups



## GroupsCalendarEventsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsCalendarEventsCreateSingleValueExtendedProperties(ctx, groupId, eventId, microsoftGraphSingleValueLegacyExtendedProperty)
Create new navigation property to singleValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsCalendarEventsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsCalendarEventsGetSingleValueExtendedProperties(ctx, groupId, eventId, singleValueLegacyExtendedPropertyId, optional)
Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***GroupsCalendarEventsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarEventsGetSingleValueExtendedPropertiesOpts struct


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


## GroupsCalendarEventsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty GroupsCalendarEventsListSingleValueExtendedProperties(ctx, groupId, eventId, optional)
Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarEventsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarEventsListSingleValueExtendedPropertiesOpts struct


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


## GroupsCalendarEventsUpdateSingleValueExtendedProperties

> GroupsCalendarEventsUpdateSingleValueExtendedProperties(ctx, groupId, eventId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)
Update the navigation property singleValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

