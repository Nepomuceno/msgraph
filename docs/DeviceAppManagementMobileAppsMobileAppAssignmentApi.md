# \DeviceAppManagementMobileAppsMobileAppAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementMobileAppsCreateAssignments**](DeviceAppManagementMobileAppsMobileAppAssignmentApi.md#DeviceAppManagementMobileAppsCreateAssignments) | **Post** /deviceAppManagement/mobileApps({mobileApp-id})/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementMobileAppsGetAssignments**](DeviceAppManagementMobileAppsMobileAppAssignmentApi.md#DeviceAppManagementMobileAppsGetAssignments) | **Get** /deviceAppManagement/mobileApps({mobileApp-id})/assignments({mobileAppAssignment-id}) | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppsListAssignments**](DeviceAppManagementMobileAppsMobileAppAssignmentApi.md#DeviceAppManagementMobileAppsListAssignments) | **Get** /deviceAppManagement/mobileApps({mobileApp-id})/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppsUpdateAssignments**](DeviceAppManagementMobileAppsMobileAppAssignmentApi.md#DeviceAppManagementMobileAppsUpdateAssignments) | **Patch** /deviceAppManagement/mobileApps({mobileApp-id})/assignments({mobileAppAssignment-id}) | Update the navigation property assignments in deviceAppManagement



## DeviceAppManagementMobileAppsCreateAssignments

> MicrosoftGraphMobileAppAssignment DeviceAppManagementMobileAppsCreateAssignments(ctx, mobileAppId, microsoftGraphMobileAppAssignment)
Create new navigation property to assignments for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**microsoftGraphMobileAppAssignment** | [**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphMobileAppAssignment**](microsoft.graph.mobileAppAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsGetAssignments

> MicrosoftGraphMobileAppAssignment DeviceAppManagementMobileAppsGetAssignments(ctx, mobileAppId, mobileAppAssignmentId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**mobileAppAssignmentId** | **string**| key: mobileAppAssignment-id of mobileAppAssignment | 
 **optional** | ***DeviceAppManagementMobileAppsGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphMobileAppAssignment**](microsoft.graph.mobileAppAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsListAssignments

> CollectionOfMobileAppAssignment DeviceAppManagementMobileAppsListAssignments(ctx, mobileAppId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
 **optional** | ***DeviceAppManagementMobileAppsListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementMobileAppsListAssignmentsOpts struct


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

[**CollectionOfMobileAppAssignment**](Collection of mobileAppAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsUpdateAssignments

> DeviceAppManagementMobileAppsUpdateAssignments(ctx, mobileAppId, mobileAppAssignmentId, microsoftGraphMobileAppAssignment)
Update the navigation property assignments in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string**| key: mobileApp-id of mobileApp | 
**mobileAppAssignmentId** | **string**| key: mobileAppAssignment-id of mobileAppAssignment | 
**microsoftGraphMobileAppAssignment** | [**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md)| New navigation property values | 

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

