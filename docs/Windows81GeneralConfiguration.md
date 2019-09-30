# Windows81GeneralConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *Windows81GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmail() bool`

GetAccountsBlockAddingNonMicrosoftAccountEmail returns the AccountsBlockAddingNonMicrosoftAccountEmail field if non-nil, zero value otherwise.

### GetAccountsBlockAddingNonMicrosoftAccountEmailOk

`func (o *Windows81GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmailOk() (bool, bool)`

GetAccountsBlockAddingNonMicrosoftAccountEmailOk returns a tuple with the AccountsBlockAddingNonMicrosoftAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *Windows81GeneralConfiguration) HasAccountsBlockAddingNonMicrosoftAccountEmail() bool`

HasAccountsBlockAddingNonMicrosoftAccountEmail returns a boolean if a field has been set.

### SetAccountsBlockAddingNonMicrosoftAccountEmail

`func (o *Windows81GeneralConfiguration) SetAccountsBlockAddingNonMicrosoftAccountEmail(v bool)`

SetAccountsBlockAddingNonMicrosoftAccountEmail gets a reference to the given bool and assigns it to the AccountsBlockAddingNonMicrosoftAccountEmail field.

### GetApplyOnlyToWindows81

`func (o *Windows81GeneralConfiguration) GetApplyOnlyToWindows81() bool`

GetApplyOnlyToWindows81 returns the ApplyOnlyToWindows81 field if non-nil, zero value otherwise.

### GetApplyOnlyToWindows81Ok

`func (o *Windows81GeneralConfiguration) GetApplyOnlyToWindows81Ok() (bool, bool)`

GetApplyOnlyToWindows81Ok returns a tuple with the ApplyOnlyToWindows81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplyOnlyToWindows81

`func (o *Windows81GeneralConfiguration) HasApplyOnlyToWindows81() bool`

HasApplyOnlyToWindows81 returns a boolean if a field has been set.

### SetApplyOnlyToWindows81

`func (o *Windows81GeneralConfiguration) SetApplyOnlyToWindows81(v bool)`

SetApplyOnlyToWindows81 gets a reference to the given bool and assigns it to the ApplyOnlyToWindows81 field.

### GetBrowserBlockAutofill

`func (o *Windows81GeneralConfiguration) GetBrowserBlockAutofill() bool`

GetBrowserBlockAutofill returns the BrowserBlockAutofill field if non-nil, zero value otherwise.

### GetBrowserBlockAutofillOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockAutofillOk() (bool, bool)`

GetBrowserBlockAutofillOk returns a tuple with the BrowserBlockAutofill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockAutofill

`func (o *Windows81GeneralConfiguration) HasBrowserBlockAutofill() bool`

HasBrowserBlockAutofill returns a boolean if a field has been set.

### SetBrowserBlockAutofill

`func (o *Windows81GeneralConfiguration) SetBrowserBlockAutofill(v bool)`

SetBrowserBlockAutofill gets a reference to the given bool and assigns it to the BrowserBlockAutofill field.

### GetBrowserBlockAutomaticDetectionOfIntranetSites

`func (o *Windows81GeneralConfiguration) GetBrowserBlockAutomaticDetectionOfIntranetSites() bool`

GetBrowserBlockAutomaticDetectionOfIntranetSites returns the BrowserBlockAutomaticDetectionOfIntranetSites field if non-nil, zero value otherwise.

### GetBrowserBlockAutomaticDetectionOfIntranetSitesOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockAutomaticDetectionOfIntranetSitesOk() (bool, bool)`

GetBrowserBlockAutomaticDetectionOfIntranetSitesOk returns a tuple with the BrowserBlockAutomaticDetectionOfIntranetSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockAutomaticDetectionOfIntranetSites

`func (o *Windows81GeneralConfiguration) HasBrowserBlockAutomaticDetectionOfIntranetSites() bool`

HasBrowserBlockAutomaticDetectionOfIntranetSites returns a boolean if a field has been set.

### SetBrowserBlockAutomaticDetectionOfIntranetSites

`func (o *Windows81GeneralConfiguration) SetBrowserBlockAutomaticDetectionOfIntranetSites(v bool)`

SetBrowserBlockAutomaticDetectionOfIntranetSites gets a reference to the given bool and assigns it to the BrowserBlockAutomaticDetectionOfIntranetSites field.

