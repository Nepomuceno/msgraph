# \DrivesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListItemsVersionsRestoreVersion**](DrivesActionsApi.md#DrivesListItemsVersionsRestoreVersion) | **Post** /drives({drive-id})/list/items({listItem-id})/versions({listItemVersion-id})/restoreVersion | Invoke action restoreVersion



## DrivesListItemsVersionsRestoreVersion

> DrivesListItemsVersionsRestoreVersion(ctx, driveId, listItemId, listItemVersionId)
Invoke action restoreVersion

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

