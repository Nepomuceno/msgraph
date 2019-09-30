# \MeManagedDevicesDeviceConfigurationStateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeManagedDevicesCreateDeviceConfigurationStates**](MeManagedDevicesDeviceConfigurationStateApi.md#MeManagedDevicesCreateDeviceConfigurationStates) | **Post** /me/managedDevices({managedDevice-id})/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for me
[**MeManagedDevicesGetDeviceConfigurationStates**](MeManagedDevicesDeviceConfigurationStateApi.md#MeManagedDevicesGetDeviceConfigurationStates) | **Get** /me/managedDevices({managedDevice-id})/deviceConfigurationStates({deviceConfigurationState-id}) | Get deviceConfigurationStates from me
[**MeManagedDevicesListDeviceConfigurationStates**](MeManagedDevicesDeviceConfigurationStateApi.md#MeManagedDevicesListDeviceConfigurationStates) | **Get** /me/managedDevices({managedDevice-id})/deviceConfigurationStates | Get deviceConfigurationStates from me
[**MeManagedDevicesUpdateDeviceConfigurationStates**](MeManagedDevicesDeviceConfigurationStateApi.md#MeManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /me/managedDevices({managedDevice-id})/deviceConfigurationStates({deviceConfigurationState-id}) | Update the navigation property deviceConfigurationStates in me



## MeManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState MeManagedDevicesCreateDeviceConfigurationStates(ctx, managedDeviceId, microsoftGraphDeviceConfigurationState)
Create new navigation property to deviceConfigurationStates for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState MeManagedDevicesGetDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId, optional)
Get deviceConfigurationStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceConfigurationStateId** | **string**| key: deviceConfigurationState-id of deviceConfigurationState | 
 **optional** | ***MeManagedDevicesGetDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesGetDeviceConfigurationStatesOpts struct


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


## MeManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState MeManagedDevicesListDeviceConfigurationStates(ctx, managedDeviceId, optional)
Get deviceConfigurationStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeManagedDevicesListDeviceConfigurationStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesListDeviceConfigurationStatesOpts struct


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


## MeManagedDevicesUpdateDeviceConfigurationStates

> MeManagedDevicesUpdateDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId, microsoftGraphDeviceConfigurationState)
Update the navigation property deviceConfigurationStates in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

