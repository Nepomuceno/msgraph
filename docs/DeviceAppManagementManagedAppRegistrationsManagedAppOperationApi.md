# \DeviceAppManagementManagedAppRegistrationsManagedAppOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedAppRegistrationsCreateOperations**](DeviceAppManagementManagedAppRegistrationsManagedAppOperationApi.md#DeviceAppManagementManagedAppRegistrationsCreateOperations) | **Post** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/operations | Create new navigation property to operations for deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsGetOperations**](DeviceAppManagementManagedAppRegistrationsManagedAppOperationApi.md#DeviceAppManagementManagedAppRegistrationsGetOperations) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/operations({managedAppOperation-id}) | Get operations from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsListOperations**](DeviceAppManagementManagedAppRegistrationsManagedAppOperationApi.md#DeviceAppManagementManagedAppRegistrationsListOperations) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/operations | Get operations from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsUpdateOperations**](DeviceAppManagementManagedAppRegistrationsManagedAppOperationApi.md#DeviceAppManagementManagedAppRegistrationsUpdateOperations) | **Patch** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/operations({managedAppOperation-id}) | Update the navigation property operations in deviceAppManagement



## DeviceAppManagementManagedAppRegistrationsCreateOperations

> MicrosoftGraphManagedAppOperation DeviceAppManagementManagedAppRegistrationsCreateOperations(ctx, managedAppRegistrationId, microsoftGraphManagedAppOperation)
Create new navigation property to operations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppOperation** | [**MicrosoftGraphManagedAppOperation**](MicrosoftGraphManagedAppOperation.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppOperation**](microsoft.graph.managedAppOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsGetOperations

> MicrosoftGraphManagedAppOperation DeviceAppManagementManagedAppRegistrationsGetOperations(ctx, managedAppRegistrationId, managedAppOperationId, optional)
Get operations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppOperationId** | **string**| key: managedAppOperation-id of managedAppOperation | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsGetOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsGetOperationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppOperation**](microsoft.graph.managedAppOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsListOperations

> CollectionOfManagedAppOperation DeviceAppManagementManagedAppRegistrationsListOperations(ctx, managedAppRegistrationId, optional)
Get operations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsListOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsListOperationsOpts struct


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

[**CollectionOfManagedAppOperation**](Collection of managedAppOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsUpdateOperations

> DeviceAppManagementManagedAppRegistrationsUpdateOperations(ctx, managedAppRegistrationId, managedAppOperationId, microsoftGraphManagedAppOperation)
Update the navigation property operations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppOperationId** | **string**| key: managedAppOperation-id of managedAppOperation | 
**microsoftGraphManagedAppOperation** | [**MicrosoftGraphManagedAppOperation**](MicrosoftGraphManagedAppOperation.md)| New navigation property values | 

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

