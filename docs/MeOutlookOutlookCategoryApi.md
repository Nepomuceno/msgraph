# \MeOutlookOutlookCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeOutlookCreateMasterCategories**](MeOutlookOutlookCategoryApi.md#MeOutlookCreateMasterCategories) | **Post** /me/outlook/masterCategories | Create new navigation property to masterCategories for me
[**MeOutlookGetMasterCategories**](MeOutlookOutlookCategoryApi.md#MeOutlookGetMasterCategories) | **Get** /me/outlook/masterCategories({outlookCategory-id}) | Get masterCategories from me
[**MeOutlookListMasterCategories**](MeOutlookOutlookCategoryApi.md#MeOutlookListMasterCategories) | **Get** /me/outlook/masterCategories | Get masterCategories from me
[**MeOutlookUpdateMasterCategories**](MeOutlookOutlookCategoryApi.md#MeOutlookUpdateMasterCategories) | **Patch** /me/outlook/masterCategories({outlookCategory-id}) | Update the navigation property masterCategories in me



## MeOutlookCreateMasterCategories

> MicrosoftGraphOutlookCategory MeOutlookCreateMasterCategories(ctx, microsoftGraphOutlookCategory)
Create new navigation property to masterCategories for me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)| New navigation property | 

### Return type

[**MicrosoftGraphOutlookCategory**](microsoft.graph.outlookCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookGetMasterCategories

> MicrosoftGraphOutlookCategory MeOutlookGetMasterCategories(ctx, outlookCategoryId, optional)
Get masterCategories from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string**| key: outlookCategory-id of outlookCategory | 
 **optional** | ***MeOutlookGetMasterCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOutlookGetMasterCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookCategory**](microsoft.graph.outlookCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookListMasterCategories

> CollectionOfOutlookCategory MeOutlookListMasterCategories(ctx, optional)
Get masterCategories from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOutlookListMasterCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOutlookListMasterCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Show only the first n items | 
 **skip** | **optional.Int32**| Skip the first n items | 
 **filter** | **optional.String**| Filter items by property values | 
 **count** | **optional.Bool**| Include count of items | 
 **orderby** | [**optional.Interface of []string**](string.md)| Order items by property values | 
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**CollectionOfOutlookCategory**](Collection of outlookCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookUpdateMasterCategories

> MeOutlookUpdateMasterCategories(ctx, outlookCategoryId, microsoftGraphOutlookCategory)
Update the navigation property masterCategories in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string**| key: outlookCategory-id of outlookCategory | 
**microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)| New navigation property values | 

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

