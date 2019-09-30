# \DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppConfigurationsGetUserStatusSummary**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserSummaryApi.md#DeviceAppManagementMobileAppConfigurationsGetUserStatusSummary) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/userStatusSummary | Get userStatusSummary from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsUpdateUserStatusSummary**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserSummaryApi.md#DeviceAppManagementMobileAppConfigurationsUpdateUserStatusSummary) | **Patch** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/userStatusSummary | Update the navigation property userStatusSummary in deviceAppManagement



## DeviceAppManagementMobileAppConfigurationsGetUserStatusSummary

> MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary DeviceAppManagementMobileAppConfigurationsGetUserStatusSummary(ctx, managedDeviceMobileAppConfigurationId, optional)
Get userStatusSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsGetUserStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsGetUserStatusSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary**](microsoft.graph.managedDeviceMobileAppConfigurationUserSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsUpdateUserStatusSummary

> DeviceAppManagementMobileAppConfigurationsUpdateUserStatusSummary(ctx, managedDeviceMobileAppConfigurationId, microsoftGraphManagedDeviceMobileAppConfigurationUserSummary)
Update the navigation property userStatusSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**microsoftGraphManagedDeviceMobileAppConfigurationUserSummary** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary**](MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary.md)| New navigation property values | 

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

