# \UsersOutlookUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGetOutlook**](UsersOutlookUserApi.md#UsersGetOutlook) | **Get** /users({user-id})/outlook | Get outlook from users
[**UsersUpdateOutlook**](UsersOutlookUserApi.md#UsersUpdateOutlook) | **Patch** /users({user-id})/outlook | Update the navigation property outlook in users



## UsersGetOutlook

> MicrosoftGraphOutlookUser UsersGetOutlook(ctx, userId, optional)
Get outlook from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
 **optional** | ***UsersGetOutlookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetOutlookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookUser**](microsoft.graph.outlookUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateOutlook

> UsersUpdateOutlook(ctx, userId, microsoftGraphOutlookUser)
Update the navigation property outlook in users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| key: user-id of user | 
**microsoftGraphOutlookUser** | [**MicrosoftGraphOutlookUser**](MicrosoftGraphOutlookUser.md)| New navigation property values | 

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

