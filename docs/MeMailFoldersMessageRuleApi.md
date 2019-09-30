# \MeMailFoldersMessageRuleApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeMailFoldersCreateMessageRules**](MeMailFoldersMessageRuleApi.md#MeMailFoldersCreateMessageRules) | **Post** /me/mailFolders({mailFolder-id})/messageRules | Create new navigation property to messageRules for me
[**MeMailFoldersGetMessageRules**](MeMailFoldersMessageRuleApi.md#MeMailFoldersGetMessageRules) | **Get** /me/mailFolders({mailFolder-id})/messageRules({messageRule-id}) | Get messageRules from me
[**MeMailFoldersListMessageRules**](MeMailFoldersMessageRuleApi.md#MeMailFoldersListMessageRules) | **Get** /me/mailFolders({mailFolder-id})/messageRules | Get messageRules from me
[**MeMailFoldersUpdateMessageRules**](MeMailFoldersMessageRuleApi.md#MeMailFoldersUpdateMessageRules) | **Patch** /me/mailFolders({mailFolder-id})/messageRules({messageRule-id}) | Update the navigation property messageRules in me



## MeMailFoldersCreateMessageRules

> MicrosoftGraphMessageRule MeMailFoldersCreateMessageRules(ctx, mailFolderId, microsoftGraphMessageRule)
Create new navigation property to messageRules for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeMailFoldersGetMessageRules

> MicrosoftGraphMessageRule MeMailFoldersGetMessageRules(ctx, mailFolderId, messageRuleId, optional)
Get messageRules from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
**messageRuleId** | **string**| key: messageRule-id of messageRule | 
 **optional** | ***MeMailFoldersGetMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersGetMessageRulesOpts struct


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


## MeMailFoldersListMessageRules

> CollectionOfMessageRule MeMailFoldersListMessageRules(ctx, mailFolderId, optional)
Get messageRules from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string**| key: mailFolder-id of mailFolder | 
 **optional** | ***MeMailFoldersListMessageRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeMailFoldersListMessageRulesOpts struct


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


## MeMailFoldersUpdateMessageRules

> MeMailFoldersUpdateMessageRules(ctx, mailFolderId, messageRuleId, microsoftGraphMessageRule)
Update the navigation property messageRules in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

