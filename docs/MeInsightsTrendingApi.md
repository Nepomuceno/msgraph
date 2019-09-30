# \MeInsightsTrendingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeInsightsCreateTrending**](MeInsightsTrendingApi.md#MeInsightsCreateTrending) | **Post** /me/insights/trending | Create new navigation property to trending for me
[**MeInsightsGetTrending**](MeInsightsTrendingApi.md#MeInsightsGetTrending) | **Get** /me/insights/trending({trending-id}) | Get trending from me
[**MeInsightsListTrending**](MeInsightsTrendingApi.md#MeInsightsListTrending) | **Get** /me/insights/trending | Get trending from me
[**MeInsightsUpdateTrending**](MeInsightsTrendingApi.md#MeInsightsUpdateTrending) | **Patch** /me/insights/trending({trending-id}) | Update the navigation property trending in me



## MeInsightsCreateTrending

> MicrosoftGraphTrending MeInsightsCreateTrending(ctx, microsoftGraphTrending)
Create new navigation property to trending for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTrending** | [**MicrosoftGraphTrending**](MicrosoftGraphTrending.md)| New navigation property | 

### Return type

[**MicrosoftGraphTrending**](microsoft.graph.trending.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsGetTrending

> MicrosoftGraphTrending MeInsightsGetTrending(ctx, trendingId, optional)
Get trending from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string**| key: trending-id of trending | 
 **optional** | ***MeInsightsGetTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsGetTrendingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTrending**](microsoft.graph.trending.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsListTrending

> CollectionOfTrending MeInsightsListTrending(ctx, optional)
Get trending from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInsightsListTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsListTrendingOpts struct


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

[**CollectionOfTrending**](Collection of trending.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsUpdateTrending

> MeInsightsUpdateTrending(ctx, trendingId, microsoftGraphTrending)
Update the navigation property trending in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string**| key: trending-id of trending | 
**microsoftGraphTrending** | [**MicrosoftGraphTrending**](MicrosoftGraphTrending.md)| New navigation property values | 

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

