# \WorkbooksWorkbookWorkbookNamedItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookCreateNames**](WorkbooksWorkbookWorkbookNamedItemApi.md#WorkbooksWorkbookCreateNames) | **Post** /workbooks({driveItem-id})/workbook/names | Create new navigation property to names for workbooks
[**WorkbooksWorkbookGetNames**](WorkbooksWorkbookWorkbookNamedItemApi.md#WorkbooksWorkbookGetNames) | **Get** /workbooks({driveItem-id})/workbook/names({workbookNamedItem-id}) | Get names from workbooks
[**WorkbooksWorkbookListNames**](WorkbooksWorkbookWorkbookNamedItemApi.md#WorkbooksWorkbookListNames) | **Get** /workbooks({driveItem-id})/workbook/names | Get names from workbooks
[**WorkbooksWorkbookUpdateNames**](WorkbooksWorkbookWorkbookNamedItemApi.md#WorkbooksWorkbookUpdateNames) | **Patch** /workbooks({driveItem-id})/workbook/names({workbookNamedItem-id}) | Update the navigation property names in workbooks



## WorkbooksWorkbookCreateNames

> MicrosoftGraphWorkbookNamedItem WorkbooksWorkbookCreateNames(ctx, driveItemId, microsoftGraphWorkbookNamedItem)
Create new navigation property to names for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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


## WorkbooksWorkbookGetNames

> MicrosoftGraphWorkbookNamedItem WorkbooksWorkbookGetNames(ctx, driveItemId, workbookNamedItemId, optional)
Get names from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookNamedItemId** | **string**| key: workbookNamedItem-id of workbookNamedItem | 
 **optional** | ***WorkbooksWorkbookGetNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookGetNamesOpts struct


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


## WorkbooksWorkbookListNames

> CollectionOfWorkbookNamedItem WorkbooksWorkbookListNames(ctx, driveItemId, optional)
Get names from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksWorkbookListNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookListNamesOpts struct


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


## WorkbooksWorkbookUpdateNames

> WorkbooksWorkbookUpdateNames(ctx, driveItemId, workbookNamedItemId, microsoftGraphWorkbookNamedItem)
Update the navigation property names in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
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

