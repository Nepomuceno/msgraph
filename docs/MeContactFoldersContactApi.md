# \MeContactFoldersContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeContactFoldersCreateContacts**](MeContactFoldersContactApi.md#MeContactFoldersCreateContacts) | **Post** /me/contactFolders({contactFolder-id})/contacts | Create new navigation property to contacts for me
[**MeContactFoldersGetContacts**](MeContactFoldersContactApi.md#MeContactFoldersGetContacts) | **Get** /me/contactFolders({contactFolder-id})/contacts({contact-id}) | Get contacts from me
[**MeContactFoldersListContacts**](MeContactFoldersContactApi.md#MeContactFoldersListContacts) | **Get** /me/contactFolders({contactFolder-id})/contacts | Get contacts from me
[**MeContactFoldersUpdateContacts**](MeContactFoldersContactApi.md#MeContactFoldersUpdateContacts) | **Patch** /me/contactFolders({contactFolder-id})/contacts({contact-id}) | Update the navigation property contacts in me



## MeContactFoldersCreateContacts

> MicrosoftGraphContact MeContactFoldersCreateContacts(ctx, contactFolderId, microsoftGraphContact)
Create new navigation property to contacts for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeContactFoldersGetContacts

> MicrosoftGraphContact MeContactFoldersGetContacts(ctx, contactFolderId, contactId, optional)
Get contacts from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***MeContactFoldersGetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeContactFoldersGetContactsOpts struct


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


## MeContactFoldersListContacts

> CollectionOfContact MeContactFoldersListContacts(ctx, contactFolderId, optional)
Get contacts from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
 **optional** | ***MeContactFoldersListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeContactFoldersListContactsOpts struct


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


## MeContactFoldersUpdateContacts

> MeContactFoldersUpdateContacts(ctx, contactFolderId, contactId, microsoftGraphContact)
Update the navigation property contacts in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

