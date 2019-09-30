# \SitesListsItemsListItemVersionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesListsItemsCreateVersions**](SitesListsItemsListItemVersionApi.md#SitesListsItemsCreateVersions) | **Post** /sites({site-id})/lists({list-id})/items({listItem-id})/versions | Create new navigation property to versions for sites
[**SitesListsItemsGetVersions**](SitesListsItemsListItemVersionApi.md#SitesListsItemsGetVersions) | **Get** /sites({site-id})/lists({list-id})/items({listItem-id})/versions({listItemVersion-id}) | Get versions from sites
[**SitesListsItemsListVersions**](SitesListsItemsListItemVersionApi.md#SitesListsItemsListVersions) | **Get** /sites({site-id})/lists({list-id})/items({listItem-id})/versions | Get versions from sites
[**SitesListsItemsUpdateVersions**](SitesListsItemsListItemVersionApi.md#SitesListsItemsUpdateVersions) | **Patch** /sites({site-id})/lists({list-id})/items({listItem-id})/versions({listItemVersion-id}) | Update the navigation property versions in sites
[**SitesListsItemsVersionsGetFields**](SitesListsItemsListItemVersionApi.md#SitesListsItemsVersionsGetFields) | **Get** /sites({site-id})/lists({list-id})/items({listItem-id})/versions({listItemVersion-id})/fields | Get fields from sites
[**SitesListsItemsVersionsUpdateFields**](SitesListsItemsListItemVersionApi.md#SitesListsItemsVersionsUpdateFields) | **Patch** /sites({site-id})/lists({list-id})/items({listItem-id})/versions({listItemVersion-id})/fields | Update the navigation property fields in sites



## SitesListsItemsCreateVersions

> MicrosoftGraphListItemVersion SitesListsItemsCreateVersions(ctx, siteId, listId, listItemId, microsoftGraphListItemVersion)
Create new navigation property to versions for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)| New navigation property | 

### Return type

[**MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetVersions

> MicrosoftGraphListItemVersion SitesListsItemsGetVersions(ctx, siteId, listId, listItemId, listItemVersionId, optional)
Get versions from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***SitesListsItemsGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsItemsGetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsListVersions

> CollectionOfListItemVersion SitesListsItemsListVersions(ctx, siteId, listId, listItemId, optional)
Get versions from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***SitesListsItemsListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsItemsListVersionsOpts struct


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

[**CollectionOfListItemVersion**](Collection of listItemVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsUpdateVersions

> SitesListsItemsUpdateVersions(ctx, siteId, listId, listItemId, listItemVersionId, microsoftGraphListItemVersion)
Update the navigation property versions in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
**microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)| New navigation property values | 

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


## SitesListsItemsVersionsGetFields

> MicrosoftGraphFieldValueSet SitesListsItemsVersionsGetFields(ctx, siteId, listId, listItemId, listItemVersionId, optional)
Get fields from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***SitesListsItemsVersionsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsItemsVersionsGetFieldsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](microsoft.graph.fieldValueSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsVersionsUpdateFields

> SitesListsItemsVersionsUpdateFields(ctx, siteId, listId, listItemId, listItemVersionId, microsoftGraphFieldValueSet)
Update the navigation property fields in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
**microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)| New navigation property values | 

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

