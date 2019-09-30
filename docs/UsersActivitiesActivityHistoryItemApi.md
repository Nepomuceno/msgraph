# \UsersActivitiesActivityHistoryItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersActivitiesCreateHistoryItems**](UsersActivitiesActivityHistoryItemApi.md#UsersActivitiesCreateHistoryItems) | **Post** /users({user-id})/activities({userActivity-id})/historyItems | Create new navigation property to historyItems for users
[**UsersActivitiesGetHistoryItems**](UsersActivitiesActivityHistoryItemApi.md#UsersActivitiesGetHistoryItems) | **Get** /users({user-id})/activities({userActivity-id})/historyItems({activityHistoryItem-id}) | Get historyItems from users
[**UsersActivitiesListHistoryItems**](UsersActivitiesActivityHistoryItemApi.md#UsersActivitiesListHistoryItems) | **Get** /users({user-id})/activities({userActivity-id})/historyItems | Get historyItems from users
[**UsersActivitiesUpdateHistoryItems**](UsersActivitiesActivityHistoryItemApi.md#UsersActivitiesUpdateHistoryItems) | **Patch** /users({user-id})/activities({userActivity-id})/historyItems({activityHistoryItem-id}) | Update the navigation property historyItems in users



## UsersActivitiesCreateHistoryItems

> MicrosoftGraphActivityHistoryItem UsersActivitiesCreateHistoryItems(ctx, userId, userActivityId, microsoftGraphActivityHistoryItem)
Create new navigation property to historyItems for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**microsoftGraphActivityHistoryItem** | [**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphActivityHistoryItem**](microsoft.graph.activityHistoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesGetHistoryItems

> MicrosoftGraphActivityHistoryItem UsersActivitiesGetHistoryItems(ctx, userId, userActivityId, activityHistoryItemId, optional)
Get historyItems from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**activityHistoryItemId** | **string**| key: activityHistoryItem-id of activityHistoryItem | 
 **optional** | ***UsersActivitiesGetHistoryItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersActivitiesGetHistoryItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphActivityHistoryItem**](microsoft.graph.activityHistoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesListHistoryItems

> CollectionOfActivityHistoryItem UsersActivitiesListHistoryItems(ctx, userId, userActivityId, optional)
Get historyItems from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
 **optional** | ***UsersActivitiesListHistoryItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersActivitiesListHistoryItemsOpts struct


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

[**CollectionOfActivityHistoryItem**](Collection of activityHistoryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersActivitiesUpdateHistoryItems

> UsersActivitiesUpdateHistoryItems(ctx, userId, userActivityId, activityHistoryItemId, microsoftGraphActivityHistoryItem)
Update the navigation property historyItems in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**activityHistoryItemId** | **string**| key: activityHistoryItem-id of activityHistoryItem | 
**microsoftGraphActivityHistoryItem** | [**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md)| New navigation property values | 

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

