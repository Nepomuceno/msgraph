# \DeviceManagementDeviceConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceConfigurations**](DeviceManagementDeviceConfigurationApi.md#DeviceManagementCreateDeviceConfigurations) | **Post** /deviceManagement/deviceConfigurations | Create new navigation property to deviceConfigurations for deviceManagement
[**DeviceManagementGetDeviceConfigurations**](DeviceManagementDeviceConfigurationApi.md#DeviceManagementGetDeviceConfigurations) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id}) | Get deviceConfigurations from deviceManagement
[**DeviceManagementListDeviceConfigurations**](DeviceManagementDeviceConfigurationApi.md#DeviceManagementListDeviceConfigurations) | **Get** /deviceManagement/deviceConfigurations | Get deviceConfigurations from deviceManagement
[**DeviceManagementUpdateDeviceConfigurations**](DeviceManagementDeviceConfigurationApi.md#DeviceManagementUpdateDeviceConfigurations) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id}) | Update the navigation property deviceConfigurations in deviceManagement



## DeviceManagementCreateDeviceConfigurations

> MicrosoftGraphDeviceConfiguration DeviceManagementCreateDeviceConfigurations(ctx, microsoftGraphDeviceConfiguration)
Create new navigation property to deviceConfigurations for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceConfiguration** | [**MicrosoftGraphDeviceConfiguration**](MicrosoftGraphDeviceConfiguration.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfiguration**](microsoft.graph.deviceConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetDeviceConfigurations

> MicrosoftGraphDeviceConfiguration DeviceManagementGetDeviceConfigurations(ctx, deviceConfigurationId, optional)
Get deviceConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementGetDeviceConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfiguration**](microsoft.graph.deviceConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceConfigurations

> CollectionOfDeviceConfiguration DeviceManagementListDeviceConfigurations(ctx, optional)
Get deviceConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceConfigurationsOpts struct


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

[**CollectionOfDeviceConfiguration**](Collection of deviceConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceConfigurations

> DeviceManagementUpdateDeviceConfigurations(ctx, deviceConfigurationId, microsoftGraphDeviceConfiguration)
Update the navigation property deviceConfigurations in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphDeviceConfiguration** | [**MicrosoftGraphDeviceConfiguration**](MicrosoftGraphDeviceConfiguration.md)| New navigation property values | 

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

