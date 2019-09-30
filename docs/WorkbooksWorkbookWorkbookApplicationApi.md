# \WorkbooksWorkbookWorkbookApplicationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookGetApplication**](WorkbooksWorkbookWorkbookApplicationApi.md#WorkbooksWorkbookGetApplication) | **Get** /workbooks({driveItem-id})/workbook/application | Get application from workbooks
[**WorkbooksWorkbookUpdateApplication**](WorkbooksWorkbookWorkbookApplicationApi.md#WorkbooksWorkbookUpdateApplication) | **Patch** /workbooks({driveItem-id})/workbook/application | Update the navigation property application in workbooks



## WorkbooksWorkbookGetApplication

> MicrosoftGraphWorkbookApplication WorkbooksWorkbookGetApplication(ctx, driveItemId, optional)
Get application from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksWorkbookGetApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookGetApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookApplication**](microsoft.graph.workbookApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookUpdateApplication

> WorkbooksWorkbookUpdateApplication(ctx, driveItemId, microsoftGraphWorkbookApplication)
Update the navigation property application in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphWorkbookApplication** | [**MicrosoftGraphWorkbookApplication**](MicrosoftGraphWorkbookApplication.md)| New navigation property values | 

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

