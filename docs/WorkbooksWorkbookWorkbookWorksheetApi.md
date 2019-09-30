# \WorkbooksWorkbookWorkbookWorksheetApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookCreateWorksheets**](WorkbooksWorkbookWorkbookWorksheetApi.md#WorkbooksWorkbookCreateWorksheets) | **Post** /workbooks({driveItem-id})/workbook/worksheets | Create new navigation property to worksheets for workbooks
[**WorkbooksWorkbookGetWorksheets**](WorkbooksWorkbookWorkbookWorksheetApi.md#WorkbooksWorkbookGetWorksheets) | **Get** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id}) | Get worksheets from workbooks
[**WorkbooksWorkbookListWorksheets**](WorkbooksWorkbookWorkbookWorksheetApi.md#WorkbooksWorkbookListWorksheets) | **Get** /workbooks({driveItem-id})/workbook/worksheets | Get worksheets from workbooks
[**WorkbooksWorkbookUpdateWorksheets**](WorkbooksWorkbookWorkbookWorksheetApi.md#WorkbooksWorkbookUpdateWorksheets) | **Patch** /workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id}) | Update the navigation property worksheets in workbooks



## WorkbooksWorkbookCreateWorksheets

> MicrosoftGraphWorkbookWorksheet WorkbooksWorkbookCreateWorksheets(ctx, driveItemId, microsoftGraphWorkbookWorksheet)
Create new navigation property to worksheets for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphWorkbookWorksheet** | [**MicrosoftGraphWorkbookWorksheet**](MicrosoftGraphWorkbookWorksheet.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookWorksheet**](microsoft.graph.workbookWorksheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookGetWorksheets

> MicrosoftGraphWorkbookWorksheet WorkbooksWorkbookGetWorksheets(ctx, driveItemId, workbookWorksheetId, optional)
Get worksheets from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
 **optional** | ***WorkbooksWorkbookGetWorksheetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookGetWorksheetsOpts struct


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


## WorkbooksWorkbookListWorksheets

> CollectionOfWorkbookWorksheet WorkbooksWorkbookListWorksheets(ctx, driveItemId, optional)
Get worksheets from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksWorkbookListWorksheetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookListWorksheetsOpts struct


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

[**CollectionOfWorkbookWorksheet**](Collection of workbookWorksheet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookUpdateWorksheets

> WorkbooksWorkbookUpdateWorksheets(ctx, driveItemId, workbookWorksheetId, microsoftGraphWorkbookWorksheet)
Update the navigation property worksheets in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookWorksheetId** | **string**| key: workbookWorksheet-id of workbookWorksheet | 
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

