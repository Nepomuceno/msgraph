# \DirectoryObjectsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryObjectsCheckMemberGroups**](DirectoryObjectsActionsApi.md#DirectoryObjectsCheckMemberGroups) | **Post** /directoryObjects({directoryObject-id})/checkMemberGroups | Invoke action checkMemberGroups
[**DirectoryObjectsGetByIds**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetByIds) | **Post** /directoryObjects/getByIds | Invoke action getByIds
[**DirectoryObjectsGetMemberGroups**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetMemberGroups) | **Post** /directoryObjects({directoryObject-id})/getMemberGroups | Invoke action getMemberGroups
[**DirectoryObjectsGetMemberObjects**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetMemberObjects) | **Post** /directoryObjects({directoryObject-id})/getMemberObjects | Invoke action getMemberObjects
[**DirectoryObjectsRestore**](DirectoryObjectsActionsApi.md#DirectoryObjectsRestore) | **Post** /directoryObjects({directoryObject-id})/restore | Invoke action restore
[**DirectoryObjectsValidateProperties**](DirectoryObjectsActionsApi.md#DirectoryObjectsValidateProperties) | **Post** /directoryObjects/validateProperties | Invoke action validateProperties



## DirectoryObjectsCheckMemberGroups

> []string DirectoryObjectsCheckMemberGroups(ctx, directoryObjectId, inlineObject24)
Invoke action checkMemberGroups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject24** | [**InlineObject24**](InlineObject24.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetByIds

> []MicrosoftGraphDirectoryObject DirectoryObjectsGetByIds(ctx, inlineObject27)
Invoke action getByIds

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject27** | [**InlineObject27**](InlineObject27.md)|  | 

### Return type

[**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetMemberGroups

> []string DirectoryObjectsGetMemberGroups(ctx, directoryObjectId, inlineObject25)
Invoke action getMemberGroups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject25** | [**InlineObject25**](InlineObject25.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsGetMemberObjects

> []string DirectoryObjectsGetMemberObjects(ctx, directoryObjectId, inlineObject26)
Invoke action getMemberObjects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 
**inlineObject26** | [**InlineObject26**](InlineObject26.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsRestore

> AnyOfmicrosoftGraphDirectoryObject DirectoryObjectsRestore(ctx, directoryObjectId)
Invoke action restore

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string**| key: directoryObject-id of directoryObject | 

### Return type

[**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsValidateProperties

> DirectoryObjectsValidateProperties(ctx, inlineObject28)
Invoke action validateProperties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject28** | [**InlineObject28**](InlineObject28.md)|  | 

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

