# MicrosoftGraphIosGeneralDeviceConfiguration

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

### GetId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAccountBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAccountBlockModification() bool`

GetAccountBlockModification returns the AccountBlockModification field if non-nil, zero value otherwise.

### GetAccountBlockModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAccountBlockModificationOk() (bool, bool)`

GetAccountBlockModificationOk returns a tuple with the AccountBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAccountBlockModification() bool`

HasAccountBlockModification returns a boolean if a field has been set.

### SetAccountBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAccountBlockModification(v bool)`

SetAccountBlockModification gets a reference to the given bool and assigns it to the AccountBlockModification field.

### GetActivationLockAllowWhenSupervised

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetActivationLockAllowWhenSupervised() bool`

GetActivationLockAllowWhenSupervised returns the ActivationLockAllowWhenSupervised field if non-nil, zero value otherwise.

### GetActivationLockAllowWhenSupervisedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetActivationLockAllowWhenSupervisedOk() (bool, bool)`

GetActivationLockAllowWhenSupervisedOk returns a tuple with the ActivationLockAllowWhenSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationLockAllowWhenSupervised

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasActivationLockAllowWhenSupervised() bool`

HasActivationLockAllowWhenSupervised returns a boolean if a field has been set.

### SetActivationLockAllowWhenSupervised

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetActivationLockAllowWhenSupervised(v bool)`

SetActivationLockAllowWhenSupervised gets a reference to the given bool and assigns it to the ActivationLockAllowWhenSupervised field.

### GetAirDropBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAirDropBlocked() bool`

GetAirDropBlocked returns the AirDropBlocked field if non-nil, zero value otherwise.

### GetAirDropBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAirDropBlockedOk() (bool, bool)`

GetAirDropBlockedOk returns a tuple with the AirDropBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAirDropBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAirDropBlocked() bool`

HasAirDropBlocked returns a boolean if a field has been set.

### SetAirDropBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAirDropBlocked(v bool)`

SetAirDropBlocked gets a reference to the given bool and assigns it to the AirDropBlocked field.

### GetAirDropForceUnmanagedDropTarget

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAirDropForceUnmanagedDropTarget() bool`

GetAirDropForceUnmanagedDropTarget returns the AirDropForceUnmanagedDropTarget field if non-nil, zero value otherwise.

### GetAirDropForceUnmanagedDropTargetOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAirDropForceUnmanagedDropTargetOk() (bool, bool)`

GetAirDropForceUnmanagedDropTargetOk returns a tuple with the AirDropForceUnmanagedDropTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAirDropForceUnmanagedDropTarget

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAirDropForceUnmanagedDropTarget() bool`

HasAirDropForceUnmanagedDropTarget returns a boolean if a field has been set.

### SetAirDropForceUnmanagedDropTarget

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAirDropForceUnmanagedDropTarget(v bool)`

SetAirDropForceUnmanagedDropTarget gets a reference to the given bool and assigns it to the AirDropForceUnmanagedDropTarget field.

### GetAirPlayForcePairingPasswordForOutgoingRequests

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAirPlayForcePairingPasswordForOutgoingRequests() bool`

GetAirPlayForcePairingPasswordForOutgoingRequests returns the AirPlayForcePairingPasswordForOutgoingRequests field if non-nil, zero value otherwise.

### GetAirPlayForcePairingPasswordForOutgoingRequestsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAirPlayForcePairingPasswordForOutgoingRequestsOk() (bool, bool)`

GetAirPlayForcePairingPasswordForOutgoingRequestsOk returns a tuple with the AirPlayForcePairingPasswordForOutgoingRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAirPlayForcePairingPasswordForOutgoingRequests

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAirPlayForcePairingPasswordForOutgoingRequests() bool`

HasAirPlayForcePairingPasswordForOutgoingRequests returns a boolean if a field has been set.

### SetAirPlayForcePairingPasswordForOutgoingRequests

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAirPlayForcePairingPasswordForOutgoingRequests(v bool)`

SetAirPlayForcePairingPasswordForOutgoingRequests gets a reference to the given bool and assigns it to the AirPlayForcePairingPasswordForOutgoingRequests field.

### GetAppleWatchBlockPairing

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppleWatchBlockPairing() bool`

GetAppleWatchBlockPairing returns the AppleWatchBlockPairing field if non-nil, zero value otherwise.

### GetAppleWatchBlockPairingOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppleWatchBlockPairingOk() (bool, bool)`

GetAppleWatchBlockPairingOk returns a tuple with the AppleWatchBlockPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleWatchBlockPairing

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppleWatchBlockPairing() bool`

HasAppleWatchBlockPairing returns a boolean if a field has been set.

### SetAppleWatchBlockPairing

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppleWatchBlockPairing(v bool)`

SetAppleWatchBlockPairing gets a reference to the given bool and assigns it to the AppleWatchBlockPairing field.

### GetAppleWatchForceWristDetection

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppleWatchForceWristDetection() bool`

GetAppleWatchForceWristDetection returns the AppleWatchForceWristDetection field if non-nil, zero value otherwise.

### GetAppleWatchForceWristDetectionOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppleWatchForceWristDetectionOk() (bool, bool)`

GetAppleWatchForceWristDetectionOk returns a tuple with the AppleWatchForceWristDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleWatchForceWristDetection

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppleWatchForceWristDetection() bool`

HasAppleWatchForceWristDetection returns a boolean if a field has been set.

### SetAppleWatchForceWristDetection

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppleWatchForceWristDetection(v bool)`

SetAppleWatchForceWristDetection gets a reference to the given bool and assigns it to the AppleWatchForceWristDetection field.

### GetAppleNewsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppleNewsBlocked() bool`

GetAppleNewsBlocked returns the AppleNewsBlocked field if non-nil, zero value otherwise.

### GetAppleNewsBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppleNewsBlockedOk() (bool, bool)`

GetAppleNewsBlockedOk returns a tuple with the AppleNewsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleNewsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppleNewsBlocked() bool`

HasAppleNewsBlocked returns a boolean if a field has been set.

### SetAppleNewsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppleNewsBlocked(v bool)`

SetAppleNewsBlocked gets a reference to the given bool and assigns it to the AppleNewsBlocked field.

### GetAppsSingleAppModeList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppsSingleAppModeList() []AnyOfmicrosoftGraphAppListItem`

GetAppsSingleAppModeList returns the AppsSingleAppModeList field if non-nil, zero value otherwise.

### GetAppsSingleAppModeListOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppsSingleAppModeListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsSingleAppModeListOk returns a tuple with the AppsSingleAppModeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsSingleAppModeList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppsSingleAppModeList() bool`

HasAppsSingleAppModeList returns a boolean if a field has been set.

### SetAppsSingleAppModeList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppsSingleAppModeList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsSingleAppModeList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsSingleAppModeList field.

### GetAppsVisibilityList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppsVisibilityList() []AnyOfmicrosoftGraphAppListItem`

GetAppsVisibilityList returns the AppsVisibilityList field if non-nil, zero value otherwise.

### GetAppsVisibilityListOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppsVisibilityListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetAppsVisibilityListOk returns a tuple with the AppsVisibilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsVisibilityList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppsVisibilityList() bool`

HasAppsVisibilityList returns a boolean if a field has been set.

### SetAppsVisibilityList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppsVisibilityList(v []AnyOfmicrosoftGraphAppListItem)`

SetAppsVisibilityList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the AppsVisibilityList field.

### GetAppsVisibilityListType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppsVisibilityListType() AnyOfmicrosoftGraphAppListType`

GetAppsVisibilityListType returns the AppsVisibilityListType field if non-nil, zero value otherwise.

### GetAppsVisibilityListTypeOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppsVisibilityListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetAppsVisibilityListTypeOk returns a tuple with the AppsVisibilityListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsVisibilityListType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppsVisibilityListType() bool`

HasAppsVisibilityListType returns a boolean if a field has been set.

### SetAppsVisibilityListType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppsVisibilityListType(v AnyOfmicrosoftGraphAppListType)`

SetAppsVisibilityListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the AppsVisibilityListType field.

### GetAppStoreBlockAutomaticDownloads

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockAutomaticDownloads() bool`

GetAppStoreBlockAutomaticDownloads returns the AppStoreBlockAutomaticDownloads field if non-nil, zero value otherwise.

### GetAppStoreBlockAutomaticDownloadsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockAutomaticDownloadsOk() (bool, bool)`

GetAppStoreBlockAutomaticDownloadsOk returns a tuple with the AppStoreBlockAutomaticDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlockAutomaticDownloads

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppStoreBlockAutomaticDownloads() bool`

HasAppStoreBlockAutomaticDownloads returns a boolean if a field has been set.

### SetAppStoreBlockAutomaticDownloads

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppStoreBlockAutomaticDownloads(v bool)`

SetAppStoreBlockAutomaticDownloads gets a reference to the given bool and assigns it to the AppStoreBlockAutomaticDownloads field.

### GetAppStoreBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlocked() bool`

GetAppStoreBlocked returns the AppStoreBlocked field if non-nil, zero value otherwise.

### GetAppStoreBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockedOk() (bool, bool)`

GetAppStoreBlockedOk returns a tuple with the AppStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppStoreBlocked() bool`

HasAppStoreBlocked returns a boolean if a field has been set.

### SetAppStoreBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppStoreBlocked(v bool)`

SetAppStoreBlocked gets a reference to the given bool and assigns it to the AppStoreBlocked field.

### GetAppStoreBlockInAppPurchases

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockInAppPurchases() bool`

GetAppStoreBlockInAppPurchases returns the AppStoreBlockInAppPurchases field if non-nil, zero value otherwise.

### GetAppStoreBlockInAppPurchasesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockInAppPurchasesOk() (bool, bool)`

GetAppStoreBlockInAppPurchasesOk returns a tuple with the AppStoreBlockInAppPurchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlockInAppPurchases

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppStoreBlockInAppPurchases() bool`

HasAppStoreBlockInAppPurchases returns a boolean if a field has been set.

### SetAppStoreBlockInAppPurchases

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppStoreBlockInAppPurchases(v bool)`

SetAppStoreBlockInAppPurchases gets a reference to the given bool and assigns it to the AppStoreBlockInAppPurchases field.

### GetAppStoreBlockUIAppInstallation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockUIAppInstallation() bool`

GetAppStoreBlockUIAppInstallation returns the AppStoreBlockUIAppInstallation field if non-nil, zero value otherwise.

### GetAppStoreBlockUIAppInstallationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreBlockUIAppInstallationOk() (bool, bool)`

GetAppStoreBlockUIAppInstallationOk returns a tuple with the AppStoreBlockUIAppInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreBlockUIAppInstallation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppStoreBlockUIAppInstallation() bool`

HasAppStoreBlockUIAppInstallation returns a boolean if a field has been set.

### SetAppStoreBlockUIAppInstallation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppStoreBlockUIAppInstallation(v bool)`

SetAppStoreBlockUIAppInstallation gets a reference to the given bool and assigns it to the AppStoreBlockUIAppInstallation field.

### GetAppStoreRequirePassword

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreRequirePassword() bool`

GetAppStoreRequirePassword returns the AppStoreRequirePassword field if non-nil, zero value otherwise.

### GetAppStoreRequirePasswordOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetAppStoreRequirePasswordOk() (bool, bool)`

GetAppStoreRequirePasswordOk returns a tuple with the AppStoreRequirePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreRequirePassword

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasAppStoreRequirePassword() bool`

HasAppStoreRequirePassword returns a boolean if a field has been set.

### SetAppStoreRequirePassword

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetAppStoreRequirePassword(v bool)`

SetAppStoreRequirePassword gets a reference to the given bool and assigns it to the AppStoreRequirePassword field.

### GetBluetoothBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetBluetoothBlockModification() bool`

GetBluetoothBlockModification returns the BluetoothBlockModification field if non-nil, zero value otherwise.

### GetBluetoothBlockModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetBluetoothBlockModificationOk() (bool, bool)`

GetBluetoothBlockModificationOk returns a tuple with the BluetoothBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasBluetoothBlockModification() bool`

HasBluetoothBlockModification returns a boolean if a field has been set.

### SetBluetoothBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetBluetoothBlockModification(v bool)`

