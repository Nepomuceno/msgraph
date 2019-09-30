# \UsersUserActivityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateActivities**](UsersUserActivityApi.md#UsersCreateActivities) | **Post** /users({user-id})/activities | Create new navigation property to activities for users
[**UsersGetActivities**](UsersUserActivityApi.md#UsersGetActivities) | **Get** /users({user-id})/activities({userActivity-id}) | Get activities from users
[**UsersListActivities**](UsersUserActivityApi.md#UsersListActivities) | **Get** /users({user-id})/activities | Get activities from users
[**UsersUpdateActivities**](UsersUserActivityApi.md#UsersUpdateActivities) | **Patch** /users({user-id})/activities({userActivity-id}) | Update the navigation property activities in users



## UsersCreateActivities

> MicrosoftGraphUserActivity UsersCreateActivities(ctx, userId, microsoftGraphUserActivity)
Create new navigation property to activities for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)| New navigation property | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetActivities

> MicrosoftGraphUserActivity UsersGetActivities(ctx, userId, userActivityId, optional)
Get activities from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
 **optional** | ***UsersGetActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetActivitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListActivities

> CollectionOfUserActivity UsersListActivities(ctx, userId, optional)
Get activities from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListActivitiesOpts struct


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

[**CollectionOfUserActivity**](Collection of userActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateActivities

> UsersUpdateActivities(ctx, userId, userActivityId, microsoftGraphUserActivity)
Update the navigation property activities in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)| New navigation property values | 

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

