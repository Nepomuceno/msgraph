# \DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceOverviewApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceConfigurationsGetDeviceStatusOverview**](DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceOverviewApi.md#DeviceManagementDeviceConfigurationsGetDeviceStatusOverview) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceStatusOverview | Get deviceStatusOverview from deviceManagement
[**DeviceManagementDeviceConfigurationsUpdateDeviceStatusOverview**](DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceOverviewApi.md#DeviceManagementDeviceConfigurationsUpdateDeviceStatusOverview) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceStatusOverview | Update the navigation property deviceStatusOverview in deviceManagement



## DeviceManagementDeviceConfigurationsGetDeviceStatusOverview

> MicrosoftGraphDeviceConfigurationDeviceOverview DeviceManagementDeviceConfigurationsGetDeviceStatusOverview(ctx, deviceConfigurationId, optional)
Get deviceStatusOverview from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementDeviceConfigurationsGetDeviceStatusOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsGetDeviceStatusOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationDeviceOverview**](microsoft.graph.deviceConfigurationDeviceOverview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsUpdateDeviceStatusOverview

> DeviceManagementDeviceConfigurationsUpdateDeviceStatusOverview(ctx, deviceConfigurationId, microsoftGraphDeviceConfigurationDeviceOverview)
Update the navigation property deviceStatusOverview in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphDeviceConfigurationDeviceOverview** | [**MicrosoftGraphDeviceConfigurationDeviceOverview**](MicrosoftGraphDeviceConfigurationDeviceOverview.md)| New navigation property values | 

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

