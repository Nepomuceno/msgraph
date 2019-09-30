# \MeOnenoteOnenotePageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOnenoteCreatePages**](MeOnenoteOnenotePageApi.md#MeOnenoteCreatePages) | **Post** /me/onenote/pages | Create new navigation property to pages for me
[**MeOnenoteGetPages**](MeOnenoteOnenotePageApi.md#MeOnenoteGetPages) | **Get** /me/onenote/pages({onenotePage-id}) | Get pages from me
[**MeOnenoteListPages**](MeOnenoteOnenotePageApi.md#MeOnenoteListPages) | **Get** /me/onenote/pages | Get pages from me
[**MeOnenoteUpdatePages**](MeOnenoteOnenotePageApi.md#MeOnenoteUpdatePages) | **Patch** /me/onenote/pages({onenotePage-id}) | Update the navigation property pages in me



## MeOnenoteCreatePages

> MicrosoftGraphOnenotePage MeOnenoteCreatePages(ctx, microsoftGraphOnenotePage)
Create new navigation property to pages for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeOnenoteGetPages

> MicrosoftGraphOnenotePage MeOnenoteGetPages(ctx, onenotePageId, optional)
Get pages from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***MeOnenoteGetPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteGetPagesOpts struct


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


## MeOnenoteListPages

> CollectionOfOnenotePage MeOnenoteListPages(ctx, optional)
Get pages from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOnenoteListPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteListPagesOpts struct


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


## MeOnenoteUpdatePages

> MeOnenoteUpdatePages(ctx, onenotePageId, microsoftGraphOnenotePage)
Update the navigation property pages in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

