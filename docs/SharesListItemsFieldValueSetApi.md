# \SharesListItemsFieldValueSetApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListItemsGetFields**](SharesListItemsFieldValueSetApi.md#SharesListItemsGetFields) | **Get** /shares({sharedDriveItem-id})/list/items({listItem-id})/fields | Get fields from shares
[**SharesListItemsUpdateFields**](SharesListItemsFieldValueSetApi.md#SharesListItemsUpdateFields) | **Patch** /shares({sharedDriveItem-id})/list/items({listItem-id})/fields | Update the navigation property fields in shares



## SharesListItemsGetFields

> MicrosoftGraphFieldValueSet SharesListItemsGetFields(ctx, sharedDriveItemId, listItemId, optional)
Get fields from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***SharesListItemsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemsGetFieldsOpts struct


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


## SharesListItemsUpdateFields

> SharesListItemsUpdateFields(ctx, sharedDriveItemId, listItemId, microsoftGraphFieldValueSet)
Update the navigation property fields in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
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

