# \DeviceManagementDeviceCompliancePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementCreateDeviceCompliancePolicies) | **Post** /deviceManagement/deviceCompliancePolicies | Create new navigation property to deviceCompliancePolicies for deviceManagement
[**DeviceManagementGetDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementGetDeviceCompliancePolicies) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id}) | Get deviceCompliancePolicies from deviceManagement
[**DeviceManagementListDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementListDeviceCompliancePolicies) | **Get** /deviceManagement/deviceCompliancePolicies | Get deviceCompliancePolicies from deviceManagement
[**DeviceManagementUpdateDeviceCompliancePolicies**](DeviceManagementDeviceCompliancePolicyApi.md#DeviceManagementUpdateDeviceCompliancePolicies) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id}) | Update the navigation property deviceCompliancePolicies in deviceManagement



## DeviceManagementCreateDeviceCompliancePolicies

> MicrosoftGraphDeviceCompliancePolicy DeviceManagementCreateDeviceCompliancePolicies(ctx, microsoftGraphDeviceCompliancePolicy)
Create new navigation property to deviceCompliancePolicies for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceCompliancePolicy** | [**MicrosoftGraphDeviceCompliancePolicy**](MicrosoftGraphDeviceCompliancePolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicy**](microsoft.graph.deviceCompliancePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetDeviceCompliancePolicies

> MicrosoftGraphDeviceCompliancePolicy DeviceManagementGetDeviceCompliancePolicies(ctx, deviceCompliancePolicyId, optional)
Get deviceCompliancePolicies from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementGetDeviceCompliancePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceCompliancePoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicy**](microsoft.graph.deviceCompliancePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceCompliancePolicies

> CollectionOfDeviceCompliancePolicy DeviceManagementListDeviceCompliancePolicies(ctx, optional)
Get deviceCompliancePolicies from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceCompliancePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceCompliancePoliciesOpts struct


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

[**CollectionOfDeviceCompliancePolicy**](Collection of deviceCompliancePolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCompliancePolicies

> DeviceManagementUpdateDeviceCompliancePolicies(ctx, deviceCompliancePolicyId, microsoftGraphDeviceCompliancePolicy)
Update the navigation property deviceCompliancePolicies in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceCompliancePolicy** | [**MicrosoftGraphDeviceCompliancePolicy**](MicrosoftGraphDeviceCompliancePolicy.md)| New navigation property values | 

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

