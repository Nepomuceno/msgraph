# \PlannerPlannerTaskApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerCreateTasks**](PlannerPlannerTaskApi.md#PlannerCreateTasks) | **Post** /planner/tasks | Create new navigation property to tasks for planner
[**PlannerGetTasks**](PlannerPlannerTaskApi.md#PlannerGetTasks) | **Get** /planner/tasks({plannerTask-id}) | Get tasks from planner
[**PlannerListTasks**](PlannerPlannerTaskApi.md#PlannerListTasks) | **Get** /planner/tasks | Get tasks from planner
[**PlannerUpdateTasks**](PlannerPlannerTaskApi.md#PlannerUpdateTasks) | **Patch** /planner/tasks({plannerTask-id}) | Update the navigation property tasks in planner



## PlannerCreateTasks

> MicrosoftGraphPlannerTask PlannerCreateTasks(ctx, microsoftGraphPlannerTask)
Create new navigation property to tasks for planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphPlannerTask** | [**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md)| New navigation property | 

### Return type

[**MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerGetTasks

> MicrosoftGraphPlannerTask PlannerGetTasks(ctx, plannerTaskId, optional)
Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
 **optional** | ***PlannerGetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerGetTasksOpts struct


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


## PlannerListTasks

> CollectionOfPlannerTask PlannerListTasks(ctx, optional)
Get tasks from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlannerListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerListTasksOpts struct


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


## PlannerUpdateTasks

> PlannerUpdateTasks(ctx, plannerTaskId, microsoftGraphPlannerTask)
Update the navigation property tasks in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string**| key: plannerTask-id of plannerTask | 
**microsoftGraphPlannerTask** | [**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md)| New navigation property values | 

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

