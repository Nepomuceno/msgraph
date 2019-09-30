# \SitesOnenoteOnenoteSectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesOnenoteCreateSections**](SitesOnenoteOnenoteSectionApi.md#SitesOnenoteCreateSections) | **Post** /sites({site-id})/onenote/sections | Create new navigation property to sections for sites
[**SitesOnenoteGetSections**](SitesOnenoteOnenoteSectionApi.md#SitesOnenoteGetSections) | **Get** /sites({site-id})/onenote/sections({onenoteSection-id}) | Get sections from sites
[**SitesOnenoteListSections**](SitesOnenoteOnenoteSectionApi.md#SitesOnenoteListSections) | **Get** /sites({site-id})/onenote/sections | Get sections from sites
[**SitesOnenoteUpdateSections**](SitesOnenoteOnenoteSectionApi.md#SitesOnenoteUpdateSections) | **Patch** /sites({site-id})/onenote/sections({onenoteSection-id}) | Update the navigation property sections in sites



## SitesOnenoteCreateSections

> MicrosoftGraphOnenoteSection SitesOnenoteCreateSections(ctx, siteId, microsoftGraphOnenoteSection)
Create new navigation property to sections for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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


## SitesOnenoteGetSections

> MicrosoftGraphOnenoteSection SitesOnenoteGetSections(ctx, siteId, onenoteSectionId, optional)
Get sections from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***SitesOnenoteGetSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteGetSectionsOpts struct


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


## SitesOnenoteListSections

> CollectionOfOnenoteSection SitesOnenoteListSections(ctx, siteId, optional)
Get sections from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
 **optional** | ***SitesOnenoteListSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteListSectionsOpts struct


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


## SitesOnenoteUpdateSections

> SitesOnenoteUpdateSections(ctx, siteId, onenoteSectionId, microsoftGraphOnenoteSection)
Update the navigation property sections in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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

