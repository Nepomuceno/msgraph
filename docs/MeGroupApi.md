# \MeGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateJoinedTeams**](MeGroupApi.md#MeCreateJoinedTeams) | **Post** /me/joinedTeams | Create new navigation property to joinedTeams for me
[**MeGetJoinedTeams**](MeGroupApi.md#MeGetJoinedTeams) | **Get** /me/joinedTeams({group-id}) | Get joinedTeams from me
[**MeListJoinedTeams**](MeGroupApi.md#MeListJoinedTeams) | **Get** /me/joinedTeams | Get joinedTeams from me
[**MeUpdateJoinedTeams**](MeGroupApi.md#MeUpdateJoinedTeams) | **Patch** /me/joinedTeams({group-id}) | Update the navigation property joinedTeams in me



## MeCreateJoinedTeams

> MicrosoftGraphGroup MeCreateJoinedTeams(ctx, microsoftGraphGroup)
Create new navigation property to joinedTeams for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeGetJoinedTeams

> MicrosoftGraphGroup MeGetJoinedTeams(ctx, groupId, optional)
Get joinedTeams from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***MeGetJoinedTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetJoinedTeamsOpts struct


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


## MeListJoinedTeams

> CollectionOfGroup MeListJoinedTeams(ctx, optional)
Get joinedTeams from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListJoinedTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListJoinedTeamsOpts struct


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


## MeUpdateJoinedTeams

> MeUpdateJoinedTeams(ctx, groupId, microsoftGraphGroup)
Update the navigation property joinedTeams in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

