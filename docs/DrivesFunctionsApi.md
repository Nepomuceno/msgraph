# \DrivesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListItemsGetActivitiesByInterval4c35**](DrivesFunctionsApi.md#DrivesListItemsGetActivitiesByInterval4c35) | **Get** /drives({drive-id})/list/items({listItem-id})/getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**DrivesRecentA1b8**](DrivesFunctionsApi.md#DrivesRecentA1b8) | **Get** /drives({drive-id})/recent() | Invoke function recent
[**DrivesSearchC788**](DrivesFunctionsApi.md#DrivesSearchC788) | **Get** /drives({drive-id})/search(q&#x3D;{q}) | Invoke function search
[**DrivesSharedWithMeB3e6**](DrivesFunctionsApi.md#DrivesSharedWithMeB3e6) | **Get** /drives({drive-id})/sharedWithMe() | Invoke function sharedWithMe



## DrivesListItemsGetActivitiesByInterval4c35

> []AnyOfmicrosoftGraphItemActivityStat DrivesListItemsGetActivitiesByInterval4c35(ctx, driveId, listItemId)
Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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


## DrivesRecentA1b8

> []AnyOfmicrosoftGraphDriveItem DrivesRecentA1b8(ctx, driveId)
Invoke function recent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 

### Return type

[**[]AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesSearchC788

> []AnyOfmicrosoftGraphDriveItem DrivesSearchC788(ctx, driveId, q)
Invoke function search

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**q** | **string**|  | 

### Return type

[**[]AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesSharedWithMeB3e6

> []AnyOfmicrosoftGraphDriveItem DrivesSharedWithMeB3e6(ctx, driveId)
Invoke function sharedWithMe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 

### Return type

[**[]AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

