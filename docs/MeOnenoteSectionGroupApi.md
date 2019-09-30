# \MeOnenoteSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOnenoteCreateSectionGroups**](MeOnenoteSectionGroupApi.md#MeOnenoteCreateSectionGroups) | **Post** /me/onenote/sectionGroups | Create new navigation property to sectionGroups for me
[**MeOnenoteGetSectionGroups**](MeOnenoteSectionGroupApi.md#MeOnenoteGetSectionGroups) | **Get** /me/onenote/sectionGroups({sectionGroup-id}) | Get sectionGroups from me
[**MeOnenoteListSectionGroups**](MeOnenoteSectionGroupApi.md#MeOnenoteListSectionGroups) | **Get** /me/onenote/sectionGroups | Get sectionGroups from me
[**MeOnenoteUpdateSectionGroups**](MeOnenoteSectionGroupApi.md#MeOnenoteUpdateSectionGroups) | **Patch** /me/onenote/sectionGroups({sectionGroup-id}) | Update the navigation property sectionGroups in me



## MeOnenoteCreateSectionGroups

> MicrosoftGraphSectionGroup MeOnenoteCreateSectionGroups(ctx, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSectionGroup** | [**MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md)| New navigation property | 

### Return type

[**MicrosoftGraphSectionGroup**](microsoft.graph.sectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteGetSectionGroups

> MicrosoftGraphSectionGroup MeOnenoteGetSectionGroups(ctx, sectionGroupId, optional)
Get sectionGroups from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***MeOnenoteGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteGetSectionGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSectionGroup**](microsoft.graph.sectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteListSectionGroups

> CollectionOfSectionGroup MeOnenoteListSectionGroups(ctx, optional)
Get sectionGroups from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOnenoteListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteListSectionGroupsOpts struct


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

[**CollectionOfSectionGroup**](Collection of sectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOnenoteUpdateSectionGroups

> MeOnenoteUpdateSectionGroups(ctx, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**microsoftGraphSectionGroup** | [**MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md)| New navigation property values | 

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

