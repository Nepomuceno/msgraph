# \GroupsThreadsPostsMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsThreadsPostsCreateMultiValueExtendedProperties**](GroupsThreadsPostsMultiValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsCreateMultiValueExtendedProperties) | **Post** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for groups
[**GroupsThreadsPostsGetMultiValueExtendedProperties**](GroupsThreadsPostsMultiValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsGetMultiValueExtendedProperties) | **Get** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from groups
[**GroupsThreadsPostsListMultiValueExtendedProperties**](GroupsThreadsPostsMultiValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsListMultiValueExtendedProperties) | **Get** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/multiValueExtendedProperties | Get multiValueExtendedProperties from groups
[**GroupsThreadsPostsUpdateMultiValueExtendedProperties**](GroupsThreadsPostsMultiValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsUpdateMultiValueExtendedProperties) | **Patch** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in groups



## GroupsThreadsPostsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsThreadsPostsCreateMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsThreadsPostsGetMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***GroupsThreadsPostsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty GroupsThreadsPostsListMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, optional)
Get multiValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsListMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfMultiValueLegacyExtendedProperty**](Collection of multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsUpdateMultiValueExtendedProperties

> GroupsThreadsPostsUpdateMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property values | 

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

