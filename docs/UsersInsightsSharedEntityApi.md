# \UsersInsightsSharedEntityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersInsightsSharedGetLastSharedMethod**](UsersInsightsSharedEntityApi.md#UsersInsightsSharedGetLastSharedMethod) | **Get** /users({user-id})/insights/shared({sharedInsight-id})/lastSharedMethod | Get lastSharedMethod from users
[**UsersInsightsSharedGetResource**](UsersInsightsSharedEntityApi.md#UsersInsightsSharedGetResource) | **Get** /users({user-id})/insights/shared({sharedInsight-id})/resource | Get resource from users



## UsersInsightsSharedGetLastSharedMethod

> MicrosoftGraphEntity UsersInsightsSharedGetLastSharedMethod(ctx, userId, sharedInsightId, optional)
Get lastSharedMethod from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***UsersInsightsSharedGetLastSharedMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsSharedGetLastSharedMethodOpts struct


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


## UsersInsightsSharedGetResource

> MicrosoftGraphEntity UsersInsightsSharedGetResource(ctx, userId, sharedInsightId, optional)
Get resource from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sharedInsightId** | **string**| key: sharedInsight-id of sharedInsight | 
 **optional** | ***UsersInsightsSharedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsSharedGetResourceOpts struct


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

