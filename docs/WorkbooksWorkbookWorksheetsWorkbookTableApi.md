# \WorkbooksWorkbookWorksheetsWorkbookTableApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookWorksheetsCreateTables**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsCreateTables) | **Post** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables | Create new navigation property to tables for workbooks
[**WorkbooksWorkbookWorksheetsGetTables**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsGetTables) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id}) | Get tables from workbooks
[**WorkbooksWorkbookWorksheetsListTables**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsListTables) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables | Get tables from workbooks
[**WorkbooksWorkbookWorksheetsTablesColumnsGetFilter**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesColumnsGetFilter) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/columns({workbookTableColumn-id})/filter | Get filter from workbooks
[**WorkbooksWorkbookWorksheetsTablesColumnsUpdateFilter**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesColumnsUpdateFilter) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/columns({workbookTableColumn-id})/filter | Update the navigation property filter in workbooks
[**WorkbooksWorkbookWorksheetsTablesCreateColumns**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesCreateColumns) | **Post** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/columns | Create new navigation property to columns for workbooks
[**WorkbooksWorkbookWorksheetsTablesCreateRows**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesCreateRows) | **Post** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/rows | Create new navigation property to rows for workbooks
[**WorkbooksWorkbookWorksheetsTablesGetColumns**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesGetColumns) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/columns({workbookTableColumn-id}) | Get columns from workbooks
[**WorkbooksWorkbookWorksheetsTablesGetRows**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesGetRows) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/rows({workbookTableRow-id}) | Get rows from workbooks
[**WorkbooksWorkbookWorksheetsTablesGetSort**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesGetSort) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/sort | Get sort from workbooks
[**WorkbooksWorkbookWorksheetsTablesGetWorksheet**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesGetWorksheet) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/worksheet | Get worksheet from workbooks
[**WorkbooksWorkbookWorksheetsTablesListColumns**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesListColumns) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/columns | Get columns from workbooks
[**WorkbooksWorkbookWorksheetsTablesListRows**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesListRows) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/rows | Get rows from workbooks
[**WorkbooksWorkbookWorksheetsTablesUpdateColumns**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesUpdateColumns) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/columns({workbookTableColumn-id}) | Update the navigation property columns in workbooks
[**WorkbooksWorkbookWorksheetsTablesUpdateRows**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesUpdateRows) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/rows({workbookTableRow-id}) | Update the navigation property rows in workbooks
[**WorkbooksWorkbookWorksheetsTablesUpdateSort**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesUpdateSort) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/sort | Update the navigation property sort in workbooks
[**WorkbooksWorkbookWorksheetsTablesUpdateWorksheet**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsTablesUpdateWorksheet) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id})/worksheet | Update the navigation property worksheet in workbooks
[**WorkbooksWorkbookWorksheetsUpdateTables**](WorkbooksWorkbookWorksheetsWorkbookTableApi.md#WorkbooksWorkbookWorksheetsUpdateTables) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/tables({workbookTable-id}) | Update the navigation property tables in workbooks



## WorkbooksWorkbookWorksheetsCreateTables

> MicrosoftGraphWorkbookTable WorkbooksWorkbookWorksheetsCreateTables(ctx, driveItemId, workbookWorksheetId, microsoftGraphWorkbookTable)
Create new navigation property to tables for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**microsoftGraphWorkbookTable** | [**MicrosoftGraphWorkbookTable**](MicrosoftGraphWorkbookTable.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookTable**](microsoft.graph.workbookTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsGetTables

> MicrosoftGraphWorkbookTable WorkbooksWorkbookWorksheetsGetTables(ctx, driveItemId, workbookWorksheetId, workbookTableId, optional)
Get tables from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsGetTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsGetTablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookTable**](microsoft.graph.workbookTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsListTables

> CollectionOfWorkbookTable WorkbooksWorkbookWorksheetsListTables(ctx, driveItemId, workbookWorksheetId, optional)
Get tables from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
 **optional** | ***WorkbooksWorkbookWorksheetsListTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsListTablesOpts struct


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

[**CollectionOfWorkbookTable**](Collection of workbookTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesColumnsGetFilter

> MicrosoftGraphWorkbookFilter WorkbooksWorkbookWorksheetsTablesColumnsGetFilter(ctx, driveItemId, workbookWorksheetId, workbookTableId, workbookTableColumnId, optional)
Get filter from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableColumnId** | **string**| key: workbookTableColumn-id of workbookTableColumn | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesColumnsGetFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesColumnsGetFilterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookFilter**](microsoft.graph.workbookFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesColumnsUpdateFilter

> WorkbooksWorkbookWorksheetsTablesColumnsUpdateFilter(ctx, driveItemId, workbookWorksheetId, workbookTableId, workbookTableColumnId, microsoftGraphWorkbookFilter)
Update the navigation property filter in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableColumnId** | **string**| key: workbookTableColumn-id of workbookTableColumn | 
**microsoftGraphWorkbookFilter** | [**MicrosoftGraphWorkbookFilter**](MicrosoftGraphWorkbookFilter.md)| New navigation property values | 

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


## WorkbooksWorkbookWorksheetsTablesCreateColumns

> MicrosoftGraphWorkbookTableColumn WorkbooksWorkbookWorksheetsTablesCreateColumns(ctx, driveItemId, workbookWorksheetId, workbookTableId, microsoftGraphWorkbookTableColumn)
Create new navigation property to columns for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**microsoftGraphWorkbookTableColumn** | [**MicrosoftGraphWorkbookTableColumn**](MicrosoftGraphWorkbookTableColumn.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookTableColumn**](microsoft.graph.workbookTableColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesCreateRows

> MicrosoftGraphWorkbookTableRow WorkbooksWorkbookWorksheetsTablesCreateRows(ctx, driveItemId, workbookWorksheetId, workbookTableId, microsoftGraphWorkbookTableRow)
Create new navigation property to rows for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**microsoftGraphWorkbookTableRow** | [**MicrosoftGraphWorkbookTableRow**](MicrosoftGraphWorkbookTableRow.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookTableRow**](microsoft.graph.workbookTableRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesGetColumns

> MicrosoftGraphWorkbookTableColumn WorkbooksWorkbookWorksheetsTablesGetColumns(ctx, driveItemId, workbookWorksheetId, workbookTableId, workbookTableColumnId, optional)
Get columns from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableColumnId** | **string**| key: workbookTableColumn-id of workbookTableColumn | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesGetColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookTableColumn**](microsoft.graph.workbookTableColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesGetRows

> MicrosoftGraphWorkbookTableRow WorkbooksWorkbookWorksheetsTablesGetRows(ctx, driveItemId, workbookWorksheetId, workbookTableId, workbookTableRowId, optional)
Get rows from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableRowId** | **string**| key: workbookTableRow-id of workbookTableRow | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesGetRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesGetRowsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookTableRow**](microsoft.graph.workbookTableRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesGetSort

> MicrosoftGraphWorkbookTableSort WorkbooksWorkbookWorksheetsTablesGetSort(ctx, driveItemId, workbookWorksheetId, workbookTableId, optional)
Get sort from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesGetSortOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesGetSortOpts struct


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


## WorkbooksWorkbookWorksheetsTablesGetWorksheet

> MicrosoftGraphWorkbookWorksheet WorkbooksWorkbookWorksheetsTablesGetWorksheet(ctx, driveItemId, workbookWorksheetId, workbookTableId, optional)
Get worksheet from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesGetWorksheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesGetWorksheetOpts struct


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


## WorkbooksWorkbookWorksheetsTablesListColumns

> CollectionOfWorkbookTableColumn WorkbooksWorkbookWorksheetsTablesListColumns(ctx, driveItemId, workbookWorksheetId, workbookTableId, optional)
Get columns from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesListColumnsOpts struct


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

[**CollectionOfWorkbookTableColumn**](Collection of workbookTableColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesListRows

> CollectionOfWorkbookTableRow WorkbooksWorkbookWorksheetsTablesListRows(ctx, driveItemId, workbookWorksheetId, workbookTableId, optional)
Get rows from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookWorksheetsTablesListRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsTablesListRowsOpts struct


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

[**CollectionOfWorkbookTableRow**](Collection of workbookTableRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsTablesUpdateColumns

> WorkbooksWorkbookWorksheetsTablesUpdateColumns(ctx, driveItemId, workbookWorksheetId, workbookTableId, workbookTableColumnId, microsoftGraphWorkbookTableColumn)
Update the navigation property columns in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableColumnId** | **string**| key: workbookTableColumn-id of workbookTableColumn | 
**microsoftGraphWorkbookTableColumn** | [**MicrosoftGraphWorkbookTableColumn**](MicrosoftGraphWorkbookTableColumn.md)| New navigation property values | 

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


## WorkbooksWorkbookWorksheetsTablesUpdateRows

> WorkbooksWorkbookWorksheetsTablesUpdateRows(ctx, driveItemId, workbookWorksheetId, workbookTableId, workbookTableRowId, microsoftGraphWorkbookTableRow)
Update the navigation property rows in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableRowId** | **string**| key: workbookTableRow-id of workbookTableRow | 
**microsoftGraphWorkbookTableRow** | [**MicrosoftGraphWorkbookTableRow**](MicrosoftGraphWorkbookTableRow.md)| New navigation property values | 

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


## WorkbooksWorkbookWorksheetsTablesUpdateSort

> WorkbooksWorkbookWorksheetsTablesUpdateSort(ctx, driveItemId, workbookWorksheetId, workbookTableId, microsoftGraphWorkbookTableSort)
Update the navigation property sort in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
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


## WorkbooksWorkbookWorksheetsTablesUpdateWorksheet

> WorkbooksWorkbookWorksheetsTablesUpdateWorksheet(ctx, driveItemId, workbookWorksheetId, workbookTableId, microsoftGraphWorkbookWorksheet)
Update the navigation property worksheet in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
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


## WorkbooksWorkbookWorksheetsUpdateTables

> WorkbooksWorkbookWorksheetsUpdateTables(ctx, driveItemId, workbookWorksheetId, workbookTableId, microsoftGraphWorkbookTable)
Update the navigation property tables in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**microsoftGraphWorkbookTable** | [**MicrosoftGraphWorkbookTable**](MicrosoftGraphWorkbookTable.md)| New navigation property values | 

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

