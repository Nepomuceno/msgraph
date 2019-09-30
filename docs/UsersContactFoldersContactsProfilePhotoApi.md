# \UsersContactFoldersContactsProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactFoldersContactsGetPhoto**](UsersContactFoldersContactsProfilePhotoApi.md#UsersContactFoldersContactsGetPhoto) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/photo | Get photo from users
[**UsersContactFoldersContactsUpdatePhoto**](UsersContactFoldersContactsProfilePhotoApi.md#UsersContactFoldersContactsUpdatePhoto) | **Patch** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/photo | Update the navigation property photo in users



## UsersContactFoldersContactsGetPhoto

> MicrosoftGraphProfilePhoto UsersContactFoldersContactsGetPhoto(ctx, userId, contactFolderId, contactId, optional)
Get photo from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactFoldersContactsGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersContactsGetPhotoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsUpdatePhoto

> UsersContactFoldersContactsUpdatePhoto(ctx, userId, contactFolderId, contactId, microsoftGraphProfilePhoto)
Update the navigation property photo in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)| New navigation property values | 

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

