# MicrosoftGraphAndroidGeneralDeviceConfiguration

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
**AppsBlockClipboardSharing** | Pointer to **bool** | Indicates whether or not to block clipboard sharing to copy and paste between applications. | [optional] 
**AppsBlockCopyPaste** | Pointer to **bool** | Indicates whether or not to block copy and paste within applications. | [optional] 
**AppsBlockYouTube** | Pointer to **bool** | Indicates whether or not to block the YouTube app. | [optional] 
**BluetoothBlocked** | Pointer to **bool** | Indicates whether or not to block Bluetooth. | [optional] 
**CameraBlocked** | Pointer to **bool** | Indicates whether or not to block the use of the camera. | [optional] 
**CellularBlockDataRoaming** | Pointer to **bool** | Indicates whether or not to block data roaming. | [optional] 
**CellularBlockMessaging** | Pointer to **bool** | Indicates whether or not to block SMS/MMS messaging. | [optional] 
**CellularBlockVoiceRoaming** | Pointer to **bool** | Indicates whether or not to block voice roaming. | [optional] 
**CellularBlockWiFiTethering** | Pointer to **bool** | Indicates whether or not to block syncing Wi-Fi tethering. | [optional] 
**CompliantAppsList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements. | [optional] 
**CompliantAppListType** | Pointer to [**AnyOfmicrosoftGraphAppListType**](anyOf&lt;microsoft.graph.appListType&gt;.md) | Type of list that is in the CompliantAppsList. | [optional] 
**DiagnosticDataBlockSubmission** | Pointer to **bool** | Indicates whether or not to block diagnostic data submission. | [optional] 
**LocationServicesBlocked** | Pointer to **bool** | Indicates whether or not to block location services. | [optional] 
**GoogleAccountBlockAutoSync** | Pointer to **bool** | Indicates whether or not to block Google account auto sync. | [optional] 
**GooglePlayStoreBlocked** | Pointer to **bool** | Indicates whether or not to block the Google Play store. | [optional] 
**KioskModeBlockSleepButton** | Pointer to **bool** | Indicates whether or not to block the screen sleep button while in Kiosk Mode. | [optional] 
**KioskModeBlockVolumeButtons** | Pointer to **bool** | Indicates whether or not to block the volume buttons while in Kiosk Mode. | [optional] 
**KioskModeApps** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | A list of apps that will be allowed to run when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements. | [optional] 
**NfcBlocked** | Pointer to **bool** | Indicates whether or not to block Near-Field Communication. | [optional] 
**PasswordBlockFingerprintUnlock** | Pointer to **bool** | Indicates whether or not to block fingerprint unlock. | [optional] 
**PasswordBlockTrustAgents** | Pointer to **bool** | Indicates whether or not to block Smart Lock and other trust agents. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. Valid values 1 to 365 | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of passwords. Valid values 4 to 16 | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before the screen times out. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 0 to 24 | [optional] 
**PasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | Number of sign in failures allowed before factory reset. Valid values 1 to 16 | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidRequiredPasswordType**](anyOf&lt;microsoft.graph.androidRequiredPasswordType&gt;.md) | Type of password that is required. | [optional] 
**PasswordRequired** | Pointer to **bool** | Indicates whether or not to require a password. | [optional] 
**PowerOffBlocked** | Pointer to **bool** | Indicates whether or not to block powering off the device. | [optional] 
**FactoryResetBlocked** | Pointer to **bool** | Indicates whether or not to block user performing a factory reset. | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether or not to block screenshots. | [optional] 
**DeviceSharingAllowed** | Pointer to **bool** | Indicates whether or not to allow device sharing mode. | [optional] 
**StorageBlockGoogleBackup** | Pointer to **bool** | Indicates whether or not to block Google Backup. | [optional] 
**StorageBlockRemovableStorage** | Pointer to **bool** | Indicates whether or not to block removable storage usage. | [optional] 
**StorageRequireDeviceEncryption** | Pointer to **bool** | Indicates whether or not to require device encryption. | [optional] 
**StorageRequireRemovableStorageEncryption** | Pointer to **bool** | Indicates whether or not to require removable storage encryption. | [optional] 
**VoiceAssistantBlocked** | Pointer to **bool** | Indicates whether or not to block the use of the Voice Assistant. | [optional] 
**VoiceDialingBlocked** | Pointer to **bool** | Indicates whether or not to block voice dialing. | [optional] 
**WebBrowserBlockPopups** | Pointer to **bool** | Indicates whether or not to block popups within the web browser. | [optional] 
**WebBrowserBlockAutofill** | Pointer to **bool** | Indicates whether or not to block the web browser&#39;s auto fill feature. | [optional] 
**WebBrowserBlockJavaScript** | Pointer to **bool** | Indicates whether or not to block JavaScript within the web browser. | [optional] 
**WebBrowserBlocked** | Pointer to **bool** | Indicates whether or not to block the web browser. | [optional] 
**WebBrowserCookieSettings** | Pointer to [**AnyOfmicrosoftGraphWebBrowserCookieSettings**](anyOf&lt;microsoft.graph.webBrowserCookieSettings&gt;.md) | Cookie settings within the web browser. | [optional] 
**WiFiBlocked** | Pointer to **bool** | Indicates whether or not to block syncing Wi-Fi. | [optional] 
**AppsInstallAllowList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps which can be installed on the KNOX device. This collection can contain a maximum of 500 elements. | [optional] 
**AppsLaunchBlockList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps which are blocked from being launched on the KNOX device. This collection can contain a maximum of 500 elements. | [optional] 
**AppsHideList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps to be hidden on the KNOX device. This collection can contain a maximum of 500 elements. | [optional] 
**SecurityRequireVerifyApps** | Pointer to **bool** | Require the Android Verify apps feature is turned on. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAppsBlockClipboardSharing

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsBlockClipboardSharing() bool`

GetAppsBlockClipboardSharing returns the AppsBlockClipboardSharing field if non-nil, zero value otherwise.

### GetAppsBlockClipboardSharingOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsBlockClipboardSharingOk() (bool, bool)`

