# \DeviceAppManagementTargetedManagedAppConfigurationsManagedMobileAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementTargetedManagedAppConfigurationsCreateApps**](DeviceAppManagementTargetedManagedAppConfigurationsManagedMobileAppApi.md#DeviceAppManagementTargetedManagedAppConfigurationsCreateApps) | **Post** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsGetApps**](DeviceAppManagementTargetedManagedAppConfigurationsManagedMobileAppApi.md#DeviceAppManagementTargetedManagedAppConfigurationsGetApps) | **Get** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/apps({managedMobileApp-id}) | Get apps from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsListApps**](DeviceAppManagementTargetedManagedAppConfigurationsManagedMobileAppApi.md#DeviceAppManagementTargetedManagedAppConfigurationsListApps) | **Get** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/apps | Get apps from deviceAppManagement
[**DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps**](DeviceAppManagementTargetedManagedAppConfigurationsManagedMobileAppApi.md#DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps) | **Patch** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/apps({managedMobileApp-id}) | Update the navigation property apps in deviceAppManagement



## DeviceAppManagementTargetedManagedAppConfigurationsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementTargetedManagedAppConfigurationsCreateApps(ctx, targetedManagedAppConfigurationId, microsoftGraphManagedMobileApp)
Create new navigation property to apps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**microsoftGraphManagedMobileApp** | [**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementTargetedManagedAppConfigurationsGetApps(ctx, targetedManagedAppConfigurationId, managedMobileAppId, optional)
Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsGetAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsGetAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsListApps

> CollectionOfManagedMobileApp DeviceAppManagementTargetedManagedAppConfigurationsListApps(ctx, targetedManagedAppConfigurationId, optional)
Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
 **optional** | ***DeviceAppManagementTargetedManagedAppConfigurationsListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementTargetedManagedAppConfigurationsListAppsOpts struct


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

[**CollectionOfManagedMobileApp**](Collection of managedMobileApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps

> DeviceAppManagementTargetedManagedAppConfigurationsUpdateApps(ctx, targetedManagedAppConfigurationId, managedMobileAppId, microsoftGraphManagedMobileApp)
Update the navigation property apps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
**microsoftGraphManagedMobileApp** | [**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md)| New navigation property values | 

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

