# IosCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetPasscodeBlockSimple

`func (o *IosCompliancePolicy) GetPasscodeBlockSimple() bool`

GetPasscodeBlockSimple returns the PasscodeBlockSimple field if non-nil, zero value otherwise.

### GetPasscodeBlockSimpleOk

`func (o *IosCompliancePolicy) GetPasscodeBlockSimpleOk() (bool, bool)`

GetPasscodeBlockSimpleOk returns a tuple with the PasscodeBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeBlockSimple

`func (o *IosCompliancePolicy) HasPasscodeBlockSimple() bool`

HasPasscodeBlockSimple returns a boolean if a field has been set.

### SetPasscodeBlockSimple

`func (o *IosCompliancePolicy) SetPasscodeBlockSimple(v bool)`

SetPasscodeBlockSimple gets a reference to the given bool and assigns it to the PasscodeBlockSimple field.

### GetPasscodeExpirationDays

`func (o *IosCompliancePolicy) GetPasscodeExpirationDays() int32`

GetPasscodeExpirationDays returns the PasscodeExpirationDays field if non-nil, zero value otherwise.

### GetPasscodeExpirationDaysOk

`func (o *IosCompliancePolicy) GetPasscodeExpirationDaysOk() (int32, bool)`

GetPasscodeExpirationDaysOk returns a tuple with the PasscodeExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeExpirationDays

`func (o *IosCompliancePolicy) HasPasscodeExpirationDays() bool`

HasPasscodeExpirationDays returns a boolean if a field has been set.

### SetPasscodeExpirationDays

`func (o *IosCompliancePolicy) SetPasscodeExpirationDays(v int32)`

SetPasscodeExpirationDays gets a reference to the given int32 and assigns it to the PasscodeExpirationDays field.

### SetPasscodeExpirationDaysExplicitNull

`func (o *IosCompliancePolicy) SetPasscodeExpirationDaysExplicitNull(b bool)`

SetPasscodeExpirationDaysExplicitNull (un)sets PasscodeExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeExpirationDays value is set to nil even if false is passed
### GetPasscodeMinimumLength

`func (o *IosCompliancePolicy) GetPasscodeMinimumLength() int32`

GetPasscodeMinimumLength returns the PasscodeMinimumLength field if non-nil, zero value otherwise.

### GetPasscodeMinimumLengthOk

`func (o *IosCompliancePolicy) GetPasscodeMinimumLengthOk() (int32, bool)`

GetPasscodeMinimumLengthOk returns a tuple with the PasscodeMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumLength

`func (o *IosCompliancePolicy) HasPasscodeMinimumLength() bool`

HasPasscodeMinimumLength returns a boolean if a field has been set.

### SetPasscodeMinimumLength

`func (o *IosCompliancePolicy) SetPasscodeMinimumLength(v int32)`

SetPasscodeMinimumLength gets a reference to the given int32 and assigns it to the PasscodeMinimumLength field.

### SetPasscodeMinimumLengthExplicitNull

`func (o *IosCompliancePolicy) SetPasscodeMinimumLengthExplicitNull(b bool)`

SetPasscodeMinimumLengthExplicitNull (un)sets PasscodeMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumLength value is set to nil even if false is passed
### GetPasscodeMinutesOfInactivityBeforeLock

`func (o *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLock() int32`

GetPasscodeMinutesOfInactivityBeforeLock returns the PasscodeMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasscodeMinutesOfInactivityBeforeLockOk

