# \GroupsOnenoteOnenoteOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteCreateOperations**](GroupsOnenoteOnenoteOperationApi.md#GroupsOnenoteCreateOperations) | **Post** /groups({group-id})/onenote/operations | Create new navigation property to operations for groups
[**GroupsOnenoteGetOperations**](GroupsOnenoteOnenoteOperationApi.md#GroupsOnenoteGetOperations) | **Get** /groups({group-id})/onenote/operations({onenoteOperation-id}) | Get operations from groups
[**GroupsOnenoteListOperations**](GroupsOnenoteOnenoteOperationApi.md#GroupsOnenoteListOperations) | **Get** /groups({group-id})/onenote/operations | Get operations from groups
[**GroupsOnenoteUpdateOperations**](GroupsOnenoteOnenoteOperationApi.md#GroupsOnenoteUpdateOperations) | **Patch** /groups({group-id})/onenote/operations({onenoteOperation-id}) | Update the navigation property operations in groups



## GroupsOnenoteCreateOperations

> MicrosoftGraphOnenoteOperation GroupsOnenoteCreateOperations(ctx, groupId, microsoftGraphOnenoteOperation)
Create new navigation property to operations for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteGetOperations

> MicrosoftGraphOnenoteOperation GroupsOnenoteGetOperations(ctx, groupId, onenoteOperationId, optional)
Get operations from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**onenoteOperationId** | **string**| key: onenoteOperation-id of onenoteOperation | 
 **optional** | ***GroupsOnenoteGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteGetOperationsOpts struct


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


## GroupsOnenoteListOperations

> CollectionOfOnenoteOperation GroupsOnenoteListOperations(ctx, groupId, optional)
Get operations from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsOnenoteListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteListOperationsOpts struct


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


## GroupsOnenoteUpdateOperations

> GroupsOnenoteUpdateOperations(ctx, groupId, onenoteOperationId, microsoftGraphOnenoteOperation)
Update the navigation property operations in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

