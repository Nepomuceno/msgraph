# \MeManagedDevicesDeviceCompliancePolicyStateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeManagedDevicesCreateDeviceCompliancePolicyStates**](MeManagedDevicesDeviceCompliancePolicyStateApi.md#MeManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /me/managedDevices({managedDevice-id})/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for me
[**MeManagedDevicesGetDeviceCompliancePolicyStates**](MeManagedDevicesDeviceCompliancePolicyStateApi.md#MeManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /me/managedDevices({managedDevice-id})/deviceCompliancePolicyStates({deviceCompliancePolicyState-id}) | Get deviceCompliancePolicyStates from me
[**MeManagedDevicesListDeviceCompliancePolicyStates**](MeManagedDevicesDeviceCompliancePolicyStateApi.md#MeManagedDevicesListDeviceCompliancePolicyStates) | **Get** /me/managedDevices({managedDevice-id})/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from me
[**MeManagedDevicesUpdateDeviceCompliancePolicyStates**](MeManagedDevicesDeviceCompliancePolicyStateApi.md#MeManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /me/managedDevices({managedDevice-id})/deviceCompliancePolicyStates({deviceCompliancePolicyState-id}) | Update the navigation property deviceCompliancePolicyStates in me



## MeManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState MeManagedDevicesCreateDeviceCompliancePolicyStates(ctx, managedDeviceId, microsoftGraphDeviceCompliancePolicyState)
Create new navigation property to deviceCompliancePolicyStates for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## MeManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState MeManagedDevicesGetDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId, optional)
Get deviceCompliancePolicyStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**deviceCompliancePolicyStateId** | **string**| key: deviceCompliancePolicyState-id of deviceCompliancePolicyState | 
 **optional** | ***MeManagedDevicesGetDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesGetDeviceCompliancePolicyStatesOpts struct


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


## MeManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState MeManagedDevicesListDeviceCompliancePolicyStates(ctx, managedDeviceId, optional)
Get deviceCompliancePolicyStates from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeManagedDevicesListDeviceCompliancePolicyStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeManagedDevicesListDeviceCompliancePolicyStatesOpts struct


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


## MeManagedDevicesUpdateDeviceCompliancePolicyStates

> MeManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId, microsoftGraphDeviceCompliancePolicyState)
Update the navigation property deviceCompliancePolicyStates in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

