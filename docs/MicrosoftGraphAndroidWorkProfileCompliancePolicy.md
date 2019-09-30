# MicrosoftGraphAndroidWorkProfileCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | Pointer to **string** | Admin provided description of the Device Configuration. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**ScheduledActionsForRule** | Pointer to [**[]MicrosoftGraphDeviceComplianceScheduledActionForRule**](microsoft.graph.deviceComplianceScheduledActionForRule.md) |  | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphDeviceComplianceDeviceStatus**](microsoft.graph.deviceComplianceDeviceStatus.md) |  | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphDeviceComplianceUserStatus**](microsoft.graph.deviceComplianceUserStatus.md) |  | [optional] 
**DeviceStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceDeviceOverview**](anyOf&lt;microsoft.graph.deviceComplianceDeviceOverview&gt;.md) |  | [optional] 
**UserStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceUserOverview**](anyOf&lt;microsoft.graph.deviceComplianceUserOverview&gt;.md) |  | [optional] 
**DeviceSettingStateSummaries** | Pointer to [**[]MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md) |  | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicyAssignment**](microsoft.graph.deviceCompliancePolicyAssignment.md) |  | [optional] 
**PasswordRequired** | Pointer to **bool** | Require a password to unlock device. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum password length. Valid values 4 to 16 | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidRequiredPasswordType**](anyOf&lt;microsoft.graph.androidRequiredPasswordType&gt;.md) | Type of characters in password | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a password is required. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. Valid values 1 to 365 | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 1 to 24 | [optional] 
**SecurityPreventInstallAppsFromUnknownSources** | Pointer to **bool** | Require that devices disallow installation of apps from unknown sources. | [optional] 
**SecurityDisableUsbDebugging** | Pointer to **bool** | Disable USB debugging on Android devices. | [optional] 
**SecurityRequireVerifyApps** | Pointer to **bool** | Require the Android Verify apps feature is turned on. | [optional] 
**DeviceThreatProtectionEnabled** | Pointer to **bool** | Require that devices have enabled device threat protection. | [optional] 
**DeviceThreatProtectionRequiredSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphDeviceThreatProtectionLevel**](anyOf&lt;microsoft.graph.deviceThreatProtectionLevel&gt;.md) | Require Mobile Threat Protection minimum risk level to report noncompliance. | [optional] 
**SecurityBlockJailbrokenDevices** | Pointer to **bool** | Devices must not be jailbroken or rooted. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum Android version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum Android version. | [optional] 
**MinAndroidSecurityPatchLevel** | Pointer to **string** | Minimum Android security patch level. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Require encryption on Android devices. | [optional] 
**SecurityRequireSafetyNetAttestationBasicIntegrity** | Pointer to **bool** | Require the device to pass the SafetyNet basic integrity check. | [optional] 
**SecurityRequireSafetyNetAttestationCertifiedDevice** | Pointer to **bool** | Require the device to pass the SafetyNet certified device check. | [optional] 
**SecurityRequireGooglePlayServices** | Pointer to **bool** | Require Google Play Services to be installed and enabled on the device. | [optional] 
**SecurityRequireUpToDateSecurityProviders** | Pointer to **bool** | Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date. | [optional] 
**SecurityRequireCompanyPortalAppIntegrity** | Pointer to **bool** | Require the device to pass the Company Portal client app runtime integrity check. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetScheduledActionsForRule

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetScheduledActionsForRule() []MicrosoftGraphDeviceComplianceScheduledActionForRule`

GetScheduledActionsForRule returns the ScheduledActionsForRule field if non-nil, zero value otherwise.

### GetScheduledActionsForRuleOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetScheduledActionsForRuleOk() ([]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool)`

GetScheduledActionsForRuleOk returns a tuple with the ScheduledActionsForRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionsForRule

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasScheduledActionsForRule() bool`

HasScheduledActionsForRule returns a boolean if a field has been set.

### SetScheduledActionsForRule

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetScheduledActionsForRule(v []MicrosoftGraphDeviceComplianceScheduledActionForRule)`

SetScheduledActionsForRule gets a reference to the given []MicrosoftGraphDeviceComplianceScheduledActionForRule and assigns it to the ScheduledActionsForRule field.

