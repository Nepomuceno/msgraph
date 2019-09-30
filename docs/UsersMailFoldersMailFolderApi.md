# \UsersMailFoldersMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMailFoldersCreateChildFolders**](UsersMailFoldersMailFolderApi.md#UsersMailFoldersCreateChildFolders) | **Post** /users({user-id})/mailFolders({mailFolder-id})/childFolders | Create new navigation property to childFolders for users
[**UsersMailFoldersGetChildFolders**](UsersMailFoldersMailFolderApi.md#UsersMailFoldersGetChildFolders) | **Get** /users({user-id})/mailFolders({mailFolder-id})/childFolders({mailFolder-id1}) | Get childFolders from users
[**UsersMailFoldersListChildFolders**](UsersMailFoldersMailFolderApi.md#UsersMailFoldersListChildFolders) | **Get** /users({user-id})/mailFolders({mailFolder-id})/childFolders | Get childFolders from users
[**UsersMailFoldersUpdateChildFolders**](UsersMailFoldersMailFolderApi.md#UsersMailFoldersUpdateChildFolders) | **Patch** /users({user-id})/mailFolders({mailFolder-id})/childFolders({mailFolder-id1}) | Update the navigation property childFolders in users



## UsersMailFoldersCreateChildFolders

> MicrosoftGraphMailFolder UsersMailFoldersCreateChildFolders(ctx, userId, mailFolderId, microsoftGraphMailFolder)
Create new navigation property to childFolders for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersGetChildFolders

> MicrosoftGraphMailFolder UsersMailFoldersGetChildFolders(ctx, userId, mailFolderId, mailFolderId1, optional)
Get childFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersGetChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetChildFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersListChildFolders

> CollectionOfMailFolder UsersMailFoldersListChildFolders(ctx, userId, mailFolderId, optional)
Get childFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListChildFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListChildFoldersOpts struct


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

[**CollectionOfMailFolder**](Collection of mailFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersUpdateChildFolders

> UsersMailFoldersUpdateChildFolders(ctx, userId, mailFolderId, mailFolderId1, microsoftGraphMailFolder)
Update the navigation property childFolders in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**mailFolderId1** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)| New navigation property values | 

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

