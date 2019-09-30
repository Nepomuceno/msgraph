# \DeviceManagementTermsAndConditionsTermsAndConditionsAcceptanceStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementTermsAndConditionsCreateAcceptanceStatuses**](DeviceManagementTermsAndConditionsTermsAndConditionsAcceptanceStatusApi.md#DeviceManagementTermsAndConditionsCreateAcceptanceStatuses) | **Post** /deviceManagement/termsAndConditions({termsAndConditions-id})/acceptanceStatuses | Create new navigation property to acceptanceStatuses for deviceManagement
[**DeviceManagementTermsAndConditionsGetAcceptanceStatuses**](DeviceManagementTermsAndConditionsTermsAndConditionsAcceptanceStatusApi.md#DeviceManagementTermsAndConditionsGetAcceptanceStatuses) | **Get** /deviceManagement/termsAndConditions({termsAndConditions-id})/acceptanceStatuses({termsAndConditionsAcceptanceStatus-id}) | Get acceptanceStatuses from deviceManagement
[**DeviceManagementTermsAndConditionsListAcceptanceStatuses**](DeviceManagementTermsAndConditionsTermsAndConditionsAcceptanceStatusApi.md#DeviceManagementTermsAndConditionsListAcceptanceStatuses) | **Get** /deviceManagement/termsAndConditions({termsAndConditions-id})/acceptanceStatuses | Get acceptanceStatuses from deviceManagement
[**DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses**](DeviceManagementTermsAndConditionsTermsAndConditionsAcceptanceStatusApi.md#DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses) | **Patch** /deviceManagement/termsAndConditions({termsAndConditions-id})/acceptanceStatuses({termsAndConditionsAcceptanceStatus-id}) | Update the navigation property acceptanceStatuses in deviceManagement



## DeviceManagementTermsAndConditionsCreateAcceptanceStatuses

> MicrosoftGraphTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsCreateAcceptanceStatuses(ctx, termsAndConditionsId, microsoftGraphTermsAndConditionsAcceptanceStatus)
Create new navigation property to acceptanceStatuses for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**microsoftGraphTermsAndConditionsAcceptanceStatus** | [**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md)| New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](microsoft.graph.termsAndConditionsAcceptanceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsGetAcceptanceStatuses

> MicrosoftGraphTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsGetAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId, optional)
Get acceptanceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string**| key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus | 
 **optional** | ***DeviceManagementTermsAndConditionsGetAcceptanceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsGetAcceptanceStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](microsoft.graph.termsAndConditionsAcceptanceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsListAcceptanceStatuses

> CollectionOfTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsListAcceptanceStatuses(ctx, termsAndConditionsId, optional)
Get acceptanceStatuses from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
 **optional** | ***DeviceManagementTermsAndConditionsListAcceptanceStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsListAcceptanceStatusesOpts struct


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

[**CollectionOfTermsAndConditionsAcceptanceStatus**](Collection of termsAndConditionsAcceptanceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses

> DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId, microsoftGraphTermsAndConditionsAcceptanceStatus)
Update the navigation property acceptanceStatuses in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string**| key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus | 
**microsoftGraphTermsAndConditionsAcceptanceStatus** | [**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md)| New navigation property values | 

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

