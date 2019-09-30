# MicrosoftGraphIosCompliancePolicy

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
**PasscodeBlockSimple** | Pointer to **bool** | Indicates whether or not to block simple passcodes. | [optional] 
**PasscodeExpirationDays** | Pointer to **int32** | Number of days before the passcode expires. Valid values 1 to 65535 | [optional] 
**PasscodeMinimumLength** | Pointer to **int32** | Minimum length of passcode. Valid values 4 to 14 | [optional] 
**PasscodeMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a passcode is required. | [optional] 
**PasscodePreviousPasscodeBlockCount** | Pointer to **int32** | Number of previous passcodes to block. Valid values 1 to 24 | [optional] 
**PasscodeMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasscodeRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required passcode type. | [optional] 
**PasscodeRequired** | Pointer to **bool** | Indicates whether or not to require a passcode. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum IOS version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum IOS version. | [optional] 
**SecurityBlockJailbrokenDevices** | Pointer to **bool** | Devices must not be jailbroken or rooted. | [optional] 
**DeviceThreatProtectionEnabled** | Pointer to **bool** | Require that devices have enabled device threat protection . | [optional] 
**DeviceThreatProtectionRequiredSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphDeviceThreatProtectionLevel**](anyOf&lt;microsoft.graph.deviceThreatProtectionLevel&gt;.md) | Require Mobile Threat Protection minimum risk level to report noncompliance. | [optional] 
**ManagedEmailProfileRequired** | Pointer to **bool** | Indicates whether or not to require a managed email profile. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphIosCompliancePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosCompliancePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosCompliancePolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphIosCompliancePolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosCompliancePolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosCompliancePolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphIosCompliancePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosCompliancePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosCompliancePolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosCompliancePolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosCompliancePolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosCompliancePolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphIosCompliancePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosCompliancePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosCompliancePolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphIosCompliancePolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphIosCompliancePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphIosCompliancePolicy) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetScheduledActionsForRule

`func (o *MicrosoftGraphIosCompliancePolicy) GetScheduledActionsForRule() []MicrosoftGraphDeviceComplianceScheduledActionForRule`

GetScheduledActionsForRule returns the ScheduledActionsForRule field if non-nil, zero value otherwise.

### GetScheduledActionsForRuleOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetScheduledActionsForRuleOk() ([]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool)`

GetScheduledActionsForRuleOk returns a tuple with the ScheduledActionsForRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionsForRule

`func (o *MicrosoftGraphIosCompliancePolicy) HasScheduledActionsForRule() bool`

HasScheduledActionsForRule returns a boolean if a field has been set.

### SetScheduledActionsForRule

`func (o *MicrosoftGraphIosCompliancePolicy) SetScheduledActionsForRule(v []MicrosoftGraphDeviceComplianceScheduledActionForRule)`

SetScheduledActionsForRule gets a reference to the given []MicrosoftGraphDeviceComplianceScheduledActionForRule and assigns it to the ScheduledActionsForRule field.

### GetDeviceStatuses

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceStatuses() []MicrosoftGraphDeviceComplianceDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceComplianceDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphIosCompliancePolicy) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphIosCompliancePolicy) SetDeviceStatuses(v []MicrosoftGraphDeviceComplianceDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphIosCompliancePolicy) GetUserStatuses() []MicrosoftGraphDeviceComplianceUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetUserStatusesOk() ([]MicrosoftGraphDeviceComplianceUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphIosCompliancePolicy) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphIosCompliancePolicy) SetUserStatuses(v []MicrosoftGraphDeviceComplianceUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceComplianceUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceComplianceDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphIosCompliancePolicy) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphIosCompliancePolicy) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphIosCompliancePolicy) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceComplianceUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceComplianceUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphIosCompliancePolicy) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphIosCompliancePolicy) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosCompliancePolicy) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosCompliancePolicy) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAssignments

`func (o *MicrosoftGraphIosCompliancePolicy) GetAssignments() []MicrosoftGraphDeviceCompliancePolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetAssignmentsOk() ([]MicrosoftGraphDeviceCompliancePolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosCompliancePolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosCompliancePolicy) SetAssignments(v []MicrosoftGraphDeviceCompliancePolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyAssignment and assigns it to the Assignments field.

### GetPasscodeBlockSimple

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeBlockSimple() bool`

GetPasscodeBlockSimple returns the PasscodeBlockSimple field if non-nil, zero value otherwise.

### GetPasscodeBlockSimpleOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeBlockSimpleOk() (bool, bool)`

GetPasscodeBlockSimpleOk returns a tuple with the PasscodeBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockSimple

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeBlockSimple() bool`

HasPasscodeBlockSimple returns a boolean if a field has been set.

### SetPasscodeBlockSimple

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeBlockSimple(v bool)`

SetPasscodeBlockSimple gets a reference to the given bool and assigns it to the PasscodeBlockSimple field.

### GetPasscodeExpirationDays

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeExpirationDays() int32`

GetPasscodeExpirationDays returns the PasscodeExpirationDays field if non-nil, zero value otherwise.

### GetPasscodeExpirationDaysOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeExpirationDaysOk() (int32, bool)`

GetPasscodeExpirationDaysOk returns a tuple with the PasscodeExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeExpirationDays

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeExpirationDays() bool`

HasPasscodeExpirationDays returns a boolean if a field has been set.

### SetPasscodeExpirationDays

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeExpirationDays(v int32)`

SetPasscodeExpirationDays gets a reference to the given int32 and assigns it to the PasscodeExpirationDays field.

### SetPasscodeExpirationDaysExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeExpirationDaysExplicitNull(b bool)`

SetPasscodeExpirationDaysExplicitNull (un)sets PasscodeExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeExpirationDays value is set to nil even if false is passed
### GetPasscodeMinimumLength

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeMinimumLength() int32`

GetPasscodeMinimumLength returns the PasscodeMinimumLength field if non-nil, zero value otherwise.

### GetPasscodeMinimumLengthOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeMinimumLengthOk() (int32, bool)`

GetPasscodeMinimumLengthOk returns a tuple with the PasscodeMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumLength

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeMinimumLength() bool`

HasPasscodeMinimumLength returns a boolean if a field has been set.

### SetPasscodeMinimumLength

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeMinimumLength(v int32)`

SetPasscodeMinimumLength gets a reference to the given int32 and assigns it to the PasscodeMinimumLength field.

### SetPasscodeMinimumLengthExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeMinimumLengthExplicitNull(b bool)`

SetPasscodeMinimumLengthExplicitNull (un)sets PasscodeMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumLength value is set to nil even if false is passed
### GetPasscodeMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLock() int32`

GetPasscodeMinutesOfInactivityBeforeLock returns the PasscodeMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasscodeMinutesOfInactivityBeforeLockOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasscodeMinutesOfInactivityBeforeLockOk returns a tuple with the PasscodeMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeMinutesOfInactivityBeforeLock() bool`

HasPasscodeMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasscodeMinutesOfInactivityBeforeLock

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLock(v int32)`

SetPasscodeMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeLock field.

### SetPasscodeMinutesOfInactivityBeforeLockExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasscodeMinutesOfInactivityBeforeLockExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasscodePreviousPasscodeBlockCount

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodePreviousPasscodeBlockCount() int32`

GetPasscodePreviousPasscodeBlockCount returns the PasscodePreviousPasscodeBlockCount field if non-nil, zero value otherwise.

### GetPasscodePreviousPasscodeBlockCountOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodePreviousPasscodeBlockCountOk() (int32, bool)`

GetPasscodePreviousPasscodeBlockCountOk returns a tuple with the PasscodePreviousPasscodeBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodePreviousPasscodeBlockCount

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodePreviousPasscodeBlockCount() bool`

HasPasscodePreviousPasscodeBlockCount returns a boolean if a field has been set.

### SetPasscodePreviousPasscodeBlockCount

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodePreviousPasscodeBlockCount(v int32)`

SetPasscodePreviousPasscodeBlockCount gets a reference to the given int32 and assigns it to the PasscodePreviousPasscodeBlockCount field.

### SetPasscodePreviousPasscodeBlockCountExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodePreviousPasscodeBlockCountExplicitNull(b bool)`

SetPasscodePreviousPasscodeBlockCountExplicitNull (un)sets PasscodePreviousPasscodeBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodePreviousPasscodeBlockCount value is set to nil even if false is passed
### GetPasscodeMinimumCharacterSetCount

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeMinimumCharacterSetCount() int32`

GetPasscodeMinimumCharacterSetCount returns the PasscodeMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasscodeMinimumCharacterSetCountOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeMinimumCharacterSetCountOk() (int32, bool)`

