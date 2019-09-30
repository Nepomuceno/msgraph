# \MeContactFoldersContactsProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeContactFoldersContactsGetPhoto**](MeContactFoldersContactsProfilePhotoApi.md#MeContactFoldersContactsGetPhoto) | **Get** /me/contactFolders({contactFolder-id})/contacts({contact-id})/photo | Get photo from me
[**MeContactFoldersContactsUpdatePhoto**](MeContactFoldersContactsProfilePhotoApi.md#MeContactFoldersContactsUpdatePhoto) | **Patch** /me/contactFolders({contactFolder-id})/contacts({contact-id})/photo | Update the navigation property photo in me



## MeContactFoldersContactsGetPhoto

> MicrosoftGraphProfilePhoto MeContactFoldersContactsGetPhoto(ctx, contactFolderId, contactId, optional)
Get photo from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***MeContactFoldersContactsGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeContactFoldersContactsGetPhotoOpts struct


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


## MeContactFoldersContactsUpdatePhoto

> MeContactFoldersContactsUpdatePhoto(ctx, contactFolderId, contactId, microsoftGraphProfilePhoto)
Update the navigation property photo in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

