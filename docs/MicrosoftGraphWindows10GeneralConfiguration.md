# MicrosoftGraphWindows10GeneralConfiguration

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
**EnterpriseCloudPrintDiscoveryEndPoint** | Pointer to **string** | Endpoint for discovering cloud printers. | [optional] 
**EnterpriseCloudPrintOAuthAuthority** | Pointer to **string** | Authentication endpoint for acquiring OAuth tokens. | [optional] 
**EnterpriseCloudPrintOAuthClientIdentifier** | Pointer to **string** | GUID of a client application authorized to retrieve OAuth tokens from the OAuth Authority. | [optional] 
**EnterpriseCloudPrintResourceIdentifier** | Pointer to **string** | OAuth resource URI for print service as configured in the Azure portal. | [optional] 
**EnterpriseCloudPrintDiscoveryMaxLimit** | Pointer to **int32** | Maximum number of printers that should be queried from a discovery endpoint. This is a mobile only setting. Valid values 1 to 65535 | [optional] 
**EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier** | Pointer to **string** | OAuth resource URI for printer discovery service as configured in Azure portal. | [optional] 
**SearchBlockDiacritics** | Pointer to **bool** | Specifies if search can use diacritics. | [optional] 
**SearchDisableAutoLanguageDetection** | Pointer to **bool** | Specifies whether to use automatic language detection when indexing content and properties. | [optional] 
**SearchDisableIndexingEncryptedItems** | Pointer to **bool** | Indicates whether or not to block indexing of WIP-protected items to prevent them from appearing in search results for Cortana or Explorer. | [optional] 
**SearchEnableRemoteQueries** | Pointer to **bool** | Indicates whether or not to block remote queries of this computer’s index. | [optional] 
**SearchDisableIndexerBackoff** | Pointer to **bool** | Indicates whether or not to disable the search indexer backoff feature. | [optional] 
**SearchDisableIndexingRemovableDrive** | Pointer to **bool** | Indicates whether or not to allow users to add locations on removable drives to libraries and to be indexed. | [optional] 
**SearchEnableAutomaticIndexSizeManangement** | Pointer to **bool** | Specifies minimum amount of hard drive space on the same drive as the index location before indexing stops. | [optional] 
**DiagnosticsDataSubmissionMode** | Pointer to [**AnyOfmicrosoftGraphDiagnosticDataSubmissionMode**](anyOf&lt;microsoft.graph.diagnosticDataSubmissionMode&gt;.md) | Gets or sets a value allowing the device to send diagnostic and usage telemetry data, such as Watson. | [optional] 
**OneDriveDisableFileSync** | Pointer to **bool** | Gets or sets a value allowing IT admins to prevent apps and features from working with files on OneDrive. | [optional] 
**SmartScreenEnableAppInstallControl** | Pointer to **bool** | Allows IT Admins to control whether users are allowed to install apps from places other than the Store. | [optional] 
**PersonalizationDesktopImageUrl** | Pointer to **string** | A http or https Url to a jpg, jpeg or png image that needs to be downloaded and used as the Desktop Image or a file Url to a local image on the file system that needs to used as the Desktop Image. | [optional] 
**PersonalizationLockScreenImageUrl** | Pointer to **string** | A http or https Url to a jpg, jpeg or png image that neeeds to be downloaded and used as the Lock Screen Image or a file Url to a local image on the file system that needs to be used as the Lock Screen Image. | [optional] 
**BluetoothAllowedServices** | Pointer to **[]string** | Specify a list of allowed Bluetooth services and profiles in hex formatted strings. | [optional] 
**BluetoothBlockAdvertising** | Pointer to **bool** | Whether or not to Block the user from using bluetooth advertising. | [optional] 
**BluetoothBlockDiscoverableMode** | Pointer to **bool** | Whether or not to Block the user from using bluetooth discoverable mode. | [optional] 
**BluetoothBlockPrePairing** | Pointer to **bool** | Whether or not to block specific bundled Bluetooth peripherals to automatically pair with the host device. | [optional] 
**EdgeBlockAutofill** | Pointer to **bool** | Indicates whether or not to block auto fill. | [optional] 
**EdgeBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from using the Edge browser. | [optional] 
**EdgeCookiePolicy** | Pointer to [**AnyOfmicrosoftGraphEdgeCookiePolicy**](anyOf&lt;microsoft.graph.edgeCookiePolicy&gt;.md) | Indicates which cookies to block in the Edge browser. | [optional] 
**EdgeBlockDeveloperTools** | Pointer to **bool** | Indicates whether or not to block developer tools in the Edge browser. | [optional] 
**EdgeBlockSendingDoNotTrackHeader** | Pointer to **bool** | Indicates whether or not to Block the user from sending the do not track header. | [optional] 
**EdgeBlockExtensions** | Pointer to **bool** | Indicates whether or not to block extensions in the Edge browser. | [optional] 
**EdgeBlockInPrivateBrowsing** | Pointer to **bool** | Indicates whether or not to block InPrivate browsing on corporate networks, in the Edge browser. | [optional] 
**EdgeBlockJavaScript** | Pointer to **bool** | Indicates whether or not to Block the user from using JavaScript. | [optional] 
**EdgeBlockPasswordManager** | Pointer to **bool** | Indicates whether or not to Block password manager. | [optional] 
**EdgeBlockAddressBarDropdown** | Pointer to **bool** | Block the address bar dropdown functionality in Microsoft Edge. Disable this settings to minimize network connections from Microsoft Edge to Microsoft services. | [optional] 
**EdgeBlockCompatibilityList** | Pointer to **bool** | Block Microsoft compatibility list in Microsoft Edge. This list from Microsoft helps Edge properly display sites with known compatibility issues. | [optional] 
**EdgeClearBrowsingDataOnExit** | Pointer to **bool** | Clear browsing data on exiting Microsoft Edge. | [optional] 
**EdgeAllowStartPagesModification** | Pointer to **bool** | Allow users to change Start pages on Edge. Use the EdgeHomepageUrls to specify the Start pages that the user would see by default when they open Edge. | [optional] 
**EdgeDisableFirstRunPage** | Pointer to **bool** | Block the Microsoft web page that opens on the first use of Microsoft Edge. This policy allows enterprises, like those enrolled in zero emissions configurations, to block this page. | [optional] 
**EdgeBlockLiveTileDataCollection** | Pointer to **bool** | Block the collection of information by Microsoft for live tile creation when users pin a site to Start from Microsoft Edge. | [optional] 
**EdgeSyncFavoritesWithInternetExplorer** | Pointer to **bool** | Enable favorites sync between Internet Explorer and Microsoft Edge. Additions, deletions, modifications and order changes to favorites are shared between browsers. | [optional] 
**CellularBlockDataWhenRoaming** | Pointer to **bool** | Whether or not to Block the user from using data over cellular while roaming. | [optional] 
**CellularBlockVpn** | Pointer to **bool** | Whether or not to Block the user from using VPN over cellular. | [optional] 
**CellularBlockVpnWhenRoaming** | Pointer to **bool** | Whether or not to Block the user from using VPN when roaming over cellular. | [optional] 
**DefenderBlockEndUserAccess** | Pointer to **bool** | Whether or not to block end user access to Defender. | [optional] 
**DefenderDaysBeforeDeletingQuarantinedMalware** | Pointer to **int32** | Number of days before deleting quarantined malware. Valid values 0 to 90 | [optional] 
**DefenderDetectedMalwareActions** | Pointer to [**AnyOfmicrosoftGraphDefenderDetectedMalwareActions**](anyOf&lt;microsoft.graph.defenderDetectedMalwareActions&gt;.md) | Gets or sets Defender’s actions to take on detected Malware per threat level. | [optional] 
**DefenderSystemScanSchedule** | Pointer to [**AnyOfmicrosoftGraphWeeklySchedule**](anyOf&lt;microsoft.graph.weeklySchedule&gt;.md) | Defender day of the week for the system scan. | [optional] 
**DefenderFilesAndFoldersToExclude** | Pointer to **[]string** | Files and folder to exclude from scans and real time protection. | [optional] 
**DefenderFileExtensionsToExclude** | Pointer to **[]string** | File extensions to exclude from scans and real time protection. | [optional] 
**DefenderScanMaxCpu** | Pointer to **int32** | Max CPU usage percentage during scan. Valid values 0 to 100 | [optional] 
**DefenderMonitorFileActivity** | Pointer to [**AnyOfmicrosoftGraphDefenderMonitorFileActivity**](anyOf&lt;microsoft.graph.defenderMonitorFileActivity&gt;.md) | Value for monitoring file activity. | [optional] 
**DefenderProcessesToExclude** | Pointer to **[]string** | Processes to exclude from scans and real time protection. | [optional] 
**DefenderPromptForSampleSubmission** | Pointer to [**AnyOfmicrosoftGraphDefenderPromptForSampleSubmission**](anyOf&lt;microsoft.graph.defenderPromptForSampleSubmission&gt;.md) | The configuration for how to prompt user for sample submission. | [optional] 
**DefenderRequireBehaviorMonitoring** | Pointer to **bool** | Indicates whether or not to require behavior monitoring. | [optional] 
**DefenderRequireCloudProtection** | Pointer to **bool** | Indicates whether or not to require cloud protection. | [optional] 
**DefenderRequireNetworkInspectionSystem** | Pointer to **bool** | Indicates whether or not to require network inspection system. | [optional] 
**DefenderRequireRealTimeMonitoring** | Pointer to **bool** | Indicates whether or not to require real time monitoring. | [optional] 
**DefenderScanArchiveFiles** | Pointer to **bool** | Indicates whether or not to scan archive files. | [optional] 
**DefenderScanDownloads** | Pointer to **bool** | Indicates whether or not to scan downloads. | [optional] 
**DefenderScanNetworkFiles** | Pointer to **bool** | Indicates whether or not to scan files opened from a network folder. | [optional] 
**DefenderScanIncomingMail** | Pointer to **bool** | Indicates whether or not to scan incoming mail messages. | [optional] 
**DefenderScanMappedNetworkDrivesDuringFullScan** | Pointer to **bool** | Indicates whether or not to scan mapped network drives during full scan. | [optional] 
**DefenderScanRemovableDrivesDuringFullScan** | Pointer to **bool** | Indicates whether or not to scan removable drives during full scan. | [optional] 
**DefenderScanScriptsLoadedInInternetExplorer** | Pointer to **bool** | Indicates whether or not to scan scripts loaded in Internet Explorer browser. | [optional] 
**DefenderSignatureUpdateIntervalInHours** | Pointer to **int32** | The signature update interval in hours. Specify 0 not to check. Valid values 0 to 24 | [optional] 
**DefenderScanType** | Pointer to [**AnyOfmicrosoftGraphDefenderScanType**](anyOf&lt;microsoft.graph.defenderScanType&gt;.md) | The defender system scan type. | [optional] 
**DefenderScheduledScanTime** | Pointer to **string** | The defender time for the system scan. | [optional] 
**DefenderScheduledQuickScanTime** | Pointer to **string** | The time to perform a daily quick scan. | [optional] 
**DefenderCloudBlockLevel** | Pointer to [**AnyOfmicrosoftGraphDefenderCloudBlockLevelType**](anyOf&lt;microsoft.graph.defenderCloudBlockLevelType&gt;.md) | Specifies the level of cloud-delivered protection. | [optional] 
**LockScreenAllowTimeoutConfiguration** | Pointer to **bool** | Specify whether to show a user-configurable setting to control the screen timeout while on the lock screen of Windows 10 Mobile devices. If this policy is set to Allow, the value set by lockScreenTimeoutInSeconds is ignored. | [optional] 
**LockScreenBlockActionCenterNotifications** | Pointer to **bool** | Indicates whether or not to block action center notifications over lock screen. | [optional] 
**LockScreenBlockCortana** | Pointer to **bool** | Indicates whether or not the user can interact with Cortana using speech while the system is locked. | [optional] 
**LockScreenBlockToastNotifications** | Pointer to **bool** | Indicates whether to allow toast notifications above the device lock screen. | [optional] 
**LockScreenTimeoutInSeconds** | Pointer to **int32** | Set the duration (in seconds) from the screen locking to the screen turning off for Windows 10 Mobile devices. Supported values are 11-1800. Valid values 11 to 1800 | [optional] 
**PasswordBlockSimple** | Pointer to **bool** | Specify whether PINs or passwords such as \&quot;1111\&quot; or \&quot;1234\&quot; are allowed. For Windows 10 desktops, it also controls the use of picture passwords. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | The password expiration in days. Valid values 0 to 730 | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | The minimum password length. Valid values 4 to 16 | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | The minutes of inactivity before the screen times out. | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | The number of previous passwords to prevent reuse of. Valid values 0 to 50 | [optional] 
**PasswordRequired** | Pointer to **bool** | Indicates whether or not to require the user to have a password. | [optional] 
**PasswordRequireWhenResumeFromIdleState** | Pointer to **bool** | Indicates whether or not to require a password upon resuming from an idle state. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required password type. | [optional] 
**PasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | The number of sign in failures before factory reset. Valid values 0 to 999 | [optional] 
**PrivacyAdvertisingId** | Pointer to [**AnyOfmicrosoftGraphStateManagementSetting**](anyOf&lt;microsoft.graph.stateManagementSetting&gt;.md) | Enables or disables the use of advertising ID. Added in Windows 10, version 1607. | [optional] 
**PrivacyAutoAcceptPairingAndConsentPrompts** | Pointer to **bool** | Indicates whether or not to allow the automatic acceptance of the pairing and privacy user consent dialog when launching apps. | [optional] 
**PrivacyBlockInputPersonalization** | Pointer to **bool** | Indicates whether or not to block the usage of cloud based speech services for Cortana, Dictation, or Store applications. | [optional] 
**StartBlockUnpinningAppsFromTaskbar** | Pointer to **bool** | Indicates whether or not to block the user from unpinning apps from taskbar. | [optional] 
**StartMenuAppListVisibility** | Pointer to [**AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType**](anyOf&lt;microsoft.graph.windowsStartMenuAppListVisibilityType&gt;.md) | Setting the value of this collapses the app list, removes the app list entirely, or disables the corresponding toggle in the Settings app. | [optional] 
**StartMenuHideChangeAccountSettings** | Pointer to **bool** | Enabling this policy hides the change account setting from appearing in the user tile in the start menu. | [optional] 
**StartMenuHideFrequentlyUsedApps** | Pointer to **bool** | Enabling this policy hides the most used apps from appearing on the start menu and disables the corresponding toggle in the Settings app. | [optional] 
**StartMenuHideHibernate** | Pointer to **bool** | Enabling this policy hides hibernate from appearing in the power button in the start menu. | [optional] 
**StartMenuHideLock** | Pointer to **bool** | Enabling this policy hides lock from appearing in the user tile in the start menu. | [optional] 
**StartMenuHidePowerButton** | Pointer to **bool** | Enabling this policy hides the power button from appearing in the start menu. | [optional] 
**StartMenuHideRecentJumpLists** | Pointer to **bool** | Enabling this policy hides recent jump lists from appearing on the start menu/taskbar and disables the corresponding toggle in the Settings app. | [optional] 
**StartMenuHideRecentlyAddedApps** | Pointer to **bool** | Enabling this policy hides recently added apps from appearing on the start menu and disables the corresponding toggle in the Settings app. | [optional] 
**StartMenuHideRestartOptions** | Pointer to **bool** | Enabling this policy hides “Restart/Update and Restart” from appearing in the power button in the start menu. | [optional] 
**StartMenuHideShutDown** | Pointer to **bool** | Enabling this policy hides shut down/update and shut down from appearing in the power button in the start menu. | [optional] 
**StartMenuHideSignOut** | Pointer to **bool** | Enabling this policy hides sign out from appearing in the user tile in the start menu. | [optional] 
**StartMenuHideSleep** | Pointer to **bool** | Enabling this policy hides sleep from appearing in the power button in the start menu. | [optional] 
**StartMenuHideSwitchAccount** | Pointer to **bool** | Enabling this policy hides switch account from appearing in the user tile in the start menu. | [optional] 
**StartMenuHideUserTile** | Pointer to **bool** | Enabling this policy hides the user tile from appearing in the start menu. | [optional] 
**StartMenuLayoutEdgeAssetsXml** | Pointer to **string** | This policy setting allows you to import Edge assets to be used with startMenuLayoutXml policy. Start layout can contain secondary tile from Edge app which looks for Edge local asset file. Edge local asset would not exist and cause Edge secondary tile to appear empty in this case. This policy only gets applied when startMenuLayoutXml policy is modified. The value should be a UTF-8 Base64 encoded byte array. | [optional] 
**StartMenuLayoutXml** | Pointer to **string** | Allows admins to override the default Start menu layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in a UTF8 encoded byte array format. | [optional] 
**StartMenuMode** | Pointer to [**AnyOfmicrosoftGraphWindowsStartMenuModeType**](anyOf&lt;microsoft.graph.windowsStartMenuModeType&gt;.md) | Allows admins to decide how the Start menu is displayed. | [optional] 
**StartMenuPinnedFolderDocuments** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Documents folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderDownloads** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Downloads folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderFileExplorer** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the FileExplorer shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderHomeGroup** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the HomeGroup folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderMusic** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Music folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderNetwork** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Network folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderPersonalFolder** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the PersonalFolder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderPictures** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Pictures folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderSettings** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Settings folder shortcut on the Start menu. | [optional] 
**StartMenuPinnedFolderVideos** | Pointer to [**AnyOfmicrosoftGraphVisibilitySetting**](anyOf&lt;microsoft.graph.visibilitySetting&gt;.md) | Enforces the visibility (Show/Hide) of the Videos folder shortcut on the Start menu. | [optional] 
**SettingsBlockSettingsApp** | Pointer to **bool** | Indicates whether or not to block access to Settings app. | [optional] 
**SettingsBlockSystemPage** | Pointer to **bool** | Indicates whether or not to block access to System in Settings app. | [optional] 
**SettingsBlockDevicesPage** | Pointer to **bool** | Indicates whether or not to block access to Devices in Settings app. | [optional] 
**SettingsBlockNetworkInternetPage** | Pointer to **bool** | Indicates whether or not to block access to Network &amp; Internet in Settings app. | [optional] 
**SettingsBlockPersonalizationPage** | Pointer to **bool** | Indicates whether or not to block access to Personalization in Settings app. | [optional] 
**SettingsBlockAccountsPage** | Pointer to **bool** | Indicates whether or not to block access to Accounts in Settings app. | [optional] 
**SettingsBlockTimeLanguagePage** | Pointer to **bool** | Indicates whether or not to block access to Time &amp; Language in Settings app. | [optional] 
**SettingsBlockEaseOfAccessPage** | Pointer to **bool** | Indicates whether or not to block access to Ease of Access in Settings app. | [optional] 
**SettingsBlockPrivacyPage** | Pointer to **bool** | Indicates whether or not to block access to Privacy in Settings app. | [optional] 
**SettingsBlockUpdateSecurityPage** | Pointer to **bool** | Indicates whether or not to block access to Update &amp; Security in Settings app. | [optional] 
**SettingsBlockAppsPage** | Pointer to **bool** | Indicates whether or not to block access to Apps in Settings app. | [optional] 
**SettingsBlockGamingPage** | Pointer to **bool** | Indicates whether or not to block access to Gaming in Settings app. | [optional] 
**WindowsSpotlightBlockConsumerSpecificFeatures** | Pointer to **bool** | Allows IT admins to block experiences that are typically for consumers only, such as Start suggestions, Membership notifications, Post-OOBE app install and redirect tiles. | [optional] 
**WindowsSpotlightBlocked** | Pointer to **bool** | Allows IT admins to turn off all Windows Spotlight features | [optional] 
**WindowsSpotlightBlockOnActionCenter** | Pointer to **bool** | Block suggestions from Microsoft that show after each OS clean install, upgrade or in an on-going basis to introduce users to what is new or changed | [optional] 
**WindowsSpotlightBlockTailoredExperiences** | Pointer to **bool** | Block personalized content in Windows spotlight based on user’s device usage. | [optional] 
**WindowsSpotlightBlockThirdPartyNotifications** | Pointer to **bool** | Block third party content delivered via Windows Spotlight | [optional] 
**WindowsSpotlightBlockWelcomeExperience** | Pointer to **bool** | Block Windows Spotlight Windows welcome experience | [optional] 
**WindowsSpotlightBlockWindowsTips** | Pointer to **bool** | Allows IT admins to turn off the popup of Windows Tips. | [optional] 
**WindowsSpotlightConfigureOnLockScreen** | Pointer to [**AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings**](anyOf&lt;microsoft.graph.windowsSpotlightEnablementSettings&gt;.md) | Specifies the type of Spotlight | [optional] 
**NetworkProxyApplySettingsDeviceWide** | Pointer to **bool** | If set, proxy settings will be applied to all processes and accounts in the device. Otherwise, it will be applied to the user account that’s enrolled into MDM. | [optional] 
**NetworkProxyDisableAutoDetect** | Pointer to **bool** | Disable automatic detection of settings. If enabled, the system will try to find the path to a proxy auto-config (PAC) script. | [optional] 
**NetworkProxyAutomaticConfigurationUrl** | Pointer to **string** | Address to the proxy auto-config (PAC) script you want to use. | [optional] 
**NetworkProxyServer** | Pointer to [**AnyOfmicrosoftGraphWindows10NetworkProxyServer**](anyOf&lt;microsoft.graph.windows10NetworkProxyServer&gt;.md) | Specifies manual proxy server settings. | [optional] 
**AccountsBlockAddingNonMicrosoftAccountEmail** | Pointer to **bool** | Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account. | [optional] 
**AntiTheftModeBlocked** | Pointer to **bool** | Indicates whether or not to block the user from selecting an AntiTheft mode preference (Windows 10 Mobile only). | [optional] 
**BluetoothBlocked** | Pointer to **bool** | Whether or not to Block the user from using bluetooth. | [optional] 
**CameraBlocked** | Pointer to **bool** | Whether or not to Block the user from accessing the camera of the device. | [optional] 
**ConnectedDevicesServiceBlocked** | Pointer to **bool** | Whether or not to block Connected Devices Service which enables discovery and connection to other devices, remote messaging, remote app sessions and other cross-device experiences. | [optional] 
**CertificatesBlockManualRootCertificateInstallation** | Pointer to **bool** | Whether or not to Block the user from doing manual root certificate installation. | [optional] 
**CopyPasteBlocked** | Pointer to **bool** | Whether or not to Block the user from using copy paste. | [optional] 
**CortanaBlocked** | Pointer to **bool** | Whether or not to Block the user from using Cortana. | [optional] 
**DeviceManagementBlockFactoryResetOnMobile** | Pointer to **bool** | Indicates whether or not to Block the user from resetting their phone. | [optional] 
**DeviceManagementBlockManualUnenroll** | Pointer to **bool** | Indicates whether or not to Block the user from doing manual un-enrollment from device management. | [optional] 
**SafeSearchFilter** | Pointer to [**AnyOfmicrosoftGraphSafeSearchFilterType**](anyOf&lt;microsoft.graph.safeSearchFilterType&gt;.md) | Specifies what filter level of safe search is required. | [optional] 
**EdgeBlockPopups** | Pointer to **bool** | Indicates whether or not to block popups. | [optional] 
**EdgeBlockSearchSuggestions** | Pointer to **bool** | Indicates whether or not to block the user from using the search suggestions in the address bar. | [optional] 
**EdgeBlockSendingIntranetTrafficToInternetExplorer** | Pointer to **bool** | Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. Note: the name of this property is misleading; the property is obsolete, use EdgeSendIntranetTrafficToInternetExplorer instead. | [optional] 
**EdgeSendIntranetTrafficToInternetExplorer** | Pointer to **bool** | Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. | [optional] 
**EdgeRequireSmartScreen** | Pointer to **bool** | Indicates whether or not to Require the user to use the smart screen filter. | [optional] 
**EdgeEnterpriseModeSiteListLocation** | Pointer to **string** | Indicates the enterprise mode site list location. Could be a local file, local network or http location. | [optional] 
**EdgeFirstRunUrl** | Pointer to **string** | The first run URL for when Edge browser is opened for the first time. | [optional] 
**EdgeSearchEngine** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their default search engine provided the AllowSearchEngineCustomization policy is not set. | [optional] 
**EdgeHomepageUrls** | Pointer to **[]string** | The list of URLs for homepages shodwn on MDM-enrolled devices on Edge browser. | [optional] 
**EdgeBlockAccessToAboutFlags** | Pointer to **bool** | Indicates whether or not to prevent access to about flags on Edge browser. | [optional] 
**SmartScreenBlockPromptOverride** | Pointer to **bool** | Indicates whether or not users can override SmartScreen Filter warnings about potentially malicious websites. | [optional] 
**SmartScreenBlockPromptOverrideForFiles** | Pointer to **bool** | Indicates whether or not users can override the SmartScreen Filter warnings about downloading unverified files | [optional] 
**WebRtcBlockLocalhostIpAddress** | Pointer to **bool** | Indicates whether or not user&#39;s localhost IP address is displayed while making phone calls using the WebRTC | [optional] 
**InternetSharingBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from using internet sharing. | [optional] 
**SettingsBlockAddProvisioningPackage** | Pointer to **bool** | Indicates whether or not to block the user from installing provisioning packages. | [optional] 
**SettingsBlockRemoveProvisioningPackage** | Pointer to **bool** | Indicates whether or not to block the runtime configuration agent from removing provisioning packages. | [optional] 
**SettingsBlockChangeSystemTime** | Pointer to **bool** | Indicates whether or not to block the user from changing date and time settings. | [optional] 
**SettingsBlockEditDeviceName** | Pointer to **bool** | Indicates whether or not to block the user from editing the device name. | [optional] 
**SettingsBlockChangeRegion** | Pointer to **bool** | Indicates whether or not to block the user from changing the region settings. | [optional] 
**SettingsBlockChangeLanguage** | Pointer to **bool** | Indicates whether or not to block the user from changing the language settings. | [optional] 
**SettingsBlockChangePowerSleep** | Pointer to **bool** | Indicates whether or not to block the user from changing power and sleep settings. | [optional] 
**LocationServicesBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from location services. | [optional] 
**MicrosoftAccountBlocked** | Pointer to **bool** | Indicates whether or not to Block a Microsoft account. | [optional] 
**MicrosoftAccountBlockSettingsSync** | Pointer to **bool** | Indicates whether or not to Block Microsoft account settings sync. | [optional] 
**NfcBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from using near field communication. | [optional] 
**ResetProtectionModeBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from reset protection mode. | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from taking Screenshots. | [optional] 
**StorageBlockRemovableStorage** | Pointer to **bool** | Indicates whether or not to Block the user from using removable storage. | [optional] 
**StorageRequireMobileDeviceEncryption** | Pointer to **bool** | Indicating whether or not to require encryption on a mobile device. | [optional] 
**UsbBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from USB connection. | [optional] 
**VoiceRecordingBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from voice recording. | [optional] 
**WiFiBlockAutomaticConnectHotspots** | Pointer to **bool** | Indicating whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked. | [optional] 
**WiFiBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from using Wi-Fi. | [optional] 
**WiFiBlockManualConfiguration** | Pointer to **bool** | Indicates whether or not to Block the user from using Wi-Fi manual configuration. | [optional] 
**WiFiScanInterval** | Pointer to **int32** | Specify how often devices scan for Wi-Fi networks. Supported values are 1-500, where 100 &#x3D; default, and 500 &#x3D; low frequency. Valid values 1 to 500 | [optional] 
**WirelessDisplayBlockProjectionToThisDevice** | Pointer to **bool** | Indicates whether or not to allow other devices from discovering this PC for projection. | [optional] 
**WirelessDisplayBlockUserInputFromReceiver** | Pointer to **bool** | Indicates whether or not to allow user input from wireless display receiver. | [optional] 
**WirelessDisplayRequirePinForPairing** | Pointer to **bool** | Indicates whether or not to require a PIN for new devices to initiate pairing. | [optional] 
**WindowsStoreBlocked** | Pointer to **bool** | Indicates whether or not to Block the user from using the Windows store. | [optional] 
**AppsAllowTrustedAppsSideloading** | Pointer to [**AnyOfmicrosoftGraphStateManagementSetting**](anyOf&lt;microsoft.graph.stateManagementSetting&gt;.md) | Indicates whether apps from AppX packages signed with a trusted certificate can be side loaded. | [optional] 
**WindowsStoreBlockAutoUpdate** | Pointer to **bool** | Indicates whether or not to block automatic update of apps from Windows Store. | [optional] 
**DeveloperUnlockSetting** | Pointer to [**AnyOfmicrosoftGraphStateManagementSetting**](anyOf&lt;microsoft.graph.stateManagementSetting&gt;.md) | Indicates whether or not to allow developer unlock. | [optional] 
**SharedUserAppDataAllowed** | Pointer to **bool** | Indicates whether or not to block multiple users of the same app to share data. | [optional] 
**AppsBlockWindowsStoreOriginatedApps** | Pointer to **bool** | Indicates whether or not to disable the launch of all apps from Windows Store that came pre-installed or were downloaded. | [optional] 
**WindowsStoreEnablePrivateStoreOnly** | Pointer to **bool** | Indicates whether or not to enable Private Store Only. | [optional] 
**StorageRestrictAppDataToSystemVolume** | Pointer to **bool** | Indicates whether application data is restricted to the system drive. | [optional] 
**StorageRestrictAppInstallToSystemVolume** | Pointer to **bool** | Indicates whether the installation of applications is restricted to the system drive. | [optional] 
**GameDvrBlocked** | Pointer to **bool** | Indicates whether or not to block DVR and broadcasting. | [optional] 
**ExperienceBlockDeviceDiscovery** | Pointer to **bool** | Indicates whether or not to enable device discovery UX. | [optional] 
**ExperienceBlockErrorDialogWhenNoSIM** | Pointer to **bool** | Indicates whether or not to allow the error dialog from displaying if no SIM card is detected. | [optional] 
**ExperienceBlockTaskSwitcher** | Pointer to **bool** | Indicates whether or not to enable task switching on the device. | [optional] 
**LogonBlockFastUserSwitching** | Pointer to **bool** | Disables the ability to quickly switch between users that are logged on simultaneously without logging off. | [optional] 
**TenantLockdownRequireNetworkDuringOutOfBoxExperience** | Pointer to **bool** | Whether the device is required to connect to the network. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetEnterpriseCloudPrintDiscoveryEndPoint

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryEndPoint() string`

GetEnterpriseCloudPrintDiscoveryEndPoint returns the EnterpriseCloudPrintDiscoveryEndPoint field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintDiscoveryEndPointOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryEndPointOk() (string, bool)`

