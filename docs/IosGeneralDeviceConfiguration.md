# IosGeneralDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBlockModification** | Pointer to **bool** | Indicates whether or not to allow account modification when the device is in supervised mode. | [optional] 
**ActivationLockAllowWhenSupervised** | Pointer to **bool** | Indicates whether or not to allow activation lock when the device is in the supervised mode. | [optional] 
**AirDropBlocked** | Pointer to **bool** | Indicates whether or not to allow AirDrop when the device is in supervised mode. | [optional] 
**AirDropForceUnmanagedDropTarget** | Pointer to **bool** | Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later). | [optional] 
**AirPlayForcePairingPasswordForOutgoingRequests** | Pointer to **bool** | Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password. | [optional] 
**AppleWatchBlockPairing** | Pointer to **bool** | Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later). | [optional] 
**AppleWatchForceWristDetection** | Pointer to **bool** | Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later). | [optional] 
**AppleNewsBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later). | [optional] 
**AppsSingleAppModeList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements. | [optional] 
**AppsVisibilityList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements. | [optional] 
**AppsVisibilityListType** | Pointer to [**AnyOfmicrosoftGraphAppListType**](anyOf&lt;microsoft.graph.appListType&gt;.md) | Type of list that is in the AppsVisibilityList. | [optional] 
**AppStoreBlockAutomaticDownloads** | Pointer to **bool** | Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later). | [optional] 
**AppStoreBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using the App Store. | [optional] 
**AppStoreBlockInAppPurchases** | Pointer to **bool** | Indicates whether or not to block the user from making in app purchases. | [optional] 
**AppStoreBlockUIAppInstallation** | Pointer to **bool** | Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later). | [optional] 
**AppStoreRequirePassword** | Pointer to **bool** | Indicates whether or not to require a password when using the app store. | [optional] 
**BluetoothBlockModification** | Pointer to **bool** | Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later). | [optional] 
**CameraBlocked** | Pointer to **bool** | Indicates whether or not to block the user from accessing the camera of the device. | [optional] 
**CellularBlockDataRoaming** | Pointer to **bool** | Indicates whether or not to block data roaming. | [optional] 
**CellularBlockGlobalBackgroundFetchWhileRoaming** | Pointer to **bool** | Indicates whether or not to block global background fetch while roaming. | [optional] 
**CellularBlockPerAppDataModification** | Pointer to **bool** | Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode. | [optional] 
**CellularBlockPersonalHotspot** | Pointer to **bool** | Indicates whether or not to block Personal Hotspot. | [optional] 
**CellularBlockVoiceRoaming** | Pointer to **bool** | Indicates whether or not to block voice roaming. | [optional] 
**CertificatesBlockUntrustedTlsCertificates** | Pointer to **bool** | Indicates whether or not to block untrusted TLS certificates. | [optional] 
**ClassroomAppBlockRemoteScreenObservation** | Pointer to **bool** | Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later). | [optional] 
**ClassroomAppForceUnpromptedScreenObservation** | Pointer to **bool** | Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student&#39;s screen without prompting when the device is in supervised mode. | [optional] 
**CompliantAppsList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements. | [optional] 
**CompliantAppListType** | Pointer to [**AnyOfmicrosoftGraphAppListType**](anyOf&lt;microsoft.graph.appListType&gt;.md) | List that is in the AppComplianceList. | [optional] 
**ConfigurationProfileBlockChanges** | Pointer to **bool** | Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode. | [optional] 
**DefinitionLookupBlocked** | Pointer to **bool** | Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ). | [optional] 
**DeviceBlockEnableRestrictions** | Pointer to **bool** | Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode. | [optional] 
**DeviceBlockEraseContentAndSettings** | Pointer to **bool** | Indicates whether or not to allow the use of the &#39;Erase all content and settings&#39; option on the device when the device is in supervised mode. | [optional] 
**DeviceBlockNameModification** | Pointer to **bool** | Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later). | [optional] 
**DiagnosticDataBlockSubmission** | Pointer to **bool** | Indicates whether or not to block diagnostic data submission. | [optional] 
**DiagnosticDataBlockSubmissionModification** | Pointer to **bool** | Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later). | [optional] 
**DocumentsBlockManagedDocumentsInUnmanagedApps** | Pointer to **bool** | Indicates whether or not to block the user from viewing managed documents in unmanaged apps. | [optional] 
**DocumentsBlockUnmanagedDocumentsInManagedApps** | Pointer to **bool** | Indicates whether or not to block the user from viewing unmanaged documents in managed apps. | [optional] 
**EmailInDomainSuffixes** | Pointer to **[]string** | An email address lacking a suffix that matches any of these strings will be considered out-of-domain. | [optional] 
**EnterpriseAppBlockTrust** | Pointer to **bool** | Indicates whether or not to block the user from trusting an enterprise app. | [optional] 
**EnterpriseAppBlockTrustModification** | Pointer to **bool** | Indicates whether or not to block the user from modifying the enterprise app trust settings. | [optional] 
**FaceTimeBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using FaceTime. | [optional] 
**FindMyFriendsBlocked** | Pointer to **bool** | Indicates whether or not to block Find My Friends when the device is in supervised mode. | [optional] 
**GamingBlockGameCenterFriends** | Pointer to **bool** | Indicates whether or not to block the user from having friends in Game Center. | [optional] 
**GamingBlockMultiplayer** | Pointer to **bool** | Indicates whether or not to block the user from using multiplayer gaming. | [optional] 
**GameCenterBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using Game Center when the device is in supervised mode. | [optional] 
**HostPairingBlocked** | Pointer to **bool** | indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode. | [optional] 
**IBooksStoreBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode. | [optional] 
**IBooksStoreBlockErotica** | Pointer to **bool** | Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica. | [optional] 
**ICloudBlockActivityContinuation** | Pointer to **bool** | Indicates whether or not to block  the the user from continuing work they started on iOS device to another iOS or macOS device. | [optional] 
**ICloudBlockBackup** | Pointer to **bool** | Indicates whether or not to block iCloud backup. | [optional] 
**ICloudBlockDocumentSync** | Pointer to **bool** | Indicates whether or not to block iCloud document sync. | [optional] 
**ICloudBlockManagedAppsSync** | Pointer to **bool** | Indicates whether or not to block Managed Apps Cloud Sync. | [optional] 
**ICloudBlockPhotoLibrary** | Pointer to **bool** | Indicates whether or not to block iCloud Photo Library. | [optional] 
**ICloudBlockPhotoStreamSync** | Pointer to **bool** | Indicates whether or not to block iCloud Photo Stream Sync. | [optional] 
**ICloudBlockSharedPhotoStream** | Pointer to **bool** | Indicates whether or not to block Shared Photo Stream. | [optional] 
**ICloudRequireEncryptedBackup** | Pointer to **bool** | Indicates whether or not to require backups to iCloud be encrypted. | [optional] 
**ITunesBlockExplicitContent** | Pointer to **bool** | Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. | [optional] 
**ITunesBlockMusicService** | Pointer to **bool** | Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later). | [optional] 
**ITunesBlockRadio** | Pointer to **bool** | Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later). | [optional] 
**KeyboardBlockAutoCorrect** | Pointer to **bool** | Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later). | [optional] 
**KeyboardBlockDictation** | Pointer to **bool** | Indicates whether or not to block the user from using dictation input when the device is in supervised mode. | [optional] 
**KeyboardBlockPredictive** | Pointer to **bool** | Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later). | [optional] 
**KeyboardBlockShortcuts** | Pointer to **bool** | Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later). | [optional] 
**KeyboardBlockSpellCheck** | Pointer to **bool** | Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later). | [optional] 
**KioskModeAllowAssistiveSpeak** | Pointer to **bool** | Indicates whether or not to allow assistive speak while in kiosk mode. | [optional] 
**KioskModeAllowAssistiveTouchSettings** | Pointer to **bool** | Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode. | [optional] 
**KioskModeAllowAutoLock** | Pointer to **bool** | Indicates whether or not to allow device auto lock while in kiosk mode. | [optional] 
**KioskModeAllowColorInversionSettings** | Pointer to **bool** | Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode. | [optional] 
**KioskModeAllowRingerSwitch** | Pointer to **bool** | Indicates whether or not to allow use of the ringer switch while in kiosk mode. | [optional] 
**KioskModeAllowScreenRotation** | Pointer to **bool** | Indicates whether or not to allow screen rotation while in kiosk mode. | [optional] 
**KioskModeAllowSleepButton** | Pointer to **bool** | Indicates whether or not to allow use of the sleep button while in kiosk mode. | [optional] 
**KioskModeAllowTouchscreen** | Pointer to **bool** | Indicates whether or not to allow use of the touchscreen while in kiosk mode. | [optional] 
**KioskModeAllowVoiceOverSettings** | Pointer to **bool** | Indicates whether or not to allow access to the voice over settings while in kiosk mode. | [optional] 
**KioskModeAllowVolumeButtons** | Pointer to **bool** | Indicates whether or not to allow use of the volume buttons while in kiosk mode. | [optional] 
**KioskModeAllowZoomSettings** | Pointer to **bool** | Indicates whether or not to allow access to the zoom settings while in kiosk mode. | [optional] 
**KioskModeAppStoreUrl** | Pointer to **string** | URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known. | [optional] 
**KioskModeBuiltInAppId** | Pointer to **string** | ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set. | [optional] 
**KioskModeRequireAssistiveTouch** | Pointer to **bool** | Indicates whether or not to require assistive touch while in kiosk mode. | [optional] 
**KioskModeRequireColorInversion** | Pointer to **bool** | Indicates whether or not to require color inversion while in kiosk mode. | [optional] 
**KioskModeRequireMonoAudio** | Pointer to **bool** | Indicates whether or not to require mono audio while in kiosk mode. | [optional] 
**KioskModeRequireVoiceOver** | Pointer to **bool** | Indicates whether or not to require voice over while in kiosk mode. | [optional] 
**KioskModeRequireZoom** | Pointer to **bool** | Indicates whether or not to require zoom while in kiosk mode. | [optional] 
**KioskModeManagedAppId** | Pointer to **string** | Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored. | [optional] 
**LockScreenBlockControlCenter** | Pointer to **bool** | Indicates whether or not to block the user from using control center on the lock screen. | [optional] 
**LockScreenBlockNotificationView** | Pointer to **bool** | Indicates whether or not to block the user from using the notification view on the lock screen. | [optional] 
**LockScreenBlockPassbook** | Pointer to **bool** | Indicates whether or not to block the user from using passbook when the device is locked. | [optional] 
**LockScreenBlockTodayView** | Pointer to **bool** | Indicates whether or not to block the user from using the Today View on the lock screen. | [optional] 
**MediaContentRatingAustralia** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingAustralia**](anyOf&lt;microsoft.graph.mediaContentRatingAustralia&gt;.md) | Media content rating settings for Australia | [optional] 
**MediaContentRatingCanada** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingCanada**](anyOf&lt;microsoft.graph.mediaContentRatingCanada&gt;.md) | Media content rating settings for Canada | [optional] 
**MediaContentRatingFrance** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingFrance**](anyOf&lt;microsoft.graph.mediaContentRatingFrance&gt;.md) | Media content rating settings for France | [optional] 
**MediaContentRatingGermany** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingGermany**](anyOf&lt;microsoft.graph.mediaContentRatingGermany&gt;.md) | Media content rating settings for Germany | [optional] 
**MediaContentRatingIreland** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingIreland**](anyOf&lt;microsoft.graph.mediaContentRatingIreland&gt;.md) | Media content rating settings for Ireland | [optional] 
**MediaContentRatingJapan** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingJapan**](anyOf&lt;microsoft.graph.mediaContentRatingJapan&gt;.md) | Media content rating settings for Japan | [optional] 
**MediaContentRatingNewZealand** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingNewZealand**](anyOf&lt;microsoft.graph.mediaContentRatingNewZealand&gt;.md) | Media content rating settings for New Zealand | [optional] 
**MediaContentRatingUnitedKingdom** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom**](anyOf&lt;microsoft.graph.mediaContentRatingUnitedKingdom&gt;.md) | Media content rating settings for United Kingdom | [optional] 
**MediaContentRatingUnitedStates** | Pointer to [**AnyOfmicrosoftGraphMediaContentRatingUnitedStates**](anyOf&lt;microsoft.graph.mediaContentRatingUnitedStates&gt;.md) | Media content rating settings for United States | [optional] 
**NetworkUsageRules** | Pointer to [**[]AnyOfmicrosoftGraphIosNetworkUsageRule**](anyOf&lt;microsoft.graph.iosNetworkUsageRule&gt;.md) | List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements. | [optional] 
**MediaContentRatingApps** | Pointer to [**AnyOfmicrosoftGraphRatingAppsType**](anyOf&lt;microsoft.graph.ratingAppsType&gt;.md) | Media content rating settings for Apps | [optional] 
**MessagesBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using the Messages app on the supervised device. | [optional] 
**NotificationsBlockSettingsModification** | Pointer to **bool** | Indicates whether or not to allow notifications settings modification (iOS 9.3 and later). | [optional] 
**PasscodeBlockFingerprintUnlock** | Pointer to **bool** | Indicates whether or not to block fingerprint unlock. | [optional] 
**PasscodeBlockFingerprintModification** | Pointer to **bool** | Block modification of registered Touch ID fingerprints when in supervised mode. | [optional] 
**PasscodeBlockModification** | Pointer to **bool** | Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later). | [optional] 
**PasscodeBlockSimple** | Pointer to **bool** | Indicates whether or not to block simple passcodes. | [optional] 
**PasscodeExpirationDays** | Pointer to **int32** | Number of days before the passcode expires. Valid values 1 to 65535 | [optional] 
**PasscodeMinimumLength** | Pointer to **int32** | Minimum length of passcode. Valid values 4 to 14 | [optional] 
**PasscodeMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a passcode is required. | [optional] 
**PasscodeMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before the screen times out. | [optional] 
**PasscodeMinimumCharacterSetCount** | Pointer to **int32** | Number of character sets a passcode must contain. Valid values 0 to 4 | [optional] 
**PasscodePreviousPasscodeBlockCount** | Pointer to **int32** | Number of previous passcodes to block. Valid values 1 to 24 | [optional] 
**PasscodeSignInFailureCountBeforeWipe** | Pointer to **int32** | Number of sign in failures allowed before wiping the device. Valid values 4 to 11 | [optional] 
**PasscodeRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | Type of passcode that is required. | [optional] 
**PasscodeRequired** | Pointer to **bool** | Indicates whether or not to require a passcode. | [optional] 
**PodcastsBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later). | [optional] 
**SafariBlockAutofill** | Pointer to **bool** | Indicates whether or not to block the user from using Auto fill in Safari. | [optional] 
**SafariBlockJavaScript** | Pointer to **bool** | Indicates whether or not to block JavaScript in Safari. | [optional] 
**SafariBlockPopups** | Pointer to **bool** | Indicates whether or not to block popups in Safari. | [optional] 
**SafariBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using Safari. | [optional] 
**SafariCookieSettings** | Pointer to [**AnyOfmicrosoftGraphWebBrowserCookieSettings**](anyOf&lt;microsoft.graph.webBrowserCookieSettings&gt;.md) | Cookie settings for Safari. | [optional] 
**SafariManagedDomains** | Pointer to **[]string** | URLs matching the patterns listed here will be considered managed. | [optional] 
**SafariPasswordAutoFillDomains** | Pointer to **[]string** | Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later). | [optional] 
**SafariRequireFraudWarning** | Pointer to **bool** | Indicates whether or not to require fraud warning in Safari. | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether or not to block the user from taking Screenshots. | [optional] 
**SiriBlocked** | Pointer to **bool** | Indicates whether or not to block the user from using Siri. | [optional] 
**SiriBlockedWhenLocked** | Pointer to **bool** | Indicates whether or not to block the user from using Siri when locked. | [optional] 
**SiriBlockUserGeneratedContent** | Pointer to **bool** | Indicates whether or not to block Siri from querying user-generated content when used on a supervised device. | [optional] 
**SiriRequireProfanityFilter** | Pointer to **bool** | Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device. | [optional] 
**SpotlightBlockInternetResults** | Pointer to **bool** | Indicates whether or not to block Spotlight search from returning internet results on supervised device. | [optional] 
**VoiceDialingBlocked** | Pointer to **bool** | Indicates whether or not to block voice dialing. | [optional] 
**WallpaperBlockModification** | Pointer to **bool** | Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) . | [optional] 
**WiFiConnectOnlyToConfiguredNetworks** | Pointer to **bool** | Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. | [optional] 

