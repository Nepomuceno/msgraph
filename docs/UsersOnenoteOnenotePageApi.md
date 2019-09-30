# \UsersOnenoteOnenotePageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteCreatePages**](UsersOnenoteOnenotePageApi.md#UsersOnenoteCreatePages) | **Post** /users({user-id})/onenote/pages | Create new navigation property to pages for users
[**UsersOnenoteGetPages**](UsersOnenoteOnenotePageApi.md#UsersOnenoteGetPages) | **Get** /users({user-id})/onenote/pages({onenotePage-id}) | Get pages from users
[**UsersOnenoteListPages**](UsersOnenoteOnenotePageApi.md#UsersOnenoteListPages) | **Get** /users({user-id})/onenote/pages | Get pages from users
[**UsersOnenoteUpdatePages**](UsersOnenoteOnenotePageApi.md#UsersOnenoteUpdatePages) | **Patch** /users({user-id})/onenote/pages({onenotePage-id}) | Update the navigation property pages in users



## UsersOnenoteCreatePages

> MicrosoftGraphOnenotePage UsersOnenoteCreatePages(ctx, userId, microsoftGraphOnenotePage)
Create new navigation property to pages for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphOnenotePage** | [**MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenotePage**](microsoft.graph.onenotePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteGetPages

> MicrosoftGraphOnenotePage UsersOnenoteGetPages(ctx, userId, onenotePageId, optional)
Get pages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***UsersOnenoteGetPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteGetPagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenotePage**](microsoft.graph.onenotePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteListPages

> CollectionOfOnenotePage UsersOnenoteListPages(ctx, userId, optional)
Get pages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersOnenoteListPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteListPagesOpts struct


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

[**CollectionOfOnenotePage**](Collection of onenotePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOnenoteUpdatePages

> UsersOnenoteUpdatePages(ctx, userId, onenotePageId, microsoftGraphOnenotePage)
Update the navigation property pages in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**microsoftGraphOnenotePage** | [**MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md)| New navigation property values | 

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

