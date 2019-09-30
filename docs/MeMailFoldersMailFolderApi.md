# \MeMailFoldersMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeMailFoldersCreateChildFolders**](MeMailFoldersMailFolderApi.md#MeMailFoldersCreateChildFolders) | **Post** /me/mailFolders({mailFolder-id})/childFolders | Create new navigation property to childFolders for me
[**MeMailFoldersGetChildFolders**](MeMailFoldersMailFolderApi.md#MeMailFoldersGetChildFolders) | **Get** /me/mailFolders({mailFolder-id})/childFolders({mailFolder-id1}) | Get childFolders from me
[**MeMailFoldersListChildFolders**](MeMailFoldersMailFolderApi.md#MeMailFoldersListChildFolders) | **Get** /me/mailFolders({mailFolder-id})/childFolders | Get childFolders from me
[**MeMailFoldersUpdateChildFolders**](MeMailFoldersMailFolderApi.md#MeMailFoldersUpdateChildFolders) | **Patch** /me/mailFolders({mailFolder-id})/childFolders({mailFolder-id1}) | Update the navigation property childFolders in me



## MeMailFoldersCreateChildFolders

> MicrosoftGraphMailFolder MeMailFoldersCreateChildFolders(ctx, mailFolderId, microsoftGraphMailFolder)
Create new navigation property to childFolders for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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


## MeMailFoldersGetChildFolders

> MicrosoftGraphMailFolder MeMailFoldersGetChildFolders(ctx, mailFolderId, mailFolderId1, optional)
Get childFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersGetChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetChildFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

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


## MeMailFoldersListChildFolders

> CollectionOfMailFolder MeMailFoldersListChildFolders(ctx, mailFolderId, optional)
Get childFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListChildFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

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


## MeMailFoldersUpdateChildFolders

> MeMailFoldersUpdateChildFolders(ctx, mailFolderId, mailFolderId1, microsoftGraphMailFolder)
Update the navigation property childFolders in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
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

