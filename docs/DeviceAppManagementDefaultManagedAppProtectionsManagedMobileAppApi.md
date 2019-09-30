# \DeviceAppManagementDefaultManagedAppProtectionsManagedMobileAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementDefaultManagedAppProtectionsCreateApps**](DeviceAppManagementDefaultManagedAppProtectionsManagedMobileAppApi.md#DeviceAppManagementDefaultManagedAppProtectionsCreateApps) | **Post** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id})/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsGetApps**](DeviceAppManagementDefaultManagedAppProtectionsManagedMobileAppApi.md#DeviceAppManagementDefaultManagedAppProtectionsGetApps) | **Get** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id})/apps({managedMobileApp-id}) | Get apps from deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsListApps**](DeviceAppManagementDefaultManagedAppProtectionsManagedMobileAppApi.md#DeviceAppManagementDefaultManagedAppProtectionsListApps) | **Get** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id})/apps | Get apps from deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsUpdateApps**](DeviceAppManagementDefaultManagedAppProtectionsManagedMobileAppApi.md#DeviceAppManagementDefaultManagedAppProtectionsUpdateApps) | **Patch** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id})/apps({managedMobileApp-id}) | Update the navigation property apps in deviceAppManagement



## DeviceAppManagementDefaultManagedAppProtectionsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementDefaultManagedAppProtectionsCreateApps(ctx, defaultManagedAppProtectionId, microsoftGraphManagedMobileApp)
Create new navigation property to apps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
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


## DeviceAppManagementDefaultManagedAppProtectionsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementDefaultManagedAppProtectionsGetApps(ctx, defaultManagedAppProtectionId, managedMobileAppId, optional)
Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
**managedMobileAppId** | **string**| key: managedMobileApp-id of managedMobileApp | 
 **optional** | ***DeviceAppManagementDefaultManagedAppProtectionsGetAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementDefaultManagedAppProtectionsGetAppsOpts struct


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


## DeviceAppManagementDefaultManagedAppProtectionsListApps

> CollectionOfManagedMobileApp DeviceAppManagementDefaultManagedAppProtectionsListApps(ctx, defaultManagedAppProtectionId, optional)
Get apps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
 **optional** | ***DeviceAppManagementDefaultManagedAppProtectionsListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementDefaultManagedAppProtectionsListAppsOpts struct


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


## DeviceAppManagementDefaultManagedAppProtectionsUpdateApps

> DeviceAppManagementDefaultManagedAppProtectionsUpdateApps(ctx, defaultManagedAppProtectionId, managedMobileAppId, microsoftGraphManagedMobileApp)
Update the navigation property apps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
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

