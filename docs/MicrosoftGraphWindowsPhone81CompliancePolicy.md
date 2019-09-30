# MicrosoftGraphWindowsPhone81CompliancePolicy

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
**PasswordBlockSimple** | Pointer to **bool** | Whether or not to block syncing the calendar. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of passwords. | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a password is required. | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required password type. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 0 to 24 | [optional] 
**PasswordRequired** | Pointer to **bool** | Whether or not to require a password. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum Windows Phone version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum Windows Phone version. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Require encryption on windows phone devices. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetScheduledActionsForRule

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetScheduledActionsForRule() []MicrosoftGraphDeviceComplianceScheduledActionForRule`

GetScheduledActionsForRule returns the ScheduledActionsForRule field if non-nil, zero value otherwise.

### GetScheduledActionsForRuleOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetScheduledActionsForRuleOk() ([]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool)`

GetScheduledActionsForRuleOk returns a tuple with the ScheduledActionsForRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionsForRule

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasScheduledActionsForRule() bool`

HasScheduledActionsForRule returns a boolean if a field has been set.

### SetScheduledActionsForRule

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetScheduledActionsForRule(v []MicrosoftGraphDeviceComplianceScheduledActionForRule)`

SetScheduledActionsForRule gets a reference to the given []MicrosoftGraphDeviceComplianceScheduledActionForRule and assigns it to the ScheduledActionsForRule field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDeviceStatuses() []MicrosoftGraphDeviceComplianceDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceComplianceDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDeviceStatuses(v []MicrosoftGraphDeviceComplianceDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetUserStatuses() []MicrosoftGraphDeviceComplianceUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetUserStatusesOk() ([]MicrosoftGraphDeviceComplianceUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetUserStatuses(v []MicrosoftGraphDeviceComplianceUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceComplianceDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceComplianceUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAssignments

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetAssignments() []MicrosoftGraphDeviceCompliancePolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetAssignmentsOk() ([]MicrosoftGraphDeviceCompliancePolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetAssignments(v []MicrosoftGraphDeviceCompliancePolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyAssignment and assigns it to the Assignments field.

### GetPasswordBlockSimple

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordRequired

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetOsMinimumVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetStorageRequireEncryption

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *MicrosoftGraphWindowsPhone81CompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


