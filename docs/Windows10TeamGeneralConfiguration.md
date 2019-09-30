# Windows10TeamGeneralConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAzureOperationalInsightsBlockTelemetry

`func (o *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsBlockTelemetry() bool`

GetAzureOperationalInsightsBlockTelemetry returns the AzureOperationalInsightsBlockTelemetry field if non-nil, zero value otherwise.

### GetAzureOperationalInsightsBlockTelemetryOk

`func (o *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsBlockTelemetryOk() (bool, bool)`

GetAzureOperationalInsightsBlockTelemetryOk returns a tuple with the AzureOperationalInsightsBlockTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureOperationalInsightsBlockTelemetry

`func (o *Windows10TeamGeneralConfiguration) HasAzureOperationalInsightsBlockTelemetry() bool`

HasAzureOperationalInsightsBlockTelemetry returns a boolean if a field has been set.

### SetAzureOperationalInsightsBlockTelemetry

`func (o *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsBlockTelemetry(v bool)`

SetAzureOperationalInsightsBlockTelemetry gets a reference to the given bool and assigns it to the AzureOperationalInsightsBlockTelemetry field.

### GetAzureOperationalInsightsWorkspaceId

`func (o *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceId() string`

GetAzureOperationalInsightsWorkspaceId returns the AzureOperationalInsightsWorkspaceId field if non-nil, zero value otherwise.

### GetAzureOperationalInsightsWorkspaceIdOk

`func (o *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceIdOk() (string, bool)`

GetAzureOperationalInsightsWorkspaceIdOk returns a tuple with the AzureOperationalInsightsWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureOperationalInsightsWorkspaceId

`func (o *Windows10TeamGeneralConfiguration) HasAzureOperationalInsightsWorkspaceId() bool`

HasAzureOperationalInsightsWorkspaceId returns a boolean if a field has been set.

### SetAzureOperationalInsightsWorkspaceId

`func (o *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceId(v string)`

SetAzureOperationalInsightsWorkspaceId gets a reference to the given string and assigns it to the AzureOperationalInsightsWorkspaceId field.

### SetAzureOperationalInsightsWorkspaceIdExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceIdExplicitNull(b bool)`

SetAzureOperationalInsightsWorkspaceIdExplicitNull (un)sets AzureOperationalInsightsWorkspaceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureOperationalInsightsWorkspaceId value is set to nil even if false is passed
### GetAzureOperationalInsightsWorkspaceKey

`func (o *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceKey() string`

GetAzureOperationalInsightsWorkspaceKey returns the AzureOperationalInsightsWorkspaceKey field if non-nil, zero value otherwise.

### GetAzureOperationalInsightsWorkspaceKeyOk

`func (o *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceKeyOk() (string, bool)`

GetAzureOperationalInsightsWorkspaceKeyOk returns a tuple with the AzureOperationalInsightsWorkspaceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureOperationalInsightsWorkspaceKey

`func (o *Windows10TeamGeneralConfiguration) HasAzureOperationalInsightsWorkspaceKey() bool`

HasAzureOperationalInsightsWorkspaceKey returns a boolean if a field has been set.

### SetAzureOperationalInsightsWorkspaceKey

`func (o *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceKey(v string)`

SetAzureOperationalInsightsWorkspaceKey gets a reference to the given string and assigns it to the AzureOperationalInsightsWorkspaceKey field.

### SetAzureOperationalInsightsWorkspaceKeyExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceKeyExplicitNull(b bool)`

SetAzureOperationalInsightsWorkspaceKeyExplicitNull (un)sets AzureOperationalInsightsWorkspaceKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureOperationalInsightsWorkspaceKey value is set to nil even if false is passed
### GetConnectAppBlockAutoLaunch

`func (o *Windows10TeamGeneralConfiguration) GetConnectAppBlockAutoLaunch() bool`

GetConnectAppBlockAutoLaunch returns the ConnectAppBlockAutoLaunch field if non-nil, zero value otherwise.

### GetConnectAppBlockAutoLaunchOk

`func (o *Windows10TeamGeneralConfiguration) GetConnectAppBlockAutoLaunchOk() (bool, bool)`

GetConnectAppBlockAutoLaunchOk returns a tuple with the ConnectAppBlockAutoLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectAppBlockAutoLaunch

`func (o *Windows10TeamGeneralConfiguration) HasConnectAppBlockAutoLaunch() bool`

HasConnectAppBlockAutoLaunch returns a boolean if a field has been set.

### SetConnectAppBlockAutoLaunch

`func (o *Windows10TeamGeneralConfiguration) SetConnectAppBlockAutoLaunch(v bool)`

