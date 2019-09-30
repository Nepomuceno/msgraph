# \WorkbooksWorkbookTablesWorkbookTableColumnApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookTablesColumnsGetFilter**](WorkbooksWorkbookTablesWorkbookTableColumnApi.md#WorkbooksWorkbookTablesColumnsGetFilter) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/columns({workbookTableColumn-id})/filter | Get filter from workbooks
[**WorkbooksWorkbookTablesColumnsUpdateFilter**](WorkbooksWorkbookTablesWorkbookTableColumnApi.md#WorkbooksWorkbookTablesColumnsUpdateFilter) | **Patch** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/columns({workbookTableColumn-id})/filter | Update the navigation property filter in workbooks
[**WorkbooksWorkbookTablesCreateColumns**](WorkbooksWorkbookTablesWorkbookTableColumnApi.md#WorkbooksWorkbookTablesCreateColumns) | **Post** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/columns | Create new navigation property to columns for workbooks
[**WorkbooksWorkbookTablesGetColumns**](WorkbooksWorkbookTablesWorkbookTableColumnApi.md#WorkbooksWorkbookTablesGetColumns) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/columns({workbookTableColumn-id}) | Get columns from workbooks
[**WorkbooksWorkbookTablesListColumns**](WorkbooksWorkbookTablesWorkbookTableColumnApi.md#WorkbooksWorkbookTablesListColumns) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/columns | Get columns from workbooks
[**WorkbooksWorkbookTablesUpdateColumns**](WorkbooksWorkbookTablesWorkbookTableColumnApi.md#WorkbooksWorkbookTablesUpdateColumns) | **Patch** /workbooks({driveItem-id})/workbook/tables({workbookTable-id})/columns({workbookTableColumn-id}) | Update the navigation property columns in workbooks



## WorkbooksWorkbookTablesColumnsGetFilter

> MicrosoftGraphWorkbookFilter WorkbooksWorkbookTablesColumnsGetFilter(ctx, driveItemId, workbookTableId, workbookTableColumnId, optional)
Get filter from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableColumnId** | **string**| key: workbookTableColumn-id of workbookTableColumn | 
 **optional** | ***WorkbooksWorkbookTablesColumnsGetFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookTablesColumnsGetFilterOpts struct


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


## WorkbooksWorkbookTablesColumnsUpdateFilter

> WorkbooksWorkbookTablesColumnsUpdateFilter(ctx, driveItemId, workbookTableId, workbookTableColumnId, microsoftGraphWorkbookFilter)
Update the navigation property filter in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksWorkbookTablesCreateColumns

> MicrosoftGraphWorkbookTableColumn WorkbooksWorkbookTablesCreateColumns(ctx, driveItemId, workbookTableId, microsoftGraphWorkbookTableColumn)
Create new navigation property to columns for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksWorkbookTablesGetColumns

> MicrosoftGraphWorkbookTableColumn WorkbooksWorkbookTablesGetColumns(ctx, driveItemId, workbookTableId, workbookTableColumnId, optional)
Get columns from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
**workbookTableColumnId** | **string**| key: workbookTableColumn-id of workbookTableColumn | 
 **optional** | ***WorkbooksWorkbookTablesGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookTablesGetColumnsOpts struct


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


## WorkbooksWorkbookTablesListColumns

> CollectionOfWorkbookTableColumn WorkbooksWorkbookTablesListColumns(ctx, driveItemId, workbookTableId, optional)
Get columns from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookTablesListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookTablesListColumnsOpts struct


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


## WorkbooksWorkbookTablesUpdateColumns

> WorkbooksWorkbookTablesUpdateColumns(ctx, driveItemId, workbookTableId, workbookTableColumnId, microsoftGraphWorkbookTableColumn)
Update the navigation property columns in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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

