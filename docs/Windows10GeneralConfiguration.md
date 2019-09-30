# Windows10GeneralConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetEnterpriseCloudPrintDiscoveryEndPoint

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryEndPoint() string`

GetEnterpriseCloudPrintDiscoveryEndPoint returns the EnterpriseCloudPrintDiscoveryEndPoint field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintDiscoveryEndPointOk

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryEndPointOk() (string, bool)`

GetEnterpriseCloudPrintDiscoveryEndPointOk returns a tuple with the EnterpriseCloudPrintDiscoveryEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintDiscoveryEndPoint

`func (o *Windows10GeneralConfiguration) HasEnterpriseCloudPrintDiscoveryEndPoint() bool`

HasEnterpriseCloudPrintDiscoveryEndPoint returns a boolean if a field has been set.

### SetEnterpriseCloudPrintDiscoveryEndPoint

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryEndPoint(v string)`

SetEnterpriseCloudPrintDiscoveryEndPoint gets a reference to the given string and assigns it to the EnterpriseCloudPrintDiscoveryEndPoint field.

### SetEnterpriseCloudPrintDiscoveryEndPointExplicitNull

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryEndPointExplicitNull(b bool)`

SetEnterpriseCloudPrintDiscoveryEndPointExplicitNull (un)sets EnterpriseCloudPrintDiscoveryEndPoint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintDiscoveryEndPoint value is set to nil even if false is passed
### GetEnterpriseCloudPrintOAuthAuthority

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthAuthority() string`

GetEnterpriseCloudPrintOAuthAuthority returns the EnterpriseCloudPrintOAuthAuthority field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintOAuthAuthorityOk

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthAuthorityOk() (string, bool)`

GetEnterpriseCloudPrintOAuthAuthorityOk returns a tuple with the EnterpriseCloudPrintOAuthAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintOAuthAuthority

`func (o *Windows10GeneralConfiguration) HasEnterpriseCloudPrintOAuthAuthority() bool`

HasEnterpriseCloudPrintOAuthAuthority returns a boolean if a field has been set.

### SetEnterpriseCloudPrintOAuthAuthority

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthAuthority(v string)`

SetEnterpriseCloudPrintOAuthAuthority gets a reference to the given string and assigns it to the EnterpriseCloudPrintOAuthAuthority field.

### SetEnterpriseCloudPrintOAuthAuthorityExplicitNull

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthAuthorityExplicitNull(b bool)`

SetEnterpriseCloudPrintOAuthAuthorityExplicitNull (un)sets EnterpriseCloudPrintOAuthAuthority to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintOAuthAuthority value is set to nil even if false is passed
### GetEnterpriseCloudPrintOAuthClientIdentifier

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthClientIdentifier() string`

GetEnterpriseCloudPrintOAuthClientIdentifier returns the EnterpriseCloudPrintOAuthClientIdentifier field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintOAuthClientIdentifierOk

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthClientIdentifierOk() (string, bool)`

GetEnterpriseCloudPrintOAuthClientIdentifierOk returns a tuple with the EnterpriseCloudPrintOAuthClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintOAuthClientIdentifier

`func (o *Windows10GeneralConfiguration) HasEnterpriseCloudPrintOAuthClientIdentifier() bool`

HasEnterpriseCloudPrintOAuthClientIdentifier returns a boolean if a field has been set.

### SetEnterpriseCloudPrintOAuthClientIdentifier

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthClientIdentifier(v string)`

SetEnterpriseCloudPrintOAuthClientIdentifier gets a reference to the given string and assigns it to the EnterpriseCloudPrintOAuthClientIdentifier field.

### SetEnterpriseCloudPrintOAuthClientIdentifierExplicitNull

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthClientIdentifierExplicitNull(b bool)`

SetEnterpriseCloudPrintOAuthClientIdentifierExplicitNull (un)sets EnterpriseCloudPrintOAuthClientIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintOAuthClientIdentifier value is set to nil even if false is passed
### GetEnterpriseCloudPrintResourceIdentifier

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintResourceIdentifier() string`

GetEnterpriseCloudPrintResourceIdentifier returns the EnterpriseCloudPrintResourceIdentifier field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintResourceIdentifierOk

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintResourceIdentifierOk() (string, bool)`

GetEnterpriseCloudPrintResourceIdentifierOk returns a tuple with the EnterpriseCloudPrintResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintResourceIdentifier

`func (o *Windows10GeneralConfiguration) HasEnterpriseCloudPrintResourceIdentifier() bool`

HasEnterpriseCloudPrintResourceIdentifier returns a boolean if a field has been set.

### SetEnterpriseCloudPrintResourceIdentifier

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintResourceIdentifier(v string)`

SetEnterpriseCloudPrintResourceIdentifier gets a reference to the given string and assigns it to the EnterpriseCloudPrintResourceIdentifier field.

### SetEnterpriseCloudPrintResourceIdentifierExplicitNull

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintResourceIdentifierExplicitNull(b bool)`

SetEnterpriseCloudPrintResourceIdentifierExplicitNull (un)sets EnterpriseCloudPrintResourceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintResourceIdentifier value is set to nil even if false is passed
### GetEnterpriseCloudPrintDiscoveryMaxLimit

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryMaxLimit() int32`

GetEnterpriseCloudPrintDiscoveryMaxLimit returns the EnterpriseCloudPrintDiscoveryMaxLimit field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintDiscoveryMaxLimitOk

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryMaxLimitOk() (int32, bool)`

GetEnterpriseCloudPrintDiscoveryMaxLimitOk returns a tuple with the EnterpriseCloudPrintDiscoveryMaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintDiscoveryMaxLimit

`func (o *Windows10GeneralConfiguration) HasEnterpriseCloudPrintDiscoveryMaxLimit() bool`

HasEnterpriseCloudPrintDiscoveryMaxLimit returns a boolean if a field has been set.

### SetEnterpriseCloudPrintDiscoveryMaxLimit

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryMaxLimit(v int32)`

SetEnterpriseCloudPrintDiscoveryMaxLimit gets a reference to the given int32 and assigns it to the EnterpriseCloudPrintDiscoveryMaxLimit field.

### SetEnterpriseCloudPrintDiscoveryMaxLimitExplicitNull

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryMaxLimitExplicitNull(b bool)`

SetEnterpriseCloudPrintDiscoveryMaxLimitExplicitNull (un)sets EnterpriseCloudPrintDiscoveryMaxLimit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintDiscoveryMaxLimit value is set to nil even if false is passed
### GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier() string`

GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier returns the EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier field if non-nil, zero value otherwise.

### GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierOk

`func (o *Windows10GeneralConfiguration) GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierOk() (string, bool)`

GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierOk returns a tuple with the EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier

`func (o *Windows10GeneralConfiguration) HasEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier() bool`

HasEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier returns a boolean if a field has been set.

### SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier(v string)`

SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier gets a reference to the given string and assigns it to the EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier field.

### SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierExplicitNull

`func (o *Windows10GeneralConfiguration) SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierExplicitNull(b bool)`

SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifierExplicitNull (un)sets EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier value is set to nil even if false is passed
### GetSearchBlockDiacritics

`func (o *Windows10GeneralConfiguration) GetSearchBlockDiacritics() bool`

GetSearchBlockDiacritics returns the SearchBlockDiacritics field if non-nil, zero value otherwise.

### GetSearchBlockDiacriticsOk

`func (o *Windows10GeneralConfiguration) GetSearchBlockDiacriticsOk() (bool, bool)`

GetSearchBlockDiacriticsOk returns a tuple with the SearchBlockDiacritics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchBlockDiacritics

`func (o *Windows10GeneralConfiguration) HasSearchBlockDiacritics() bool`

HasSearchBlockDiacritics returns a boolean if a field has been set.

### SetSearchBlockDiacritics

`func (o *Windows10GeneralConfiguration) SetSearchBlockDiacritics(v bool)`

SetSearchBlockDiacritics gets a reference to the given bool and assigns it to the SearchBlockDiacritics field.

### GetSearchDisableAutoLanguageDetection

`func (o *Windows10GeneralConfiguration) GetSearchDisableAutoLanguageDetection() bool`

GetSearchDisableAutoLanguageDetection returns the SearchDisableAutoLanguageDetection field if non-nil, zero value otherwise.

### GetSearchDisableAutoLanguageDetectionOk

`func (o *Windows10GeneralConfiguration) GetSearchDisableAutoLanguageDetectionOk() (bool, bool)`

GetSearchDisableAutoLanguageDetectionOk returns a tuple with the SearchDisableAutoLanguageDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableAutoLanguageDetection

`func (o *Windows10GeneralConfiguration) HasSearchDisableAutoLanguageDetection() bool`

HasSearchDisableAutoLanguageDetection returns a boolean if a field has been set.

### SetSearchDisableAutoLanguageDetection

`func (o *Windows10GeneralConfiguration) SetSearchDisableAutoLanguageDetection(v bool)`

SetSearchDisableAutoLanguageDetection gets a reference to the given bool and assigns it to the SearchDisableAutoLanguageDetection field.

### GetSearchDisableIndexingEncryptedItems

`func (o *Windows10GeneralConfiguration) GetSearchDisableIndexingEncryptedItems() bool`

GetSearchDisableIndexingEncryptedItems returns the SearchDisableIndexingEncryptedItems field if non-nil, zero value otherwise.

### GetSearchDisableIndexingEncryptedItemsOk

`func (o *Windows10GeneralConfiguration) GetSearchDisableIndexingEncryptedItemsOk() (bool, bool)`

GetSearchDisableIndexingEncryptedItemsOk returns a tuple with the SearchDisableIndexingEncryptedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableIndexingEncryptedItems

`func (o *Windows10GeneralConfiguration) HasSearchDisableIndexingEncryptedItems() bool`

HasSearchDisableIndexingEncryptedItems returns a boolean if a field has been set.

### SetSearchDisableIndexingEncryptedItems

`func (o *Windows10GeneralConfiguration) SetSearchDisableIndexingEncryptedItems(v bool)`

SetSearchDisableIndexingEncryptedItems gets a reference to the given bool and assigns it to the SearchDisableIndexingEncryptedItems field.

### GetSearchEnableRemoteQueries

`func (o *Windows10GeneralConfiguration) GetSearchEnableRemoteQueries() bool`

GetSearchEnableRemoteQueries returns the SearchEnableRemoteQueries field if non-nil, zero value otherwise.

### GetSearchEnableRemoteQueriesOk

`func (o *Windows10GeneralConfiguration) GetSearchEnableRemoteQueriesOk() (bool, bool)`

GetSearchEnableRemoteQueriesOk returns a tuple with the SearchEnableRemoteQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchEnableRemoteQueries

`func (o *Windows10GeneralConfiguration) HasSearchEnableRemoteQueries() bool`

HasSearchEnableRemoteQueries returns a boolean if a field has been set.

### SetSearchEnableRemoteQueries

`func (o *Windows10GeneralConfiguration) SetSearchEnableRemoteQueries(v bool)`

SetSearchEnableRemoteQueries gets a reference to the given bool and assigns it to the SearchEnableRemoteQueries field.

### GetSearchDisableIndexerBackoff

`func (o *Windows10GeneralConfiguration) GetSearchDisableIndexerBackoff() bool`

GetSearchDisableIndexerBackoff returns the SearchDisableIndexerBackoff field if non-nil, zero value otherwise.

### GetSearchDisableIndexerBackoffOk

`func (o *Windows10GeneralConfiguration) GetSearchDisableIndexerBackoffOk() (bool, bool)`

GetSearchDisableIndexerBackoffOk returns a tuple with the SearchDisableIndexerBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableIndexerBackoff

`func (o *Windows10GeneralConfiguration) HasSearchDisableIndexerBackoff() bool`

HasSearchDisableIndexerBackoff returns a boolean if a field has been set.

### SetSearchDisableIndexerBackoff

`func (o *Windows10GeneralConfiguration) SetSearchDisableIndexerBackoff(v bool)`

SetSearchDisableIndexerBackoff gets a reference to the given bool and assigns it to the SearchDisableIndexerBackoff field.

### GetSearchDisableIndexingRemovableDrive

`func (o *Windows10GeneralConfiguration) GetSearchDisableIndexingRemovableDrive() bool`

GetSearchDisableIndexingRemovableDrive returns the SearchDisableIndexingRemovableDrive field if non-nil, zero value otherwise.

### GetSearchDisableIndexingRemovableDriveOk

`func (o *Windows10GeneralConfiguration) GetSearchDisableIndexingRemovableDriveOk() (bool, bool)`

GetSearchDisableIndexingRemovableDriveOk returns a tuple with the SearchDisableIndexingRemovableDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchDisableIndexingRemovableDrive

`func (o *Windows10GeneralConfiguration) HasSearchDisableIndexingRemovableDrive() bool`

HasSearchDisableIndexingRemovableDrive returns a boolean if a field has been set.

### SetSearchDisableIndexingRemovableDrive

`func (o *Windows10GeneralConfiguration) SetSearchDisableIndexingRemovableDrive(v bool)`

SetSearchDisableIndexingRemovableDrive gets a reference to the given bool and assigns it to the SearchDisableIndexingRemovableDrive field.

### GetSearchEnableAutomaticIndexSizeManangement

`func (o *Windows10GeneralConfiguration) GetSearchEnableAutomaticIndexSizeManangement() bool`

GetSearchEnableAutomaticIndexSizeManangement returns the SearchEnableAutomaticIndexSizeManangement field if non-nil, zero value otherwise.

### GetSearchEnableAutomaticIndexSizeManangementOk

`func (o *Windows10GeneralConfiguration) GetSearchEnableAutomaticIndexSizeManangementOk() (bool, bool)`

GetSearchEnableAutomaticIndexSizeManangementOk returns a tuple with the SearchEnableAutomaticIndexSizeManangement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchEnableAutomaticIndexSizeManangement

`func (o *Windows10GeneralConfiguration) HasSearchEnableAutomaticIndexSizeManangement() bool`

HasSearchEnableAutomaticIndexSizeManangement returns a boolean if a field has been set.

### SetSearchEnableAutomaticIndexSizeManangement

`func (o *Windows10GeneralConfiguration) SetSearchEnableAutomaticIndexSizeManangement(v bool)`

SetSearchEnableAutomaticIndexSizeManangement gets a reference to the given bool and assigns it to the SearchEnableAutomaticIndexSizeManangement field.

### GetDiagnosticsDataSubmissionMode

`func (o *Windows10GeneralConfiguration) GetDiagnosticsDataSubmissionMode() AnyOfmicrosoftGraphDiagnosticDataSubmissionMode`

GetDiagnosticsDataSubmissionMode returns the DiagnosticsDataSubmissionMode field if non-nil, zero value otherwise.

### GetDiagnosticsDataSubmissionModeOk

`func (o *Windows10GeneralConfiguration) GetDiagnosticsDataSubmissionModeOk() (AnyOfmicrosoftGraphDiagnosticDataSubmissionMode, bool)`

GetDiagnosticsDataSubmissionModeOk returns a tuple with the DiagnosticsDataSubmissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticsDataSubmissionMode

`func (o *Windows10GeneralConfiguration) HasDiagnosticsDataSubmissionMode() bool`

HasDiagnosticsDataSubmissionMode returns a boolean if a field has been set.

### SetDiagnosticsDataSubmissionMode

`func (o *Windows10GeneralConfiguration) SetDiagnosticsDataSubmissionMode(v AnyOfmicrosoftGraphDiagnosticDataSubmissionMode)`

SetDiagnosticsDataSubmissionMode gets a reference to the given AnyOfmicrosoftGraphDiagnosticDataSubmissionMode and assigns it to the DiagnosticsDataSubmissionMode field.

### GetOneDriveDisableFileSync

`func (o *Windows10GeneralConfiguration) GetOneDriveDisableFileSync() bool`

GetOneDriveDisableFileSync returns the OneDriveDisableFileSync field if non-nil, zero value otherwise.

### GetOneDriveDisableFileSyncOk

`func (o *Windows10GeneralConfiguration) GetOneDriveDisableFileSyncOk() (bool, bool)`

GetOneDriveDisableFileSyncOk returns a tuple with the OneDriveDisableFileSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOneDriveDisableFileSync

`func (o *Windows10GeneralConfiguration) HasOneDriveDisableFileSync() bool`

HasOneDriveDisableFileSync returns a boolean if a field has been set.

### SetOneDriveDisableFileSync

`func (o *Windows10GeneralConfiguration) SetOneDriveDisableFileSync(v bool)`

SetOneDriveDisableFileSync gets a reference to the given bool and assigns it to the OneDriveDisableFileSync field.

### GetSmartScreenEnableAppInstallControl

`func (o *Windows10GeneralConfiguration) GetSmartScreenEnableAppInstallControl() bool`

GetSmartScreenEnableAppInstallControl returns the SmartScreenEnableAppInstallControl field if non-nil, zero value otherwise.

### GetSmartScreenEnableAppInstallControlOk

`func (o *Windows10GeneralConfiguration) GetSmartScreenEnableAppInstallControlOk() (bool, bool)`

GetSmartScreenEnableAppInstallControlOk returns a tuple with the SmartScreenEnableAppInstallControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenEnableAppInstallControl

`func (o *Windows10GeneralConfiguration) HasSmartScreenEnableAppInstallControl() bool`

HasSmartScreenEnableAppInstallControl returns a boolean if a field has been set.

### SetSmartScreenEnableAppInstallControl

`func (o *Windows10GeneralConfiguration) SetSmartScreenEnableAppInstallControl(v bool)`

SetSmartScreenEnableAppInstallControl gets a reference to the given bool and assigns it to the SmartScreenEnableAppInstallControl field.

### GetPersonalizationDesktopImageUrl

`func (o *Windows10GeneralConfiguration) GetPersonalizationDesktopImageUrl() string`

GetPersonalizationDesktopImageUrl returns the PersonalizationDesktopImageUrl field if non-nil, zero value otherwise.

### GetPersonalizationDesktopImageUrlOk

`func (o *Windows10GeneralConfiguration) GetPersonalizationDesktopImageUrlOk() (string, bool)`

GetPersonalizationDesktopImageUrlOk returns a tuple with the PersonalizationDesktopImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalizationDesktopImageUrl

`func (o *Windows10GeneralConfiguration) HasPersonalizationDesktopImageUrl() bool`

HasPersonalizationDesktopImageUrl returns a boolean if a field has been set.

### SetPersonalizationDesktopImageUrl

`func (o *Windows10GeneralConfiguration) SetPersonalizationDesktopImageUrl(v string)`

SetPersonalizationDesktopImageUrl gets a reference to the given string and assigns it to the PersonalizationDesktopImageUrl field.

### SetPersonalizationDesktopImageUrlExplicitNull

`func (o *Windows10GeneralConfiguration) SetPersonalizationDesktopImageUrlExplicitNull(b bool)`

SetPersonalizationDesktopImageUrlExplicitNull (un)sets PersonalizationDesktopImageUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonalizationDesktopImageUrl value is set to nil even if false is passed
### GetPersonalizationLockScreenImageUrl

`func (o *Windows10GeneralConfiguration) GetPersonalizationLockScreenImageUrl() string`

GetPersonalizationLockScreenImageUrl returns the PersonalizationLockScreenImageUrl field if non-nil, zero value otherwise.

### GetPersonalizationLockScreenImageUrlOk

`func (o *Windows10GeneralConfiguration) GetPersonalizationLockScreenImageUrlOk() (string, bool)`

GetPersonalizationLockScreenImageUrlOk returns a tuple with the PersonalizationLockScreenImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalizationLockScreenImageUrl

`func (o *Windows10GeneralConfiguration) HasPersonalizationLockScreenImageUrl() bool`

HasPersonalizationLockScreenImageUrl returns a boolean if a field has been set.

### SetPersonalizationLockScreenImageUrl

`func (o *Windows10GeneralConfiguration) SetPersonalizationLockScreenImageUrl(v string)`

SetPersonalizationLockScreenImageUrl gets a reference to the given string and assigns it to the PersonalizationLockScreenImageUrl field.

### SetPersonalizationLockScreenImageUrlExplicitNull

`func (o *Windows10GeneralConfiguration) SetPersonalizationLockScreenImageUrlExplicitNull(b bool)`

SetPersonalizationLockScreenImageUrlExplicitNull (un)sets PersonalizationLockScreenImageUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonalizationLockScreenImageUrl value is set to nil even if false is passed
### GetBluetoothAllowedServices

`func (o *Windows10GeneralConfiguration) GetBluetoothAllowedServices() []string`

GetBluetoothAllowedServices returns the BluetoothAllowedServices field if non-nil, zero value otherwise.

### GetBluetoothAllowedServicesOk

`func (o *Windows10GeneralConfiguration) GetBluetoothAllowedServicesOk() ([]string, bool)`

GetBluetoothAllowedServicesOk returns a tuple with the BluetoothAllowedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothAllowedServices

`func (o *Windows10GeneralConfiguration) HasBluetoothAllowedServices() bool`

HasBluetoothAllowedServices returns a boolean if a field has been set.

### SetBluetoothAllowedServices

`func (o *Windows10GeneralConfiguration) SetBluetoothAllowedServices(v []string)`

SetBluetoothAllowedServices gets a reference to the given []string and assigns it to the BluetoothAllowedServices field.

### GetBluetoothBlockAdvertising

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockAdvertising() bool`

GetBluetoothBlockAdvertising returns the BluetoothBlockAdvertising field if non-nil, zero value otherwise.

### GetBluetoothBlockAdvertisingOk

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockAdvertisingOk() (bool, bool)`

GetBluetoothBlockAdvertisingOk returns a tuple with the BluetoothBlockAdvertising field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockAdvertising

`func (o *Windows10GeneralConfiguration) HasBluetoothBlockAdvertising() bool`

HasBluetoothBlockAdvertising returns a boolean if a field has been set.

### SetBluetoothBlockAdvertising

`func (o *Windows10GeneralConfiguration) SetBluetoothBlockAdvertising(v bool)`

SetBluetoothBlockAdvertising gets a reference to the given bool and assigns it to the BluetoothBlockAdvertising field.

### GetBluetoothBlockDiscoverableMode

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockDiscoverableMode() bool`

GetBluetoothBlockDiscoverableMode returns the BluetoothBlockDiscoverableMode field if non-nil, zero value otherwise.

### GetBluetoothBlockDiscoverableModeOk

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockDiscoverableModeOk() (bool, bool)`

GetBluetoothBlockDiscoverableModeOk returns a tuple with the BluetoothBlockDiscoverableMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockDiscoverableMode

`func (o *Windows10GeneralConfiguration) HasBluetoothBlockDiscoverableMode() bool`

HasBluetoothBlockDiscoverableMode returns a boolean if a field has been set.

### SetBluetoothBlockDiscoverableMode

`func (o *Windows10GeneralConfiguration) SetBluetoothBlockDiscoverableMode(v bool)`

SetBluetoothBlockDiscoverableMode gets a reference to the given bool and assigns it to the BluetoothBlockDiscoverableMode field.

### GetBluetoothBlockPrePairing

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockPrePairing() bool`

GetBluetoothBlockPrePairing returns the BluetoothBlockPrePairing field if non-nil, zero value otherwise.

### GetBluetoothBlockPrePairingOk

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockPrePairingOk() (bool, bool)`

GetBluetoothBlockPrePairingOk returns a tuple with the BluetoothBlockPrePairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlockPrePairing

`func (o *Windows10GeneralConfiguration) HasBluetoothBlockPrePairing() bool`

HasBluetoothBlockPrePairing returns a boolean if a field has been set.

### SetBluetoothBlockPrePairing

`func (o *Windows10GeneralConfiguration) SetBluetoothBlockPrePairing(v bool)`

SetBluetoothBlockPrePairing gets a reference to the given bool and assigns it to the BluetoothBlockPrePairing field.

