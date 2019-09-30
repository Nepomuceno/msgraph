# MicrosoftGraphWindows10MobileCompliancePolicy

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
**PasswordRequired** | Pointer to **bool** | Require a password to unlock Windows Phone device. | [optional] 
**PasswordBlockSimple** | Pointer to **bool** | Whether or not to block syncing the calendar. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum password length. Valid values 4 to 16 | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required password type. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | The number of previous passwords to prevent re-use of. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before password expiration. Valid values 1 to 255 | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a password is required. | [optional] 
**PasswordRequireToUnlockFromIdle** | Pointer to **bool** | Require a password to unlock an idle device. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum Windows Phone version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum Windows Phone version. | [optional] 
**EarlyLaunchAntiMalwareDriverEnabled** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled. | [optional] 
**BitLockerEnabled** | Pointer to **bool** | Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled | [optional] 
**SecureBootEnabled** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled. | [optional] 
**CodeIntegrityEnabled** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Require encryption on windows devices. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetScheduledActionsForRule

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetScheduledActionsForRule() []MicrosoftGraphDeviceComplianceScheduledActionForRule`

GetScheduledActionsForRule returns the ScheduledActionsForRule field if non-nil, zero value otherwise.

### GetScheduledActionsForRuleOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetScheduledActionsForRuleOk() ([]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool)`

GetScheduledActionsForRuleOk returns a tuple with the ScheduledActionsForRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionsForRule

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasScheduledActionsForRule() bool`

HasScheduledActionsForRule returns a boolean if a field has been set.

### SetScheduledActionsForRule

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetScheduledActionsForRule(v []MicrosoftGraphDeviceComplianceScheduledActionForRule)`

SetScheduledActionsForRule gets a reference to the given []MicrosoftGraphDeviceComplianceScheduledActionForRule and assigns it to the ScheduledActionsForRule field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDeviceStatuses() []MicrosoftGraphDeviceComplianceDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceComplianceDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDeviceStatuses(v []MicrosoftGraphDeviceComplianceDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetUserStatuses() []MicrosoftGraphDeviceComplianceUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetUserStatusesOk() ([]MicrosoftGraphDeviceComplianceUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetUserStatuses(v []MicrosoftGraphDeviceComplianceUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceComplianceDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceComplianceUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAssignments

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetAssignments() []MicrosoftGraphDeviceCompliancePolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetAssignmentsOk() ([]MicrosoftGraphDeviceCompliancePolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetAssignments(v []MicrosoftGraphDeviceCompliancePolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyAssignment and assigns it to the Assignments field.

### GetPasswordRequired

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordBlockSimple

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordMinimumLength

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordExpirationDays

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordRequireToUnlockFromIdle

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordRequireToUnlockFromIdle() bool`

GetPasswordRequireToUnlockFromIdle returns the PasswordRequireToUnlockFromIdle field if non-nil, zero value otherwise.

### GetPasswordRequireToUnlockFromIdleOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetPasswordRequireToUnlockFromIdleOk() (bool, bool)`

GetPasswordRequireToUnlockFromIdleOk returns a tuple with the PasswordRequireToUnlockFromIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequireToUnlockFromIdle

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasPasswordRequireToUnlockFromIdle() bool`

HasPasswordRequireToUnlockFromIdle returns a boolean if a field has been set.

### SetPasswordRequireToUnlockFromIdle

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetPasswordRequireToUnlockFromIdle(v bool)`

SetPasswordRequireToUnlockFromIdle gets a reference to the given bool and assigns it to the PasswordRequireToUnlockFromIdle field.

### GetOsMinimumVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetEarlyLaunchAntiMalwareDriverEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabled() bool`

GetEarlyLaunchAntiMalwareDriverEnabled returns the EarlyLaunchAntiMalwareDriverEnabled field if non-nil, zero value otherwise.

### GetEarlyLaunchAntiMalwareDriverEnabledOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabledOk() (bool, bool)`

GetEarlyLaunchAntiMalwareDriverEnabledOk returns a tuple with the EarlyLaunchAntiMalwareDriverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEarlyLaunchAntiMalwareDriverEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasEarlyLaunchAntiMalwareDriverEnabled() bool`

HasEarlyLaunchAntiMalwareDriverEnabled returns a boolean if a field has been set.

### SetEarlyLaunchAntiMalwareDriverEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetEarlyLaunchAntiMalwareDriverEnabled(v bool)`

SetEarlyLaunchAntiMalwareDriverEnabled gets a reference to the given bool and assigns it to the EarlyLaunchAntiMalwareDriverEnabled field.

### GetBitLockerEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetBitLockerEnabled() bool`

GetBitLockerEnabled returns the BitLockerEnabled field if non-nil, zero value otherwise.

### GetBitLockerEnabledOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetBitLockerEnabledOk() (bool, bool)`

GetBitLockerEnabledOk returns a tuple with the BitLockerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasBitLockerEnabled() bool`

HasBitLockerEnabled returns a boolean if a field has been set.

### SetBitLockerEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetBitLockerEnabled(v bool)`

SetBitLockerEnabled gets a reference to the given bool and assigns it to the BitLockerEnabled field.

### GetSecureBootEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetSecureBootEnabledOk() (bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecureBootEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### SetSecureBootEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled gets a reference to the given bool and assigns it to the SecureBootEnabled field.

### GetCodeIntegrityEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetCodeIntegrityEnabled() bool`

GetCodeIntegrityEnabled returns the CodeIntegrityEnabled field if non-nil, zero value otherwise.

### GetCodeIntegrityEnabledOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetCodeIntegrityEnabledOk() (bool, bool)`

GetCodeIntegrityEnabledOk returns a tuple with the CodeIntegrityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCodeIntegrityEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasCodeIntegrityEnabled() bool`

HasCodeIntegrityEnabled returns a boolean if a field has been set.

### SetCodeIntegrityEnabled

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetCodeIntegrityEnabled(v bool)`

SetCodeIntegrityEnabled gets a reference to the given bool and assigns it to the CodeIntegrityEnabled field.

### GetStorageRequireEncryption

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *MicrosoftGraphWindows10MobileCompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


