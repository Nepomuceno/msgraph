# \SharesListContentTypeApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListCreateContentTypes**](SharesListContentTypeApi.md#SharesListCreateContentTypes) | **Post** /shares({sharedDriveItem-id})/list/contentTypes | Create new navigation property to contentTypes for shares
[**SharesListGetContentTypes**](SharesListContentTypeApi.md#SharesListGetContentTypes) | **Get** /shares({sharedDriveItem-id})/list/contentTypes({contentType-id}) | Get contentTypes from shares
[**SharesListListContentTypes**](SharesListContentTypeApi.md#SharesListListContentTypes) | **Get** /shares({sharedDriveItem-id})/list/contentTypes | Get contentTypes from shares
[**SharesListUpdateContentTypes**](SharesListContentTypeApi.md#SharesListUpdateContentTypes) | **Patch** /shares({sharedDriveItem-id})/list/contentTypes({contentType-id}) | Update the navigation property contentTypes in shares



## SharesListCreateContentTypes

> MicrosoftGraphContentType SharesListCreateContentTypes(ctx, sharedDriveItemId, microsoftGraphContentType)
Create new navigation property to contentTypes for shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)| New navigation property | 

### Return type

[**MicrosoftGraphContentType**](microsoft.graph.contentType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListGetContentTypes

> MicrosoftGraphContentType SharesListGetContentTypes(ctx, sharedDriveItemId, contentTypeId, optional)
Get contentTypes from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***SharesListGetContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListGetContentTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphContentType**](microsoft.graph.contentType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListListContentTypes

> CollectionOfContentType SharesListListContentTypes(ctx, sharedDriveItemId, optional)
Get contentTypes from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
 **optional** | ***SharesListListContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListListContentTypesOpts struct


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

[**CollectionOfContentType**](Collection of contentType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListUpdateContentTypes

> SharesListUpdateContentTypes(ctx, sharedDriveItemId, contentTypeId, microsoftGraphContentType)
Update the navigation property contentTypes in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)| New navigation property values | 

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

