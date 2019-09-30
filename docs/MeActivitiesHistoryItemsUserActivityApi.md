# \MeActivitiesHistoryItemsUserActivityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeActivitiesHistoryItemsGetActivity**](MeActivitiesHistoryItemsUserActivityApi.md#MeActivitiesHistoryItemsGetActivity) | **Get** /me/activities({userActivity-id})/historyItems({activityHistoryItem-id})/activity | Get activity from me



## MeActivitiesHistoryItemsGetActivity

> MicrosoftGraphUserActivity MeActivitiesHistoryItemsGetActivity(ctx, userActivityId, activityHistoryItemId, optional)
Get activity from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**activityHistoryItemId** | **string**| key: activityHistoryItem-id of activityHistoryItem | 
 **optional** | ***MeActivitiesHistoryItemsGetActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeActivitiesHistoryItemsGetActivityOpts struct


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

