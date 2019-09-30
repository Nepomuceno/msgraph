# \UsersContactFoldersContactsSingleValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactFoldersContactsCreateSingleValueExtendedProperties**](UsersContactFoldersContactsSingleValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsCreateSingleValueExtendedProperties) | **Post** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersContactFoldersContactsGetSingleValueExtendedProperties**](UsersContactFoldersContactsSingleValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsGetSingleValueExtendedProperties) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Get singleValueExtendedProperties from users
[**UsersContactFoldersContactsListSingleValueExtendedProperties**](UsersContactFoldersContactsSingleValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsListSingleValueExtendedProperties) | **Get** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersContactFoldersContactsUpdateSingleValueExtendedProperties**](UsersContactFoldersContactsSingleValueLegacyExtendedPropertyApi.md#UsersContactFoldersContactsUpdateSingleValueExtendedProperties) | **Patch** /users({user-id})/contactFolders({contactFolder-id})/contacts({contact-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Update the navigation property singleValueExtendedProperties in users



## UsersContactFoldersContactsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersContactFoldersContactsCreateSingleValueExtendedProperties(ctx, userId, contactFolderId, contactId, microsoftGraphSingleValueLegacyExtendedProperty)
Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersContactFoldersContactsGetSingleValueExtendedProperties(ctx, userId, contactFolderId, contactId, singleValueLegacyExtendedPropertyId, optional)
Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersContactFoldersContactsGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersContactsGetSingleValueExtendedPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersContactFoldersContactsListSingleValueExtendedProperties(ctx, userId, contactFolderId, contactId, optional)
Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***UsersContactFoldersContactsListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersContactFoldersContactsListSingleValueExtendedPropertiesOpts struct


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

[**CollectionOfSingleValueLegacyExtendedProperty**](Collection of singleValueLegacyExtendedProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactFoldersContactsUpdateSingleValueExtendedProperties

> UsersContactFoldersContactsUpdateSingleValueExtendedProperties(ctx, userId, contactFolderId, contactId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)
Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**contactFolderId** | **string**| key: contactFolder-id of contactFolder | 
**contactId** | **string**| key: contact-id of contact | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
**microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)| New navigation property values | 

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

