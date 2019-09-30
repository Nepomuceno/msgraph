# \SitesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesGetActivitiesByInterval4c35**](SitesFunctionsApi.md#SitesGetActivitiesByInterval4c35) | **Get** /sites({site-id})/getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SitesGetByPath1173**](SitesFunctionsApi.md#SitesGetByPath1173) | **Get** /sites({site-id})/getByPath(path&#x3D;{path}) | Invoke function getByPath
[**SitesListsItemsGetActivitiesByInterval4c35**](SitesFunctionsApi.md#SitesListsItemsGetActivitiesByInterval4c35) | **Get** /sites({site-id})/lists({list-id})/items({listItem-id})/getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SitesOnenoteNotebooksGetRecentNotebooks1d97**](SitesFunctionsApi.md#SitesOnenoteNotebooksGetRecentNotebooks1d97) | **Get** /sites({site-id})/onenote/notebooks/getRecentNotebooks(includePersonalNotebooks&#x3D;{includePersonalNotebooks}) | Invoke function getRecentNotebooks
[**SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**SitesOnenoteNotebooksSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenoteNotebooksSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/notebooks({notebook-id})/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/pages({onenotePage-id})/parentNotebook/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id1})/preview() | Invoke function preview
[**SitesOnenotePagesParentNotebookSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenotePagesParentNotebookSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/pages({onenotePage-id})/parentNotebook/sections({onenoteSection-id})/pages({onenotePage-id1})/preview() | Invoke function preview
[**SitesOnenotePagesParentSectionPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenotePagesParentSectionPagesPreview12f3) | **Get** /sites({site-id})/onenote/pages({onenotePage-id})/parentSection/pages({onenotePage-id1})/preview() | Invoke function preview
[**SitesOnenotePagesPreview12f3**](SitesFunctionsApi.md#SitesOnenotePagesPreview12f3) | **Get** /sites({site-id})/onenote/pages({onenotePage-id})/preview() | Invoke function preview
[**SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/parentNotebook/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**SitesOnenoteSectionGroupsSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenoteSectionGroupsSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**SitesOnenoteSectionsPagesPreview12f3**](SitesFunctionsApi.md#SitesOnenoteSectionsPagesPreview12f3) | **Get** /sites({site-id})/onenote/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview



## SitesGetActivitiesByInterval4c35

> []AnyOfmicrosoftGraphItemActivityStat SitesGetActivitiesByInterval4c35(ctx, siteId)
Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 

### Return type

[**[]AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesGetByPath1173

> AnyOfmicrosoftGraphSite SitesGetByPath1173(ctx, siteId, path)
Invoke function getByPath

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**path** | **string**|  | 

### Return type

[**AnyOfmicrosoftGraphSite**](anyOf&lt;microsoft.graph.site&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetActivitiesByInterval4c35

> []AnyOfmicrosoftGraphItemActivityStat SitesListsItemsGetActivitiesByInterval4c35(ctx, siteId, listId, listItemId)
Invoke function getActivitiesByInterval

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**listId** | **string**| key: list-id of list | 
**listItemId** | **string**| key: listItem-id of listItem | 

### Return type

[**[]AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteNotebooksGetRecentNotebooks1d97

> []AnyOfmicrosoftGraphRecentNotebook SitesOnenoteNotebooksGetRecentNotebooks1d97(ctx, siteId, includePersonalNotebooks)
Invoke function getRecentNotebooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**includePersonalNotebooks** | **bool**|  | [default to false]

### Return type

[**[]AnyOfmicrosoftGraphRecentNotebook**](anyOf&lt;microsoft.graph.recentNotebook&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3(ctx, siteId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteNotebooksSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenoteNotebooksSectionsPagesPreview12f3(ctx, siteId, notebookId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**notebookId** | **string**| key: notebook-id of notebook | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3(ctx, siteId, onenotePageId, sectionGroupId, onenoteSectionId, onenotePageId1)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId1** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesParentNotebookSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenotePagesParentNotebookSectionsPagesPreview12f3(ctx, siteId, onenotePageId, onenoteSectionId, onenotePageId1)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId1** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesParentSectionPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenotePagesParentSectionPagesPreview12f3(ctx, siteId, onenotePageId, onenotePageId1)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
**onenotePageId1** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenotePagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenotePagesPreview12f3(ctx, siteId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3(ctx, siteId, sectionGroupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteSectionGroupsSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenoteSectionGroupsSectionsPagesPreview12f3(ctx, siteId, sectionGroupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesOnenoteSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview SitesOnenoteSectionsPagesPreview12f3(ctx, siteId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string**| key: site-id of site | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 

### Return type

[**AnyOfmicrosoftGraphOnenotePagePreview**](anyOf&lt;microsoft.graph.onenotePagePreview&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