### GetBrowserBlockEnterpriseModeAccess

`func (o *Windows81GeneralConfiguration) GetBrowserBlockEnterpriseModeAccess() bool`

GetBrowserBlockEnterpriseModeAccess returns the BrowserBlockEnterpriseModeAccess field if non-nil, zero value otherwise.

### GetBrowserBlockEnterpriseModeAccessOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockEnterpriseModeAccessOk() (bool, bool)`

GetBrowserBlockEnterpriseModeAccessOk returns a tuple with the BrowserBlockEnterpriseModeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockEnterpriseModeAccess

`func (o *Windows81GeneralConfiguration) HasBrowserBlockEnterpriseModeAccess() bool`

HasBrowserBlockEnterpriseModeAccess returns a boolean if a field has been set.

### SetBrowserBlockEnterpriseModeAccess

`func (o *Windows81GeneralConfiguration) SetBrowserBlockEnterpriseModeAccess(v bool)`

SetBrowserBlockEnterpriseModeAccess gets a reference to the given bool and assigns it to the BrowserBlockEnterpriseModeAccess field.

### GetBrowserBlockJavaScript

`func (o *Windows81GeneralConfiguration) GetBrowserBlockJavaScript() bool`

GetBrowserBlockJavaScript returns the BrowserBlockJavaScript field if non-nil, zero value otherwise.

### GetBrowserBlockJavaScriptOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockJavaScriptOk() (bool, bool)`

GetBrowserBlockJavaScriptOk returns a tuple with the BrowserBlockJavaScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockJavaScript

`func (o *Windows81GeneralConfiguration) HasBrowserBlockJavaScript() bool`

HasBrowserBlockJavaScript returns a boolean if a field has been set.

### SetBrowserBlockJavaScript

`func (o *Windows81GeneralConfiguration) SetBrowserBlockJavaScript(v bool)`

SetBrowserBlockJavaScript gets a reference to the given bool and assigns it to the BrowserBlockJavaScript field.

### GetBrowserBlockPlugins

`func (o *Windows81GeneralConfiguration) GetBrowserBlockPlugins() bool`

GetBrowserBlockPlugins returns the BrowserBlockPlugins field if non-nil, zero value otherwise.

### GetBrowserBlockPluginsOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockPluginsOk() (bool, bool)`

GetBrowserBlockPluginsOk returns a tuple with the BrowserBlockPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockPlugins

`func (o *Windows81GeneralConfiguration) HasBrowserBlockPlugins() bool`

HasBrowserBlockPlugins returns a boolean if a field has been set.

### SetBrowserBlockPlugins

`func (o *Windows81GeneralConfiguration) SetBrowserBlockPlugins(v bool)`

SetBrowserBlockPlugins gets a reference to the given bool and assigns it to the BrowserBlockPlugins field.

### GetBrowserBlockPopups

`func (o *Windows81GeneralConfiguration) GetBrowserBlockPopups() bool`

GetBrowserBlockPopups returns the BrowserBlockPopups field if non-nil, zero value otherwise.

### GetBrowserBlockPopupsOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockPopupsOk() (bool, bool)`

GetBrowserBlockPopupsOk returns a tuple with the BrowserBlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockPopups

`func (o *Windows81GeneralConfiguration) HasBrowserBlockPopups() bool`

HasBrowserBlockPopups returns a boolean if a field has been set.

### SetBrowserBlockPopups

`func (o *Windows81GeneralConfiguration) SetBrowserBlockPopups(v bool)`

SetBrowserBlockPopups gets a reference to the given bool and assigns it to the BrowserBlockPopups field.

### GetBrowserBlockSendingDoNotTrackHeader

`func (o *Windows81GeneralConfiguration) GetBrowserBlockSendingDoNotTrackHeader() bool`

GetBrowserBlockSendingDoNotTrackHeader returns the BrowserBlockSendingDoNotTrackHeader field if non-nil, zero value otherwise.

### GetBrowserBlockSendingDoNotTrackHeaderOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockSendingDoNotTrackHeaderOk() (bool, bool)`

GetBrowserBlockSendingDoNotTrackHeaderOk returns a tuple with the BrowserBlockSendingDoNotTrackHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockSendingDoNotTrackHeader

