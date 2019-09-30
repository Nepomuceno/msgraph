# \DeviceAppManagementTargetedManagedAppConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementCreateTargetedManagedAppConfigurations) | **Post** /deviceAppManagement/targetedManagedAppConfigurations | Create new navigation property to targetedManagedAppConfigurations for deviceAppManagement
[**DeviceAppManagementGetTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementGetTargetedManagedAppConfigurations) | **Get** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id}) | Get targetedManagedAppConfigurations from deviceAppManagement
[**DeviceAppManagementListTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementListTargetedManagedAppConfigurations) | **Get** /deviceAppManagement/targetedManagedAppConfigurations | Get targetedManagedAppConfigurations from deviceAppManagement
[**DeviceAppManagementUpdateTargetedManagedAppConfigurations**](DeviceAppManagementTargetedManagedAppConfigurationApi.md#DeviceAppManagementUpdateTargetedManagedAppConfigurations) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id}) | Update the navigation property targetedManagedAppConfigurations in deviceAppManagement



## DeviceAppManagementCreateTargetedManagedAppConfigurations

> MicrosoftGraphTargetedManagedAppConfiguration DeviceAppManagementCreateTargetedManagedAppConfigurations(ctx, microsoftGraphTargetedManagedAppConfiguration)
Create new navigation property to targetedManagedAppConfigurations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTargetedManagedAppConfiguration** | [**MicrosoftGraphTargetedManagedAppConfiguration**](MicrosoftGraphTargetedManagedAppConfiguration.md)| New navigation property | 

### Return type

[**MicrosoftGraphTargetedManagedAppConfiguration**](microsoft.graph.targetedManagedAppConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetTargetedManagedAppConfigurations

> MicrosoftGraphTargetedManagedAppConfiguration DeviceAppManagementGetTargetedManagedAppConfigurations(ctx, targetedManagedAppConfigurationId, optional)
Get targetedManagedAppConfigurations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementGetTargetedManagedAppConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetTargetedManagedAppConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTargetedManagedAppConfiguration**](microsoft.graph.targetedManagedAppConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListTargetedManagedAppConfigurations

> CollectionOfTargetedManagedAppConfiguration DeviceAppManagementListTargetedManagedAppConfigurations(ctx, optional)
Get targetedManagedAppConfigurations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListTargetedManagedAppConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListTargetedManagedAppConfigurationsOpts struct


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

[**CollectionOfTargetedManagedAppConfiguration**](Collection of targetedManagedAppConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateTargetedManagedAppConfigurations

> DeviceAppManagementUpdateTargetedManagedAppConfigurations(ctx, targetedManagedAppConfigurationId, microsoftGraphTargetedManagedAppConfiguration)
Update the navigation property targetedManagedAppConfigurations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**microsoftGraphTargetedManagedAppConfiguration** | [**MicrosoftGraphTargetedManagedAppConfiguration**](MicrosoftGraphTargetedManagedAppConfiguration.md)| New navigation property values | 

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

