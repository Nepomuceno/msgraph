# Windows10SecureAssessmentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaunchUri** | Pointer to **string** | Url link to an assessment that&#39;s automatically loaded when the secure assessment browser is launched. It has to be a valid Url (http[s]://msdn.microsoft.com/). | [optional] 
**ConfigurationAccount** | Pointer to **string** | The account used to configure the Windows device for taking the test. The user can be a domain account (domain\\user), an AAD account (username@tenant.com) or a local account (username). | [optional] 
**AllowPrinting** | Pointer to **bool** | Indicates whether or not to allow the app from printing during the test. | [optional] 
**AllowScreenCapture** | Pointer to **bool** | Indicates whether or not to allow screen capture capability during a test. | [optional] 
**AllowTextSuggestion** | Pointer to **bool** | Indicates whether or not to allow text suggestions during the test. | [optional] 

## Methods

### GetLaunchUri

`func (o *Windows10SecureAssessmentConfiguration) GetLaunchUri() string`

GetLaunchUri returns the LaunchUri field if non-nil, zero value otherwise.

### GetLaunchUriOk

`func (o *Windows10SecureAssessmentConfiguration) GetLaunchUriOk() (string, bool)`

GetLaunchUriOk returns a tuple with the LaunchUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLaunchUri

`func (o *Windows10SecureAssessmentConfiguration) HasLaunchUri() bool`

HasLaunchUri returns a boolean if a field has been set.

### SetLaunchUri

`func (o *Windows10SecureAssessmentConfiguration) SetLaunchUri(v string)`

SetLaunchUri gets a reference to the given string and assigns it to the LaunchUri field.

### SetLaunchUriExplicitNull

`func (o *Windows10SecureAssessmentConfiguration) SetLaunchUriExplicitNull(b bool)`

SetLaunchUriExplicitNull (un)sets LaunchUri to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LaunchUri value is set to nil even if false is passed
### GetConfigurationAccount

`func (o *Windows10SecureAssessmentConfiguration) GetConfigurationAccount() string`

GetConfigurationAccount returns the ConfigurationAccount field if non-nil, zero value otherwise.

### GetConfigurationAccountOk

`func (o *Windows10SecureAssessmentConfiguration) GetConfigurationAccountOk() (string, bool)`

GetConfigurationAccountOk returns a tuple with the ConfigurationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationAccount

`func (o *Windows10SecureAssessmentConfiguration) HasConfigurationAccount() bool`

HasConfigurationAccount returns a boolean if a field has been set.

### SetConfigurationAccount

`func (o *Windows10SecureAssessmentConfiguration) SetConfigurationAccount(v string)`

SetConfigurationAccount gets a reference to the given string and assigns it to the ConfigurationAccount field.

### SetConfigurationAccountExplicitNull

`func (o *Windows10SecureAssessmentConfiguration) SetConfigurationAccountExplicitNull(b bool)`

SetConfigurationAccountExplicitNull (un)sets ConfigurationAccount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConfigurationAccount value is set to nil even if false is passed
### GetAllowPrinting

`func (o *Windows10SecureAssessmentConfiguration) GetAllowPrinting() bool`

GetAllowPrinting returns the AllowPrinting field if non-nil, zero value otherwise.

### GetAllowPrintingOk

`func (o *Windows10SecureAssessmentConfiguration) GetAllowPrintingOk() (bool, bool)`

GetAllowPrintingOk returns a tuple with the AllowPrinting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowPrinting

`func (o *Windows10SecureAssessmentConfiguration) HasAllowPrinting() bool`

HasAllowPrinting returns a boolean if a field has been set.

### SetAllowPrinting

`func (o *Windows10SecureAssessmentConfiguration) SetAllowPrinting(v bool)`

SetAllowPrinting gets a reference to the given bool and assigns it to the AllowPrinting field.

### GetAllowScreenCapture

`func (o *Windows10SecureAssessmentConfiguration) GetAllowScreenCapture() bool`

GetAllowScreenCapture returns the AllowScreenCapture field if non-nil, zero value otherwise.

### GetAllowScreenCaptureOk

`func (o *Windows10SecureAssessmentConfiguration) GetAllowScreenCaptureOk() (bool, bool)`

GetAllowScreenCaptureOk returns a tuple with the AllowScreenCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowScreenCapture

`func (o *Windows10SecureAssessmentConfiguration) HasAllowScreenCapture() bool`

HasAllowScreenCapture returns a boolean if a field has been set.

### SetAllowScreenCapture

`func (o *Windows10SecureAssessmentConfiguration) SetAllowScreenCapture(v bool)`

SetAllowScreenCapture gets a reference to the given bool and assigns it to the AllowScreenCapture field.

### GetAllowTextSuggestion

`func (o *Windows10SecureAssessmentConfiguration) GetAllowTextSuggestion() bool`

GetAllowTextSuggestion returns the AllowTextSuggestion field if non-nil, zero value otherwise.

### GetAllowTextSuggestionOk

`func (o *Windows10SecureAssessmentConfiguration) GetAllowTextSuggestionOk() (bool, bool)`

GetAllowTextSuggestionOk returns a tuple with the AllowTextSuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowTextSuggestion

`func (o *Windows10SecureAssessmentConfiguration) HasAllowTextSuggestion() bool`

HasAllowTextSuggestion returns a boolean if a field has been set.

### SetAllowTextSuggestion

`func (o *Windows10SecureAssessmentConfiguration) SetAllowTextSuggestion(v bool)`

SetAllowTextSuggestion gets a reference to the given bool and assigns it to the AllowTextSuggestion field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


