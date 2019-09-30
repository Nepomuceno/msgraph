# \WorkbooksListItemVersionsFieldValueSetApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksListItemVersionsGetFields**](WorkbooksListItemVersionsFieldValueSetApi.md#WorkbooksListItemVersionsGetFields) | **Get** /workbooks({driveItem-id})/listItem/versions({listItemVersion-id})/fields | Get fields from workbooks
[**WorkbooksListItemVersionsUpdateFields**](WorkbooksListItemVersionsFieldValueSetApi.md#WorkbooksListItemVersionsUpdateFields) | **Patch** /workbooks({driveItem-id})/listItem/versions({listItemVersion-id})/fields | Update the navigation property fields in workbooks



## WorkbooksListItemVersionsGetFields

> MicrosoftGraphFieldValueSet WorkbooksListItemVersionsGetFields(ctx, driveItemId, listItemVersionId, optional)
Get fields from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***WorkbooksListItemVersionsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksListItemVersionsGetFieldsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](microsoft.graph.fieldValueSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemVersionsUpdateFields

> WorkbooksListItemVersionsUpdateFields(ctx, driveItemId, listItemVersionId, microsoftGraphFieldValueSet)
Update the navigation property fields in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
**microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)| New navigation property values | 

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