`func (o *Windows81GeneralConfiguration) HasBrowserBlockSendingDoNotTrackHeader() bool`

HasBrowserBlockSendingDoNotTrackHeader returns a boolean if a field has been set.

### SetBrowserBlockSendingDoNotTrackHeader

`func (o *Windows81GeneralConfiguration) SetBrowserBlockSendingDoNotTrackHeader(v bool)`

SetBrowserBlockSendingDoNotTrackHeader gets a reference to the given bool and assigns it to the BrowserBlockSendingDoNotTrackHeader field.

### GetBrowserBlockSingleWordEntryOnIntranetSites

`func (o *Windows81GeneralConfiguration) GetBrowserBlockSingleWordEntryOnIntranetSites() bool`

GetBrowserBlockSingleWordEntryOnIntranetSites returns the BrowserBlockSingleWordEntryOnIntranetSites field if non-nil, zero value otherwise.

### GetBrowserBlockSingleWordEntryOnIntranetSitesOk

`func (o *Windows81GeneralConfiguration) GetBrowserBlockSingleWordEntryOnIntranetSitesOk() (bool, bool)`

GetBrowserBlockSingleWordEntryOnIntranetSitesOk returns a tuple with the BrowserBlockSingleWordEntryOnIntranetSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserBlockSingleWordEntryOnIntranetSites

`func (o *Windows81GeneralConfiguration) HasBrowserBlockSingleWordEntryOnIntranetSites() bool`

HasBrowserBlockSingleWordEntryOnIntranetSites returns a boolean if a field has been set.

### SetBrowserBlockSingleWordEntryOnIntranetSites

`func (o *Windows81GeneralConfiguration) SetBrowserBlockSingleWordEntryOnIntranetSites(v bool)`

SetBrowserBlockSingleWordEntryOnIntranetSites gets a reference to the given bool and assigns it to the BrowserBlockSingleWordEntryOnIntranetSites field.

### GetBrowserRequireSmartScreen

`func (o *Windows81GeneralConfiguration) GetBrowserRequireSmartScreen() bool`

GetBrowserRequireSmartScreen returns the BrowserRequireSmartScreen field if non-nil, zero value otherwise.

### GetBrowserRequireSmartScreenOk

`func (o *Windows81GeneralConfiguration) GetBrowserRequireSmartScreenOk() (bool, bool)`

GetBrowserRequireSmartScreenOk returns a tuple with the BrowserRequireSmartScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireSmartScreen

`func (o *Windows81GeneralConfiguration) HasBrowserRequireSmartScreen() bool`

HasBrowserRequireSmartScreen returns a boolean if a field has been set.

### SetBrowserRequireSmartScreen

`func (o *Windows81GeneralConfiguration) SetBrowserRequireSmartScreen(v bool)`

SetBrowserRequireSmartScreen gets a reference to the given bool and assigns it to the BrowserRequireSmartScreen field.

### GetBrowserEnterpriseModeSiteListLocation

`func (o *Windows81GeneralConfiguration) GetBrowserEnterpriseModeSiteListLocation() string`

GetBrowserEnterpriseModeSiteListLocation returns the BrowserEnterpriseModeSiteListLocation field if non-nil, zero value otherwise.

### GetBrowserEnterpriseModeSiteListLocationOk

`func (o *Windows81GeneralConfiguration) GetBrowserEnterpriseModeSiteListLocationOk() (string, bool)`

GetBrowserEnterpriseModeSiteListLocationOk returns a tuple with the BrowserEnterpriseModeSiteListLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserEnterpriseModeSiteListLocation

`func (o *Windows81GeneralConfiguration) HasBrowserEnterpriseModeSiteListLocation() bool`

HasBrowserEnterpriseModeSiteListLocation returns a boolean if a field has been set.

### SetBrowserEnterpriseModeSiteListLocation

`func (o *Windows81GeneralConfiguration) SetBrowserEnterpriseModeSiteListLocation(v string)`

SetBrowserEnterpriseModeSiteListLocation gets a reference to the given string and assigns it to the BrowserEnterpriseModeSiteListLocation field.

### SetBrowserEnterpriseModeSiteListLocationExplicitNull

`func (o *Windows81GeneralConfiguration) SetBrowserEnterpriseModeSiteListLocationExplicitNull(b bool)`

