# \DeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceComplianceSettingStateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceComplianceSettingStateApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates) | **Post** /deviceManagement/deviceCompliancePolicySettingStateSummaries({deviceCompliancePolicySettingStateSummary-id})/deviceComplianceSettingStates | Create new navigation property to deviceComplianceSettingStates for deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceComplianceSettingStateApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries({deviceCompliancePolicySettingStateSummary-id})/deviceComplianceSettingStates({deviceComplianceSettingState-id}) | Get deviceComplianceSettingStates from deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceComplianceSettingStateApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries({deviceCompliancePolicySettingStateSummary-id})/deviceComplianceSettingStates | Get deviceComplianceSettingStates from deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceComplianceSettingStateApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates) | **Patch** /deviceManagement/deviceCompliancePolicySettingStateSummaries({deviceCompliancePolicySettingStateSummary-id})/deviceComplianceSettingStates({deviceComplianceSettingState-id}) | Update the navigation property deviceComplianceSettingStates in deviceManagement



## DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates

> MicrosoftGraphDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, microsoftGraphDeviceComplianceSettingState)
Create new navigation property to deviceComplianceSettingStates for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**microsoftGraphDeviceComplianceSettingState** | [**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md)| New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceSettingState**](microsoft.graph.deviceComplianceSettingState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates

> MicrosoftGraphDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId, optional)
Get deviceComplianceSettingStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string**| key: deviceComplianceSettingState-id of deviceComplianceSettingState | 
 **optional** | ***DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceSettingState**](microsoft.graph.deviceComplianceSettingState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates

> CollectionOfDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, optional)
Get deviceComplianceSettingStates from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
 **optional** | ***DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStatesOpts struct


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

[**CollectionOfDeviceComplianceSettingState**](Collection of deviceComplianceSettingState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates

> DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId, microsoftGraphDeviceComplianceSettingState)
Update the navigation property deviceComplianceSettingStates in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string**| key: deviceCompliancePolicySettingStateSummary-id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string**| key: deviceComplianceSettingState-id of deviceComplianceSettingState | 
**microsoftGraphDeviceComplianceSettingState** | [**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md)| New navigation property values | 

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

