# \DomainsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsForceDelete**](DomainsActionsApi.md#DomainsForceDelete) | **Post** /domains({domain-id})/forceDelete | Invoke action forceDelete
[**DomainsVerify**](DomainsActionsApi.md#DomainsVerify) | **Post** /domains({domain-id})/verify | Invoke action verify



## DomainsForceDelete

> DomainsForceDelete(ctx, domainId, inlineObject29)
Invoke action forceDelete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 
**inlineObject29** | [**InlineObject29**](InlineObject29.md)|  | 

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


## DomainsVerify

> AnyOfmicrosoftGraphDomain DomainsVerify(ctx, domainId)
Invoke action verify

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string**| key: domain-id of domain | 

### Return type

[**AnyOfmicrosoftGraphDomain**](anyOf&lt;microsoft.graph.domain&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

