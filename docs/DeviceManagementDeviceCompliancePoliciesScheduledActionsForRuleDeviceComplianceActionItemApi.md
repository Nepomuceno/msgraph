# \DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleDeviceComplianceActionItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations**](DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleDeviceComplianceActionItemApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations) | **Post** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule({deviceComplianceScheduledActionForRule-id})/scheduledActionConfigurations | Create new navigation property to scheduledActionConfigurations for deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations**](DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleDeviceComplianceActionItemApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule({deviceComplianceScheduledActionForRule-id})/scheduledActionConfigurations({deviceComplianceActionItem-id}) | Get scheduledActionConfigurations from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations**](DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleDeviceComplianceActionItemApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations) | **Get** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule({deviceComplianceScheduledActionForRule-id})/scheduledActionConfigurations | Get scheduledActionConfigurations from deviceManagement
[**DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations**](DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleDeviceComplianceActionItemApi.md#DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations) | **Patch** /deviceManagement/deviceCompliancePolicies({deviceCompliancePolicy-id})/scheduledActionsForRule({deviceComplianceScheduledActionForRule-id})/scheduledActionConfigurations({deviceComplianceActionItem-id}) | Update the navigation property scheduledActionConfigurations in deviceManagement



## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations

> MicrosoftGraphDeviceComplianceActionItem DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleCreateScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, microsoftGraphDeviceComplianceActionItem)
Create new navigation property to scheduledActionConfigurations for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**microsoftGraphDeviceComplianceActionItem** | [**MicrosoftGraphDeviceComplianceActionItem**](MicrosoftGraphDeviceComplianceActionItem.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceActionItem**](microsoft.graph.deviceComplianceActionItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations

> MicrosoftGraphDeviceComplianceActionItem DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, deviceComplianceActionItemId, optional)
Get scheduledActionConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**deviceComplianceActionItemId** | **string**| key: deviceComplianceActionItem-id of deviceComplianceActionItem | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleGetScheduledActionConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceActionItem**](microsoft.graph.deviceComplianceActionItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations

> CollectionOfDeviceComplianceActionItem DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, optional)
Get scheduledActionConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
 **optional** | ***DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleListScheduledActionConfigurationsOpts struct


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

[**CollectionOfDeviceComplianceActionItem**](Collection of deviceComplianceActionItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations

> DeviceManagementDeviceCompliancePoliciesScheduledActionsForRuleUpdateScheduledActionConfigurations(ctx, deviceCompliancePolicyId, deviceComplianceScheduledActionForRuleId, deviceComplianceActionItemId, microsoftGraphDeviceComplianceActionItem)
Update the navigation property scheduledActionConfigurations in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicyId** | **string**| key: deviceCompliancePolicy-id of deviceCompliancePolicy | 
**deviceComplianceScheduledActionForRuleId** | **string**| key: deviceComplianceScheduledActionForRule-id of deviceComplianceScheduledActionForRule | 
**deviceComplianceActionItemId** | **string**| key: deviceComplianceActionItem-id of deviceComplianceActionItem | 
**microsoftGraphDeviceComplianceActionItem** | [**MicrosoftGraphDeviceComplianceActionItem**](MicrosoftGraphDeviceComplianceActionItem.md)| New navigation property values | 

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

