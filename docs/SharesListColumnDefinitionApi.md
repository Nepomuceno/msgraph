# \SharesListColumnDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListCreateColumns**](SharesListColumnDefinitionApi.md#SharesListCreateColumns) | **Post** /shares({sharedDriveItem-id})/list/columns | Create new navigation property to columns for shares
[**SharesListGetColumns**](SharesListColumnDefinitionApi.md#SharesListGetColumns) | **Get** /shares({sharedDriveItem-id})/list/columns({columnDefinition-id}) | Get columns from shares
[**SharesListListColumns**](SharesListColumnDefinitionApi.md#SharesListListColumns) | **Get** /shares({sharedDriveItem-id})/list/columns | Get columns from shares
[**SharesListUpdateColumns**](SharesListColumnDefinitionApi.md#SharesListUpdateColumns) | **Patch** /shares({sharedDriveItem-id})/list/columns({columnDefinition-id}) | Update the navigation property columns in shares



## SharesListCreateColumns

> MicrosoftGraphColumnDefinition SharesListCreateColumns(ctx, sharedDriveItemId, microsoftGraphColumnDefinition)
Create new navigation property to columns for shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListGetColumns

> MicrosoftGraphColumnDefinition SharesListGetColumns(ctx, sharedDriveItemId, columnDefinitionId, optional)
Get columns from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
 **optional** | ***SharesListGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListGetColumnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListListColumns

> CollectionOfColumnDefinition SharesListListColumns(ctx, sharedDriveItemId, optional)
Get columns from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesListListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListListColumnsOpts struct


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

[**CollectionOfColumnDefinition**](Collection of columnDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListUpdateColumns

> SharesListUpdateColumns(ctx, sharedDriveItemId, columnDefinitionId, microsoftGraphColumnDefinition)
Update the navigation property columns in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
**microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)| New navigation property values | 

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