SetBluetoothBlockModification gets a reference to the given bool and assigns it to the BluetoothBlockModification field.

### GetCameraBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetCellularBlockDataRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockDataRoaming() bool`

GetCellularBlockDataRoaming returns the CellularBlockDataRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataRoamingOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockDataRoamingOk() (bool, bool)`

GetCellularBlockDataRoamingOk returns a tuple with the CellularBlockDataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCellularBlockDataRoaming() bool`

HasCellularBlockDataRoaming returns a boolean if a field has been set.

### SetCellularBlockDataRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCellularBlockDataRoaming(v bool)`

SetCellularBlockDataRoaming gets a reference to the given bool and assigns it to the CellularBlockDataRoaming field.

### GetCellularBlockGlobalBackgroundFetchWhileRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockGlobalBackgroundFetchWhileRoaming() bool`

GetCellularBlockGlobalBackgroundFetchWhileRoaming returns the CellularBlockGlobalBackgroundFetchWhileRoaming field if non-nil, zero value otherwise.

### GetCellularBlockGlobalBackgroundFetchWhileRoamingOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockGlobalBackgroundFetchWhileRoamingOk() (bool, bool)`

GetCellularBlockGlobalBackgroundFetchWhileRoamingOk returns a tuple with the CellularBlockGlobalBackgroundFetchWhileRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockGlobalBackgroundFetchWhileRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCellularBlockGlobalBackgroundFetchWhileRoaming() bool`

HasCellularBlockGlobalBackgroundFetchWhileRoaming returns a boolean if a field has been set.

### SetCellularBlockGlobalBackgroundFetchWhileRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCellularBlockGlobalBackgroundFetchWhileRoaming(v bool)`

SetCellularBlockGlobalBackgroundFetchWhileRoaming gets a reference to the given bool and assigns it to the CellularBlockGlobalBackgroundFetchWhileRoaming field.

### GetCellularBlockPerAppDataModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockPerAppDataModification() bool`

GetCellularBlockPerAppDataModification returns the CellularBlockPerAppDataModification field if non-nil, zero value otherwise.

### GetCellularBlockPerAppDataModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockPerAppDataModificationOk() (bool, bool)`

GetCellularBlockPerAppDataModificationOk returns a tuple with the CellularBlockPerAppDataModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockPerAppDataModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCellularBlockPerAppDataModification() bool`

HasCellularBlockPerAppDataModification returns a boolean if a field has been set.

### SetCellularBlockPerAppDataModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCellularBlockPerAppDataModification(v bool)`

SetCellularBlockPerAppDataModification gets a reference to the given bool and assigns it to the CellularBlockPerAppDataModification field.

### GetCellularBlockPersonalHotspot

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspot() bool`

GetCellularBlockPersonalHotspot returns the CellularBlockPersonalHotspot field if non-nil, zero value otherwise.

### GetCellularBlockPersonalHotspotOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspotOk() (bool, bool)`

GetCellularBlockPersonalHotspotOk returns a tuple with the CellularBlockPersonalHotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockPersonalHotspot

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCellularBlockPersonalHotspot() bool`

HasCellularBlockPersonalHotspot returns a boolean if a field has been set.

### SetCellularBlockPersonalHotspot

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCellularBlockPersonalHotspot(v bool)`

SetCellularBlockPersonalHotspot gets a reference to the given bool and assigns it to the CellularBlockPersonalHotspot field.

### GetCellularBlockVoiceRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming() bool`

GetCellularBlockVoiceRoaming returns the CellularBlockVoiceRoaming field if non-nil, zero value otherwise.

### GetCellularBlockVoiceRoamingOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCellularBlockVoiceRoamingOk() (bool, bool)`

GetCellularBlockVoiceRoamingOk returns a tuple with the CellularBlockVoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVoiceRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCellularBlockVoiceRoaming() bool`

HasCellularBlockVoiceRoaming returns a boolean if a field has been set.

### SetCellularBlockVoiceRoaming

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(v bool)`

SetCellularBlockVoiceRoaming gets a reference to the given bool and assigns it to the CellularBlockVoiceRoaming field.

### GetCertificatesBlockUntrustedTlsCertificates

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCertificatesBlockUntrustedTlsCertificates() bool`

GetCertificatesBlockUntrustedTlsCertificates returns the CertificatesBlockUntrustedTlsCertificates field if non-nil, zero value otherwise.

### GetCertificatesBlockUntrustedTlsCertificatesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCertificatesBlockUntrustedTlsCertificatesOk() (bool, bool)`

GetCertificatesBlockUntrustedTlsCertificatesOk returns a tuple with the CertificatesBlockUntrustedTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificatesBlockUntrustedTlsCertificates

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCertificatesBlockUntrustedTlsCertificates() bool`

HasCertificatesBlockUntrustedTlsCertificates returns a boolean if a field has been set.

### SetCertificatesBlockUntrustedTlsCertificates

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCertificatesBlockUntrustedTlsCertificates(v bool)`

SetCertificatesBlockUntrustedTlsCertificates gets a reference to the given bool and assigns it to the CertificatesBlockUntrustedTlsCertificates field.

### GetClassroomAppBlockRemoteScreenObservation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservation() bool`

GetClassroomAppBlockRemoteScreenObservation returns the ClassroomAppBlockRemoteScreenObservation field if non-nil, zero value otherwise.

### GetClassroomAppBlockRemoteScreenObservationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservationOk() (bool, bool)`

GetClassroomAppBlockRemoteScreenObservationOk returns a tuple with the ClassroomAppBlockRemoteScreenObservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassroomAppBlockRemoteScreenObservation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasClassroomAppBlockRemoteScreenObservation() bool`

HasClassroomAppBlockRemoteScreenObservation returns a boolean if a field has been set.

### SetClassroomAppBlockRemoteScreenObservation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetClassroomAppBlockRemoteScreenObservation(v bool)`

SetClassroomAppBlockRemoteScreenObservation gets a reference to the given bool and assigns it to the ClassroomAppBlockRemoteScreenObservation field.

### GetClassroomAppForceUnpromptedScreenObservation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservation() bool`

GetClassroomAppForceUnpromptedScreenObservation returns the ClassroomAppForceUnpromptedScreenObservation field if non-nil, zero value otherwise.

### GetClassroomAppForceUnpromptedScreenObservationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservationOk() (bool, bool)`

GetClassroomAppForceUnpromptedScreenObservationOk returns a tuple with the ClassroomAppForceUnpromptedScreenObservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassroomAppForceUnpromptedScreenObservation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasClassroomAppForceUnpromptedScreenObservation() bool`

HasClassroomAppForceUnpromptedScreenObservation returns a boolean if a field has been set.

### SetClassroomAppForceUnpromptedScreenObservation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetClassroomAppForceUnpromptedScreenObservation(v bool)`

SetClassroomAppForceUnpromptedScreenObservation gets a reference to the given bool and assigns it to the ClassroomAppForceUnpromptedScreenObservation field.

### GetCompliantAppsList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetConfigurationProfileBlockChanges

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetConfigurationProfileBlockChanges() bool`

GetConfigurationProfileBlockChanges returns the ConfigurationProfileBlockChanges field if non-nil, zero value otherwise.

### GetConfigurationProfileBlockChangesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetConfigurationProfileBlockChangesOk() (bool, bool)`

GetConfigurationProfileBlockChangesOk returns a tuple with the ConfigurationProfileBlockChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationProfileBlockChanges

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasConfigurationProfileBlockChanges() bool`

HasConfigurationProfileBlockChanges returns a boolean if a field has been set.

### SetConfigurationProfileBlockChanges

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetConfigurationProfileBlockChanges(v bool)`

SetConfigurationProfileBlockChanges gets a reference to the given bool and assigns it to the ConfigurationProfileBlockChanges field.

### GetDefinitionLookupBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDefinitionLookupBlocked() bool`

GetDefinitionLookupBlocked returns the DefinitionLookupBlocked field if non-nil, zero value otherwise.

### GetDefinitionLookupBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDefinitionLookupBlockedOk() (bool, bool)`

GetDefinitionLookupBlockedOk returns a tuple with the DefinitionLookupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefinitionLookupBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDefinitionLookupBlocked() bool`

HasDefinitionLookupBlocked returns a boolean if a field has been set.

### SetDefinitionLookupBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDefinitionLookupBlocked(v bool)`

SetDefinitionLookupBlocked gets a reference to the given bool and assigns it to the DefinitionLookupBlocked field.

### GetDeviceBlockEnableRestrictions

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceBlockEnableRestrictions() bool`

GetDeviceBlockEnableRestrictions returns the DeviceBlockEnableRestrictions field if non-nil, zero value otherwise.

### GetDeviceBlockEnableRestrictionsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceBlockEnableRestrictionsOk() (bool, bool)`

GetDeviceBlockEnableRestrictionsOk returns a tuple with the DeviceBlockEnableRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceBlockEnableRestrictions

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDeviceBlockEnableRestrictions() bool`

HasDeviceBlockEnableRestrictions returns a boolean if a field has been set.

### SetDeviceBlockEnableRestrictions

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceBlockEnableRestrictions(v bool)`

SetDeviceBlockEnableRestrictions gets a reference to the given bool and assigns it to the DeviceBlockEnableRestrictions field.

### GetDeviceBlockEraseContentAndSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceBlockEraseContentAndSettings() bool`

GetDeviceBlockEraseContentAndSettings returns the DeviceBlockEraseContentAndSettings field if non-nil, zero value otherwise.

### GetDeviceBlockEraseContentAndSettingsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceBlockEraseContentAndSettingsOk() (bool, bool)`

GetDeviceBlockEraseContentAndSettingsOk returns a tuple with the DeviceBlockEraseContentAndSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceBlockEraseContentAndSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDeviceBlockEraseContentAndSettings() bool`

HasDeviceBlockEraseContentAndSettings returns a boolean if a field has been set.

### SetDeviceBlockEraseContentAndSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceBlockEraseContentAndSettings(v bool)`

SetDeviceBlockEraseContentAndSettings gets a reference to the given bool and assigns it to the DeviceBlockEraseContentAndSettings field.

### GetDeviceBlockNameModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceBlockNameModification() bool`

GetDeviceBlockNameModification returns the DeviceBlockNameModification field if non-nil, zero value otherwise.

### GetDeviceBlockNameModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDeviceBlockNameModificationOk() (bool, bool)`

GetDeviceBlockNameModificationOk returns a tuple with the DeviceBlockNameModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceBlockNameModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDeviceBlockNameModification() bool`

HasDeviceBlockNameModification returns a boolean if a field has been set.

### SetDeviceBlockNameModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDeviceBlockNameModification(v bool)`

SetDeviceBlockNameModification gets a reference to the given bool and assigns it to the DeviceBlockNameModification field.

### GetDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission() bool`

GetDiagnosticDataBlockSubmission returns the DiagnosticDataBlockSubmission field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionOk returns a tuple with the DiagnosticDataBlockSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDiagnosticDataBlockSubmission() bool`

HasDiagnosticDataBlockSubmission returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(v bool)`

SetDiagnosticDataBlockSubmission gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmission field.

### GetDiagnosticDataBlockSubmissionModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionModification() bool`

GetDiagnosticDataBlockSubmissionModification returns the DiagnosticDataBlockSubmissionModification field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionModificationOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionModificationOk returns a tuple with the DiagnosticDataBlockSubmissionModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmissionModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDiagnosticDataBlockSubmissionModification() bool`

HasDiagnosticDataBlockSubmissionModification returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmissionModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmissionModification(v bool)`

SetDiagnosticDataBlockSubmissionModification gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmissionModification field.

### GetDocumentsBlockManagedDocumentsInUnmanagedApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDocumentsBlockManagedDocumentsInUnmanagedApps() bool`

GetDocumentsBlockManagedDocumentsInUnmanagedApps returns the DocumentsBlockManagedDocumentsInUnmanagedApps field if non-nil, zero value otherwise.

