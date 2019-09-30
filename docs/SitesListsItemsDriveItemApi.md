# \SitesListsItemsDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesListsItemsGetDriveItem**](SitesListsItemsDriveItemApi.md#SitesListsItemsGetDriveItem) | **Get** /sites({site-id})/lists({list-id})/items({listItem-id})/driveItem | Get driveItem from sites
[**SitesListsItemsUpdateDriveItem**](SitesListsItemsDriveItemApi.md#SitesListsItemsUpdateDriveItem) | **Patch** /sites({site-id})/lists({list-id})/items({listItem-id})/driveItem | Update the navigation property driveItem in sites



## SitesListsItemsGetDriveItem

> MicrosoftGraphDriveItem SitesListsItemsGetDriveItem(ctx, siteId, listId, listItemId, optional)
Get driveItem from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***SitesListsItemsGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesListsItemsGetDriveItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsUpdateDriveItem

> SitesListsItemsUpdateDriveItem(ctx, siteId, listId, listItemId, microsoftGraphDriveItem)
Update the navigation property driveItem in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 
**microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)| New navigation property values | 

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

