# \AppCatalogsTeamsAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCatalogsCreateTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsCreateTeamsApps) | **Post** /appCatalogs/teamsApps | Create new navigation property to teamsApps for appCatalogs
[**AppCatalogsGetTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsGetTeamsApps) | **Get** /appCatalogs/teamsApps({teamsApp-id}) | Get teamsApps from appCatalogs
[**AppCatalogsListTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsListTeamsApps) | **Get** /appCatalogs/teamsApps | Get teamsApps from appCatalogs
[**AppCatalogsUpdateTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsUpdateTeamsApps) | **Patch** /appCatalogs/teamsApps({teamsApp-id}) | Update the navigation property teamsApps in appCatalogs



## AppCatalogsCreateTeamsApps

> MicrosoftGraphTeamsApp AppCatalogsCreateTeamsApps(ctx, microsoftGraphTeamsApp)
Create new navigation property to teamsApps for appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTeamsApp** | [**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md)| New navigation property | 

### Return type

[**MicrosoftGraphTeamsApp**](microsoft.graph.teamsApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsGetTeamsApps

> MicrosoftGraphTeamsApp AppCatalogsGetTeamsApps(ctx, teamsAppId, optional)
Get teamsApps from appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string**| key: teamsApp-id of teamsApp | 
 **optional** | ***AppCatalogsGetTeamsAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCatalogsGetTeamsAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsApp**](microsoft.graph.teamsApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsListTeamsApps

> CollectionOfTeamsApp AppCatalogsListTeamsApps(ctx, optional)
Get teamsApps from appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppCatalogsListTeamsAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCatalogsListTeamsAppsOpts struct


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

[**CollectionOfTeamsApp**](Collection of teamsApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsUpdateTeamsApps

> AppCatalogsUpdateTeamsApps(ctx, teamsAppId, microsoftGraphTeamsApp)
Update the navigation property teamsApps in appCatalogs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string**| key: teamsApp-id of teamsApp | 
**microsoftGraphTeamsApp** | [**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md)| New navigation property values | 

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

