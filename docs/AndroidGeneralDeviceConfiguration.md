# AndroidGeneralDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAppsBlockClipboardSharing

`func (o *AndroidGeneralDeviceConfiguration) GetAppsBlockClipboardSharing() bool`

GetAppsBlockClipboardSharing returns the AppsBlockClipboardSharing field if non-nil, zero value otherwise.

### GetAppsBlockClipboardSharingOk

`func (o *AndroidGeneralDeviceConfiguration) GetAppsBlockClipboardSharingOk() (bool, bool)`

GetAppsBlockClipboardSharingOk returns a tuple with the AppsBlockClipboardSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockClipboardSharing

`func (o *AndroidGeneralDeviceConfiguration) HasAppsBlockClipboardSharing() bool`

HasAppsBlockClipboardSharing returns a boolean if a field has been set.

### SetAppsBlockClipboardSharing

`func (o *AndroidGeneralDeviceConfiguration) SetAppsBlockClipboardSharing(v bool)`

SetAppsBlockClipboardSharing gets a reference to the given bool and assigns it to the AppsBlockClipboardSharing field.

### GetAppsBlockCopyPaste

`func (o *AndroidGeneralDeviceConfiguration) GetAppsBlockCopyPaste() bool`

GetAppsBlockCopyPaste returns the AppsBlockCopyPaste field if non-nil, zero value otherwise.

### GetAppsBlockCopyPasteOk

`func (o *AndroidGeneralDeviceConfiguration) GetAppsBlockCopyPasteOk() (bool, bool)`

GetAppsBlockCopyPasteOk returns a tuple with the AppsBlockCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockCopyPaste

`func (o *AndroidGeneralDeviceConfiguration) HasAppsBlockCopyPaste() bool`

HasAppsBlockCopyPaste returns a boolean if a field has been set.

### SetAppsBlockCopyPaste

`func (o *AndroidGeneralDeviceConfiguration) SetAppsBlockCopyPaste(v bool)`

SetAppsBlockCopyPaste gets a reference to the given bool and assigns it to the AppsBlockCopyPaste field.

### GetAppsBlockYouTube

`func (o *AndroidGeneralDeviceConfiguration) GetAppsBlockYouTube() bool`

GetAppsBlockYouTube returns the AppsBlockYouTube field if non-nil, zero value otherwise.

### GetAppsBlockYouTubeOk

`func (o *AndroidGeneralDeviceConfiguration) GetAppsBlockYouTubeOk() (bool, bool)`

GetAppsBlockYouTubeOk returns a tuple with the AppsBlockYouTube field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockYouTube

`func (o *AndroidGeneralDeviceConfiguration) HasAppsBlockYouTube() bool`

HasAppsBlockYouTube returns a boolean if a field has been set.

### SetAppsBlockYouTube

`func (o *AndroidGeneralDeviceConfiguration) SetAppsBlockYouTube(v bool)`

SetAppsBlockYouTube gets a reference to the given bool and assigns it to the AppsBlockYouTube field.

### GetBluetoothBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetBluetoothBlocked() bool`

GetBluetoothBlocked returns the BluetoothBlocked field if non-nil, zero value otherwise.

### GetBluetoothBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetBluetoothBlockedOk() (bool, bool)`

GetBluetoothBlockedOk returns a tuple with the BluetoothBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasBluetoothBlocked() bool`

HasBluetoothBlocked returns a boolean if a field has been set.

### SetBluetoothBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetBluetoothBlocked(v bool)`

SetBluetoothBlocked gets a reference to the given bool and assigns it to the BluetoothBlocked field.

### GetCameraBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetCellularBlockDataRoaming

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockDataRoaming() bool`

GetCellularBlockDataRoaming returns the CellularBlockDataRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataRoamingOk

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockDataRoamingOk() (bool, bool)`

GetCellularBlockDataRoamingOk returns a tuple with the CellularBlockDataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataRoaming

`func (o *AndroidGeneralDeviceConfiguration) HasCellularBlockDataRoaming() bool`

