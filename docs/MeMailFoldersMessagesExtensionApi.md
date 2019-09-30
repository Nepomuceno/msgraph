# \MeMailFoldersMessagesExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeMailFoldersMessagesCreateExtensions**](MeMailFoldersMessagesExtensionApi.md#MeMailFoldersMessagesCreateExtensions) | **Post** /me/mailFolders({mailFolder-id})/messages({message-id})/extensions | Create new navigation property to extensions for me
[**MeMailFoldersMessagesGetExtensions**](MeMailFoldersMessagesExtensionApi.md#MeMailFoldersMessagesGetExtensions) | **Get** /me/mailFolders({mailFolder-id})/messages({message-id})/extensions({extension-id}) | Get extensions from me
[**MeMailFoldersMessagesListExtensions**](MeMailFoldersMessagesExtensionApi.md#MeMailFoldersMessagesListExtensions) | **Get** /me/mailFolders({mailFolder-id})/messages({message-id})/extensions | Get extensions from me
[**MeMailFoldersMessagesUpdateExtensions**](MeMailFoldersMessagesExtensionApi.md#MeMailFoldersMessagesUpdateExtensions) | **Patch** /me/mailFolders({mailFolder-id})/messages({message-id})/extensions({extension-id}) | Update the navigation property extensions in me



## MeMailFoldersMessagesCreateExtensions

> MicrosoftGraphExtension MeMailFoldersMessagesCreateExtensions(ctx, mailFolderId, messageId, microsoftGraphExtension)
Create new navigation property to extensions for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersMessagesGetExtensions

> MicrosoftGraphExtension MeMailFoldersMessagesGetExtensions(ctx, mailFolderId, messageId, extensionId, optional)
Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**extensionId** | **string**| key: extension-id of extension | 
 **optional** | ***MeMailFoldersMessagesGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesGetExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](microsoft.graph.extension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersMessagesListExtensions

> CollectionOfExtension MeMailFoldersMessagesListExtensions(ctx, mailFolderId, messageId, optional)
Get extensions from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersMessagesListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesListExtensionsOpts struct


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

[**CollectionOfExtension**](Collection of extension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersMessagesUpdateExtensions

> MeMailFoldersMessagesUpdateExtensions(ctx, mailFolderId, messageId, extensionId, microsoftGraphExtension)
Update the navigation property extensions in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**extensionId** | **string**| key: extension-id of extension | 
**microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)| New navigation property values | 

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

