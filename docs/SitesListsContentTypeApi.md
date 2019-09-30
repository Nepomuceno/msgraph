# \SitesListsContentTypeApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesListsCreateContentTypes**](SitesListsContentTypeApi.md#SitesListsCreateContentTypes) | **Post** /sites({site-id})/lists({list-id})/contentTypes | Create new navigation property to contentTypes for sites
[**SitesListsGetContentTypes**](SitesListsContentTypeApi.md#SitesListsGetContentTypes) | **Get** /sites({site-id})/lists({list-id})/contentTypes({contentType-id}) | Get contentTypes from sites
[**SitesListsListContentTypes**](SitesListsContentTypeApi.md#SitesListsListContentTypes) | **Get** /sites({site-id})/lists({list-id})/contentTypes | Get contentTypes from sites
[**SitesListsUpdateContentTypes**](SitesListsContentTypeApi.md#SitesListsUpdateContentTypes) | **Patch** /sites({site-id})/lists({list-id})/contentTypes({contentType-id}) | Update the navigation property contentTypes in sites



## SitesListsCreateContentTypes

> MicrosoftGraphContentType SitesListsCreateContentTypes(ctx, siteId, listId, microsoftGraphContentType)
Create new navigation property to contentTypes for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)| New navigation property | 

### Return type

[**MicrosoftGraphContentType**](microsoft.graph.contentType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsGetContentTypes

> MicrosoftGraphContentType SitesListsGetContentTypes(ctx, siteId, listId, contentTypeId, optional)
Get contentTypes from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***SitesListsGetContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsGetContentTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphContentType**](microsoft.graph.contentType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsListContentTypes

> CollectionOfContentType SitesListsListContentTypes(ctx, siteId, listId, optional)
Get contentTypes from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
 **optional** | ***SitesListsListContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsListContentTypesOpts struct


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

[**CollectionOfContentType**](Collection of contentType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsUpdateContentTypes

> SitesListsUpdateContentTypes(ctx, siteId, listId, contentTypeId, microsoftGraphContentType)
Update the navigation property contentTypes in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)| New navigation property values | 

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

