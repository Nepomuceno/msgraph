# \SecuritySecureScoreControlProfileApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCreateSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityCreateSecureScoreControlProfiles) | **Post** /Security/secureScoreControlProfiles | Create new navigation property to secureScoreControlProfiles for Security
[**SecurityGetSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityGetSecureScoreControlProfiles) | **Get** /Security/secureScoreControlProfiles({secureScoreControlProfile-id}) | Get secureScoreControlProfiles from Security
[**SecurityListSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityListSecureScoreControlProfiles) | **Get** /Security/secureScoreControlProfiles | Get secureScoreControlProfiles from Security
[**SecurityUpdateSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityUpdateSecureScoreControlProfiles) | **Patch** /Security/secureScoreControlProfiles({secureScoreControlProfile-id}) | Update the navigation property secureScoreControlProfiles in Security



## SecurityCreateSecureScoreControlProfiles

> MicrosoftGraphSecureScoreControlProfile SecurityCreateSecureScoreControlProfiles(ctx, microsoftGraphSecureScoreControlProfile)
Create new navigation property to secureScoreControlProfiles for Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftGraphSecureScoreControlProfile** | [**MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md)| New navigation property | 

### Return type

[**MicrosoftGraphSecureScoreControlProfile**](microsoft.graph.secureScoreControlProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGetSecureScoreControlProfiles

> MicrosoftGraphSecureScoreControlProfile SecurityGetSecureScoreControlProfiles(ctx, secureScoreControlProfileId, optional)
Get secureScoreControlProfiles from Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreControlProfileId** | **string**| key: secureScoreControlProfile-id of secureScoreControlProfile | 
 **optional** | ***SecurityGetSecureScoreControlProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityGetSecureScoreControlProfilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Select properties to be returned | 
 **expand** | [**optional.Interface of []string**](string.md)| Expand related entities | 

### Return type

[**MicrosoftGraphSecureScoreControlProfile**](microsoft.graph.secureScoreControlProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityListSecureScoreControlProfiles

> CollectionOfSecureScoreControlProfile SecurityListSecureScoreControlProfiles(ctx, optional)
Get secureScoreControlProfiles from Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecurityListSecureScoreControlProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SecurityListSecureScoreControlProfilesOpts struct


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

[**CollectionOfSecureScoreControlProfile**](Collection of secureScoreControlProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityUpdateSecureScoreControlProfiles

> SecurityUpdateSecureScoreControlProfiles(ctx, secureScoreControlProfileId, microsoftGraphSecureScoreControlProfile)
Update the navigation property secureScoreControlProfiles in Security

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreControlProfileId** | **string**| key: secureScoreControlProfile-id of secureScoreControlProfile | 
**microsoftGraphSecureScoreControlProfile** | [**MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md)| New navigation property values | 

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
