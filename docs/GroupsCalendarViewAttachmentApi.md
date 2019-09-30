# \GroupsCalendarViewAttachmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarViewCreateAttachments**](GroupsCalendarViewAttachmentApi.md#GroupsCalendarViewCreateAttachments) | **Post** /groups({group-id})/calendarView({event-id})/attachments | Create new navigation property to attachments for groups
[**GroupsCalendarViewGetAttachments**](GroupsCalendarViewAttachmentApi.md#GroupsCalendarViewGetAttachments) | **Get** /groups({group-id})/calendarView({event-id})/attachments({attachment-id}) | Get attachments from groups
[**GroupsCalendarViewListAttachments**](GroupsCalendarViewAttachmentApi.md#GroupsCalendarViewListAttachments) | **Get** /groups({group-id})/calendarView({event-id})/attachments | Get attachments from groups
[**GroupsCalendarViewUpdateAttachments**](GroupsCalendarViewAttachmentApi.md#GroupsCalendarViewUpdateAttachments) | **Patch** /groups({group-id})/calendarView({event-id})/attachments({attachment-id}) | Update the navigation property attachments in groups



## GroupsCalendarViewCreateAttachments

> MicrosoftGraphAttachment GroupsCalendarViewCreateAttachments(ctx, groupId, eventId, microsoftGraphAttachment)
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


## GroupsCalendarViewGetAttachments

> MicrosoftGraphAttachment GroupsCalendarViewGetAttachments(ctx, groupId, eventId, attachmentId, optional)
Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***GroupsCalendarViewGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarViewGetAttachmentsOpts struct


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


## GroupsCalendarViewListAttachments

> CollectionOfAttachment GroupsCalendarViewListAttachments(ctx, groupId, eventId, optional)
Get attachments from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***GroupsCalendarViewListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsCalendarViewListAttachmentsOpts struct


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


## GroupsCalendarViewUpdateAttachments

> GroupsCalendarViewUpdateAttachments(ctx, groupId, eventId, attachmentId, microsoftGraphAttachment)
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

