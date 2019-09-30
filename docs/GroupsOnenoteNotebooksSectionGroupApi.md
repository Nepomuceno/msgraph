# \GroupsOnenoteNotebooksSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsOnenoteNotebooksCreateSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksCreateSectionGroups) | **Post** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups | Create new navigation property to sectionGroups for groups
[**GroupsOnenoteNotebooksGetSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksGetSectionGroups) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id}) | Get sectionGroups from groups
[**GroupsOnenoteNotebooksListSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksListSectionGroups) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups | Get sectionGroups from groups
[**GroupsOnenoteNotebooksSectionGroupsCreateSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsCreateSectionGroups) | **Post** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups | Create new navigation property to sectionGroups for groups
[**GroupsOnenoteNotebooksSectionGroupsCreateSections**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsCreateSections) | **Post** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections | Create new navigation property to sections for groups
[**GroupsOnenoteNotebooksSectionGroupsGetParentNotebook**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsGetParentNotebook) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentNotebook | Get parentNotebook from groups
[**GroupsOnenoteNotebooksSectionGroupsGetParentSectionGroup**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsGetParentSectionGroup) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentSectionGroup | Get parentSectionGroup from groups
[**GroupsOnenoteNotebooksSectionGroupsGetSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsGetSectionGroups) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Get sectionGroups from groups
[**GroupsOnenoteNotebooksSectionGroupsGetSections**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsGetSections) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id}) | Get sections from groups
[**GroupsOnenoteNotebooksSectionGroupsListSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsListSectionGroups) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups | Get sectionGroups from groups
[**GroupsOnenoteNotebooksSectionGroupsListSections**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsListSections) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections | Get sections from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsCreatePages**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsCreatePages) | **Post** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages | Create new navigation property to pages for groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsGetPages**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsGetPages) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id}) | Get pages from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsGetParentNotebook**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsGetParentNotebook) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentNotebook | Get parentNotebook from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentSectionGroup | Get parentSectionGroup from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsListPages**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsListPages) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages | Get pages from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentNotebook | Get parentNotebook from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentSection | Get parentSection from groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentNotebook | Update the navigation property parentNotebook in groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentSection | Update the navigation property parentSection in groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsUpdatePages**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsUpdatePages) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id}) | Update the navigation property pages in groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentNotebook | Update the navigation property parentNotebook in groups
[**GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentSectionGroup | Update the navigation property parentSectionGroup in groups
[**GroupsOnenoteNotebooksSectionGroupsUpdateParentNotebook**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsUpdateParentNotebook) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentNotebook | Update the navigation property parentNotebook in groups
[**GroupsOnenoteNotebooksSectionGroupsUpdateParentSectionGroup**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsUpdateParentSectionGroup) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentSectionGroup | Update the navigation property parentSectionGroup in groups
[**GroupsOnenoteNotebooksSectionGroupsUpdateSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsUpdateSectionGroups) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Update the navigation property sectionGroups in groups
[**GroupsOnenoteNotebooksSectionGroupsUpdateSections**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksSectionGroupsUpdateSections) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id}) | Update the navigation property sections in groups
[**GroupsOnenoteNotebooksUpdateSectionGroups**](GroupsOnenoteNotebooksSectionGroupApi.md#GroupsOnenoteNotebooksUpdateSectionGroups) | **Patch** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id}) | Update the navigation property sectionGroups in groups



## GroupsOnenoteNotebooksCreateSectionGroups

> MicrosoftGraphSectionGroup GroupsOnenoteNotebooksCreateSectionGroups(ctx, groupId, notebookId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
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


## GroupsOnenoteNotebooksGetSectionGroups

> MicrosoftGraphSectionGroup GroupsOnenoteNotebooksGetSectionGroups(ctx, groupId, notebookId, sectionGroupId, optional)
Get sectionGroups from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteNotebooksGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksGetSectionGroupsOpts struct


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


## GroupsOnenoteNotebooksListSectionGroups

> CollectionOfSectionGroup GroupsOnenoteNotebooksListSectionGroups(ctx, groupId, notebookId, optional)
Get sectionGroups from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
 **optional** | ***GroupsOnenoteNotebooksListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksListSectionGroupsOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsCreateSectionGroups

> MicrosoftGraphSectionGroup GroupsOnenoteNotebooksSectionGroupsCreateSectionGroups(ctx, groupId, notebookId, sectionGroupId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
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


## GroupsOnenoteNotebooksSectionGroupsCreateSections

> MicrosoftGraphOnenoteSection GroupsOnenoteNotebooksSectionGroupsCreateSections(ctx, groupId, notebookId, sectionGroupId, microsoftGraphOnenoteSection)
Create new navigation property to sections for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
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


## GroupsOnenoteNotebooksSectionGroupsGetParentNotebook

> MicrosoftGraphNotebook GroupsOnenoteNotebooksSectionGroupsGetParentNotebook(ctx, groupId, notebookId, sectionGroupId, optional)
Get parentNotebook from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsGetParentNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsGetParentNotebookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteNotebooksSectionGroupsGetParentSectionGroup

> MicrosoftGraphSectionGroup GroupsOnenoteNotebooksSectionGroupsGetParentSectionGroup(ctx, groupId, notebookId, sectionGroupId, optional)
Get parentSectionGroup from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsGetParentSectionGroupOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsGetSectionGroups

> MicrosoftGraphSectionGroup GroupsOnenoteNotebooksSectionGroupsGetSectionGroups(ctx, groupId, notebookId, sectionGroupId, sectionGroupId1, optional)
Get sectionGroups from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsGetSectionGroupsOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsGetSections

> MicrosoftGraphOnenoteSection GroupsOnenoteNotebooksSectionGroupsGetSections(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get sections from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsGetSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsGetSectionsOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsListSectionGroups

> CollectionOfSectionGroup GroupsOnenoteNotebooksSectionGroupsListSectionGroups(ctx, groupId, notebookId, sectionGroupId, optional)
Get sectionGroups from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsListSectionGroupsOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsListSections

> CollectionOfOnenoteSection GroupsOnenoteNotebooksSectionGroupsListSections(ctx, groupId, notebookId, sectionGroupId, optional)
Get sections from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsListSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsListSectionsOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsSectionsCreatePages

> MicrosoftGraphOnenotePage GroupsOnenoteNotebooksSectionGroupsSectionsCreatePages(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphOnenotePage)
Create new navigation property to pages for groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
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


## GroupsOnenoteNotebooksSectionGroupsSectionsGetPages

> MicrosoftGraphOnenotePage GroupsOnenoteNotebooksSectionGroupsSectionsGetPages(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, optional)
Get pages from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsSectionsGetPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsSectionsGetPagesOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsSectionsGetParentNotebook

> MicrosoftGraphNotebook GroupsOnenoteNotebooksSectionGroupsSectionsGetParentNotebook(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get parentNotebook from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsSectionsGetParentNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsSectionsGetParentNotebookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup

> MicrosoftGraphSectionGroup GroupsOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get parentSectionGroup from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroupOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsSectionsListPages

> CollectionOfOnenotePage GroupsOnenoteNotebooksSectionGroupsSectionsListPages(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get pages from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsSectionsListPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsSectionsListPagesOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook

> MicrosoftGraphNotebook GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, optional)
Get parentNotebook from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphNotebook**](microsoft.graph.notebook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection

> MicrosoftGraphOnenoteSection GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, optional)
Get parentSection from groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsOnenoteNotebooksSectionGroupsSectionsPagesGetParentSectionOpts struct


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


## GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook

> GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, microsoftGraphNotebook)
Update the navigation property parentNotebook in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property values | 

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


## GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection

> GroupsOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, microsoftGraphOnenoteSection)
Update the navigation property parentSection in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
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


## GroupsOnenoteNotebooksSectionGroupsSectionsUpdatePages

> GroupsOnenoteNotebooksSectionGroupsSectionsUpdatePages(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, microsoftGraphOnenotePage)
Update the navigation property pages in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
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


## GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook

> GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphNotebook)
Update the navigation property parentNotebook in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property values | 

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


## GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup

> GroupsOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
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


## GroupsOnenoteNotebooksSectionGroupsUpdateParentNotebook

> GroupsOnenoteNotebooksSectionGroupsUpdateParentNotebook(ctx, groupId, notebookId, sectionGroupId, microsoftGraphNotebook)
Update the navigation property parentNotebook in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**microsoftGraphNotebook** | [**MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md)| New navigation property values | 

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


## GroupsOnenoteNotebooksSectionGroupsUpdateParentSectionGroup

> GroupsOnenoteNotebooksSectionGroupsUpdateParentSectionGroup(ctx, groupId, notebookId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
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


## GroupsOnenoteNotebooksSectionGroupsUpdateSectionGroups

> GroupsOnenoteNotebooksSectionGroupsUpdateSectionGroups(ctx, groupId, notebookId, sectionGroupId, sectionGroupId1, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
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


## GroupsOnenoteNotebooksSectionGroupsUpdateSections

> GroupsOnenoteNotebooksSectionGroupsUpdateSections(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphOnenoteSection)
Update the navigation property sections in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
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


## GroupsOnenoteNotebooksUpdateSectionGroups

> GroupsOnenoteNotebooksUpdateSectionGroups(ctx, groupId, notebookId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**notebookId** | **string**| key: notebook-id of notebook | 
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

