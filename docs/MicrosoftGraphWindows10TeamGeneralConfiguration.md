# MicrosoftGraphWindows10TeamGeneralConfiguration

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
**AzureOperationalInsightsBlockTelemetry** | Pointer to **bool** | Indicates whether or not to Block Azure Operational Insights. | [optional] 
**AzureOperationalInsightsWorkspaceId** | Pointer to **string** | The Azure Operational Insights workspace id. | [optional] 
**AzureOperationalInsightsWorkspaceKey** | Pointer to **string** | The Azure Operational Insights Workspace key. | [optional] 
**ConnectAppBlockAutoLaunch** | Pointer to **bool** | Specifies whether to automatically launch the Connect app whenever a projection is initiated. | [optional] 
**MaintenanceWindowBlocked** | Pointer to **bool** | Indicates whether or not to Block setting a maintenance window for device updates. | [optional] 
**MaintenanceWindowDurationInHours** | Pointer to **int32** | Maintenance window duration for device updates. Valid values 0 to 5 | [optional] 
**MaintenanceWindowStartTime** | Pointer to **string** | Maintenance window start time for device updates. | [optional] 
**MiracastChannel** | Pointer to [**AnyOfmicrosoftGraphMiracastChannel**](anyOf&lt;microsoft.graph.miracastChannel&gt;.md) | The channel. | [optional] 
**MiracastBlocked** | Pointer to **bool** | Indicates whether or not to Block wireless projection. | [optional] 
**MiracastRequirePin** | Pointer to **bool** | Indicates whether or not to require a pin for wireless projection. | [optional] 
**SettingsBlockMyMeetingsAndFiles** | Pointer to **bool** | Specifies whether to disable the \&quot;My meetings and files\&quot; feature in the Start menu, which shows the signed-in user&#39;s meetings and files from Office 365. | [optional] 
**SettingsBlockSessionResume** | Pointer to **bool** | Specifies whether to allow the ability to resume a session when the session times out. | [optional] 
**SettingsBlockSigninSuggestions** | Pointer to **bool** | Specifies whether to disable auto-populating of the sign-in dialog with invitees from scheduled meetings. | [optional] 
**SettingsDefaultVolume** | Pointer to **int32** | Specifies the default volume value for a new session. Permitted values are 0-100. The default is 45. Valid values 0 to 100 | [optional] 
**SettingsScreenTimeoutInMinutes** | Pointer to **int32** | Specifies the number of minutes until the Hub screen turns off. | [optional] 
**SettingsSessionTimeoutInMinutes** | Pointer to **int32** | Specifies the number of minutes until the session times out. | [optional] 
**SettingsSleepTimeoutInMinutes** | Pointer to **int32** | Specifies the number of minutes until the Hub enters sleep mode. | [optional] 
**WelcomeScreenBlockAutomaticWakeUp** | Pointer to **bool** | Indicates whether or not to Block the welcome screen from waking up automatically when someone enters the room. | [optional] 
**WelcomeScreenBackgroundImageUrl** | Pointer to **string** | The welcome screen background image URL. The URL must use the HTTPS protocol and return a PNG image. | [optional] 
**WelcomeScreenMeetingInformation** | Pointer to [**AnyOfmicrosoftGraphWelcomeScreenMeetingInformation**](anyOf&lt;microsoft.graph.welcomeScreenMeetingInformation&gt;.md) | The welcome screen meeting information shown. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAzureOperationalInsightsBlockTelemetry

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAzureOperationalInsightsBlockTelemetry() bool`

GetAzureOperationalInsightsBlockTelemetry returns the AzureOperationalInsightsBlockTelemetry field if non-nil, zero value otherwise.

### GetAzureOperationalInsightsBlockTelemetryOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAzureOperationalInsightsBlockTelemetryOk() (bool, bool)`

GetAzureOperationalInsightsBlockTelemetryOk returns a tuple with the AzureOperationalInsightsBlockTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureOperationalInsightsBlockTelemetry

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasAzureOperationalInsightsBlockTelemetry() bool`

HasAzureOperationalInsightsBlockTelemetry returns a boolean if a field has been set.

### SetAzureOperationalInsightsBlockTelemetry

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetAzureOperationalInsightsBlockTelemetry(v bool)`

SetAzureOperationalInsightsBlockTelemetry gets a reference to the given bool and assigns it to the AzureOperationalInsightsBlockTelemetry field.

### GetAzureOperationalInsightsWorkspaceId

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceId() string`

GetAzureOperationalInsightsWorkspaceId returns the AzureOperationalInsightsWorkspaceId field if non-nil, zero value otherwise.

### GetAzureOperationalInsightsWorkspaceIdOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceIdOk() (string, bool)`