### GetEdgeBlockAutofill

`func (o *Windows10GeneralConfiguration) GetEdgeBlockAutofill() bool`

GetEdgeBlockAutofill returns the EdgeBlockAutofill field if non-nil, zero value otherwise.

### GetEdgeBlockAutofillOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockAutofillOk() (bool, bool)`

GetEdgeBlockAutofillOk returns a tuple with the EdgeBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockAutofill

`func (o *Windows10GeneralConfiguration) HasEdgeBlockAutofill() bool`

HasEdgeBlockAutofill returns a boolean if a field has been set.

### SetEdgeBlockAutofill

`func (o *Windows10GeneralConfiguration) SetEdgeBlockAutofill(v bool)`

SetEdgeBlockAutofill gets a reference to the given bool and assigns it to the EdgeBlockAutofill field.

### GetEdgeBlocked

`func (o *Windows10GeneralConfiguration) GetEdgeBlocked() bool`

GetEdgeBlocked returns the EdgeBlocked field if non-nil, zero value otherwise.

### GetEdgeBlockedOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockedOk() (bool, bool)`

GetEdgeBlockedOk returns a tuple with the EdgeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlocked

`func (o *Windows10GeneralConfiguration) HasEdgeBlocked() bool`

HasEdgeBlocked returns a boolean if a field has been set.

### SetEdgeBlocked

`func (o *Windows10GeneralConfiguration) SetEdgeBlocked(v bool)`

SetEdgeBlocked gets a reference to the given bool and assigns it to the EdgeBlocked field.

### GetEdgeCookiePolicy

`func (o *Windows10GeneralConfiguration) GetEdgeCookiePolicy() AnyOfmicrosoftGraphEdgeCookiePolicy`

GetEdgeCookiePolicy returns the EdgeCookiePolicy field if non-nil, zero value otherwise.

### GetEdgeCookiePolicyOk

`func (o *Windows10GeneralConfiguration) GetEdgeCookiePolicyOk() (AnyOfmicrosoftGraphEdgeCookiePolicy, bool)`

GetEdgeCookiePolicyOk returns a tuple with the EdgeCookiePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeCookiePolicy

`func (o *Windows10GeneralConfiguration) HasEdgeCookiePolicy() bool`

HasEdgeCookiePolicy returns a boolean if a field has been set.

### SetEdgeCookiePolicy

`func (o *Windows10GeneralConfiguration) SetEdgeCookiePolicy(v AnyOfmicrosoftGraphEdgeCookiePolicy)`

SetEdgeCookiePolicy gets a reference to the given AnyOfmicrosoftGraphEdgeCookiePolicy and assigns it to the EdgeCookiePolicy field.

### GetEdgeBlockDeveloperTools

`func (o *Windows10GeneralConfiguration) GetEdgeBlockDeveloperTools() bool`

GetEdgeBlockDeveloperTools returns the EdgeBlockDeveloperTools field if non-nil, zero value otherwise.

### GetEdgeBlockDeveloperToolsOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockDeveloperToolsOk() (bool, bool)`

GetEdgeBlockDeveloperToolsOk returns a tuple with the EdgeBlockDeveloperTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockDeveloperTools

`func (o *Windows10GeneralConfiguration) HasEdgeBlockDeveloperTools() bool`

HasEdgeBlockDeveloperTools returns a boolean if a field has been set.

### SetEdgeBlockDeveloperTools

`func (o *Windows10GeneralConfiguration) SetEdgeBlockDeveloperTools(v bool)`

SetEdgeBlockDeveloperTools gets a reference to the given bool and assigns it to the EdgeBlockDeveloperTools field.

### GetEdgeBlockSendingDoNotTrackHeader

`func (o *Windows10GeneralConfiguration) GetEdgeBlockSendingDoNotTrackHeader() bool`

GetEdgeBlockSendingDoNotTrackHeader returns the EdgeBlockSendingDoNotTrackHeader field if non-nil, zero value otherwise.

### GetEdgeBlockSendingDoNotTrackHeaderOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockSendingDoNotTrackHeaderOk() (bool, bool)`

GetEdgeBlockSendingDoNotTrackHeaderOk returns a tuple with the EdgeBlockSendingDoNotTrackHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockSendingDoNotTrackHeader

`func (o *Windows10GeneralConfiguration) HasEdgeBlockSendingDoNotTrackHeader() bool`

HasEdgeBlockSendingDoNotTrackHeader returns a boolean if a field has been set.

### SetEdgeBlockSendingDoNotTrackHeader

`func (o *Windows10GeneralConfiguration) SetEdgeBlockSendingDoNotTrackHeader(v bool)`

SetEdgeBlockSendingDoNotTrackHeader gets a reference to the given bool and assigns it to the EdgeBlockSendingDoNotTrackHeader field.

### GetEdgeBlockExtensions

`func (o *Windows10GeneralConfiguration) GetEdgeBlockExtensions() bool`

GetEdgeBlockExtensions returns the EdgeBlockExtensions field if non-nil, zero value otherwise.

### GetEdgeBlockExtensionsOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockExtensionsOk() (bool, bool)`

GetEdgeBlockExtensionsOk returns a tuple with the EdgeBlockExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockExtensions

`func (o *Windows10GeneralConfiguration) HasEdgeBlockExtensions() bool`

HasEdgeBlockExtensions returns a boolean if a field has been set.

### SetEdgeBlockExtensions

`func (o *Windows10GeneralConfiguration) SetEdgeBlockExtensions(v bool)`

SetEdgeBlockExtensions gets a reference to the given bool and assigns it to the EdgeBlockExtensions field.

### GetEdgeBlockInPrivateBrowsing

`func (o *Windows10GeneralConfiguration) GetEdgeBlockInPrivateBrowsing() bool`

GetEdgeBlockInPrivateBrowsing returns the EdgeBlockInPrivateBrowsing field if non-nil, zero value otherwise.

### GetEdgeBlockInPrivateBrowsingOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockInPrivateBrowsingOk() (bool, bool)`

GetEdgeBlockInPrivateBrowsingOk returns a tuple with the EdgeBlockInPrivateBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockInPrivateBrowsing

`func (o *Windows10GeneralConfiguration) HasEdgeBlockInPrivateBrowsing() bool`

HasEdgeBlockInPrivateBrowsing returns a boolean if a field has been set.

### SetEdgeBlockInPrivateBrowsing

`func (o *Windows10GeneralConfiguration) SetEdgeBlockInPrivateBrowsing(v bool)`

SetEdgeBlockInPrivateBrowsing gets a reference to the given bool and assigns it to the EdgeBlockInPrivateBrowsing field.

### GetEdgeBlockJavaScript

`func (o *Windows10GeneralConfiguration) GetEdgeBlockJavaScript() bool`

GetEdgeBlockJavaScript returns the EdgeBlockJavaScript field if non-nil, zero value otherwise.

### GetEdgeBlockJavaScriptOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockJavaScriptOk() (bool, bool)`

GetEdgeBlockJavaScriptOk returns a tuple with the EdgeBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockJavaScript

`func (o *Windows10GeneralConfiguration) HasEdgeBlockJavaScript() bool`

HasEdgeBlockJavaScript returns a boolean if a field has been set.

### SetEdgeBlockJavaScript

`func (o *Windows10GeneralConfiguration) SetEdgeBlockJavaScript(v bool)`

SetEdgeBlockJavaScript gets a reference to the given bool and assigns it to the EdgeBlockJavaScript field.

### GetEdgeBlockPasswordManager

`func (o *Windows10GeneralConfiguration) GetEdgeBlockPasswordManager() bool`

GetEdgeBlockPasswordManager returns the EdgeBlockPasswordManager field if non-nil, zero value otherwise.

### GetEdgeBlockPasswordManagerOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockPasswordManagerOk() (bool, bool)`

GetEdgeBlockPasswordManagerOk returns a tuple with the EdgeBlockPasswordManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockPasswordManager

`func (o *Windows10GeneralConfiguration) HasEdgeBlockPasswordManager() bool`

HasEdgeBlockPasswordManager returns a boolean if a field has been set.

### SetEdgeBlockPasswordManager

`func (o *Windows10GeneralConfiguration) SetEdgeBlockPasswordManager(v bool)`

SetEdgeBlockPasswordManager gets a reference to the given bool and assigns it to the EdgeBlockPasswordManager field.

### GetEdgeBlockAddressBarDropdown

`func (o *Windows10GeneralConfiguration) GetEdgeBlockAddressBarDropdown() bool`

GetEdgeBlockAddressBarDropdown returns the EdgeBlockAddressBarDropdown field if non-nil, zero value otherwise.

### GetEdgeBlockAddressBarDropdownOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockAddressBarDropdownOk() (bool, bool)`

GetEdgeBlockAddressBarDropdownOk returns a tuple with the EdgeBlockAddressBarDropdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockAddressBarDropdown

`func (o *Windows10GeneralConfiguration) HasEdgeBlockAddressBarDropdown() bool`

HasEdgeBlockAddressBarDropdown returns a boolean if a field has been set.

### SetEdgeBlockAddressBarDropdown

`func (o *Windows10GeneralConfiguration) SetEdgeBlockAddressBarDropdown(v bool)`

SetEdgeBlockAddressBarDropdown gets a reference to the given bool and assigns it to the EdgeBlockAddressBarDropdown field.

### GetEdgeBlockCompatibilityList

`func (o *Windows10GeneralConfiguration) GetEdgeBlockCompatibilityList() bool`

GetEdgeBlockCompatibilityList returns the EdgeBlockCompatibilityList field if non-nil, zero value otherwise.

### GetEdgeBlockCompatibilityListOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockCompatibilityListOk() (bool, bool)`

GetEdgeBlockCompatibilityListOk returns a tuple with the EdgeBlockCompatibilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockCompatibilityList

`func (o *Windows10GeneralConfiguration) HasEdgeBlockCompatibilityList() bool`

HasEdgeBlockCompatibilityList returns a boolean if a field has been set.

### SetEdgeBlockCompatibilityList

`func (o *Windows10GeneralConfiguration) SetEdgeBlockCompatibilityList(v bool)`

SetEdgeBlockCompatibilityList gets a reference to the given bool and assigns it to the EdgeBlockCompatibilityList field.

### GetEdgeClearBrowsingDataOnExit

`func (o *Windows10GeneralConfiguration) GetEdgeClearBrowsingDataOnExit() bool`

GetEdgeClearBrowsingDataOnExit returns the EdgeClearBrowsingDataOnExit field if non-nil, zero value otherwise.

### GetEdgeClearBrowsingDataOnExitOk

`func (o *Windows10GeneralConfiguration) GetEdgeClearBrowsingDataOnExitOk() (bool, bool)`

GetEdgeClearBrowsingDataOnExitOk returns a tuple with the EdgeClearBrowsingDataOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeClearBrowsingDataOnExit

`func (o *Windows10GeneralConfiguration) HasEdgeClearBrowsingDataOnExit() bool`

HasEdgeClearBrowsingDataOnExit returns a boolean if a field has been set.

### SetEdgeClearBrowsingDataOnExit

`func (o *Windows10GeneralConfiguration) SetEdgeClearBrowsingDataOnExit(v bool)`

SetEdgeClearBrowsingDataOnExit gets a reference to the given bool and assigns it to the EdgeClearBrowsingDataOnExit field.

### GetEdgeAllowStartPagesModification

`func (o *Windows10GeneralConfiguration) GetEdgeAllowStartPagesModification() bool`

GetEdgeAllowStartPagesModification returns the EdgeAllowStartPagesModification field if non-nil, zero value otherwise.

### GetEdgeAllowStartPagesModificationOk

`func (o *Windows10GeneralConfiguration) GetEdgeAllowStartPagesModificationOk() (bool, bool)`

GetEdgeAllowStartPagesModificationOk returns a tuple with the EdgeAllowStartPagesModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeAllowStartPagesModification

`func (o *Windows10GeneralConfiguration) HasEdgeAllowStartPagesModification() bool`

HasEdgeAllowStartPagesModification returns a boolean if a field has been set.

### SetEdgeAllowStartPagesModification

`func (o *Windows10GeneralConfiguration) SetEdgeAllowStartPagesModification(v bool)`

SetEdgeAllowStartPagesModification gets a reference to the given bool and assigns it to the EdgeAllowStartPagesModification field.

### GetEdgeDisableFirstRunPage

`func (o *Windows10GeneralConfiguration) GetEdgeDisableFirstRunPage() bool`

GetEdgeDisableFirstRunPage returns the EdgeDisableFirstRunPage field if non-nil, zero value otherwise.

### GetEdgeDisableFirstRunPageOk

`func (o *Windows10GeneralConfiguration) GetEdgeDisableFirstRunPageOk() (bool, bool)`

GetEdgeDisableFirstRunPageOk returns a tuple with the EdgeDisableFirstRunPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeDisableFirstRunPage

`func (o *Windows10GeneralConfiguration) HasEdgeDisableFirstRunPage() bool`

HasEdgeDisableFirstRunPage returns a boolean if a field has been set.

### SetEdgeDisableFirstRunPage

`func (o *Windows10GeneralConfiguration) SetEdgeDisableFirstRunPage(v bool)`

SetEdgeDisableFirstRunPage gets a reference to the given bool and assigns it to the EdgeDisableFirstRunPage field.

### GetEdgeBlockLiveTileDataCollection

`func (o *Windows10GeneralConfiguration) GetEdgeBlockLiveTileDataCollection() bool`

GetEdgeBlockLiveTileDataCollection returns the EdgeBlockLiveTileDataCollection field if non-nil, zero value otherwise.

### GetEdgeBlockLiveTileDataCollectionOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockLiveTileDataCollectionOk() (bool, bool)`

GetEdgeBlockLiveTileDataCollectionOk returns a tuple with the EdgeBlockLiveTileDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockLiveTileDataCollection

`func (o *Windows10GeneralConfiguration) HasEdgeBlockLiveTileDataCollection() bool`

HasEdgeBlockLiveTileDataCollection returns a boolean if a field has been set.

### SetEdgeBlockLiveTileDataCollection

`func (o *Windows10GeneralConfiguration) SetEdgeBlockLiveTileDataCollection(v bool)`

SetEdgeBlockLiveTileDataCollection gets a reference to the given bool and assigns it to the EdgeBlockLiveTileDataCollection field.

### GetEdgeSyncFavoritesWithInternetExplorer

`func (o *Windows10GeneralConfiguration) GetEdgeSyncFavoritesWithInternetExplorer() bool`

GetEdgeSyncFavoritesWithInternetExplorer returns the EdgeSyncFavoritesWithInternetExplorer field if non-nil, zero value otherwise.

### GetEdgeSyncFavoritesWithInternetExplorerOk

`func (o *Windows10GeneralConfiguration) GetEdgeSyncFavoritesWithInternetExplorerOk() (bool, bool)`

GetEdgeSyncFavoritesWithInternetExplorerOk returns a tuple with the EdgeSyncFavoritesWithInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeSyncFavoritesWithInternetExplorer

`func (o *Windows10GeneralConfiguration) HasEdgeSyncFavoritesWithInternetExplorer() bool`

HasEdgeSyncFavoritesWithInternetExplorer returns a boolean if a field has been set.

### SetEdgeSyncFavoritesWithInternetExplorer

`func (o *Windows10GeneralConfiguration) SetEdgeSyncFavoritesWithInternetExplorer(v bool)`

SetEdgeSyncFavoritesWithInternetExplorer gets a reference to the given bool and assigns it to the EdgeSyncFavoritesWithInternetExplorer field.

### GetCellularBlockDataWhenRoaming

`func (o *Windows10GeneralConfiguration) GetCellularBlockDataWhenRoaming() bool`

GetCellularBlockDataWhenRoaming returns the CellularBlockDataWhenRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataWhenRoamingOk

`func (o *Windows10GeneralConfiguration) GetCellularBlockDataWhenRoamingOk() (bool, bool)`

GetCellularBlockDataWhenRoamingOk returns a tuple with the CellularBlockDataWhenRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataWhenRoaming

`func (o *Windows10GeneralConfiguration) HasCellularBlockDataWhenRoaming() bool`

HasCellularBlockDataWhenRoaming returns a boolean if a field has been set.

### SetCellularBlockDataWhenRoaming

`func (o *Windows10GeneralConfiguration) SetCellularBlockDataWhenRoaming(v bool)`

SetCellularBlockDataWhenRoaming gets a reference to the given bool and assigns it to the CellularBlockDataWhenRoaming field.

### GetCellularBlockVpn

`func (o *Windows10GeneralConfiguration) GetCellularBlockVpn() bool`

GetCellularBlockVpn returns the CellularBlockVpn field if non-nil, zero value otherwise.

### GetCellularBlockVpnOk

`func (o *Windows10GeneralConfiguration) GetCellularBlockVpnOk() (bool, bool)`

GetCellularBlockVpnOk returns a tuple with the CellularBlockVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVpn

`func (o *Windows10GeneralConfiguration) HasCellularBlockVpn() bool`

HasCellularBlockVpn returns a boolean if a field has been set.

### SetCellularBlockVpn

`func (o *Windows10GeneralConfiguration) SetCellularBlockVpn(v bool)`

SetCellularBlockVpn gets a reference to the given bool and assigns it to the CellularBlockVpn field.

### GetCellularBlockVpnWhenRoaming

`func (o *Windows10GeneralConfiguration) GetCellularBlockVpnWhenRoaming() bool`

GetCellularBlockVpnWhenRoaming returns the CellularBlockVpnWhenRoaming field if non-nil, zero value otherwise.

### GetCellularBlockVpnWhenRoamingOk

`func (o *Windows10GeneralConfiguration) GetCellularBlockVpnWhenRoamingOk() (bool, bool)`

GetCellularBlockVpnWhenRoamingOk returns a tuple with the CellularBlockVpnWhenRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockVpnWhenRoaming

`func (o *Windows10GeneralConfiguration) HasCellularBlockVpnWhenRoaming() bool`

HasCellularBlockVpnWhenRoaming returns a boolean if a field has been set.

### SetCellularBlockVpnWhenRoaming

`func (o *Windows10GeneralConfiguration) SetCellularBlockVpnWhenRoaming(v bool)`

SetCellularBlockVpnWhenRoaming gets a reference to the given bool and assigns it to the CellularBlockVpnWhenRoaming field.

### GetDefenderBlockEndUserAccess

`func (o *Windows10GeneralConfiguration) GetDefenderBlockEndUserAccess() bool`

GetDefenderBlockEndUserAccess returns the DefenderBlockEndUserAccess field if non-nil, zero value otherwise.

### GetDefenderBlockEndUserAccessOk

`func (o *Windows10GeneralConfiguration) GetDefenderBlockEndUserAccessOk() (bool, bool)`

GetDefenderBlockEndUserAccessOk returns a tuple with the DefenderBlockEndUserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderBlockEndUserAccess

`func (o *Windows10GeneralConfiguration) HasDefenderBlockEndUserAccess() bool`

HasDefenderBlockEndUserAccess returns a boolean if a field has been set.

### SetDefenderBlockEndUserAccess

`func (o *Windows10GeneralConfiguration) SetDefenderBlockEndUserAccess(v bool)`

SetDefenderBlockEndUserAccess gets a reference to the given bool and assigns it to the DefenderBlockEndUserAccess field.

### GetDefenderDaysBeforeDeletingQuarantinedMalware

`func (o *Windows10GeneralConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalware() int32`

GetDefenderDaysBeforeDeletingQuarantinedMalware returns the DefenderDaysBeforeDeletingQuarantinedMalware field if non-nil, zero value otherwise.

### GetDefenderDaysBeforeDeletingQuarantinedMalwareOk

`func (o *Windows10GeneralConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalwareOk() (int32, bool)`

GetDefenderDaysBeforeDeletingQuarantinedMalwareOk returns a tuple with the DefenderDaysBeforeDeletingQuarantinedMalware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderDaysBeforeDeletingQuarantinedMalware

`func (o *Windows10GeneralConfiguration) HasDefenderDaysBeforeDeletingQuarantinedMalware() bool`

HasDefenderDaysBeforeDeletingQuarantinedMalware returns a boolean if a field has been set.

### SetDefenderDaysBeforeDeletingQuarantinedMalware

`func (o *Windows10GeneralConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalware(v int32)`

SetDefenderDaysBeforeDeletingQuarantinedMalware gets a reference to the given int32 and assigns it to the DefenderDaysBeforeDeletingQuarantinedMalware field.

### SetDefenderDaysBeforeDeletingQuarantinedMalwareExplicitNull

`func (o *Windows10GeneralConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalwareExplicitNull(b bool)`

SetDefenderDaysBeforeDeletingQuarantinedMalwareExplicitNull (un)sets DefenderDaysBeforeDeletingQuarantinedMalware to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderDaysBeforeDeletingQuarantinedMalware value is set to nil even if false is passed
### GetDefenderDetectedMalwareActions

`func (o *Windows10GeneralConfiguration) GetDefenderDetectedMalwareActions() AnyOfmicrosoftGraphDefenderDetectedMalwareActions`

GetDefenderDetectedMalwareActions returns the DefenderDetectedMalwareActions field if non-nil, zero value otherwise.

### GetDefenderDetectedMalwareActionsOk

`func (o *Windows10GeneralConfiguration) GetDefenderDetectedMalwareActionsOk() (AnyOfmicrosoftGraphDefenderDetectedMalwareActions, bool)`

GetDefenderDetectedMalwareActionsOk returns a tuple with the DefenderDetectedMalwareActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderDetectedMalwareActions

`func (o *Windows10GeneralConfiguration) HasDefenderDetectedMalwareActions() bool`

HasDefenderDetectedMalwareActions returns a boolean if a field has been set.

### SetDefenderDetectedMalwareActions

`func (o *Windows10GeneralConfiguration) SetDefenderDetectedMalwareActions(v AnyOfmicrosoftGraphDefenderDetectedMalwareActions)`

SetDefenderDetectedMalwareActions gets a reference to the given AnyOfmicrosoftGraphDefenderDetectedMalwareActions and assigns it to the DefenderDetectedMalwareActions field.

### SetDefenderDetectedMalwareActionsExplicitNull

`func (o *Windows10GeneralConfiguration) SetDefenderDetectedMalwareActionsExplicitNull(b bool)`

SetDefenderDetectedMalwareActionsExplicitNull (un)sets DefenderDetectedMalwareActions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderDetectedMalwareActions value is set to nil even if false is passed
### GetDefenderSystemScanSchedule

`func (o *Windows10GeneralConfiguration) GetDefenderSystemScanSchedule() AnyOfmicrosoftGraphWeeklySchedule`

GetDefenderSystemScanSchedule returns the DefenderSystemScanSchedule field if non-nil, zero value otherwise.

### GetDefenderSystemScanScheduleOk

`func (o *Windows10GeneralConfiguration) GetDefenderSystemScanScheduleOk() (AnyOfmicrosoftGraphWeeklySchedule, bool)`

GetDefenderSystemScanScheduleOk returns a tuple with the DefenderSystemScanSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderSystemScanSchedule

`func (o *Windows10GeneralConfiguration) HasDefenderSystemScanSchedule() bool`

HasDefenderSystemScanSchedule returns a boolean if a field has been set.

### SetDefenderSystemScanSchedule

`func (o *Windows10GeneralConfiguration) SetDefenderSystemScanSchedule(v AnyOfmicrosoftGraphWeeklySchedule)`

SetDefenderSystemScanSchedule gets a reference to the given AnyOfmicrosoftGraphWeeklySchedule and assigns it to the DefenderSystemScanSchedule field.

### GetDefenderFilesAndFoldersToExclude

`func (o *Windows10GeneralConfiguration) GetDefenderFilesAndFoldersToExclude() []string`

GetDefenderFilesAndFoldersToExclude returns the DefenderFilesAndFoldersToExclude field if non-nil, zero value otherwise.