### GetDocumentsBlockManagedDocumentsInUnmanagedAppsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDocumentsBlockManagedDocumentsInUnmanagedAppsOk() (bool, bool)`

GetDocumentsBlockManagedDocumentsInUnmanagedAppsOk returns a tuple with the DocumentsBlockManagedDocumentsInUnmanagedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDocumentsBlockManagedDocumentsInUnmanagedApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDocumentsBlockManagedDocumentsInUnmanagedApps() bool`

HasDocumentsBlockManagedDocumentsInUnmanagedApps returns a boolean if a field has been set.

### SetDocumentsBlockManagedDocumentsInUnmanagedApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDocumentsBlockManagedDocumentsInUnmanagedApps(v bool)`

SetDocumentsBlockManagedDocumentsInUnmanagedApps gets a reference to the given bool and assigns it to the DocumentsBlockManagedDocumentsInUnmanagedApps field.

### GetDocumentsBlockUnmanagedDocumentsInManagedApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDocumentsBlockUnmanagedDocumentsInManagedApps() bool`

GetDocumentsBlockUnmanagedDocumentsInManagedApps returns the DocumentsBlockUnmanagedDocumentsInManagedApps field if non-nil, zero value otherwise.

### GetDocumentsBlockUnmanagedDocumentsInManagedAppsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetDocumentsBlockUnmanagedDocumentsInManagedAppsOk() (bool, bool)`

GetDocumentsBlockUnmanagedDocumentsInManagedAppsOk returns a tuple with the DocumentsBlockUnmanagedDocumentsInManagedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDocumentsBlockUnmanagedDocumentsInManagedApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasDocumentsBlockUnmanagedDocumentsInManagedApps() bool`

HasDocumentsBlockUnmanagedDocumentsInManagedApps returns a boolean if a field has been set.

### SetDocumentsBlockUnmanagedDocumentsInManagedApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetDocumentsBlockUnmanagedDocumentsInManagedApps(v bool)`

SetDocumentsBlockUnmanagedDocumentsInManagedApps gets a reference to the given bool and assigns it to the DocumentsBlockUnmanagedDocumentsInManagedApps field.

### GetEmailInDomainSuffixes

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetEmailInDomainSuffixes() []string`

GetEmailInDomainSuffixes returns the EmailInDomainSuffixes field if non-nil, zero value otherwise.

### GetEmailInDomainSuffixesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetEmailInDomainSuffixesOk() ([]string, bool)`

GetEmailInDomainSuffixesOk returns a tuple with the EmailInDomainSuffixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailInDomainSuffixes

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasEmailInDomainSuffixes() bool`

HasEmailInDomainSuffixes returns a boolean if a field has been set.

### SetEmailInDomainSuffixes

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetEmailInDomainSuffixes(v []string)`

SetEmailInDomainSuffixes gets a reference to the given []string and assigns it to the EmailInDomainSuffixes field.

### GetEnterpriseAppBlockTrust

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrust() bool`

GetEnterpriseAppBlockTrust returns the EnterpriseAppBlockTrust field if non-nil, zero value otherwise.

### GetEnterpriseAppBlockTrustOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustOk() (bool, bool)`

GetEnterpriseAppBlockTrustOk returns a tuple with the EnterpriseAppBlockTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseAppBlockTrust

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasEnterpriseAppBlockTrust() bool`

HasEnterpriseAppBlockTrust returns a boolean if a field has been set.

### SetEnterpriseAppBlockTrust

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrust(v bool)`

SetEnterpriseAppBlockTrust gets a reference to the given bool and assigns it to the EnterpriseAppBlockTrust field.

### GetEnterpriseAppBlockTrustModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustModification() bool`

GetEnterpriseAppBlockTrustModification returns the EnterpriseAppBlockTrustModification field if non-nil, zero value otherwise.

### GetEnterpriseAppBlockTrustModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustModificationOk() (bool, bool)`

GetEnterpriseAppBlockTrustModificationOk returns a tuple with the EnterpriseAppBlockTrustModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseAppBlockTrustModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasEnterpriseAppBlockTrustModification() bool`

HasEnterpriseAppBlockTrustModification returns a boolean if a field has been set.

### SetEnterpriseAppBlockTrustModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrustModification(v bool)`

SetEnterpriseAppBlockTrustModification gets a reference to the given bool and assigns it to the EnterpriseAppBlockTrustModification field.

### GetFaceTimeBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetFaceTimeBlocked() bool`

GetFaceTimeBlocked returns the FaceTimeBlocked field if non-nil, zero value otherwise.

### GetFaceTimeBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetFaceTimeBlockedOk() (bool, bool)`

GetFaceTimeBlockedOk returns a tuple with the FaceTimeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaceTimeBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasFaceTimeBlocked() bool`

HasFaceTimeBlocked returns a boolean if a field has been set.

### SetFaceTimeBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetFaceTimeBlocked(v bool)`

SetFaceTimeBlocked gets a reference to the given bool and assigns it to the FaceTimeBlocked field.

### GetFindMyFriendsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetFindMyFriendsBlocked() bool`

GetFindMyFriendsBlocked returns the FindMyFriendsBlocked field if non-nil, zero value otherwise.

### GetFindMyFriendsBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetFindMyFriendsBlockedOk() (bool, bool)`

GetFindMyFriendsBlockedOk returns a tuple with the FindMyFriendsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFindMyFriendsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasFindMyFriendsBlocked() bool`

HasFindMyFriendsBlocked returns a boolean if a field has been set.

### SetFindMyFriendsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetFindMyFriendsBlocked(v bool)`

SetFindMyFriendsBlocked gets a reference to the given bool and assigns it to the FindMyFriendsBlocked field.

### GetGamingBlockGameCenterFriends

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetGamingBlockGameCenterFriends() bool`

GetGamingBlockGameCenterFriends returns the GamingBlockGameCenterFriends field if non-nil, zero value otherwise.

### GetGamingBlockGameCenterFriendsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetGamingBlockGameCenterFriendsOk() (bool, bool)`

GetGamingBlockGameCenterFriendsOk returns a tuple with the GamingBlockGameCenterFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGamingBlockGameCenterFriends

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasGamingBlockGameCenterFriends() bool`

HasGamingBlockGameCenterFriends returns a boolean if a field has been set.

### SetGamingBlockGameCenterFriends

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetGamingBlockGameCenterFriends(v bool)`

SetGamingBlockGameCenterFriends gets a reference to the given bool and assigns it to the GamingBlockGameCenterFriends field.

### GetGamingBlockMultiplayer

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetGamingBlockMultiplayer() bool`

GetGamingBlockMultiplayer returns the GamingBlockMultiplayer field if non-nil, zero value otherwise.

### GetGamingBlockMultiplayerOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetGamingBlockMultiplayerOk() (bool, bool)`

GetGamingBlockMultiplayerOk returns a tuple with the GamingBlockMultiplayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGamingBlockMultiplayer

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasGamingBlockMultiplayer() bool`

HasGamingBlockMultiplayer returns a boolean if a field has been set.

### SetGamingBlockMultiplayer

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetGamingBlockMultiplayer(v bool)`

SetGamingBlockMultiplayer gets a reference to the given bool and assigns it to the GamingBlockMultiplayer field.

### GetGameCenterBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetGameCenterBlocked() bool`

GetGameCenterBlocked returns the GameCenterBlocked field if non-nil, zero value otherwise.

### GetGameCenterBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetGameCenterBlockedOk() (bool, bool)`

GetGameCenterBlockedOk returns a tuple with the GameCenterBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGameCenterBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasGameCenterBlocked() bool`

HasGameCenterBlocked returns a boolean if a field has been set.

### SetGameCenterBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetGameCenterBlocked(v bool)`

SetGameCenterBlocked gets a reference to the given bool and assigns it to the GameCenterBlocked field.

### GetHostPairingBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetHostPairingBlocked() bool`

GetHostPairingBlocked returns the HostPairingBlocked field if non-nil, zero value otherwise.

### GetHostPairingBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetHostPairingBlockedOk() (bool, bool)`

GetHostPairingBlockedOk returns a tuple with the HostPairingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostPairingBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasHostPairingBlocked() bool`

HasHostPairingBlocked returns a boolean if a field has been set.

### SetHostPairingBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetHostPairingBlocked(v bool)`

SetHostPairingBlocked gets a reference to the given bool and assigns it to the HostPairingBlocked field.

### GetIBooksStoreBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetIBooksStoreBlocked() bool`

GetIBooksStoreBlocked returns the IBooksStoreBlocked field if non-nil, zero value otherwise.

### GetIBooksStoreBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetIBooksStoreBlockedOk() (bool, bool)`

GetIBooksStoreBlockedOk returns a tuple with the IBooksStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIBooksStoreBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasIBooksStoreBlocked() bool`

HasIBooksStoreBlocked returns a boolean if a field has been set.

### SetIBooksStoreBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetIBooksStoreBlocked(v bool)`

SetIBooksStoreBlocked gets a reference to the given bool and assigns it to the IBooksStoreBlocked field.

### GetIBooksStoreBlockErotica

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetIBooksStoreBlockErotica() bool`

GetIBooksStoreBlockErotica returns the IBooksStoreBlockErotica field if non-nil, zero value otherwise.

### GetIBooksStoreBlockEroticaOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetIBooksStoreBlockEroticaOk() (bool, bool)`

GetIBooksStoreBlockEroticaOk returns a tuple with the IBooksStoreBlockErotica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIBooksStoreBlockErotica

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasIBooksStoreBlockErotica() bool`

HasIBooksStoreBlockErotica returns a boolean if a field has been set.

### SetIBooksStoreBlockErotica

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetIBooksStoreBlockErotica(v bool)`

SetIBooksStoreBlockErotica gets a reference to the given bool and assigns it to the IBooksStoreBlockErotica field.

### GetICloudBlockActivityContinuation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockActivityContinuation() bool`

GetICloudBlockActivityContinuation returns the ICloudBlockActivityContinuation field if non-nil, zero value otherwise.

### GetICloudBlockActivityContinuationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockActivityContinuationOk() (bool, bool)`

GetICloudBlockActivityContinuationOk returns a tuple with the ICloudBlockActivityContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockActivityContinuation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockActivityContinuation() bool`

HasICloudBlockActivityContinuation returns a boolean if a field has been set.

### SetICloudBlockActivityContinuation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockActivityContinuation(v bool)`

SetICloudBlockActivityContinuation gets a reference to the given bool and assigns it to the ICloudBlockActivityContinuation field.

### GetICloudBlockBackup

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockBackup() bool`

GetICloudBlockBackup returns the ICloudBlockBackup field if non-nil, zero value otherwise.

### GetICloudBlockBackupOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockBackupOk() (bool, bool)`

GetICloudBlockBackupOk returns a tuple with the ICloudBlockBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockBackup

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockBackup() bool`

HasICloudBlockBackup returns a boolean if a field has been set.

### SetICloudBlockBackup

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockBackup(v bool)`

SetICloudBlockBackup gets a reference to the given bool and assigns it to the ICloudBlockBackup field.

### GetICloudBlockDocumentSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockDocumentSync() bool`

GetICloudBlockDocumentSync returns the ICloudBlockDocumentSync field if non-nil, zero value otherwise.

### GetICloudBlockDocumentSyncOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockDocumentSyncOk() (bool, bool)`

GetICloudBlockDocumentSyncOk returns a tuple with the ICloudBlockDocumentSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockDocumentSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockDocumentSync() bool`

HasICloudBlockDocumentSync returns a boolean if a field has been set.

### SetICloudBlockDocumentSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockDocumentSync(v bool)`

SetICloudBlockDocumentSync gets a reference to the given bool and assigns it to the ICloudBlockDocumentSync field.

### GetICloudBlockManagedAppsSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockManagedAppsSync() bool`

GetICloudBlockManagedAppsSync returns the ICloudBlockManagedAppsSync field if non-nil, zero value otherwise.

### GetICloudBlockManagedAppsSyncOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockManagedAppsSyncOk() (bool, bool)`