### GetDeviceStatuses

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceStatuses() []MicrosoftGraphDeviceComplianceDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceComplianceDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDeviceStatuses(v []MicrosoftGraphDeviceComplianceDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetUserStatuses() []MicrosoftGraphDeviceComplianceUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetUserStatusesOk() ([]MicrosoftGraphDeviceComplianceUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetUserStatuses(v []MicrosoftGraphDeviceComplianceUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceComplianceDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceComplianceUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAssignments

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetAssignments() []MicrosoftGraphDeviceCompliancePolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetAssignmentsOk() ([]MicrosoftGraphDeviceCompliancePolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetAssignments(v []MicrosoftGraphDeviceCompliancePolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyAssignment and assigns it to the Assignments field.

### GetPasswordRequired

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphAndroidRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphAndroidRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetSecurityPreventInstallAppsFromUnknownSources

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSources() bool`

GetSecurityPreventInstallAppsFromUnknownSources returns the SecurityPreventInstallAppsFromUnknownSources field if non-nil, zero value otherwise.

### GetSecurityPreventInstallAppsFromUnknownSourcesOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSourcesOk() (bool, bool)`

GetSecurityPreventInstallAppsFromUnknownSourcesOk returns a tuple with the SecurityPreventInstallAppsFromUnknownSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityPreventInstallAppsFromUnknownSources

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityPreventInstallAppsFromUnknownSources() bool`

HasSecurityPreventInstallAppsFromUnknownSources returns a boolean if a field has been set.

### SetSecurityPreventInstallAppsFromUnknownSources

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityPreventInstallAppsFromUnknownSources(v bool)`

SetSecurityPreventInstallAppsFromUnknownSources gets a reference to the given bool and assigns it to the SecurityPreventInstallAppsFromUnknownSources field.

### GetSecurityDisableUsbDebugging

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityDisableUsbDebugging() bool`

GetSecurityDisableUsbDebugging returns the SecurityDisableUsbDebugging field if non-nil, zero value otherwise.

### GetSecurityDisableUsbDebuggingOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityDisableUsbDebuggingOk() (bool, bool)`

GetSecurityDisableUsbDebuggingOk returns a tuple with the SecurityDisableUsbDebugging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityDisableUsbDebugging

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityDisableUsbDebugging() bool`

HasSecurityDisableUsbDebugging returns a boolean if a field has been set.

### SetSecurityDisableUsbDebugging

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityDisableUsbDebugging(v bool)`

SetSecurityDisableUsbDebugging gets a reference to the given bool and assigns it to the SecurityDisableUsbDebugging field.

### GetSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireVerifyApps() bool`

GetSecurityRequireVerifyApps returns the SecurityRequireVerifyApps field if non-nil, zero value otherwise.

### GetSecurityRequireVerifyAppsOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireVerifyAppsOk() (bool, bool)`

GetSecurityRequireVerifyAppsOk returns a tuple with the SecurityRequireVerifyApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityRequireVerifyApps() bool`

HasSecurityRequireVerifyApps returns a boolean if a field has been set.

### SetSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityRequireVerifyApps(v bool)`

SetSecurityRequireVerifyApps gets a reference to the given bool and assigns it to the SecurityRequireVerifyApps field.

### GetDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceThreatProtectionEnabled() bool`

GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionEnabledOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool)`

GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDeviceThreatProtectionEnabled() bool`

HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.

### SetDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool)`

SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.

### GetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel`

GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionRequiredSecurityLevelOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool)`

GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool`

HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.

### SetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel)`

SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.

### GetSecurityBlockJailbrokenDevices

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityBlockJailbrokenDevices() bool`

GetSecurityBlockJailbrokenDevices returns the SecurityBlockJailbrokenDevices field if non-nil, zero value otherwise.

### GetSecurityBlockJailbrokenDevicesOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityBlockJailbrokenDevicesOk() (bool, bool)`

GetSecurityBlockJailbrokenDevicesOk returns a tuple with the SecurityBlockJailbrokenDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityBlockJailbrokenDevices

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityBlockJailbrokenDevices() bool`

HasSecurityBlockJailbrokenDevices returns a boolean if a field has been set.

### SetSecurityBlockJailbrokenDevices

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityBlockJailbrokenDevices(v bool)`

SetSecurityBlockJailbrokenDevices gets a reference to the given bool and assigns it to the SecurityBlockJailbrokenDevices field.

### GetOsMinimumVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetMinAndroidSecurityPatchLevel

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetMinAndroidSecurityPatchLevel() string`

GetMinAndroidSecurityPatchLevel returns the MinAndroidSecurityPatchLevel field if non-nil, zero value otherwise.

### GetMinAndroidSecurityPatchLevelOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetMinAndroidSecurityPatchLevelOk() (string, bool)`

GetMinAndroidSecurityPatchLevelOk returns a tuple with the MinAndroidSecurityPatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinAndroidSecurityPatchLevel

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasMinAndroidSecurityPatchLevel() bool`

HasMinAndroidSecurityPatchLevel returns a boolean if a field has been set.

### SetMinAndroidSecurityPatchLevel

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetMinAndroidSecurityPatchLevel(v string)`

SetMinAndroidSecurityPatchLevel gets a reference to the given string and assigns it to the MinAndroidSecurityPatchLevel field.

### SetMinAndroidSecurityPatchLevelExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetMinAndroidSecurityPatchLevelExplicitNull(b bool)`

SetMinAndroidSecurityPatchLevelExplicitNull (un)sets MinAndroidSecurityPatchLevel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinAndroidSecurityPatchLevel value is set to nil even if false is passed
### GetStorageRequireEncryption

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.

### GetSecurityRequireSafetyNetAttestationBasicIntegrity

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity() bool`

GetSecurityRequireSafetyNetAttestationBasicIntegrity returns the SecurityRequireSafetyNetAttestationBasicIntegrity field if non-nil, zero value otherwise.

### GetSecurityRequireSafetyNetAttestationBasicIntegrityOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrityOk() (bool, bool)`

GetSecurityRequireSafetyNetAttestationBasicIntegrityOk returns a tuple with the SecurityRequireSafetyNetAttestationBasicIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireSafetyNetAttestationBasicIntegrity

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityRequireSafetyNetAttestationBasicIntegrity() bool`

HasSecurityRequireSafetyNetAttestationBasicIntegrity returns a boolean if a field has been set.

### SetSecurityRequireSafetyNetAttestationBasicIntegrity

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(v bool)`

SetSecurityRequireSafetyNetAttestationBasicIntegrity gets a reference to the given bool and assigns it to the SecurityRequireSafetyNetAttestationBasicIntegrity field.

### GetSecurityRequireSafetyNetAttestationCertifiedDevice

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice() bool`

GetSecurityRequireSafetyNetAttestationCertifiedDevice returns the SecurityRequireSafetyNetAttestationCertifiedDevice field if non-nil, zero value otherwise.

### GetSecurityRequireSafetyNetAttestationCertifiedDeviceOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDeviceOk() (bool, bool)`

GetSecurityRequireSafetyNetAttestationCertifiedDeviceOk returns a tuple with the SecurityRequireSafetyNetAttestationCertifiedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireSafetyNetAttestationCertifiedDevice

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityRequireSafetyNetAttestationCertifiedDevice() bool`

HasSecurityRequireSafetyNetAttestationCertifiedDevice returns a boolean if a field has been set.

### SetSecurityRequireSafetyNetAttestationCertifiedDevice

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(v bool)`

SetSecurityRequireSafetyNetAttestationCertifiedDevice gets a reference to the given bool and assigns it to the SecurityRequireSafetyNetAttestationCertifiedDevice field.

### GetSecurityRequireGooglePlayServices

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireGooglePlayServices() bool`

GetSecurityRequireGooglePlayServices returns the SecurityRequireGooglePlayServices field if non-nil, zero value otherwise.

### GetSecurityRequireGooglePlayServicesOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireGooglePlayServicesOk() (bool, bool)`

GetSecurityRequireGooglePlayServicesOk returns a tuple with the SecurityRequireGooglePlayServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireGooglePlayServices

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityRequireGooglePlayServices() bool`

HasSecurityRequireGooglePlayServices returns a boolean if a field has been set.

### SetSecurityRequireGooglePlayServices

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityRequireGooglePlayServices(v bool)`

SetSecurityRequireGooglePlayServices gets a reference to the given bool and assigns it to the SecurityRequireGooglePlayServices field.

### GetSecurityRequireUpToDateSecurityProviders

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireUpToDateSecurityProviders() bool`

GetSecurityRequireUpToDateSecurityProviders returns the SecurityRequireUpToDateSecurityProviders field if non-nil, zero value otherwise.

### GetSecurityRequireUpToDateSecurityProvidersOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireUpToDateSecurityProvidersOk() (bool, bool)`

GetSecurityRequireUpToDateSecurityProvidersOk returns a tuple with the SecurityRequireUpToDateSecurityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireUpToDateSecurityProviders

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityRequireUpToDateSecurityProviders() bool`

HasSecurityRequireUpToDateSecurityProviders returns a boolean if a field has been set.

### SetSecurityRequireUpToDateSecurityProviders

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityRequireUpToDateSecurityProviders(v bool)`

SetSecurityRequireUpToDateSecurityProviders gets a reference to the given bool and assigns it to the SecurityRequireUpToDateSecurityProviders field.

### GetSecurityRequireCompanyPortalAppIntegrity

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrity() bool`

GetSecurityRequireCompanyPortalAppIntegrity returns the SecurityRequireCompanyPortalAppIntegrity field if non-nil, zero value otherwise.

### GetSecurityRequireCompanyPortalAppIntegrityOk

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrityOk() (bool, bool)`

GetSecurityRequireCompanyPortalAppIntegrityOk returns a tuple with the SecurityRequireCompanyPortalAppIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireCompanyPortalAppIntegrity

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) HasSecurityRequireCompanyPortalAppIntegrity() bool`

HasSecurityRequireCompanyPortalAppIntegrity returns a boolean if a field has been set.

### SetSecurityRequireCompanyPortalAppIntegrity

`func (o *MicrosoftGraphAndroidWorkProfileCompliancePolicy) SetSecurityRequireCompanyPortalAppIntegrity(v bool)`

SetSecurityRequireCompanyPortalAppIntegrity gets a reference to the given bool and assigns it to the SecurityRequireCompanyPortalAppIntegrity field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


