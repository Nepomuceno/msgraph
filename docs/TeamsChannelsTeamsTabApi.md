# \TeamsChannelsTeamsTabApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsChannelsCreateTabs**](TeamsChannelsTeamsTabApi.md#TeamsChannelsCreateTabs) | **Post** /teams({team-id})/channels({channel-id})/tabs | Create new navigation property to tabs for teams
[**TeamsChannelsGetTabs**](TeamsChannelsTeamsTabApi.md#TeamsChannelsGetTabs) | **Get** /teams({team-id})/channels({channel-id})/tabs({teamsTab-id}) | Get tabs from teams
[**TeamsChannelsListTabs**](TeamsChannelsTeamsTabApi.md#TeamsChannelsListTabs) | **Get** /teams({team-id})/channels({channel-id})/tabs | Get tabs from teams
[**TeamsChannelsUpdateTabs**](TeamsChannelsTeamsTabApi.md#TeamsChannelsUpdateTabs) | **Patch** /teams({team-id})/channels({channel-id})/tabs({teamsTab-id}) | Update the navigation property tabs in teams



## TeamsChannelsCreateTabs

> MicrosoftGraphTeamsTab TeamsChannelsCreateTabs(ctx, teamId, channelId, microsoftGraphTeamsTab)
Create new navigation property to tabs for teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
**microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)| New navigation property | 

### Return type

[**MicrosoftGraphTeamsTab**](microsoft.graph.teamsTab.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsGetTabs

> MicrosoftGraphTeamsTab TeamsChannelsGetTabs(ctx, teamId, channelId, teamsTabId, optional)
Get tabs from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
**teamsTabId** | **string**| key: teamsTab-id of teamsTab | 
 **optional** | ***TeamsChannelsGetTabsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsChannelsGetTabsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTeamsTab**](microsoft.graph.teamsTab.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsListTabs

> CollectionOfTeamsTab TeamsChannelsListTabs(ctx, teamId, channelId, optional)
Get tabs from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
 **optional** | ***TeamsChannelsListTabsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsChannelsListTabsOpts struct


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

[**CollectionOfTeamsTab**](Collection of teamsTab.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsUpdateTabs

> TeamsChannelsUpdateTabs(ctx, teamId, channelId, teamsTabId, microsoftGraphTeamsTab)
Update the navigation property tabs in teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
**teamsTabId** | **string**| key: teamsTab-id of teamsTab | 
**microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)| New navigation property values | 

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

