# \DeviceManagementTermsAndConditionsAcceptanceStatusesTermsAndConditionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions**](DeviceManagementTermsAndConditionsAcceptanceStatusesTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions) | **Get** /deviceManagement/termsAndConditions({termsAndConditions-id})/acceptanceStatuses({termsAndConditionsAcceptanceStatus-id})/termsAndConditions | Get termsAndConditions from deviceManagement



## DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId, optional)
Get termsAndConditions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string**| key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus | 
 **optional** | ***DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditions**](microsoft.graph.termsAndConditions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

