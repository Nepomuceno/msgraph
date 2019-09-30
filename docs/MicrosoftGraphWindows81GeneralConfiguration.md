# MicrosoftGraphWindows81GeneralConfiguration

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
**AccountsBlockAddingNonMicrosoftAccountEmail** | Pointer to **bool** | Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account. | [optional] 
**ApplyOnlyToWindows81** | Pointer to **bool** | Value indicating whether this policy only applies to Windows 8.1. This property is read-only. | [optional] 
**BrowserBlockAutofill** | Pointer to **bool** | Indicates whether or not to block auto fill. | [optional] 
**BrowserBlockAutomaticDetectionOfIntranetSites** | Pointer to **bool** | Indicates whether or not to block automatic detection of Intranet sites. | [optional] 
**BrowserBlockEnterpriseModeAccess** | Pointer to **bool** | Indicates whether or not to block enterprise mode access. | [optional] 
**BrowserBlockJavaScript** | Pointer to **bool** | Indicates whether or not to Block the user from using JavaScript. | [optional] 
**BrowserBlockPlugins** | Pointer to **bool** | Indicates whether or not to block plug-ins. | [optional] 
**BrowserBlockPopups** | Pointer to **bool** | Indicates whether or not to block popups. | [optional] 
**BrowserBlockSendingDoNotTrackHeader** | Pointer to **bool** | Indicates whether or not to Block the user from sending the do not track header. | [optional] 
**BrowserBlockSingleWordEntryOnIntranetSites** | Pointer to **bool** | Indicates whether or not to block a single word entry on Intranet sites. | [optional] 
**BrowserRequireSmartScreen** | Pointer to **bool** | Indicates whether or not to require the user to use the smart screen filter. | [optional] 
**BrowserEnterpriseModeSiteListLocation** | Pointer to **string** | The enterprise mode site list location. Could be a local file, local network or http location. | [optional] 
**BrowserInternetSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphInternetSiteSecurityLevel**](anyOf&lt;microsoft.graph.internetSiteSecurityLevel&gt;.md) | The internet security level. | [optional] 
**BrowserIntranetSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphSiteSecurityLevel**](anyOf&lt;microsoft.graph.siteSecurityLevel&gt;.md) | The Intranet security level. | [optional] 
**BrowserLoggingReportLocation** | Pointer to **string** | The logging report location. | [optional] 
**BrowserRequireHighSecurityForRestrictedSites** | Pointer to **bool** | Indicates whether or not to require high security for restricted sites. | [optional] 
**BrowserRequireFirewall** | Pointer to **bool** | Indicates whether or not to require a firewall. | [optional] 
**BrowserRequireFraudWarning** | Pointer to **bool** | Indicates whether or not to require fraud warning. | [optional] 
**BrowserTrustedSitesSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphSiteSecurityLevel**](anyOf&lt;microsoft.graph.siteSecurityLevel&gt;.md) | The trusted sites security level. | [optional] 
**CellularBlockDataRoaming** | Pointer to **bool** | Indicates whether or not to block data roaming. | [optional] 
**DiagnosticsBlockDataSubmission** | Pointer to **bool** | Indicates whether or not to block diagnostic data submission. | [optional] 
**PasswordBlockPicturePasswordAndPin** | Pointer to **bool** | Indicates whether or not to Block the user from using a pictures password and pin. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Password expiration in days. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | The minimum password length. | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | The minutes of inactivity before the screen times out. | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | The number of previous passwords to prevent re-use of. Valid values 0 to 24 | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required password type. | [optional] 
**PasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | The number of sign in failures before factory reset. | [optional] 
**StorageRequireDeviceEncryption** | Pointer to **bool** | Indicates whether or not to require encryption on a mobile device. | [optional] 
**UpdatesRequireAutomaticUpdates** | Pointer to **bool** | Indicates whether or not to require automatic updates. | [optional] 
**UserAccountControlSettings** | Pointer to [**AnyOfmicrosoftGraphWindowsUserAccountControlSettings**](anyOf&lt;microsoft.graph.windowsUserAccountControlSettings&gt;.md) | The user account control settings. | [optional] 
**WorkFoldersUrl** | Pointer to **string** | The work folders url. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmail() bool`

GetAccountsBlockAddingNonMicrosoftAccountEmail returns the AccountsBlockAddingNonMicrosoftAccountEmail field if non-nil, zero value otherwise.

### GetAccountsBlockAddingNonMicrosoftAccountEmailOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmailOk() (bool, bool)`

