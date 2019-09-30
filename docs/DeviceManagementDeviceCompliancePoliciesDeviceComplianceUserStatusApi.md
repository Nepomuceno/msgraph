# \DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesCreateUserStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserStatusApi.md#DeviceManagementDeviceCompliancePoliciesCreateUserStatuses) | **Post** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/userStatuses | Create new navigation property to userStatuses for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetUserStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserStatusApi.md#DeviceManagementDeviceCompliancePoliciesGetUserStatuses) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/userStatuses({deviceComplianceUserStatus-id}) | Get userStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListUserStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserStatusApi.md#DeviceManagementDeviceCompliancePoliciesListUserStatuses) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/userStatuses | Get userStatuses from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceUserStatusApi.md#DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/userStatuses({deviceComplianceUserStatus-id}) | Update the navigation property userStatuses in deviceManagement



## DeviceManagementDeviceCompliancePoliciesCreateUserStatuses

> MicrosoftGraphDeviceComplianceUserStatus DeviceManagementDeviceCompliancePoliciesCreateUserStatuses(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceUserStatus)
Create new navigation property to userStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceUserStatus** | [**MicrosoftGraphDeviceComplianceUserStatus**](MicrosoftGraphDeviceComplianceUserStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceUserStatus**](microsoft.graph.deviceComplianceUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetUserStatuses

> MicrosoftGraphDeviceComplianceUserStatus DeviceManagementDeviceCompliancePoliciesGetUserStatuses(ctx, deviceCompliancePolicyId, deviceComplianceUserStatusId, optional)
Get userStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceUserStatusId** | **string**| key: deviceComplianceUserStatus-id of deviceComplianceUserStatus | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetUserStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceUserStatus**](microsoft.graph.deviceComplianceUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListUserStatuses

> CollectionOfDeviceComplianceUserStatus DeviceManagementDeviceCompliancePoliciesListUserStatuses(ctx, deviceCompliancePolicyId, optional)
Get userStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListUserStatusesOpts struct


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

[**CollectionOfDeviceComplianceUserStatus**](Collection of deviceComplianceUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses

> DeviceManagementDeviceCompliancePoliciesUpdateUserStatuses(ctx, deviceCompliancePolicyId, deviceComplianceUserStatusId, microsoftGraphDeviceComplianceUserStatus)
Update the navigation property userStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceUserStatusId** | **string**| key: deviceComplianceUserStatus-id of deviceComplianceUserStatus | 
**microsoftGraphDeviceComplianceUserStatus** | [**MicrosoftGraphDeviceComplianceUserStatus**](MicrosoftGraphDeviceComplianceUserStatus.md)| New navigation property values | 

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

