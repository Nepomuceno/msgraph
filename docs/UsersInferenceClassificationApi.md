# \UsersInferenceClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetInferenceClassification**](UsersInferenceClassificationApi.md#UsersGetInferenceClassification) | **Get** /users({user-id})/inferenceClassification | Get inferenceClassification from users
[**UsersUpdateInferenceClassification**](UsersInferenceClassificationApi.md#UsersUpdateInferenceClassification) | **Patch** /users({user-id})/inferenceClassification | Update the navigation property inferenceClassification in users



## UsersGetInferenceClassification

> MicrosoftGraphInferenceClassification UsersGetInferenceClassification(ctx, userId, optional)
Get inferenceClassification from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetInferenceClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetInferenceClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphInferenceClassification**](microsoft.graph.inferenceClassification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateInferenceClassification

> UsersUpdateInferenceClassification(ctx, userId, microsoftGraphInferenceClassification)
Update the navigation property inferenceClassification in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphInferenceClassification** | [**MicrosoftGraphInferenceClassification**](MicrosoftGraphInferenceClassification.md)| New navigation property values | 

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