GetAzureOperationalInsightsWorkspaceIdOk returns a tuple with the AzureOperationalInsightsWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureOperationalInsightsWorkspaceId

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasAzureOperationalInsightsWorkspaceId() bool`

HasAzureOperationalInsightsWorkspaceId returns a boolean if a field has been set.

### SetAzureOperationalInsightsWorkspaceId

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceId(v string)`

SetAzureOperationalInsightsWorkspaceId gets a reference to the given string and assigns it to the AzureOperationalInsightsWorkspaceId field.

### SetAzureOperationalInsightsWorkspaceIdExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceIdExplicitNull(b bool)`

SetAzureOperationalInsightsWorkspaceIdExplicitNull (un)sets AzureOperationalInsightsWorkspaceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureOperationalInsightsWorkspaceId value is set to nil even if false is passed
### GetAzureOperationalInsightsWorkspaceKey

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceKey() string`

GetAzureOperationalInsightsWorkspaceKey returns the AzureOperationalInsightsWorkspaceKey field if non-nil, zero value otherwise.

### GetAzureOperationalInsightsWorkspaceKeyOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceKeyOk() (string, bool)`

GetAzureOperationalInsightsWorkspaceKeyOk returns a tuple with the AzureOperationalInsightsWorkspaceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureOperationalInsightsWorkspaceKey

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasAzureOperationalInsightsWorkspaceKey() bool`

HasAzureOperationalInsightsWorkspaceKey returns a boolean if a field has been set.

### SetAzureOperationalInsightsWorkspaceKey

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceKey(v string)`

SetAzureOperationalInsightsWorkspaceKey gets a reference to the given string and assigns it to the AzureOperationalInsightsWorkspaceKey field.

### SetAzureOperationalInsightsWorkspaceKeyExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceKeyExplicitNull(b bool)`

SetAzureOperationalInsightsWorkspaceKeyExplicitNull (un)sets AzureOperationalInsightsWorkspaceKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureOperationalInsightsWorkspaceKey value is set to nil even if false is passed
### GetConnectAppBlockAutoLaunch

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetConnectAppBlockAutoLaunch() bool`

GetConnectAppBlockAutoLaunch returns the ConnectAppBlockAutoLaunch field if non-nil, zero value otherwise.

### GetConnectAppBlockAutoLaunchOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetConnectAppBlockAutoLaunchOk() (bool, bool)`

GetConnectAppBlockAutoLaunchOk returns a tuple with the ConnectAppBlockAutoLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectAppBlockAutoLaunch

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasConnectAppBlockAutoLaunch() bool`

HasConnectAppBlockAutoLaunch returns a boolean if a field has been set.

### SetConnectAppBlockAutoLaunch

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetConnectAppBlockAutoLaunch(v bool)`

SetConnectAppBlockAutoLaunch gets a reference to the given bool and assigns it to the ConnectAppBlockAutoLaunch field.

### GetMaintenanceWindowBlocked

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMaintenanceWindowBlocked() bool`

GetMaintenanceWindowBlocked returns the MaintenanceWindowBlocked field if non-nil, zero value otherwise.

### GetMaintenanceWindowBlockedOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMaintenanceWindowBlockedOk() (bool, bool)`

GetMaintenanceWindowBlockedOk returns a tuple with the MaintenanceWindowBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceWindowBlocked

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasMaintenanceWindowBlocked() bool`

HasMaintenanceWindowBlocked returns a boolean if a field has been set.

### SetMaintenanceWindowBlocked

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMaintenanceWindowBlocked(v bool)`

SetMaintenanceWindowBlocked gets a reference to the given bool and assigns it to the MaintenanceWindowBlocked field.

### GetMaintenanceWindowDurationInHours

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMaintenanceWindowDurationInHours() int32`

GetMaintenanceWindowDurationInHours returns the MaintenanceWindowDurationInHours field if non-nil, zero value otherwise.

### GetMaintenanceWindowDurationInHoursOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMaintenanceWindowDurationInHoursOk() (int32, bool)`

GetMaintenanceWindowDurationInHoursOk returns a tuple with the MaintenanceWindowDurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceWindowDurationInHours

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasMaintenanceWindowDurationInHours() bool`

HasMaintenanceWindowDurationInHours returns a boolean if a field has been set.

### SetMaintenanceWindowDurationInHours

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMaintenanceWindowDurationInHours(v int32)`

