# \DeviceManagementDeviceCompliancePoliciesDeviceComplianceScheduledActionForRuleApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceScheduledActionForRuleApi.md#DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule) | **Post** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule | Create new navigation property to scheduledActionsForRule for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceScheduledActionForRuleApi.md#DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule({deviceComplianceScheduledActionForRule-id}) | Get scheduledActionsForRule from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceScheduledActionForRuleApi.md#DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule | Get scheduledActionsForRule from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule**](DeviceManagementDeviceCompliancePoliciesDeviceComplianceScheduledActionForRuleApi.md#DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule({deviceComplianceScheduledActionForRule-id}) | Update the navigation property scheduledActionsForRule in deviceManagement



## DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule

> MicrosoftGraphDeviceComplianceScheduledActionForRule DeviceManagementDeviceCompliancePoliciesCreateScheduledActionsForRule(ctx, deviceCompliancePolicyId, microsoftGraphDeviceComplianceScheduledActionForRule)
Create new navigation property to scheduledActionsForRule for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**microsoftGraphDeviceComplianceScheduledActionForRule** | [**MicrosoftGraphDeviceComplianceScheduledActionForRule**](MicrosoftGraphDeviceComplianceScheduledActionForRule.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceScheduledActionForRule**](microsoft.graph.deviceComplianceScheduledActionForRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule

> MicrosoftGraphDeviceComplianceScheduledActionForRule DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRule(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, optional)
Get scheduledActionsForRule from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesGetScheduledActionsForRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceScheduledActionForRule**](microsoft.graph.deviceComplianceScheduledActionForRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule

> CollectionOfDeviceComplianceScheduledActionForRule DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRule(ctx, deviceCompliancePolicyId, optional)
Get scheduledActionsForRule from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesListScheduledActionsForRuleOpts struct


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

[**CollectionOfDeviceComplianceScheduledActionForRule**](Collection of deviceComplianceScheduledActionForRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule

> DeviceManagementDeviceCompliancePoliciesUpdateScheduledActionsForRule(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, microsoftGraphDeviceComplianceScheduledActionForRule)
Update the navigation property scheduledActionsForRule in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**microsoftGraphDeviceComplianceScheduledActionForRule** | [**MicrosoftGraphDeviceComplianceScheduledActionForRule**](MicrosoftGraphDeviceComplianceScheduledActionForRule.md)| New navigation property values | 

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

