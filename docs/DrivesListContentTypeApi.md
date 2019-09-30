# \DrivesListContentTypeApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListCreateContentTypes**](DrivesListContentTypeApi.md#DrivesListCreateContentTypes) | **Post** /drives({drive-id})/list/contentTypes | Create new navigation property to contentTypes for drives
[**DrivesListGetContentTypes**](DrivesListContentTypeApi.md#DrivesListGetContentTypes) | **Get** /drives({drive-id})/list/contentTypes({contentType-id}) | Get contentTypes from drives
[**DrivesListListContentTypes**](DrivesListContentTypeApi.md#DrivesListListContentTypes) | **Get** /drives({drive-id})/list/contentTypes | Get contentTypes from drives
[**DrivesListUpdateContentTypes**](DrivesListContentTypeApi.md#DrivesListUpdateContentTypes) | **Patch** /drives({drive-id})/list/contentTypes({contentType-id}) | Update the navigation property contentTypes in drives



## DrivesListCreateContentTypes

> MicrosoftGraphContentType DrivesListCreateContentTypes(ctx, driveId, microsoftGraphContentType)
Create new navigation property to contentTypes for drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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


## DrivesListGetContentTypes

> MicrosoftGraphContentType DrivesListGetContentTypes(ctx, driveId, contentTypeId, optional)
Get contentTypes from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***DrivesListGetContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListGetContentTypesOpts struct


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


## DrivesListListContentTypes

> CollectionOfContentType DrivesListListContentTypes(ctx, driveId, optional)
Get contentTypes from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
 **optional** | ***DrivesListListContentTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListListContentTypesOpts struct


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


## DrivesListUpdateContentTypes

> DrivesListUpdateContentTypes(ctx, driveId, contentTypeId, microsoftGraphContentType)
Update the navigation property contentTypes in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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

