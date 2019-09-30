# Windows10CompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordRequired** | Pointer to **bool** | Require a password to unlock Windows device. | [optional] 
**PasswordBlockSimple** | Pointer to **bool** | Indicates whether or not to block simple password. | [optional] 
**PasswordRequiredToUnlockFromIdle** | Pointer to **bool** | Require a password to unlock an idle device. | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a password is required. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | The password expiration in days. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | The minimum password length. | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | The number of character sets required in the password. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | The required password type. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | The number of previous passwords to prevent re-use of. | [optional] 
**RequireHealthyDeviceReport** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum Windows 10 version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum Windows 10 version. | [optional] 
**MobileOsMinimumVersion** | Pointer to **string** | Minimum Windows Phone version. | [optional] 
**MobileOsMaximumVersion** | Pointer to **string** | Maximum Windows Phone version. | [optional] 
**EarlyLaunchAntiMalwareDriverEnabled** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled. | [optional] 
**BitLockerEnabled** | Pointer to **bool** | Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled | [optional] 
**SecureBootEnabled** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled. | [optional] 
**CodeIntegrityEnabled** | Pointer to **bool** | Require devices to be reported as healthy by Windows Device Health Attestation. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Require encryption on windows devices. | [optional] 

## Methods

### GetPasswordRequired

`func (o *Windows10CompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *Windows10CompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *Windows10CompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *Windows10CompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordBlockSimple

`func (o *Windows10CompliancePolicy) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *Windows10CompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *Windows10CompliancePolicy) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *Windows10CompliancePolicy) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordRequiredToUnlockFromIdle

`func (o *Windows10CompliancePolicy) GetPasswordRequiredToUnlockFromIdle() bool`

GetPasswordRequiredToUnlockFromIdle returns the PasswordRequiredToUnlockFromIdle field if non-nil, zero value otherwise.

### GetPasswordRequiredToUnlockFromIdleOk

`func (o *Windows10CompliancePolicy) GetPasswordRequiredToUnlockFromIdleOk() (bool, bool)`

GetPasswordRequiredToUnlockFromIdleOk returns a tuple with the PasswordRequiredToUnlockFromIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredToUnlockFromIdle

`func (o *Windows10CompliancePolicy) HasPasswordRequiredToUnlockFromIdle() bool`

HasPasswordRequiredToUnlockFromIdle returns a boolean if a field has been set.

### SetPasswordRequiredToUnlockFromIdle

`func (o *Windows10CompliancePolicy) SetPasswordRequiredToUnlockFromIdle(v bool)`

SetPasswordRequiredToUnlockFromIdle gets a reference to the given bool and assigns it to the PasswordRequiredToUnlockFromIdle field.

### GetPasswordMinutesOfInactivityBeforeLock

`func (o *Windows10CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *Windows10CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *Windows10CompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *Windows10CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *Windows10CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordExpirationDays

`func (o *Windows10CompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *Windows10CompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *Windows10CompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *Windows10CompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *Windows10CompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *Windows10CompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *Windows10CompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *Windows10CompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *Windows10CompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *Windows10CompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *Windows10CompliancePolicy) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *Windows10CompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *Windows10CompliancePolicy) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *Windows10CompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *Windows10CompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *Windows10CompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *Windows10CompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *Windows10CompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *Windows10CompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordPreviousPasswordBlockCount

