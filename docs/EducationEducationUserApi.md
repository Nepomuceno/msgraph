# \EducationEducationUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationCreateUsers**](EducationEducationUserApi.md#EducationCreateUsers) | **Post** /education/users | Create new navigation property to users for education
[**EducationGetMe**](EducationEducationUserApi.md#EducationGetMe) | **Get** /education/me | Get me from education
[**EducationGetUsers**](EducationEducationUserApi.md#EducationGetUsers) | **Get** /education/users({educationUser-id}) | Get users from education
[**EducationListUsers**](EducationEducationUserApi.md#EducationListUsers) | **Get** /education/users | Get users from education
[**EducationUpdateMe**](EducationEducationUserApi.md#EducationUpdateMe) | **Patch** /education/me | Update the navigation property me in education
[**EducationUpdateUsers**](EducationEducationUserApi.md#EducationUpdateUsers) | **Patch** /education/users({educationUser-id}) | Update the navigation property users in education



## EducationCreateUsers

> MicrosoftGraphEducationUser EducationCreateUsers(ctx, microsoftGraphEducationUser)
Create new navigation property to users for education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphEducationUser** | [**MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md)| New navigation property | 

### Return type

[**MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationGetMe

> MicrosoftGraphEducationUser EducationGetMe(ctx, optional)
Get me from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EducationGetMeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationGetMeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationGetUsers

> MicrosoftGraphEducationUser EducationGetUsers(ctx, educationUserId, optional)
Get users from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string**| key: educationUser-id of educationUser | 
 **optional** | ***EducationGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationGetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationListUsers

> CollectionOfEducationUser EducationListUsers(ctx, optional)
Get users from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EducationListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationListUsersOpts struct


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

[**CollectionOfEducationUser**](Collection of educationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUpdateMe

> EducationUpdateMe(ctx, microsoftGraphEducationUser)
Update the navigation property me in education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphEducationUser** | [**MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md)| New navigation property values | 

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


## EducationUpdateUsers

> EducationUpdateUsers(ctx, educationUserId, microsoftGraphEducationUser)
Update the navigation property users in education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string**| key: educationUser-id of educationUser | 
**microsoftGraphEducationUser** | [**MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md)| New navigation property values | 

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

