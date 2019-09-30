# \GroupsFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCalendarCalendarViewDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarCalendarViewDeltaFa14) | **Get** /groups({group-id})/calendar/calendarView/delta() | Invoke function delta
[**GroupsCalendarCalendarViewInstancesDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarCalendarViewInstancesDeltaFa14) | **Get** /groups({group-id})/calendar/calendarView({event-id})/instances/delta() | Invoke function delta
[**GroupsCalendarEventsDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarEventsDeltaFa14) | **Get** /groups({group-id})/calendar/events/delta() | Invoke function delta
[**GroupsCalendarEventsInstancesDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarEventsInstancesDeltaFa14) | **Get** /groups({group-id})/calendar/events({event-id})/instances/delta() | Invoke function delta
[**GroupsCalendarViewCalendarCalendarViewDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarViewCalendarCalendarViewDeltaFa14) | **Get** /groups({group-id})/calendarView({event-id})/calendar/calendarView/delta() | Invoke function delta
[**GroupsCalendarViewCalendarEventsDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarViewCalendarEventsDeltaFa14) | **Get** /groups({group-id})/calendarView({event-id})/calendar/events/delta() | Invoke function delta
[**GroupsCalendarViewDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarViewDeltaFa14) | **Get** /groups({group-id})/calendarView/delta() | Invoke function delta
[**GroupsCalendarViewInstancesDeltaFa14**](GroupsFunctionsApi.md#GroupsCalendarViewInstancesDeltaFa14) | **Get** /groups({group-id})/calendarView({event-id})/instances/delta() | Invoke function delta
[**GroupsDeltaFa14**](GroupsFunctionsApi.md#GroupsDeltaFa14) | **Get** /groups/delta() | Invoke function delta
[**GroupsEventsCalendarCalendarViewDeltaFa14**](GroupsFunctionsApi.md#GroupsEventsCalendarCalendarViewDeltaFa14) | **Get** /groups({group-id})/events({event-id})/calendar/calendarView/delta() | Invoke function delta
[**GroupsEventsCalendarEventsDeltaFa14**](GroupsFunctionsApi.md#GroupsEventsCalendarEventsDeltaFa14) | **Get** /groups({group-id})/events({event-id})/calendar/events/delta() | Invoke function delta
[**GroupsEventsDeltaFa14**](GroupsFunctionsApi.md#GroupsEventsDeltaFa14) | **Get** /groups({group-id})/events/delta() | Invoke function delta
[**GroupsEventsInstancesDeltaFa14**](GroupsFunctionsApi.md#GroupsEventsInstancesDeltaFa14) | **Get** /groups({group-id})/events({event-id})/instances/delta() | Invoke function delta
[**GroupsOnenoteNotebooksGetRecentNotebooks1d97**](GroupsFunctionsApi.md#GroupsOnenoteNotebooksGetRecentNotebooks1d97) | **Get** /groups({group-id})/onenote/notebooks/getRecentNotebooks(includePersonalNotebooks&#x3D;{includePersonalNotebooks}) | Invoke function getRecentNotebooks
[**GroupsOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**GroupsOnenoteNotebooksSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenoteNotebooksSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/notebooks({notebook-id})/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**GroupsOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/pages({onenotePage-id})/parentNotebook/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id1})/preview() | Invoke function preview
[**GroupsOnenotePagesParentNotebookSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenotePagesParentNotebookSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/pages({onenotePage-id})/parentNotebook/sections({onenoteSection-id})/pages({onenotePage-id1})/preview() | Invoke function preview
[**GroupsOnenotePagesParentSectionPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenotePagesParentSectionPagesPreview12f3) | **Get** /groups({group-id})/onenote/pages({onenotePage-id})/parentSection/pages({onenotePage-id1})/preview() | Invoke function preview
[**GroupsOnenotePagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenotePagesPreview12f3) | **Get** /groups({group-id})/onenote/pages({onenotePage-id})/preview() | Invoke function preview
[**GroupsOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/parentNotebook/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**GroupsOnenoteSectionGroupsSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenoteSectionGroupsSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview
[**GroupsOnenoteSectionsPagesPreview12f3**](GroupsFunctionsApi.md#GroupsOnenoteSectionsPagesPreview12f3) | **Get** /groups({group-id})/onenote/sections({onenoteSection-id})/pages({onenotePage-id})/preview() | Invoke function preview



## GroupsCalendarCalendarViewDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarCalendarViewDeltaFa14(ctx, groupId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarCalendarViewInstancesDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarCalendarViewInstancesDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarEventsDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarEventsDeltaFa14(ctx, groupId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarEventsInstancesDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarEventsInstancesDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarViewCalendarCalendarViewDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarViewCalendarCalendarViewDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarViewCalendarEventsDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarViewCalendarEventsDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarViewDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarViewDeltaFa14(ctx, groupId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCalendarViewInstancesDeltaFa14

> []MicrosoftGraphEvent GroupsCalendarViewInstancesDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDeltaFa14

> []AnyOfmicrosoftGraphGroup GroupsDeltaFa14(ctx, )
Invoke function delta

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AnyOfmicrosoftGraphGroup**](anyOf&lt;microsoft.graph.group&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsEventsCalendarCalendarViewDeltaFa14

> []MicrosoftGraphEvent GroupsEventsCalendarCalendarViewDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsEventsCalendarEventsDeltaFa14

> []MicrosoftGraphEvent GroupsEventsCalendarEventsDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsEventsDeltaFa14

> []MicrosoftGraphEvent GroupsEventsDeltaFa14(ctx, groupId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsEventsInstancesDeltaFa14

> []MicrosoftGraphEvent GroupsEventsInstancesDeltaFa14(ctx, groupId, eventId)
Invoke function delta

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
**eventId** | **string**| key: event-id of event | 

### Return type

[**[]MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsOnenoteNotebooksGetRecentNotebooks1d97

> []AnyOfmicrosoftGraphRecentNotebook GroupsOnenoteNotebooksGetRecentNotebooks1d97(ctx, groupId, includePersonalNotebooks)
Invoke function getRecentNotebooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenoteNotebooksSectionGroupsSectionsPagesPreview12f3(ctx, groupId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteNotebooksSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenoteNotebooksSectionsPagesPreview12f3(ctx, groupId, notebookId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenotePagesParentNotebookSectionGroupsSectionsPagesPreview12f3(ctx, groupId, onenotePageId, sectionGroupId, onenoteSectionId, onenotePageId1)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenotePagesParentNotebookSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenotePagesParentNotebookSectionsPagesPreview12f3(ctx, groupId, onenotePageId, onenoteSectionId, onenotePageId1)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenotePagesParentSectionPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenotePagesParentSectionPagesPreview12f3(ctx, groupId, onenotePageId, onenotePageId1)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenotePagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenotePagesPreview12f3(ctx, groupId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenoteSectionGroupsParentNotebookSectionsPagesPreview12f3(ctx, groupId, sectionGroupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteSectionGroupsSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenoteSectionGroupsSectionsPagesPreview12f3(ctx, groupId, sectionGroupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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


## GroupsOnenoteSectionsPagesPreview12f3

> AnyOfmicrosoftGraphOnenotePagePreview GroupsOnenoteSectionsPagesPreview12f3(ctx, groupId, onenoteSectionId, onenotePageId)
Invoke function preview

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| key: group-id of group | 
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

