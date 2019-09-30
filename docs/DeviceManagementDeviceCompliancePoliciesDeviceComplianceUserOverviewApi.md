# \DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserOverviewApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserOverviewApi.md#DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/userStatusOverview | Get userStatusOverview from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserOverviewApi.md#DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/userStatusOverview | Update the navigation property userStatusOverview in deviceManagement



## DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview

> MicrosoftGraphDeviceComplianceUserOverview DeviceManagementDeviceCompliancePoliciesGetUserStatusOverview(ctx, deviceCompliancePolicyId, optional)
Get userStatusOverview from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetUserStatusOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetUserStatusOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceUserOverview**](microsoft.graph.deviceComplianceUserOverview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview

> DeviceManagementDeviceCompliancePoliciesUpdateUserStatusOverview(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceUserOverview)
Update the navigation property userStatusOverview in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceUserOverview** | [**MicrosoftGraphDeviceComplianceUserOverview**](MicrosoftGraphDeviceComplianceUserOverview.md)| New navigation property values | 

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

