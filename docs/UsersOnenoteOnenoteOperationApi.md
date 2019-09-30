# \UsersOnenoteOnenoteOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteCreateOperations**](UsersOnenoteOnenoteOperationApi.md#UsersOnenoteCreateOperations) | **Post** /users({user-id})/onenote/operations | Create new navigation property to operations for users
[**UsersOnenoteGetOperations**](UsersOnenoteOnenoteOperationApi.md#UsersOnenoteGetOperations) | **Get** /users({user-id})/onenote/operations({onenoteOperation-id}) | Get operations from users
[**UsersOnenoteListOperations**](UsersOnenoteOnenoteOperationApi.md#UsersOnenoteListOperations) | **Get** /users({user-id})/onenote/operations | Get operations from users
[**UsersOnenoteUpdateOperations**](UsersOnenoteOnenoteOperationApi.md#UsersOnenoteUpdateOperations) | **Patch** /users({user-id})/onenote/operations({onenoteOperation-id}) | Update the navigation property operations in users



## UsersOnenoteCreateOperations

> MicrosoftGraphOnenoteOperation UsersOnenoteCreateOperations(ctx, userId, microsoftGraphOnenoteOperation)
Create new navigation property to operations for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphOnenoteOperation** | [**MicrosoftGraphOnenoteOperation**](MicrosoftGraphOnenoteOperation.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenoteOperation**](microsoft.graph.onenoteOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteGetOperations

> MicrosoftGraphOnenoteOperation UsersOnenoteGetOperations(ctx, userId, onenoteOperationId, optional)
Get operations from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onenoteOperationId** | **string**| key: onenoteOperation-id of onenoteOperation | 
 **optional** | ***UsersOnenoteGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteGetOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenoteOperation**](microsoft.graph.onenoteOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteListOperations

> CollectionOfOnenoteOperation UsersOnenoteListOperations(ctx, userId, optional)
Get operations from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersOnenoteListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteListOperationsOpts struct


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

[**CollectionOfOnenoteOperation**](Collection of onenoteOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteUpdateOperations

> UsersOnenoteUpdateOperations(ctx, userId, onenoteOperationId, microsoftGraphOnenoteOperation)
Update the navigation property operations in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onenoteOperationId** | **string**| key: onenoteOperation-id of onenoteOperation | 
**microsoftGraphOnenoteOperation** | [**MicrosoftGraphOnenoteOperation**](MicrosoftGraphOnenoteOperation.md)| New navigation property values | 

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

