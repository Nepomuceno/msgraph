# MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration

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
**PasswordBlockFingerprintUnlock** | Pointer to **bool** | Indicates whether or not to block fingerprint unlock. | [optional] 
**PasswordBlockTrustAgents** | Pointer to **bool** | Indicates whether or not to block Smart Lock and other trust agents. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. Valid values 1 to 365 | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of passwords. Valid values 4 to 16 | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before the screen times out. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 0 to 24 | [optional] 
**PasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | Number of sign in failures allowed before factory reset. Valid values 1 to 16 | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType**](anyOf&lt;microsoft.graph.androidWorkProfileRequiredPasswordType&gt;.md) | Type of password that is required. | [optional] 
**WorkProfileDataSharingType** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType**](anyOf&lt;microsoft.graph.androidWorkProfileCrossProfileDataSharingType&gt;.md) | Type of data sharing that is allowed. | [optional] 
**WorkProfileBlockNotificationsWhileDeviceLocked** | Pointer to **bool** | Indicates whether or not to block notifications while device locked. | [optional] 
**WorkProfileBlockAddingAccounts** | Pointer to **bool** | Block users from adding/removing accounts in work profile. | [optional] 
**WorkProfileBluetoothEnableContactSharing** | Pointer to **bool** | Allow bluetooth devices to access enterprise contacts. | [optional] 
**WorkProfileBlockScreenCapture** | Pointer to **bool** | Block screen capture in work profile. | [optional] 
**WorkProfileBlockCrossProfileCallerId** | Pointer to **bool** | Block display work profile caller ID in personal profile. | [optional] 
**WorkProfileBlockCamera** | Pointer to **bool** | Block work profile camera. | [optional] 
**WorkProfileBlockCrossProfileContactsSearch** | Pointer to **bool** | Block work profile contacts availability in personal profile. | [optional] 
**WorkProfileBlockCrossProfileCopyPaste** | Pointer to **bool** | Boolean that indicates if the setting disallow cross profile copy/paste is enabled. | [optional] 
**WorkProfileDefaultAppPermissionPolicy** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType**](anyOf&lt;microsoft.graph.androidWorkProfileDefaultAppPermissionPolicyType&gt;.md) | Type of password that is required. | [optional] 
**WorkProfilePasswordBlockFingerprintUnlock** | Pointer to **bool** | Indicates whether or not to block fingerprint unlock for work profile. | [optional] 
**WorkProfilePasswordBlockTrustAgents** | Pointer to **bool** | Indicates whether or not to block Smart Lock and other trust agents for work profile. | [optional] 
**WorkProfilePasswordExpirationDays** | Pointer to **int32** | Number of days before the work profile password expires. Valid values 1 to 365 | [optional] 
**WorkProfilePasswordMinimumLength** | Pointer to **int32** | Minimum length of work profile password. Valid values 4 to 16 | [optional] 
**WorkProfilePasswordMinNumericCharacters** | Pointer to **int32** | Minimum # of numeric characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinNonLetterCharacters** | Pointer to **int32** | Minimum # of non-letter characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinLetterCharacters** | Pointer to **int32** | Minimum # of letter characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinLowerCaseCharacters** | Pointer to **int32** | Minimum # of lower-case characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinUpperCaseCharacters** | Pointer to **int32** | Minimum # of upper-case characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinSymbolCharacters** | Pointer to **int32** | Minimum # of symbols required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before the screen times out. | [optional] 
**WorkProfilePasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous work profile passwords to block. Valid values 0 to 24 | [optional] 
**WorkProfilePasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16 | [optional] 
**WorkProfilePasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType**](anyOf&lt;microsoft.graph.androidWorkProfileRequiredPasswordType&gt;.md) | Type of work profile password that is required. | [optional] 
**WorkProfileRequirePassword** | Pointer to **bool** | Password is required or not for work profile | [optional] 
**SecurityRequireVerifyApps** | Pointer to **bool** | Require the Android Verify apps feature is turned on. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetPasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock() bool`

GetPasswordBlockFingerprintUnlock returns the PasswordBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetPasswordBlockFingerprintUnlockOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlockOk() (bool, bool)`