GetAppsBlockClipboardSharingOk returns a tuple with the AppsBlockClipboardSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockClipboardSharing

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAppsBlockClipboardSharing() bool`

HasAppsBlockClipboardSharing returns a boolean if a field has been set.

### SetAppsBlockClipboardSharing

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAppsBlockClipboardSharing(v bool)`

SetAppsBlockClipboardSharing gets a reference to the given bool and assigns it to the AppsBlockClipboardSharing field.

### GetAppsBlockCopyPaste

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsBlockCopyPaste() bool`

GetAppsBlockCopyPaste returns the AppsBlockCopyPaste field if non-nil, zero value otherwise.

### GetAppsBlockCopyPasteOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsBlockCopyPasteOk() (bool, bool)`

GetAppsBlockCopyPasteOk returns a tuple with the AppsBlockCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockCopyPaste

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAppsBlockCopyPaste() bool`

HasAppsBlockCopyPaste returns a boolean if a field has been set.

### SetAppsBlockCopyPaste

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAppsBlockCopyPaste(v bool)`

SetAppsBlockCopyPaste gets a reference to the given bool and assigns it to the AppsBlockCopyPaste field.

### GetAppsBlockYouTube

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsBlockYouTube() bool`

GetAppsBlockYouTube returns the AppsBlockYouTube field if non-nil, zero value otherwise.

### GetAppsBlockYouTubeOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsBlockYouTubeOk() (bool, bool)`

GetAppsBlockYouTubeOk returns a tuple with the AppsBlockYouTube field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockYouTube

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAppsBlockYouTube() bool`

HasAppsBlockYouTube returns a boolean if a field has been set.

### SetAppsBlockYouTube

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAppsBlockYouTube(v bool)`