SetConnectAppBlockAutoLaunch gets a reference to the given bool and assigns it to the ConnectAppBlockAutoLaunch field.

### GetMaintenanceWindowBlocked

`func (o *Windows10TeamGeneralConfiguration) GetMaintenanceWindowBlocked() bool`

GetMaintenanceWindowBlocked returns the MaintenanceWindowBlocked field if non-nil, zero value otherwise.

### GetMaintenanceWindowBlockedOk

`func (o *Windows10TeamGeneralConfiguration) GetMaintenanceWindowBlockedOk() (bool, bool)`

GetMaintenanceWindowBlockedOk returns a tuple with the MaintenanceWindowBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceWindowBlocked

`func (o *Windows10TeamGeneralConfiguration) HasMaintenanceWindowBlocked() bool`

HasMaintenanceWindowBlocked returns a boolean if a field has been set.

### SetMaintenanceWindowBlocked

`func (o *Windows10TeamGeneralConfiguration) SetMaintenanceWindowBlocked(v bool)`

SetMaintenanceWindowBlocked gets a reference to the given bool and assigns it to the MaintenanceWindowBlocked field.

### GetMaintenanceWindowDurationInHours

`func (o *Windows10TeamGeneralConfiguration) GetMaintenanceWindowDurationInHours() int32`

GetMaintenanceWindowDurationInHours returns the MaintenanceWindowDurationInHours field if non-nil, zero value otherwise.

### GetMaintenanceWindowDurationInHoursOk

`func (o *Windows10TeamGeneralConfiguration) GetMaintenanceWindowDurationInHoursOk() (int32, bool)`

GetMaintenanceWindowDurationInHoursOk returns a tuple with the MaintenanceWindowDurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceWindowDurationInHours

`func (o *Windows10TeamGeneralConfiguration) HasMaintenanceWindowDurationInHours() bool`

HasMaintenanceWindowDurationInHours returns a boolean if a field has been set.

### SetMaintenanceWindowDurationInHours

`func (o *Windows10TeamGeneralConfiguration) SetMaintenanceWindowDurationInHours(v int32)`

SetMaintenanceWindowDurationInHours gets a reference to the given int32 and assigns it to the MaintenanceWindowDurationInHours field.

### SetMaintenanceWindowDurationInHoursExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetMaintenanceWindowDurationInHoursExplicitNull(b bool)`

SetMaintenanceWindowDurationInHoursExplicitNull (un)sets MaintenanceWindowDurationInHours to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaintenanceWindowDurationInHours value is set to nil even if false is passed
### GetMaintenanceWindowStartTime

`func (o *Windows10TeamGeneralConfiguration) GetMaintenanceWindowStartTime() string`

GetMaintenanceWindowStartTime returns the MaintenanceWindowStartTime field if non-nil, zero value otherwise.

### GetMaintenanceWindowStartTimeOk

`func (o *Windows10TeamGeneralConfiguration) GetMaintenanceWindowStartTimeOk() (string, bool)`

GetMaintenanceWindowStartTimeOk returns a tuple with the MaintenanceWindowStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceWindowStartTime

`func (o *Windows10TeamGeneralConfiguration) HasMaintenanceWindowStartTime() bool`

HasMaintenanceWindowStartTime returns a boolean if a field has been set.

### SetMaintenanceWindowStartTime

`func (o *Windows10TeamGeneralConfiguration) SetMaintenanceWindowStartTime(v string)`

SetMaintenanceWindowStartTime gets a reference to the given string and assigns it to the MaintenanceWindowStartTime field.

### SetMaintenanceWindowStartTimeExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetMaintenanceWindowStartTimeExplicitNull(b bool)`

SetMaintenanceWindowStartTimeExplicitNull (un)sets MaintenanceWindowStartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaintenanceWindowStartTime value is set to nil even if false is passed
### GetMiracastChannel

`func (o *Windows10TeamGeneralConfiguration) GetMiracastChannel() AnyOfmicrosoftGraphMiracastChannel`

GetMiracastChannel returns the MiracastChannel field if non-nil, zero value otherwise.

### GetMiracastChannelOk

`func (o *Windows10TeamGeneralConfiguration) GetMiracastChannelOk() (AnyOfmicrosoftGraphMiracastChannel, bool)`

GetMiracastChannelOk returns a tuple with the MiracastChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiracastChannel

`func (o *Windows10TeamGeneralConfiguration) HasMiracastChannel() bool`

HasMiracastChannel returns a boolean if a field has been set.

### SetMiracastChannel

