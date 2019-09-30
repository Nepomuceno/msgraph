# \DeviceManagementNotificationMessageTemplatesLocalizedNotificationMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplatesLocalizedNotificationMessageApi.md#DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages) | **Post** /deviceManagement/notificationMessageTemplates({notificationMessageTemplate-id})/localizedNotificationMessages | Create new navigation property to localizedNotificationMessages for deviceManagement
[**DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplatesLocalizedNotificationMessageApi.md#DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages) | **Get** /deviceManagement/notificationMessageTemplates({notificationMessageTemplate-id})/localizedNotificationMessages({localizedNotificationMessage-id}) | Get localizedNotificationMessages from deviceManagement
[**DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplatesLocalizedNotificationMessageApi.md#DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages) | **Get** /deviceManagement/notificationMessageTemplates({notificationMessageTemplate-id})/localizedNotificationMessages | Get localizedNotificationMessages from deviceManagement
[**DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplatesLocalizedNotificationMessageApi.md#DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages) | **Patch** /deviceManagement/notificationMessageTemplates({notificationMessageTemplate-id})/localizedNotificationMessages({localizedNotificationMessage-id}) | Update the navigation property localizedNotificationMessages in deviceManagement



## DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages

> MicrosoftGraphLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages(ctx, notificationMessageTemplateId, microsoftGraphLocalizedNotificationMessage)
Create new navigation property to localizedNotificationMessages for deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**microsoftGraphLocalizedNotificationMessage** | [**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md)| New navigation property | 

### Return type

[**MicrosoftGraphLocalizedNotificationMessage**](microsoft.graph.localizedNotificationMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages

> MicrosoftGraphLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId, optional)
Get localizedNotificationMessages from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string**| key: localizedNotificationMessage-id of localizedNotificationMessage | 
 **optional** | ***DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphLocalizedNotificationMessage**](microsoft.graph.localizedNotificationMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages

> CollectionOfLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages(ctx, notificationMessageTemplateId, optional)
Get localizedNotificationMessages from deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
 **optional** | ***DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessagesOpts struct


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

[**CollectionOfLocalizedNotificationMessage**](Collection of localizedNotificationMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages

> DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId, microsoftGraphLocalizedNotificationMessage)
Update the navigation property localizedNotificationMessages in deviceManagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string**| key: notificationMessageTemplate-id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string**| key: localizedNotificationMessage-id of localizedNotificationMessage | 
**microsoftGraphLocalizedNotificationMessage** | [**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md)| New navigation property values | 

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

