# \DeviceAppManagementMobileAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementCreateMobileApps) | **Post** /deviceAppManagement/mobileApps | Create new navigation property to mobileApps for deviceAppManagement
[**DeviceAppManagementGetMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementGetMobileApps) | **Get** /deviceAppManagement/mobileApps({mobileApp-id}) | Get mobileApps from deviceAppManagement
[**DeviceAppManagementListMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementListMobileApps) | **Get** /deviceAppManagement/mobileApps | Get mobileApps from deviceAppManagement
[**DeviceAppManagementUpdateMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementUpdateMobileApps) | **Patch** /deviceAppManagement/mobileApps({mobileApp-id}) | Update the navigation property mobileApps in deviceAppManagement



## DeviceAppManagementCreateMobileApps

> MicrosoftGraphMobileApp DeviceAppManagementCreateMobileApps(ctx, microsoftGraphMobileApp)
Create new navigation property to mobileApps for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphMobileApp** | [**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md)| New navigation property | 

### Return type

[**MicrosoftGraphMobileApp**](microsoft.graph.mobileApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetMobileApps

> MicrosoftGraphMobileApp DeviceAppManagementGetMobileApps(ctx, mobileAppId, optional)
Get mobileApps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
 **optional** | ***DeviceAppManagementGetMobileAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetMobileAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileApp**](microsoft.graph.mobileApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMobileApps

> CollectionOfMobileApp DeviceAppManagementListMobileApps(ctx, optional)
Get mobileApps from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListMobileAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListMobileAppsOpts struct


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

[**CollectionOfMobileApp**](Collection of mobileApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateMobileApps

> DeviceAppManagementUpdateMobileApps(ctx, mobileAppId, microsoftGraphMobileApp)
Update the navigation property mobileApps in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**microsoftGraphMobileApp** | [**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md)| New navigation property values | 

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

