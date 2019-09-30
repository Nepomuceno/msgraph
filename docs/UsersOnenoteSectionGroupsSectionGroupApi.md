# \UsersOnenoteSectionGroupsSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteSectionGroupsCreateSectionGroups**](UsersOnenoteSectionGroupsSectionGroupApi.md#UsersOnenoteSectionGroupsCreateSectionGroups) | **Post** /users({user-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Create new navigation property to sectionGroups for users
[**UsersOnenoteSectionGroupsGetParentSectionGroup**](UsersOnenoteSectionGroupsSectionGroupApi.md#UsersOnenoteSectionGroupsGetParentSectionGroup) | **Get** /users({user-id})/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Get parentSectionGroup from users
[**UsersOnenoteSectionGroupsGetSectionGroups**](UsersOnenoteSectionGroupsSectionGroupApi.md#UsersOnenoteSectionGroupsGetSectionGroups) | **Get** /users({user-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Get sectionGroups from users
[**UsersOnenoteSectionGroupsListSectionGroups**](UsersOnenoteSectionGroupsSectionGroupApi.md#UsersOnenoteSectionGroupsListSectionGroups) | **Get** /users({user-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups | Get sectionGroups from users
[**UsersOnenoteSectionGroupsUpdateParentSectionGroup**](UsersOnenoteSectionGroupsSectionGroupApi.md#UsersOnenoteSectionGroupsUpdateParentSectionGroup) | **Patch** /users({user-id})/onenote/sectionGroups({sectionGroup-id})/parentSectionGroup | Update the navigation property parentSectionGroup in users
[**UsersOnenoteSectionGroupsUpdateSectionGroups**](UsersOnenoteSectionGroupsSectionGroupApi.md#UsersOnenoteSectionGroupsUpdateSectionGroups) | **Patch** /users({user-id})/onenote/sectionGroups({sectionGroup-id})/sectionGroups({sectionGroup-id1}) | Update the navigation property sectionGroups in users



## UsersOnenoteSectionGroupsCreateSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteSectionGroupsCreateSectionGroups(ctx, userId, sectionGroupId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteSectionGroupsGetParentSectionGroup

> MicrosoftGraphSectionGroup UsersOnenoteSectionGroupsGetParentSectionGroup(ctx, userId, sectionGroupId, optional)
Get parentSectionGroup from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteSectionGroupsGetParentSectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteSectionGroupsGetParentSectionGroupOpts struct


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


## UsersOnenoteSectionGroupsGetSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteSectionGroupsGetSectionGroups(ctx, userId, sectionGroupId, sectionGroupId1, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
**sectionGroupId1** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteSectionGroupsGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteSectionGroupsGetSectionGroupsOpts struct


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


## UsersOnenoteSectionGroupsListSectionGroups

> CollectionOfSectionGroup UsersOnenoteSectionGroupsListSectionGroups(ctx, userId, sectionGroupId, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteSectionGroupsListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteSectionGroupsListSectionGroupsOpts struct


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


## UsersOnenoteSectionGroupsUpdateParentSectionGroup

> UsersOnenoteSectionGroupsUpdateParentSectionGroup(ctx, userId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property parentSectionGroup in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteSectionGroupsUpdateSectionGroups

> UsersOnenoteSectionGroupsUpdateSectionGroups(ctx, userId, sectionGroupId, sectionGroupId1, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

