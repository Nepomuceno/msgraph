# \UsersOnenoteOnenoteResourceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteCreateResources**](UsersOnenoteOnenoteResourceApi.md#UsersOnenoteCreateResources) | **Post** /users({user-id})/onenote/resources | Create new navigation property to resources for users
[**UsersOnenoteGetResources**](UsersOnenoteOnenoteResourceApi.md#UsersOnenoteGetResources) | **Get** /users({user-id})/onenote/resources({onenoteResource-id}) | Get resources from users
[**UsersOnenoteListResources**](UsersOnenoteOnenoteResourceApi.md#UsersOnenoteListResources) | **Get** /users({user-id})/onenote/resources | Get resources from users
[**UsersOnenoteUpdateResources**](UsersOnenoteOnenoteResourceApi.md#UsersOnenoteUpdateResources) | **Patch** /users({user-id})/onenote/resources({onenoteResource-id}) | Update the navigation property resources in users



## UsersOnenoteCreateResources

> MicrosoftGraphOnenoteResource UsersOnenoteCreateResources(ctx, userId, microsoftGraphOnenoteResource)
Create new navigation property to resources for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphOnenoteResource** | [**MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenoteResource**](microsoft.graph.onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteGetResources

> MicrosoftGraphOnenoteResource UsersOnenoteGetResources(ctx, userId, onenoteResourceId, optional)
Get resources from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
 **optional** | ***UsersOnenoteGetResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteGetResourcesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenoteResource**](microsoft.graph.onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteListResources

> CollectionOfOnenoteResource UsersOnenoteListResources(ctx, userId, optional)
Get resources from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersOnenoteListResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteListResourcesOpts struct


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

[**CollectionOfOnenoteResource**](Collection of onenoteResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteUpdateResources

> UsersOnenoteUpdateResources(ctx, userId, onenoteResourceId, microsoftGraphOnenoteResource)
Update the navigation property resources in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onenoteResourceId** | **string**| key: onenoteResource-id of onenoteResource | 
**microsoftGraphOnenoteResource** | [**MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md)| New navigation property values | 

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