SetBrowserEnterpriseModeSiteListLocationExplicitNull (un)sets BrowserEnterpriseModeSiteListLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BrowserEnterpriseModeSiteListLocation value is set to nil even if false is passed
### GetBrowserInternetSecurityLevel

`func (o *Windows81GeneralConfiguration) GetBrowserInternetSecurityLevel() AnyOfmicrosoftGraphInternetSiteSecurityLevel`

GetBrowserInternetSecurityLevel returns the BrowserInternetSecurityLevel field if non-nil, zero value otherwise.

### GetBrowserInternetSecurityLevelOk

`func (o *Windows81GeneralConfiguration) GetBrowserInternetSecurityLevelOk() (AnyOfmicrosoftGraphInternetSiteSecurityLevel, bool)`

GetBrowserInternetSecurityLevelOk returns a tuple with the BrowserInternetSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserInternetSecurityLevel

`func (o *Windows81GeneralConfiguration) HasBrowserInternetSecurityLevel() bool`

HasBrowserInternetSecurityLevel returns a boolean if a field has been set.

### SetBrowserInternetSecurityLevel

`func (o *Windows81GeneralConfiguration) SetBrowserInternetSecurityLevel(v AnyOfmicrosoftGraphInternetSiteSecurityLevel)`

SetBrowserInternetSecurityLevel gets a reference to the given AnyOfmicrosoftGraphInternetSiteSecurityLevel and assigns it to the BrowserInternetSecurityLevel field.

### GetBrowserIntranetSecurityLevel

`func (o *Windows81GeneralConfiguration) GetBrowserIntranetSecurityLevel() AnyOfmicrosoftGraphSiteSecurityLevel`

GetBrowserIntranetSecurityLevel returns the BrowserIntranetSecurityLevel field if non-nil, zero value otherwise.

### GetBrowserIntranetSecurityLevelOk

`func (o *Windows81GeneralConfiguration) GetBrowserIntranetSecurityLevelOk() (AnyOfmicrosoftGraphSiteSecurityLevel, bool)`

GetBrowserIntranetSecurityLevelOk returns a tuple with the BrowserIntranetSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserIntranetSecurityLevel

`func (o *Windows81GeneralConfiguration) HasBrowserIntranetSecurityLevel() bool`

HasBrowserIntranetSecurityLevel returns a boolean if a field has been set.

### SetBrowserIntranetSecurityLevel

`func (o *Windows81GeneralConfiguration) SetBrowserIntranetSecurityLevel(v AnyOfmicrosoftGraphSiteSecurityLevel)`

SetBrowserIntranetSecurityLevel gets a reference to the given AnyOfmicrosoftGraphSiteSecurityLevel and assigns it to the BrowserIntranetSecurityLevel field.

### GetBrowserLoggingReportLocation

`func (o *Windows81GeneralConfiguration) GetBrowserLoggingReportLocation() string`

GetBrowserLoggingReportLocation returns the BrowserLoggingReportLocation field if non-nil, zero value otherwise.

### GetBrowserLoggingReportLocationOk

`func (o *Windows81GeneralConfiguration) GetBrowserLoggingReportLocationOk() (string, bool)`

GetBrowserLoggingReportLocationOk returns a tuple with the BrowserLoggingReportLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserLoggingReportLocation

`func (o *Windows81GeneralConfiguration) HasBrowserLoggingReportLocation() bool`

HasBrowserLoggingReportLocation returns a boolean if a field has been set.

### SetBrowserLoggingReportLocation

`func (o *Windows81GeneralConfiguration) SetBrowserLoggingReportLocation(v string)`

SetBrowserLoggingReportLocation gets a reference to the given string and assigns it to the BrowserLoggingReportLocation field.

### SetBrowserLoggingReportLocationExplicitNull

`func (o *Windows81GeneralConfiguration) SetBrowserLoggingReportLocationExplicitNull(b bool)`

SetBrowserLoggingReportLocationExplicitNull (un)sets BrowserLoggingReportLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BrowserLoggingReportLocation value is set to nil even if false is passed
### GetBrowserRequireHighSecurityForRestrictedSites

`func (o *Windows81GeneralConfiguration) GetBrowserRequireHighSecurityForRestrictedSites() bool`