`func (o *Windows10TeamGeneralConfiguration) SetMiracastChannel(v AnyOfmicrosoftGraphMiracastChannel)`

SetMiracastChannel gets a reference to the given AnyOfmicrosoftGraphMiracastChannel and assigns it to the MiracastChannel field.

### GetMiracastBlocked

`func (o *Windows10TeamGeneralConfiguration) GetMiracastBlocked() bool`

GetMiracastBlocked returns the MiracastBlocked field if non-nil, zero value otherwise.

### GetMiracastBlockedOk

`func (o *Windows10TeamGeneralConfiguration) GetMiracastBlockedOk() (bool, bool)`

GetMiracastBlockedOk returns a tuple with the MiracastBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiracastBlocked

`func (o *Windows10TeamGeneralConfiguration) HasMiracastBlocked() bool`

HasMiracastBlocked returns a boolean if a field has been set.

### SetMiracastBlocked

`func (o *Windows10TeamGeneralConfiguration) SetMiracastBlocked(v bool)`

SetMiracastBlocked gets a reference to the given bool and assigns it to the MiracastBlocked field.

### GetMiracastRequirePin

`func (o *Windows10TeamGeneralConfiguration) GetMiracastRequirePin() bool`

GetMiracastRequirePin returns the MiracastRequirePin field if non-nil, zero value otherwise.

### GetMiracastRequirePinOk

`func (o *Windows10TeamGeneralConfiguration) GetMiracastRequirePinOk() (bool, bool)`

GetMiracastRequirePinOk returns a tuple with the MiracastRequirePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiracastRequirePin

`func (o *Windows10TeamGeneralConfiguration) HasMiracastRequirePin() bool`

HasMiracastRequirePin returns a boolean if a field has been set.

### SetMiracastRequirePin

`func (o *Windows10TeamGeneralConfiguration) SetMiracastRequirePin(v bool)`

SetMiracastRequirePin gets a reference to the given bool and assigns it to the MiracastRequirePin field.

### GetSettingsBlockMyMeetingsAndFiles

`func (o *Windows10TeamGeneralConfiguration) GetSettingsBlockMyMeetingsAndFiles() bool`

GetSettingsBlockMyMeetingsAndFiles returns the SettingsBlockMyMeetingsAndFiles field if non-nil, zero value otherwise.

### GetSettingsBlockMyMeetingsAndFilesOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsBlockMyMeetingsAndFilesOk() (bool, bool)`

GetSettingsBlockMyMeetingsAndFilesOk returns a tuple with the SettingsBlockMyMeetingsAndFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockMyMeetingsAndFiles

`func (o *Windows10TeamGeneralConfiguration) HasSettingsBlockMyMeetingsAndFiles() bool`

HasSettingsBlockMyMeetingsAndFiles returns a boolean if a field has been set.

### SetSettingsBlockMyMeetingsAndFiles

`func (o *Windows10TeamGeneralConfiguration) SetSettingsBlockMyMeetingsAndFiles(v bool)`

SetSettingsBlockMyMeetingsAndFiles gets a reference to the given bool and assigns it to the SettingsBlockMyMeetingsAndFiles field.

### GetSettingsBlockSessionResume

`func (o *Windows10TeamGeneralConfiguration) GetSettingsBlockSessionResume() bool`

GetSettingsBlockSessionResume returns the SettingsBlockSessionResume field if non-nil, zero value otherwise.

### GetSettingsBlockSessionResumeOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsBlockSessionResumeOk() (bool, bool)`

GetSettingsBlockSessionResumeOk returns a tuple with the SettingsBlockSessionResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSessionResume

`func (o *Windows10TeamGeneralConfiguration) HasSettingsBlockSessionResume() bool`

HasSettingsBlockSessionResume returns a boolean if a field has been set.

### SetSettingsBlockSessionResume

`func (o *Windows10TeamGeneralConfiguration) SetSettingsBlockSessionResume(v bool)`

SetSettingsBlockSessionResume gets a reference to the given bool and assigns it to the SettingsBlockSessionResume field.

### GetSettingsBlockSigninSuggestions

`func (o *Windows10TeamGeneralConfiguration) GetSettingsBlockSigninSuggestions() bool`

GetSettingsBlockSigninSuggestions returns the SettingsBlockSigninSuggestions field if non-nil, zero value otherwise.

### GetSettingsBlockSigninSuggestionsOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsBlockSigninSuggestionsOk() (bool, bool)`

GetSettingsBlockSigninSuggestionsOk returns a tuple with the SettingsBlockSigninSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSigninSuggestions

`func (o *Windows10TeamGeneralConfiguration) HasSettingsBlockSigninSuggestions() bool`

