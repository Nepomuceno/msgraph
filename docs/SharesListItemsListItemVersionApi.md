# \SharesListItemsListItemVersionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesListItemsCreateVersions**](SharesListItemsListItemVersionApi.md#SharesListItemsCreateVersions) | **Post** /shares({sharedDriveItem-id})/list/items({listItem-id})/versions | Create new navigation property to versions for shares
[**SharesListItemsGetVersions**](SharesListItemsListItemVersionApi.md#SharesListItemsGetVersions) | **Get** /shares({sharedDriveItem-id})/list/items({listItem-id})/versions({listItemVersion-id}) | Get versions from shares
[**SharesListItemsListVersions**](SharesListItemsListItemVersionApi.md#SharesListItemsListVersions) | **Get** /shares({sharedDriveItem-id})/list/items({listItem-id})/versions | Get versions from shares
[**SharesListItemsUpdateVersions**](SharesListItemsListItemVersionApi.md#SharesListItemsUpdateVersions) | **Patch** /shares({sharedDriveItem-id})/list/items({listItem-id})/versions({listItemVersion-id}) | Update the navigation property versions in shares
[**SharesListItemsVersionsGetFields**](SharesListItemsListItemVersionApi.md#SharesListItemsVersionsGetFields) | **Get** /shares({sharedDriveItem-id})/list/items({listItem-id})/versions({listItemVersion-id})/fields | Get fields from shares
[**SharesListItemsVersionsUpdateFields**](SharesListItemsListItemVersionApi.md#SharesListItemsVersionsUpdateFields) | **Patch** /shares({sharedDriveItem-id})/list/items({listItem-id})/versions({listItemVersion-id})/fields | Update the navigation property fields in shares



## SharesListItemsCreateVersions

> MicrosoftGraphListItemVersion SharesListItemsCreateVersions(ctx, sharedDriveItemId, listItemId, microsoftGraphListItemVersion)
Create new navigation property to versions for shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
**microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)| New navigation property | 

### Return type

[**MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemsGetVersions

> MicrosoftGraphListItemVersion SharesListItemsGetVersions(ctx, sharedDriveItemId, listItemId, listItemVersionId, optional)
Get versions from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***SharesListItemsGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemsGetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemsListVersions

> CollectionOfListItemVersion SharesListItemsListVersions(ctx, sharedDriveItemId, listItemId, optional)
Get versions from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
 **optional** | ***SharesListItemsListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemsListVersionsOpts struct


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

[**CollectionOfListItemVersion**](Collection of listItemVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesListItemsUpdateVersions

> SharesListItemsUpdateVersions(ctx, sharedDriveItemId, listItemId, listItemVersionId, microsoftGraphListItemVersion)
Update the navigation property versions in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
**microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)| New navigation property values | 

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


## SharesListItemsVersionsGetFields

> MicrosoftGraphFieldValueSet SharesListItemsVersionsGetFields(ctx, sharedDriveItemId, listItemId, listItemVersionId, optional)
Get fields from shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
 **optional** | ***SharesListItemsVersionsGetFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SharesListItemsVersionsGetFieldsOpts struct


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


## SharesListItemsVersionsUpdateFields

> SharesListItemsVersionsUpdateFields(ctx, sharedDriveItemId, listItemId, listItemVersionId, microsoftGraphFieldValueSet)
Update the navigation property fields in shares

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string**| key: sharedDriveItem-id of sharedDriveItem | 
**listItemId** | **string**| key: listItem-id of listItem | 
**listItemVersionId** | **string**| key: listItemVersion-id of listItemVersion | 
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

