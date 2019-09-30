# \DrivesListColumnDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListCreateColumns**](DrivesListColumnDefinitionApi.md#DrivesListCreateColumns) | **Post** /drives({drive-id})/list/columns | Create new navigation property to columns for drives
[**DrivesListGetColumns**](DrivesListColumnDefinitionApi.md#DrivesListGetColumns) | **Get** /drives({drive-id})/list/columns({columnDefinition-id}) | Get columns from drives
[**DrivesListListColumns**](DrivesListColumnDefinitionApi.md#DrivesListListColumns) | **Get** /drives({drive-id})/list/columns | Get columns from drives
[**DrivesListUpdateColumns**](DrivesListColumnDefinitionApi.md#DrivesListUpdateColumns) | **Patch** /drives({drive-id})/list/columns({columnDefinition-id}) | Update the navigation property columns in drives



## DrivesListCreateColumns

> MicrosoftGraphColumnDefinition DrivesListCreateColumns(ctx, driveId, microsoftGraphColumnDefinition)
Create new navigation property to columns for drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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


## DrivesListGetColumns

> MicrosoftGraphColumnDefinition DrivesListGetColumns(ctx, driveId, columnDefinitionId, optional)
Get columns from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**columnDefinitionId** | **string**| key: columnDefinition-id of columnDefinition | 
 **optional** | ***DrivesListGetColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListGetColumnsOpts struct


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


## DrivesListListColumns

> CollectionOfColumnDefinition DrivesListListColumns(ctx, driveId, optional)
Get columns from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesListListColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListListColumnsOpts struct


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


## DrivesListUpdateColumns

> DrivesListUpdateColumns(ctx, driveId, columnDefinitionId, microsoftGraphColumnDefinition)
Update the navigation property columns in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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

