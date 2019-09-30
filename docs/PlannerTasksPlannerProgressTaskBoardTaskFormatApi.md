# \PlannerTasksPlannerProgressTaskBoardTaskFormatApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerTasksGetProgressTaskBoardFormat**](PlannerTasksPlannerProgressTaskBoardTaskFormatApi.md#PlannerTasksGetProgressTaskBoardFormat) | **Get** /planner/tasks({plannerTask-id})/progressTaskBoardFormat | Get progressTaskBoardFormat from planner
[**PlannerTasksUpdateProgressTaskBoardFormat**](PlannerTasksPlannerProgressTaskBoardTaskFormatApi.md#PlannerTasksUpdateProgressTaskBoardFormat) | **Patch** /planner/tasks({plannerTask-id})/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in planner



## PlannerTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat PlannerTasksGetProgressTaskBoardFormat(ctx, plannerTaskId, optional)
Get progressTaskBoardFormat from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetProgressTaskBoardFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetProgressTaskBoardFormatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerProgressTaskBoardTaskFormat**](microsoft.graph.plannerProgressTaskBoardTaskFormat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksUpdateProgressTaskBoardFormat

> PlannerTasksUpdateProgressTaskBoardFormat(ctx, plannerTaskId, microsoftGraphPlannerProgressTaskBoardTaskFormat)
Update the navigation property progressTaskBoardFormat in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerProgressTaskBoardTaskFormat** | [**MicrosoftGraphPlannerProgressTaskBoardTaskFormat**](MicrosoftGraphPlannerProgressTaskBoardTaskFormat.md)| New navigation property values | 

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

