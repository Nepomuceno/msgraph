# \DeviceManagementFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestA373**](DeviceManagementFunctionsApi.md#DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestA373) | **Get** /deviceManagement/applePushNotificationCertificate/downloadApplePushNotificationCertificateSigningRequest() | Invoke function downloadApplePushNotificationCertificateSigningRequest
[**DeviceManagementGetEffectivePermissions04fa**](DeviceManagementFunctionsApi.md#DeviceManagementGetEffectivePermissions04fa) | **Get** /deviceManagement/getEffectivePermissions(scope&#x3D;{scope}) | Invoke function getEffectivePermissions
[**DeviceManagementVerifyWindowsEnrollmentAutoDiscovery135e**](DeviceManagementFunctionsApi.md#DeviceManagementVerifyWindowsEnrollmentAutoDiscovery135e) | **Get** /deviceManagement/verifyWindowsEnrollmentAutoDiscovery(domainName&#x3D;{domainName}) | Invoke function verifyWindowsEnrollmentAutoDiscovery



## DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestA373

> string DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestA373(ctx, )
Invoke function downloadApplePushNotificationCertificateSigningRequest

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetEffectivePermissions04fa

> []AnyOfmicrosoftGraphRolePermission DeviceManagementGetEffectivePermissions04fa(ctx, scope)
Invoke function getEffectivePermissions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**|  | 

### Return type

[**[]AnyOfmicrosoftGraphRolePermission**](anyOf&lt;microsoft.graph.rolePermission&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementVerifyWindowsEnrollmentAutoDiscovery135e

> bool DeviceManagementVerifyWindowsEnrollmentAutoDiscovery135e(ctx, domainName)
Invoke function verifyWindowsEnrollmentAutoDiscovery

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

