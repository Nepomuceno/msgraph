# \WorkbooksWorkbookCommentsWorkbookCommentReplyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksWorkbookCommentsCreateReplies**](WorkbooksWorkbookCommentsWorkbookCommentReplyApi.md#WorkbooksWorkbookCommentsCreateReplies) | **Post** /workbooks({driveItem-id})/workbook/comments({workbookComment-id})/replies | Create new navigation property to replies for workbooks
[**WorkbooksWorkbookCommentsGetReplies**](WorkbooksWorkbookCommentsWorkbookCommentReplyApi.md#WorkbooksWorkbookCommentsGetReplies) | **Get** /workbooks({driveItem-id})/workbook/comments({workbookComment-id})/replies({workbookCommentReply-id}) | Get replies from workbooks
[**WorkbooksWorkbookCommentsListReplies**](WorkbooksWorkbookCommentsWorkbookCommentReplyApi.md#WorkbooksWorkbookCommentsListReplies) | **Get** /workbooks({driveItem-id})/workbook/comments({workbookComment-id})/replies | Get replies from workbooks
[**WorkbooksWorkbookCommentsUpdateReplies**](WorkbooksWorkbookCommentsWorkbookCommentReplyApi.md#WorkbooksWorkbookCommentsUpdateReplies) | **Patch** /workbooks({driveItem-id})/workbook/comments({workbookComment-id})/replies({workbookCommentReply-id}) | Update the navigation property replies in workbooks



## WorkbooksWorkbookCommentsCreateReplies

> MicrosoftGraphWorkbookCommentReply WorkbooksWorkbookCommentsCreateReplies(ctx, driveItemId, workbookCommentId, microsoftGraphWorkbookCommentReply)
Create new navigation property to replies for workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookCommentId** | **string**| key: workbookComment-id of workbookComment | 
**microsoftGraphWorkbookCommentReply** | [**MicrosoftGraphWorkbookCommentReply**](MicrosoftGraphWorkbookCommentReply.md)| New navigation property | 

### Return type

[**MicrosoftGraphWorkbookCommentReply**](microsoft.graph.workbookCommentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookCommentsGetReplies

> MicrosoftGraphWorkbookCommentReply WorkbooksWorkbookCommentsGetReplies(ctx, driveItemId, workbookCommentId, workbookCommentReplyId, optional)
Get replies from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookCommentId** | **string**| key: workbookComment-id of workbookComment | 
**workbookCommentReplyId** | **string**| key: workbookCommentReply-id of workbookCommentReply | 
 **optional** | ***WorkbooksWorkbookCommentsGetRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookCommentsGetRepliesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphWorkbookCommentReply**](microsoft.graph.workbookCommentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookCommentsListReplies

> CollectionOfWorkbookCommentReply WorkbooksWorkbookCommentsListReplies(ctx, driveItemId, workbookCommentId, optional)
Get replies from workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookCommentId** | **string**| key: workbookComment-id of workbookComment | 
 **optional** | ***WorkbooksWorkbookCommentsListRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkbooksWorkbookCommentsListRepliesOpts struct


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

[**CollectionOfWorkbookCommentReply**](Collection of workbookCommentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksWorkbookCommentsUpdateReplies

> WorkbooksWorkbookCommentsUpdateReplies(ctx, driveItemId, workbookCommentId, workbookCommentReplyId, microsoftGraphWorkbookCommentReply)
Update the navigation property replies in workbooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string**| key: driveItem-id of driveItem | 
**workbookCommentId** | **string**| key: workbookComment-id of workbookComment | 
**workbookCommentReplyId** | **string**| key: workbookCommentReply-id of workbookCommentReply | 
**microsoftGraphWorkbookCommentReply** | [**MicrosoftGraphWorkbookCommentReply**](MicrosoftGraphWorkbookCommentReply.md)| New navigation property values | 

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

