# \DeviceManagementDeviceConfigurationsDeviceConfigurationAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceConfigurationsCreateAssignments**](DeviceManagementDeviceConfigurationsDeviceConfigurationAssignmentApi.md#DeviceManagementDeviceConfigurationsCreateAssignments) | **Post** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementDeviceConfigurationsGetAssignments**](DeviceManagementDeviceConfigurationsDeviceConfigurationAssignmentApi.md#DeviceManagementDeviceConfigurationsGetAssignments) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/assignments({deviceConfigurationAssignment-id}) | Get assignments from deviceManagement
[**DeviceManagementDeviceConfigurationsListAssignments**](DeviceManagementDeviceConfigurationsDeviceConfigurationAssignmentApi.md#DeviceManagementDeviceConfigurationsListAssignments) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/assignments | Get assignments from deviceManagement
[**DeviceManagementDeviceConfigurationsUpdateAssignments**](DeviceManagementDeviceConfigurationsDeviceConfigurationAssignmentApi.md#DeviceManagementDeviceConfigurationsUpdateAssignments) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/assignments({deviceConfigurationAssignment-id}) | Update the navigation property assignments in deviceManagement



## DeviceManagementDeviceConfigurationsCreateAssignments

> MicrosoftGraphDeviceConfigurationAssignment DeviceManagementDeviceConfigurationsCreateAssignments(ctx, deviceConfigurationId, microsoftGraphDeviceConfigurationAssignment)
Create new navigation property to assignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphDeviceConfigurationAssignment** | [**MicrosoftGraphDeviceConfigurationAssignment**](MicrosoftGraphDeviceConfigurationAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfigurationAssignment**](microsoft.graph.deviceConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsGetAssignments

> MicrosoftGraphDeviceConfigurationAssignment DeviceManagementDeviceConfigurationsGetAssignments(ctx, deviceConfigurationId, deviceConfigurationAssignmentId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**deviceConfigurationAssignmentId** | **string**| key: deviceConfigurationAssignment-id of deviceConfigurationAssignment | 
 **optional** | ***DeviceManagementDeviceConfigurationsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationAssignment**](microsoft.graph.deviceConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsListAssignments

> CollectionOfDeviceConfigurationAssignment DeviceManagementDeviceConfigurationsListAssignments(ctx, deviceConfigurationId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementDeviceConfigurationsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsListAssignmentsOpts struct


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

[**CollectionOfDeviceConfigurationAssignment**](Collection of deviceConfigurationAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsUpdateAssignments

> DeviceManagementDeviceConfigurationsUpdateAssignments(ctx, deviceConfigurationId, deviceConfigurationAssignmentId, microsoftGraphDeviceConfigurationAssignment)
Update the navigation property assignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**deviceConfigurationAssignmentId** | **string**| key: deviceConfigurationAssignment-id of deviceConfigurationAssignment | 
**microsoftGraphDeviceConfigurationAssignment** | [**MicrosoftGraphDeviceConfigurationAssignment**](MicrosoftGraphDeviceConfigurationAssignment.md)| New navigation property values | 

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