GetICloudBlockManagedAppsSyncOk returns a tuple with the ICloudBlockManagedAppsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockManagedAppsSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockManagedAppsSync() bool`

HasICloudBlockManagedAppsSync returns a boolean if a field has been set.

### SetICloudBlockManagedAppsSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockManagedAppsSync(v bool)`

SetICloudBlockManagedAppsSync gets a reference to the given bool and assigns it to the ICloudBlockManagedAppsSync field.

### GetICloudBlockPhotoLibrary

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockPhotoLibrary() bool`

GetICloudBlockPhotoLibrary returns the ICloudBlockPhotoLibrary field if non-nil, zero value otherwise.

### GetICloudBlockPhotoLibraryOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockPhotoLibraryOk() (bool, bool)`

GetICloudBlockPhotoLibraryOk returns a tuple with the ICloudBlockPhotoLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockPhotoLibrary

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockPhotoLibrary() bool`

HasICloudBlockPhotoLibrary returns a boolean if a field has been set.

### SetICloudBlockPhotoLibrary

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockPhotoLibrary(v bool)`

SetICloudBlockPhotoLibrary gets a reference to the given bool and assigns it to the ICloudBlockPhotoLibrary field.

### GetICloudBlockPhotoStreamSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockPhotoStreamSync() bool`

GetICloudBlockPhotoStreamSync returns the ICloudBlockPhotoStreamSync field if non-nil, zero value otherwise.

### GetICloudBlockPhotoStreamSyncOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockPhotoStreamSyncOk() (bool, bool)`

GetICloudBlockPhotoStreamSyncOk returns a tuple with the ICloudBlockPhotoStreamSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockPhotoStreamSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockPhotoStreamSync() bool`

HasICloudBlockPhotoStreamSync returns a boolean if a field has been set.

### SetICloudBlockPhotoStreamSync

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockPhotoStreamSync(v bool)`

SetICloudBlockPhotoStreamSync gets a reference to the given bool and assigns it to the ICloudBlockPhotoStreamSync field.

### GetICloudBlockSharedPhotoStream

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockSharedPhotoStream() bool`

GetICloudBlockSharedPhotoStream returns the ICloudBlockSharedPhotoStream field if non-nil, zero value otherwise.

### GetICloudBlockSharedPhotoStreamOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudBlockSharedPhotoStreamOk() (bool, bool)`

GetICloudBlockSharedPhotoStreamOk returns a tuple with the ICloudBlockSharedPhotoStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudBlockSharedPhotoStream

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudBlockSharedPhotoStream() bool`

HasICloudBlockSharedPhotoStream returns a boolean if a field has been set.

### SetICloudBlockSharedPhotoStream

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudBlockSharedPhotoStream(v bool)`

SetICloudBlockSharedPhotoStream gets a reference to the given bool and assigns it to the ICloudBlockSharedPhotoStream field.

### GetICloudRequireEncryptedBackup

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudRequireEncryptedBackup() bool`

GetICloudRequireEncryptedBackup returns the ICloudRequireEncryptedBackup field if non-nil, zero value otherwise.

### GetICloudRequireEncryptedBackupOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetICloudRequireEncryptedBackupOk() (bool, bool)`

GetICloudRequireEncryptedBackupOk returns a tuple with the ICloudRequireEncryptedBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICloudRequireEncryptedBackup

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasICloudRequireEncryptedBackup() bool`

HasICloudRequireEncryptedBackup returns a boolean if a field has been set.

### SetICloudRequireEncryptedBackup

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetICloudRequireEncryptedBackup(v bool)`

SetICloudRequireEncryptedBackup gets a reference to the given bool and assigns it to the ICloudRequireEncryptedBackup field.

### GetITunesBlockExplicitContent

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetITunesBlockExplicitContent() bool`

GetITunesBlockExplicitContent returns the ITunesBlockExplicitContent field if non-nil, zero value otherwise.

### GetITunesBlockExplicitContentOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetITunesBlockExplicitContentOk() (bool, bool)`

GetITunesBlockExplicitContentOk returns a tuple with the ITunesBlockExplicitContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasITunesBlockExplicitContent

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasITunesBlockExplicitContent() bool`

HasITunesBlockExplicitContent returns a boolean if a field has been set.

### SetITunesBlockExplicitContent

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetITunesBlockExplicitContent(v bool)`

SetITunesBlockExplicitContent gets a reference to the given bool and assigns it to the ITunesBlockExplicitContent field.

### GetITunesBlockMusicService

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetITunesBlockMusicService() bool`

GetITunesBlockMusicService returns the ITunesBlockMusicService field if non-nil, zero value otherwise.

### GetITunesBlockMusicServiceOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetITunesBlockMusicServiceOk() (bool, bool)`

GetITunesBlockMusicServiceOk returns a tuple with the ITunesBlockMusicService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasITunesBlockMusicService

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasITunesBlockMusicService() bool`

HasITunesBlockMusicService returns a boolean if a field has been set.

### SetITunesBlockMusicService

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetITunesBlockMusicService(v bool)`

SetITunesBlockMusicService gets a reference to the given bool and assigns it to the ITunesBlockMusicService field.

### GetITunesBlockRadio

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetITunesBlockRadio() bool`

GetITunesBlockRadio returns the ITunesBlockRadio field if non-nil, zero value otherwise.

### GetITunesBlockRadioOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetITunesBlockRadioOk() (bool, bool)`

GetITunesBlockRadioOk returns a tuple with the ITunesBlockRadio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasITunesBlockRadio

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasITunesBlockRadio() bool`

HasITunesBlockRadio returns a boolean if a field has been set.

### SetITunesBlockRadio

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetITunesBlockRadio(v bool)`

SetITunesBlockRadio gets a reference to the given bool and assigns it to the ITunesBlockRadio field.

### GetKeyboardBlockAutoCorrect

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockAutoCorrect() bool`

GetKeyboardBlockAutoCorrect returns the KeyboardBlockAutoCorrect field if non-nil, zero value otherwise.

### GetKeyboardBlockAutoCorrectOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockAutoCorrectOk() (bool, bool)`

GetKeyboardBlockAutoCorrectOk returns a tuple with the KeyboardBlockAutoCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockAutoCorrect

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKeyboardBlockAutoCorrect() bool`

HasKeyboardBlockAutoCorrect returns a boolean if a field has been set.

### SetKeyboardBlockAutoCorrect

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKeyboardBlockAutoCorrect(v bool)`

SetKeyboardBlockAutoCorrect gets a reference to the given bool and assigns it to the KeyboardBlockAutoCorrect field.

### GetKeyboardBlockDictation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockDictation() bool`

GetKeyboardBlockDictation returns the KeyboardBlockDictation field if non-nil, zero value otherwise.

### GetKeyboardBlockDictationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockDictationOk() (bool, bool)`

GetKeyboardBlockDictationOk returns a tuple with the KeyboardBlockDictation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockDictation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKeyboardBlockDictation() bool`

HasKeyboardBlockDictation returns a boolean if a field has been set.

### SetKeyboardBlockDictation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKeyboardBlockDictation(v bool)`

SetKeyboardBlockDictation gets a reference to the given bool and assigns it to the KeyboardBlockDictation field.

### GetKeyboardBlockPredictive

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockPredictive() bool`

GetKeyboardBlockPredictive returns the KeyboardBlockPredictive field if non-nil, zero value otherwise.

### GetKeyboardBlockPredictiveOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockPredictiveOk() (bool, bool)`

GetKeyboardBlockPredictiveOk returns a tuple with the KeyboardBlockPredictive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockPredictive

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKeyboardBlockPredictive() bool`

HasKeyboardBlockPredictive returns a boolean if a field has been set.

### SetKeyboardBlockPredictive

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKeyboardBlockPredictive(v bool)`

SetKeyboardBlockPredictive gets a reference to the given bool and assigns it to the KeyboardBlockPredictive field.

### GetKeyboardBlockShortcuts

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockShortcuts() bool`

GetKeyboardBlockShortcuts returns the KeyboardBlockShortcuts field if non-nil, zero value otherwise.

### GetKeyboardBlockShortcutsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockShortcutsOk() (bool, bool)`

GetKeyboardBlockShortcutsOk returns a tuple with the KeyboardBlockShortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockShortcuts

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKeyboardBlockShortcuts() bool`

HasKeyboardBlockShortcuts returns a boolean if a field has been set.

### SetKeyboardBlockShortcuts

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKeyboardBlockShortcuts(v bool)`

SetKeyboardBlockShortcuts gets a reference to the given bool and assigns it to the KeyboardBlockShortcuts field.

### GetKeyboardBlockSpellCheck

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockSpellCheck() bool`

GetKeyboardBlockSpellCheck returns the KeyboardBlockSpellCheck field if non-nil, zero value otherwise.

### GetKeyboardBlockSpellCheckOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKeyboardBlockSpellCheckOk() (bool, bool)`

GetKeyboardBlockSpellCheckOk returns a tuple with the KeyboardBlockSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeyboardBlockSpellCheck

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKeyboardBlockSpellCheck() bool`

HasKeyboardBlockSpellCheck returns a boolean if a field has been set.

### SetKeyboardBlockSpellCheck

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKeyboardBlockSpellCheck(v bool)`

SetKeyboardBlockSpellCheck gets a reference to the given bool and assigns it to the KeyboardBlockSpellCheck field.

### GetKioskModeAllowAssistiveSpeak

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveSpeak() bool`

GetKioskModeAllowAssistiveSpeak returns the KioskModeAllowAssistiveSpeak field if non-nil, zero value otherwise.

### GetKioskModeAllowAssistiveSpeakOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveSpeakOk() (bool, bool)`

GetKioskModeAllowAssistiveSpeakOk returns a tuple with the KioskModeAllowAssistiveSpeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowAssistiveSpeak

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowAssistiveSpeak() bool`

HasKioskModeAllowAssistiveSpeak returns a boolean if a field has been set.

### SetKioskModeAllowAssistiveSpeak

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveSpeak(v bool)`

SetKioskModeAllowAssistiveSpeak gets a reference to the given bool and assigns it to the KioskModeAllowAssistiveSpeak field.

### GetKioskModeAllowAssistiveTouchSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveTouchSettings() bool`

GetKioskModeAllowAssistiveTouchSettings returns the KioskModeAllowAssistiveTouchSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowAssistiveTouchSettingsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveTouchSettingsOk() (bool, bool)`

GetKioskModeAllowAssistiveTouchSettingsOk returns a tuple with the KioskModeAllowAssistiveTouchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowAssistiveTouchSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowAssistiveTouchSettings() bool`

HasKioskModeAllowAssistiveTouchSettings returns a boolean if a field has been set.

### SetKioskModeAllowAssistiveTouchSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveTouchSettings(v bool)`

SetKioskModeAllowAssistiveTouchSettings gets a reference to the given bool and assigns it to the KioskModeAllowAssistiveTouchSettings field.

### GetKioskModeAllowAutoLock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowAutoLock() bool`

GetKioskModeAllowAutoLock returns the KioskModeAllowAutoLock field if non-nil, zero value otherwise.

### GetKioskModeAllowAutoLockOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowAutoLockOk() (bool, bool)`

GetKioskModeAllowAutoLockOk returns a tuple with the KioskModeAllowAutoLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowAutoLock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowAutoLock() bool`

HasKioskModeAllowAutoLock returns a boolean if a field has been set.

### SetKioskModeAllowAutoLock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowAutoLock(v bool)`

SetKioskModeAllowAutoLock gets a reference to the given bool and assigns it to the KioskModeAllowAutoLock field.

### GetKioskModeAllowColorInversionSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowColorInversionSettings() bool`

GetKioskModeAllowColorInversionSettings returns the KioskModeAllowColorInversionSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowColorInversionSettingsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowColorInversionSettingsOk() (bool, bool)`

GetKioskModeAllowColorInversionSettingsOk returns a tuple with the KioskModeAllowColorInversionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowColorInversionSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowColorInversionSettings() bool`

HasKioskModeAllowColorInversionSettings returns a boolean if a field has been set.

### SetKioskModeAllowColorInversionSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowColorInversionSettings(v bool)`

SetKioskModeAllowColorInversionSettings gets a reference to the given bool and assigns it to the KioskModeAllowColorInversionSettings field.

