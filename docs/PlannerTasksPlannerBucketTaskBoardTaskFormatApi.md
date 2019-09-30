# \PlannerTasksPlannerBucketTaskBoardTaskFormatApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerTasksGetBucketTaskBoardFormat**](PlannerTasksPlannerBucketTaskBoardTaskFormatApi.md#PlannerTasksGetBucketTaskBoardFormat) | **Get** /planner/tasks({plannerTask-id})/bucketTaskBoardFormat | Get bucketTaskBoardFormat from planner
[**PlannerTasksUpdateBucketTaskBoardFormat**](PlannerTasksPlannerBucketTaskBoardTaskFormatApi.md#PlannerTasksUpdateBucketTaskBoardFormat) | **Patch** /planner/tasks({plannerTask-id})/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in planner



## PlannerTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat PlannerTasksGetBucketTaskBoardFormat(ctx, plannerTaskId, optional)
Get bucketTaskBoardFormat from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetBucketTaskBoardFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetBucketTaskBoardFormatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerBucketTaskBoardTaskFormat**](microsoft.graph.plannerBucketTaskBoardTaskFormat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksUpdateBucketTaskBoardFormat

> PlannerTasksUpdateBucketTaskBoardFormat(ctx, plannerTaskId, microsoftGraphPlannerBucketTaskBoardTaskFormat)
Update the navigation property bucketTaskBoardFormat in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerBucketTaskBoardTaskFormat** | [**MicrosoftGraphPlannerBucketTaskBoardTaskFormat**](MicrosoftGraphPlannerBucketTaskBoardTaskFormat.md)| New navigation property values | 

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

