# \DeviceManagementDeviceEnrollmentConfigurationsEnrollmentConfigurationAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments**](DeviceManagementDeviceEnrollmentConfigurationsEnrollmentConfigurationAssignmentApi.md#DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments) | **Post** /deviceManagement/deviceEnrollmentConfigurations({deviceEnrollmentConfiguration-id})/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsGetAssignments**](DeviceManagementDeviceEnrollmentConfigurationsEnrollmentConfigurationAssignmentApi.md#DeviceManagementDeviceEnrollmentConfigurationsGetAssignments) | **Get** /deviceManagement/deviceEnrollmentConfigurations({deviceEnrollmentConfiguration-id})/assignments({enrollmentConfigurationAssignment-id}) | Get assignments from deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsListAssignments**](DeviceManagementDeviceEnrollmentConfigurationsEnrollmentConfigurationAssignmentApi.md#DeviceManagementDeviceEnrollmentConfigurationsListAssignments) | **Get** /deviceManagement/deviceEnrollmentConfigurations({deviceEnrollmentConfiguration-id})/assignments | Get assignments from deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments**](DeviceManagementDeviceEnrollmentConfigurationsEnrollmentConfigurationAssignmentApi.md#DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments) | **Patch** /deviceManagement/deviceEnrollmentConfigurations({deviceEnrollmentConfiguration-id})/assignments({enrollmentConfigurationAssignment-id}) | Update the navigation property assignments in deviceManagement



## DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments

> MicrosoftGraphEnrollmentConfigurationAssignment DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments(ctx, deviceEnrollmentConfigurationId, microsoftGraphEnrollmentConfigurationAssignment)
Create new navigation property to assignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string**| key: deviceEnrollmentConfiguration-id of deviceEnrollmentConfiguration | 
**microsoftGraphEnrollmentConfigurationAssignment** | [**MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphEnrollmentConfigurationAssignment**](microsoft.graph.enrollmentConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceEnrollmentConfigurationsGetAssignments

> MicrosoftGraphEnrollmentConfigurationAssignment DeviceManagementDeviceEnrollmentConfigurationsGetAssignments(ctx, deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string**| key: deviceEnrollmentConfiguration-id of deviceEnrollmentConfiguration | 
**enrollmentConfigurationAssignmentId** | **string**| key: enrollmentConfigurationAssignment-id of enrollmentConfigurationAssignment | 
 **optional** | ***DeviceManagementDeviceEnrollmentConfigurationsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceEnrollmentConfigurationsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEnrollmentConfigurationAssignment**](microsoft.graph.enrollmentConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceEnrollmentConfigurationsListAssignments

> CollectionOfEnrollmentConfigurationAssignment DeviceManagementDeviceEnrollmentConfigurationsListAssignments(ctx, deviceEnrollmentConfigurationId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string**| key: deviceEnrollmentConfiguration-id of deviceEnrollmentConfiguration | 
 **optional** | ***DeviceManagementDeviceEnrollmentConfigurationsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceEnrollmentConfigurationsListAssignmentsOpts struct


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

[**CollectionOfEnrollmentConfigurationAssignment**](Collection of enrollmentConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments

> DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments(ctx, deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId, microsoftGraphEnrollmentConfigurationAssignment)
Update the navigation property assignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string**| key: deviceEnrollmentConfiguration-id of deviceEnrollmentConfiguration | 
**enrollmentConfigurationAssignmentId** | **string**| key: enrollmentConfigurationAssignment-id of enrollmentConfigurationAssignment | 
**microsoftGraphEnrollmentConfigurationAssignment** | [**MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md)| New navigation property values | 

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

