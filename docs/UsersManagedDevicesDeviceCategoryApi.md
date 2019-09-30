# \UsersManagedDevicesDeviceCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersManagedDevicesGetDeviceCategory**](UsersManagedDevicesDeviceCategoryApi.md#UsersManagedDevicesGetDeviceCategory) | **Get** /users({user-id})/managedDevices({managedDevice-id})/deviceCategory | Get deviceCategory from users
[**UsersManagedDevicesUpdateDeviceCategory**](UsersManagedDevicesDeviceCategoryApi.md#UsersManagedDevicesUpdateDeviceCategory) | **Patch** /users({user-id})/managedDevices({managedDevice-id})/deviceCategory | Update the navigation property deviceCategory in users



## UsersManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory UsersManagedDevicesGetDeviceCategory(ctx, userId, managedDeviceId, optional)
Get deviceCategory from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***UsersManagedDevicesGetDeviceCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersManagedDevicesGetDeviceCategoryOpts struct


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


## UsersManagedDevicesUpdateDeviceCategory

> UsersManagedDevicesUpdateDeviceCategory(ctx, userId, managedDeviceId, microsoftGraphDeviceCategory)
Update the navigation property deviceCategory in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
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

