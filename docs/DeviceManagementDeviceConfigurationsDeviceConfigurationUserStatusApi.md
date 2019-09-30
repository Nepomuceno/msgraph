# \DeviceManagementDeviceConfigurationsDeviceConfigurationUserStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceConfigurationsCreateUserStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationUserStatusApi.md#DeviceManagementDeviceConfigurationsCreateUserStatuses) | **Post** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/userStatuses | Create new navigation property to userStatuses for deviceManagement
[**DeviceManagementDeviceConfigurationsGetUserStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationUserStatusApi.md#DeviceManagementDeviceConfigurationsGetUserStatuses) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/userStatuses({deviceConfigurationUserStatus-id}) | Get userStatuses from deviceManagement
[**DeviceManagementDeviceConfigurationsListUserStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationUserStatusApi.md#DeviceManagementDeviceConfigurationsListUserStatuses) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/userStatuses | Get userStatuses from deviceManagement
[**DeviceManagementDeviceConfigurationsUpdateUserStatuses**](DeviceManagementDeviceConfigurationsDeviceConfigurationUserStatusApi.md#DeviceManagementDeviceConfigurationsUpdateUserStatuses) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/userStatuses({deviceConfigurationUserStatus-id}) | Update the navigation property userStatuses in deviceManagement



## DeviceManagementDeviceConfigurationsCreateUserStatuses

> MicrosoftGraphDeviceConfigurationUserStatus DeviceManagementDeviceConfigurationsCreateUserStatuses(ctx, deviceConfigurationId, microsoftGraphDeviceConfigurationUserStatus)
Create new navigation property to userStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphDeviceConfigurationUserStatus** | [**MicrosoftGraphDeviceConfigurationUserStatus**](MicrosoftGraphDeviceConfigurationUserStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfigurationUserStatus**](microsoft.graph.deviceConfigurationUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsGetUserStatuses

> MicrosoftGraphDeviceConfigurationUserStatus DeviceManagementDeviceConfigurationsGetUserStatuses(ctx, deviceConfigurationId, deviceConfigurationUserStatusId, optional)
Get userStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**deviceConfigurationUserStatusId** | **string**| key: deviceConfigurationUserStatus-id of deviceConfigurationUserStatus | 
 **optional** | ***DeviceManagementDeviceConfigurationsGetUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsGetUserStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationUserStatus**](microsoft.graph.deviceConfigurationUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsListUserStatuses

> CollectionOfDeviceConfigurationUserStatus DeviceManagementDeviceConfigurationsListUserStatuses(ctx, deviceConfigurationId, optional)
Get userStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementDeviceConfigurationsListUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsListUserStatusesOpts struct


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

[**CollectionOfDeviceConfigurationUserStatus**](Collection of deviceConfigurationUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsUpdateUserStatuses

> DeviceManagementDeviceConfigurationsUpdateUserStatuses(ctx, deviceConfigurationId, deviceConfigurationUserStatusId, microsoftGraphDeviceConfigurationUserStatus)
Update the navigation property userStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**deviceConfigurationUserStatusId** | **string**| key: deviceConfigurationUserStatus-id of deviceConfigurationUserStatus | 
**microsoftGraphDeviceConfigurationUserStatus** | [**MicrosoftGraphDeviceConfigurationUserStatus**](MicrosoftGraphDeviceConfigurationUserStatus.md)| New navigation property values | 

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

