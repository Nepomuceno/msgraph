# MicrosoftGraphWindows10CustomConfiguration

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
**OmaSettings** | Pointer to [**[]AnyOfmicrosoftGraphOmaSetting**](anyOf&lt;microsoft.graph.omaSetting&gt;.md) | OMA settings. This collection can contain a maximum of 1000 elements. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetOmaSettings

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetOmaSettings() []AnyOfmicrosoftGraphOmaSetting`

GetOmaSettings returns the OmaSettings field if non-nil, zero value otherwise.

### GetOmaSettingsOk

`func (o *MicrosoftGraphWindows10CustomConfiguration) GetOmaSettingsOk() ([]AnyOfmicrosoftGraphOmaSetting, bool)`

GetOmaSettingsOk returns a tuple with the OmaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOmaSettings

`func (o *MicrosoftGraphWindows10CustomConfiguration) HasOmaSettings() bool`

HasOmaSettings returns a boolean if a field has been set.

### SetOmaSettings

`func (o *MicrosoftGraphWindows10CustomConfiguration) SetOmaSettings(v []AnyOfmicrosoftGraphOmaSetting)`

SetOmaSettings gets a reference to the given []AnyOfmicrosoftGraphOmaSetting and assigns it to the OmaSettings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


