# \GroupsConversationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateConversations**](GroupsConversationApi.md#GroupsCreateConversations) | **Post** /groups({group-id})/conversations | Create new navigation property to conversations for groups
[**GroupsGetConversations**](GroupsConversationApi.md#GroupsGetConversations) | **Get** /groups({group-id})/conversations({conversation-id}) | Get conversations from groups
[**GroupsListConversations**](GroupsConversationApi.md#GroupsListConversations) | **Get** /groups({group-id})/conversations | Get conversations from groups
[**GroupsUpdateConversations**](GroupsConversationApi.md#GroupsUpdateConversations) | **Patch** /groups({group-id})/conversations({conversation-id}) | Update the navigation property conversations in groups



## GroupsCreateConversations

> MicrosoftGraphConversation GroupsCreateConversations(ctx, groupId, microsoftGraphConversation)
Create new navigation property to conversations for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphConversation** | [**MicrosoftGraphConversation**](MicrosoftGraphConversation.md)| New navigation property | 

### Return type

[**MicrosoftGraphConversation**](microsoft.graph.conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetConversations

> MicrosoftGraphConversation GroupsGetConversations(ctx, groupId, conversationId, optional)
Get conversations from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
 **optional** | ***GroupsGetConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetConversationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphConversation**](microsoft.graph.conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListConversations

> CollectionOfConversation GroupsListConversations(ctx, groupId, optional)
Get conversations from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListConversationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfConversation**](Collection of conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateConversations

> GroupsUpdateConversations(ctx, groupId, conversationId, microsoftGraphConversation)
Update the navigation property conversations in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationId** | **string**| key: conversation-id of conversation | 
**microsoftGraphConversation** | [**MicrosoftGraphConversation**](MicrosoftGraphConversation.md)| New navigation property values | 

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

