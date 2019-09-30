# \MeContactFoldersContactFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeContactFoldersCreateChildFolders**](MeContactFoldersContactFolderApi.md#MeContactFoldersCreateChildFolders) | **Post** /me/contactFolders({contactFolder-id})/childFolders | Create new navigation property to childFolders for me
[**MeContactFoldersGetChildFolders**](MeContactFoldersContactFolderApi.md#MeContactFoldersGetChildFolders) | **Get** /me/contactFolders({contactFolder-id})/childFolders({contactFolder-id1}) | Get childFolders from me
[**MeContactFoldersListChildFolders**](MeContactFoldersContactFolderApi.md#MeContactFoldersListChildFolders) | **Get** /me/contactFolders({contactFolder-id})/childFolders | Get childFolders from me
[**MeContactFoldersUpdateChildFolders**](MeContactFoldersContactFolderApi.md#MeContactFoldersUpdateChildFolders) | **Patch** /me/contactFolders({contactFolder-id})/childFolders({contactFolder-id1}) | Update the navigation property childFolders in me



## MeContactFoldersCreateChildFolders

> MicrosoftGraphContactFolder MeContactFoldersCreateChildFolders(ctx, contactFolderId, microsoftGraphContactFolder)
Create new navigation property to childFolders for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**microsoftGraphContactFolder** | [**MicrosoftGraphContactFolder**](MicrosoftGraphContactFolder.md)| New navigation property | 

### Return type

[**MicrosoftGraphContactFolder**](microsoft.graph.contactFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeContactFoldersGetChildFolders

> MicrosoftGraphContactFolder MeContactFoldersGetChildFolders(ctx, contactFolderId, contactFolderId1, optional)
Get childFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactFolderId1** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***MeContactFoldersGetChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeContactFoldersGetChildFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphContactFolder**](microsoft.graph.contactFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeContactFoldersListChildFolders

> CollectionOfContactFolder MeContactFoldersListChildFolders(ctx, contactFolderId, optional)
Get childFolders from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***MeContactFoldersListChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeContactFoldersListChildFoldersOpts struct


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

[**CollectionOfContactFolder**](Collection of contactFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeContactFoldersUpdateChildFolders

> MeContactFoldersUpdateChildFolders(ctx, contactFolderId, contactFolderId1, microsoftGraphContactFolder)
Update the navigation property childFolders in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactFolderId1** | **string**| key: contactFolder-id of contactFolder | 
**microsoftGraphContactFolder** | [**MicrosoftGraphContactFolder**](MicrosoftGraphContactFolder.md)| New navigation property values | 

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

