# \UsersCalendarsEventsAttachmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarsEventsCreateAttachments**](UsersCalendarsEventsAttachmentApi.md#UsersCalendarsEventsCreateAttachments) | **Post** /users({user-id})/calendars({calendar-id})/events({event-id})/attachments | Create new navigation property to attachments for users
[**UsersCalendarsEventsGetAttachments**](UsersCalendarsEventsAttachmentApi.md#UsersCalendarsEventsGetAttachments) | **Get** /users({user-id})/calendars({calendar-id})/events({event-id})/attachments({attachment-id}) | Get attachments from users
[**UsersCalendarsEventsListAttachments**](UsersCalendarsEventsAttachmentApi.md#UsersCalendarsEventsListAttachments) | **Get** /users({user-id})/calendars({calendar-id})/events({event-id})/attachments | Get attachments from users
[**UsersCalendarsEventsUpdateAttachments**](UsersCalendarsEventsAttachmentApi.md#UsersCalendarsEventsUpdateAttachments) | **Patch** /users({user-id})/calendars({calendar-id})/events({event-id})/attachments({attachment-id}) | Update the navigation property attachments in users



## UsersCalendarsEventsCreateAttachments

> MicrosoftGraphAttachment UsersCalendarsEventsCreateAttachments(ctx, userId, calendarId, eventId, microsoftGraphAttachment)
Create new navigation property to attachments for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
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


## UsersCalendarsEventsGetAttachments

> MicrosoftGraphAttachment UsersCalendarsEventsGetAttachments(ctx, userId, calendarId, eventId, attachmentId, optional)
Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
**attachmentId** | **string**| key: attachment-id of attachment | 
 **optional** | ***UsersCalendarsEventsGetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarsEventsGetAttachmentsOpts struct


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


## UsersCalendarsEventsListAttachments

> CollectionOfAttachment UsersCalendarsEventsListAttachments(ctx, userId, calendarId, eventId, optional)
Get attachments from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarsEventsListAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarsEventsListAttachmentsOpts struct


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


## UsersCalendarsEventsUpdateAttachments

> UsersCalendarsEventsUpdateAttachments(ctx, userId, calendarId, eventId, attachmentId, microsoftGraphAttachment)
Update the navigation property attachments in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**calendarId** | **string**| key: calendar-id of calendar | 
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

