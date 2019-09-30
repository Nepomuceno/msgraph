# \MeOnenoteOnenoteOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOnenoteCreateOperations**](MeOnenoteOnenoteOperationApi.md#MeOnenoteCreateOperations) | **Post** /me/onenote/operations | Create new navigation property to operations for me
[**MeOnenoteGetOperations**](MeOnenoteOnenoteOperationApi.md#MeOnenoteGetOperations) | **Get** /me/onenote/operations({onenoteOperation-id}) | Get operations from me
[**MeOnenoteListOperations**](MeOnenoteOnenoteOperationApi.md#MeOnenoteListOperations) | **Get** /me/onenote/operations | Get operations from me
[**MeOnenoteUpdateOperations**](MeOnenoteOnenoteOperationApi.md#MeOnenoteUpdateOperations) | **Patch** /me/onenote/operations({onenoteOperation-id}) | Update the navigation property operations in me



## MeOnenoteCreateOperations

> MicrosoftGraphOnenoteOperation MeOnenoteCreateOperations(ctx, microsoftGraphOnenoteOperation)
Create new navigation property to operations for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOnenoteOperation** | [**MicrosoftGraphOnenoteOperation**](MicrosoftGraphOnenoteOperation.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenoteOperation**](microsoft.graph.onenoteOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteGetOperations

> MicrosoftGraphOnenoteOperation MeOnenoteGetOperations(ctx, onenoteOperationId, optional)
Get operations from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenoteOperationId** | **string**| key: onenoteOperation-id of onenoteOperation | 
 **optional** | ***MeOnenoteGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteGetOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenoteOperation**](microsoft.graph.onenoteOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteListOperations

> CollectionOfOnenoteOperation MeOnenoteListOperations(ctx, optional)
Get operations from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOnenoteListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteListOperationsOpts struct


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

[**CollectionOfOnenoteOperation**](Collection of onenoteOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteUpdateOperations

> MeOnenoteUpdateOperations(ctx, onenoteOperationId, microsoftGraphOnenoteOperation)
Update the navigation property operations in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenoteOperationId** | **string**| key: onenoteOperation-id of onenoteOperation | 
**microsoftGraphOnenoteOperation** | [**MicrosoftGraphOnenoteOperation**](MicrosoftGraphOnenoteOperation.md)| New navigation property values | 

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

