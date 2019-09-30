# \DeviceAppManagementDefaultManagedAppProtectionsManagedAppPolicyDeploymentSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementDefaultManagedAppProtectionsManagedAppPolicyDeploymentSummaryApi.md#DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id})/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementDefaultManagedAppProtectionsManagedAppPolicyDeploymentSummaryApi.md#DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id})/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement



## DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary(ctx, defaultManagedAppProtectionId, optional)
Get deploymentSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
 **optional** | ***DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummaryOpts struct


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


## DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary(ctx, defaultManagedAppProtectionId, microsoftGraphManagedAppPolicyDeploymentSummary)
Update the navigation property deploymentSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
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

