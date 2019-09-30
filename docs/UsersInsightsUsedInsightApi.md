# \UsersInsightsUsedInsightApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersInsightsCreateUsed**](UsersInsightsUsedInsightApi.md#UsersInsightsCreateUsed) | **Post** /users({user-id})/insights/used | Create new navigation property to used for users
[**UsersInsightsGetUsed**](UsersInsightsUsedInsightApi.md#UsersInsightsGetUsed) | **Get** /users({user-id})/insights/used({usedInsight-id}) | Get used from users
[**UsersInsightsListUsed**](UsersInsightsUsedInsightApi.md#UsersInsightsListUsed) | **Get** /users({user-id})/insights/used | Get used from users
[**UsersInsightsUpdateUsed**](UsersInsightsUsedInsightApi.md#UsersInsightsUpdateUsed) | **Patch** /users({user-id})/insights/used({usedInsight-id}) | Update the navigation property used in users



## UsersInsightsCreateUsed

> MicrosoftGraphUsedInsight UsersInsightsCreateUsed(ctx, userId, microsoftGraphUsedInsight)
Create new navigation property to used for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersInsightsGetUsed

> MicrosoftGraphUsedInsight UsersInsightsGetUsed(ctx, userId, usedInsightId, optional)
Get used from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***UsersInsightsGetUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsGetUsedOpts struct


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


## UsersInsightsListUsed

> CollectionOfUsedInsight UsersInsightsListUsed(ctx, userId, optional)
Get used from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInsightsListUsedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsListUsedOpts struct


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


## UsersInsightsUpdateUsed

> UsersInsightsUpdateUsed(ctx, userId, usedInsightId, microsoftGraphUsedInsight)
Update the navigation property used in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

