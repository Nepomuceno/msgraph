# \DeviceAppManagementManagedEBooksEBookInstallSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedEBooksGetInstallSummary**](DeviceAppManagementManagedEBooksEBookInstallSummaryApi.md#DeviceAppManagementManagedEBooksGetInstallSummary) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/installSummary | Get installSummary from deviceAppManagement
[**DeviceAppManagementManagedEBooksUpdateInstallSummary**](DeviceAppManagementManagedEBooksEBookInstallSummaryApi.md#DeviceAppManagementManagedEBooksUpdateInstallSummary) | **Patch** /deviceAppManagement/managedEBooks({managedEBook-id})/installSummary | Update the navigation property installSummary in deviceAppManagement



## DeviceAppManagementManagedEBooksGetInstallSummary

> MicrosoftGraphEBookInstallSummary DeviceAppManagementManagedEBooksGetInstallSummary(ctx, managedEBookId, optional)
Get installSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
 **optional** | ***DeviceAppManagementManagedEBooksGetInstallSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksGetInstallSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEBookInstallSummary**](microsoft.graph.eBookInstallSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksUpdateInstallSummary

> DeviceAppManagementManagedEBooksUpdateInstallSummary(ctx, managedEBookId, microsoftGraphEBookInstallSummary)
Update the navigation property installSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**microsoftGraphEBookInstallSummary** | [**MicrosoftGraphEBookInstallSummary**](MicrosoftGraphEBookInstallSummary.md)| New navigation property values | 

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

