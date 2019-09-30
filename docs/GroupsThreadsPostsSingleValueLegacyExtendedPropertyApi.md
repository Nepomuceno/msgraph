# \GroupsThreadsPostsSingleValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsThreadsPostsCreateSingleValueExtendedProperties**](GroupsThreadsPostsSingleValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsCreateSingleValueExtendedProperties) | **Post** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for groups
[**GroupsThreadsPostsGetSingleValueExtendedProperties**](GroupsThreadsPostsSingleValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsGetSingleValueExtendedProperties) | **Get** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Get singleValueExtendedProperties from groups
[**GroupsThreadsPostsListSingleValueExtendedProperties**](GroupsThreadsPostsSingleValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsListSingleValueExtendedProperties) | **Get** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/singleValueExtendedProperties | Get singleValueExtendedProperties from groups
[**GroupsThreadsPostsUpdateSingleValueExtendedProperties**](GroupsThreadsPostsSingleValueLegacyExtendedPropertyApi.md#GroupsThreadsPostsUpdateSingleValueExtendedProperties) | **Patch** /groups({group-id})/threads({conversationThread-id})/posts({post-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Update the navigation property singleValueExtendedProperties in groups



## GroupsThreadsPostsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsThreadsPostsCreateSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, microsoftGraphSingleValueLegacyExtendedProperty)
Create new navigation property to singleValueExtendedProperties for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsThreadsPostsGetSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId, optional)
Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***GroupsThreadsPostsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsGetSingleValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty GroupsThreadsPostsListSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, optional)
Get singleValueExtendedProperties from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
 **optional** | ***GroupsThreadsPostsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsThreadsPostsListSingleValueExtendedPropertiesOpts struct


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

[**CollectionOfSingleValueLegacyExtendedProperty**](Collection of singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsThreadsPostsUpdateSingleValueExtendedProperties

> GroupsThreadsPostsUpdateSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)
Update the navigation property singleValueExtendedProperties in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**conversationThreadId** | **string**| key: conversationThread-id of conversationThread | 
**postId** | **string**| key: post-id of post | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property values | 

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