GetAccountsBlockAddingNonMicrosoftAccountEmailOk returns a tuple with the AccountsBlockAddingNonMicrosoftAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasAccountsBlockAddingNonMicrosoftAccountEmail() bool`

HasAccountsBlockAddingNonMicrosoftAccountEmail returns a boolean if a field has been set.

### SetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetAccountsBlockAddingNonMicrosoftAccountEmail(v bool)`

SetAccountsBlockAddingNonMicrosoftAccountEmail gets a reference to the given bool and assigns it to the AccountsBlockAddingNonMicrosoftAccountEmail field.

### GetApplyOnlyToWindows81

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetApplyOnlyToWindows81() bool`

GetApplyOnlyToWindows81 returns the ApplyOnlyToWindows81 field if non-nil, zero value otherwise.

### GetApplyOnlyToWindows81Ok

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetApplyOnlyToWindows81Ok() (bool, bool)`

GetApplyOnlyToWindows81Ok returns a tuple with the ApplyOnlyToWindows81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplyOnlyToWindows81

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasApplyOnlyToWindows81() bool`

HasApplyOnlyToWindows81 returns a boolean if a field has been set.

### SetApplyOnlyToWindows81

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetApplyOnlyToWindows81(v bool)`

SetApplyOnlyToWindows81 gets a reference to the given bool and assigns it to the ApplyOnlyToWindows81 field.

### GetBrowserBlockAutofill

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockAutofill() bool`

GetBrowserBlockAutofill returns the BrowserBlockAutofill field if non-nil, zero value otherwise.

### GetBrowserBlockAutofillOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockAutofillOk() (bool, bool)`

GetBrowserBlockAutofillOk returns a tuple with the BrowserBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockAutofill

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockAutofill() bool`

HasBrowserBlockAutofill returns a boolean if a field has been set.

### SetBrowserBlockAutofill

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockAutofill(v bool)`

SetBrowserBlockAutofill gets a reference to the given bool and assigns it to the BrowserBlockAutofill field.

### GetBrowserBlockAutomaticDetectionOfIntranetSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockAutomaticDetectionOfIntranetSites() bool`

GetBrowserBlockAutomaticDetectionOfIntranetSites returns the BrowserBlockAutomaticDetectionOfIntranetSites field if non-nil, zero value otherwise.

### GetBrowserBlockAutomaticDetectionOfIntranetSitesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockAutomaticDetectionOfIntranetSitesOk() (bool, bool)`

GetBrowserBlockAutomaticDetectionOfIntranetSitesOk returns a tuple with the BrowserBlockAutomaticDetectionOfIntranetSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockAutomaticDetectionOfIntranetSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockAutomaticDetectionOfIntranetSites() bool`

HasBrowserBlockAutomaticDetectionOfIntranetSites returns a boolean if a field has been set.

### SetBrowserBlockAutomaticDetectionOfIntranetSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockAutomaticDetectionOfIntranetSites(v bool)`

SetBrowserBlockAutomaticDetectionOfIntranetSites gets a reference to the given bool and assigns it to the BrowserBlockAutomaticDetectionOfIntranetSites field.

### GetBrowserBlockEnterpriseModeAccess

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockEnterpriseModeAccess() bool`

GetBrowserBlockEnterpriseModeAccess returns the BrowserBlockEnterpriseModeAccess field if non-nil, zero value otherwise.

### GetBrowserBlockEnterpriseModeAccessOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockEnterpriseModeAccessOk() (bool, bool)`

