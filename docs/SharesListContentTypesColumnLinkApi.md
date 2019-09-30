# \SharesListContentTypesColumnLinkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListContentTypesCreateColumnLinks**](SharesListContentTypesColumnLinkApi.md#SharesListContentTypesCreateColumnLinks) | **Post** /shares({sharedDriveItem-id})/list/contentTypes({contentType-id})/columnLinks | Create new navigation property to columnLinks for shares
[**SharesListContentTypesGetColumnLinks**](SharesListContentTypesColumnLinkApi.md#SharesListContentTypesGetColumnLinks) | **Get** /shares({sharedDriveItem-id})/list/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Get columnLinks from shares
[**SharesListContentTypesListColumnLinks**](SharesListContentTypesColumnLinkApi.md#SharesListContentTypesListColumnLinks) | **Get** /shares({sharedDriveItem-id})/list/contentTypes({contentType-id})/columnLinks | Get columnLinks from shares
[**SharesListContentTypesUpdateColumnLinks**](SharesListContentTypesColumnLinkApi.md#SharesListContentTypesUpdateColumnLinks) | **Patch** /shares({sharedDriveItem-id})/list/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Update the navigation property columnLinks in shares



## SharesListContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink SharesListContentTypesCreateColumnLinks(ctx, sharedDriveItemId, contentTypeId, microsoftGraphColumnLink)
Create new navigation property to columnLinks for shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
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


## SharesListContentTypesGetColumnLinks

> MicrosoftGraphColumnLink SharesListContentTypesGetColumnLinks(ctx, sharedDriveItemId, contentTypeId, columnLinkId, optional)
Get columnLinks from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
 **optional** | ***SharesListContentTypesGetColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListContentTypesGetColumnLinksOpts struct


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


## SharesListContentTypesListColumnLinks

> CollectionOfColumnLink SharesListContentTypesListColumnLinks(ctx, sharedDriveItemId, contentTypeId, optional)
Get columnLinks from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***SharesListContentTypesListColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListContentTypesListColumnLinksOpts struct


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


## SharesListContentTypesUpdateColumnLinks

> SharesListContentTypesUpdateColumnLinks(ctx, sharedDriveItemId, contentTypeId, columnLinkId, microsoftGraphColumnLink)
Update the navigation property columnLinks in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
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

