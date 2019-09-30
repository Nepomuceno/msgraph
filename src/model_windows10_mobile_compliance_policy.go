/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
import (
	"encoding/json"
)
// Windows10MobileCompliancePolicy This class contains compliance settings for Windows 10 Mobile.
type Windows10MobileCompliancePolicy struct {
	// Require a password to unlock Windows Phone device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Whether or not to block syncing the calendar.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// Minimum password length. Valid values 4 to 16
	PasswordMinimumLength *int32 `json:"passwordMinimumLength,omitempty"`
	isExplicitNullPasswordMinimumLength bool `json:"-"`
	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int32 `json:"passwordMinimumCharacterSetCount,omitempty"`
	isExplicitNullPasswordMinimumCharacterSetCount bool `json:"-"`
	// The required password type.
	PasswordRequiredType *AnyOfmicrosoftGraphRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// The number of previous passwords to prevent re-use of.
	PasswordPreviousPasswordBlockCount *int32 `json:"passwordPreviousPasswordBlockCount,omitempty"`
	isExplicitNullPasswordPreviousPasswordBlockCount bool `json:"-"`
	// Number of days before password expiration. Valid values 1 to 255
	PasswordExpirationDays *int32 `json:"passwordExpirationDays,omitempty"`
	isExplicitNullPasswordExpirationDays bool `json:"-"`
	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int32 `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	isExplicitNullPasswordMinutesOfInactivityBeforeLock bool `json:"-"`
	// Require a password to unlock an idle device.
	PasswordRequireToUnlockFromIdle *bool `json:"passwordRequireToUnlockFromIdle,omitempty"`

	// Minimum Windows Phone version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	isExplicitNullOsMinimumVersion bool `json:"-"`
	// Maximum Windows Phone version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	isExplicitNullOsMaximumVersion bool `json:"-"`
	// Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled.
	EarlyLaunchAntiMalwareDriverEnabled *bool `json:"earlyLaunchAntiMalwareDriverEnabled,omitempty"`

	// Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
	BitLockerEnabled *bool `json:"bitLockerEnabled,omitempty"`

	// Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`

	// Require devices to be reported as healthy by Windows Device Health Attestation.
	CodeIntegrityEnabled *bool `json:"codeIntegrityEnabled,omitempty"`

	// Require encryption on windows devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`

}

// GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordRequired() bool {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret
	}
	return *o.PasswordRequired
}

// GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordRequiredOk() (bool, bool) {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordRequired, true
}

// HasPasswordRequired returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordRequired() bool {
	if o != nil && o.PasswordRequired != nil {
		return true
	}

	return false
}

// SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.
func (o *Windows10MobileCompliancePolicy) SetPasswordRequired(v bool) {
	o.PasswordRequired = &v
}

// GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordBlockSimple() bool {
	if o == nil || o.PasswordBlockSimple == nil {
		var ret bool
		return ret
	}
	return *o.PasswordBlockSimple
}

// GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool) {
	if o == nil || o.PasswordBlockSimple == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordBlockSimple, true
}

// HasPasswordBlockSimple returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordBlockSimple() bool {
	if o != nil && o.PasswordBlockSimple != nil {
		return true
	}

	return false
}

// SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.
func (o *Windows10MobileCompliancePolicy) SetPasswordBlockSimple(v bool) {
	o.PasswordBlockSimple = &v
}

// GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumLength() int32 {
	if o == nil || o.PasswordMinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinimumLength
}

// GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool) {
	if o == nil || o.PasswordMinimumLength == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinimumLength, true
}

// HasPasswordMinimumLength returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordMinimumLength() bool {
	if o != nil && o.PasswordMinimumLength != nil {
		return true
	}

	return false
}

// SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.
func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumLength(v int32) {
	o.PasswordMinimumLength = &v
}

// SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinimumLength value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool) {
	o.PasswordMinimumLength = nil
	o.isExplicitNullPasswordMinimumLength = b
}
// GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumCharacterSetCount() int32 {
	if o == nil || o.PasswordMinimumCharacterSetCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinimumCharacterSetCount
}

// GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool) {
	if o == nil || o.PasswordMinimumCharacterSetCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinimumCharacterSetCount, true
}

// HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordMinimumCharacterSetCount() bool {
	if o != nil && o.PasswordMinimumCharacterSetCount != nil {
		return true
	}

	return false
}

// SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.
func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32) {
	o.PasswordMinimumCharacterSetCount = &v
}

// SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool) {
	o.PasswordMinimumCharacterSetCount = nil
	o.isExplicitNullPasswordMinimumCharacterSetCount = b
}
// GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType {
	if o == nil || o.PasswordRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret
	}
	return *o.PasswordRequiredType
}

// GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool) {
	if o == nil || o.PasswordRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret, false
	}
	return *o.PasswordRequiredType, true
}

// HasPasswordRequiredType returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordRequiredType() bool {
	if o != nil && o.PasswordRequiredType != nil {
		return true
	}

	return false
}

// SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.
func (o *Windows10MobileCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType) {
	o.PasswordRequiredType = &v
}

// GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32 {
	if o == nil || o.PasswordPreviousPasswordBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordPreviousPasswordBlockCount
}

// GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool) {
	if o == nil || o.PasswordPreviousPasswordBlockCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordPreviousPasswordBlockCount, true
}

// HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool {
	if o != nil && o.PasswordPreviousPasswordBlockCount != nil {
		return true
	}

	return false
}

// SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.
func (o *Windows10MobileCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32) {
	o.PasswordPreviousPasswordBlockCount = &v
}

// SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool) {
	o.PasswordPreviousPasswordBlockCount = nil
	o.isExplicitNullPasswordPreviousPasswordBlockCount = b
}
// GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordExpirationDays() int32 {
	if o == nil || o.PasswordExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.PasswordExpirationDays
}

// GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool) {
	if o == nil || o.PasswordExpirationDays == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordExpirationDays, true
}

// HasPasswordExpirationDays returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordExpirationDays() bool {
	if o != nil && o.PasswordExpirationDays != nil {
		return true
	}

	return false
}

// SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.
func (o *Windows10MobileCompliancePolicy) SetPasswordExpirationDays(v int32) {
	o.PasswordExpirationDays = &v
}

// SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordExpirationDays value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool) {
	o.PasswordExpirationDays = nil
	o.isExplicitNullPasswordExpirationDays = b
}
// GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32 {
	if o == nil || o.PasswordMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinutesOfInactivityBeforeLock
}

// GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool) {
	if o == nil || o.PasswordMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinutesOfInactivityBeforeLock, true
}

// HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool {
	if o != nil && o.PasswordMinutesOfInactivityBeforeLock != nil {
		return true
	}

	return false
}

// SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.
func (o *Windows10MobileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32) {
	o.PasswordMinutesOfInactivityBeforeLock = &v
}

// SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool) {
	o.PasswordMinutesOfInactivityBeforeLock = nil
	o.isExplicitNullPasswordMinutesOfInactivityBeforeLock = b
}
// GetPasswordRequireToUnlockFromIdle returns the PasswordRequireToUnlockFromIdle field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetPasswordRequireToUnlockFromIdle() bool {
	if o == nil || o.PasswordRequireToUnlockFromIdle == nil {
		var ret bool
		return ret
	}
	return *o.PasswordRequireToUnlockFromIdle
}

// GetPasswordRequireToUnlockFromIdleOk returns a tuple with the PasswordRequireToUnlockFromIdle field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetPasswordRequireToUnlockFromIdleOk() (bool, bool) {
	if o == nil || o.PasswordRequireToUnlockFromIdle == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordRequireToUnlockFromIdle, true
}

// HasPasswordRequireToUnlockFromIdle returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasPasswordRequireToUnlockFromIdle() bool {
	if o != nil && o.PasswordRequireToUnlockFromIdle != nil {
		return true
	}

	return false
}

// SetPasswordRequireToUnlockFromIdle gets a reference to the given bool and assigns it to the PasswordRequireToUnlockFromIdle field.
func (o *Windows10MobileCompliancePolicy) SetPasswordRequireToUnlockFromIdle(v bool) {
	o.PasswordRequireToUnlockFromIdle = &v
}

// GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetOsMinimumVersion() string {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMinimumVersion
}

// GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetOsMinimumVersionOk() (string, bool) {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMinimumVersion, true
}

// HasOsMinimumVersion returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasOsMinimumVersion() bool {
	if o != nil && o.OsMinimumVersion != nil {
		return true
	}

	return false
}

// SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.
func (o *Windows10MobileCompliancePolicy) SetOsMinimumVersion(v string) {
	o.OsMinimumVersion = &v
}

// SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMinimumVersion value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool) {
	o.OsMinimumVersion = nil
	o.isExplicitNullOsMinimumVersion = b
}
// GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetOsMaximumVersion() string {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMaximumVersion
}

// GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetOsMaximumVersionOk() (string, bool) {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMaximumVersion, true
}

// HasOsMaximumVersion returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasOsMaximumVersion() bool {
	if o != nil && o.OsMaximumVersion != nil {
		return true
	}

	return false
}

// SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.
func (o *Windows10MobileCompliancePolicy) SetOsMaximumVersion(v string) {
	o.OsMaximumVersion = &v
}

// SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMaximumVersion value is set to nil even if false is passed
func (o *Windows10MobileCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool) {
	o.OsMaximumVersion = nil
	o.isExplicitNullOsMaximumVersion = b
}
// GetEarlyLaunchAntiMalwareDriverEnabled returns the EarlyLaunchAntiMalwareDriverEnabled field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabled() bool {
	if o == nil || o.EarlyLaunchAntiMalwareDriverEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EarlyLaunchAntiMalwareDriverEnabled
}

// GetEarlyLaunchAntiMalwareDriverEnabledOk returns a tuple with the EarlyLaunchAntiMalwareDriverEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabledOk() (bool, bool) {
	if o == nil || o.EarlyLaunchAntiMalwareDriverEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.EarlyLaunchAntiMalwareDriverEnabled, true
}

// HasEarlyLaunchAntiMalwareDriverEnabled returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasEarlyLaunchAntiMalwareDriverEnabled() bool {
	if o != nil && o.EarlyLaunchAntiMalwareDriverEnabled != nil {
		return true
	}

	return false
}

// SetEarlyLaunchAntiMalwareDriverEnabled gets a reference to the given bool and assigns it to the EarlyLaunchAntiMalwareDriverEnabled field.
func (o *Windows10MobileCompliancePolicy) SetEarlyLaunchAntiMalwareDriverEnabled(v bool) {
	o.EarlyLaunchAntiMalwareDriverEnabled = &v
}

// GetBitLockerEnabled returns the BitLockerEnabled field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetBitLockerEnabled() bool {
	if o == nil || o.BitLockerEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BitLockerEnabled
}

// GetBitLockerEnabledOk returns a tuple with the BitLockerEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetBitLockerEnabledOk() (bool, bool) {
	if o == nil || o.BitLockerEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.BitLockerEnabled, true
}

// HasBitLockerEnabled returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasBitLockerEnabled() bool {
	if o != nil && o.BitLockerEnabled != nil {
		return true
	}

	return false
}

// SetBitLockerEnabled gets a reference to the given bool and assigns it to the BitLockerEnabled field.
func (o *Windows10MobileCompliancePolicy) SetBitLockerEnabled(v bool) {
	o.BitLockerEnabled = &v
}

// GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetSecureBootEnabled() bool {
	if o == nil || o.SecureBootEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SecureBootEnabled
}

// GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetSecureBootEnabledOk() (bool, bool) {
	if o == nil || o.SecureBootEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.SecureBootEnabled, true
}

// HasSecureBootEnabled returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasSecureBootEnabled() bool {
	if o != nil && o.SecureBootEnabled != nil {
		return true
	}

	return false
}

// SetSecureBootEnabled gets a reference to the given bool and assigns it to the SecureBootEnabled field.
func (o *Windows10MobileCompliancePolicy) SetSecureBootEnabled(v bool) {
	o.SecureBootEnabled = &v
}

// GetCodeIntegrityEnabled returns the CodeIntegrityEnabled field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetCodeIntegrityEnabled() bool {
	if o == nil || o.CodeIntegrityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CodeIntegrityEnabled
}

// GetCodeIntegrityEnabledOk returns a tuple with the CodeIntegrityEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetCodeIntegrityEnabledOk() (bool, bool) {
	if o == nil || o.CodeIntegrityEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.CodeIntegrityEnabled, true
}

// HasCodeIntegrityEnabled returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasCodeIntegrityEnabled() bool {
	if o != nil && o.CodeIntegrityEnabled != nil {
		return true
	}

	return false
}

// SetCodeIntegrityEnabled gets a reference to the given bool and assigns it to the CodeIntegrityEnabled field.
func (o *Windows10MobileCompliancePolicy) SetCodeIntegrityEnabled(v bool) {
	o.CodeIntegrityEnabled = &v
}

// GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.
func (o *Windows10MobileCompliancePolicy) GetStorageRequireEncryption() bool {
	if o == nil || o.StorageRequireEncryption == nil {
		var ret bool
		return ret
	}
	return *o.StorageRequireEncryption
}

// GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Windows10MobileCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool) {
	if o == nil || o.StorageRequireEncryption == nil {
		var ret bool
		return ret, false
	}
	return *o.StorageRequireEncryption, true
}

// HasStorageRequireEncryption returns a boolean if a field has been set.
func (o *Windows10MobileCompliancePolicy) HasStorageRequireEncryption() bool {
	if o != nil && o.StorageRequireEncryption != nil {
		return true
	}

	return false
}

// SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.
func (o *Windows10MobileCompliancePolicy) SetStorageRequireEncryption(v bool) {
	o.StorageRequireEncryption = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o Windows10MobileCompliancePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordRequired != nil {
		toSerialize["passwordRequired"] = o.PasswordRequired
	}
	if o.PasswordBlockSimple != nil {
		toSerialize["passwordBlockSimple"] = o.PasswordBlockSimple
	}
	if o.PasswordMinimumLength == nil {
		if o.isExplicitNullPasswordMinimumLength {
			toSerialize["passwordMinimumLength"] = o.PasswordMinimumLength
		}
	} else {
		toSerialize["passwordMinimumLength"] = o.PasswordMinimumLength
	}
	if o.PasswordMinimumCharacterSetCount == nil {
		if o.isExplicitNullPasswordMinimumCharacterSetCount {
			toSerialize["passwordMinimumCharacterSetCount"] = o.PasswordMinimumCharacterSetCount
		}
	} else {
		toSerialize["passwordMinimumCharacterSetCount"] = o.PasswordMinimumCharacterSetCount
	}
	if o.PasswordRequiredType != nil {
		toSerialize["passwordRequiredType"] = o.PasswordRequiredType
	}
	if o.PasswordPreviousPasswordBlockCount == nil {
		if o.isExplicitNullPasswordPreviousPasswordBlockCount {
			toSerialize["passwordPreviousPasswordBlockCount"] = o.PasswordPreviousPasswordBlockCount
		}
	} else {
		toSerialize["passwordPreviousPasswordBlockCount"] = o.PasswordPreviousPasswordBlockCount
	}
	if o.PasswordExpirationDays == nil {
		if o.isExplicitNullPasswordExpirationDays {
			toSerialize["passwordExpirationDays"] = o.PasswordExpirationDays
		}
	} else {
		toSerialize["passwordExpirationDays"] = o.PasswordExpirationDays
	}
	if o.PasswordMinutesOfInactivityBeforeLock == nil {
		if o.isExplicitNullPasswordMinutesOfInactivityBeforeLock {
			toSerialize["passwordMinutesOfInactivityBeforeLock"] = o.PasswordMinutesOfInactivityBeforeLock
		}
	} else {
		toSerialize["passwordMinutesOfInactivityBeforeLock"] = o.PasswordMinutesOfInactivityBeforeLock
	}
	if o.PasswordRequireToUnlockFromIdle != nil {
		toSerialize["passwordRequireToUnlockFromIdle"] = o.PasswordRequireToUnlockFromIdle
	}
	if o.OsMinimumVersion == nil {
		if o.isExplicitNullOsMinimumVersion {
			toSerialize["osMinimumVersion"] = o.OsMinimumVersion
		}
	} else {
		toSerialize["osMinimumVersion"] = o.OsMinimumVersion
	}
	if o.OsMaximumVersion == nil {
		if o.isExplicitNullOsMaximumVersion {
			toSerialize["osMaximumVersion"] = o.OsMaximumVersion
		}
	} else {
		toSerialize["osMaximumVersion"] = o.OsMaximumVersion
	}
	if o.EarlyLaunchAntiMalwareDriverEnabled != nil {
		toSerialize["earlyLaunchAntiMalwareDriverEnabled"] = o.EarlyLaunchAntiMalwareDriverEnabled
	}
	if o.BitLockerEnabled != nil {
		toSerialize["bitLockerEnabled"] = o.BitLockerEnabled
	}
	if o.SecureBootEnabled != nil {
		toSerialize["secureBootEnabled"] = o.SecureBootEnabled
	}
	if o.CodeIntegrityEnabled != nil {
		toSerialize["codeIntegrityEnabled"] = o.CodeIntegrityEnabled
	}
	if o.StorageRequireEncryption != nil {
		toSerialize["storageRequireEncryption"] = o.StorageRequireEncryption
	}
	return json.Marshal(toSerialize)
}


