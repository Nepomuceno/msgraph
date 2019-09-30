# \DeviceManagementDeviceConfigurationsDeviceConfigurationUserOverviewApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceConfigurationsGetUserStatusOverview**](DeviceManagementDeviceConfigurationsDeviceConfigurationUserOverviewApi.md#DeviceManagementDeviceConfigurationsGetUserStatusOverview) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/userStatusOverview | Get userStatusOverview from deviceManagement
[**DeviceManagementDeviceConfigurationsUpdateUserStatusOverview**](DeviceManagementDeviceConfigurationsDeviceConfigurationUserOverviewApi.md#DeviceManagementDeviceConfigurationsUpdateUserStatusOverview) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/userStatusOverview | Update the navigation property userStatusOverview in deviceManagement



## DeviceManagementDeviceConfigurationsGetUserStatusOverview

> MicrosoftGraphDeviceConfigurationUserOverview DeviceManagementDeviceConfigurationsGetUserStatusOverview(ctx, deviceConfigurationId, optional)
Get userStatusOverview from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementDeviceConfigurationsGetUserStatusOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsGetUserStatusOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationUserOverview**](microsoft.graph.deviceConfigurationUserOverview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsUpdateUserStatusOverview

> DeviceManagementDeviceConfigurationsUpdateUserStatusOverview(ctx, deviceConfigurationId, microsoftGraphDeviceConfigurationUserOverview)
Update the navigation property userStatusOverview in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphDeviceConfigurationUserOverview** | [**MicrosoftGraphDeviceConfigurationUserOverview**](MicrosoftGraphDeviceConfigurationUserOverview.md)| New navigation property values | 

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

