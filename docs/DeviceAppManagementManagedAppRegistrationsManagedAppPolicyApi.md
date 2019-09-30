# \DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies) | **Post** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/appliedPolicies | Create new navigation property to appliedPolicies for deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies) | **Post** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/intendedPolicies | Create new navigation property to intendedPolicies for deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/appliedPolicies({managedAppPolicy-id}) | Get appliedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/intendedPolicies({managedAppPolicy-id}) | Get intendedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsListAppliedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsListAppliedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/appliedPolicies | Get appliedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsListIntendedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsListIntendedPolicies) | **Get** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/intendedPolicies | Get intendedPolicies from deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies) | **Patch** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/appliedPolicies({managedAppPolicy-id}) | Update the navigation property appliedPolicies in deviceAppManagement
[**DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies**](DeviceAppManagementManagedAppRegistrationsManagedAppPolicyApi.md#DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies) | **Patch** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/intendedPolicies({managedAppPolicy-id}) | Update the navigation property intendedPolicies in deviceAppManagement



## DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsCreateAppliedPolicies(ctx, managedAppRegistrationId, microsoftGraphManagedAppPolicy)
Create new navigation property to appliedPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsCreateIntendedPolicies(ctx, managedAppRegistrationId, microsoftGraphManagedAppPolicy)
Create new navigation property to intendedPolicies for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsGetAppliedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, optional)
Get appliedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsGetAppliedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsGetAppliedPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementManagedAppRegistrationsGetIntendedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, optional)
Get intendedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsGetIntendedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsGetIntendedPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsListAppliedPolicies

> CollectionOfManagedAppPolicy DeviceAppManagementManagedAppRegistrationsListAppliedPolicies(ctx, managedAppRegistrationId, optional)
Get appliedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsListAppliedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsListAppliedPoliciesOpts struct


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

[**CollectionOfManagedAppPolicy**](Collection of managedAppPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsListIntendedPolicies

> CollectionOfManagedAppPolicy DeviceAppManagementManagedAppRegistrationsListIntendedPolicies(ctx, managedAppRegistrationId, optional)
Get intendedPolicies from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
 **optional** | ***DeviceAppManagementManagedAppRegistrationsListIntendedPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedAppRegistrationsListIntendedPoliciesOpts struct


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

[**CollectionOfManagedAppPolicy**](Collection of managedAppPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies

> DeviceAppManagementManagedAppRegistrationsUpdateAppliedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, microsoftGraphManagedAppPolicy)
Update the navigation property appliedPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property values | 

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


## DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies

> DeviceAppManagementManagedAppRegistrationsUpdateIntendedPolicies(ctx, managedAppRegistrationId, managedAppPolicyId, microsoftGraphManagedAppPolicy)
Update the navigation property intendedPolicies in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)| New navigation property values | 

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

