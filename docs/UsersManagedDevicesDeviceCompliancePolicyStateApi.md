# \UsersManagedDevicesDeviceCompliancePolicyStateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersManagedDevicesCreateDeviceCompliancePolicyStates**](UsersManagedDevicesDeviceCompliancePolicyStateApi.md#UsersManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /users({user-id})/managedDevices({managedDevice-id})/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for users
[**UsersManagedDevicesGetDeviceCompliancePolicyStates**](UsersManagedDevicesDeviceCompliancePolicyStateApi.md#UsersManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /users({user-id})/managedDevices({managedDevice-id})/deviceCompliancePolicyStates({deviceCompliancePolicyState-id}) | Get deviceCompliancePolicyStates from users
[**UsersManagedDevicesListDeviceCompliancePolicyStates**](UsersManagedDevicesDeviceCompliancePolicyStateApi.md#UsersManagedDevicesListDeviceCompliancePolicyStates) | **Get** /users({user-id})/managedDevices({managedDevice-id})/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from users
[**UsersManagedDevicesUpdateDeviceCompliancePolicyStates**](UsersManagedDevicesDeviceCompliancePolicyStateApi.md#UsersManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /users({user-id})/managedDevices({managedDevice-id})/deviceCompliancePolicyStates({deviceCompliancePolicyState-id}) | Update the navigation property deviceCompliancePolicyStates in users



## UsersManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState UsersManagedDevicesCreateDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, microsoftGraphDeviceCompliancePolicyState)
Create new navigation property to deviceCompliancePolicyStates for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphDeviceCompliancePolicyState** | [**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyState**](microsoft.graph.deviceCompliancePolicyState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState UsersManagedDevicesGetDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId, optional)
Get deviceCompliancePolicyStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
 **optional** | ***UsersManagedDevicesGetDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesGetDeviceCompliancePolicyStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyState**](microsoft.graph.deviceCompliancePolicyState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState UsersManagedDevicesListDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, optional)
Get deviceCompliancePolicyStates from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersManagedDevicesListDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesListDeviceCompliancePolicyStatesOpts struct


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

[**CollectionOfDeviceCompliancePolicyState**](Collection of deviceCompliancePolicyState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersManagedDevicesUpdateDeviceCompliancePolicyStates

> UsersManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId, microsoftGraphDeviceCompliancePolicyState)
Update the navigation property deviceCompliancePolicyStates in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
**microsoftGraphDeviceCompliancePolicyState** | [**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md)| New navigation property values | 

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

