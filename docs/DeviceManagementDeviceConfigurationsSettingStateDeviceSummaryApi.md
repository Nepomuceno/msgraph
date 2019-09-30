# \DeviceManagementDeviceConfigurationsSettingStateDeviceSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceConfigurationsCreateDeviceSettingStateSummaries**](DeviceManagementDeviceConfigurationsSettingStateDeviceSummaryApi.md#DeviceManagementDeviceConfigurationsCreateDeviceSettingStateSummaries) | **Post** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceSettingStateSummaries | Create new navigation property to deviceSettingStateSummaries for deviceManagement
[**DeviceManagementDeviceConfigurationsGetDeviceSettingStateSummaries**](DeviceManagementDeviceConfigurationsSettingStateDeviceSummaryApi.md#DeviceManagementDeviceConfigurationsGetDeviceSettingStateSummaries) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceSettingStateSummaries({settingStateDeviceSummary-id}) | Get deviceSettingStateSummaries from deviceManagement
[**DeviceManagementDeviceConfigurationsListDeviceSettingStateSummaries**](DeviceManagementDeviceConfigurationsSettingStateDeviceSummaryApi.md#DeviceManagementDeviceConfigurationsListDeviceSettingStateSummaries) | **Get** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceSettingStateSummaries | Get deviceSettingStateSummaries from deviceManagement
[**DeviceManagementDeviceConfigurationsUpdateDeviceSettingStateSummaries**](DeviceManagementDeviceConfigurationsSettingStateDeviceSummaryApi.md#DeviceManagementDeviceConfigurationsUpdateDeviceSettingStateSummaries) | **Patch** /deviceManagement/deviceConfigurations({deviceConfiguration-id})/deviceSettingStateSummaries({settingStateDeviceSummary-id}) | Update the navigation property deviceSettingStateSummaries in deviceManagement



## DeviceManagementDeviceConfigurationsCreateDeviceSettingStateSummaries

> MicrosoftGraphSettingStateDeviceSummary DeviceManagementDeviceConfigurationsCreateDeviceSettingStateSummaries(ctx, deviceConfigurationId, microsoftGraphSettingStateDeviceSummary)
Create new navigation property to deviceSettingStateSummaries for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**microsoftGraphSettingStateDeviceSummary** | [**MicrosoftGraphSettingStateDeviceSummary**](MicrosoftGraphSettingStateDeviceSummary.md)| New navigation property | 

### Return type

[**MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsGetDeviceSettingStateSummaries

> MicrosoftGraphSettingStateDeviceSummary DeviceManagementDeviceConfigurationsGetDeviceSettingStateSummaries(ctx, deviceConfigurationId, settingStateDeviceSummaryId, optional)
Get deviceSettingStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**settingStateDeviceSummaryId** | **string**| key: settingStateDeviceSummary-id of settingStateDeviceSummary | 
 **optional** | ***DeviceManagementDeviceConfigurationsGetDeviceSettingStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsGetDeviceSettingStateSummariesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsListDeviceSettingStateSummaries

> CollectionOfSettingStateDeviceSummary DeviceManagementDeviceConfigurationsListDeviceSettingStateSummaries(ctx, deviceConfigurationId, optional)
Get deviceSettingStateSummaries from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
 **optional** | ***DeviceManagementDeviceConfigurationsListDeviceSettingStateSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementDeviceConfigurationsListDeviceSettingStateSummariesOpts struct


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

[**CollectionOfSettingStateDeviceSummary**](Collection of settingStateDeviceSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsUpdateDeviceSettingStateSummaries

> DeviceManagementDeviceConfigurationsUpdateDeviceSettingStateSummaries(ctx, deviceConfigurationId, settingStateDeviceSummaryId, microsoftGraphSettingStateDeviceSummary)
Update the navigation property deviceSettingStateSummaries in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string**| key: deviceConfiguration-id of deviceConfiguration | 
**settingStateDeviceSummaryId** | **string**| key: settingStateDeviceSummary-id of settingStateDeviceSummary | 
**microsoftGraphSettingStateDeviceSummary** | [**MicrosoftGraphSettingStateDeviceSummary**](MicrosoftGraphSettingStateDeviceSummary.md)| New navigation property values | 

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

