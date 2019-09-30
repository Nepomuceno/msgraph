# \SitesOnenoteOnenoteResourceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesOnenoteCreateResources**](SitesOnenoteOnenoteResourceApi.md#SitesOnenoteCreateResources) | **Post** /sites({site-id})/onenote/resources | Create new navigation property to resources for sites
[**SitesOnenoteGetResources**](SitesOnenoteOnenoteResourceApi.md#SitesOnenoteGetResources) | **Get** /sites({site-id})/onenote/resources({onenoteResource-id}) | Get resources from sites
[**SitesOnenoteListResources**](SitesOnenoteOnenoteResourceApi.md#SitesOnenoteListResources) | **Get** /sites({site-id})/onenote/resources | Get resources from sites
[**SitesOnenoteUpdateResources**](SitesOnenoteOnenoteResourceApi.md#SitesOnenoteUpdateResources) | **Patch** /sites({site-id})/onenote/resources({onenoteResource-id}) | Update the navigation property resources in sites



## SitesOnenoteCreateResources

> MicrosoftGraphOnenoteResource SitesOnenoteCreateResources(ctx, siteId, microsoftGraphOnenoteResource)
Create new navigation property to resources for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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


## SitesOnenoteGetResources

> MicrosoftGraphOnenoteResource SitesOnenoteGetResources(ctx, siteId, onenoteResourceId, optional)
Get resources from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
 **optional** | ***SitesOnenoteGetResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteGetResourcesOpts struct


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


## SitesOnenoteListResources

> CollectionOfOnenoteResource SitesOnenoteListResources(ctx, siteId, optional)
Get resources from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesOnenoteListResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteListResourcesOpts struct


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


## SitesOnenoteUpdateResources

> SitesOnenoteUpdateResources(ctx, siteId, onenoteResourceId, microsoftGraphOnenoteResource)
Update the navigation property resources in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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

