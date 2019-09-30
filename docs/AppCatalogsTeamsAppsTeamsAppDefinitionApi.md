# \AppCatalogsTeamsAppsTeamsAppDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCatalogsTeamsAppsCreateAppDefinitions**](AppCatalogsTeamsAppsTeamsAppDefinitionApi.md#AppCatalogsTeamsAppsCreateAppDefinitions) | **Post** /appCatalogs/teamsApps({teamsApp-id})/appDefinitions | Create new navigation property to appDefinitions for appCatalogs
[**AppCatalogsTeamsAppsGetAppDefinitions**](AppCatalogsTeamsAppsTeamsAppDefinitionApi.md#AppCatalogsTeamsAppsGetAppDefinitions) | **Get** /appCatalogs/teamsApps({teamsApp-id})/appDefinitions({teamsAppDefinition-id}) | Get appDefinitions from appCatalogs
[**AppCatalogsTeamsAppsListAppDefinitions**](AppCatalogsTeamsAppsTeamsAppDefinitionApi.md#AppCatalogsTeamsAppsListAppDefinitions) | **Get** /appCatalogs/teamsApps({teamsApp-id})/appDefinitions | Get appDefinitions from appCatalogs
[**AppCatalogsTeamsAppsUpdateAppDefinitions**](AppCatalogsTeamsAppsTeamsAppDefinitionApi.md#AppCatalogsTeamsAppsUpdateAppDefinitions) | **Patch** /appCatalogs/teamsApps({teamsApp-id})/appDefinitions({teamsAppDefinition-id}) | Update the navigation property appDefinitions in appCatalogs



## AppCatalogsTeamsAppsCreateAppDefinitions

> MicrosoftGraphTeamsAppDefinition AppCatalogsTeamsAppsCreateAppDefinitions(ctx, teamsAppId, microsoftGraphTeamsAppDefinition)
Create new navigation property to appDefinitions for appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string**| key: teamsApp-id of teamsApp | 
**microsoftGraphTeamsAppDefinition** | [**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md)| New navigation property | 

### Return type

[**MicrosoftGraphTeamsAppDefinition**](microsoft.graph.teamsAppDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsGetAppDefinitions

> MicrosoftGraphTeamsAppDefinition AppCatalogsTeamsAppsGetAppDefinitions(ctx, teamsAppId, teamsAppDefinitionId, optional)
Get appDefinitions from appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string**| key: teamsApp-id of teamsApp | 
**teamsAppDefinitionId** | **string**| key: teamsAppDefinition-id of teamsAppDefinition | 
 **optional** | ***AppCatalogsTeamsAppsGetAppDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCatalogsTeamsAppsGetAppDefinitionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppDefinition**](microsoft.graph.teamsAppDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsListAppDefinitions

> CollectionOfTeamsAppDefinition AppCatalogsTeamsAppsListAppDefinitions(ctx, teamsAppId, optional)
Get appDefinitions from appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string**| key: teamsApp-id of teamsApp | 
 **optional** | ***AppCatalogsTeamsAppsListAppDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCatalogsTeamsAppsListAppDefinitionsOpts struct


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

[**CollectionOfTeamsAppDefinition**](Collection of teamsAppDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsUpdateAppDefinitions

> AppCatalogsTeamsAppsUpdateAppDefinitions(ctx, teamsAppId, teamsAppDefinitionId, microsoftGraphTeamsAppDefinition)
Update the navigation property appDefinitions in appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string**| key: teamsApp-id of teamsApp | 
**teamsAppDefinitionId** | **string**| key: teamsAppDefinition-id of teamsAppDefinition | 
**microsoftGraphTeamsAppDefinition** | [**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md)| New navigation property values | 

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