SetAppsBlockYouTube gets a reference to the given bool and assigns it to the AppsBlockYouTube field.

### GetBluetoothBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetBluetoothBlocked() bool`

GetBluetoothBlocked returns the BluetoothBlocked field if non-nil, zero value otherwise.

### GetBluetoothBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetBluetoothBlockedOk() (bool, bool)`

GetBluetoothBlockedOk returns a tuple with the BluetoothBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasBluetoothBlocked() bool`

HasBluetoothBlocked returns a boolean if a field has been set.

### SetBluetoothBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetBluetoothBlocked(v bool)`

SetBluetoothBlocked gets a reference to the given bool and assigns it to the BluetoothBlocked field.

### GetCameraBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetCellularBlockDataRoaming

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockDataRoaming() bool`

GetCellularBlockDataRoaming returns the CellularBlockDataRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataRoamingOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockDataRoamingOk() (bool, bool)`

GetCellularBlockDataRoamingOk returns a tuple with the CellularBlockDataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataRoaming

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCellularBlockDataRoaming() bool`

HasCellularBlockDataRoaming returns a boolean if a field has been set.

### SetCellularBlockDataRoaming

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCellularBlockDataRoaming(v bool)`

SetCellularBlockDataRoaming gets a reference to the given bool and assigns it to the CellularBlockDataRoaming field.

### GetCellularBlockMessaging

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockMessaging() bool`

GetCellularBlockMessaging returns the CellularBlockMessaging field if non-nil, zero value otherwise.

### GetCellularBlockMessagingOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockMessagingOk() (bool, bool)`

GetCellularBlockMessagingOk returns a tuple with the CellularBlockMessaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockMessaging

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCellularBlockMessaging() bool`

HasCellularBlockMessaging returns a boolean if a field has been set.

### SetCellularBlockMessaging

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCellularBlockMessaging(v bool)`

SetCellularBlockMessaging gets a reference to the given bool and assigns it to the CellularBlockMessaging field.

### GetCellularBlockVoiceRoaming

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming() bool`

GetCellularBlockVoiceRoaming returns the CellularBlockVoiceRoaming field if non-nil, zero value otherwise.

### GetCellularBlockVoiceRoamingOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockVoiceRoamingOk() (bool, bool)`

GetCellularBlockVoiceRoamingOk returns a tuple with the CellularBlockVoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVoiceRoaming

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCellularBlockVoiceRoaming() bool`

HasCellularBlockVoiceRoaming returns a boolean if a field has been set.

### SetCellularBlockVoiceRoaming

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(v bool)`

SetCellularBlockVoiceRoaming gets a reference to the given bool and assigns it to the CellularBlockVoiceRoaming field.

### GetCellularBlockWiFiTethering

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockWiFiTethering() bool`

GetCellularBlockWiFiTethering returns the CellularBlockWiFiTethering field if non-nil, zero value otherwise.

### GetCellularBlockWiFiTetheringOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCellularBlockWiFiTetheringOk() (bool, bool)`

GetCellularBlockWiFiTetheringOk returns a tuple with the CellularBlockWiFiTethering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockWiFiTethering

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCellularBlockWiFiTethering() bool`

HasCellularBlockWiFiTethering returns a boolean if a field has been set.

### SetCellularBlockWiFiTethering

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCellularBlockWiFiTethering(v bool)`

SetCellularBlockWiFiTethering gets a reference to the given bool and assigns it to the CellularBlockWiFiTethering field.

### GetCompliantAppsList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission() bool`

GetDiagnosticDataBlockSubmission returns the DiagnosticDataBlockSubmission field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionOk returns a tuple with the DiagnosticDataBlockSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDiagnosticDataBlockSubmission() bool`

HasDiagnosticDataBlockSubmission returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(v bool)`

SetDiagnosticDataBlockSubmission gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmission field.

### GetLocationServicesBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetLocationServicesBlocked() bool`

GetLocationServicesBlocked returns the LocationServicesBlocked field if non-nil, zero value otherwise.

### GetLocationServicesBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetLocationServicesBlockedOk() (bool, bool)`

GetLocationServicesBlockedOk returns a tuple with the LocationServicesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationServicesBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasLocationServicesBlocked() bool`

HasLocationServicesBlocked returns a boolean if a field has been set.

### SetLocationServicesBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetLocationServicesBlocked(v bool)`

SetLocationServicesBlocked gets a reference to the given bool and assigns it to the LocationServicesBlocked field.

### GetGoogleAccountBlockAutoSync

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetGoogleAccountBlockAutoSync() bool`

GetGoogleAccountBlockAutoSync returns the GoogleAccountBlockAutoSync field if non-nil, zero value otherwise.

### GetGoogleAccountBlockAutoSyncOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetGoogleAccountBlockAutoSyncOk() (bool, bool)`

GetGoogleAccountBlockAutoSyncOk returns a tuple with the GoogleAccountBlockAutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGoogleAccountBlockAutoSync

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasGoogleAccountBlockAutoSync() bool`

HasGoogleAccountBlockAutoSync returns a boolean if a field has been set.

### SetGoogleAccountBlockAutoSync

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetGoogleAccountBlockAutoSync(v bool)`

SetGoogleAccountBlockAutoSync gets a reference to the given bool and assigns it to the GoogleAccountBlockAutoSync field.

### GetGooglePlayStoreBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetGooglePlayStoreBlocked() bool`

GetGooglePlayStoreBlocked returns the GooglePlayStoreBlocked field if non-nil, zero value otherwise.

### GetGooglePlayStoreBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetGooglePlayStoreBlockedOk() (bool, bool)`

GetGooglePlayStoreBlockedOk returns a tuple with the GooglePlayStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGooglePlayStoreBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasGooglePlayStoreBlocked() bool`

HasGooglePlayStoreBlocked returns a boolean if a field has been set.

### SetGooglePlayStoreBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetGooglePlayStoreBlocked(v bool)`

SetGooglePlayStoreBlocked gets a reference to the given bool and assigns it to the GooglePlayStoreBlocked field.

### GetKioskModeBlockSleepButton

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetKioskModeBlockSleepButton() bool`

GetKioskModeBlockSleepButton returns the KioskModeBlockSleepButton field if non-nil, zero value otherwise.

### GetKioskModeBlockSleepButtonOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetKioskModeBlockSleepButtonOk() (bool, bool)`

GetKioskModeBlockSleepButtonOk returns a tuple with the KioskModeBlockSleepButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeBlockSleepButton

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasKioskModeBlockSleepButton() bool`

HasKioskModeBlockSleepButton returns a boolean if a field has been set.

### SetKioskModeBlockSleepButton

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetKioskModeBlockSleepButton(v bool)`

SetKioskModeBlockSleepButton gets a reference to the given bool and assigns it to the KioskModeBlockSleepButton field.

### GetKioskModeBlockVolumeButtons

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetKioskModeBlockVolumeButtons() bool`

GetKioskModeBlockVolumeButtons returns the KioskModeBlockVolumeButtons field if non-nil, zero value otherwise.

### GetKioskModeBlockVolumeButtonsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetKioskModeBlockVolumeButtonsOk() (bool, bool)`

GetKioskModeBlockVolumeButtonsOk returns a tuple with the KioskModeBlockVolumeButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeBlockVolumeButtons

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasKioskModeBlockVolumeButtons() bool`

HasKioskModeBlockVolumeButtons returns a boolean if a field has been set.

### SetKioskModeBlockVolumeButtons

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetKioskModeBlockVolumeButtons(v bool)`

SetKioskModeBlockVolumeButtons gets a reference to the given bool and assigns it to the KioskModeBlockVolumeButtons field.