## Methods

### GetAccountBlockModification

`func (o *IosGeneralDeviceConfiguration) GetAccountBlockModification() bool`

GetAccountBlockModification returns the AccountBlockModification field if non-nil, zero value otherwise.

### GetAccountBlockModificationOk

`func (o *IosGeneralDeviceConfiguration) GetAccountBlockModificationOk() (bool, bool)`

GetAccountBlockModificationOk returns a tuple with the AccountBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountBlockModification

`func (o *IosGeneralDeviceConfiguration) HasAccountBlockModification() bool`

HasAccountBlockModification returns a boolean if a field has been set.

### SetAccountBlockModification

`func (o *IosGeneralDeviceConfiguration) SetAccountBlockModification(v bool)`

SetAccountBlockModification gets a reference to the given bool and assigns it to the AccountBlockModification field.

### GetActivationLockAllowWhenSupervised

`func (o *IosGeneralDeviceConfiguration) GetActivationLockAllowWhenSupervised() bool`

GetActivationLockAllowWhenSupervised returns the ActivationLockAllowWhenSupervised field if non-nil, zero value otherwise.

### GetActivationLockAllowWhenSupervisedOk

`func (o *IosGeneralDeviceConfiguration) GetActivationLockAllowWhenSupervisedOk() (bool, bool)`

GetActivationLockAllowWhenSupervisedOk returns a tuple with the ActivationLockAllowWhenSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationLockAllowWhenSupervised

`func (o *IosGeneralDeviceConfiguration) HasActivationLockAllowWhenSupervised() bool`

HasActivationLockAllowWhenSupervised returns a boolean if a field has been set.

### SetActivationLockAllowWhenSupervised

`func (o *IosGeneralDeviceConfiguration) SetActivationLockAllowWhenSupervised(v bool)`

SetActivationLockAllowWhenSupervised gets a reference to the given bool and assigns it to the ActivationLockAllowWhenSupervised field.

### GetAirDropBlocked

`func (o *IosGeneralDeviceConfiguration) GetAirDropBlocked() bool`

GetAirDropBlocked returns the AirDropBlocked field if non-nil, zero value otherwise.

### GetAirDropBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetAirDropBlockedOk() (bool, bool)`

GetAirDropBlockedOk returns a tuple with the AirDropBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAirDropBlocked

`func (o *IosGeneralDeviceConfiguration) HasAirDropBlocked() bool`

HasAirDropBlocked returns a boolean if a field has been set.

### SetAirDropBlocked

`func (o *IosGeneralDeviceConfiguration) SetAirDropBlocked(v bool)`

SetAirDropBlocked gets a reference to the given bool and assigns it to the AirDropBlocked field.

### GetAirDropForceUnmanagedDropTarget

`func (o *IosGeneralDeviceConfiguration) GetAirDropForceUnmanagedDropTarget() bool`

GetAirDropForceUnmanagedDropTarget returns the AirDropForceUnmanagedDropTarget field if non-nil, zero value otherwise.

### GetAirDropForceUnmanagedDropTargetOk

`func (o *IosGeneralDeviceConfiguration) GetAirDropForceUnmanagedDropTargetOk() (bool, bool)`

GetAirDropForceUnmanagedDropTargetOk returns a tuple with the AirDropForceUnmanagedDropTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAirDropForceUnmanagedDropTarget

`func (o *IosGeneralDeviceConfiguration) HasAirDropForceUnmanagedDropTarget() bool`

HasAirDropForceUnmanagedDropTarget returns a boolean if a field has been set.

### SetAirDropForceUnmanagedDropTarget

`func (o *IosGeneralDeviceConfiguration) SetAirDropForceUnmanagedDropTarget(v bool)`

SetAirDropForceUnmanagedDropTarget gets a reference to the given bool and assigns it to the AirDropForceUnmanagedDropTarget field.

### GetAirPlayForcePairingPasswordForOutgoingRequests

`func (o *IosGeneralDeviceConfiguration) GetAirPlayForcePairingPasswordForOutgoingRequests() bool`

GetAirPlayForcePairingPasswordForOutgoingRequests returns the AirPlayForcePairingPasswordForOutgoingRequests field if non-nil, zero value otherwise.

### GetAirPlayForcePairingPasswordForOutgoingRequestsOk

`func (o *IosGeneralDeviceConfiguration) GetAirPlayForcePairingPasswordForOutgoingRequestsOk() (bool, bool)`

GetAirPlayForcePairingPasswordForOutgoingRequestsOk returns a tuple with the AirPlayForcePairingPasswordForOutgoingRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAirPlayForcePairingPasswordForOutgoingRequests

`func (o *IosGeneralDeviceConfiguration) HasAirPlayForcePairingPasswordForOutgoingRequests() bool`

HasAirPlayForcePairingPasswordForOutgoingRequests returns a boolean if a field has been set.

### SetAirPlayForcePairingPasswordForOutgoingRequests

`func (o *IosGeneralDeviceConfiguration) SetAirPlayForcePairingPasswordForOutgoingRequests(v bool)`

SetAirPlayForcePairingPasswordForOutgoingRequests gets a reference to the given bool and assigns it to the AirPlayForcePairingPasswordForOutgoingRequests field.

### GetAppleWatchBlockPairing

`func (o *IosGeneralDeviceConfiguration) GetAppleWatchBlockPairing() bool`

GetAppleWatchBlockPairing returns the AppleWatchBlockPairing field if non-nil, zero value otherwise.

### GetAppleWatchBlockPairingOk

`func (o *IosGeneralDeviceConfiguration) GetAppleWatchBlockPairingOk() (bool, bool)`

GetAppleWatchBlockPairingOk returns a tuple with the AppleWatchBlockPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleWatchBlockPairing

`func (o *IosGeneralDeviceConfiguration) HasAppleWatchBlockPairing() bool`

HasAppleWatchBlockPairing returns a boolean if a field has been set.

### SetAppleWatchBlockPairing

`func (o *IosGeneralDeviceConfiguration) SetAppleWatchBlockPairing(v bool)`

SetAppleWatchBlockPairing gets a reference to the given bool and assigns it to the AppleWatchBlockPairing field.

### GetAppleWatchForceWristDetection

`func (o *IosGeneralDeviceConfiguration) GetAppleWatchForceWristDetection() bool`

GetAppleWatchForceWristDetection returns the AppleWatchForceWristDetection field if non-nil, zero value otherwise.

### GetAppleWatchForceWristDetectionOk

`func (o *IosGeneralDeviceConfiguration) GetAppleWatchForceWristDetectionOk() (bool, bool)`

GetAppleWatchForceWristDetectionOk returns a tuple with the AppleWatchForceWristDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleWatchForceWristDetection

`func (o *IosGeneralDeviceConfiguration) HasAppleWatchForceWristDetection() bool`

HasAppleWatchForceWristDetection returns a boolean if a field has been set.

### SetAppleWatchForceWristDetection

`func (o *IosGeneralDeviceConfiguration) SetAppleWatchForceWristDetection(v bool)`

SetAppleWatchForceWristDetection gets a reference to the given bool and assigns it to the AppleWatchForceWristDetection field.

### GetAppleNewsBlocked

`func (o *IosGeneralDeviceConfiguration) GetAppleNewsBlocked() bool`

GetAppleNewsBlocked returns the AppleNewsBlocked field if non-nil, zero value otherwise.

### GetAppleNewsBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetAppleNewsBlockedOk() (bool, bool)`

GetAppleNewsBlockedOk returns a tuple with the AppleNewsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleNewsBlocked

`func (o *IosGeneralDeviceConfiguration) HasAppleNewsBlocked() bool`

HasAppleNewsBlocked returns a boolean if a field has been set.

### SetAppleNewsBlocked

`func (o *IosGeneralDeviceConfiguration) SetAppleNewsBlocked(v bool)`

SetAppleNewsBlocked gets a reference to the given bool and assigns it to the AppleNewsBlocked field.

### GetAppsSingleAppModeList

`func (o *IosGeneralDeviceConfiguration) GetAppsSingleAppModeList() []AnyOfmicrosoftGraphAppListItem`

GetAppsSingleAppModeList returns the AppsSingleAppModeList field if non-nil, zero value otherwise.

### GetAppsSingleAppModeListOk

`func (o *IosGeneralDeviceConfiguration) GetAppsSingleAppModeListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsSingleAppModeListOk returns a tuple with the AppsSingleAppModeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsSingleAppModeList

`func (o *IosGeneralDeviceConfiguration) HasAppsSingleAppModeList() bool`

HasAppsSingleAppModeList returns a boolean if a field has been set.

### SetAppsSingleAppModeList

`func (o *IosGeneralDeviceConfiguration) SetAppsSingleAppModeList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsSingleAppModeList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsSingleAppModeList field.

### GetAppsVisibilityList

`func (o *IosGeneralDeviceConfiguration) GetAppsVisibilityList() []AnyOfmicrosoftGraphAppListItem`

GetAppsVisibilityList returns the AppsVisibilityList field if non-nil, zero value otherwise.

### GetAppsVisibilityListOk

`func (o *IosGeneralDeviceConfiguration) GetAppsVisibilityListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsVisibilityListOk returns a tuple with the AppsVisibilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsVisibilityList

`func (o *IosGeneralDeviceConfiguration) HasAppsVisibilityList() bool`

HasAppsVisibilityList returns a boolean if a field has been set.

### SetAppsVisibilityList

`func (o *IosGeneralDeviceConfiguration) SetAppsVisibilityList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsVisibilityList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsVisibilityList field.

### GetAppsVisibilityListType

`func (o *IosGeneralDeviceConfiguration) GetAppsVisibilityListType() AnyOfmicrosoftGraphAppListType`

GetAppsVisibilityListType returns the AppsVisibilityListType field if non-nil, zero value otherwise.

### GetAppsVisibilityListTypeOk

`func (o *IosGeneralDeviceConfiguration) GetAppsVisibilityListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetAppsVisibilityListTypeOk returns a tuple with the AppsVisibilityListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsVisibilityListType

`func (o *IosGeneralDeviceConfiguration) HasAppsVisibilityListType() bool`

HasAppsVisibilityListType returns a boolean if a field has been set.

### SetAppsVisibilityListType

`func (o *IosGeneralDeviceConfiguration) SetAppsVisibilityListType(v AnyOfmicrosoftGraphAppListType)`

SetAppsVisibilityListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the AppsVisibilityListType field.

### GetAppStoreBlockAutomaticDownloads

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockAutomaticDownloads() bool`

GetAppStoreBlockAutomaticDownloads returns the AppStoreBlockAutomaticDownloads field if non-nil, zero value otherwise.

### GetAppStoreBlockAutomaticDownloadsOk

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockAutomaticDownloadsOk() (bool, bool)`

GetAppStoreBlockAutomaticDownloadsOk returns a tuple with the AppStoreBlockAutomaticDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlockAutomaticDownloads

`func (o *IosGeneralDeviceConfiguration) HasAppStoreBlockAutomaticDownloads() bool`

HasAppStoreBlockAutomaticDownloads returns a boolean if a field has been set.

### SetAppStoreBlockAutomaticDownloads

`func (o *IosGeneralDeviceConfiguration) SetAppStoreBlockAutomaticDownloads(v bool)`

SetAppStoreBlockAutomaticDownloads gets a reference to the given bool and assigns it to the AppStoreBlockAutomaticDownloads field.

### GetAppStoreBlocked

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlocked() bool`

GetAppStoreBlocked returns the AppStoreBlocked field if non-nil, zero value otherwise.

### GetAppStoreBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockedOk() (bool, bool)`

GetAppStoreBlockedOk returns a tuple with the AppStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlocked

`func (o *IosGeneralDeviceConfiguration) HasAppStoreBlocked() bool`

HasAppStoreBlocked returns a boolean if a field has been set.

### SetAppStoreBlocked

`func (o *IosGeneralDeviceConfiguration) SetAppStoreBlocked(v bool)`

SetAppStoreBlocked gets a reference to the given bool and assigns it to the AppStoreBlocked field.

### GetAppStoreBlockInAppPurchases

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockInAppPurchases() bool`

GetAppStoreBlockInAppPurchases returns the AppStoreBlockInAppPurchases field if non-nil, zero value otherwise.

### GetAppStoreBlockInAppPurchasesOk

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockInAppPurchasesOk() (bool, bool)`

GetAppStoreBlockInAppPurchasesOk returns a tuple with the AppStoreBlockInAppPurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlockInAppPurchases

`func (o *IosGeneralDeviceConfiguration) HasAppStoreBlockInAppPurchases() bool`

HasAppStoreBlockInAppPurchases returns a boolean if a field has been set.

### SetAppStoreBlockInAppPurchases

`func (o *IosGeneralDeviceConfiguration) SetAppStoreBlockInAppPurchases(v bool)`

SetAppStoreBlockInAppPurchases gets a reference to the given bool and assigns it to the AppStoreBlockInAppPurchases field.

