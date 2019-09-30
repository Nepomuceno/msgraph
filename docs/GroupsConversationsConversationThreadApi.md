# \GroupsConversationsConversationThreadApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsConversationsCreateThreads**](GroupsConversationsConversationThreadApi.md#GroupsConversationsCreateThreads) | **Post** /groups({group-id})/conversations({conversation-id})/threads | Create new navigation property to threads for groups
[**GroupsConversationsGetThreads**](GroupsConversationsConversationThreadApi.md#GroupsConversationsGetThreads) | **Get** /groups({group-id})/conversations({conversation-id})/threads({conversationThread-id}) | Get threads from groups
[**GroupsConversationsListThreads**](GroupsConversationsConversationThreadApi.md#GroupsConversationsListThreads) | **Get** /groups({group-id})/conversations({conversation-id})/threads | Get threads from groups
[**GroupsConversationsUpdateThreads**](GroupsConversationsConversationThreadApi.md#GroupsConversationsUpdateThreads) | **Patch** /groups({group-id})/conversations({conversation-id})/threads({conversationThread-id}) | Update the navigation property threads in groups



## GroupsConversationsCreateThreads

> MicrosoftGraphConversationThread GroupsConversationsCreateThreads(ctx, groupId, conversationId, microsoftGraphConversationThread)
Create new navigation property to threads for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**microsoftGraphConversationThread** | [**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md)| New navigation property | 

### Return type

[**MicrosoftGraphConversationThread**](microsoft.graph.conversationThread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsGetThreads

> MicrosoftGraphConversationThread GroupsConversationsGetThreads(ctx, groupId, conversationId, conversationThreadId, optional)
Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsConversationsGetThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsGetThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphConversationThread**](microsoft.graph.conversationThread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsListThreads

> CollectionOfConversationThread GroupsConversationsListThreads(ctx, groupId, conversationId, optional)
Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
 **optional** | ***GroupsConversationsListThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsConversationsListThreadsOpts struct


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

[**CollectionOfConversationThread**](Collection of conversationThread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsUpdateThreads

> GroupsConversationsUpdateThreads(ctx, groupId, conversationId, conversationThreadId, microsoftGraphConversationThread)
Update the navigation property threads in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**microsoftGraphConversationThread** | [**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md)| New navigation property values | 

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

