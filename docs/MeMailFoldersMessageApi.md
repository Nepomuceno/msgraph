# \MeMailFoldersMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeMailFoldersCreateMessages**](MeMailFoldersMessageApi.md#MeMailFoldersCreateMessages) | **Post** /me/mailFolders({mailFolder-id})/messages | Create new navigation property to messages for me
[**MeMailFoldersGetMessages**](MeMailFoldersMessageApi.md#MeMailFoldersGetMessages) | **Get** /me/mailFolders({mailFolder-id})/messages({message-id}) | Get messages from me
[**MeMailFoldersListMessages**](MeMailFoldersMessageApi.md#MeMailFoldersListMessages) | **Get** /me/mailFolders({mailFolder-id})/messages | Get messages from me
[**MeMailFoldersUpdateMessages**](MeMailFoldersMessageApi.md#MeMailFoldersUpdateMessages) | **Patch** /me/mailFolders({mailFolder-id})/messages({message-id}) | Update the navigation property messages in me



## MeMailFoldersCreateMessages

> MicrosoftGraphMessage MeMailFoldersCreateMessages(ctx, mailFolderId, microsoftGraphMessage)
Create new navigation property to messages for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMessage** | [**MicrosoftGraphMessage**](MicrosoftGraphMessage.md)| New navigation property | 

### Return type

[**MicrosoftGraphMessage**](microsoft.graph.message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersGetMessages

> MicrosoftGraphMessage MeMailFoldersGetMessages(ctx, mailFolderId, messageId, optional)
Get messages from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersGetMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMessage**](microsoft.graph.message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersListMessages

> CollectionOfMessage MeMailFoldersListMessages(ctx, mailFolderId, optional)
Get messages from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListMessagesOpts struct


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

[**CollectionOfMessage**](Collection of message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersUpdateMessages

> MeMailFoldersUpdateMessages(ctx, mailFolderId, messageId, microsoftGraphMessage)
Update the navigation property messages in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**microsoftGraphMessage** | [**MicrosoftGraphMessage**](MicrosoftGraphMessage.md)| New navigation property values | 

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

