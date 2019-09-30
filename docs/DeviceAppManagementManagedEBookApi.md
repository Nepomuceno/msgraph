# \DeviceAppManagementManagedEBookApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedEBooks**](DeviceAppManagementManagedEBookApi.md#DeviceAppManagementCreateManagedEBooks) | **Post** /deviceAppManagement/managedEBooks | Create new navigation property to managedEBooks for deviceAppManagement
[**DeviceAppManagementGetManagedEBooks**](DeviceAppManagementManagedEBookApi.md#DeviceAppManagementGetManagedEBooks) | **Get** /deviceAppManagement/managedEBooks({managedEBook-id}) | Get managedEBooks from deviceAppManagement
[**DeviceAppManagementListManagedEBooks**](DeviceAppManagementManagedEBookApi.md#DeviceAppManagementListManagedEBooks) | **Get** /deviceAppManagement/managedEBooks | Get managedEBooks from deviceAppManagement
[**DeviceAppManagementUpdateManagedEBooks**](DeviceAppManagementManagedEBookApi.md#DeviceAppManagementUpdateManagedEBooks) | **Patch** /deviceAppManagement/managedEBooks({managedEBook-id}) | Update the navigation property managedEBooks in deviceAppManagement



## DeviceAppManagementCreateManagedEBooks

> MicrosoftGraphManagedEBook DeviceAppManagementCreateManagedEBooks(ctx, microsoftGraphManagedEBook)
Create new navigation property to managedEBooks for deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphManagedEBook** | [**MicrosoftGraphManagedEBook**](MicrosoftGraphManagedEBook.md)| New navigation property | 

### Return type

[**MicrosoftGraphManagedEBook**](microsoft.graph.managedEBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementGetManagedEBooks

> MicrosoftGraphManagedEBook DeviceAppManagementGetManagedEBooks(ctx, managedEBookId, optional)
Get managedEBooks from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
 **optional** | ***DeviceAppManagementGetManagedEBooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementGetManagedEBooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphManagedEBook**](microsoft.graph.managedEBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedEBooks

> CollectionOfManagedEBook DeviceAppManagementListManagedEBooks(ctx, optional)
Get managedEBooks from deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceAppManagementListManagedEBooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceAppManagementListManagedEBooksOpts struct


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

[**CollectionOfManagedEBook**](Collection of managedEBook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateManagedEBooks

> DeviceAppManagementUpdateManagedEBooks(ctx, managedEBookId, microsoftGraphManagedEBook)
Update the navigation property managedEBooks in deviceAppManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedEBookId** | **string**| key: managedEBook-id of managedEBook | 
**microsoftGraphManagedEBook** | [**MicrosoftGraphManagedEBook**](MicrosoftGraphManagedEBook.md)| New navigation property values | 

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

