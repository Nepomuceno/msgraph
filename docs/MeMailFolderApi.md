# \MeMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateMailFolders**](MeMailFolderApi.md#MeCreateMailFolders) | **Post** /me/mailFolders | Create new navigation property to mailFolders for me
[**MeGetMailFolders**](MeMailFolderApi.md#MeGetMailFolders) | **Get** /me/mailFolders({mailFolder-id}) | Get mailFolders from me
[**MeListMailFolders**](MeMailFolderApi.md#MeListMailFolders) | **Get** /me/mailFolders | Get mailFolders from me
[**MeUpdateMailFolders**](MeMailFolderApi.md#MeUpdateMailFolders) | **Patch** /me/mailFolders({mailFolder-id}) | Update the navigation property mailFolders in me



## MeCreateMailFolders

> MicrosoftGraphMailFolder MeCreateMailFolders(ctx, microsoftGraphMailFolder)
Create new navigation property to mailFolders for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetMailFolders

> MicrosoftGraphMailFolder MeGetMailFolders(ctx, mailFolderId, optional)
Get mailFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeGetMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetMailFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListMailFolders

> CollectionOfMailFolder MeListMailFolders(ctx, optional)
Get mailFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListMailFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfMailFolder**](Collection of mailFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateMailFolders

> MeUpdateMailFolders(ctx, mailFolderId, microsoftGraphMailFolder)
Update the navigation property mailFolders in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property values | 

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

