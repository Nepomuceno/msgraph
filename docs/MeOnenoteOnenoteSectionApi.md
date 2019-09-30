# \MeOnenoteOnenoteSectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOnenoteCreateSections**](MeOnenoteOnenoteSectionApi.md#MeOnenoteCreateSections) | **Post** /me/onenote/sections | Create new navigation property to sections for me
[**MeOnenoteGetSections**](MeOnenoteOnenoteSectionApi.md#MeOnenoteGetSections) | **Get** /me/onenote/sections({onenoteSection-id}) | Get sections from me
[**MeOnenoteListSections**](MeOnenoteOnenoteSectionApi.md#MeOnenoteListSections) | **Get** /me/onenote/sections | Get sections from me
[**MeOnenoteUpdateSections**](MeOnenoteOnenoteSectionApi.md#MeOnenoteUpdateSections) | **Patch** /me/onenote/sections({onenoteSection-id}) | Update the navigation property sections in me



## MeOnenoteCreateSections

> MicrosoftGraphOnenoteSection MeOnenoteCreateSections(ctx, microsoftGraphOnenoteSection)
Create new navigation property to sections for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOnenoteSection** | [**MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md)| New navigation property | 

### Return type

[**MicrosoftGraphOnenoteSection**](microsoft.graph.onenoteSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteGetSections

> MicrosoftGraphOnenoteSection MeOnenoteGetSections(ctx, onenoteSectionId, optional)
Get sections from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***MeOnenoteGetSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteGetSectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphOnenoteSection**](microsoft.graph.onenoteSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteListSections

> CollectionOfOnenoteSection MeOnenoteListSections(ctx, optional)
Get sections from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOnenoteListSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteListSectionsOpts struct


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

[**CollectionOfOnenoteSection**](Collection of onenoteSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteUpdateSections

> MeOnenoteUpdateSections(ctx, onenoteSectionId, microsoftGraphOnenoteSection)
Update the navigation property sections in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**microsoftGraphOnenoteSection** | [**MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md)| New navigation property values | 

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

