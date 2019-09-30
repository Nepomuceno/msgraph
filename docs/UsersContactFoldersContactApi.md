# \UsersContactFoldersContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactFoldersCreateContacts**](UsersContactFoldersContactApi.md#UsersContactFoldersCreateContacts) | **Post** /users({user-id})/contactFolders({contactFolder-id})/contacts | Create new navigation property to contacts for users
[**UsersContactFoldersGetContacts**](UsersContactFoldersContactApi.md#UsersContactFoldersGetContacts) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id}) | Get contacts from users
[**UsersContactFoldersListContacts**](UsersContactFoldersContactApi.md#UsersContactFoldersListContacts) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts | Get contacts from users
[**UsersContactFoldersUpdateContacts**](UsersContactFoldersContactApi.md#UsersContactFoldersUpdateContacts) | **Patch** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id}) | Update the navigation property contacts in users



## UsersContactFoldersCreateContacts

> MicrosoftGraphContact UsersContactFoldersCreateContacts(ctx, userId, contactFolderId, microsoftGraphContact)
Create new navigation property to contacts for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**microsoftGraphContact** | [**MicrosoftGraphContact**](MicrosoftGraphContact.md)| New navigation property | 

### Return type

[**MicrosoftGraphContact**](microsoft.graph.contact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersGetContacts

> MicrosoftGraphContact UsersContactFoldersGetContacts(ctx, userId, contactFolderId, contactId, optional)
Get contacts from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactFoldersGetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersGetContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphContact**](microsoft.graph.contact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersListContacts

> CollectionOfContact UsersContactFoldersListContacts(ctx, userId, contactFolderId, optional)
Get contacts from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***UsersContactFoldersListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersListContactsOpts struct


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

[**CollectionOfContact**](Collection of contact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersUpdateContacts

> UsersContactFoldersUpdateContacts(ctx, userId, contactFolderId, contactId, microsoftGraphContact)
Update the navigation property contacts in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphContact** | [**MicrosoftGraphContact**](MicrosoftGraphContact.md)| New navigation property values | 

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

