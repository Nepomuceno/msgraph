# \UsersGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateJoinedTeams**](UsersGroupApi.md#UsersCreateJoinedTeams) | **Post** /users({user-id})/joinedTeams | Create new navigation property to joinedTeams for users
[**UsersGetJoinedTeams**](UsersGroupApi.md#UsersGetJoinedTeams) | **Get** /users({user-id})/joinedTeams({group-id}) | Get joinedTeams from users
[**UsersListJoinedTeams**](UsersGroupApi.md#UsersListJoinedTeams) | **Get** /users({user-id})/joinedTeams | Get joinedTeams from users
[**UsersUpdateJoinedTeams**](UsersGroupApi.md#UsersUpdateJoinedTeams) | **Patch** /users({user-id})/joinedTeams({group-id}) | Update the navigation property joinedTeams in users



## UsersCreateJoinedTeams

> MicrosoftGraphGroup UsersCreateJoinedTeams(ctx, userId, microsoftGraphGroup)
Create new navigation property to joinedTeams for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphGroup** | [**MicrosoftGraphGroup**](MicrosoftGraphGroup.md)| New navigation property | 

### Return type

[**MicrosoftGraphGroup**](microsoft.graph.group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetJoinedTeams

> MicrosoftGraphGroup UsersGetJoinedTeams(ctx, userId, groupId, optional)
Get joinedTeams from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**groupId** | **string**| key: group-id of group | 
 **optional** | ***UsersGetJoinedTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetJoinedTeamsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphGroup**](microsoft.graph.group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListJoinedTeams

> CollectionOfGroup UsersListJoinedTeams(ctx, userId, optional)
Get joinedTeams from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersListJoinedTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersListJoinedTeamsOpts struct


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

[**CollectionOfGroup**](Collection of group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateJoinedTeams

> UsersUpdateJoinedTeams(ctx, userId, groupId, microsoftGraphGroup)
Update the navigation property joinedTeams in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**groupId** | **string**| key: group-id of group | 
**microsoftGraphGroup** | [**MicrosoftGraphGroup**](MicrosoftGraphGroup.md)| New navigation property values | 

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