GetBrowserBlockEnterpriseModeAccessOk returns a tuple with the BrowserBlockEnterpriseModeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockEnterpriseModeAccess

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockEnterpriseModeAccess() bool`

HasBrowserBlockEnterpriseModeAccess returns a boolean if a field has been set.

### SetBrowserBlockEnterpriseModeAccess

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockEnterpriseModeAccess(v bool)`

SetBrowserBlockEnterpriseModeAccess gets a reference to the given bool and assigns it to the BrowserBlockEnterpriseModeAccess field.

### GetBrowserBlockJavaScript

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockJavaScript() bool`

GetBrowserBlockJavaScript returns the BrowserBlockJavaScript field if non-nil, zero value otherwise.

### GetBrowserBlockJavaScriptOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockJavaScriptOk() (bool, bool)`

GetBrowserBlockJavaScriptOk returns a tuple with the BrowserBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockJavaScript

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockJavaScript() bool`

HasBrowserBlockJavaScript returns a boolean if a field has been set.

### SetBrowserBlockJavaScript

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockJavaScript(v bool)`

SetBrowserBlockJavaScript gets a reference to the given bool and assigns it to the BrowserBlockJavaScript field.

### GetBrowserBlockPlugins

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockPlugins() bool`

GetBrowserBlockPlugins returns the BrowserBlockPlugins field if non-nil, zero value otherwise.

### GetBrowserBlockPluginsOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockPluginsOk() (bool, bool)`

GetBrowserBlockPluginsOk returns a tuple with the BrowserBlockPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockPlugins

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockPlugins() bool`

HasBrowserBlockPlugins returns a boolean if a field has been set.

### SetBrowserBlockPlugins

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockPlugins(v bool)`

SetBrowserBlockPlugins gets a reference to the given bool and assigns it to the BrowserBlockPlugins field.

### GetBrowserBlockPopups

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockPopups() bool`

GetBrowserBlockPopups returns the BrowserBlockPopups field if non-nil, zero value otherwise.

### GetBrowserBlockPopupsOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockPopupsOk() (bool, bool)`

GetBrowserBlockPopupsOk returns a tuple with the BrowserBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockPopups

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockPopups() bool`

HasBrowserBlockPopups returns a boolean if a field has been set.

### SetBrowserBlockPopups

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockPopups(v bool)`

SetBrowserBlockPopups gets a reference to the given bool and assigns it to the BrowserBlockPopups field.

### GetBrowserBlockSendingDoNotTrackHeader

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockSendingDoNotTrackHeader() bool`

GetBrowserBlockSendingDoNotTrackHeader returns the BrowserBlockSendingDoNotTrackHeader field if non-nil, zero value otherwise.

### GetBrowserBlockSendingDoNotTrackHeaderOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockSendingDoNotTrackHeaderOk() (bool, bool)`

GetBrowserBlockSendingDoNotTrackHeaderOk returns a tuple with the BrowserBlockSendingDoNotTrackHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockSendingDoNotTrackHeader

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockSendingDoNotTrackHeader() bool`

HasBrowserBlockSendingDoNotTrackHeader returns a boolean if a field has been set.

### SetBrowserBlockSendingDoNotTrackHeader

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockSendingDoNotTrackHeader(v bool)`

SetBrowserBlockSendingDoNotTrackHeader gets a reference to the given bool and assigns it to the BrowserBlockSendingDoNotTrackHeader field.

### GetBrowserBlockSingleWordEntryOnIntranetSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockSingleWordEntryOnIntranetSites() bool`

GetBrowserBlockSingleWordEntryOnIntranetSites returns the BrowserBlockSingleWordEntryOnIntranetSites field if non-nil, zero value otherwise.

