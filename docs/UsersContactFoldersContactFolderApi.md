# \UsersContactFoldersContactFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactFoldersCreateChildFolders**](UsersContactFoldersContactFolderApi.md#UsersContactFoldersCreateChildFolders) | **Post** /users({user-id})/contactFolders({contactFolder-id})/childFolders | Create new navigation property to childFolders for users
[**UsersContactFoldersGetChildFolders**](UsersContactFoldersContactFolderApi.md#UsersContactFoldersGetChildFolders) | **Get** /users({user-id})/contactFolders({contactFolder-id})/childFolders({contactFolder-id1}) | Get childFolders from users
[**UsersContactFoldersListChildFolders**](UsersContactFoldersContactFolderApi.md#UsersContactFoldersListChildFolders) | **Get** /users({user-id})/contactFolders({contactFolder-id})/childFolders | Get childFolders from users
[**UsersContactFoldersUpdateChildFolders**](UsersContactFoldersContactFolderApi.md#UsersContactFoldersUpdateChildFolders) | **Patch** /users({user-id})/contactFolders({contactFolder-id})/childFolders({contactFolder-id1}) | Update the navigation property childFolders in users



## UsersContactFoldersCreateChildFolders

> MicrosoftGraphContactFolder UsersContactFoldersCreateChildFolders(ctx, userId, contactFolderId, microsoftGraphContactFolder)
Create new navigation property to childFolders for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersContactFoldersGetChildFolders

> MicrosoftGraphContactFolder UsersContactFoldersGetChildFolders(ctx, userId, contactFolderId, contactFolderId1, optional)
Get childFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactFolderId1** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***UsersContactFoldersGetChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersGetChildFoldersOpts struct


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


## UsersContactFoldersListChildFolders

> CollectionOfContactFolder UsersContactFoldersListChildFolders(ctx, userId, contactFolderId, optional)
Get childFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***UsersContactFoldersListChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersListChildFoldersOpts struct


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


## UsersContactFoldersUpdateChildFolders

> UsersContactFoldersUpdateChildFolders(ctx, userId, contactFolderId, contactFolderId1, microsoftGraphContactFolder)
Update the navigation property childFolders in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

