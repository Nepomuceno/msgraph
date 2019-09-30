# \UsersOnenoteSectionGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOnenoteCreateSectionGroups**](UsersOnenoteSectionGroupApi.md#UsersOnenoteCreateSectionGroups) | **Post** /users({user-id})/onenote/sectionGroups | Create new navigation property to sectionGroups for users
[**UsersOnenoteGetSectionGroups**](UsersOnenoteSectionGroupApi.md#UsersOnenoteGetSectionGroups) | **Get** /users({user-id})/onenote/sectionGroups({sectionGroup-id}) | Get sectionGroups from users
[**UsersOnenoteListSectionGroups**](UsersOnenoteSectionGroupApi.md#UsersOnenoteListSectionGroups) | **Get** /users({user-id})/onenote/sectionGroups | Get sectionGroups from users
[**UsersOnenoteUpdateSectionGroups**](UsersOnenoteSectionGroupApi.md#UsersOnenoteUpdateSectionGroups) | **Patch** /users({user-id})/onenote/sectionGroups({sectionGroup-id}) | Update the navigation property sectionGroups in users



## UsersOnenoteCreateSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteCreateSectionGroups(ctx, userId, microsoftGraphSectionGroup)
Create new navigation property to sectionGroups for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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


## UsersOnenoteGetSectionGroups

> MicrosoftGraphSectionGroup UsersOnenoteGetSectionGroups(ctx, userId, sectionGroupId, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**sectionGroupId** | **string**| key: sectionGroup-id of sectionGroup | 
 **optional** | ***UsersOnenoteGetSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteGetSectionGroupsOpts struct


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


## UsersOnenoteListSectionGroups

> CollectionOfSectionGroup UsersOnenoteListSectionGroups(ctx, userId, optional)
Get sectionGroups from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersOnenoteListSectionGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersOnenoteListSectionGroupsOpts struct


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


## UsersOnenoteUpdateSectionGroups

> UsersOnenoteUpdateSectionGroups(ctx, userId, sectionGroupId, microsoftGraphSectionGroup)
Update the navigation property sectionGroups in users

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

