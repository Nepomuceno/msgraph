# Windows10MobileCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetPasswordRequired

`func (o *Windows10MobileCompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *Windows10MobileCompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *Windows10MobileCompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordBlockSimple

`func (o *Windows10MobileCompliancePolicy) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *Windows10MobileCompliancePolicy) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *Windows10MobileCompliancePolicy) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordMinimumLength

`func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *Windows10MobileCompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *Windows10MobileCompliancePolicy) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *Windows10MobileCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *Windows10MobileCompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *Windows10MobileCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordPreviousPasswordBlockCount

`func (o *Windows10MobileCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *Windows10MobileCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *Windows10MobileCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordExpirationDays

`func (o *Windows10MobileCompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *Windows10MobileCompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *Windows10MobileCompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeLock

`func (o *Windows10MobileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *Windows10MobileCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *Windows10MobileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordRequireToUnlockFromIdle

`func (o *Windows10MobileCompliancePolicy) GetPasswordRequireToUnlockFromIdle() bool`

GetPasswordRequireToUnlockFromIdle returns the PasswordRequireToUnlockFromIdle field if non-nil, zero value otherwise.

### GetPasswordRequireToUnlockFromIdleOk

`func (o *Windows10MobileCompliancePolicy) GetPasswordRequireToUnlockFromIdleOk() (bool, bool)`

GetPasswordRequireToUnlockFromIdleOk returns a tuple with the PasswordRequireToUnlockFromIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequireToUnlockFromIdle

`func (o *Windows10MobileCompliancePolicy) HasPasswordRequireToUnlockFromIdle() bool`

HasPasswordRequireToUnlockFromIdle returns a boolean if a field has been set.

### SetPasswordRequireToUnlockFromIdle

`func (o *Windows10MobileCompliancePolicy) SetPasswordRequireToUnlockFromIdle(v bool)`

SetPasswordRequireToUnlockFromIdle gets a reference to the given bool and assigns it to the PasswordRequireToUnlockFromIdle field.

### GetOsMinimumVersion

`func (o *Windows10MobileCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *Windows10MobileCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *Windows10MobileCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *Windows10MobileCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *Windows10MobileCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *Windows10MobileCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *Windows10MobileCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *Windows10MobileCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *Windows10MobileCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetEarlyLaunchAntiMalwareDriverEnabled

`func (o *Windows10MobileCompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabled() bool`

GetEarlyLaunchAntiMalwareDriverEnabled returns the EarlyLaunchAntiMalwareDriverEnabled field if non-nil, zero value otherwise.

### GetEarlyLaunchAntiMalwareDriverEnabledOk

`func (o *Windows10MobileCompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabledOk() (bool, bool)`

GetEarlyLaunchAntiMalwareDriverEnabledOk returns a tuple with the EarlyLaunchAntiMalwareDriverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEarlyLaunchAntiMalwareDriverEnabled

`func (o *Windows10MobileCompliancePolicy) HasEarlyLaunchAntiMalwareDriverEnabled() bool`

HasEarlyLaunchAntiMalwareDriverEnabled returns a boolean if a field has been set.

### SetEarlyLaunchAntiMalwareDriverEnabled

`func (o *Windows10MobileCompliancePolicy) SetEarlyLaunchAntiMalwareDriverEnabled(v bool)`

SetEarlyLaunchAntiMalwareDriverEnabled gets a reference to the given bool and assigns it to the EarlyLaunchAntiMalwareDriverEnabled field.

### GetBitLockerEnabled

`func (o *Windows10MobileCompliancePolicy) GetBitLockerEnabled() bool`

GetBitLockerEnabled returns the BitLockerEnabled field if non-nil, zero value otherwise.

### GetBitLockerEnabledOk

`func (o *Windows10MobileCompliancePolicy) GetBitLockerEnabledOk() (bool, bool)`

GetBitLockerEnabledOk returns a tuple with the BitLockerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerEnabled

`func (o *Windows10MobileCompliancePolicy) HasBitLockerEnabled() bool`

HasBitLockerEnabled returns a boolean if a field has been set.

### SetBitLockerEnabled

`func (o *Windows10MobileCompliancePolicy) SetBitLockerEnabled(v bool)`

SetBitLockerEnabled gets a reference to the given bool and assigns it to the BitLockerEnabled field.

### GetSecureBootEnabled

`func (o *Windows10MobileCompliancePolicy) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *Windows10MobileCompliancePolicy) GetSecureBootEnabledOk() (bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecureBootEnabled

`func (o *Windows10MobileCompliancePolicy) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### SetSecureBootEnabled

`func (o *Windows10MobileCompliancePolicy) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled gets a reference to the given bool and assigns it to the SecureBootEnabled field.

### GetCodeIntegrityEnabled

`func (o *Windows10MobileCompliancePolicy) GetCodeIntegrityEnabled() bool`

GetCodeIntegrityEnabled returns the CodeIntegrityEnabled field if non-nil, zero value otherwise.

### GetCodeIntegrityEnabledOk

`func (o *Windows10MobileCompliancePolicy) GetCodeIntegrityEnabledOk() (bool, bool)`

GetCodeIntegrityEnabledOk returns a tuple with the CodeIntegrityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCodeIntegrityEnabled

`func (o *Windows10MobileCompliancePolicy) HasCodeIntegrityEnabled() bool`

HasCodeIntegrityEnabled returns a boolean if a field has been set.

### SetCodeIntegrityEnabled

`func (o *Windows10MobileCompliancePolicy) SetCodeIntegrityEnabled(v bool)`

SetCodeIntegrityEnabled gets a reference to the given bool and assigns it to the CodeIntegrityEnabled field.

### GetStorageRequireEncryption

`func (o *Windows10MobileCompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *Windows10MobileCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *Windows10MobileCompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *Windows10MobileCompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