### GetAppStoreBlockUIAppInstallation

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockUIAppInstallation() bool`

GetAppStoreBlockUIAppInstallation returns the AppStoreBlockUIAppInstallation field if non-nil, zero value otherwise.

### GetAppStoreBlockUIAppInstallationOk

`func (o *IosGeneralDeviceConfiguration) GetAppStoreBlockUIAppInstallationOk() (bool, bool)`

GetAppStoreBlockUIAppInstallationOk returns a tuple with the AppStoreBlockUIAppInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlockUIAppInstallation

`func (o *IosGeneralDeviceConfiguration) HasAppStoreBlockUIAppInstallation() bool`

HasAppStoreBlockUIAppInstallation returns a boolean if a field has been set.

### SetAppStoreBlockUIAppInstallation

`func (o *IosGeneralDeviceConfiguration) SetAppStoreBlockUIAppInstallation(v bool)`

SetAppStoreBlockUIAppInstallation gets a reference to the given bool and assigns it to the AppStoreBlockUIAppInstallation field.

### GetAppStoreRequirePassword

`func (o *IosGeneralDeviceConfiguration) GetAppStoreRequirePassword() bool`

GetAppStoreRequirePassword returns the AppStoreRequirePassword field if non-nil, zero value otherwise.

### GetAppStoreRequirePasswordOk

`func (o *IosGeneralDeviceConfiguration) GetAppStoreRequirePasswordOk() (bool, bool)`

GetAppStoreRequirePasswordOk returns a tuple with the AppStoreRequirePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreRequirePassword

`func (o *IosGeneralDeviceConfiguration) HasAppStoreRequirePassword() bool`

HasAppStoreRequirePassword returns a boolean if a field has been set.

### SetAppStoreRequirePassword

`func (o *IosGeneralDeviceConfiguration) SetAppStoreRequirePassword(v bool)`

SetAppStoreRequirePassword gets a reference to the given bool and assigns it to the AppStoreRequirePassword field.

### GetBluetoothBlockModification

`func (o *IosGeneralDeviceConfiguration) GetBluetoothBlockModification() bool`

GetBluetoothBlockModification returns the BluetoothBlockModification field if non-nil, zero value otherwise.

### GetBluetoothBlockModificationOk

`func (o *IosGeneralDeviceConfiguration) GetBluetoothBlockModificationOk() (bool, bool)`

GetBluetoothBlockModificationOk returns a tuple with the BluetoothBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockModification

`func (o *IosGeneralDeviceConfiguration) HasBluetoothBlockModification() bool`

HasBluetoothBlockModification returns a boolean if a field has been set.

### SetBluetoothBlockModification

`func (o *IosGeneralDeviceConfiguration) SetBluetoothBlockModification(v bool)`

SetBluetoothBlockModification gets a reference to the given bool and assigns it to the BluetoothBlockModification field.

### GetCameraBlocked

`func (o *IosGeneralDeviceConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *IosGeneralDeviceConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *IosGeneralDeviceConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetCellularBlockDataRoaming

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockDataRoaming() bool`

GetCellularBlockDataRoaming returns the CellularBlockDataRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataRoamingOk

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockDataRoamingOk() (bool, bool)`

GetCellularBlockDataRoamingOk returns a tuple with the CellularBlockDataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataRoaming

`func (o *IosGeneralDeviceConfiguration) HasCellularBlockDataRoaming() bool`

HasCellularBlockDataRoaming returns a boolean if a field has been set.

### SetCellularBlockDataRoaming

`func (o *IosGeneralDeviceConfiguration) SetCellularBlockDataRoaming(v bool)`

SetCellularBlockDataRoaming gets a reference to the given bool and assigns it to the CellularBlockDataRoaming field.

### GetCellularBlockGlobalBackgroundFetchWhileRoaming

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockGlobalBackgroundFetchWhileRoaming() bool`

GetCellularBlockGlobalBackgroundFetchWhileRoaming returns the CellularBlockGlobalBackgroundFetchWhileRoaming field if non-nil, zero value otherwise.

### GetCellularBlockGlobalBackgroundFetchWhileRoamingOk

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockGlobalBackgroundFetchWhileRoamingOk() (bool, bool)`

GetCellularBlockGlobalBackgroundFetchWhileRoamingOk returns a tuple with the CellularBlockGlobalBackgroundFetchWhileRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockGlobalBackgroundFetchWhileRoaming

`func (o *IosGeneralDeviceConfiguration) HasCellularBlockGlobalBackgroundFetchWhileRoaming() bool`

HasCellularBlockGlobalBackgroundFetchWhileRoaming returns a boolean if a field has been set.

### SetCellularBlockGlobalBackgroundFetchWhileRoaming

`func (o *IosGeneralDeviceConfiguration) SetCellularBlockGlobalBackgroundFetchWhileRoaming(v bool)`

SetCellularBlockGlobalBackgroundFetchWhileRoaming gets a reference to the given bool and assigns it to the CellularBlockGlobalBackgroundFetchWhileRoaming field.

### GetCellularBlockPerAppDataModification

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockPerAppDataModification() bool`

GetCellularBlockPerAppDataModification returns the CellularBlockPerAppDataModification field if non-nil, zero value otherwise.

### GetCellularBlockPerAppDataModificationOk

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockPerAppDataModificationOk() (bool, bool)`

GetCellularBlockPerAppDataModificationOk returns a tuple with the CellularBlockPerAppDataModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockPerAppDataModification

`func (o *IosGeneralDeviceConfiguration) HasCellularBlockPerAppDataModification() bool`

HasCellularBlockPerAppDataModification returns a boolean if a field has been set.

### SetCellularBlockPerAppDataModification

`func (o *IosGeneralDeviceConfiguration) SetCellularBlockPerAppDataModification(v bool)`

SetCellularBlockPerAppDataModification gets a reference to the given bool and assigns it to the CellularBlockPerAppDataModification field.

### GetCellularBlockPersonalHotspot

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspot() bool`

GetCellularBlockPersonalHotspot returns the CellularBlockPersonalHotspot field if non-nil, zero value otherwise.

### GetCellularBlockPersonalHotspotOk

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspotOk() (bool, bool)`

GetCellularBlockPersonalHotspotOk returns a tuple with the CellularBlockPersonalHotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockPersonalHotspot

`func (o *IosGeneralDeviceConfiguration) HasCellularBlockPersonalHotspot() bool`

HasCellularBlockPersonalHotspot returns a boolean if a field has been set.

### SetCellularBlockPersonalHotspot

`func (o *IosGeneralDeviceConfiguration) SetCellularBlockPersonalHotspot(v bool)`

SetCellularBlockPersonalHotspot gets a reference to the given bool and assigns it to the CellularBlockPersonalHotspot field.

### GetCellularBlockVoiceRoaming

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming() bool`

GetCellularBlockVoiceRoaming returns the CellularBlockVoiceRoaming field if non-nil, zero value otherwise.

### GetCellularBlockVoiceRoamingOk

`func (o *IosGeneralDeviceConfiguration) GetCellularBlockVoiceRoamingOk() (bool, bool)`

GetCellularBlockVoiceRoamingOk returns a tuple with the CellularBlockVoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVoiceRoaming

`func (o *IosGeneralDeviceConfiguration) HasCellularBlockVoiceRoaming() bool`

HasCellularBlockVoiceRoaming returns a boolean if a field has been set.

### SetCellularBlockVoiceRoaming

`func (o *IosGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(v bool)`

SetCellularBlockVoiceRoaming gets a reference to the given bool and assigns it to the CellularBlockVoiceRoaming field.

### GetCertificatesBlockUntrustedTlsCertificates

`func (o *IosGeneralDeviceConfiguration) GetCertificatesBlockUntrustedTlsCertificates() bool`

GetCertificatesBlockUntrustedTlsCertificates returns the CertificatesBlockUntrustedTlsCertificates field if non-nil, zero value otherwise.

### GetCertificatesBlockUntrustedTlsCertificatesOk

`func (o *IosGeneralDeviceConfiguration) GetCertificatesBlockUntrustedTlsCertificatesOk() (bool, bool)`

GetCertificatesBlockUntrustedTlsCertificatesOk returns a tuple with the CertificatesBlockUntrustedTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificatesBlockUntrustedTlsCertificates

`func (o *IosGeneralDeviceConfiguration) HasCertificatesBlockUntrustedTlsCertificates() bool`

HasCertificatesBlockUntrustedTlsCertificates returns a boolean if a field has been set.

### SetCertificatesBlockUntrustedTlsCertificates

`func (o *IosGeneralDeviceConfiguration) SetCertificatesBlockUntrustedTlsCertificates(v bool)`

SetCertificatesBlockUntrustedTlsCertificates gets a reference to the given bool and assigns it to the CertificatesBlockUntrustedTlsCertificates field.

### GetClassroomAppBlockRemoteScreenObservation

`func (o *IosGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservation() bool`

GetClassroomAppBlockRemoteScreenObservation returns the ClassroomAppBlockRemoteScreenObservation field if non-nil, zero value otherwise.

### GetClassroomAppBlockRemoteScreenObservationOk

`func (o *IosGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservationOk() (bool, bool)`

GetClassroomAppBlockRemoteScreenObservationOk returns a tuple with the ClassroomAppBlockRemoteScreenObservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassroomAppBlockRemoteScreenObservation

`func (o *IosGeneralDeviceConfiguration) HasClassroomAppBlockRemoteScreenObservation() bool`

HasClassroomAppBlockRemoteScreenObservation returns a boolean if a field has been set.

### SetClassroomAppBlockRemoteScreenObservation

`func (o *IosGeneralDeviceConfiguration) SetClassroomAppBlockRemoteScreenObservation(v bool)`

SetClassroomAppBlockRemoteScreenObservation gets a reference to the given bool and assigns it to the ClassroomAppBlockRemoteScreenObservation field.

### GetClassroomAppForceUnpromptedScreenObservation

`func (o *IosGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservation() bool`

GetClassroomAppForceUnpromptedScreenObservation returns the ClassroomAppForceUnpromptedScreenObservation field if non-nil, zero value otherwise.

### GetClassroomAppForceUnpromptedScreenObservationOk

`func (o *IosGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservationOk() (bool, bool)`

GetClassroomAppForceUnpromptedScreenObservationOk returns a tuple with the ClassroomAppForceUnpromptedScreenObservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassroomAppForceUnpromptedScreenObservation

`func (o *IosGeneralDeviceConfiguration) HasClassroomAppForceUnpromptedScreenObservation() bool`

HasClassroomAppForceUnpromptedScreenObservation returns a boolean if a field has been set.

### SetClassroomAppForceUnpromptedScreenObservation

`func (o *IosGeneralDeviceConfiguration) SetClassroomAppForceUnpromptedScreenObservation(v bool)`

SetClassroomAppForceUnpromptedScreenObservation gets a reference to the given bool and assigns it to the ClassroomAppForceUnpromptedScreenObservation field.

### GetCompliantAppsList

`func (o *IosGeneralDeviceConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *IosGeneralDeviceConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *IosGeneralDeviceConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *IosGeneralDeviceConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *IosGeneralDeviceConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *IosGeneralDeviceConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *IosGeneralDeviceConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *IosGeneralDeviceConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetConfigurationProfileBlockChanges

`func (o *IosGeneralDeviceConfiguration) GetConfigurationProfileBlockChanges() bool`

GetConfigurationProfileBlockChanges returns the ConfigurationProfileBlockChanges field if non-nil, zero value otherwise.

### GetConfigurationProfileBlockChangesOk

`func (o *IosGeneralDeviceConfiguration) GetConfigurationProfileBlockChangesOk() (bool, bool)`

GetConfigurationProfileBlockChangesOk returns a tuple with the ConfigurationProfileBlockChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationProfileBlockChanges

`func (o *IosGeneralDeviceConfiguration) HasConfigurationProfileBlockChanges() bool`

HasConfigurationProfileBlockChanges returns a boolean if a field has been set.

### SetConfigurationProfileBlockChanges

`func (o *IosGeneralDeviceConfiguration) SetConfigurationProfileBlockChanges(v bool)`

SetConfigurationProfileBlockChanges gets a reference to the given bool and assigns it to the ConfigurationProfileBlockChanges field.

### GetDefinitionLookupBlocked

`func (o *IosGeneralDeviceConfiguration) GetDefinitionLookupBlocked() bool`

GetDefinitionLookupBlocked returns the DefinitionLookupBlocked field if non-nil, zero value otherwise.

### GetDefinitionLookupBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetDefinitionLookupBlockedOk() (bool, bool)`

GetDefinitionLookupBlockedOk returns a tuple with the DefinitionLookupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefinitionLookupBlocked

`func (o *IosGeneralDeviceConfiguration) HasDefinitionLookupBlocked() bool`

HasDefinitionLookupBlocked returns a boolean if a field has been set.

### SetDefinitionLookupBlocked

`func (o *IosGeneralDeviceConfiguration) SetDefinitionLookupBlocked(v bool)`

SetDefinitionLookupBlocked gets a reference to the given bool and assigns it to the DefinitionLookupBlocked field.

### GetDeviceBlockEnableRestrictions

`func (o *IosGeneralDeviceConfiguration) GetDeviceBlockEnableRestrictions() bool`

GetDeviceBlockEnableRestrictions returns the DeviceBlockEnableRestrictions field if non-nil, zero value otherwise.

### GetDeviceBlockEnableRestrictionsOk

`func (o *IosGeneralDeviceConfiguration) GetDeviceBlockEnableRestrictionsOk() (bool, bool)`

GetDeviceBlockEnableRestrictionsOk returns a tuple with the DeviceBlockEnableRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceBlockEnableRestrictions

`func (o *IosGeneralDeviceConfiguration) HasDeviceBlockEnableRestrictions() bool`

HasDeviceBlockEnableRestrictions returns a boolean if a field has been set.

### SetDeviceBlockEnableRestrictions

`func (o *IosGeneralDeviceConfiguration) SetDeviceBlockEnableRestrictions(v bool)`

SetDeviceBlockEnableRestrictions gets a reference to the given bool and assigns it to the DeviceBlockEnableRestrictions field.

### GetDeviceBlockEraseContentAndSettings

`func (o *IosGeneralDeviceConfiguration) GetDeviceBlockEraseContentAndSettings() bool`

GetDeviceBlockEraseContentAndSettings returns the DeviceBlockEraseContentAndSettings field if non-nil, zero value otherwise.

### GetDeviceBlockEraseContentAndSettingsOk

`func (o *IosGeneralDeviceConfiguration) GetDeviceBlockEraseContentAndSettingsOk() (bool, bool)`

GetDeviceBlockEraseContentAndSettingsOk returns a tuple with the DeviceBlockEraseContentAndSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceBlockEraseContentAndSettings

`func (o *IosGeneralDeviceConfiguration) HasDeviceBlockEraseContentAndSettings() bool`

HasDeviceBlockEraseContentAndSettings returns a boolean if a field has been set.

### SetDeviceBlockEraseContentAndSettings

`func (o *IosGeneralDeviceConfiguration) SetDeviceBlockEraseContentAndSettings(v bool)`

SetDeviceBlockEraseContentAndSettings gets a reference to the given bool and assigns it to the DeviceBlockEraseContentAndSettings field.

### GetDeviceBlockNameModification

`func (o *IosGeneralDeviceConfiguration) GetDeviceBlockNameModification() bool`

GetDeviceBlockNameModification returns the DeviceBlockNameModification field if non-nil, zero value otherwise.

### GetDeviceBlockNameModificationOk

`func (o *IosGeneralDeviceConfiguration) GetDeviceBlockNameModificationOk() (bool, bool)`

GetDeviceBlockNameModificationOk returns a tuple with the DeviceBlockNameModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceBlockNameModification

`func (o *IosGeneralDeviceConfiguration) HasDeviceBlockNameModification() bool`

HasDeviceBlockNameModification returns a boolean if a field has been set.

### SetDeviceBlockNameModification

`func (o *IosGeneralDeviceConfiguration) SetDeviceBlockNameModification(v bool)`

SetDeviceBlockNameModification gets a reference to the given bool and assigns it to the DeviceBlockNameModification field.

### GetDiagnosticDataBlockSubmission

`func (o *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission() bool`

GetDiagnosticDataBlockSubmission returns the DiagnosticDataBlockSubmission field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionOk

`func (o *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionOk returns a tuple with the DiagnosticDataBlockSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmission

`func (o *IosGeneralDeviceConfiguration) HasDiagnosticDataBlockSubmission() bool`

HasDiagnosticDataBlockSubmission returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmission

`func (o *IosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(v bool)`

SetDiagnosticDataBlockSubmission gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmission field.

### GetDiagnosticDataBlockSubmissionModification

`func (o *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionModification() bool`

GetDiagnosticDataBlockSubmissionModification returns the DiagnosticDataBlockSubmissionModification field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionModificationOk

`func (o *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionModificationOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionModificationOk returns a tuple with the DiagnosticDataBlockSubmissionModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmissionModification

`func (o *IosGeneralDeviceConfiguration) HasDiagnosticDataBlockSubmissionModification() bool`

HasDiagnosticDataBlockSubmissionModification returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmissionModification

`func (o *IosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmissionModification(v bool)`

SetDiagnosticDataBlockSubmissionModification gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmissionModification field.

### GetDocumentsBlockManagedDocumentsInUnmanagedApps

`func (o *IosGeneralDeviceConfiguration) GetDocumentsBlockManagedDocumentsInUnmanagedApps() bool`

GetDocumentsBlockManagedDocumentsInUnmanagedApps returns the DocumentsBlockManagedDocumentsInUnmanagedApps field if non-nil, zero value otherwise.

### GetDocumentsBlockManagedDocumentsInUnmanagedAppsOk

`func (o *IosGeneralDeviceConfiguration) GetDocumentsBlockManagedDocumentsInUnmanagedAppsOk() (bool, bool)`

GetDocumentsBlockManagedDocumentsInUnmanagedAppsOk returns a tuple with the DocumentsBlockManagedDocumentsInUnmanagedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDocumentsBlockManagedDocumentsInUnmanagedApps

`func (o *IosGeneralDeviceConfiguration) HasDocumentsBlockManagedDocumentsInUnmanagedApps() bool`

HasDocumentsBlockManagedDocumentsInUnmanagedApps returns a boolean if a field has been set.

### SetDocumentsBlockManagedDocumentsInUnmanagedApps

`func (o *IosGeneralDeviceConfiguration) SetDocumentsBlockManagedDocumentsInUnmanagedApps(v bool)`

SetDocumentsBlockManagedDocumentsInUnmanagedApps gets a reference to the given bool and assigns it to the DocumentsBlockManagedDocumentsInUnmanagedApps field.

### GetDocumentsBlockUnmanagedDocumentsInManagedApps

`func (o *IosGeneralDeviceConfiguration) GetDocumentsBlockUnmanagedDocumentsInManagedApps() bool`

GetDocumentsBlockUnmanagedDocumentsInManagedApps returns the DocumentsBlockUnmanagedDocumentsInManagedApps field if non-nil, zero value otherwise.

### GetDocumentsBlockUnmanagedDocumentsInManagedAppsOk

`func (o *IosGeneralDeviceConfiguration) GetDocumentsBlockUnmanagedDocumentsInManagedAppsOk() (bool, bool)`

GetDocumentsBlockUnmanagedDocumentsInManagedAppsOk returns a tuple with the DocumentsBlockUnmanagedDocumentsInManagedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDocumentsBlockUnmanagedDocumentsInManagedApps

`func (o *IosGeneralDeviceConfiguration) HasDocumentsBlockUnmanagedDocumentsInManagedApps() bool`

HasDocumentsBlockUnmanagedDocumentsInManagedApps returns a boolean if a field has been set.

### SetDocumentsBlockUnmanagedDocumentsInManagedApps

`func (o *IosGeneralDeviceConfiguration) SetDocumentsBlockUnmanagedDocumentsInManagedApps(v bool)`

SetDocumentsBlockUnmanagedDocumentsInManagedApps gets a reference to the given bool and assigns it to the DocumentsBlockUnmanagedDocumentsInManagedApps field.

### GetEmailInDomainSuffixes

`func (o *IosGeneralDeviceConfiguration) GetEmailInDomainSuffixes() []string`

GetEmailInDomainSuffixes returns the EmailInDomainSuffixes field if non-nil, zero value otherwise.

### GetEmailInDomainSuffixesOk

`func (o *IosGeneralDeviceConfiguration) GetEmailInDomainSuffixesOk() ([]string, bool)`

GetEmailInDomainSuffixesOk returns a tuple with the EmailInDomainSuffixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailInDomainSuffixes

`func (o *IosGeneralDeviceConfiguration) HasEmailInDomainSuffixes() bool`

HasEmailInDomainSuffixes returns a boolean if a field has been set.

### SetEmailInDomainSuffixes

`func (o *IosGeneralDeviceConfiguration) SetEmailInDomainSuffixes(v []string)`

SetEmailInDomainSuffixes gets a reference to the given []string and assigns it to the EmailInDomainSuffixes field.

### GetEnterpriseAppBlockTrust

`func (o *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrust() bool`

GetEnterpriseAppBlockTrust returns the EnterpriseAppBlockTrust field if non-nil, zero value otherwise.

### GetEnterpriseAppBlockTrustOk

`func (o *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustOk() (bool, bool)`

GetEnterpriseAppBlockTrustOk returns a tuple with the EnterpriseAppBlockTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseAppBlockTrust

`func (o *IosGeneralDeviceConfiguration) HasEnterpriseAppBlockTrust() bool`

HasEnterpriseAppBlockTrust returns a boolean if a field has been set.

### SetEnterpriseAppBlockTrust

`func (o *IosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrust(v bool)`

SetEnterpriseAppBlockTrust gets a reference to the given bool and assigns it to the EnterpriseAppBlockTrust field.

### GetEnterpriseAppBlockTrustModification

`func (o *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustModification() bool`

GetEnterpriseAppBlockTrustModification returns the EnterpriseAppBlockTrustModification field if non-nil, zero value otherwise.

### GetEnterpriseAppBlockTrustModificationOk

`func (o *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustModificationOk() (bool, bool)`

GetEnterpriseAppBlockTrustModificationOk returns a tuple with the EnterpriseAppBlockTrustModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseAppBlockTrustModification

`func (o *IosGeneralDeviceConfiguration) HasEnterpriseAppBlockTrustModification() bool`

HasEnterpriseAppBlockTrustModification returns a boolean if a field has been set.

### SetEnterpriseAppBlockTrustModification

`func (o *IosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrustModification(v bool)`

SetEnterpriseAppBlockTrustModification gets a reference to the given bool and assigns it to the EnterpriseAppBlockTrustModification field.

### GetFaceTimeBlocked

`func (o *IosGeneralDeviceConfiguration) GetFaceTimeBlocked() bool`

GetFaceTimeBlocked returns the FaceTimeBlocked field if non-nil, zero value otherwise.

### GetFaceTimeBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetFaceTimeBlockedOk() (bool, bool)`

GetFaceTimeBlockedOk returns a tuple with the FaceTimeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaceTimeBlocked

`func (o *IosGeneralDeviceConfiguration) HasFaceTimeBlocked() bool`

HasFaceTimeBlocked returns a boolean if a field has been set.

### SetFaceTimeBlocked

`func (o *IosGeneralDeviceConfiguration) SetFaceTimeBlocked(v bool)`

SetFaceTimeBlocked gets a reference to the given bool and assigns it to the FaceTimeBlocked field.

### GetFindMyFriendsBlocked

`func (o *IosGeneralDeviceConfiguration) GetFindMyFriendsBlocked() bool`

GetFindMyFriendsBlocked returns the FindMyFriendsBlocked field if non-nil, zero value otherwise.

### GetFindMyFriendsBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetFindMyFriendsBlockedOk() (bool, bool)`

GetFindMyFriendsBlockedOk returns a tuple with the FindMyFriendsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFindMyFriendsBlocked

`func (o *IosGeneralDeviceConfiguration) HasFindMyFriendsBlocked() bool`

HasFindMyFriendsBlocked returns a boolean if a field has been set.

### SetFindMyFriendsBlocked

`func (o *IosGeneralDeviceConfiguration) SetFindMyFriendsBlocked(v bool)`

SetFindMyFriendsBlocked gets a reference to the given bool and assigns it to the FindMyFriendsBlocked field.

### GetGamingBlockGameCenterFriends

`func (o *IosGeneralDeviceConfiguration) GetGamingBlockGameCenterFriends() bool`

GetGamingBlockGameCenterFriends returns the GamingBlockGameCenterFriends field if non-nil, zero value otherwise.

### GetGamingBlockGameCenterFriendsOk

`func (o *IosGeneralDeviceConfiguration) GetGamingBlockGameCenterFriendsOk() (bool, bool)`

GetGamingBlockGameCenterFriendsOk returns a tuple with the GamingBlockGameCenterFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGamingBlockGameCenterFriends

`func (o *IosGeneralDeviceConfiguration) HasGamingBlockGameCenterFriends() bool`

HasGamingBlockGameCenterFriends returns a boolean if a field has been set.

### SetGamingBlockGameCenterFriends

`func (o *IosGeneralDeviceConfiguration) SetGamingBlockGameCenterFriends(v bool)`

SetGamingBlockGameCenterFriends gets a reference to the given bool and assigns it to the GamingBlockGameCenterFriends field.

### GetGamingBlockMultiplayer

`func (o *IosGeneralDeviceConfiguration) GetGamingBlockMultiplayer() bool`

GetGamingBlockMultiplayer returns the GamingBlockMultiplayer field if non-nil, zero value otherwise.

### GetGamingBlockMultiplayerOk

`func (o *IosGeneralDeviceConfiguration) GetGamingBlockMultiplayerOk() (bool, bool)`

GetGamingBlockMultiplayerOk returns a tuple with the GamingBlockMultiplayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGamingBlockMultiplayer

`func (o *IosGeneralDeviceConfiguration) HasGamingBlockMultiplayer() bool`

HasGamingBlockMultiplayer returns a boolean if a field has been set.

### SetGamingBlockMultiplayer

`func (o *IosGeneralDeviceConfiguration) SetGamingBlockMultiplayer(v bool)`

SetGamingBlockMultiplayer gets a reference to the given bool and assigns it to the GamingBlockMultiplayer field.

### GetGameCenterBlocked

`func (o *IosGeneralDeviceConfiguration) GetGameCenterBlocked() bool`

GetGameCenterBlocked returns the GameCenterBlocked field if non-nil, zero value otherwise.

### GetGameCenterBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetGameCenterBlockedOk() (bool, bool)`

GetGameCenterBlockedOk returns a tuple with the GameCenterBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGameCenterBlocked

`func (o *IosGeneralDeviceConfiguration) HasGameCenterBlocked() bool`

HasGameCenterBlocked returns a boolean if a field has been set.

### SetGameCenterBlocked

`func (o *IosGeneralDeviceConfiguration) SetGameCenterBlocked(v bool)`

SetGameCenterBlocked gets a reference to the given bool and assigns it to the GameCenterBlocked field.

### GetHostPairingBlocked

`func (o *IosGeneralDeviceConfiguration) GetHostPairingBlocked() bool`

GetHostPairingBlocked returns the HostPairingBlocked field if non-nil, zero value otherwise.

### GetHostPairingBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetHostPairingBlockedOk() (bool, bool)`

GetHostPairingBlockedOk returns a tuple with the HostPairingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostPairingBlocked

`func (o *IosGeneralDeviceConfiguration) HasHostPairingBlocked() bool`

HasHostPairingBlocked returns a boolean if a field has been set.

### SetHostPairingBlocked

`func (o *IosGeneralDeviceConfiguration) SetHostPairingBlocked(v bool)`

SetHostPairingBlocked gets a reference to the given bool and assigns it to the HostPairingBlocked field.

### GetIBooksStoreBlocked

`func (o *IosGeneralDeviceConfiguration) GetIBooksStoreBlocked() bool`

GetIBooksStoreBlocked returns the IBooksStoreBlocked field if non-nil, zero value otherwise.

### GetIBooksStoreBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetIBooksStoreBlockedOk() (bool, bool)`

GetIBooksStoreBlockedOk returns a tuple with the IBooksStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIBooksStoreBlocked

`func (o *IosGeneralDeviceConfiguration) HasIBooksStoreBlocked() bool`

HasIBooksStoreBlocked returns a boolean if a field has been set.

### SetIBooksStoreBlocked

`func (o *IosGeneralDeviceConfiguration) SetIBooksStoreBlocked(v bool)`

SetIBooksStoreBlocked gets a reference to the given bool and assigns it to the IBooksStoreBlocked field.

### GetIBooksStoreBlockErotica

`func (o *IosGeneralDeviceConfiguration) GetIBooksStoreBlockErotica() bool`

GetIBooksStoreBlockErotica returns the IBooksStoreBlockErotica field if non-nil, zero value otherwise.

### GetIBooksStoreBlockEroticaOk

`func (o *IosGeneralDeviceConfiguration) GetIBooksStoreBlockEroticaOk() (bool, bool)`

GetIBooksStoreBlockEroticaOk returns a tuple with the IBooksStoreBlockErotica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIBooksStoreBlockErotica

`func (o *IosGeneralDeviceConfiguration) HasIBooksStoreBlockErotica() bool`

HasIBooksStoreBlockErotica returns a boolean if a field has been set.

### SetIBooksStoreBlockErotica

`func (o *IosGeneralDeviceConfiguration) SetIBooksStoreBlockErotica(v bool)`

SetIBooksStoreBlockErotica gets a reference to the given bool and assigns it to the IBooksStoreBlockErotica field.

### GetICloudBlockActivityContinuation

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockActivityContinuation() bool`

GetICloudBlockActivityContinuation returns the ICloudBlockActivityContinuation field if non-nil, zero value otherwise.

### GetICloudBlockActivityContinuationOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockActivityContinuationOk() (bool, bool)`

GetICloudBlockActivityContinuationOk returns a tuple with the ICloudBlockActivityContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockActivityContinuation

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockActivityContinuation() bool`

HasICloudBlockActivityContinuation returns a boolean if a field has been set.

### SetICloudBlockActivityContinuation

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockActivityContinuation(v bool)`

SetICloudBlockActivityContinuation gets a reference to the given bool and assigns it to the ICloudBlockActivityContinuation field.

### GetICloudBlockBackup

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockBackup() bool`

GetICloudBlockBackup returns the ICloudBlockBackup field if non-nil, zero value otherwise.

### GetICloudBlockBackupOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockBackupOk() (bool, bool)`

GetICloudBlockBackupOk returns a tuple with the ICloudBlockBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockBackup

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockBackup() bool`

HasICloudBlockBackup returns a boolean if a field has been set.

### SetICloudBlockBackup

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockBackup(v bool)`

SetICloudBlockBackup gets a reference to the given bool and assigns it to the ICloudBlockBackup field.

### GetICloudBlockDocumentSync

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockDocumentSync() bool`

GetICloudBlockDocumentSync returns the ICloudBlockDocumentSync field if non-nil, zero value otherwise.

### GetICloudBlockDocumentSyncOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockDocumentSyncOk() (bool, bool)`

GetICloudBlockDocumentSyncOk returns a tuple with the ICloudBlockDocumentSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockDocumentSync

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockDocumentSync() bool`

HasICloudBlockDocumentSync returns a boolean if a field has been set.

### SetICloudBlockDocumentSync

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockDocumentSync(v bool)`

SetICloudBlockDocumentSync gets a reference to the given bool and assigns it to the ICloudBlockDocumentSync field.

### GetICloudBlockManagedAppsSync

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockManagedAppsSync() bool`

GetICloudBlockManagedAppsSync returns the ICloudBlockManagedAppsSync field if non-nil, zero value otherwise.

### GetICloudBlockManagedAppsSyncOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockManagedAppsSyncOk() (bool, bool)`

GetICloudBlockManagedAppsSyncOk returns a tuple with the ICloudBlockManagedAppsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockManagedAppsSync

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockManagedAppsSync() bool`

HasICloudBlockManagedAppsSync returns a boolean if a field has been set.

### SetICloudBlockManagedAppsSync

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockManagedAppsSync(v bool)`

SetICloudBlockManagedAppsSync gets a reference to the given bool and assigns it to the ICloudBlockManagedAppsSync field.

### GetICloudBlockPhotoLibrary

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockPhotoLibrary() bool`

GetICloudBlockPhotoLibrary returns the ICloudBlockPhotoLibrary field if non-nil, zero value otherwise.

### GetICloudBlockPhotoLibraryOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockPhotoLibraryOk() (bool, bool)`

GetICloudBlockPhotoLibraryOk returns a tuple with the ICloudBlockPhotoLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockPhotoLibrary

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockPhotoLibrary() bool`

HasICloudBlockPhotoLibrary returns a boolean if a field has been set.

### SetICloudBlockPhotoLibrary

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockPhotoLibrary(v bool)`

SetICloudBlockPhotoLibrary gets a reference to the given bool and assigns it to the ICloudBlockPhotoLibrary field.

### GetICloudBlockPhotoStreamSync

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockPhotoStreamSync() bool`

GetICloudBlockPhotoStreamSync returns the ICloudBlockPhotoStreamSync field if non-nil, zero value otherwise.

### GetICloudBlockPhotoStreamSyncOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockPhotoStreamSyncOk() (bool, bool)`

GetICloudBlockPhotoStreamSyncOk returns a tuple with the ICloudBlockPhotoStreamSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockPhotoStreamSync

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockPhotoStreamSync() bool`

HasICloudBlockPhotoStreamSync returns a boolean if a field has been set.

### SetICloudBlockPhotoStreamSync

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockPhotoStreamSync(v bool)`

SetICloudBlockPhotoStreamSync gets a reference to the given bool and assigns it to the ICloudBlockPhotoStreamSync field.

### GetICloudBlockSharedPhotoStream

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockSharedPhotoStream() bool`

GetICloudBlockSharedPhotoStream returns the ICloudBlockSharedPhotoStream field if non-nil, zero value otherwise.

### GetICloudBlockSharedPhotoStreamOk

`func (o *IosGeneralDeviceConfiguration) GetICloudBlockSharedPhotoStreamOk() (bool, bool)`

GetICloudBlockSharedPhotoStreamOk returns a tuple with the ICloudBlockSharedPhotoStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockSharedPhotoStream

`func (o *IosGeneralDeviceConfiguration) HasICloudBlockSharedPhotoStream() bool`

HasICloudBlockSharedPhotoStream returns a boolean if a field has been set.

### SetICloudBlockSharedPhotoStream

`func (o *IosGeneralDeviceConfiguration) SetICloudBlockSharedPhotoStream(v bool)`

SetICloudBlockSharedPhotoStream gets a reference to the given bool and assigns it to the ICloudBlockSharedPhotoStream field.

### GetICloudRequireEncryptedBackup

`func (o *IosGeneralDeviceConfiguration) GetICloudRequireEncryptedBackup() bool`

GetICloudRequireEncryptedBackup returns the ICloudRequireEncryptedBackup field if non-nil, zero value otherwise.

### GetICloudRequireEncryptedBackupOk

`func (o *IosGeneralDeviceConfiguration) GetICloudRequireEncryptedBackupOk() (bool, bool)`

GetICloudRequireEncryptedBackupOk returns a tuple with the ICloudRequireEncryptedBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudRequireEncryptedBackup

`func (o *IosGeneralDeviceConfiguration) HasICloudRequireEncryptedBackup() bool`

HasICloudRequireEncryptedBackup returns a boolean if a field has been set.

### SetICloudRequireEncryptedBackup

`func (o *IosGeneralDeviceConfiguration) SetICloudRequireEncryptedBackup(v bool)`

SetICloudRequireEncryptedBackup gets a reference to the given bool and assigns it to the ICloudRequireEncryptedBackup field.

### GetITunesBlockExplicitContent

`func (o *IosGeneralDeviceConfiguration) GetITunesBlockExplicitContent() bool`

GetITunesBlockExplicitContent returns the ITunesBlockExplicitContent field if non-nil, zero value otherwise.

### GetITunesBlockExplicitContentOk

`func (o *IosGeneralDeviceConfiguration) GetITunesBlockExplicitContentOk() (bool, bool)`

GetITunesBlockExplicitContentOk returns a tuple with the ITunesBlockExplicitContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasITunesBlockExplicitContent

`func (o *IosGeneralDeviceConfiguration) HasITunesBlockExplicitContent() bool`

HasITunesBlockExplicitContent returns a boolean if a field has been set.

### SetITunesBlockExplicitContent

`func (o *IosGeneralDeviceConfiguration) SetITunesBlockExplicitContent(v bool)`

SetITunesBlockExplicitContent gets a reference to the given bool and assigns it to the ITunesBlockExplicitContent field.

### GetITunesBlockMusicService

`func (o *IosGeneralDeviceConfiguration) GetITunesBlockMusicService() bool`

GetITunesBlockMusicService returns the ITunesBlockMusicService field if non-nil, zero value otherwise.

### GetITunesBlockMusicServiceOk

`func (o *IosGeneralDeviceConfiguration) GetITunesBlockMusicServiceOk() (bool, bool)`

GetITunesBlockMusicServiceOk returns a tuple with the ITunesBlockMusicService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasITunesBlockMusicService

`func (o *IosGeneralDeviceConfiguration) HasITunesBlockMusicService() bool`

HasITunesBlockMusicService returns a boolean if a field has been set.

### SetITunesBlockMusicService

`func (o *IosGeneralDeviceConfiguration) SetITunesBlockMusicService(v bool)`

SetITunesBlockMusicService gets a reference to the given bool and assigns it to the ITunesBlockMusicService field.

### GetITunesBlockRadio

`func (o *IosGeneralDeviceConfiguration) GetITunesBlockRadio() bool`

GetITunesBlockRadio returns the ITunesBlockRadio field if non-nil, zero value otherwise.

### GetITunesBlockRadioOk

`func (o *IosGeneralDeviceConfiguration) GetITunesBlockRadioOk() (bool, bool)`

GetITunesBlockRadioOk returns a tuple with the ITunesBlockRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasITunesBlockRadio

`func (o *IosGeneralDeviceConfiguration) HasITunesBlockRadio() bool`

HasITunesBlockRadio returns a boolean if a field has been set.

### SetITunesBlockRadio

`func (o *IosGeneralDeviceConfiguration) SetITunesBlockRadio(v bool)`

SetITunesBlockRadio gets a reference to the given bool and assigns it to the ITunesBlockRadio field.

### GetKeyboardBlockAutoCorrect

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockAutoCorrect() bool`

GetKeyboardBlockAutoCorrect returns the KeyboardBlockAutoCorrect field if non-nil, zero value otherwise.

### GetKeyboardBlockAutoCorrectOk

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockAutoCorrectOk() (bool, bool)`

GetKeyboardBlockAutoCorrectOk returns a tuple with the KeyboardBlockAutoCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockAutoCorrect

`func (o *IosGeneralDeviceConfiguration) HasKeyboardBlockAutoCorrect() bool`

HasKeyboardBlockAutoCorrect returns a boolean if a field has been set.

### SetKeyboardBlockAutoCorrect

`func (o *IosGeneralDeviceConfiguration) SetKeyboardBlockAutoCorrect(v bool)`

SetKeyboardBlockAutoCorrect gets a reference to the given bool and assigns it to the KeyboardBlockAutoCorrect field.

### GetKeyboardBlockDictation

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockDictation() bool`

GetKeyboardBlockDictation returns the KeyboardBlockDictation field if non-nil, zero value otherwise.

### GetKeyboardBlockDictationOk

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockDictationOk() (bool, bool)`

GetKeyboardBlockDictationOk returns a tuple with the KeyboardBlockDictation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockDictation

`func (o *IosGeneralDeviceConfiguration) HasKeyboardBlockDictation() bool`

HasKeyboardBlockDictation returns a boolean if a field has been set.

### SetKeyboardBlockDictation

`func (o *IosGeneralDeviceConfiguration) SetKeyboardBlockDictation(v bool)`

SetKeyboardBlockDictation gets a reference to the given bool and assigns it to the KeyboardBlockDictation field.

### GetKeyboardBlockPredictive

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockPredictive() bool`

GetKeyboardBlockPredictive returns the KeyboardBlockPredictive field if non-nil, zero value otherwise.

### GetKeyboardBlockPredictiveOk

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockPredictiveOk() (bool, bool)`

GetKeyboardBlockPredictiveOk returns a tuple with the KeyboardBlockPredictive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockPredictive

`func (o *IosGeneralDeviceConfiguration) HasKeyboardBlockPredictive() bool`

HasKeyboardBlockPredictive returns a boolean if a field has been set.

### SetKeyboardBlockPredictive

`func (o *IosGeneralDeviceConfiguration) SetKeyboardBlockPredictive(v bool)`

SetKeyboardBlockPredictive gets a reference to the given bool and assigns it to the KeyboardBlockPredictive field.

### GetKeyboardBlockShortcuts

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockShortcuts() bool`

GetKeyboardBlockShortcuts returns the KeyboardBlockShortcuts field if non-nil, zero value otherwise.

### GetKeyboardBlockShortcutsOk

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockShortcutsOk() (bool, bool)`

GetKeyboardBlockShortcutsOk returns a tuple with the KeyboardBlockShortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockShortcuts

`func (o *IosGeneralDeviceConfiguration) HasKeyboardBlockShortcuts() bool`

HasKeyboardBlockShortcuts returns a boolean if a field has been set.

### SetKeyboardBlockShortcuts

`func (o *IosGeneralDeviceConfiguration) SetKeyboardBlockShortcuts(v bool)`

SetKeyboardBlockShortcuts gets a reference to the given bool and assigns it to the KeyboardBlockShortcuts field.

### GetKeyboardBlockSpellCheck

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockSpellCheck() bool`

GetKeyboardBlockSpellCheck returns the KeyboardBlockSpellCheck field if non-nil, zero value otherwise.

### GetKeyboardBlockSpellCheckOk

`func (o *IosGeneralDeviceConfiguration) GetKeyboardBlockSpellCheckOk() (bool, bool)`

GetKeyboardBlockSpellCheckOk returns a tuple with the KeyboardBlockSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockSpellCheck

`func (o *IosGeneralDeviceConfiguration) HasKeyboardBlockSpellCheck() bool`

HasKeyboardBlockSpellCheck returns a boolean if a field has been set.

### SetKeyboardBlockSpellCheck

`func (o *IosGeneralDeviceConfiguration) SetKeyboardBlockSpellCheck(v bool)`

SetKeyboardBlockSpellCheck gets a reference to the given bool and assigns it to the KeyboardBlockSpellCheck field.

### GetKioskModeAllowAssistiveSpeak

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveSpeak() bool`

GetKioskModeAllowAssistiveSpeak returns the KioskModeAllowAssistiveSpeak field if non-nil, zero value otherwise.

### GetKioskModeAllowAssistiveSpeakOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveSpeakOk() (bool, bool)`

GetKioskModeAllowAssistiveSpeakOk returns a tuple with the KioskModeAllowAssistiveSpeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowAssistiveSpeak

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowAssistiveSpeak() bool`

HasKioskModeAllowAssistiveSpeak returns a boolean if a field has been set.

### SetKioskModeAllowAssistiveSpeak

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveSpeak(v bool)`

SetKioskModeAllowAssistiveSpeak gets a reference to the given bool and assigns it to the KioskModeAllowAssistiveSpeak field.

### GetKioskModeAllowAssistiveTouchSettings

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveTouchSettings() bool`

GetKioskModeAllowAssistiveTouchSettings returns the KioskModeAllowAssistiveTouchSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowAssistiveTouchSettingsOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveTouchSettingsOk() (bool, bool)`

GetKioskModeAllowAssistiveTouchSettingsOk returns a tuple with the KioskModeAllowAssistiveTouchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowAssistiveTouchSettings

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowAssistiveTouchSettings() bool`

HasKioskModeAllowAssistiveTouchSettings returns a boolean if a field has been set.

### SetKioskModeAllowAssistiveTouchSettings

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveTouchSettings(v bool)`

SetKioskModeAllowAssistiveTouchSettings gets a reference to the given bool and assigns it to the KioskModeAllowAssistiveTouchSettings field.

### GetKioskModeAllowAutoLock

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowAutoLock() bool`

GetKioskModeAllowAutoLock returns the KioskModeAllowAutoLock field if non-nil, zero value otherwise.

### GetKioskModeAllowAutoLockOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowAutoLockOk() (bool, bool)`

GetKioskModeAllowAutoLockOk returns a tuple with the KioskModeAllowAutoLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowAutoLock

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowAutoLock() bool`

HasKioskModeAllowAutoLock returns a boolean if a field has been set.

### SetKioskModeAllowAutoLock

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowAutoLock(v bool)`

SetKioskModeAllowAutoLock gets a reference to the given bool and assigns it to the KioskModeAllowAutoLock field.

### GetKioskModeAllowColorInversionSettings

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowColorInversionSettings() bool`

GetKioskModeAllowColorInversionSettings returns the KioskModeAllowColorInversionSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowColorInversionSettingsOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowColorInversionSettingsOk() (bool, bool)`

GetKioskModeAllowColorInversionSettingsOk returns a tuple with the KioskModeAllowColorInversionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowColorInversionSettings

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowColorInversionSettings() bool`

HasKioskModeAllowColorInversionSettings returns a boolean if a field has been set.

### SetKioskModeAllowColorInversionSettings

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowColorInversionSettings(v bool)`

SetKioskModeAllowColorInversionSettings gets a reference to the given bool and assigns it to the KioskModeAllowColorInversionSettings field.

### GetKioskModeAllowRingerSwitch

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowRingerSwitch() bool`

GetKioskModeAllowRingerSwitch returns the KioskModeAllowRingerSwitch field if non-nil, zero value otherwise.

### GetKioskModeAllowRingerSwitchOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowRingerSwitchOk() (bool, bool)`

GetKioskModeAllowRingerSwitchOk returns a tuple with the KioskModeAllowRingerSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowRingerSwitch

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowRingerSwitch() bool`

HasKioskModeAllowRingerSwitch returns a boolean if a field has been set.

### SetKioskModeAllowRingerSwitch

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowRingerSwitch(v bool)`

SetKioskModeAllowRingerSwitch gets a reference to the given bool and assigns it to the KioskModeAllowRingerSwitch field.

### GetKioskModeAllowScreenRotation

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowScreenRotation() bool`

GetKioskModeAllowScreenRotation returns the KioskModeAllowScreenRotation field if non-nil, zero value otherwise.

### GetKioskModeAllowScreenRotationOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowScreenRotationOk() (bool, bool)`

GetKioskModeAllowScreenRotationOk returns a tuple with the KioskModeAllowScreenRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowScreenRotation

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowScreenRotation() bool`

HasKioskModeAllowScreenRotation returns a boolean if a field has been set.

### SetKioskModeAllowScreenRotation

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowScreenRotation(v bool)`

SetKioskModeAllowScreenRotation gets a reference to the given bool and assigns it to the KioskModeAllowScreenRotation field.

### GetKioskModeAllowSleepButton

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowSleepButton() bool`

GetKioskModeAllowSleepButton returns the KioskModeAllowSleepButton field if non-nil, zero value otherwise.

### GetKioskModeAllowSleepButtonOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowSleepButtonOk() (bool, bool)`

GetKioskModeAllowSleepButtonOk returns a tuple with the KioskModeAllowSleepButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowSleepButton

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowSleepButton() bool`

HasKioskModeAllowSleepButton returns a boolean if a field has been set.

### SetKioskModeAllowSleepButton

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowSleepButton(v bool)`

SetKioskModeAllowSleepButton gets a reference to the given bool and assigns it to the KioskModeAllowSleepButton field.

### GetKioskModeAllowTouchscreen

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowTouchscreen() bool`

GetKioskModeAllowTouchscreen returns the KioskModeAllowTouchscreen field if non-nil, zero value otherwise.

### GetKioskModeAllowTouchscreenOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowTouchscreenOk() (bool, bool)`

GetKioskModeAllowTouchscreenOk returns a tuple with the KioskModeAllowTouchscreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowTouchscreen

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowTouchscreen() bool`

HasKioskModeAllowTouchscreen returns a boolean if a field has been set.

### SetKioskModeAllowTouchscreen

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowTouchscreen(v bool)`

SetKioskModeAllowTouchscreen gets a reference to the given bool and assigns it to the KioskModeAllowTouchscreen field.

### GetKioskModeAllowVoiceOverSettings

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowVoiceOverSettings() bool`

GetKioskModeAllowVoiceOverSettings returns the KioskModeAllowVoiceOverSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowVoiceOverSettingsOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowVoiceOverSettingsOk() (bool, bool)`

GetKioskModeAllowVoiceOverSettingsOk returns a tuple with the KioskModeAllowVoiceOverSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowVoiceOverSettings

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowVoiceOverSettings() bool`

HasKioskModeAllowVoiceOverSettings returns a boolean if a field has been set.

### SetKioskModeAllowVoiceOverSettings

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowVoiceOverSettings(v bool)`

SetKioskModeAllowVoiceOverSettings gets a reference to the given bool and assigns it to the KioskModeAllowVoiceOverSettings field.

### GetKioskModeAllowVolumeButtons

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowVolumeButtons() bool`

GetKioskModeAllowVolumeButtons returns the KioskModeAllowVolumeButtons field if non-nil, zero value otherwise.

### GetKioskModeAllowVolumeButtonsOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowVolumeButtonsOk() (bool, bool)`

GetKioskModeAllowVolumeButtonsOk returns a tuple with the KioskModeAllowVolumeButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowVolumeButtons

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowVolumeButtons() bool`

HasKioskModeAllowVolumeButtons returns a boolean if a field has been set.

### SetKioskModeAllowVolumeButtons

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowVolumeButtons(v bool)`

SetKioskModeAllowVolumeButtons gets a reference to the given bool and assigns it to the KioskModeAllowVolumeButtons field.

### GetKioskModeAllowZoomSettings

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowZoomSettings() bool`

GetKioskModeAllowZoomSettings returns the KioskModeAllowZoomSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowZoomSettingsOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAllowZoomSettingsOk() (bool, bool)`

GetKioskModeAllowZoomSettingsOk returns a tuple with the KioskModeAllowZoomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowZoomSettings

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAllowZoomSettings() bool`

HasKioskModeAllowZoomSettings returns a boolean if a field has been set.

### SetKioskModeAllowZoomSettings

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAllowZoomSettings(v bool)`

SetKioskModeAllowZoomSettings gets a reference to the given bool and assigns it to the KioskModeAllowZoomSettings field.

### GetKioskModeAppStoreUrl

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAppStoreUrl() string`

GetKioskModeAppStoreUrl returns the KioskModeAppStoreUrl field if non-nil, zero value otherwise.

### GetKioskModeAppStoreUrlOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeAppStoreUrlOk() (string, bool)`

GetKioskModeAppStoreUrlOk returns a tuple with the KioskModeAppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAppStoreUrl

`func (o *IosGeneralDeviceConfiguration) HasKioskModeAppStoreUrl() bool`

HasKioskModeAppStoreUrl returns a boolean if a field has been set.

### SetKioskModeAppStoreUrl

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAppStoreUrl(v string)`

SetKioskModeAppStoreUrl gets a reference to the given string and assigns it to the KioskModeAppStoreUrl field.

### SetKioskModeAppStoreUrlExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetKioskModeAppStoreUrlExplicitNull(b bool)`

SetKioskModeAppStoreUrlExplicitNull (un)sets KioskModeAppStoreUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskModeAppStoreUrl value is set to nil even if false is passed
### GetKioskModeBuiltInAppId

`func (o *IosGeneralDeviceConfiguration) GetKioskModeBuiltInAppId() string`

GetKioskModeBuiltInAppId returns the KioskModeBuiltInAppId field if non-nil, zero value otherwise.

### GetKioskModeBuiltInAppIdOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeBuiltInAppIdOk() (string, bool)`

GetKioskModeBuiltInAppIdOk returns a tuple with the KioskModeBuiltInAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeBuiltInAppId

`func (o *IosGeneralDeviceConfiguration) HasKioskModeBuiltInAppId() bool`

HasKioskModeBuiltInAppId returns a boolean if a field has been set.

### SetKioskModeBuiltInAppId

`func (o *IosGeneralDeviceConfiguration) SetKioskModeBuiltInAppId(v string)`

SetKioskModeBuiltInAppId gets a reference to the given string and assigns it to the KioskModeBuiltInAppId field.

### SetKioskModeBuiltInAppIdExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetKioskModeBuiltInAppIdExplicitNull(b bool)`

SetKioskModeBuiltInAppIdExplicitNull (un)sets KioskModeBuiltInAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskModeBuiltInAppId value is set to nil even if false is passed
### GetKioskModeRequireAssistiveTouch

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireAssistiveTouch() bool`

GetKioskModeRequireAssistiveTouch returns the KioskModeRequireAssistiveTouch field if non-nil, zero value otherwise.

### GetKioskModeRequireAssistiveTouchOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireAssistiveTouchOk() (bool, bool)`

GetKioskModeRequireAssistiveTouchOk returns a tuple with the KioskModeRequireAssistiveTouch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireAssistiveTouch

`func (o *IosGeneralDeviceConfiguration) HasKioskModeRequireAssistiveTouch() bool`

HasKioskModeRequireAssistiveTouch returns a boolean if a field has been set.

### SetKioskModeRequireAssistiveTouch

`func (o *IosGeneralDeviceConfiguration) SetKioskModeRequireAssistiveTouch(v bool)`

SetKioskModeRequireAssistiveTouch gets a reference to the given bool and assigns it to the KioskModeRequireAssistiveTouch field.

### GetKioskModeRequireColorInversion

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireColorInversion() bool`

GetKioskModeRequireColorInversion returns the KioskModeRequireColorInversion field if non-nil, zero value otherwise.

### GetKioskModeRequireColorInversionOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireColorInversionOk() (bool, bool)`

GetKioskModeRequireColorInversionOk returns a tuple with the KioskModeRequireColorInversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireColorInversion

`func (o *IosGeneralDeviceConfiguration) HasKioskModeRequireColorInversion() bool`

HasKioskModeRequireColorInversion returns a boolean if a field has been set.

### SetKioskModeRequireColorInversion

`func (o *IosGeneralDeviceConfiguration) SetKioskModeRequireColorInversion(v bool)`

SetKioskModeRequireColorInversion gets a reference to the given bool and assigns it to the KioskModeRequireColorInversion field.

### GetKioskModeRequireMonoAudio

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireMonoAudio() bool`

GetKioskModeRequireMonoAudio returns the KioskModeRequireMonoAudio field if non-nil, zero value otherwise.

### GetKioskModeRequireMonoAudioOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireMonoAudioOk() (bool, bool)`

GetKioskModeRequireMonoAudioOk returns a tuple with the KioskModeRequireMonoAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireMonoAudio

`func (o *IosGeneralDeviceConfiguration) HasKioskModeRequireMonoAudio() bool`

HasKioskModeRequireMonoAudio returns a boolean if a field has been set.

### SetKioskModeRequireMonoAudio

`func (o *IosGeneralDeviceConfiguration) SetKioskModeRequireMonoAudio(v bool)`

SetKioskModeRequireMonoAudio gets a reference to the given bool and assigns it to the KioskModeRequireMonoAudio field.

### GetKioskModeRequireVoiceOver

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireVoiceOver() bool`

GetKioskModeRequireVoiceOver returns the KioskModeRequireVoiceOver field if non-nil, zero value otherwise.

### GetKioskModeRequireVoiceOverOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireVoiceOverOk() (bool, bool)`

GetKioskModeRequireVoiceOverOk returns a tuple with the KioskModeRequireVoiceOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireVoiceOver

`func (o *IosGeneralDeviceConfiguration) HasKioskModeRequireVoiceOver() bool`

HasKioskModeRequireVoiceOver returns a boolean if a field has been set.

### SetKioskModeRequireVoiceOver

`func (o *IosGeneralDeviceConfiguration) SetKioskModeRequireVoiceOver(v bool)`

SetKioskModeRequireVoiceOver gets a reference to the given bool and assigns it to the KioskModeRequireVoiceOver field.

### GetKioskModeRequireZoom

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireZoom() bool`

GetKioskModeRequireZoom returns the KioskModeRequireZoom field if non-nil, zero value otherwise.

### GetKioskModeRequireZoomOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeRequireZoomOk() (bool, bool)`

GetKioskModeRequireZoomOk returns a tuple with the KioskModeRequireZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireZoom

`func (o *IosGeneralDeviceConfiguration) HasKioskModeRequireZoom() bool`

HasKioskModeRequireZoom returns a boolean if a field has been set.

### SetKioskModeRequireZoom

`func (o *IosGeneralDeviceConfiguration) SetKioskModeRequireZoom(v bool)`

SetKioskModeRequireZoom gets a reference to the given bool and assigns it to the KioskModeRequireZoom field.

### GetKioskModeManagedAppId

`func (o *IosGeneralDeviceConfiguration) GetKioskModeManagedAppId() string`

GetKioskModeManagedAppId returns the KioskModeManagedAppId field if non-nil, zero value otherwise.

### GetKioskModeManagedAppIdOk

`func (o *IosGeneralDeviceConfiguration) GetKioskModeManagedAppIdOk() (string, bool)`

GetKioskModeManagedAppIdOk returns a tuple with the KioskModeManagedAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeManagedAppId

`func (o *IosGeneralDeviceConfiguration) HasKioskModeManagedAppId() bool`

HasKioskModeManagedAppId returns a boolean if a field has been set.

### SetKioskModeManagedAppId

`func (o *IosGeneralDeviceConfiguration) SetKioskModeManagedAppId(v string)`

SetKioskModeManagedAppId gets a reference to the given string and assigns it to the KioskModeManagedAppId field.

### SetKioskModeManagedAppIdExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetKioskModeManagedAppIdExplicitNull(b bool)`

SetKioskModeManagedAppIdExplicitNull (un)sets KioskModeManagedAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskModeManagedAppId value is set to nil even if false is passed
### GetLockScreenBlockControlCenter

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockControlCenter() bool`

GetLockScreenBlockControlCenter returns the LockScreenBlockControlCenter field if non-nil, zero value otherwise.

### GetLockScreenBlockControlCenterOk

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockControlCenterOk() (bool, bool)`

GetLockScreenBlockControlCenterOk returns a tuple with the LockScreenBlockControlCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockControlCenter

`func (o *IosGeneralDeviceConfiguration) HasLockScreenBlockControlCenter() bool`

HasLockScreenBlockControlCenter returns a boolean if a field has been set.

### SetLockScreenBlockControlCenter

`func (o *IosGeneralDeviceConfiguration) SetLockScreenBlockControlCenter(v bool)`

SetLockScreenBlockControlCenter gets a reference to the given bool and assigns it to the LockScreenBlockControlCenter field.

### GetLockScreenBlockNotificationView

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockNotificationView() bool`

GetLockScreenBlockNotificationView returns the LockScreenBlockNotificationView field if non-nil, zero value otherwise.

### GetLockScreenBlockNotificationViewOk

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockNotificationViewOk() (bool, bool)`

GetLockScreenBlockNotificationViewOk returns a tuple with the LockScreenBlockNotificationView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockNotificationView

`func (o *IosGeneralDeviceConfiguration) HasLockScreenBlockNotificationView() bool`

HasLockScreenBlockNotificationView returns a boolean if a field has been set.

### SetLockScreenBlockNotificationView

`func (o *IosGeneralDeviceConfiguration) SetLockScreenBlockNotificationView(v bool)`

SetLockScreenBlockNotificationView gets a reference to the given bool and assigns it to the LockScreenBlockNotificationView field.

### GetLockScreenBlockPassbook

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockPassbook() bool`

GetLockScreenBlockPassbook returns the LockScreenBlockPassbook field if non-nil, zero value otherwise.

### GetLockScreenBlockPassbookOk

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockPassbookOk() (bool, bool)`

GetLockScreenBlockPassbookOk returns a tuple with the LockScreenBlockPassbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockPassbook

`func (o *IosGeneralDeviceConfiguration) HasLockScreenBlockPassbook() bool`

HasLockScreenBlockPassbook returns a boolean if a field has been set.

### SetLockScreenBlockPassbook

`func (o *IosGeneralDeviceConfiguration) SetLockScreenBlockPassbook(v bool)`

SetLockScreenBlockPassbook gets a reference to the given bool and assigns it to the LockScreenBlockPassbook field.

### GetLockScreenBlockTodayView

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockTodayView() bool`

GetLockScreenBlockTodayView returns the LockScreenBlockTodayView field if non-nil, zero value otherwise.

### GetLockScreenBlockTodayViewOk

`func (o *IosGeneralDeviceConfiguration) GetLockScreenBlockTodayViewOk() (bool, bool)`

GetLockScreenBlockTodayViewOk returns a tuple with the LockScreenBlockTodayView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockTodayView

`func (o *IosGeneralDeviceConfiguration) HasLockScreenBlockTodayView() bool`

HasLockScreenBlockTodayView returns a boolean if a field has been set.

### SetLockScreenBlockTodayView

`func (o *IosGeneralDeviceConfiguration) SetLockScreenBlockTodayView(v bool)`

SetLockScreenBlockTodayView gets a reference to the given bool and assigns it to the LockScreenBlockTodayView field.

### GetMediaContentRatingAustralia

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingAustralia() AnyOfmicrosoftGraphMediaContentRatingAustralia`

GetMediaContentRatingAustralia returns the MediaContentRatingAustralia field if non-nil, zero value otherwise.

### GetMediaContentRatingAustraliaOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingAustraliaOk() (AnyOfmicrosoftGraphMediaContentRatingAustralia, bool)`

GetMediaContentRatingAustraliaOk returns a tuple with the MediaContentRatingAustralia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingAustralia

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingAustralia() bool`

HasMediaContentRatingAustralia returns a boolean if a field has been set.

### SetMediaContentRatingAustralia

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingAustralia(v AnyOfmicrosoftGraphMediaContentRatingAustralia)`

SetMediaContentRatingAustralia gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingAustralia and assigns it to the MediaContentRatingAustralia field.

### SetMediaContentRatingAustraliaExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingAustraliaExplicitNull(b bool)`

SetMediaContentRatingAustraliaExplicitNull (un)sets MediaContentRatingAustralia to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingAustralia value is set to nil even if false is passed
### GetMediaContentRatingCanada

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingCanada() AnyOfmicrosoftGraphMediaContentRatingCanada`

GetMediaContentRatingCanada returns the MediaContentRatingCanada field if non-nil, zero value otherwise.

### GetMediaContentRatingCanadaOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingCanadaOk() (AnyOfmicrosoftGraphMediaContentRatingCanada, bool)`

GetMediaContentRatingCanadaOk returns a tuple with the MediaContentRatingCanada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingCanada

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingCanada() bool`

HasMediaContentRatingCanada returns a boolean if a field has been set.

### SetMediaContentRatingCanada

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingCanada(v AnyOfmicrosoftGraphMediaContentRatingCanada)`

SetMediaContentRatingCanada gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingCanada and assigns it to the MediaContentRatingCanada field.

### SetMediaContentRatingCanadaExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingCanadaExplicitNull(b bool)`

SetMediaContentRatingCanadaExplicitNull (un)sets MediaContentRatingCanada to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingCanada value is set to nil even if false is passed
### GetMediaContentRatingFrance

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingFrance() AnyOfmicrosoftGraphMediaContentRatingFrance`

GetMediaContentRatingFrance returns the MediaContentRatingFrance field if non-nil, zero value otherwise.

### GetMediaContentRatingFranceOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingFranceOk() (AnyOfmicrosoftGraphMediaContentRatingFrance, bool)`

GetMediaContentRatingFranceOk returns a tuple with the MediaContentRatingFrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingFrance

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingFrance() bool`

HasMediaContentRatingFrance returns a boolean if a field has been set.

### SetMediaContentRatingFrance

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingFrance(v AnyOfmicrosoftGraphMediaContentRatingFrance)`

SetMediaContentRatingFrance gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingFrance and assigns it to the MediaContentRatingFrance field.

### SetMediaContentRatingFranceExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingFranceExplicitNull(b bool)`

SetMediaContentRatingFranceExplicitNull (un)sets MediaContentRatingFrance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingFrance value is set to nil even if false is passed
### GetMediaContentRatingGermany

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingGermany() AnyOfmicrosoftGraphMediaContentRatingGermany`

GetMediaContentRatingGermany returns the MediaContentRatingGermany field if non-nil, zero value otherwise.

### GetMediaContentRatingGermanyOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingGermanyOk() (AnyOfmicrosoftGraphMediaContentRatingGermany, bool)`

GetMediaContentRatingGermanyOk returns a tuple with the MediaContentRatingGermany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingGermany

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingGermany() bool`

HasMediaContentRatingGermany returns a boolean if a field has been set.

### SetMediaContentRatingGermany

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingGermany(v AnyOfmicrosoftGraphMediaContentRatingGermany)`

SetMediaContentRatingGermany gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingGermany and assigns it to the MediaContentRatingGermany field.

### SetMediaContentRatingGermanyExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingGermanyExplicitNull(b bool)`

SetMediaContentRatingGermanyExplicitNull (un)sets MediaContentRatingGermany to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingGermany value is set to nil even if false is passed
### GetMediaContentRatingIreland

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingIreland() AnyOfmicrosoftGraphMediaContentRatingIreland`

GetMediaContentRatingIreland returns the MediaContentRatingIreland field if non-nil, zero value otherwise.

### GetMediaContentRatingIrelandOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingIrelandOk() (AnyOfmicrosoftGraphMediaContentRatingIreland, bool)`

GetMediaContentRatingIrelandOk returns a tuple with the MediaContentRatingIreland field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingIreland

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingIreland() bool`

HasMediaContentRatingIreland returns a boolean if a field has been set.

### SetMediaContentRatingIreland

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingIreland(v AnyOfmicrosoftGraphMediaContentRatingIreland)`

SetMediaContentRatingIreland gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingIreland and assigns it to the MediaContentRatingIreland field.

### SetMediaContentRatingIrelandExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingIrelandExplicitNull(b bool)`

SetMediaContentRatingIrelandExplicitNull (un)sets MediaContentRatingIreland to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingIreland value is set to nil even if false is passed
### GetMediaContentRatingJapan

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingJapan() AnyOfmicrosoftGraphMediaContentRatingJapan`

GetMediaContentRatingJapan returns the MediaContentRatingJapan field if non-nil, zero value otherwise.

### GetMediaContentRatingJapanOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingJapanOk() (AnyOfmicrosoftGraphMediaContentRatingJapan, bool)`

GetMediaContentRatingJapanOk returns a tuple with the MediaContentRatingJapan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingJapan

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingJapan() bool`

HasMediaContentRatingJapan returns a boolean if a field has been set.

### SetMediaContentRatingJapan

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingJapan(v AnyOfmicrosoftGraphMediaContentRatingJapan)`

SetMediaContentRatingJapan gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingJapan and assigns it to the MediaContentRatingJapan field.

### SetMediaContentRatingJapanExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingJapanExplicitNull(b bool)`

SetMediaContentRatingJapanExplicitNull (un)sets MediaContentRatingJapan to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingJapan value is set to nil even if false is passed
### GetMediaContentRatingNewZealand

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingNewZealand() AnyOfmicrosoftGraphMediaContentRatingNewZealand`

GetMediaContentRatingNewZealand returns the MediaContentRatingNewZealand field if non-nil, zero value otherwise.

### GetMediaContentRatingNewZealandOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingNewZealandOk() (AnyOfmicrosoftGraphMediaContentRatingNewZealand, bool)`

GetMediaContentRatingNewZealandOk returns a tuple with the MediaContentRatingNewZealand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingNewZealand

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingNewZealand() bool`

HasMediaContentRatingNewZealand returns a boolean if a field has been set.

### SetMediaContentRatingNewZealand

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingNewZealand(v AnyOfmicrosoftGraphMediaContentRatingNewZealand)`

SetMediaContentRatingNewZealand gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingNewZealand and assigns it to the MediaContentRatingNewZealand field.

### SetMediaContentRatingNewZealandExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingNewZealandExplicitNull(b bool)`

SetMediaContentRatingNewZealandExplicitNull (un)sets MediaContentRatingNewZealand to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingNewZealand value is set to nil even if false is passed
### GetMediaContentRatingUnitedKingdom

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedKingdom() AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom`

GetMediaContentRatingUnitedKingdom returns the MediaContentRatingUnitedKingdom field if non-nil, zero value otherwise.

### GetMediaContentRatingUnitedKingdomOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedKingdomOk() (AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom, bool)`

GetMediaContentRatingUnitedKingdomOk returns a tuple with the MediaContentRatingUnitedKingdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingUnitedKingdom

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingUnitedKingdom() bool`

HasMediaContentRatingUnitedKingdom returns a boolean if a field has been set.

### SetMediaContentRatingUnitedKingdom

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedKingdom(v AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom)`

SetMediaContentRatingUnitedKingdom gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom and assigns it to the MediaContentRatingUnitedKingdom field.

### SetMediaContentRatingUnitedKingdomExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedKingdomExplicitNull(b bool)`

SetMediaContentRatingUnitedKingdomExplicitNull (un)sets MediaContentRatingUnitedKingdom to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingUnitedKingdom value is set to nil even if false is passed
### GetMediaContentRatingUnitedStates

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedStates() AnyOfmicrosoftGraphMediaContentRatingUnitedStates`

GetMediaContentRatingUnitedStates returns the MediaContentRatingUnitedStates field if non-nil, zero value otherwise.

### GetMediaContentRatingUnitedStatesOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedStatesOk() (AnyOfmicrosoftGraphMediaContentRatingUnitedStates, bool)`

GetMediaContentRatingUnitedStatesOk returns a tuple with the MediaContentRatingUnitedStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingUnitedStates

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingUnitedStates() bool`

HasMediaContentRatingUnitedStates returns a boolean if a field has been set.

### SetMediaContentRatingUnitedStates

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedStates(v AnyOfmicrosoftGraphMediaContentRatingUnitedStates)`

SetMediaContentRatingUnitedStates gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingUnitedStates and assigns it to the MediaContentRatingUnitedStates field.

### SetMediaContentRatingUnitedStatesExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedStatesExplicitNull(b bool)`

SetMediaContentRatingUnitedStatesExplicitNull (un)sets MediaContentRatingUnitedStates to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingUnitedStates value is set to nil even if false is passed
### GetNetworkUsageRules

`func (o *IosGeneralDeviceConfiguration) GetNetworkUsageRules() []AnyOfmicrosoftGraphIosNetworkUsageRule`

GetNetworkUsageRules returns the NetworkUsageRules field if non-nil, zero value otherwise.

### GetNetworkUsageRulesOk

`func (o *IosGeneralDeviceConfiguration) GetNetworkUsageRulesOk() ([]AnyOfmicrosoftGraphIosNetworkUsageRule, bool)`

GetNetworkUsageRulesOk returns a tuple with the NetworkUsageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkUsageRules

`func (o *IosGeneralDeviceConfiguration) HasNetworkUsageRules() bool`

HasNetworkUsageRules returns a boolean if a field has been set.

### SetNetworkUsageRules

`func (o *IosGeneralDeviceConfiguration) SetNetworkUsageRules(v []AnyOfmicrosoftGraphIosNetworkUsageRule)`

SetNetworkUsageRules gets a reference to the given []AnyOfmicrosoftGraphIosNetworkUsageRule and assigns it to the NetworkUsageRules field.

### GetMediaContentRatingApps

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingApps() AnyOfmicrosoftGraphRatingAppsType`

GetMediaContentRatingApps returns the MediaContentRatingApps field if non-nil, zero value otherwise.

### GetMediaContentRatingAppsOk

`func (o *IosGeneralDeviceConfiguration) GetMediaContentRatingAppsOk() (AnyOfmicrosoftGraphRatingAppsType, bool)`

GetMediaContentRatingAppsOk returns a tuple with the MediaContentRatingApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingApps

`func (o *IosGeneralDeviceConfiguration) HasMediaContentRatingApps() bool`

HasMediaContentRatingApps returns a boolean if a field has been set.

### SetMediaContentRatingApps

`func (o *IosGeneralDeviceConfiguration) SetMediaContentRatingApps(v AnyOfmicrosoftGraphRatingAppsType)`

SetMediaContentRatingApps gets a reference to the given AnyOfmicrosoftGraphRatingAppsType and assigns it to the MediaContentRatingApps field.

### GetMessagesBlocked

`func (o *IosGeneralDeviceConfiguration) GetMessagesBlocked() bool`

GetMessagesBlocked returns the MessagesBlocked field if non-nil, zero value otherwise.

### GetMessagesBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetMessagesBlockedOk() (bool, bool)`

GetMessagesBlockedOk returns a tuple with the MessagesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessagesBlocked

`func (o *IosGeneralDeviceConfiguration) HasMessagesBlocked() bool`

HasMessagesBlocked returns a boolean if a field has been set.

### SetMessagesBlocked

`func (o *IosGeneralDeviceConfiguration) SetMessagesBlocked(v bool)`

SetMessagesBlocked gets a reference to the given bool and assigns it to the MessagesBlocked field.

### GetNotificationsBlockSettingsModification

`func (o *IosGeneralDeviceConfiguration) GetNotificationsBlockSettingsModification() bool`

GetNotificationsBlockSettingsModification returns the NotificationsBlockSettingsModification field if non-nil, zero value otherwise.

### GetNotificationsBlockSettingsModificationOk

`func (o *IosGeneralDeviceConfiguration) GetNotificationsBlockSettingsModificationOk() (bool, bool)`

GetNotificationsBlockSettingsModificationOk returns a tuple with the NotificationsBlockSettingsModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationsBlockSettingsModification

`func (o *IosGeneralDeviceConfiguration) HasNotificationsBlockSettingsModification() bool`

HasNotificationsBlockSettingsModification returns a boolean if a field has been set.

### SetNotificationsBlockSettingsModification

`func (o *IosGeneralDeviceConfiguration) SetNotificationsBlockSettingsModification(v bool)`

SetNotificationsBlockSettingsModification gets a reference to the given bool and assigns it to the NotificationsBlockSettingsModification field.

### GetPasscodeBlockFingerprintUnlock

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintUnlock() bool`

GetPasscodeBlockFingerprintUnlock returns the PasscodeBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetPasscodeBlockFingerprintUnlockOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintUnlockOk() (bool, bool)`

GetPasscodeBlockFingerprintUnlockOk returns a tuple with the PasscodeBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockFingerprintUnlock

`func (o *IosGeneralDeviceConfiguration) HasPasscodeBlockFingerprintUnlock() bool`

HasPasscodeBlockFingerprintUnlock returns a boolean if a field has been set.

### SetPasscodeBlockFingerprintUnlock

`func (o *IosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintUnlock(v bool)`

SetPasscodeBlockFingerprintUnlock gets a reference to the given bool and assigns it to the PasscodeBlockFingerprintUnlock field.

### GetPasscodeBlockFingerprintModification

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintModification() bool`

GetPasscodeBlockFingerprintModification returns the PasscodeBlockFingerprintModification field if non-nil, zero value otherwise.

### GetPasscodeBlockFingerprintModificationOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintModificationOk() (bool, bool)`

GetPasscodeBlockFingerprintModificationOk returns a tuple with the PasscodeBlockFingerprintModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockFingerprintModification

`func (o *IosGeneralDeviceConfiguration) HasPasscodeBlockFingerprintModification() bool`

HasPasscodeBlockFingerprintModification returns a boolean if a field has been set.

### SetPasscodeBlockFingerprintModification

`func (o *IosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintModification(v bool)`

SetPasscodeBlockFingerprintModification gets a reference to the given bool and assigns it to the PasscodeBlockFingerprintModification field.

### GetPasscodeBlockModification

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockModification() bool`

GetPasscodeBlockModification returns the PasscodeBlockModification field if non-nil, zero value otherwise.

### GetPasscodeBlockModificationOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockModificationOk() (bool, bool)`

GetPasscodeBlockModificationOk returns a tuple with the PasscodeBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockModification

`func (o *IosGeneralDeviceConfiguration) HasPasscodeBlockModification() bool`

HasPasscodeBlockModification returns a boolean if a field has been set.

### SetPasscodeBlockModification

`func (o *IosGeneralDeviceConfiguration) SetPasscodeBlockModification(v bool)`

SetPasscodeBlockModification gets a reference to the given bool and assigns it to the PasscodeBlockModification field.

### GetPasscodeBlockSimple

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockSimple() bool`

GetPasscodeBlockSimple returns the PasscodeBlockSimple field if non-nil, zero value otherwise.

### GetPasscodeBlockSimpleOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeBlockSimpleOk() (bool, bool)`

GetPasscodeBlockSimpleOk returns a tuple with the PasscodeBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockSimple

`func (o *IosGeneralDeviceConfiguration) HasPasscodeBlockSimple() bool`

HasPasscodeBlockSimple returns a boolean if a field has been set.

### SetPasscodeBlockSimple

`func (o *IosGeneralDeviceConfiguration) SetPasscodeBlockSimple(v bool)`

SetPasscodeBlockSimple gets a reference to the given bool and assigns it to the PasscodeBlockSimple field.

### GetPasscodeExpirationDays

`func (o *IosGeneralDeviceConfiguration) GetPasscodeExpirationDays() int32`

GetPasscodeExpirationDays returns the PasscodeExpirationDays field if non-nil, zero value otherwise.

### GetPasscodeExpirationDaysOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeExpirationDaysOk() (int32, bool)`

GetPasscodeExpirationDaysOk returns a tuple with the PasscodeExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeExpirationDays

`func (o *IosGeneralDeviceConfiguration) HasPasscodeExpirationDays() bool`

HasPasscodeExpirationDays returns a boolean if a field has been set.

### SetPasscodeExpirationDays

`func (o *IosGeneralDeviceConfiguration) SetPasscodeExpirationDays(v int32)`

SetPasscodeExpirationDays gets a reference to the given int32 and assigns it to the PasscodeExpirationDays field.

### SetPasscodeExpirationDaysExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodeExpirationDaysExplicitNull(b bool)`

SetPasscodeExpirationDaysExplicitNull (un)sets PasscodeExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeExpirationDays value is set to nil even if false is passed
### GetPasscodeMinimumLength

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinimumLength() int32`

GetPasscodeMinimumLength returns the PasscodeMinimumLength field if non-nil, zero value otherwise.

### GetPasscodeMinimumLengthOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinimumLengthOk() (int32, bool)`

GetPasscodeMinimumLengthOk returns a tuple with the PasscodeMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumLength

`func (o *IosGeneralDeviceConfiguration) HasPasscodeMinimumLength() bool`

HasPasscodeMinimumLength returns a boolean if a field has been set.

### SetPasscodeMinimumLength

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinimumLength(v int32)`

SetPasscodeMinimumLength gets a reference to the given int32 and assigns it to the PasscodeMinimumLength field.

### SetPasscodeMinimumLengthExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinimumLengthExplicitNull(b bool)`

SetPasscodeMinimumLengthExplicitNull (un)sets PasscodeMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumLength value is set to nil even if false is passed
### GetPasscodeMinutesOfInactivityBeforeLock

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeLock() int32`

GetPasscodeMinutesOfInactivityBeforeLock returns the PasscodeMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasscodeMinutesOfInactivityBeforeLockOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasscodeMinutesOfInactivityBeforeLockOk returns a tuple with the PasscodeMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinutesOfInactivityBeforeLock

`func (o *IosGeneralDeviceConfiguration) HasPasscodeMinutesOfInactivityBeforeLock() bool`

HasPasscodeMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasscodeMinutesOfInactivityBeforeLock

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeLock(v int32)`

SetPasscodeMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeLock field.

### SetPasscodeMinutesOfInactivityBeforeLockExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasscodeMinutesOfInactivityBeforeLockExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasscodeMinutesOfInactivityBeforeScreenTimeout

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasscodeMinutesOfInactivityBeforeScreenTimeout returns the PasscodeMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasscodeMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasscodeMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasscodeMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinutesOfInactivityBeforeScreenTimeout

`func (o *IosGeneralDeviceConfiguration) HasPasscodeMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasscodeMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasscodeMinutesOfInactivityBeforeScreenTimeout

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasscodeMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeScreenTimeout field.

### SetPasscodeMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasscodeMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasscodeMinimumCharacterSetCount

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinimumCharacterSetCount() int32`

GetPasscodeMinimumCharacterSetCount returns the PasscodeMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasscodeMinimumCharacterSetCountOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeMinimumCharacterSetCountOk() (int32, bool)`

GetPasscodeMinimumCharacterSetCountOk returns a tuple with the PasscodeMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumCharacterSetCount

`func (o *IosGeneralDeviceConfiguration) HasPasscodeMinimumCharacterSetCount() bool`

HasPasscodeMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasscodeMinimumCharacterSetCount

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinimumCharacterSetCount(v int32)`

SetPasscodeMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasscodeMinimumCharacterSetCount field.

### SetPasscodeMinimumCharacterSetCountExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodeMinimumCharacterSetCountExplicitNull(b bool)`

SetPasscodeMinimumCharacterSetCountExplicitNull (un)sets PasscodeMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasscodePreviousPasscodeBlockCount

`func (o *IosGeneralDeviceConfiguration) GetPasscodePreviousPasscodeBlockCount() int32`

GetPasscodePreviousPasscodeBlockCount returns the PasscodePreviousPasscodeBlockCount field if non-nil, zero value otherwise.

### GetPasscodePreviousPasscodeBlockCountOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodePreviousPasscodeBlockCountOk() (int32, bool)`

GetPasscodePreviousPasscodeBlockCountOk returns a tuple with the PasscodePreviousPasscodeBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodePreviousPasscodeBlockCount

`func (o *IosGeneralDeviceConfiguration) HasPasscodePreviousPasscodeBlockCount() bool`

HasPasscodePreviousPasscodeBlockCount returns a boolean if a field has been set.

### SetPasscodePreviousPasscodeBlockCount

`func (o *IosGeneralDeviceConfiguration) SetPasscodePreviousPasscodeBlockCount(v int32)`

SetPasscodePreviousPasscodeBlockCount gets a reference to the given int32 and assigns it to the PasscodePreviousPasscodeBlockCount field.

### SetPasscodePreviousPasscodeBlockCountExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodePreviousPasscodeBlockCountExplicitNull(b bool)`

SetPasscodePreviousPasscodeBlockCountExplicitNull (un)sets PasscodePreviousPasscodeBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodePreviousPasscodeBlockCount value is set to nil even if false is passed
### GetPasscodeSignInFailureCountBeforeWipe

`func (o *IosGeneralDeviceConfiguration) GetPasscodeSignInFailureCountBeforeWipe() int32`

GetPasscodeSignInFailureCountBeforeWipe returns the PasscodeSignInFailureCountBeforeWipe field if non-nil, zero value otherwise.

### GetPasscodeSignInFailureCountBeforeWipeOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeSignInFailureCountBeforeWipeOk() (int32, bool)`

GetPasscodeSignInFailureCountBeforeWipeOk returns a tuple with the PasscodeSignInFailureCountBeforeWipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeSignInFailureCountBeforeWipe

`func (o *IosGeneralDeviceConfiguration) HasPasscodeSignInFailureCountBeforeWipe() bool`

HasPasscodeSignInFailureCountBeforeWipe returns a boolean if a field has been set.

### SetPasscodeSignInFailureCountBeforeWipe

`func (o *IosGeneralDeviceConfiguration) SetPasscodeSignInFailureCountBeforeWipe(v int32)`

SetPasscodeSignInFailureCountBeforeWipe gets a reference to the given int32 and assigns it to the PasscodeSignInFailureCountBeforeWipe field.

### SetPasscodeSignInFailureCountBeforeWipeExplicitNull

`func (o *IosGeneralDeviceConfiguration) SetPasscodeSignInFailureCountBeforeWipeExplicitNull(b bool)`

SetPasscodeSignInFailureCountBeforeWipeExplicitNull (un)sets PasscodeSignInFailureCountBeforeWipe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeSignInFailureCountBeforeWipe value is set to nil even if false is passed
### GetPasscodeRequiredType

`func (o *IosGeneralDeviceConfiguration) GetPasscodeRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasscodeRequiredType returns the PasscodeRequiredType field if non-nil, zero value otherwise.

### GetPasscodeRequiredTypeOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasscodeRequiredTypeOk returns a tuple with the PasscodeRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequiredType

`func (o *IosGeneralDeviceConfiguration) HasPasscodeRequiredType() bool`

HasPasscodeRequiredType returns a boolean if a field has been set.

### SetPasscodeRequiredType

`func (o *IosGeneralDeviceConfiguration) SetPasscodeRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasscodeRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasscodeRequiredType field.

### GetPasscodeRequired

`func (o *IosGeneralDeviceConfiguration) GetPasscodeRequired() bool`

GetPasscodeRequired returns the PasscodeRequired field if non-nil, zero value otherwise.

### GetPasscodeRequiredOk

`func (o *IosGeneralDeviceConfiguration) GetPasscodeRequiredOk() (bool, bool)`

GetPasscodeRequiredOk returns a tuple with the PasscodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequired

`func (o *IosGeneralDeviceConfiguration) HasPasscodeRequired() bool`

HasPasscodeRequired returns a boolean if a field has been set.

### SetPasscodeRequired

`func (o *IosGeneralDeviceConfiguration) SetPasscodeRequired(v bool)`

SetPasscodeRequired gets a reference to the given bool and assigns it to the PasscodeRequired field.

### GetPodcastsBlocked

`func (o *IosGeneralDeviceConfiguration) GetPodcastsBlocked() bool`

GetPodcastsBlocked returns the PodcastsBlocked field if non-nil, zero value otherwise.

### GetPodcastsBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetPodcastsBlockedOk() (bool, bool)`

GetPodcastsBlockedOk returns a tuple with the PodcastsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPodcastsBlocked

`func (o *IosGeneralDeviceConfiguration) HasPodcastsBlocked() bool`

HasPodcastsBlocked returns a boolean if a field has been set.

### SetPodcastsBlocked

`func (o *IosGeneralDeviceConfiguration) SetPodcastsBlocked(v bool)`

SetPodcastsBlocked gets a reference to the given bool and assigns it to the PodcastsBlocked field.

### GetSafariBlockAutofill

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockAutofill() bool`

GetSafariBlockAutofill returns the SafariBlockAutofill field if non-nil, zero value otherwise.

### GetSafariBlockAutofillOk

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockAutofillOk() (bool, bool)`

GetSafariBlockAutofillOk returns a tuple with the SafariBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlockAutofill

`func (o *IosGeneralDeviceConfiguration) HasSafariBlockAutofill() bool`

HasSafariBlockAutofill returns a boolean if a field has been set.

### SetSafariBlockAutofill

`func (o *IosGeneralDeviceConfiguration) SetSafariBlockAutofill(v bool)`

SetSafariBlockAutofill gets a reference to the given bool and assigns it to the SafariBlockAutofill field.

### GetSafariBlockJavaScript

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockJavaScript() bool`

GetSafariBlockJavaScript returns the SafariBlockJavaScript field if non-nil, zero value otherwise.

### GetSafariBlockJavaScriptOk

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockJavaScriptOk() (bool, bool)`

GetSafariBlockJavaScriptOk returns a tuple with the SafariBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlockJavaScript

`func (o *IosGeneralDeviceConfiguration) HasSafariBlockJavaScript() bool`

HasSafariBlockJavaScript returns a boolean if a field has been set.

### SetSafariBlockJavaScript

`func (o *IosGeneralDeviceConfiguration) SetSafariBlockJavaScript(v bool)`

SetSafariBlockJavaScript gets a reference to the given bool and assigns it to the SafariBlockJavaScript field.

### GetSafariBlockPopups

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockPopups() bool`

GetSafariBlockPopups returns the SafariBlockPopups field if non-nil, zero value otherwise.

### GetSafariBlockPopupsOk

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockPopupsOk() (bool, bool)`

GetSafariBlockPopupsOk returns a tuple with the SafariBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlockPopups

`func (o *IosGeneralDeviceConfiguration) HasSafariBlockPopups() bool`

HasSafariBlockPopups returns a boolean if a field has been set.

### SetSafariBlockPopups

`func (o *IosGeneralDeviceConfiguration) SetSafariBlockPopups(v bool)`

SetSafariBlockPopups gets a reference to the given bool and assigns it to the SafariBlockPopups field.

### GetSafariBlocked

`func (o *IosGeneralDeviceConfiguration) GetSafariBlocked() bool`

GetSafariBlocked returns the SafariBlocked field if non-nil, zero value otherwise.

### GetSafariBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetSafariBlockedOk() (bool, bool)`

GetSafariBlockedOk returns a tuple with the SafariBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlocked

`func (o *IosGeneralDeviceConfiguration) HasSafariBlocked() bool`

HasSafariBlocked returns a boolean if a field has been set.

### SetSafariBlocked

`func (o *IosGeneralDeviceConfiguration) SetSafariBlocked(v bool)`

SetSafariBlocked gets a reference to the given bool and assigns it to the SafariBlocked field.

### GetSafariCookieSettings

`func (o *IosGeneralDeviceConfiguration) GetSafariCookieSettings() AnyOfmicrosoftGraphWebBrowserCookieSettings`

GetSafariCookieSettings returns the SafariCookieSettings field if non-nil, zero value otherwise.

### GetSafariCookieSettingsOk

`func (o *IosGeneralDeviceConfiguration) GetSafariCookieSettingsOk() (AnyOfmicrosoftGraphWebBrowserCookieSettings, bool)`

GetSafariCookieSettingsOk returns a tuple with the SafariCookieSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariCookieSettings

`func (o *IosGeneralDeviceConfiguration) HasSafariCookieSettings() bool`

HasSafariCookieSettings returns a boolean if a field has been set.

### SetSafariCookieSettings

`func (o *IosGeneralDeviceConfiguration) SetSafariCookieSettings(v AnyOfmicrosoftGraphWebBrowserCookieSettings)`

SetSafariCookieSettings gets a reference to the given AnyOfmicrosoftGraphWebBrowserCookieSettings and assigns it to the SafariCookieSettings field.

### GetSafariManagedDomains

`func (o *IosGeneralDeviceConfiguration) GetSafariManagedDomains() []string`

GetSafariManagedDomains returns the SafariManagedDomains field if non-nil, zero value otherwise.

### GetSafariManagedDomainsOk

`func (o *IosGeneralDeviceConfiguration) GetSafariManagedDomainsOk() ([]string, bool)`

GetSafariManagedDomainsOk returns a tuple with the SafariManagedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariManagedDomains

`func (o *IosGeneralDeviceConfiguration) HasSafariManagedDomains() bool`

HasSafariManagedDomains returns a boolean if a field has been set.

### SetSafariManagedDomains

`func (o *IosGeneralDeviceConfiguration) SetSafariManagedDomains(v []string)`

SetSafariManagedDomains gets a reference to the given []string and assigns it to the SafariManagedDomains field.

### GetSafariPasswordAutoFillDomains

`func (o *IosGeneralDeviceConfiguration) GetSafariPasswordAutoFillDomains() []string`

GetSafariPasswordAutoFillDomains returns the SafariPasswordAutoFillDomains field if non-nil, zero value otherwise.

### GetSafariPasswordAutoFillDomainsOk

`func (o *IosGeneralDeviceConfiguration) GetSafariPasswordAutoFillDomainsOk() ([]string, bool)`

GetSafariPasswordAutoFillDomainsOk returns a tuple with the SafariPasswordAutoFillDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariPasswordAutoFillDomains

`func (o *IosGeneralDeviceConfiguration) HasSafariPasswordAutoFillDomains() bool`

HasSafariPasswordAutoFillDomains returns a boolean if a field has been set.

### SetSafariPasswordAutoFillDomains

`func (o *IosGeneralDeviceConfiguration) SetSafariPasswordAutoFillDomains(v []string)`

SetSafariPasswordAutoFillDomains gets a reference to the given []string and assigns it to the SafariPasswordAutoFillDomains field.

### GetSafariRequireFraudWarning

`func (o *IosGeneralDeviceConfiguration) GetSafariRequireFraudWarning() bool`

GetSafariRequireFraudWarning returns the SafariRequireFraudWarning field if non-nil, zero value otherwise.

### GetSafariRequireFraudWarningOk

`func (o *IosGeneralDeviceConfiguration) GetSafariRequireFraudWarningOk() (bool, bool)`

GetSafariRequireFraudWarningOk returns a tuple with the SafariRequireFraudWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariRequireFraudWarning

`func (o *IosGeneralDeviceConfiguration) HasSafariRequireFraudWarning() bool`

HasSafariRequireFraudWarning returns a boolean if a field has been set.

### SetSafariRequireFraudWarning

`func (o *IosGeneralDeviceConfiguration) SetSafariRequireFraudWarning(v bool)`

SetSafariRequireFraudWarning gets a reference to the given bool and assigns it to the SafariRequireFraudWarning field.

### GetScreenCaptureBlocked

`func (o *IosGeneralDeviceConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *IosGeneralDeviceConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *IosGeneralDeviceConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetSiriBlocked

`func (o *IosGeneralDeviceConfiguration) GetSiriBlocked() bool`

GetSiriBlocked returns the SiriBlocked field if non-nil, zero value otherwise.

### GetSiriBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetSiriBlockedOk() (bool, bool)`

GetSiriBlockedOk returns a tuple with the SiriBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriBlocked

`func (o *IosGeneralDeviceConfiguration) HasSiriBlocked() bool`

HasSiriBlocked returns a boolean if a field has been set.

### SetSiriBlocked

`func (o *IosGeneralDeviceConfiguration) SetSiriBlocked(v bool)`

SetSiriBlocked gets a reference to the given bool and assigns it to the SiriBlocked field.

### GetSiriBlockedWhenLocked

`func (o *IosGeneralDeviceConfiguration) GetSiriBlockedWhenLocked() bool`

GetSiriBlockedWhenLocked returns the SiriBlockedWhenLocked field if non-nil, zero value otherwise.

### GetSiriBlockedWhenLockedOk

`func (o *IosGeneralDeviceConfiguration) GetSiriBlockedWhenLockedOk() (bool, bool)`

GetSiriBlockedWhenLockedOk returns a tuple with the SiriBlockedWhenLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriBlockedWhenLocked

`func (o *IosGeneralDeviceConfiguration) HasSiriBlockedWhenLocked() bool`

HasSiriBlockedWhenLocked returns a boolean if a field has been set.

### SetSiriBlockedWhenLocked

`func (o *IosGeneralDeviceConfiguration) SetSiriBlockedWhenLocked(v bool)`

SetSiriBlockedWhenLocked gets a reference to the given bool and assigns it to the SiriBlockedWhenLocked field.

### GetSiriBlockUserGeneratedContent

`func (o *IosGeneralDeviceConfiguration) GetSiriBlockUserGeneratedContent() bool`

GetSiriBlockUserGeneratedContent returns the SiriBlockUserGeneratedContent field if non-nil, zero value otherwise.

### GetSiriBlockUserGeneratedContentOk

`func (o *IosGeneralDeviceConfiguration) GetSiriBlockUserGeneratedContentOk() (bool, bool)`

GetSiriBlockUserGeneratedContentOk returns a tuple with the SiriBlockUserGeneratedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriBlockUserGeneratedContent

`func (o *IosGeneralDeviceConfiguration) HasSiriBlockUserGeneratedContent() bool`

HasSiriBlockUserGeneratedContent returns a boolean if a field has been set.

### SetSiriBlockUserGeneratedContent

`func (o *IosGeneralDeviceConfiguration) SetSiriBlockUserGeneratedContent(v bool)`

SetSiriBlockUserGeneratedContent gets a reference to the given bool and assigns it to the SiriBlockUserGeneratedContent field.

### GetSiriRequireProfanityFilter

`func (o *IosGeneralDeviceConfiguration) GetSiriRequireProfanityFilter() bool`

GetSiriRequireProfanityFilter returns the SiriRequireProfanityFilter field if non-nil, zero value otherwise.

### GetSiriRequireProfanityFilterOk

`func (o *IosGeneralDeviceConfiguration) GetSiriRequireProfanityFilterOk() (bool, bool)`

GetSiriRequireProfanityFilterOk returns a tuple with the SiriRequireProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriRequireProfanityFilter

`func (o *IosGeneralDeviceConfiguration) HasSiriRequireProfanityFilter() bool`

HasSiriRequireProfanityFilter returns a boolean if a field has been set.

### SetSiriRequireProfanityFilter

`func (o *IosGeneralDeviceConfiguration) SetSiriRequireProfanityFilter(v bool)`

SetSiriRequireProfanityFilter gets a reference to the given bool and assigns it to the SiriRequireProfanityFilter field.

### GetSpotlightBlockInternetResults

`func (o *IosGeneralDeviceConfiguration) GetSpotlightBlockInternetResults() bool`

GetSpotlightBlockInternetResults returns the SpotlightBlockInternetResults field if non-nil, zero value otherwise.

### GetSpotlightBlockInternetResultsOk

`func (o *IosGeneralDeviceConfiguration) GetSpotlightBlockInternetResultsOk() (bool, bool)`

GetSpotlightBlockInternetResultsOk returns a tuple with the SpotlightBlockInternetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpotlightBlockInternetResults

`func (o *IosGeneralDeviceConfiguration) HasSpotlightBlockInternetResults() bool`

HasSpotlightBlockInternetResults returns a boolean if a field has been set.

### SetSpotlightBlockInternetResults

`func (o *IosGeneralDeviceConfiguration) SetSpotlightBlockInternetResults(v bool)`

SetSpotlightBlockInternetResults gets a reference to the given bool and assigns it to the SpotlightBlockInternetResults field.

### GetVoiceDialingBlocked

`func (o *IosGeneralDeviceConfiguration) GetVoiceDialingBlocked() bool`

GetVoiceDialingBlocked returns the VoiceDialingBlocked field if non-nil, zero value otherwise.

### GetVoiceDialingBlockedOk

`func (o *IosGeneralDeviceConfiguration) GetVoiceDialingBlockedOk() (bool, bool)`

GetVoiceDialingBlockedOk returns a tuple with the VoiceDialingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceDialingBlocked

`func (o *IosGeneralDeviceConfiguration) HasVoiceDialingBlocked() bool`

HasVoiceDialingBlocked returns a boolean if a field has been set.

### SetVoiceDialingBlocked

`func (o *IosGeneralDeviceConfiguration) SetVoiceDialingBlocked(v bool)`

SetVoiceDialingBlocked gets a reference to the given bool and assigns it to the VoiceDialingBlocked field.

### GetWallpaperBlockModification

`func (o *IosGeneralDeviceConfiguration) GetWallpaperBlockModification() bool`

GetWallpaperBlockModification returns the WallpaperBlockModification field if non-nil, zero value otherwise.

### GetWallpaperBlockModificationOk

`func (o *IosGeneralDeviceConfiguration) GetWallpaperBlockModificationOk() (bool, bool)`

GetWallpaperBlockModificationOk returns a tuple with the WallpaperBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWallpaperBlockModification

`func (o *IosGeneralDeviceConfiguration) HasWallpaperBlockModification() bool`

HasWallpaperBlockModification returns a boolean if a field has been set.

### SetWallpaperBlockModification

`func (o *IosGeneralDeviceConfiguration) SetWallpaperBlockModification(v bool)`

SetWallpaperBlockModification gets a reference to the given bool and assigns it to the WallpaperBlockModification field.

### GetWiFiConnectOnlyToConfiguredNetworks

`func (o *IosGeneralDeviceConfiguration) GetWiFiConnectOnlyToConfiguredNetworks() bool`

GetWiFiConnectOnlyToConfiguredNetworks returns the WiFiConnectOnlyToConfiguredNetworks field if non-nil, zero value otherwise.

### GetWiFiConnectOnlyToConfiguredNetworksOk

`func (o *IosGeneralDeviceConfiguration) GetWiFiConnectOnlyToConfiguredNetworksOk() (bool, bool)`

GetWiFiConnectOnlyToConfiguredNetworksOk returns a tuple with the WiFiConnectOnlyToConfiguredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiConnectOnlyToConfiguredNetworks

`func (o *IosGeneralDeviceConfiguration) HasWiFiConnectOnlyToConfiguredNetworks() bool`

HasWiFiConnectOnlyToConfiguredNetworks returns a boolean if a field has been set.

### SetWiFiConnectOnlyToConfiguredNetworks

`func (o *IosGeneralDeviceConfiguration) SetWiFiConnectOnlyToConfiguredNetworks(v bool)`

SetWiFiConnectOnlyToConfiguredNetworks gets a reference to the given bool and assigns it to the WiFiConnectOnlyToConfiguredNetworks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


