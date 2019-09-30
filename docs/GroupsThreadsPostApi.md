# \GroupsThreadsPostApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsThreadsCreatePosts**](GroupsThreadsPostApi.md#GroupsThreadsCreatePosts) | **Post** /groups({group-id})/threads({conversationThread-id})/posts | Create new navigation property to posts for groups
[**GroupsThreadsGetPosts**](GroupsThreadsPostApi.md#GroupsThreadsGetPosts) | **Get** /groups({group-id})/threads({conversationThread-id})/posts({post-id}) | Get posts from groups
[**GroupsThreadsListPosts**](GroupsThreadsPostApi.md#GroupsThreadsListPosts) | **Get** /groups({group-id})/threads({conversationThread-id})/posts | Get posts from groups
[**GroupsThreadsUpdatePosts**](GroupsThreadsPostApi.md#GroupsThreadsUpdatePosts) | **Patch** /groups({group-id})/threads({conversationThread-id})/posts({post-id}) | Update the navigation property posts in groups



## GroupsThreadsCreatePosts

> MicrosoftGraphPost GroupsThreadsCreatePosts(ctx, groupId, conversationThreadId, microsoftGraphPost)
Create new navigation property to posts for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md)| New navigation property | 

### Return type

[**MicrosoftGraphPost**](microsoft.graph.post.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsGetPosts

> MicrosoftGraphPost GroupsThreadsGetPosts(ctx, groupId, conversationThreadId, postId, optional)
Get posts from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsGetPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsGetPostsOpts struct


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


## GroupsThreadsListPosts

> CollectionOfPost GroupsThreadsListPosts(ctx, groupId, conversationThreadId, optional)
Get posts from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
 **optional** | ***GroupsThreadsListPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsListPostsOpts struct


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

[**CollectionOfPost**](Collection of post.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsUpdatePosts

> GroupsThreadsUpdatePosts(ctx, groupId, conversationThreadId, postId, microsoftGraphPost)
Update the navigation property posts in groups

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

