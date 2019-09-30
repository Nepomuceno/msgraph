# \UsersMailFoldersMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMailFoldersCreateMessages**](UsersMailFoldersMessageApi.md#UsersMailFoldersCreateMessages) | **Post** /users({user-id})/mailFolders({mailFolder-id})/messages | Create new navigation property to messages for users
[**UsersMailFoldersGetMessages**](UsersMailFoldersMessageApi.md#UsersMailFoldersGetMessages) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id}) | Get messages from users
[**UsersMailFoldersListMessages**](UsersMailFoldersMessageApi.md#UsersMailFoldersListMessages) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages | Get messages from users
[**UsersMailFoldersUpdateMessages**](UsersMailFoldersMessageApi.md#UsersMailFoldersUpdateMessages) | **Patch** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id}) | Update the navigation property messages in users



## UsersMailFoldersCreateMessages

> MicrosoftGraphMessage UsersMailFoldersCreateMessages(ctx, userId, mailFolderId, microsoftGraphMessage)
Create new navigation property to messages for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersMailFoldersGetMessages

> MicrosoftGraphMessage UsersMailFoldersGetMessages(ctx, userId, mailFolderId, messageId, optional)
Get messages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersGetMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetMessagesOpts struct


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


## UsersMailFoldersListMessages

> CollectionOfMessage UsersMailFoldersListMessages(ctx, userId, mailFolderId, optional)
Get messages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListMessagesOpts struct


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


## UsersMailFoldersUpdateMessages

> UsersMailFoldersUpdateMessages(ctx, userId, mailFolderId, messageId, microsoftGraphMessage)
Update the navigation property messages in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

