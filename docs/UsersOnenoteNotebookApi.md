# \UsersOnenoteNotebookApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteCreateNotebooks**](UsersOnenoteNotebookApi.md#UsersOnenoteCreateNotebooks) | **Post** /users({user-id})/onenote/notebooks | Create new navigation property to notebooks for users
[**UsersOnenoteGetNotebooks**](UsersOnenoteNotebookApi.md#UsersOnenoteGetNotebooks) | **Get** /users({user-id})/onenote/notebooks({notebook-id}) | Get notebooks from users
[**UsersOnenoteListNotebooks**](UsersOnenoteNotebookApi.md#UsersOnenoteListNotebooks) | **Get** /users({user-id})/onenote/notebooks | Get notebooks from users
[**UsersOnenoteUpdateNotebooks**](UsersOnenoteNotebookApi.md#UsersOnenoteUpdateNotebooks) | **Patch** /users({user-id})/onenote/notebooks({notebook-id}) | Update the navigation property notebooks in users



## UsersOnenoteCreateNotebooks

> MicrosoftGraphNotebook UsersOnenoteCreateNotebooks(ctx, userId, microsoftGraphNotebook)
Create new navigation property to notebooks for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteGetNotebooks

> MicrosoftGraphNotebook UsersOnenoteGetNotebooks(ctx, userId, notebookId, optional)
Get notebooks from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
 **optional** | ***UsersOnenoteGetNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteGetNotebooksOpts struct


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


## UsersOnenoteListNotebooks

> CollectionOfNotebook UsersOnenoteListNotebooks(ctx, userId, optional)
Get notebooks from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersOnenoteListNotebooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteListNotebooksOpts struct


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


## UsersOnenoteUpdateNotebooks

> UsersOnenoteUpdateNotebooks(ctx, userId, notebookId, microsoftGraphNotebook)
Update the navigation property notebooks in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

