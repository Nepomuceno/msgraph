# \DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceOverviewApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceOverviewApi.md#DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/deviceStatusOverview | Get deviceStatusOverview from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceOverviewApi.md#DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/deviceStatusOverview | Update the navigation property deviceStatusOverview in deviceManagement



## DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview

> MicrosoftGraphDeviceComplianceDeviceOverview DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverview(ctx, deviceCompliancePolicyId, optional)
Get deviceStatusOverview from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetDeviceStatusOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceDeviceOverview**](microsoft.graph.deviceComplianceDeviceOverview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview

> DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatusOverview(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceDeviceOverview)
Update the navigation property deviceStatusOverview in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceDeviceOverview** | [**MicrosoftGraphDeviceComplianceDeviceOverview**](MicrosoftGraphDeviceComplianceDeviceOverview.md)| New navigation property values | 

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

