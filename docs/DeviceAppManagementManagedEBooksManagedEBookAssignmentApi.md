# \DeviceAppManagementManagedEBooksManagedEBookAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedEBooksCreateAssignments**](DeviceAppManagementManagedEBooksManagedEBookAssignmentApi.md#DeviceAppManagementManagedEBooksCreateAssignments) | **Post** /deviceAppManagement/managedEBooks({managedEBook-id})/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementManagedEBooksGetAssignments**](DeviceAppManagementManagedEBooksManagedEBookAssignmentApi.md#DeviceAppManagementManagedEBooksGetAssignments) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/assignments({managedEBookAssignment-id}) | Get assignments from deviceAppManagement
[**DeviceAppManagementManagedEBooksListAssignments**](DeviceAppManagementManagedEBooksManagedEBookAssignmentApi.md#DeviceAppManagementManagedEBooksListAssignments) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id})/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementManagedEBooksUpdateAssignments**](DeviceAppManagementManagedEBooksManagedEBookAssignmentApi.md#DeviceAppManagementManagedEBooksUpdateAssignments) | **Patch** /deviceAppManagement/managedEBooks({managedEBook-id})/assignments({managedEBookAssignment-id}) | Update the navigation property assignments in deviceAppManagement



## DeviceAppManagementManagedEBooksCreateAssignments

> MicrosoftGraphManagedEBookAssignment DeviceAppManagementManagedEBooksCreateAssignments(ctx, managedEBookId, microsoftGraphManagedEBookAssignment)
Create new navigation property to assignments for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**microsoftGraphManagedEBookAssignment** | [**MicrosoftGraphManagedEBookAssignment**](MicrosoftGraphManagedEBookAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedEBookAssignment**](microsoft.graph.managedEBookAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksGetAssignments

> MicrosoftGraphManagedEBookAssignment DeviceAppManagementManagedEBooksGetAssignments(ctx, managedEBookId, managedEBookAssignmentId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**managedEBookAssignmentId** | **string**| key: managedEBookAssignment-id of managedEBookAssignment | 
 **optional** | ***DeviceAppManagementManagedEBooksGetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksGetAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedEBookAssignment**](microsoft.graph.managedEBookAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksListAssignments

> CollectionOfManagedEBookAssignment DeviceAppManagementManagedEBooksListAssignments(ctx, managedEBookId, optional)
Get assignments from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
 **optional** | ***DeviceAppManagementManagedEBooksListAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementManagedEBooksListAssignmentsOpts struct


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

[**CollectionOfManagedEBookAssignment**](Collection of managedEBookAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementManagedEBooksUpdateAssignments

> DeviceAppManagementManagedEBooksUpdateAssignments(ctx, managedEBookId, managedEBookAssignmentId, microsoftGraphManagedEBookAssignment)
Update the navigation property assignments in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**managedEBookAssignmentId** | **string**| key: managedEBookAssignment-id of managedEBookAssignment | 
**microsoftGraphManagedEBookAssignment** | [**MicrosoftGraphManagedEBookAssignment**](MicrosoftGraphManagedEBookAssignment.md)| New navigation property values | 

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

