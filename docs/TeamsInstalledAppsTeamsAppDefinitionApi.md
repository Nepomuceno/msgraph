# \TeamsInstalledAppsTeamsAppDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsInstalledAppsGetTeamsAppDefinition**](TeamsInstalledAppsTeamsAppDefinitionApi.md#TeamsInstalledAppsGetTeamsAppDefinition) | **Get** /teams({team-id})/installedApps({teamsAppInstallation-id})/teamsAppDefinition | Get teamsAppDefinition from teams



## TeamsInstalledAppsGetTeamsAppDefinition

> MicrosoftGraphTeamsAppDefinition TeamsInstalledAppsGetTeamsAppDefinition(ctx, teamId, teamsAppInstallationId, optional)
Get teamsAppDefinition from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 
 **optional** | ***TeamsInstalledAppsGetTeamsAppDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsInstalledAppsGetTeamsAppDefinitionOpts struct


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