GetPasswordBlockFingerprintUnlockOk returns a tuple with the PasswordBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordBlockFingerprintUnlock() bool`

HasPasswordBlockFingerprintUnlock returns a boolean if a field has been set.

### SetPasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(v bool)`

SetPasswordBlockFingerprintUnlock gets a reference to the given bool and assigns it to the PasswordBlockFingerprintUnlock field.

### GetPasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockTrustAgents() bool`

GetPasswordBlockTrustAgents returns the PasswordBlockTrustAgents field if non-nil, zero value otherwise.

### GetPasswordBlockTrustAgentsOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockTrustAgentsOk() (bool, bool)`

GetPasswordBlockTrustAgentsOk returns a tuple with the PasswordBlockTrustAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordBlockTrustAgents() bool`

HasPasswordBlockTrustAgents returns a boolean if a field has been set.

### SetPasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(v bool)`

SetPasswordBlockTrustAgents gets a reference to the given bool and assigns it to the PasswordBlockTrustAgents field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetWorkProfileDataSharingType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDataSharingType() AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType`

GetWorkProfileDataSharingType returns the WorkProfileDataSharingType field if non-nil, zero value otherwise.

### GetWorkProfileDataSharingTypeOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDataSharingTypeOk() (AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType, bool)`

GetWorkProfileDataSharingTypeOk returns a tuple with the WorkProfileDataSharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileDataSharingType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileDataSharingType() bool`

HasWorkProfileDataSharingType returns a boolean if a field has been set.

### SetWorkProfileDataSharingType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDataSharingType(v AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType)`

SetWorkProfileDataSharingType gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType and assigns it to the WorkProfileDataSharingType field.

### GetWorkProfileBlockNotificationsWhileDeviceLocked

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLocked() bool`

GetWorkProfileBlockNotificationsWhileDeviceLocked returns the WorkProfileBlockNotificationsWhileDeviceLocked field if non-nil, zero value otherwise.

### GetWorkProfileBlockNotificationsWhileDeviceLockedOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLockedOk() (bool, bool)`

GetWorkProfileBlockNotificationsWhileDeviceLockedOk returns a tuple with the WorkProfileBlockNotificationsWhileDeviceLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockNotificationsWhileDeviceLocked

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockNotificationsWhileDeviceLocked() bool`

HasWorkProfileBlockNotificationsWhileDeviceLocked returns a boolean if a field has been set.

### SetWorkProfileBlockNotificationsWhileDeviceLocked

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockNotificationsWhileDeviceLocked(v bool)`

SetWorkProfileBlockNotificationsWhileDeviceLocked gets a reference to the given bool and assigns it to the WorkProfileBlockNotificationsWhileDeviceLocked field.

### GetWorkProfileBlockAddingAccounts

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccounts() bool`

GetWorkProfileBlockAddingAccounts returns the WorkProfileBlockAddingAccounts field if non-nil, zero value otherwise.

### GetWorkProfileBlockAddingAccountsOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccountsOk() (bool, bool)`

GetWorkProfileBlockAddingAccountsOk returns a tuple with the WorkProfileBlockAddingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockAddingAccounts

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockAddingAccounts() bool`

HasWorkProfileBlockAddingAccounts returns a boolean if a field has been set.

### SetWorkProfileBlockAddingAccounts

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockAddingAccounts(v bool)`

SetWorkProfileBlockAddingAccounts gets a reference to the given bool and assigns it to the WorkProfileBlockAddingAccounts field.

### GetWorkProfileBluetoothEnableContactSharing

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharing() bool`

GetWorkProfileBluetoothEnableContactSharing returns the WorkProfileBluetoothEnableContactSharing field if non-nil, zero value otherwise.

### GetWorkProfileBluetoothEnableContactSharingOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharingOk() (bool, bool)`

