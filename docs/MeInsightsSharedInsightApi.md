# \MeInsightsSharedInsightApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeInsightsCreateShared**](MeInsightsSharedInsightApi.md#MeInsightsCreateShared) | **Post** /me/insights/shared | Create new navigation property to shared for me
[**MeInsightsGetShared**](MeInsightsSharedInsightApi.md#MeInsightsGetShared) | **Get** /me/insights/shared({sharedInsight-id}) | Get shared from me
[**MeInsightsListShared**](MeInsightsSharedInsightApi.md#MeInsightsListShared) | **Get** /me/insights/shared | Get shared from me
[**MeInsightsUpdateShared**](MeInsightsSharedInsightApi.md#MeInsightsUpdateShared) | **Patch** /me/insights/shared({sharedInsight-id}) | Update the navigation property shared in me



## MeInsightsCreateShared

> MicrosoftGraphSharedInsight MeInsightsCreateShared(ctx, microsoftGraphSharedInsight)
Create new navigation property to shared for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSharedInsight** | [**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md)| New navigation property | 

### Return type

[**MicrosoftGraphSharedInsight**](microsoft.graph.sharedInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsGetShared

> MicrosoftGraphSharedInsight MeInsightsGetShared(ctx, sharedInsightId, optional)
Get shared from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***MeInsightsGetSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsGetSharedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSharedInsight**](microsoft.graph.sharedInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsListShared

> CollectionOfSharedInsight MeInsightsListShared(ctx, optional)
Get shared from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeInsightsListSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeInsightsListSharedOpts struct


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

[**CollectionOfSharedInsight**](Collection of sharedInsight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsUpdateShared

> MeInsightsUpdateShared(ctx, sharedInsightId, microsoftGraphSharedInsight)
Update the navigation property shared in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
**microsoftGraphSharedInsight** | [**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md)| New navigation property values | 

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

