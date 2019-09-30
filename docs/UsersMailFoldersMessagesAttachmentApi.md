# \UsersMailFoldersMessagesAttachmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMailFoldersMessagesCreateAttachments**](UsersMailFoldersMessagesAttachmentApi.md#UsersMailFoldersMessagesCreateAttachments) | **Post** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/attachments | Create new navigation property to attachments for users
[**UsersMailFoldersMessagesGetAttachments**](UsersMailFoldersMessagesAttachmentApi.md#UsersMailFoldersMessagesGetAttachments) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/attachments({attachment-id}) | Get attachments from users
[**UsersMailFoldersMessagesListAttachments**](UsersMailFoldersMessagesAttachmentApi.md#UsersMailFoldersMessagesListAttachments) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/attachments | Get attachments from users
[**UsersMailFoldersMessagesUpdateAttachments**](UsersMailFoldersMessagesAttachmentApi.md#UsersMailFoldersMessagesUpdateAttachments) | **Patch** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/attachments({attachment-id}) | Update the navigation property attachments in users



## UsersMailFoldersMessagesCreateAttachments

> MicrosoftGraphAttachment UsersMailFoldersMessagesCreateAttachments(ctx, userId, mailFolderId, messageId, microsoftGraphAttachment)
Create new navigation property to attachments for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**microsoftGraphAttachment** | [**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md)| New navigation property | 

### Return type

[**MicrosoftGraphAttachment**](microsoft.graph.attachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersMessagesGetAttachments

> MicrosoftGraphAttachment UsersMailFoldersMessagesGetAttachments(ctx, userId, mailFolderId, messageId, attachmentId, optional)
Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***UsersMailFoldersMessagesGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetAttachmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAttachment**](microsoft.graph.attachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersMessagesListAttachments

> CollectionOfAttachment UsersMailFoldersMessagesListAttachments(ctx, userId, mailFolderId, messageId, optional)
Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListAttachmentsOpts struct


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

[**CollectionOfAttachment**](Collection of attachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersMessagesUpdateAttachments

> UsersMailFoldersMessagesUpdateAttachments(ctx, userId, mailFolderId, messageId, attachmentId, microsoftGraphAttachment)
Update the navigation property attachments in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
**microsoftGraphAttachment** | [**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md)| New navigation property values | 

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

