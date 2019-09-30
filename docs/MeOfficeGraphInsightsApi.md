# \MeOfficeGraphInsightsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetInsights**](MeOfficeGraphInsightsApi.md#MeGetInsights) | **Get** /me/insights | Get insights from me
[**MeUpdateInsights**](MeOfficeGraphInsightsApi.md#MeUpdateInsights) | **Patch** /me/insights | Update the navigation property insights in me



## MeGetInsights

> MicrosoftGraphOfficeGraphInsights MeGetInsights(ctx, optional)
Get insights from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetInsightsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOfficeGraphInsights**](microsoft.graph.officeGraphInsights.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateInsights

> MeUpdateInsights(ctx, microsoftGraphOfficeGraphInsights)
Update the navigation property insights in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOfficeGraphInsights** | [**MicrosoftGraphOfficeGraphInsights**](MicrosoftGraphOfficeGraphInsights.md)| New navigation property values | 

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