HasCellularBlockDataRoaming returns a boolean if a field has been set.

### SetCellularBlockDataRoaming

`func (o *AndroidGeneralDeviceConfiguration) SetCellularBlockDataRoaming(v bool)`

SetCellularBlockDataRoaming gets a reference to the given bool and assigns it to the CellularBlockDataRoaming field.

### GetCellularBlockMessaging

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockMessaging() bool`

GetCellularBlockMessaging returns the CellularBlockMessaging field if non-nil, zero value otherwise.

### GetCellularBlockMessagingOk

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockMessagingOk() (bool, bool)`

GetCellularBlockMessagingOk returns a tuple with the CellularBlockMessaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockMessaging

`func (o *AndroidGeneralDeviceConfiguration) HasCellularBlockMessaging() bool`

HasCellularBlockMessaging returns a boolean if a field has been set.

### SetCellularBlockMessaging

`func (o *AndroidGeneralDeviceConfiguration) SetCellularBlockMessaging(v bool)`

SetCellularBlockMessaging gets a reference to the given bool and assigns it to the CellularBlockMessaging field.

### GetCellularBlockVoiceRoaming

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming() bool`

GetCellularBlockVoiceRoaming returns the CellularBlockVoiceRoaming field if non-nil, zero value otherwise.

### GetCellularBlockVoiceRoamingOk

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockVoiceRoamingOk() (bool, bool)`

GetCellularBlockVoiceRoamingOk returns a tuple with the CellularBlockVoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVoiceRoaming

`func (o *AndroidGeneralDeviceConfiguration) HasCellularBlockVoiceRoaming() bool`

HasCellularBlockVoiceRoaming returns a boolean if a field has been set.

### SetCellularBlockVoiceRoaming

`func (o *AndroidGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(v bool)`

SetCellularBlockVoiceRoaming gets a reference to the given bool and assigns it to the CellularBlockVoiceRoaming field.

