# \DeviceAppManagementAndroidManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementCreateAndroidManagedAppProtections) | **Post** /deviceAppManagement/androidManagedAppProtections | Create new navigation property to androidManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementGetAndroidManagedAppProtections) | **Get** /deviceAppManagement/androidManagedAppProtections({androidManagedAppProtection-id}) | Get androidManagedAppProtections from deviceAppManagement
[**DeviceAppManagementListAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementListAndroidManagedAppProtections) | **Get** /deviceAppManagement/androidManagedAppProtections | Get androidManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementUpdateAndroidManagedAppProtections) | **Patch** /deviceAppManagement/androidManagedAppProtections({androidManagedAppProtection-id}) | Update the navigation property androidManagedAppProtections in deviceAppManagement



## DeviceAppManagementCreateAndroidManagedAppProtections

> MicrosoftGraphAndroidManagedAppProtection DeviceAppManagementCreateAndroidManagedAppProtections(ctx, microsoftGraphAndroidManagedAppProtection)
Create new navigation property to androidManagedAppProtections for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphAndroidManagedAppProtection** | [**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md)| New navigation property | 

### Return type

[**MicrosoftGraphAndroidManagedAppProtection**](microsoft.graph.androidManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetAndroidManagedAppProtections

> MicrosoftGraphAndroidManagedAppProtection DeviceAppManagementGetAndroidManagedAppProtections(ctx, androidManagedAppProtectionId, optional)
Get androidManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
 **optional** | ***DeviceAppManagementGetAndroidManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetAndroidManagedAppProtectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphAndroidManagedAppProtection**](microsoft.graph.androidManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListAndroidManagedAppProtections

> CollectionOfAndroidManagedAppProtection DeviceAppManagementListAndroidManagedAppProtections(ctx, optional)
Get androidManagedAppProtections from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListAndroidManagedAppProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListAndroidManagedAppProtectionsOpts struct


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

[**CollectionOfAndroidManagedAppProtection**](Collection of androidManagedAppProtection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateAndroidManagedAppProtections

> DeviceAppManagementUpdateAndroidManagedAppProtections(ctx, androidManagedAppProtectionId, microsoftGraphAndroidManagedAppProtection)
Update the navigation property androidManagedAppProtections in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string**| key: androidManagedAppProtection-id of androidManagedAppProtection | 
**microsoftGraphAndroidManagedAppProtection** | [**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md)| New navigation property values | 

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