SetMaintenanceWindowDurationInHours gets a reference to the given int32 and assigns it to the MaintenanceWindowDurationInHours field.

### SetMaintenanceWindowDurationInHoursExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMaintenanceWindowDurationInHoursExplicitNull(b bool)`

SetMaintenanceWindowDurationInHoursExplicitNull (un)sets MaintenanceWindowDurationInHours to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaintenanceWindowDurationInHours value is set to nil even if false is passed
### GetMaintenanceWindowStartTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMaintenanceWindowStartTime() string`

GetMaintenanceWindowStartTime returns the MaintenanceWindowStartTime field if non-nil, zero value otherwise.

### GetMaintenanceWindowStartTimeOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMaintenanceWindowStartTimeOk() (string, bool)`

GetMaintenanceWindowStartTimeOk returns a tuple with the MaintenanceWindowStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceWindowStartTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasMaintenanceWindowStartTime() bool`

HasMaintenanceWindowStartTime returns a boolean if a field has been set.

### SetMaintenanceWindowStartTime

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMaintenanceWindowStartTime(v string)`

SetMaintenanceWindowStartTime gets a reference to the given string and assigns it to the MaintenanceWindowStartTime field.

### SetMaintenanceWindowStartTimeExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMaintenanceWindowStartTimeExplicitNull(b bool)`

SetMaintenanceWindowStartTimeExplicitNull (un)sets MaintenanceWindowStartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaintenanceWindowStartTime value is set to nil even if false is passed
### GetMiracastChannel

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMiracastChannel() AnyOfmicrosoftGraphMiracastChannel`

GetMiracastChannel returns the MiracastChannel field if non-nil, zero value otherwise.

### GetMiracastChannelOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMiracastChannelOk() (AnyOfmicrosoftGraphMiracastChannel, bool)`

GetMiracastChannelOk returns a tuple with the MiracastChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiracastChannel

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasMiracastChannel() bool`

HasMiracastChannel returns a boolean if a field has been set.

### SetMiracastChannel

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMiracastChannel(v AnyOfmicrosoftGraphMiracastChannel)`

SetMiracastChannel gets a reference to the given AnyOfmicrosoftGraphMiracastChannel and assigns it to the MiracastChannel field.

### GetMiracastBlocked

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMiracastBlocked() bool`

GetMiracastBlocked returns the MiracastBlocked field if non-nil, zero value otherwise.

### GetMiracastBlockedOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMiracastBlockedOk() (bool, bool)`

GetMiracastBlockedOk returns a tuple with the MiracastBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiracastBlocked

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasMiracastBlocked() bool`

HasMiracastBlocked returns a boolean if a field has been set.

### SetMiracastBlocked

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMiracastBlocked(v bool)`

SetMiracastBlocked gets a reference to the given bool and assigns it to the MiracastBlocked field.

### GetMiracastRequirePin

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMiracastRequirePin() bool`

GetMiracastRequirePin returns the MiracastRequirePin field if non-nil, zero value otherwise.

### GetMiracastRequirePinOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetMiracastRequirePinOk() (bool, bool)`

GetMiracastRequirePinOk returns a tuple with the MiracastRequirePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiracastRequirePin

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasMiracastRequirePin() bool`

HasMiracastRequirePin returns a boolean if a field has been set.

### SetMiracastRequirePin

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetMiracastRequirePin(v bool)`

SetMiracastRequirePin gets a reference to the given bool and assigns it to the MiracastRequirePin field.

### GetSettingsBlockMyMeetingsAndFiles

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsBlockMyMeetingsAndFiles() bool`

GetSettingsBlockMyMeetingsAndFiles returns the SettingsBlockMyMeetingsAndFiles field if non-nil, zero value otherwise.

### GetSettingsBlockMyMeetingsAndFilesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsBlockMyMeetingsAndFilesOk() (bool, bool)`

GetSettingsBlockMyMeetingsAndFilesOk returns a tuple with the SettingsBlockMyMeetingsAndFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockMyMeetingsAndFiles

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsBlockMyMeetingsAndFiles() bool`

HasSettingsBlockMyMeetingsAndFiles returns a boolean if a field has been set.

### SetSettingsBlockMyMeetingsAndFiles

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsBlockMyMeetingsAndFiles(v bool)`

SetSettingsBlockMyMeetingsAndFiles gets a reference to the given bool and assigns it to the SettingsBlockMyMeetingsAndFiles field.

### GetSettingsBlockSessionResume

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsBlockSessionResume() bool`

GetSettingsBlockSessionResume returns the SettingsBlockSessionResume field if non-nil, zero value otherwise.

### GetSettingsBlockSessionResumeOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsBlockSessionResumeOk() (bool, bool)`

