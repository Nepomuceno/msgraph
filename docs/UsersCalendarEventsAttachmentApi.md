# \UsersCalendarEventsAttachmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarEventsCreateAttachments**](UsersCalendarEventsAttachmentApi.md#UsersCalendarEventsCreateAttachments) | **Post** /users({user-id})/calendar/events({event-id})/attachments | Create new navigation property to attachments for users
[**UsersCalendarEventsGetAttachments**](UsersCalendarEventsAttachmentApi.md#UsersCalendarEventsGetAttachments) | **Get** /users({user-id})/calendar/events({event-id})/attachments({attachment-id}) | Get attachments from users
[**UsersCalendarEventsListAttachments**](UsersCalendarEventsAttachmentApi.md#UsersCalendarEventsListAttachments) | **Get** /users({user-id})/calendar/events({event-id})/attachments | Get attachments from users
[**UsersCalendarEventsUpdateAttachments**](UsersCalendarEventsAttachmentApi.md#UsersCalendarEventsUpdateAttachments) | **Patch** /users({user-id})/calendar/events({event-id})/attachments({attachment-id}) | Update the navigation property attachments in users



## UsersCalendarEventsCreateAttachments

> MicrosoftGraphAttachment UsersCalendarEventsCreateAttachments(ctx, userId, eventId, microsoftGraphAttachment)
Create new navigation property to attachments for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersCalendarEventsGetAttachments

> MicrosoftGraphAttachment UsersCalendarEventsGetAttachments(ctx, userId, eventId, attachmentId, optional)
Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***UsersCalendarEventsGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarEventsGetAttachmentsOpts struct


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


## UsersCalendarEventsListAttachments

> CollectionOfAttachment UsersCalendarEventsListAttachments(ctx, userId, eventId, optional)
Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarEventsListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarEventsListAttachmentsOpts struct


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


## UsersCalendarEventsUpdateAttachments

> UsersCalendarEventsUpdateAttachments(ctx, userId, eventId, attachmentId, microsoftGraphAttachment)
Update the navigation property attachments in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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