### GetDefenderFilesAndFoldersToExcludeOk

`func (o *Windows10GeneralConfiguration) GetDefenderFilesAndFoldersToExcludeOk() ([]string, bool)`

GetDefenderFilesAndFoldersToExcludeOk returns a tuple with the DefenderFilesAndFoldersToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderFilesAndFoldersToExclude

`func (o *Windows10GeneralConfiguration) HasDefenderFilesAndFoldersToExclude() bool`

HasDefenderFilesAndFoldersToExclude returns a boolean if a field has been set.

### SetDefenderFilesAndFoldersToExclude

`func (o *Windows10GeneralConfiguration) SetDefenderFilesAndFoldersToExclude(v []string)`

SetDefenderFilesAndFoldersToExclude gets a reference to the given []string and assigns it to the DefenderFilesAndFoldersToExclude field.

### GetDefenderFileExtensionsToExclude

`func (o *Windows10GeneralConfiguration) GetDefenderFileExtensionsToExclude() []string`

GetDefenderFileExtensionsToExclude returns the DefenderFileExtensionsToExclude field if non-nil, zero value otherwise.

### GetDefenderFileExtensionsToExcludeOk

`func (o *Windows10GeneralConfiguration) GetDefenderFileExtensionsToExcludeOk() ([]string, bool)`

GetDefenderFileExtensionsToExcludeOk returns a tuple with the DefenderFileExtensionsToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderFileExtensionsToExclude

`func (o *Windows10GeneralConfiguration) HasDefenderFileExtensionsToExclude() bool`

HasDefenderFileExtensionsToExclude returns a boolean if a field has been set.

### SetDefenderFileExtensionsToExclude

`func (o *Windows10GeneralConfiguration) SetDefenderFileExtensionsToExclude(v []string)`

SetDefenderFileExtensionsToExclude gets a reference to the given []string and assigns it to the DefenderFileExtensionsToExclude field.

### GetDefenderScanMaxCpu

`func (o *Windows10GeneralConfiguration) GetDefenderScanMaxCpu() int32`

GetDefenderScanMaxCpu returns the DefenderScanMaxCpu field if non-nil, zero value otherwise.

### GetDefenderScanMaxCpuOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanMaxCpuOk() (int32, bool)`

GetDefenderScanMaxCpuOk returns a tuple with the DefenderScanMaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanMaxCpu

`func (o *Windows10GeneralConfiguration) HasDefenderScanMaxCpu() bool`

HasDefenderScanMaxCpu returns a boolean if a field has been set.

### SetDefenderScanMaxCpu

`func (o *Windows10GeneralConfiguration) SetDefenderScanMaxCpu(v int32)`

SetDefenderScanMaxCpu gets a reference to the given int32 and assigns it to the DefenderScanMaxCpu field.

### SetDefenderScanMaxCpuExplicitNull

`func (o *Windows10GeneralConfiguration) SetDefenderScanMaxCpuExplicitNull(b bool)`

SetDefenderScanMaxCpuExplicitNull (un)sets DefenderScanMaxCpu to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderScanMaxCpu value is set to nil even if false is passed
### GetDefenderMonitorFileActivity

`func (o *Windows10GeneralConfiguration) GetDefenderMonitorFileActivity() AnyOfmicrosoftGraphDefenderMonitorFileActivity`

GetDefenderMonitorFileActivity returns the DefenderMonitorFileActivity field if non-nil, zero value otherwise.

### GetDefenderMonitorFileActivityOk

`func (o *Windows10GeneralConfiguration) GetDefenderMonitorFileActivityOk() (AnyOfmicrosoftGraphDefenderMonitorFileActivity, bool)`

GetDefenderMonitorFileActivityOk returns a tuple with the DefenderMonitorFileActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderMonitorFileActivity

`func (o *Windows10GeneralConfiguration) HasDefenderMonitorFileActivity() bool`

HasDefenderMonitorFileActivity returns a boolean if a field has been set.

### SetDefenderMonitorFileActivity

`func (o *Windows10GeneralConfiguration) SetDefenderMonitorFileActivity(v AnyOfmicrosoftGraphDefenderMonitorFileActivity)`

SetDefenderMonitorFileActivity gets a reference to the given AnyOfmicrosoftGraphDefenderMonitorFileActivity and assigns it to the DefenderMonitorFileActivity field.

### GetDefenderProcessesToExclude

`func (o *Windows10GeneralConfiguration) GetDefenderProcessesToExclude() []string`

GetDefenderProcessesToExclude returns the DefenderProcessesToExclude field if non-nil, zero value otherwise.

### GetDefenderProcessesToExcludeOk

`func (o *Windows10GeneralConfiguration) GetDefenderProcessesToExcludeOk() ([]string, bool)`

GetDefenderProcessesToExcludeOk returns a tuple with the DefenderProcessesToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderProcessesToExclude

`func (o *Windows10GeneralConfiguration) HasDefenderProcessesToExclude() bool`

HasDefenderProcessesToExclude returns a boolean if a field has been set.

### SetDefenderProcessesToExclude

`func (o *Windows10GeneralConfiguration) SetDefenderProcessesToExclude(v []string)`

SetDefenderProcessesToExclude gets a reference to the given []string and assigns it to the DefenderProcessesToExclude field.

### GetDefenderPromptForSampleSubmission

`func (o *Windows10GeneralConfiguration) GetDefenderPromptForSampleSubmission() AnyOfmicrosoftGraphDefenderPromptForSampleSubmission`

GetDefenderPromptForSampleSubmission returns the DefenderPromptForSampleSubmission field if non-nil, zero value otherwise.

### GetDefenderPromptForSampleSubmissionOk

`func (o *Windows10GeneralConfiguration) GetDefenderPromptForSampleSubmissionOk() (AnyOfmicrosoftGraphDefenderPromptForSampleSubmission, bool)`

GetDefenderPromptForSampleSubmissionOk returns a tuple with the DefenderPromptForSampleSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderPromptForSampleSubmission

`func (o *Windows10GeneralConfiguration) HasDefenderPromptForSampleSubmission() bool`

HasDefenderPromptForSampleSubmission returns a boolean if a field has been set.

### SetDefenderPromptForSampleSubmission

`func (o *Windows10GeneralConfiguration) SetDefenderPromptForSampleSubmission(v AnyOfmicrosoftGraphDefenderPromptForSampleSubmission)`

SetDefenderPromptForSampleSubmission gets a reference to the given AnyOfmicrosoftGraphDefenderPromptForSampleSubmission and assigns it to the DefenderPromptForSampleSubmission field.

### GetDefenderRequireBehaviorMonitoring

`func (o *Windows10GeneralConfiguration) GetDefenderRequireBehaviorMonitoring() bool`

GetDefenderRequireBehaviorMonitoring returns the DefenderRequireBehaviorMonitoring field if non-nil, zero value otherwise.

### GetDefenderRequireBehaviorMonitoringOk

`func (o *Windows10GeneralConfiguration) GetDefenderRequireBehaviorMonitoringOk() (bool, bool)`

GetDefenderRequireBehaviorMonitoringOk returns a tuple with the DefenderRequireBehaviorMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireBehaviorMonitoring

`func (o *Windows10GeneralConfiguration) HasDefenderRequireBehaviorMonitoring() bool`

HasDefenderRequireBehaviorMonitoring returns a boolean if a field has been set.

### SetDefenderRequireBehaviorMonitoring

`func (o *Windows10GeneralConfiguration) SetDefenderRequireBehaviorMonitoring(v bool)`

SetDefenderRequireBehaviorMonitoring gets a reference to the given bool and assigns it to the DefenderRequireBehaviorMonitoring field.

### GetDefenderRequireCloudProtection

`func (o *Windows10GeneralConfiguration) GetDefenderRequireCloudProtection() bool`

GetDefenderRequireCloudProtection returns the DefenderRequireCloudProtection field if non-nil, zero value otherwise.

### GetDefenderRequireCloudProtectionOk

`func (o *Windows10GeneralConfiguration) GetDefenderRequireCloudProtectionOk() (bool, bool)`

GetDefenderRequireCloudProtectionOk returns a tuple with the DefenderRequireCloudProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireCloudProtection

`func (o *Windows10GeneralConfiguration) HasDefenderRequireCloudProtection() bool`

HasDefenderRequireCloudProtection returns a boolean if a field has been set.

### SetDefenderRequireCloudProtection

`func (o *Windows10GeneralConfiguration) SetDefenderRequireCloudProtection(v bool)`

SetDefenderRequireCloudProtection gets a reference to the given bool and assigns it to the DefenderRequireCloudProtection field.

### GetDefenderRequireNetworkInspectionSystem

`func (o *Windows10GeneralConfiguration) GetDefenderRequireNetworkInspectionSystem() bool`

GetDefenderRequireNetworkInspectionSystem returns the DefenderRequireNetworkInspectionSystem field if non-nil, zero value otherwise.

### GetDefenderRequireNetworkInspectionSystemOk

`func (o *Windows10GeneralConfiguration) GetDefenderRequireNetworkInspectionSystemOk() (bool, bool)`

GetDefenderRequireNetworkInspectionSystemOk returns a tuple with the DefenderRequireNetworkInspectionSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireNetworkInspectionSystem

`func (o *Windows10GeneralConfiguration) HasDefenderRequireNetworkInspectionSystem() bool`

HasDefenderRequireNetworkInspectionSystem returns a boolean if a field has been set.

### SetDefenderRequireNetworkInspectionSystem

`func (o *Windows10GeneralConfiguration) SetDefenderRequireNetworkInspectionSystem(v bool)`

SetDefenderRequireNetworkInspectionSystem gets a reference to the given bool and assigns it to the DefenderRequireNetworkInspectionSystem field.

### GetDefenderRequireRealTimeMonitoring

`func (o *Windows10GeneralConfiguration) GetDefenderRequireRealTimeMonitoring() bool`

GetDefenderRequireRealTimeMonitoring returns the DefenderRequireRealTimeMonitoring field if non-nil, zero value otherwise.

### GetDefenderRequireRealTimeMonitoringOk

`func (o *Windows10GeneralConfiguration) GetDefenderRequireRealTimeMonitoringOk() (bool, bool)`

GetDefenderRequireRealTimeMonitoringOk returns a tuple with the DefenderRequireRealTimeMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderRequireRealTimeMonitoring

`func (o *Windows10GeneralConfiguration) HasDefenderRequireRealTimeMonitoring() bool`

HasDefenderRequireRealTimeMonitoring returns a boolean if a field has been set.

### SetDefenderRequireRealTimeMonitoring

`func (o *Windows10GeneralConfiguration) SetDefenderRequireRealTimeMonitoring(v bool)`

SetDefenderRequireRealTimeMonitoring gets a reference to the given bool and assigns it to the DefenderRequireRealTimeMonitoring field.

### GetDefenderScanArchiveFiles

`func (o *Windows10GeneralConfiguration) GetDefenderScanArchiveFiles() bool`

GetDefenderScanArchiveFiles returns the DefenderScanArchiveFiles field if non-nil, zero value otherwise.

### GetDefenderScanArchiveFilesOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanArchiveFilesOk() (bool, bool)`

GetDefenderScanArchiveFilesOk returns a tuple with the DefenderScanArchiveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanArchiveFiles

`func (o *Windows10GeneralConfiguration) HasDefenderScanArchiveFiles() bool`

HasDefenderScanArchiveFiles returns a boolean if a field has been set.

### SetDefenderScanArchiveFiles

`func (o *Windows10GeneralConfiguration) SetDefenderScanArchiveFiles(v bool)`

SetDefenderScanArchiveFiles gets a reference to the given bool and assigns it to the DefenderScanArchiveFiles field.

### GetDefenderScanDownloads

`func (o *Windows10GeneralConfiguration) GetDefenderScanDownloads() bool`

GetDefenderScanDownloads returns the DefenderScanDownloads field if non-nil, zero value otherwise.

### GetDefenderScanDownloadsOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanDownloadsOk() (bool, bool)`

GetDefenderScanDownloadsOk returns a tuple with the DefenderScanDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanDownloads

`func (o *Windows10GeneralConfiguration) HasDefenderScanDownloads() bool`

HasDefenderScanDownloads returns a boolean if a field has been set.

### SetDefenderScanDownloads

`func (o *Windows10GeneralConfiguration) SetDefenderScanDownloads(v bool)`

SetDefenderScanDownloads gets a reference to the given bool and assigns it to the DefenderScanDownloads field.

### GetDefenderScanNetworkFiles

`func (o *Windows10GeneralConfiguration) GetDefenderScanNetworkFiles() bool`

GetDefenderScanNetworkFiles returns the DefenderScanNetworkFiles field if non-nil, zero value otherwise.

### GetDefenderScanNetworkFilesOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanNetworkFilesOk() (bool, bool)`

GetDefenderScanNetworkFilesOk returns a tuple with the DefenderScanNetworkFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanNetworkFiles

`func (o *Windows10GeneralConfiguration) HasDefenderScanNetworkFiles() bool`

HasDefenderScanNetworkFiles returns a boolean if a field has been set.

### SetDefenderScanNetworkFiles

`func (o *Windows10GeneralConfiguration) SetDefenderScanNetworkFiles(v bool)`

SetDefenderScanNetworkFiles gets a reference to the given bool and assigns it to the DefenderScanNetworkFiles field.

### GetDefenderScanIncomingMail

`func (o *Windows10GeneralConfiguration) GetDefenderScanIncomingMail() bool`

GetDefenderScanIncomingMail returns the DefenderScanIncomingMail field if non-nil, zero value otherwise.

### GetDefenderScanIncomingMailOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanIncomingMailOk() (bool, bool)`

GetDefenderScanIncomingMailOk returns a tuple with the DefenderScanIncomingMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanIncomingMail

`func (o *Windows10GeneralConfiguration) HasDefenderScanIncomingMail() bool`

HasDefenderScanIncomingMail returns a boolean if a field has been set.

### SetDefenderScanIncomingMail

`func (o *Windows10GeneralConfiguration) SetDefenderScanIncomingMail(v bool)`

SetDefenderScanIncomingMail gets a reference to the given bool and assigns it to the DefenderScanIncomingMail field.

### GetDefenderScanMappedNetworkDrivesDuringFullScan

`func (o *Windows10GeneralConfiguration) GetDefenderScanMappedNetworkDrivesDuringFullScan() bool`

GetDefenderScanMappedNetworkDrivesDuringFullScan returns the DefenderScanMappedNetworkDrivesDuringFullScan field if non-nil, zero value otherwise.

### GetDefenderScanMappedNetworkDrivesDuringFullScanOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanMappedNetworkDrivesDuringFullScanOk() (bool, bool)`

GetDefenderScanMappedNetworkDrivesDuringFullScanOk returns a tuple with the DefenderScanMappedNetworkDrivesDuringFullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanMappedNetworkDrivesDuringFullScan

`func (o *Windows10GeneralConfiguration) HasDefenderScanMappedNetworkDrivesDuringFullScan() bool`

HasDefenderScanMappedNetworkDrivesDuringFullScan returns a boolean if a field has been set.

### SetDefenderScanMappedNetworkDrivesDuringFullScan

`func (o *Windows10GeneralConfiguration) SetDefenderScanMappedNetworkDrivesDuringFullScan(v bool)`

SetDefenderScanMappedNetworkDrivesDuringFullScan gets a reference to the given bool and assigns it to the DefenderScanMappedNetworkDrivesDuringFullScan field.

### GetDefenderScanRemovableDrivesDuringFullScan

`func (o *Windows10GeneralConfiguration) GetDefenderScanRemovableDrivesDuringFullScan() bool`

GetDefenderScanRemovableDrivesDuringFullScan returns the DefenderScanRemovableDrivesDuringFullScan field if non-nil, zero value otherwise.

### GetDefenderScanRemovableDrivesDuringFullScanOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanRemovableDrivesDuringFullScanOk() (bool, bool)`

GetDefenderScanRemovableDrivesDuringFullScanOk returns a tuple with the DefenderScanRemovableDrivesDuringFullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanRemovableDrivesDuringFullScan

`func (o *Windows10GeneralConfiguration) HasDefenderScanRemovableDrivesDuringFullScan() bool`

HasDefenderScanRemovableDrivesDuringFullScan returns a boolean if a field has been set.

### SetDefenderScanRemovableDrivesDuringFullScan

`func (o *Windows10GeneralConfiguration) SetDefenderScanRemovableDrivesDuringFullScan(v bool)`

SetDefenderScanRemovableDrivesDuringFullScan gets a reference to the given bool and assigns it to the DefenderScanRemovableDrivesDuringFullScan field.

### GetDefenderScanScriptsLoadedInInternetExplorer

`func (o *Windows10GeneralConfiguration) GetDefenderScanScriptsLoadedInInternetExplorer() bool`

GetDefenderScanScriptsLoadedInInternetExplorer returns the DefenderScanScriptsLoadedInInternetExplorer field if non-nil, zero value otherwise.

### GetDefenderScanScriptsLoadedInInternetExplorerOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanScriptsLoadedInInternetExplorerOk() (bool, bool)`

GetDefenderScanScriptsLoadedInInternetExplorerOk returns a tuple with the DefenderScanScriptsLoadedInInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanScriptsLoadedInInternetExplorer

`func (o *Windows10GeneralConfiguration) HasDefenderScanScriptsLoadedInInternetExplorer() bool`

HasDefenderScanScriptsLoadedInInternetExplorer returns a boolean if a field has been set.

### SetDefenderScanScriptsLoadedInInternetExplorer

`func (o *Windows10GeneralConfiguration) SetDefenderScanScriptsLoadedInInternetExplorer(v bool)`

SetDefenderScanScriptsLoadedInInternetExplorer gets a reference to the given bool and assigns it to the DefenderScanScriptsLoadedInInternetExplorer field.

### GetDefenderSignatureUpdateIntervalInHours

`func (o *Windows10GeneralConfiguration) GetDefenderSignatureUpdateIntervalInHours() int32`

GetDefenderSignatureUpdateIntervalInHours returns the DefenderSignatureUpdateIntervalInHours field if non-nil, zero value otherwise.

### GetDefenderSignatureUpdateIntervalInHoursOk

`func (o *Windows10GeneralConfiguration) GetDefenderSignatureUpdateIntervalInHoursOk() (int32, bool)`

GetDefenderSignatureUpdateIntervalInHoursOk returns a tuple with the DefenderSignatureUpdateIntervalInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderSignatureUpdateIntervalInHours

`func (o *Windows10GeneralConfiguration) HasDefenderSignatureUpdateIntervalInHours() bool`

HasDefenderSignatureUpdateIntervalInHours returns a boolean if a field has been set.

### SetDefenderSignatureUpdateIntervalInHours

`func (o *Windows10GeneralConfiguration) SetDefenderSignatureUpdateIntervalInHours(v int32)`

SetDefenderSignatureUpdateIntervalInHours gets a reference to the given int32 and assigns it to the DefenderSignatureUpdateIntervalInHours field.

### SetDefenderSignatureUpdateIntervalInHoursExplicitNull

`func (o *Windows10GeneralConfiguration) SetDefenderSignatureUpdateIntervalInHoursExplicitNull(b bool)`

SetDefenderSignatureUpdateIntervalInHoursExplicitNull (un)sets DefenderSignatureUpdateIntervalInHours to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderSignatureUpdateIntervalInHours value is set to nil even if false is passed
### GetDefenderScanType

`func (o *Windows10GeneralConfiguration) GetDefenderScanType() AnyOfmicrosoftGraphDefenderScanType`

GetDefenderScanType returns the DefenderScanType field if non-nil, zero value otherwise.

### GetDefenderScanTypeOk

`func (o *Windows10GeneralConfiguration) GetDefenderScanTypeOk() (AnyOfmicrosoftGraphDefenderScanType, bool)`

GetDefenderScanTypeOk returns a tuple with the DefenderScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScanType

`func (o *Windows10GeneralConfiguration) HasDefenderScanType() bool`

HasDefenderScanType returns a boolean if a field has been set.

### SetDefenderScanType

`func (o *Windows10GeneralConfiguration) SetDefenderScanType(v AnyOfmicrosoftGraphDefenderScanType)`

SetDefenderScanType gets a reference to the given AnyOfmicrosoftGraphDefenderScanType and assigns it to the DefenderScanType field.

### GetDefenderScheduledScanTime

`func (o *Windows10GeneralConfiguration) GetDefenderScheduledScanTime() string`

GetDefenderScheduledScanTime returns the DefenderScheduledScanTime field if non-nil, zero value otherwise.

### GetDefenderScheduledScanTimeOk

`func (o *Windows10GeneralConfiguration) GetDefenderScheduledScanTimeOk() (string, bool)`

GetDefenderScheduledScanTimeOk returns a tuple with the DefenderScheduledScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScheduledScanTime

`func (o *Windows10GeneralConfiguration) HasDefenderScheduledScanTime() bool`

HasDefenderScheduledScanTime returns a boolean if a field has been set.

### SetDefenderScheduledScanTime

`func (o *Windows10GeneralConfiguration) SetDefenderScheduledScanTime(v string)`

SetDefenderScheduledScanTime gets a reference to the given string and assigns it to the DefenderScheduledScanTime field.

### SetDefenderScheduledScanTimeExplicitNull

`func (o *Windows10GeneralConfiguration) SetDefenderScheduledScanTimeExplicitNull(b bool)`

SetDefenderScheduledScanTimeExplicitNull (un)sets DefenderScheduledScanTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderScheduledScanTime value is set to nil even if false is passed
### GetDefenderScheduledQuickScanTime

`func (o *Windows10GeneralConfiguration) GetDefenderScheduledQuickScanTime() string`

GetDefenderScheduledQuickScanTime returns the DefenderScheduledQuickScanTime field if non-nil, zero value otherwise.

### GetDefenderScheduledQuickScanTimeOk

`func (o *Windows10GeneralConfiguration) GetDefenderScheduledQuickScanTimeOk() (string, bool)`

GetDefenderScheduledQuickScanTimeOk returns a tuple with the DefenderScheduledQuickScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderScheduledQuickScanTime

`func (o *Windows10GeneralConfiguration) HasDefenderScheduledQuickScanTime() bool`

HasDefenderScheduledQuickScanTime returns a boolean if a field has been set.

### SetDefenderScheduledQuickScanTime

`func (o *Windows10GeneralConfiguration) SetDefenderScheduledQuickScanTime(v string)`

SetDefenderScheduledQuickScanTime gets a reference to the given string and assigns it to the DefenderScheduledQuickScanTime field.

### SetDefenderScheduledQuickScanTimeExplicitNull

`func (o *Windows10GeneralConfiguration) SetDefenderScheduledQuickScanTimeExplicitNull(b bool)`

SetDefenderScheduledQuickScanTimeExplicitNull (un)sets DefenderScheduledQuickScanTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderScheduledQuickScanTime value is set to nil even if false is passed
### GetDefenderCloudBlockLevel

`func (o *Windows10GeneralConfiguration) GetDefenderCloudBlockLevel() AnyOfmicrosoftGraphDefenderCloudBlockLevelType`

GetDefenderCloudBlockLevel returns the DefenderCloudBlockLevel field if non-nil, zero value otherwise.

### GetDefenderCloudBlockLevelOk

`func (o *Windows10GeneralConfiguration) GetDefenderCloudBlockLevelOk() (AnyOfmicrosoftGraphDefenderCloudBlockLevelType, bool)`

GetDefenderCloudBlockLevelOk returns a tuple with the DefenderCloudBlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderCloudBlockLevel

`func (o *Windows10GeneralConfiguration) HasDefenderCloudBlockLevel() bool`

HasDefenderCloudBlockLevel returns a boolean if a field has been set.

### SetDefenderCloudBlockLevel

`func (o *Windows10GeneralConfiguration) SetDefenderCloudBlockLevel(v AnyOfmicrosoftGraphDefenderCloudBlockLevelType)`

