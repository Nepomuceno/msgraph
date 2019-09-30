# \MeInsightsUsedInsightApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeInsightsCreateUsed**](MeInsightsUsedInsightApi.md#MeInsightsCreateUsed) | **Post** /me/insights/used | Create new navigation property to used for me
[**MeInsightsGetUsed**](MeInsightsUsedInsightApi.md#MeInsightsGetUsed) | **Get** /me/insights/used({usedInsight-id}) | Get used from me
[**MeInsightsListUsed**](MeInsightsUsedInsightApi.md#MeInsightsListUsed) | **Get** /me/insights/used | Get used from me
[**MeInsightsUpdateUsed**](MeInsightsUsedInsightApi.md#MeInsightsUpdateUsed) | **Patch** /me/insights/used({usedInsight-id}) | Update the navigation property used in me



## MeInsightsCreateUsed

> MicrosoftGraphUsedInsight MeInsightsCreateUsed(ctx, microsoftGraphUsedInsight)
Create new navigation property to used for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphUsedInsight** | [**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md)| New navigation property | 

### Return type

[**MicrosoftGraphUsedInsight**](microsoft.graph.usedInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsGetUsed

> MicrosoftGraphUsedInsight MeInsightsGetUsed(ctx, usedInsightId, optional)
Get used from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***MeInsightsGetUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsGetUsedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUsedInsight**](microsoft.graph.usedInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsListUsed

> CollectionOfUsedInsight MeInsightsListUsed(ctx, optional)
Get used from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInsightsListUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsListUsedOpts struct


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

[**CollectionOfUsedInsight**](Collection of usedInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsUpdateUsed

> MeInsightsUpdateUsed(ctx, usedInsightId, microsoftGraphUsedInsight)
Update the navigation property used in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
**microsoftGraphUsedInsight** | [**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md)| New navigation property values | 

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

