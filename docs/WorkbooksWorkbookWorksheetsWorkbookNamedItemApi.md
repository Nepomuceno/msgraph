# \WorkbooksWorkbookWorksheetsWorkbookNamedItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookWorksheetsCreateNames**](WorkbooksWorkbookWorksheetsWorkbookNamedItemApi.md#WorkbooksWorkbookWorksheetsCreateNames) | **Post** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/names | Create new navigation property to names for workbooks
[**WorkbooksWorkbookWorksheetsGetNames**](WorkbooksWorkbookWorksheetsWorkbookNamedItemApi.md#WorkbooksWorkbookWorksheetsGetNames) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/names({workbookNamedItem-id}) | Get names from workbooks
[**WorkbooksWorkbookWorksheetsListNames**](WorkbooksWorkbookWorksheetsWorkbookNamedItemApi.md#WorkbooksWorkbookWorksheetsListNames) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/names | Get names from workbooks
[**WorkbooksWorkbookWorksheetsNamesGetWorksheet**](WorkbooksWorkbookWorksheetsWorkbookNamedItemApi.md#WorkbooksWorkbookWorksheetsNamesGetWorksheet) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/names({workbookNamedItem-id})/worksheet | Get worksheet from workbooks
[**WorkbooksWorkbookWorksheetsNamesUpdateWorksheet**](WorkbooksWorkbookWorksheetsWorkbookNamedItemApi.md#WorkbooksWorkbookWorksheetsNamesUpdateWorksheet) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/names({workbookNamedItem-id})/worksheet | Update the navigation property worksheet in workbooks
[**WorkbooksWorkbookWorksheetsUpdateNames**](WorkbooksWorkbookWorksheetsWorkbookNamedItemApi.md#WorkbooksWorkbookWorksheetsUpdateNames) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/names({workbookNamedItem-id}) | Update the navigation property names in workbooks



## WorkbooksWorkbookWorksheetsCreateNames

> MicrosoftGraphWorkbookNamedItem WorkbooksWorkbookWorksheetsCreateNames(ctx, driveItemId, workbookWorksheetId, microsoftGraphWorkbookNamedItem)
Create new navigation property to names for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**microsoftGraphWorkbookNamedItem** | [**MicrosoftGraphWorkbookNamedItem**](MicrosoftGraphWorkbookNamedItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookNamedItem**](microsoft.graph.workbookNamedItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsGetNames

> MicrosoftGraphWorkbookNamedItem WorkbooksWorkbookWorksheetsGetNames(ctx, driveItemId, workbookWorksheetId, workbookNamedItemId, optional)
Get names from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookNamedItemId** | **string**| key: workbookNamedItem-id of workbookNamedItem | 
 **optional** | ***WorkbooksWorkbookWorksheetsGetNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsGetNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookNamedItem**](microsoft.graph.workbookNamedItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsListNames

> CollectionOfWorkbookNamedItem WorkbooksWorkbookWorksheetsListNames(ctx, driveItemId, workbookWorksheetId, optional)
Get names from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
 **optional** | ***WorkbooksWorkbookWorksheetsListNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsListNamesOpts struct


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

[**CollectionOfWorkbookNamedItem**](Collection of workbookNamedItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookWorksheetsNamesGetWorksheet

> MicrosoftGraphWorkbookWorksheet WorkbooksWorkbookWorksheetsNamesGetWorksheet(ctx, driveItemId, workbookWorksheetId, workbookNamedItemId, optional)
Get worksheet from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookNamedItemId** | **string**| key: workbookNamedItem-id of workbookNamedItem | 
 **optional** | ***WorkbooksWorkbookWorksheetsNamesGetWorksheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookWorksheetsNamesGetWorksheetOpts struct


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


## WorkbooksWorkbookWorksheetsNamesUpdateWorksheet

> WorkbooksWorkbookWorksheetsNamesUpdateWorksheet(ctx, driveItemId, workbookWorksheetId, workbookNamedItemId, microsoftGraphWorkbookWorksheet)
Update the navigation property worksheet in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookNamedItemId** | **string**| key: workbookNamedItem-id of workbookNamedItem | 
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


## WorkbooksWorkbookWorksheetsUpdateNames

> WorkbooksWorkbookWorksheetsUpdateNames(ctx, driveItemId, workbookWorksheetId, workbookNamedItemId, microsoftGraphWorkbookNamedItem)
Update the navigation property names in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
**workbookNamedItemId** | **string**| key: workbookNamedItem-id of workbookNamedItem | 
**microsoftGraphWorkbookNamedItem** | [**MicrosoftGraphWorkbookNamedItem**](MicrosoftGraphWorkbookNamedItem.md)| New navigation property values | 

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

