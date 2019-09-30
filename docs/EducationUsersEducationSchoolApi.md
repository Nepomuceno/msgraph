# \EducationUsersEducationSchoolApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationUsersGetSchools**](EducationUsersEducationSchoolApi.md#EducationUsersGetSchools) | **Get** /education/users({educationUser-id})/schools({educationSchool-id}) | Get schools from education
[**EducationUsersListSchools**](EducationUsersEducationSchoolApi.md#EducationUsersListSchools) | **Get** /education/users({educationUser-id})/schools | Get schools from education



## EducationUsersGetSchools

> MicrosoftGraphEducationSchool EducationUsersGetSchools(ctx, educationUserId, educationSchoolId, optional)
Get schools from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string**| key: educationUser-id of educationUser | 
**educationSchoolId** | **string**| key: educationSchool-id of educationSchool | 
 **optional** | ***EducationUsersGetSchoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationUsersGetSchoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationSchool**](microsoft.graph.educationSchool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUsersListSchools

> CollectionOfEducationSchool EducationUsersListSchools(ctx, educationUserId, optional)
Get schools from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string**| key: educationUser-id of educationUser | 
 **optional** | ***EducationUsersListSchoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationUsersListSchoolsOpts struct


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

[**CollectionOfEducationSchool**](Collection of educationSchool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

