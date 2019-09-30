# \DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppPolicyAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments**](DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppPolicyAssignmentApi.md#DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments) | **Post** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments**](DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppPolicyAssignmentApi.md#DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments) | **Get** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/assignments({targetedManagedAppPolicyAssignment-id}) | Get assignments from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsListAssignments**](DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppPolicyAssignmentApi.md#DeviceAppManagementTargetedManagedAppConfigurationsListAssignments) | **Get** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments**](DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppPolicyAssignmentApi.md#DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/assignments({targetedManagedAppPolicyAssignment-id}) | Update the navigation property assignments in deviceAppManagement



## DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments

> MicrosoftGraphTargetedManagedAppPolicyAssignment DeviceAppManagementTargetedManagedAppConfigurationsCreateAssignments(ctx, targetedManagedAppConfigurationId, microsoftGraphTargetedManagedAppPolicyAssignment)
Create new navigation property to assignments for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**microsoftGraphTargetedManagedAppPolicyAssignment** | [**MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments

> MicrosoftGraphTargetedManagedAppPolicyAssignment DeviceAppManagementTargetedManagedAppConfigurationsGetAssignments(ctx, targetedManagedAppConfigurationId, targetedManagedAppPolicyAssignmentId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**targetedManagedAppPolicyAssignmentId** | **string**| key: targetedManagedAppPolicyAssignment-id of targetedManagedAppPolicyAssignment | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsListAssignments

> CollectionOfTargetedManagedAppPolicyAssignment DeviceAppManagementTargetedManagedAppConfigurationsListAssignments(ctx, targetedManagedAppConfigurationId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsListAssignmentsOpts struct


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

[**CollectionOfTargetedManagedAppPolicyAssignment**](Collection of targetedManagedAppPolicyAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments

> DeviceAppManagementTargetedManagedAppConfigurationsUpdateAssignments(ctx, targetedManagedAppConfigurationId, targetedManagedAppPolicyAssignmentId, microsoftGraphTargetedManagedAppPolicyAssignment)
Update the navigation property assignments in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**targetedManagedAppPolicyAssignmentId** | **string**| key: targetedManagedAppPolicyAssignment-id of targetedManagedAppPolicyAssignment | 
**microsoftGraphTargetedManagedAppPolicyAssignment** | [**MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md)| New navigation property values | 

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

