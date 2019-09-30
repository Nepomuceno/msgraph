# \DeviceAppManagementActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedAppPoliciesTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedAppPoliciesTargetApps) | **Post** /deviceAppManagement/managedAppPolicies({managedAppPolicy-id})/targetApps | Invoke action targetApps
[**DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps) | **Post** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/appliedPolicies({managedAppPolicy-id})/targetApps | Invoke action targetApps
[**DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps) | **Post** /deviceAppManagement/managedAppRegistrations({managedAppRegistration-id})/intendedPolicies({managedAppPolicy-id})/targetApps | Invoke action targetApps
[**DeviceAppManagementManagedEBooksAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementManagedEBooksAssign) | **Post** /deviceAppManagement/managedEBooks({managedEBook-id})/assign | Invoke action assign
[**DeviceAppManagementMobileAppConfigurationsAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementMobileAppConfigurationsAssign) | **Post** /deviceAppManagement/mobileAppConfigurations({managedDeviceMobileAppConfiguration-id})/assign | Invoke action assign
[**DeviceAppManagementMobileAppsAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementMobileAppsAssign) | **Post** /deviceAppManagement/mobileApps({mobileApp-id})/assign | Invoke action assign
[**DeviceAppManagementSyncMicrosoftStoreForBusinessApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementSyncMicrosoftStoreForBusinessApps) | **Post** /deviceAppManagement/syncMicrosoftStoreForBusinessApps | Invoke action syncMicrosoftStoreForBusinessApps
[**DeviceAppManagementTargetedManagedAppConfigurationsAssign**](DeviceAppManagementActionsApi.md#DeviceAppManagementTargetedManagedAppConfigurationsAssign) | **Post** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/assign | Invoke action assign
[**DeviceAppManagementTargetedManagedAppConfigurationsTargetApps**](DeviceAppManagementActionsApi.md#DeviceAppManagementTargetedManagedAppConfigurationsTargetApps) | **Post** /deviceAppManagement/targetedManagedAppConfigurations({targetedManagedAppConfiguration-id})/targetApps | Invoke action targetApps
[**DeviceAppManagementVppTokensSyncLicenses**](DeviceAppManagementActionsApi.md#DeviceAppManagementVppTokensSyncLicenses) | **Post** /deviceAppManagement/vppTokens({vppToken-id})/syncLicenses | Invoke action syncLicenses



## DeviceAppManagementManagedAppPoliciesTargetApps

> DeviceAppManagementManagedAppPoliciesTargetApps(ctx, managedAppPolicyId, inlineObject)
Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

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


## DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps

> DeviceAppManagementManagedAppRegistrationsAppliedPoliciesTargetApps(ctx, managedAppRegistrationId, managedAppPolicyId, inlineObject1)
Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

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


## DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps

> DeviceAppManagementManagedAppRegistrationsIntendedPoliciesTargetApps(ctx, managedAppRegistrationId, managedAppPolicyId, inlineObject2)
Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppRegistrationId** | **string**| key: managedAppRegistration-id of managedAppRegistration | 
**managedAppPolicyId** | **string**| key: managedAppPolicy-id of managedAppPolicy | 
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

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


## DeviceAppManagementManagedEBooksAssign

> DeviceAppManagementManagedEBooksAssign(ctx, managedEBookId, inlineObject3)
Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

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


## DeviceAppManagementMobileAppConfigurationsAssign

> DeviceAppManagementMobileAppConfigurationsAssign(ctx, managedDeviceMobileAppConfigurationId, inlineObject4)
Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceMobileAppConfigurationId** | **string**| key: managedDeviceMobileAppConfiguration-id of managedDeviceMobileAppConfiguration | 
**inlineObject4** | [**InlineObject4**](InlineObject4.md)|  | 

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


## DeviceAppManagementMobileAppsAssign

> DeviceAppManagementMobileAppsAssign(ctx, mobileAppId, inlineObject5)
Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | 

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


## DeviceAppManagementSyncMicrosoftStoreForBusinessApps

> DeviceAppManagementSyncMicrosoftStoreForBusinessApps(ctx, )
Invoke action syncMicrosoftStoreForBusinessApps

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementTargetedManagedAppConfigurationsAssign

> DeviceAppManagementTargetedManagedAppConfigurationsAssign(ctx, targetedManagedAppConfigurationId, inlineObject6)
Invoke action assign

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**inlineObject6** | [**InlineObject6**](InlineObject6.md)|  | 

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


## DeviceAppManagementTargetedManagedAppConfigurationsTargetApps

> DeviceAppManagementTargetedManagedAppConfigurationsTargetApps(ctx, targetedManagedAppConfigurationId, inlineObject7)
Invoke action targetApps

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetedManagedAppConfigurationId** | **string**| key: targetedManagedAppConfiguration-id of targetedManagedAppConfiguration | 
**inlineObject7** | [**InlineObject7**](InlineObject7.md)|  | 

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


## DeviceAppManagementVppTokensSyncLicenses

> AnyOfmicrosoftGraphVppToken DeviceAppManagementVppTokensSyncLicenses(ctx, vppTokenId)
Invoke action syncLicenses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string**| key: vppToken-id of vppToken | 

### Return type

[**AnyOfmicrosoftGraphVppToken**](anyOf&lt;microsoft.graph.vppToken&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

