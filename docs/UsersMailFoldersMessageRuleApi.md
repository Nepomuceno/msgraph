# \UsersMailFoldersMessageRuleApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMailFoldersCreateMessageRules**](UsersMailFoldersMessageRuleApi.md#UsersMailFoldersCreateMessageRules) | **Post** /users({user-id})/mailFolders({mailFolder-id})/messageRules | Create new navigation property to messageRules for users
[**UsersMailFoldersGetMessageRules**](UsersMailFoldersMessageRuleApi.md#UsersMailFoldersGetMessageRules) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messageRules({messageRule-id}) | Get messageRules from users
[**UsersMailFoldersListMessageRules**](UsersMailFoldersMessageRuleApi.md#UsersMailFoldersListMessageRules) | **Get** /users({user-id})/mailFolders({mailFolder-id})/messageRules | Get messageRules from users
[**UsersMailFoldersUpdateMessageRules**](UsersMailFoldersMessageRuleApi.md#UsersMailFoldersUpdateMessageRules) | **Patch** /users({user-id})/mailFolders({mailFolder-id})/messageRules({messageRule-id}) | Update the navigation property messageRules in users



## UsersMailFoldersCreateMessageRules

> MicrosoftGraphMessageRule UsersMailFoldersCreateMessageRules(ctx, userId, mailFolderId, microsoftGraphMessageRule)
Create new navigation property to messageRules for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**microsoftGraphMessageRule** | [**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md)| New navigation property | 

### Return type

[**MicrosoftGraphMessageRule**](microsoft.graph.messageRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersGetMessageRules

> MicrosoftGraphMessageRule UsersMailFoldersGetMessageRules(ctx, userId, mailFolderId, messageRuleId, optional)
Get messageRules from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageRuleId** | **string**| key: messageRule-id of messageRule | 
 **optional** | ***UsersMailFoldersGetMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersGetMessageRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphMessageRule**](microsoft.graph.messageRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersListMessageRules

> CollectionOfMessageRule UsersMailFoldersListMessageRules(ctx, userId, mailFolderId, optional)
Get messageRules from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***UsersMailFoldersListMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersMailFoldersListMessageRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfMessageRule**](Collection of messageRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMailFoldersUpdateMessageRules

> UsersMailFoldersUpdateMessageRules(ctx, userId, mailFolderId, messageRuleId, microsoftGraphMessageRule)
Update the navigation property messageRules in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageRuleId** | **string**| key: messageRule-id of messageRule | 
**microsoftGraphMessageRule** | [**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md)| New navigation property values | 

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

