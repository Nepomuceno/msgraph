# \MeOnenoteSectionGroupsSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOnenoteSectionGroupsCreateSectionGroups**](MeOnenoteSectionGroupsSectionGroupApi.md#MeOnenoteSectionGroupsCreateSectionGroups) | **Post** /me/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Create new navigation property to sectionGroups for me
[**MeOnenoteSectionGroupsGetParentSectionGroup**](MeOnenoteSectionGroupsSectionGroupApi.md#MeOnenoteSectionGroupsGetParentSectionGroup) | **Get** /me/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Get parentSectionGroup from me
[**MeOnenoteSectionGroupsGetSectionGroups**](MeOnenoteSectionGroupsSectionGroupApi.md#MeOnenoteSectionGroupsGetSectionGroups) | **Get** /me/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Get sectionGroups from me
[**MeOnenoteSectionGroupsListSectionGroups**](MeOnenoteSectionGroupsSectionGroupApi.md#MeOnenoteSectionGroupsListSectionGroups) | **Get** /me/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Get sectionGroups from me
[**MeOnenoteSectionGroupsUpdateParentSectionGroup**](MeOnenoteSectionGroupsSectionGroupApi.md#MeOnenoteSectionGroupsUpdateParentSectionGroup) | **Patch** /me/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Update the navigation property parentSectionGroup in me
[**MeOnenoteSectionGroupsUpdateSectionGroups**](MeOnenoteSectionGroupsSectionGroupApi.md#MeOnenoteSectionGroupsUpdateSectionGroups) | **Patch** /me/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Update the navigation property sectionGroups in me



## MeOnenoteSectionGroupsCreateSectionGroups

> MicrosoftGraphSectionGroup MeOnenoteSectionGroupsCreateSectionGroups(ctx, sectionGroupId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
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


## MeOnenoteSectionGroupsGetParentSectionGroup

> MicrosoftGraphSectionGroup MeOnenoteSectionGroupsGetParentSectionGroup(ctx, sectionGroupId, optional)
Get parentSectionGroup from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***MeOnenoteSectionGroupsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteSectionGroupsGetParentSectionGroupOpts struct


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


## MeOnenoteSectionGroupsGetSectionGroups

> MicrosoftGraphSectionGroup MeOnenoteSectionGroupsGetSectionGroups(ctx, sectionGroupId, sectionGroupId1, optional)
Get sectionGroups from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***MeOnenoteSectionGroupsGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteSectionGroupsGetSectionGroupsOpts struct


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


## MeOnenoteSectionGroupsListSectionGroups

> CollectionOfSectionGroup MeOnenoteSectionGroupsListSectionGroups(ctx, sectionGroupId, optional)
Get sectionGroups from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***MeOnenoteSectionGroupsListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOnenoteSectionGroupsListSectionGroupsOpts struct


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


## MeOnenoteSectionGroupsUpdateParentSectionGroup

> MeOnenoteSectionGroupsUpdateParentSectionGroup(ctx, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in me

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


## MeOnenoteSectionGroupsUpdateSectionGroups

> MeOnenoteSectionGroupsUpdateSectionGroups(ctx, sectionGroupId, sectionGroupId1, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
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

