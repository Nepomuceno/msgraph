# \DeviceManagementTermsAndConditionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementCreateTermsAndConditions) | **Post** /deviceManagement/termsAndConditions | Create new navigation property to termsAndConditions for deviceManagement
[**DeviceManagementGetTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementGetTermsAndConditions) | **Get** /deviceManagement/termsAndConditions({termsAndConditions-id}) | Get termsAndConditions from deviceManagement
[**DeviceManagementListTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementListTermsAndConditions) | **Get** /deviceManagement/termsAndConditions | Get termsAndConditions from deviceManagement
[**DeviceManagementUpdateTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementUpdateTermsAndConditions) | **Patch** /deviceManagement/termsAndConditions({termsAndConditions-id}) | Update the navigation property termsAndConditions in deviceManagement



## DeviceManagementCreateTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementCreateTermsAndConditions(ctx, microsoftGraphTermsAndConditions)
Create new navigation property to termsAndConditions for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphTermsAndConditions** | [**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)| New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditions**](microsoft.graph.termsAndConditions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementGetTermsAndConditions(ctx, termsAndConditionsId, optional)
Get termsAndConditions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
 **optional** | ***DeviceManagementGetTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetTermsAndConditionsOpts struct


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


## DeviceManagementListTermsAndConditions

> CollectionOfTermsAndConditions DeviceManagementListTermsAndConditions(ctx, optional)
Get termsAndConditions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListTermsAndConditionsOpts struct


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

[**CollectionOfTermsAndConditions**](Collection of termsAndConditions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateTermsAndConditions

> DeviceManagementUpdateTermsAndConditions(ctx, termsAndConditionsId, microsoftGraphTermsAndConditions)
Update the navigation property termsAndConditions in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**microsoftGraphTermsAndConditions** | [**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)| New navigation property values | 

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