GetEnterpriseCloudPrintDiscoveryEndPointOk returns a tuple with the EnterpriseCloudPrintDiscoveryEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintDiscoveryEndPoint

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEnterpriseCloudPrintDiscoveryEndPoint() bool`

HasEnterpriseCloudPrintDiscoveryEndPoint returns a boolean if a field has been set.

### SetEnterpriseCloudPrintDiscoveryEndPoint

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryEndPoint(v string)`

SetEnterpriseCloudPrintDiscoveryEndPoint gets a reference to the given string and assigns it to the EnterpriseCloudPrintDiscoveryEndPoint field.

### SetEnterpriseCloudPrintDiscoveryEndPointExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryEndPointExplicitNull(b bool)`

SetEnterpriseCloudPrintDiscoveryEndPointExplicitNull (un)sets EnterpriseCloudPrintDiscoveryEndPoint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintDiscoveryEndPoint value is set to nil even if false is passed
### GetEnterpriseCloudPrintOAuthAuthority

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthAuthority() string`

GetEnterpriseCloudPrintOAuthAuthority returns the EnterpriseCloudPrintOAuthAuthority field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintOAuthAuthorityOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthAuthorityOk() (string, bool)`

GetEnterpriseCloudPrintOAuthAuthorityOk returns a tuple with the EnterpriseCloudPrintOAuthAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintOAuthAuthority

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEnterpriseCloudPrintOAuthAuthority() bool`

HasEnterpriseCloudPrintOAuthAuthority returns a boolean if a field has been set.

### SetEnterpriseCloudPrintOAuthAuthority

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthAuthority(v string)`

SetEnterpriseCloudPrintOAuthAuthority gets a reference to the given string and assigns it to the EnterpriseCloudPrintOAuthAuthority field.

### SetEnterpriseCloudPrintOAuthAuthorityExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthAuthorityExplicitNull(b bool)`

SetEnterpriseCloudPrintOAuthAuthorityExplicitNull (un)sets EnterpriseCloudPrintOAuthAuthority to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintOAuthAuthority value is set to nil even if false is passed
### GetEnterpriseCloudPrintOAuthClientIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthClientIdentifier() string`

GetEnterpriseCloudPrintOAuthClientIdentifier returns the EnterpriseCloudPrintOAuthClientIdentifier field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintOAuthClientIdentifierOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthClientIdentifierOk() (string, bool)`

GetEnterpriseCloudPrintOAuthClientIdentifierOk returns a tuple with the EnterpriseCloudPrintOAuthClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintOAuthClientIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEnterpriseCloudPrintOAuthClientIdentifier() bool`

HasEnterpriseCloudPrintOAuthClientIdentifier returns a boolean if a field has been set.

### SetEnterpriseCloudPrintOAuthClientIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthClientIdentifier(v string)`

SetEnterpriseCloudPrintOAuthClientIdentifier gets a reference to the given string and assigns it to the EnterpriseCloudPrintOAuthClientIdentifier field.

### SetEnterpriseCloudPrintOAuthClientIdentifierExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthClientIdentifierExplicitNull(b bool)`

SetEnterpriseCloudPrintOAuthClientIdentifierExplicitNull (un)sets EnterpriseCloudPrintOAuthClientIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintOAuthClientIdentifier value is set to nil even if false is passed
### GetEnterpriseCloudPrintResourceIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintResourceIdentifier() string`

GetEnterpriseCloudPrintResourceIdentifier returns the EnterpriseCloudPrintResourceIdentifier field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintResourceIdentifierOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintResourceIdentifierOk() (string, bool)`

GetEnterpriseCloudPrintResourceIdentifierOk returns a tuple with the EnterpriseCloudPrintResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintResourceIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEnterpriseCloudPrintResourceIdentifier() bool`

HasEnterpriseCloudPrintResourceIdentifier returns a boolean if a field has been set.

### SetEnterpriseCloudPrintResourceIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintResourceIdentifier(v string)`

SetEnterpriseCloudPrintResourceIdentifier gets a reference to the given string and assigns it to the EnterpriseCloudPrintResourceIdentifier field.

### SetEnterpriseCloudPrintResourceIdentifierExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintResourceIdentifierExplicitNull(b bool)`

SetEnterpriseCloudPrintResourceIdentifierExplicitNull (un)sets EnterpriseCloudPrintResourceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintResourceIdentifier value is set to nil even if false is passed
### GetEnterpriseCloudPrintDiscoveryMaxLimit

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryMaxLimit() int32`

GetEnterpriseCloudPrintDiscoveryMaxLimit returns the EnterpriseCloudPrintDiscoveryMaxLimit field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintDiscoveryMaxLimitOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryMaxLimitOk() (int32, bool)`

GetEnterpriseCloudPrintDiscoveryMaxLimitOk returns a tuple with the EnterpriseCloudPrintDiscoveryMaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintDiscoveryMaxLimit

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEnterpriseCloudPrintDiscoveryMaxLimit() bool`

HasEnterpriseCloudPrintDiscoveryMaxLimit returns a boolean if a field has been set.

### SetEnterpriseCloudPrintDiscoveryMaxLimit

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryMaxLimit(v int32)`

SetEnterpriseCloudPrintDiscoveryMaxLimit gets a reference to the given int32 and assigns it to the EnterpriseCloudPrintDiscoveryMaxLimit field.

### SetEnterpriseCloudPrintDiscoveryMaxLimitExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryMaxLimitExplicitNull(b bool)`

SetEnterpriseCloudPrintDiscoveryMaxLimitExplicitNull (un)sets EnterpriseCloudPrintDiscoveryMaxLimit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintDiscoveryMaxLimit value is set to nil even if false is passed
### GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier() string`

GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier returns the EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierOk() (string, bool)`

GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierOk returns a tuple with the EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier() bool`

HasEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier returns a boolean if a field has been set.

### SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier(v string)`

SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier gets a reference to the given string and assigns it to the EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier field.

### SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierExplicitNull(b bool)`

SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierExplicitNull (un)sets EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier value is set to nil even if false is passed
### GetSearchBlockDiacritics

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchBlockDiacritics() bool`

GetSearchBlockDiacritics returns the SearchBlockDiacritics field if non-nil, zero value otherwise.

### GetSearchBlockDiacriticsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchBlockDiacriticsOk() (bool, bool)`

GetSearchBlockDiacriticsOk returns a tuple with the SearchBlockDiacritics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchBlockDiacritics

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchBlockDiacritics() bool`

HasSearchBlockDiacritics returns a boolean if a field has been set.

### SetSearchBlockDiacritics

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchBlockDiacritics(v bool)`

SetSearchBlockDiacritics gets a reference to the given bool and assigns it to the SearchBlockDiacritics field.

### GetSearchDisableAutoLanguageDetection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableAutoLanguageDetection() bool`

GetSearchDisableAutoLanguageDetection returns the SearchDisableAutoLanguageDetection field if non-nil, zero value otherwise.

### GetSearchDisableAutoLanguageDetectionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableAutoLanguageDetectionOk() (bool, bool)`

GetSearchDisableAutoLanguageDetectionOk returns a tuple with the SearchDisableAutoLanguageDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableAutoLanguageDetection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchDisableAutoLanguageDetection() bool`

HasSearchDisableAutoLanguageDetection returns a boolean if a field has been set.

### SetSearchDisableAutoLanguageDetection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchDisableAutoLanguageDetection(v bool)`

SetSearchDisableAutoLanguageDetection gets a reference to the given bool and assigns it to the SearchDisableAutoLanguageDetection field.

### GetSearchDisableIndexingEncryptedItems

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableIndexingEncryptedItems() bool`

GetSearchDisableIndexingEncryptedItems returns the SearchDisableIndexingEncryptedItems field if non-nil, zero value otherwise.

### GetSearchDisableIndexingEncryptedItemsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableIndexingEncryptedItemsOk() (bool, bool)`

GetSearchDisableIndexingEncryptedItemsOk returns a tuple with the SearchDisableIndexingEncryptedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableIndexingEncryptedItems

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchDisableIndexingEncryptedItems() bool`

HasSearchDisableIndexingEncryptedItems returns a boolean if a field has been set.

### SetSearchDisableIndexingEncryptedItems

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchDisableIndexingEncryptedItems(v bool)`

SetSearchDisableIndexingEncryptedItems gets a reference to the given bool and assigns it to the SearchDisableIndexingEncryptedItems field.

### GetSearchEnableRemoteQueries

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchEnableRemoteQueries() bool`

GetSearchEnableRemoteQueries returns the SearchEnableRemoteQueries field if non-nil, zero value otherwise.

### GetSearchEnableRemoteQueriesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchEnableRemoteQueriesOk() (bool, bool)`

GetSearchEnableRemoteQueriesOk returns a tuple with the SearchEnableRemoteQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchEnableRemoteQueries

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchEnableRemoteQueries() bool`

HasSearchEnableRemoteQueries returns a boolean if a field has been set.

### SetSearchEnableRemoteQueries

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchEnableRemoteQueries(v bool)`

SetSearchEnableRemoteQueries gets a reference to the given bool and assigns it to the SearchEnableRemoteQueries field.

### GetSearchDisableIndexerBackoff

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableIndexerBackoff() bool`

GetSearchDisableIndexerBackoff returns the SearchDisableIndexerBackoff field if non-nil, zero value otherwise.

### GetSearchDisableIndexerBackoffOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableIndexerBackoffOk() (bool, bool)`

GetSearchDisableIndexerBackoffOk returns a tuple with the SearchDisableIndexerBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableIndexerBackoff

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchDisableIndexerBackoff() bool`

HasSearchDisableIndexerBackoff returns a boolean if a field has been set.

### SetSearchDisableIndexerBackoff

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchDisableIndexerBackoff(v bool)`

SetSearchDisableIndexerBackoff gets a reference to the given bool and assigns it to the SearchDisableIndexerBackoff field.

### GetSearchDisableIndexingRemovableDrive

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableIndexingRemovableDrive() bool`

GetSearchDisableIndexingRemovableDrive returns the SearchDisableIndexingRemovableDrive field if non-nil, zero value otherwise.

### GetSearchDisableIndexingRemovableDriveOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchDisableIndexingRemovableDriveOk() (bool, bool)`

GetSearchDisableIndexingRemovableDriveOk returns a tuple with the SearchDisableIndexingRemovableDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableIndexingRemovableDrive

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchDisableIndexingRemovableDrive() bool`

HasSearchDisableIndexingRemovableDrive returns a boolean if a field has been set.

### SetSearchDisableIndexingRemovableDrive

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchDisableIndexingRemovableDrive(v bool)`

SetSearchDisableIndexingRemovableDrive gets a reference to the given bool and assigns it to the SearchDisableIndexingRemovableDrive field.

### GetSearchEnableAutomaticIndexSizeManangement

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchEnableAutomaticIndexSizeManangement() bool`

GetSearchEnableAutomaticIndexSizeManangement returns the SearchEnableAutomaticIndexSizeManangement field if non-nil, zero value otherwise.

### GetSearchEnableAutomaticIndexSizeManangementOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSearchEnableAutomaticIndexSizeManangementOk() (bool, bool)`

GetSearchEnableAutomaticIndexSizeManangementOk returns a tuple with the SearchEnableAutomaticIndexSizeManangement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchEnableAutomaticIndexSizeManangement

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSearchEnableAutomaticIndexSizeManangement() bool`

HasSearchEnableAutomaticIndexSizeManangement returns a boolean if a field has been set.

### SetSearchEnableAutomaticIndexSizeManangement

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSearchEnableAutomaticIndexSizeManangement(v bool)`

SetSearchEnableAutomaticIndexSizeManangement gets a reference to the given bool and assigns it to the SearchEnableAutomaticIndexSizeManangement field.

### GetDiagnosticsDataSubmissionMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDiagnosticsDataSubmissionMode() AnyOfmicrosoftGraphDiagnosticDataSubmissionMode`

GetDiagnosticsDataSubmissionMode returns the DiagnosticsDataSubmissionMode field if non-nil, zero value otherwise.

### GetDiagnosticsDataSubmissionModeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDiagnosticsDataSubmissionModeOk() (AnyOfmicrosoftGraphDiagnosticDataSubmissionMode, bool)`

GetDiagnosticsDataSubmissionModeOk returns a tuple with the DiagnosticsDataSubmissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticsDataSubmissionMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDiagnosticsDataSubmissionMode() bool`

HasDiagnosticsDataSubmissionMode returns a boolean if a field has been set.

### SetDiagnosticsDataSubmissionMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDiagnosticsDataSubmissionMode(v AnyOfmicrosoftGraphDiagnosticDataSubmissionMode)`

SetDiagnosticsDataSubmissionMode gets a reference to the given AnyOfmicrosoftGraphDiagnosticDataSubmissionMode and assigns it to the DiagnosticsDataSubmissionMode field.

### GetOneDriveDisableFileSync

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetOneDriveDisableFileSync() bool`

GetOneDriveDisableFileSync returns the OneDriveDisableFileSync field if non-nil, zero value otherwise.

### GetOneDriveDisableFileSyncOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetOneDriveDisableFileSyncOk() (bool, bool)`

GetOneDriveDisableFileSyncOk returns a tuple with the OneDriveDisableFileSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOneDriveDisableFileSync

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasOneDriveDisableFileSync() bool`

HasOneDriveDisableFileSync returns a boolean if a field has been set.

### SetOneDriveDisableFileSync

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetOneDriveDisableFileSync(v bool)`

SetOneDriveDisableFileSync gets a reference to the given bool and assigns it to the OneDriveDisableFileSync field.

### GetSmartScreenEnableAppInstallControl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSmartScreenEnableAppInstallControl() bool`

GetSmartScreenEnableAppInstallControl returns the SmartScreenEnableAppInstallControl field if non-nil, zero value otherwise.

### GetSmartScreenEnableAppInstallControlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSmartScreenEnableAppInstallControlOk() (bool, bool)`

GetSmartScreenEnableAppInstallControlOk returns a tuple with the SmartScreenEnableAppInstallControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenEnableAppInstallControl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSmartScreenEnableAppInstallControl() bool`

HasSmartScreenEnableAppInstallControl returns a boolean if a field has been set.

### SetSmartScreenEnableAppInstallControl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSmartScreenEnableAppInstallControl(v bool)`

SetSmartScreenEnableAppInstallControl gets a reference to the given bool and assigns it to the SmartScreenEnableAppInstallControl field.

### GetPersonalizationDesktopImageUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPersonalizationDesktopImageUrl() string`

GetPersonalizationDesktopImageUrl returns the PersonalizationDesktopImageUrl field if non-nil, zero value otherwise.

### GetPersonalizationDesktopImageUrlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPersonalizationDesktopImageUrlOk() (string, bool)`

GetPersonalizationDesktopImageUrlOk returns a tuple with the PersonalizationDesktopImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalizationDesktopImageUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPersonalizationDesktopImageUrl() bool`

HasPersonalizationDesktopImageUrl returns a boolean if a field has been set.

### SetPersonalizationDesktopImageUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPersonalizationDesktopImageUrl(v string)`

SetPersonalizationDesktopImageUrl gets a reference to the given string and assigns it to the PersonalizationDesktopImageUrl field.

### SetPersonalizationDesktopImageUrlExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPersonalizationDesktopImageUrlExplicitNull(b bool)`

SetPersonalizationDesktopImageUrlExplicitNull (un)sets PersonalizationDesktopImageUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonalizationDesktopImageUrl value is set to nil even if false is passed
### GetPersonalizationLockScreenImageUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPersonalizationLockScreenImageUrl() string`

GetPersonalizationLockScreenImageUrl returns the PersonalizationLockScreenImageUrl field if non-nil, zero value otherwise.

### GetPersonalizationLockScreenImageUrlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPersonalizationLockScreenImageUrlOk() (string, bool)`

GetPersonalizationLockScreenImageUrlOk returns a tuple with the PersonalizationLockScreenImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalizationLockScreenImageUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPersonalizationLockScreenImageUrl() bool`

HasPersonalizationLockScreenImageUrl returns a boolean if a field has been set.

### SetPersonalizationLockScreenImageUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPersonalizationLockScreenImageUrl(v string)`

SetPersonalizationLockScreenImageUrl gets a reference to the given string and assigns it to the PersonalizationLockScreenImageUrl field.

### SetPersonalizationLockScreenImageUrlExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPersonalizationLockScreenImageUrlExplicitNull(b bool)`

SetPersonalizationLockScreenImageUrlExplicitNull (un)sets PersonalizationLockScreenImageUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonalizationLockScreenImageUrl value is set to nil even if false is passed
### GetBluetoothAllowedServices

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothAllowedServices() []string`

GetBluetoothAllowedServices returns the BluetoothAllowedServices field if non-nil, zero value otherwise.

### GetBluetoothAllowedServicesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothAllowedServicesOk() ([]string, bool)`

GetBluetoothAllowedServicesOk returns a tuple with the BluetoothAllowedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothAllowedServices

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasBluetoothAllowedServices() bool`

HasBluetoothAllowedServices returns a boolean if a field has been set.

### SetBluetoothAllowedServices

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetBluetoothAllowedServices(v []string)`

SetBluetoothAllowedServices gets a reference to the given []string and assigns it to the BluetoothAllowedServices field.

### GetBluetoothBlockAdvertising

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockAdvertising() bool`

GetBluetoothBlockAdvertising returns the BluetoothBlockAdvertising field if non-nil, zero value otherwise.

### GetBluetoothBlockAdvertisingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockAdvertisingOk() (bool, bool)`

GetBluetoothBlockAdvertisingOk returns a tuple with the BluetoothBlockAdvertising field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockAdvertising

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasBluetoothBlockAdvertising() bool`

HasBluetoothBlockAdvertising returns a boolean if a field has been set.

### SetBluetoothBlockAdvertising

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetBluetoothBlockAdvertising(v bool)`

SetBluetoothBlockAdvertising gets a reference to the given bool and assigns it to the BluetoothBlockAdvertising field.

### GetBluetoothBlockDiscoverableMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockDiscoverableMode() bool`

GetBluetoothBlockDiscoverableMode returns the BluetoothBlockDiscoverableMode field if non-nil, zero value otherwise.

### GetBluetoothBlockDiscoverableModeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockDiscoverableModeOk() (bool, bool)`

GetBluetoothBlockDiscoverableModeOk returns a tuple with the BluetoothBlockDiscoverableMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockDiscoverableMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasBluetoothBlockDiscoverableMode() bool`

HasBluetoothBlockDiscoverableMode returns a boolean if a field has been set.

### SetBluetoothBlockDiscoverableMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetBluetoothBlockDiscoverableMode(v bool)`

SetBluetoothBlockDiscoverableMode gets a reference to the given bool and assigns it to the BluetoothBlockDiscoverableMode field.

### GetBluetoothBlockPrePairing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockPrePairing() bool`

GetBluetoothBlockPrePairing returns the BluetoothBlockPrePairing field if non-nil, zero value otherwise.

### GetBluetoothBlockPrePairingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockPrePairingOk() (bool, bool)`

GetBluetoothBlockPrePairingOk returns a tuple with the BluetoothBlockPrePairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockPrePairing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasBluetoothBlockPrePairing() bool`

HasBluetoothBlockPrePairing returns a boolean if a field has been set.

### SetBluetoothBlockPrePairing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetBluetoothBlockPrePairing(v bool)`

SetBluetoothBlockPrePairing gets a reference to the given bool and assigns it to the BluetoothBlockPrePairing field.

### GetEdgeBlockAutofill

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockAutofill() bool`

GetEdgeBlockAutofill returns the EdgeBlockAutofill field if non-nil, zero value otherwise.

### GetEdgeBlockAutofillOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockAutofillOk() (bool, bool)`

GetEdgeBlockAutofillOk returns a tuple with the EdgeBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockAutofill

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockAutofill() bool`

HasEdgeBlockAutofill returns a boolean if a field has been set.

### SetEdgeBlockAutofill

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockAutofill(v bool)`

SetEdgeBlockAutofill gets a reference to the given bool and assigns it to the EdgeBlockAutofill field.

### GetEdgeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlocked() bool`

GetEdgeBlocked returns the EdgeBlocked field if non-nil, zero value otherwise.

### GetEdgeBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockedOk() (bool, bool)`

GetEdgeBlockedOk returns a tuple with the EdgeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlocked() bool`

HasEdgeBlocked returns a boolean if a field has been set.

### SetEdgeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlocked(v bool)`

SetEdgeBlocked gets a reference to the given bool and assigns it to the EdgeBlocked field.

### GetEdgeCookiePolicy

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeCookiePolicy() AnyOfmicrosoftGraphEdgeCookiePolicy`

GetEdgeCookiePolicy returns the EdgeCookiePolicy field if non-nil, zero value otherwise.

### GetEdgeCookiePolicyOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeCookiePolicyOk() (AnyOfmicrosoftGraphEdgeCookiePolicy, bool)`

GetEdgeCookiePolicyOk returns a tuple with the EdgeCookiePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeCookiePolicy

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeCookiePolicy() bool`

HasEdgeCookiePolicy returns a boolean if a field has been set.

### SetEdgeCookiePolicy

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeCookiePolicy(v AnyOfmicrosoftGraphEdgeCookiePolicy)`

SetEdgeCookiePolicy gets a reference to the given AnyOfmicrosoftGraphEdgeCookiePolicy and assigns it to the EdgeCookiePolicy field.

### GetEdgeBlockDeveloperTools

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockDeveloperTools() bool`

GetEdgeBlockDeveloperTools returns the EdgeBlockDeveloperTools field if non-nil, zero value otherwise.

### GetEdgeBlockDeveloperToolsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockDeveloperToolsOk() (bool, bool)`

GetEdgeBlockDeveloperToolsOk returns a tuple with the EdgeBlockDeveloperTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockDeveloperTools

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockDeveloperTools() bool`

HasEdgeBlockDeveloperTools returns a boolean if a field has been set.

### SetEdgeBlockDeveloperTools

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockDeveloperTools(v bool)`

SetEdgeBlockDeveloperTools gets a reference to the given bool and assigns it to the EdgeBlockDeveloperTools field.

### GetEdgeBlockSendingDoNotTrackHeader

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockSendingDoNotTrackHeader() bool`

GetEdgeBlockSendingDoNotTrackHeader returns the EdgeBlockSendingDoNotTrackHeader field if non-nil, zero value otherwise.

### GetEdgeBlockSendingDoNotTrackHeaderOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockSendingDoNotTrackHeaderOk() (bool, bool)`

GetEdgeBlockSendingDoNotTrackHeaderOk returns a tuple with the EdgeBlockSendingDoNotTrackHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockSendingDoNotTrackHeader

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockSendingDoNotTrackHeader() bool`

HasEdgeBlockSendingDoNotTrackHeader returns a boolean if a field has been set.

### SetEdgeBlockSendingDoNotTrackHeader

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockSendingDoNotTrackHeader(v bool)`

SetEdgeBlockSendingDoNotTrackHeader gets a reference to the given bool and assigns it to the EdgeBlockSendingDoNotTrackHeader field.

### GetEdgeBlockExtensions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockExtensions() bool`

GetEdgeBlockExtensions returns the EdgeBlockExtensions field if non-nil, zero value otherwise.

### GetEdgeBlockExtensionsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockExtensionsOk() (bool, bool)`

GetEdgeBlockExtensionsOk returns a tuple with the EdgeBlockExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockExtensions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockExtensions() bool`

HasEdgeBlockExtensions returns a boolean if a field has been set.

### SetEdgeBlockExtensions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockExtensions(v bool)`

SetEdgeBlockExtensions gets a reference to the given bool and assigns it to the EdgeBlockExtensions field.

### GetEdgeBlockInPrivateBrowsing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockInPrivateBrowsing() bool`

GetEdgeBlockInPrivateBrowsing returns the EdgeBlockInPrivateBrowsing field if non-nil, zero value otherwise.

### GetEdgeBlockInPrivateBrowsingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockInPrivateBrowsingOk() (bool, bool)`

GetEdgeBlockInPrivateBrowsingOk returns a tuple with the EdgeBlockInPrivateBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockInPrivateBrowsing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockInPrivateBrowsing() bool`

HasEdgeBlockInPrivateBrowsing returns a boolean if a field has been set.

### SetEdgeBlockInPrivateBrowsing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockInPrivateBrowsing(v bool)`

SetEdgeBlockInPrivateBrowsing gets a reference to the given bool and assigns it to the EdgeBlockInPrivateBrowsing field.

### GetEdgeBlockJavaScript

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockJavaScript() bool`

GetEdgeBlockJavaScript returns the EdgeBlockJavaScript field if non-nil, zero value otherwise.

### GetEdgeBlockJavaScriptOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockJavaScriptOk() (bool, bool)`

GetEdgeBlockJavaScriptOk returns a tuple with the EdgeBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockJavaScript

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockJavaScript() bool`

HasEdgeBlockJavaScript returns a boolean if a field has been set.

### SetEdgeBlockJavaScript

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockJavaScript(v bool)`

SetEdgeBlockJavaScript gets a reference to the given bool and assigns it to the EdgeBlockJavaScript field.

### GetEdgeBlockPasswordManager

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockPasswordManager() bool`

GetEdgeBlockPasswordManager returns the EdgeBlockPasswordManager field if non-nil, zero value otherwise.

### GetEdgeBlockPasswordManagerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockPasswordManagerOk() (bool, bool)`

GetEdgeBlockPasswordManagerOk returns a tuple with the EdgeBlockPasswordManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockPasswordManager

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockPasswordManager() bool`

HasEdgeBlockPasswordManager returns a boolean if a field has been set.

### SetEdgeBlockPasswordManager

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockPasswordManager(v bool)`

SetEdgeBlockPasswordManager gets a reference to the given bool and assigns it to the EdgeBlockPasswordManager field.

### GetEdgeBlockAddressBarDropdown

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockAddressBarDropdown() bool`

GetEdgeBlockAddressBarDropdown returns the EdgeBlockAddressBarDropdown field if non-nil, zero value otherwise.

### GetEdgeBlockAddressBarDropdownOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockAddressBarDropdownOk() (bool, bool)`

GetEdgeBlockAddressBarDropdownOk returns a tuple with the EdgeBlockAddressBarDropdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockAddressBarDropdown

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockAddressBarDropdown() bool`

HasEdgeBlockAddressBarDropdown returns a boolean if a field has been set.

### SetEdgeBlockAddressBarDropdown

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockAddressBarDropdown(v bool)`

SetEdgeBlockAddressBarDropdown gets a reference to the given bool and assigns it to the EdgeBlockAddressBarDropdown field.

### GetEdgeBlockCompatibilityList

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockCompatibilityList() bool`

GetEdgeBlockCompatibilityList returns the EdgeBlockCompatibilityList field if non-nil, zero value otherwise.

### GetEdgeBlockCompatibilityListOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockCompatibilityListOk() (bool, bool)`

GetEdgeBlockCompatibilityListOk returns a tuple with the EdgeBlockCompatibilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockCompatibilityList

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockCompatibilityList() bool`

HasEdgeBlockCompatibilityList returns a boolean if a field has been set.

### SetEdgeBlockCompatibilityList

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockCompatibilityList(v bool)`

SetEdgeBlockCompatibilityList gets a reference to the given bool and assigns it to the EdgeBlockCompatibilityList field.

### GetEdgeClearBrowsingDataOnExit

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeClearBrowsingDataOnExit() bool`

GetEdgeClearBrowsingDataOnExit returns the EdgeClearBrowsingDataOnExit field if non-nil, zero value otherwise.

### GetEdgeClearBrowsingDataOnExitOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeClearBrowsingDataOnExitOk() (bool, bool)`

GetEdgeClearBrowsingDataOnExitOk returns a tuple with the EdgeClearBrowsingDataOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeClearBrowsingDataOnExit

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeClearBrowsingDataOnExit() bool`

HasEdgeClearBrowsingDataOnExit returns a boolean if a field has been set.

### SetEdgeClearBrowsingDataOnExit

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeClearBrowsingDataOnExit(v bool)`

SetEdgeClearBrowsingDataOnExit gets a reference to the given bool and assigns it to the EdgeClearBrowsingDataOnExit field.

### GetEdgeAllowStartPagesModification

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeAllowStartPagesModification() bool`

GetEdgeAllowStartPagesModification returns the EdgeAllowStartPagesModification field if non-nil, zero value otherwise.

### GetEdgeAllowStartPagesModificationOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeAllowStartPagesModificationOk() (bool, bool)`

GetEdgeAllowStartPagesModificationOk returns a tuple with the EdgeAllowStartPagesModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeAllowStartPagesModification

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeAllowStartPagesModification() bool`

HasEdgeAllowStartPagesModification returns a boolean if a field has been set.

### SetEdgeAllowStartPagesModification

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeAllowStartPagesModification(v bool)`

SetEdgeAllowStartPagesModification gets a reference to the given bool and assigns it to the EdgeAllowStartPagesModification field.

### GetEdgeDisableFirstRunPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeDisableFirstRunPage() bool`

GetEdgeDisableFirstRunPage returns the EdgeDisableFirstRunPage field if non-nil, zero value otherwise.

### GetEdgeDisableFirstRunPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeDisableFirstRunPageOk() (bool, bool)`

GetEdgeDisableFirstRunPageOk returns a tuple with the EdgeDisableFirstRunPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeDisableFirstRunPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeDisableFirstRunPage() bool`

HasEdgeDisableFirstRunPage returns a boolean if a field has been set.

### SetEdgeDisableFirstRunPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeDisableFirstRunPage(v bool)`

SetEdgeDisableFirstRunPage gets a reference to the given bool and assigns it to the EdgeDisableFirstRunPage field.

### GetEdgeBlockLiveTileDataCollection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockLiveTileDataCollection() bool`

GetEdgeBlockLiveTileDataCollection returns the EdgeBlockLiveTileDataCollection field if non-nil, zero value otherwise.

### GetEdgeBlockLiveTileDataCollectionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockLiveTileDataCollectionOk() (bool, bool)`

GetEdgeBlockLiveTileDataCollectionOk returns a tuple with the EdgeBlockLiveTileDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockLiveTileDataCollection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockLiveTileDataCollection() bool`

HasEdgeBlockLiveTileDataCollection returns a boolean if a field has been set.

### SetEdgeBlockLiveTileDataCollection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockLiveTileDataCollection(v bool)`

SetEdgeBlockLiveTileDataCollection gets a reference to the given bool and assigns it to the EdgeBlockLiveTileDataCollection field.

### GetEdgeSyncFavoritesWithInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeSyncFavoritesWithInternetExplorer() bool`

GetEdgeSyncFavoritesWithInternetExplorer returns the EdgeSyncFavoritesWithInternetExplorer field if non-nil, zero value otherwise.

### GetEdgeSyncFavoritesWithInternetExplorerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeSyncFavoritesWithInternetExplorerOk() (bool, bool)`

GetEdgeSyncFavoritesWithInternetExplorerOk returns a tuple with the EdgeSyncFavoritesWithInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeSyncFavoritesWithInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeSyncFavoritesWithInternetExplorer() bool`

HasEdgeSyncFavoritesWithInternetExplorer returns a boolean if a field has been set.

### SetEdgeSyncFavoritesWithInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeSyncFavoritesWithInternetExplorer(v bool)`

SetEdgeSyncFavoritesWithInternetExplorer gets a reference to the given bool and assigns it to the EdgeSyncFavoritesWithInternetExplorer field.

### GetCellularBlockDataWhenRoaming

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCellularBlockDataWhenRoaming() bool`

GetCellularBlockDataWhenRoaming returns the CellularBlockDataWhenRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataWhenRoamingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCellularBlockDataWhenRoamingOk() (bool, bool)`

GetCellularBlockDataWhenRoamingOk returns a tuple with the CellularBlockDataWhenRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataWhenRoaming

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCellularBlockDataWhenRoaming() bool`

HasCellularBlockDataWhenRoaming returns a boolean if a field has been set.

### SetCellularBlockDataWhenRoaming

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCellularBlockDataWhenRoaming(v bool)`

SetCellularBlockDataWhenRoaming gets a reference to the given bool and assigns it to the CellularBlockDataWhenRoaming field.

### GetCellularBlockVpn

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCellularBlockVpn() bool`

GetCellularBlockVpn returns the CellularBlockVpn field if non-nil, zero value otherwise.

### GetCellularBlockVpnOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCellularBlockVpnOk() (bool, bool)`

GetCellularBlockVpnOk returns a tuple with the CellularBlockVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVpn

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCellularBlockVpn() bool`

HasCellularBlockVpn returns a boolean if a field has been set.

### SetCellularBlockVpn

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCellularBlockVpn(v bool)`

SetCellularBlockVpn gets a reference to the given bool and assigns it to the CellularBlockVpn field.

### GetCellularBlockVpnWhenRoaming

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCellularBlockVpnWhenRoaming() bool`

GetCellularBlockVpnWhenRoaming returns the CellularBlockVpnWhenRoaming field if non-nil, zero value otherwise.

### GetCellularBlockVpnWhenRoamingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCellularBlockVpnWhenRoamingOk() (bool, bool)`

GetCellularBlockVpnWhenRoamingOk returns a tuple with the CellularBlockVpnWhenRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVpnWhenRoaming

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCellularBlockVpnWhenRoaming() bool`

HasCellularBlockVpnWhenRoaming returns a boolean if a field has been set.

### SetCellularBlockVpnWhenRoaming

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCellularBlockVpnWhenRoaming(v bool)`

SetCellularBlockVpnWhenRoaming gets a reference to the given bool and assigns it to the CellularBlockVpnWhenRoaming field.

### GetDefenderBlockEndUserAccess

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderBlockEndUserAccess() bool`

GetDefenderBlockEndUserAccess returns the DefenderBlockEndUserAccess field if non-nil, zero value otherwise.

### GetDefenderBlockEndUserAccessOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderBlockEndUserAccessOk() (bool, bool)`

GetDefenderBlockEndUserAccessOk returns a tuple with the DefenderBlockEndUserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderBlockEndUserAccess

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderBlockEndUserAccess() bool`

HasDefenderBlockEndUserAccess returns a boolean if a field has been set.

### SetDefenderBlockEndUserAccess

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderBlockEndUserAccess(v bool)`

SetDefenderBlockEndUserAccess gets a reference to the given bool and assigns it to the DefenderBlockEndUserAccess field.

### GetDefenderDaysBeforeDeletingQuarantinedMalware

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalware() int32`

GetDefenderDaysBeforeDeletingQuarantinedMalware returns the DefenderDaysBeforeDeletingQuarantinedMalware field if non-nil, zero value otherwise.

### GetDefenderDaysBeforeDeletingQuarantinedMalwareOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalwareOk() (int32, bool)`

GetDefenderDaysBeforeDeletingQuarantinedMalwareOk returns a tuple with the DefenderDaysBeforeDeletingQuarantinedMalware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderDaysBeforeDeletingQuarantinedMalware

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderDaysBeforeDeletingQuarantinedMalware() bool`

HasDefenderDaysBeforeDeletingQuarantinedMalware returns a boolean if a field has been set.

### SetDefenderDaysBeforeDeletingQuarantinedMalware

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalware(v int32)`

SetDefenderDaysBeforeDeletingQuarantinedMalware gets a reference to the given int32 and assigns it to the DefenderDaysBeforeDeletingQuarantinedMalware field.

### SetDefenderDaysBeforeDeletingQuarantinedMalwareExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalwareExplicitNull(b bool)`

SetDefenderDaysBeforeDeletingQuarantinedMalwareExplicitNull (un)sets DefenderDaysBeforeDeletingQuarantinedMalware to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderDaysBeforeDeletingQuarantinedMalware value is set to nil even if false is passed
### GetDefenderDetectedMalwareActions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderDetectedMalwareActions() AnyOfmicrosoftGraphDefenderDetectedMalwareActions`

GetDefenderDetectedMalwareActions returns the DefenderDetectedMalwareActions field if non-nil, zero value otherwise.

### GetDefenderDetectedMalwareActionsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderDetectedMalwareActionsOk() (AnyOfmicrosoftGraphDefenderDetectedMalwareActions, bool)`

GetDefenderDetectedMalwareActionsOk returns a tuple with the DefenderDetectedMalwareActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderDetectedMalwareActions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderDetectedMalwareActions() bool`

HasDefenderDetectedMalwareActions returns a boolean if a field has been set.

### SetDefenderDetectedMalwareActions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderDetectedMalwareActions(v AnyOfmicrosoftGraphDefenderDetectedMalwareActions)`

SetDefenderDetectedMalwareActions gets a reference to the given AnyOfmicrosoftGraphDefenderDetectedMalwareActions and assigns it to the DefenderDetectedMalwareActions field.

### SetDefenderDetectedMalwareActionsExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderDetectedMalwareActionsExplicitNull(b bool)`

SetDefenderDetectedMalwareActionsExplicitNull (un)sets DefenderDetectedMalwareActions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderDetectedMalwareActions value is set to nil even if false is passed
### GetDefenderSystemScanSchedule

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderSystemScanSchedule() AnyOfmicrosoftGraphWeeklySchedule`

GetDefenderSystemScanSchedule returns the DefenderSystemScanSchedule field if non-nil, zero value otherwise.

### GetDefenderSystemScanScheduleOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderSystemScanScheduleOk() (AnyOfmicrosoftGraphWeeklySchedule, bool)`

GetDefenderSystemScanScheduleOk returns a tuple with the DefenderSystemScanSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderSystemScanSchedule

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderSystemScanSchedule() bool`

HasDefenderSystemScanSchedule returns a boolean if a field has been set.

### SetDefenderSystemScanSchedule

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderSystemScanSchedule(v AnyOfmicrosoftGraphWeeklySchedule)`

SetDefenderSystemScanSchedule gets a reference to the given AnyOfmicrosoftGraphWeeklySchedule and assigns it to the DefenderSystemScanSchedule field.

### GetDefenderFilesAndFoldersToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderFilesAndFoldersToExclude() []string`

GetDefenderFilesAndFoldersToExclude returns the DefenderFilesAndFoldersToExclude field if non-nil, zero value otherwise.

### GetDefenderFilesAndFoldersToExcludeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderFilesAndFoldersToExcludeOk() ([]string, bool)`