### GetKioskModeAllowRingerSwitch

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowRingerSwitch() bool`

GetKioskModeAllowRingerSwitch returns the KioskModeAllowRingerSwitch field if non-nil, zero value otherwise.

### GetKioskModeAllowRingerSwitchOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowRingerSwitchOk() (bool, bool)`

GetKioskModeAllowRingerSwitchOk returns a tuple with the KioskModeAllowRingerSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowRingerSwitch

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowRingerSwitch() bool`

HasKioskModeAllowRingerSwitch returns a boolean if a field has been set.

### SetKioskModeAllowRingerSwitch

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowRingerSwitch(v bool)`

SetKioskModeAllowRingerSwitch gets a reference to the given bool and assigns it to the KioskModeAllowRingerSwitch field.

### GetKioskModeAllowScreenRotation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowScreenRotation() bool`

GetKioskModeAllowScreenRotation returns the KioskModeAllowScreenRotation field if non-nil, zero value otherwise.

### GetKioskModeAllowScreenRotationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowScreenRotationOk() (bool, bool)`

GetKioskModeAllowScreenRotationOk returns a tuple with the KioskModeAllowScreenRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowScreenRotation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowScreenRotation() bool`

HasKioskModeAllowScreenRotation returns a boolean if a field has been set.

### SetKioskModeAllowScreenRotation

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowScreenRotation(v bool)`

SetKioskModeAllowScreenRotation gets a reference to the given bool and assigns it to the KioskModeAllowScreenRotation field.

### GetKioskModeAllowSleepButton

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowSleepButton() bool`

GetKioskModeAllowSleepButton returns the KioskModeAllowSleepButton field if non-nil, zero value otherwise.

### GetKioskModeAllowSleepButtonOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowSleepButtonOk() (bool, bool)`

GetKioskModeAllowSleepButtonOk returns a tuple with the KioskModeAllowSleepButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowSleepButton

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowSleepButton() bool`

HasKioskModeAllowSleepButton returns a boolean if a field has been set.

### SetKioskModeAllowSleepButton

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowSleepButton(v bool)`

SetKioskModeAllowSleepButton gets a reference to the given bool and assigns it to the KioskModeAllowSleepButton field.

### GetKioskModeAllowTouchscreen

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowTouchscreen() bool`

GetKioskModeAllowTouchscreen returns the KioskModeAllowTouchscreen field if non-nil, zero value otherwise.

### GetKioskModeAllowTouchscreenOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowTouchscreenOk() (bool, bool)`

GetKioskModeAllowTouchscreenOk returns a tuple with the KioskModeAllowTouchscreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowTouchscreen

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowTouchscreen() bool`

HasKioskModeAllowTouchscreen returns a boolean if a field has been set.

### SetKioskModeAllowTouchscreen

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowTouchscreen(v bool)`

SetKioskModeAllowTouchscreen gets a reference to the given bool and assigns it to the KioskModeAllowTouchscreen field.

### GetKioskModeAllowVoiceOverSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowVoiceOverSettings() bool`

GetKioskModeAllowVoiceOverSettings returns the KioskModeAllowVoiceOverSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowVoiceOverSettingsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowVoiceOverSettingsOk() (bool, bool)`

GetKioskModeAllowVoiceOverSettingsOk returns a tuple with the KioskModeAllowVoiceOverSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowVoiceOverSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowVoiceOverSettings() bool`

HasKioskModeAllowVoiceOverSettings returns a boolean if a field has been set.

### SetKioskModeAllowVoiceOverSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowVoiceOverSettings(v bool)`

SetKioskModeAllowVoiceOverSettings gets a reference to the given bool and assigns it to the KioskModeAllowVoiceOverSettings field.

### GetKioskModeAllowVolumeButtons

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowVolumeButtons() bool`

GetKioskModeAllowVolumeButtons returns the KioskModeAllowVolumeButtons field if non-nil, zero value otherwise.

### GetKioskModeAllowVolumeButtonsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowVolumeButtonsOk() (bool, bool)`

GetKioskModeAllowVolumeButtonsOk returns a tuple with the KioskModeAllowVolumeButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowVolumeButtons

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowVolumeButtons() bool`

HasKioskModeAllowVolumeButtons returns a boolean if a field has been set.

### SetKioskModeAllowVolumeButtons

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowVolumeButtons(v bool)`

SetKioskModeAllowVolumeButtons gets a reference to the given bool and assigns it to the KioskModeAllowVolumeButtons field.

### GetKioskModeAllowZoomSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowZoomSettings() bool`

GetKioskModeAllowZoomSettings returns the KioskModeAllowZoomSettings field if non-nil, zero value otherwise.

### GetKioskModeAllowZoomSettingsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAllowZoomSettingsOk() (bool, bool)`

GetKioskModeAllowZoomSettingsOk returns a tuple with the KioskModeAllowZoomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAllowZoomSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAllowZoomSettings() bool`

HasKioskModeAllowZoomSettings returns a boolean if a field has been set.

### SetKioskModeAllowZoomSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAllowZoomSettings(v bool)`

SetKioskModeAllowZoomSettings gets a reference to the given bool and assigns it to the KioskModeAllowZoomSettings field.

### GetKioskModeAppStoreUrl

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAppStoreUrl() string`

GetKioskModeAppStoreUrl returns the KioskModeAppStoreUrl field if non-nil, zero value otherwise.

### GetKioskModeAppStoreUrlOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeAppStoreUrlOk() (string, bool)`

GetKioskModeAppStoreUrlOk returns a tuple with the KioskModeAppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeAppStoreUrl

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeAppStoreUrl() bool`

HasKioskModeAppStoreUrl returns a boolean if a field has been set.

### SetKioskModeAppStoreUrl

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAppStoreUrl(v string)`

SetKioskModeAppStoreUrl gets a reference to the given string and assigns it to the KioskModeAppStoreUrl field.

### SetKioskModeAppStoreUrlExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeAppStoreUrlExplicitNull(b bool)`

SetKioskModeAppStoreUrlExplicitNull (un)sets KioskModeAppStoreUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskModeAppStoreUrl value is set to nil even if false is passed
### GetKioskModeBuiltInAppId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeBuiltInAppId() string`

GetKioskModeBuiltInAppId returns the KioskModeBuiltInAppId field if non-nil, zero value otherwise.

### GetKioskModeBuiltInAppIdOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeBuiltInAppIdOk() (string, bool)`

GetKioskModeBuiltInAppIdOk returns a tuple with the KioskModeBuiltInAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeBuiltInAppId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeBuiltInAppId() bool`

HasKioskModeBuiltInAppId returns a boolean if a field has been set.

### SetKioskModeBuiltInAppId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeBuiltInAppId(v string)`

SetKioskModeBuiltInAppId gets a reference to the given string and assigns it to the KioskModeBuiltInAppId field.

### SetKioskModeBuiltInAppIdExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeBuiltInAppIdExplicitNull(b bool)`

SetKioskModeBuiltInAppIdExplicitNull (un)sets KioskModeBuiltInAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskModeBuiltInAppId value is set to nil even if false is passed
### GetKioskModeRequireAssistiveTouch

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireAssistiveTouch() bool`

GetKioskModeRequireAssistiveTouch returns the KioskModeRequireAssistiveTouch field if non-nil, zero value otherwise.

### GetKioskModeRequireAssistiveTouchOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireAssistiveTouchOk() (bool, bool)`

GetKioskModeRequireAssistiveTouchOk returns a tuple with the KioskModeRequireAssistiveTouch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireAssistiveTouch

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeRequireAssistiveTouch() bool`

HasKioskModeRequireAssistiveTouch returns a boolean if a field has been set.

### SetKioskModeRequireAssistiveTouch

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeRequireAssistiveTouch(v bool)`

SetKioskModeRequireAssistiveTouch gets a reference to the given bool and assigns it to the KioskModeRequireAssistiveTouch field.

### GetKioskModeRequireColorInversion

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireColorInversion() bool`

GetKioskModeRequireColorInversion returns the KioskModeRequireColorInversion field if non-nil, zero value otherwise.

### GetKioskModeRequireColorInversionOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireColorInversionOk() (bool, bool)`

GetKioskModeRequireColorInversionOk returns a tuple with the KioskModeRequireColorInversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireColorInversion

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeRequireColorInversion() bool`

HasKioskModeRequireColorInversion returns a boolean if a field has been set.

### SetKioskModeRequireColorInversion

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeRequireColorInversion(v bool)`

SetKioskModeRequireColorInversion gets a reference to the given bool and assigns it to the KioskModeRequireColorInversion field.

### GetKioskModeRequireMonoAudio

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireMonoAudio() bool`

GetKioskModeRequireMonoAudio returns the KioskModeRequireMonoAudio field if non-nil, zero value otherwise.

### GetKioskModeRequireMonoAudioOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireMonoAudioOk() (bool, bool)`

GetKioskModeRequireMonoAudioOk returns a tuple with the KioskModeRequireMonoAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireMonoAudio

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeRequireMonoAudio() bool`

HasKioskModeRequireMonoAudio returns a boolean if a field has been set.

### SetKioskModeRequireMonoAudio

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeRequireMonoAudio(v bool)`

SetKioskModeRequireMonoAudio gets a reference to the given bool and assigns it to the KioskModeRequireMonoAudio field.

### GetKioskModeRequireVoiceOver

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireVoiceOver() bool`

GetKioskModeRequireVoiceOver returns the KioskModeRequireVoiceOver field if non-nil, zero value otherwise.

### GetKioskModeRequireVoiceOverOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireVoiceOverOk() (bool, bool)`

GetKioskModeRequireVoiceOverOk returns a tuple with the KioskModeRequireVoiceOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireVoiceOver

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeRequireVoiceOver() bool`

HasKioskModeRequireVoiceOver returns a boolean if a field has been set.

### SetKioskModeRequireVoiceOver

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeRequireVoiceOver(v bool)`

SetKioskModeRequireVoiceOver gets a reference to the given bool and assigns it to the KioskModeRequireVoiceOver field.

### GetKioskModeRequireZoom

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireZoom() bool`

GetKioskModeRequireZoom returns the KioskModeRequireZoom field if non-nil, zero value otherwise.

### GetKioskModeRequireZoomOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeRequireZoomOk() (bool, bool)`

GetKioskModeRequireZoomOk returns a tuple with the KioskModeRequireZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeRequireZoom

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeRequireZoom() bool`

HasKioskModeRequireZoom returns a boolean if a field has been set.

### SetKioskModeRequireZoom

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeRequireZoom(v bool)`

SetKioskModeRequireZoom gets a reference to the given bool and assigns it to the KioskModeRequireZoom field.

### GetKioskModeManagedAppId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeManagedAppId() string`

GetKioskModeManagedAppId returns the KioskModeManagedAppId field if non-nil, zero value otherwise.

### GetKioskModeManagedAppIdOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetKioskModeManagedAppIdOk() (string, bool)`

GetKioskModeManagedAppIdOk returns a tuple with the KioskModeManagedAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskModeManagedAppId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasKioskModeManagedAppId() bool`

HasKioskModeManagedAppId returns a boolean if a field has been set.

### SetKioskModeManagedAppId

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeManagedAppId(v string)`

SetKioskModeManagedAppId gets a reference to the given string and assigns it to the KioskModeManagedAppId field.

### SetKioskModeManagedAppIdExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetKioskModeManagedAppIdExplicitNull(b bool)`

SetKioskModeManagedAppIdExplicitNull (un)sets KioskModeManagedAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskModeManagedAppId value is set to nil even if false is passed
### GetLockScreenBlockControlCenter

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockControlCenter() bool`

GetLockScreenBlockControlCenter returns the LockScreenBlockControlCenter field if non-nil, zero value otherwise.

### GetLockScreenBlockControlCenterOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockControlCenterOk() (bool, bool)`

GetLockScreenBlockControlCenterOk returns a tuple with the LockScreenBlockControlCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockControlCenter

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasLockScreenBlockControlCenter() bool`

HasLockScreenBlockControlCenter returns a boolean if a field has been set.

### SetLockScreenBlockControlCenter

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetLockScreenBlockControlCenter(v bool)`

