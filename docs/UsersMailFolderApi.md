# \UsersMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateMailFolders**](UsersMailFolderApi.md#UsersCreateMailFolders) | **Post** /users({user-id})/mailFolders | Create new navigation property to mailFolders for users
[**UsersGetMailFolders**](UsersMailFolderApi.md#UsersGetMailFolders) | **Get** /users({user-id})/mailFolders({mailFolder-id}) | Get mailFolders from users
[**UsersListMailFolders**](UsersMailFolderApi.md#UsersListMailFolders) | **Get** /users({user-id})/mailFolders | Get mailFolders from users
[**UsersUpdateMailFolders**](UsersMailFolderApi.md#UsersUpdateMailFolders) | **Patch** /users({user-id})/mailFolders({mailFolder-id}) | Update the navigation property mailFolders in users



## UsersCreateMailFolders

> MicrosoftGraphMailFolder UsersCreateMailFolders(ctx, userId, microsoftGraphMailFolder)
Create new navigation property to mailFolders for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersGetMailFolders

> MicrosoftGraphMailFolder UsersGetMailFolders(ctx, userId, mailFolderId, optional)
Get mailFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersGetMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetMailFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## UsersListMailFolders

> CollectionOfMailFolder UsersListMailFolders(ctx, userId, optional)
Get mailFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListMailFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListMailFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## UsersUpdateMailFolders

> UsersUpdateMailFolders(ctx, userId, mailFolderId, microsoftGraphMailFolder)
Update the navigation property mailFolders in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
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