GetDefenderFilesAndFoldersToExcludeOk returns a tuple with the DefenderFilesAndFoldersToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderFilesAndFoldersToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderFilesAndFoldersToExclude() bool`

HasDefenderFilesAndFoldersToExclude returns a boolean if a field has been set.

### SetDefenderFilesAndFoldersToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderFilesAndFoldersToExclude(v []string)`

SetDefenderFilesAndFoldersToExclude gets a reference to the given []string and assigns it to the DefenderFilesAndFoldersToExclude field.

### GetDefenderFileExtensionsToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderFileExtensionsToExclude() []string`

GetDefenderFileExtensionsToExclude returns the DefenderFileExtensionsToExclude field if non-nil, zero value otherwise.

### GetDefenderFileExtensionsToExcludeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderFileExtensionsToExcludeOk() ([]string, bool)`

GetDefenderFileExtensionsToExcludeOk returns a tuple with the DefenderFileExtensionsToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderFileExtensionsToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderFileExtensionsToExclude() bool`

HasDefenderFileExtensionsToExclude returns a boolean if a field has been set.

### SetDefenderFileExtensionsToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderFileExtensionsToExclude(v []string)`

SetDefenderFileExtensionsToExclude gets a reference to the given []string and assigns it to the DefenderFileExtensionsToExclude field.

### GetDefenderScanMaxCpu

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanMaxCpu() int32`

GetDefenderScanMaxCpu returns the DefenderScanMaxCpu field if non-nil, zero value otherwise.

### GetDefenderScanMaxCpuOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanMaxCpuOk() (int32, bool)`

GetDefenderScanMaxCpuOk returns a tuple with the DefenderScanMaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanMaxCpu

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanMaxCpu() bool`

HasDefenderScanMaxCpu returns a boolean if a field has been set.

### SetDefenderScanMaxCpu

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanMaxCpu(v int32)`

SetDefenderScanMaxCpu gets a reference to the given int32 and assigns it to the DefenderScanMaxCpu field.

### SetDefenderScanMaxCpuExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanMaxCpuExplicitNull(b bool)`

SetDefenderScanMaxCpuExplicitNull (un)sets DefenderScanMaxCpu to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderScanMaxCpu value is set to nil even if false is passed
### GetDefenderMonitorFileActivity

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderMonitorFileActivity() AnyOfmicrosoftGraphDefenderMonitorFileActivity`

GetDefenderMonitorFileActivity returns the DefenderMonitorFileActivity field if non-nil, zero value otherwise.

### GetDefenderMonitorFileActivityOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderMonitorFileActivityOk() (AnyOfmicrosoftGraphDefenderMonitorFileActivity, bool)`

GetDefenderMonitorFileActivityOk returns a tuple with the DefenderMonitorFileActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderMonitorFileActivity

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderMonitorFileActivity() bool`

HasDefenderMonitorFileActivity returns a boolean if a field has been set.

### SetDefenderMonitorFileActivity

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderMonitorFileActivity(v AnyOfmicrosoftGraphDefenderMonitorFileActivity)`

SetDefenderMonitorFileActivity gets a reference to the given AnyOfmicrosoftGraphDefenderMonitorFileActivity and assigns it to the DefenderMonitorFileActivity field.

### GetDefenderProcessesToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderProcessesToExclude() []string`

GetDefenderProcessesToExclude returns the DefenderProcessesToExclude field if non-nil, zero value otherwise.

### GetDefenderProcessesToExcludeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderProcessesToExcludeOk() ([]string, bool)`

GetDefenderProcessesToExcludeOk returns a tuple with the DefenderProcessesToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderProcessesToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderProcessesToExclude() bool`

HasDefenderProcessesToExclude returns a boolean if a field has been set.

### SetDefenderProcessesToExclude

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderProcessesToExclude(v []string)`

SetDefenderProcessesToExclude gets a reference to the given []string and assigns it to the DefenderProcessesToExclude field.

### GetDefenderPromptForSampleSubmission

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderPromptForSampleSubmission() AnyOfmicrosoftGraphDefenderPromptForSampleSubmission`

GetDefenderPromptForSampleSubmission returns the DefenderPromptForSampleSubmission field if non-nil, zero value otherwise.

### GetDefenderPromptForSampleSubmissionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderPromptForSampleSubmissionOk() (AnyOfmicrosoftGraphDefenderPromptForSampleSubmission, bool)`

GetDefenderPromptForSampleSubmissionOk returns a tuple with the DefenderPromptForSampleSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderPromptForSampleSubmission

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderPromptForSampleSubmission() bool`

HasDefenderPromptForSampleSubmission returns a boolean if a field has been set.

### SetDefenderPromptForSampleSubmission

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderPromptForSampleSubmission(v AnyOfmicrosoftGraphDefenderPromptForSampleSubmission)`

SetDefenderPromptForSampleSubmission gets a reference to the given AnyOfmicrosoftGraphDefenderPromptForSampleSubmission and assigns it to the DefenderPromptForSampleSubmission field.

### GetDefenderRequireBehaviorMonitoring

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireBehaviorMonitoring() bool`

GetDefenderRequireBehaviorMonitoring returns the DefenderRequireBehaviorMonitoring field if non-nil, zero value otherwise.

### GetDefenderRequireBehaviorMonitoringOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireBehaviorMonitoringOk() (bool, bool)`

GetDefenderRequireBehaviorMonitoringOk returns a tuple with the DefenderRequireBehaviorMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireBehaviorMonitoring

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderRequireBehaviorMonitoring() bool`

HasDefenderRequireBehaviorMonitoring returns a boolean if a field has been set.

### SetDefenderRequireBehaviorMonitoring

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderRequireBehaviorMonitoring(v bool)`

SetDefenderRequireBehaviorMonitoring gets a reference to the given bool and assigns it to the DefenderRequireBehaviorMonitoring field.

### GetDefenderRequireCloudProtection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireCloudProtection() bool`

GetDefenderRequireCloudProtection returns the DefenderRequireCloudProtection field if non-nil, zero value otherwise.

### GetDefenderRequireCloudProtectionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireCloudProtectionOk() (bool, bool)`

GetDefenderRequireCloudProtectionOk returns a tuple with the DefenderRequireCloudProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireCloudProtection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderRequireCloudProtection() bool`

HasDefenderRequireCloudProtection returns a boolean if a field has been set.

### SetDefenderRequireCloudProtection

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderRequireCloudProtection(v bool)`

SetDefenderRequireCloudProtection gets a reference to the given bool and assigns it to the DefenderRequireCloudProtection field.

### GetDefenderRequireNetworkInspectionSystem

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireNetworkInspectionSystem() bool`

GetDefenderRequireNetworkInspectionSystem returns the DefenderRequireNetworkInspectionSystem field if non-nil, zero value otherwise.

### GetDefenderRequireNetworkInspectionSystemOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireNetworkInspectionSystemOk() (bool, bool)`

GetDefenderRequireNetworkInspectionSystemOk returns a tuple with the DefenderRequireNetworkInspectionSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireNetworkInspectionSystem

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderRequireNetworkInspectionSystem() bool`

HasDefenderRequireNetworkInspectionSystem returns a boolean if a field has been set.

### SetDefenderRequireNetworkInspectionSystem

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderRequireNetworkInspectionSystem(v bool)`

SetDefenderRequireNetworkInspectionSystem gets a reference to the given bool and assigns it to the DefenderRequireNetworkInspectionSystem field.

### GetDefenderRequireRealTimeMonitoring

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireRealTimeMonitoring() bool`

GetDefenderRequireRealTimeMonitoring returns the DefenderRequireRealTimeMonitoring field if non-nil, zero value otherwise.

### GetDefenderRequireRealTimeMonitoringOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderRequireRealTimeMonitoringOk() (bool, bool)`

GetDefenderRequireRealTimeMonitoringOk returns a tuple with the DefenderRequireRealTimeMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireRealTimeMonitoring

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderRequireRealTimeMonitoring() bool`

HasDefenderRequireRealTimeMonitoring returns a boolean if a field has been set.

### SetDefenderRequireRealTimeMonitoring

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderRequireRealTimeMonitoring(v bool)`

SetDefenderRequireRealTimeMonitoring gets a reference to the given bool and assigns it to the DefenderRequireRealTimeMonitoring field.

### GetDefenderScanArchiveFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanArchiveFiles() bool`

GetDefenderScanArchiveFiles returns the DefenderScanArchiveFiles field if non-nil, zero value otherwise.

### GetDefenderScanArchiveFilesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanArchiveFilesOk() (bool, bool)`

GetDefenderScanArchiveFilesOk returns a tuple with the DefenderScanArchiveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanArchiveFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanArchiveFiles() bool`

HasDefenderScanArchiveFiles returns a boolean if a field has been set.

### SetDefenderScanArchiveFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanArchiveFiles(v bool)`

SetDefenderScanArchiveFiles gets a reference to the given bool and assigns it to the DefenderScanArchiveFiles field.

### GetDefenderScanDownloads

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanDownloads() bool`

GetDefenderScanDownloads returns the DefenderScanDownloads field if non-nil, zero value otherwise.

### GetDefenderScanDownloadsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanDownloadsOk() (bool, bool)`

GetDefenderScanDownloadsOk returns a tuple with the DefenderScanDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanDownloads

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanDownloads() bool`

HasDefenderScanDownloads returns a boolean if a field has been set.

### SetDefenderScanDownloads

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanDownloads(v bool)`

SetDefenderScanDownloads gets a reference to the given bool and assigns it to the DefenderScanDownloads field.

### GetDefenderScanNetworkFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanNetworkFiles() bool`

GetDefenderScanNetworkFiles returns the DefenderScanNetworkFiles field if non-nil, zero value otherwise.

### GetDefenderScanNetworkFilesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanNetworkFilesOk() (bool, bool)`

GetDefenderScanNetworkFilesOk returns a tuple with the DefenderScanNetworkFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanNetworkFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanNetworkFiles() bool`

HasDefenderScanNetworkFiles returns a boolean if a field has been set.

### SetDefenderScanNetworkFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanNetworkFiles(v bool)`

SetDefenderScanNetworkFiles gets a reference to the given bool and assigns it to the DefenderScanNetworkFiles field.

### GetDefenderScanIncomingMail

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanIncomingMail() bool`

GetDefenderScanIncomingMail returns the DefenderScanIncomingMail field if non-nil, zero value otherwise.

### GetDefenderScanIncomingMailOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanIncomingMailOk() (bool, bool)`

GetDefenderScanIncomingMailOk returns a tuple with the DefenderScanIncomingMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanIncomingMail

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanIncomingMail() bool`

HasDefenderScanIncomingMail returns a boolean if a field has been set.

### SetDefenderScanIncomingMail

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanIncomingMail(v bool)`

SetDefenderScanIncomingMail gets a reference to the given bool and assigns it to the DefenderScanIncomingMail field.

### GetDefenderScanMappedNetworkDrivesDuringFullScan

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanMappedNetworkDrivesDuringFullScan() bool`

GetDefenderScanMappedNetworkDrivesDuringFullScan returns the DefenderScanMappedNetworkDrivesDuringFullScan field if non-nil, zero value otherwise.

### GetDefenderScanMappedNetworkDrivesDuringFullScanOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanMappedNetworkDrivesDuringFullScanOk() (bool, bool)`

GetDefenderScanMappedNetworkDrivesDuringFullScanOk returns a tuple with the DefenderScanMappedNetworkDrivesDuringFullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanMappedNetworkDrivesDuringFullScan

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanMappedNetworkDrivesDuringFullScan() bool`

HasDefenderScanMappedNetworkDrivesDuringFullScan returns a boolean if a field has been set.

### SetDefenderScanMappedNetworkDrivesDuringFullScan

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanMappedNetworkDrivesDuringFullScan(v bool)`

SetDefenderScanMappedNetworkDrivesDuringFullScan gets a reference to the given bool and assigns it to the DefenderScanMappedNetworkDrivesDuringFullScan field.

### GetDefenderScanRemovableDrivesDuringFullScan

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanRemovableDrivesDuringFullScan() bool`

GetDefenderScanRemovableDrivesDuringFullScan returns the DefenderScanRemovableDrivesDuringFullScan field if non-nil, zero value otherwise.

### GetDefenderScanRemovableDrivesDuringFullScanOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanRemovableDrivesDuringFullScanOk() (bool, bool)`

GetDefenderScanRemovableDrivesDuringFullScanOk returns a tuple with the DefenderScanRemovableDrivesDuringFullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanRemovableDrivesDuringFullScan

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanRemovableDrivesDuringFullScan() bool`

HasDefenderScanRemovableDrivesDuringFullScan returns a boolean if a field has been set.

### SetDefenderScanRemovableDrivesDuringFullScan

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanRemovableDrivesDuringFullScan(v bool)`

SetDefenderScanRemovableDrivesDuringFullScan gets a reference to the given bool and assigns it to the DefenderScanRemovableDrivesDuringFullScan field.

### GetDefenderScanScriptsLoadedInInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanScriptsLoadedInInternetExplorer() bool`

GetDefenderScanScriptsLoadedInInternetExplorer returns the DefenderScanScriptsLoadedInInternetExplorer field if non-nil, zero value otherwise.

### GetDefenderScanScriptsLoadedInInternetExplorerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanScriptsLoadedInInternetExplorerOk() (bool, bool)`

GetDefenderScanScriptsLoadedInInternetExplorerOk returns a tuple with the DefenderScanScriptsLoadedInInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanScriptsLoadedInInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanScriptsLoadedInInternetExplorer() bool`

HasDefenderScanScriptsLoadedInInternetExplorer returns a boolean if a field has been set.

### SetDefenderScanScriptsLoadedInInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanScriptsLoadedInInternetExplorer(v bool)`

SetDefenderScanScriptsLoadedInInternetExplorer gets a reference to the given bool and assigns it to the DefenderScanScriptsLoadedInInternetExplorer field.

### GetDefenderSignatureUpdateIntervalInHours

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderSignatureUpdateIntervalInHours() int32`

GetDefenderSignatureUpdateIntervalInHours returns the DefenderSignatureUpdateIntervalInHours field if non-nil, zero value otherwise.

### GetDefenderSignatureUpdateIntervalInHoursOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderSignatureUpdateIntervalInHoursOk() (int32, bool)`

GetDefenderSignatureUpdateIntervalInHoursOk returns a tuple with the DefenderSignatureUpdateIntervalInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderSignatureUpdateIntervalInHours

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderSignatureUpdateIntervalInHours() bool`

HasDefenderSignatureUpdateIntervalInHours returns a boolean if a field has been set.

### SetDefenderSignatureUpdateIntervalInHours

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderSignatureUpdateIntervalInHours(v int32)`

SetDefenderSignatureUpdateIntervalInHours gets a reference to the given int32 and assigns it to the DefenderSignatureUpdateIntervalInHours field.

### SetDefenderSignatureUpdateIntervalInHoursExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderSignatureUpdateIntervalInHoursExplicitNull(b bool)`

SetDefenderSignatureUpdateIntervalInHoursExplicitNull (un)sets DefenderSignatureUpdateIntervalInHours to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderSignatureUpdateIntervalInHours value is set to nil even if false is passed
### GetDefenderScanType

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanType() AnyOfmicrosoftGraphDefenderScanType`

GetDefenderScanType returns the DefenderScanType field if non-nil, zero value otherwise.

### GetDefenderScanTypeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScanTypeOk() (AnyOfmicrosoftGraphDefenderScanType, bool)`

GetDefenderScanTypeOk returns a tuple with the DefenderScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanType

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScanType() bool`

HasDefenderScanType returns a boolean if a field has been set.

### SetDefenderScanType

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScanType(v AnyOfmicrosoftGraphDefenderScanType)`

SetDefenderScanType gets a reference to the given AnyOfmicrosoftGraphDefenderScanType and assigns it to the DefenderScanType field.

### GetDefenderScheduledScanTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScheduledScanTime() string`

GetDefenderScheduledScanTime returns the DefenderScheduledScanTime field if non-nil, zero value otherwise.

### GetDefenderScheduledScanTimeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScheduledScanTimeOk() (string, bool)`

GetDefenderScheduledScanTimeOk returns a tuple with the DefenderScheduledScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScheduledScanTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScheduledScanTime() bool`

HasDefenderScheduledScanTime returns a boolean if a field has been set.

### SetDefenderScheduledScanTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScheduledScanTime(v string)`

SetDefenderScheduledScanTime gets a reference to the given string and assigns it to the DefenderScheduledScanTime field.

### SetDefenderScheduledScanTimeExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScheduledScanTimeExplicitNull(b bool)`

SetDefenderScheduledScanTimeExplicitNull (un)sets DefenderScheduledScanTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderScheduledScanTime value is set to nil even if false is passed
### GetDefenderScheduledQuickScanTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScheduledQuickScanTime() string`

GetDefenderScheduledQuickScanTime returns the DefenderScheduledQuickScanTime field if non-nil, zero value otherwise.

### GetDefenderScheduledQuickScanTimeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderScheduledQuickScanTimeOk() (string, bool)`

GetDefenderScheduledQuickScanTimeOk returns a tuple with the DefenderScheduledQuickScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScheduledQuickScanTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderScheduledQuickScanTime() bool`

HasDefenderScheduledQuickScanTime returns a boolean if a field has been set.

### SetDefenderScheduledQuickScanTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScheduledQuickScanTime(v string)`

SetDefenderScheduledQuickScanTime gets a reference to the given string and assigns it to the DefenderScheduledQuickScanTime field.

### SetDefenderScheduledQuickScanTimeExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderScheduledQuickScanTimeExplicitNull(b bool)`

SetDefenderScheduledQuickScanTimeExplicitNull (un)sets DefenderScheduledQuickScanTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderScheduledQuickScanTime value is set to nil even if false is passed
### GetDefenderCloudBlockLevel

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderCloudBlockLevel() AnyOfmicrosoftGraphDefenderCloudBlockLevelType`

GetDefenderCloudBlockLevel returns the DefenderCloudBlockLevel field if non-nil, zero value otherwise.

### GetDefenderCloudBlockLevelOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDefenderCloudBlockLevelOk() (AnyOfmicrosoftGraphDefenderCloudBlockLevelType, bool)`

GetDefenderCloudBlockLevelOk returns a tuple with the DefenderCloudBlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderCloudBlockLevel

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDefenderCloudBlockLevel() bool`

HasDefenderCloudBlockLevel returns a boolean if a field has been set.

### SetDefenderCloudBlockLevel

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDefenderCloudBlockLevel(v AnyOfmicrosoftGraphDefenderCloudBlockLevelType)`

SetDefenderCloudBlockLevel gets a reference to the given AnyOfmicrosoftGraphDefenderCloudBlockLevelType and assigns it to the DefenderCloudBlockLevel field.

### GetLockScreenAllowTimeoutConfiguration

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenAllowTimeoutConfiguration() bool`

GetLockScreenAllowTimeoutConfiguration returns the LockScreenAllowTimeoutConfiguration field if non-nil, zero value otherwise.

### GetLockScreenAllowTimeoutConfigurationOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenAllowTimeoutConfigurationOk() (bool, bool)`

GetLockScreenAllowTimeoutConfigurationOk returns a tuple with the LockScreenAllowTimeoutConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenAllowTimeoutConfiguration

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLockScreenAllowTimeoutConfiguration() bool`

HasLockScreenAllowTimeoutConfiguration returns a boolean if a field has been set.

### SetLockScreenAllowTimeoutConfiguration

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLockScreenAllowTimeoutConfiguration(v bool)`

SetLockScreenAllowTimeoutConfiguration gets a reference to the given bool and assigns it to the LockScreenAllowTimeoutConfiguration field.

### GetLockScreenBlockActionCenterNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenBlockActionCenterNotifications() bool`

GetLockScreenBlockActionCenterNotifications returns the LockScreenBlockActionCenterNotifications field if non-nil, zero value otherwise.

### GetLockScreenBlockActionCenterNotificationsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenBlockActionCenterNotificationsOk() (bool, bool)`

GetLockScreenBlockActionCenterNotificationsOk returns a tuple with the LockScreenBlockActionCenterNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockActionCenterNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLockScreenBlockActionCenterNotifications() bool`

HasLockScreenBlockActionCenterNotifications returns a boolean if a field has been set.

### SetLockScreenBlockActionCenterNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLockScreenBlockActionCenterNotifications(v bool)`

SetLockScreenBlockActionCenterNotifications gets a reference to the given bool and assigns it to the LockScreenBlockActionCenterNotifications field.

### GetLockScreenBlockCortana

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenBlockCortana() bool`

GetLockScreenBlockCortana returns the LockScreenBlockCortana field if non-nil, zero value otherwise.

### GetLockScreenBlockCortanaOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenBlockCortanaOk() (bool, bool)`

GetLockScreenBlockCortanaOk returns a tuple with the LockScreenBlockCortana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockCortana

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLockScreenBlockCortana() bool`

HasLockScreenBlockCortana returns a boolean if a field has been set.

### SetLockScreenBlockCortana

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLockScreenBlockCortana(v bool)`

SetLockScreenBlockCortana gets a reference to the given bool and assigns it to the LockScreenBlockCortana field.

### GetLockScreenBlockToastNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenBlockToastNotifications() bool`

GetLockScreenBlockToastNotifications returns the LockScreenBlockToastNotifications field if non-nil, zero value otherwise.

### GetLockScreenBlockToastNotificationsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenBlockToastNotificationsOk() (bool, bool)`

GetLockScreenBlockToastNotificationsOk returns a tuple with the LockScreenBlockToastNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockToastNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLockScreenBlockToastNotifications() bool`

HasLockScreenBlockToastNotifications returns a boolean if a field has been set.

### SetLockScreenBlockToastNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLockScreenBlockToastNotifications(v bool)`

SetLockScreenBlockToastNotifications gets a reference to the given bool and assigns it to the LockScreenBlockToastNotifications field.

### GetLockScreenTimeoutInSeconds

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenTimeoutInSeconds() int32`

GetLockScreenTimeoutInSeconds returns the LockScreenTimeoutInSeconds field if non-nil, zero value otherwise.

### GetLockScreenTimeoutInSecondsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLockScreenTimeoutInSecondsOk() (int32, bool)`

GetLockScreenTimeoutInSecondsOk returns a tuple with the LockScreenTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenTimeoutInSeconds

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLockScreenTimeoutInSeconds() bool`

HasLockScreenTimeoutInSeconds returns a boolean if a field has been set.

### SetLockScreenTimeoutInSeconds

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLockScreenTimeoutInSeconds(v int32)`

SetLockScreenTimeoutInSeconds gets a reference to the given int32 and assigns it to the LockScreenTimeoutInSeconds field.

### SetLockScreenTimeoutInSecondsExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLockScreenTimeoutInSecondsExplicitNull(b bool)`

SetLockScreenTimeoutInSecondsExplicitNull (un)sets LockScreenTimeoutInSeconds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LockScreenTimeoutInSeconds value is set to nil even if false is passed
### GetPasswordBlockSimple

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordRequired

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordRequireWhenResumeFromIdleState

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordRequireWhenResumeFromIdleState() bool`

GetPasswordRequireWhenResumeFromIdleState returns the PasswordRequireWhenResumeFromIdleState field if non-nil, zero value otherwise.

### GetPasswordRequireWhenResumeFromIdleStateOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordRequireWhenResumeFromIdleStateOk() (bool, bool)`

