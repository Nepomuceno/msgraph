# \WorkbooksWorkbookWorkbookCommentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookCreateComments**](WorkbooksWorkbookWorkbookCommentApi.md#WorkbooksWorkbookCreateComments) | **Post** /workbooks({driveItem-id})/workbook/comments | Create new navigation property to comments for workbooks
[**WorkbooksWorkbookGetComments**](WorkbooksWorkbookWorkbookCommentApi.md#WorkbooksWorkbookGetComments) | **Get** /workbooks({driveItem-id})/workbook/comments({workbookComment-id}) | Get comments from workbooks
[**WorkbooksWorkbookListComments**](WorkbooksWorkbookWorkbookCommentApi.md#WorkbooksWorkbookListComments) | **Get** /workbooks({driveItem-id})/workbook/comments | Get comments from workbooks
[**WorkbooksWorkbookUpdateComments**](WorkbooksWorkbookWorkbookCommentApi.md#WorkbooksWorkbookUpdateComments) | **Patch** /workbooks({driveItem-id})/workbook/comments({workbookComment-id}) | Update the navigation property comments in workbooks



## WorkbooksWorkbookCreateComments

> MicrosoftGraphWorkbookComment WorkbooksWorkbookCreateComments(ctx, driveItemId, microsoftGraphWorkbookComment)
Create new navigation property to comments for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**microsoftGraphWorkbookComment** | [**MicrosoftGraphWorkbookComment**](MicrosoftGraphWorkbookComment.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookComment**](microsoft.graph.workbookComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookGetComments

> MicrosoftGraphWorkbookComment WorkbooksWorkbookGetComments(ctx, driveItemId, workbookCommentId, optional)
Get comments from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookCommentId** | **string**| key: workbookComment-id of workbookComment | 
 **optional** | ***WorkbooksWorkbookGetCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookGetCommentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookComment**](microsoft.graph.workbookComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookListComments

> CollectionOfWorkbookComment WorkbooksWorkbookListComments(ctx, driveItemId, optional)
Get comments from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
 **optional** | ***WorkbooksWorkbookListCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookListCommentsOpts struct


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

[**CollectionOfWorkbookComment**](Collection of workbookComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookUpdateComments

> WorkbooksWorkbookUpdateComments(ctx, driveItemId, workbookCommentId, microsoftGraphWorkbookComment)
Update the navigation property comments in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookCommentId** | **string**| key: workbookComment-id of workbookComment | 
**microsoftGraphWorkbookComment** | [**MicrosoftGraphWorkbookComment**](MicrosoftGraphWorkbookComment.md)| New navigation property values | 

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