GetWorkProfileBluetoothEnableContactSharingOk returns a tuple with the WorkProfileBluetoothEnableContactSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBluetoothEnableContactSharing

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBluetoothEnableContactSharing() bool`

HasWorkProfileBluetoothEnableContactSharing returns a boolean if a field has been set.

### SetWorkProfileBluetoothEnableContactSharing

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBluetoothEnableContactSharing(v bool)`

SetWorkProfileBluetoothEnableContactSharing gets a reference to the given bool and assigns it to the WorkProfileBluetoothEnableContactSharing field.

### GetWorkProfileBlockScreenCapture

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockScreenCapture() bool`

GetWorkProfileBlockScreenCapture returns the WorkProfileBlockScreenCapture field if non-nil, zero value otherwise.

### GetWorkProfileBlockScreenCaptureOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockScreenCaptureOk() (bool, bool)`

GetWorkProfileBlockScreenCaptureOk returns a tuple with the WorkProfileBlockScreenCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockScreenCapture

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockScreenCapture() bool`

HasWorkProfileBlockScreenCapture returns a boolean if a field has been set.

### SetWorkProfileBlockScreenCapture

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockScreenCapture(v bool)`

SetWorkProfileBlockScreenCapture gets a reference to the given bool and assigns it to the WorkProfileBlockScreenCapture field.

### GetWorkProfileBlockCrossProfileCallerId

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerId() bool`

GetWorkProfileBlockCrossProfileCallerId returns the WorkProfileBlockCrossProfileCallerId field if non-nil, zero value otherwise.

### GetWorkProfileBlockCrossProfileCallerIdOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerIdOk() (bool, bool)`

GetWorkProfileBlockCrossProfileCallerIdOk returns a tuple with the WorkProfileBlockCrossProfileCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCrossProfileCallerId

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCrossProfileCallerId() bool`

HasWorkProfileBlockCrossProfileCallerId returns a boolean if a field has been set.

### SetWorkProfileBlockCrossProfileCallerId

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCallerId(v bool)`

SetWorkProfileBlockCrossProfileCallerId gets a reference to the given bool and assigns it to the WorkProfileBlockCrossProfileCallerId field.

### GetWorkProfileBlockCamera

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCamera() bool`

GetWorkProfileBlockCamera returns the WorkProfileBlockCamera field if non-nil, zero value otherwise.

### GetWorkProfileBlockCameraOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCameraOk() (bool, bool)`

GetWorkProfileBlockCameraOk returns a tuple with the WorkProfileBlockCamera field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCamera

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCamera() bool`

HasWorkProfileBlockCamera returns a boolean if a field has been set.

### SetWorkProfileBlockCamera

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCamera(v bool)`

SetWorkProfileBlockCamera gets a reference to the given bool and assigns it to the WorkProfileBlockCamera field.

### GetWorkProfileBlockCrossProfileContactsSearch

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearch() bool`

GetWorkProfileBlockCrossProfileContactsSearch returns the WorkProfileBlockCrossProfileContactsSearch field if non-nil, zero value otherwise.

### GetWorkProfileBlockCrossProfileContactsSearchOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearchOk() (bool, bool)`

GetWorkProfileBlockCrossProfileContactsSearchOk returns a tuple with the WorkProfileBlockCrossProfileContactsSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCrossProfileContactsSearch

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCrossProfileContactsSearch() bool`

HasWorkProfileBlockCrossProfileContactsSearch returns a boolean if a field has been set.

### SetWorkProfileBlockCrossProfileContactsSearch

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileContactsSearch(v bool)`

SetWorkProfileBlockCrossProfileContactsSearch gets a reference to the given bool and assigns it to the WorkProfileBlockCrossProfileContactsSearch field.

### GetWorkProfileBlockCrossProfileCopyPaste

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPaste() bool`

GetWorkProfileBlockCrossProfileCopyPaste returns the WorkProfileBlockCrossProfileCopyPaste field if non-nil, zero value otherwise.

### GetWorkProfileBlockCrossProfileCopyPasteOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPasteOk() (bool, bool)`

