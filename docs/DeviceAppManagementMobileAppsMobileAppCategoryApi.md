# \DeviceAppManagementMobileAppsMobileAppCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppsGetCategories**](DeviceAppManagementMobileAppsMobileAppCategoryApi.md#DeviceAppManagementMobileAppsGetCategories) | **Get** /deviceAppManagement/mobileApps({mobileApp-id})/categories({mobileAppCategory-id}) | Get categories from deviceAppManagement
[**DeviceAppManagementMobileAppsListCategories**](DeviceAppManagementMobileAppsMobileAppCategoryApi.md#DeviceAppManagementMobileAppsListCategories) | **Get** /deviceAppManagement/mobileApps({mobileApp-id})/categories | Get categories from deviceAppManagement



## DeviceAppManagementMobileAppsGetCategories

> MicrosoftGraphMobileAppCategory DeviceAppManagementMobileAppsGetCategories(ctx, mobileAppId, mobileAppCategoryId, optional)
Get categories from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**mobileAppCategoryId** | **string**| key: mobileAppCategory-id of mobileAppCategory | 
 **optional** | ***DeviceAppManagementMobileAppsGetCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsGetCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileAppCategory**](microsoft.graph.mobileAppCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsListCategories

> CollectionOfMobileAppCategory DeviceAppManagementMobileAppsListCategories(ctx, mobileAppId, optional)
Get categories from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
 **optional** | ***DeviceAppManagementMobileAppsListCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsListCategoriesOpts struct


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

[**CollectionOfMobileAppCategory**](Collection of mobileAppCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

