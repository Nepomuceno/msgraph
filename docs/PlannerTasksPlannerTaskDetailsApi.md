# \PlannerTasksPlannerTaskDetailsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerTasksGetDetails**](PlannerTasksPlannerTaskDetailsApi.md#PlannerTasksGetDetails) | **Get** /planner/tasks({plannerTask-id})/details | Get details from planner
[**PlannerTasksUpdateDetails**](PlannerTasksPlannerTaskDetailsApi.md#PlannerTasksUpdateDetails) | **Patch** /planner/tasks({plannerTask-id})/details | Update the navigation property details in planner



## PlannerTasksGetDetails

> MicrosoftGraphPlannerTaskDetails PlannerTasksGetDetails(ctx, plannerTaskId, optional)
Get details from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerTasksGetDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerTasksGetDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerTaskDetails**](microsoft.graph.plannerTaskDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksUpdateDetails

> PlannerTasksUpdateDetails(ctx, plannerTaskId, microsoftGraphPlannerTaskDetails)
Update the navigation property details in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerTaskDetails** | [**MicrosoftGraphPlannerTaskDetails**](MicrosoftGraphPlannerTaskDetails.md)| New navigation property values | 

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

