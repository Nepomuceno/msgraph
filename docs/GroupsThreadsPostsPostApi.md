# \GroupsThreadsPostsPostApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsThreadsPostsGetInReplyTo**](GroupsThreadsPostsPostApi.md#GroupsThreadsPostsGetInReplyTo) | **Get** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/inReplyTo | Get inReplyTo from groups
[**GroupsThreadsPostsUpdateInReplyTo**](GroupsThreadsPostsPostApi.md#GroupsThreadsPostsUpdateInReplyTo) | **Patch** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/inReplyTo | Update the navigation property inReplyTo in groups



## GroupsThreadsPostsGetInReplyTo

> MicrosoftGraphPost GroupsThreadsPostsGetInReplyTo(ctx, groupId, conversationThreadId, postId, optional)
Get inReplyTo from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsGetInReplyToOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetInReplyToOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPost**](microsoft.graph.post.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsUpdateInReplyTo

> GroupsThreadsPostsUpdateInReplyTo(ctx, groupId, conversationThreadId, postId, microsoftGraphPost)
Update the navigation property inReplyTo in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md)| New navigation property values | 

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

