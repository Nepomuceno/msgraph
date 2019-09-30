# \MeInsightsSharedEntityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeInsightsSharedGetLastSharedMethod**](MeInsightsSharedEntityApi.md#MeInsightsSharedGetLastSharedMethod) | **Get** /me/insights/shared({sharedInsight-id})/lastSharedMethod | Get lastSharedMethod from me
[**MeInsightsSharedGetResource**](MeInsightsSharedEntityApi.md#MeInsightsSharedGetResource) | **Get** /me/insights/shared({sharedInsight-id})/resource | Get resource from me



## MeInsightsSharedGetLastSharedMethod

> MicrosoftGraphEntity MeInsightsSharedGetLastSharedMethod(ctx, sharedInsightId, optional)
Get lastSharedMethod from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***MeInsightsSharedGetLastSharedMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsSharedGetLastSharedMethodOpts struct


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


## MeInsightsSharedGetResource

> MicrosoftGraphEntity MeInsightsSharedGetResource(ctx, sharedInsightId, optional)
Get resource from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***MeInsightsSharedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsSharedGetResourceOpts struct


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