GetPasswordRequireWhenResumeFromIdleStateOk returns a tuple with the PasswordRequireWhenResumeFromIdleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequireWhenResumeFromIdleState

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordRequireWhenResumeFromIdleState() bool`

HasPasswordRequireWhenResumeFromIdleState returns a boolean if a field has been set.

### SetPasswordRequireWhenResumeFromIdleState

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordRequireWhenResumeFromIdleState(v bool)`

SetPasswordRequireWhenResumeFromIdleState gets a reference to the given bool and assigns it to the PasswordRequireWhenResumeFromIdleState field.

### GetPasswordRequiredType

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPrivacyAdvertisingId

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPrivacyAdvertisingId() AnyOfmicrosoftGraphStateManagementSetting`

GetPrivacyAdvertisingId returns the PrivacyAdvertisingId field if non-nil, zero value otherwise.

### GetPrivacyAdvertisingIdOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPrivacyAdvertisingIdOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetPrivacyAdvertisingIdOk returns a tuple with the PrivacyAdvertisingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyAdvertisingId

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPrivacyAdvertisingId() bool`

HasPrivacyAdvertisingId returns a boolean if a field has been set.

### SetPrivacyAdvertisingId

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPrivacyAdvertisingId(v AnyOfmicrosoftGraphStateManagementSetting)`

SetPrivacyAdvertisingId gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the PrivacyAdvertisingId field.

### GetPrivacyAutoAcceptPairingAndConsentPrompts

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPrivacyAutoAcceptPairingAndConsentPrompts() bool`

GetPrivacyAutoAcceptPairingAndConsentPrompts returns the PrivacyAutoAcceptPairingAndConsentPrompts field if non-nil, zero value otherwise.

### GetPrivacyAutoAcceptPairingAndConsentPromptsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPrivacyAutoAcceptPairingAndConsentPromptsOk() (bool, bool)`

GetPrivacyAutoAcceptPairingAndConsentPromptsOk returns a tuple with the PrivacyAutoAcceptPairingAndConsentPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyAutoAcceptPairingAndConsentPrompts

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPrivacyAutoAcceptPairingAndConsentPrompts() bool`

HasPrivacyAutoAcceptPairingAndConsentPrompts returns a boolean if a field has been set.

### SetPrivacyAutoAcceptPairingAndConsentPrompts

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPrivacyAutoAcceptPairingAndConsentPrompts(v bool)`

SetPrivacyAutoAcceptPairingAndConsentPrompts gets a reference to the given bool and assigns it to the PrivacyAutoAcceptPairingAndConsentPrompts field.

### GetPrivacyBlockInputPersonalization

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPrivacyBlockInputPersonalization() bool`

GetPrivacyBlockInputPersonalization returns the PrivacyBlockInputPersonalization field if non-nil, zero value otherwise.

### GetPrivacyBlockInputPersonalizationOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetPrivacyBlockInputPersonalizationOk() (bool, bool)`

GetPrivacyBlockInputPersonalizationOk returns a tuple with the PrivacyBlockInputPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyBlockInputPersonalization

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasPrivacyBlockInputPersonalization() bool`

HasPrivacyBlockInputPersonalization returns a boolean if a field has been set.

### SetPrivacyBlockInputPersonalization

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetPrivacyBlockInputPersonalization(v bool)`

SetPrivacyBlockInputPersonalization gets a reference to the given bool and assigns it to the PrivacyBlockInputPersonalization field.

### GetStartBlockUnpinningAppsFromTaskbar

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartBlockUnpinningAppsFromTaskbar() bool`

GetStartBlockUnpinningAppsFromTaskbar returns the StartBlockUnpinningAppsFromTaskbar field if non-nil, zero value otherwise.

### GetStartBlockUnpinningAppsFromTaskbarOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartBlockUnpinningAppsFromTaskbarOk() (bool, bool)`

GetStartBlockUnpinningAppsFromTaskbarOk returns a tuple with the StartBlockUnpinningAppsFromTaskbar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartBlockUnpinningAppsFromTaskbar

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartBlockUnpinningAppsFromTaskbar() bool`

HasStartBlockUnpinningAppsFromTaskbar returns a boolean if a field has been set.

### SetStartBlockUnpinningAppsFromTaskbar

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartBlockUnpinningAppsFromTaskbar(v bool)`

SetStartBlockUnpinningAppsFromTaskbar gets a reference to the given bool and assigns it to the StartBlockUnpinningAppsFromTaskbar field.

### GetStartMenuAppListVisibility

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuAppListVisibility() AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType`

GetStartMenuAppListVisibility returns the StartMenuAppListVisibility field if non-nil, zero value otherwise.

### GetStartMenuAppListVisibilityOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuAppListVisibilityOk() (AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType, bool)`

GetStartMenuAppListVisibilityOk returns a tuple with the StartMenuAppListVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuAppListVisibility

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuAppListVisibility() bool`

HasStartMenuAppListVisibility returns a boolean if a field has been set.

### SetStartMenuAppListVisibility

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuAppListVisibility(v AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType)`

SetStartMenuAppListVisibility gets a reference to the given AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType and assigns it to the StartMenuAppListVisibility field.

### GetStartMenuHideChangeAccountSettings

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideChangeAccountSettings() bool`

GetStartMenuHideChangeAccountSettings returns the StartMenuHideChangeAccountSettings field if non-nil, zero value otherwise.

### GetStartMenuHideChangeAccountSettingsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideChangeAccountSettingsOk() (bool, bool)`

GetStartMenuHideChangeAccountSettingsOk returns a tuple with the StartMenuHideChangeAccountSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideChangeAccountSettings

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideChangeAccountSettings() bool`

HasStartMenuHideChangeAccountSettings returns a boolean if a field has been set.

### SetStartMenuHideChangeAccountSettings

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideChangeAccountSettings(v bool)`

SetStartMenuHideChangeAccountSettings gets a reference to the given bool and assigns it to the StartMenuHideChangeAccountSettings field.

### GetStartMenuHideFrequentlyUsedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideFrequentlyUsedApps() bool`

GetStartMenuHideFrequentlyUsedApps returns the StartMenuHideFrequentlyUsedApps field if non-nil, zero value otherwise.

### GetStartMenuHideFrequentlyUsedAppsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideFrequentlyUsedAppsOk() (bool, bool)`

GetStartMenuHideFrequentlyUsedAppsOk returns a tuple with the StartMenuHideFrequentlyUsedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideFrequentlyUsedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideFrequentlyUsedApps() bool`

HasStartMenuHideFrequentlyUsedApps returns a boolean if a field has been set.

### SetStartMenuHideFrequentlyUsedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideFrequentlyUsedApps(v bool)`

SetStartMenuHideFrequentlyUsedApps gets a reference to the given bool and assigns it to the StartMenuHideFrequentlyUsedApps field.

### GetStartMenuHideHibernate

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideHibernate() bool`

GetStartMenuHideHibernate returns the StartMenuHideHibernate field if non-nil, zero value otherwise.

### GetStartMenuHideHibernateOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideHibernateOk() (bool, bool)`

GetStartMenuHideHibernateOk returns a tuple with the StartMenuHideHibernate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideHibernate

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideHibernate() bool`

HasStartMenuHideHibernate returns a boolean if a field has been set.

### SetStartMenuHideHibernate

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideHibernate(v bool)`

SetStartMenuHideHibernate gets a reference to the given bool and assigns it to the StartMenuHideHibernate field.

### GetStartMenuHideLock

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideLock() bool`

GetStartMenuHideLock returns the StartMenuHideLock field if non-nil, zero value otherwise.

### GetStartMenuHideLockOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideLockOk() (bool, bool)`

GetStartMenuHideLockOk returns a tuple with the StartMenuHideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideLock

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideLock() bool`

HasStartMenuHideLock returns a boolean if a field has been set.

### SetStartMenuHideLock

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideLock(v bool)`

SetStartMenuHideLock gets a reference to the given bool and assigns it to the StartMenuHideLock field.

### GetStartMenuHidePowerButton

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHidePowerButton() bool`

GetStartMenuHidePowerButton returns the StartMenuHidePowerButton field if non-nil, zero value otherwise.

### GetStartMenuHidePowerButtonOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHidePowerButtonOk() (bool, bool)`

GetStartMenuHidePowerButtonOk returns a tuple with the StartMenuHidePowerButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHidePowerButton

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHidePowerButton() bool`

HasStartMenuHidePowerButton returns a boolean if a field has been set.

### SetStartMenuHidePowerButton

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHidePowerButton(v bool)`

SetStartMenuHidePowerButton gets a reference to the given bool and assigns it to the StartMenuHidePowerButton field.

### GetStartMenuHideRecentJumpLists

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideRecentJumpLists() bool`

GetStartMenuHideRecentJumpLists returns the StartMenuHideRecentJumpLists field if non-nil, zero value otherwise.

### GetStartMenuHideRecentJumpListsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideRecentJumpListsOk() (bool, bool)`

GetStartMenuHideRecentJumpListsOk returns a tuple with the StartMenuHideRecentJumpLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideRecentJumpLists

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideRecentJumpLists() bool`

HasStartMenuHideRecentJumpLists returns a boolean if a field has been set.

### SetStartMenuHideRecentJumpLists

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideRecentJumpLists(v bool)`

SetStartMenuHideRecentJumpLists gets a reference to the given bool and assigns it to the StartMenuHideRecentJumpLists field.

### GetStartMenuHideRecentlyAddedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideRecentlyAddedApps() bool`

GetStartMenuHideRecentlyAddedApps returns the StartMenuHideRecentlyAddedApps field if non-nil, zero value otherwise.

### GetStartMenuHideRecentlyAddedAppsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideRecentlyAddedAppsOk() (bool, bool)`

GetStartMenuHideRecentlyAddedAppsOk returns a tuple with the StartMenuHideRecentlyAddedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideRecentlyAddedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideRecentlyAddedApps() bool`

HasStartMenuHideRecentlyAddedApps returns a boolean if a field has been set.

### SetStartMenuHideRecentlyAddedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideRecentlyAddedApps(v bool)`

SetStartMenuHideRecentlyAddedApps gets a reference to the given bool and assigns it to the StartMenuHideRecentlyAddedApps field.

### GetStartMenuHideRestartOptions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideRestartOptions() bool`

GetStartMenuHideRestartOptions returns the StartMenuHideRestartOptions field if non-nil, zero value otherwise.

### GetStartMenuHideRestartOptionsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideRestartOptionsOk() (bool, bool)`

GetStartMenuHideRestartOptionsOk returns a tuple with the StartMenuHideRestartOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideRestartOptions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideRestartOptions() bool`

HasStartMenuHideRestartOptions returns a boolean if a field has been set.

### SetStartMenuHideRestartOptions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideRestartOptions(v bool)`

SetStartMenuHideRestartOptions gets a reference to the given bool and assigns it to the StartMenuHideRestartOptions field.

### GetStartMenuHideShutDown

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideShutDown() bool`

GetStartMenuHideShutDown returns the StartMenuHideShutDown field if non-nil, zero value otherwise.

### GetStartMenuHideShutDownOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideShutDownOk() (bool, bool)`

GetStartMenuHideShutDownOk returns a tuple with the StartMenuHideShutDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideShutDown

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideShutDown() bool`

HasStartMenuHideShutDown returns a boolean if a field has been set.

### SetStartMenuHideShutDown

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideShutDown(v bool)`

SetStartMenuHideShutDown gets a reference to the given bool and assigns it to the StartMenuHideShutDown field.

### GetStartMenuHideSignOut

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideSignOut() bool`

GetStartMenuHideSignOut returns the StartMenuHideSignOut field if non-nil, zero value otherwise.

### GetStartMenuHideSignOutOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideSignOutOk() (bool, bool)`

GetStartMenuHideSignOutOk returns a tuple with the StartMenuHideSignOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideSignOut

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideSignOut() bool`

HasStartMenuHideSignOut returns a boolean if a field has been set.

### SetStartMenuHideSignOut

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideSignOut(v bool)`

SetStartMenuHideSignOut gets a reference to the given bool and assigns it to the StartMenuHideSignOut field.

### GetStartMenuHideSleep

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideSleep() bool`

GetStartMenuHideSleep returns the StartMenuHideSleep field if non-nil, zero value otherwise.

### GetStartMenuHideSleepOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideSleepOk() (bool, bool)`

GetStartMenuHideSleepOk returns a tuple with the StartMenuHideSleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideSleep

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideSleep() bool`

HasStartMenuHideSleep returns a boolean if a field has been set.

### SetStartMenuHideSleep

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideSleep(v bool)`

SetStartMenuHideSleep gets a reference to the given bool and assigns it to the StartMenuHideSleep field.

### GetStartMenuHideSwitchAccount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideSwitchAccount() bool`

GetStartMenuHideSwitchAccount returns the StartMenuHideSwitchAccount field if non-nil, zero value otherwise.

### GetStartMenuHideSwitchAccountOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideSwitchAccountOk() (bool, bool)`

GetStartMenuHideSwitchAccountOk returns a tuple with the StartMenuHideSwitchAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideSwitchAccount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideSwitchAccount() bool`

HasStartMenuHideSwitchAccount returns a boolean if a field has been set.

### SetStartMenuHideSwitchAccount

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideSwitchAccount(v bool)`

SetStartMenuHideSwitchAccount gets a reference to the given bool and assigns it to the StartMenuHideSwitchAccount field.

### GetStartMenuHideUserTile

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideUserTile() bool`

GetStartMenuHideUserTile returns the StartMenuHideUserTile field if non-nil, zero value otherwise.

### GetStartMenuHideUserTileOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuHideUserTileOk() (bool, bool)`

GetStartMenuHideUserTileOk returns a tuple with the StartMenuHideUserTile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideUserTile

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuHideUserTile() bool`

HasStartMenuHideUserTile returns a boolean if a field has been set.

### SetStartMenuHideUserTile

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuHideUserTile(v bool)`

SetStartMenuHideUserTile gets a reference to the given bool and assigns it to the StartMenuHideUserTile field.

### GetStartMenuLayoutEdgeAssetsXml

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuLayoutEdgeAssetsXml() string`

GetStartMenuLayoutEdgeAssetsXml returns the StartMenuLayoutEdgeAssetsXml field if non-nil, zero value otherwise.

### GetStartMenuLayoutEdgeAssetsXmlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuLayoutEdgeAssetsXmlOk() (string, bool)`

GetStartMenuLayoutEdgeAssetsXmlOk returns a tuple with the StartMenuLayoutEdgeAssetsXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuLayoutEdgeAssetsXml

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuLayoutEdgeAssetsXml() bool`

HasStartMenuLayoutEdgeAssetsXml returns a boolean if a field has been set.

### SetStartMenuLayoutEdgeAssetsXml

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuLayoutEdgeAssetsXml(v string)`

SetStartMenuLayoutEdgeAssetsXml gets a reference to the given string and assigns it to the StartMenuLayoutEdgeAssetsXml field.

### SetStartMenuLayoutEdgeAssetsXmlExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuLayoutEdgeAssetsXmlExplicitNull(b bool)`

SetStartMenuLayoutEdgeAssetsXmlExplicitNull (un)sets StartMenuLayoutEdgeAssetsXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartMenuLayoutEdgeAssetsXml value is set to nil even if false is passed
### GetStartMenuLayoutXml

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuLayoutXml() string`

GetStartMenuLayoutXml returns the StartMenuLayoutXml field if non-nil, zero value otherwise.

### GetStartMenuLayoutXmlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuLayoutXmlOk() (string, bool)`

GetStartMenuLayoutXmlOk returns a tuple with the StartMenuLayoutXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuLayoutXml

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuLayoutXml() bool`

HasStartMenuLayoutXml returns a boolean if a field has been set.

### SetStartMenuLayoutXml

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuLayoutXml(v string)`

SetStartMenuLayoutXml gets a reference to the given string and assigns it to the StartMenuLayoutXml field.

### SetStartMenuLayoutXmlExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuLayoutXmlExplicitNull(b bool)`

SetStartMenuLayoutXmlExplicitNull (un)sets StartMenuLayoutXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartMenuLayoutXml value is set to nil even if false is passed
### GetStartMenuMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuMode() AnyOfmicrosoftGraphWindowsStartMenuModeType`

GetStartMenuMode returns the StartMenuMode field if non-nil, zero value otherwise.

### GetStartMenuModeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuModeOk() (AnyOfmicrosoftGraphWindowsStartMenuModeType, bool)`

GetStartMenuModeOk returns a tuple with the StartMenuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuMode() bool`

HasStartMenuMode returns a boolean if a field has been set.

### SetStartMenuMode

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuMode(v AnyOfmicrosoftGraphWindowsStartMenuModeType)`

SetStartMenuMode gets a reference to the given AnyOfmicrosoftGraphWindowsStartMenuModeType and assigns it to the StartMenuMode field.

### GetStartMenuPinnedFolderDocuments

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderDocuments() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderDocuments returns the StartMenuPinnedFolderDocuments field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderDocumentsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderDocumentsOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderDocumentsOk returns a tuple with the StartMenuPinnedFolderDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderDocuments

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderDocuments() bool`

HasStartMenuPinnedFolderDocuments returns a boolean if a field has been set.

### SetStartMenuPinnedFolderDocuments

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderDocuments(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderDocuments gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderDocuments field.

### GetStartMenuPinnedFolderDownloads

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderDownloads() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderDownloads returns the StartMenuPinnedFolderDownloads field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderDownloadsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderDownloadsOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderDownloadsOk returns a tuple with the StartMenuPinnedFolderDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderDownloads

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderDownloads() bool`

HasStartMenuPinnedFolderDownloads returns a boolean if a field has been set.

