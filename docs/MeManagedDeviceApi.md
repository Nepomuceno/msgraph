# \MeManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateManagedDevices**](MeManagedDeviceApi.md#MeCreateManagedDevices) | **Post** /me/managedDevices | Create new navigation property to managedDevices for me
[**MeGetManagedDevices**](MeManagedDeviceApi.md#MeGetManagedDevices) | **Get** /me/managedDevices({managedDevice-id}) | Get managedDevices from me
[**MeListManagedDevices**](MeManagedDeviceApi.md#MeListManagedDevices) | **Get** /me/managedDevices | Get managedDevices from me
[**MeUpdateManagedDevices**](MeManagedDeviceApi.md#MeUpdateManagedDevices) | **Patch** /me/managedDevices({managedDevice-id}) | Update the navigation property managedDevices in me



## MeCreateManagedDevices

> MicrosoftGraphManagedDevice MeCreateManagedDevices(ctx, microsoftGraphManagedDevice)
Create new navigation property to managedDevices for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGetManagedDevices

> MicrosoftGraphManagedDevice MeGetManagedDevices(ctx, managedDeviceId, optional)
Get managedDevices from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
 **optional** | ***MeGetManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeGetManagedDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListManagedDevices

> CollectionOfManagedDevice MeListManagedDevices(ctx, optional)
Get managedDevices from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeListManagedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeListManagedDevicesOpts struct


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

[**CollectionOfManagedDevice**](Collection of managedDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateManagedDevices

> MeUpdateManagedDevices(ctx, managedDeviceId, microsoftGraphManagedDevice)
Update the navigation property managedDevices in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string**| key: managedDevice-id of managedDevice | 
**microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)| New navigation property values | 

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

