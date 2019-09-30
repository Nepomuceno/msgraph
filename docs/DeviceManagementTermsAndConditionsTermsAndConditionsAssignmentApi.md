# \DeviceManagementTermsAndConditionsTermsAndConditionsAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementTermsAndConditionsCreateAssignments**](DeviceManagementTermsAndConditionsTermsAndConditionsAssignmentApi.md#DeviceManagementTermsAndConditionsCreateAssignments) | **Post** /deviceManagement/termsAndConditions({termsAndConditions-id})/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementTermsAndConditionsGetAssignments**](DeviceManagementTermsAndConditionsTermsAndConditionsAssignmentApi.md#DeviceManagementTermsAndConditionsGetAssignments) | **Get** /deviceManagement/termsAndConditions({termsAndConditions-id})/assignments({termsAndConditionsAssignment-id}) | Get assignments from deviceManagement
[**DeviceManagementTermsAndConditionsListAssignments**](DeviceManagementTermsAndConditionsTermsAndConditionsAssignmentApi.md#DeviceManagementTermsAndConditionsListAssignments) | **Get** /deviceManagement/termsAndConditions({termsAndConditions-id})/assignments | Get assignments from deviceManagement
[**DeviceManagementTermsAndConditionsUpdateAssignments**](DeviceManagementTermsAndConditionsTermsAndConditionsAssignmentApi.md#DeviceManagementTermsAndConditionsUpdateAssignments) | **Patch** /deviceManagement/termsAndConditions({termsAndConditions-id})/assignments({termsAndConditionsAssignment-id}) | Update the navigation property assignments in deviceManagement



## DeviceManagementTermsAndConditionsCreateAssignments

> MicrosoftGraphTermsAndConditionsAssignment DeviceManagementTermsAndConditionsCreateAssignments(ctx, termsAndConditionsId, microsoftGraphTermsAndConditionsAssignment)
Create new navigation property to assignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**microsoftGraphTermsAndConditionsAssignment** | [**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditionsAssignment**](microsoft.graph.termsAndConditionsAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsGetAssignments

> MicrosoftGraphTermsAndConditionsAssignment DeviceManagementTermsAndConditionsGetAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string**| key: termsAndConditionsAssignment-id of termsAndConditionsAssignment | 
 **optional** | ***DeviceManagementTermsAndConditionsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditionsAssignment**](microsoft.graph.termsAndConditionsAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsListAssignments

> CollectionOfTermsAndConditionsAssignment DeviceManagementTermsAndConditionsListAssignments(ctx, termsAndConditionsId, optional)
Get assignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
 **optional** | ***DeviceManagementTermsAndConditionsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementTermsAndConditionsListAssignmentsOpts struct


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

[**CollectionOfTermsAndConditionsAssignment**](Collection of termsAndConditionsAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsUpdateAssignments

> DeviceManagementTermsAndConditionsUpdateAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId, microsoftGraphTermsAndConditionsAssignment)
Update the navigation property assignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string**| key: termsAndConditions-id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string**| key: termsAndConditionsAssignment-id of termsAndConditionsAssignment | 
**microsoftGraphTermsAndConditionsAssignment** | [**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md)| New navigation property values | 

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

