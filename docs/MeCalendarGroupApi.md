# \MeCalendarGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateCalendarGroups**](MeCalendarGroupApi.md#MeCreateCalendarGroups) | **Post** /me/calendarGroups | Create new navigation property to calendarGroups for me
[**MeGetCalendarGroups**](MeCalendarGroupApi.md#MeGetCalendarGroups) | **Get** /me/calendarGroups({calendarGroup-id}) | Get calendarGroups from me
[**MeListCalendarGroups**](MeCalendarGroupApi.md#MeListCalendarGroups) | **Get** /me/calendarGroups | Get calendarGroups from me
[**MeUpdateCalendarGroups**](MeCalendarGroupApi.md#MeUpdateCalendarGroups) | **Patch** /me/calendarGroups({calendarGroup-id}) | Update the navigation property calendarGroups in me



## MeCreateCalendarGroups

> MicrosoftGraphCalendarGroup MeCreateCalendarGroups(ctx, microsoftGraphCalendarGroup)
Create new navigation property to calendarGroups for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphCalendarGroup** | [**MicrosoftGraphCalendarGroup**](MicrosoftGraphCalendarGroup.md)| New navigation property | 

### Return type

[**MicrosoftGraphCalendarGroup**](microsoft.graph.calendarGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetCalendarGroups

> MicrosoftGraphCalendarGroup MeGetCalendarGroups(ctx, calendarGroupId, optional)
Get calendarGroups from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
 **optional** | ***MeGetCalendarGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetCalendarGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphCalendarGroup**](microsoft.graph.calendarGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListCalendarGroups

> CollectionOfCalendarGroup MeListCalendarGroups(ctx, optional)
Get calendarGroups from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListCalendarGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListCalendarGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfCalendarGroup**](Collection of calendarGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateCalendarGroups

> MeUpdateCalendarGroups(ctx, calendarGroupId, microsoftGraphCalendarGroup)
Update the navigation property calendarGroups in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calendarGroupId** | **string**| key: calendarGroup-id of calendarGroup | 
**microsoftGraphCalendarGroup** | [**MicrosoftGraphCalendarGroup**](MicrosoftGraphCalendarGroup.md)| New navigation property values | 

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

