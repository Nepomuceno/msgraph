# \UsersOnenoteNotebooksSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteNotebooksCreateSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksCreateSectionGroups) | **Post** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups | Create new navigation property to sectionGroups for users
[**UsersOnenoteNotebooksGetSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksGetSectionGroups) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id}) | Get sectionGroups from users
[**UsersOnenoteNotebooksListSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksListSectionGroups) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups | Get sectionGroups from users
[**UsersOnenoteNotebooksSectionGroupsCreateSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsCreateSectionGroups) | **Post** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups | Create new navigation property to sectionGroups for users
[**UsersOnenoteNotebooksSectionGroupsCreateSections**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsCreateSections) | **Post** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections | Create new navigation property to sections for users
[**UsersOnenoteNotebooksSectionGroupsGetParentNotebook**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsGetParentNotebook) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentNotebook | Get parentNotebook from users
[**UsersOnenoteNotebooksSectionGroupsGetParentSectionGroup**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsGetParentSectionGroup) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentSectionGroup | Get parentSectionGroup from users
[**UsersOnenoteNotebooksSectionGroupsGetSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsGetSectionGroups) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Get sectionGroups from users
[**UsersOnenoteNotebooksSectionGroupsGetSections**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsGetSections) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id}) | Get sections from users
[**UsersOnenoteNotebooksSectionGroupsListSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsListSectionGroups) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups | Get sectionGroups from users
[**UsersOnenoteNotebooksSectionGroupsListSections**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsListSections) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections | Get sections from users
[**UsersOnenoteNotebooksSectionGroupsSectionsCreatePages**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsCreatePages) | **Post** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages | Create new navigation property to pages for users
[**UsersOnenoteNotebooksSectionGroupsSectionsGetPages**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsGetPages) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id}) | Get pages from users
[**UsersOnenoteNotebooksSectionGroupsSectionsGetParentNotebook**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsGetParentNotebook) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentNotebook | Get parentNotebook from users
[**UsersOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentSectionGroup | Get parentSectionGroup from users
[**UsersOnenoteNotebooksSectionGroupsSectionsListPages**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsListPages) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages | Get pages from users
[**UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentNotebook | Get parentNotebook from users
[**UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection) | **Get** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentSection | Get parentSection from users
[**UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentNotebook | Update the navigation property parentNotebook in users
[**UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id})/parentSection | Update the navigation property parentSection in users
[**UsersOnenoteNotebooksSectionGroupsSectionsUpdatePages**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsUpdatePages) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/pages({onenotePage-id}) | Update the navigation property pages in users
[**UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentNotebook | Update the navigation property parentNotebook in users
[**UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id})/parentSectionGroup | Update the navigation property parentSectionGroup in users
[**UsersOnenoteNotebooksSectionGroupsUpdateParentNotebook**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsUpdateParentNotebook) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentNotebook | Update the navigation property parentNotebook in users
[**UsersOnenoteNotebooksSectionGroupsUpdateParentSectionGroup**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsUpdateParentSectionGroup) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/parentSectionGroup | Update the navigation property parentSectionGroup in users
[**UsersOnenoteNotebooksSectionGroupsUpdateSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsUpdateSectionGroups) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Update the navigation property sectionGroups in users
[**UsersOnenoteNotebooksSectionGroupsUpdateSections**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksSectionGroupsUpdateSections) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id})/sections({onenoteSection-id}) | Update the navigation property sections in users
[**UsersOnenoteNotebooksUpdateSectionGroups**](UsersOnenoteNotebooksSectionGroupApi.md#UsersOnenoteNotebooksUpdateSectionGroups) | **Patch** /users({user-id})/onenote/notebooks({notebook-id})/sectionGroups({sectionGroup-id}) | Update the navigation property sectionGroups in users



## UsersOnenoteNotebooksCreateSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteNotebooksCreateSectionGroups(ctx, userId, notebookId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksGetSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteNotebooksGetSectionGroups(ctx, userId, notebookId, sectionGroupId, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteNotebooksGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksGetSectionGroupsOpts struct


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


## UsersOnenoteNotebooksListSectionGroups

> CollectionOfSectionGroup UsersOnenoteNotebooksListSectionGroups(ctx, userId, notebookId, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
 **optional** | ***UsersOnenoteNotebooksListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksListSectionGroupsOpts struct


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


## UsersOnenoteNotebooksSectionGroupsCreateSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteNotebooksSectionGroupsCreateSectionGroups(ctx, userId, notebookId, sectionGroupId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsCreateSections

> MicrosoftGraphOnenoteSection UsersOnenoteNotebooksSectionGroupsCreateSections(ctx, userId, notebookId, sectionGroupId, microsoftGraphOnenoteSection)
Create new navigation property to sections for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsGetParentNotebook

> MicrosoftGraphNotebook UsersOnenoteNotebooksSectionGroupsGetParentNotebook(ctx, userId, notebookId, sectionGroupId, optional)
Get parentNotebook from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsGetParentNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsGetParentNotebookOpts struct


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


## UsersOnenoteNotebooksSectionGroupsGetParentSectionGroup

> MicrosoftGraphSectionGroup UsersOnenoteNotebooksSectionGroupsGetParentSectionGroup(ctx, userId, notebookId, sectionGroupId, optional)
Get parentSectionGroup from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsGetParentSectionGroupOpts struct


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


## UsersOnenoteNotebooksSectionGroupsGetSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteNotebooksSectionGroupsGetSectionGroups(ctx, userId, notebookId, sectionGroupId, sectionGroupId1, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsGetSectionGroupsOpts struct


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


## UsersOnenoteNotebooksSectionGroupsGetSections

> MicrosoftGraphOnenoteSection UsersOnenoteNotebooksSectionGroupsGetSections(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get sections from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsGetSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsGetSectionsOpts struct


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


## UsersOnenoteNotebooksSectionGroupsListSectionGroups

> CollectionOfSectionGroup UsersOnenoteNotebooksSectionGroupsListSectionGroups(ctx, userId, notebookId, sectionGroupId, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsListSectionGroupsOpts struct


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


## UsersOnenoteNotebooksSectionGroupsListSections

> CollectionOfOnenoteSection UsersOnenoteNotebooksSectionGroupsListSections(ctx, userId, notebookId, sectionGroupId, optional)
Get sections from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsListSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsListSectionsOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsCreatePages

> MicrosoftGraphOnenotePage UsersOnenoteNotebooksSectionGroupsSectionsCreatePages(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphOnenotePage)
Create new navigation property to pages for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsSectionsGetPages

> MicrosoftGraphOnenotePage UsersOnenoteNotebooksSectionGroupsSectionsGetPages(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, optional)
Get pages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsSectionsGetPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsSectionsGetPagesOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsGetParentNotebook

> MicrosoftGraphNotebook UsersOnenoteNotebooksSectionGroupsSectionsGetParentNotebook(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get parentNotebook from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsSectionsGetParentNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsSectionsGetParentNotebookOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup

> MicrosoftGraphSectionGroup UsersOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroup(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get parentSectionGroup from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsSectionsGetParentSectionGroupOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsListPages

> CollectionOfOnenotePage UsersOnenoteNotebooksSectionGroupsSectionsListPages(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, optional)
Get pages from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsSectionsListPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsSectionsListPagesOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook

> MicrosoftGraphNotebook UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebook(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, optional)
Get parentNotebook from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentNotebookOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection

> MicrosoftGraphOnenoteSection UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentSection(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, optional)
Get parentSection from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**notebookId** | **string**| key: notebook-id of notebook | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**onenoteSectionId** | **string**| key: onenoteSection-id of onenoteSection | 
**onenotePageId** | **string**| key: onenotePage-id of onenotePage | 
 **optional** | ***UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteNotebooksSectionGroupsSectionsPagesGetParentSectionOpts struct


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


## UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook

> UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentNotebook(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, microsoftGraphNotebook)
Update the navigation property parentNotebook in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection

> UsersOnenoteNotebooksSectionGroupsSectionsPagesUpdateParentSection(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, microsoftGraphOnenoteSection)
Update the navigation property parentSection in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsSectionsUpdatePages

> UsersOnenoteNotebooksSectionGroupsSectionsUpdatePages(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, onenotePageId, microsoftGraphOnenotePage)
Update the navigation property pages in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook

> UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentNotebook(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphNotebook)
Update the navigation property parentNotebook in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup

> UsersOnenoteNotebooksSectionGroupsSectionsUpdateParentSectionGroup(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsUpdateParentNotebook

> UsersOnenoteNotebooksSectionGroupsUpdateParentNotebook(ctx, userId, notebookId, sectionGroupId, microsoftGraphNotebook)
Update the navigation property parentNotebook in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsUpdateParentSectionGroup

> UsersOnenoteNotebooksSectionGroupsUpdateParentSectionGroup(ctx, userId, notebookId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsUpdateSectionGroups

> UsersOnenoteNotebooksSectionGroupsUpdateSectionGroups(ctx, userId, notebookId, sectionGroupId, sectionGroupId1, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksSectionGroupsUpdateSections

> UsersOnenoteNotebooksSectionGroupsUpdateSections(ctx, userId, notebookId, sectionGroupId, onenoteSectionId, microsoftGraphOnenoteSection)
Update the navigation property sections in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteNotebooksUpdateSectionGroups

> UsersOnenoteNotebooksUpdateSectionGroups(ctx, userId, notebookId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

