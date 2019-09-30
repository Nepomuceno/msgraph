# \TeamsChannelsTabsTeamsAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsChannelsTabsGetTeamsApp**](TeamsChannelsTabsTeamsAppApi.md#TeamsChannelsTabsGetTeamsApp) | **Get** /teams({team-id})/channels({channel-id})/tabs({teamsTab-id})/teamsApp | Get teamsApp from teams



## TeamsChannelsTabsGetTeamsApp

> MicrosoftGraphTeamsApp TeamsChannelsTabsGetTeamsApp(ctx, teamId, channelId, teamsTabId, optional)
Get teamsApp from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
**teamsTabId** | **string**| key: teamsTab-id of teamsTab | 
 **optional** | ***TeamsChannelsTabsGetTeamsAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsChannelsTabsGetTeamsAppOpts struct


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