### GetKioskModeApps

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetKioskModeApps() []AnyOfmicrosoftGraphAppListItem`

GetKioskModeApps returns the KioskModeApps field if non-nil, zero value otherwise.

### GetKioskModeAppsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetKioskModeAppsOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetKioskModeAppsOk returns a tuple with the KioskModeApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeApps

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasKioskModeApps() bool`

HasKioskModeApps returns a boolean if a field has been set.

### SetKioskModeApps

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetKioskModeApps(v []AnyOfmicrosoftGraphAppListItem)`

SetKioskModeApps gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the KioskModeApps field.

### GetNfcBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetNfcBlocked() bool`

GetNfcBlocked returns the NfcBlocked field if non-nil, zero value otherwise.

### GetNfcBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetNfcBlockedOk() (bool, bool)`

GetNfcBlockedOk returns a tuple with the NfcBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNfcBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasNfcBlocked() bool`

HasNfcBlocked returns a boolean if a field has been set.

### SetNfcBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetNfcBlocked(v bool)`

SetNfcBlocked gets a reference to the given bool and assigns it to the NfcBlocked field.

### GetPasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock() bool`

GetPasswordBlockFingerprintUnlock returns the PasswordBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetPasswordBlockFingerprintUnlockOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlockOk() (bool, bool)`

GetPasswordBlockFingerprintUnlockOk returns a tuple with the PasswordBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordBlockFingerprintUnlock() bool`

HasPasswordBlockFingerprintUnlock returns a boolean if a field has been set.

### SetPasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(v bool)`

SetPasswordBlockFingerprintUnlock gets a reference to the given bool and assigns it to the PasswordBlockFingerprintUnlock field.

### GetPasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordBlockTrustAgents() bool`

GetPasswordBlockTrustAgents returns the PasswordBlockTrustAgents field if non-nil, zero value otherwise.

### GetPasswordBlockTrustAgentsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordBlockTrustAgentsOk() (bool, bool)`

GetPasswordBlockTrustAgentsOk returns a tuple with the PasswordBlockTrustAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordBlockTrustAgents() bool`

HasPasswordBlockTrustAgents returns a boolean if a field has been set.

### SetPasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(v bool)`

SetPasswordBlockTrustAgents gets a reference to the given bool and assigns it to the PasswordBlockTrustAgents field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphAndroidRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphAndroidRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordRequired

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPowerOffBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPowerOffBlocked() bool`

GetPowerOffBlocked returns the PowerOffBlocked field if non-nil, zero value otherwise.

### GetPowerOffBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetPowerOffBlockedOk() (bool, bool)`

GetPowerOffBlockedOk returns a tuple with the PowerOffBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPowerOffBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasPowerOffBlocked() bool`

HasPowerOffBlocked returns a boolean if a field has been set.

### SetPowerOffBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetPowerOffBlocked(v bool)`

SetPowerOffBlocked gets a reference to the given bool and assigns it to the PowerOffBlocked field.

### GetFactoryResetBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetFactoryResetBlocked() bool`

GetFactoryResetBlocked returns the FactoryResetBlocked field if non-nil, zero value otherwise.

### GetFactoryResetBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetFactoryResetBlockedOk() (bool, bool)`

GetFactoryResetBlockedOk returns a tuple with the FactoryResetBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFactoryResetBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasFactoryResetBlocked() bool`

HasFactoryResetBlocked returns a boolean if a field has been set.

### SetFactoryResetBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetFactoryResetBlocked(v bool)`

SetFactoryResetBlocked gets a reference to the given bool and assigns it to the FactoryResetBlocked field.

### GetScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetDeviceSharingAllowed

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceSharingAllowed() bool`

GetDeviceSharingAllowed returns the DeviceSharingAllowed field if non-nil, zero value otherwise.

### GetDeviceSharingAllowedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetDeviceSharingAllowedOk() (bool, bool)`

