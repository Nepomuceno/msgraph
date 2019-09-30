# \DeviceManagementDeviceEnrollmentConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementCreateDeviceEnrollmentConfigurations) | **Post** /deviceManagement/deviceEnrollmentConfigurations | Create new navigation property to deviceEnrollmentConfigurations for deviceManagement
[**DeviceManagementGetDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementGetDeviceEnrollmentConfigurations) | **Get** /deviceManagement/deviceEnrollmentConfigurations({deviceEnrollmentConfiguration-id}) | Get deviceEnrollmentConfigurations from deviceManagement
[**DeviceManagementListDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementListDeviceEnrollmentConfigurations) | **Get** /deviceManagement/deviceEnrollmentConfigurations | Get deviceEnrollmentConfigurations from deviceManagement
[**DeviceManagementUpdateDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementUpdateDeviceEnrollmentConfigurations) | **Patch** /deviceManagement/deviceEnrollmentConfigurations({deviceEnrollmentConfiguration-id}) | Update the navigation property deviceEnrollmentConfigurations in deviceManagement



## DeviceManagementCreateDeviceEnrollmentConfigurations

> MicrosoftGraphDeviceEnrollmentConfiguration DeviceManagementCreateDeviceEnrollmentConfigurations(ctx, microsoftGraphDeviceEnrollmentConfiguration)
Create new navigation property to deviceEnrollmentConfigurations for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDeviceEnrollmentConfiguration** | [**MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceEnrollmentConfiguration**](microsoft.graph.deviceEnrollmentConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetDeviceEnrollmentConfigurations

> MicrosoftGraphDeviceEnrollmentConfiguration DeviceManagementGetDeviceEnrollmentConfigurations(ctx, deviceEnrollmentConfigurationId, optional)
Get deviceEnrollmentConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string**| key: deviceEnrollmentConfiguration-id of deviceEnrollmentConfiguration | 
 **optional** | ***DeviceManagementGetDeviceEnrollmentConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetDeviceEnrollmentConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceEnrollmentConfiguration**](microsoft.graph.deviceEnrollmentConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceEnrollmentConfigurations

> CollectionOfDeviceEnrollmentConfiguration DeviceManagementListDeviceEnrollmentConfigurations(ctx, optional)
Get deviceEnrollmentConfigurations from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListDeviceEnrollmentConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListDeviceEnrollmentConfigurationsOpts struct


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

[**CollectionOfDeviceEnrollmentConfiguration**](Collection of deviceEnrollmentConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceEnrollmentConfigurations

> DeviceManagementUpdateDeviceEnrollmentConfigurations(ctx, deviceEnrollmentConfigurationId, microsoftGraphDeviceEnrollmentConfiguration)
Update the navigation property deviceEnrollmentConfigurations in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string**| key: deviceEnrollmentConfiguration-id of deviceEnrollmentConfiguration | 
**microsoftGraphDeviceEnrollmentConfiguration** | [**MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md)| New navigation property values | 

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

