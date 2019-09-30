# \DeviceManagementManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementCreateManagedDevices) | **Post** /deviceManagement/managedDevices | Create new navigation property to managedDevices for deviceManagement
[**DeviceManagementGetManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementGetManagedDevices) | **Get** /deviceManagement/managedDevices({managedDevice-id}) | Get managedDevices from deviceManagement
[**DeviceManagementListManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementListManagedDevices) | **Get** /deviceManagement/managedDevices | Get managedDevices from deviceManagement
[**DeviceManagementUpdateManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementUpdateManagedDevices) | **Patch** /deviceManagement/managedDevices({managedDevice-id}) | Update the navigation property managedDevices in deviceManagement



## DeviceManagementCreateManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementCreateManagedDevices(ctx, microsoftGraphManagedDevice)
Create new navigation property to managedDevices for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementGetManagedDevices(ctx, managedDeviceId, optional)
Get managedDevices from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementGetManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetManagedDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListManagedDevices

> CollectionOfManagedDevice DeviceManagementListManagedDevices(ctx, optional)
Get managedDevices from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListManagedDevicesOpts struct


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

[**CollectionOfManagedDevice**](Collection of managedDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateManagedDevices

> DeviceManagementUpdateManagedDevices(ctx, managedDeviceId, microsoftGraphManagedDevice)
Update the navigation property managedDevices in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)| New navigation property values | 

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

