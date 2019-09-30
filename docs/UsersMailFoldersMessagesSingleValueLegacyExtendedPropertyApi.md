# \UsersMailFoldersMessagesSingleValueLegacyExtendedPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMailFoldersMessagesCreateSingleValueExtendedProperties**](UsersMailFoldersMessagesSingleValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesCreateSingleValueExtendedProperties) | **Post** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMailFoldersMessagesGetSingleValueExtendedProperties**](UsersMailFoldersMessagesSingleValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesGetSingleValueExtendedProperties) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesListSingleValueExtendedProperties**](UsersMailFoldersMessagesSingleValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesListSingleValueExtendedProperties) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesUpdateSingleValueExtendedProperties**](UsersMailFoldersMessagesSingleValueLegacyExtendedPropertyApi.md#UsersMailFoldersMessagesUpdateSingleValueExtendedProperties) | **Patch** /users({user-id})/mailFolders({mailFolder-id})/messages({message-id})/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id}) | Update the navigation property singleValueExtendedProperties in users



## UsersMailFoldersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersMessagesCreateSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, microsoftGraphSingleValueLegacyExtendedProperty)
Create new navigation property to singleValueExtendedProperties for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
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


## UsersMailFoldersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersMessagesGetSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId, optional)
Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
**singleValueLegacyExtendedPropertyId** | **string**| key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty | 
 **optional** | ***UsersMailFoldersMessagesGetSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesGetSingleValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMailFoldersMessagesListSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, optional)
Get singleValueExtendedProperties from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
 **optional** | ***UsersMailFoldersMessagesListSingleValueExtendedPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersMessagesListSingleValueExtendedPropertiesOpts struct


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


## UsersMailFoldersMessagesUpdateSingleValueExtendedProperties

> UsersMailFoldersMessagesUpdateSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId, microsoftGraphSingleValueLegacyExtendedProperty)
Update the navigation property singleValueExtendedProperties in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageId** | **string**| key: message-id of message | 
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

