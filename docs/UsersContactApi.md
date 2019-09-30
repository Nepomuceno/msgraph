# \UsersContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateContacts**](UsersContactApi.md#UsersCreateContacts) | **Post** /users({user-id})/contacts | Create new navigation property to contacts for users
[**UsersGetContacts**](UsersContactApi.md#UsersGetContacts) | **Get** /users({user-id})/contacts({contact-id}) | Get contacts from users
[**UsersListContacts**](UsersContactApi.md#UsersListContacts) | **Get** /users({user-id})/contacts | Get contacts from users
[**UsersUpdateContacts**](UsersContactApi.md#UsersUpdateContacts) | **Patch** /users({user-id})/contacts({contact-id}) | Update the navigation property contacts in users



## UsersCreateContacts

> MicrosoftGraphContact UsersCreateContacts(ctx, userId, microsoftGraphContact)
Create new navigation property to contacts for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersGetContacts

> MicrosoftGraphContact UsersGetContacts(ctx, userId, contactId, optional)
Get contacts from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersGetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## UsersListContacts

> CollectionOfContact UsersListContacts(ctx, userId, optional)
Get contacts from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

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


## UsersUpdateContacts

> UsersUpdateContacts(ctx, userId, contactId, microsoftGraphContact)
Update the navigation property contacts in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

