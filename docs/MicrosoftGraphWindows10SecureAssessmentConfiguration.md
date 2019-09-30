# MicrosoftGraphWindows10SecureAssessmentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | Pointer to **string** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphDeviceConfigurationAssignment**](microsoft.graph.deviceConfigurationAssignment.md) |  | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationDeviceStatus**](microsoft.graph.deviceConfigurationDeviceStatus.md) |  | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationUserStatus**](microsoft.graph.deviceConfigurationUserStatus.md) |  | [optional] 
**DeviceStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview**](anyOf&lt;microsoft.graph.deviceConfigurationDeviceOverview&gt;.md) |  | [optional] 
**UserStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationUserOverview**](anyOf&lt;microsoft.graph.deviceConfigurationUserOverview&gt;.md) |  | [optional] 
**DeviceSettingStateSummaries** | Pointer to [**[]MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md) |  | [optional] 
**LaunchUri** | Pointer to **string** | Url link to an assessment that&#39;s automatically loaded when the secure assessment browser is launched. It has to be a valid Url (http[s]://msdn.microsoft.com/). | [optional] 
**ConfigurationAccount** | Pointer to **string** | The account used to configure the Windows device for taking the test. The user can be a domain account (domain\\user), an AAD account (username@tenant.com) or a local account (username). | [optional] 
**AllowPrinting** | Pointer to **bool** | Indicates whether or not to allow the app from printing during the test. | [optional] 
**AllowScreenCapture** | Pointer to **bool** | Indicates whether or not to allow screen capture capability during a test. | [optional] 
**AllowTextSuggestion** | Pointer to **bool** | Indicates whether or not to allow text suggestions during the test. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetLaunchUri

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetLaunchUri() string`

GetLaunchUri returns the LaunchUri field if non-nil, zero value otherwise.

### GetLaunchUriOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetLaunchUriOk() (string, bool)`

GetLaunchUriOk returns a tuple with the LaunchUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLaunchUri

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasLaunchUri() bool`

HasLaunchUri returns a boolean if a field has been set.

### SetLaunchUri

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetLaunchUri(v string)`

SetLaunchUri gets a reference to the given string and assigns it to the LaunchUri field.

### SetLaunchUriExplicitNull

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetLaunchUriExplicitNull(b bool)`

SetLaunchUriExplicitNull (un)sets LaunchUri to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LaunchUri value is set to nil even if false is passed
### GetConfigurationAccount

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetConfigurationAccount() string`

GetConfigurationAccount returns the ConfigurationAccount field if non-nil, zero value otherwise.

### GetConfigurationAccountOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetConfigurationAccountOk() (string, bool)`

GetConfigurationAccountOk returns a tuple with the ConfigurationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationAccount

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasConfigurationAccount() bool`

HasConfigurationAccount returns a boolean if a field has been set.

### SetConfigurationAccount

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetConfigurationAccount(v string)`

SetConfigurationAccount gets a reference to the given string and assigns it to the ConfigurationAccount field.

### SetConfigurationAccountExplicitNull

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetConfigurationAccountExplicitNull(b bool)`

SetConfigurationAccountExplicitNull (un)sets ConfigurationAccount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConfigurationAccount value is set to nil even if false is passed
### GetAllowPrinting

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAllowPrinting() bool`

GetAllowPrinting returns the AllowPrinting field if non-nil, zero value otherwise.

### GetAllowPrintingOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAllowPrintingOk() (bool, bool)`

GetAllowPrintingOk returns a tuple with the AllowPrinting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowPrinting

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasAllowPrinting() bool`

HasAllowPrinting returns a boolean if a field has been set.

### SetAllowPrinting

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetAllowPrinting(v bool)`

SetAllowPrinting gets a reference to the given bool and assigns it to the AllowPrinting field.

### GetAllowScreenCapture

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAllowScreenCapture() bool`

GetAllowScreenCapture returns the AllowScreenCapture field if non-nil, zero value otherwise.

### GetAllowScreenCaptureOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAllowScreenCaptureOk() (bool, bool)`

GetAllowScreenCaptureOk returns a tuple with the AllowScreenCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowScreenCapture

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasAllowScreenCapture() bool`

HasAllowScreenCapture returns a boolean if a field has been set.

### SetAllowScreenCapture

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetAllowScreenCapture(v bool)`

SetAllowScreenCapture gets a reference to the given bool and assigns it to the AllowScreenCapture field.

### GetAllowTextSuggestion

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAllowTextSuggestion() bool`

GetAllowTextSuggestion returns the AllowTextSuggestion field if non-nil, zero value otherwise.

### GetAllowTextSuggestionOk

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) GetAllowTextSuggestionOk() (bool, bool)`

GetAllowTextSuggestionOk returns a tuple with the AllowTextSuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowTextSuggestion

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) HasAllowTextSuggestion() bool`

HasAllowTextSuggestion returns a boolean if a field has been set.

### SetAllowTextSuggestion

`func (o *MicrosoftGraphWindows10SecureAssessmentConfiguration) SetAllowTextSuggestion(v bool)`

SetAllowTextSuggestion gets a reference to the given bool and assigns it to the AllowTextSuggestion field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


