# \WorkbooksWorkbookWorksheetsWorkbookPivotTableApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookWorksheetsCreatePivotTables**](WorkbooksWorkbookWorksheetsWorkbookPivotTableApi.md#WorkbooksWorkbookWorksheetsCreatePivotTables) | **Post** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/pivotTables | Create new navigation property to pivotTables for workbooks
[**WorkbooksWorkbookWorksheetsGetPivotTables**](WorkbooksWorkbookWorksheetsWorkbookPivotTableApi.md#WorkbooksWorkbookWorksheetsGetPivotTables) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/pivotTables({workbookPivotTable-id}) | Get pivotTables from workbooks
[**WorkbooksWorkbookWorksheetsListPivotTables**](WorkbooksWorkbookWorksheetsWorkbookPivotTableApi.md#WorkbooksWorkbookWorksheetsListPivotTables) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/pivotTables | Get pivotTables from workbooks
[**WorkbooksWorkbookWorksheetsPivotTablesGetWorksheet**](WorkbooksWorkbookWorksheetsWorkbookPivotTableApi.md#WorkbooksWorkbookWorksheetsPivotTablesGetWorksheet) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/pivotTables({workbookPivotTable-id})/worksheet | Get worksheet from workbooks
[**WorkbooksWorkbookWorksheetsPivotTablesUpdateWorksheet**](WorkbooksWorkbookWorksheetsWorkbookPivotTableApi.md#WorkbooksWorkbookWorksheetsPivotTablesUpdateWorksheet) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/pivotTables({workbookPivotTable-id})/worksheet | Update the navigation property worksheet in workbooks
[**WorkbooksWorkbookWorksheetsUpdatePivotTables**](WorkbooksWorkbookWorksheetsWorkbookPivotTableApi.md#WorkbooksWorkbookWorksheetsUpdatePivotTables) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/pivotTables({workbookPivotTable-id}) | Update the navigation property pivotTables in workbooks



## WorkbooksWorkbookWorksheetsCreatePivotTables

> MicrosoftGraphWorkbookPivotTable WorkbooksWorkbookWorksheetsCreatePivotTables(ctx, driveItemId, workbookWorksheetId, microsoftGraphWorkbookPivotTable)
Create new navigation property to pivotTables for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**microsoftGraphWorkbookPivotTable** | [**MicrosoftGraphWorkbookPivotTable**](MicrosoftGraphWorkbookPivotTable.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookPivotTable**](microsoft.graph.workbookPivotTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsGetPivotTables

> MicrosoftGraphWorkbookPivotTable WorkbooksWorkbookWorksheetsGetPivotTables(ctx, driveItemId, workbookWorksheetId, workbookPivotTableId, optional)
Get pivotTables from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookPivotTableId** | **string**| key: workbookPivotTable-id of workbookPivotTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsGetPivotTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsGetPivotTablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookPivotTable**](microsoft.graph.workbookPivotTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsListPivotTables

> CollectionOfWorkbookPivotTable WorkbooksWorkbookWorksheetsListPivotTables(ctx, driveItemId, workbookWorksheetId, optional)
Get pivotTables from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
 **optional** | ***WorkbooksWorkbookWorksheetsListPivotTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsListPivotTablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfWorkbookPivotTable**](Collection of workbookPivotTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsPivotTablesGetWorksheet

> MicrosoftGraphWorkbookWorksheet WorkbooksWorkbookWorksheetsPivotTablesGetWorksheet(ctx, driveItemId, workbookWorksheetId, workbookPivotTableId, optional)
Get worksheet from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookPivotTableId** | **string**| key: workbookPivotTable-id of workbookPivotTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsPivotTablesGetWorksheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsPivotTablesGetWorksheetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookWorksheet**](microsoft.graph.workbookWorksheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsPivotTablesUpdateWorksheet

> WorkbooksWorkbookWorksheetsPivotTablesUpdateWorksheet(ctx, driveItemId, workbookWorksheetId, workbookPivotTableId, microsoftGraphWorkbookWorksheet)
Update the navigation property worksheet in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookPivotTableId** | **string**| key: workbookPivotTable-id of workbookPivotTable | 
**microsoftGraphWorkbookWorksheet** | [**MicrosoftGraphWorkbookWorksheet**](MicrosoftGraphWorkbookWorksheet.md)| New navigation property values | 

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


## WorkbooksWorkbookWorksheetsUpdatePivotTables

> WorkbooksWorkbookWorksheetsUpdatePivotTables(ctx, driveItemId, workbookWorksheetId, workbookPivotTableId, microsoftGraphWorkbookPivotTable)
Update the navigation property pivotTables in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookPivotTableId** | **string**| key: workbookPivotTable-id of workbookPivotTable | 
**microsoftGraphWorkbookPivotTable** | [**MicrosoftGraphWorkbookPivotTable**](MicrosoftGraphWorkbookPivotTable.md)| New navigation property values | 

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

