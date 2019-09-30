# \SitesOnenoteSectionGroupsSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesOnenoteSectionGroupsCreateSectionGroups**](SitesOnenoteSectionGroupsSectionGroupApi.md#SitesOnenoteSectionGroupsCreateSectionGroups) | **Post** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Create new navigation property to sectionGroups for sites
[**SitesOnenoteSectionGroupsGetParentSectionGroup**](SitesOnenoteSectionGroupsSectionGroupApi.md#SitesOnenoteSectionGroupsGetParentSectionGroup) | **Get** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Get parentSectionGroup from sites
[**SitesOnenoteSectionGroupsGetSectionGroups**](SitesOnenoteSectionGroupsSectionGroupApi.md#SitesOnenoteSectionGroupsGetSectionGroups) | **Get** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Get sectionGroups from sites
[**SitesOnenoteSectionGroupsListSectionGroups**](SitesOnenoteSectionGroupsSectionGroupApi.md#SitesOnenoteSectionGroupsListSectionGroups) | **Get** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Get sectionGroups from sites
[**SitesOnenoteSectionGroupsUpdateParentSectionGroup**](SitesOnenoteSectionGroupsSectionGroupApi.md#SitesOnenoteSectionGroupsUpdateParentSectionGroup) | **Patch** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Update the navigation property parentSectionGroup in sites
[**SitesOnenoteSectionGroupsUpdateSectionGroups**](SitesOnenoteSectionGroupsSectionGroupApi.md#SitesOnenoteSectionGroupsUpdateSectionGroups) | **Patch** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Update the navigation property sectionGroups in sites



## SitesOnenoteSectionGroupsCreateSectionGroups

> MicrosoftGraphSectionGroup SitesOnenoteSectionGroupsCreateSectionGroups(ctx, siteId, sectionGroupId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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


## SitesOnenoteSectionGroupsGetParentSectionGroup

> MicrosoftGraphSectionGroup SitesOnenoteSectionGroupsGetParentSectionGroup(ctx, siteId, sectionGroupId, optional)
Get parentSectionGroup from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***SitesOnenoteSectionGroupsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteSectionGroupsGetParentSectionGroupOpts struct


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


## SitesOnenoteSectionGroupsGetSectionGroups

> MicrosoftGraphSectionGroup SitesOnenoteSectionGroupsGetSectionGroups(ctx, siteId, sectionGroupId, sectionGroupId1, optional)
Get sectionGroups from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***SitesOnenoteSectionGroupsGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteSectionGroupsGetSectionGroupsOpts struct


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


## SitesOnenoteSectionGroupsListSectionGroups

> CollectionOfSectionGroup SitesOnenoteSectionGroupsListSectionGroups(ctx, siteId, sectionGroupId, optional)
Get sectionGroups from sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***SitesOnenoteSectionGroupsListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SitesOnenoteSectionGroupsListSectionGroupsOpts struct


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


## SitesOnenoteSectionGroupsUpdateParentSectionGroup

> SitesOnenoteSectionGroupsUpdateParentSectionGroup(ctx, siteId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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


## SitesOnenoteSectionGroupsUpdateSectionGroups

> SitesOnenoteSectionGroupsUpdateSectionGroups(ctx, siteId, sectionGroupId, sectionGroupId1, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in sites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
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

