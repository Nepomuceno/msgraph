# \TeamsChannelApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsCreateChannels**](TeamsChannelApi.md#TeamsCreateChannels) | **Post** /teams({team-id})/channels | Create new navigation property to channels for teams
[**TeamsGetChannels**](TeamsChannelApi.md#TeamsGetChannels) | **Get** /teams({team-id})/channels({channel-id}) | Get channels from teams
[**TeamsListChannels**](TeamsChannelApi.md#TeamsListChannels) | **Get** /teams({team-id})/channels | Get channels from teams
[**TeamsUpdateChannels**](TeamsChannelApi.md#TeamsUpdateChannels) | **Patch** /teams({team-id})/channels({channel-id}) | Update the navigation property channels in teams



## TeamsCreateChannels

> MicrosoftGraphChannel TeamsCreateChannels(ctx, teamId, microsoftGraphChannel)
Create new navigation property to channels for teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**microsoftGraphChannel** | [**MicrosoftGraphChannel**](MicrosoftGraphChannel.md)| New navigation property | 

### Return type

[**MicrosoftGraphChannel**](microsoft.graph.channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetChannels

> MicrosoftGraphChannel TeamsGetChannels(ctx, teamId, channelId, optional)
Get channels from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
 **optional** | ***TeamsGetChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsGetChannelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphChannel**](microsoft.graph.channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListChannels

> CollectionOfChannel TeamsListChannels(ctx, teamId, optional)
Get channels from teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
 **optional** | ***TeamsListChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TeamsListChannelsOpts struct


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

[**CollectionOfChannel**](Collection of channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateChannels

> TeamsUpdateChannels(ctx, teamId, channelId, microsoftGraphChannel)
Update the navigation property channels in teams

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| key: team-id of team | 
**channelId** | **string**| key: channel-id of channel | 
**microsoftGraphChannel** | [**MicrosoftGraphChannel**](MicrosoftGraphChannel.md)| New navigation property values | 

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

