# \PlannerPlansPlannerPlanDetailsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerPlansGetDetails**](PlannerPlansPlannerPlanDetailsApi.md#PlannerPlansGetDetails) | **Get** /planner/plans({plannerPlan-id})/details | Get details from planner
[**PlannerPlansUpdateDetails**](PlannerPlansPlannerPlanDetailsApi.md#PlannerPlansUpdateDetails) | **Patch** /planner/plans({plannerPlan-id})/details | Update the navigation property details in planner



## PlannerPlansGetDetails

> MicrosoftGraphPlannerPlanDetails PlannerPlansGetDetails(ctx, plannerPlanId, optional)
Get details from planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
 **optional** | ***PlannerPlansGetDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlannerPlansGetDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphPlannerPlanDetails**](microsoft.graph.plannerPlanDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerPlansUpdateDetails

> PlannerPlansUpdateDetails(ctx, plannerPlanId, microsoftGraphPlannerPlanDetails)
Update the navigation property details in planner

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string**| key: plannerPlan-id of plannerPlan | 
**microsoftGraphPlannerPlanDetails** | [**MicrosoftGraphPlannerPlanDetails**](MicrosoftGraphPlannerPlanDetails.md)| New navigation property values | 

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