GetWorkProfileBlockCrossProfileCopyPasteOk returns a tuple with the WorkProfileBlockCrossProfileCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCrossProfileCopyPaste

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCrossProfileCopyPaste() bool`

HasWorkProfileBlockCrossProfileCopyPaste returns a boolean if a field has been set.

### SetWorkProfileBlockCrossProfileCopyPaste

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCopyPaste(v bool)`

SetWorkProfileBlockCrossProfileCopyPaste gets a reference to the given bool and assigns it to the WorkProfileBlockCrossProfileCopyPaste field.

### GetWorkProfileDefaultAppPermissionPolicy

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicy() AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType`

GetWorkProfileDefaultAppPermissionPolicy returns the WorkProfileDefaultAppPermissionPolicy field if non-nil, zero value otherwise.

### GetWorkProfileDefaultAppPermissionPolicyOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicyOk() (AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType, bool)`

GetWorkProfileDefaultAppPermissionPolicyOk returns a tuple with the WorkProfileDefaultAppPermissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileDefaultAppPermissionPolicy

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileDefaultAppPermissionPolicy() bool`

HasWorkProfileDefaultAppPermissionPolicy returns a boolean if a field has been set.

### SetWorkProfileDefaultAppPermissionPolicy

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDefaultAppPermissionPolicy(v AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType)`

SetWorkProfileDefaultAppPermissionPolicy gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType and assigns it to the WorkProfileDefaultAppPermissionPolicy field.

### GetWorkProfilePasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlock() bool`

GetWorkProfilePasswordBlockFingerprintUnlock returns the WorkProfilePasswordBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetWorkProfilePasswordBlockFingerprintUnlockOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlockOk() (bool, bool)`

GetWorkProfilePasswordBlockFingerprintUnlockOk returns a tuple with the WorkProfilePasswordBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordBlockFingerprintUnlock() bool`

HasWorkProfilePasswordBlockFingerprintUnlock returns a boolean if a field has been set.

### SetWorkProfilePasswordBlockFingerprintUnlock

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFingerprintUnlock(v bool)`

SetWorkProfilePasswordBlockFingerprintUnlock gets a reference to the given bool and assigns it to the WorkProfilePasswordBlockFingerprintUnlock field.

### GetWorkProfilePasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgents() bool`

GetWorkProfilePasswordBlockTrustAgents returns the WorkProfilePasswordBlockTrustAgents field if non-nil, zero value otherwise.

### GetWorkProfilePasswordBlockTrustAgentsOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgentsOk() (bool, bool)`

GetWorkProfilePasswordBlockTrustAgentsOk returns a tuple with the WorkProfilePasswordBlockTrustAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordBlockTrustAgents() bool`

HasWorkProfilePasswordBlockTrustAgents returns a boolean if a field has been set.

### SetWorkProfilePasswordBlockTrustAgents

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockTrustAgents(v bool)`

SetWorkProfilePasswordBlockTrustAgents gets a reference to the given bool and assigns it to the WorkProfilePasswordBlockTrustAgents field.

### GetWorkProfilePasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays() int32`

GetWorkProfilePasswordExpirationDays returns the WorkProfilePasswordExpirationDays field if non-nil, zero value otherwise.

### GetWorkProfilePasswordExpirationDaysOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDaysOk() (int32, bool)`

GetWorkProfilePasswordExpirationDaysOk returns a tuple with the WorkProfilePasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordExpirationDays() bool`

HasWorkProfilePasswordExpirationDays returns a boolean if a field has been set.

### SetWorkProfilePasswordExpirationDays

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(v int32)`

SetWorkProfilePasswordExpirationDays gets a reference to the given int32 and assigns it to the WorkProfilePasswordExpirationDays field.

### SetWorkProfilePasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDaysExplicitNull(b bool)`

SetWorkProfilePasswordExpirationDaysExplicitNull (un)sets WorkProfilePasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordExpirationDays value is set to nil even if false is passed
### GetWorkProfilePasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength() int32`