SetLockScreenBlockControlCenter gets a reference to the given bool and assigns it to the LockScreenBlockControlCenter field.

### GetLockScreenBlockNotificationView

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockNotificationView() bool`

GetLockScreenBlockNotificationView returns the LockScreenBlockNotificationView field if non-nil, zero value otherwise.

### GetLockScreenBlockNotificationViewOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockNotificationViewOk() (bool, bool)`

GetLockScreenBlockNotificationViewOk returns a tuple with the LockScreenBlockNotificationView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockNotificationView

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasLockScreenBlockNotificationView() bool`

HasLockScreenBlockNotificationView returns a boolean if a field has been set.

### SetLockScreenBlockNotificationView

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetLockScreenBlockNotificationView(v bool)`

SetLockScreenBlockNotificationView gets a reference to the given bool and assigns it to the LockScreenBlockNotificationView field.

### GetLockScreenBlockPassbook

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockPassbook() bool`

GetLockScreenBlockPassbook returns the LockScreenBlockPassbook field if non-nil, zero value otherwise.

### GetLockScreenBlockPassbookOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockPassbookOk() (bool, bool)`

GetLockScreenBlockPassbookOk returns a tuple with the LockScreenBlockPassbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockPassbook

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasLockScreenBlockPassbook() bool`

HasLockScreenBlockPassbook returns a boolean if a field has been set.

### SetLockScreenBlockPassbook

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetLockScreenBlockPassbook(v bool)`

SetLockScreenBlockPassbook gets a reference to the given bool and assigns it to the LockScreenBlockPassbook field.

### GetLockScreenBlockTodayView

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockTodayView() bool`

GetLockScreenBlockTodayView returns the LockScreenBlockTodayView field if non-nil, zero value otherwise.

### GetLockScreenBlockTodayViewOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetLockScreenBlockTodayViewOk() (bool, bool)`

GetLockScreenBlockTodayViewOk returns a tuple with the LockScreenBlockTodayView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockTodayView

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasLockScreenBlockTodayView() bool`

HasLockScreenBlockTodayView returns a boolean if a field has been set.

### SetLockScreenBlockTodayView

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetLockScreenBlockTodayView(v bool)`

SetLockScreenBlockTodayView gets a reference to the given bool and assigns it to the LockScreenBlockTodayView field.

### GetMediaContentRatingAustralia

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingAustralia() AnyOfmicrosoftGraphMediaContentRatingAustralia`

GetMediaContentRatingAustralia returns the MediaContentRatingAustralia field if non-nil, zero value otherwise.

### GetMediaContentRatingAustraliaOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingAustraliaOk() (AnyOfmicrosoftGraphMediaContentRatingAustralia, bool)`

GetMediaContentRatingAustraliaOk returns a tuple with the MediaContentRatingAustralia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingAustralia

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingAustralia() bool`

HasMediaContentRatingAustralia returns a boolean if a field has been set.

### SetMediaContentRatingAustralia

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingAustralia(v AnyOfmicrosoftGraphMediaContentRatingAustralia)`

SetMediaContentRatingAustralia gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingAustralia and assigns it to the MediaContentRatingAustralia field.

### SetMediaContentRatingAustraliaExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingAustraliaExplicitNull(b bool)`

SetMediaContentRatingAustraliaExplicitNull (un)sets MediaContentRatingAustralia to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingAustralia value is set to nil even if false is passed
### GetMediaContentRatingCanada

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingCanada() AnyOfmicrosoftGraphMediaContentRatingCanada`

GetMediaContentRatingCanada returns the MediaContentRatingCanada field if non-nil, zero value otherwise.

### GetMediaContentRatingCanadaOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingCanadaOk() (AnyOfmicrosoftGraphMediaContentRatingCanada, bool)`

GetMediaContentRatingCanadaOk returns a tuple with the MediaContentRatingCanada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingCanada

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingCanada() bool`

HasMediaContentRatingCanada returns a boolean if a field has been set.

### SetMediaContentRatingCanada

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingCanada(v AnyOfmicrosoftGraphMediaContentRatingCanada)`

SetMediaContentRatingCanada gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingCanada and assigns it to the MediaContentRatingCanada field.

### SetMediaContentRatingCanadaExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingCanadaExplicitNull(b bool)`

SetMediaContentRatingCanadaExplicitNull (un)sets MediaContentRatingCanada to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingCanada value is set to nil even if false is passed
### GetMediaContentRatingFrance

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingFrance() AnyOfmicrosoftGraphMediaContentRatingFrance`

GetMediaContentRatingFrance returns the MediaContentRatingFrance field if non-nil, zero value otherwise.

### GetMediaContentRatingFranceOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingFranceOk() (AnyOfmicrosoftGraphMediaContentRatingFrance, bool)`

GetMediaContentRatingFranceOk returns a tuple with the MediaContentRatingFrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingFrance

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingFrance() bool`

HasMediaContentRatingFrance returns a boolean if a field has been set.

### SetMediaContentRatingFrance

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingFrance(v AnyOfmicrosoftGraphMediaContentRatingFrance)`

SetMediaContentRatingFrance gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingFrance and assigns it to the MediaContentRatingFrance field.

### SetMediaContentRatingFranceExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingFranceExplicitNull(b bool)`

SetMediaContentRatingFranceExplicitNull (un)sets MediaContentRatingFrance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingFrance value is set to nil even if false is passed
### GetMediaContentRatingGermany

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingGermany() AnyOfmicrosoftGraphMediaContentRatingGermany`

GetMediaContentRatingGermany returns the MediaContentRatingGermany field if non-nil, zero value otherwise.

### GetMediaContentRatingGermanyOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingGermanyOk() (AnyOfmicrosoftGraphMediaContentRatingGermany, bool)`

GetMediaContentRatingGermanyOk returns a tuple with the MediaContentRatingGermany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingGermany

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingGermany() bool`

HasMediaContentRatingGermany returns a boolean if a field has been set.

### SetMediaContentRatingGermany

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingGermany(v AnyOfmicrosoftGraphMediaContentRatingGermany)`

SetMediaContentRatingGermany gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingGermany and assigns it to the MediaContentRatingGermany field.

### SetMediaContentRatingGermanyExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingGermanyExplicitNull(b bool)`

SetMediaContentRatingGermanyExplicitNull (un)sets MediaContentRatingGermany to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingGermany value is set to nil even if false is passed
### GetMediaContentRatingIreland

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingIreland() AnyOfmicrosoftGraphMediaContentRatingIreland`

GetMediaContentRatingIreland returns the MediaContentRatingIreland field if non-nil, zero value otherwise.

### GetMediaContentRatingIrelandOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingIrelandOk() (AnyOfmicrosoftGraphMediaContentRatingIreland, bool)`

GetMediaContentRatingIrelandOk returns a tuple with the MediaContentRatingIreland field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingIreland

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingIreland() bool`

HasMediaContentRatingIreland returns a boolean if a field has been set.

### SetMediaContentRatingIreland

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingIreland(v AnyOfmicrosoftGraphMediaContentRatingIreland)`

SetMediaContentRatingIreland gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingIreland and assigns it to the MediaContentRatingIreland field.

### SetMediaContentRatingIrelandExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingIrelandExplicitNull(b bool)`

SetMediaContentRatingIrelandExplicitNull (un)sets MediaContentRatingIreland to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingIreland value is set to nil even if false is passed
### GetMediaContentRatingJapan

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingJapan() AnyOfmicrosoftGraphMediaContentRatingJapan`

GetMediaContentRatingJapan returns the MediaContentRatingJapan field if non-nil, zero value otherwise.

### GetMediaContentRatingJapanOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingJapanOk() (AnyOfmicrosoftGraphMediaContentRatingJapan, bool)`

GetMediaContentRatingJapanOk returns a tuple with the MediaContentRatingJapan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingJapan

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingJapan() bool`

HasMediaContentRatingJapan returns a boolean if a field has been set.

### SetMediaContentRatingJapan

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingJapan(v AnyOfmicrosoftGraphMediaContentRatingJapan)`

SetMediaContentRatingJapan gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingJapan and assigns it to the MediaContentRatingJapan field.

### SetMediaContentRatingJapanExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingJapanExplicitNull(b bool)`

SetMediaContentRatingJapanExplicitNull (un)sets MediaContentRatingJapan to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingJapan value is set to nil even if false is passed
### GetMediaContentRatingNewZealand

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingNewZealand() AnyOfmicrosoftGraphMediaContentRatingNewZealand`

GetMediaContentRatingNewZealand returns the MediaContentRatingNewZealand field if non-nil, zero value otherwise.

### GetMediaContentRatingNewZealandOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingNewZealandOk() (AnyOfmicrosoftGraphMediaContentRatingNewZealand, bool)`

GetMediaContentRatingNewZealandOk returns a tuple with the MediaContentRatingNewZealand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingNewZealand

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingNewZealand() bool`

HasMediaContentRatingNewZealand returns a boolean if a field has been set.

### SetMediaContentRatingNewZealand

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingNewZealand(v AnyOfmicrosoftGraphMediaContentRatingNewZealand)`

SetMediaContentRatingNewZealand gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingNewZealand and assigns it to the MediaContentRatingNewZealand field.

### SetMediaContentRatingNewZealandExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingNewZealandExplicitNull(b bool)`

SetMediaContentRatingNewZealandExplicitNull (un)sets MediaContentRatingNewZealand to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingNewZealand value is set to nil even if false is passed
### GetMediaContentRatingUnitedKingdom

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingUnitedKingdom() AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom`

GetMediaContentRatingUnitedKingdom returns the MediaContentRatingUnitedKingdom field if non-nil, zero value otherwise.

### GetMediaContentRatingUnitedKingdomOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingUnitedKingdomOk() (AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom, bool)`

GetMediaContentRatingUnitedKingdomOk returns a tuple with the MediaContentRatingUnitedKingdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingUnitedKingdom

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingUnitedKingdom() bool`

HasMediaContentRatingUnitedKingdom returns a boolean if a field has been set.

### SetMediaContentRatingUnitedKingdom

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingUnitedKingdom(v AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom)`

SetMediaContentRatingUnitedKingdom gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingUnitedKingdom and assigns it to the MediaContentRatingUnitedKingdom field.

### SetMediaContentRatingUnitedKingdomExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingUnitedKingdomExplicitNull(b bool)`

SetMediaContentRatingUnitedKingdomExplicitNull (un)sets MediaContentRatingUnitedKingdom to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingUnitedKingdom value is set to nil even if false is passed
### GetMediaContentRatingUnitedStates

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingUnitedStates() AnyOfmicrosoftGraphMediaContentRatingUnitedStates`

GetMediaContentRatingUnitedStates returns the MediaContentRatingUnitedStates field if non-nil, zero value otherwise.

### GetMediaContentRatingUnitedStatesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingUnitedStatesOk() (AnyOfmicrosoftGraphMediaContentRatingUnitedStates, bool)`

GetMediaContentRatingUnitedStatesOk returns a tuple with the MediaContentRatingUnitedStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingUnitedStates

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingUnitedStates() bool`

HasMediaContentRatingUnitedStates returns a boolean if a field has been set.

### SetMediaContentRatingUnitedStates

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingUnitedStates(v AnyOfmicrosoftGraphMediaContentRatingUnitedStates)`

SetMediaContentRatingUnitedStates gets a reference to the given AnyOfmicrosoftGraphMediaContentRatingUnitedStates and assigns it to the MediaContentRatingUnitedStates field.

### SetMediaContentRatingUnitedStatesExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingUnitedStatesExplicitNull(b bool)`

SetMediaContentRatingUnitedStatesExplicitNull (un)sets MediaContentRatingUnitedStates to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MediaContentRatingUnitedStates value is set to nil even if false is passed
### GetNetworkUsageRules

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetNetworkUsageRules() []AnyOfmicrosoftGraphIosNetworkUsageRule`

GetNetworkUsageRules returns the NetworkUsageRules field if non-nil, zero value otherwise.

### GetNetworkUsageRulesOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetNetworkUsageRulesOk() ([]AnyOfmicrosoftGraphIosNetworkUsageRule, bool)`

