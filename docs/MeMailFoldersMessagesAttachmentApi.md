# \MeMailFoldersMessagesAttachmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeMailFoldersMessagesCreateAttachments**](MeMailFoldersMessagesAttachmentApi.md#MeMailFoldersMessagesCreateAttachments) | **Post** /me/mailFolders({mailFolder-id})/messages({message-id})/attachments | Create new navigation property to attachments for me
[**MeMailFoldersMessagesGetAttachments**](MeMailFoldersMessagesAttachmentApi.md#MeMailFoldersMessagesGetAttachments) | **Get** /me/mailFolders({mailFolder-id})/messages({message-id})/attachments({attachment-id}) | Get attachments from me
[**MeMailFoldersMessagesListAttachments**](MeMailFoldersMessagesAttachmentApi.md#MeMailFoldersMessagesListAttachments) | **Get** /me/mailFolders({mailFolder-id})/messages({message-id})/attachments | Get attachments from me
[**MeMailFoldersMessagesUpdateAttachments**](MeMailFoldersMessagesAttachmentApi.md#MeMailFoldersMessagesUpdateAttachments) | **Patch** /me/mailFolders({mailFolder-id})/messages({message-id})/attachments({attachment-id}) | Update the navigation property attachments in me



## MeMailFoldersMessagesCreateAttachments

> MicrosoftGraphAttachment MeMailFoldersMessagesCreateAttachments(ctx, mailFolderId, messageId, microsoftGraphAttachment)
Create new navigation property to attachments for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersMessagesGetAttachments

> MicrosoftGraphAttachment MeMailFoldersMessagesGetAttachments(ctx, mailFolderId, messageId, attachmentId, optional)
Get attachments from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***MeMailFoldersMessagesGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesGetAttachmentsOpts struct


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


## MeMailFoldersMessagesListAttachments

> CollectionOfAttachment MeMailFoldersMessagesListAttachments(ctx, mailFolderId, messageId, optional)
Get attachments from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***MeMailFoldersMessagesListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersMessagesListAttachmentsOpts struct


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


## MeMailFoldersMessagesUpdateAttachments

> MeMailFoldersMessagesUpdateAttachments(ctx, mailFolderId, messageId, attachmentId, microsoftGraphAttachment)
Update the navigation property attachments in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

