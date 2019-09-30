# \GroupsEventsAttachmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsEventsCreateAttachments**](GroupsEventsAttachmentApi.md#GroupsEventsCreateAttachments) | **Post** /groups({group-id})/events({event-id})/attachments | Create new navigation property to attachments for groups
[**GroupsEventsGetAttachments**](GroupsEventsAttachmentApi.md#GroupsEventsGetAttachments) | **Get** /groups({group-id})/events({event-id})/attachments({attachment-id}) | Get attachments from groups
[**GroupsEventsListAttachments**](GroupsEventsAttachmentApi.md#GroupsEventsListAttachments) | **Get** /groups({group-id})/events({event-id})/attachments | Get attachments from groups
[**GroupsEventsUpdateAttachments**](GroupsEventsAttachmentApi.md#GroupsEventsUpdateAttachments) | **Patch** /groups({group-id})/events({event-id})/attachments({attachment-id}) | Update the navigation property attachments in groups



## GroupsEventsCreateAttachments

> MicrosoftGraphAttachment GroupsEventsCreateAttachments(ctx, groupId, eventId, microsoftGraphAttachment)
Create new navigation property to attachments for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
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


## GroupsEventsGetAttachments

> MicrosoftGraphAttachment GroupsEventsGetAttachments(ctx, groupId, eventId, attachmentId, optional)
Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***GroupsEventsGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsEventsGetAttachmentsOpts struct


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


## GroupsEventsListAttachments

> CollectionOfAttachment GroupsEventsListAttachments(ctx, groupId, eventId, optional)
Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsEventsListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsEventsListAttachmentsOpts struct


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


## GroupsEventsUpdateAttachments

> GroupsEventsUpdateAttachments(ctx, groupId, eventId, attachmentId, microsoftGraphAttachment)
Update the navigation property attachments in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
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