HasSettingsBlockSigninSuggestions returns a boolean if a field has been set.

### SetSettingsBlockSigninSuggestions

`func (o *Windows10TeamGeneralConfiguration) SetSettingsBlockSigninSuggestions(v bool)`

SetSettingsBlockSigninSuggestions gets a reference to the given bool and assigns it to the SettingsBlockSigninSuggestions field.

### GetSettingsDefaultVolume

`func (o *Windows10TeamGeneralConfiguration) GetSettingsDefaultVolume() int32`

GetSettingsDefaultVolume returns the SettingsDefaultVolume field if non-nil, zero value otherwise.

### GetSettingsDefaultVolumeOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsDefaultVolumeOk() (int32, bool)`

GetSettingsDefaultVolumeOk returns a tuple with the SettingsDefaultVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsDefaultVolume

`func (o *Windows10TeamGeneralConfiguration) HasSettingsDefaultVolume() bool`

HasSettingsDefaultVolume returns a boolean if a field has been set.

### SetSettingsDefaultVolume

`func (o *Windows10TeamGeneralConfiguration) SetSettingsDefaultVolume(v int32)`

SetSettingsDefaultVolume gets a reference to the given int32 and assigns it to the SettingsDefaultVolume field.

### SetSettingsDefaultVolumeExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetSettingsDefaultVolumeExplicitNull(b bool)`

SetSettingsDefaultVolumeExplicitNull (un)sets SettingsDefaultVolume to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsDefaultVolume value is set to nil even if false is passed
### GetSettingsScreenTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) GetSettingsScreenTimeoutInMinutes() int32`

GetSettingsScreenTimeoutInMinutes returns the SettingsScreenTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSettingsScreenTimeoutInMinutesOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsScreenTimeoutInMinutesOk() (int32, bool)`

GetSettingsScreenTimeoutInMinutesOk returns a tuple with the SettingsScreenTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsScreenTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) HasSettingsScreenTimeoutInMinutes() bool`

HasSettingsScreenTimeoutInMinutes returns a boolean if a field has been set.

### SetSettingsScreenTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) SetSettingsScreenTimeoutInMinutes(v int32)`

SetSettingsScreenTimeoutInMinutes gets a reference to the given int32 and assigns it to the SettingsScreenTimeoutInMinutes field.

### SetSettingsScreenTimeoutInMinutesExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetSettingsScreenTimeoutInMinutesExplicitNull(b bool)`

SetSettingsScreenTimeoutInMinutesExplicitNull (un)sets SettingsScreenTimeoutInMinutes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsScreenTimeoutInMinutes value is set to nil even if false is passed
### GetSettingsSessionTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) GetSettingsSessionTimeoutInMinutes() int32`

GetSettingsSessionTimeoutInMinutes returns the SettingsSessionTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSettingsSessionTimeoutInMinutesOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsSessionTimeoutInMinutesOk() (int32, bool)`

GetSettingsSessionTimeoutInMinutesOk returns a tuple with the SettingsSessionTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsSessionTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) HasSettingsSessionTimeoutInMinutes() bool`

HasSettingsSessionTimeoutInMinutes returns a boolean if a field has been set.

### SetSettingsSessionTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) SetSettingsSessionTimeoutInMinutes(v int32)`

SetSettingsSessionTimeoutInMinutes gets a reference to the given int32 and assigns it to the SettingsSessionTimeoutInMinutes field.

### SetSettingsSessionTimeoutInMinutesExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetSettingsSessionTimeoutInMinutesExplicitNull(b bool)`

SetSettingsSessionTimeoutInMinutesExplicitNull (un)sets SettingsSessionTimeoutInMinutes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsSessionTimeoutInMinutes value is set to nil even if false is passed
### GetSettingsSleepTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) GetSettingsSleepTimeoutInMinutes() int32`

GetSettingsSleepTimeoutInMinutes returns the SettingsSleepTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSettingsSleepTimeoutInMinutesOk

`func (o *Windows10TeamGeneralConfiguration) GetSettingsSleepTimeoutInMinutesOk() (int32, bool)`

GetSettingsSleepTimeoutInMinutesOk returns a tuple with the SettingsSleepTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsSleepTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) HasSettingsSleepTimeoutInMinutes() bool`

HasSettingsSleepTimeoutInMinutes returns a boolean if a field has been set.

### SetSettingsSleepTimeoutInMinutes

`func (o *Windows10TeamGeneralConfiguration) SetSettingsSleepTimeoutInMinutes(v int32)`

SetSettingsSleepTimeoutInMinutes gets a reference to the given int32 and assigns it to the SettingsSleepTimeoutInMinutes field.

### SetSettingsSleepTimeoutInMinutesExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetSettingsSleepTimeoutInMinutesExplicitNull(b bool)`

