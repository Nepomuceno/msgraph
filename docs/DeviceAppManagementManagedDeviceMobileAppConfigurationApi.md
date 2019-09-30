# \DeviceAppManagementManagedDeviceMobileAppConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMobileAppConfigurations**](DeviceAppManagementManagedDeviceMobileAppConfigurationApi.md#DeviceAppManagementCreateMobileAppConfigurations) | **Post** /deviceAppManagement/mobileAppConfigurations | Create new navigation property to mobileAppConfigurations for deviceAppManagement
[**DeviceAppManagementGetMobileAppConfigurations**](DeviceAppManagementManagedDeviceMobileAppConfigurationApi.md#DeviceAppManagementGetMobileAppConfigurations) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id}) | Get mobileAppConfigurations from deviceAppManagement
[**DeviceAppManagementListMobileAppConfigurations**](DeviceAppManagementManagedDeviceMobileAppConfigurationApi.md#DeviceAppManagementListMobileAppConfigurations) | **Get** /deviceAppManagement/mobileAppConfigurations | Get mobileAppConfigurations from deviceAppManagement
[**DeviceAppManagementUpdateMobileAppConfigurations**](DeviceAppManagementManagedDeviceMobileAppConfigurationApi.md#DeviceAppManagementUpdateMobileAppConfigurations) | **Patch** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id}) | Update the navigation property mobileAppConfigurations in deviceAppManagement



## DeviceAppManagementCreateMobileAppConfigurations

> MicrosoftGraphManagedDeviceMobileAppConfiguration DeviceAppManagementCreateMobileAppConfigurations(ctx, microsoftGraphManagedDeviceMobileAppConfiguration)
Create new navigation property to mobileAppConfigurations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedDeviceMobileAppConfiguration** | [**MicrosoftGraphManagedDeviceMobileAppConfiguration**](MicrosoftGraphManagedDeviceMobileAppConfiguration.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfiguration**](microsoft.graph.managedDeviceMobileAppConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetMobileAppConfigurations

> MicrosoftGraphManagedDeviceMobileAppConfiguration DeviceAppManagementGetMobileAppConfigurations(ctx, managedDeviceMobileAppConfigurationId, optional)
Get mobileAppConfigurations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
 **optional** | ***DeviceAppManagementGetMobileAppConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetMobileAppConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfiguration**](microsoft.graph.managedDeviceMobileAppConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMobileAppConfigurations

> CollectionOfManagedDeviceMobileAppConfiguration DeviceAppManagementListMobileAppConfigurations(ctx, optional)
Get mobileAppConfigurations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListMobileAppConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListMobileAppConfigurationsOpts struct


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

[**CollectionOfManagedDeviceMobileAppConfiguration**](Collection of managedDeviceMobileAppConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateMobileAppConfigurations

> DeviceAppManagementUpdateMobileAppConfigurations(ctx, managedDeviceMobileAppConfigurationId, microsoftGraphManagedDeviceMobileAppConfiguration)
Update the navigation property mobileAppConfigurations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**microsoftGraphManagedDeviceMobileAppConfiguration** | [**MicrosoftGraphManagedDeviceMobileAppConfiguration**](MicrosoftGraphManagedDeviceMobileAppConfiguration.md)| New navigation property values | 

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