GetSettingsBlockSessionResumeOk returns a tuple with the SettingsBlockSessionResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSessionResume

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsBlockSessionResume() bool`

HasSettingsBlockSessionResume returns a boolean if a field has been set.

### SetSettingsBlockSessionResume

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsBlockSessionResume(v bool)`

SetSettingsBlockSessionResume gets a reference to the given bool and assigns it to the SettingsBlockSessionResume field.

### GetSettingsBlockSigninSuggestions

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsBlockSigninSuggestions() bool`

GetSettingsBlockSigninSuggestions returns the SettingsBlockSigninSuggestions field if non-nil, zero value otherwise.

### GetSettingsBlockSigninSuggestionsOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsBlockSigninSuggestionsOk() (bool, bool)`

GetSettingsBlockSigninSuggestionsOk returns a tuple with the SettingsBlockSigninSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSigninSuggestions

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsBlockSigninSuggestions() bool`

HasSettingsBlockSigninSuggestions returns a boolean if a field has been set.

### SetSettingsBlockSigninSuggestions

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsBlockSigninSuggestions(v bool)`

SetSettingsBlockSigninSuggestions gets a reference to the given bool and assigns it to the SettingsBlockSigninSuggestions field.

### GetSettingsDefaultVolume

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsDefaultVolume() int32`

GetSettingsDefaultVolume returns the SettingsDefaultVolume field if non-nil, zero value otherwise.

### GetSettingsDefaultVolumeOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsDefaultVolumeOk() (int32, bool)`

GetSettingsDefaultVolumeOk returns a tuple with the SettingsDefaultVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsDefaultVolume

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsDefaultVolume() bool`

HasSettingsDefaultVolume returns a boolean if a field has been set.

### SetSettingsDefaultVolume

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsDefaultVolume(v int32)`

SetSettingsDefaultVolume gets a reference to the given int32 and assigns it to the SettingsDefaultVolume field.

### SetSettingsDefaultVolumeExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsDefaultVolumeExplicitNull(b bool)`

SetSettingsDefaultVolumeExplicitNull (un)sets SettingsDefaultVolume to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsDefaultVolume value is set to nil even if false is passed
### GetSettingsScreenTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsScreenTimeoutInMinutes() int32`

GetSettingsScreenTimeoutInMinutes returns the SettingsScreenTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSettingsScreenTimeoutInMinutesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsScreenTimeoutInMinutesOk() (int32, bool)`

GetSettingsScreenTimeoutInMinutesOk returns a tuple with the SettingsScreenTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsScreenTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsScreenTimeoutInMinutes() bool`

HasSettingsScreenTimeoutInMinutes returns a boolean if a field has been set.

### SetSettingsScreenTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsScreenTimeoutInMinutes(v int32)`

SetSettingsScreenTimeoutInMinutes gets a reference to the given int32 and assigns it to the SettingsScreenTimeoutInMinutes field.

### SetSettingsScreenTimeoutInMinutesExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsScreenTimeoutInMinutesExplicitNull(b bool)`

SetSettingsScreenTimeoutInMinutesExplicitNull (un)sets SettingsScreenTimeoutInMinutes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsScreenTimeoutInMinutes value is set to nil even if false is passed
### GetSettingsSessionTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsSessionTimeoutInMinutes() int32`

GetSettingsSessionTimeoutInMinutes returns the SettingsSessionTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSettingsSessionTimeoutInMinutesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsSessionTimeoutInMinutesOk() (int32, bool)`

GetSettingsSessionTimeoutInMinutesOk returns a tuple with the SettingsSessionTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsSessionTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsSessionTimeoutInMinutes() bool`

HasSettingsSessionTimeoutInMinutes returns a boolean if a field has been set.

### SetSettingsSessionTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsSessionTimeoutInMinutes(v int32)`

SetSettingsSessionTimeoutInMinutes gets a reference to the given int32 and assigns it to the SettingsSessionTimeoutInMinutes field.

### SetSettingsSessionTimeoutInMinutesExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsSessionTimeoutInMinutesExplicitNull(b bool)`

SetSettingsSessionTimeoutInMinutesExplicitNull (un)sets SettingsSessionTimeoutInMinutes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsSessionTimeoutInMinutes value is set to nil even if false is passed
### GetSettingsSleepTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsSleepTimeoutInMinutes() int32`

