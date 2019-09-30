# \DeviceManagementRoleDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementCreateRoleDefinitions) | **Post** /deviceManagement/roleDefinitions | Create new navigation property to roleDefinitions for deviceManagement
[**DeviceManagementGetRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementGetRoleDefinitions) | **Get** /deviceManagement/roleDefinitions({roleDefinition-id}) | Get roleDefinitions from deviceManagement
[**DeviceManagementListRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementListRoleDefinitions) | **Get** /deviceManagement/roleDefinitions | Get roleDefinitions from deviceManagement
[**DeviceManagementUpdateRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementUpdateRoleDefinitions) | **Patch** /deviceManagement/roleDefinitions({roleDefinition-id}) | Update the navigation property roleDefinitions in deviceManagement



## DeviceManagementCreateRoleDefinitions

> MicrosoftGraphRoleDefinition DeviceManagementCreateRoleDefinitions(ctx, microsoftGraphRoleDefinition)
Create new navigation property to roleDefinitions for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphRoleDefinition** | [**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md)| New navigation property | 

### Return type

[**MicrosoftGraphRoleDefinition**](microsoft.graph.roleDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetRoleDefinitions

> MicrosoftGraphRoleDefinition DeviceManagementGetRoleDefinitions(ctx, roleDefinitionId, optional)
Get roleDefinitions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
 **optional** | ***DeviceManagementGetRoleDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementGetRoleDefinitionsOpts struct


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


## DeviceManagementListRoleDefinitions

> CollectionOfRoleDefinition DeviceManagementListRoleDefinitions(ctx, optional)
Get roleDefinitions from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceManagementListRoleDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementListRoleDefinitionsOpts struct


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

[**CollectionOfRoleDefinition**](Collection of roleDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateRoleDefinitions

> DeviceManagementUpdateRoleDefinitions(ctx, roleDefinitionId, microsoftGraphRoleDefinition)
Update the navigation property roleDefinitions in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string**| key: roleDefinition-id of roleDefinition | 
**microsoftGraphRoleDefinition** | [**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md)| New navigation property values | 

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