GetBrowserRequireHighSecurityForRestrictedSites returns the BrowserRequireHighSecurityForRestrictedSites field if non-nil, zero value otherwise.

### GetBrowserRequireHighSecurityForRestrictedSitesOk

`func (o *Windows81GeneralConfiguration) GetBrowserRequireHighSecurityForRestrictedSitesOk() (bool, bool)`

GetBrowserRequireHighSecurityForRestrictedSitesOk returns a tuple with the BrowserRequireHighSecurityForRestrictedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireHighSecurityForRestrictedSites

`func (o *Windows81GeneralConfiguration) HasBrowserRequireHighSecurityForRestrictedSites() bool`

HasBrowserRequireHighSecurityForRestrictedSites returns a boolean if a field has been set.

### SetBrowserRequireHighSecurityForRestrictedSites

`func (o *Windows81GeneralConfiguration) SetBrowserRequireHighSecurityForRestrictedSites(v bool)`

SetBrowserRequireHighSecurityForRestrictedSites gets a reference to the given bool and assigns it to the BrowserRequireHighSecurityForRestrictedSites field.

### GetBrowserRequireFirewall

`func (o *Windows81GeneralConfiguration) GetBrowserRequireFirewall() bool`

GetBrowserRequireFirewall returns the BrowserRequireFirewall field if non-nil, zero value otherwise.

### GetBrowserRequireFirewallOk

`func (o *Windows81GeneralConfiguration) GetBrowserRequireFirewallOk() (bool, bool)`

GetBrowserRequireFirewallOk returns a tuple with the BrowserRequireFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireFirewall

`func (o *Windows81GeneralConfiguration) HasBrowserRequireFirewall() bool`

HasBrowserRequireFirewall returns a boolean if a field has been set.

### SetBrowserRequireFirewall

`func (o *Windows81GeneralConfiguration) SetBrowserRequireFirewall(v bool)`

SetBrowserRequireFirewall gets a reference to the given bool and assigns it to the BrowserRequireFirewall field.

### GetBrowserRequireFraudWarning

`func (o *Windows81GeneralConfiguration) GetBrowserRequireFraudWarning() bool`

GetBrowserRequireFraudWarning returns the BrowserRequireFraudWarning field if non-nil, zero value otherwise.

### GetBrowserRequireFraudWarningOk

`func (o *Windows81GeneralConfiguration) GetBrowserRequireFraudWarningOk() (bool, bool)`

GetBrowserRequireFraudWarningOk returns a tuple with the BrowserRequireFraudWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserRequireFraudWarning

`func (o *Windows81GeneralConfiguration) HasBrowserRequireFraudWarning() bool`

HasBrowserRequireFraudWarning returns a boolean if a field has been set.

### SetBrowserRequireFraudWarning

`func (o *Windows81GeneralConfiguration) SetBrowserRequireFraudWarning(v bool)`

SetBrowserRequireFraudWarning gets a reference to the given bool and assigns it to the BrowserRequireFraudWarning field.

### GetBrowserTrustedSitesSecurityLevel

`func (o *Windows81GeneralConfiguration) GetBrowserTrustedSitesSecurityLevel() AnyOfmicrosoftGraphSiteSecurityLevel`

GetBrowserTrustedSitesSecurityLevel returns the BrowserTrustedSitesSecurityLevel field if non-nil, zero value otherwise.

### GetBrowserTrustedSitesSecurityLevelOk

`func (o *Windows81GeneralConfiguration) GetBrowserTrustedSitesSecurityLevelOk() (AnyOfmicrosoftGraphSiteSecurityLevel, bool)`

GetBrowserTrustedSitesSecurityLevelOk returns a tuple with the BrowserTrustedSitesSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowserTrustedSitesSecurityLevel

`func (o *Windows81GeneralConfiguration) HasBrowserTrustedSitesSecurityLevel() bool`

HasBrowserTrustedSitesSecurityLevel returns a boolean if a field has been set.

### SetBrowserTrustedSitesSecurityLevel

`func (o *Windows81GeneralConfiguration) SetBrowserTrustedSitesSecurityLevel(v AnyOfmicrosoftGraphSiteSecurityLevel)`

SetBrowserTrustedSitesSecurityLevel gets a reference to the given AnyOfmicrosoftGraphSiteSecurityLevel and assigns it to the BrowserTrustedSitesSecurityLevel field.