GetDeviceSharingAllowedOk returns a tuple with the DeviceSharingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSharingAllowed

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasDeviceSharingAllowed() bool`

HasDeviceSharingAllowed returns a boolean if a field has been set.

### SetDeviceSharingAllowed

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetDeviceSharingAllowed(v bool)`

SetDeviceSharingAllowed gets a reference to the given bool and assigns it to the DeviceSharingAllowed field.

### GetStorageBlockGoogleBackup

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageBlockGoogleBackup() bool`

GetStorageBlockGoogleBackup returns the StorageBlockGoogleBackup field if non-nil, zero value otherwise.

### GetStorageBlockGoogleBackupOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageBlockGoogleBackupOk() (bool, bool)`

GetStorageBlockGoogleBackupOk returns a tuple with the StorageBlockGoogleBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockGoogleBackup

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasStorageBlockGoogleBackup() bool`

HasStorageBlockGoogleBackup returns a boolean if a field has been set.

### SetStorageBlockGoogleBackup

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetStorageBlockGoogleBackup(v bool)`

SetStorageBlockGoogleBackup gets a reference to the given bool and assigns it to the StorageBlockGoogleBackup field.

### GetStorageBlockRemovableStorage

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageBlockRemovableStorage() bool`

GetStorageBlockRemovableStorage returns the StorageBlockRemovableStorage field if non-nil, zero value otherwise.

### GetStorageBlockRemovableStorageOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageBlockRemovableStorageOk() (bool, bool)`

GetStorageBlockRemovableStorageOk returns a tuple with the StorageBlockRemovableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockRemovableStorage

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasStorageBlockRemovableStorage() bool`

HasStorageBlockRemovableStorage returns a boolean if a field has been set.

### SetStorageBlockRemovableStorage

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetStorageBlockRemovableStorage(v bool)`

SetStorageBlockRemovableStorage gets a reference to the given bool and assigns it to the StorageBlockRemovableStorage field.

### GetStorageRequireDeviceEncryption

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageRequireDeviceEncryption() bool`

GetStorageRequireDeviceEncryption returns the StorageRequireDeviceEncryption field if non-nil, zero value otherwise.

### GetStorageRequireDeviceEncryptionOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageRequireDeviceEncryptionOk() (bool, bool)`

GetStorageRequireDeviceEncryptionOk returns a tuple with the StorageRequireDeviceEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireDeviceEncryption

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasStorageRequireDeviceEncryption() bool`

HasStorageRequireDeviceEncryption returns a boolean if a field has been set.

### SetStorageRequireDeviceEncryption

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetStorageRequireDeviceEncryption(v bool)`

SetStorageRequireDeviceEncryption gets a reference to the given bool and assigns it to the StorageRequireDeviceEncryption field.

### GetStorageRequireRemovableStorageEncryption

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageRequireRemovableStorageEncryption() bool`

GetStorageRequireRemovableStorageEncryption returns the StorageRequireRemovableStorageEncryption field if non-nil, zero value otherwise.

### GetStorageRequireRemovableStorageEncryptionOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetStorageRequireRemovableStorageEncryptionOk() (bool, bool)`

GetStorageRequireRemovableStorageEncryptionOk returns a tuple with the StorageRequireRemovableStorageEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireRemovableStorageEncryption

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasStorageRequireRemovableStorageEncryption() bool`

HasStorageRequireRemovableStorageEncryption returns a boolean if a field has been set.

### SetStorageRequireRemovableStorageEncryption

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetStorageRequireRemovableStorageEncryption(v bool)`

SetStorageRequireRemovableStorageEncryption gets a reference to the given bool and assigns it to the StorageRequireRemovableStorageEncryption field.

### GetVoiceAssistantBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetVoiceAssistantBlocked() bool`

GetVoiceAssistantBlocked returns the VoiceAssistantBlocked field if non-nil, zero value otherwise.

### GetVoiceAssistantBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetVoiceAssistantBlockedOk() (bool, bool)`

GetVoiceAssistantBlockedOk returns a tuple with the VoiceAssistantBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceAssistantBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasVoiceAssistantBlocked() bool`

HasVoiceAssistantBlocked returns a boolean if a field has been set.

### SetVoiceAssistantBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetVoiceAssistantBlocked(v bool)`

SetVoiceAssistantBlocked gets a reference to the given bool and assigns it to the VoiceAssistantBlocked field.

### GetVoiceDialingBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetVoiceDialingBlocked() bool`

GetVoiceDialingBlocked returns the VoiceDialingBlocked field if non-nil, zero value otherwise.

### GetVoiceDialingBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetVoiceDialingBlockedOk() (bool, bool)`

GetVoiceDialingBlockedOk returns a tuple with the VoiceDialingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceDialingBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasVoiceDialingBlocked() bool`

HasVoiceDialingBlocked returns a boolean if a field has been set.

### SetVoiceDialingBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetVoiceDialingBlocked(v bool)`

SetVoiceDialingBlocked gets a reference to the given bool and assigns it to the VoiceDialingBlocked field.

### GetWebBrowserBlockPopups

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockPopups() bool`

GetWebBrowserBlockPopups returns the WebBrowserBlockPopups field if non-nil, zero value otherwise.

### GetWebBrowserBlockPopupsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockPopupsOk() (bool, bool)`

GetWebBrowserBlockPopupsOk returns a tuple with the WebBrowserBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlockPopups

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasWebBrowserBlockPopups() bool`

HasWebBrowserBlockPopups returns a boolean if a field has been set.

### SetWebBrowserBlockPopups

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetWebBrowserBlockPopups(v bool)`

SetWebBrowserBlockPopups gets a reference to the given bool and assigns it to the WebBrowserBlockPopups field.

### GetWebBrowserBlockAutofill

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockAutofill() bool`

GetWebBrowserBlockAutofill returns the WebBrowserBlockAutofill field if non-nil, zero value otherwise.

### GetWebBrowserBlockAutofillOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockAutofillOk() (bool, bool)`

GetWebBrowserBlockAutofillOk returns a tuple with the WebBrowserBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlockAutofill

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasWebBrowserBlockAutofill() bool`

HasWebBrowserBlockAutofill returns a boolean if a field has been set.

### SetWebBrowserBlockAutofill

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetWebBrowserBlockAutofill(v bool)`

SetWebBrowserBlockAutofill gets a reference to the given bool and assigns it to the WebBrowserBlockAutofill field.

### GetWebBrowserBlockJavaScript

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockJavaScript() bool`

GetWebBrowserBlockJavaScript returns the WebBrowserBlockJavaScript field if non-nil, zero value otherwise.

### GetWebBrowserBlockJavaScriptOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockJavaScriptOk() (bool, bool)`

GetWebBrowserBlockJavaScriptOk returns a tuple with the WebBrowserBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlockJavaScript

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasWebBrowserBlockJavaScript() bool`

HasWebBrowserBlockJavaScript returns a boolean if a field has been set.

### SetWebBrowserBlockJavaScript

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetWebBrowserBlockJavaScript(v bool)`

SetWebBrowserBlockJavaScript gets a reference to the given bool and assigns it to the WebBrowserBlockJavaScript field.

### GetWebBrowserBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlocked() bool`

GetWebBrowserBlocked returns the WebBrowserBlocked field if non-nil, zero value otherwise.

### GetWebBrowserBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserBlockedOk() (bool, bool)`

GetWebBrowserBlockedOk returns a tuple with the WebBrowserBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasWebBrowserBlocked() bool`

HasWebBrowserBlocked returns a boolean if a field has been set.

### SetWebBrowserBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetWebBrowserBlocked(v bool)`

SetWebBrowserBlocked gets a reference to the given bool and assigns it to the WebBrowserBlocked field.