SetSettingsSleepTimeoutInMinutesExplicitNull (un)sets SettingsSleepTimeoutInMinutes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingsSleepTimeoutInMinutes value is set to nil even if false is passed
### GetWelcomeScreenBlockAutomaticWakeUp

`func (o *Windows10TeamGeneralConfiguration) GetWelcomeScreenBlockAutomaticWakeUp() bool`

GetWelcomeScreenBlockAutomaticWakeUp returns the WelcomeScreenBlockAutomaticWakeUp field if non-nil, zero value otherwise.

### GetWelcomeScreenBlockAutomaticWakeUpOk

`func (o *Windows10TeamGeneralConfiguration) GetWelcomeScreenBlockAutomaticWakeUpOk() (bool, bool)`

GetWelcomeScreenBlockAutomaticWakeUpOk returns a tuple with the WelcomeScreenBlockAutomaticWakeUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWelcomeScreenBlockAutomaticWakeUp

`func (o *Windows10TeamGeneralConfiguration) HasWelcomeScreenBlockAutomaticWakeUp() bool`

HasWelcomeScreenBlockAutomaticWakeUp returns a boolean if a field has been set.

### SetWelcomeScreenBlockAutomaticWakeUp

`func (o *Windows10TeamGeneralConfiguration) SetWelcomeScreenBlockAutomaticWakeUp(v bool)`

SetWelcomeScreenBlockAutomaticWakeUp gets a reference to the given bool and assigns it to the WelcomeScreenBlockAutomaticWakeUp field.

### GetWelcomeScreenBackgroundImageUrl

`func (o *Windows10TeamGeneralConfiguration) GetWelcomeScreenBackgroundImageUrl() string`

GetWelcomeScreenBackgroundImageUrl returns the WelcomeScreenBackgroundImageUrl field if non-nil, zero value otherwise.

### GetWelcomeScreenBackgroundImageUrlOk

`func (o *Windows10TeamGeneralConfiguration) GetWelcomeScreenBackgroundImageUrlOk() (string, bool)`

GetWelcomeScreenBackgroundImageUrlOk returns a tuple with the WelcomeScreenBackgroundImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWelcomeScreenBackgroundImageUrl

`func (o *Windows10TeamGeneralConfiguration) HasWelcomeScreenBackgroundImageUrl() bool`

HasWelcomeScreenBackgroundImageUrl returns a boolean if a field has been set.

### SetWelcomeScreenBackgroundImageUrl

`func (o *Windows10TeamGeneralConfiguration) SetWelcomeScreenBackgroundImageUrl(v string)`

SetWelcomeScreenBackgroundImageUrl gets a reference to the given string and assigns it to the WelcomeScreenBackgroundImageUrl field.

### SetWelcomeScreenBackgroundImageUrlExplicitNull

`func (o *Windows10TeamGeneralConfiguration) SetWelcomeScreenBackgroundImageUrlExplicitNull(b bool)`

SetWelcomeScreenBackgroundImageUrlExplicitNull (un)sets WelcomeScreenBackgroundImageUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WelcomeScreenBackgroundImageUrl value is set to nil even if false is passed
### GetWelcomeScreenMeetingInformation

`func (o *Windows10TeamGeneralConfiguration) GetWelcomeScreenMeetingInformation() AnyOfmicrosoftGraphWelcomeScreenMeetingInformation`

GetWelcomeScreenMeetingInformation returns the WelcomeScreenMeetingInformation field if non-nil, zero value otherwise.

### GetWelcomeScreenMeetingInformationOk

`func (o *Windows10TeamGeneralConfiguration) GetWelcomeScreenMeetingInformationOk() (AnyOfmicrosoftGraphWelcomeScreenMeetingInformation, bool)`

GetWelcomeScreenMeetingInformationOk returns a tuple with the WelcomeScreenMeetingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWelcomeScreenMeetingInformation

`func (o *Windows10TeamGeneralConfiguration) HasWelcomeScreenMeetingInformation() bool`

HasWelcomeScreenMeetingInformation returns a boolean if a field has been set.

### SetWelcomeScreenMeetingInformation

`func (o *Windows10TeamGeneralConfiguration) SetWelcomeScreenMeetingInformation(v AnyOfmicrosoftGraphWelcomeScreenMeetingInformation)`

SetWelcomeScreenMeetingInformation gets a reference to the given AnyOfmicrosoftGraphWelcomeScreenMeetingInformation and assigns it to the WelcomeScreenMeetingInformation field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


