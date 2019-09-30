# \SitesListsItemsItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesListsItemsGetAnalytics**](SitesListsItemsItemAnalyticsApi.md#SitesListsItemsGetAnalytics) | **Get** /sites({site-id})/lists({list-id})/items({listItem-id})/analytics | Get analytics from sites



## SitesListsItemsGetAnalytics

> MicrosoftGraphItemAnalytics SitesListsItemsGetAnalytics(ctx, siteId, listId, listItemId, optional)
Get analytics from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***SitesListsItemsGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsItemsGetAnalyticsOpts struct


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