GetNetworkUsageRulesOk returns a tuple with the NetworkUsageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkUsageRules

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasNetworkUsageRules() bool`

HasNetworkUsageRules returns a boolean if a field has been set.

### SetNetworkUsageRules

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetNetworkUsageRules(v []AnyOfmicrosoftGraphIosNetworkUsageRule)`

SetNetworkUsageRules gets a reference to the given []AnyOfmicrosoftGraphIosNetworkUsageRule and assigns it to the NetworkUsageRules field.

### GetMediaContentRatingApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingApps() AnyOfmicrosoftGraphRatingAppsType`

GetMediaContentRatingApps returns the MediaContentRatingApps field if non-nil, zero value otherwise.

### GetMediaContentRatingAppsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMediaContentRatingAppsOk() (AnyOfmicrosoftGraphRatingAppsType, bool)`

GetMediaContentRatingAppsOk returns a tuple with the MediaContentRatingApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediaContentRatingApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMediaContentRatingApps() bool`

HasMediaContentRatingApps returns a boolean if a field has been set.

### SetMediaContentRatingApps

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMediaContentRatingApps(v AnyOfmicrosoftGraphRatingAppsType)`

SetMediaContentRatingApps gets a reference to the given AnyOfmicrosoftGraphRatingAppsType and assigns it to the MediaContentRatingApps field.

### GetMessagesBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMessagesBlocked() bool`

GetMessagesBlocked returns the MessagesBlocked field if non-nil, zero value otherwise.

### GetMessagesBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetMessagesBlockedOk() (bool, bool)`

GetMessagesBlockedOk returns a tuple with the MessagesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessagesBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasMessagesBlocked() bool`

HasMessagesBlocked returns a boolean if a field has been set.

### SetMessagesBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetMessagesBlocked(v bool)`

SetMessagesBlocked gets a reference to the given bool and assigns it to the MessagesBlocked field.

### GetNotificationsBlockSettingsModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetNotificationsBlockSettingsModification() bool`

GetNotificationsBlockSettingsModification returns the NotificationsBlockSettingsModification field if non-nil, zero value otherwise.

### GetNotificationsBlockSettingsModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetNotificationsBlockSettingsModificationOk() (bool, bool)`

GetNotificationsBlockSettingsModificationOk returns a tuple with the NotificationsBlockSettingsModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationsBlockSettingsModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasNotificationsBlockSettingsModification() bool`

HasNotificationsBlockSettingsModification returns a boolean if a field has been set.

### SetNotificationsBlockSettingsModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetNotificationsBlockSettingsModification(v bool)`

SetNotificationsBlockSettingsModification gets a reference to the given bool and assigns it to the NotificationsBlockSettingsModification field.

### GetPasscodeBlockFingerprintUnlock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintUnlock() bool`

GetPasscodeBlockFingerprintUnlock returns the PasscodeBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetPasscodeBlockFingerprintUnlockOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintUnlockOk() (bool, bool)`

GetPasscodeBlockFingerprintUnlockOk returns a tuple with the PasscodeBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockFingerprintUnlock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeBlockFingerprintUnlock() bool`

HasPasscodeBlockFingerprintUnlock returns a boolean if a field has been set.

### SetPasscodeBlockFingerprintUnlock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintUnlock(v bool)`

SetPasscodeBlockFingerprintUnlock gets a reference to the given bool and assigns it to the PasscodeBlockFingerprintUnlock field.

### GetPasscodeBlockFingerprintModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintModification() bool`

GetPasscodeBlockFingerprintModification returns the PasscodeBlockFingerprintModification field if non-nil, zero value otherwise.

### GetPasscodeBlockFingerprintModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintModificationOk() (bool, bool)`

GetPasscodeBlockFingerprintModificationOk returns a tuple with the PasscodeBlockFingerprintModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockFingerprintModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeBlockFingerprintModification() bool`

HasPasscodeBlockFingerprintModification returns a boolean if a field has been set.

### SetPasscodeBlockFingerprintModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintModification(v bool)`

SetPasscodeBlockFingerprintModification gets a reference to the given bool and assigns it to the PasscodeBlockFingerprintModification field.

### GetPasscodeBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockModification() bool`

GetPasscodeBlockModification returns the PasscodeBlockModification field if non-nil, zero value otherwise.

### GetPasscodeBlockModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockModificationOk() (bool, bool)`

GetPasscodeBlockModificationOk returns a tuple with the PasscodeBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeBlockModification() bool`

HasPasscodeBlockModification returns a boolean if a field has been set.

### SetPasscodeBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeBlockModification(v bool)`

SetPasscodeBlockModification gets a reference to the given bool and assigns it to the PasscodeBlockModification field.

### GetPasscodeBlockSimple

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockSimple() bool`

GetPasscodeBlockSimple returns the PasscodeBlockSimple field if non-nil, zero value otherwise.

### GetPasscodeBlockSimpleOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeBlockSimpleOk() (bool, bool)`

GetPasscodeBlockSimpleOk returns a tuple with the PasscodeBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockSimple

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeBlockSimple() bool`

HasPasscodeBlockSimple returns a boolean if a field has been set.

### SetPasscodeBlockSimple

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeBlockSimple(v bool)`

SetPasscodeBlockSimple gets a reference to the given bool and assigns it to the PasscodeBlockSimple field.

### GetPasscodeExpirationDays

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeExpirationDays() int32`

GetPasscodeExpirationDays returns the PasscodeExpirationDays field if non-nil, zero value otherwise.

### GetPasscodeExpirationDaysOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeExpirationDaysOk() (int32, bool)`

GetPasscodeExpirationDaysOk returns a tuple with the PasscodeExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeExpirationDays

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeExpirationDays() bool`

HasPasscodeExpirationDays returns a boolean if a field has been set.

### SetPasscodeExpirationDays

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeExpirationDays(v int32)`

SetPasscodeExpirationDays gets a reference to the given int32 and assigns it to the PasscodeExpirationDays field.

### SetPasscodeExpirationDaysExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeExpirationDaysExplicitNull(b bool)`

SetPasscodeExpirationDaysExplicitNull (un)sets PasscodeExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeExpirationDays value is set to nil even if false is passed
### GetPasscodeMinimumLength

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinimumLength() int32`

GetPasscodeMinimumLength returns the PasscodeMinimumLength field if non-nil, zero value otherwise.

### GetPasscodeMinimumLengthOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinimumLengthOk() (int32, bool)`

GetPasscodeMinimumLengthOk returns a tuple with the PasscodeMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumLength

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeMinimumLength() bool`

HasPasscodeMinimumLength returns a boolean if a field has been set.

### SetPasscodeMinimumLength

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinimumLength(v int32)`

SetPasscodeMinimumLength gets a reference to the given int32 and assigns it to the PasscodeMinimumLength field.

### SetPasscodeMinimumLengthExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinimumLengthExplicitNull(b bool)`

SetPasscodeMinimumLengthExplicitNull (un)sets PasscodeMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumLength value is set to nil even if false is passed
### GetPasscodeMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeLock() int32`

GetPasscodeMinutesOfInactivityBeforeLock returns the PasscodeMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasscodeMinutesOfInactivityBeforeLockOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasscodeMinutesOfInactivityBeforeLockOk returns a tuple with the PasscodeMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeMinutesOfInactivityBeforeLock() bool`

HasPasscodeMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasscodeMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeLock(v int32)`

SetPasscodeMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeLock field.

### SetPasscodeMinutesOfInactivityBeforeLockExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasscodeMinutesOfInactivityBeforeLockExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasscodeMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasscodeMinutesOfInactivityBeforeScreenTimeout returns the PasscodeMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasscodeMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasscodeMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasscodeMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasscodeMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasscodeMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasscodeMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeScreenTimeout field.

### SetPasscodeMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasscodeMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasscodeMinimumCharacterSetCount

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinimumCharacterSetCount() int32`

GetPasscodeMinimumCharacterSetCount returns the PasscodeMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasscodeMinimumCharacterSetCountOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeMinimumCharacterSetCountOk() (int32, bool)`

GetPasscodeMinimumCharacterSetCountOk returns a tuple with the PasscodeMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumCharacterSetCount

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeMinimumCharacterSetCount() bool`

HasPasscodeMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasscodeMinimumCharacterSetCount

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinimumCharacterSetCount(v int32)`

SetPasscodeMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasscodeMinimumCharacterSetCount field.

### SetPasscodeMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeMinimumCharacterSetCountExplicitNull(b bool)`

SetPasscodeMinimumCharacterSetCountExplicitNull (un)sets PasscodeMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasscodePreviousPasscodeBlockCount

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodePreviousPasscodeBlockCount() int32`

GetPasscodePreviousPasscodeBlockCount returns the PasscodePreviousPasscodeBlockCount field if non-nil, zero value otherwise.

### GetPasscodePreviousPasscodeBlockCountOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodePreviousPasscodeBlockCountOk() (int32, bool)`

GetPasscodePreviousPasscodeBlockCountOk returns a tuple with the PasscodePreviousPasscodeBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodePreviousPasscodeBlockCount

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodePreviousPasscodeBlockCount() bool`

HasPasscodePreviousPasscodeBlockCount returns a boolean if a field has been set.

### SetPasscodePreviousPasscodeBlockCount

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodePreviousPasscodeBlockCount(v int32)`

SetPasscodePreviousPasscodeBlockCount gets a reference to the given int32 and assigns it to the PasscodePreviousPasscodeBlockCount field.

### SetPasscodePreviousPasscodeBlockCountExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodePreviousPasscodeBlockCountExplicitNull(b bool)`

SetPasscodePreviousPasscodeBlockCountExplicitNull (un)sets PasscodePreviousPasscodeBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodePreviousPasscodeBlockCount value is set to nil even if false is passed
### GetPasscodeSignInFailureCountBeforeWipe

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeSignInFailureCountBeforeWipe() int32`

GetPasscodeSignInFailureCountBeforeWipe returns the PasscodeSignInFailureCountBeforeWipe field if non-nil, zero value otherwise.

### GetPasscodeSignInFailureCountBeforeWipeOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeSignInFailureCountBeforeWipeOk() (int32, bool)`

GetPasscodeSignInFailureCountBeforeWipeOk returns a tuple with the PasscodeSignInFailureCountBeforeWipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeSignInFailureCountBeforeWipe

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeSignInFailureCountBeforeWipe() bool`

HasPasscodeSignInFailureCountBeforeWipe returns a boolean if a field has been set.

### SetPasscodeSignInFailureCountBeforeWipe

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeSignInFailureCountBeforeWipe(v int32)`

SetPasscodeSignInFailureCountBeforeWipe gets a reference to the given int32 and assigns it to the PasscodeSignInFailureCountBeforeWipe field.

### SetPasscodeSignInFailureCountBeforeWipeExplicitNull

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeSignInFailureCountBeforeWipeExplicitNull(b bool)`

SetPasscodeSignInFailureCountBeforeWipeExplicitNull (un)sets PasscodeSignInFailureCountBeforeWipe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeSignInFailureCountBeforeWipe value is set to nil even if false is passed
### GetPasscodeRequiredType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasscodeRequiredType returns the PasscodeRequiredType field if non-nil, zero value otherwise.

### GetPasscodeRequiredTypeOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasscodeRequiredTypeOk returns a tuple with the PasscodeRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequiredType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeRequiredType() bool`

HasPasscodeRequiredType returns a boolean if a field has been set.

### SetPasscodeRequiredType

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasscodeRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasscodeRequiredType field.

### GetPasscodeRequired

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeRequired() bool`

GetPasscodeRequired returns the PasscodeRequired field if non-nil, zero value otherwise.

### GetPasscodeRequiredOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPasscodeRequiredOk() (bool, bool)`

GetPasscodeRequiredOk returns a tuple with the PasscodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequired

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPasscodeRequired() bool`

HasPasscodeRequired returns a boolean if a field has been set.

### SetPasscodeRequired

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPasscodeRequired(v bool)`

