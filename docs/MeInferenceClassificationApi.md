# \MeInferenceClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeGetInferenceClassification**](MeInferenceClassificationApi.md#MeGetInferenceClassification) | **Get** /me/inferenceClassification | Get inferenceClassification from me
[**MeUpdateInferenceClassification**](MeInferenceClassificationApi.md#MeUpdateInferenceClassification) | **Patch** /me/inferenceClassification | Update the navigation property inferenceClassification in me



## MeGetInferenceClassification

> MicrosoftGraphInferenceClassification MeGetInferenceClassification(ctx, optional)
Get inferenceClassification from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeGetInferenceClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetInferenceClassificationOpts struct


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


## MeUpdateInferenceClassification

> MeUpdateInferenceClassification(ctx, microsoftGraphInferenceClassification)
Update the navigation property inferenceClassification in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

