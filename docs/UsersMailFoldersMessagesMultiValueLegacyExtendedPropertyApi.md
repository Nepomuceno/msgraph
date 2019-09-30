# \UsersMailFoldersMessagesMultiValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMailFoldersMessagesCreateMultiValueExtendedProperties**](UsersMailFoldersMessagesMultiValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesCreateMultiValueExtendedProperties) | **Post** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMailFoldersMessagesGetMultiValueExtendedProperties**](UsersMailFoldersMessagesMultiValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesGetMultiValueExtendedProperties) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Get multiValueExtendedProperties from users
[**UsersMailFoldersMessagesListMultiValueExtendedProperties**](UsersMailFoldersMessagesMultiValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesListMultiValueExtendedProperties) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMailFoldersMessagesUpdateMultiValueExtendedProperties**](UsersMailFoldersMessagesMultiValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesUpdateMultiValueExtendedProperties) | **Patch** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/multiValueExtendedProperties({multiValueLegacyExtendedProperty-id}) | Update the navigation property multiValueExtendedProperties in users



## UsersMailFoldersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersMessagesCreateMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, microsoftGraphMultiValueLegacyExtendedProperty)
Create new navigation property to multiValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMailFoldersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersMessagesGetMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**multiValueLegacyExtendedPropertyId** | **string**| key: multiValueLegacyExtendedProperty-id of multiValueLegacyExtendedProperty | 
 **optional** | ***UsersMailFoldersMessagesGetMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetMultiValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMailFoldersMessagesListMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, optional)
Get multiValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListMultiValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListMultiValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesUpdateMultiValueExtendedProperties

> UsersMailFoldersMessagesUpdateMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId, microsoftGraphMultiValueLegacyExtendedProperty)
Update the navigation property multiValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
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

