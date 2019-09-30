# \GroupsOnenoteOnenoteSectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteCreateSections**](GroupsOnenoteOnenoteSectionApi.md#GroupsOnenoteCreateSections) | **Post** /groups({group-id})/onenote/sections | Create new navigation property to sections for groups
[**GroupsOnenoteGetSections**](GroupsOnenoteOnenoteSectionApi.md#GroupsOnenoteGetSections) | **Get** /groups({group-id})/onenote/sections({onenoteSection-id}) | Get sections from groups
[**GroupsOnenoteListSections**](GroupsOnenoteOnenoteSectionApi.md#GroupsOnenoteListSections) | **Get** /groups({group-id})/onenote/sections | Get sections from groups
[**GroupsOnenoteUpdateSections**](GroupsOnenoteOnenoteSectionApi.md#GroupsOnenoteUpdateSections) | **Patch** /groups({group-id})/onenote/sections({onenoteSection-id}) | Update the navigation property sections in groups



## GroupsOnenoteCreateSections

> MicrosoftGraphOnenoteSection GroupsOnenoteCreateSections(ctx, groupId, microsoftGraphOnenoteSection)
Create new navigation property to sections for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteGetSections

> MicrosoftGraphOnenoteSection GroupsOnenoteGetSections(ctx, groupId, onenoteSectionId, optional)
Get sections from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***GroupsOnenoteGetSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteGetSectionsOpts struct


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


## GroupsOnenoteListSections

> CollectionOfOnenoteSection GroupsOnenoteListSections(ctx, groupId, optional)
Get sections from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
 **optional** | ***GroupsOnenoteListSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteListSectionsOpts struct


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


## GroupsOnenoteUpdateSections

> GroupsOnenoteUpdateSections(ctx, groupId, onenoteSectionId, microsoftGraphOnenoteSection)
Update the navigation property sections in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