GetSettingsSleepTimeoutInMinutes returns the SettingsSleepTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSettingsSleepTimeoutInMinutesOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetSettingsSleepTimeoutInMinutesOk() (int32, bool)`

GetSettingsSleepTimeoutInMinutesOk returns a tuple with the SettingsSleepTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsSleepTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasSettingsSleepTimeoutInMinutes() bool`

HasSettingsSleepTimeoutInMinutes returns a boolean if a field has been set.

### SetSettingsSleepTimeoutInMinutes

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsSleepTimeoutInMinutes(v int32)`

SetSettingsSleepTimeoutInMinutes gets a reference to the given int32 and assigns it to the SettingsSleepTimeoutInMinutes field.

### SetSettingsSleepTimeoutInMinutesExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetSettingsSleepTimeoutInMinutesExplicitNull(b bool)`

SetSettingsSleepTimeoutInMinutesExplicitNull (un)sets SettingsSleepTimeoutInMinutes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsSleepTimeoutInMinutes value is set to nil even if false is passed
### GetWelcomeScreenBlockAutomaticWakeUp

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetWelcomeScreenBlockAutomaticWakeUp() bool`

GetWelcomeScreenBlockAutomaticWakeUp returns the WelcomeScreenBlockAutomaticWakeUp field if non-nil, zero value otherwise.

### GetWelcomeScreenBlockAutomaticWakeUpOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetWelcomeScreenBlockAutomaticWakeUpOk() (bool, bool)`

GetWelcomeScreenBlockAutomaticWakeUpOk returns a tuple with the WelcomeScreenBlockAutomaticWakeUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWelcomeScreenBlockAutomaticWakeUp

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasWelcomeScreenBlockAutomaticWakeUp() bool`

HasWelcomeScreenBlockAutomaticWakeUp returns a boolean if a field has been set.

### SetWelcomeScreenBlockAutomaticWakeUp

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetWelcomeScreenBlockAutomaticWakeUp(v bool)`

SetWelcomeScreenBlockAutomaticWakeUp gets a reference to the given bool and assigns it to the WelcomeScreenBlockAutomaticWakeUp field.

### GetWelcomeScreenBackgroundImageUrl

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetWelcomeScreenBackgroundImageUrl() string`

GetWelcomeScreenBackgroundImageUrl returns the WelcomeScreenBackgroundImageUrl field if non-nil, zero value otherwise.

### GetWelcomeScreenBackgroundImageUrlOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetWelcomeScreenBackgroundImageUrlOk() (string, bool)`

GetWelcomeScreenBackgroundImageUrlOk returns a tuple with the WelcomeScreenBackgroundImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWelcomeScreenBackgroundImageUrl

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasWelcomeScreenBackgroundImageUrl() bool`

HasWelcomeScreenBackgroundImageUrl returns a boolean if a field has been set.

### SetWelcomeScreenBackgroundImageUrl

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetWelcomeScreenBackgroundImageUrl(v string)`

SetWelcomeScreenBackgroundImageUrl gets a reference to the given string and assigns it to the WelcomeScreenBackgroundImageUrl field.

### SetWelcomeScreenBackgroundImageUrlExplicitNull

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetWelcomeScreenBackgroundImageUrlExplicitNull(b bool)`

SetWelcomeScreenBackgroundImageUrlExplicitNull (un)sets WelcomeScreenBackgroundImageUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WelcomeScreenBackgroundImageUrl value is set to nil even if false is passed
### GetWelcomeScreenMeetingInformation

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetWelcomeScreenMeetingInformation() AnyOfmicrosoftGraphWelcomeScreenMeetingInformation`

GetWelcomeScreenMeetingInformation returns the WelcomeScreenMeetingInformation field if non-nil, zero value otherwise.

### GetWelcomeScreenMeetingInformationOk

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) GetWelcomeScreenMeetingInformationOk() (AnyOfmicrosoftGraphWelcomeScreenMeetingInformation, bool)`

GetWelcomeScreenMeetingInformationOk returns a tuple with the WelcomeScreenMeetingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWelcomeScreenMeetingInformation

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) HasWelcomeScreenMeetingInformation() bool`

HasWelcomeScreenMeetingInformation returns a boolean if a field has been set.

### SetWelcomeScreenMeetingInformation

`func (o *MicrosoftGraphWindows10TeamGeneralConfiguration) SetWelcomeScreenMeetingInformation(v AnyOfmicrosoftGraphWelcomeScreenMeetingInformation)`

SetWelcomeScreenMeetingInformation gets a reference to the given AnyOfmicrosoftGraphWelcomeScreenMeetingInformation and assigns it to the WelcomeScreenMeetingInformation field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


