# \DeviceAppManagementAndroidManagedAppProtectionsManagedAppPolicyDeploymentSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionsManagedAppPolicyDeploymentSummaryApi.md#DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/androidManagedAppProtections({androidManagedAppProtection-id})/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionsManagedAppPolicyDeploymentSummaryApi.md#DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/androidManagedAppProtections({androidManagedAppProtection-id})/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement



## DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary(ctx, androidManagedAppProtectionId, optional)
Get deploymentSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
 **optional** | ***DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummaryOpts struct


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


## DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary(ctx, androidManagedAppProtectionId, microsoftGraphManagedAppPolicyDeploymentSummary)
Update the navigation property deploymentSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
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