### SetStartMenuPinnedFolderDownloads

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderDownloads(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderDownloads gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderDownloads field.

### GetStartMenuPinnedFolderFileExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderFileExplorer() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderFileExplorer returns the StartMenuPinnedFolderFileExplorer field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderFileExplorerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderFileExplorerOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderFileExplorerOk returns a tuple with the StartMenuPinnedFolderFileExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderFileExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderFileExplorer() bool`

HasStartMenuPinnedFolderFileExplorer returns a boolean if a field has been set.

### SetStartMenuPinnedFolderFileExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderFileExplorer(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderFileExplorer gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderFileExplorer field.

### GetStartMenuPinnedFolderHomeGroup

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderHomeGroup() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderHomeGroup returns the StartMenuPinnedFolderHomeGroup field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderHomeGroupOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderHomeGroupOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderHomeGroupOk returns a tuple with the StartMenuPinnedFolderHomeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderHomeGroup

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderHomeGroup() bool`

HasStartMenuPinnedFolderHomeGroup returns a boolean if a field has been set.

### SetStartMenuPinnedFolderHomeGroup

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderHomeGroup(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderHomeGroup gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderHomeGroup field.

### GetStartMenuPinnedFolderMusic

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderMusic() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderMusic returns the StartMenuPinnedFolderMusic field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderMusicOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderMusicOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderMusicOk returns a tuple with the StartMenuPinnedFolderMusic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderMusic

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderMusic() bool`

HasStartMenuPinnedFolderMusic returns a boolean if a field has been set.

### SetStartMenuPinnedFolderMusic

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderMusic(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderMusic gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderMusic field.

### GetStartMenuPinnedFolderNetwork

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderNetwork() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderNetwork returns the StartMenuPinnedFolderNetwork field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderNetworkOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderNetworkOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderNetworkOk returns a tuple with the StartMenuPinnedFolderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderNetwork

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderNetwork() bool`

HasStartMenuPinnedFolderNetwork returns a boolean if a field has been set.

### SetStartMenuPinnedFolderNetwork

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderNetwork(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderNetwork gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderNetwork field.

### GetStartMenuPinnedFolderPersonalFolder

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderPersonalFolder() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderPersonalFolder returns the StartMenuPinnedFolderPersonalFolder field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderPersonalFolderOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderPersonalFolderOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderPersonalFolderOk returns a tuple with the StartMenuPinnedFolderPersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderPersonalFolder

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderPersonalFolder() bool`

HasStartMenuPinnedFolderPersonalFolder returns a boolean if a field has been set.

### SetStartMenuPinnedFolderPersonalFolder

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderPersonalFolder(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderPersonalFolder gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderPersonalFolder field.

### GetStartMenuPinnedFolderPictures

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderPictures() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderPictures returns the StartMenuPinnedFolderPictures field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderPicturesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderPicturesOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderPicturesOk returns a tuple with the StartMenuPinnedFolderPictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderPictures

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderPictures() bool`

HasStartMenuPinnedFolderPictures returns a boolean if a field has been set.

### SetStartMenuPinnedFolderPictures

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderPictures(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderPictures gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderPictures field.

### GetStartMenuPinnedFolderSettings

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderSettings() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderSettings returns the StartMenuPinnedFolderSettings field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderSettingsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderSettingsOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderSettingsOk returns a tuple with the StartMenuPinnedFolderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderSettings

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderSettings() bool`

HasStartMenuPinnedFolderSettings returns a boolean if a field has been set.

### SetStartMenuPinnedFolderSettings

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderSettings(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderSettings gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderSettings field.

### GetStartMenuPinnedFolderVideos

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderVideos() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderVideos returns the StartMenuPinnedFolderVideos field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderVideosOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStartMenuPinnedFolderVideosOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderVideosOk returns a tuple with the StartMenuPinnedFolderVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderVideos

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStartMenuPinnedFolderVideos() bool`

HasStartMenuPinnedFolderVideos returns a boolean if a field has been set.

### SetStartMenuPinnedFolderVideos

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStartMenuPinnedFolderVideos(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderVideos gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderVideos field.

### GetSettingsBlockSettingsApp

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockSettingsApp() bool`

GetSettingsBlockSettingsApp returns the SettingsBlockSettingsApp field if non-nil, zero value otherwise.

### GetSettingsBlockSettingsAppOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockSettingsAppOk() (bool, bool)`

GetSettingsBlockSettingsAppOk returns a tuple with the SettingsBlockSettingsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSettingsApp

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockSettingsApp() bool`

HasSettingsBlockSettingsApp returns a boolean if a field has been set.

### SetSettingsBlockSettingsApp

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockSettingsApp(v bool)`

SetSettingsBlockSettingsApp gets a reference to the given bool and assigns it to the SettingsBlockSettingsApp field.

### GetSettingsBlockSystemPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockSystemPage() bool`

GetSettingsBlockSystemPage returns the SettingsBlockSystemPage field if non-nil, zero value otherwise.

### GetSettingsBlockSystemPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockSystemPageOk() (bool, bool)`

GetSettingsBlockSystemPageOk returns a tuple with the SettingsBlockSystemPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSystemPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockSystemPage() bool`

HasSettingsBlockSystemPage returns a boolean if a field has been set.

### SetSettingsBlockSystemPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockSystemPage(v bool)`

SetSettingsBlockSystemPage gets a reference to the given bool and assigns it to the SettingsBlockSystemPage field.

### GetSettingsBlockDevicesPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockDevicesPage() bool`

GetSettingsBlockDevicesPage returns the SettingsBlockDevicesPage field if non-nil, zero value otherwise.

### GetSettingsBlockDevicesPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockDevicesPageOk() (bool, bool)`

GetSettingsBlockDevicesPageOk returns a tuple with the SettingsBlockDevicesPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockDevicesPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockDevicesPage() bool`

HasSettingsBlockDevicesPage returns a boolean if a field has been set.

### SetSettingsBlockDevicesPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockDevicesPage(v bool)`

SetSettingsBlockDevicesPage gets a reference to the given bool and assigns it to the SettingsBlockDevicesPage field.

### GetSettingsBlockNetworkInternetPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockNetworkInternetPage() bool`

GetSettingsBlockNetworkInternetPage returns the SettingsBlockNetworkInternetPage field if non-nil, zero value otherwise.

### GetSettingsBlockNetworkInternetPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockNetworkInternetPageOk() (bool, bool)`

GetSettingsBlockNetworkInternetPageOk returns a tuple with the SettingsBlockNetworkInternetPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockNetworkInternetPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockNetworkInternetPage() bool`

HasSettingsBlockNetworkInternetPage returns a boolean if a field has been set.

### SetSettingsBlockNetworkInternetPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockNetworkInternetPage(v bool)`

SetSettingsBlockNetworkInternetPage gets a reference to the given bool and assigns it to the SettingsBlockNetworkInternetPage field.

### GetSettingsBlockPersonalizationPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockPersonalizationPage() bool`

GetSettingsBlockPersonalizationPage returns the SettingsBlockPersonalizationPage field if non-nil, zero value otherwise.

### GetSettingsBlockPersonalizationPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockPersonalizationPageOk() (bool, bool)`

GetSettingsBlockPersonalizationPageOk returns a tuple with the SettingsBlockPersonalizationPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockPersonalizationPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockPersonalizationPage() bool`

HasSettingsBlockPersonalizationPage returns a boolean if a field has been set.

### SetSettingsBlockPersonalizationPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockPersonalizationPage(v bool)`

SetSettingsBlockPersonalizationPage gets a reference to the given bool and assigns it to the SettingsBlockPersonalizationPage field.

### GetSettingsBlockAccountsPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockAccountsPage() bool`

GetSettingsBlockAccountsPage returns the SettingsBlockAccountsPage field if non-nil, zero value otherwise.

### GetSettingsBlockAccountsPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockAccountsPageOk() (bool, bool)`

GetSettingsBlockAccountsPageOk returns a tuple with the SettingsBlockAccountsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockAccountsPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockAccountsPage() bool`

HasSettingsBlockAccountsPage returns a boolean if a field has been set.

### SetSettingsBlockAccountsPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockAccountsPage(v bool)`

SetSettingsBlockAccountsPage gets a reference to the given bool and assigns it to the SettingsBlockAccountsPage field.

### GetSettingsBlockTimeLanguagePage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockTimeLanguagePage() bool`

GetSettingsBlockTimeLanguagePage returns the SettingsBlockTimeLanguagePage field if non-nil, zero value otherwise.

### GetSettingsBlockTimeLanguagePageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockTimeLanguagePageOk() (bool, bool)`

GetSettingsBlockTimeLanguagePageOk returns a tuple with the SettingsBlockTimeLanguagePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockTimeLanguagePage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockTimeLanguagePage() bool`

HasSettingsBlockTimeLanguagePage returns a boolean if a field has been set.

### SetSettingsBlockTimeLanguagePage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockTimeLanguagePage(v bool)`

SetSettingsBlockTimeLanguagePage gets a reference to the given bool and assigns it to the SettingsBlockTimeLanguagePage field.

### GetSettingsBlockEaseOfAccessPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockEaseOfAccessPage() bool`

GetSettingsBlockEaseOfAccessPage returns the SettingsBlockEaseOfAccessPage field if non-nil, zero value otherwise.

### GetSettingsBlockEaseOfAccessPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockEaseOfAccessPageOk() (bool, bool)`

GetSettingsBlockEaseOfAccessPageOk returns a tuple with the SettingsBlockEaseOfAccessPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockEaseOfAccessPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockEaseOfAccessPage() bool`

HasSettingsBlockEaseOfAccessPage returns a boolean if a field has been set.

### SetSettingsBlockEaseOfAccessPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockEaseOfAccessPage(v bool)`

SetSettingsBlockEaseOfAccessPage gets a reference to the given bool and assigns it to the SettingsBlockEaseOfAccessPage field.

### GetSettingsBlockPrivacyPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockPrivacyPage() bool`

GetSettingsBlockPrivacyPage returns the SettingsBlockPrivacyPage field if non-nil, zero value otherwise.

### GetSettingsBlockPrivacyPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockPrivacyPageOk() (bool, bool)`

GetSettingsBlockPrivacyPageOk returns a tuple with the SettingsBlockPrivacyPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockPrivacyPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockPrivacyPage() bool`

HasSettingsBlockPrivacyPage returns a boolean if a field has been set.

### SetSettingsBlockPrivacyPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockPrivacyPage(v bool)`

SetSettingsBlockPrivacyPage gets a reference to the given bool and assigns it to the SettingsBlockPrivacyPage field.

### GetSettingsBlockUpdateSecurityPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockUpdateSecurityPage() bool`

GetSettingsBlockUpdateSecurityPage returns the SettingsBlockUpdateSecurityPage field if non-nil, zero value otherwise.

### GetSettingsBlockUpdateSecurityPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockUpdateSecurityPageOk() (bool, bool)`

GetSettingsBlockUpdateSecurityPageOk returns a tuple with the SettingsBlockUpdateSecurityPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockUpdateSecurityPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockUpdateSecurityPage() bool`

HasSettingsBlockUpdateSecurityPage returns a boolean if a field has been set.

### SetSettingsBlockUpdateSecurityPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockUpdateSecurityPage(v bool)`

SetSettingsBlockUpdateSecurityPage gets a reference to the given bool and assigns it to the SettingsBlockUpdateSecurityPage field.

### GetSettingsBlockAppsPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockAppsPage() bool`

GetSettingsBlockAppsPage returns the SettingsBlockAppsPage field if non-nil, zero value otherwise.

### GetSettingsBlockAppsPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockAppsPageOk() (bool, bool)`

GetSettingsBlockAppsPageOk returns a tuple with the SettingsBlockAppsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockAppsPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockAppsPage() bool`

HasSettingsBlockAppsPage returns a boolean if a field has been set.

### SetSettingsBlockAppsPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockAppsPage(v bool)`

SetSettingsBlockAppsPage gets a reference to the given bool and assigns it to the SettingsBlockAppsPage field.

### GetSettingsBlockGamingPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockGamingPage() bool`

GetSettingsBlockGamingPage returns the SettingsBlockGamingPage field if non-nil, zero value otherwise.

### GetSettingsBlockGamingPageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockGamingPageOk() (bool, bool)`

GetSettingsBlockGamingPageOk returns a tuple with the SettingsBlockGamingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockGamingPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockGamingPage() bool`

HasSettingsBlockGamingPage returns a boolean if a field has been set.

### SetSettingsBlockGamingPage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockGamingPage(v bool)`

SetSettingsBlockGamingPage gets a reference to the given bool and assigns it to the SettingsBlockGamingPage field.

### GetWindowsSpotlightBlockConsumerSpecificFeatures

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockConsumerSpecificFeatures() bool`

GetWindowsSpotlightBlockConsumerSpecificFeatures returns the WindowsSpotlightBlockConsumerSpecificFeatures field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockConsumerSpecificFeaturesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockConsumerSpecificFeaturesOk() (bool, bool)`

GetWindowsSpotlightBlockConsumerSpecificFeaturesOk returns a tuple with the WindowsSpotlightBlockConsumerSpecificFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockConsumerSpecificFeatures

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlockConsumerSpecificFeatures() bool`

HasWindowsSpotlightBlockConsumerSpecificFeatures returns a boolean if a field has been set.

### SetWindowsSpotlightBlockConsumerSpecificFeatures

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlockConsumerSpecificFeatures(v bool)`

SetWindowsSpotlightBlockConsumerSpecificFeatures gets a reference to the given bool and assigns it to the WindowsSpotlightBlockConsumerSpecificFeatures field.

### GetWindowsSpotlightBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlocked() bool`

GetWindowsSpotlightBlocked returns the WindowsSpotlightBlocked field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockedOk() (bool, bool)`

GetWindowsSpotlightBlockedOk returns a tuple with the WindowsSpotlightBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlocked() bool`

HasWindowsSpotlightBlocked returns a boolean if a field has been set.

### SetWindowsSpotlightBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlocked(v bool)`

SetWindowsSpotlightBlocked gets a reference to the given bool and assigns it to the WindowsSpotlightBlocked field.

### GetWindowsSpotlightBlockOnActionCenter

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockOnActionCenter() bool`

GetWindowsSpotlightBlockOnActionCenter returns the WindowsSpotlightBlockOnActionCenter field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockOnActionCenterOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockOnActionCenterOk() (bool, bool)`

GetWindowsSpotlightBlockOnActionCenterOk returns a tuple with the WindowsSpotlightBlockOnActionCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockOnActionCenter

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlockOnActionCenter() bool`

HasWindowsSpotlightBlockOnActionCenter returns a boolean if a field has been set.

### SetWindowsSpotlightBlockOnActionCenter

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlockOnActionCenter(v bool)`

SetWindowsSpotlightBlockOnActionCenter gets a reference to the given bool and assigns it to the WindowsSpotlightBlockOnActionCenter field.

### GetWindowsSpotlightBlockTailoredExperiences

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockTailoredExperiences() bool`

GetWindowsSpotlightBlockTailoredExperiences returns the WindowsSpotlightBlockTailoredExperiences field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockTailoredExperiencesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockTailoredExperiencesOk() (bool, bool)`

GetWindowsSpotlightBlockTailoredExperiencesOk returns a tuple with the WindowsSpotlightBlockTailoredExperiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockTailoredExperiences

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlockTailoredExperiences() bool`

HasWindowsSpotlightBlockTailoredExperiences returns a boolean if a field has been set.

### SetWindowsSpotlightBlockTailoredExperiences

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlockTailoredExperiences(v bool)`

SetWindowsSpotlightBlockTailoredExperiences gets a reference to the given bool and assigns it to the WindowsSpotlightBlockTailoredExperiences field.

### GetWindowsSpotlightBlockThirdPartyNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockThirdPartyNotifications() bool`

GetWindowsSpotlightBlockThirdPartyNotifications returns the WindowsSpotlightBlockThirdPartyNotifications field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockThirdPartyNotificationsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockThirdPartyNotificationsOk() (bool, bool)`

GetWindowsSpotlightBlockThirdPartyNotificationsOk returns a tuple with the WindowsSpotlightBlockThirdPartyNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockThirdPartyNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlockThirdPartyNotifications() bool`

HasWindowsSpotlightBlockThirdPartyNotifications returns a boolean if a field has been set.

### SetWindowsSpotlightBlockThirdPartyNotifications

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlockThirdPartyNotifications(v bool)`

SetWindowsSpotlightBlockThirdPartyNotifications gets a reference to the given bool and assigns it to the WindowsSpotlightBlockThirdPartyNotifications field.

### GetWindowsSpotlightBlockWelcomeExperience

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockWelcomeExperience() bool`

GetWindowsSpotlightBlockWelcomeExperience returns the WindowsSpotlightBlockWelcomeExperience field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockWelcomeExperienceOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockWelcomeExperienceOk() (bool, bool)`

GetWindowsSpotlightBlockWelcomeExperienceOk returns a tuple with the WindowsSpotlightBlockWelcomeExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockWelcomeExperience

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlockWelcomeExperience() bool`

HasWindowsSpotlightBlockWelcomeExperience returns a boolean if a field has been set.

### SetWindowsSpotlightBlockWelcomeExperience

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlockWelcomeExperience(v bool)`

SetWindowsSpotlightBlockWelcomeExperience gets a reference to the given bool and assigns it to the WindowsSpotlightBlockWelcomeExperience field.

### GetWindowsSpotlightBlockWindowsTips

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockWindowsTips() bool`

GetWindowsSpotlightBlockWindowsTips returns the WindowsSpotlightBlockWindowsTips field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockWindowsTipsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightBlockWindowsTipsOk() (bool, bool)`

GetWindowsSpotlightBlockWindowsTipsOk returns a tuple with the WindowsSpotlightBlockWindowsTips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockWindowsTips

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightBlockWindowsTips() bool`

HasWindowsSpotlightBlockWindowsTips returns a boolean if a field has been set.

### SetWindowsSpotlightBlockWindowsTips

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightBlockWindowsTips(v bool)`

SetWindowsSpotlightBlockWindowsTips gets a reference to the given bool and assigns it to the WindowsSpotlightBlockWindowsTips field.

### GetWindowsSpotlightConfigureOnLockScreen

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightConfigureOnLockScreen() AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings`

GetWindowsSpotlightConfigureOnLockScreen returns the WindowsSpotlightConfigureOnLockScreen field if non-nil, zero value otherwise.

### GetWindowsSpotlightConfigureOnLockScreenOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsSpotlightConfigureOnLockScreenOk() (AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings, bool)`

GetWindowsSpotlightConfigureOnLockScreenOk returns a tuple with the WindowsSpotlightConfigureOnLockScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightConfigureOnLockScreen

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsSpotlightConfigureOnLockScreen() bool`

HasWindowsSpotlightConfigureOnLockScreen returns a boolean if a field has been set.

### SetWindowsSpotlightConfigureOnLockScreen

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsSpotlightConfigureOnLockScreen(v AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings)`

SetWindowsSpotlightConfigureOnLockScreen gets a reference to the given AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings and assigns it to the WindowsSpotlightConfigureOnLockScreen field.

### GetNetworkProxyApplySettingsDeviceWide

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyApplySettingsDeviceWide() bool`

GetNetworkProxyApplySettingsDeviceWide returns the NetworkProxyApplySettingsDeviceWide field if non-nil, zero value otherwise.

### GetNetworkProxyApplySettingsDeviceWideOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyApplySettingsDeviceWideOk() (bool, bool)`

GetNetworkProxyApplySettingsDeviceWideOk returns a tuple with the NetworkProxyApplySettingsDeviceWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyApplySettingsDeviceWide

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasNetworkProxyApplySettingsDeviceWide() bool`

HasNetworkProxyApplySettingsDeviceWide returns a boolean if a field has been set.

### SetNetworkProxyApplySettingsDeviceWide

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNetworkProxyApplySettingsDeviceWide(v bool)`

SetNetworkProxyApplySettingsDeviceWide gets a reference to the given bool and assigns it to the NetworkProxyApplySettingsDeviceWide field.

### GetNetworkProxyDisableAutoDetect

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyDisableAutoDetect() bool`

GetNetworkProxyDisableAutoDetect returns the NetworkProxyDisableAutoDetect field if non-nil, zero value otherwise.

### GetNetworkProxyDisableAutoDetectOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyDisableAutoDetectOk() (bool, bool)`

GetNetworkProxyDisableAutoDetectOk returns a tuple with the NetworkProxyDisableAutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyDisableAutoDetect

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasNetworkProxyDisableAutoDetect() bool`

HasNetworkProxyDisableAutoDetect returns a boolean if a field has been set.

### SetNetworkProxyDisableAutoDetect

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNetworkProxyDisableAutoDetect(v bool)`

SetNetworkProxyDisableAutoDetect gets a reference to the given bool and assigns it to the NetworkProxyDisableAutoDetect field.

### GetNetworkProxyAutomaticConfigurationUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyAutomaticConfigurationUrl() string`

GetNetworkProxyAutomaticConfigurationUrl returns the NetworkProxyAutomaticConfigurationUrl field if non-nil, zero value otherwise.

### GetNetworkProxyAutomaticConfigurationUrlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyAutomaticConfigurationUrlOk() (string, bool)`

GetNetworkProxyAutomaticConfigurationUrlOk returns a tuple with the NetworkProxyAutomaticConfigurationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyAutomaticConfigurationUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasNetworkProxyAutomaticConfigurationUrl() bool`

HasNetworkProxyAutomaticConfigurationUrl returns a boolean if a field has been set.

### SetNetworkProxyAutomaticConfigurationUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNetworkProxyAutomaticConfigurationUrl(v string)`

SetNetworkProxyAutomaticConfigurationUrl gets a reference to the given string and assigns it to the NetworkProxyAutomaticConfigurationUrl field.

### SetNetworkProxyAutomaticConfigurationUrlExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNetworkProxyAutomaticConfigurationUrlExplicitNull(b bool)`

SetNetworkProxyAutomaticConfigurationUrlExplicitNull (un)sets NetworkProxyAutomaticConfigurationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NetworkProxyAutomaticConfigurationUrl value is set to nil even if false is passed
### GetNetworkProxyServer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyServer() AnyOfmicrosoftGraphWindows10NetworkProxyServer`

GetNetworkProxyServer returns the NetworkProxyServer field if non-nil, zero value otherwise.

### GetNetworkProxyServerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNetworkProxyServerOk() (AnyOfmicrosoftGraphWindows10NetworkProxyServer, bool)`

GetNetworkProxyServerOk returns a tuple with the NetworkProxyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyServer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasNetworkProxyServer() bool`

HasNetworkProxyServer returns a boolean if a field has been set.

### SetNetworkProxyServer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNetworkProxyServer(v AnyOfmicrosoftGraphWindows10NetworkProxyServer)`

SetNetworkProxyServer gets a reference to the given AnyOfmicrosoftGraphWindows10NetworkProxyServer and assigns it to the NetworkProxyServer field.

### SetNetworkProxyServerExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNetworkProxyServerExplicitNull(b bool)`

SetNetworkProxyServerExplicitNull (un)sets NetworkProxyServer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NetworkProxyServer value is set to nil even if false is passed
### GetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmail() bool`

GetAccountsBlockAddingNonMicrosoftAccountEmail returns the AccountsBlockAddingNonMicrosoftAccountEmail field if non-nil, zero value otherwise.

### GetAccountsBlockAddingNonMicrosoftAccountEmailOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmailOk() (bool, bool)`

GetAccountsBlockAddingNonMicrosoftAccountEmailOk returns a tuple with the AccountsBlockAddingNonMicrosoftAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasAccountsBlockAddingNonMicrosoftAccountEmail() bool`

HasAccountsBlockAddingNonMicrosoftAccountEmail returns a boolean if a field has been set.

### SetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetAccountsBlockAddingNonMicrosoftAccountEmail(v bool)`

SetAccountsBlockAddingNonMicrosoftAccountEmail gets a reference to the given bool and assigns it to the AccountsBlockAddingNonMicrosoftAccountEmail field.

### GetAntiTheftModeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAntiTheftModeBlocked() bool`

GetAntiTheftModeBlocked returns the AntiTheftModeBlocked field if non-nil, zero value otherwise.

### GetAntiTheftModeBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAntiTheftModeBlockedOk() (bool, bool)`

GetAntiTheftModeBlockedOk returns a tuple with the AntiTheftModeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAntiTheftModeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasAntiTheftModeBlocked() bool`

HasAntiTheftModeBlocked returns a boolean if a field has been set.

### SetAntiTheftModeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetAntiTheftModeBlocked(v bool)`

SetAntiTheftModeBlocked gets a reference to the given bool and assigns it to the AntiTheftModeBlocked field.

### GetBluetoothBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlocked() bool`

GetBluetoothBlocked returns the BluetoothBlocked field if non-nil, zero value otherwise.

### GetBluetoothBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetBluetoothBlockedOk() (bool, bool)`

GetBluetoothBlockedOk returns a tuple with the BluetoothBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasBluetoothBlocked() bool`

HasBluetoothBlocked returns a boolean if a field has been set.

### SetBluetoothBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetBluetoothBlocked(v bool)`

SetBluetoothBlocked gets a reference to the given bool and assigns it to the BluetoothBlocked field.

### GetCameraBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetConnectedDevicesServiceBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetConnectedDevicesServiceBlocked() bool`

GetConnectedDevicesServiceBlocked returns the ConnectedDevicesServiceBlocked field if non-nil, zero value otherwise.

### GetConnectedDevicesServiceBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetConnectedDevicesServiceBlockedOk() (bool, bool)`

GetConnectedDevicesServiceBlockedOk returns a tuple with the ConnectedDevicesServiceBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectedDevicesServiceBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasConnectedDevicesServiceBlocked() bool`

HasConnectedDevicesServiceBlocked returns a boolean if a field has been set.

### SetConnectedDevicesServiceBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetConnectedDevicesServiceBlocked(v bool)`

SetConnectedDevicesServiceBlocked gets a reference to the given bool and assigns it to the ConnectedDevicesServiceBlocked field.

### GetCertificatesBlockManualRootCertificateInstallation

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCertificatesBlockManualRootCertificateInstallation() bool`

GetCertificatesBlockManualRootCertificateInstallation returns the CertificatesBlockManualRootCertificateInstallation field if non-nil, zero value otherwise.

### GetCertificatesBlockManualRootCertificateInstallationOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCertificatesBlockManualRootCertificateInstallationOk() (bool, bool)`

GetCertificatesBlockManualRootCertificateInstallationOk returns a tuple with the CertificatesBlockManualRootCertificateInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificatesBlockManualRootCertificateInstallation

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCertificatesBlockManualRootCertificateInstallation() bool`

HasCertificatesBlockManualRootCertificateInstallation returns a boolean if a field has been set.

### SetCertificatesBlockManualRootCertificateInstallation

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCertificatesBlockManualRootCertificateInstallation(v bool)`

SetCertificatesBlockManualRootCertificateInstallation gets a reference to the given bool and assigns it to the CertificatesBlockManualRootCertificateInstallation field.

### GetCopyPasteBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCopyPasteBlocked() bool`

GetCopyPasteBlocked returns the CopyPasteBlocked field if non-nil, zero value otherwise.

### GetCopyPasteBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCopyPasteBlockedOk() (bool, bool)`

GetCopyPasteBlockedOk returns a tuple with the CopyPasteBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCopyPasteBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCopyPasteBlocked() bool`

HasCopyPasteBlocked returns a boolean if a field has been set.

### SetCopyPasteBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCopyPasteBlocked(v bool)`

SetCopyPasteBlocked gets a reference to the given bool and assigns it to the CopyPasteBlocked field.

### GetCortanaBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCortanaBlocked() bool`

GetCortanaBlocked returns the CortanaBlocked field if non-nil, zero value otherwise.

### GetCortanaBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetCortanaBlockedOk() (bool, bool)`

GetCortanaBlockedOk returns a tuple with the CortanaBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCortanaBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasCortanaBlocked() bool`

HasCortanaBlocked returns a boolean if a field has been set.

### SetCortanaBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetCortanaBlocked(v bool)`

SetCortanaBlocked gets a reference to the given bool and assigns it to the CortanaBlocked field.

### GetDeviceManagementBlockFactoryResetOnMobile

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceManagementBlockFactoryResetOnMobile() bool`

GetDeviceManagementBlockFactoryResetOnMobile returns the DeviceManagementBlockFactoryResetOnMobile field if non-nil, zero value otherwise.

### GetDeviceManagementBlockFactoryResetOnMobileOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceManagementBlockFactoryResetOnMobileOk() (bool, bool)`

GetDeviceManagementBlockFactoryResetOnMobileOk returns a tuple with the DeviceManagementBlockFactoryResetOnMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementBlockFactoryResetOnMobile

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDeviceManagementBlockFactoryResetOnMobile() bool`

HasDeviceManagementBlockFactoryResetOnMobile returns a boolean if a field has been set.

### SetDeviceManagementBlockFactoryResetOnMobile

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeviceManagementBlockFactoryResetOnMobile(v bool)`

SetDeviceManagementBlockFactoryResetOnMobile gets a reference to the given bool and assigns it to the DeviceManagementBlockFactoryResetOnMobile field.

### GetDeviceManagementBlockManualUnenroll

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceManagementBlockManualUnenroll() bool`

GetDeviceManagementBlockManualUnenroll returns the DeviceManagementBlockManualUnenroll field if non-nil, zero value otherwise.

### GetDeviceManagementBlockManualUnenrollOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeviceManagementBlockManualUnenrollOk() (bool, bool)`

GetDeviceManagementBlockManualUnenrollOk returns a tuple with the DeviceManagementBlockManualUnenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementBlockManualUnenroll

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDeviceManagementBlockManualUnenroll() bool`

HasDeviceManagementBlockManualUnenroll returns a boolean if a field has been set.

### SetDeviceManagementBlockManualUnenroll

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeviceManagementBlockManualUnenroll(v bool)`

SetDeviceManagementBlockManualUnenroll gets a reference to the given bool and assigns it to the DeviceManagementBlockManualUnenroll field.

### GetSafeSearchFilter

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSafeSearchFilter() AnyOfmicrosoftGraphSafeSearchFilterType`

GetSafeSearchFilter returns the SafeSearchFilter field if non-nil, zero value otherwise.

### GetSafeSearchFilterOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSafeSearchFilterOk() (AnyOfmicrosoftGraphSafeSearchFilterType, bool)`

GetSafeSearchFilterOk returns a tuple with the SafeSearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafeSearchFilter

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSafeSearchFilter() bool`

HasSafeSearchFilter returns a boolean if a field has been set.

### SetSafeSearchFilter

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSafeSearchFilter(v AnyOfmicrosoftGraphSafeSearchFilterType)`

SetSafeSearchFilter gets a reference to the given AnyOfmicrosoftGraphSafeSearchFilterType and assigns it to the SafeSearchFilter field.

### GetEdgeBlockPopups

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockPopups() bool`

GetEdgeBlockPopups returns the EdgeBlockPopups field if non-nil, zero value otherwise.

### GetEdgeBlockPopupsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockPopupsOk() (bool, bool)`

GetEdgeBlockPopupsOk returns a tuple with the EdgeBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockPopups

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockPopups() bool`

HasEdgeBlockPopups returns a boolean if a field has been set.

### SetEdgeBlockPopups

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockPopups(v bool)`

SetEdgeBlockPopups gets a reference to the given bool and assigns it to the EdgeBlockPopups field.

### GetEdgeBlockSearchSuggestions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockSearchSuggestions() bool`

GetEdgeBlockSearchSuggestions returns the EdgeBlockSearchSuggestions field if non-nil, zero value otherwise.

### GetEdgeBlockSearchSuggestionsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockSearchSuggestionsOk() (bool, bool)`

GetEdgeBlockSearchSuggestionsOk returns a tuple with the EdgeBlockSearchSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockSearchSuggestions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockSearchSuggestions() bool`

HasEdgeBlockSearchSuggestions returns a boolean if a field has been set.

### SetEdgeBlockSearchSuggestions

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockSearchSuggestions(v bool)`

SetEdgeBlockSearchSuggestions gets a reference to the given bool and assigns it to the EdgeBlockSearchSuggestions field.

### GetEdgeBlockSendingIntranetTrafficToInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockSendingIntranetTrafficToInternetExplorer() bool`

GetEdgeBlockSendingIntranetTrafficToInternetExplorer returns the EdgeBlockSendingIntranetTrafficToInternetExplorer field if non-nil, zero value otherwise.

### GetEdgeBlockSendingIntranetTrafficToInternetExplorerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockSendingIntranetTrafficToInternetExplorerOk() (bool, bool)`

GetEdgeBlockSendingIntranetTrafficToInternetExplorerOk returns a tuple with the EdgeBlockSendingIntranetTrafficToInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockSendingIntranetTrafficToInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockSendingIntranetTrafficToInternetExplorer() bool`

HasEdgeBlockSendingIntranetTrafficToInternetExplorer returns a boolean if a field has been set.

### SetEdgeBlockSendingIntranetTrafficToInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockSendingIntranetTrafficToInternetExplorer(v bool)`

SetEdgeBlockSendingIntranetTrafficToInternetExplorer gets a reference to the given bool and assigns it to the EdgeBlockSendingIntranetTrafficToInternetExplorer field.

### GetEdgeSendIntranetTrafficToInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeSendIntranetTrafficToInternetExplorer() bool`

GetEdgeSendIntranetTrafficToInternetExplorer returns the EdgeSendIntranetTrafficToInternetExplorer field if non-nil, zero value otherwise.

### GetEdgeSendIntranetTrafficToInternetExplorerOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeSendIntranetTrafficToInternetExplorerOk() (bool, bool)`

GetEdgeSendIntranetTrafficToInternetExplorerOk returns a tuple with the EdgeSendIntranetTrafficToInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeSendIntranetTrafficToInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeSendIntranetTrafficToInternetExplorer() bool`

HasEdgeSendIntranetTrafficToInternetExplorer returns a boolean if a field has been set.

### SetEdgeSendIntranetTrafficToInternetExplorer

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeSendIntranetTrafficToInternetExplorer(v bool)`

SetEdgeSendIntranetTrafficToInternetExplorer gets a reference to the given bool and assigns it to the EdgeSendIntranetTrafficToInternetExplorer field.

### GetEdgeRequireSmartScreen

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeRequireSmartScreen() bool`

GetEdgeRequireSmartScreen returns the EdgeRequireSmartScreen field if non-nil, zero value otherwise.

### GetEdgeRequireSmartScreenOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeRequireSmartScreenOk() (bool, bool)`

GetEdgeRequireSmartScreenOk returns a tuple with the EdgeRequireSmartScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeRequireSmartScreen

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeRequireSmartScreen() bool`

HasEdgeRequireSmartScreen returns a boolean if a field has been set.

### SetEdgeRequireSmartScreen

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeRequireSmartScreen(v bool)`

SetEdgeRequireSmartScreen gets a reference to the given bool and assigns it to the EdgeRequireSmartScreen field.

### GetEdgeEnterpriseModeSiteListLocation

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeEnterpriseModeSiteListLocation() string`

GetEdgeEnterpriseModeSiteListLocation returns the EdgeEnterpriseModeSiteListLocation field if non-nil, zero value otherwise.

### GetEdgeEnterpriseModeSiteListLocationOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeEnterpriseModeSiteListLocationOk() (string, bool)`

GetEdgeEnterpriseModeSiteListLocationOk returns a tuple with the EdgeEnterpriseModeSiteListLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeEnterpriseModeSiteListLocation

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeEnterpriseModeSiteListLocation() bool`

HasEdgeEnterpriseModeSiteListLocation returns a boolean if a field has been set.

### SetEdgeEnterpriseModeSiteListLocation

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeEnterpriseModeSiteListLocation(v string)`

SetEdgeEnterpriseModeSiteListLocation gets a reference to the given string and assigns it to the EdgeEnterpriseModeSiteListLocation field.

### SetEdgeEnterpriseModeSiteListLocationExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeEnterpriseModeSiteListLocationExplicitNull(b bool)`

SetEdgeEnterpriseModeSiteListLocationExplicitNull (un)sets EdgeEnterpriseModeSiteListLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EdgeEnterpriseModeSiteListLocation value is set to nil even if false is passed
### GetEdgeFirstRunUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeFirstRunUrl() string`

GetEdgeFirstRunUrl returns the EdgeFirstRunUrl field if non-nil, zero value otherwise.

### GetEdgeFirstRunUrlOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeFirstRunUrlOk() (string, bool)`

GetEdgeFirstRunUrlOk returns a tuple with the EdgeFirstRunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeFirstRunUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeFirstRunUrl() bool`

HasEdgeFirstRunUrl returns a boolean if a field has been set.

### SetEdgeFirstRunUrl

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeFirstRunUrl(v string)`

SetEdgeFirstRunUrl gets a reference to the given string and assigns it to the EdgeFirstRunUrl field.

### SetEdgeFirstRunUrlExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeFirstRunUrlExplicitNull(b bool)`

SetEdgeFirstRunUrlExplicitNull (un)sets EdgeFirstRunUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EdgeFirstRunUrl value is set to nil even if false is passed
### GetEdgeSearchEngine

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeSearchEngine() AnyOfobject`

GetEdgeSearchEngine returns the EdgeSearchEngine field if non-nil, zero value otherwise.

### GetEdgeSearchEngineOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeSearchEngineOk() (AnyOfobject, bool)`

GetEdgeSearchEngineOk returns a tuple with the EdgeSearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeSearchEngine

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeSearchEngine() bool`

HasEdgeSearchEngine returns a boolean if a field has been set.

### SetEdgeSearchEngine

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeSearchEngine(v AnyOfobject)`

SetEdgeSearchEngine gets a reference to the given AnyOfobject and assigns it to the EdgeSearchEngine field.

### SetEdgeSearchEngineExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeSearchEngineExplicitNull(b bool)`

SetEdgeSearchEngineExplicitNull (un)sets EdgeSearchEngine to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EdgeSearchEngine value is set to nil even if false is passed
### GetEdgeHomepageUrls

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeHomepageUrls() []string`

GetEdgeHomepageUrls returns the EdgeHomepageUrls field if non-nil, zero value otherwise.

### GetEdgeHomepageUrlsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeHomepageUrlsOk() ([]string, bool)`

GetEdgeHomepageUrlsOk returns a tuple with the EdgeHomepageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeHomepageUrls

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeHomepageUrls() bool`

HasEdgeHomepageUrls returns a boolean if a field has been set.

### SetEdgeHomepageUrls

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeHomepageUrls(v []string)`

SetEdgeHomepageUrls gets a reference to the given []string and assigns it to the EdgeHomepageUrls field.

### GetEdgeBlockAccessToAboutFlags

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockAccessToAboutFlags() bool`

GetEdgeBlockAccessToAboutFlags returns the EdgeBlockAccessToAboutFlags field if non-nil, zero value otherwise.

### GetEdgeBlockAccessToAboutFlagsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetEdgeBlockAccessToAboutFlagsOk() (bool, bool)`

GetEdgeBlockAccessToAboutFlagsOk returns a tuple with the EdgeBlockAccessToAboutFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockAccessToAboutFlags

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasEdgeBlockAccessToAboutFlags() bool`

HasEdgeBlockAccessToAboutFlags returns a boolean if a field has been set.

### SetEdgeBlockAccessToAboutFlags

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetEdgeBlockAccessToAboutFlags(v bool)`

SetEdgeBlockAccessToAboutFlags gets a reference to the given bool and assigns it to the EdgeBlockAccessToAboutFlags field.

### GetSmartScreenBlockPromptOverride

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSmartScreenBlockPromptOverride() bool`

GetSmartScreenBlockPromptOverride returns the SmartScreenBlockPromptOverride field if non-nil, zero value otherwise.

### GetSmartScreenBlockPromptOverrideOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideOk() (bool, bool)`

GetSmartScreenBlockPromptOverrideOk returns a tuple with the SmartScreenBlockPromptOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenBlockPromptOverride

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSmartScreenBlockPromptOverride() bool`

HasSmartScreenBlockPromptOverride returns a boolean if a field has been set.

### SetSmartScreenBlockPromptOverride

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSmartScreenBlockPromptOverride(v bool)`

SetSmartScreenBlockPromptOverride gets a reference to the given bool and assigns it to the SmartScreenBlockPromptOverride field.

### GetSmartScreenBlockPromptOverrideForFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideForFiles() bool`

GetSmartScreenBlockPromptOverrideForFiles returns the SmartScreenBlockPromptOverrideForFiles field if non-nil, zero value otherwise.

### GetSmartScreenBlockPromptOverrideForFilesOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideForFilesOk() (bool, bool)`

GetSmartScreenBlockPromptOverrideForFilesOk returns a tuple with the SmartScreenBlockPromptOverrideForFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenBlockPromptOverrideForFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSmartScreenBlockPromptOverrideForFiles() bool`

HasSmartScreenBlockPromptOverrideForFiles returns a boolean if a field has been set.

### SetSmartScreenBlockPromptOverrideForFiles

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSmartScreenBlockPromptOverrideForFiles(v bool)`

SetSmartScreenBlockPromptOverrideForFiles gets a reference to the given bool and assigns it to the SmartScreenBlockPromptOverrideForFiles field.

### GetWebRtcBlockLocalhostIpAddress

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWebRtcBlockLocalhostIpAddress() bool`

GetWebRtcBlockLocalhostIpAddress returns the WebRtcBlockLocalhostIpAddress field if non-nil, zero value otherwise.

### GetWebRtcBlockLocalhostIpAddressOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWebRtcBlockLocalhostIpAddressOk() (bool, bool)`

GetWebRtcBlockLocalhostIpAddressOk returns a tuple with the WebRtcBlockLocalhostIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebRtcBlockLocalhostIpAddress

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWebRtcBlockLocalhostIpAddress() bool`

HasWebRtcBlockLocalhostIpAddress returns a boolean if a field has been set.

### SetWebRtcBlockLocalhostIpAddress

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWebRtcBlockLocalhostIpAddress(v bool)`

SetWebRtcBlockLocalhostIpAddress gets a reference to the given bool and assigns it to the WebRtcBlockLocalhostIpAddress field.

### GetInternetSharingBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetInternetSharingBlocked() bool`

GetInternetSharingBlocked returns the InternetSharingBlocked field if non-nil, zero value otherwise.

### GetInternetSharingBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetInternetSharingBlockedOk() (bool, bool)`

GetInternetSharingBlockedOk returns a tuple with the InternetSharingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetSharingBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasInternetSharingBlocked() bool`

HasInternetSharingBlocked returns a boolean if a field has been set.

### SetInternetSharingBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetInternetSharingBlocked(v bool)`

SetInternetSharingBlocked gets a reference to the given bool and assigns it to the InternetSharingBlocked field.

### GetSettingsBlockAddProvisioningPackage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockAddProvisioningPackage() bool`

GetSettingsBlockAddProvisioningPackage returns the SettingsBlockAddProvisioningPackage field if non-nil, zero value otherwise.

### GetSettingsBlockAddProvisioningPackageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockAddProvisioningPackageOk() (bool, bool)`

GetSettingsBlockAddProvisioningPackageOk returns a tuple with the SettingsBlockAddProvisioningPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockAddProvisioningPackage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockAddProvisioningPackage() bool`

HasSettingsBlockAddProvisioningPackage returns a boolean if a field has been set.

### SetSettingsBlockAddProvisioningPackage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockAddProvisioningPackage(v bool)`

SetSettingsBlockAddProvisioningPackage gets a reference to the given bool and assigns it to the SettingsBlockAddProvisioningPackage field.

### GetSettingsBlockRemoveProvisioningPackage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockRemoveProvisioningPackage() bool`

GetSettingsBlockRemoveProvisioningPackage returns the SettingsBlockRemoveProvisioningPackage field if non-nil, zero value otherwise.

### GetSettingsBlockRemoveProvisioningPackageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockRemoveProvisioningPackageOk() (bool, bool)`

GetSettingsBlockRemoveProvisioningPackageOk returns a tuple with the SettingsBlockRemoveProvisioningPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockRemoveProvisioningPackage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockRemoveProvisioningPackage() bool`

HasSettingsBlockRemoveProvisioningPackage returns a boolean if a field has been set.

### SetSettingsBlockRemoveProvisioningPackage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockRemoveProvisioningPackage(v bool)`

SetSettingsBlockRemoveProvisioningPackage gets a reference to the given bool and assigns it to the SettingsBlockRemoveProvisioningPackage field.

### GetSettingsBlockChangeSystemTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangeSystemTime() bool`

GetSettingsBlockChangeSystemTime returns the SettingsBlockChangeSystemTime field if non-nil, zero value otherwise.

### GetSettingsBlockChangeSystemTimeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangeSystemTimeOk() (bool, bool)`

GetSettingsBlockChangeSystemTimeOk returns a tuple with the SettingsBlockChangeSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangeSystemTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockChangeSystemTime() bool`

HasSettingsBlockChangeSystemTime returns a boolean if a field has been set.

### SetSettingsBlockChangeSystemTime

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockChangeSystemTime(v bool)`

SetSettingsBlockChangeSystemTime gets a reference to the given bool and assigns it to the SettingsBlockChangeSystemTime field.

### GetSettingsBlockEditDeviceName

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockEditDeviceName() bool`

GetSettingsBlockEditDeviceName returns the SettingsBlockEditDeviceName field if non-nil, zero value otherwise.

### GetSettingsBlockEditDeviceNameOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockEditDeviceNameOk() (bool, bool)`

GetSettingsBlockEditDeviceNameOk returns a tuple with the SettingsBlockEditDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockEditDeviceName

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockEditDeviceName() bool`

HasSettingsBlockEditDeviceName returns a boolean if a field has been set.

### SetSettingsBlockEditDeviceName

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockEditDeviceName(v bool)`

SetSettingsBlockEditDeviceName gets a reference to the given bool and assigns it to the SettingsBlockEditDeviceName field.

### GetSettingsBlockChangeRegion

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangeRegion() bool`

GetSettingsBlockChangeRegion returns the SettingsBlockChangeRegion field if non-nil, zero value otherwise.

### GetSettingsBlockChangeRegionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangeRegionOk() (bool, bool)`

GetSettingsBlockChangeRegionOk returns a tuple with the SettingsBlockChangeRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangeRegion

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockChangeRegion() bool`

HasSettingsBlockChangeRegion returns a boolean if a field has been set.

### SetSettingsBlockChangeRegion

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockChangeRegion(v bool)`

SetSettingsBlockChangeRegion gets a reference to the given bool and assigns it to the SettingsBlockChangeRegion field.

### GetSettingsBlockChangeLanguage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangeLanguage() bool`

GetSettingsBlockChangeLanguage returns the SettingsBlockChangeLanguage field if non-nil, zero value otherwise.

### GetSettingsBlockChangeLanguageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangeLanguageOk() (bool, bool)`

GetSettingsBlockChangeLanguageOk returns a tuple with the SettingsBlockChangeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangeLanguage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockChangeLanguage() bool`

HasSettingsBlockChangeLanguage returns a boolean if a field has been set.

### SetSettingsBlockChangeLanguage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockChangeLanguage(v bool)`

SetSettingsBlockChangeLanguage gets a reference to the given bool and assigns it to the SettingsBlockChangeLanguage field.

### GetSettingsBlockChangePowerSleep

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangePowerSleep() bool`

GetSettingsBlockChangePowerSleep returns the SettingsBlockChangePowerSleep field if non-nil, zero value otherwise.

### GetSettingsBlockChangePowerSleepOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSettingsBlockChangePowerSleepOk() (bool, bool)`

GetSettingsBlockChangePowerSleepOk returns a tuple with the SettingsBlockChangePowerSleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangePowerSleep

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSettingsBlockChangePowerSleep() bool`

HasSettingsBlockChangePowerSleep returns a boolean if a field has been set.

### SetSettingsBlockChangePowerSleep

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSettingsBlockChangePowerSleep(v bool)`

SetSettingsBlockChangePowerSleep gets a reference to the given bool and assigns it to the SettingsBlockChangePowerSleep field.

### GetLocationServicesBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLocationServicesBlocked() bool`

GetLocationServicesBlocked returns the LocationServicesBlocked field if non-nil, zero value otherwise.

### GetLocationServicesBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLocationServicesBlockedOk() (bool, bool)`

GetLocationServicesBlockedOk returns a tuple with the LocationServicesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationServicesBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLocationServicesBlocked() bool`

HasLocationServicesBlocked returns a boolean if a field has been set.

### SetLocationServicesBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLocationServicesBlocked(v bool)`

SetLocationServicesBlocked gets a reference to the given bool and assigns it to the LocationServicesBlocked field.

### GetMicrosoftAccountBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetMicrosoftAccountBlocked() bool`

GetMicrosoftAccountBlocked returns the MicrosoftAccountBlocked field if non-nil, zero value otherwise.

### GetMicrosoftAccountBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetMicrosoftAccountBlockedOk() (bool, bool)`

GetMicrosoftAccountBlockedOk returns a tuple with the MicrosoftAccountBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftAccountBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasMicrosoftAccountBlocked() bool`

HasMicrosoftAccountBlocked returns a boolean if a field has been set.

### SetMicrosoftAccountBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetMicrosoftAccountBlocked(v bool)`

SetMicrosoftAccountBlocked gets a reference to the given bool and assigns it to the MicrosoftAccountBlocked field.

### GetMicrosoftAccountBlockSettingsSync

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetMicrosoftAccountBlockSettingsSync() bool`

GetMicrosoftAccountBlockSettingsSync returns the MicrosoftAccountBlockSettingsSync field if non-nil, zero value otherwise.

### GetMicrosoftAccountBlockSettingsSyncOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetMicrosoftAccountBlockSettingsSyncOk() (bool, bool)`

GetMicrosoftAccountBlockSettingsSyncOk returns a tuple with the MicrosoftAccountBlockSettingsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftAccountBlockSettingsSync

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasMicrosoftAccountBlockSettingsSync() bool`

HasMicrosoftAccountBlockSettingsSync returns a boolean if a field has been set.

### SetMicrosoftAccountBlockSettingsSync

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetMicrosoftAccountBlockSettingsSync(v bool)`

SetMicrosoftAccountBlockSettingsSync gets a reference to the given bool and assigns it to the MicrosoftAccountBlockSettingsSync field.

### GetNfcBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNfcBlocked() bool`

GetNfcBlocked returns the NfcBlocked field if non-nil, zero value otherwise.

### GetNfcBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetNfcBlockedOk() (bool, bool)`

GetNfcBlockedOk returns a tuple with the NfcBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNfcBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasNfcBlocked() bool`

HasNfcBlocked returns a boolean if a field has been set.

### SetNfcBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetNfcBlocked(v bool)`

SetNfcBlocked gets a reference to the given bool and assigns it to the NfcBlocked field.

### GetResetProtectionModeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetResetProtectionModeBlocked() bool`

GetResetProtectionModeBlocked returns the ResetProtectionModeBlocked field if non-nil, zero value otherwise.

### GetResetProtectionModeBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetResetProtectionModeBlockedOk() (bool, bool)`

GetResetProtectionModeBlockedOk returns a tuple with the ResetProtectionModeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResetProtectionModeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasResetProtectionModeBlocked() bool`

HasResetProtectionModeBlocked returns a boolean if a field has been set.

### SetResetProtectionModeBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetResetProtectionModeBlocked(v bool)`

SetResetProtectionModeBlocked gets a reference to the given bool and assigns it to the ResetProtectionModeBlocked field.

### GetScreenCaptureBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetStorageBlockRemovableStorage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageBlockRemovableStorage() bool`

GetStorageBlockRemovableStorage returns the StorageBlockRemovableStorage field if non-nil, zero value otherwise.

### GetStorageBlockRemovableStorageOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageBlockRemovableStorageOk() (bool, bool)`

GetStorageBlockRemovableStorageOk returns a tuple with the StorageBlockRemovableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockRemovableStorage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStorageBlockRemovableStorage() bool`

HasStorageBlockRemovableStorage returns a boolean if a field has been set.

### SetStorageBlockRemovableStorage

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStorageBlockRemovableStorage(v bool)`

SetStorageBlockRemovableStorage gets a reference to the given bool and assigns it to the StorageBlockRemovableStorage field.

### GetStorageRequireMobileDeviceEncryption

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageRequireMobileDeviceEncryption() bool`

GetStorageRequireMobileDeviceEncryption returns the StorageRequireMobileDeviceEncryption field if non-nil, zero value otherwise.

### GetStorageRequireMobileDeviceEncryptionOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageRequireMobileDeviceEncryptionOk() (bool, bool)`

GetStorageRequireMobileDeviceEncryptionOk returns a tuple with the StorageRequireMobileDeviceEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireMobileDeviceEncryption

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStorageRequireMobileDeviceEncryption() bool`

HasStorageRequireMobileDeviceEncryption returns a boolean if a field has been set.

### SetStorageRequireMobileDeviceEncryption

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStorageRequireMobileDeviceEncryption(v bool)`

SetStorageRequireMobileDeviceEncryption gets a reference to the given bool and assigns it to the StorageRequireMobileDeviceEncryption field.

### GetUsbBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetUsbBlocked() bool`

GetUsbBlocked returns the UsbBlocked field if non-nil, zero value otherwise.

### GetUsbBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetUsbBlockedOk() (bool, bool)`

GetUsbBlockedOk returns a tuple with the UsbBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsbBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasUsbBlocked() bool`

HasUsbBlocked returns a boolean if a field has been set.

### SetUsbBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetUsbBlocked(v bool)`

SetUsbBlocked gets a reference to the given bool and assigns it to the UsbBlocked field.

### GetVoiceRecordingBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetVoiceRecordingBlocked() bool`

GetVoiceRecordingBlocked returns the VoiceRecordingBlocked field if non-nil, zero value otherwise.

### GetVoiceRecordingBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetVoiceRecordingBlockedOk() (bool, bool)`

GetVoiceRecordingBlockedOk returns a tuple with the VoiceRecordingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceRecordingBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasVoiceRecordingBlocked() bool`

HasVoiceRecordingBlocked returns a boolean if a field has been set.

### SetVoiceRecordingBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetVoiceRecordingBlocked(v bool)`

SetVoiceRecordingBlocked gets a reference to the given bool and assigns it to the VoiceRecordingBlocked field.

### GetWiFiBlockAutomaticConnectHotspots

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiBlockAutomaticConnectHotspots() bool`

GetWiFiBlockAutomaticConnectHotspots returns the WiFiBlockAutomaticConnectHotspots field if non-nil, zero value otherwise.

### GetWiFiBlockAutomaticConnectHotspotsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiBlockAutomaticConnectHotspotsOk() (bool, bool)`

GetWiFiBlockAutomaticConnectHotspotsOk returns a tuple with the WiFiBlockAutomaticConnectHotspots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlockAutomaticConnectHotspots

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWiFiBlockAutomaticConnectHotspots() bool`

HasWiFiBlockAutomaticConnectHotspots returns a boolean if a field has been set.

### SetWiFiBlockAutomaticConnectHotspots

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWiFiBlockAutomaticConnectHotspots(v bool)`

SetWiFiBlockAutomaticConnectHotspots gets a reference to the given bool and assigns it to the WiFiBlockAutomaticConnectHotspots field.

### GetWiFiBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiBlocked() bool`

GetWiFiBlocked returns the WiFiBlocked field if non-nil, zero value otherwise.

### GetWiFiBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiBlockedOk() (bool, bool)`

GetWiFiBlockedOk returns a tuple with the WiFiBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWiFiBlocked() bool`

HasWiFiBlocked returns a boolean if a field has been set.

### SetWiFiBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWiFiBlocked(v bool)`

SetWiFiBlocked gets a reference to the given bool and assigns it to the WiFiBlocked field.

### GetWiFiBlockManualConfiguration

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiBlockManualConfiguration() bool`

GetWiFiBlockManualConfiguration returns the WiFiBlockManualConfiguration field if non-nil, zero value otherwise.

### GetWiFiBlockManualConfigurationOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiBlockManualConfigurationOk() (bool, bool)`

GetWiFiBlockManualConfigurationOk returns a tuple with the WiFiBlockManualConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlockManualConfiguration

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWiFiBlockManualConfiguration() bool`

HasWiFiBlockManualConfiguration returns a boolean if a field has been set.

### SetWiFiBlockManualConfiguration

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWiFiBlockManualConfiguration(v bool)`

SetWiFiBlockManualConfiguration gets a reference to the given bool and assigns it to the WiFiBlockManualConfiguration field.

### GetWiFiScanInterval

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiScanInterval() int32`

GetWiFiScanInterval returns the WiFiScanInterval field if non-nil, zero value otherwise.

### GetWiFiScanIntervalOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWiFiScanIntervalOk() (int32, bool)`

GetWiFiScanIntervalOk returns a tuple with the WiFiScanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiScanInterval

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWiFiScanInterval() bool`

HasWiFiScanInterval returns a boolean if a field has been set.

### SetWiFiScanInterval

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWiFiScanInterval(v int32)`

SetWiFiScanInterval gets a reference to the given int32 and assigns it to the WiFiScanInterval field.

### SetWiFiScanIntervalExplicitNull

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWiFiScanIntervalExplicitNull(b bool)`

SetWiFiScanIntervalExplicitNull (un)sets WiFiScanInterval to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WiFiScanInterval value is set to nil even if false is passed
### GetWirelessDisplayBlockProjectionToThisDevice

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWirelessDisplayBlockProjectionToThisDevice() bool`

GetWirelessDisplayBlockProjectionToThisDevice returns the WirelessDisplayBlockProjectionToThisDevice field if non-nil, zero value otherwise.

### GetWirelessDisplayBlockProjectionToThisDeviceOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWirelessDisplayBlockProjectionToThisDeviceOk() (bool, bool)`

GetWirelessDisplayBlockProjectionToThisDeviceOk returns a tuple with the WirelessDisplayBlockProjectionToThisDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWirelessDisplayBlockProjectionToThisDevice

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWirelessDisplayBlockProjectionToThisDevice() bool`

HasWirelessDisplayBlockProjectionToThisDevice returns a boolean if a field has been set.

### SetWirelessDisplayBlockProjectionToThisDevice

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWirelessDisplayBlockProjectionToThisDevice(v bool)`

SetWirelessDisplayBlockProjectionToThisDevice gets a reference to the given bool and assigns it to the WirelessDisplayBlockProjectionToThisDevice field.

### GetWirelessDisplayBlockUserInputFromReceiver

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWirelessDisplayBlockUserInputFromReceiver() bool`

GetWirelessDisplayBlockUserInputFromReceiver returns the WirelessDisplayBlockUserInputFromReceiver field if non-nil, zero value otherwise.

### GetWirelessDisplayBlockUserInputFromReceiverOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWirelessDisplayBlockUserInputFromReceiverOk() (bool, bool)`

GetWirelessDisplayBlockUserInputFromReceiverOk returns a tuple with the WirelessDisplayBlockUserInputFromReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWirelessDisplayBlockUserInputFromReceiver

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWirelessDisplayBlockUserInputFromReceiver() bool`

HasWirelessDisplayBlockUserInputFromReceiver returns a boolean if a field has been set.

### SetWirelessDisplayBlockUserInputFromReceiver

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWirelessDisplayBlockUserInputFromReceiver(v bool)`

SetWirelessDisplayBlockUserInputFromReceiver gets a reference to the given bool and assigns it to the WirelessDisplayBlockUserInputFromReceiver field.

### GetWirelessDisplayRequirePinForPairing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWirelessDisplayRequirePinForPairing() bool`

GetWirelessDisplayRequirePinForPairing returns the WirelessDisplayRequirePinForPairing field if non-nil, zero value otherwise.

### GetWirelessDisplayRequirePinForPairingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWirelessDisplayRequirePinForPairingOk() (bool, bool)`

GetWirelessDisplayRequirePinForPairingOk returns a tuple with the WirelessDisplayRequirePinForPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWirelessDisplayRequirePinForPairing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWirelessDisplayRequirePinForPairing() bool`

HasWirelessDisplayRequirePinForPairing returns a boolean if a field has been set.

### SetWirelessDisplayRequirePinForPairing

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWirelessDisplayRequirePinForPairing(v bool)`

SetWirelessDisplayRequirePinForPairing gets a reference to the given bool and assigns it to the WirelessDisplayRequirePinForPairing field.

### GetWindowsStoreBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsStoreBlocked() bool`

GetWindowsStoreBlocked returns the WindowsStoreBlocked field if non-nil, zero value otherwise.

### GetWindowsStoreBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsStoreBlockedOk() (bool, bool)`

GetWindowsStoreBlockedOk returns a tuple with the WindowsStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsStoreBlocked() bool`

HasWindowsStoreBlocked returns a boolean if a field has been set.

### SetWindowsStoreBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsStoreBlocked(v bool)`

SetWindowsStoreBlocked gets a reference to the given bool and assigns it to the WindowsStoreBlocked field.

### GetAppsAllowTrustedAppsSideloading

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAppsAllowTrustedAppsSideloading() AnyOfmicrosoftGraphStateManagementSetting`

GetAppsAllowTrustedAppsSideloading returns the AppsAllowTrustedAppsSideloading field if non-nil, zero value otherwise.

### GetAppsAllowTrustedAppsSideloadingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAppsAllowTrustedAppsSideloadingOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetAppsAllowTrustedAppsSideloadingOk returns a tuple with the AppsAllowTrustedAppsSideloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsAllowTrustedAppsSideloading

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasAppsAllowTrustedAppsSideloading() bool`

HasAppsAllowTrustedAppsSideloading returns a boolean if a field has been set.

### SetAppsAllowTrustedAppsSideloading

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetAppsAllowTrustedAppsSideloading(v AnyOfmicrosoftGraphStateManagementSetting)`

SetAppsAllowTrustedAppsSideloading gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the AppsAllowTrustedAppsSideloading field.

### GetWindowsStoreBlockAutoUpdate

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsStoreBlockAutoUpdate() bool`

GetWindowsStoreBlockAutoUpdate returns the WindowsStoreBlockAutoUpdate field if non-nil, zero value otherwise.

### GetWindowsStoreBlockAutoUpdateOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsStoreBlockAutoUpdateOk() (bool, bool)`

GetWindowsStoreBlockAutoUpdateOk returns a tuple with the WindowsStoreBlockAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreBlockAutoUpdate

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsStoreBlockAutoUpdate() bool`

HasWindowsStoreBlockAutoUpdate returns a boolean if a field has been set.

### SetWindowsStoreBlockAutoUpdate

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsStoreBlockAutoUpdate(v bool)`

SetWindowsStoreBlockAutoUpdate gets a reference to the given bool and assigns it to the WindowsStoreBlockAutoUpdate field.

### GetDeveloperUnlockSetting

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeveloperUnlockSetting() AnyOfmicrosoftGraphStateManagementSetting`

GetDeveloperUnlockSetting returns the DeveloperUnlockSetting field if non-nil, zero value otherwise.

### GetDeveloperUnlockSettingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetDeveloperUnlockSettingOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetDeveloperUnlockSettingOk returns a tuple with the DeveloperUnlockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloperUnlockSetting

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasDeveloperUnlockSetting() bool`

HasDeveloperUnlockSetting returns a boolean if a field has been set.

### SetDeveloperUnlockSetting

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetDeveloperUnlockSetting(v AnyOfmicrosoftGraphStateManagementSetting)`

SetDeveloperUnlockSetting gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the DeveloperUnlockSetting field.

### GetSharedUserAppDataAllowed

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSharedUserAppDataAllowed() bool`

GetSharedUserAppDataAllowed returns the SharedUserAppDataAllowed field if non-nil, zero value otherwise.

### GetSharedUserAppDataAllowedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetSharedUserAppDataAllowedOk() (bool, bool)`

GetSharedUserAppDataAllowedOk returns a tuple with the SharedUserAppDataAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedUserAppDataAllowed

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasSharedUserAppDataAllowed() bool`

HasSharedUserAppDataAllowed returns a boolean if a field has been set.

### SetSharedUserAppDataAllowed

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetSharedUserAppDataAllowed(v bool)`

SetSharedUserAppDataAllowed gets a reference to the given bool and assigns it to the SharedUserAppDataAllowed field.

### GetAppsBlockWindowsStoreOriginatedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAppsBlockWindowsStoreOriginatedApps() bool`

GetAppsBlockWindowsStoreOriginatedApps returns the AppsBlockWindowsStoreOriginatedApps field if non-nil, zero value otherwise.

### GetAppsBlockWindowsStoreOriginatedAppsOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetAppsBlockWindowsStoreOriginatedAppsOk() (bool, bool)`

GetAppsBlockWindowsStoreOriginatedAppsOk returns a tuple with the AppsBlockWindowsStoreOriginatedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockWindowsStoreOriginatedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasAppsBlockWindowsStoreOriginatedApps() bool`

HasAppsBlockWindowsStoreOriginatedApps returns a boolean if a field has been set.

### SetAppsBlockWindowsStoreOriginatedApps

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetAppsBlockWindowsStoreOriginatedApps(v bool)`

SetAppsBlockWindowsStoreOriginatedApps gets a reference to the given bool and assigns it to the AppsBlockWindowsStoreOriginatedApps field.

### GetWindowsStoreEnablePrivateStoreOnly

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsStoreEnablePrivateStoreOnly() bool`

GetWindowsStoreEnablePrivateStoreOnly returns the WindowsStoreEnablePrivateStoreOnly field if non-nil, zero value otherwise.

### GetWindowsStoreEnablePrivateStoreOnlyOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetWindowsStoreEnablePrivateStoreOnlyOk() (bool, bool)`

GetWindowsStoreEnablePrivateStoreOnlyOk returns a tuple with the WindowsStoreEnablePrivateStoreOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreEnablePrivateStoreOnly

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasWindowsStoreEnablePrivateStoreOnly() bool`

HasWindowsStoreEnablePrivateStoreOnly returns a boolean if a field has been set.

### SetWindowsStoreEnablePrivateStoreOnly

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetWindowsStoreEnablePrivateStoreOnly(v bool)`

SetWindowsStoreEnablePrivateStoreOnly gets a reference to the given bool and assigns it to the WindowsStoreEnablePrivateStoreOnly field.

### GetStorageRestrictAppDataToSystemVolume

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageRestrictAppDataToSystemVolume() bool`

GetStorageRestrictAppDataToSystemVolume returns the StorageRestrictAppDataToSystemVolume field if non-nil, zero value otherwise.

### GetStorageRestrictAppDataToSystemVolumeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageRestrictAppDataToSystemVolumeOk() (bool, bool)`

GetStorageRestrictAppDataToSystemVolumeOk returns a tuple with the StorageRestrictAppDataToSystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRestrictAppDataToSystemVolume

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStorageRestrictAppDataToSystemVolume() bool`

HasStorageRestrictAppDataToSystemVolume returns a boolean if a field has been set.

### SetStorageRestrictAppDataToSystemVolume

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStorageRestrictAppDataToSystemVolume(v bool)`

SetStorageRestrictAppDataToSystemVolume gets a reference to the given bool and assigns it to the StorageRestrictAppDataToSystemVolume field.

### GetStorageRestrictAppInstallToSystemVolume

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageRestrictAppInstallToSystemVolume() bool`

GetStorageRestrictAppInstallToSystemVolume returns the StorageRestrictAppInstallToSystemVolume field if non-nil, zero value otherwise.

### GetStorageRestrictAppInstallToSystemVolumeOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetStorageRestrictAppInstallToSystemVolumeOk() (bool, bool)`

GetStorageRestrictAppInstallToSystemVolumeOk returns a tuple with the StorageRestrictAppInstallToSystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRestrictAppInstallToSystemVolume

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasStorageRestrictAppInstallToSystemVolume() bool`

HasStorageRestrictAppInstallToSystemVolume returns a boolean if a field has been set.

### SetStorageRestrictAppInstallToSystemVolume

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetStorageRestrictAppInstallToSystemVolume(v bool)`

SetStorageRestrictAppInstallToSystemVolume gets a reference to the given bool and assigns it to the StorageRestrictAppInstallToSystemVolume field.

### GetGameDvrBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetGameDvrBlocked() bool`

GetGameDvrBlocked returns the GameDvrBlocked field if non-nil, zero value otherwise.

### GetGameDvrBlockedOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetGameDvrBlockedOk() (bool, bool)`

GetGameDvrBlockedOk returns a tuple with the GameDvrBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGameDvrBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasGameDvrBlocked() bool`

HasGameDvrBlocked returns a boolean if a field has been set.

### SetGameDvrBlocked

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetGameDvrBlocked(v bool)`

SetGameDvrBlocked gets a reference to the given bool and assigns it to the GameDvrBlocked field.

### GetExperienceBlockDeviceDiscovery

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetExperienceBlockDeviceDiscovery() bool`

GetExperienceBlockDeviceDiscovery returns the ExperienceBlockDeviceDiscovery field if non-nil, zero value otherwise.

### GetExperienceBlockDeviceDiscoveryOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetExperienceBlockDeviceDiscoveryOk() (bool, bool)`

GetExperienceBlockDeviceDiscoveryOk returns a tuple with the ExperienceBlockDeviceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperienceBlockDeviceDiscovery

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasExperienceBlockDeviceDiscovery() bool`

HasExperienceBlockDeviceDiscovery returns a boolean if a field has been set.

### SetExperienceBlockDeviceDiscovery

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetExperienceBlockDeviceDiscovery(v bool)`

SetExperienceBlockDeviceDiscovery gets a reference to the given bool and assigns it to the ExperienceBlockDeviceDiscovery field.

### GetExperienceBlockErrorDialogWhenNoSIM

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetExperienceBlockErrorDialogWhenNoSIM() bool`

GetExperienceBlockErrorDialogWhenNoSIM returns the ExperienceBlockErrorDialogWhenNoSIM field if non-nil, zero value otherwise.

### GetExperienceBlockErrorDialogWhenNoSIMOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetExperienceBlockErrorDialogWhenNoSIMOk() (bool, bool)`

GetExperienceBlockErrorDialogWhenNoSIMOk returns a tuple with the ExperienceBlockErrorDialogWhenNoSIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperienceBlockErrorDialogWhenNoSIM

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasExperienceBlockErrorDialogWhenNoSIM() bool`

HasExperienceBlockErrorDialogWhenNoSIM returns a boolean if a field has been set.

### SetExperienceBlockErrorDialogWhenNoSIM

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetExperienceBlockErrorDialogWhenNoSIM(v bool)`

SetExperienceBlockErrorDialogWhenNoSIM gets a reference to the given bool and assigns it to the ExperienceBlockErrorDialogWhenNoSIM field.

### GetExperienceBlockTaskSwitcher

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetExperienceBlockTaskSwitcher() bool`

GetExperienceBlockTaskSwitcher returns the ExperienceBlockTaskSwitcher field if non-nil, zero value otherwise.

### GetExperienceBlockTaskSwitcherOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetExperienceBlockTaskSwitcherOk() (bool, bool)`

GetExperienceBlockTaskSwitcherOk returns a tuple with the ExperienceBlockTaskSwitcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperienceBlockTaskSwitcher

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasExperienceBlockTaskSwitcher() bool`

HasExperienceBlockTaskSwitcher returns a boolean if a field has been set.

### SetExperienceBlockTaskSwitcher

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetExperienceBlockTaskSwitcher(v bool)`

SetExperienceBlockTaskSwitcher gets a reference to the given bool and assigns it to the ExperienceBlockTaskSwitcher field.

### GetLogonBlockFastUserSwitching

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLogonBlockFastUserSwitching() bool`

GetLogonBlockFastUserSwitching returns the LogonBlockFastUserSwitching field if non-nil, zero value otherwise.

### GetLogonBlockFastUserSwitchingOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetLogonBlockFastUserSwitchingOk() (bool, bool)`

GetLogonBlockFastUserSwitchingOk returns a tuple with the LogonBlockFastUserSwitching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonBlockFastUserSwitching

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasLogonBlockFastUserSwitching() bool`

HasLogonBlockFastUserSwitching returns a boolean if a field has been set.

### SetLogonBlockFastUserSwitching

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetLogonBlockFastUserSwitching(v bool)`

SetLogonBlockFastUserSwitching gets a reference to the given bool and assigns it to the LogonBlockFastUserSwitching field.

### GetTenantLockdownRequireNetworkDuringOutOfBoxExperience

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetTenantLockdownRequireNetworkDuringOutOfBoxExperience() bool`

GetTenantLockdownRequireNetworkDuringOutOfBoxExperience returns the TenantLockdownRequireNetworkDuringOutOfBoxExperience field if non-nil, zero value otherwise.

### GetTenantLockdownRequireNetworkDuringOutOfBoxExperienceOk

`func (o *MicrosoftGraphWindows10GeneralConfiguration) GetTenantLockdownRequireNetworkDuringOutOfBoxExperienceOk() (bool, bool)`

GetTenantLockdownRequireNetworkDuringOutOfBoxExperienceOk returns a tuple with the TenantLockdownRequireNetworkDuringOutOfBoxExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTenantLockdownRequireNetworkDuringOutOfBoxExperience

`func (o *MicrosoftGraphWindows10GeneralConfiguration) HasTenantLockdownRequireNetworkDuringOutOfBoxExperience() bool`

HasTenantLockdownRequireNetworkDuringOutOfBoxExperience returns a boolean if a field has been set.

### SetTenantLockdownRequireNetworkDuringOutOfBoxExperience

`func (o *MicrosoftGraphWindows10GeneralConfiguration) SetTenantLockdownRequireNetworkDuringOutOfBoxExperience(v bool)`

SetTenantLockdownRequireNetworkDuringOutOfBoxExperience gets a reference to the given bool and assigns it to the TenantLockdownRequireNetworkDuringOutOfBoxExperience field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


