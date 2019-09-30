# \SharesListItemsDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListItemsGetDriveItem**](SharesListItemsDriveItemApi.md#SharesListItemsGetDriveItem) | **Get** /shares({sharedDriveItem-id})/list/items({listItem-id})/driveItem | Get driveItem from shares
[**SharesListItemsUpdateDriveItem**](SharesListItemsDriveItemApi.md#SharesListItemsUpdateDriveItem) | **Patch** /shares({sharedDriveItem-id})/list/items({listItem-id})/driveItem | Update the navigation property driveItem in shares



## SharesListItemsGetDriveItem

> MicrosoftGraphDriveItem SharesListItemsGetDriveItem(ctx, sharedDriveItemId, listItemId, optional)
Get driveItem from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***SharesListItemsGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemsGetDriveItemOpts struct


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


## SharesListItemsUpdateDriveItem

> SharesListItemsUpdateDriveItem(ctx, sharedDriveItemId, listItemId, microsoftGraphDriveItem)
Update the navigation property driveItem in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
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

