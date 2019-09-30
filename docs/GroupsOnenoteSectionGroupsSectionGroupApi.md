# \GroupsOnenoteSectionGroupsSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteSectionGroupsCreateSectionGroups**](GroupsOnenoteSectionGroupsSectionGroupApi.md#GroupsOnenoteSectionGroupsCreateSectionGroups) | **Post** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Create new navigation property to sectionGroups for groups
[**GroupsOnenoteSectionGroupsGetParentSectionGroup**](GroupsOnenoteSectionGroupsSectionGroupApi.md#GroupsOnenoteSectionGroupsGetParentSectionGroup) | **Get** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Get parentSectionGroup from groups
[**GroupsOnenoteSectionGroupsGetSectionGroups**](GroupsOnenoteSectionGroupsSectionGroupApi.md#GroupsOnenoteSectionGroupsGetSectionGroups) | **Get** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Get sectionGroups from groups
[**GroupsOnenoteSectionGroupsListSectionGroups**](GroupsOnenoteSectionGroupsSectionGroupApi.md#GroupsOnenoteSectionGroupsListSectionGroups) | **Get** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Get sectionGroups from groups
[**GroupsOnenoteSectionGroupsUpdateParentSectionGroup**](GroupsOnenoteSectionGroupsSectionGroupApi.md#GroupsOnenoteSectionGroupsUpdateParentSectionGroup) | **Patch** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Update the navigation property parentSectionGroup in groups
[**GroupsOnenoteSectionGroupsUpdateSectionGroups**](GroupsOnenoteSectionGroupsSectionGroupApi.md#GroupsOnenoteSectionGroupsUpdateSectionGroups) | **Patch** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Update the navigation property sectionGroups in groups



## GroupsOnenoteSectionGroupsCreateSectionGroups

> MicrosoftGraphSectionGroup GroupsOnenoteSectionGroupsCreateSectionGroups(ctx, groupId, sectionGroupId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteSectionGroupsGetParentSectionGroup

> MicrosoftGraphSectionGroup GroupsOnenoteSectionGroupsGetParentSectionGroup(ctx, groupId, sectionGroupId, optional)
Get parentSectionGroup from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteSectionGroupsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteSectionGroupsGetParentSectionGroupOpts struct


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


## GroupsOnenoteSectionGroupsGetSectionGroups

> MicrosoftGraphSectionGroup GroupsOnenoteSectionGroupsGetSectionGroups(ctx, groupId, sectionGroupId, sectionGroupId1, optional)
Get sectionGroups from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteSectionGroupsGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteSectionGroupsGetSectionGroupsOpts struct


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


## GroupsOnenoteSectionGroupsListSectionGroups

> CollectionOfSectionGroup GroupsOnenoteSectionGroupsListSectionGroups(ctx, groupId, sectionGroupId, optional)
Get sectionGroups from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteSectionGroupsListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteSectionGroupsListSectionGroupsOpts struct


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


## GroupsOnenoteSectionGroupsUpdateParentSectionGroup

> GroupsOnenoteSectionGroupsUpdateParentSectionGroup(ctx, groupId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteSectionGroupsUpdateSectionGroups

> GroupsOnenoteSectionGroupsUpdateSectionGroups(ctx, groupId, sectionGroupId, sectionGroupId1, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

