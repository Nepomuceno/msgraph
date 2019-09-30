# \DeviceManagementRoleDefinitionsRoleAssignmentsRoleDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition**](DeviceManagementRoleDefinitionsRoleAssignmentsRoleDefinitionApi.md#DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition) | **Get** /deviceManagement/roleDefinitions({roleDefinition-id})/roleAssignments({roleAssignment-id})/roleDefinition | Get roleDefinition from deviceManagement



## DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition

> MicrosoftGraphRoleDefinition DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition(ctx, roleDefinitionId, roleAssignmentId, optional)
Get roleDefinition from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
**roleAssignmentId** | **string**| key: roleAssignment-id of roleAssignment | 
 **optional** | ***DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinitionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphRoleDefinition**](microsoft.graph.roleDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

