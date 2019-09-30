# \GroupsOnenoteOnenoteResourceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteCreateResources**](GroupsOnenoteOnenoteResourceApi.md#GroupsOnenoteCreateResources) | **Post** /groups({group-id})/onenote/resources | Create new navigation property to resources for groups
[**GroupsOnenoteGetResources**](GroupsOnenoteOnenoteResourceApi.md#GroupsOnenoteGetResources) | **Get** /groups({group-id})/onenote/resources({onenoteResource-id}) | Get resources from groups
[**GroupsOnenoteListResources**](GroupsOnenoteOnenoteResourceApi.md#GroupsOnenoteListResources) | **Get** /groups({group-id})/onenote/resources | Get resources from groups
[**GroupsOnenoteUpdateResources**](GroupsOnenoteOnenoteResourceApi.md#GroupsOnenoteUpdateResources) | **Patch** /groups({group-id})/onenote/resources({onenoteResource-id}) | Update the navigation property resources in groups



## GroupsOnenoteCreateResources

> MicrosoftGraphOnenoteResource GroupsOnenoteCreateResources(ctx, groupId, microsoftGraphOnenoteResource)
Create new navigation property to resources for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphOnenoteResource** | [**MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenoteResource**](microsoft.graph.onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteGetResources

> MicrosoftGraphOnenoteResource GroupsOnenoteGetResources(ctx, groupId, onenoteResourceId, optional)
Get resources from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
 **optional** | ***GroupsOnenoteGetResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteGetResourcesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenoteResource**](microsoft.graph.onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteListResources

> CollectionOfOnenoteResource GroupsOnenoteListResources(ctx, groupId, optional)
Get resources from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsOnenoteListResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteListResourcesOpts struct


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

[**CollectionOfOnenoteResource**](Collection of onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteUpdateResources

> GroupsOnenoteUpdateResources(ctx, groupId, onenoteResourceId, microsoftGraphOnenoteResource)
Update the navigation property resources in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
**microsoftGraphOnenoteResource** | [**MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md)| New navigation property values | 

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

