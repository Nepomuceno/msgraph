# \DeviceAppManagementManagedEBooksUserStateSummaryDeviceInstallStateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedEBooksUserStateSummaryCreateDeviceStates**](DeviceAppManagementManagedEBooksUserStateSummaryDeviceInstallStateApi.md#DeviceAppManagementManagedEBooksUserStateSummaryCreateDeviceStates) | **Post** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary({userInstallStateSummary-id})/deviceStates | Create new navigation property to deviceStates for deviceAppManagement
[**DeviceAppManagementManagedEBooksUserStateSummaryGetDeviceStates**](DeviceAppManagementManagedEBooksUserStateSummaryDeviceInstallStateApi.md#DeviceAppManagementManagedEBooksUserStateSummaryGetDeviceStates) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary({userInstallStateSummary-id})/deviceStates({deviceInstallState-id}) | Get deviceStates from deviceAppManagement
[**DeviceAppManagementManagedEBooksUserStateSummaryListDeviceStates**](DeviceAppManagementManagedEBooksUserStateSummaryDeviceInstallStateApi.md#DeviceAppManagementManagedEBooksUserStateSummaryListDeviceStates) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary({userInstallStateSummary-id})/deviceStates | Get deviceStates from deviceAppManagement
[**DeviceAppManagementManagedEBooksUserStateSummaryUpdateDeviceStates**](DeviceAppManagementManagedEBooksUserStateSummaryDeviceInstallStateApi.md#DeviceAppManagementManagedEBooksUserStateSummaryUpdateDeviceStates) | **Patch** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary({userInstallStateSummary-id})/deviceStates({deviceInstallState-id}) | Update the navigation property deviceStates in deviceAppManagement



## DeviceAppManagementManagedEBooksUserStateSummaryCreateDeviceStates

> MicrosoftGraphDeviceInstallState DeviceAppManagementManagedEBooksUserStateSummaryCreateDeviceStates(ctx, managedEBookId, userInstallStateSummaryId, microsoftGraphDeviceInstallState)
Create new navigation property to deviceStates for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**userInstallStateSummaryId** | **string**| key: userInstallStateSummary-id of userInstallStateSummary | 
**microsoftGraphDeviceInstallState** | [**MicrosoftGraphDeviceInstallState**](MicrosoftGraphDeviceInstallState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceInstallState**](microsoft.graph.deviceInstallState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksUserStateSummaryGetDeviceStates

> MicrosoftGraphDeviceInstallState DeviceAppManagementManagedEBooksUserStateSummaryGetDeviceStates(ctx, managedEBookId, userInstallStateSummaryId, deviceInstallStateId, optional)
Get deviceStates from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**userInstallStateSummaryId** | **string**| key: userInstallStateSummary-id of userInstallStateSummary | 
**deviceInstallStateId** | **string**| key: deviceInstallState-id of deviceInstallState | 
 **optional** | ***DeviceAppManagementManagedEBooksUserStateSummaryGetDeviceStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksUserStateSummaryGetDeviceStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceInstallState**](microsoft.graph.deviceInstallState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksUserStateSummaryListDeviceStates

> CollectionOfDeviceInstallState DeviceAppManagementManagedEBooksUserStateSummaryListDeviceStates(ctx, managedEBookId, userInstallStateSummaryId, optional)
Get deviceStates from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**userInstallStateSummaryId** | **string**| key: userInstallStateSummary-id of userInstallStateSummary | 
 **optional** | ***DeviceAppManagementManagedEBooksUserStateSummaryListDeviceStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksUserStateSummaryListDeviceStatesOpts struct


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

[**CollectionOfDeviceInstallState**](Collection of deviceInstallState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksUserStateSummaryUpdateDeviceStates

> DeviceAppManagementManagedEBooksUserStateSummaryUpdateDeviceStates(ctx, managedEBookId, userInstallStateSummaryId, deviceInstallStateId, microsoftGraphDeviceInstallState)
Update the navigation property deviceStates in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**userInstallStateSummaryId** | **string**| key: userInstallStateSummary-id of userInstallStateSummary | 
**deviceInstallStateId** | **string**| key: deviceInstallState-id of deviceInstallState | 
**microsoftGraphDeviceInstallState** | [**MicrosoftGraphDeviceInstallState**](MicrosoftGraphDeviceInstallState.md)| New navigation property values | 

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

