# \SharesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListItemGetActivitiesByInterval4c35**](SharesFunctionsApi.md#SharesListItemGetActivitiesByInterval4c35) | **Get** /shares({sharedDriveItem-id})/listItem/getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SharesListItemsGetActivitiesByInterval4c35**](SharesFunctionsApi.md#SharesListItemsGetActivitiesByInterval4c35) | **Get** /shares({sharedDriveItem-id})/list/items({listItem-id})/getActivitiesByInterval() | Invoke function getActivitiesByInterval



## SharesListItemGetActivitiesByInterval4c35

> []AnyOfmicrosoftGraphItemActivityStat SharesListItemGetActivitiesByInterval4c35(ctx, sharedDriveItemId)
Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 

### Return type

[**[]AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemsGetActivitiesByInterval4c35

> []AnyOfmicrosoftGraphItemActivityStat SharesListItemsGetActivitiesByInterval4c35(ctx, sharedDriveItemId, listItemId)
Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 

### Return type

[**[]AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

