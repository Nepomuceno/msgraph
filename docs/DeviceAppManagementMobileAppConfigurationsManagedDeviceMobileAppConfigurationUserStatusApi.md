# \DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppConfigurationsCreateUserStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserStatusApi.md#DeviceAppManagementMobileAppConfigurationsCreateUserStatuses) | **Post** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/userStatuses | Create new navigation property to userStatuses for deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsGetUserStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserStatusApi.md#DeviceAppManagementMobileAppConfigurationsGetUserStatuses) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/userStatuses({managedDeviceMobileAppConfigurationUserStatus-id}) | Get userStatuses from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsListUserStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserStatusApi.md#DeviceAppManagementMobileAppConfigurationsListUserStatuses) | **Get** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/userStatuses | Get userStatuses from deviceAppManagement
[**DeviceAppManagementMobileAppConfigurationsUpdateUserStatuses**](DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationUserStatusApi.md#DeviceAppManagementMobileAppConfigurationsUpdateUserStatuses) | **Patch** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/userStatuses({managedDeviceMobileAppConfigurationUserStatus-id}) | Update the navigation property userStatuses in deviceAppManagement



## DeviceAppManagementMobileAppConfigurationsCreateUserStatuses

> MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus DeviceAppManagementMobileAppConfigurationsCreateUserStatuses(ctx, managedDeviceMobileAppConfigurationId, microsoftGraphManagedDeviceMobileAppConfigurationUserStatus)
Create new navigation property to userStatuses for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**microsoftGraphManagedDeviceMobileAppConfigurationUserStatus** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus**](MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus**](microsoft.graph.managedDeviceMobileAppConfigurationUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsGetUserStatuses

> MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus DeviceAppManagementMobileAppConfigurationsGetUserStatuses(ctx, managedDeviceMobileAppConfigurationId, managedDeviceMobileAppConfigurationUserStatusId, optional)
Get userStatuses from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**managedDeviceMobileAppConfigurationUserStatusId** | **string**| key: managedDeviceMobileAppConfigurationUserStatus-id of managedDeviceMobileAppConfigurationUserStatus | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsGetUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsGetUserStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus**](microsoft.graph.managedDeviceMobileAppConfigurationUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsListUserStatuses

> CollectionOfManagedDeviceMobileAppConfigurationUserStatus DeviceAppManagementMobileAppConfigurationsListUserStatuses(ctx, managedDeviceMobileAppConfigurationId, optional)
Get userStatuses from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
 **optional** | ***DeviceAppManagementMobileAppConfigurationsListUserStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppConfigurationsListUserStatusesOpts struct


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

[**CollectionOfManagedDeviceMobileAppConfigurationUserStatus**](Collection of managedDeviceMobileAppConfigurationUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppConfigurationsUpdateUserStatuses

> DeviceAppManagementMobileAppConfigurationsUpdateUserStatuses(ctx, managedDeviceMobileAppConfigurationId, managedDeviceMobileAppConfigurationUserStatusId, microsoftGraphManagedDeviceMobileAppConfigurationUserStatus)
Update the navigation property userStatuses in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**managedDeviceMobileAppConfigurationUserStatusId** | **string**| key: managedDeviceMobileAppConfigurationUserStatus-id of managedDeviceMobileAppConfigurationUserStatus | 
**microsoftGraphManagedDeviceMobileAppConfigurationUserStatus** | [**MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus**](MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus.md)| New navigation property values | 

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

