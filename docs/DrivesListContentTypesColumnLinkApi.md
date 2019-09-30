# \DrivesListContentTypesColumnLinkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesListContentTypesCreateColumnLinks**](DrivesListContentTypesColumnLinkApi.md#DrivesListContentTypesCreateColumnLinks) | **Post** /drives({drive-id})/list/contentTypes({contentType-id})/columnLinks | Create new navigation property to columnLinks for drives
[**DrivesListContentTypesGetColumnLinks**](DrivesListContentTypesColumnLinkApi.md#DrivesListContentTypesGetColumnLinks) | **Get** /drives({drive-id})/list/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Get columnLinks from drives
[**DrivesListContentTypesListColumnLinks**](DrivesListContentTypesColumnLinkApi.md#DrivesListContentTypesListColumnLinks) | **Get** /drives({drive-id})/list/contentTypes({contentType-id})/columnLinks | Get columnLinks from drives
[**DrivesListContentTypesUpdateColumnLinks**](DrivesListContentTypesColumnLinkApi.md#DrivesListContentTypesUpdateColumnLinks) | **Patch** /drives({drive-id})/list/contentTypes({contentType-id})/columnLinks({columnLink-id}) | Update the navigation property columnLinks in drives



## DrivesListContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink DrivesListContentTypesCreateColumnLinks(ctx, driveId, contentTypeId, microsoftGraphColumnLink)
Create new navigation property to columnLinks for drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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


## DrivesListContentTypesGetColumnLinks

> MicrosoftGraphColumnLink DrivesListContentTypesGetColumnLinks(ctx, driveId, contentTypeId, columnLinkId, optional)
Get columnLinks from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
**columnLinkId** | **string**| key: columnLink-id of columnLink | 
 **optional** | ***DrivesListContentTypesGetColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListContentTypesGetColumnLinksOpts struct


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


## DrivesListContentTypesListColumnLinks

> CollectionOfColumnLink DrivesListContentTypesListColumnLinks(ctx, driveId, contentTypeId, optional)
Get columnLinks from drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
**contentTypeId** | **string**| key: contentType-id of contentType | 
 **optional** | ***DrivesListContentTypesListColumnLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DrivesListContentTypesListColumnLinksOpts struct


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


## DrivesListContentTypesUpdateColumnLinks

> DrivesListContentTypesUpdateColumnLinks(ctx, driveId, contentTypeId, columnLinkId, microsoftGraphColumnLink)
Update the navigation property columnLinks in drives

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string**| key: drive-id of drive | 
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

