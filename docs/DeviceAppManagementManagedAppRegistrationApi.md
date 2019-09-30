# \DeviceAppManagementManagedAppRegistrationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementCreateManagedAppRegistrations) | **Post** /deviceAppManagement/managedAppRegistrations | Create new navigation property to managedAppRegistrations for deviceAppManagement
[**DeviceAppManagementGetManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementGetManagedAppRegistrations) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id}) | Get managedAppRegistrations from deviceAppManagement
[**DeviceAppManagementListManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementListManagedAppRegistrations) | **Get** /deviceAppManagement/managedAppRegistrations | Get managedAppRegistrations from deviceAppManagement
[**DeviceAppManagementUpdateManagedAppRegistrations**](DeviceAppManagementManagedAppRegistrationApi.md#DeviceAppManagementUpdateManagedAppRegistrations) | **Patch** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id}) | Update the navigation property managedAppRegistrations in deviceAppManagement



## DeviceAppManagementCreateManagedAppRegistrations

> MicrosoftGraphManagedAppRegistration DeviceAppManagementCreateManagedAppRegistrations(ctx, microsoftGraphManagedAppRegistration)
Create new navigation property to managedAppRegistrations for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedAppRegistration** | [**MicrosoftGraphManagedAppRegistration**](MicrosoftGraphManagedAppRegistration.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetManagedAppRegistrations

> MicrosoftGraphManagedAppRegistration DeviceAppManagementGetManagedAppRegistrations(ctx, managedAppRegistrationId, optional)
Get managedAppRegistrations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementGetManagedAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetManagedAppRegistrationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedAppRegistrations

> CollectionOfManagedAppRegistration DeviceAppManagementListManagedAppRegistrations(ctx, optional)
Get managedAppRegistrations from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListManagedAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListManagedAppRegistrationsOpts struct


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

[**CollectionOfManagedAppRegistration**](Collection of managedAppRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateManagedAppRegistrations

> DeviceAppManagementUpdateManagedAppRegistrations(ctx, managedAppRegistrationId, microsoftGraphManagedAppRegistration)
Update the navigation property managedAppRegistrations in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppRegistration** | [**MicrosoftGraphManagedAppRegistration**](MicrosoftGraphManagedAppRegistration.md)| New navigation property values | 

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