### GetBrowserBlockSingleWordEntryOnIntranetSitesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserBlockSingleWordEntryOnIntranetSitesOk() (bool, bool)`

GetBrowserBlockSingleWordEntryOnIntranetSitesOk returns a tuple with the BrowserBlockSingleWordEntryOnIntranetSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockSingleWordEntryOnIntranetSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserBlockSingleWordEntryOnIntranetSites() bool`

HasBrowserBlockSingleWordEntryOnIntranetSites returns a boolean if a field has been set.

### SetBrowserBlockSingleWordEntryOnIntranetSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserBlockSingleWordEntryOnIntranetSites(v bool)`

SetBrowserBlockSingleWordEntryOnIntranetSites gets a reference to the given bool and assigns it to the BrowserBlockSingleWordEntryOnIntranetSites field.

### GetBrowserRequireSmartScreen

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireSmartScreen() bool`

GetBrowserRequireSmartScreen returns the BrowserRequireSmartScreen field if non-nil, zero value otherwise.

### GetBrowserRequireSmartScreenOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireSmartScreenOk() (bool, bool)`

GetBrowserRequireSmartScreenOk returns a tuple with the BrowserRequireSmartScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireSmartScreen

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserRequireSmartScreen() bool`

HasBrowserRequireSmartScreen returns a boolean if a field has been set.

### SetBrowserRequireSmartScreen

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserRequireSmartScreen(v bool)`

SetBrowserRequireSmartScreen gets a reference to the given bool and assigns it to the BrowserRequireSmartScreen field.

### GetBrowserEnterpriseModeSiteListLocation

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserEnterpriseModeSiteListLocation() string`

GetBrowserEnterpriseModeSiteListLocation returns the BrowserEnterpriseModeSiteListLocation field if non-nil, zero value otherwise.

### GetBrowserEnterpriseModeSiteListLocationOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserEnterpriseModeSiteListLocationOk() (string, bool)`

GetBrowserEnterpriseModeSiteListLocationOk returns a tuple with the BrowserEnterpriseModeSiteListLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserEnterpriseModeSiteListLocation

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserEnterpriseModeSiteListLocation() bool`

HasBrowserEnterpriseModeSiteListLocation returns a boolean if a field has been set.

### SetBrowserEnterpriseModeSiteListLocation

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserEnterpriseModeSiteListLocation(v string)`

SetBrowserEnterpriseModeSiteListLocation gets a reference to the given string and assigns it to the BrowserEnterpriseModeSiteListLocation field.

### SetBrowserEnterpriseModeSiteListLocationExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserEnterpriseModeSiteListLocationExplicitNull(b bool)`

SetBrowserEnterpriseModeSiteListLocationExplicitNull (un)sets BrowserEnterpriseModeSiteListLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BrowserEnterpriseModeSiteListLocation value is set to nil even if false is passed
### GetBrowserInternetSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserInternetSecurityLevel() AnyOfmicrosoftGraphInternetSiteSecurityLevel`

GetBrowserInternetSecurityLevel returns the BrowserInternetSecurityLevel field if non-nil, zero value otherwise.

### GetBrowserInternetSecurityLevelOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserInternetSecurityLevelOk() (AnyOfmicrosoftGraphInternetSiteSecurityLevel, bool)`

GetBrowserInternetSecurityLevelOk returns a tuple with the BrowserInternetSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserInternetSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserInternetSecurityLevel() bool`

HasBrowserInternetSecurityLevel returns a boolean if a field has been set.

### SetBrowserInternetSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserInternetSecurityLevel(v AnyOfmicrosoftGraphInternetSiteSecurityLevel)`

SetBrowserInternetSecurityLevel gets a reference to the given AnyOfmicrosoftGraphInternetSiteSecurityLevel and assigns it to the BrowserInternetSecurityLevel field.

### GetBrowserIntranetSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserIntranetSecurityLevel() AnyOfmicrosoftGraphSiteSecurityLevel`

