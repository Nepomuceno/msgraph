# \DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesCreateAssignments**](DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyAssignmentApi.md#DeviceManagementDeviceCompliancePoliciesCreateAssignments) | **Post** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetAssignments**](DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyAssignmentApi.md#DeviceManagementDeviceCompliancePoliciesGetAssignments) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/assignments({deviceCompliancePolicyAssignment-id}) | Get assignments from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListAssignments**](DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyAssignmentApi.md#DeviceManagementDeviceCompliancePoliciesListAssignments) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/assignments | Get assignments from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateAssignments**](DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyAssignmentApi.md#DeviceManagementDeviceCompliancePoliciesUpdateAssignments) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/assignments({deviceCompliancePolicyAssignment-id}) | Update the navigation property assignments in deviceManagement



## DeviceManagementDeviceCompliancePoliciesCreateAssignments

> MicrosoftGraphDeviceCompliancePolicyAssignment DeviceManagementDeviceCompliancePoliciesCreateAssignments(ctx, deviceCompliancePolicyId, microsoftGraphDeviceCompliancePolicyAssignment)
Create new navigation property to assignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceCompliancePolicyAssignment** | [**MicrosoftGraphDeviceCompliancePolicyAssignment**](MicrosoftGraphDeviceCompliancePolicyAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyAssignment**](microsoft.graph.deviceCompliancePolicyAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetAssignments

> MicrosoftGraphDeviceCompliancePolicyAssignment DeviceManagementDeviceCompliancePoliciesGetAssignments(ctx, deviceCompliancePolicyId, deviceCompliancePolicyAssignmentId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceCompliancePolicyAssignmentId** | **string**| key: deviceCompliancePolicyAssignment-id of deviceCompliancePolicyAssignment | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyAssignment**](microsoft.graph.deviceCompliancePolicyAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListAssignments

> CollectionOfDeviceCompliancePolicyAssignment DeviceManagementDeviceCompliancePoliciesListAssignments(ctx, deviceCompliancePolicyId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **search** | **optional.String**| Search items by search phrases | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**CollectionOfDeviceCompliancePolicyAssignment**](Collection of deviceCompliancePolicyAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesUpdateAssignments

> DeviceManagementDeviceCompliancePoliciesUpdateAssignments(ctx, deviceCompliancePolicyId, deviceCompliancePolicyAssignmentId, microsoftGraphDeviceCompliancePolicyAssignment)
Update the navigation property assignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceCompliancePolicyAssignmentId** | **string**| key: deviceCompliancePolicyAssignment-id of deviceCompliancePolicyAssignment | 
**microsoftGraphDeviceCompliancePolicyAssignment** | [**MicrosoftGraphDeviceCompliancePolicyAssignment**](MicrosoftGraphDeviceCompliancePolicyAssignment.md)| New navigation property values | 

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

