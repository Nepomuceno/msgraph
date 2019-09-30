# \DeviceAppManagementManagedEBooksUserInstallStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedEBooksCreateUserStateSummary**](DeviceAppManagementManagedEBooksUserInstallStateSummaryApi.md#DeviceAppManagementManagedEBooksCreateUserStateSummary) | **Post** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary | Create new navigation property to userStateSummary for deviceAppManagement
[**DeviceAppManagementManagedEBooksGetUserStateSummary**](DeviceAppManagementManagedEBooksUserInstallStateSummaryApi.md#DeviceAppManagementManagedEBooksGetUserStateSummary) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary({userInstallStateSummary-id}) | Get userStateSummary from deviceAppManagement
[**DeviceAppManagementManagedEBooksListUserStateSummary**](DeviceAppManagementManagedEBooksUserInstallStateSummaryApi.md#DeviceAppManagementManagedEBooksListUserStateSummary) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary | Get userStateSummary from deviceAppManagement
[**DeviceAppManagementManagedEBooksUpdateUserStateSummary**](DeviceAppManagementManagedEBooksUserInstallStateSummaryApi.md#DeviceAppManagementManagedEBooksUpdateUserStateSummary) | **Patch** /deviceAppManagement/managedEBooks({managedEBook-id})/userStateSummary({userInstallStateSummary-id}) | Update the navigation property userStateSummary in deviceAppManagement



## DeviceAppManagementManagedEBooksCreateUserStateSummary

> MicrosoftGraphUserInstallStateSummary DeviceAppManagementManagedEBooksCreateUserStateSummary(ctx, managedEBookId, microsoftGraphUserInstallStateSummary)
Create new navigation property to userStateSummary for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**microsoftGraphUserInstallStateSummary** | [**MicrosoftGraphUserInstallStateSummary**](MicrosoftGraphUserInstallStateSummary.md)| New navigation property | 

### Return type

[**MicrosoftGraphUserInstallStateSummary**](microsoft.graph.userInstallStateSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksGetUserStateSummary

> MicrosoftGraphUserInstallStateSummary DeviceAppManagementManagedEBooksGetUserStateSummary(ctx, managedEBookId, userInstallStateSummaryId, optional)
Get userStateSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**userInstallStateSummaryId** | **string**| key: userInstallStateSummary-id of userInstallStateSummary | 
 **optional** | ***DeviceAppManagementManagedEBooksGetUserStateSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksGetUserStateSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphUserInstallStateSummary**](microsoft.graph.userInstallStateSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksListUserStateSummary

> CollectionOfUserInstallStateSummary DeviceAppManagementManagedEBooksListUserStateSummary(ctx, managedEBookId, optional)
Get userStateSummary from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
 **optional** | ***DeviceAppManagementManagedEBooksListUserStateSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksListUserStateSummaryOpts struct


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

[**CollectionOfUserInstallStateSummary**](Collection of userInstallStateSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksUpdateUserStateSummary

> DeviceAppManagementManagedEBooksUpdateUserStateSummary(ctx, managedEBookId, userInstallStateSummaryId, microsoftGraphUserInstallStateSummary)
Update the navigation property userStateSummary in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**userInstallStateSummaryId** | **string**| key: userInstallStateSummary-id of userInstallStateSummary | 
**microsoftGraphUserInstallStateSummary** | [**MicrosoftGraphUserInstallStateSummary**](MicrosoftGraphUserInstallStateSummary.md)| New navigation property values | 

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