GetBrowserIntranetSecurityLevel returns the BrowserIntranetSecurityLevel field if non-nil, zero value otherwise.

### GetBrowserIntranetSecurityLevelOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserIntranetSecurityLevelOk() (AnyOfmicrosoftGraphSiteSecurityLevel, bool)`

GetBrowserIntranetSecurityLevelOk returns a tuple with the BrowserIntranetSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserIntranetSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserIntranetSecurityLevel() bool`

HasBrowserIntranetSecurityLevel returns a boolean if a field has been set.

### SetBrowserIntranetSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserIntranetSecurityLevel(v AnyOfmicrosoftGraphSiteSecurityLevel)`

SetBrowserIntranetSecurityLevel gets a reference to the given AnyOfmicrosoftGraphSiteSecurityLevel and assigns it to the BrowserIntranetSecurityLevel field.

### GetBrowserLoggingReportLocation

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserLoggingReportLocation() string`

GetBrowserLoggingReportLocation returns the BrowserLoggingReportLocation field if non-nil, zero value otherwise.

### GetBrowserLoggingReportLocationOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserLoggingReportLocationOk() (string, bool)`

GetBrowserLoggingReportLocationOk returns a tuple with the BrowserLoggingReportLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserLoggingReportLocation

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserLoggingReportLocation() bool`

HasBrowserLoggingReportLocation returns a boolean if a field has been set.

### SetBrowserLoggingReportLocation

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserLoggingReportLocation(v string)`

SetBrowserLoggingReportLocation gets a reference to the given string and assigns it to the BrowserLoggingReportLocation field.

### SetBrowserLoggingReportLocationExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserLoggingReportLocationExplicitNull(b bool)`

SetBrowserLoggingReportLocationExplicitNull (un)sets BrowserLoggingReportLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BrowserLoggingReportLocation value is set to nil even if false is passed
### GetBrowserRequireHighSecurityForRestrictedSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireHighSecurityForRestrictedSites() bool`

GetBrowserRequireHighSecurityForRestrictedSites returns the BrowserRequireHighSecurityForRestrictedSites field if non-nil, zero value otherwise.

### GetBrowserRequireHighSecurityForRestrictedSitesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireHighSecurityForRestrictedSitesOk() (bool, bool)`

GetBrowserRequireHighSecurityForRestrictedSitesOk returns a tuple with the BrowserRequireHighSecurityForRestrictedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireHighSecurityForRestrictedSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserRequireHighSecurityForRestrictedSites() bool`

HasBrowserRequireHighSecurityForRestrictedSites returns a boolean if a field has been set.

### SetBrowserRequireHighSecurityForRestrictedSites

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserRequireHighSecurityForRestrictedSites(v bool)`

SetBrowserRequireHighSecurityForRestrictedSites gets a reference to the given bool and assigns it to the BrowserRequireHighSecurityForRestrictedSites field.

### GetBrowserRequireFirewall

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireFirewall() bool`

GetBrowserRequireFirewall returns the BrowserRequireFirewall field if non-nil, zero value otherwise.

### GetBrowserRequireFirewallOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireFirewallOk() (bool, bool)`

GetBrowserRequireFirewallOk returns a tuple with the BrowserRequireFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireFirewall

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserRequireFirewall() bool`

HasBrowserRequireFirewall returns a boolean if a field has been set.

### SetBrowserRequireFirewall

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserRequireFirewall(v bool)`

SetBrowserRequireFirewall gets a reference to the given bool and assigns it to the BrowserRequireFirewall field.

### GetBrowserRequireFraudWarning

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireFraudWarning() bool`

GetBrowserRequireFraudWarning returns the BrowserRequireFraudWarning field if non-nil, zero value otherwise.

### GetBrowserRequireFraudWarningOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserRequireFraudWarningOk() (bool, bool)`

