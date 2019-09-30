# \DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppConfigurationsCreateAssignments**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationAssignmentApi.md#DeviceAppManagementMobileAppConfigurationsCreateAssignments) | **Post** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsGetAssignments**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationAssignmentApi.md#DeviceAppManagementMobileAppConfigurationsGetAssignments) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/assignments({managedDeviceMobileAppConfigurationAssignment-id}) | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsListAssignments**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationAssignmentApi.md#DeviceAppManagementMobileAppConfigurationsListAssignments) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsUpdateAssignments**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationAssignmentApi.md#DeviceAppManagementMobileAppConfigurationsUpdateAssignments) | **Patch** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/assignments({managedDeviceMobileAppConfigurationAssignment-id}) | Update the navigation property assignments in deviceAppManagement



## DeviceAppManagementMobileAppConfigurationsCreateAssignments

> MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment DeviceAppManagementMobileAppConfigurationsCreateAssignments(ctx, managedDeviceMobileAppConfigurationId, microsoftGraphManagedDeviceMobileAppConfigurationAssignment)
Create new navigation property to assignments for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**microsoftGraphManagedDeviceMobileAppConfigurationAssignment** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment**](MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment**](microsoft.graph.managedDeviceMobileAppConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsGetAssignments

> MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment DeviceAppManagementMobileAppConfigurationsGetAssignments(ctx, managedDeviceMobileAppConfigurationId, managedDeviceMobileAppConfigurationAssignmentId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**managedDeviceMobileAppConfigurationAssignmentId** | **string**| key: managedDeviceMobileAppConfigurationAssignment-id of managedDeviceMobileAppConfigurationAssignment | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment**](microsoft.graph.managedDeviceMobileAppConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsListAssignments

> CollectionOfManagedDeviceMobileAppConfigurationAssignment DeviceAppManagementMobileAppConfigurationsListAssignments(ctx, managedDeviceMobileAppConfigurationId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsListAssignmentsOpts struct


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

[**CollectionOfManagedDeviceMobileAppConfigurationAssignment**](Collection of managedDeviceMobileAppConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsUpdateAssignments

> DeviceAppManagementMobileAppConfigurationsUpdateAssignments(ctx, managedDeviceMobileAppConfigurationId, managedDeviceMobileAppConfigurationAssignmentId, microsoftGraphManagedDeviceMobileAppConfigurationAssignment)
Update the navigation property assignments in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**managedDeviceMobileAppConfigurationAssignmentId** | **string**| key: managedDeviceMobileAppConfigurationAssignment-id of managedDeviceMobileAppConfigurationAssignment | 
**microsoftGraphManagedDeviceMobileAppConfigurationAssignment** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment**](MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment.md)| New navigation property values | 

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

