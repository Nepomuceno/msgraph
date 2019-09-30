# \DrivesListItemsDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListItemsGetDriveItem**](DrivesListItemsDriveItemApi.md#DrivesListItemsGetDriveItem) | **Get** /drives({drive-id})/list/items({listItem-id})/driveItem | Get driveItem from drives
[**DrivesListItemsUpdateDriveItem**](DrivesListItemsDriveItemApi.md#DrivesListItemsUpdateDriveItem) | **Patch** /drives({drive-id})/list/items({listItem-id})/driveItem | Update the navigation property driveItem in drives



## DrivesListItemsGetDriveItem

> MicrosoftGraphDriveItem DrivesListItemsGetDriveItem(ctx, driveId, listItemId, optional)
Get driveItem from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DrivesListItemsGetDriveItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListItemsGetDriveItemOpts struct


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


## DrivesListItemsUpdateDriveItem

> DrivesListItemsUpdateDriveItem(ctx, driveId, listItemId, microsoftGraphDriveItem)
Update the navigation property driveItem in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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