### GetCellularBlockDataRoaming

`func (o *Windows81GeneralConfiguration) GetCellularBlockDataRoaming() bool`

GetCellularBlockDataRoaming returns the CellularBlockDataRoaming field if non-nil, zero value otherwise.

### GetCellularBlockDataRoamingOk

`func (o *Windows81GeneralConfiguration) GetCellularBlockDataRoamingOk() (bool, bool)`

GetCellularBlockDataRoamingOk returns a tuple with the CellularBlockDataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockDataRoaming

`func (o *Windows81GeneralConfiguration) HasCellularBlockDataRoaming() bool`

HasCellularBlockDataRoaming returns a boolean if a field has been set.

### SetCellularBlockDataRoaming

`func (o *Windows81GeneralConfiguration) SetCellularBlockDataRoaming(v bool)`

SetCellularBlockDataRoaming gets a reference to the given bool and assigns it to the CellularBlockDataRoaming field.

### GetDiagnosticsBlockDataSubmission

`func (o *Windows81GeneralConfiguration) GetDiagnosticsBlockDataSubmission() bool`

GetDiagnosticsBlockDataSubmission returns the DiagnosticsBlockDataSubmission field if non-nil, zero value otherwise.

### GetDiagnosticsBlockDataSubmissionOk

`func (o *Windows81GeneralConfiguration) GetDiagnosticsBlockDataSubmissionOk() (bool, bool)`

GetDiagnosticsBlockDataSubmissionOk returns a tuple with the DiagnosticsBlockDataSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticsBlockDataSubmission

`func (o *Windows81GeneralConfiguration) HasDiagnosticsBlockDataSubmission() bool`

HasDiagnosticsBlockDataSubmission returns a boolean if a field has been set.

### SetDiagnosticsBlockDataSubmission

`func (o *Windows81GeneralConfiguration) SetDiagnosticsBlockDataSubmission(v bool)`

SetDiagnosticsBlockDataSubmission gets a reference to the given bool and assigns it to the DiagnosticsBlockDataSubmission field.

### GetPasswordBlockPicturePasswordAndPin

`func (o *Windows81GeneralConfiguration) GetPasswordBlockPicturePasswordAndPin() bool`

GetPasswordBlockPicturePasswordAndPin returns the PasswordBlockPicturePasswordAndPin field if non-nil, zero value otherwise.

### GetPasswordBlockPicturePasswordAndPinOk

`func (o *Windows81GeneralConfiguration) GetPasswordBlockPicturePasswordAndPinOk() (bool, bool)`

GetPasswordBlockPicturePasswordAndPinOk returns a tuple with the PasswordBlockPicturePasswordAndPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockPicturePasswordAndPin

`func (o *Windows81GeneralConfiguration) HasPasswordBlockPicturePasswordAndPin() bool`

HasPasswordBlockPicturePasswordAndPin returns a boolean if a field has been set.

### SetPasswordBlockPicturePasswordAndPin

`func (o *Windows81GeneralConfiguration) SetPasswordBlockPicturePasswordAndPin(v bool)`

SetPasswordBlockPicturePasswordAndPin gets a reference to the given bool and assigns it to the PasswordBlockPicturePasswordAndPin field.

### GetPasswordExpirationDays

