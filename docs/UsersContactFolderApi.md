# \UsersContactFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateContactFolders**](UsersContactFolderApi.md#UsersCreateContactFolders) | **Post** /users({user-id})/contactFolders | Create new navigation property to contactFolders for users
[**UsersGetContactFolders**](UsersContactFolderApi.md#UsersGetContactFolders) | **Get** /users({user-id})/contactFolders({contactFolder-id}) | Get contactFolders from users
[**UsersListContactFolders**](UsersContactFolderApi.md#UsersListContactFolders) | **Get** /users({user-id})/contactFolders | Get contactFolders from users
[**UsersUpdateContactFolders**](UsersContactFolderApi.md#UsersUpdateContactFolders) | **Patch** /users({user-id})/contactFolders({contactFolder-id}) | Update the navigation property contactFolders in users



## UsersCreateContactFolders

> MicrosoftGraphContactFolder UsersCreateContactFolders(ctx, userId, microsoftGraphContactFolder)
Create new navigation property to contactFolders for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersGetContactFolders

> MicrosoftGraphContactFolder UsersGetContactFolders(ctx, userId, contactFolderId, optional)
Get contactFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***UsersGetContactFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetContactFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## UsersListContactFolders

> CollectionOfContactFolder UsersListContactFolders(ctx, userId, optional)
Get contactFolders from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListContactFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListContactFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## UsersUpdateContactFolders

> UsersUpdateContactFolders(ctx, userId, contactFolderId, microsoftGraphContactFolder)
Update the navigation property contactFolders in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
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

