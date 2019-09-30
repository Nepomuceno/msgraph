# \DriveListContentTypesColumnLinkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListContentTypesCreateColumnLinks**](DriveListContentTypesColumnLinkApi.md#DriveListContentTypesCreateColumnLinks) | **Post** /drive/list/contentTypes({contentType-id})/columnLinks | Create new navigation property to columnLinks for drive
[**DriveListContentTypesGetColumnLinks**](DriveListContentTypesColumnLinkApi.md#DriveListContentTypesGetColumnLinks) | **Get** /drive/list/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Get columnLinks from drive
[**DriveListContentTypesListColumnLinks**](DriveListContentTypesColumnLinkApi.md#DriveListContentTypesListColumnLinks) | **Get** /drive/list/contentTypes({contentType-id})/columnLinks | Get columnLinks from drive
[**DriveListContentTypesUpdateColumnLinks**](DriveListContentTypesColumnLinkApi.md#DriveListContentTypesUpdateColumnLinks) | **Patch** /drive/list/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Update the navigation property columnLinks in drive



## DriveListContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink DriveListContentTypesCreateColumnLinks(ctx, contentTypeId, microsoftGraphColumnLink)
Create new navigation property to columnLinks for drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)| New navigation property | 

### Return type

[**MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesGetColumnLinks

> MicrosoftGraphColumnLink DriveListContentTypesGetColumnLinks(ctx, contentTypeId, columnLinkId, optional)
Get columnLinks from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
 **optional** | ***DriveListContentTypesGetColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListContentTypesGetColumnLinksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesListColumnLinks

> CollectionOfColumnLink DriveListContentTypesListColumnLinks(ctx, contentTypeId, optional)
Get columnLinks from drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***DriveListContentTypesListColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DriveListContentTypesListColumnLinksOpts struct


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

[**CollectionOfColumnLink**](Collection of columnLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesUpdateColumnLinks

> DriveListContentTypesUpdateColumnLinks(ctx, contentTypeId, columnLinkId, microsoftGraphColumnLink)
Update the navigation property columnLinks in drive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
**microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)| New navigation property values | 

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

