# \DriveListListItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListCreateItems**](DriveListListItemApi.md#DriveListCreateItems) | **Post** /drive/list/items | Create new navigation property to items for drive
[**DriveListGetItems**](DriveListListItemApi.md#DriveListGetItems) | **Get** /drive/list/items({listItem-id}) | Get items from drive
[**DriveListListItems**](DriveListListItemApi.md#DriveListListItems) | **Get** /drive/list/items | Get items from drive
[**DriveListUpdateItems**](DriveListListItemApi.md#DriveListUpdateItems) | **Patch** /drive/list/items({listItem-id}) | Update the navigation property items in drive



## DriveListCreateItems

> MicrosoftGraphListItem DriveListCreateItems(ctx, microsoftGraphListItem)
Create new navigation property to items for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphListItem**](microsoft.graph.listItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListGetItems

> MicrosoftGraphListItem DriveListGetItems(ctx, listItemId, optional)
Get items from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***DriveListGetItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphListItem**](microsoft.graph.listItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListListItems

> CollectionOfListItem DriveListListItems(ctx, optional)
Get items from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListListItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListListItemsOpts struct


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

[**CollectionOfListItem**](Collection of listItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListUpdateItems

> DriveListUpdateItems(ctx, listItemId, microsoftGraphListItem)
Update the navigation property items in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string**| key: listItem-id of listItem | 
**microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)| New navigation property values | 

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

