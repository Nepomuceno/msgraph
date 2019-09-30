# \UsersCalendarGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateCalendarGroups**](UsersCalendarGroupApi.md#UsersCreateCalendarGroups) | **Post** /users({user-id})/calendarGroups | Create new navigation property to calendarGroups for users
[**UsersGetCalendarGroups**](UsersCalendarGroupApi.md#UsersGetCalendarGroups) | **Get** /users({user-id})/calendarGroups({calendarGroup-id}) | Get calendarGroups from users
[**UsersListCalendarGroups**](UsersCalendarGroupApi.md#UsersListCalendarGroups) | **Get** /users({user-id})/calendarGroups | Get calendarGroups from users
[**UsersUpdateCalendarGroups**](UsersCalendarGroupApi.md#UsersUpdateCalendarGroups) | **Patch** /users({user-id})/calendarGroups({calendarGroup-id}) | Update the navigation property calendarGroups in users



## UsersCreateCalendarGroups

> MicrosoftGraphCalendarGroup UsersCreateCalendarGroups(ctx, userId, microsoftGraphCalendarGroup)
Create new navigation property to calendarGroups for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphCalendarGroup** | [**MicrosoftGraphCalendarGroup**](MicrosoftGraphCalendarGroup.md)| New navigation property | 

### Return type

[**MicrosoftGraphCalendarGroup**](microsoft.graph.calendarGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetCalendarGroups

> MicrosoftGraphCalendarGroup UsersGetCalendarGroups(ctx, userId, calendarGroupId, optional)
Get calendarGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
 **optional** | ***UsersGetCalendarGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetCalendarGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphCalendarGroup**](microsoft.graph.calendarGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListCalendarGroups

> CollectionOfCalendarGroup UsersListCalendarGroups(ctx, userId, optional)
Get calendarGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListCalendarGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListCalendarGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfCalendarGroup**](Collection of calendarGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateCalendarGroups

> UsersUpdateCalendarGroups(ctx, userId, calendarGroupId, microsoftGraphCalendarGroup)
Update the navigation property calendarGroups in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**microsoftGraphCalendarGroup** | [**MicrosoftGraphCalendarGroup**](MicrosoftGraphCalendarGroup.md)| New navigation property values | 

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

