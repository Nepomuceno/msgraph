# \TeamsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsArchive**](TeamsActionsApi.md#TeamsArchive) | **Post** /teams({team-id})/archive | Invoke action archive
[**TeamsClone**](TeamsActionsApi.md#TeamsClone) | **Post** /teams({team-id})/clone | Invoke action clone
[**TeamsInstalledAppsUpgrade**](TeamsActionsApi.md#TeamsInstalledAppsUpgrade) | **Post** /teams({team-id})/installedApps({teamsAppInstallation-id})/upgrade | Invoke action upgrade
[**TeamsUnarchive**](TeamsActionsApi.md#TeamsUnarchive) | **Post** /teams({team-id})/unarchive | Invoke action unarchive



## TeamsArchive

> TeamsArchive(ctx, teamId, inlineObject514)
Invoke action archive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**inlineObject514** | [**InlineObject514**](InlineObject514.md)|  | 

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


## TeamsClone

> TeamsClone(ctx, teamId, inlineObject515)
Invoke action clone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**inlineObject515** | [**InlineObject515**](InlineObject515.md)|  | 

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


## TeamsInstalledAppsUpgrade

> TeamsInstalledAppsUpgrade(ctx, teamId, teamsAppInstallationId)
Invoke action upgrade

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**teamsAppInstallationId** | **string**| key: teamsAppInstallation-id of teamsAppInstallation | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUnarchive

> TeamsUnarchive(ctx, teamId)
Invoke action unarchive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

