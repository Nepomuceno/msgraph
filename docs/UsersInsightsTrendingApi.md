# \UsersInsightsTrendingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersInsightsCreateTrending**](UsersInsightsTrendingApi.md#UsersInsightsCreateTrending) | **Post** /users({user-id})/insights/trending | Create new navigation property to trending for users
[**UsersInsightsGetTrending**](UsersInsightsTrendingApi.md#UsersInsightsGetTrending) | **Get** /users({user-id})/insights/trending({trending-id}) | Get trending from users
[**UsersInsightsListTrending**](UsersInsightsTrendingApi.md#UsersInsightsListTrending) | **Get** /users({user-id})/insights/trending | Get trending from users
[**UsersInsightsUpdateTrending**](UsersInsightsTrendingApi.md#UsersInsightsUpdateTrending) | **Patch** /users({user-id})/insights/trending({trending-id}) | Update the navigation property trending in users



## UsersInsightsCreateTrending

> MicrosoftGraphTrending UsersInsightsCreateTrending(ctx, userId, microsoftGraphTrending)
Create new navigation property to trending for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersInsightsGetTrending

> MicrosoftGraphTrending UsersInsightsGetTrending(ctx, userId, trendingId, optional)
Get trending from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**trendingId** | **string**| key: trending-id of trending | 
 **optional** | ***UsersInsightsGetTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsGetTrendingOpts struct


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


## UsersInsightsListTrending

> CollectionOfTrending UsersInsightsListTrending(ctx, userId, optional)
Get trending from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInsightsListTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsListTrendingOpts struct


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


## UsersInsightsUpdateTrending

> UsersInsightsUpdateTrending(ctx, userId, trendingId, microsoftGraphTrending)
Update the navigation property trending in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

