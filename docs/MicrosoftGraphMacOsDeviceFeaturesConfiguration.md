# MicrosoftGraphMacOsDeviceFeaturesConfiguration

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

## Methods

### GetId

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphMacOsDeviceFeaturesConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


