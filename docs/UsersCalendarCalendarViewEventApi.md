# \UsersCalendarCalendarViewEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCalendarCalendarViewCreateInstances**](UsersCalendarCalendarViewEventApi.md#UsersCalendarCalendarViewCreateInstances) | **Post** /users({user-id})/calendar/calendarView({event-id})/instances | Create new navigation property to instances for users
[**UsersCalendarCalendarViewGetInstances**](UsersCalendarCalendarViewEventApi.md#UsersCalendarCalendarViewGetInstances) | **Get** /users({user-id})/calendar/calendarView({event-id})/instances({event-id1}) | Get instances from users
[**UsersCalendarCalendarViewListInstances**](UsersCalendarCalendarViewEventApi.md#UsersCalendarCalendarViewListInstances) | **Get** /users({user-id})/calendar/calendarView({event-id})/instances | Get instances from users
[**UsersCalendarCalendarViewUpdateInstances**](UsersCalendarCalendarViewEventApi.md#UsersCalendarCalendarViewUpdateInstances) | **Patch** /users({user-id})/calendar/calendarView({event-id})/instances({event-id1}) | Update the navigation property instances in users



## UsersCalendarCalendarViewCreateInstances

> MicrosoftGraphEvent UsersCalendarCalendarViewCreateInstances(ctx, userId, eventId, microsoftGraphEvent)
Create new navigation property to instances for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**microsoftGraphEvent** | [**MicrosoftGraphEvent**](MicrosoftGraphEvent.md)| New navigation property | 

### Return type

[**MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarCalendarViewGetInstances

> MicrosoftGraphEvent UsersCalendarCalendarViewGetInstances(ctx, userId, eventId, eventId1, optional)
Get instances from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarCalendarViewGetInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarCalendarViewGetInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphEvent**](microsoft.graph.event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarCalendarViewListInstances

> CollectionOfEvent UsersCalendarCalendarViewListInstances(ctx, userId, eventId, optional)
Get instances from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
 **optional** | ***UsersCalendarCalendarViewListInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersCalendarCalendarViewListInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfEvent**](Collection of event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCalendarCalendarViewUpdateInstances

> UsersCalendarCalendarViewUpdateInstances(ctx, userId, eventId, eventId1, microsoftGraphEvent)
Update the navigation property instances in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**eventId** | **string**| key: event-id of event | 
**eventId1** | **string**| key: event-id of event | 
**microsoftGraphEvent** | [**MicrosoftGraphEvent**](MicrosoftGraphEvent.md)| New navigation property values | 

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

