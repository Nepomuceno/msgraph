# \DriveListItemsItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListItemsGetAnalytics**](DriveListItemsItemAnalyticsApi.md#DriveListItemsGetAnalytics) | **Get** /drive/list/items({listItem-id})/analytics | Get analytics from drive



## DriveListItemsGetAnalytics

> MicrosoftGraphItemAnalytics DriveListItemsGetAnalytics(ctx, listItemId, optional)
Get analytics from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListItemsGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListItemsGetAnalyticsOpts struct


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

