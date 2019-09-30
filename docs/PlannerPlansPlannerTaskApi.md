# \PlannerPlansPlannerTaskApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerPlansGetTasks**](PlannerPlansPlannerTaskApi.md#PlannerPlansGetTasks) | **Get** /planner/plans({plannerPlan-id})/tasks({plannerTask-id}) | Get tasks from planner
[**PlannerPlansListTasks**](PlannerPlansPlannerTaskApi.md#PlannerPlansListTasks) | **Get** /planner/plans({plannerPlan-id})/tasks | Get tasks from planner



## PlannerPlansGetTasks

> MicrosoftGraphPlannerTask PlannerPlansGetTasks(ctx, plannerPlanId, plannerTaskId, optional)
Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerPlansGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansGetTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerPlansListTasks

> CollectionOfPlannerTask PlannerPlansListTasks(ctx, plannerPlanId, optional)
Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***PlannerPlansListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansListTasksOpts struct


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

[**CollectionOfPlannerTask**](Collection of plannerTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

