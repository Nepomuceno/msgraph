# \GroupLifecyclePoliciesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupLifecyclePoliciesAddGroup**](GroupLifecyclePoliciesActionsApi.md#GroupLifecyclePoliciesAddGroup) | **Post** /groupLifecyclePolicies({groupLifecyclePolicy-id})/addGroup | Invoke action addGroup
[**GroupLifecyclePoliciesRemoveGroup**](GroupLifecyclePoliciesActionsApi.md#GroupLifecyclePoliciesRemoveGroup) | **Post** /groupLifecyclePolicies({groupLifecyclePolicy-id})/removeGroup | Invoke action removeGroup



## GroupLifecyclePoliciesAddGroup

> bool GroupLifecyclePoliciesAddGroup(ctx, groupLifecyclePolicyId, inlineObject30)
Invoke action addGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
**inlineObject30** | [**InlineObject30**](InlineObject30.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesRemoveGroup

> bool GroupLifecyclePoliciesRemoveGroup(ctx, groupLifecyclePolicyId, inlineObject31)
Invoke action removeGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string**| key: groupLifecyclePolicy-id of groupLifecyclePolicy | 
**inlineObject31** | [**InlineObject31**](InlineObject31.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

