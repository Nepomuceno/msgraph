# \WorkbooksWorkbookTablesWorkbookTableRowApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookTablesCreateRows**](WorkbooksWorkbookTablesWorkbookTableRowApi.md#WorkbooksWorkbookTablesCreateRows) | **Post** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/rows | Create new navigation property to rows for workbooks
[**WorkbooksWorkbookTablesGetRows**](WorkbooksWorkbookTablesWorkbookTableRowApi.md#WorkbooksWorkbookTablesGetRows) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/rows({workbookTableRow-id}) | Get rows from workbooks
[**WorkbooksWorkbookTablesListRows**](WorkbooksWorkbookTablesWorkbookTableRowApi.md#WorkbooksWorkbookTablesListRows) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/rows | Get rows from workbooks
[**WorkbooksWorkbookTablesUpdateRows**](WorkbooksWorkbookTablesWorkbookTableRowApi.md#WorkbooksWorkbookTablesUpdateRows) | **Patch** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/rows({workbookTableRow-id}) | Update the navigation property rows in workbooks



## WorkbooksWorkbookTablesCreateRows

> MicrosoftGraphWorkbookTableRow WorkbooksWorkbookTablesCreateRows(ctx, driveItemId, workbookTableId, microsoftGraphWorkbookTableRow)
Create new navigation property to rows for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksWorkbookTablesGetRows

> MicrosoftGraphWorkbookTableRow WorkbooksWorkbookTablesGetRows(ctx, driveItemId, workbookTableId, workbookTableRowId, optional)
Get rows from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableRowId** | **string**| key: workbookTableRow-id of workbookTableRow | 
 **optional** | ***WorkbooksWorkbookTablesGetRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookTablesGetRowsOpts struct


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


## WorkbooksWorkbookTablesListRows

> CollectionOfWorkbookTableRow WorkbooksWorkbookTablesListRows(ctx, driveItemId, workbookTableId, optional)
Get rows from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookTablesListRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookTablesListRowsOpts struct


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


## WorkbooksWorkbookTablesUpdateRows

> WorkbooksWorkbookTablesUpdateRows(ctx, driveItemId, workbookTableId, workbookTableRowId, microsoftGraphWorkbookTableRow)
Update the navigation property rows in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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