GetPasscodeMinimumCharacterSetCountOk returns a tuple with the PasscodeMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumCharacterSetCount

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeMinimumCharacterSetCount() bool`

HasPasscodeMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasscodeMinimumCharacterSetCount

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeMinimumCharacterSetCount(v int32)`

SetPasscodeMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasscodeMinimumCharacterSetCount field.

### SetPasscodeMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeMinimumCharacterSetCountExplicitNull(b bool)`

SetPasscodeMinimumCharacterSetCountExplicitNull (un)sets PasscodeMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasscodeRequiredType

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasscodeRequiredType returns the PasscodeRequiredType field if non-nil, zero value otherwise.

### GetPasscodeRequiredTypeOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasscodeRequiredTypeOk returns a tuple with the PasscodeRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequiredType

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeRequiredType() bool`

HasPasscodeRequiredType returns a boolean if a field has been set.

### SetPasscodeRequiredType

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasscodeRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasscodeRequiredType field.

### GetPasscodeRequired

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeRequired() bool`

GetPasscodeRequired returns the PasscodeRequired field if non-nil, zero value otherwise.

### GetPasscodeRequiredOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetPasscodeRequiredOk() (bool, bool)`

GetPasscodeRequiredOk returns a tuple with the PasscodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequired

`func (o *MicrosoftGraphIosCompliancePolicy) HasPasscodeRequired() bool`

HasPasscodeRequired returns a boolean if a field has been set.

### SetPasscodeRequired

`func (o *MicrosoftGraphIosCompliancePolicy) SetPasscodeRequired(v bool)`

SetPasscodeRequired gets a reference to the given bool and assigns it to the PasscodeRequired field.

### GetOsMinimumVersion

`func (o *MicrosoftGraphIosCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MicrosoftGraphIosCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MicrosoftGraphIosCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MicrosoftGraphIosCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MicrosoftGraphIosCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MicrosoftGraphIosCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MicrosoftGraphIosCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetSecurityBlockJailbrokenDevices

`func (o *MicrosoftGraphIosCompliancePolicy) GetSecurityBlockJailbrokenDevices() bool`

GetSecurityBlockJailbrokenDevices returns the SecurityBlockJailbrokenDevices field if non-nil, zero value otherwise.

### GetSecurityBlockJailbrokenDevicesOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetSecurityBlockJailbrokenDevicesOk() (bool, bool)`

GetSecurityBlockJailbrokenDevicesOk returns a tuple with the SecurityBlockJailbrokenDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityBlockJailbrokenDevices

`func (o *MicrosoftGraphIosCompliancePolicy) HasSecurityBlockJailbrokenDevices() bool`

HasSecurityBlockJailbrokenDevices returns a boolean if a field has been set.

### SetSecurityBlockJailbrokenDevices

`func (o *MicrosoftGraphIosCompliancePolicy) SetSecurityBlockJailbrokenDevices(v bool)`

SetSecurityBlockJailbrokenDevices gets a reference to the given bool and assigns it to the SecurityBlockJailbrokenDevices field.

### GetDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceThreatProtectionEnabled() bool`

GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionEnabledOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool)`

GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphIosCompliancePolicy) HasDeviceThreatProtectionEnabled() bool`

HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.

### SetDeviceThreatProtectionEnabled

`func (o *MicrosoftGraphIosCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool)`

SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.

### GetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel`

GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionRequiredSecurityLevelOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool)`

GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphIosCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool`

HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.

### SetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MicrosoftGraphIosCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel)`

SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.

### GetManagedEmailProfileRequired

`func (o *MicrosoftGraphIosCompliancePolicy) GetManagedEmailProfileRequired() bool`

GetManagedEmailProfileRequired returns the ManagedEmailProfileRequired field if non-nil, zero value otherwise.

### GetManagedEmailProfileRequiredOk

`func (o *MicrosoftGraphIosCompliancePolicy) GetManagedEmailProfileRequiredOk() (bool, bool)`

GetManagedEmailProfileRequiredOk returns a tuple with the ManagedEmailProfileRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedEmailProfileRequired

`func (o *MicrosoftGraphIosCompliancePolicy) HasManagedEmailProfileRequired() bool`

HasManagedEmailProfileRequired returns a boolean if a field has been set.

### SetManagedEmailProfileRequired

`func (o *MicrosoftGraphIosCompliancePolicy) SetManagedEmailProfileRequired(v bool)`

SetManagedEmailProfileRequired gets a reference to the given bool and assigns it to the ManagedEmailProfileRequired field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