GetBrowserRequireFraudWarningOk returns a tuple with the BrowserRequireFraudWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireFraudWarning

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserRequireFraudWarning() bool`

HasBrowserRequireFraudWarning returns a boolean if a field has been set.

### SetBrowserRequireFraudWarning

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserRequireFraudWarning(v bool)`

SetBrowserRequireFraudWarning gets a reference to the given bool and assigns it to the BrowserRequireFraudWarning field.

### GetBrowserTrustedSitesSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserTrustedSitesSecurityLevel() AnyOfmicrosoftGraphSiteSecurityLevel`

GetBrowserTrustedSitesSecurityLevel returns the BrowserTrustedSitesSecurityLevel field if non-nil, zero value otherwise.

### GetBrowserTrustedSitesSecurityLevelOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetBrowserTrustedSitesSecurityLevelOk() (AnyOfmicrosoftGraphSiteSecurityLevel, bool)`

GetBrowserTrustedSitesSecurityLevelOk returns a tuple with the BrowserTrustedSitesSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserTrustedSitesSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasBrowserTrustedSitesSecurityLevel() bool`

HasBrowserTrustedSitesSecurityLevel returns a boolean if a field has been set.

### SetBrowserTrustedSitesSecurityLevel

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetBrowserTrustedSitesSecurityLevel(v AnyOfmicrosoftGraphSiteSecurityLevel)`

SetBrowserTrustedSitesSecurityLevel gets a reference to the given AnyOfmicrosoftGraphSiteSecurityLevel and assigns it to the BrowserTrustedSitesSecurityLevel field.

### GetCellularBlockDataRoaming

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetCellularBlockDataRoaming() bool`

GetCellularBlockDataRoaming returns the CellularBlockDataRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataRoamingOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetCellularBlockDataRoamingOk() (bool, bool)`

GetCellularBlockDataRoamingOk returns a tuple with the CellularBlockDataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataRoaming

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasCellularBlockDataRoaming() bool`

HasCellularBlockDataRoaming returns a boolean if a field has been set.

### SetCellularBlockDataRoaming

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetCellularBlockDataRoaming(v bool)`

SetCellularBlockDataRoaming gets a reference to the given bool and assigns it to the CellularBlockDataRoaming field.

### GetDiagnosticsBlockDataSubmission

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDiagnosticsBlockDataSubmission() bool`

GetDiagnosticsBlockDataSubmission returns the DiagnosticsBlockDataSubmission field if non-nil, zero value otherwise.

### GetDiagnosticsBlockDataSubmissionOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetDiagnosticsBlockDataSubmissionOk() (bool, bool)`

GetDiagnosticsBlockDataSubmissionOk returns a tuple with the DiagnosticsBlockDataSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticsBlockDataSubmission

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasDiagnosticsBlockDataSubmission() bool`

HasDiagnosticsBlockDataSubmission returns a boolean if a field has been set.

### SetDiagnosticsBlockDataSubmission

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetDiagnosticsBlockDataSubmission(v bool)`

SetDiagnosticsBlockDataSubmission gets a reference to the given bool and assigns it to the DiagnosticsBlockDataSubmission field.

### GetPasswordBlockPicturePasswordAndPin

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordBlockPicturePasswordAndPin() bool`

GetPasswordBlockPicturePasswordAndPin returns the PasswordBlockPicturePasswordAndPin field if non-nil, zero value otherwise.

### GetPasswordBlockPicturePasswordAndPinOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordBlockPicturePasswordAndPinOk() (bool, bool)`

GetPasswordBlockPicturePasswordAndPinOk returns a tuple with the PasswordBlockPicturePasswordAndPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockPicturePasswordAndPin

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordBlockPicturePasswordAndPin() bool`

HasPasswordBlockPicturePasswordAndPin returns a boolean if a field has been set.

### SetPasswordBlockPicturePasswordAndPin

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordBlockPicturePasswordAndPin(v bool)`

