# \MeContactFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateContactFolders**](MeContactFolderApi.md#MeCreateContactFolders) | **Post** /me/contactFolders | Create new navigation property to contactFolders for me
[**MeGetContactFolders**](MeContactFolderApi.md#MeGetContactFolders) | **Get** /me/contactFolders({contactFolder-id}) | Get contactFolders from me
[**MeListContactFolders**](MeContactFolderApi.md#MeListContactFolders) | **Get** /me/contactFolders | Get contactFolders from me
[**MeUpdateContactFolders**](MeContactFolderApi.md#MeUpdateContactFolders) | **Patch** /me/contactFolders({contactFolder-id}) | Update the navigation property contactFolders in me



## MeCreateContactFolders

> MicrosoftGraphContactFolder MeCreateContactFolders(ctx, microsoftGraphContactFolder)
Create new navigation property to contactFolders for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphContactFolder** | [**MicrosoftGraphContactFolder**](MicrosoftGraphContactFolder.md)| New navigation property | 

### Return type

[**MicrosoftGraphContactFolder**](microsoft.graph.contactFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetContactFolders

> MicrosoftGraphContactFolder MeGetContactFolders(ctx, contactFolderId, optional)
Get contactFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***MeGetContactFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetContactFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphContactFolder**](microsoft.graph.contactFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListContactFolders

> CollectionOfContactFolder MeListContactFolders(ctx, optional)
Get contactFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListContactFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListContactFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfContactFolder**](Collection of contactFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateContactFolders

> MeUpdateContactFolders(ctx, contactFolderId, microsoftGraphContactFolder)
Update the navigation property contactFolders in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**microsoftGraphContactFolder** | [**MicrosoftGraphContactFolder**](MicrosoftGraphContactFolder.md)| New navigation property values | 

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

