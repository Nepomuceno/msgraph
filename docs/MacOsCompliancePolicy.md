# MacOsCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetPasswordRequired

`func (o *MacOsCompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MacOsCompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MacOsCompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MacOsCompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordBlockSimple

`func (o *MacOsCompliancePolicy) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MacOsCompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MacOsCompliancePolicy) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MacOsCompliancePolicy) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *MacOsCompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MacOsCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MacOsCompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MacOsCompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MacOsCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MacOsCompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MacOsCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MacOsCompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MacOsCompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MacOsCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeLock

`func (o *MacOsCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *MacOsCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *MacOsCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *MacOsCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *MacOsCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MacOsCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MacOsCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MacOsCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MacOsCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MacOsCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MacOsCompliancePolicy) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MacOsCompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MacOsCompliancePolicy) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MacOsCompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MacOsCompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MacOsCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MacOsCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MacOsCompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MacOsCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetOsMinimumVersion

`func (o *MacOsCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MacOsCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MacOsCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MacOsCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MacOsCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MacOsCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MacOsCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MacOsCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MacOsCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MacOsCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetSystemIntegrityProtectionEnabled

`func (o *MacOsCompliancePolicy) GetSystemIntegrityProtectionEnabled() bool`

GetSystemIntegrityProtectionEnabled returns the SystemIntegrityProtectionEnabled field if non-nil, zero value otherwise.

### GetSystemIntegrityProtectionEnabledOk

`func (o *MacOsCompliancePolicy) GetSystemIntegrityProtectionEnabledOk() (bool, bool)`

GetSystemIntegrityProtectionEnabledOk returns a tuple with the SystemIntegrityProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystemIntegrityProtectionEnabled

`func (o *MacOsCompliancePolicy) HasSystemIntegrityProtectionEnabled() bool`

HasSystemIntegrityProtectionEnabled returns a boolean if a field has been set.

### SetSystemIntegrityProtectionEnabled

`func (o *MacOsCompliancePolicy) SetSystemIntegrityProtectionEnabled(v bool)`

SetSystemIntegrityProtectionEnabled gets a reference to the given bool and assigns it to the SystemIntegrityProtectionEnabled field.

### GetDeviceThreatProtectionEnabled

`func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionEnabled() bool`

GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionEnabledOk

`func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool)`

GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionEnabled

`func (o *MacOsCompliancePolicy) HasDeviceThreatProtectionEnabled() bool`

HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.

### SetDeviceThreatProtectionEnabled

`func (o *MacOsCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool)`

SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.

### GetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel`

GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionRequiredSecurityLevelOk

`func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool)`

GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionRequiredSecurityLevel

`func (o *MacOsCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool`

HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.

### SetDeviceThreatProtectionRequiredSecurityLevel

`func (o *MacOsCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel)`

SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.

### GetStorageRequireEncryption

`func (o *MacOsCompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *MacOsCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *MacOsCompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *MacOsCompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.

### GetFirewallEnabled

`func (o *MacOsCompliancePolicy) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *MacOsCompliancePolicy) GetFirewallEnabledOk() (bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallEnabled

`func (o *MacOsCompliancePolicy) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### SetFirewallEnabled

`func (o *MacOsCompliancePolicy) SetFirewallEnabled(v bool)`

SetFirewallEnabled gets a reference to the given bool and assigns it to the FirewallEnabled field.

### GetFirewallBlockAllIncoming

`func (o *MacOsCompliancePolicy) GetFirewallBlockAllIncoming() bool`

GetFirewallBlockAllIncoming returns the FirewallBlockAllIncoming field if non-nil, zero value otherwise.

### GetFirewallBlockAllIncomingOk

`func (o *MacOsCompliancePolicy) GetFirewallBlockAllIncomingOk() (bool, bool)`

GetFirewallBlockAllIncomingOk returns a tuple with the FirewallBlockAllIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallBlockAllIncoming

`func (o *MacOsCompliancePolicy) HasFirewallBlockAllIncoming() bool`

HasFirewallBlockAllIncoming returns a boolean if a field has been set.

### SetFirewallBlockAllIncoming

`func (o *MacOsCompliancePolicy) SetFirewallBlockAllIncoming(v bool)`

SetFirewallBlockAllIncoming gets a reference to the given bool and assigns it to the FirewallBlockAllIncoming field.

### GetFirewallEnableStealthMode

`func (o *MacOsCompliancePolicy) GetFirewallEnableStealthMode() bool`

GetFirewallEnableStealthMode returns the FirewallEnableStealthMode field if non-nil, zero value otherwise.

### GetFirewallEnableStealthModeOk

`func (o *MacOsCompliancePolicy) GetFirewallEnableStealthModeOk() (bool, bool)`

GetFirewallEnableStealthModeOk returns a tuple with the FirewallEnableStealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallEnableStealthMode

`func (o *MacOsCompliancePolicy) HasFirewallEnableStealthMode() bool`

HasFirewallEnableStealthMode returns a boolean if a field has been set.

### SetFirewallEnableStealthMode

`func (o *MacOsCompliancePolicy) SetFirewallEnableStealthMode(v bool)`

SetFirewallEnableStealthMode gets a reference to the given bool and assigns it to the FirewallEnableStealthMode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


