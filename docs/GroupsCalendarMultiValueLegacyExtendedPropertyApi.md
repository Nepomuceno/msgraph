# \GroupsCalendarMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarCreateMultiValueExtendedProperties**](GroupsCalendarMultiValueLegacyExtendedPropertyApi.md#GroupsCalendarCreateMultiValueExtendedProperties) | **Post** /groups({group-id})/calendar/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for groups
[**GroupsCalendarGetMultiValueExtendedProperties**](GroupsCalendarMultiValueLegacyExtendedPropertyApi.md#GroupsCalendarGetMultiValueExtendedProperties) | **Get** /groups({group-id})/calendar/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from groups
[**GroupsCalendarListMultiValueExtendedProperties**](GroupsCalendarMultiValueLegacyExtendedPropertyApi.md#GroupsCalendarListMultiValueExtendedProperties) | **Get** /groups({group-id})/calendar/multiValueExtendedProperties | Get multiValueExtendedProperties from groups
[**GroupsCalendarUpdateMultiValueExtendedProperties**](GroupsCalendarMultiValueLegacyExtendedPropertyApi.md#GroupsCalendarUpdateMultiValueExtendedProperties) | **Patch** /groups({group-id})/calendar/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in groups



## GroupsCalendarCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsCalendarCreateMultiValueExtendedProperties(ctx, groupId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsCalendarGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsCalendarGetMultiValueExtendedProperties(ctx, groupId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***GroupsCalendarGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarGetMultiValueExtendedPropertiesOpts struct


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


## GroupsCalendarListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty GroupsCalendarListMultiValueExtendedProperties(ctx, groupId, optional)
Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsCalendarListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarListMultiValueExtendedPropertiesOpts struct


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


## GroupsCalendarUpdateMultiValueExtendedProperties

> GroupsCalendarUpdateMultiValueExtendedProperties(ctx, groupId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

