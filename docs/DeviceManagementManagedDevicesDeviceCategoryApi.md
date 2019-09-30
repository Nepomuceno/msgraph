# \DeviceManagementManagedDevicesDeviceCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementManagedDevicesGetDeviceCategory**](DeviceManagementManagedDevicesDeviceCategoryApi.md#DeviceManagementManagedDevicesGetDeviceCategory) | **Get** /deviceManagement/managedDevices({managedDevice-id})/deviceCategory | Get deviceCategory from deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceCategory**](DeviceManagementManagedDevicesDeviceCategoryApi.md#DeviceManagementManagedDevicesUpdateDeviceCategory) | **Patch** /deviceManagement/managedDevices({managedDevice-id})/deviceCategory | Update the navigation property deviceCategory in deviceManagement



## DeviceManagementManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory DeviceManagementManagedDevicesGetDeviceCategory(ctx, managedDeviceId, optional)
Get deviceCategory from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***DeviceManagementManagedDevicesGetDeviceCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementManagedDevicesGetDeviceCategoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCategory**](microsoft.graph.deviceCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesUpdateDeviceCategory

> DeviceManagementManagedDevicesUpdateDeviceCategory(ctx, managedDeviceId, microsoftGraphDeviceCategory)
Update the navigation property deviceCategory in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphDeviceCategory** | [**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md)| New navigation property values | 

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

