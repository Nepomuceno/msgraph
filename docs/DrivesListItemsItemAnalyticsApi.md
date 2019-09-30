# \DrivesListItemsItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListItemsGetAnalytics**](DrivesListItemsItemAnalyticsApi.md#DrivesListItemsGetAnalytics) | **Get** /drives({drive-id})/list/items({listItem-id})/analytics | Get analytics from drives



## DrivesListItemsGetAnalytics

> MicrosoftGraphItemAnalytics DrivesListItemsGetAnalytics(ctx, driveId, listItemId, optional)
Get analytics from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DrivesListItemsGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListItemsGetAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphItemAnalytics**](microsoft.graph.itemAnalytics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

