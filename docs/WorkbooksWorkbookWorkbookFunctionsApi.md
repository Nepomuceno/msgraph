# \WorkbooksWorkbookWorkbookFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookGetFunctions**](WorkbooksWorkbookWorkbookFunctionsApi.md#WorkbooksWorkbookGetFunctions) | **Get** /workbooks({driveItem-id})/workbook/functions | Get functions from workbooks
[**WorkbooksWorkbookUpdateFunctions**](WorkbooksWorkbookWorkbookFunctionsApi.md#WorkbooksWorkbookUpdateFunctions) | **Patch** /workbooks({driveItem-id})/workbook/functions | Update the navigation property functions in workbooks



## WorkbooksWorkbookGetFunctions

> MicrosoftGraphWorkbookFunctions WorkbooksWorkbookGetFunctions(ctx, driveItemId, optional)
Get functions from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksWorkbookGetFunctionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookGetFunctionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookFunctions**](microsoft.graph.workbookFunctions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookUpdateFunctions

> WorkbooksWorkbookUpdateFunctions(ctx, driveItemId, microsoftGraphWorkbookFunctions)
Update the navigation property functions in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphWorkbookFunctions** | [**MicrosoftGraphWorkbookFunctions**](MicrosoftGraphWorkbookFunctions.md)| New navigation property values | 

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

