# \DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceConfigurationsCreateDeviceStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceStatusApi.md#DeviceManagementDeviceConfigurationsCreateDeviceStatuses) | **Post** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceStatuses | Create new navigation property to deviceStatuses for deviceManagement
[**DeviceManagementDeviceConfigurationsGetDeviceStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceStatusApi.md#DeviceManagementDeviceConfigurationsGetDeviceStatuses) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceStatuses({deviceConfigurationDeviceStatus-id}) | Get deviceStatuses from deviceManagement
[**DeviceManagementDeviceConfigurationsListDeviceStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceStatusApi.md#DeviceManagementDeviceConfigurationsListDeviceStatuses) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceStatuses | Get deviceStatuses from deviceManagement
[**DeviceManagementDeviceConfigurationsUpdateDeviceStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationDeviceStatusApi.md#DeviceManagementDeviceConfigurationsUpdateDeviceStatuses) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceStatuses({deviceConfigurationDeviceStatus-id}) | Update the navigation property deviceStatuses in deviceManagement



## DeviceManagementDeviceConfigurationsCreateDeviceStatuses

> MicrosoftGraphDeviceConfigurationDeviceStatus DeviceManagementDeviceConfigurationsCreateDeviceStatuses(ctx, deviceConfigurationId, microsoftGraphDeviceConfigurationDeviceStatus)
Create new navigation property to deviceStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphDeviceConfigurationDeviceStatus** | [**MicrosoftGraphDeviceConfigurationDeviceStatus**](MicrosoftGraphDeviceConfigurationDeviceStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfigurationDeviceStatus**](microsoft.graph.deviceConfigurationDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsGetDeviceStatuses

> MicrosoftGraphDeviceConfigurationDeviceStatus DeviceManagementDeviceConfigurationsGetDeviceStatuses(ctx, deviceConfigurationId, deviceConfigurationDeviceStatusId, optional)
Get deviceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**deviceConfigurationDeviceStatusId** | **string**| key: deviceConfigurationDeviceStatus-id of deviceConfigurationDeviceStatus | 
 **optional** | ***DeviceManagementDeviceConfigurationsGetDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsGetDeviceStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationDeviceStatus**](microsoft.graph.deviceConfigurationDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsListDeviceStatuses

> CollectionOfDeviceConfigurationDeviceStatus DeviceManagementDeviceConfigurationsListDeviceStatuses(ctx, deviceConfigurationId, optional)
Get deviceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementDeviceConfigurationsListDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsListDeviceStatusesOpts struct


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

[**CollectionOfDeviceConfigurationDeviceStatus**](Collection of deviceConfigurationDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsUpdateDeviceStatuses

> DeviceManagementDeviceConfigurationsUpdateDeviceStatuses(ctx, deviceConfigurationId, deviceConfigurationDeviceStatusId, microsoftGraphDeviceConfigurationDeviceStatus)
Update the navigation property deviceStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**deviceConfigurationDeviceStatusId** | **string**| key: deviceConfigurationDeviceStatus-id of deviceConfigurationDeviceStatus | 
**microsoftGraphDeviceConfigurationDeviceStatus** | [**MicrosoftGraphDeviceConfigurationDeviceStatus**](MicrosoftGraphDeviceConfigurationDeviceStatus.md)| New navigation property values | 

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

