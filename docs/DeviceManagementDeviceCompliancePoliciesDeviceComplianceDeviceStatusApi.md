# \DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceStatusApi.md#DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses) | **Post** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/deviceStatuses | Create new navigation property to deviceStatuses for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceStatusApi.md#DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/deviceStatuses({deviceComplianceDeviceStatus-id}) | Get deviceStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListDeviceStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceStatusApi.md#DeviceManagementDeviceCompliancePoliciesListDeviceStatuses) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/deviceStatuses | Get deviceStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceDeviceStatusApi.md#DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/deviceStatuses({deviceComplianceDeviceStatus-id}) | Update the navigation property deviceStatuses in deviceManagement



## DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses

> MicrosoftGraphDeviceComplianceDeviceStatus DeviceManagementDeviceCompliancePoliciesCreateDeviceStatuses(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceDeviceStatus)
Create new navigation property to deviceStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceDeviceStatus** | [**MicrosoftGraphDeviceComplianceDeviceStatus**](MicrosoftGraphDeviceComplianceDeviceStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceDeviceStatus**](microsoft.graph.deviceComplianceDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses

> MicrosoftGraphDeviceComplianceDeviceStatus DeviceManagementDeviceCompliancePoliciesGetDeviceStatuses(ctx, deviceCompliancePolicyId, deviceComplianceDeviceStatusId, optional)
Get deviceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceDeviceStatusId** | **string**| key: deviceComplianceDeviceStatus-id of deviceComplianceDeviceStatus | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetDeviceStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceDeviceStatus**](microsoft.graph.deviceComplianceDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListDeviceStatuses

> CollectionOfDeviceComplianceDeviceStatus DeviceManagementDeviceCompliancePoliciesListDeviceStatuses(ctx, deviceCompliancePolicyId, optional)
Get deviceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListDeviceStatusesOpts struct


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

[**CollectionOfDeviceComplianceDeviceStatus**](Collection of deviceComplianceDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses

> DeviceManagementDeviceCompliancePoliciesUpdateDeviceStatuses(ctx, deviceCompliancePolicyId, deviceComplianceDeviceStatusId, microsoftGraphDeviceComplianceDeviceStatus)
Update the navigation property deviceStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceDeviceStatusId** | **string**| key: deviceComplianceDeviceStatus-id of deviceComplianceDeviceStatus | 
**microsoftGraphDeviceComplianceDeviceStatus** | [**MicrosoftGraphDeviceComplianceDeviceStatus**](MicrosoftGraphDeviceComplianceDeviceStatus.md)| New navigation property values | 

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

