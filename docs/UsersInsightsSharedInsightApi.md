# \UsersInsightsSharedInsightApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersInsightsCreateShared**](UsersInsightsSharedInsightApi.md#UsersInsightsCreateShared) | **Post** /users({user-id})/insights/shared | Create new navigation property to shared for users
[**UsersInsightsGetShared**](UsersInsightsSharedInsightApi.md#UsersInsightsGetShared) | **Get** /users({user-id})/insights/shared({sharedInsight-id}) | Get shared from users
[**UsersInsightsListShared**](UsersInsightsSharedInsightApi.md#UsersInsightsListShared) | **Get** /users({user-id})/insights/shared | Get shared from users
[**UsersInsightsUpdateShared**](UsersInsightsSharedInsightApi.md#UsersInsightsUpdateShared) | **Patch** /users({user-id})/insights/shared({sharedInsight-id}) | Update the navigation property shared in users



## UsersInsightsCreateShared

> MicrosoftGraphSharedInsight UsersInsightsCreateShared(ctx, userId, microsoftGraphSharedInsight)
Create new navigation property to shared for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersInsightsGetShared

> MicrosoftGraphSharedInsight UsersInsightsGetShared(ctx, userId, sharedInsightId, optional)
Get shared from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***UsersInsightsGetSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsGetSharedOpts struct


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


## UsersInsightsListShared

> CollectionOfSharedInsight UsersInsightsListShared(ctx, userId, optional)
Get shared from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersInsightsListSharedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsListSharedOpts struct


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


## UsersInsightsUpdateShared

> UsersInsightsUpdateShared(ctx, userId, sharedInsightId, microsoftGraphSharedInsight)
Update the navigation property shared in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

