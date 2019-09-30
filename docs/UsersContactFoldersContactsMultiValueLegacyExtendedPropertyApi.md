# \UsersContactFoldersContactsMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactFoldersContactsCreateMultiValueExtendedProperties**](UsersContactFoldersContactsMultiValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsCreateMultiValueExtendedProperties) | **Post** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersContactFoldersContactsGetMultiValueExtendedProperties**](UsersContactFoldersContactsMultiValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsGetMultiValueExtendedProperties) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from users
[**UsersContactFoldersContactsListMultiValueExtendedProperties**](UsersContactFoldersContactsMultiValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsListMultiValueExtendedProperties) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersContactFoldersContactsUpdateMultiValueExtendedProperties**](UsersContactFoldersContactsMultiValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsUpdateMultiValueExtendedProperties) | **Patch** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in users



## UsersContactFoldersContactsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersContactFoldersContactsCreateMultiValueExtendedProperties(ctx, userId, contactFolderId, contactId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersContactFoldersContactsGetMultiValueExtendedProperties(ctx, userId, contactFolderId, contactId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersContactFoldersContactsGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersContactsGetMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersContactFoldersContactsListMultiValueExtendedProperties(ctx, userId, contactFolderId, contactId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactFoldersContactsListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersContactsListMultiValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfMultiValueLegacyExtendedProperty**](Collection of multiValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsUpdateMultiValueExtendedProperties

> UsersContactFoldersContactsUpdateMultiValueExtendedProperties(ctx, userId, contactFolderId, contactId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
**microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)| New navigation property values | 

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

