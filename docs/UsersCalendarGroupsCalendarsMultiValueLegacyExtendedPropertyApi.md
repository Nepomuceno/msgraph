# \UsersCalendarGroupsCalendarsMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarGroupsCalendarsCreateMultiValueExtendedProperties**](UsersCalendarGroupsCalendarsMultiValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsCreateMultiValueExtendedProperties) | **Post** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersCalendarGroupsCalendarsGetMultiValueExtendedProperties**](UsersCalendarGroupsCalendarsMultiValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsGetMultiValueExtendedProperties) | **Get** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from users
[**UsersCalendarGroupsCalendarsListMultiValueExtendedProperties**](UsersCalendarGroupsCalendarsMultiValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsListMultiValueExtendedProperties) | **Get** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersCalendarGroupsCalendarsUpdateMultiValueExtendedProperties**](UsersCalendarGroupsCalendarsMultiValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsUpdateMultiValueExtendedProperties) | **Patch** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in users



## UsersCalendarGroupsCalendarsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersCalendarGroupsCalendarsCreateMultiValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
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


## UsersCalendarGroupsCalendarsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersCalendarGroupsCalendarsGetMultiValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersCalendarGroupsCalendarsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarGroupsCalendarsGetMultiValueExtendedPropertiesOpts struct


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


## UsersCalendarGroupsCalendarsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersCalendarGroupsCalendarsListMultiValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
 **optional** | ***UsersCalendarGroupsCalendarsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarGroupsCalendarsListMultiValueExtendedPropertiesOpts struct


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


## UsersCalendarGroupsCalendarsUpdateMultiValueExtendedProperties

> UsersCalendarGroupsCalendarsUpdateMultiValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
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

