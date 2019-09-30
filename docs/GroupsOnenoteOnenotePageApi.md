# \GroupsOnenoteOnenotePageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteCreatePages**](GroupsOnenoteOnenotePageApi.md#GroupsOnenoteCreatePages) | **Post** /groups({group-id})/onenote/pages | Create new navigation property to pages for groups
[**GroupsOnenoteGetPages**](GroupsOnenoteOnenotePageApi.md#GroupsOnenoteGetPages) | **Get** /groups({group-id})/onenote/pages({onenotePage-id}) | Get pages from groups
[**GroupsOnenoteListPages**](GroupsOnenoteOnenotePageApi.md#GroupsOnenoteListPages) | **Get** /groups({group-id})/onenote/pages | Get pages from groups
[**GroupsOnenoteUpdatePages**](GroupsOnenoteOnenotePageApi.md#GroupsOnenoteUpdatePages) | **Patch** /groups({group-id})/onenote/pages({onenotePage-id}) | Update the navigation property pages in groups



## GroupsOnenoteCreatePages

> MicrosoftGraphOnenotePage GroupsOnenoteCreatePages(ctx, groupId, microsoftGraphOnenotePage)
Create new navigation property to pages for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphOnenotePage** | [**MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenotePage**](microsoft.graph.onenotePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteGetPages

> MicrosoftGraphOnenotePage GroupsOnenoteGetPages(ctx, groupId, onenotePageId, optional)
Get pages from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***GroupsOnenoteGetPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteGetPagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenotePage**](microsoft.graph.onenotePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteListPages

> CollectionOfOnenotePage GroupsOnenoteListPages(ctx, groupId, optional)
Get pages from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsOnenoteListPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteListPagesOpts struct


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

[**CollectionOfOnenotePage**](Collection of onenotePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteUpdatePages

> GroupsOnenoteUpdatePages(ctx, groupId, onenotePageId, microsoftGraphOnenotePage)
Update the navigation property pages in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**microsoftGraphOnenotePage** | [**MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md)| New navigation property values | 

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