SetPasscodeRequired gets a reference to the given bool and assigns it to the PasscodeRequired field.

### GetPodcastsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPodcastsBlocked() bool`

GetPodcastsBlocked returns the PodcastsBlocked field if non-nil, zero value otherwise.

### GetPodcastsBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetPodcastsBlockedOk() (bool, bool)`

GetPodcastsBlockedOk returns a tuple with the PodcastsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPodcastsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasPodcastsBlocked() bool`

HasPodcastsBlocked returns a boolean if a field has been set.

### SetPodcastsBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetPodcastsBlocked(v bool)`

SetPodcastsBlocked gets a reference to the given bool and assigns it to the PodcastsBlocked field.

### GetSafariBlockAutofill

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockAutofill() bool`

GetSafariBlockAutofill returns the SafariBlockAutofill field if non-nil, zero value otherwise.

### GetSafariBlockAutofillOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockAutofillOk() (bool, bool)`

GetSafariBlockAutofillOk returns a tuple with the SafariBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlockAutofill

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariBlockAutofill() bool`

HasSafariBlockAutofill returns a boolean if a field has been set.

### SetSafariBlockAutofill

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariBlockAutofill(v bool)`

SetSafariBlockAutofill gets a reference to the given bool and assigns it to the SafariBlockAutofill field.

### GetSafariBlockJavaScript

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockJavaScript() bool`

GetSafariBlockJavaScript returns the SafariBlockJavaScript field if non-nil, zero value otherwise.

### GetSafariBlockJavaScriptOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockJavaScriptOk() (bool, bool)`

GetSafariBlockJavaScriptOk returns a tuple with the SafariBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlockJavaScript

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariBlockJavaScript() bool`

HasSafariBlockJavaScript returns a boolean if a field has been set.

### SetSafariBlockJavaScript

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariBlockJavaScript(v bool)`

SetSafariBlockJavaScript gets a reference to the given bool and assigns it to the SafariBlockJavaScript field.

### GetSafariBlockPopups

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockPopups() bool`

GetSafariBlockPopups returns the SafariBlockPopups field if non-nil, zero value otherwise.

### GetSafariBlockPopupsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockPopupsOk() (bool, bool)`

GetSafariBlockPopupsOk returns a tuple with the SafariBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlockPopups

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariBlockPopups() bool`

HasSafariBlockPopups returns a boolean if a field has been set.

### SetSafariBlockPopups

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariBlockPopups(v bool)`

SetSafariBlockPopups gets a reference to the given bool and assigns it to the SafariBlockPopups field.

### GetSafariBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlocked() bool`

GetSafariBlocked returns the SafariBlocked field if non-nil, zero value otherwise.

### GetSafariBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariBlockedOk() (bool, bool)`

GetSafariBlockedOk returns a tuple with the SafariBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariBlocked() bool`

HasSafariBlocked returns a boolean if a field has been set.

### SetSafariBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariBlocked(v bool)`

SetSafariBlocked gets a reference to the given bool and assigns it to the SafariBlocked field.

### GetSafariCookieSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariCookieSettings() AnyOfmicrosoftGraphWebBrowserCookieSettings`

GetSafariCookieSettings returns the SafariCookieSettings field if non-nil, zero value otherwise.

### GetSafariCookieSettingsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariCookieSettingsOk() (AnyOfmicrosoftGraphWebBrowserCookieSettings, bool)`

GetSafariCookieSettingsOk returns a tuple with the SafariCookieSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariCookieSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariCookieSettings() bool`

HasSafariCookieSettings returns a boolean if a field has been set.

### SetSafariCookieSettings

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariCookieSettings(v AnyOfmicrosoftGraphWebBrowserCookieSettings)`

SetSafariCookieSettings gets a reference to the given AnyOfmicrosoftGraphWebBrowserCookieSettings and assigns it to the SafariCookieSettings field.

### GetSafariManagedDomains

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariManagedDomains() []string`

GetSafariManagedDomains returns the SafariManagedDomains field if non-nil, zero value otherwise.

### GetSafariManagedDomainsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariManagedDomainsOk() ([]string, bool)`

GetSafariManagedDomainsOk returns a tuple with the SafariManagedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariManagedDomains

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariManagedDomains() bool`

HasSafariManagedDomains returns a boolean if a field has been set.

### SetSafariManagedDomains

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariManagedDomains(v []string)`

SetSafariManagedDomains gets a reference to the given []string and assigns it to the SafariManagedDomains field.

### GetSafariPasswordAutoFillDomains

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariPasswordAutoFillDomains() []string`

GetSafariPasswordAutoFillDomains returns the SafariPasswordAutoFillDomains field if non-nil, zero value otherwise.

### GetSafariPasswordAutoFillDomainsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariPasswordAutoFillDomainsOk() ([]string, bool)`

GetSafariPasswordAutoFillDomainsOk returns a tuple with the SafariPasswordAutoFillDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariPasswordAutoFillDomains

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariPasswordAutoFillDomains() bool`

HasSafariPasswordAutoFillDomains returns a boolean if a field has been set.

### SetSafariPasswordAutoFillDomains

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariPasswordAutoFillDomains(v []string)`

SetSafariPasswordAutoFillDomains gets a reference to the given []string and assigns it to the SafariPasswordAutoFillDomains field.

### GetSafariRequireFraudWarning

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariRequireFraudWarning() bool`

GetSafariRequireFraudWarning returns the SafariRequireFraudWarning field if non-nil, zero value otherwise.

### GetSafariRequireFraudWarningOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSafariRequireFraudWarningOk() (bool, bool)`

GetSafariRequireFraudWarningOk returns a tuple with the SafariRequireFraudWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafariRequireFraudWarning

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSafariRequireFraudWarning() bool`

HasSafariRequireFraudWarning returns a boolean if a field has been set.

### SetSafariRequireFraudWarning

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSafariRequireFraudWarning(v bool)`

SetSafariRequireFraudWarning gets a reference to the given bool and assigns it to the SafariRequireFraudWarning field.

### GetScreenCaptureBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetSiriBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriBlocked() bool`

GetSiriBlocked returns the SiriBlocked field if non-nil, zero value otherwise.

### GetSiriBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriBlockedOk() (bool, bool)`

GetSiriBlockedOk returns a tuple with the SiriBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSiriBlocked() bool`

HasSiriBlocked returns a boolean if a field has been set.

### SetSiriBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSiriBlocked(v bool)`

SetSiriBlocked gets a reference to the given bool and assigns it to the SiriBlocked field.

### GetSiriBlockedWhenLocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriBlockedWhenLocked() bool`

GetSiriBlockedWhenLocked returns the SiriBlockedWhenLocked field if non-nil, zero value otherwise.

### GetSiriBlockedWhenLockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriBlockedWhenLockedOk() (bool, bool)`

GetSiriBlockedWhenLockedOk returns a tuple with the SiriBlockedWhenLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriBlockedWhenLocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSiriBlockedWhenLocked() bool`

HasSiriBlockedWhenLocked returns a boolean if a field has been set.

### SetSiriBlockedWhenLocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSiriBlockedWhenLocked(v bool)`

SetSiriBlockedWhenLocked gets a reference to the given bool and assigns it to the SiriBlockedWhenLocked field.

### GetSiriBlockUserGeneratedContent

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriBlockUserGeneratedContent() bool`

GetSiriBlockUserGeneratedContent returns the SiriBlockUserGeneratedContent field if non-nil, zero value otherwise.

### GetSiriBlockUserGeneratedContentOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriBlockUserGeneratedContentOk() (bool, bool)`

GetSiriBlockUserGeneratedContentOk returns a tuple with the SiriBlockUserGeneratedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriBlockUserGeneratedContent

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSiriBlockUserGeneratedContent() bool`

HasSiriBlockUserGeneratedContent returns a boolean if a field has been set.

### SetSiriBlockUserGeneratedContent

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSiriBlockUserGeneratedContent(v bool)`

SetSiriBlockUserGeneratedContent gets a reference to the given bool and assigns it to the SiriBlockUserGeneratedContent field.

### GetSiriRequireProfanityFilter

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriRequireProfanityFilter() bool`

GetSiriRequireProfanityFilter returns the SiriRequireProfanityFilter field if non-nil, zero value otherwise.

### GetSiriRequireProfanityFilterOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSiriRequireProfanityFilterOk() (bool, bool)`

GetSiriRequireProfanityFilterOk returns a tuple with the SiriRequireProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiriRequireProfanityFilter

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSiriRequireProfanityFilter() bool`

HasSiriRequireProfanityFilter returns a boolean if a field has been set.

### SetSiriRequireProfanityFilter

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSiriRequireProfanityFilter(v bool)`

SetSiriRequireProfanityFilter gets a reference to the given bool and assigns it to the SiriRequireProfanityFilter field.

### GetSpotlightBlockInternetResults

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSpotlightBlockInternetResults() bool`

GetSpotlightBlockInternetResults returns the SpotlightBlockInternetResults field if non-nil, zero value otherwise.

### GetSpotlightBlockInternetResultsOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetSpotlightBlockInternetResultsOk() (bool, bool)`

GetSpotlightBlockInternetResultsOk returns a tuple with the SpotlightBlockInternetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpotlightBlockInternetResults

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasSpotlightBlockInternetResults() bool`

HasSpotlightBlockInternetResults returns a boolean if a field has been set.

### SetSpotlightBlockInternetResults

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetSpotlightBlockInternetResults(v bool)`

SetSpotlightBlockInternetResults gets a reference to the given bool and assigns it to the SpotlightBlockInternetResults field.

### GetVoiceDialingBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetVoiceDialingBlocked() bool`

GetVoiceDialingBlocked returns the VoiceDialingBlocked field if non-nil, zero value otherwise.

### GetVoiceDialingBlockedOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetVoiceDialingBlockedOk() (bool, bool)`

GetVoiceDialingBlockedOk returns a tuple with the VoiceDialingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceDialingBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasVoiceDialingBlocked() bool`

HasVoiceDialingBlocked returns a boolean if a field has been set.

### SetVoiceDialingBlocked

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetVoiceDialingBlocked(v bool)`

SetVoiceDialingBlocked gets a reference to the given bool and assigns it to the VoiceDialingBlocked field.

### GetWallpaperBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetWallpaperBlockModification() bool`

GetWallpaperBlockModification returns the WallpaperBlockModification field if non-nil, zero value otherwise.

### GetWallpaperBlockModificationOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetWallpaperBlockModificationOk() (bool, bool)`

GetWallpaperBlockModificationOk returns a tuple with the WallpaperBlockModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWallpaperBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasWallpaperBlockModification() bool`

HasWallpaperBlockModification returns a boolean if a field has been set.

### SetWallpaperBlockModification

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetWallpaperBlockModification(v bool)`

SetWallpaperBlockModification gets a reference to the given bool and assigns it to the WallpaperBlockModification field.

### GetWiFiConnectOnlyToConfiguredNetworks

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetWiFiConnectOnlyToConfiguredNetworks() bool`

GetWiFiConnectOnlyToConfiguredNetworks returns the WiFiConnectOnlyToConfiguredNetworks field if non-nil, zero value otherwise.

### GetWiFiConnectOnlyToConfiguredNetworksOk

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) GetWiFiConnectOnlyToConfiguredNetworksOk() (bool, bool)`

GetWiFiConnectOnlyToConfiguredNetworksOk returns a tuple with the WiFiConnectOnlyToConfiguredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiConnectOnlyToConfiguredNetworks

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) HasWiFiConnectOnlyToConfiguredNetworks() bool`

HasWiFiConnectOnlyToConfiguredNetworks returns a boolean if a field has been set.

### SetWiFiConnectOnlyToConfiguredNetworks

`func (o *MicrosoftGraphIosGeneralDeviceConfiguration) SetWiFiConnectOnlyToConfiguredNetworks(v bool)`

SetWiFiConnectOnlyToConfiguredNetworks gets a reference to the given bool and assigns it to the WiFiConnectOnlyToConfiguredNetworks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


