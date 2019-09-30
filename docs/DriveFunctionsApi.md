# \DriveFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListItemsGetActivitiesByInterval4c35**](DriveFunctionsApi.md#DriveListItemsGetActivitiesByInterval4c35) | **Get** /drive/list/items({listItem-id})/getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**DriveRecentA1b8**](DriveFunctionsApi.md#DriveRecentA1b8) | **Get** /drive/recent() | Invoke function recent
[**DriveSearchC788**](DriveFunctionsApi.md#DriveSearchC788) | **Get** /drive/search(q&#x3D;{q}) | Invoke function search
[**DriveSharedWithMeB3e6**](DriveFunctionsApi.md#DriveSharedWithMeB3e6) | **Get** /drive/sharedWithMe() | Invoke function sharedWithMe



## DriveListItemsGetActivitiesByInterval4c35

> []AnyOfmicrosoftGraphItemActivityStat DriveListItemsGetActivitiesByInterval4c35(ctx, listItemId)
Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## DriveRecentA1b8

> []AnyOfmicrosoftGraphDriveItem DriveRecentA1b8(ctx, )
Invoke function recent

### Required Parameters

This endpoint does not need any parameter.

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


## DriveSearchC788

> []AnyOfmicrosoftGraphDriveItem DriveSearchC788(ctx, q)
Invoke function search

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## DriveSharedWithMeB3e6

> []AnyOfmicrosoftGraphDriveItem DriveSharedWithMeB3e6(ctx, )
Invoke function sharedWithMe

### Required Parameters

This endpoint does not need any parameter.

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