### GetWebBrowserCookieSettings

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserCookieSettings() AnyOfmicrosoftGraphWebBrowserCookieSettings`

GetWebBrowserCookieSettings returns the WebBrowserCookieSettings field if non-nil, zero value otherwise.

### GetWebBrowserCookieSettingsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWebBrowserCookieSettingsOk() (AnyOfmicrosoftGraphWebBrowserCookieSettings, bool)`

GetWebBrowserCookieSettingsOk returns a tuple with the WebBrowserCookieSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserCookieSettings

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasWebBrowserCookieSettings() bool`

HasWebBrowserCookieSettings returns a boolean if a field has been set.

### SetWebBrowserCookieSettings

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetWebBrowserCookieSettings(v AnyOfmicrosoftGraphWebBrowserCookieSettings)`

SetWebBrowserCookieSettings gets a reference to the given AnyOfmicrosoftGraphWebBrowserCookieSettings and assigns it to the WebBrowserCookieSettings field.

### GetWiFiBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWiFiBlocked() bool`

GetWiFiBlocked returns the WiFiBlocked field if non-nil, zero value otherwise.

### GetWiFiBlockedOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetWiFiBlockedOk() (bool, bool)`

GetWiFiBlockedOk returns a tuple with the WiFiBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasWiFiBlocked() bool`

HasWiFiBlocked returns a boolean if a field has been set.

### SetWiFiBlocked

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetWiFiBlocked(v bool)`

SetWiFiBlocked gets a reference to the given bool and assigns it to the WiFiBlocked field.

### GetAppsInstallAllowList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsInstallAllowList() []AnyOfmicrosoftGraphAppListItem`

GetAppsInstallAllowList returns the AppsInstallAllowList field if non-nil, zero value otherwise.

### GetAppsInstallAllowListOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsInstallAllowListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsInstallAllowListOk returns a tuple with the AppsInstallAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsInstallAllowList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAppsInstallAllowList() bool`

HasAppsInstallAllowList returns a boolean if a field has been set.

### SetAppsInstallAllowList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAppsInstallAllowList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsInstallAllowList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsInstallAllowList field.

### GetAppsLaunchBlockList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsLaunchBlockList() []AnyOfmicrosoftGraphAppListItem`

GetAppsLaunchBlockList returns the AppsLaunchBlockList field if non-nil, zero value otherwise.

### GetAppsLaunchBlockListOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsLaunchBlockListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsLaunchBlockListOk returns a tuple with the AppsLaunchBlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsLaunchBlockList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAppsLaunchBlockList() bool`

HasAppsLaunchBlockList returns a boolean if a field has been set.

### SetAppsLaunchBlockList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAppsLaunchBlockList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsLaunchBlockList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsLaunchBlockList field.

### GetAppsHideList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsHideList() []AnyOfmicrosoftGraphAppListItem`

GetAppsHideList returns the AppsHideList field if non-nil, zero value otherwise.

### GetAppsHideListOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetAppsHideListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsHideListOk returns a tuple with the AppsHideList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsHideList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasAppsHideList() bool`

HasAppsHideList returns a boolean if a field has been set.

### SetAppsHideList

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetAppsHideList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsHideList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsHideList field.

### GetSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetSecurityRequireVerifyApps() bool`

GetSecurityRequireVerifyApps returns the SecurityRequireVerifyApps field if non-nil, zero value otherwise.

### GetSecurityRequireVerifyAppsOk

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) GetSecurityRequireVerifyAppsOk() (bool, bool)`

GetSecurityRequireVerifyAppsOk returns a tuple with the SecurityRequireVerifyApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) HasSecurityRequireVerifyApps() bool`

HasSecurityRequireVerifyApps returns a boolean if a field has been set.

### SetSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(v bool)`

SetSecurityRequireVerifyApps gets a reference to the given bool and assigns it to the SecurityRequireVerifyApps field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


