# \SitesOnenoteNotebookApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesOnenoteCreateNotebooks**](SitesOnenoteNotebookApi.md#SitesOnenoteCreateNotebooks) | **Post** /sites({site-id})/onenote/notebooks | Create new navigation property to notebooks for sites
[**SitesOnenoteGetNotebooks**](SitesOnenoteNotebookApi.md#SitesOnenoteGetNotebooks) | **Get** /sites({site-id})/onenote/notebooks({notebook-id}) | Get notebooks from sites
[**SitesOnenoteListNotebooks**](SitesOnenoteNotebookApi.md#SitesOnenoteListNotebooks) | **Get** /sites({site-id})/onenote/notebooks | Get notebooks from sites
[**SitesOnenoteUpdateNotebooks**](SitesOnenoteNotebookApi.md#SitesOnenoteUpdateNotebooks) | **Patch** /sites({site-id})/onenote/notebooks({notebook-id}) | Update the navigation property notebooks in sites



## SitesOnenoteCreateNotebooks

> MicrosoftGraphNotebook SitesOnenoteCreateNotebooks(ctx, siteId, microsoftGraphNotebook)
Create new navigation property to notebooks for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteGetNotebooks

> MicrosoftGraphNotebook SitesOnenoteGetNotebooks(ctx, siteId, notebookId, optional)
Get notebooks from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**notebookId** | **string**| key: notebook-id of notebook | 
 **optional** | ***SitesOnenoteGetNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteGetNotebooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteListNotebooks

> CollectionOfNotebook SitesOnenoteListNotebooks(ctx, siteId, optional)
Get notebooks from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesOnenoteListNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteListNotebooksOpts struct


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

[**CollectionOfNotebook**](Collection of notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteUpdateNotebooks

> SitesOnenoteUpdateNotebooks(ctx, siteId, notebookId, microsoftGraphNotebook)
Update the navigation property notebooks in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**notebookId** | **string**| key: notebook-id of notebook | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property values | 

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

