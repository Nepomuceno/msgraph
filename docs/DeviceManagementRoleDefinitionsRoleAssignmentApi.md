# \DeviceManagementRoleDefinitionsRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementRoleDefinitionsCreateRoleAssignments**](DeviceManagementRoleDefinitionsRoleAssignmentApi.md#DeviceManagementRoleDefinitionsCreateRoleAssignments) | **Post** /deviceManagement/roleDefinitions({roleDefinition-id})/roleAssignments | Create new navigation property to roleAssignments for deviceManagement
[**DeviceManagementRoleDefinitionsGetRoleAssignments**](DeviceManagementRoleDefinitionsRoleAssignmentApi.md#DeviceManagementRoleDefinitionsGetRoleAssignments) | **Get** /deviceManagement/roleDefinitions({roleDefinition-id})/roleAssignments({roleAssignment-id}) | Get roleAssignments from deviceManagement
[**DeviceManagementRoleDefinitionsListRoleAssignments**](DeviceManagementRoleDefinitionsRoleAssignmentApi.md#DeviceManagementRoleDefinitionsListRoleAssignments) | **Get** /deviceManagement/roleDefinitions({roleDefinition-id})/roleAssignments | Get roleAssignments from deviceManagement
[**DeviceManagementRoleDefinitionsUpdateRoleAssignments**](DeviceManagementRoleDefinitionsRoleAssignmentApi.md#DeviceManagementRoleDefinitionsUpdateRoleAssignments) | **Patch** /deviceManagement/roleDefinitions({roleDefinition-id})/roleAssignments({roleAssignment-id}) | Update the navigation property roleAssignments in deviceManagement



## DeviceManagementRoleDefinitionsCreateRoleAssignments

> MicrosoftGraphRoleAssignment DeviceManagementRoleDefinitionsCreateRoleAssignments(ctx, roleDefinitionId, microsoftGraphRoleAssignment)
Create new navigation property to roleAssignments for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
**microsoftGraphRoleAssignment** | [**MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md)| New navigation property | 

### Return type

[**MicrosoftGraphRoleAssignment**](microsoft.graph.roleAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsGetRoleAssignments

> MicrosoftGraphRoleAssignment DeviceManagementRoleDefinitionsGetRoleAssignments(ctx, roleDefinitionId, roleAssignmentId, optional)
Get roleAssignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
**roleAssignmentId** | **string**| key: roleAssignment-id of roleAssignment | 
 **optional** | ***DeviceManagementRoleDefinitionsGetRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementRoleDefinitionsGetRoleAssignmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphRoleAssignment**](microsoft.graph.roleAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsListRoleAssignments

> CollectionOfRoleAssignment DeviceManagementRoleDefinitionsListRoleAssignments(ctx, roleDefinitionId, optional)
Get roleAssignments from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
 **optional** | ***DeviceManagementRoleDefinitionsListRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementRoleDefinitionsListRoleAssignmentsOpts struct


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

[**CollectionOfRoleAssignment**](Collection of roleAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsUpdateRoleAssignments

> DeviceManagementRoleDefinitionsUpdateRoleAssignments(ctx, roleDefinitionId, roleAssignmentId, microsoftGraphRoleAssignment)
Update the navigation property roleAssignments in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
**roleAssignmentId** | **string**| key: roleAssignment-id of roleAssignment | 
**microsoftGraphRoleAssignment** | [**MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md)| New navigation property values | 

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

