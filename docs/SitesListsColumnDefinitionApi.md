# \SitesListsColumnDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesListsCreateColumns**](SitesListsColumnDefinitionApi.md#SitesListsCreateColumns) | **Post** /sites({site-id})/lists({list-id})/columns | Create new navigation property to columns for sites
[**SitesListsGetColumns**](SitesListsColumnDefinitionApi.md#SitesListsGetColumns) | **Get** /sites({site-id})/lists({list-id})/columns({columnDefinition-id}) | Get columns from sites
[**SitesListsListColumns**](SitesListsColumnDefinitionApi.md#SitesListsListColumns) | **Get** /sites({site-id})/lists({list-id})/columns | Get columns from sites
[**SitesListsUpdateColumns**](SitesListsColumnDefinitionApi.md#SitesListsUpdateColumns) | **Patch** /sites({site-id})/lists({list-id})/columns({columnDefinition-id}) | Update the navigation property columns in sites



## SitesListsCreateColumns

> MicrosoftGraphColumnDefinition SitesListsCreateColumns(ctx, siteId, listId, microsoftGraphColumnDefinition)
Create new navigation property to columns for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsGetColumns

> MicrosoftGraphColumnDefinition SitesListsGetColumns(ctx, siteId, listId, columnDefinitionId, optional)
Get columns from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
 **optional** | ***SitesListsGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsGetColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsListColumns

> CollectionOfColumnDefinition SitesListsListColumns(ctx, siteId, listId, optional)
Get columns from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
 **optional** | ***SitesListsListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsListColumnsOpts struct


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

[**CollectionOfColumnDefinition**](Collection of columnDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsUpdateColumns

> SitesListsUpdateColumns(ctx, siteId, listId, columnDefinitionId, microsoftGraphColumnDefinition)
Update the navigation property columns in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property values | 

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

