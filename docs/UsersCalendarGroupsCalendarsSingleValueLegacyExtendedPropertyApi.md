# \UsersCalendarGroupsCalendarsSingleValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarGroupsCalendarsCreateSingleValueExtendedProperties**](UsersCalendarGroupsCalendarsSingleValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsCreateSingleValueExtendedProperties) | **Post** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersCalendarGroupsCalendarsGetSingleValueExtendedProperties**](UsersCalendarGroupsCalendarsSingleValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsGetSingleValueExtendedProperties) | **Get** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Get singleValueExtendedProperties from users
[**UsersCalendarGroupsCalendarsListSingleValueExtendedProperties**](UsersCalendarGroupsCalendarsSingleValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsListSingleValueExtendedProperties) | **Get** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersCalendarGroupsCalendarsUpdateSingleValueExtendedProperties**](UsersCalendarGroupsCalendarsSingleValueLegacyExtendedPropertyApi.md#UsersCalendarGroupsCalendarsUpdateSingleValueExtendedProperties) | **Patch** /users({user-id})/calendarGroups({calendarGroup-id})/calendars({calendar-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Update the navigation property singleValueExtendedProperties in users



## UsersCalendarGroupsCalendarsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersCalendarGroupsCalendarsCreateSingleValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, microsoftGraphSingleValueLegacyExtendedProperty)
Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
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


## UsersCalendarGroupsCalendarsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersCalendarGroupsCalendarsGetSingleValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, singleValueLegacyExtendedPropertyId, optional)
Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersCalendarGroupsCalendarsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarGroupsCalendarsGetSingleValueExtendedPropertiesOpts struct


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


## UsersCalendarGroupsCalendarsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersCalendarGroupsCalendarsListSingleValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, optional)
Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
 **optional** | ***UsersCalendarGroupsCalendarsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarGroupsCalendarsListSingleValueExtendedPropertiesOpts struct


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


## UsersCalendarGroupsCalendarsUpdateSingleValueExtendedProperties

> UsersCalendarGroupsCalendarsUpdateSingleValueExtendedProperties(ctx, userId, calendarGroupId, calendarId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)
Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**calendarId** | **string**| key: calendar-id of calendar | 
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

