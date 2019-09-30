# \MeInsightsUsedEntityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeInsightsUsedGetResource**](MeInsightsUsedEntityApi.md#MeInsightsUsedGetResource) | **Get** /me/insights/used({usedInsight-id})/resource | Get resource from me



## MeInsightsUsedGetResource

> MicrosoftGraphEntity MeInsightsUsedGetResource(ctx, usedInsightId, optional)
Get resource from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***MeInsightsUsedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsUsedGetResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](microsoft.graph.entity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

