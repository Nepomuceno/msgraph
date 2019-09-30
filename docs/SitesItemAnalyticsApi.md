# \SitesItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesGetAnalytics**](SitesItemAnalyticsApi.md#SitesGetAnalytics) | **Get** /sites({site-id})/analytics | Get analytics from sites



## SitesGetAnalytics

> MicrosoftGraphItemAnalytics SitesGetAnalytics(ctx, siteId, optional)
Get analytics from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesGetAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesGetAnalyticsOpts struct


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
