# \SitesOnenoteOnenoteOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesOnenoteCreateOperations**](SitesOnenoteOnenoteOperationApi.md#SitesOnenoteCreateOperations) | **Post** /sites({site-id})/onenote/operations | Create new navigation property to operations for sites
[**SitesOnenoteGetOperations**](SitesOnenoteOnenoteOperationApi.md#SitesOnenoteGetOperations) | **Get** /sites({site-id})/onenote/operations({onenoteOperation-id}) | Get operations from sites
[**SitesOnenoteListOperations**](SitesOnenoteOnenoteOperationApi.md#SitesOnenoteListOperations) | **Get** /sites({site-id})/onenote/operations | Get operations from sites
[**SitesOnenoteUpdateOperations**](SitesOnenoteOnenoteOperationApi.md#SitesOnenoteUpdateOperations) | **Patch** /sites({site-id})/onenote/operations({onenoteOperation-id}) | Update the navigation property operations in sites



## SitesOnenoteCreateOperations

> MicrosoftGraphOnenoteOperation SitesOnenoteCreateOperations(ctx, siteId, microsoftGraphOnenoteOperation)
Create new navigation property to operations for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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


## SitesOnenoteGetOperations

> MicrosoftGraphOnenoteOperation SitesOnenoteGetOperations(ctx, siteId, onenoteOperationId, optional)
Get operations from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenoteOperationId** | **string**| key: onenoteOperation-id of onenoteOperation | 
 **optional** | ***SitesOnenoteGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteGetOperationsOpts struct


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


## SitesOnenoteListOperations

> CollectionOfOnenoteOperation SitesOnenoteListOperations(ctx, siteId, optional)
Get operations from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesOnenoteListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteListOperationsOpts struct


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


## SitesOnenoteUpdateOperations

> SitesOnenoteUpdateOperations(ctx, siteId, onenoteOperationId, microsoftGraphOnenoteOperation)
Update the navigation property operations in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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