SetDefenderCloudBlockLevel gets a reference to the given AnyOfmicrosoftGraphDefenderCloudBlockLevelType and assigns it to the DefenderCloudBlockLevel field.

### GetLockScreenAllowTimeoutConfiguration

`func (o *Windows10GeneralConfiguration) GetLockScreenAllowTimeoutConfiguration() bool`

GetLockScreenAllowTimeoutConfiguration returns the LockScreenAllowTimeoutConfiguration field if non-nil, zero value otherwise.

### GetLockScreenAllowTimeoutConfigurationOk

`func (o *Windows10GeneralConfiguration) GetLockScreenAllowTimeoutConfigurationOk() (bool, bool)`

GetLockScreenAllowTimeoutConfigurationOk returns a tuple with the LockScreenAllowTimeoutConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenAllowTimeoutConfiguration

`func (o *Windows10GeneralConfiguration) HasLockScreenAllowTimeoutConfiguration() bool`

HasLockScreenAllowTimeoutConfiguration returns a boolean if a field has been set.

### SetLockScreenAllowTimeoutConfiguration

`func (o *Windows10GeneralConfiguration) SetLockScreenAllowTimeoutConfiguration(v bool)`

SetLockScreenAllowTimeoutConfiguration gets a reference to the given bool and assigns it to the LockScreenAllowTimeoutConfiguration field.

### GetLockScreenBlockActionCenterNotifications

`func (o *Windows10GeneralConfiguration) GetLockScreenBlockActionCenterNotifications() bool`

GetLockScreenBlockActionCenterNotifications returns the LockScreenBlockActionCenterNotifications field if non-nil, zero value otherwise.

### GetLockScreenBlockActionCenterNotificationsOk

`func (o *Windows10GeneralConfiguration) GetLockScreenBlockActionCenterNotificationsOk() (bool, bool)`

GetLockScreenBlockActionCenterNotificationsOk returns a tuple with the LockScreenBlockActionCenterNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockActionCenterNotifications

`func (o *Windows10GeneralConfiguration) HasLockScreenBlockActionCenterNotifications() bool`

HasLockScreenBlockActionCenterNotifications returns a boolean if a field has been set.

### SetLockScreenBlockActionCenterNotifications

`func (o *Windows10GeneralConfiguration) SetLockScreenBlockActionCenterNotifications(v bool)`

SetLockScreenBlockActionCenterNotifications gets a reference to the given bool and assigns it to the LockScreenBlockActionCenterNotifications field.

### GetLockScreenBlockCortana

`func (o *Windows10GeneralConfiguration) GetLockScreenBlockCortana() bool`

GetLockScreenBlockCortana returns the LockScreenBlockCortana field if non-nil, zero value otherwise.

### GetLockScreenBlockCortanaOk

`func (o *Windows10GeneralConfiguration) GetLockScreenBlockCortanaOk() (bool, bool)`

GetLockScreenBlockCortanaOk returns a tuple with the LockScreenBlockCortana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockCortana

`func (o *Windows10GeneralConfiguration) HasLockScreenBlockCortana() bool`

HasLockScreenBlockCortana returns a boolean if a field has been set.

### SetLockScreenBlockCortana

`func (o *Windows10GeneralConfiguration) SetLockScreenBlockCortana(v bool)`

SetLockScreenBlockCortana gets a reference to the given bool and assigns it to the LockScreenBlockCortana field.

### GetLockScreenBlockToastNotifications

`func (o *Windows10GeneralConfiguration) GetLockScreenBlockToastNotifications() bool`

GetLockScreenBlockToastNotifications returns the LockScreenBlockToastNotifications field if non-nil, zero value otherwise.

### GetLockScreenBlockToastNotificationsOk

`func (o *Windows10GeneralConfiguration) GetLockScreenBlockToastNotificationsOk() (bool, bool)`

GetLockScreenBlockToastNotificationsOk returns a tuple with the LockScreenBlockToastNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenBlockToastNotifications

`func (o *Windows10GeneralConfiguration) HasLockScreenBlockToastNotifications() bool`

HasLockScreenBlockToastNotifications returns a boolean if a field has been set.

### SetLockScreenBlockToastNotifications

`func (o *Windows10GeneralConfiguration) SetLockScreenBlockToastNotifications(v bool)`

SetLockScreenBlockToastNotifications gets a reference to the given bool and assigns it to the LockScreenBlockToastNotifications field.

### GetLockScreenTimeoutInSeconds

`func (o *Windows10GeneralConfiguration) GetLockScreenTimeoutInSeconds() int32`

GetLockScreenTimeoutInSeconds returns the LockScreenTimeoutInSeconds field if non-nil, zero value otherwise.

### GetLockScreenTimeoutInSecondsOk

`func (o *Windows10GeneralConfiguration) GetLockScreenTimeoutInSecondsOk() (int32, bool)`

GetLockScreenTimeoutInSecondsOk returns a tuple with the LockScreenTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenTimeoutInSeconds

`func (o *Windows10GeneralConfiguration) HasLockScreenTimeoutInSeconds() bool`

HasLockScreenTimeoutInSeconds returns a boolean if a field has been set.

### SetLockScreenTimeoutInSeconds

`func (o *Windows10GeneralConfiguration) SetLockScreenTimeoutInSeconds(v int32)`

SetLockScreenTimeoutInSeconds gets a reference to the given int32 and assigns it to the LockScreenTimeoutInSeconds field.

### SetLockScreenTimeoutInSecondsExplicitNull

`func (o *Windows10GeneralConfiguration) SetLockScreenTimeoutInSecondsExplicitNull(b bool)`

SetLockScreenTimeoutInSecondsExplicitNull (un)sets LockScreenTimeoutInSeconds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LockScreenTimeoutInSeconds value is set to nil even if false is passed
### GetPasswordBlockSimple