`func (o *Windows81GeneralConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *Windows81GeneralConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *Windows81GeneralConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *Windows81GeneralConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *Windows81GeneralConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *Windows81GeneralConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *Windows81GeneralConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *Windows81GeneralConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *Windows81GeneralConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *Windows81GeneralConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *Windows81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *Windows81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *Windows81GeneralConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *Windows81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *Windows81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *Windows81GeneralConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *Windows81GeneralConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *Windows81GeneralConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *Windows81GeneralConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *Windows81GeneralConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *Windows81GeneralConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *Windows81GeneralConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *Windows81GeneralConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *Windows81GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *Windows81GeneralConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *Windows81GeneralConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *Windows81GeneralConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *Windows81GeneralConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *Windows81GeneralConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *Windows81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *Windows81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *Windows81GeneralConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *Windows81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *Windows81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetStorageRequireDeviceEncryption

`func (o *Windows81GeneralConfiguration) GetStorageRequireDeviceEncryption() bool`

GetStorageRequireDeviceEncryption returns the StorageRequireDeviceEncryption field if non-nil, zero value otherwise.

### GetStorageRequireDeviceEncryptionOk

`func (o *Windows81GeneralConfiguration) GetStorageRequireDeviceEncryptionOk() (bool, bool)`

GetStorageRequireDeviceEncryptionOk returns a tuple with the StorageRequireDeviceEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireDeviceEncryption

`func (o *Windows81GeneralConfiguration) HasStorageRequireDeviceEncryption() bool`

HasStorageRequireDeviceEncryption returns a boolean if a field has been set.

### SetStorageRequireDeviceEncryption

`func (o *Windows81GeneralConfiguration) SetStorageRequireDeviceEncryption(v bool)`

SetStorageRequireDeviceEncryption gets a reference to the given bool and assigns it to the StorageRequireDeviceEncryption field.

### GetUpdatesRequireAutomaticUpdates

`func (o *Windows81GeneralConfiguration) GetUpdatesRequireAutomaticUpdates() bool`

GetUpdatesRequireAutomaticUpdates returns the UpdatesRequireAutomaticUpdates field if non-nil, zero value otherwise.

### GetUpdatesRequireAutomaticUpdatesOk

`func (o *Windows81GeneralConfiguration) GetUpdatesRequireAutomaticUpdatesOk() (bool, bool)`

GetUpdatesRequireAutomaticUpdatesOk returns a tuple with the UpdatesRequireAutomaticUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatesRequireAutomaticUpdates

`func (o *Windows81GeneralConfiguration) HasUpdatesRequireAutomaticUpdates() bool`

HasUpdatesRequireAutomaticUpdates returns a boolean if a field has been set.

### SetUpdatesRequireAutomaticUpdates

`func (o *Windows81GeneralConfiguration) SetUpdatesRequireAutomaticUpdates(v bool)`

SetUpdatesRequireAutomaticUpdates gets a reference to the given bool and assigns it to the UpdatesRequireAutomaticUpdates field.

### GetUserAccountControlSettings

`func (o *Windows81GeneralConfiguration) GetUserAccountControlSettings() AnyOfmicrosoftGraphWindowsUserAccountControlSettings`

GetUserAccountControlSettings returns the UserAccountControlSettings field if non-nil, zero value otherwise.

### GetUserAccountControlSettingsOk

`func (o *Windows81GeneralConfiguration) GetUserAccountControlSettingsOk() (AnyOfmicrosoftGraphWindowsUserAccountControlSettings, bool)`

GetUserAccountControlSettingsOk returns a tuple with the UserAccountControlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserAccountControlSettings

`func (o *Windows81GeneralConfiguration) HasUserAccountControlSettings() bool`

HasUserAccountControlSettings returns a boolean if a field has been set.

### SetUserAccountControlSettings

`func (o *Windows81GeneralConfiguration) SetUserAccountControlSettings(v AnyOfmicrosoftGraphWindowsUserAccountControlSettings)`

SetUserAccountControlSettings gets a reference to the given AnyOfmicrosoftGraphWindowsUserAccountControlSettings and assigns it to the UserAccountControlSettings field.

### GetWorkFoldersUrl

`func (o *Windows81GeneralConfiguration) GetWorkFoldersUrl() string`

GetWorkFoldersUrl returns the WorkFoldersUrl field if non-nil, zero value otherwise.

### GetWorkFoldersUrlOk

`func (o *Windows81GeneralConfiguration) GetWorkFoldersUrlOk() (string, bool)`

GetWorkFoldersUrlOk returns a tuple with the WorkFoldersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkFoldersUrl

`func (o *Windows81GeneralConfiguration) HasWorkFoldersUrl() bool`

HasWorkFoldersUrl returns a boolean if a field has been set.

### SetWorkFoldersUrl

`func (o *Windows81GeneralConfiguration) SetWorkFoldersUrl(v string)`

SetWorkFoldersUrl gets a reference to the given string and assigns it to the WorkFoldersUrl field.

### SetWorkFoldersUrlExplicitNull

`func (o *Windows81GeneralConfiguration) SetWorkFoldersUrlExplicitNull(b bool)`

SetWorkFoldersUrlExplicitNull (un)sets WorkFoldersUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkFoldersUrl value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


