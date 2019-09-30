# \MeContactsProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeContactsGetPhoto**](MeContactsProfilePhotoApi.md#MeContactsGetPhoto) | **Get** /me/contacts({contact-id})/photo | Get photo from me
[**MeContactsUpdatePhoto**](MeContactsProfilePhotoApi.md#MeContactsUpdatePhoto) | **Patch** /me/contacts({contact-id})/photo | Update the navigation property photo in me



## MeContactsGetPhoto

> MicrosoftGraphProfilePhoto MeContactsGetPhoto(ctx, contactId, optional)
Get photo from me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string**| key: contact-id of contact | 
 **optional** | ***MeContactsGetPhotoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeContactsGetPhotoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeContactsUpdatePhoto

> MeContactsUpdatePhoto(ctx, contactId, microsoftGraphProfilePhoto)
Update the navigation property photo in me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string**| key: contact-id of contact | 
**microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)| New navigation property values | 

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