GetWorkProfilePasswordMinimumLength returns the WorkProfilePasswordMinimumLength field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinimumLengthOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLengthOk() (int32, bool)`

GetWorkProfilePasswordMinimumLengthOk returns a tuple with the WorkProfilePasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinimumLength() bool`

HasWorkProfilePasswordMinimumLength returns a boolean if a field has been set.

### SetWorkProfilePasswordMinimumLength

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(v int32)`

SetWorkProfilePasswordMinimumLength gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinimumLength field.

### SetWorkProfilePasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLengthExplicitNull(b bool)`

SetWorkProfilePasswordMinimumLengthExplicitNull (un)sets WorkProfilePasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinimumLength value is set to nil even if false is passed
### GetWorkProfilePasswordMinNumericCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharacters() int32`

GetWorkProfilePasswordMinNumericCharacters returns the WorkProfilePasswordMinNumericCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinNumericCharactersOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinNumericCharactersOk returns a tuple with the WorkProfilePasswordMinNumericCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinNumericCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinNumericCharacters() bool`

HasWorkProfilePasswordMinNumericCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinNumericCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharacters(v int32)`

SetWorkProfilePasswordMinNumericCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinNumericCharacters field.

### SetWorkProfilePasswordMinNumericCharactersExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinNumericCharactersExplicitNull (un)sets WorkProfilePasswordMinNumericCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinNumericCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinNonLetterCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharacters() int32`

GetWorkProfilePasswordMinNonLetterCharacters returns the WorkProfilePasswordMinNonLetterCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinNonLetterCharactersOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinNonLetterCharactersOk returns a tuple with the WorkProfilePasswordMinNonLetterCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinNonLetterCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinNonLetterCharacters() bool`

HasWorkProfilePasswordMinNonLetterCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinNonLetterCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharacters(v int32)`

SetWorkProfilePasswordMinNonLetterCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinNonLetterCharacters field.

### SetWorkProfilePasswordMinNonLetterCharactersExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinNonLetterCharactersExplicitNull (un)sets WorkProfilePasswordMinNonLetterCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinNonLetterCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinLetterCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharacters() int32`

GetWorkProfilePasswordMinLetterCharacters returns the WorkProfilePasswordMinLetterCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinLetterCharactersOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinLetterCharactersOk returns a tuple with the WorkProfilePasswordMinLetterCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinLetterCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinLetterCharacters() bool`

HasWorkProfilePasswordMinLetterCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinLetterCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharacters(v int32)`

SetWorkProfilePasswordMinLetterCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinLetterCharacters field.

### SetWorkProfilePasswordMinLetterCharactersExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinLetterCharactersExplicitNull (un)sets WorkProfilePasswordMinLetterCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinLetterCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinLowerCaseCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharacters() int32`

GetWorkProfilePasswordMinLowerCaseCharacters returns the WorkProfilePasswordMinLowerCaseCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinLowerCaseCharactersOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinLowerCaseCharactersOk returns a tuple with the WorkProfilePasswordMinLowerCaseCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinLowerCaseCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinLowerCaseCharacters() bool`

HasWorkProfilePasswordMinLowerCaseCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinLowerCaseCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharacters(v int32)`

SetWorkProfilePasswordMinLowerCaseCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinLowerCaseCharacters field.

### SetWorkProfilePasswordMinLowerCaseCharactersExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinLowerCaseCharactersExplicitNull (un)sets WorkProfilePasswordMinLowerCaseCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinLowerCaseCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinUpperCaseCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharacters() int32`

GetWorkProfilePasswordMinUpperCaseCharacters returns the WorkProfilePasswordMinUpperCaseCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinUpperCaseCharactersOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinUpperCaseCharactersOk returns a tuple with the WorkProfilePasswordMinUpperCaseCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinUpperCaseCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinUpperCaseCharacters() bool`

HasWorkProfilePasswordMinUpperCaseCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinUpperCaseCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharacters(v int32)`

SetWorkProfilePasswordMinUpperCaseCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinUpperCaseCharacters field.

