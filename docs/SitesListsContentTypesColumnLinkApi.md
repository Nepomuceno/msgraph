# \SitesListsContentTypesColumnLinkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesListsContentTypesCreateColumnLinks**](SitesListsContentTypesColumnLinkApi.md#SitesListsContentTypesCreateColumnLinks) | **Post** /sites({site-id})/lists({list-id})/contentTypes({contentType-id})/columnLinks | Create new navigation property to columnLinks for sites
[**SitesListsContentTypesGetColumnLinks**](SitesListsContentTypesColumnLinkApi.md#SitesListsContentTypesGetColumnLinks) | **Get** /sites({site-id})/lists({list-id})/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Get columnLinks from sites
[**SitesListsContentTypesListColumnLinks**](SitesListsContentTypesColumnLinkApi.md#SitesListsContentTypesListColumnLinks) | **Get** /sites({site-id})/lists({list-id})/contentTypes({contentType-id})/columnLinks | Get columnLinks from sites
[**SitesListsContentTypesUpdateColumnLinks**](SitesListsContentTypesColumnLinkApi.md#SitesListsContentTypesUpdateColumnLinks) | **Patch** /sites({site-id})/lists({list-id})/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Update the navigation property columnLinks in sites



## SitesListsContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink SitesListsContentTypesCreateColumnLinks(ctx, siteId, listId, contentTypeId, microsoftGraphColumnLink)
Create new navigation property to columnLinks for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsContentTypesGetColumnLinks

> MicrosoftGraphColumnLink SitesListsContentTypesGetColumnLinks(ctx, siteId, listId, contentTypeId, columnLinkId, optional)
Get columnLinks from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
 **optional** | ***SitesListsContentTypesGetColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsContentTypesGetColumnLinksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsContentTypesListColumnLinks

> CollectionOfColumnLink SitesListsContentTypesListColumnLinks(ctx, siteId, listId, contentTypeId, optional)
Get columnLinks from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***SitesListsContentTypesListColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsContentTypesListColumnLinksOpts struct


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

[**CollectionOfColumnLink**](Collection of columnLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsContentTypesUpdateColumnLinks

> SitesListsContentTypesUpdateColumnLinks(ctx, siteId, listId, contentTypeId, columnLinkId, microsoftGraphColumnLink)
Update the navigation property columnLinks in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
**microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)| New navigation property values | 

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

