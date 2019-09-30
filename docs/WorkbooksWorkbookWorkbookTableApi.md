# \WorkbooksWorkbookWorkbookTableApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookCreateTables**](WorkbooksWorkbookWorkbookTableApi.md#WorkbooksWorkbookCreateTables) | **Post** /workbooks({driveItem-id})/workbook/tables | Create new navigation property to tables for workbooks
[**WorkbooksWorkbookGetTables**](WorkbooksWorkbookWorkbookTableApi.md#WorkbooksWorkbookGetTables) | **Get** /workbooks({driveItem-id})/workbook/tables({workbookTable-id}) | Get tables from workbooks
[**WorkbooksWorkbookListTables**](WorkbooksWorkbookWorkbookTableApi.md#WorkbooksWorkbookListTables) | **Get** /workbooks({driveItem-id})/workbook/tables | Get tables from workbooks
[**WorkbooksWorkbookUpdateTables**](WorkbooksWorkbookWorkbookTableApi.md#WorkbooksWorkbookUpdateTables) | **Patch** /workbooks({driveItem-id})/workbook/tables({workbookTable-id}) | Update the navigation property tables in workbooks



## WorkbooksWorkbookCreateTables

> MicrosoftGraphWorkbookTable WorkbooksWorkbookCreateTables(ctx, driveItemId, microsoftGraphWorkbookTable)
Create new navigation property to tables for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksWorkbookGetTables

> MicrosoftGraphWorkbookTable WorkbooksWorkbookGetTables(ctx, driveItemId, workbookTableId, optional)
Get tables from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookTableId** | **string**| key: workbookTable-id of workbookTable | 
 **optional** | ***WorkbooksWorkbookGetTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookGetTablesOpts struct


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


## WorkbooksWorkbookListTables

> CollectionOfWorkbookTable WorkbooksWorkbookListTables(ctx, driveItemId, optional)
Get tables from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksWorkbookListTablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookListTablesOpts struct


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


## WorkbooksWorkbookUpdateTables

> WorkbooksWorkbookUpdateTables(ctx, driveItemId, workbookTableId, microsoftGraphWorkbookTable)
Update the navigation property tables in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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

