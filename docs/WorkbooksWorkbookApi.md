# \WorkbooksWorkbookApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksGetWorkbook**](WorkbooksWorkbookApi.md#WorkbooksGetWorkbook) | **Get** /workbooks({driveItem-id})/workbook | Get workbook from workbooks
[**WorkbooksUpdateWorkbook**](WorkbooksWorkbookApi.md#WorkbooksUpdateWorkbook) | **Patch** /workbooks({driveItem-id})/workbook | Update the navigation property workbook in workbooks



## WorkbooksGetWorkbook

> MicrosoftGraphWorkbook WorkbooksGetWorkbook(ctx, driveItemId, optional)
Get workbook from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksGetWorkbookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksGetWorkbookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbook**](microsoft.graph.workbook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateWorkbook

> WorkbooksUpdateWorkbook(ctx, driveItemId, microsoftGraphWorkbook)
Update the navigation property workbook in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphWorkbook** | [**MicrosoftGraphWorkbook**](MicrosoftGraphWorkbook.md)| New navigation property values | 

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

