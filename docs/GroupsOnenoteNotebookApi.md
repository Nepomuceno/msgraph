# \GroupsOnenoteNotebookApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteCreateNotebooks**](GroupsOnenoteNotebookApi.md#GroupsOnenoteCreateNotebooks) | **Post** /groups({group-id})/onenote/notebooks | Create new navigation property to notebooks for groups
[**GroupsOnenoteGetNotebooks**](GroupsOnenoteNotebookApi.md#GroupsOnenoteGetNotebooks) | **Get** /groups({group-id})/onenote/notebooks({notebook-id}) | Get notebooks from groups
[**GroupsOnenoteListNotebooks**](GroupsOnenoteNotebookApi.md#GroupsOnenoteListNotebooks) | **Get** /groups({group-id})/onenote/notebooks | Get notebooks from groups
[**GroupsOnenoteUpdateNotebooks**](GroupsOnenoteNotebookApi.md#GroupsOnenoteUpdateNotebooks) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id}) | Update the navigation property notebooks in groups



## GroupsOnenoteCreateNotebooks

> MicrosoftGraphNotebook GroupsOnenoteCreateNotebooks(ctx, groupId, microsoftGraphNotebook)
Create new navigation property to notebooks for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteGetNotebooks

> MicrosoftGraphNotebook GroupsOnenoteGetNotebooks(ctx, groupId, notebookId, optional)
Get notebooks from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
 **optional** | ***GroupsOnenoteGetNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteGetNotebooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteListNotebooks

> CollectionOfNotebook GroupsOnenoteListNotebooks(ctx, groupId, optional)
Get notebooks from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsOnenoteListNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteListNotebooksOpts struct


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

[**CollectionOfNotebook**](Collection of notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteUpdateNotebooks

> GroupsOnenoteUpdateNotebooks(ctx, groupId, notebookId, microsoftGraphNotebook)
Update the navigation property notebooks in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property values | 

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

