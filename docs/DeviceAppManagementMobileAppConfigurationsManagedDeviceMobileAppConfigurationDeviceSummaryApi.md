# \DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppConfigurationsGetDeviceStatusSummary**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceSummaryApi.md#DeviceAppManagementMobileAppConfigurationsGetDeviceStatusSummary) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/deviceStatusSummary | Get deviceStatusSummary from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatusSummary**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceSummaryApi.md#DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatusSummary) | **Patch** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/deviceStatusSummary | Update the navigation property deviceStatusSummary in deviceAppManagement



## DeviceAppManagementMobileAppConfigurationsGetDeviceStatusSummary

> MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary DeviceAppManagementMobileAppConfigurationsGetDeviceStatusSummary(ctx, managedDeviceMobileAppConfigurationId, optional)
Get deviceStatusSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsGetDeviceStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsGetDeviceStatusSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary**](microsoft.graph.managedDeviceMobileAppConfigurationDeviceSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatusSummary

> DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatusSummary(ctx, managedDeviceMobileAppConfigurationId, microsoftGraphManagedDeviceMobileAppConfigurationDeviceSummary)
Update the navigation property deviceStatusSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**microsoftGraphManagedDeviceMobileAppConfigurationDeviceSummary** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary**](MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary.md)| New navigation property values | 

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

