# \DriveListContentTypeApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListCreateContentTypes**](DriveListContentTypeApi.md#DriveListCreateContentTypes) | **Post** /drive/list/contentTypes | Create new navigation property to contentTypes for drive
[**DriveListGetContentTypes**](DriveListContentTypeApi.md#DriveListGetContentTypes) | **Get** /drive/list/contentTypes({contentType-id}) | Get contentTypes from drive
[**DriveListListContentTypes**](DriveListContentTypeApi.md#DriveListListContentTypes) | **Get** /drive/list/contentTypes | Get contentTypes from drive
[**DriveListUpdateContentTypes**](DriveListContentTypeApi.md#DriveListUpdateContentTypes) | **Patch** /drive/list/contentTypes({contentType-id}) | Update the navigation property contentTypes in drive



## DriveListCreateContentTypes

> MicrosoftGraphContentType DriveListCreateContentTypes(ctx, microsoftGraphContentType)
Create new navigation property to contentTypes for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## DriveListGetContentTypes

> MicrosoftGraphContentType DriveListGetContentTypes(ctx, contentTypeId, optional)
Get contentTypes from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***DriveListGetContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListGetContentTypesOpts struct


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


## DriveListListContentTypes

> CollectionOfContentType DriveListListContentTypes(ctx, optional)
Get contentTypes from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriveListListContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListListContentTypesOpts struct


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


## DriveListUpdateContentTypes

> DriveListUpdateContentTypes(ctx, contentTypeId, microsoftGraphContentType)
Update the navigation property contentTypes in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