SetPasswordBlockPicturePasswordAndPin gets a reference to the given bool and assigns it to the PasswordBlockPicturePasswordAndPin field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetStorageRequireDeviceEncryption

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetStorageRequireDeviceEncryption() bool`

GetStorageRequireDeviceEncryption returns the StorageRequireDeviceEncryption field if non-nil, zero value otherwise.

### GetStorageRequireDeviceEncryptionOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetStorageRequireDeviceEncryptionOk() (bool, bool)`

GetStorageRequireDeviceEncryptionOk returns a tuple with the StorageRequireDeviceEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireDeviceEncryption

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasStorageRequireDeviceEncryption() bool`

HasStorageRequireDeviceEncryption returns a boolean if a field has been set.

### SetStorageRequireDeviceEncryption

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetStorageRequireDeviceEncryption(v bool)`

SetStorageRequireDeviceEncryption gets a reference to the given bool and assigns it to the StorageRequireDeviceEncryption field.

### GetUpdatesRequireAutomaticUpdates

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUpdatesRequireAutomaticUpdates() bool`

GetUpdatesRequireAutomaticUpdates returns the UpdatesRequireAutomaticUpdates field if non-nil, zero value otherwise.

### GetUpdatesRequireAutomaticUpdatesOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUpdatesRequireAutomaticUpdatesOk() (bool, bool)`

GetUpdatesRequireAutomaticUpdatesOk returns a tuple with the UpdatesRequireAutomaticUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatesRequireAutomaticUpdates

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasUpdatesRequireAutomaticUpdates() bool`

HasUpdatesRequireAutomaticUpdates returns a boolean if a field has been set.

### SetUpdatesRequireAutomaticUpdates

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetUpdatesRequireAutomaticUpdates(v bool)`

SetUpdatesRequireAutomaticUpdates gets a reference to the given bool and assigns it to the UpdatesRequireAutomaticUpdates field.

### GetUserAccountControlSettings

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUserAccountControlSettings() AnyOfmicrosoftGraphWindowsUserAccountControlSettings`

GetUserAccountControlSettings returns the UserAccountControlSettings field if non-nil, zero value otherwise.

### GetUserAccountControlSettingsOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetUserAccountControlSettingsOk() (AnyOfmicrosoftGraphWindowsUserAccountControlSettings, bool)`

GetUserAccountControlSettingsOk returns a tuple with the UserAccountControlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserAccountControlSettings

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasUserAccountControlSettings() bool`

HasUserAccountControlSettings returns a boolean if a field has been set.

### SetUserAccountControlSettings

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetUserAccountControlSettings(v AnyOfmicrosoftGraphWindowsUserAccountControlSettings)`

SetUserAccountControlSettings gets a reference to the given AnyOfmicrosoftGraphWindowsUserAccountControlSettings and assigns it to the UserAccountControlSettings field.

### GetWorkFoldersUrl

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetWorkFoldersUrl() string`

GetWorkFoldersUrl returns the WorkFoldersUrl field if non-nil, zero value otherwise.

### GetWorkFoldersUrlOk

`func (o *MicrosoftGraphWindows81GeneralConfiguration) GetWorkFoldersUrlOk() (string, bool)`

GetWorkFoldersUrlOk returns a tuple with the WorkFoldersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkFoldersUrl

`func (o *MicrosoftGraphWindows81GeneralConfiguration) HasWorkFoldersUrl() bool`

HasWorkFoldersUrl returns a boolean if a field has been set.

### SetWorkFoldersUrl

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetWorkFoldersUrl(v string)`

SetWorkFoldersUrl gets a reference to the given string and assigns it to the WorkFoldersUrl field.

### SetWorkFoldersUrlExplicitNull

`func (o *MicrosoftGraphWindows81GeneralConfiguration) SetWorkFoldersUrlExplicitNull(b bool)`

SetWorkFoldersUrlExplicitNull (un)sets WorkFoldersUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkFoldersUrl value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


