# \DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppConfigurationsCreateDeviceStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceStatusApi.md#DeviceAppManagementMobileAppConfigurationsCreateDeviceStatuses) | **Post** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/deviceStatuses | Create new navigation property to deviceStatuses for deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsGetDeviceStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceStatusApi.md#DeviceAppManagementMobileAppConfigurationsGetDeviceStatuses) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/deviceStatuses({managedDeviceMobileAppConfigurationDeviceStatus-id}) | Get deviceStatuses from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsListDeviceStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceStatusApi.md#DeviceAppManagementMobileAppConfigurationsListDeviceStatuses) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/deviceStatuses | Get deviceStatuses from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationDeviceStatusApi.md#DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatuses) | **Patch** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/deviceStatuses({managedDeviceMobileAppConfigurationDeviceStatus-id}) | Update the navigation property deviceStatuses in deviceAppManagement



## DeviceAppManagementMobileAppConfigurationsCreateDeviceStatuses

> MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus DeviceAppManagementMobileAppConfigurationsCreateDeviceStatuses(ctx, managedDeviceMobileAppConfigurationId, microsoftGraphManagedDeviceMobileAppConfigurationDeviceStatus)
Create new navigation property to deviceStatuses for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**microsoftGraphManagedDeviceMobileAppConfigurationDeviceStatus** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus**](MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus**](microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsGetDeviceStatuses

> MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus DeviceAppManagementMobileAppConfigurationsGetDeviceStatuses(ctx, managedDeviceMobileAppConfigurationId, managedDeviceMobileAppConfigurationDeviceStatusId, optional)
Get deviceStatuses from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**managedDeviceMobileAppConfigurationDeviceStatusId** | **string**| key: managedDeviceMobileAppConfigurationDeviceStatus-id of managedDeviceMobileAppConfigurationDeviceStatus | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsGetDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsGetDeviceStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus**](microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsListDeviceStatuses

> CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus DeviceAppManagementMobileAppConfigurationsListDeviceStatuses(ctx, managedDeviceMobileAppConfigurationId, optional)
Get deviceStatuses from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsListDeviceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsListDeviceStatusesOpts struct


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

[**CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus**](Collection of managedDeviceMobileAppConfigurationDeviceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatuses

> DeviceAppManagementMobileAppConfigurationsUpdateDeviceStatuses(ctx, managedDeviceMobileAppConfigurationId, managedDeviceMobileAppConfigurationDeviceStatusId, microsoftGraphManagedDeviceMobileAppConfigurationDeviceStatus)
Update the navigation property deviceStatuses in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**managedDeviceMobileAppConfigurationDeviceStatusId** | **string**| key: managedDeviceMobileAppConfigurationDeviceStatus-id of managedDeviceMobileAppConfigurationDeviceStatus | 
**microsoftGraphManagedDeviceMobileAppConfigurationDeviceStatus** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus**](MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus.md)| New navigation property values | 

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

