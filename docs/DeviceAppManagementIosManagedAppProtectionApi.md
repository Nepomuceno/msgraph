# \DeviceAppManagementIosManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementCreateIosManagedAppProtections) | **Post** /deviceAppManagement/iosManagedAppProtections | Create new navigation property to iosManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementGetIosManagedAppProtections) | **Get** /deviceAppManagement/iosManagedAppProtections({iosManagedAppProtection-id}) | Get iosManagedAppProtections from deviceAppManagement
[**DeviceAppManagementListIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementListIosManagedAppProtections) | **Get** /deviceAppManagement/iosManagedAppProtections | Get iosManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementUpdateIosManagedAppProtections) | **Patch** /deviceAppManagement/iosManagedAppProtections({iosManagedAppProtection-id}) | Update the navigation property iosManagedAppProtections in deviceAppManagement



## DeviceAppManagementCreateIosManagedAppProtections

> MicrosoftGraphIosManagedAppProtection DeviceAppManagementCreateIosManagedAppProtections(ctx, microsoftGraphIosManagedAppProtection)
Create new navigation property to iosManagedAppProtections for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphIosManagedAppProtection** | [**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md)| New navigation property | 

### Return type

[**MicrosoftGraphIosManagedAppProtection**](microsoft.graph.iosManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetIosManagedAppProtections

> MicrosoftGraphIosManagedAppProtection DeviceAppManagementGetIosManagedAppProtections(ctx, iosManagedAppProtectionId, optional)
Get iosManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
 **optional** | ***DeviceAppManagementGetIosManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetIosManagedAppProtectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphIosManagedAppProtection**](microsoft.graph.iosManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListIosManagedAppProtections

> CollectionOfIosManagedAppProtection DeviceAppManagementListIosManagedAppProtections(ctx, optional)
Get iosManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListIosManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListIosManagedAppProtectionsOpts struct


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

[**CollectionOfIosManagedAppProtection**](Collection of iosManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateIosManagedAppProtections

> DeviceAppManagementUpdateIosManagedAppProtections(ctx, iosManagedAppProtectionId, microsoftGraphIosManagedAppProtection)
Update the navigation property iosManagedAppProtections in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string**| key: iosManagedAppProtection-id of iosManagedAppProtection | 
**microsoftGraphIosManagedAppProtection** | [**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md)| New navigation property values | 

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

