# \UsersManagedDevicesDeviceConfigurationStateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersManagedDevicesCreateDeviceConfigurationStates**](UsersManagedDevicesDeviceConfigurationStateApi.md#UsersManagedDevicesCreateDeviceConfigurationStates) | **Post** /users({user-id})/managedDevices({managedDevice-id})/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for users
[**UsersManagedDevicesGetDeviceConfigurationStates**](UsersManagedDevicesDeviceConfigurationStateApi.md#UsersManagedDevicesGetDeviceConfigurationStates) | **Get** /users({user-id})/managedDevices({managedDevice-id})/deviceConfigurationStates({deviceConfigurationState-id}) | Get deviceConfigurationStates from users
[**UsersManagedDevicesListDeviceConfigurationStates**](UsersManagedDevicesDeviceConfigurationStateApi.md#UsersManagedDevicesListDeviceConfigurationStates) | **Get** /users({user-id})/managedDevices({managedDevice-id})/deviceConfigurationStates | Get deviceConfigurationStates from users
[**UsersManagedDevicesUpdateDeviceConfigurationStates**](UsersManagedDevicesDeviceConfigurationStateApi.md#UsersManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /users({user-id})/managedDevices({managedDevice-id})/deviceConfigurationStates({deviceConfigurationState-id}) | Update the navigation property deviceConfigurationStates in users



## UsersManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState UsersManagedDevicesCreateDeviceConfigurationStates(ctx, userId, managedDeviceId, microsoftGraphDeviceConfigurationState)
Create new navigation property to deviceConfigurationStates for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphDeviceConfigurationState** | [**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfigurationState**](microsoft.graph.deviceConfigurationState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState UsersManagedDevicesGetDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId, optional)
Get deviceConfigurationStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
 **optional** | ***UsersManagedDevicesGetDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesGetDeviceConfigurationStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationState**](microsoft.graph.deviceConfigurationState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState UsersManagedDevicesListDeviceConfigurationStates(ctx, userId, managedDeviceId, optional)
Get deviceConfigurationStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersManagedDevicesListDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesListDeviceConfigurationStatesOpts struct


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

[**CollectionOfDeviceConfigurationState**](Collection of deviceConfigurationState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesUpdateDeviceConfigurationStates

> UsersManagedDevicesUpdateDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId, microsoftGraphDeviceConfigurationState)
Update the navigation property deviceConfigurationStates in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
**microsoftGraphDeviceConfigurationState** | [**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md)| New navigation property values | 

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

