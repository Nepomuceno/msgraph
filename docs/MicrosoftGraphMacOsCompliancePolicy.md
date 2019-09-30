# MicrosoftGraphMacOsCompliancePolicy

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
**PasswordRequired** | Pointer to **bool** | Whether or not to require a password. | [optional] 
**PasswordBlockSimple** | Pointer to **bool** | Indicates whether or not to block simple passwords. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. Valid values 1 to 65535 | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of password. Valid values 4 to 14 | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a password is required. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 1 to 24 | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required password type. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum MacOS version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum MacOS version. | [optional] 
**SystemIntegrityProtectionEnabled** | Pointer to **bool** | Require that devices have enabled system integrity protection. | [optional] 
**DeviceThreatProtectionEnabled** | Pointer to **bool** | Require that devices have enabled device threat protection. | [optional] 
**DeviceThreatProtectionRequiredSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphDeviceThreatProtectionLevel**](anyOf&lt;microsoft.graph.deviceThreatProtectionLevel&gt;.md) | Require Mobile Threat Protection minimum risk level to report noncompliance. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Require encryption on Mac OS devices. | [optional] 
**FirewallEnabled** | Pointer to **bool** | Whether the firewall should be enabled or not. | [optional] 
**FirewallBlockAllIncoming** | Pointer to **bool** | Corresponds to the “Block all incoming connections” option. | [optional] 
**FirewallEnableStealthMode** | Pointer to **bool** | Corresponds to “Enable stealth mode.” | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetScheduledActionsForRule

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetScheduledActionsForRule() []MicrosoftGraphDeviceComplianceScheduledActionForRule`

GetScheduledActionsForRule returns the ScheduledActionsForRule field if non-nil, zero value otherwise.

### GetScheduledActionsForRuleOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetScheduledActionsForRuleOk() ([]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool)`

GetScheduledActionsForRuleOk returns a tuple with the ScheduledActionsForRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionsForRule

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasScheduledActionsForRule() bool`

HasScheduledActionsForRule returns a boolean if a field has been set.

### SetScheduledActionsForRule

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetScheduledActionsForRule(v []MicrosoftGraphDeviceComplianceScheduledActionForRule)`

SetScheduledActionsForRule gets a reference to the given []MicrosoftGraphDeviceComplianceScheduledActionForRule and assigns it to the ScheduledActionsForRule field.

### GetDeviceStatuses

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceStatuses() []MicrosoftGraphDeviceComplianceDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceComplianceDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDeviceStatuses(v []MicrosoftGraphDeviceComplianceDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetUserStatuses() []MicrosoftGraphDeviceComplianceUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetUserStatusesOk() ([]MicrosoftGraphDeviceComplianceUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetUserStatuses(v []MicrosoftGraphDeviceComplianceUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceComplianceDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceComplianceUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAssignments

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetAssignments() []MicrosoftGraphDeviceCompliancePolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetAssignmentsOk() ([]MicrosoftGraphDeviceCompliancePolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetAssignments(v []MicrosoftGraphDeviceCompliancePolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyAssignment and assigns it to the Assignments field.

### GetPasswordRequired

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordBlockSimple

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetOsMinimumVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetSystemIntegrityProtectionEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetSystemIntegrityProtectionEnabled() bool`

GetSystemIntegrityProtectionEnabled returns the SystemIntegrityProtectionEnabled field if non-nil, zero value otherwise.

### GetSystemIntegrityProtectionEnabledOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetSystemIntegrityProtectionEnabledOk() (bool, bool)`

GetSystemIntegrityProtectionEnabledOk returns a tuple with the SystemIntegrityProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystemIntegrityProtectionEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasSystemIntegrityProtectionEnabled() bool`

HasSystemIntegrityProtectionEnabled returns a boolean if a field has been set.

### SetSystemIntegrityProtectionEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetSystemIntegrityProtectionEnabled(v bool)`

SetSystemIntegrityProtectionEnabled gets a reference to the given bool and assigns it to the SystemIntegrityProtectionEnabled field.

### GetDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceThreatProtectionEnabled() bool`

GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionEnabledOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool)`

GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDeviceThreatProtectionEnabled() bool`

HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.

### SetDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool)`

SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.

### GetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel`

GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionRequiredSecurityLevelOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool)`

GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool`

HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.

### SetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel)`

SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.

### GetStorageRequireEncryption

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.

### GetFirewallEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetFirewallEnabledOk() (bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### SetFirewallEnabled

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetFirewallEnabled(v bool)`

SetFirewallEnabled gets a reference to the given bool and assigns it to the FirewallEnabled field.

### GetFirewallBlockAllIncoming

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetFirewallBlockAllIncoming() bool`

GetFirewallBlockAllIncoming returns the FirewallBlockAllIncoming field if non-nil, zero value otherwise.

### GetFirewallBlockAllIncomingOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetFirewallBlockAllIncomingOk() (bool, bool)`

GetFirewallBlockAllIncomingOk returns a tuple with the FirewallBlockAllIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallBlockAllIncoming

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasFirewallBlockAllIncoming() bool`

HasFirewallBlockAllIncoming returns a boolean if a field has been set.

### SetFirewallBlockAllIncoming

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetFirewallBlockAllIncoming(v bool)`

SetFirewallBlockAllIncoming gets a reference to the given bool and assigns it to the FirewallBlockAllIncoming field.

### GetFirewallEnableStealthMode

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetFirewallEnableStealthMode() bool`

GetFirewallEnableStealthMode returns the FirewallEnableStealthMode field if non-nil, zero value otherwise.

### GetFirewallEnableStealthModeOk

`func (o *MicrosoftGraphMacOsCompliancePolicy) GetFirewallEnableStealthModeOk() (bool, bool)`

GetFirewallEnableStealthModeOk returns a tuple with the FirewallEnableStealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallEnableStealthMode

`func (o *MicrosoftGraphMacOsCompliancePolicy) HasFirewallEnableStealthMode() bool`

HasFirewallEnableStealthMode returns a boolean if a field has been set.

### SetFirewallEnableStealthMode

`func (o *MicrosoftGraphMacOsCompliancePolicy) SetFirewallEnableStealthMode(v bool)`

SetFirewallEnableStealthMode gets a reference to the given bool and assigns it to the FirewallEnableStealthMode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


