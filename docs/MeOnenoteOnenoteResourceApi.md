# \MeOnenoteOnenoteResourceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOnenoteCreateResources**](MeOnenoteOnenoteResourceApi.md#MeOnenoteCreateResources) | **Post** /me/onenote/resources | Create new navigation property to resources for me
[**MeOnenoteGetResources**](MeOnenoteOnenoteResourceApi.md#MeOnenoteGetResources) | **Get** /me/onenote/resources({onenoteResource-id}) | Get resources from me
[**MeOnenoteListResources**](MeOnenoteOnenoteResourceApi.md#MeOnenoteListResources) | **Get** /me/onenote/resources | Get resources from me
[**MeOnenoteUpdateResources**](MeOnenoteOnenoteResourceApi.md#MeOnenoteUpdateResources) | **Patch** /me/onenote/resources({onenoteResource-id}) | Update the navigation property resources in me



## MeOnenoteCreateResources

> MicrosoftGraphOnenoteResource MeOnenoteCreateResources(ctx, microsoftGraphOnenoteResource)
Create new navigation property to resources for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOnenoteResource** | [**MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenoteResource**](microsoft.graph.onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteGetResources

> MicrosoftGraphOnenoteResource MeOnenoteGetResources(ctx, onenoteResourceId, optional)
Get resources from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
 **optional** | ***MeOnenoteGetResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteGetResourcesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenoteResource**](microsoft.graph.onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteListResources

> CollectionOfOnenoteResource MeOnenoteListResources(ctx, optional)
Get resources from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOnenoteListResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteListResourcesOpts struct


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

[**CollectionOfOnenoteResource**](Collection of onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteUpdateResources

> MeOnenoteUpdateResources(ctx, onenoteResourceId, microsoftGraphOnenoteResource)
Update the navigation property resources in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
**microsoftGraphOnenoteResource** | [**MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md)| New navigation property values | 

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

