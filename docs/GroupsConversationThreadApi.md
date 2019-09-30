# \GroupsConversationThreadApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateThreads**](GroupsConversationThreadApi.md#GroupsCreateThreads) | **Post** /groups({group-id})/threads | Create new navigation property to threads for groups
[**GroupsGetThreads**](GroupsConversationThreadApi.md#GroupsGetThreads) | **Get** /groups({group-id})/threads({conversationThread-id}) | Get threads from groups
[**GroupsListThreads**](GroupsConversationThreadApi.md#GroupsListThreads) | **Get** /groups({group-id})/threads | Get threads from groups
[**GroupsUpdateThreads**](GroupsConversationThreadApi.md#GroupsUpdateThreads) | **Patch** /groups({group-id})/threads({conversationThread-id}) | Update the navigation property threads in groups



## GroupsCreateThreads

> MicrosoftGraphConversationThread GroupsCreateThreads(ctx, groupId, microsoftGraphConversationThread)
Create new navigation property to threads for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsGetThreads

> MicrosoftGraphConversationThread GroupsGetThreads(ctx, groupId, conversationThreadId, optional)
Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsGetThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsGetThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## GroupsListThreads

> CollectionOfConversationThread GroupsListThreads(ctx, groupId, optional)
Get threads from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsListThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsListThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## GroupsUpdateThreads

> GroupsUpdateThreads(ctx, groupId, conversationThreadId, microsoftGraphConversationThread)
Update the navigation property threads in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