`func (o *Windows10GeneralConfiguration) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *Windows10GeneralConfiguration) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *Windows10GeneralConfiguration) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *Windows10GeneralConfiguration) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *Windows10GeneralConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *Windows10GeneralConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *Windows10GeneralConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *Windows10GeneralConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *Windows10GeneralConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *Windows10GeneralConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *Windows10GeneralConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *Windows10GeneralConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *Windows10GeneralConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *Windows10GeneralConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *Windows10GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *Windows10GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *Windows10GeneralConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *Windows10GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *Windows10GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *Windows10GeneralConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *Windows10GeneralConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *Windows10GeneralConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *Windows10GeneralConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *Windows10GeneralConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *Windows10GeneralConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *Windows10GeneralConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *Windows10GeneralConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *Windows10GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *Windows10GeneralConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordRequired

`func (o *Windows10GeneralConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *Windows10GeneralConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *Windows10GeneralConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *Windows10GeneralConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordRequireWhenResumeFromIdleState

`func (o *Windows10GeneralConfiguration) GetPasswordRequireWhenResumeFromIdleState() bool`

GetPasswordRequireWhenResumeFromIdleState returns the PasswordRequireWhenResumeFromIdleState field if non-nil, zero value otherwise.

### GetPasswordRequireWhenResumeFromIdleStateOk

`func (o *Windows10GeneralConfiguration) GetPasswordRequireWhenResumeFromIdleStateOk() (bool, bool)`

GetPasswordRequireWhenResumeFromIdleStateOk returns a tuple with the PasswordRequireWhenResumeFromIdleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequireWhenResumeFromIdleState

`func (o *Windows10GeneralConfiguration) HasPasswordRequireWhenResumeFromIdleState() bool`

HasPasswordRequireWhenResumeFromIdleState returns a boolean if a field has been set.

### SetPasswordRequireWhenResumeFromIdleState

`func (o *Windows10GeneralConfiguration) SetPasswordRequireWhenResumeFromIdleState(v bool)`

SetPasswordRequireWhenResumeFromIdleState gets a reference to the given bool and assigns it to the PasswordRequireWhenResumeFromIdleState field.

### GetPasswordRequiredType

`func (o *Windows10GeneralConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *Windows10GeneralConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *Windows10GeneralConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *Windows10GeneralConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *Windows10GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *Windows10GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *Windows10GeneralConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *Windows10GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *Windows10GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPrivacyAdvertisingId

`func (o *Windows10GeneralConfiguration) GetPrivacyAdvertisingId() AnyOfmicrosoftGraphStateManagementSetting`

GetPrivacyAdvertisingId returns the PrivacyAdvertisingId field if non-nil, zero value otherwise.

### GetPrivacyAdvertisingIdOk

`func (o *Windows10GeneralConfiguration) GetPrivacyAdvertisingIdOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetPrivacyAdvertisingIdOk returns a tuple with the PrivacyAdvertisingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyAdvertisingId

`func (o *Windows10GeneralConfiguration) HasPrivacyAdvertisingId() bool`

HasPrivacyAdvertisingId returns a boolean if a field has been set.

### SetPrivacyAdvertisingId

`func (o *Windows10GeneralConfiguration) SetPrivacyAdvertisingId(v AnyOfmicrosoftGraphStateManagementSetting)`

SetPrivacyAdvertisingId gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the PrivacyAdvertisingId field.

### GetPrivacyAutoAcceptPairingAndConsentPrompts

`func (o *Windows10GeneralConfiguration) GetPrivacyAutoAcceptPairingAndConsentPrompts() bool`

GetPrivacyAutoAcceptPairingAndConsentPrompts returns the PrivacyAutoAcceptPairingAndConsentPrompts field if non-nil, zero value otherwise.

### GetPrivacyAutoAcceptPairingAndConsentPromptsOk

`func (o *Windows10GeneralConfiguration) GetPrivacyAutoAcceptPairingAndConsentPromptsOk() (bool, bool)`

GetPrivacyAutoAcceptPairingAndConsentPromptsOk returns a tuple with the PrivacyAutoAcceptPairingAndConsentPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyAutoAcceptPairingAndConsentPrompts

`func (o *Windows10GeneralConfiguration) HasPrivacyAutoAcceptPairingAndConsentPrompts() bool`

HasPrivacyAutoAcceptPairingAndConsentPrompts returns a boolean if a field has been set.

### SetPrivacyAutoAcceptPairingAndConsentPrompts

`func (o *Windows10GeneralConfiguration) SetPrivacyAutoAcceptPairingAndConsentPrompts(v bool)`

SetPrivacyAutoAcceptPairingAndConsentPrompts gets a reference to the given bool and assigns it to the PrivacyAutoAcceptPairingAndConsentPrompts field.

### GetPrivacyBlockInputPersonalization

`func (o *Windows10GeneralConfiguration) GetPrivacyBlockInputPersonalization() bool`

GetPrivacyBlockInputPersonalization returns the PrivacyBlockInputPersonalization field if non-nil, zero value otherwise.

### GetPrivacyBlockInputPersonalizationOk

`func (o *Windows10GeneralConfiguration) GetPrivacyBlockInputPersonalizationOk() (bool, bool)`

GetPrivacyBlockInputPersonalizationOk returns a tuple with the PrivacyBlockInputPersonalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyBlockInputPersonalization

`func (o *Windows10GeneralConfiguration) HasPrivacyBlockInputPersonalization() bool`

HasPrivacyBlockInputPersonalization returns a boolean if a field has been set.

### SetPrivacyBlockInputPersonalization

`func (o *Windows10GeneralConfiguration) SetPrivacyBlockInputPersonalization(v bool)`

SetPrivacyBlockInputPersonalization gets a reference to the given bool and assigns it to the PrivacyBlockInputPersonalization field.

### GetStartBlockUnpinningAppsFromTaskbar

`func (o *Windows10GeneralConfiguration) GetStartBlockUnpinningAppsFromTaskbar() bool`

GetStartBlockUnpinningAppsFromTaskbar returns the StartBlockUnpinningAppsFromTaskbar field if non-nil, zero value otherwise.

### GetStartBlockUnpinningAppsFromTaskbarOk

`func (o *Windows10GeneralConfiguration) GetStartBlockUnpinningAppsFromTaskbarOk() (bool, bool)`

GetStartBlockUnpinningAppsFromTaskbarOk returns a tuple with the StartBlockUnpinningAppsFromTaskbar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartBlockUnpinningAppsFromTaskbar

`func (o *Windows10GeneralConfiguration) HasStartBlockUnpinningAppsFromTaskbar() bool`

HasStartBlockUnpinningAppsFromTaskbar returns a boolean if a field has been set.

### SetStartBlockUnpinningAppsFromTaskbar

`func (o *Windows10GeneralConfiguration) SetStartBlockUnpinningAppsFromTaskbar(v bool)`

SetStartBlockUnpinningAppsFromTaskbar gets a reference to the given bool and assigns it to the StartBlockUnpinningAppsFromTaskbar field.

### GetStartMenuAppListVisibility

`func (o *Windows10GeneralConfiguration) GetStartMenuAppListVisibility() AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType`

GetStartMenuAppListVisibility returns the StartMenuAppListVisibility field if non-nil, zero value otherwise.

### GetStartMenuAppListVisibilityOk

`func (o *Windows10GeneralConfiguration) GetStartMenuAppListVisibilityOk() (AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType, bool)`

GetStartMenuAppListVisibilityOk returns a tuple with the StartMenuAppListVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuAppListVisibility

`func (o *Windows10GeneralConfiguration) HasStartMenuAppListVisibility() bool`

HasStartMenuAppListVisibility returns a boolean if a field has been set.

### SetStartMenuAppListVisibility

`func (o *Windows10GeneralConfiguration) SetStartMenuAppListVisibility(v AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType)`

SetStartMenuAppListVisibility gets a reference to the given AnyOfmicrosoftGraphWindowsStartMenuAppListVisibilityType and assigns it to the StartMenuAppListVisibility field.

### GetStartMenuHideChangeAccountSettings

`func (o *Windows10GeneralConfiguration) GetStartMenuHideChangeAccountSettings() bool`

GetStartMenuHideChangeAccountSettings returns the StartMenuHideChangeAccountSettings field if non-nil, zero value otherwise.

### GetStartMenuHideChangeAccountSettingsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideChangeAccountSettingsOk() (bool, bool)`

GetStartMenuHideChangeAccountSettingsOk returns a tuple with the StartMenuHideChangeAccountSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideChangeAccountSettings

`func (o *Windows10GeneralConfiguration) HasStartMenuHideChangeAccountSettings() bool`

HasStartMenuHideChangeAccountSettings returns a boolean if a field has been set.

### SetStartMenuHideChangeAccountSettings

`func (o *Windows10GeneralConfiguration) SetStartMenuHideChangeAccountSettings(v bool)`

SetStartMenuHideChangeAccountSettings gets a reference to the given bool and assigns it to the StartMenuHideChangeAccountSettings field.

### GetStartMenuHideFrequentlyUsedApps

`func (o *Windows10GeneralConfiguration) GetStartMenuHideFrequentlyUsedApps() bool`

GetStartMenuHideFrequentlyUsedApps returns the StartMenuHideFrequentlyUsedApps field if non-nil, zero value otherwise.

### GetStartMenuHideFrequentlyUsedAppsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideFrequentlyUsedAppsOk() (bool, bool)`

GetStartMenuHideFrequentlyUsedAppsOk returns a tuple with the StartMenuHideFrequentlyUsedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideFrequentlyUsedApps

`func (o *Windows10GeneralConfiguration) HasStartMenuHideFrequentlyUsedApps() bool`

HasStartMenuHideFrequentlyUsedApps returns a boolean if a field has been set.

### SetStartMenuHideFrequentlyUsedApps

`func (o *Windows10GeneralConfiguration) SetStartMenuHideFrequentlyUsedApps(v bool)`

SetStartMenuHideFrequentlyUsedApps gets a reference to the given bool and assigns it to the StartMenuHideFrequentlyUsedApps field.

### GetStartMenuHideHibernate

`func (o *Windows10GeneralConfiguration) GetStartMenuHideHibernate() bool`

GetStartMenuHideHibernate returns the StartMenuHideHibernate field if non-nil, zero value otherwise.

### GetStartMenuHideHibernateOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideHibernateOk() (bool, bool)`

GetStartMenuHideHibernateOk returns a tuple with the StartMenuHideHibernate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideHibernate

`func (o *Windows10GeneralConfiguration) HasStartMenuHideHibernate() bool`

HasStartMenuHideHibernate returns a boolean if a field has been set.

### SetStartMenuHideHibernate

`func (o *Windows10GeneralConfiguration) SetStartMenuHideHibernate(v bool)`

SetStartMenuHideHibernate gets a reference to the given bool and assigns it to the StartMenuHideHibernate field.

### GetStartMenuHideLock

`func (o *Windows10GeneralConfiguration) GetStartMenuHideLock() bool`

GetStartMenuHideLock returns the StartMenuHideLock field if non-nil, zero value otherwise.

### GetStartMenuHideLockOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideLockOk() (bool, bool)`

GetStartMenuHideLockOk returns a tuple with the StartMenuHideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideLock

`func (o *Windows10GeneralConfiguration) HasStartMenuHideLock() bool`

HasStartMenuHideLock returns a boolean if a field has been set.

### SetStartMenuHideLock

`func (o *Windows10GeneralConfiguration) SetStartMenuHideLock(v bool)`

SetStartMenuHideLock gets a reference to the given bool and assigns it to the StartMenuHideLock field.

### GetStartMenuHidePowerButton

`func (o *Windows10GeneralConfiguration) GetStartMenuHidePowerButton() bool`

GetStartMenuHidePowerButton returns the StartMenuHidePowerButton field if non-nil, zero value otherwise.

### GetStartMenuHidePowerButtonOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHidePowerButtonOk() (bool, bool)`

GetStartMenuHidePowerButtonOk returns a tuple with the StartMenuHidePowerButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHidePowerButton

`func (o *Windows10GeneralConfiguration) HasStartMenuHidePowerButton() bool`

HasStartMenuHidePowerButton returns a boolean if a field has been set.

### SetStartMenuHidePowerButton

`func (o *Windows10GeneralConfiguration) SetStartMenuHidePowerButton(v bool)`

SetStartMenuHidePowerButton gets a reference to the given bool and assigns it to the StartMenuHidePowerButton field.

### GetStartMenuHideRecentJumpLists

`func (o *Windows10GeneralConfiguration) GetStartMenuHideRecentJumpLists() bool`

GetStartMenuHideRecentJumpLists returns the StartMenuHideRecentJumpLists field if non-nil, zero value otherwise.

### GetStartMenuHideRecentJumpListsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideRecentJumpListsOk() (bool, bool)`

GetStartMenuHideRecentJumpListsOk returns a tuple with the StartMenuHideRecentJumpLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideRecentJumpLists

`func (o *Windows10GeneralConfiguration) HasStartMenuHideRecentJumpLists() bool`

HasStartMenuHideRecentJumpLists returns a boolean if a field has been set.

### SetStartMenuHideRecentJumpLists

`func (o *Windows10GeneralConfiguration) SetStartMenuHideRecentJumpLists(v bool)`

SetStartMenuHideRecentJumpLists gets a reference to the given bool and assigns it to the StartMenuHideRecentJumpLists field.

### GetStartMenuHideRecentlyAddedApps

`func (o *Windows10GeneralConfiguration) GetStartMenuHideRecentlyAddedApps() bool`

GetStartMenuHideRecentlyAddedApps returns the StartMenuHideRecentlyAddedApps field if non-nil, zero value otherwise.

### GetStartMenuHideRecentlyAddedAppsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideRecentlyAddedAppsOk() (bool, bool)`

GetStartMenuHideRecentlyAddedAppsOk returns a tuple with the StartMenuHideRecentlyAddedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideRecentlyAddedApps

`func (o *Windows10GeneralConfiguration) HasStartMenuHideRecentlyAddedApps() bool`

HasStartMenuHideRecentlyAddedApps returns a boolean if a field has been set.

### SetStartMenuHideRecentlyAddedApps

`func (o *Windows10GeneralConfiguration) SetStartMenuHideRecentlyAddedApps(v bool)`

SetStartMenuHideRecentlyAddedApps gets a reference to the given bool and assigns it to the StartMenuHideRecentlyAddedApps field.

### GetStartMenuHideRestartOptions

`func (o *Windows10GeneralConfiguration) GetStartMenuHideRestartOptions() bool`

GetStartMenuHideRestartOptions returns the StartMenuHideRestartOptions field if non-nil, zero value otherwise.

### GetStartMenuHideRestartOptionsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideRestartOptionsOk() (bool, bool)`

GetStartMenuHideRestartOptionsOk returns a tuple with the StartMenuHideRestartOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideRestartOptions

`func (o *Windows10GeneralConfiguration) HasStartMenuHideRestartOptions() bool`

HasStartMenuHideRestartOptions returns a boolean if a field has been set.

### SetStartMenuHideRestartOptions

`func (o *Windows10GeneralConfiguration) SetStartMenuHideRestartOptions(v bool)`

SetStartMenuHideRestartOptions gets a reference to the given bool and assigns it to the StartMenuHideRestartOptions field.

### GetStartMenuHideShutDown

`func (o *Windows10GeneralConfiguration) GetStartMenuHideShutDown() bool`

GetStartMenuHideShutDown returns the StartMenuHideShutDown field if non-nil, zero value otherwise.

### GetStartMenuHideShutDownOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideShutDownOk() (bool, bool)`

GetStartMenuHideShutDownOk returns a tuple with the StartMenuHideShutDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideShutDown

`func (o *Windows10GeneralConfiguration) HasStartMenuHideShutDown() bool`

HasStartMenuHideShutDown returns a boolean if a field has been set.

### SetStartMenuHideShutDown

`func (o *Windows10GeneralConfiguration) SetStartMenuHideShutDown(v bool)`

SetStartMenuHideShutDown gets a reference to the given bool and assigns it to the StartMenuHideShutDown field.

### GetStartMenuHideSignOut

`func (o *Windows10GeneralConfiguration) GetStartMenuHideSignOut() bool`

GetStartMenuHideSignOut returns the StartMenuHideSignOut field if non-nil, zero value otherwise.

### GetStartMenuHideSignOutOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideSignOutOk() (bool, bool)`

GetStartMenuHideSignOutOk returns a tuple with the StartMenuHideSignOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideSignOut

`func (o *Windows10GeneralConfiguration) HasStartMenuHideSignOut() bool`

HasStartMenuHideSignOut returns a boolean if a field has been set.

### SetStartMenuHideSignOut

`func (o *Windows10GeneralConfiguration) SetStartMenuHideSignOut(v bool)`

SetStartMenuHideSignOut gets a reference to the given bool and assigns it to the StartMenuHideSignOut field.

### GetStartMenuHideSleep

`func (o *Windows10GeneralConfiguration) GetStartMenuHideSleep() bool`

GetStartMenuHideSleep returns the StartMenuHideSleep field if non-nil, zero value otherwise.

### GetStartMenuHideSleepOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideSleepOk() (bool, bool)`

GetStartMenuHideSleepOk returns a tuple with the StartMenuHideSleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideSleep

`func (o *Windows10GeneralConfiguration) HasStartMenuHideSleep() bool`

HasStartMenuHideSleep returns a boolean if a field has been set.

### SetStartMenuHideSleep

`func (o *Windows10GeneralConfiguration) SetStartMenuHideSleep(v bool)`

SetStartMenuHideSleep gets a reference to the given bool and assigns it to the StartMenuHideSleep field.

### GetStartMenuHideSwitchAccount

`func (o *Windows10GeneralConfiguration) GetStartMenuHideSwitchAccount() bool`

GetStartMenuHideSwitchAccount returns the StartMenuHideSwitchAccount field if non-nil, zero value otherwise.

### GetStartMenuHideSwitchAccountOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideSwitchAccountOk() (bool, bool)`

GetStartMenuHideSwitchAccountOk returns a tuple with the StartMenuHideSwitchAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideSwitchAccount

`func (o *Windows10GeneralConfiguration) HasStartMenuHideSwitchAccount() bool`

HasStartMenuHideSwitchAccount returns a boolean if a field has been set.

### SetStartMenuHideSwitchAccount

`func (o *Windows10GeneralConfiguration) SetStartMenuHideSwitchAccount(v bool)`

SetStartMenuHideSwitchAccount gets a reference to the given bool and assigns it to the StartMenuHideSwitchAccount field.

### GetStartMenuHideUserTile

`func (o *Windows10GeneralConfiguration) GetStartMenuHideUserTile() bool`

GetStartMenuHideUserTile returns the StartMenuHideUserTile field if non-nil, zero value otherwise.

### GetStartMenuHideUserTileOk

`func (o *Windows10GeneralConfiguration) GetStartMenuHideUserTileOk() (bool, bool)`

GetStartMenuHideUserTileOk returns a tuple with the StartMenuHideUserTile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuHideUserTile

`func (o *Windows10GeneralConfiguration) HasStartMenuHideUserTile() bool`

HasStartMenuHideUserTile returns a boolean if a field has been set.

### SetStartMenuHideUserTile

`func (o *Windows10GeneralConfiguration) SetStartMenuHideUserTile(v bool)`

SetStartMenuHideUserTile gets a reference to the given bool and assigns it to the StartMenuHideUserTile field.

### GetStartMenuLayoutEdgeAssetsXml

`func (o *Windows10GeneralConfiguration) GetStartMenuLayoutEdgeAssetsXml() string`

GetStartMenuLayoutEdgeAssetsXml returns the StartMenuLayoutEdgeAssetsXml field if non-nil, zero value otherwise.

### GetStartMenuLayoutEdgeAssetsXmlOk

`func (o *Windows10GeneralConfiguration) GetStartMenuLayoutEdgeAssetsXmlOk() (string, bool)`

GetStartMenuLayoutEdgeAssetsXmlOk returns a tuple with the StartMenuLayoutEdgeAssetsXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuLayoutEdgeAssetsXml

`func (o *Windows10GeneralConfiguration) HasStartMenuLayoutEdgeAssetsXml() bool`

HasStartMenuLayoutEdgeAssetsXml returns a boolean if a field has been set.

### SetStartMenuLayoutEdgeAssetsXml

`func (o *Windows10GeneralConfiguration) SetStartMenuLayoutEdgeAssetsXml(v string)`

SetStartMenuLayoutEdgeAssetsXml gets a reference to the given string and assigns it to the StartMenuLayoutEdgeAssetsXml field.

### SetStartMenuLayoutEdgeAssetsXmlExplicitNull

`func (o *Windows10GeneralConfiguration) SetStartMenuLayoutEdgeAssetsXmlExplicitNull(b bool)`

SetStartMenuLayoutEdgeAssetsXmlExplicitNull (un)sets StartMenuLayoutEdgeAssetsXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartMenuLayoutEdgeAssetsXml value is set to nil even if false is passed
### GetStartMenuLayoutXml

`func (o *Windows10GeneralConfiguration) GetStartMenuLayoutXml() string`

GetStartMenuLayoutXml returns the StartMenuLayoutXml field if non-nil, zero value otherwise.

### GetStartMenuLayoutXmlOk

`func (o *Windows10GeneralConfiguration) GetStartMenuLayoutXmlOk() (string, bool)`

GetStartMenuLayoutXmlOk returns a tuple with the StartMenuLayoutXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuLayoutXml

`func (o *Windows10GeneralConfiguration) HasStartMenuLayoutXml() bool`

HasStartMenuLayoutXml returns a boolean if a field has been set.

### SetStartMenuLayoutXml

`func (o *Windows10GeneralConfiguration) SetStartMenuLayoutXml(v string)`

SetStartMenuLayoutXml gets a reference to the given string and assigns it to the StartMenuLayoutXml field.

### SetStartMenuLayoutXmlExplicitNull

`func (o *Windows10GeneralConfiguration) SetStartMenuLayoutXmlExplicitNull(b bool)`

SetStartMenuLayoutXmlExplicitNull (un)sets StartMenuLayoutXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartMenuLayoutXml value is set to nil even if false is passed
### GetStartMenuMode

`func (o *Windows10GeneralConfiguration) GetStartMenuMode() AnyOfmicrosoftGraphWindowsStartMenuModeType`

GetStartMenuMode returns the StartMenuMode field if non-nil, zero value otherwise.

### GetStartMenuModeOk

`func (o *Windows10GeneralConfiguration) GetStartMenuModeOk() (AnyOfmicrosoftGraphWindowsStartMenuModeType, bool)`

GetStartMenuModeOk returns a tuple with the StartMenuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuMode

`func (o *Windows10GeneralConfiguration) HasStartMenuMode() bool`

HasStartMenuMode returns a boolean if a field has been set.

### SetStartMenuMode

`func (o *Windows10GeneralConfiguration) SetStartMenuMode(v AnyOfmicrosoftGraphWindowsStartMenuModeType)`

SetStartMenuMode gets a reference to the given AnyOfmicrosoftGraphWindowsStartMenuModeType and assigns it to the StartMenuMode field.

### GetStartMenuPinnedFolderDocuments

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderDocuments() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderDocuments returns the StartMenuPinnedFolderDocuments field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderDocumentsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderDocumentsOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderDocumentsOk returns a tuple with the StartMenuPinnedFolderDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderDocuments

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderDocuments() bool`

HasStartMenuPinnedFolderDocuments returns a boolean if a field has been set.

### SetStartMenuPinnedFolderDocuments

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderDocuments(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderDocuments gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderDocuments field.

### GetStartMenuPinnedFolderDownloads

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderDownloads() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderDownloads returns the StartMenuPinnedFolderDownloads field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderDownloadsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderDownloadsOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderDownloadsOk returns a tuple with the StartMenuPinnedFolderDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderDownloads

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderDownloads() bool`

HasStartMenuPinnedFolderDownloads returns a boolean if a field has been set.

### SetStartMenuPinnedFolderDownloads

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderDownloads(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderDownloads gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderDownloads field.

### GetStartMenuPinnedFolderFileExplorer

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderFileExplorer() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderFileExplorer returns the StartMenuPinnedFolderFileExplorer field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderFileExplorerOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderFileExplorerOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderFileExplorerOk returns a tuple with the StartMenuPinnedFolderFileExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderFileExplorer

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderFileExplorer() bool`

HasStartMenuPinnedFolderFileExplorer returns a boolean if a field has been set.

### SetStartMenuPinnedFolderFileExplorer

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderFileExplorer(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderFileExplorer gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderFileExplorer field.

### GetStartMenuPinnedFolderHomeGroup

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderHomeGroup() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderHomeGroup returns the StartMenuPinnedFolderHomeGroup field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderHomeGroupOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderHomeGroupOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderHomeGroupOk returns a tuple with the StartMenuPinnedFolderHomeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderHomeGroup

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderHomeGroup() bool`

HasStartMenuPinnedFolderHomeGroup returns a boolean if a field has been set.

### SetStartMenuPinnedFolderHomeGroup

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderHomeGroup(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderHomeGroup gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderHomeGroup field.

### GetStartMenuPinnedFolderMusic

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderMusic() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderMusic returns the StartMenuPinnedFolderMusic field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderMusicOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderMusicOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderMusicOk returns a tuple with the StartMenuPinnedFolderMusic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderMusic

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderMusic() bool`

HasStartMenuPinnedFolderMusic returns a boolean if a field has been set.

### SetStartMenuPinnedFolderMusic

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderMusic(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderMusic gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderMusic field.

### GetStartMenuPinnedFolderNetwork

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderNetwork() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderNetwork returns the StartMenuPinnedFolderNetwork field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderNetworkOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderNetworkOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderNetworkOk returns a tuple with the StartMenuPinnedFolderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderNetwork

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderNetwork() bool`

HasStartMenuPinnedFolderNetwork returns a boolean if a field has been set.

### SetStartMenuPinnedFolderNetwork

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderNetwork(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderNetwork gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderNetwork field.

### GetStartMenuPinnedFolderPersonalFolder

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderPersonalFolder() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderPersonalFolder returns the StartMenuPinnedFolderPersonalFolder field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderPersonalFolderOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderPersonalFolderOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderPersonalFolderOk returns a tuple with the StartMenuPinnedFolderPersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderPersonalFolder

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderPersonalFolder() bool`

HasStartMenuPinnedFolderPersonalFolder returns a boolean if a field has been set.

### SetStartMenuPinnedFolderPersonalFolder

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderPersonalFolder(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderPersonalFolder gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderPersonalFolder field.

### GetStartMenuPinnedFolderPictures

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderPictures() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderPictures returns the StartMenuPinnedFolderPictures field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderPicturesOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderPicturesOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderPicturesOk returns a tuple with the StartMenuPinnedFolderPictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderPictures

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderPictures() bool`

HasStartMenuPinnedFolderPictures returns a boolean if a field has been set.

### SetStartMenuPinnedFolderPictures

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderPictures(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderPictures gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderPictures field.

### GetStartMenuPinnedFolderSettings

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderSettings() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderSettings returns the StartMenuPinnedFolderSettings field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderSettingsOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderSettingsOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderSettingsOk returns a tuple with the StartMenuPinnedFolderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderSettings

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderSettings() bool`

HasStartMenuPinnedFolderSettings returns a boolean if a field has been set.

### SetStartMenuPinnedFolderSettings

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderSettings(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderSettings gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderSettings field.

### GetStartMenuPinnedFolderVideos

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderVideos() AnyOfmicrosoftGraphVisibilitySetting`

GetStartMenuPinnedFolderVideos returns the StartMenuPinnedFolderVideos field if non-nil, zero value otherwise.

### GetStartMenuPinnedFolderVideosOk

`func (o *Windows10GeneralConfiguration) GetStartMenuPinnedFolderVideosOk() (AnyOfmicrosoftGraphVisibilitySetting, bool)`

GetStartMenuPinnedFolderVideosOk returns a tuple with the StartMenuPinnedFolderVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartMenuPinnedFolderVideos

`func (o *Windows10GeneralConfiguration) HasStartMenuPinnedFolderVideos() bool`

HasStartMenuPinnedFolderVideos returns a boolean if a field has been set.

### SetStartMenuPinnedFolderVideos

`func (o *Windows10GeneralConfiguration) SetStartMenuPinnedFolderVideos(v AnyOfmicrosoftGraphVisibilitySetting)`

SetStartMenuPinnedFolderVideos gets a reference to the given AnyOfmicrosoftGraphVisibilitySetting and assigns it to the StartMenuPinnedFolderVideos field.

### GetSettingsBlockSettingsApp

`func (o *Windows10GeneralConfiguration) GetSettingsBlockSettingsApp() bool`

GetSettingsBlockSettingsApp returns the SettingsBlockSettingsApp field if non-nil, zero value otherwise.

### GetSettingsBlockSettingsAppOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockSettingsAppOk() (bool, bool)`

GetSettingsBlockSettingsAppOk returns a tuple with the SettingsBlockSettingsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSettingsApp

`func (o *Windows10GeneralConfiguration) HasSettingsBlockSettingsApp() bool`

HasSettingsBlockSettingsApp returns a boolean if a field has been set.

### SetSettingsBlockSettingsApp

`func (o *Windows10GeneralConfiguration) SetSettingsBlockSettingsApp(v bool)`

SetSettingsBlockSettingsApp gets a reference to the given bool and assigns it to the SettingsBlockSettingsApp field.

### GetSettingsBlockSystemPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockSystemPage() bool`

GetSettingsBlockSystemPage returns the SettingsBlockSystemPage field if non-nil, zero value otherwise.

### GetSettingsBlockSystemPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockSystemPageOk() (bool, bool)`

GetSettingsBlockSystemPageOk returns a tuple with the SettingsBlockSystemPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockSystemPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockSystemPage() bool`

HasSettingsBlockSystemPage returns a boolean if a field has been set.

### SetSettingsBlockSystemPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockSystemPage(v bool)`

SetSettingsBlockSystemPage gets a reference to the given bool and assigns it to the SettingsBlockSystemPage field.

### GetSettingsBlockDevicesPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockDevicesPage() bool`

GetSettingsBlockDevicesPage returns the SettingsBlockDevicesPage field if non-nil, zero value otherwise.

### GetSettingsBlockDevicesPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockDevicesPageOk() (bool, bool)`

GetSettingsBlockDevicesPageOk returns a tuple with the SettingsBlockDevicesPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockDevicesPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockDevicesPage() bool`

HasSettingsBlockDevicesPage returns a boolean if a field has been set.

### SetSettingsBlockDevicesPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockDevicesPage(v bool)`

SetSettingsBlockDevicesPage gets a reference to the given bool and assigns it to the SettingsBlockDevicesPage field.

### GetSettingsBlockNetworkInternetPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockNetworkInternetPage() bool`

GetSettingsBlockNetworkInternetPage returns the SettingsBlockNetworkInternetPage field if non-nil, zero value otherwise.

### GetSettingsBlockNetworkInternetPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockNetworkInternetPageOk() (bool, bool)`

GetSettingsBlockNetworkInternetPageOk returns a tuple with the SettingsBlockNetworkInternetPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockNetworkInternetPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockNetworkInternetPage() bool`

HasSettingsBlockNetworkInternetPage returns a boolean if a field has been set.

### SetSettingsBlockNetworkInternetPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockNetworkInternetPage(v bool)`

SetSettingsBlockNetworkInternetPage gets a reference to the given bool and assigns it to the SettingsBlockNetworkInternetPage field.

### GetSettingsBlockPersonalizationPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockPersonalizationPage() bool`

GetSettingsBlockPersonalizationPage returns the SettingsBlockPersonalizationPage field if non-nil, zero value otherwise.

### GetSettingsBlockPersonalizationPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockPersonalizationPageOk() (bool, bool)`

GetSettingsBlockPersonalizationPageOk returns a tuple with the SettingsBlockPersonalizationPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockPersonalizationPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockPersonalizationPage() bool`

HasSettingsBlockPersonalizationPage returns a boolean if a field has been set.

### SetSettingsBlockPersonalizationPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockPersonalizationPage(v bool)`

SetSettingsBlockPersonalizationPage gets a reference to the given bool and assigns it to the SettingsBlockPersonalizationPage field.

### GetSettingsBlockAccountsPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockAccountsPage() bool`

GetSettingsBlockAccountsPage returns the SettingsBlockAccountsPage field if non-nil, zero value otherwise.

### GetSettingsBlockAccountsPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockAccountsPageOk() (bool, bool)`

GetSettingsBlockAccountsPageOk returns a tuple with the SettingsBlockAccountsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockAccountsPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockAccountsPage() bool`

HasSettingsBlockAccountsPage returns a boolean if a field has been set.

### SetSettingsBlockAccountsPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockAccountsPage(v bool)`

SetSettingsBlockAccountsPage gets a reference to the given bool and assigns it to the SettingsBlockAccountsPage field.

### GetSettingsBlockTimeLanguagePage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockTimeLanguagePage() bool`

GetSettingsBlockTimeLanguagePage returns the SettingsBlockTimeLanguagePage field if non-nil, zero value otherwise.

### GetSettingsBlockTimeLanguagePageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockTimeLanguagePageOk() (bool, bool)`

GetSettingsBlockTimeLanguagePageOk returns a tuple with the SettingsBlockTimeLanguagePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockTimeLanguagePage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockTimeLanguagePage() bool`

HasSettingsBlockTimeLanguagePage returns a boolean if a field has been set.

### SetSettingsBlockTimeLanguagePage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockTimeLanguagePage(v bool)`

SetSettingsBlockTimeLanguagePage gets a reference to the given bool and assigns it to the SettingsBlockTimeLanguagePage field.

### GetSettingsBlockEaseOfAccessPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockEaseOfAccessPage() bool`

GetSettingsBlockEaseOfAccessPage returns the SettingsBlockEaseOfAccessPage field if non-nil, zero value otherwise.

### GetSettingsBlockEaseOfAccessPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockEaseOfAccessPageOk() (bool, bool)`

GetSettingsBlockEaseOfAccessPageOk returns a tuple with the SettingsBlockEaseOfAccessPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockEaseOfAccessPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockEaseOfAccessPage() bool`

HasSettingsBlockEaseOfAccessPage returns a boolean if a field has been set.

### SetSettingsBlockEaseOfAccessPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockEaseOfAccessPage(v bool)`

SetSettingsBlockEaseOfAccessPage gets a reference to the given bool and assigns it to the SettingsBlockEaseOfAccessPage field.

### GetSettingsBlockPrivacyPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockPrivacyPage() bool`

GetSettingsBlockPrivacyPage returns the SettingsBlockPrivacyPage field if non-nil, zero value otherwise.

### GetSettingsBlockPrivacyPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockPrivacyPageOk() (bool, bool)`

GetSettingsBlockPrivacyPageOk returns a tuple with the SettingsBlockPrivacyPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockPrivacyPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockPrivacyPage() bool`

HasSettingsBlockPrivacyPage returns a boolean if a field has been set.

### SetSettingsBlockPrivacyPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockPrivacyPage(v bool)`

SetSettingsBlockPrivacyPage gets a reference to the given bool and assigns it to the SettingsBlockPrivacyPage field.

### GetSettingsBlockUpdateSecurityPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockUpdateSecurityPage() bool`

GetSettingsBlockUpdateSecurityPage returns the SettingsBlockUpdateSecurityPage field if non-nil, zero value otherwise.

### GetSettingsBlockUpdateSecurityPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockUpdateSecurityPageOk() (bool, bool)`

GetSettingsBlockUpdateSecurityPageOk returns a tuple with the SettingsBlockUpdateSecurityPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockUpdateSecurityPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockUpdateSecurityPage() bool`

HasSettingsBlockUpdateSecurityPage returns a boolean if a field has been set.

### SetSettingsBlockUpdateSecurityPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockUpdateSecurityPage(v bool)`

SetSettingsBlockUpdateSecurityPage gets a reference to the given bool and assigns it to the SettingsBlockUpdateSecurityPage field.

### GetSettingsBlockAppsPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockAppsPage() bool`

GetSettingsBlockAppsPage returns the SettingsBlockAppsPage field if non-nil, zero value otherwise.

### GetSettingsBlockAppsPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockAppsPageOk() (bool, bool)`

GetSettingsBlockAppsPageOk returns a tuple with the SettingsBlockAppsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockAppsPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockAppsPage() bool`

HasSettingsBlockAppsPage returns a boolean if a field has been set.

### SetSettingsBlockAppsPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockAppsPage(v bool)`

SetSettingsBlockAppsPage gets a reference to the given bool and assigns it to the SettingsBlockAppsPage field.

### GetSettingsBlockGamingPage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockGamingPage() bool`

GetSettingsBlockGamingPage returns the SettingsBlockGamingPage field if non-nil, zero value otherwise.

### GetSettingsBlockGamingPageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockGamingPageOk() (bool, bool)`

GetSettingsBlockGamingPageOk returns a tuple with the SettingsBlockGamingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockGamingPage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockGamingPage() bool`

HasSettingsBlockGamingPage returns a boolean if a field has been set.

### SetSettingsBlockGamingPage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockGamingPage(v bool)`

SetSettingsBlockGamingPage gets a reference to the given bool and assigns it to the SettingsBlockGamingPage field.

### GetWindowsSpotlightBlockConsumerSpecificFeatures

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockConsumerSpecificFeatures() bool`

GetWindowsSpotlightBlockConsumerSpecificFeatures returns the WindowsSpotlightBlockConsumerSpecificFeatures field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockConsumerSpecificFeaturesOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockConsumerSpecificFeaturesOk() (bool, bool)`

GetWindowsSpotlightBlockConsumerSpecificFeaturesOk returns a tuple with the WindowsSpotlightBlockConsumerSpecificFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockConsumerSpecificFeatures

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlockConsumerSpecificFeatures() bool`

HasWindowsSpotlightBlockConsumerSpecificFeatures returns a boolean if a field has been set.

### SetWindowsSpotlightBlockConsumerSpecificFeatures

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlockConsumerSpecificFeatures(v bool)`

SetWindowsSpotlightBlockConsumerSpecificFeatures gets a reference to the given bool and assigns it to the WindowsSpotlightBlockConsumerSpecificFeatures field.

### GetWindowsSpotlightBlocked

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlocked() bool`

GetWindowsSpotlightBlocked returns the WindowsSpotlightBlocked field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockedOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockedOk() (bool, bool)`

GetWindowsSpotlightBlockedOk returns a tuple with the WindowsSpotlightBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlocked

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlocked() bool`

HasWindowsSpotlightBlocked returns a boolean if a field has been set.

### SetWindowsSpotlightBlocked

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlocked(v bool)`

SetWindowsSpotlightBlocked gets a reference to the given bool and assigns it to the WindowsSpotlightBlocked field.

### GetWindowsSpotlightBlockOnActionCenter

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockOnActionCenter() bool`

GetWindowsSpotlightBlockOnActionCenter returns the WindowsSpotlightBlockOnActionCenter field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockOnActionCenterOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockOnActionCenterOk() (bool, bool)`

GetWindowsSpotlightBlockOnActionCenterOk returns a tuple with the WindowsSpotlightBlockOnActionCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockOnActionCenter

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlockOnActionCenter() bool`

HasWindowsSpotlightBlockOnActionCenter returns a boolean if a field has been set.

### SetWindowsSpotlightBlockOnActionCenter

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlockOnActionCenter(v bool)`

SetWindowsSpotlightBlockOnActionCenter gets a reference to the given bool and assigns it to the WindowsSpotlightBlockOnActionCenter field.

### GetWindowsSpotlightBlockTailoredExperiences

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockTailoredExperiences() bool`

GetWindowsSpotlightBlockTailoredExperiences returns the WindowsSpotlightBlockTailoredExperiences field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockTailoredExperiencesOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockTailoredExperiencesOk() (bool, bool)`

GetWindowsSpotlightBlockTailoredExperiencesOk returns a tuple with the WindowsSpotlightBlockTailoredExperiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockTailoredExperiences

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlockTailoredExperiences() bool`

HasWindowsSpotlightBlockTailoredExperiences returns a boolean if a field has been set.

### SetWindowsSpotlightBlockTailoredExperiences

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlockTailoredExperiences(v bool)`

SetWindowsSpotlightBlockTailoredExperiences gets a reference to the given bool and assigns it to the WindowsSpotlightBlockTailoredExperiences field.

### GetWindowsSpotlightBlockThirdPartyNotifications

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockThirdPartyNotifications() bool`

GetWindowsSpotlightBlockThirdPartyNotifications returns the WindowsSpotlightBlockThirdPartyNotifications field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockThirdPartyNotificationsOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockThirdPartyNotificationsOk() (bool, bool)`

GetWindowsSpotlightBlockThirdPartyNotificationsOk returns a tuple with the WindowsSpotlightBlockThirdPartyNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockThirdPartyNotifications

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlockThirdPartyNotifications() bool`

HasWindowsSpotlightBlockThirdPartyNotifications returns a boolean if a field has been set.

### SetWindowsSpotlightBlockThirdPartyNotifications

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlockThirdPartyNotifications(v bool)`

SetWindowsSpotlightBlockThirdPartyNotifications gets a reference to the given bool and assigns it to the WindowsSpotlightBlockThirdPartyNotifications field.

### GetWindowsSpotlightBlockWelcomeExperience

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockWelcomeExperience() bool`

GetWindowsSpotlightBlockWelcomeExperience returns the WindowsSpotlightBlockWelcomeExperience field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockWelcomeExperienceOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockWelcomeExperienceOk() (bool, bool)`

GetWindowsSpotlightBlockWelcomeExperienceOk returns a tuple with the WindowsSpotlightBlockWelcomeExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockWelcomeExperience

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlockWelcomeExperience() bool`

HasWindowsSpotlightBlockWelcomeExperience returns a boolean if a field has been set.

### SetWindowsSpotlightBlockWelcomeExperience

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlockWelcomeExperience(v bool)`

SetWindowsSpotlightBlockWelcomeExperience gets a reference to the given bool and assigns it to the WindowsSpotlightBlockWelcomeExperience field.

### GetWindowsSpotlightBlockWindowsTips

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockWindowsTips() bool`

GetWindowsSpotlightBlockWindowsTips returns the WindowsSpotlightBlockWindowsTips field if non-nil, zero value otherwise.

### GetWindowsSpotlightBlockWindowsTipsOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightBlockWindowsTipsOk() (bool, bool)`

GetWindowsSpotlightBlockWindowsTipsOk returns a tuple with the WindowsSpotlightBlockWindowsTips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightBlockWindowsTips

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightBlockWindowsTips() bool`

HasWindowsSpotlightBlockWindowsTips returns a boolean if a field has been set.

### SetWindowsSpotlightBlockWindowsTips

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightBlockWindowsTips(v bool)`

SetWindowsSpotlightBlockWindowsTips gets a reference to the given bool and assigns it to the WindowsSpotlightBlockWindowsTips field.

### GetWindowsSpotlightConfigureOnLockScreen

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightConfigureOnLockScreen() AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings`

GetWindowsSpotlightConfigureOnLockScreen returns the WindowsSpotlightConfigureOnLockScreen field if non-nil, zero value otherwise.

### GetWindowsSpotlightConfigureOnLockScreenOk

`func (o *Windows10GeneralConfiguration) GetWindowsSpotlightConfigureOnLockScreenOk() (AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings, bool)`

GetWindowsSpotlightConfigureOnLockScreenOk returns a tuple with the WindowsSpotlightConfigureOnLockScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsSpotlightConfigureOnLockScreen

`func (o *Windows10GeneralConfiguration) HasWindowsSpotlightConfigureOnLockScreen() bool`

HasWindowsSpotlightConfigureOnLockScreen returns a boolean if a field has been set.

### SetWindowsSpotlightConfigureOnLockScreen

`func (o *Windows10GeneralConfiguration) SetWindowsSpotlightConfigureOnLockScreen(v AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings)`

SetWindowsSpotlightConfigureOnLockScreen gets a reference to the given AnyOfmicrosoftGraphWindowsSpotlightEnablementSettings and assigns it to the WindowsSpotlightConfigureOnLockScreen field.

### GetNetworkProxyApplySettingsDeviceWide

`func (o *Windows10GeneralConfiguration) GetNetworkProxyApplySettingsDeviceWide() bool`

GetNetworkProxyApplySettingsDeviceWide returns the NetworkProxyApplySettingsDeviceWide field if non-nil, zero value otherwise.

### GetNetworkProxyApplySettingsDeviceWideOk

`func (o *Windows10GeneralConfiguration) GetNetworkProxyApplySettingsDeviceWideOk() (bool, bool)`

GetNetworkProxyApplySettingsDeviceWideOk returns a tuple with the NetworkProxyApplySettingsDeviceWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyApplySettingsDeviceWide

`func (o *Windows10GeneralConfiguration) HasNetworkProxyApplySettingsDeviceWide() bool`

HasNetworkProxyApplySettingsDeviceWide returns a boolean if a field has been set.

### SetNetworkProxyApplySettingsDeviceWide

`func (o *Windows10GeneralConfiguration) SetNetworkProxyApplySettingsDeviceWide(v bool)`

SetNetworkProxyApplySettingsDeviceWide gets a reference to the given bool and assigns it to the NetworkProxyApplySettingsDeviceWide field.

### GetNetworkProxyDisableAutoDetect

`func (o *Windows10GeneralConfiguration) GetNetworkProxyDisableAutoDetect() bool`

GetNetworkProxyDisableAutoDetect returns the NetworkProxyDisableAutoDetect field if non-nil, zero value otherwise.

### GetNetworkProxyDisableAutoDetectOk

`func (o *Windows10GeneralConfiguration) GetNetworkProxyDisableAutoDetectOk() (bool, bool)`

GetNetworkProxyDisableAutoDetectOk returns a tuple with the NetworkProxyDisableAutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyDisableAutoDetect

`func (o *Windows10GeneralConfiguration) HasNetworkProxyDisableAutoDetect() bool`

HasNetworkProxyDisableAutoDetect returns a boolean if a field has been set.

### SetNetworkProxyDisableAutoDetect

`func (o *Windows10GeneralConfiguration) SetNetworkProxyDisableAutoDetect(v bool)`

SetNetworkProxyDisableAutoDetect gets a reference to the given bool and assigns it to the NetworkProxyDisableAutoDetect field.

### GetNetworkProxyAutomaticConfigurationUrl

`func (o *Windows10GeneralConfiguration) GetNetworkProxyAutomaticConfigurationUrl() string`

GetNetworkProxyAutomaticConfigurationUrl returns the NetworkProxyAutomaticConfigurationUrl field if non-nil, zero value otherwise.

### GetNetworkProxyAutomaticConfigurationUrlOk

`func (o *Windows10GeneralConfiguration) GetNetworkProxyAutomaticConfigurationUrlOk() (string, bool)`

GetNetworkProxyAutomaticConfigurationUrlOk returns a tuple with the NetworkProxyAutomaticConfigurationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyAutomaticConfigurationUrl

`func (o *Windows10GeneralConfiguration) HasNetworkProxyAutomaticConfigurationUrl() bool`

HasNetworkProxyAutomaticConfigurationUrl returns a boolean if a field has been set.

### SetNetworkProxyAutomaticConfigurationUrl

`func (o *Windows10GeneralConfiguration) SetNetworkProxyAutomaticConfigurationUrl(v string)`

SetNetworkProxyAutomaticConfigurationUrl gets a reference to the given string and assigns it to the NetworkProxyAutomaticConfigurationUrl field.

### SetNetworkProxyAutomaticConfigurationUrlExplicitNull

`func (o *Windows10GeneralConfiguration) SetNetworkProxyAutomaticConfigurationUrlExplicitNull(b bool)`

SetNetworkProxyAutomaticConfigurationUrlExplicitNull (un)sets NetworkProxyAutomaticConfigurationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NetworkProxyAutomaticConfigurationUrl value is set to nil even if false is passed
### GetNetworkProxyServer

`func (o *Windows10GeneralConfiguration) GetNetworkProxyServer() AnyOfmicrosoftGraphWindows10NetworkProxyServer`

GetNetworkProxyServer returns the NetworkProxyServer field if non-nil, zero value otherwise.

### GetNetworkProxyServerOk

`func (o *Windows10GeneralConfiguration) GetNetworkProxyServerOk() (AnyOfmicrosoftGraphWindows10NetworkProxyServer, bool)`

GetNetworkProxyServerOk returns a tuple with the NetworkProxyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkProxyServer

`func (o *Windows10GeneralConfiguration) HasNetworkProxyServer() bool`

HasNetworkProxyServer returns a boolean if a field has been set.

### SetNetworkProxyServer

`func (o *Windows10GeneralConfiguration) SetNetworkProxyServer(v AnyOfmicrosoftGraphWindows10NetworkProxyServer)`

SetNetworkProxyServer gets a reference to the given AnyOfmicrosoftGraphWindows10NetworkProxyServer and assigns it to the NetworkProxyServer field.

### SetNetworkProxyServerExplicitNull

`func (o *Windows10GeneralConfiguration) SetNetworkProxyServerExplicitNull(b bool)`

SetNetworkProxyServerExplicitNull (un)sets NetworkProxyServer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NetworkProxyServer value is set to nil even if false is passed
### GetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *Windows10GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmail() bool`

GetAccountsBlockAddingNonMicrosoftAccountEmail returns the AccountsBlockAddingNonMicrosoftAccountEmail field if non-nil, zero value otherwise.

### GetAccountsBlockAddingNonMicrosoftAccountEmailOk

`func (o *Windows10GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmailOk() (bool, bool)`

GetAccountsBlockAddingNonMicrosoftAccountEmailOk returns a tuple with the AccountsBlockAddingNonMicrosoftAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *Windows10GeneralConfiguration) HasAccountsBlockAddingNonMicrosoftAccountEmail() bool`

HasAccountsBlockAddingNonMicrosoftAccountEmail returns a boolean if a field has been set.

### SetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *Windows10GeneralConfiguration) SetAccountsBlockAddingNonMicrosoftAccountEmail(v bool)`

SetAccountsBlockAddingNonMicrosoftAccountEmail gets a reference to the given bool and assigns it to the AccountsBlockAddingNonMicrosoftAccountEmail field.

### GetAntiTheftModeBlocked

`func (o *Windows10GeneralConfiguration) GetAntiTheftModeBlocked() bool`

GetAntiTheftModeBlocked returns the AntiTheftModeBlocked field if non-nil, zero value otherwise.

### GetAntiTheftModeBlockedOk

`func (o *Windows10GeneralConfiguration) GetAntiTheftModeBlockedOk() (bool, bool)`

GetAntiTheftModeBlockedOk returns a tuple with the AntiTheftModeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAntiTheftModeBlocked

`func (o *Windows10GeneralConfiguration) HasAntiTheftModeBlocked() bool`

HasAntiTheftModeBlocked returns a boolean if a field has been set.

### SetAntiTheftModeBlocked

`func (o *Windows10GeneralConfiguration) SetAntiTheftModeBlocked(v bool)`

SetAntiTheftModeBlocked gets a reference to the given bool and assigns it to the AntiTheftModeBlocked field.

### GetBluetoothBlocked

`func (o *Windows10GeneralConfiguration) GetBluetoothBlocked() bool`

GetBluetoothBlocked returns the BluetoothBlocked field if non-nil, zero value otherwise.

### GetBluetoothBlockedOk

`func (o *Windows10GeneralConfiguration) GetBluetoothBlockedOk() (bool, bool)`

GetBluetoothBlockedOk returns a tuple with the BluetoothBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlocked

`func (o *Windows10GeneralConfiguration) HasBluetoothBlocked() bool`

HasBluetoothBlocked returns a boolean if a field has been set.

### SetBluetoothBlocked

`func (o *Windows10GeneralConfiguration) SetBluetoothBlocked(v bool)`

SetBluetoothBlocked gets a reference to the given bool and assigns it to the BluetoothBlocked field.

### GetCameraBlocked

`func (o *Windows10GeneralConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *Windows10GeneralConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *Windows10GeneralConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *Windows10GeneralConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetConnectedDevicesServiceBlocked

`func (o *Windows10GeneralConfiguration) GetConnectedDevicesServiceBlocked() bool`

GetConnectedDevicesServiceBlocked returns the ConnectedDevicesServiceBlocked field if non-nil, zero value otherwise.

### GetConnectedDevicesServiceBlockedOk

`func (o *Windows10GeneralConfiguration) GetConnectedDevicesServiceBlockedOk() (bool, bool)`

GetConnectedDevicesServiceBlockedOk returns a tuple with the ConnectedDevicesServiceBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectedDevicesServiceBlocked

`func (o *Windows10GeneralConfiguration) HasConnectedDevicesServiceBlocked() bool`

HasConnectedDevicesServiceBlocked returns a boolean if a field has been set.

### SetConnectedDevicesServiceBlocked

`func (o *Windows10GeneralConfiguration) SetConnectedDevicesServiceBlocked(v bool)`

SetConnectedDevicesServiceBlocked gets a reference to the given bool and assigns it to the ConnectedDevicesServiceBlocked field.

### GetCertificatesBlockManualRootCertificateInstallation

`func (o *Windows10GeneralConfiguration) GetCertificatesBlockManualRootCertificateInstallation() bool`

GetCertificatesBlockManualRootCertificateInstallation returns the CertificatesBlockManualRootCertificateInstallation field if non-nil, zero value otherwise.

### GetCertificatesBlockManualRootCertificateInstallationOk

`func (o *Windows10GeneralConfiguration) GetCertificatesBlockManualRootCertificateInstallationOk() (bool, bool)`

GetCertificatesBlockManualRootCertificateInstallationOk returns a tuple with the CertificatesBlockManualRootCertificateInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificatesBlockManualRootCertificateInstallation

`func (o *Windows10GeneralConfiguration) HasCertificatesBlockManualRootCertificateInstallation() bool`

HasCertificatesBlockManualRootCertificateInstallation returns a boolean if a field has been set.

### SetCertificatesBlockManualRootCertificateInstallation

`func (o *Windows10GeneralConfiguration) SetCertificatesBlockManualRootCertificateInstallation(v bool)`

SetCertificatesBlockManualRootCertificateInstallation gets a reference to the given bool and assigns it to the CertificatesBlockManualRootCertificateInstallation field.

### GetCopyPasteBlocked

`func (o *Windows10GeneralConfiguration) GetCopyPasteBlocked() bool`

GetCopyPasteBlocked returns the CopyPasteBlocked field if non-nil, zero value otherwise.

### GetCopyPasteBlockedOk

`func (o *Windows10GeneralConfiguration) GetCopyPasteBlockedOk() (bool, bool)`

GetCopyPasteBlockedOk returns a tuple with the CopyPasteBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCopyPasteBlocked

`func (o *Windows10GeneralConfiguration) HasCopyPasteBlocked() bool`

HasCopyPasteBlocked returns a boolean if a field has been set.

### SetCopyPasteBlocked

`func (o *Windows10GeneralConfiguration) SetCopyPasteBlocked(v bool)`

SetCopyPasteBlocked gets a reference to the given bool and assigns it to the CopyPasteBlocked field.

### GetCortanaBlocked

`func (o *Windows10GeneralConfiguration) GetCortanaBlocked() bool`

GetCortanaBlocked returns the CortanaBlocked field if non-nil, zero value otherwise.

### GetCortanaBlockedOk

`func (o *Windows10GeneralConfiguration) GetCortanaBlockedOk() (bool, bool)`

GetCortanaBlockedOk returns a tuple with the CortanaBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCortanaBlocked

`func (o *Windows10GeneralConfiguration) HasCortanaBlocked() bool`

HasCortanaBlocked returns a boolean if a field has been set.

### SetCortanaBlocked

`func (o *Windows10GeneralConfiguration) SetCortanaBlocked(v bool)`

SetCortanaBlocked gets a reference to the given bool and assigns it to the CortanaBlocked field.

### GetDeviceManagementBlockFactoryResetOnMobile

`func (o *Windows10GeneralConfiguration) GetDeviceManagementBlockFactoryResetOnMobile() bool`

GetDeviceManagementBlockFactoryResetOnMobile returns the DeviceManagementBlockFactoryResetOnMobile field if non-nil, zero value otherwise.

### GetDeviceManagementBlockFactoryResetOnMobileOk

`func (o *Windows10GeneralConfiguration) GetDeviceManagementBlockFactoryResetOnMobileOk() (bool, bool)`

GetDeviceManagementBlockFactoryResetOnMobileOk returns a tuple with the DeviceManagementBlockFactoryResetOnMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementBlockFactoryResetOnMobile

`func (o *Windows10GeneralConfiguration) HasDeviceManagementBlockFactoryResetOnMobile() bool`

HasDeviceManagementBlockFactoryResetOnMobile returns a boolean if a field has been set.

### SetDeviceManagementBlockFactoryResetOnMobile

`func (o *Windows10GeneralConfiguration) SetDeviceManagementBlockFactoryResetOnMobile(v bool)`

SetDeviceManagementBlockFactoryResetOnMobile gets a reference to the given bool and assigns it to the DeviceManagementBlockFactoryResetOnMobile field.

### GetDeviceManagementBlockManualUnenroll

`func (o *Windows10GeneralConfiguration) GetDeviceManagementBlockManualUnenroll() bool`

GetDeviceManagementBlockManualUnenroll returns the DeviceManagementBlockManualUnenroll field if non-nil, zero value otherwise.

### GetDeviceManagementBlockManualUnenrollOk

`func (o *Windows10GeneralConfiguration) GetDeviceManagementBlockManualUnenrollOk() (bool, bool)`

GetDeviceManagementBlockManualUnenrollOk returns a tuple with the DeviceManagementBlockManualUnenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementBlockManualUnenroll

`func (o *Windows10GeneralConfiguration) HasDeviceManagementBlockManualUnenroll() bool`

HasDeviceManagementBlockManualUnenroll returns a boolean if a field has been set.

### SetDeviceManagementBlockManualUnenroll

`func (o *Windows10GeneralConfiguration) SetDeviceManagementBlockManualUnenroll(v bool)`

SetDeviceManagementBlockManualUnenroll gets a reference to the given bool and assigns it to the DeviceManagementBlockManualUnenroll field.

### GetSafeSearchFilter

`func (o *Windows10GeneralConfiguration) GetSafeSearchFilter() AnyOfmicrosoftGraphSafeSearchFilterType`

GetSafeSearchFilter returns the SafeSearchFilter field if non-nil, zero value otherwise.

### GetSafeSearchFilterOk

`func (o *Windows10GeneralConfiguration) GetSafeSearchFilterOk() (AnyOfmicrosoftGraphSafeSearchFilterType, bool)`

GetSafeSearchFilterOk returns a tuple with the SafeSearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafeSearchFilter

`func (o *Windows10GeneralConfiguration) HasSafeSearchFilter() bool`

HasSafeSearchFilter returns a boolean if a field has been set.

### SetSafeSearchFilter

`func (o *Windows10GeneralConfiguration) SetSafeSearchFilter(v AnyOfmicrosoftGraphSafeSearchFilterType)`

SetSafeSearchFilter gets a reference to the given AnyOfmicrosoftGraphSafeSearchFilterType and assigns it to the SafeSearchFilter field.

### GetEdgeBlockPopups

`func (o *Windows10GeneralConfiguration) GetEdgeBlockPopups() bool`

GetEdgeBlockPopups returns the EdgeBlockPopups field if non-nil, zero value otherwise.

### GetEdgeBlockPopupsOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockPopupsOk() (bool, bool)`

GetEdgeBlockPopupsOk returns a tuple with the EdgeBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockPopups

`func (o *Windows10GeneralConfiguration) HasEdgeBlockPopups() bool`

HasEdgeBlockPopups returns a boolean if a field has been set.

### SetEdgeBlockPopups

`func (o *Windows10GeneralConfiguration) SetEdgeBlockPopups(v bool)`

SetEdgeBlockPopups gets a reference to the given bool and assigns it to the EdgeBlockPopups field.

### GetEdgeBlockSearchSuggestions

`func (o *Windows10GeneralConfiguration) GetEdgeBlockSearchSuggestions() bool`

GetEdgeBlockSearchSuggestions returns the EdgeBlockSearchSuggestions field if non-nil, zero value otherwise.

### GetEdgeBlockSearchSuggestionsOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockSearchSuggestionsOk() (bool, bool)`

GetEdgeBlockSearchSuggestionsOk returns a tuple with the EdgeBlockSearchSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockSearchSuggestions

`func (o *Windows10GeneralConfiguration) HasEdgeBlockSearchSuggestions() bool`

HasEdgeBlockSearchSuggestions returns a boolean if a field has been set.

### SetEdgeBlockSearchSuggestions

`func (o *Windows10GeneralConfiguration) SetEdgeBlockSearchSuggestions(v bool)`

SetEdgeBlockSearchSuggestions gets a reference to the given bool and assigns it to the EdgeBlockSearchSuggestions field.

### GetEdgeBlockSendingIntranetTrafficToInternetExplorer

`func (o *Windows10GeneralConfiguration) GetEdgeBlockSendingIntranetTrafficToInternetExplorer() bool`

GetEdgeBlockSendingIntranetTrafficToInternetExplorer returns the EdgeBlockSendingIntranetTrafficToInternetExplorer field if non-nil, zero value otherwise.

### GetEdgeBlockSendingIntranetTrafficToInternetExplorerOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockSendingIntranetTrafficToInternetExplorerOk() (bool, bool)`

GetEdgeBlockSendingIntranetTrafficToInternetExplorerOk returns a tuple with the EdgeBlockSendingIntranetTrafficToInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockSendingIntranetTrafficToInternetExplorer

`func (o *Windows10GeneralConfiguration) HasEdgeBlockSendingIntranetTrafficToInternetExplorer() bool`

HasEdgeBlockSendingIntranetTrafficToInternetExplorer returns a boolean if a field has been set.

### SetEdgeBlockSendingIntranetTrafficToInternetExplorer

`func (o *Windows10GeneralConfiguration) SetEdgeBlockSendingIntranetTrafficToInternetExplorer(v bool)`

SetEdgeBlockSendingIntranetTrafficToInternetExplorer gets a reference to the given bool and assigns it to the EdgeBlockSendingIntranetTrafficToInternetExplorer field.

### GetEdgeSendIntranetTrafficToInternetExplorer

`func (o *Windows10GeneralConfiguration) GetEdgeSendIntranetTrafficToInternetExplorer() bool`

GetEdgeSendIntranetTrafficToInternetExplorer returns the EdgeSendIntranetTrafficToInternetExplorer field if non-nil, zero value otherwise.

### GetEdgeSendIntranetTrafficToInternetExplorerOk

`func (o *Windows10GeneralConfiguration) GetEdgeSendIntranetTrafficToInternetExplorerOk() (bool, bool)`

GetEdgeSendIntranetTrafficToInternetExplorerOk returns a tuple with the EdgeSendIntranetTrafficToInternetExplorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeSendIntranetTrafficToInternetExplorer

`func (o *Windows10GeneralConfiguration) HasEdgeSendIntranetTrafficToInternetExplorer() bool`

HasEdgeSendIntranetTrafficToInternetExplorer returns a boolean if a field has been set.

### SetEdgeSendIntranetTrafficToInternetExplorer

`func (o *Windows10GeneralConfiguration) SetEdgeSendIntranetTrafficToInternetExplorer(v bool)`

SetEdgeSendIntranetTrafficToInternetExplorer gets a reference to the given bool and assigns it to the EdgeSendIntranetTrafficToInternetExplorer field.

### GetEdgeRequireSmartScreen

`func (o *Windows10GeneralConfiguration) GetEdgeRequireSmartScreen() bool`

GetEdgeRequireSmartScreen returns the EdgeRequireSmartScreen field if non-nil, zero value otherwise.

### GetEdgeRequireSmartScreenOk

`func (o *Windows10GeneralConfiguration) GetEdgeRequireSmartScreenOk() (bool, bool)`

GetEdgeRequireSmartScreenOk returns a tuple with the EdgeRequireSmartScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeRequireSmartScreen

`func (o *Windows10GeneralConfiguration) HasEdgeRequireSmartScreen() bool`

HasEdgeRequireSmartScreen returns a boolean if a field has been set.

### SetEdgeRequireSmartScreen

`func (o *Windows10GeneralConfiguration) SetEdgeRequireSmartScreen(v bool)`

SetEdgeRequireSmartScreen gets a reference to the given bool and assigns it to the EdgeRequireSmartScreen field.

### GetEdgeEnterpriseModeSiteListLocation

`func (o *Windows10GeneralConfiguration) GetEdgeEnterpriseModeSiteListLocation() string`

GetEdgeEnterpriseModeSiteListLocation returns the EdgeEnterpriseModeSiteListLocation field if non-nil, zero value otherwise.

### GetEdgeEnterpriseModeSiteListLocationOk

`func (o *Windows10GeneralConfiguration) GetEdgeEnterpriseModeSiteListLocationOk() (string, bool)`

GetEdgeEnterpriseModeSiteListLocationOk returns a tuple with the EdgeEnterpriseModeSiteListLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeEnterpriseModeSiteListLocation

`func (o *Windows10GeneralConfiguration) HasEdgeEnterpriseModeSiteListLocation() bool`

HasEdgeEnterpriseModeSiteListLocation returns a boolean if a field has been set.

### SetEdgeEnterpriseModeSiteListLocation

`func (o *Windows10GeneralConfiguration) SetEdgeEnterpriseModeSiteListLocation(v string)`

SetEdgeEnterpriseModeSiteListLocation gets a reference to the given string and assigns it to the EdgeEnterpriseModeSiteListLocation field.

### SetEdgeEnterpriseModeSiteListLocationExplicitNull

`func (o *Windows10GeneralConfiguration) SetEdgeEnterpriseModeSiteListLocationExplicitNull(b bool)`

SetEdgeEnterpriseModeSiteListLocationExplicitNull (un)sets EdgeEnterpriseModeSiteListLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EdgeEnterpriseModeSiteListLocation value is set to nil even if false is passed
### GetEdgeFirstRunUrl

`func (o *Windows10GeneralConfiguration) GetEdgeFirstRunUrl() string`

GetEdgeFirstRunUrl returns the EdgeFirstRunUrl field if non-nil, zero value otherwise.

### GetEdgeFirstRunUrlOk

`func (o *Windows10GeneralConfiguration) GetEdgeFirstRunUrlOk() (string, bool)`

GetEdgeFirstRunUrlOk returns a tuple with the EdgeFirstRunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeFirstRunUrl

`func (o *Windows10GeneralConfiguration) HasEdgeFirstRunUrl() bool`

HasEdgeFirstRunUrl returns a boolean if a field has been set.

### SetEdgeFirstRunUrl

`func (o *Windows10GeneralConfiguration) SetEdgeFirstRunUrl(v string)`

SetEdgeFirstRunUrl gets a reference to the given string and assigns it to the EdgeFirstRunUrl field.

### SetEdgeFirstRunUrlExplicitNull

`func (o *Windows10GeneralConfiguration) SetEdgeFirstRunUrlExplicitNull(b bool)`

SetEdgeFirstRunUrlExplicitNull (un)sets EdgeFirstRunUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EdgeFirstRunUrl value is set to nil even if false is passed
### GetEdgeSearchEngine

`func (o *Windows10GeneralConfiguration) GetEdgeSearchEngine() AnyOfobject`

GetEdgeSearchEngine returns the EdgeSearchEngine field if non-nil, zero value otherwise.

### GetEdgeSearchEngineOk

`func (o *Windows10GeneralConfiguration) GetEdgeSearchEngineOk() (AnyOfobject, bool)`

GetEdgeSearchEngineOk returns a tuple with the EdgeSearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeSearchEngine

`func (o *Windows10GeneralConfiguration) HasEdgeSearchEngine() bool`

HasEdgeSearchEngine returns a boolean if a field has been set.

### SetEdgeSearchEngine

`func (o *Windows10GeneralConfiguration) SetEdgeSearchEngine(v AnyOfobject)`

SetEdgeSearchEngine gets a reference to the given AnyOfobject and assigns it to the EdgeSearchEngine field.

### SetEdgeSearchEngineExplicitNull

`func (o *Windows10GeneralConfiguration) SetEdgeSearchEngineExplicitNull(b bool)`

SetEdgeSearchEngineExplicitNull (un)sets EdgeSearchEngine to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EdgeSearchEngine value is set to nil even if false is passed
### GetEdgeHomepageUrls

`func (o *Windows10GeneralConfiguration) GetEdgeHomepageUrls() []string`

GetEdgeHomepageUrls returns the EdgeHomepageUrls field if non-nil, zero value otherwise.

### GetEdgeHomepageUrlsOk

`func (o *Windows10GeneralConfiguration) GetEdgeHomepageUrlsOk() ([]string, bool)`

GetEdgeHomepageUrlsOk returns a tuple with the EdgeHomepageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeHomepageUrls

`func (o *Windows10GeneralConfiguration) HasEdgeHomepageUrls() bool`

HasEdgeHomepageUrls returns a boolean if a field has been set.

### SetEdgeHomepageUrls

`func (o *Windows10GeneralConfiguration) SetEdgeHomepageUrls(v []string)`

SetEdgeHomepageUrls gets a reference to the given []string and assigns it to the EdgeHomepageUrls field.

### GetEdgeBlockAccessToAboutFlags

`func (o *Windows10GeneralConfiguration) GetEdgeBlockAccessToAboutFlags() bool`

GetEdgeBlockAccessToAboutFlags returns the EdgeBlockAccessToAboutFlags field if non-nil, zero value otherwise.

### GetEdgeBlockAccessToAboutFlagsOk

`func (o *Windows10GeneralConfiguration) GetEdgeBlockAccessToAboutFlagsOk() (bool, bool)`

GetEdgeBlockAccessToAboutFlagsOk returns a tuple with the EdgeBlockAccessToAboutFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdgeBlockAccessToAboutFlags

`func (o *Windows10GeneralConfiguration) HasEdgeBlockAccessToAboutFlags() bool`

HasEdgeBlockAccessToAboutFlags returns a boolean if a field has been set.

### SetEdgeBlockAccessToAboutFlags

`func (o *Windows10GeneralConfiguration) SetEdgeBlockAccessToAboutFlags(v bool)`

SetEdgeBlockAccessToAboutFlags gets a reference to the given bool and assigns it to the EdgeBlockAccessToAboutFlags field.

### GetSmartScreenBlockPromptOverride

`func (o *Windows10GeneralConfiguration) GetSmartScreenBlockPromptOverride() bool`

GetSmartScreenBlockPromptOverride returns the SmartScreenBlockPromptOverride field if non-nil, zero value otherwise.

### GetSmartScreenBlockPromptOverrideOk

`func (o *Windows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideOk() (bool, bool)`

GetSmartScreenBlockPromptOverrideOk returns a tuple with the SmartScreenBlockPromptOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenBlockPromptOverride

`func (o *Windows10GeneralConfiguration) HasSmartScreenBlockPromptOverride() bool`

HasSmartScreenBlockPromptOverride returns a boolean if a field has been set.

### SetSmartScreenBlockPromptOverride

`func (o *Windows10GeneralConfiguration) SetSmartScreenBlockPromptOverride(v bool)`

SetSmartScreenBlockPromptOverride gets a reference to the given bool and assigns it to the SmartScreenBlockPromptOverride field.

### GetSmartScreenBlockPromptOverrideForFiles

`func (o *Windows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideForFiles() bool`

GetSmartScreenBlockPromptOverrideForFiles returns the SmartScreenBlockPromptOverrideForFiles field if non-nil, zero value otherwise.

### GetSmartScreenBlockPromptOverrideForFilesOk

`func (o *Windows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideForFilesOk() (bool, bool)`

GetSmartScreenBlockPromptOverrideForFilesOk returns a tuple with the SmartScreenBlockPromptOverrideForFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenBlockPromptOverrideForFiles

`func (o *Windows10GeneralConfiguration) HasSmartScreenBlockPromptOverrideForFiles() bool`

HasSmartScreenBlockPromptOverrideForFiles returns a boolean if a field has been set.

### SetSmartScreenBlockPromptOverrideForFiles

`func (o *Windows10GeneralConfiguration) SetSmartScreenBlockPromptOverrideForFiles(v bool)`

SetSmartScreenBlockPromptOverrideForFiles gets a reference to the given bool and assigns it to the SmartScreenBlockPromptOverrideForFiles field.

### GetWebRtcBlockLocalhostIpAddress

`func (o *Windows10GeneralConfiguration) GetWebRtcBlockLocalhostIpAddress() bool`

GetWebRtcBlockLocalhostIpAddress returns the WebRtcBlockLocalhostIpAddress field if non-nil, zero value otherwise.

### GetWebRtcBlockLocalhostIpAddressOk

`func (o *Windows10GeneralConfiguration) GetWebRtcBlockLocalhostIpAddressOk() (bool, bool)`

GetWebRtcBlockLocalhostIpAddressOk returns a tuple with the WebRtcBlockLocalhostIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebRtcBlockLocalhostIpAddress

`func (o *Windows10GeneralConfiguration) HasWebRtcBlockLocalhostIpAddress() bool`

HasWebRtcBlockLocalhostIpAddress returns a boolean if a field has been set.

### SetWebRtcBlockLocalhostIpAddress

`func (o *Windows10GeneralConfiguration) SetWebRtcBlockLocalhostIpAddress(v bool)`

SetWebRtcBlockLocalhostIpAddress gets a reference to the given bool and assigns it to the WebRtcBlockLocalhostIpAddress field.

### GetInternetSharingBlocked

`func (o *Windows10GeneralConfiguration) GetInternetSharingBlocked() bool`

GetInternetSharingBlocked returns the InternetSharingBlocked field if non-nil, zero value otherwise.

### GetInternetSharingBlockedOk

`func (o *Windows10GeneralConfiguration) GetInternetSharingBlockedOk() (bool, bool)`

GetInternetSharingBlockedOk returns a tuple with the InternetSharingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetSharingBlocked

`func (o *Windows10GeneralConfiguration) HasInternetSharingBlocked() bool`

HasInternetSharingBlocked returns a boolean if a field has been set.

### SetInternetSharingBlocked

`func (o *Windows10GeneralConfiguration) SetInternetSharingBlocked(v bool)`

SetInternetSharingBlocked gets a reference to the given bool and assigns it to the InternetSharingBlocked field.

### GetSettingsBlockAddProvisioningPackage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockAddProvisioningPackage() bool`

GetSettingsBlockAddProvisioningPackage returns the SettingsBlockAddProvisioningPackage field if non-nil, zero value otherwise.

### GetSettingsBlockAddProvisioningPackageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockAddProvisioningPackageOk() (bool, bool)`

GetSettingsBlockAddProvisioningPackageOk returns a tuple with the SettingsBlockAddProvisioningPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockAddProvisioningPackage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockAddProvisioningPackage() bool`

HasSettingsBlockAddProvisioningPackage returns a boolean if a field has been set.

### SetSettingsBlockAddProvisioningPackage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockAddProvisioningPackage(v bool)`

SetSettingsBlockAddProvisioningPackage gets a reference to the given bool and assigns it to the SettingsBlockAddProvisioningPackage field.

### GetSettingsBlockRemoveProvisioningPackage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockRemoveProvisioningPackage() bool`

GetSettingsBlockRemoveProvisioningPackage returns the SettingsBlockRemoveProvisioningPackage field if non-nil, zero value otherwise.

### GetSettingsBlockRemoveProvisioningPackageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockRemoveProvisioningPackageOk() (bool, bool)`

GetSettingsBlockRemoveProvisioningPackageOk returns a tuple with the SettingsBlockRemoveProvisioningPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockRemoveProvisioningPackage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockRemoveProvisioningPackage() bool`

HasSettingsBlockRemoveProvisioningPackage returns a boolean if a field has been set.

### SetSettingsBlockRemoveProvisioningPackage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockRemoveProvisioningPackage(v bool)`

SetSettingsBlockRemoveProvisioningPackage gets a reference to the given bool and assigns it to the SettingsBlockRemoveProvisioningPackage field.

### GetSettingsBlockChangeSystemTime

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangeSystemTime() bool`

GetSettingsBlockChangeSystemTime returns the SettingsBlockChangeSystemTime field if non-nil, zero value otherwise.

### GetSettingsBlockChangeSystemTimeOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangeSystemTimeOk() (bool, bool)`

GetSettingsBlockChangeSystemTimeOk returns a tuple with the SettingsBlockChangeSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangeSystemTime

`func (o *Windows10GeneralConfiguration) HasSettingsBlockChangeSystemTime() bool`

HasSettingsBlockChangeSystemTime returns a boolean if a field has been set.

### SetSettingsBlockChangeSystemTime

`func (o *Windows10GeneralConfiguration) SetSettingsBlockChangeSystemTime(v bool)`

SetSettingsBlockChangeSystemTime gets a reference to the given bool and assigns it to the SettingsBlockChangeSystemTime field.

### GetSettingsBlockEditDeviceName

`func (o *Windows10GeneralConfiguration) GetSettingsBlockEditDeviceName() bool`

GetSettingsBlockEditDeviceName returns the SettingsBlockEditDeviceName field if non-nil, zero value otherwise.

### GetSettingsBlockEditDeviceNameOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockEditDeviceNameOk() (bool, bool)`

GetSettingsBlockEditDeviceNameOk returns a tuple with the SettingsBlockEditDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockEditDeviceName

`func (o *Windows10GeneralConfiguration) HasSettingsBlockEditDeviceName() bool`

HasSettingsBlockEditDeviceName returns a boolean if a field has been set.

### SetSettingsBlockEditDeviceName

`func (o *Windows10GeneralConfiguration) SetSettingsBlockEditDeviceName(v bool)`

SetSettingsBlockEditDeviceName gets a reference to the given bool and assigns it to the SettingsBlockEditDeviceName field.

### GetSettingsBlockChangeRegion

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangeRegion() bool`

GetSettingsBlockChangeRegion returns the SettingsBlockChangeRegion field if non-nil, zero value otherwise.

### GetSettingsBlockChangeRegionOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangeRegionOk() (bool, bool)`

GetSettingsBlockChangeRegionOk returns a tuple with the SettingsBlockChangeRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangeRegion

`func (o *Windows10GeneralConfiguration) HasSettingsBlockChangeRegion() bool`

HasSettingsBlockChangeRegion returns a boolean if a field has been set.

### SetSettingsBlockChangeRegion

`func (o *Windows10GeneralConfiguration) SetSettingsBlockChangeRegion(v bool)`

SetSettingsBlockChangeRegion gets a reference to the given bool and assigns it to the SettingsBlockChangeRegion field.

### GetSettingsBlockChangeLanguage

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangeLanguage() bool`

GetSettingsBlockChangeLanguage returns the SettingsBlockChangeLanguage field if non-nil, zero value otherwise.

### GetSettingsBlockChangeLanguageOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangeLanguageOk() (bool, bool)`

GetSettingsBlockChangeLanguageOk returns a tuple with the SettingsBlockChangeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangeLanguage

`func (o *Windows10GeneralConfiguration) HasSettingsBlockChangeLanguage() bool`

HasSettingsBlockChangeLanguage returns a boolean if a field has been set.

### SetSettingsBlockChangeLanguage

`func (o *Windows10GeneralConfiguration) SetSettingsBlockChangeLanguage(v bool)`

SetSettingsBlockChangeLanguage gets a reference to the given bool and assigns it to the SettingsBlockChangeLanguage field.

### GetSettingsBlockChangePowerSleep

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangePowerSleep() bool`

GetSettingsBlockChangePowerSleep returns the SettingsBlockChangePowerSleep field if non-nil, zero value otherwise.

### GetSettingsBlockChangePowerSleepOk

`func (o *Windows10GeneralConfiguration) GetSettingsBlockChangePowerSleepOk() (bool, bool)`

GetSettingsBlockChangePowerSleepOk returns a tuple with the SettingsBlockChangePowerSleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingsBlockChangePowerSleep

`func (o *Windows10GeneralConfiguration) HasSettingsBlockChangePowerSleep() bool`

HasSettingsBlockChangePowerSleep returns a boolean if a field has been set.

### SetSettingsBlockChangePowerSleep

`func (o *Windows10GeneralConfiguration) SetSettingsBlockChangePowerSleep(v bool)`

SetSettingsBlockChangePowerSleep gets a reference to the given bool and assigns it to the SettingsBlockChangePowerSleep field.

### GetLocationServicesBlocked

`func (o *Windows10GeneralConfiguration) GetLocationServicesBlocked() bool`

GetLocationServicesBlocked returns the LocationServicesBlocked field if non-nil, zero value otherwise.

### GetLocationServicesBlockedOk

`func (o *Windows10GeneralConfiguration) GetLocationServicesBlockedOk() (bool, bool)`

GetLocationServicesBlockedOk returns a tuple with the LocationServicesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationServicesBlocked

`func (o *Windows10GeneralConfiguration) HasLocationServicesBlocked() bool`

HasLocationServicesBlocked returns a boolean if a field has been set.

### SetLocationServicesBlocked

`func (o *Windows10GeneralConfiguration) SetLocationServicesBlocked(v bool)`

SetLocationServicesBlocked gets a reference to the given bool and assigns it to the LocationServicesBlocked field.

### GetMicrosoftAccountBlocked

`func (o *Windows10GeneralConfiguration) GetMicrosoftAccountBlocked() bool`

GetMicrosoftAccountBlocked returns the MicrosoftAccountBlocked field if non-nil, zero value otherwise.

### GetMicrosoftAccountBlockedOk

`func (o *Windows10GeneralConfiguration) GetMicrosoftAccountBlockedOk() (bool, bool)`

GetMicrosoftAccountBlockedOk returns a tuple with the MicrosoftAccountBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftAccountBlocked

`func (o *Windows10GeneralConfiguration) HasMicrosoftAccountBlocked() bool`

HasMicrosoftAccountBlocked returns a boolean if a field has been set.

### SetMicrosoftAccountBlocked

`func (o *Windows10GeneralConfiguration) SetMicrosoftAccountBlocked(v bool)`

SetMicrosoftAccountBlocked gets a reference to the given bool and assigns it to the MicrosoftAccountBlocked field.

### GetMicrosoftAccountBlockSettingsSync

`func (o *Windows10GeneralConfiguration) GetMicrosoftAccountBlockSettingsSync() bool`

GetMicrosoftAccountBlockSettingsSync returns the MicrosoftAccountBlockSettingsSync field if non-nil, zero value otherwise.

### GetMicrosoftAccountBlockSettingsSyncOk

`func (o *Windows10GeneralConfiguration) GetMicrosoftAccountBlockSettingsSyncOk() (bool, bool)`

GetMicrosoftAccountBlockSettingsSyncOk returns a tuple with the MicrosoftAccountBlockSettingsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftAccountBlockSettingsSync

`func (o *Windows10GeneralConfiguration) HasMicrosoftAccountBlockSettingsSync() bool`

HasMicrosoftAccountBlockSettingsSync returns a boolean if a field has been set.

### SetMicrosoftAccountBlockSettingsSync

`func (o *Windows10GeneralConfiguration) SetMicrosoftAccountBlockSettingsSync(v bool)`

SetMicrosoftAccountBlockSettingsSync gets a reference to the given bool and assigns it to the MicrosoftAccountBlockSettingsSync field.

### GetNfcBlocked

`func (o *Windows10GeneralConfiguration) GetNfcBlocked() bool`

GetNfcBlocked returns the NfcBlocked field if non-nil, zero value otherwise.

### GetNfcBlockedOk

`func (o *Windows10GeneralConfiguration) GetNfcBlockedOk() (bool, bool)`

GetNfcBlockedOk returns a tuple with the NfcBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNfcBlocked

`func (o *Windows10GeneralConfiguration) HasNfcBlocked() bool`

HasNfcBlocked returns a boolean if a field has been set.

### SetNfcBlocked

`func (o *Windows10GeneralConfiguration) SetNfcBlocked(v bool)`

SetNfcBlocked gets a reference to the given bool and assigns it to the NfcBlocked field.

### GetResetProtectionModeBlocked

`func (o *Windows10GeneralConfiguration) GetResetProtectionModeBlocked() bool`

GetResetProtectionModeBlocked returns the ResetProtectionModeBlocked field if non-nil, zero value otherwise.

### GetResetProtectionModeBlockedOk

`func (o *Windows10GeneralConfiguration) GetResetProtectionModeBlockedOk() (bool, bool)`

GetResetProtectionModeBlockedOk returns a tuple with the ResetProtectionModeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResetProtectionModeBlocked

`func (o *Windows10GeneralConfiguration) HasResetProtectionModeBlocked() bool`

HasResetProtectionModeBlocked returns a boolean if a field has been set.

### SetResetProtectionModeBlocked

`func (o *Windows10GeneralConfiguration) SetResetProtectionModeBlocked(v bool)`

SetResetProtectionModeBlocked gets a reference to the given bool and assigns it to the ResetProtectionModeBlocked field.

### GetScreenCaptureBlocked

`func (o *Windows10GeneralConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *Windows10GeneralConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *Windows10GeneralConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *Windows10GeneralConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetStorageBlockRemovableStorage

`func (o *Windows10GeneralConfiguration) GetStorageBlockRemovableStorage() bool`

GetStorageBlockRemovableStorage returns the StorageBlockRemovableStorage field if non-nil, zero value otherwise.

### GetStorageBlockRemovableStorageOk

`func (o *Windows10GeneralConfiguration) GetStorageBlockRemovableStorageOk() (bool, bool)`

GetStorageBlockRemovableStorageOk returns a tuple with the StorageBlockRemovableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockRemovableStorage

`func (o *Windows10GeneralConfiguration) HasStorageBlockRemovableStorage() bool`

HasStorageBlockRemovableStorage returns a boolean if a field has been set.

### SetStorageBlockRemovableStorage

`func (o *Windows10GeneralConfiguration) SetStorageBlockRemovableStorage(v bool)`

SetStorageBlockRemovableStorage gets a reference to the given bool and assigns it to the StorageBlockRemovableStorage field.

### GetStorageRequireMobileDeviceEncryption

`func (o *Windows10GeneralConfiguration) GetStorageRequireMobileDeviceEncryption() bool`

GetStorageRequireMobileDeviceEncryption returns the StorageRequireMobileDeviceEncryption field if non-nil, zero value otherwise.

### GetStorageRequireMobileDeviceEncryptionOk

`func (o *Windows10GeneralConfiguration) GetStorageRequireMobileDeviceEncryptionOk() (bool, bool)`

GetStorageRequireMobileDeviceEncryptionOk returns a tuple with the StorageRequireMobileDeviceEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireMobileDeviceEncryption

`func (o *Windows10GeneralConfiguration) HasStorageRequireMobileDeviceEncryption() bool`

HasStorageRequireMobileDeviceEncryption returns a boolean if a field has been set.

### SetStorageRequireMobileDeviceEncryption

`func (o *Windows10GeneralConfiguration) SetStorageRequireMobileDeviceEncryption(v bool)`

SetStorageRequireMobileDeviceEncryption gets a reference to the given bool and assigns it to the StorageRequireMobileDeviceEncryption field.

### GetUsbBlocked

`func (o *Windows10GeneralConfiguration) GetUsbBlocked() bool`

GetUsbBlocked returns the UsbBlocked field if non-nil, zero value otherwise.

### GetUsbBlockedOk

`func (o *Windows10GeneralConfiguration) GetUsbBlockedOk() (bool, bool)`

GetUsbBlockedOk returns a tuple with the UsbBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsbBlocked

`func (o *Windows10GeneralConfiguration) HasUsbBlocked() bool`

HasUsbBlocked returns a boolean if a field has been set.

### SetUsbBlocked

`func (o *Windows10GeneralConfiguration) SetUsbBlocked(v bool)`

SetUsbBlocked gets a reference to the given bool and assigns it to the UsbBlocked field.

### GetVoiceRecordingBlocked

`func (o *Windows10GeneralConfiguration) GetVoiceRecordingBlocked() bool`

GetVoiceRecordingBlocked returns the VoiceRecordingBlocked field if non-nil, zero value otherwise.

### GetVoiceRecordingBlockedOk

`func (o *Windows10GeneralConfiguration) GetVoiceRecordingBlockedOk() (bool, bool)`

GetVoiceRecordingBlockedOk returns a tuple with the VoiceRecordingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVoiceRecordingBlocked

`func (o *Windows10GeneralConfiguration) HasVoiceRecordingBlocked() bool`

HasVoiceRecordingBlocked returns a boolean if a field has been set.

### SetVoiceRecordingBlocked

`func (o *Windows10GeneralConfiguration) SetVoiceRecordingBlocked(v bool)`

SetVoiceRecordingBlocked gets a reference to the given bool and assigns it to the VoiceRecordingBlocked field.

### GetWiFiBlockAutomaticConnectHotspots

`func (o *Windows10GeneralConfiguration) GetWiFiBlockAutomaticConnectHotspots() bool`

GetWiFiBlockAutomaticConnectHotspots returns the WiFiBlockAutomaticConnectHotspots field if non-nil, zero value otherwise.

### GetWiFiBlockAutomaticConnectHotspotsOk

`func (o *Windows10GeneralConfiguration) GetWiFiBlockAutomaticConnectHotspotsOk() (bool, bool)`

GetWiFiBlockAutomaticConnectHotspotsOk returns a tuple with the WiFiBlockAutomaticConnectHotspots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlockAutomaticConnectHotspots

`func (o *Windows10GeneralConfiguration) HasWiFiBlockAutomaticConnectHotspots() bool`

HasWiFiBlockAutomaticConnectHotspots returns a boolean if a field has been set.

### SetWiFiBlockAutomaticConnectHotspots

`func (o *Windows10GeneralConfiguration) SetWiFiBlockAutomaticConnectHotspots(v bool)`

SetWiFiBlockAutomaticConnectHotspots gets a reference to the given bool and assigns it to the WiFiBlockAutomaticConnectHotspots field.

### GetWiFiBlocked

`func (o *Windows10GeneralConfiguration) GetWiFiBlocked() bool`

GetWiFiBlocked returns the WiFiBlocked field if non-nil, zero value otherwise.

### GetWiFiBlockedOk

`func (o *Windows10GeneralConfiguration) GetWiFiBlockedOk() (bool, bool)`

GetWiFiBlockedOk returns a tuple with the WiFiBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlocked

`func (o *Windows10GeneralConfiguration) HasWiFiBlocked() bool`

HasWiFiBlocked returns a boolean if a field has been set.

### SetWiFiBlocked

`func (o *Windows10GeneralConfiguration) SetWiFiBlocked(v bool)`

SetWiFiBlocked gets a reference to the given bool and assigns it to the WiFiBlocked field.

### GetWiFiBlockManualConfiguration

`func (o *Windows10GeneralConfiguration) GetWiFiBlockManualConfiguration() bool`

GetWiFiBlockManualConfiguration returns the WiFiBlockManualConfiguration field if non-nil, zero value otherwise.

### GetWiFiBlockManualConfigurationOk

`func (o *Windows10GeneralConfiguration) GetWiFiBlockManualConfigurationOk() (bool, bool)`

GetWiFiBlockManualConfigurationOk returns a tuple with the WiFiBlockManualConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiBlockManualConfiguration

`func (o *Windows10GeneralConfiguration) HasWiFiBlockManualConfiguration() bool`

HasWiFiBlockManualConfiguration returns a boolean if a field has been set.

### SetWiFiBlockManualConfiguration

`func (o *Windows10GeneralConfiguration) SetWiFiBlockManualConfiguration(v bool)`

SetWiFiBlockManualConfiguration gets a reference to the given bool and assigns it to the WiFiBlockManualConfiguration field.

### GetWiFiScanInterval

`func (o *Windows10GeneralConfiguration) GetWiFiScanInterval() int32`

GetWiFiScanInterval returns the WiFiScanInterval field if non-nil, zero value otherwise.

### GetWiFiScanIntervalOk

`func (o *Windows10GeneralConfiguration) GetWiFiScanIntervalOk() (int32, bool)`

GetWiFiScanIntervalOk returns a tuple with the WiFiScanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiScanInterval

`func (o *Windows10GeneralConfiguration) HasWiFiScanInterval() bool`

HasWiFiScanInterval returns a boolean if a field has been set.

### SetWiFiScanInterval

`func (o *Windows10GeneralConfiguration) SetWiFiScanInterval(v int32)`

SetWiFiScanInterval gets a reference to the given int32 and assigns it to the WiFiScanInterval field.

### SetWiFiScanIntervalExplicitNull

`func (o *Windows10GeneralConfiguration) SetWiFiScanIntervalExplicitNull(b bool)`

SetWiFiScanIntervalExplicitNull (un)sets WiFiScanInterval to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WiFiScanInterval value is set to nil even if false is passed
### GetWirelessDisplayBlockProjectionToThisDevice

`func (o *Windows10GeneralConfiguration) GetWirelessDisplayBlockProjectionToThisDevice() bool`

GetWirelessDisplayBlockProjectionToThisDevice returns the WirelessDisplayBlockProjectionToThisDevice field if non-nil, zero value otherwise.

### GetWirelessDisplayBlockProjectionToThisDeviceOk

`func (o *Windows10GeneralConfiguration) GetWirelessDisplayBlockProjectionToThisDeviceOk() (bool, bool)`

GetWirelessDisplayBlockProjectionToThisDeviceOk returns a tuple with the WirelessDisplayBlockProjectionToThisDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWirelessDisplayBlockProjectionToThisDevice

`func (o *Windows10GeneralConfiguration) HasWirelessDisplayBlockProjectionToThisDevice() bool`

HasWirelessDisplayBlockProjectionToThisDevice returns a boolean if a field has been set.

### SetWirelessDisplayBlockProjectionToThisDevice

`func (o *Windows10GeneralConfiguration) SetWirelessDisplayBlockProjectionToThisDevice(v bool)`

SetWirelessDisplayBlockProjectionToThisDevice gets a reference to the given bool and assigns it to the WirelessDisplayBlockProjectionToThisDevice field.

### GetWirelessDisplayBlockUserInputFromReceiver

`func (o *Windows10GeneralConfiguration) GetWirelessDisplayBlockUserInputFromReceiver() bool`

GetWirelessDisplayBlockUserInputFromReceiver returns the WirelessDisplayBlockUserInputFromReceiver field if non-nil, zero value otherwise.

### GetWirelessDisplayBlockUserInputFromReceiverOk

`func (o *Windows10GeneralConfiguration) GetWirelessDisplayBlockUserInputFromReceiverOk() (bool, bool)`

GetWirelessDisplayBlockUserInputFromReceiverOk returns a tuple with the WirelessDisplayBlockUserInputFromReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWirelessDisplayBlockUserInputFromReceiver

`func (o *Windows10GeneralConfiguration) HasWirelessDisplayBlockUserInputFromReceiver() bool`

HasWirelessDisplayBlockUserInputFromReceiver returns a boolean if a field has been set.

### SetWirelessDisplayBlockUserInputFromReceiver

`func (o *Windows10GeneralConfiguration) SetWirelessDisplayBlockUserInputFromReceiver(v bool)`

SetWirelessDisplayBlockUserInputFromReceiver gets a reference to the given bool and assigns it to the WirelessDisplayBlockUserInputFromReceiver field.

### GetWirelessDisplayRequirePinForPairing

`func (o *Windows10GeneralConfiguration) GetWirelessDisplayRequirePinForPairing() bool`

GetWirelessDisplayRequirePinForPairing returns the WirelessDisplayRequirePinForPairing field if non-nil, zero value otherwise.

### GetWirelessDisplayRequirePinForPairingOk

`func (o *Windows10GeneralConfiguration) GetWirelessDisplayRequirePinForPairingOk() (bool, bool)`

GetWirelessDisplayRequirePinForPairingOk returns a tuple with the WirelessDisplayRequirePinForPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWirelessDisplayRequirePinForPairing

`func (o *Windows10GeneralConfiguration) HasWirelessDisplayRequirePinForPairing() bool`

HasWirelessDisplayRequirePinForPairing returns a boolean if a field has been set.

### SetWirelessDisplayRequirePinForPairing

`func (o *Windows10GeneralConfiguration) SetWirelessDisplayRequirePinForPairing(v bool)`

SetWirelessDisplayRequirePinForPairing gets a reference to the given bool and assigns it to the WirelessDisplayRequirePinForPairing field.

### GetWindowsStoreBlocked

`func (o *Windows10GeneralConfiguration) GetWindowsStoreBlocked() bool`

GetWindowsStoreBlocked returns the WindowsStoreBlocked field if non-nil, zero value otherwise.

### GetWindowsStoreBlockedOk

`func (o *Windows10GeneralConfiguration) GetWindowsStoreBlockedOk() (bool, bool)`

GetWindowsStoreBlockedOk returns a tuple with the WindowsStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreBlocked

`func (o *Windows10GeneralConfiguration) HasWindowsStoreBlocked() bool`

HasWindowsStoreBlocked returns a boolean if a field has been set.

### SetWindowsStoreBlocked

`func (o *Windows10GeneralConfiguration) SetWindowsStoreBlocked(v bool)`

SetWindowsStoreBlocked gets a reference to the given bool and assigns it to the WindowsStoreBlocked field.

### GetAppsAllowTrustedAppsSideloading

`func (o *Windows10GeneralConfiguration) GetAppsAllowTrustedAppsSideloading() AnyOfmicrosoftGraphStateManagementSetting`

GetAppsAllowTrustedAppsSideloading returns the AppsAllowTrustedAppsSideloading field if non-nil, zero value otherwise.

### GetAppsAllowTrustedAppsSideloadingOk

`func (o *Windows10GeneralConfiguration) GetAppsAllowTrustedAppsSideloadingOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetAppsAllowTrustedAppsSideloadingOk returns a tuple with the AppsAllowTrustedAppsSideloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsAllowTrustedAppsSideloading

`func (o *Windows10GeneralConfiguration) HasAppsAllowTrustedAppsSideloading() bool`

HasAppsAllowTrustedAppsSideloading returns a boolean if a field has been set.

### SetAppsAllowTrustedAppsSideloading

`func (o *Windows10GeneralConfiguration) SetAppsAllowTrustedAppsSideloading(v AnyOfmicrosoftGraphStateManagementSetting)`

SetAppsAllowTrustedAppsSideloading gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the AppsAllowTrustedAppsSideloading field.

### GetWindowsStoreBlockAutoUpdate

`func (o *Windows10GeneralConfiguration) GetWindowsStoreBlockAutoUpdate() bool`

GetWindowsStoreBlockAutoUpdate returns the WindowsStoreBlockAutoUpdate field if non-nil, zero value otherwise.

### GetWindowsStoreBlockAutoUpdateOk

`func (o *Windows10GeneralConfiguration) GetWindowsStoreBlockAutoUpdateOk() (bool, bool)`

GetWindowsStoreBlockAutoUpdateOk returns a tuple with the WindowsStoreBlockAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreBlockAutoUpdate

`func (o *Windows10GeneralConfiguration) HasWindowsStoreBlockAutoUpdate() bool`

HasWindowsStoreBlockAutoUpdate returns a boolean if a field has been set.

### SetWindowsStoreBlockAutoUpdate

`func (o *Windows10GeneralConfiguration) SetWindowsStoreBlockAutoUpdate(v bool)`

SetWindowsStoreBlockAutoUpdate gets a reference to the given bool and assigns it to the WindowsStoreBlockAutoUpdate field.

### GetDeveloperUnlockSetting

`func (o *Windows10GeneralConfiguration) GetDeveloperUnlockSetting() AnyOfmicrosoftGraphStateManagementSetting`

GetDeveloperUnlockSetting returns the DeveloperUnlockSetting field if non-nil, zero value otherwise.

### GetDeveloperUnlockSettingOk

`func (o *Windows10GeneralConfiguration) GetDeveloperUnlockSettingOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetDeveloperUnlockSettingOk returns a tuple with the DeveloperUnlockSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloperUnlockSetting

`func (o *Windows10GeneralConfiguration) HasDeveloperUnlockSetting() bool`

HasDeveloperUnlockSetting returns a boolean if a field has been set.

### SetDeveloperUnlockSetting

`func (o *Windows10GeneralConfiguration) SetDeveloperUnlockSetting(v AnyOfmicrosoftGraphStateManagementSetting)`

SetDeveloperUnlockSetting gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the DeveloperUnlockSetting field.

### GetSharedUserAppDataAllowed

`func (o *Windows10GeneralConfiguration) GetSharedUserAppDataAllowed() bool`

GetSharedUserAppDataAllowed returns the SharedUserAppDataAllowed field if non-nil, zero value otherwise.

### GetSharedUserAppDataAllowedOk

`func (o *Windows10GeneralConfiguration) GetSharedUserAppDataAllowedOk() (bool, bool)`

GetSharedUserAppDataAllowedOk returns a tuple with the SharedUserAppDataAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedUserAppDataAllowed

`func (o *Windows10GeneralConfiguration) HasSharedUserAppDataAllowed() bool`

HasSharedUserAppDataAllowed returns a boolean if a field has been set.

### SetSharedUserAppDataAllowed

`func (o *Windows10GeneralConfiguration) SetSharedUserAppDataAllowed(v bool)`

SetSharedUserAppDataAllowed gets a reference to the given bool and assigns it to the SharedUserAppDataAllowed field.

### GetAppsBlockWindowsStoreOriginatedApps

`func (o *Windows10GeneralConfiguration) GetAppsBlockWindowsStoreOriginatedApps() bool`

GetAppsBlockWindowsStoreOriginatedApps returns the AppsBlockWindowsStoreOriginatedApps field if non-nil, zero value otherwise.

### GetAppsBlockWindowsStoreOriginatedAppsOk

`func (o *Windows10GeneralConfiguration) GetAppsBlockWindowsStoreOriginatedAppsOk() (bool, bool)`

GetAppsBlockWindowsStoreOriginatedAppsOk returns a tuple with the AppsBlockWindowsStoreOriginatedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockWindowsStoreOriginatedApps

`func (o *Windows10GeneralConfiguration) HasAppsBlockWindowsStoreOriginatedApps() bool`

HasAppsBlockWindowsStoreOriginatedApps returns a boolean if a field has been set.

### SetAppsBlockWindowsStoreOriginatedApps

`func (o *Windows10GeneralConfiguration) SetAppsBlockWindowsStoreOriginatedApps(v bool)`

SetAppsBlockWindowsStoreOriginatedApps gets a reference to the given bool and assigns it to the AppsBlockWindowsStoreOriginatedApps field.

### GetWindowsStoreEnablePrivateStoreOnly

`func (o *Windows10GeneralConfiguration) GetWindowsStoreEnablePrivateStoreOnly() bool`

GetWindowsStoreEnablePrivateStoreOnly returns the WindowsStoreEnablePrivateStoreOnly field if non-nil, zero value otherwise.

### GetWindowsStoreEnablePrivateStoreOnlyOk

`func (o *Windows10GeneralConfiguration) GetWindowsStoreEnablePrivateStoreOnlyOk() (bool, bool)`

GetWindowsStoreEnablePrivateStoreOnlyOk returns a tuple with the WindowsStoreEnablePrivateStoreOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreEnablePrivateStoreOnly

`func (o *Windows10GeneralConfiguration) HasWindowsStoreEnablePrivateStoreOnly() bool`

HasWindowsStoreEnablePrivateStoreOnly returns a boolean if a field has been set.

### SetWindowsStoreEnablePrivateStoreOnly

`func (o *Windows10GeneralConfiguration) SetWindowsStoreEnablePrivateStoreOnly(v bool)`

SetWindowsStoreEnablePrivateStoreOnly gets a reference to the given bool and assigns it to the WindowsStoreEnablePrivateStoreOnly field.

### GetStorageRestrictAppDataToSystemVolume

`func (o *Windows10GeneralConfiguration) GetStorageRestrictAppDataToSystemVolume() bool`

GetStorageRestrictAppDataToSystemVolume returns the StorageRestrictAppDataToSystemVolume field if non-nil, zero value otherwise.

### GetStorageRestrictAppDataToSystemVolumeOk

`func (o *Windows10GeneralConfiguration) GetStorageRestrictAppDataToSystemVolumeOk() (bool, bool)`

GetStorageRestrictAppDataToSystemVolumeOk returns a tuple with the StorageRestrictAppDataToSystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRestrictAppDataToSystemVolume

`func (o *Windows10GeneralConfiguration) HasStorageRestrictAppDataToSystemVolume() bool`

HasStorageRestrictAppDataToSystemVolume returns a boolean if a field has been set.

### SetStorageRestrictAppDataToSystemVolume

`func (o *Windows10GeneralConfiguration) SetStorageRestrictAppDataToSystemVolume(v bool)`

SetStorageRestrictAppDataToSystemVolume gets a reference to the given bool and assigns it to the StorageRestrictAppDataToSystemVolume field.

### GetStorageRestrictAppInstallToSystemVolume

`func (o *Windows10GeneralConfiguration) GetStorageRestrictAppInstallToSystemVolume() bool`

GetStorageRestrictAppInstallToSystemVolume returns the StorageRestrictAppInstallToSystemVolume field if non-nil, zero value otherwise.

### GetStorageRestrictAppInstallToSystemVolumeOk

`func (o *Windows10GeneralConfiguration) GetStorageRestrictAppInstallToSystemVolumeOk() (bool, bool)`

GetStorageRestrictAppInstallToSystemVolumeOk returns a tuple with the StorageRestrictAppInstallToSystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRestrictAppInstallToSystemVolume

`func (o *Windows10GeneralConfiguration) HasStorageRestrictAppInstallToSystemVolume() bool`

HasStorageRestrictAppInstallToSystemVolume returns a boolean if a field has been set.

### SetStorageRestrictAppInstallToSystemVolume

`func (o *Windows10GeneralConfiguration) SetStorageRestrictAppInstallToSystemVolume(v bool)`

SetStorageRestrictAppInstallToSystemVolume gets a reference to the given bool and assigns it to the StorageRestrictAppInstallToSystemVolume field.

### GetGameDvrBlocked

`func (o *Windows10GeneralConfiguration) GetGameDvrBlocked() bool`

GetGameDvrBlocked returns the GameDvrBlocked field if non-nil, zero value otherwise.

### GetGameDvrBlockedOk

`func (o *Windows10GeneralConfiguration) GetGameDvrBlockedOk() (bool, bool)`

GetGameDvrBlockedOk returns a tuple with the GameDvrBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGameDvrBlocked

`func (o *Windows10GeneralConfiguration) HasGameDvrBlocked() bool`

HasGameDvrBlocked returns a boolean if a field has been set.

### SetGameDvrBlocked

`func (o *Windows10GeneralConfiguration) SetGameDvrBlocked(v bool)`

SetGameDvrBlocked gets a reference to the given bool and assigns it to the GameDvrBlocked field.

### GetExperienceBlockDeviceDiscovery

`func (o *Windows10GeneralConfiguration) GetExperienceBlockDeviceDiscovery() bool`

GetExperienceBlockDeviceDiscovery returns the ExperienceBlockDeviceDiscovery field if non-nil, zero value otherwise.

### GetExperienceBlockDeviceDiscoveryOk

`func (o *Windows10GeneralConfiguration) GetExperienceBlockDeviceDiscoveryOk() (bool, bool)`

GetExperienceBlockDeviceDiscoveryOk returns a tuple with the ExperienceBlockDeviceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperienceBlockDeviceDiscovery

`func (o *Windows10GeneralConfiguration) HasExperienceBlockDeviceDiscovery() bool`

HasExperienceBlockDeviceDiscovery returns a boolean if a field has been set.

### SetExperienceBlockDeviceDiscovery

`func (o *Windows10GeneralConfiguration) SetExperienceBlockDeviceDiscovery(v bool)`

SetExperienceBlockDeviceDiscovery gets a reference to the given bool and assigns it to the ExperienceBlockDeviceDiscovery field.

### GetExperienceBlockErrorDialogWhenNoSIM

`func (o *Windows10GeneralConfiguration) GetExperienceBlockErrorDialogWhenNoSIM() bool`

GetExperienceBlockErrorDialogWhenNoSIM returns the ExperienceBlockErrorDialogWhenNoSIM field if non-nil, zero value otherwise.

### GetExperienceBlockErrorDialogWhenNoSIMOk

`func (o *Windows10GeneralConfiguration) GetExperienceBlockErrorDialogWhenNoSIMOk() (bool, bool)`

GetExperienceBlockErrorDialogWhenNoSIMOk returns a tuple with the ExperienceBlockErrorDialogWhenNoSIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperienceBlockErrorDialogWhenNoSIM

`func (o *Windows10GeneralConfiguration) HasExperienceBlockErrorDialogWhenNoSIM() bool`

HasExperienceBlockErrorDialogWhenNoSIM returns a boolean if a field has been set.

### SetExperienceBlockErrorDialogWhenNoSIM

`func (o *Windows10GeneralConfiguration) SetExperienceBlockErrorDialogWhenNoSIM(v bool)`

SetExperienceBlockErrorDialogWhenNoSIM gets a reference to the given bool and assigns it to the ExperienceBlockErrorDialogWhenNoSIM field.

### GetExperienceBlockTaskSwitcher

`func (o *Windows10GeneralConfiguration) GetExperienceBlockTaskSwitcher() bool`

GetExperienceBlockTaskSwitcher returns the ExperienceBlockTaskSwitcher field if non-nil, zero value otherwise.

### GetExperienceBlockTaskSwitcherOk

`func (o *Windows10GeneralConfiguration) GetExperienceBlockTaskSwitcherOk() (bool, bool)`

GetExperienceBlockTaskSwitcherOk returns a tuple with the ExperienceBlockTaskSwitcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperienceBlockTaskSwitcher

`func (o *Windows10GeneralConfiguration) HasExperienceBlockTaskSwitcher() bool`

HasExperienceBlockTaskSwitcher returns a boolean if a field has been set.

### SetExperienceBlockTaskSwitcher

`func (o *Windows10GeneralConfiguration) SetExperienceBlockTaskSwitcher(v bool)`

SetExperienceBlockTaskSwitcher gets a reference to the given bool and assigns it to the ExperienceBlockTaskSwitcher field.

### GetLogonBlockFastUserSwitching

`func (o *Windows10GeneralConfiguration) GetLogonBlockFastUserSwitching() bool`

GetLogonBlockFastUserSwitching returns the LogonBlockFastUserSwitching field if non-nil, zero value otherwise.

### GetLogonBlockFastUserSwitchingOk

`func (o *Windows10GeneralConfiguration) GetLogonBlockFastUserSwitchingOk() (bool, bool)`

GetLogonBlockFastUserSwitchingOk returns a tuple with the LogonBlockFastUserSwitching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonBlockFastUserSwitching

`func (o *Windows10GeneralConfiguration) HasLogonBlockFastUserSwitching() bool`

HasLogonBlockFastUserSwitching returns a boolean if a field has been set.

### SetLogonBlockFastUserSwitching

`func (o *Windows10GeneralConfiguration) SetLogonBlockFastUserSwitching(v bool)`

SetLogonBlockFastUserSwitching gets a reference to the given bool and assigns it to the LogonBlockFastUserSwitching field.

### GetTenantLockdownRequireNetworkDuringOutOfBoxExperience

`func (o *Windows10GeneralConfiguration) GetTenantLockdownRequireNetworkDuringOutOfBoxExperience() bool`

GetTenantLockdownRequireNetworkDuringOutOfBoxExperience returns the TenantLockdownRequireNetworkDuringOutOfBoxExperience field if non-nil, zero value otherwise.

### GetTenantLockdownRequireNetworkDuringOutOfBoxExperienceOk

`func (o *Windows10GeneralConfiguration) GetTenantLockdownRequireNetworkDuringOutOfBoxExperienceOk() (bool, bool)`

GetTenantLockdownRequireNetworkDuringOutOfBoxExperienceOk returns a tuple with the TenantLockdownRequireNetworkDuringOutOfBoxExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTenantLockdownRequireNetworkDuringOutOfBoxExperience

`func (o *Windows10GeneralConfiguration) HasTenantLockdownRequireNetworkDuringOutOfBoxExperience() bool`

HasTenantLockdownRequireNetworkDuringOutOfBoxExperience returns a boolean if a field has been set.

### SetTenantLockdownRequireNetworkDuringOutOfBoxExperience

`func (o *Windows10GeneralConfiguration) SetTenantLockdownRequireNetworkDuringOutOfBoxExperience(v bool)`

SetTenantLockdownRequireNetworkDuringOutOfBoxExperience gets a reference to the given bool and assigns it to the TenantLockdownRequireNetworkDuringOutOfBoxExperience field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


