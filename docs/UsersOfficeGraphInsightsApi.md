# \UsersOfficeGraphInsightsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetInsights**](UsersOfficeGraphInsightsApi.md#UsersGetInsights) | **Get** /users({user-id})/insights | Get insights from users
[**UsersUpdateInsights**](UsersOfficeGraphInsightsApi.md#UsersUpdateInsights) | **Patch** /users({user-id})/insights | Update the navigation property insights in users



## UsersGetInsights

> MicrosoftGraphOfficeGraphInsights UsersGetInsights(ctx, userId, optional)
Get insights from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetInsightsOpts struct


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


## UsersUpdateInsights

> UsersUpdateInsights(ctx, userId, microsoftGraphOfficeGraphInsights)
Update the navigation property insights in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

