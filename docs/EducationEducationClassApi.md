# \EducationEducationClassApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationCreateClasses**](EducationEducationClassApi.md#EducationCreateClasses) | **Post** /education/classes | Create new navigation property to classes for education
[**EducationGetClasses**](EducationEducationClassApi.md#EducationGetClasses) | **Get** /education/classes({educationClass-id}) | Get classes from education
[**EducationListClasses**](EducationEducationClassApi.md#EducationListClasses) | **Get** /education/classes | Get classes from education
[**EducationUpdateClasses**](EducationEducationClassApi.md#EducationUpdateClasses) | **Patch** /education/classes({educationClass-id}) | Update the navigation property classes in education



## EducationCreateClasses

> MicrosoftGraphEducationClass EducationCreateClasses(ctx, microsoftGraphEducationClass)
Create new navigation property to classes for education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphEducationClass** | [**MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md)| New navigation property | 

### Return type

[**MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationGetClasses

> MicrosoftGraphEducationClass EducationGetClasses(ctx, educationClassId, optional)
Get classes from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
 **optional** | ***EducationGetClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationGetClassesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationListClasses

> CollectionOfEducationClass EducationListClasses(ctx, optional)
Get classes from education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EducationListClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EducationListClassesOpts struct


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

[**CollectionOfEducationClass**](Collection of educationClass.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUpdateClasses

> EducationUpdateClasses(ctx, educationClassId, microsoftGraphEducationClass)
Update the navigation property classes in education

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string**| key: educationClass-id of educationClass | 
**microsoftGraphEducationClass** | [**MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md)| New navigation property values | 

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

