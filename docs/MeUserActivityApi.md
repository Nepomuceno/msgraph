# \MeUserActivityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateActivities**](MeUserActivityApi.md#MeCreateActivities) | **Post** /me/activities | Create new navigation property to activities for me
[**MeGetActivities**](MeUserActivityApi.md#MeGetActivities) | **Get** /me/activities({userActivity-id}) | Get activities from me
[**MeListActivities**](MeUserActivityApi.md#MeListActivities) | **Get** /me/activities | Get activities from me
[**MeUpdateActivities**](MeUserActivityApi.md#MeUpdateActivities) | **Patch** /me/activities({userActivity-id}) | Update the navigation property activities in me



## MeCreateActivities

> MicrosoftGraphUserActivity MeCreateActivities(ctx, microsoftGraphUserActivity)
Create new navigation property to activities for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)| New navigation property | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetActivities

> MicrosoftGraphUserActivity MeGetActivities(ctx, userActivityId, optional)
Get activities from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string**| key: userActivity-id of userActivity | 
 **optional** | ***MeGetActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetActivitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListActivities

> CollectionOfUserActivity MeListActivities(ctx, optional)
Get activities from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListActivitiesOpts struct


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

[**CollectionOfUserActivity**](Collection of userActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateActivities

> MeUpdateActivities(ctx, userActivityId, microsoftGraphUserActivity)
Update the navigation property activities in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string**| key: userActivity-id of userActivity | 
**microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)| New navigation property values | 

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