### GetCellularBlockWiFiTethering

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockWiFiTethering() bool`

GetCellularBlockWiFiTethering returns the CellularBlockWiFiTethering field if non-nil, zero value otherwise.

### GetCellularBlockWiFiTetheringOk

`func (o *AndroidGeneralDeviceConfiguration) GetCellularBlockWiFiTetheringOk() (bool, bool)`

GetCellularBlockWiFiTetheringOk returns a tuple with the CellularBlockWiFiTethering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockWiFiTethering

`func (o *AndroidGeneralDeviceConfiguration) HasCellularBlockWiFiTethering() bool`

HasCellularBlockWiFiTethering returns a boolean if a field has been set.

### SetCellularBlockWiFiTethering

`func (o *AndroidGeneralDeviceConfiguration) SetCellularBlockWiFiTethering(v bool)`

SetCellularBlockWiFiTethering gets a reference to the given bool and assigns it to the CellularBlockWiFiTethering field.

### GetCompliantAppsList

`func (o *AndroidGeneralDeviceConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *AndroidGeneralDeviceConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *AndroidGeneralDeviceConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *AndroidGeneralDeviceConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *AndroidGeneralDeviceConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *AndroidGeneralDeviceConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *AndroidGeneralDeviceConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *AndroidGeneralDeviceConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetDiagnosticDataBlockSubmission

`func (o *AndroidGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission() bool`

GetDiagnosticDataBlockSubmission returns the DiagnosticDataBlockSubmission field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionOk

`func (o *AndroidGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionOk returns a tuple with the DiagnosticDataBlockSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmission

`func (o *AndroidGeneralDeviceConfiguration) HasDiagnosticDataBlockSubmission() bool`

HasDiagnosticDataBlockSubmission returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmission

`func (o *AndroidGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(v bool)`

SetDiagnosticDataBlockSubmission gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmission field.

### GetLocationServicesBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetLocationServicesBlocked() bool`

GetLocationServicesBlocked returns the LocationServicesBlocked field if non-nil, zero value otherwise.

### GetLocationServicesBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetLocationServicesBlockedOk() (bool, bool)`

GetLocationServicesBlockedOk returns a tuple with the LocationServicesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationServicesBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasLocationServicesBlocked() bool`

HasLocationServicesBlocked returns a boolean if a field has been set.

### SetLocationServicesBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetLocationServicesBlocked(v bool)`

SetLocationServicesBlocked gets a reference to the given bool and assigns it to the LocationServicesBlocked field.

### GetGoogleAccountBlockAutoSync

`func (o *AndroidGeneralDeviceConfiguration) GetGoogleAccountBlockAutoSync() bool`

GetGoogleAccountBlockAutoSync returns the GoogleAccountBlockAutoSync field if non-nil, zero value otherwise.

### GetGoogleAccountBlockAutoSyncOk

`func (o *AndroidGeneralDeviceConfiguration) GetGoogleAccountBlockAutoSyncOk() (bool, bool)`

GetGoogleAccountBlockAutoSyncOk returns a tuple with the GoogleAccountBlockAutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGoogleAccountBlockAutoSync

`func (o *AndroidGeneralDeviceConfiguration) HasGoogleAccountBlockAutoSync() bool`

HasGoogleAccountBlockAutoSync returns a boolean if a field has been set.

### SetGoogleAccountBlockAutoSync

`func (o *AndroidGeneralDeviceConfiguration) SetGoogleAccountBlockAutoSync(v bool)`

SetGoogleAccountBlockAutoSync gets a reference to the given bool and assigns it to the GoogleAccountBlockAutoSync field.

### GetGooglePlayStoreBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetGooglePlayStoreBlocked() bool`

GetGooglePlayStoreBlocked returns the GooglePlayStoreBlocked field if non-nil, zero value otherwise.

### GetGooglePlayStoreBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetGooglePlayStoreBlockedOk() (bool, bool)`

GetGooglePlayStoreBlockedOk returns a tuple with the GooglePlayStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGooglePlayStoreBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasGooglePlayStoreBlocked() bool`

HasGooglePlayStoreBlocked returns a boolean if a field has been set.

### SetGooglePlayStoreBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetGooglePlayStoreBlocked(v bool)`

SetGooglePlayStoreBlocked gets a reference to the given bool and assigns it to the GooglePlayStoreBlocked field.

### GetKioskModeBlockSleepButton

`func (o *AndroidGeneralDeviceConfiguration) GetKioskModeBlockSleepButton() bool`

GetKioskModeBlockSleepButton returns the KioskModeBlockSleepButton field if non-nil, zero value otherwise.

### GetKioskModeBlockSleepButtonOk

`func (o *AndroidGeneralDeviceConfiguration) GetKioskModeBlockSleepButtonOk() (bool, bool)`

GetKioskModeBlockSleepButtonOk returns a tuple with the KioskModeBlockSleepButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeBlockSleepButton

`func (o *AndroidGeneralDeviceConfiguration) HasKioskModeBlockSleepButton() bool`

HasKioskModeBlockSleepButton returns a boolean if a field has been set.

### SetKioskModeBlockSleepButton

`func (o *AndroidGeneralDeviceConfiguration) SetKioskModeBlockSleepButton(v bool)`

SetKioskModeBlockSleepButton gets a reference to the given bool and assigns it to the KioskModeBlockSleepButton field.

### GetKioskModeBlockVolumeButtons

`func (o *AndroidGeneralDeviceConfiguration) GetKioskModeBlockVolumeButtons() bool`

GetKioskModeBlockVolumeButtons returns the KioskModeBlockVolumeButtons field if non-nil, zero value otherwise.

### GetKioskModeBlockVolumeButtonsOk

`func (o *AndroidGeneralDeviceConfiguration) GetKioskModeBlockVolumeButtonsOk() (bool, bool)`

GetKioskModeBlockVolumeButtonsOk returns a tuple with the KioskModeBlockVolumeButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeBlockVolumeButtons

`func (o *AndroidGeneralDeviceConfiguration) HasKioskModeBlockVolumeButtons() bool`

HasKioskModeBlockVolumeButtons returns a boolean if a field has been set.

### SetKioskModeBlockVolumeButtons

`func (o *AndroidGeneralDeviceConfiguration) SetKioskModeBlockVolumeButtons(v bool)`

SetKioskModeBlockVolumeButtons gets a reference to the given bool and assigns it to the KioskModeBlockVolumeButtons field.

### GetKioskModeApps

`func (o *AndroidGeneralDeviceConfiguration) GetKioskModeApps() []AnyOfmicrosoftGraphAppListItem`

GetKioskModeApps returns the KioskModeApps field if non-nil, zero value otherwise.

### GetKioskModeAppsOk

`func (o *AndroidGeneralDeviceConfiguration) GetKioskModeAppsOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetKioskModeAppsOk returns a tuple with the KioskModeApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeApps

`func (o *AndroidGeneralDeviceConfiguration) HasKioskModeApps() bool`

HasKioskModeApps returns a boolean if a field has been set.

### SetKioskModeApps

`func (o *AndroidGeneralDeviceConfiguration) SetKioskModeApps(v []AnyOfmicrosoftGraphAppListItem)`

SetKioskModeApps gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the KioskModeApps field.

### GetNfcBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetNfcBlocked() bool`

GetNfcBlocked returns the NfcBlocked field if non-nil, zero value otherwise.

### GetNfcBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetNfcBlockedOk() (bool, bool)`

GetNfcBlockedOk returns a tuple with the NfcBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNfcBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasNfcBlocked() bool`

HasNfcBlocked returns a boolean if a field has been set.

### SetNfcBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetNfcBlocked(v bool)`

SetNfcBlocked gets a reference to the given bool and assigns it to the NfcBlocked field.

### GetPasswordBlockFingerprintUnlock

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock() bool`

GetPasswordBlockFingerprintUnlock returns the PasswordBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetPasswordBlockFingerprintUnlockOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlockOk() (bool, bool)`

GetPasswordBlockFingerprintUnlockOk returns a tuple with the PasswordBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockFingerprintUnlock

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordBlockFingerprintUnlock() bool`

HasPasswordBlockFingerprintUnlock returns a boolean if a field has been set.

### SetPasswordBlockFingerprintUnlock

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(v bool)`

SetPasswordBlockFingerprintUnlock gets a reference to the given bool and assigns it to the PasswordBlockFingerprintUnlock field.

### GetPasswordBlockTrustAgents

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordBlockTrustAgents() bool`

GetPasswordBlockTrustAgents returns the PasswordBlockTrustAgents field if non-nil, zero value otherwise.

### GetPasswordBlockTrustAgentsOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordBlockTrustAgentsOk() (bool, bool)`

GetPasswordBlockTrustAgentsOk returns a tuple with the PasswordBlockTrustAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockTrustAgents

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordBlockTrustAgents() bool`

HasPasswordBlockTrustAgents returns a boolean if a field has been set.

### SetPasswordBlockTrustAgents

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(v bool)`

SetPasswordBlockTrustAgents gets a reference to the given bool and assigns it to the PasswordBlockTrustAgents field.

### GetPasswordExpirationDays

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphAndroidRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphAndroidRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordRequired

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *AndroidGeneralDeviceConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *AndroidGeneralDeviceConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *AndroidGeneralDeviceConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPowerOffBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetPowerOffBlocked() bool`

GetPowerOffBlocked returns the PowerOffBlocked field if non-nil, zero value otherwise.

### GetPowerOffBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetPowerOffBlockedOk() (bool, bool)`

GetPowerOffBlockedOk returns a tuple with the PowerOffBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPowerOffBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasPowerOffBlocked() bool`

HasPowerOffBlocked returns a boolean if a field has been set.

### SetPowerOffBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetPowerOffBlocked(v bool)`

SetPowerOffBlocked gets a reference to the given bool and assigns it to the PowerOffBlocked field.

### GetFactoryResetBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetFactoryResetBlocked() bool`

GetFactoryResetBlocked returns the FactoryResetBlocked field if non-nil, zero value otherwise.

### GetFactoryResetBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetFactoryResetBlockedOk() (bool, bool)`

GetFactoryResetBlockedOk returns a tuple with the FactoryResetBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFactoryResetBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasFactoryResetBlocked() bool`

HasFactoryResetBlocked returns a boolean if a field has been set.

### SetFactoryResetBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetFactoryResetBlocked(v bool)`

SetFactoryResetBlocked gets a reference to the given bool and assigns it to the FactoryResetBlocked field.

### GetScreenCaptureBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetDeviceSharingAllowed

`func (o *AndroidGeneralDeviceConfiguration) GetDeviceSharingAllowed() bool`

GetDeviceSharingAllowed returns the DeviceSharingAllowed field if non-nil, zero value otherwise.

### GetDeviceSharingAllowedOk

`func (o *AndroidGeneralDeviceConfiguration) GetDeviceSharingAllowedOk() (bool, bool)`

GetDeviceSharingAllowedOk returns a tuple with the DeviceSharingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSharingAllowed

`func (o *AndroidGeneralDeviceConfiguration) HasDeviceSharingAllowed() bool`

HasDeviceSharingAllowed returns a boolean if a field has been set.

### SetDeviceSharingAllowed

`func (o *AndroidGeneralDeviceConfiguration) SetDeviceSharingAllowed(v bool)`

SetDeviceSharingAllowed gets a reference to the given bool and assigns it to the DeviceSharingAllowed field.

### GetStorageBlockGoogleBackup

`func (o *AndroidGeneralDeviceConfiguration) GetStorageBlockGoogleBackup() bool`

GetStorageBlockGoogleBackup returns the StorageBlockGoogleBackup field if non-nil, zero value otherwise.

### GetStorageBlockGoogleBackupOk

`func (o *AndroidGeneralDeviceConfiguration) GetStorageBlockGoogleBackupOk() (bool, bool)`

GetStorageBlockGoogleBackupOk returns a tuple with the StorageBlockGoogleBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockGoogleBackup

`func (o *AndroidGeneralDeviceConfiguration) HasStorageBlockGoogleBackup() bool`

HasStorageBlockGoogleBackup returns a boolean if a field has been set.

### SetStorageBlockGoogleBackup

`func (o *AndroidGeneralDeviceConfiguration) SetStorageBlockGoogleBackup(v bool)`

SetStorageBlockGoogleBackup gets a reference to the given bool and assigns it to the StorageBlockGoogleBackup field.

### GetStorageBlockRemovableStorage

`func (o *AndroidGeneralDeviceConfiguration) GetStorageBlockRemovableStorage() bool`

GetStorageBlockRemovableStorage returns the StorageBlockRemovableStorage field if non-nil, zero value otherwise.

### GetStorageBlockRemovableStorageOk

`func (o *AndroidGeneralDeviceConfiguration) GetStorageBlockRemovableStorageOk() (bool, bool)`

GetStorageBlockRemovableStorageOk returns a tuple with the StorageBlockRemovableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockRemovableStorage

`func (o *AndroidGeneralDeviceConfiguration) HasStorageBlockRemovableStorage() bool`

HasStorageBlockRemovableStorage returns a boolean if a field has been set.

### SetStorageBlockRemovableStorage

`func (o *AndroidGeneralDeviceConfiguration) SetStorageBlockRemovableStorage(v bool)`

SetStorageBlockRemovableStorage gets a reference to the given bool and assigns it to the StorageBlockRemovableStorage field.

### GetStorageRequireDeviceEncryption

`func (o *AndroidGeneralDeviceConfiguration) GetStorageRequireDeviceEncryption() bool`

GetStorageRequireDeviceEncryption returns the StorageRequireDeviceEncryption field if non-nil, zero value otherwise.

### GetStorageRequireDeviceEncryptionOk

`func (o *AndroidGeneralDeviceConfiguration) GetStorageRequireDeviceEncryptionOk() (bool, bool)`

GetStorageRequireDeviceEncryptionOk returns a tuple with the StorageRequireDeviceEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireDeviceEncryption

`func (o *AndroidGeneralDeviceConfiguration) HasStorageRequireDeviceEncryption() bool`

HasStorageRequireDeviceEncryption returns a boolean if a field has been set.

### SetStorageRequireDeviceEncryption

`func (o *AndroidGeneralDeviceConfiguration) SetStorageRequireDeviceEncryption(v bool)`

SetStorageRequireDeviceEncryption gets a reference to the given bool and assigns it to the StorageRequireDeviceEncryption field.

### GetStorageRequireRemovableStorageEncryption

`func (o *AndroidGeneralDeviceConfiguration) GetStorageRequireRemovableStorageEncryption() bool`

GetStorageRequireRemovableStorageEncryption returns the StorageRequireRemovableStorageEncryption field if non-nil, zero value otherwise.

### GetStorageRequireRemovableStorageEncryptionOk

`func (o *AndroidGeneralDeviceConfiguration) GetStorageRequireRemovableStorageEncryptionOk() (bool, bool)`

GetStorageRequireRemovableStorageEncryptionOk returns a tuple with the StorageRequireRemovableStorageEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireRemovableStorageEncryption

`func (o *AndroidGeneralDeviceConfiguration) HasStorageRequireRemovableStorageEncryption() bool`

HasStorageRequireRemovableStorageEncryption returns a boolean if a field has been set.

### SetStorageRequireRemovableStorageEncryption

`func (o *AndroidGeneralDeviceConfiguration) SetStorageRequireRemovableStorageEncryption(v bool)`

SetStorageRequireRemovableStorageEncryption gets a reference to the given bool and assigns it to the StorageRequireRemovableStorageEncryption field.

### GetVoiceAssistantBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetVoiceAssistantBlocked() bool`

GetVoiceAssistantBlocked returns the VoiceAssistantBlocked field if non-nil, zero value otherwise.

### GetVoiceAssistantBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetVoiceAssistantBlockedOk() (bool, bool)`

GetVoiceAssistantBlockedOk returns a tuple with the VoiceAssistantBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceAssistantBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasVoiceAssistantBlocked() bool`

HasVoiceAssistantBlocked returns a boolean if a field has been set.

### SetVoiceAssistantBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetVoiceAssistantBlocked(v bool)`

SetVoiceAssistantBlocked gets a reference to the given bool and assigns it to the VoiceAssistantBlocked field.

### GetVoiceDialingBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetVoiceDialingBlocked() bool`

GetVoiceDialingBlocked returns the VoiceDialingBlocked field if non-nil, zero value otherwise.

### GetVoiceDialingBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetVoiceDialingBlockedOk() (bool, bool)`

GetVoiceDialingBlockedOk returns a tuple with the VoiceDialingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceDialingBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasVoiceDialingBlocked() bool`

HasVoiceDialingBlocked returns a boolean if a field has been set.

### SetVoiceDialingBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetVoiceDialingBlocked(v bool)`

SetVoiceDialingBlocked gets a reference to the given bool and assigns it to the VoiceDialingBlocked field.

### GetWebBrowserBlockPopups

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockPopups() bool`

GetWebBrowserBlockPopups returns the WebBrowserBlockPopups field if non-nil, zero value otherwise.

### GetWebBrowserBlockPopupsOk

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockPopupsOk() (bool, bool)`

GetWebBrowserBlockPopupsOk returns a tuple with the WebBrowserBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlockPopups

`func (o *AndroidGeneralDeviceConfiguration) HasWebBrowserBlockPopups() bool`

HasWebBrowserBlockPopups returns a boolean if a field has been set.

### SetWebBrowserBlockPopups

`func (o *AndroidGeneralDeviceConfiguration) SetWebBrowserBlockPopups(v bool)`

SetWebBrowserBlockPopups gets a reference to the given bool and assigns it to the WebBrowserBlockPopups field.

### GetWebBrowserBlockAutofill

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockAutofill() bool`

GetWebBrowserBlockAutofill returns the WebBrowserBlockAutofill field if non-nil, zero value otherwise.

### GetWebBrowserBlockAutofillOk

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockAutofillOk() (bool, bool)`

GetWebBrowserBlockAutofillOk returns a tuple with the WebBrowserBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlockAutofill

`func (o *AndroidGeneralDeviceConfiguration) HasWebBrowserBlockAutofill() bool`

HasWebBrowserBlockAutofill returns a boolean if a field has been set.

### SetWebBrowserBlockAutofill

`func (o *AndroidGeneralDeviceConfiguration) SetWebBrowserBlockAutofill(v bool)`

SetWebBrowserBlockAutofill gets a reference to the given bool and assigns it to the WebBrowserBlockAutofill field.

### GetWebBrowserBlockJavaScript

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockJavaScript() bool`

GetWebBrowserBlockJavaScript returns the WebBrowserBlockJavaScript field if non-nil, zero value otherwise.

### GetWebBrowserBlockJavaScriptOk

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockJavaScriptOk() (bool, bool)`

GetWebBrowserBlockJavaScriptOk returns a tuple with the WebBrowserBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlockJavaScript

`func (o *AndroidGeneralDeviceConfiguration) HasWebBrowserBlockJavaScript() bool`

HasWebBrowserBlockJavaScript returns a boolean if a field has been set.

### SetWebBrowserBlockJavaScript

`func (o *AndroidGeneralDeviceConfiguration) SetWebBrowserBlockJavaScript(v bool)`

SetWebBrowserBlockJavaScript gets a reference to the given bool and assigns it to the WebBrowserBlockJavaScript field.

### GetWebBrowserBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlocked() bool`

GetWebBrowserBlocked returns the WebBrowserBlocked field if non-nil, zero value otherwise.

### GetWebBrowserBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockedOk() (bool, bool)`

GetWebBrowserBlockedOk returns a tuple with the WebBrowserBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasWebBrowserBlocked() bool`

HasWebBrowserBlocked returns a boolean if a field has been set.

### SetWebBrowserBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetWebBrowserBlocked(v bool)`

SetWebBrowserBlocked gets a reference to the given bool and assigns it to the WebBrowserBlocked field.

### GetWebBrowserCookieSettings

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserCookieSettings() AnyOfmicrosoftGraphWebBrowserCookieSettings`

GetWebBrowserCookieSettings returns the WebBrowserCookieSettings field if non-nil, zero value otherwise.

### GetWebBrowserCookieSettingsOk

`func (o *AndroidGeneralDeviceConfiguration) GetWebBrowserCookieSettingsOk() (AnyOfmicrosoftGraphWebBrowserCookieSettings, bool)`

GetWebBrowserCookieSettingsOk returns a tuple with the WebBrowserCookieSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserCookieSettings

`func (o *AndroidGeneralDeviceConfiguration) HasWebBrowserCookieSettings() bool`

HasWebBrowserCookieSettings returns a boolean if a field has been set.

### SetWebBrowserCookieSettings

`func (o *AndroidGeneralDeviceConfiguration) SetWebBrowserCookieSettings(v AnyOfmicrosoftGraphWebBrowserCookieSettings)`

SetWebBrowserCookieSettings gets a reference to the given AnyOfmicrosoftGraphWebBrowserCookieSettings and assigns it to the WebBrowserCookieSettings field.

### GetWiFiBlocked

`func (o *AndroidGeneralDeviceConfiguration) GetWiFiBlocked() bool`

GetWiFiBlocked returns the WiFiBlocked field if non-nil, zero value otherwise.

### GetWiFiBlockedOk

`func (o *AndroidGeneralDeviceConfiguration) GetWiFiBlockedOk() (bool, bool)`

GetWiFiBlockedOk returns a tuple with the WiFiBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlocked

`func (o *AndroidGeneralDeviceConfiguration) HasWiFiBlocked() bool`

HasWiFiBlocked returns a boolean if a field has been set.

### SetWiFiBlocked

`func (o *AndroidGeneralDeviceConfiguration) SetWiFiBlocked(v bool)`

SetWiFiBlocked gets a reference to the given bool and assigns it to the WiFiBlocked field.

### GetAppsInstallAllowList

`func (o *AndroidGeneralDeviceConfiguration) GetAppsInstallAllowList() []AnyOfmicrosoftGraphAppListItem`

GetAppsInstallAllowList returns the AppsInstallAllowList field if non-nil, zero value otherwise.

### GetAppsInstallAllowListOk

`func (o *AndroidGeneralDeviceConfiguration) GetAppsInstallAllowListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsInstallAllowListOk returns a tuple with the AppsInstallAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsInstallAllowList

`func (o *AndroidGeneralDeviceConfiguration) HasAppsInstallAllowList() bool`

HasAppsInstallAllowList returns a boolean if a field has been set.

### SetAppsInstallAllowList

`func (o *AndroidGeneralDeviceConfiguration) SetAppsInstallAllowList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsInstallAllowList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsInstallAllowList field.

### GetAppsLaunchBlockList

`func (o *AndroidGeneralDeviceConfiguration) GetAppsLaunchBlockList() []AnyOfmicrosoftGraphAppListItem`

GetAppsLaunchBlockList returns the AppsLaunchBlockList field if non-nil, zero value otherwise.

### GetAppsLaunchBlockListOk

`func (o *AndroidGeneralDeviceConfiguration) GetAppsLaunchBlockListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsLaunchBlockListOk returns a tuple with the AppsLaunchBlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsLaunchBlockList

`func (o *AndroidGeneralDeviceConfiguration) HasAppsLaunchBlockList() bool`

HasAppsLaunchBlockList returns a boolean if a field has been set.

### SetAppsLaunchBlockList

`func (o *AndroidGeneralDeviceConfiguration) SetAppsLaunchBlockList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsLaunchBlockList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsLaunchBlockList field.

### GetAppsHideList

`func (o *AndroidGeneralDeviceConfiguration) GetAppsHideList() []AnyOfmicrosoftGraphAppListItem`

GetAppsHideList returns the AppsHideList field if non-nil, zero value otherwise.

### GetAppsHideListOk

`func (o *AndroidGeneralDeviceConfiguration) GetAppsHideListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsHideListOk returns a tuple with the AppsHideList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsHideList

`func (o *AndroidGeneralDeviceConfiguration) HasAppsHideList() bool`

HasAppsHideList returns a boolean if a field has been set.

### SetAppsHideList

`func (o *AndroidGeneralDeviceConfiguration) SetAppsHideList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsHideList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsHideList field.

### GetSecurityRequireVerifyApps

`func (o *AndroidGeneralDeviceConfiguration) GetSecurityRequireVerifyApps() bool`

GetSecurityRequireVerifyApps returns the SecurityRequireVerifyApps field if non-nil, zero value otherwise.

### GetSecurityRequireVerifyAppsOk

`func (o *AndroidGeneralDeviceConfiguration) GetSecurityRequireVerifyAppsOk() (bool, bool)`

GetSecurityRequireVerifyAppsOk returns a tuple with the SecurityRequireVerifyApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireVerifyApps

`func (o *AndroidGeneralDeviceConfiguration) HasSecurityRequireVerifyApps() bool`

HasSecurityRequireVerifyApps returns a boolean if a field has been set.

### SetSecurityRequireVerifyApps

`func (o *AndroidGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(v bool)`

SetSecurityRequireVerifyApps gets a reference to the given bool and assigns it to the SecurityRequireVerifyApps field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


