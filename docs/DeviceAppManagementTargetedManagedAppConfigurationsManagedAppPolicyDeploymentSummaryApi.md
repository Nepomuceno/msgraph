# \DeviceAppManagementTargetedManagedAppConfigurationsManagedAppPolicyDeploymentSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary**](DeviceAppManagementTargetedManagedAppConfigurationsManagedAppPolicyDeploymentSummaryApi.md#DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary) | **Get** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary**](DeviceAppManagementTargetedManagedAppConfigurationsManagedAppPolicyDeploymentSummaryApi.md#DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement



## DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummary(ctx, targetedManagedAppConfigurationId, optional)
Get deploymentSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsGetDeploymentSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicyDeploymentSummary**](microsoft.graph.managedAppPolicyDeploymentSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary

> DeviceAppManagementTargetedManagedAppConfigurationsUpdateDeploymentSummary(ctx, targetedManagedAppConfigurationId, microsoftGraphManagedAppPolicyDeploymentSummary)
Update the navigation property deploymentSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**microsoftGraphManagedAppPolicyDeploymentSummary** | [**MicrosoftGraphManagedAppPolicyDeploymentSummary**](MicrosoftGraphManagedAppPolicyDeploymentSummary.md)| New navigation property values | 

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