### SetWorkProfilePasswordMinUpperCaseCharactersExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinUpperCaseCharactersExplicitNull (un)sets WorkProfilePasswordMinUpperCaseCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinUpperCaseCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinSymbolCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharacters() int32`

GetWorkProfilePasswordMinSymbolCharacters returns the WorkProfilePasswordMinSymbolCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinSymbolCharactersOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinSymbolCharactersOk returns a tuple with the WorkProfilePasswordMinSymbolCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinSymbolCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinSymbolCharacters() bool`

HasWorkProfilePasswordMinSymbolCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinSymbolCharacters

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharacters(v int32)`

SetWorkProfilePasswordMinSymbolCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinSymbolCharacters field.

### SetWorkProfilePasswordMinSymbolCharactersExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinSymbolCharactersExplicitNull (un)sets WorkProfilePasswordMinSymbolCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinSymbolCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout returns the WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetWorkProfilePasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCount() int32`

GetWorkProfilePasswordPreviousPasswordBlockCount returns the WorkProfilePasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetWorkProfilePasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetWorkProfilePasswordPreviousPasswordBlockCountOk returns a tuple with the WorkProfilePasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordPreviousPasswordBlockCount() bool`

HasWorkProfilePasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetWorkProfilePasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCount(v int32)`

SetWorkProfilePasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the WorkProfilePasswordPreviousPasswordBlockCount field.

### SetWorkProfilePasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetWorkProfilePasswordPreviousPasswordBlockCountExplicitNull (un)sets WorkProfilePasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset() int32`

GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset returns the WorkProfilePasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetWorkProfilePasswordSignInFailureCountBeforeFactoryResetOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetWorkProfilePasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the WorkProfilePasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordSignInFailureCountBeforeFactoryReset() bool`

HasWorkProfilePasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the WorkProfilePasswordSignInFailureCountBeforeFactoryReset field.

### SetWorkProfilePasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetWorkProfilePasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets WorkProfilePasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetWorkProfilePasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType() AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType`

GetWorkProfilePasswordRequiredType returns the WorkProfilePasswordRequiredType field if non-nil, zero value otherwise.

### GetWorkProfilePasswordRequiredTypeOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType, bool)`

GetWorkProfilePasswordRequiredTypeOk returns a tuple with the WorkProfilePasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordRequiredType() bool`

HasWorkProfilePasswordRequiredType returns a boolean if a field has been set.

### SetWorkProfilePasswordRequiredType

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(v AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType)`

SetWorkProfilePasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType and assigns it to the WorkProfilePasswordRequiredType field.

### GetWorkProfileRequirePassword

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequirePassword() bool`

GetWorkProfileRequirePassword returns the WorkProfileRequirePassword field if non-nil, zero value otherwise.

### GetWorkProfileRequirePasswordOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequirePasswordOk() (bool, bool)`

GetWorkProfileRequirePasswordOk returns a tuple with the WorkProfileRequirePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileRequirePassword

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileRequirePassword() bool`

HasWorkProfileRequirePassword returns a boolean if a field has been set.

### SetWorkProfileRequirePassword

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileRequirePassword(v bool)`

SetWorkProfileRequirePassword gets a reference to the given bool and assigns it to the WorkProfileRequirePassword field.

### GetSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetSecurityRequireVerifyApps() bool`

GetSecurityRequireVerifyApps returns the SecurityRequireVerifyApps field if non-nil, zero value otherwise.

### GetSecurityRequireVerifyAppsOk

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) GetSecurityRequireVerifyAppsOk() (bool, bool)`

GetSecurityRequireVerifyAppsOk returns a tuple with the SecurityRequireVerifyApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) HasSecurityRequireVerifyApps() bool`

HasSecurityRequireVerifyApps returns a boolean if a field has been set.

### SetSecurityRequireVerifyApps

`func (o *MicrosoftGraphAndroidWorkProfileGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(v bool)`

SetSecurityRequireVerifyApps gets a reference to the given bool and assigns it to the SecurityRequireVerifyApps field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


