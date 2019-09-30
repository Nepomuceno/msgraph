# \PlannerTasksPlannerAssignedToTaskBoardTaskFormatApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerTasksGetAssignedToTaskBoardFormat**](PlannerTasksPlannerAssignedToTaskBoardTaskFormatApi.md#PlannerTasksGetAssignedToTaskBoardFormat) | **Get** /planner/tasks({plannerTask-id})/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from planner
[**PlannerTasksUpdateAssignedToTaskBoardFormat**](PlannerTasksPlannerAssignedToTaskBoardTaskFormatApi.md#PlannerTasksUpdateAssignedToTaskBoardFormat) | **Patch** /planner/tasks({plannerTask-id})/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in planner



## PlannerTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat PlannerTasksGetAssignedToTaskBoardFormat(ctx, plannerTaskId, optional)
Get assignedToTaskBoardFormat from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetAssignedToTaskBoardFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetAssignedToTaskBoardFormatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](microsoft.graph.plannerAssignedToTaskBoardTaskFormat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksUpdateAssignedToTaskBoardFormat

> PlannerTasksUpdateAssignedToTaskBoardFormat(ctx, plannerTaskId, microsoftGraphPlannerAssignedToTaskBoardTaskFormat)
Update the navigation property assignedToTaskBoardFormat in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerAssignedToTaskBoardTaskFormat** | [**MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat.md)| New navigation property values | 

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