`func (o *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasscodeMinutesOfInactivityBeforeLockOk returns a tuple with the PasscodeMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinutesOfInactivityBeforeLock

`func (o *IosCompliancePolicy) HasPasscodeMinutesOfInactivityBeforeLock() bool`

HasPasscodeMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasscodeMinutesOfInactivityBeforeLock

`func (o *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLock(v int32)`

SetPasscodeMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeLock field.

### SetPasscodeMinutesOfInactivityBeforeLockExplicitNull

`func (o *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasscodeMinutesOfInactivityBeforeLockExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasscodePreviousPasscodeBlockCount

`func (o *IosCompliancePolicy) GetPasscodePreviousPasscodeBlockCount() int32`

GetPasscodePreviousPasscodeBlockCount returns the PasscodePreviousPasscodeBlockCount field if non-nil, zero value otherwise.

### GetPasscodePreviousPasscodeBlockCountOk

`func (o *IosCompliancePolicy) GetPasscodePreviousPasscodeBlockCountOk() (int32, bool)`

GetPasscodePreviousPasscodeBlockCountOk returns a tuple with the PasscodePreviousPasscodeBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodePreviousPasscodeBlockCount

`func (o *IosCompliancePolicy) HasPasscodePreviousPasscodeBlockCount() bool`

HasPasscodePreviousPasscodeBlockCount returns a boolean if a field has been set.

### SetPasscodePreviousPasscodeBlockCount

`func (o *IosCompliancePolicy) SetPasscodePreviousPasscodeBlockCount(v int32)`

SetPasscodePreviousPasscodeBlockCount gets a reference to the given int32 and assigns it to the PasscodePreviousPasscodeBlockCount field.

### SetPasscodePreviousPasscodeBlockCountExplicitNull

`func (o *IosCompliancePolicy) SetPasscodePreviousPasscodeBlockCountExplicitNull(b bool)`

SetPasscodePreviousPasscodeBlockCountExplicitNull (un)sets PasscodePreviousPasscodeBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodePreviousPasscodeBlockCount value is set to nil even if false is passed
### GetPasscodeMinimumCharacterSetCount

`func (o *IosCompliancePolicy) GetPasscodeMinimumCharacterSetCount() int32`

GetPasscodeMinimumCharacterSetCount returns the PasscodeMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasscodeMinimumCharacterSetCountOk

`func (o *IosCompliancePolicy) GetPasscodeMinimumCharacterSetCountOk() (int32, bool)`

GetPasscodeMinimumCharacterSetCountOk returns a tuple with the PasscodeMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeMinimumCharacterSetCount

`func (o *IosCompliancePolicy) HasPasscodeMinimumCharacterSetCount() bool`

HasPasscodeMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasscodeMinimumCharacterSetCount

`func (o *IosCompliancePolicy) SetPasscodeMinimumCharacterSetCount(v int32)`

SetPasscodeMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasscodeMinimumCharacterSetCount field.

### SetPasscodeMinimumCharacterSetCountExplicitNull

`func (o *IosCompliancePolicy) SetPasscodeMinimumCharacterSetCountExplicitNull(b bool)`

SetPasscodeMinimumCharacterSetCountExplicitNull (un)sets PasscodeMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasscodeMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasscodeRequiredType

`func (o *IosCompliancePolicy) GetPasscodeRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasscodeRequiredType returns the PasscodeRequiredType field if non-nil, zero value otherwise.

### GetPasscodeRequiredTypeOk

`func (o *IosCompliancePolicy) GetPasscodeRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasscodeRequiredTypeOk returns a tuple with the PasscodeRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequiredType

`func (o *IosCompliancePolicy) HasPasscodeRequiredType() bool`

HasPasscodeRequiredType returns a boolean if a field has been set.

### SetPasscodeRequiredType

`func (o *IosCompliancePolicy) SetPasscodeRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasscodeRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasscodeRequiredType field.

### GetPasscodeRequired

`func (o *IosCompliancePolicy) GetPasscodeRequired() bool`

GetPasscodeRequired returns the PasscodeRequired field if non-nil, zero value otherwise.

### GetPasscodeRequiredOk

`func (o *IosCompliancePolicy) GetPasscodeRequiredOk() (bool, bool)`

GetPasscodeRequiredOk returns a tuple with the PasscodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscodeRequired

`func (o *IosCompliancePolicy) HasPasscodeRequired() bool`

HasPasscodeRequired returns a boolean if a field has been set.

### SetPasscodeRequired

`func (o *IosCompliancePolicy) SetPasscodeRequired(v bool)`

SetPasscodeRequired gets a reference to the given bool and assigns it to the PasscodeRequired field.

### GetOsMinimumVersion

`func (o *IosCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *IosCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *IosCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *IosCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *IosCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *IosCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *IosCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *IosCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *IosCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *IosCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetSecurityBlockJailbrokenDevices

`func (o *IosCompliancePolicy) GetSecurityBlockJailbrokenDevices() bool`

GetSecurityBlockJailbrokenDevices returns the SecurityBlockJailbrokenDevices field if non-nil, zero value otherwise.

### GetSecurityBlockJailbrokenDevicesOk

`func (o *IosCompliancePolicy) GetSecurityBlockJailbrokenDevicesOk() (bool, bool)`

GetSecurityBlockJailbrokenDevicesOk returns a tuple with the SecurityBlockJailbrokenDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityBlockJailbrokenDevices

`func (o *IosCompliancePolicy) HasSecurityBlockJailbrokenDevices() bool`

HasSecurityBlockJailbrokenDevices returns a boolean if a field has been set.

### SetSecurityBlockJailbrokenDevices

`func (o *IosCompliancePolicy) SetSecurityBlockJailbrokenDevices(v bool)`

SetSecurityBlockJailbrokenDevices gets a reference to the given bool and assigns it to the SecurityBlockJailbrokenDevices field.

### GetDeviceThreatProtectionEnabled

`func (o *IosCompliancePolicy) GetDeviceThreatProtectionEnabled() bool`

GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionEnabledOk

`func (o *IosCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool)`

GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionEnabled

`func (o *IosCompliancePolicy) HasDeviceThreatProtectionEnabled() bool`

HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.

### SetDeviceThreatProtectionEnabled

`func (o *IosCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool)`

SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.

### GetDeviceThreatProtectionRequiredSecurityLevel

`func (o *IosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel`

GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionRequiredSecurityLevelOk

`func (o *IosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool)`

GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionRequiredSecurityLevel

`func (o *IosCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool`

HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.

### SetDeviceThreatProtectionRequiredSecurityLevel

`func (o *IosCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel)`

SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.

### GetManagedEmailProfileRequired

`func (o *IosCompliancePolicy) GetManagedEmailProfileRequired() bool`

GetManagedEmailProfileRequired returns the ManagedEmailProfileRequired field if non-nil, zero value otherwise.

### GetManagedEmailProfileRequiredOk

`func (o *IosCompliancePolicy) GetManagedEmailProfileRequiredOk() (bool, bool)`

GetManagedEmailProfileRequiredOk returns a tuple with the ManagedEmailProfileRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedEmailProfileRequired

`func (o *IosCompliancePolicy) HasManagedEmailProfileRequired() bool`

HasManagedEmailProfileRequired returns a boolean if a field has been set.

### SetManagedEmailProfileRequired

`func (o *IosCompliancePolicy) SetManagedEmailProfileRequired(v bool)`

SetManagedEmailProfileRequired gets a reference to the given bool and assigns it to the ManagedEmailProfileRequired field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


