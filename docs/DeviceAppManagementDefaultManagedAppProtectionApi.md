# \DeviceAppManagementDefaultManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementCreateDefaultManagedAppProtections) | **Post** /deviceAppManagement/defaultManagedAppProtections | Create new navigation property to defaultManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementGetDefaultManagedAppProtections) | **Get** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id}) | Get defaultManagedAppProtections from deviceAppManagement
[**DeviceAppManagementListDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementListDefaultManagedAppProtections) | **Get** /deviceAppManagement/defaultManagedAppProtections | Get defaultManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementUpdateDefaultManagedAppProtections) | **Patch** /deviceAppManagement/defaultManagedAppProtections({defaultManagedAppProtection-id}) | Update the navigation property defaultManagedAppProtections in deviceAppManagement



## DeviceAppManagementCreateDefaultManagedAppProtections

> MicrosoftGraphDefaultManagedAppProtection DeviceAppManagementCreateDefaultManagedAppProtections(ctx, microsoftGraphDefaultManagedAppProtection)
Create new navigation property to defaultManagedAppProtections for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphDefaultManagedAppProtection** | [**MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md)| New navigation property | 

### Return type

[**MicrosoftGraphDefaultManagedAppProtection**](microsoft.graph.defaultManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetDefaultManagedAppProtections

> MicrosoftGraphDefaultManagedAppProtection DeviceAppManagementGetDefaultManagedAppProtections(ctx, defaultManagedAppProtectionId, optional)
Get defaultManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
 **optional** | ***DeviceAppManagementGetDefaultManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetDefaultManagedAppProtectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDefaultManagedAppProtection**](microsoft.graph.defaultManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListDefaultManagedAppProtections

> CollectionOfDefaultManagedAppProtection DeviceAppManagementListDefaultManagedAppProtections(ctx, optional)
Get defaultManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListDefaultManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListDefaultManagedAppProtectionsOpts struct


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

[**CollectionOfDefaultManagedAppProtection**](Collection of defaultManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateDefaultManagedAppProtections

> DeviceAppManagementUpdateDefaultManagedAppProtections(ctx, defaultManagedAppProtectionId, microsoftGraphDefaultManagedAppProtection)
Update the navigation property defaultManagedAppProtections in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string**| key: defaultManagedAppProtection-id of defaultManagedAppProtection | 
**microsoftGraphDefaultManagedAppProtection** | [**MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md)| New navigation property values | 

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

