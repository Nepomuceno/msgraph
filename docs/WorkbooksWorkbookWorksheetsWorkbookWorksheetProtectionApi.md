# \WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookWorksheetsGetProtection**](WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApi.md#WorkbooksWorkbookWorksheetsGetProtection) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/protection | Get protection from workbooks
[**WorkbooksWorkbookWorksheetsUpdateProtection**](WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApi.md#WorkbooksWorkbookWorksheetsUpdateProtection) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/protection | Update the navigation property protection in workbooks



## WorkbooksWorkbookWorksheetsGetProtection

> MicrosoftGraphWorkbookWorksheetProtection WorkbooksWorkbookWorksheetsGetProtection(ctx, driveItemId, workbookWorksheetId, optional)
Get protection from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
 **optional** | ***WorkbooksWorkbookWorksheetsGetProtectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsGetProtectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookWorksheetProtection**](microsoft.graph.workbookWorksheetProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsUpdateProtection

> WorkbooksWorkbookWorksheetsUpdateProtection(ctx, driveItemId, workbookWorksheetId, microsoftGraphWorkbookWorksheetProtection)
Update the navigation property protection in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**microsoftGraphWorkbookWorksheetProtection** | [**MicrosoftGraphWorkbookWorksheetProtection**](MicrosoftGraphWorkbookWorksheetProtection.md)| New navigation property values | 

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