`func (o *Windows10CompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *Windows10CompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *Windows10CompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *Windows10CompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *Windows10CompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetRequireHealthyDeviceReport

`func (o *Windows10CompliancePolicy) GetRequireHealthyDeviceReport() bool`

GetRequireHealthyDeviceReport returns the RequireHealthyDeviceReport field if non-nil, zero value otherwise.

### GetRequireHealthyDeviceReportOk

`func (o *Windows10CompliancePolicy) GetRequireHealthyDeviceReportOk() (bool, bool)`

GetRequireHealthyDeviceReportOk returns a tuple with the RequireHealthyDeviceReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequireHealthyDeviceReport

`func (o *Windows10CompliancePolicy) HasRequireHealthyDeviceReport() bool`

HasRequireHealthyDeviceReport returns a boolean if a field has been set.

### SetRequireHealthyDeviceReport

`func (o *Windows10CompliancePolicy) SetRequireHealthyDeviceReport(v bool)`

SetRequireHealthyDeviceReport gets a reference to the given bool and assigns it to the RequireHealthyDeviceReport field.

### GetOsMinimumVersion

`func (o *Windows10CompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *Windows10CompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *Windows10CompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *Windows10CompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *Windows10CompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *Windows10CompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *Windows10CompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *Windows10CompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *Windows10CompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *Windows10CompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetMobileOsMinimumVersion

`func (o *Windows10CompliancePolicy) GetMobileOsMinimumVersion() string`

GetMobileOsMinimumVersion returns the MobileOsMinimumVersion field if non-nil, zero value otherwise.

### GetMobileOsMinimumVersionOk

`func (o *Windows10CompliancePolicy) GetMobileOsMinimumVersionOk() (string, bool)`

GetMobileOsMinimumVersionOk returns a tuple with the MobileOsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileOsMinimumVersion

`func (o *Windows10CompliancePolicy) HasMobileOsMinimumVersion() bool`

HasMobileOsMinimumVersion returns a boolean if a field has been set.

### SetMobileOsMinimumVersion

`func (o *Windows10CompliancePolicy) SetMobileOsMinimumVersion(v string)`

SetMobileOsMinimumVersion gets a reference to the given string and assigns it to the MobileOsMinimumVersion field.

### SetMobileOsMinimumVersionExplicitNull

`func (o *Windows10CompliancePolicy) SetMobileOsMinimumVersionExplicitNull(b bool)`

SetMobileOsMinimumVersionExplicitNull (un)sets MobileOsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobileOsMinimumVersion value is set to nil even if false is passed
### GetMobileOsMaximumVersion

`func (o *Windows10CompliancePolicy) GetMobileOsMaximumVersion() string`

GetMobileOsMaximumVersion returns the MobileOsMaximumVersion field if non-nil, zero value otherwise.

### GetMobileOsMaximumVersionOk

`func (o *Windows10CompliancePolicy) GetMobileOsMaximumVersionOk() (string, bool)`

GetMobileOsMaximumVersionOk returns a tuple with the MobileOsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileOsMaximumVersion

`func (o *Windows10CompliancePolicy) HasMobileOsMaximumVersion() bool`

HasMobileOsMaximumVersion returns a boolean if a field has been set.

### SetMobileOsMaximumVersion

`func (o *Windows10CompliancePolicy) SetMobileOsMaximumVersion(v string)`

SetMobileOsMaximumVersion gets a reference to the given string and assigns it to the MobileOsMaximumVersion field.

### SetMobileOsMaximumVersionExplicitNull

`func (o *Windows10CompliancePolicy) SetMobileOsMaximumVersionExplicitNull(b bool)`

SetMobileOsMaximumVersionExplicitNull (un)sets MobileOsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobileOsMaximumVersion value is set to nil even if false is passed
### GetEarlyLaunchAntiMalwareDriverEnabled

`func (o *Windows10CompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabled() bool`

GetEarlyLaunchAntiMalwareDriverEnabled returns the EarlyLaunchAntiMalwareDriverEnabled field if non-nil, zero value otherwise.

### GetEarlyLaunchAntiMalwareDriverEnabledOk

`func (o *Windows10CompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabledOk() (bool, bool)`

GetEarlyLaunchAntiMalwareDriverEnabledOk returns a tuple with the EarlyLaunchAntiMalwareDriverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEarlyLaunchAntiMalwareDriverEnabled

`func (o *Windows10CompliancePolicy) HasEarlyLaunchAntiMalwareDriverEnabled() bool`

HasEarlyLaunchAntiMalwareDriverEnabled returns a boolean if a field has been set.

### SetEarlyLaunchAntiMalwareDriverEnabled

`func (o *Windows10CompliancePolicy) SetEarlyLaunchAntiMalwareDriverEnabled(v bool)`

SetEarlyLaunchAntiMalwareDriverEnabled gets a reference to the given bool and assigns it to the EarlyLaunchAntiMalwareDriverEnabled field.

### GetBitLockerEnabled

`func (o *Windows10CompliancePolicy) GetBitLockerEnabled() bool`

GetBitLockerEnabled returns the BitLockerEnabled field if non-nil, zero value otherwise.

### GetBitLockerEnabledOk

`func (o *Windows10CompliancePolicy) GetBitLockerEnabledOk() (bool, bool)`

GetBitLockerEnabledOk returns a tuple with the BitLockerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerEnabled

`func (o *Windows10CompliancePolicy) HasBitLockerEnabled() bool`

HasBitLockerEnabled returns a boolean if a field has been set.

### SetBitLockerEnabled

`func (o *Windows10CompliancePolicy) SetBitLockerEnabled(v bool)`

SetBitLockerEnabled gets a reference to the given bool and assigns it to the BitLockerEnabled field.

### GetSecureBootEnabled

`func (o *Windows10CompliancePolicy) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *Windows10CompliancePolicy) GetSecureBootEnabledOk() (bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecureBootEnabled

`func (o *Windows10CompliancePolicy) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### SetSecureBootEnabled

`func (o *Windows10CompliancePolicy) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled gets a reference to the given bool and assigns it to the SecureBootEnabled field.

### GetCodeIntegrityEnabled

`func (o *Windows10CompliancePolicy) GetCodeIntegrityEnabled() bool`

GetCodeIntegrityEnabled returns the CodeIntegrityEnabled field if non-nil, zero value otherwise.

### GetCodeIntegrityEnabledOk

`func (o *Windows10CompliancePolicy) GetCodeIntegrityEnabledOk() (bool, bool)`

GetCodeIntegrityEnabledOk returns a tuple with the CodeIntegrityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCodeIntegrityEnabled

`func (o *Windows10CompliancePolicy) HasCodeIntegrityEnabled() bool`

HasCodeIntegrityEnabled returns a boolean if a field has been set.

### SetCodeIntegrityEnabled

`func (o *Windows10CompliancePolicy) SetCodeIntegrityEnabled(v bool)`

SetCodeIntegrityEnabled gets a reference to the given bool and assigns it to the CodeIntegrityEnabled field.

### GetStorageRequireEncryption

`func (o *Windows10CompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *Windows10CompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *Windows10CompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *Windows10CompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


