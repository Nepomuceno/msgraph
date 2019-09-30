# \WorkbooksWorkbookTablesWorkbookTableSortApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookTablesGetSort**](WorkbooksWorkbookTablesWorkbookTableSortApi.md#WorkbooksWorkbookTablesGetSort) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/sort | Get sort from workbooks
[**WorkbooksWorkbookTablesUpdateSort**](WorkbooksWorkbookTablesWorkbookTableSortApi.md#WorkbooksWorkbookTablesUpdateSort) | **Patch** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/sort | Update the navigation property sort in workbooks



## WorkbooksWorkbookTablesGetSort

> MicrosoftGraphWorkbookTableSort WorkbooksWorkbookTablesGetSort(ctx, driveItemId, workbookTableId, optional)
Get sort from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookTablesGetSortOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookTablesGetSortOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookTableSort**](microsoft.graph.workbookTableSort.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookTablesUpdateSort

> WorkbooksWorkbookTablesUpdateSort(ctx, driveItemId, workbookTableId, microsoftGraphWorkbookTableSort)
Update the navigation property sort in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**microsoftGraphWorkbookTableSort** | [**MicrosoftGraphWorkbookTableSort**](MicrosoftGraphWorkbookTableSort.md)| New navigation property values | 

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

