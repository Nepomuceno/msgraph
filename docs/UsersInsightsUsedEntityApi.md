# \UsersInsightsUsedEntityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersInsightsUsedGetResource**](UsersInsightsUsedEntityApi.md#UsersInsightsUsedGetResource) | **Get** /users({user-id})/insights/used({usedInsight-id})/resource | Get resource from users



## UsersInsightsUsedGetResource

> MicrosoftGraphEntity UsersInsightsUsedGetResource(ctx, userId, usedInsightId, optional)
Get resource from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**usedInsightId** | **string**| key: usedInsight-id of usedInsight | 
 **optional** | ***UsersInsightsUsedGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersInsightsUsedGetResourceOpts struct


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

